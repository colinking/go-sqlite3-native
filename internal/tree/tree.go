package tree

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/colinking/go-sqlite3-native/internal/pager"
	"github.com/segmentio/events/v2"
	"github.com/segmentio/textio"
)

// Tree is an implementation of a B-tree.
//
// This type of tree is composed of internal and leaf pages. Each tree page maps
// to exactly one database page.
//
// Leaf pages store multiple tuples of (key, data) where the data is an arbitrary
// []byte and the key is a 64-bit signed integer.
//
// Interior pages map to multiple children pages, which could be either other internal
// pages or leaf pages. This mapping is represented by a 32 bit database page index.
//
// For more information on the file format that backs this tree, see:
// https://www.sqlite.org/fileformat2.html#b_tree_pages
type Tree struct {
	pager          *pager.Pager
	rootPageNumber int
	root           *node

	// cursor location
	cursor      *node
	cursorStack []int

	// Error produced by Next() and fetched by Err()
	err error
}

func (t *Tree) String() string {
	var buf bytes.Buffer

	err := t.nodeString(&buf, t.rootPageNumber, t.root)
	if err != nil {
		t.setError(err)

		return ""
	}

	return buf.String()
}

func (t *Tree) nodeString(w io.Writer, pageNumber int, node *node) error {
	tw := textio.NewTreeWriter(w)
	defer tw.Close()

	_, err := tw.WriteString(fmt.Sprintf("Page %d (%s)", pageNumber, node.typ.String()))
	if err != nil {
		return err
	}

	for _, record := range node.records {
		var row []string
		for _, c := range record.columns {
			row = append(row, c.String())
		}
		_, err := io.WriteString(textio.NewTreeWriter(tw), fmt.Sprintf("rowid=%d row=[ %s ]", record.key, strings.Join(row, " | ")))
		if err != nil {
			return err
		}
	}

	for _, chld := range node.children {
		if chld.node == nil {
			// We lazy-load children pages until we need them:
			chld.node, err = newNode(chld.pageNumber, t.pager, node)
			if err != nil {
				return err
			}
		}
		err := t.nodeString(textio.NewTreeWriter(tw), chld.pageNumber, chld.node)
		if err != nil {
			return err
		}
	}

	return nil
}

func newTree(rootPageNumber int, pgr *pager.Pager) (*Tree, error) {
	root, err := newNode(rootPageNumber, pgr /* parent= */, nil)
	if err != nil {
		return nil, err
	}

	t := &Tree{
		pager:          pgr,
		rootPageNumber: rootPageNumber,
		root:           root,
	}
	t.ResetCursor()

	return t, nil
}

//go:generate stringer -type=TreeType
type TreeType int

const (
	TreeTypeUnknown TreeType = iota
	// PTF_LEAFDATA + PTF_INTKEY
	TreeTypeTableInterior
	// PTF_LEAF + PTF_LEAFDATA + PTF_INTKEY
	TreeTypeTableLeaf
	// PTF_ZERODATA
	TreeTypeIndexInterior
	// PTF_LEAF + PTF_ZERODATA
	TreeTypeIndexLeaf
)

func ToTreeType(b byte) TreeType {
	switch b {
	case 0x02:
		return TreeTypeIndexInterior
	case 0x05:
		return TreeTypeTableInterior
	case 0x0a:
		return TreeTypeIndexLeaf
	case 0x0d:
		return TreeTypeTableLeaf
	default:
		return TreeTypeUnknown
	}
}

func (t *Tree) ResetCursor() {
	t.cursor = t.root
	t.cursorStack = []int{-1}
}

func (t *Tree) Next() bool {
	for {
		// Move the cursor to the next record/child in this node:
		idx := t.cursorStackPeek() + 1

		switch t.cursor.typ {
		case TreeTypeTableLeaf:
			// If there is a record at this index, then we've found a next record:
			if idx < len(t.cursor.records) {
				// Store this new index so we can access it on the next Get() call:
				t.cursorStack[len(t.cursorStack)-1] = idx
				return true
			}
			// Otherwise, we've exahusted all records in this leaf node. This means we
			// should move to the next leaf node.
			if t.cursor.parent == nil {
				// This means there are no more nodes to look for a next record in.
				return false
			}
			t.cursorStackPop()
			t.cursor = t.cursor.parent
		case TreeTypeTableInterior:
			if idx < len(t.cursor.children) {
				// Move our cursor to this child where we will continue the search
				t.cursorStack[len(t.cursorStack)-1] = idx
				t.cursorStack = append(t.cursorStack, -1)
				if t.cursor.children[idx].node == nil {
					// We lazy-load children pages until we need them:
					node, err := newNode(t.cursor.children[idx].pageNumber, t.pager, t.cursor)
					if err != nil {
						t.setError(err)
						return false
					}
					t.cursor.children[idx].node = node
				}
				t.cursor = t.cursor.children[idx].node
			} else {
				// Otherwise, we should pop to this node's parent to continue the search there.
				if t.cursor.parent == nil {
					// This means there are no more nodes to look for a next record in.
					return false
				}
				t.cursorStackPop()
				t.cursor = t.cursor.parent
			}
		}
	}
}

// cursorStackPeek returns the last index in cursorStack
func (t *Tree) cursorStackPeek() int {
	return t.cursorStack[len(t.cursorStack)-1]
}

// cursorStackPop returns the last index in cursorStack
func (t *Tree) cursorStackPop() {
	t.cursorStack = t.cursorStack[:len(t.cursorStack)-1]
}

func (t *Tree) Get() Record {
	return t.cursor.records[t.cursorStackPeek()]
}

func (t *Tree) setError(err error) {
	if t.err != nil {
		events.Debug("second error is shadowing previous error: %+v", t.err)
	}

	t.err = err
}

func (t *Tree) Err() error {
	return t.err
}

func (t *Tree) Close() error {
	return t.root.Close()
}

package tree

import (
	"database/sql/driver"
	"encoding/binary"
	"errors"
	"fmt"
	"math"

	"github.com/colinking/go-sqlite3-native/internal/pager"
	"github.com/mohae/uvarint"
	"github.com/segmentio/events/v2"
)

type TreeManager struct {
	pager *pager.Pager
}

func NewManager(pager *pager.Pager) *TreeManager {
	return &TreeManager{
		pager: pager,
	}
}

func (tm *TreeManager) Header() (pager.SQLiteHeader, error) {
	return tm.pager.Header()
}

func (tm *TreeManager) Open(rootPage int) (*Tree, error) {
	return newTree(rootPage, tm.pager)
}

func (tm *TreeManager) Close() error {
	return tm.pager.Close()
}

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
	pager *pager.Pager
	root  *node

	// cursor location
	cursor      *node
	cursorStack []int
}

func newTree(rootPageNumber int, pgr *pager.Pager) (*Tree, error) {
	root, err := newNode(rootPageNumber, pgr /* parent= */, nil)
	if err != nil {
		return nil, err
	}

	events.Debug("root node (tree=%d): %+v", rootPageNumber, root)

	t := &Tree{
		pager: pgr,
		root:  root,
	}
	t.ResetCursor()

	return t, nil
}

//go:generate stringer -type=TreeType
type TreeType int

const (
	TreeTypeUnknown TreeType = iota
	TreeTypeTableInterior
	TreeTypeTableLeaf
	TreeTypeIndexInterior
	TreeTypeIndexLeaf
)

func ToTreeType(b byte) TreeType {
	switch b {
	case 0x02:
		return TreeTypeIndexInterior
	case 0x05:
		return TreeTypeTableInterior
	case 0x10:
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

func (t *Tree) Close() error {
	return t.root.Close()
}

type node struct {
	pager *pager.Pager

	// the following fields are kept for debugging purposes:
	typ                    TreeType
	freeblockOffset        int
	numCells               int
	contentOffset          int
	numFreeBytes           int
	nextPageNumber         int
	cellPointerArrayOffset int

	// the content of this node
	// these are ordered in increasing order
	records  []Record
	children []child

	// parent pointer for backtracking when traversing a tree
	// nil if there is no parent (the root page)
	parent *node
}

type child struct {
	key    int
	offset int
	node   *node
}

type Record struct {
	key     int
	columns []Column
}

func (r Record) GetColumn(idx int) Column {
	if idx < len(r.columns) {
		return r.columns[idx]
	}

	// NULL
	return Column{
		typ: 0,
	}
}

type Column struct {
	typ     int
	content []byte
}

func (c Column) Value() driver.Value {
	// TODO: validate this works with negative integers (2's complement)

	switch c.typ {
	case 0:
		return nil
	case 1:
		return int64(c.content[0])
	case 2:
		return int64(binary.BigEndian.Uint16(c.content))
	case 3:
		// stdlib binary does not have a 24-bit option
		b := c.content
		u := uint32(b[2]) | uint32(b[1])<<8 | uint32(b[0])<<16
		return int64(u)
	case 4:
		return int64(binary.BigEndian.Uint32(c.content))
	case 5:
		// stdlib binary does not have a 48-bit option
		b := c.content
		u := uint64(b[5]) | uint64(b[4])<<8 | uint64(b[3])<<16 | uint64(b[2])<<24 | uint64(b[1])<<32 | uint64(b[0])<<40
		return int64(u)
	case 6:
		return int64(binary.BigEndian.Uint64(c.content))
	case 7:
		b := binary.BigEndian.Uint64(c.content)
		return math.Float64frombits(b)
	case 8:
		return int64(0)
	case 9:
		return int64(1)
	default:
		if c.typ%2 == 0 {
			// blob
			return c.content
		} else {
			// string
			return string(c.content)
		}
	}
}

// https://www.sqlite.org/fileformat2.html#serialtype
func columnContentSize(typ int) int {
	switch typ {
	case 0, 8, 9:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return 4
	case 5:
		return 6
	case 6, 7:
		return 8
	default:
		if typ%2 == 0 {
			return (typ - 12) / 2
		} else {
			return (typ - 13) / 2
		}
	}
}

func newNode(pageNumber int, pgr *pager.Pager, parent *node) (n *node, err error) {
	page, err := pgr.Get(pageNumber)
	if err != nil {
		return nil, err
	}
	defer func() {
		if rerr := pgr.ReleasePage(); rerr != nil {
			// TODO: multi errors
			err = rerr
		}
	}()

	offset := 0
	if pageNumber == 1 { // the root table (aka sqlite_schema)
		// Note: page 1 is an exception, since it includes the database header which
		// is contained in the first 100 bytes of this page.
		offset += 100
	}

	// Read the tree header, which is stored in the either the first 8 bytes (leaf pages)
	// or 12 bytes (interior pages) of the page.
	typ := ToTreeType(page[offset])
	if typ == TreeTypeUnknown {
		return nil, fmt.Errorf("unknown tree page type for page=%d: %+v", pageNumber, page[offset])
	}
	// From this point forward, we'll assume we're working on a table tree.
	if typ == TreeTypeIndexInterior || typ == TreeTypeIndexLeaf {
		return nil, fmt.Errorf("indexes not supported yet")
	}
	offset += 1

	freeblockOffset := int(binary.BigEndian.Uint16(page[offset : offset+2]))
	offset += 2

	numCells := int(binary.BigEndian.Uint16(page[offset : offset+2]))
	offset += 2

	contentOffset := int(binary.BigEndian.Uint16(page[offset : offset+2]))
	offset += 2

	numFreeBytes := int(page[offset])
	offset += 1

	// If this is an interior page, then the header is 4 bytes longer because the next four bytes
	// store the pointer of the right-most node in this page.
	nextPageNumber := 0
	if typ == TreeTypeTableInterior {
		nextPageNumber = int(binary.BigEndian.Uint32(page[offset : offset+4]))
		offset += 4
	}

	// TODO: validate long keys (>25% of the page size) which will test overflow pages,
	// however this only happens on index trees.

	// Read all cells into memory.
	records := []Record{}
	children := []child{}
	for i := 0; i < numCells; i++ {
		// read the cell pointer
		cellOffset := int(binary.BigEndian.Uint16(page[offset : offset+2]))
		offset += 2

		switch typ {
		case TreeTypeTableInterior:
			// A 4-byte big-endian page number which is the left child pointer.
			childOffset := int(binary.BigEndian.Uint32(page[cellOffset : cellOffset+4]))
			cellOffset += 4
			// A varint which is the integer key
			key, _ := uvarint.Uvarint(page[cellOffset:])
			children = append(children, child{
				offset: childOffset,
				key:    int(key),
			})
		case TreeTypeTableLeaf:
			// A varint which is the total number of bytes of payload, including any overflow
			numPayloadBytes, size := uvarint.Uvarint(page[cellOffset:])
			cellOffset += size

			// A varint which is the integer key, a.k.a. "rowid"
			rowID, size := uvarint.Uvarint(page[cellOffset:])
			cellOffset += size

			// The initial portion of the payload that does not spill to overflow pages.
			end := cellOffset + int(numPayloadBytes)
			if end > len(page) {
				return nil, errors.New("unsupported: content overflowing")
			}
			content := page[cellOffset:end]

			// A 4-byte big-endian integer page number for the first page of the overflow page list - omitted if all payload fits on the b-tree page.
			// TODO: support overflowing content

			// read columns using the SQLite record format
			// https://www.sqlite.org/fileformat2.html#record_format
			contentOffset := 0
			headerSize, size := uvarint.Uvarint(content[contentOffset:])
			contentOffset += size
			columnTypes := []int{}
			for contentOffset < int(headerSize) {
				serialType, size := uvarint.Uvarint(content[contentOffset:])
				contentOffset += size
				columnTypes = append(columnTypes, int(serialType))
			}
			columns := make([]Column, 0, len(columnTypes))
			for _, typ := range columnTypes {
				size := columnContentSize(typ)
				columns = append(columns, Column{
					typ:     typ,
					content: content[contentOffset : contentOffset+size],
				})
				contentOffset += size

				// TODO: support ALTER COLUMN ADD COLUMN where we should use default values here
			}

			records = append(records, Record{
				key:     int(rowID),
				columns: columns,
			})
		default:
			return nil, fmt.Errorf("unsupported tree page type: %+v", typ)
		}
	}

	return &node{
		pager: pgr,

		records:  records,
		children: children,

		typ:                    typ,
		freeblockOffset:        freeblockOffset,
		numCells:               numCells,
		contentOffset:          contentOffset,
		numFreeBytes:           numFreeBytes,
		nextPageNumber:         nextPageNumber,
		cellPointerArrayOffset: offset,

		parent: parent,
	}, nil
}

func (n *node) Close() error {
	// TODO: release all children nodes
	// for _, child := range n.children {
	// 	child
	// }

	return nil
}

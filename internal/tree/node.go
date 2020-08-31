package tree

import (
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"

	"github.com/colinking/go-sqlite3-native/internal"
	"github.com/colinking/go-sqlite3-native/internal/pager"
)

type node struct {
	pager *pager.Pager

	// the following fields are kept for debugging purposes:
	typ                    TreeType
	freeblockOffset        int
	numCells               int
	contentOffset          int
	numFreeBytes           int
	cellPointerArrayOffset int

	// the content of this node
	// these are ordered in increasing order
	records  []Record
	children []*child

	// parent pointer for backtracking when traversing a tree
	// nil if there is no parent (the root page)
	parent *node
}

type child struct {
	key        []byte
	pageNumber int
	node       *node
}

type Record struct {
	key     []byte
	columns []Column
}

func (r Record) String() string {
	var row []string
	for _, c := range r.columns {
		row = append(row, c.String())
	}

	// TODO: parse key as varint
	return fmt.Sprintf("key=%+v row=[ %s ]", r.key, strings.Join(row, " | "))
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

func (c Column) String() string {
	v := c.Value()
	switch vt := v.(type) {
	case string:
		return vt
	case []byte:
		return "\"" + string(vt) + "\""
	case float64:
		return fmt.Sprintf("%f", vt)
	case int64:
		return fmt.Sprintf("%d", vt)
	case time.Time:
		return vt.String()
	case bool:
		return fmt.Sprintf("%v", vt)
	case nil:
		return "nil"
	default:
		return fmt.Sprintf("<unknown column type = %s>", reflect.TypeOf(v))
	}
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
	case 10, 11:
		// https://github.com/sqlite/sqlite/blob/96e3c39bd58ede59150c00e4f8609cbac674ffae/tool/offsets.c#L216
		// return 0
		panic(fmt.Errorf("cannot support columns of type=%d", typ))
	default:
		if typ%2 == 0 {
			return (typ - 12) / 2
		} else {
			// https://github.com/sqlite/sqlite/blob/96e3c39bd58ede59150c00e4f8609cbac674ffae/tool/offsets.c#L216
			// should this be 12?
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
	offset += 1

	// Assert that tree pages are of the same type (table vs. index) and that
	// parents are always of type interior.
	if parent != nil {
		if typ == TreeTypeTableInterior || typ == TreeTypeTableLeaf {
			if parent.typ != TreeTypeTableInterior {
				return nil, fmt.Errorf("invalid node type: parent=%s child=%s", parent.typ.String(), typ.String())
			}
		} else if typ == TreeTypeIndexInterior || typ == TreeTypeIndexLeaf {
			if parent.typ != TreeTypeIndexInterior {
				return nil, fmt.Errorf("invalid node type: parent=%s child=%s", parent.typ.String(), typ.String())
			}
		}
	}

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
	if typ == TreeTypeTableInterior || typ == TreeTypeIndexInterior {
		nextPageNumber = int(binary.BigEndian.Uint32(page[offset : offset+4]))
		offset += 4
	}

	// TODO: validate long keys (>25% of the page size) which will test overflow pages,
	// however this only happens on index trees.

	// Read all cell pointers into memory.
	ptrs := []int{}
	for i := 0; i < numCells; i++ {
		// read the cell pointer
		b := page[offset : offset+2]
		v := binary.BigEndian.Uint16(b)
		ptrs = append(ptrs, int(v))
		offset += 2
	}

	records := []Record{}
	children := []*child{}
	for _, ptr := range ptrs {
		// read the cell contents
		switch typ {
		case TreeTypeTableInterior, TreeTypeIndexInterior:
			// A 4-byte big-endian page number which is the left child pointer.
			childPageNumber := int(binary.BigEndian.Uint32(page[ptr : ptr+4]))
			ptr += 4

			var key []byte
			if typ == TreeTypeTableInterior {
				// A varint which is the integer key
				_, size := internal.Varint(page[ptr:])
				// We ignore the varint value because we want to store the key as []byte
				cpy := make([]byte, size)
				copy(cpy, page[ptr:ptr+size])
				key = cpy
			} else {
				var keySizeBytes int
				ptr += internal.PutVarint(page[ptr:], &keySizeBytes)
				// TODO: support overflowing index keys
				cpy := make([]byte, keySizeBytes)
				copy(cpy, page[ptr:ptr+keySizeBytes])
				key = cpy
			}

			children = append(children, &child{
				pageNumber: childPageNumber,
				key:        key,
			})
		case TreeTypeTableLeaf, TreeTypeIndexLeaf:
			// A varint which is the total number of bytes of payload, including any overflow
			var numPayloadBytes int
			ptr += internal.PutVarint(page[ptr:], &numPayloadBytes)

			var key []byte
			if typ == TreeTypeTableLeaf {
				// A varint which is the integer key, a.k.a. "rowid"
				_, size := internal.Varint(page[ptr:])
				cpy := make([]byte, size)
				copy(cpy, page[ptr:ptr+size])
				key = cpy
				ptr += size
			}

			// The initial portion of the payload that does not spill to overflow pages.
			// TODO: support overflowing using the U/M/P/X calculations:
			//     U = 4096
			//     P = ?
			//     X = U-35 = 4096 - 35 = 4061
			//     M = ((U-12)*32/255)-23 = ((4096-12)*32/255)-23 = 489.5
			//     K = M+((P-M)%(U-4)) = 489.5+((P-489.5)%(4092))
			//     if P <= X: no overflow
			//     elif P > X and K <= X: store K bytes on page, remaining overflows
			//     else: store first M bytes, remaining overflows
			end := ptr + int(numPayloadBytes)
			content := page[ptr:end]

			// read columns using the SQLite record format
			// https://www.sqlite.org/fileformat2.html#record_format
			contentOffset := 0

			var headerSize int
			contentOffset += internal.PutVarint(content[contentOffset:], &headerSize)
			columnTypes := []int{}
			for contentOffset < int(headerSize) {
				var serialType int
				contentOffset += internal.PutVarint(content[contentOffset:], &serialType)
				columnTypes = append(columnTypes, int(serialType))
			}
			if contentOffset > int(headerSize) {
				return nil, fmt.Errorf("consumed more header than expected! (%d>%d)", contentOffset, int(headerSize))
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

			if contentOffset != int(numPayloadBytes) {
				return nil, fmt.Errorf("did not consume all bytes in record (%d!=%d)", contentOffset, numPayloadBytes)
			}

			records = append(records, Record{
				key:     key,
				columns: columns,
			})
		default:
			return nil, fmt.Errorf("unsupported tree page type: %+v", typ)
		}
	}

	if nextPageNumber > 0 {
		children = append(children, &child{
			pageNumber: nextPageNumber,
		})
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

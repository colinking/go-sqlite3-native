package pager

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

// DefaultPageSize is the default size for DB pages, in bytes.
//
// Note that was originally 1024, but was updated to 4096 in 2016:
// https://www.sqlite.org/pgszchng2016.html
var DefaultPageSize int64 = 4096

// SQLiteHeader stores metadata on a SQLite database. This header is
// persisted in the first 100 bytes of the database file and therefore
// always available from page 1.
type SQLiteHeader struct {
	// 0+16: magic string ("SQLite format 3" in UTF-8)
	// ignored
	// 16+2
	PageSizeBytes int
	// 18+1
	FileFormatWriteVersion int
	// 19+1
	FileFormatReadVersion int
	// 20+1
	EOFPageByteReservation int
	// 21+1
	EmbeddedPayloadFractionMax int
	// 22+1
	EmbeddedPayloadFractionMin int
	// 23+1
	LeafPayloadFractionMin int
	// 24+4
	FileChangeCounter int
	// 28+4
	DatabaseSizePages int
	// 32+4
	FreelistFirstPageIndex int
	// 36+4
	FreelistNumPages int
	// 40+4
	SchemaCookieNumber int
	// 44+56
	VariadicElements []int
}

func (p *Pager) readHeader() error {
	// The header is entirely stored in first 100 bytes of the file.
	bytes := make([]byte, 100)
	_, err := p.file.ReadAt(bytes, 0)
	if err == io.EOF {
		// TODO: a zero length file _is_ valid, so we need to support the same defaults.
		return errors.New("reading empty files is not supported")
	} else if err != nil {
		return err
	}

	offset := 0

	magicString := bytes[offset : offset+16]
	offset += 16
	if string(magicString) != "SQLite format 3\x00" {
		return fmt.Errorf("invalid magic string (found: '%s'", magicString)
	}
	p.Header.PageSizeBytes = int(binary.BigEndian.Uint16(bytes[offset : offset+2]))
	offset += 2
	p.Header.FileFormatWriteVersion = int(bytes[offset])
	offset += 1
	p.Header.FileFormatReadVersion = int(bytes[offset])
	offset += 1
	p.Header.EOFPageByteReservation = int(bytes[offset])
	offset += 1
	p.Header.EmbeddedPayloadFractionMax = int(bytes[offset])
	offset += 1
	p.Header.EmbeddedPayloadFractionMin = int(bytes[offset])
	offset += 1
	p.Header.LeafPayloadFractionMin = int(bytes[offset])
	offset += 1
	p.Header.FileChangeCounter = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	p.Header.DatabaseSizePages = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	p.Header.FreelistFirstPageIndex = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	p.Header.FreelistNumPages = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	p.Header.SchemaCookieNumber = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	// offset += 4

	return nil
}

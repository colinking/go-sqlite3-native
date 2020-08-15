package internal

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
	PageSizeBytes int16
	// 18+1
	FileFormatWriteVersion int8
	// 19+1
	FileFormatReadVersion int8
	// 20+1
	EOFPageByteReservation int8
	// 21+1
	EmbeddedPayloadFractionMax int8
	// 22+1
	EmbeddedPayloadFractionMin int8
	// 23+1
	LeafPayloadFractionMin int8
	// 24+4
	FileChangeCounter int32
	// 28+4
	DatabaseSizePages int32
	// 32+4
	FreelistFirstPageIndex int32
	// 36+4
	FreelistNumPages int32
	// 40+4
	SchemaCookieNumber int32
	// 44+56
	VariadicElements []int32
}

func (db *DB) ReadHeader() error {
	// The header is entirely stored in first 100 bytes of the file.
	bytes := make([]byte, 100)
	_, err := db.file.ReadAt(bytes, 0)
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
	db.Header.PageSizeBytes = int16(binary.BigEndian.Uint16(bytes[offset : offset+2]))
	offset += 2
	db.Header.FileFormatWriteVersion = int8(bytes[offset])
	offset += 1
	db.Header.FileFormatReadVersion = int8(bytes[offset])
	offset += 1
	db.Header.EOFPageByteReservation = int8(bytes[offset])
	offset += 1
	db.Header.EmbeddedPayloadFractionMax = int8(bytes[offset])
	offset += 1
	db.Header.EmbeddedPayloadFractionMin = int8(bytes[offset])
	offset += 1
	db.Header.LeafPayloadFractionMin = int8(bytes[offset])
	offset += 1
	db.Header.FileChangeCounter = int32(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	db.Header.DatabaseSizePages = int32(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	db.Header.FreelistFirstPageIndex = int32(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	db.Header.FreelistNumPages = int32(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	db.Header.SchemaCookieNumber = int32(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	// offset += 4

	return nil
}

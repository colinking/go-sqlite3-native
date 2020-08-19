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
//
// https://www.sqlite.org/fileformat.html#the_database_header
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
	// 44+4
	SchemaFormatNumber int
	// 48+4
	DefaultPageCacheSize int
	// 52+4
	VacuumLargestRootTreePage int
	// 56+4
	TextEncoding int
	// 60+4
	UserVersion int
	// 64+4
	IncrementalVacuumEnabled int
	// 68+4
	ApplicationID int
	// 72+20
	// ignored
	// 92+4
	VersionValidFor int
	// 96+4
	SQLiteVersionNumber int
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
	p.header.PageSizeBytes = int(binary.BigEndian.Uint16(bytes[offset : offset+2]))
	offset += 2
	p.header.FileFormatWriteVersion = int(bytes[offset])
	offset += 1
	p.header.FileFormatReadVersion = int(bytes[offset])
	offset += 1
	p.header.EOFPageByteReservation = int(bytes[offset])
	offset += 1
	p.header.EmbeddedPayloadFractionMax = int(bytes[offset])
	offset += 1
	p.header.EmbeddedPayloadFractionMin = int(bytes[offset])
	offset += 1
	p.header.LeafPayloadFractionMin = int(bytes[offset])
	offset += 1
	p.header.FileChangeCounter = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	p.header.DatabaseSizePages = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	p.header.FreelistFirstPageIndex = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	p.header.FreelistNumPages = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	p.header.SchemaCookieNumber = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	p.header.SchemaFormatNumber = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	p.header.DefaultPageCacheSize = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	p.header.VacuumLargestRootTreePage = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	p.header.TextEncoding = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	p.header.UserVersion = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	p.header.IncrementalVacuumEnabled = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	p.header.ApplicationID = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4 + 20 // skip next 20 bytes, which are unused
	p.header.VersionValidFor = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	p.header.SQLiteVersionNumber = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	// offset += 4

	if p.header.FileFormatReadVersion != 2 || p.header.FileFormatWriteVersion != 2 {
		// PRAGMA journal_mode=WAL;
		return fmt.Errorf("legacy (rollback) journaling is unsupported: please enable WAL instead")
	}

	if p.header.TextEncoding != 1 {
		return fmt.Errorf("non-UTF-8 encodings are unsupported")
	}

	if p.header.FileChangeCounter != p.header.VersionValidFor {
		// https://www.sqlite.org/fileformat.html#in_header_database_size
		// if this happens, this means we cannot trust the database size in pages in the header.
		// In the future we could fallback to computing this based on the file size, which is
		// how the SQLite client handles this.
		return fmt.Errorf("this DB was modified by an old version of SQLite (<3.7.0)")
	}

	if p.header.SchemaFormatNumber != 4 {
		// https://www.sqlite.org/fileformat.html#schema_format_number
		// format #4 became the default in ~2006.
		return fmt.Errorf("unsupported schema format (%d)", p.header.SchemaFormatNumber)
	}

	return nil
}

package pager

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"

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
// https://www.sqlite.org/fileformat2.html#the_database_header
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
	EndOfPageByteReservation int
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

	header := SQLiteHeader{}

	magicString := bytes[offset : offset+16]
	offset += 16
	if string(magicString) != "SQLite format 3\x00" {
		return fmt.Errorf("invalid magic string (found: '%s'", magicString)
	}
	header.PageSizeBytes = int(binary.BigEndian.Uint16(bytes[offset : offset+2]))
	offset += 2
	header.FileFormatWriteVersion = int(bytes[offset])
	offset += 1
	header.FileFormatReadVersion = int(bytes[offset])
	offset += 1
	header.EndOfPageByteReservation = int(bytes[offset])
	offset += 1
	header.EmbeddedPayloadFractionMax = int(bytes[offset])
	offset += 1
	header.EmbeddedPayloadFractionMin = int(bytes[offset])
	offset += 1
	header.LeafPayloadFractionMin = int(bytes[offset])
	offset += 1
	header.FileChangeCounter = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	header.DatabaseSizePages = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	header.FreelistFirstPageIndex = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	header.FreelistNumPages = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	header.SchemaCookieNumber = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	header.SchemaFormatNumber = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	header.DefaultPageCacheSize = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	header.VacuumLargestRootTreePage = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	header.TextEncoding = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	header.UserVersion = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	header.IncrementalVacuumEnabled = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	header.ApplicationID = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4 + 20 // skip next 20 bytes, which are unused
	header.VersionValidFor = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	offset += 4
	header.SQLiteVersionNumber = int(binary.BigEndian.Uint32(bytes[offset : offset+4]))
	// offset += 4

	if header.FileFormatReadVersion != 2 || header.FileFormatWriteVersion != 2 {
		// PRAGMA journal_mode=WAL;
		// TODO: supporting rollback journaling should not require any work! confirm this
		return fmt.Errorf("legacy (rollback) journaling is unsupported: please enable WAL instead")
	}

	if header.TextEncoding != 1 {
		return fmt.Errorf("non-UTF-8 encodings are unsupported")
	}

	if header.FileChangeCounter != header.VersionValidFor {
		// https://www.sqlite.org/fileformat2.html#in_header_database_size
		// if this happens, this means we cannot trust the database size in pages in the header.
		// In the future we could fallback to computing this based on the file size, which is
		// how the SQLite client handles this.
		return fmt.Errorf("this DB was modified by an old version of SQLite (<3.7.0)")
	}

	if header.SchemaFormatNumber != 4 {
		// https://www.sqlite.org/fileformat2.html#schema_format_number
		// format #4 became the default in ~2006.
		return fmt.Errorf("unsupported schema format (%d)", header.SchemaFormatNumber)
	}

	if header.EndOfPageByteReservation > 0 {
		return fmt.Errorf("end-of-page reservations are not supported (f.e. SQLite encryption)")
	}

	// TODO: validate that vacuuming increases the schema cookie, which therefore means vacuuming causes no issues.

	p.header = &header

	return nil
}

func (h SQLiteHeader) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
Magic String:                      SQLite format 3
Page Size (bytes):                 %-20d
File Format (write):               %-20d # 2 = WAL, 1 = Rollback
File Format (read):                %-20d # 2 = WAL, 1 = Rollback
End of Page Reservation (bytes):   %-20d
Max Embedded Payload (%%):          %-20d
Min Embedded Payload (%%):          %-20d
Min Leaf Payload (%%):              %-20d
File Change Counter (FCC):         %-20d
Database Size (pages):             %-20d
Freelist First Page (page no):     %-20d
Freelist Size (pages):             %-20d
Schema Cookie Number:              %-20d
Schema Format Number:              %-20d
Suggested Page Cache Size (bytes): %-20d # Deprecated!
Vacuuming Next Tree (page no):     %-20d
Text Encoding:                     %-20d # 1 = UTF-8
User Version:                      %-20d
Incremental Vacuuming Enabled:     %-20d
Application ID:                    %-20d
Version Valid For FCC:             %-20d
SQLite Last Write Version:         %-20d
	`),
		h.PageSizeBytes,
		h.FileFormatWriteVersion,
		h.FileFormatReadVersion,
		h.EndOfPageByteReservation,
		h.EmbeddedPayloadFractionMax,
		h.EmbeddedPayloadFractionMin,
		h.LeafPayloadFractionMin,
		h.FileChangeCounter,
		h.DatabaseSizePages,
		h.FreelistFirstPageIndex,
		h.FreelistNumPages,
		h.SchemaCookieNumber,
		h.SchemaFormatNumber,
		h.DefaultPageCacheSize,
		h.VacuumLargestRootTreePage,
		h.TextEncoding,
		h.UserVersion,
		h.IncrementalVacuumEnabled,
		h.ApplicationID,
		h.VersionValidFor,
		h.SQLiteVersionNumber,
	)
}

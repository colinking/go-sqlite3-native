package internal

import (
	"os"

	"github.com/pkg/errors"
	"github.com/segmentio/events/v2"
)

type DB struct {
	Path   string
	Header SQLiteHeader

	file *os.File
	fd   uintptr
	pid  int32
}

func Open(path string) (*DB, error) {
	var err error
	db := &DB{
		Path: path,

		pid: int32(os.Getpid()),
	}

	events.Debug("opening SQLite DB at: %s", path)

	db.file, err = os.Open(path)
	if err != nil {
		return &DB{}, errors.Wrap(err, "opening file")
	}
	db.fd = db.file.Fd()

	if err = db.ReadHeader(); err != nil {
		return &DB{}, err
	}

	return db, nil
}

func (db *DB) Close() error {
	return db.file.Close()
}

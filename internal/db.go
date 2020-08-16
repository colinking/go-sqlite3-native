package internal

import (
	"github.com/colinking/sqlite3-experiments/internal/pager"
)

type DB struct {
	pager *pager.Pager
}

func Open(path string) (*DB, error) {
	var err error
	db := &DB{}

	db.pager, err = pager.NewPager(path)
	if err != nil {
		return &DB{}, err
	}

	return db, nil
}

func (db *DB) Header() pager.SQLiteHeader {
	return db.pager.Header
}

func (db *DB) Close() error {
	return db.pager.Close()
}

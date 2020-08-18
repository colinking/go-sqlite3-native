package main

import (
	"context"
	"database/sql"
	_ "net/http/pprof"
	"time"

	sqlite3native "github.com/colinking/go-sqlite3-native"
	"github.com/colinking/go-sqlite3-native/internal/pager"
	"github.com/segmentio/cli"
	"github.com/segmentio/events/v2"
	_ "github.com/segmentio/events/v2/ecslogs"
	_ "github.com/segmentio/events/v2/sigevents"
	_ "github.com/segmentio/events/v2/text"
)

func main() {
	cli.Exec(cli.CommandSet{
		"printHeader": cli.Command(printHeader),
		"lockStats":   cli.Command(lockStats),
		"query":       cli.Command(query),
	})
}

// Dumps the contents of a SQLite DB header.
func printHeader(_ struct{}, path string) (int, error) {
	db, err := sql.Open("sqlite3-native", path)
	if err != nil {
		return 1, err
	}
	defer db.Close()

	conn, err := db.Conn(context.Background())
	if err != nil {
		return 1, err
	}
	defer conn.Close()

	err = conn.Raw(func(driverConn interface{}) error {
		_conn := driverConn.(*sqlite3native.Conn)

		events.Log("header: %+v", _conn.Header())

		return nil
	})
	if err != nil {
		return 1, err
	}

	return 0, nil
}

// Records stats on lock usage for a SQLite DB.
func lockStats(_ struct{}, path string) (int, error) {
	p, err := pager.NewPager(path)
	if err != nil {
		return 1, err
	}
	defer p.Close()

	err = p.Lock(pager.LockTypeShared)
	if err != nil {
		return 1, err
	}
	defer func() {
		if err := p.Unlock(pager.LockTypeNoLock); err != nil {
			events.Log("failed to unlock: %+v", err)
		}
	}()

	time.Sleep(100 * time.Second)

	return 0, nil
}

// query does a sample query on the DB
func query(_ struct{}, path string) (int, error) {
	db, err := sql.Open("sqlite3-native", path)
	if err != nil {
		return 1, err
	}
	defer db.Close()

	rows, err := db.QueryContext(context.Background(), `SELECT column1 FROM table1;`)
	if err != nil {
		return 1, err
	}

	for rows.Next() {
		var column1 int
		if err := rows.Scan(&column1); err != nil {
			return 1, err
		}
		events.Log("column1: %d", column1)
	}
	if err := rows.Err(); err != nil {
		return 1, err
	}

	return 0, nil
}

package main

import (
	"context"
	"database/sql"
	_ "net/http/pprof"
	"time"

	"github.com/colinking/sqlite3native"
	"github.com/colinking/sqlite3native/internal/pager"
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

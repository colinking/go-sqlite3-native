package main

import (
	_ "net/http/pprof"

	_ "github.com/colinking/go-sqlite3-native"
	"github.com/colinking/go-sqlite3-native/internal/pager"
	"github.com/segmentio/cli"
	"github.com/segmentio/events/v2"
	_ "github.com/segmentio/events/v2/ecslogs"
	_ "github.com/segmentio/events/v2/sigevents"
	_ "github.com/segmentio/events/v2/text"
)

func main() {
	// TODO: make configurable with a flag
	events.DefaultLogger.EnableDebug = true

	// These commands are used for debugging the Go client on a SQLite DB.
	cli.Exec(cli.CommandSet{
		"printHeader": cli.Command(printHeader),
		// "lockStats":   cli.Command(lockStats),
		// TODO: generate trees as a graph
		// TODO: pretty print bytecode
		// TODO: generate parse tree as a graph
	})
}

// Dumps the contents of a SQLite DB header.
func printHeader(_ struct{}, path string) (int, error) {
	p, err := pager.NewPager(path)
	if err != nil {
		return 1, err
	}
	defer func() {
		if err := p.Close(); err != nil {
			events.Log("%+v", err)
		}
	}()

	header, err := p.Header()
	if err != nil {
		return 1, err
	}

	events.Log("header: %+v", header)

	return 0, nil
}

// Records stats on lock usage for a SQLite DB.
// func lockStats(_ struct{}, path string) (int, error) {
// 	p, err := pager.NewPager(path)
// 	if err != nil {
// 		return 1, err
// 	}
// 	defer p.Close()

// 	err = p.Lock(pager.LockTypeShared)
// 	if err != nil {
// 		return 1, err
// 	}
// 	defer func() {
// 		if err := p.Unlock(pager.LockTypeNoLock); err != nil {
// 			events.Log("failed to unlock: %+v", err)
// 		}
// 	}()

// 	time.Sleep(100 * time.Second)

// 	return 0, nil
// }

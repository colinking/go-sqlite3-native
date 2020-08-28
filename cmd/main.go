package main

import (
	"fmt"
	_ "net/http/pprof"

	_ "github.com/colinking/go-sqlite3-native"
	"github.com/colinking/go-sqlite3-native/internal/pager"
	"github.com/colinking/go-sqlite3-native/internal/tree"
	"github.com/segmentio/cli"
	"github.com/segmentio/events/v2"
	_ "github.com/segmentio/events/v2/ecslogs"
	_ "github.com/segmentio/events/v2/sigevents"
	_ "github.com/segmentio/events/v2/text"
)

func main() {
	// TODO: make configurable with a flag
	events.DefaultLogger.EnableDebug = false

	// These commands are used for debugging the Go client on a SQLite DB.
	cli.Exec(cli.CommandSet{
		"printHeader": cli.Command(printHeader),
		"printTree":   cli.Command(printTree),
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

	fmt.Printf("%+v\n", header)

	return 0, nil
}

// Dumps the contents of a tree.
func printTree(_ struct{}, path string) (int, error) {
	p, err := pager.NewPager(path)
	if err != nil {
		return 1, err
	}

	tm := tree.NewManager(p)
	defer func() {
		if err := tm.Close(); err != nil {
			events.Log("tree manager: %+v", err)
		}
	}()

	page := 2
	t, err := tm.Open(page)
	if err != nil {
		return 1, err
	}

	fmt.Printf("%+v\n", t)

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

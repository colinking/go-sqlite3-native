package main

import (
	"errors"
	"fmt"
	_ "net/http/pprof"
	"os"
	"strings"
	"time"

	"database/sql"
	"database/sql/driver"

	_ "github.com/colinking/go-sqlite3-native"
	"github.com/colinking/go-sqlite3-native/internal/pager"
	"github.com/colinking/go-sqlite3-native/internal/tree"
	"github.com/segmentio/cli"
	"github.com/segmentio/events/v2"
	_ "github.com/segmentio/events/v2/ecslogs"
	_ "github.com/segmentio/events/v2/sigevents"
	"github.com/segmentio/events/v2/text"
	_ "github.com/segmentio/events/v2/text"
)

func main() {
	// These commands are used for debugging the Go client on a SQLite DB.
	cli.Exec(cli.CommandSet{
		"printHeader": cli.Command(printHeader),
		"printTree":   cli.Command(printTree),
		"query":       cli.Command(query),
		// "lockStats":   cli.Command(lockStats),
		// TODO: generate trees as a graph
		// TODO: pretty print bytecode
		// TODO: generate parse tree as a graph
	})
}

type defaultFlags struct {
	Debug bool `flag:"--debug" help:"Enables more verbose debug logging"`
}

func (f defaultFlags) Load() {
	events.DefaultLogger.EnableDebug = f.Debug

	// Disable prefixes for text logging
	if events.IsTerminal(1) {
		events.DefaultHandler = &text.Handler{
			Output: os.Stdout,
		}
	}
}

// Dumps the contents of a SQLite DB header.
func printHeader(flags defaultFlags, path string) (int, error) {
	flags.Load()

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
func printTree(flags defaultFlags, path string, page int) (int, error) {
	flags.Load()

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

	t, err := tm.Open(page)
	if err != nil {
		return 1, err
	}

	fmt.Printf("%+v\n", t.String())
	if err := t.Err(); err != nil {
		return 1, err
	}

	return 0, nil
}

// Executes a query on a given DB.
func query(flags defaultFlags, path, query string) (int, error) {
	flags.Load()

	db, err := sql.Open("sqlite3-native", path)
	if err != nil {
		return 1, err
	}
	defer func() {
		if err := db.Close(); err != nil {
			events.Log("failed to close db: %+v", err)
		}
	}()

	stmt, err := db.Prepare(query)
	if err != nil {
		return 1, err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			events.Log("failed to close stmt: %+v", err)
		}
	}()

	rows, err := stmt.Query()
	if err != nil {
		return 1, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			events.Log("failed to close rows: %+v", err)
		}
	}()

	cols, err := rows.Columns()
	if err != nil {
		return 1, err
	}
	fmt.Printf("%s\n", strings.Join(cols, "|"))
	for rows.Next() {
		// To scan in an arbitrary length of items, we need to do some
		// pointer juggling:
		row := make([]driver.Value, len(cols))
		ptrs := make([]interface{}, len(cols))
		for i := range cols {
			ptrs[i] = &row[i]
		}
		err := rows.Scan(ptrs...)
		if err != nil {
			return 1, err
		}

		var content []string
		for _, col := range row {
			var s string
			switch c := col.(type) {
			case string:
				s = c
			case []byte:
				s = string(c)
			case float64:
				s = fmt.Sprintf("%f", c)
			case int64:
				s = fmt.Sprintf("%d", c)
			case time.Time:
				s = c.String()
			case bool:
				s = fmt.Sprintf("%v", c)
			default:
				return 1, errors.New("unknown column type")
			}
			content = append(content, s)
		}

		fmt.Printf("%s\n", strings.Join(content, "|"))
	}

	if err := rows.Err(); err != nil {
		return 1, err
	}

	return 0, nil
}

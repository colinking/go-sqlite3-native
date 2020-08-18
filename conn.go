package sqlite3native

import (
	"context"
	"database/sql/driver"

	"github.com/colinking/go-sqlite3-native/internal/pager"
	"github.com/colinking/go-sqlite3-native/internal/parser"
)

type Conn struct {
	pager *pager.Pager
}

var _ driver.Conn = &Conn{}
var _ driver.QueryerContext = &Conn{}
var _ driver.ConnPrepareContext = &Conn{}
var _ driver.ConnBeginTx = &Conn{}

// TODO: support these other, recommended, driver methods:
// var _ driver.Pinger = &Conn{}
// var _ driver.SessionResetter = &Conn{}
// var _ driver.ExecerContext = &Conn{}

// TODO: for Go 1.15, also include:
// var _ driver.Validator = &Conn{}

// func (c *Conn) Ping(ctx context.Context) error {
// 	panic("Ping not implemented")
// }

// func (c *Conn) ResetSession(ctx context.Context) error {
// 	panic("ResetSession not implemented")
// }

// func (c *Conn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
// 	panic("ExecContext not implemented")
// }

func (c *Conn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	_, err := parser.Parse(query)
	if err != nil {
		return nil, err
	}

	panic("QueryContext not implemented")
}

func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	return c.PrepareContext(context.Background(), query)
}

func (c *Conn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	panic("PrepareContext not implemented")
}

func (c *Conn) Begin() (driver.Tx, error) {
	return c.BeginTx(context.Background(), driver.TxOptions{})
}

func (c *Conn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	panic("BeginTx not implemented")
}

func (c *Conn) Close() error {
	return c.pager.Close()
}

func (c *Conn) Header() pager.SQLiteHeader {
	return c.pager.Header
}

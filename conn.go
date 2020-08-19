package sqlite3native

import (
	"context"
	"database/sql/driver"

	"github.com/colinking/go-sqlite3-native/internal/parser"
	"github.com/colinking/go-sqlite3-native/internal/vm"
	"github.com/segmentio/events/v2"
)

type Conn struct {
	vm *vm.VM
}

var _ driver.Conn = &Conn{}
var _ driver.QueryerContext = &Conn{}
var _ driver.ConnPrepareContext = &Conn{}
var _ driver.ConnBeginTx = &Conn{}
var _ driver.ExecerContext = &Conn{}

// TODO: support these other, recommended, driver methods:
// var _ driver.Pinger = &Conn{}
// var _ driver.SessionResetter = &Conn{}

// TODO: for Go 1.15, also include:
// var _ driver.Validator = &Conn{}

// func (c *Conn) Ping(ctx context.Context) error {
// 	panic("Ping not implemented")
// }

// func (c *Conn) ResetSession(ctx context.Context) error {
// 	panic("ResetSession not implemented")
// }

func (c *Conn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	stmt, err := c.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	s := stmt.(driver.StmtQueryContext)
	return s.QueryContext(ctx, args)
}

func (c *Conn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	stmt, err := c.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	s := stmt.(driver.StmtExecContext)
	return s.ExecContext(ctx, args)
}

func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	return c.PrepareContext(context.Background(), query)
}

func (c *Conn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	_, err := parser.Parse(query)
	if err != nil {
		return nil, err
	}

	// TODO: Right now, we don't use the parsed output since we haven't implemented the full parser.
	// Instead, we hardcode a VM program so that we can get the backend working.
	events.Log("Warning: ignoring query and using hardcoded VM program for testing")
	// This program only works on tmp/simple.db. It is equivalent to:
	/*
		sqlite> explain select * from table1;
		addr  opcode         p1    p2    p3    p4             p5  comment
		----  -------------  ----  ----  ----  -------------  --  -------------
		0     Init           0     7     0                    00  Start at 7
		1     OpenRead       0     2     0     1              00  root=2 iDb=0; table1
		2     Rewind         0     6     0                    00
		3       Column         0     0     1                    00  r[1]=table1.column1
		4       ResultRow      1     1     0                    00  output=r[1]
		5     Next           0     3     0                    01
		6     Halt           0     0     0                    00
		7     Transaction    0     0     1     0              01  usesStmtJournal=0
		8     Goto           0     1     0                    00
	*/
	program := vm.Program{
		Instructions: []vm.Instruction{
			{vm.OpcodeInit, 0, 7, 0, 0, 0},
			{vm.OpcodeOpenRead, 0, 2, 0, 1, 0},
			{vm.OpcodeRewind, 0, 6, 0, 0, 0},
			{vm.OpcodeColumn, 0, 0, 1, 0, 0},
			{vm.OpcodeResultRow, 1, 1, 0, 0, 0},
			{vm.OpcodeNext, 0, 3, 0, 0, 1},
			{vm.OpcodeHalt, 0, 0, 0, 0, 0},
			{vm.OpcodeTransaction, 0, 0, 1, 0, 1},
			{vm.OpcodeGoto, 0, 1, 0, 0, 0},
		},
		NumPlaceholders: 0,
		Columns:         []string{"column1"},
	}

	return &Stmt{
		conn:    c,
		program: program,
	}, nil
}

func (c *Conn) Begin() (driver.Tx, error) {
	return c.BeginTx(context.Background(), driver.TxOptions{})
}

func (c *Conn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	// TODO: implement begin
	panic("BeginTx not implemented")
}

func (c *Conn) Close() error {
	return c.vm.Close()
}

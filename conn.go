package sqlite3native

import (
	"context"
	"database/sql/driver"
	"fmt"

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
	var program vm.Program
	switch query {
	case "select * from table1;":
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
		program = vm.Program{
			Instructions: []vm.Instruction{
				vm.NewInstruction(vm.OpcodeInit, 0, 7, 0, 0, 0),
				vm.NewInstruction(vm.OpcodeOpenRead, 0, 2, 0, 1, 0),
				vm.NewInstruction(vm.OpcodeRewind, 0, 6, 0, 0, 0),
				vm.NewInstruction(vm.OpcodeColumn, 0, 0, 1, 0, 0),
				vm.NewInstruction(vm.OpcodeResultRow, 1, 1, 0, 0, 0),
				vm.NewInstruction(vm.OpcodeNext, 0, 3, 0, 0, 1),
				vm.NewInstruction(vm.OpcodeHalt, 0, 0, 0, 0, 0),
				vm.NewInstruction(vm.OpcodeTransaction, 0, 0, 1, 0, 1),
				vm.NewInstruction(vm.OpcodeGoto, 0, 1, 0, 0, 0),
			},
			NumPlaceholders: 0,
			Columns:         []string{"column1"},
		}
	case "select * from core___source_id_write_key_mapping where write_key=CAST('eE7e8Kpd7Xv6WJ8gzCofFh' AS BLOB);":
		// This only works on tmp/stage.db
		/*
			sqlite> explain select * from core___source_id_write_key_mapping where write_key=CAST('eE7e8Kpd7Xv6WJ8gzCofFh' AS BLOB);
			addr  opcode         p1    p2    p3    p4             p5  comment
			----  -------------  ----  ----  ----  -------------  --  -------------
			0     Init           0     13    0                    00  Start at 13
			1     OpenRead       0     81    0     2              00  root=81 iDb=0; core___source_id_write_key_mapping
			2     OpenRead       1     82    0     k(2,,)         02  root=82 iDb=0; sqlite_autoindex_core___source_id_write_key_mapping_1
			3     String8        0     1     0     eE7e8Kpd7Xv6WJ8gzCofFh  00  r[1]='eE7e8Kpd7Xv6WJ8gzCofFh'
			4     Cast           1     65    0                    00  affinity(r[1])
			5     IsNull         1     12    0                    00  if r[1]==NULL goto 12
			6     SeekGE         1     12    1     1              00  key=r[1]
			7     IdxGT          1     12    1     1              00  key=r[1]
			8     DeferredSeek   1     0     0                    00  Move 0 to 1.rowid if needed
			9     Column         1     0     2                    00  r[2]=core___source_id_write_key_mapping.write_key
			10    Column         0     1     3                    00  r[3]=core___source_id_write_key_mapping.source_id
			11    ResultRow      2     2     0                    00  output=r[2..3]
			12    Halt           0     0     0                    00
			13    Transaction    0     0     167193  0              01  usesStmtJournal=0
			14    Goto           0     1     0                    00
		*/
		program = vm.Program{
			Instructions: []vm.Instruction{
				vm.NewInstruction(vm.OpcodeInit, 0, 13, 0, 0, 0),
				vm.NewInstruction(vm.OpcodeOpenRead, 0, 81, 0, 2, 0),
				vm.NewInstruction(vm.OpcodeOpenRead, 1, 82, 0, 2, 2),
				vm.NewInstructionStr(vm.OpcodeString8, 0, 1, 0, "eE7e8Kpd7Xv6WJ8gzCofFh", 0),
				vm.NewInstruction(vm.OpcodeCast, 1, 65, 0, 0, 0),
				vm.NewInstruction(vm.OpcodeIsNull, 1, 12, 0, 0, 0),
				vm.NewInstruction(vm.OpcodeSeekGE, 1, 12, 1, 1, 0),
				vm.NewInstruction(vm.OpcodeIdxGT, 1, 12, 1, 1, 0),
				vm.NewInstruction(vm.OpcodeDeferredSeek, 1, 0, 0, 0, 0),
				vm.NewInstruction(vm.OpcodeColumn, 1, 0, 2, 0, 0),
				vm.NewInstruction(vm.OpcodeColumn, 0, 1, 3, 0, 0),
				vm.NewInstruction(vm.OpcodeResultRow, 2, 2, 0, 0, 0),
				vm.NewInstruction(vm.OpcodeHalt, 0, 0, 0, 0, 0),
				vm.NewInstruction(vm.OpcodeTransaction, 0, 0, 167193, 0, 1),
				vm.NewInstruction(vm.OpcodeGoto, 0, 1, 0, 0, 0),
			},
			NumPlaceholders: 0,
			Columns:         []string{"write_key", "source_id"},
		}
	case "select * from core___source_id_write_key_mapping;":
		// This only works on tmp/stage.db
		/*
			sqlite> explain select * from core___source_id_write_key_mapping;
			addr  opcode         p1    p2    p3    p4             p5  comment
			----  -------------  ----  ----  ----  -------------  --  -------------
			0     Init           0     8     0                    00  Start at 8
			1     OpenRead       0     81    0     2              00  root=81 iDb=0; core___source_id_write_key_mapping
			2     Rewind         0     7     0                    00
			3       Column         0     0     1                    00  r[1]=core___source_id_write_key_mapping.write_key
			4       Column         0     1     2                    00  r[2]=core___source_id_write_key_mapping.source_id
			5       ResultRow      1     2     0                    00  output=r[1..2]
			6     Next           0     3     0                    01
			7     Halt           0     0     0                    00
			8     Transaction    0     0     167193  0              01  usesStmtJournal=0
			9     Goto           0     1     0                    00
		*/
		program = vm.Program{
			Instructions: []vm.Instruction{
				vm.NewInstruction(vm.OpcodeInit, 0, 8, 0, 0, 0),
				vm.NewInstruction(vm.OpcodeOpenRead, 0, 81, 0, 2, 0),
				vm.NewInstruction(vm.OpcodeRewind, 0, 7, 0, 0, 0),
				vm.NewInstruction(vm.OpcodeColumn, 0, 0, 1, 0, 0),
				vm.NewInstruction(vm.OpcodeColumn, 0, 1, 2, 0, 0),
				vm.NewInstruction(vm.OpcodeResultRow, 1, 2, 0, 0, 0),
				vm.NewInstruction(vm.OpcodeNext, 0, 3, 0, 0, 1),
				vm.NewInstruction(vm.OpcodeHalt, 0, 0, 0, 0, 0),
				vm.NewInstruction(vm.OpcodeTransaction, 0, 0, 167193, 0, 1),
				vm.NewInstruction(vm.OpcodeGoto, 0, 1, 0, 0, 0),
			},
			NumPlaceholders: 0,
			Columns:         []string{"write_key", "source_id"},
		}
	default:
		return nil, fmt.Errorf("unsupported query: '%s'", query)
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

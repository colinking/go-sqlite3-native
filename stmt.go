package sqlite3native

import (
	"context"
	"database/sql/driver"

	"github.com/colinking/go-sqlite3-native/internal/vm"
)

type Stmt struct {
	conn            *Conn
	program         vm.Program
	numPlaceholders int
}

var _ driver.Stmt = &Stmt{}
var _ driver.StmtQueryContext = &Stmt{}
var _ driver.StmtExecContext = &Stmt{}

func (s *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	namedValues := make([]driver.NamedValue, len(args))
	for i, v := range args {
		namedValues[i] = driver.NamedValue{
			Ordinal: i + 1, // 1-indexed
			Value:   v,
		}
	}

	return s.QueryContext(context.Background(), namedValues)
}

func (s *Stmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	if err := s.conn.vm.Run(s.program); err != nil {
		return nil, err
	}

	panic("QueryContext not implemented")
}

func (s *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	namedValues := make([]driver.NamedValue, len(args))
	for i, v := range args {
		namedValues[i] = driver.NamedValue{
			Ordinal: i + 1, // 1-indexed
			Value:   v,
		}
	}

	return s.ExecContext(context.Background(), namedValues)
}

func (s *Stmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	panic("ExecContext not implemented")
}

func (s *Stmt) NumInput() int {
	return s.numPlaceholders
}

func (s *Stmt) Close() error {
	return nil
}

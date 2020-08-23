package sqlite3native

import (
	"database/sql/driver"
	"io"

	"github.com/colinking/go-sqlite3-native/internal/vm"
	"github.com/segmentio/events/v2"
)

type Rows struct {
	program   vm.Program
	execution *vm.Execution
}

var _ driver.Rows = &Rows{}

func (r *Rows) Columns() []string {
	return r.program.Columns
}

func (r *Rows) Next(dest []driver.Value) error {
	t, err := r.execution.Next()
	events.Debug("received row: %+v %+v", t, err)
	if err != nil {
		return err
	}

	if t == nil {
		return io.EOF
	}
	columns := *t

	for i := range columns {
		dest[i] = columns[i].Value()
	}

	return nil
}

func (r *Rows) Close() error {
	return r.execution.Close()
}

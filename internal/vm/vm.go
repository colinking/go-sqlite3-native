package vm

import (
	"fmt"

	"github.com/colinking/go-sqlite3-native/internal/tree"
	"github.com/segmentio/events/v2"
)

const (
	BufferSize = 100
)

type VM struct {
	tm *tree.TreeManager
}

func NewVM(tm *tree.TreeManager) *VM {
	return &VM{
		tm: tm,
	}
}

func (m *VM) Close() error {
	return m.tm.Close()
}

type Execution struct {
	program Program
	tm      *tree.TreeManager
	results chan []tree.Column
	done    chan error
}

func (m *VM) Execute(program Program) *Execution {
	e := &Execution{
		program: program,

		tm:      m.tm,
		results: make(chan []tree.Column, BufferSize),
		done:    make(chan error),
	}

	// A VM program is executed in a separate goroutine where results are
	// buffered up and returned to the upstream caller.
	go e.run()

	return e
}

func (e *Execution) run() {
	activeTreeIndex := -1
	trees := []*tree.Tree{}
	row := []tree.Column{}

	events.Debug("Executing program:\n%s", e.program)

	for pc := 0; pc < len(e.program.Instructions); pc++ {
		inst := e.program.Instructions[pc]

		// Opcodes are explained in the SQLite docs here: https://www.sqlite.org/opcode.html
		switch inst.Op {
		case OpcodeInit: // https://www.sqlite.org/opcode.html#Init
			if inst.P2 > 0 {
				pc = inst.P2
				pc-- // negate pc++
			}
		case OpcodeHalt: // https://www.sqlite.org/opcode.html#Halt
			pc = len(e.program.Instructions)
		case OpcodeTransaction: // https://www.sqlite.org/opcode.html#Transaction
			// TODO: issue a Begin query to the pager to acquire the read lock

			// Start a read transaction by looking up the header:
			header, err := e.tm.Header()
			if err != nil {
				e.done <- err
				return
			}

			// By definition, we only check these header values if P5!=0.
			if inst.P5 != 0 {
				// Verify the schema cookie
				if inst.P3 != header.SchemaCookieNumber {
					// this is a SQLITE_SCHEMA error indicating the program should be re-compiled.
					e.done <- fmt.Errorf("invalid schema cookie number: expected %d got %d", inst.P3, header.SchemaCookieNumber)
					return
				}

				// TODO: there's some kind of "schema generation counter" to validate here that is not well-defined.
			}
		case OpcodeGoto: // https://www.sqlite.org/opcode.html#Goto
			pc = inst.P2
			pc-- // negate pc++
		case OpcodeOpenRead: // https://www.sqlite.org/opcode.html#OpenRead
			if inst.P3 != 0 {
				// We don't need temporary tables because we don't support complex JOINs.
				// And we don't support ATTACH-ing other databases. Therefore this should always
				// be on the main database.
				e.done <- fmt.Errorf("operations on databases other than main are not supported: %d", inst.P3)
				return
			}

			cursorID := inst.P1
			rootPageNumber := inst.P2
			// numColumns := inst.P4

			t, err := e.tm.Open(rootPageNumber)
			if err != nil {
				e.done <- err
				return
			}

			if cursorID < cap(trees) {
				trees[cursorID] = t
			} else if cursorID == cap(trees) {
				trees = append(trees, t)
			} else {
				e.done <- fmt.Errorf("failed to insert tree (len=%d cap=%d cursorID=%d)", len(trees), cap(trees), cursorID)
				return
			}

			// TODO: assert the opened b-tree has numColumns

		case OpcodeRewind: // https://www.sqlite.org/opcode.html#Rewind
			treeIdx := inst.P1
			activeTreeIndex = treeIdx // signal which tree to use, to the next call to Column/etc.
			tree := trees[treeIdx]
			tree.ResetCursor()

			if !tree.Next() {
				// If there are _no_ more rows to read, skip to:
				pc = inst.P2
				pc-- // negate pc++
			}

		case OpcodeColumn: // https://www.sqlite.org/opcode.html#Column
			tree := trees[activeTreeIndex]
			columnIdx := inst.P2
			column := tree.Get().GetColumn(columnIdx)
			row = append(row, column)

		case OpcodeResultRow: // https://www.sqlite.org/opcode.html#ResultRow
			events.Debug("result row: %+v", row)
			e.results <- row

			// Reset the row
			row = []tree.Column{}

		case OpcodeNext: // https://www.sqlite.org/opcode.html#Next
			cursorID := inst.P1
			tree := trees[cursorID]
			if tree.Next() {
				// If there are _more_ rows to read, skip to:
				pc = inst.P2
				pc-- // negate pc++
			}

		default:
			e.done <- fmt.Errorf("unknown opcode! %+v", inst)
			return
		}
	}

	e.done <- nil
}

// Next returns the next available tuple produced by executing this VM program.
//
// If a nil error and nil tuple are returned, that means that all rows have been
// read successfully. A non-nil error indicates an error while executing the bytecode.
func (e *Execution) Next() (*[]tree.Column, error) {
	select {
	case t := <-e.results:
		return &t, nil
	case err := <-e.done:
		return nil, err
	}
}

func (e *Execution) Close() error {
	close(e.results)
	close(e.done)

	return nil
}

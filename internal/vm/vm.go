package vm

import (
	"fmt"

	"github.com/colinking/go-sqlite3-native/internal/tree"
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
		done:    make(chan error, 1),
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
	registers := &Registers{}

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

			// TODO: consider incorporating P5's OPFLAG_SEEKEQ to optimize tree lookups

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

		case OpcodeString8: // https://www.sqlite.org/opcode.html#String8
			s := inst.P4.s
			idx := inst.P2
			registers.SetString(idx, s)

		case OpcodeCast: // https://www.sqlite.org/opcode.html#Cast
			idx := inst.P1
			typ := inst.P2

			var err error
			switch typ {
			case 'A': // BLOB
				err = registers.CastAsBlob(idx)
			case 'B': // TEXT
				err = registers.CastAsString(idx)
			case 'D': // INTEGER
				err = registers.CastAsInt(idx)
			case 'E': // REAL
				err = registers.CastAsFloat(idx)
			default:
				e.done <- fmt.Errorf("unknown/unsupported typ=%+v", typ)
				return
			}

			if err != nil {
				e.done <- err
				return
			}

		case OpcodeIsNull: // https://www.sqlite.org/opcode.html#IsNull
			idx := inst.P1
			r := registers.Get(idx)
			if r.typ == RegisterTypeNull {
				pc = inst.P2
				pc-- // negate pc++
			}

		case OpcodeSeekGE: // https://www.sqlite.org/opcode.html#SeekGE
			cursorIdx := inst.P1
			keyIdx := inst.P3
			// TODO: nKeys := inst.P4

			cursor := trees[cursorIdx]
			key := registers.Get(keyIdx).Blob
			err := cursor.SeekGE(key)
			if err != nil {
				e.done <- err
				return
			}

		case OpcodeIdxGT: // https://www.sqlite.org/opcode.html#IdxGT
			e.done <- fmt.Errorf("todo: support IdxGT! %+v", inst)
			return

		case OpcodeDeferredSeek: // https://www.sqlite.org/opcode.html#DeferredSeek
			e.done <- fmt.Errorf("todo: support DeferredSeek! %+v", inst)
			return

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
		// If we call Next() when both e.results and e.done are ready, then we'll receive
		// a result at random. We want to guarantee that if both are ready, we always prioritize
		// receiving from e.results. To do that, if we get a result from e.done we do a non-blocking
		// receive on e.results to see if it was ready to. If so, then we return that instead.
		select {
		case t := <-e.results:
			// Place the result we got from e.done back into its channel so a future Next() call
			// will receive it:
			e.done <- err

			return &t, nil
		default:
			return nil, err
		}
	}
}

func (e *Execution) Close() error {
	close(e.results)
	close(e.done)

	return nil
}

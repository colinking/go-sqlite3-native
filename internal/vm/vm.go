package vm

import (
	"fmt"

	"github.com/colinking/go-sqlite3-native/internal/tree"
)

const TupleBufferSize = 100

type VM struct {
	tree *tree.Tree
}

func NewVM(tree *tree.Tree) *VM {
	return &VM{
		tree: tree,
	}
}

func (m *VM) Close() error {
	return m.tree.Close()
}

type Execution struct {
	program Program
	tree    *tree.Tree
	results chan *Tuple
	done    chan error
}

type Tuple struct {
	// TODO
}

func (m *VM) Execute(program Program) *Execution {
	e := &Execution{
		program: program,

		tree:    m.tree,
		results: make(chan *Tuple, TupleBufferSize),
		done:    make(chan error),
	}

	// A VM program is executed in a separate goroutine where results are
	// buffered up and returned to the upstream caller.
	go e.run()

	return e
}

func (e *Execution) run() {
	// TODO: array of open tree cursors
	// TODO: array of memory locations, referenced by the program

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
			// Start a read transaction by looking up the header:
			header, err := e.tree.Header()
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

				// NOTE: there's some kind of "schema generation counter" to validate here that
				// is not well-defined.
			}
		case OpcodeGoto: // https://www.sqlite.org/opcode.html#Goto
			pc = inst.P2
			pc-- // negate pc++

		// TODO: other opcodes
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
func (e *Execution) Next() (*Tuple, error) {
	select {
	case err := <-e.done:
		return nil, err
	case t := <-e.results:
		return t, nil
	}
}

func (e *Execution) Close() error {
	close(e.results)
	close(e.done)

	return nil
}

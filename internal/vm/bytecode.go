package vm

import "fmt"

type Program struct {
	Instructions    []Instruction
	NumPlaceholders int
	Columns         []string
}

func (p Program) String() string {
	s := ""
	for i, inst := range p.Instructions {
		s += fmt.Sprintf("%d: %s\n", i, inst)
	}
	return s
}

type Instruction struct {
	Op Opcode

	P1 int
	P2 int
	P3 int
	P4 int
	P5 int
}

func (i Instruction) String() string {
	return fmt.Sprintf("%s(P1: %d, P2: %d, P3: %d, P4: %d, P5: %d)", i.Op.String(), i.P1, i.P2, i.P3, i.P4, i.P5)
}

//go:generate stringer -type=Opcode
type Opcode int

const (
	OpcodeInit Opcode = iota
	OpcodeOpenRead
	OpcodeString8
	OpcodeCast
	OpcodeIsNull
	OpcodeSeekGE
	OpcodeIdxGT
	OpcodeDeferredSeek
	OpcodeColumn
	OpcodeResultRow
	OpcodeHalt
	OpcodeTransaction
	OpcodeGoto
	OpcodeNext
	OpcodeRewind
)

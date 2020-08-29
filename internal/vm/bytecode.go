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
	P4 struct {
		i int
		s string
	}
	P5 int
}

func NewInstruction(op Opcode, p1, p2, p3, p4, p5 int) Instruction {
	in := Instruction{
		Op: op,
		P1: p1,
		P2: p2,
		P3: p3,
		P5: p5,
	}
	in.P4.i = p4

	return in
}

func NewInstructionStr(op Opcode, p1, p2, p3 int, p4 string, p5 int) Instruction {
	in := Instruction{
		Op: op,
		P1: p1,
		P2: p2,
		P3: p3,
		P5: p5,
	}
	in.P4.s = p4

	return in
}

func (i Instruction) String() string {
	p4 := i.P4.s
	if p4 == "" {
		p4 = fmt.Sprintf("%d", i.P4.i)
	}

	return fmt.Sprintf("%s(P1: %d, P2: %d, P3: %d, P4: %s, P5: %d)", i.Op.String(), i.P1, i.P2, i.P3, p4, i.P5)
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

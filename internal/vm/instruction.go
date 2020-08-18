package vm

type Instruction struct {
	Op Opcode

	P1 int
	P2 int
	P3 int
	P4 int
	P5 int
}

type Program struct {
	Instructions []Instruction
}

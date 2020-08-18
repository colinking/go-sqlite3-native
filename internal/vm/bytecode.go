package vm

type Program struct {
	Instructions []Instruction
}

type Instruction struct {
	Op Opcode

	P1 int
	P2 int
	P3 int
	P4 int
	P5 int
}

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

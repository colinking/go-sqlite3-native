package vm

import "github.com/segmentio/events/v2"

type VM struct {
	// TODO
}

func NewVM() *VM {
	return &VM{}
}

func (m *VM) Run(program Program) error {
	events.Log("program: %+v", program)

	return nil
}

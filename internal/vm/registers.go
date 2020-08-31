package vm

import "fmt"

type Registers struct {
	Registers []Register
}

//go:generate stringer -type=RegisterType
type RegisterType int

const (
	RegisterTypeUnknown RegisterType = iota
	RegisterTypeNull
	RegisterTypeInt
	RegisterTypeFloat
	RegisterTypeString
	RegisterTypeBlob
)

type Register struct {
	typ RegisterType

	Int    int
	Float  float64
	String string
	Blob   []byte
}

func (r *Registers) Get(idx int) Register {
	if idx < cap(r.Registers) {
		return r.Registers[idx]
	}

	return Register{}
}

func (r *Registers) SetInt(idx int, i int) {
	r.resize(idx)
	r.Registers[idx].typ = RegisterTypeInt
	r.Registers[idx].Int = i
}

func (r *Registers) CastAsInt(idx int) error {
	to := RegisterTypeInt

	if idx >= cap(r.Registers) {
		return fmt.Errorf("unknown register at idx=%d, unable to cast as %s", idx, to.String())
	}

	switch r.Registers[idx].typ {
	case RegisterTypeNull:
		// casting a null is a no-op
	case RegisterTypeFloat:
		r.Registers[idx].typ = to
		r.Registers[idx].Int = int(r.Registers[idx].Float)
	// case RegisterTypeString:
	// 	r.Registers[idx].typ = to
	// 	r.Registers[idx].Int = int(r.Registers[idx].String)
	// case RegisterTypeBlob:
	// 	r.Registers[idx].typ = to
	// 	r.Registers[idx].Int = int(r.Registers[idx].Blob)
	default:
		return fmt.Errorf("unsupported cast at idx=%d from typ=%s to typ=%s", idx, r.Registers[idx].typ.String(), to.String())
	}

	return nil
}

func (r *Registers) SetFloat(idx int, f float64) {
	r.resize(idx)
	r.Registers[idx].typ = RegisterTypeFloat
	r.Registers[idx].Float = f
}

func (r *Registers) CastAsFloat(idx int) error {
	to := RegisterTypeFloat

	if idx >= cap(r.Registers) {
		return fmt.Errorf("unknown register at idx=%d, unable to cast as %s", idx, to.String())
	}

	switch r.Registers[idx].typ {
	case RegisterTypeNull:
		// casting a null is a no-op
	case RegisterTypeInt:
		r.Registers[idx].typ = to
		r.Registers[idx].Float = float64(r.Registers[idx].Int)
	// case RegisterTypeString:
	// 	r.Registers[idx].typ = to
	// 	r.Registers[idx].Float = float64(r.Registers[idx].String)
	// case RegisterTypeBlob:
	// 	r.Registers[idx].typ = to
	// 	r.Registers[idx].Float = float64(r.Registers[idx].Blob)
	default:
		return fmt.Errorf("unsupported cast at idx=%d from typ=%s to typ=%s", idx, r.Registers[idx].typ.String(), to.String())
	}

	return nil
}

func (r *Registers) SetString(idx int, s string) {
	r.resize(idx)
	r.Registers[idx].typ = RegisterTypeString
	r.Registers[idx].String = s
}

func (r *Registers) CastAsString(idx int) error {
	to := RegisterTypeString

	if idx >= cap(r.Registers) {
		return fmt.Errorf("unknown register at idx=%d, unable to cast as %s", idx, to.String())
	}

	switch r.Registers[idx].typ {
	case RegisterTypeNull:
		// casting a null is a no-op
	case RegisterTypeInt:
		r.Registers[idx].typ = to
		r.Registers[idx].String = fmt.Sprintf("%d", r.Registers[idx].Int)
	case RegisterTypeFloat:
		r.Registers[idx].typ = to
		r.Registers[idx].String = fmt.Sprintf("%f", r.Registers[idx].Float)
	case RegisterTypeBlob:
		r.Registers[idx].typ = to
		r.Registers[idx].String = string(r.Registers[idx].Blob)
	default:
		return fmt.Errorf("unsupported cast at idx=%d from typ=%s to typ=%s", idx, r.Registers[idx].typ.String(), to.String())
	}

	return nil
}

func (r *Registers) SetBlob(idx int, b []byte) {
	r.resize(idx)
	r.Registers[idx].typ = RegisterTypeBlob
	r.Registers[idx].Blob = b
}

func (r *Registers) CastAsBlob(idx int) error {
	to := RegisterTypeBlob

	if idx >= cap(r.Registers) {
		return fmt.Errorf("unknown register at idx=%d, unable to cast as %s", idx, to.String())
	}

	switch r.Registers[idx].typ {
	case RegisterTypeNull:
		// casting a null is a no-op
	// case RegisterTypeInt:
	// 	r.Registers[idx].typ = to
	// 	r.Registers[idx].Blob = []byte(r.Registers[idx].Int)
	// case RegisterTypeFloat:
	// 	r.Registers[idx].typ = to
	// 	r.Registers[idx].Blob = []byte(r.Registers[idx].Float)
	case RegisterTypeString:
		r.Registers[idx].typ = to
		r.Registers[idx].Blob = []byte(r.Registers[idx].String)
	default:
		return fmt.Errorf("unsupported cast at idx=%d from typ=%s to typ=%s", idx, r.Registers[idx].typ.String(), to.String())
	}

	return nil
}

func (r *Registers) SetNull(idx int) {
	r.resize(idx)
	r.Registers[idx].typ = RegisterTypeNull
}

func (r *Registers) resize(idx int) {
	if idx < cap(r.Registers) {
		return
	}

	// Double (or more) the storage capacity, as necessary.
	n := len(r.Registers)
	for idx >= n {
		if n == 0 {
			n = 1
		} else {
			n = 2 * n
		}
	}
	regs := make([]Register, n)
	copy(regs, r.Registers)
	r.Registers = regs
}

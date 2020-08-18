package parser

import (
	"testing"

	"github.com/colinking/go-sqlite3-native/internal/vm"
	"github.com/stretchr/testify/require"
)

func TestParser(tt *testing.T) {
	p := NewParser()

	for _, test := range []struct {
		name    string
		sql     string
		program []vm.Instruction
	}{
		{
			name:    "simple select",
			sql:     `SELECT * FROM table1`,
			program: []vm.Instruction{},
		},
	} {
		tt.Run(test.name, func(t *testing.T) {
			result, err := p.Parse(test.sql)
			require.NoError(t, err)

			require.Equal(t, test.program, result)
		})
	}
}

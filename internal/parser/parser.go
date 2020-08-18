package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/colinking/go-sqlite3-native/internal/parser/generated"
)

//go:generate antlr -Dlanguage=Go -o generated -package generated SQL.g4

type Parser struct {
	// TODO
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(query string) (struct{}, error) {
	// TODO
	is := antlr.NewInputStream(query)
	_ = generated.NewSQLLexer(is)

	// for {
	// 	t := lexer.NextToken()
	// }

	return struct{}{}, nil
}

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
	is := antlr.NewInputStream(query)

	// Create a lexer which can take arbitrary user-supplied strings and convert them
	// into tokens that we can produce a parse tree on.
	gLexer := generated.NewSQLLexer(is)
	stream := antlr.NewCommonTokenStream(gLexer, antlr.LexerDefaultTokenChannel)

	// Create a parser that can consume the list of tokens and produce a parse tree that we can walk:
	gParser := generated.NewSQLParser(stream)

	// Walk through the parse tree. This walk will invoke methods on the listener
	// which we can catch in order to produce our bytecode program.
	antlr.ParseTreeWalkerDefault.Walk(&generated.BaseSQLListener{}, gParser.Start())

	return struct{}{}, nil
}

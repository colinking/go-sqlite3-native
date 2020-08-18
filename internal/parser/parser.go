package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/colinking/go-sqlite3-native/internal/parser/generated"
	"github.com/colinking/go-sqlite3-native/internal/vm"
)

//go:generate antlr -Dlanguage=Go -o generated -package generated SQL.g4

func Parse(query string) ([]vm.Instruction, error) {
	// This parser is based on the antlr language and uses the official Go antlr runtime.
	// For more information on how this works, see: https://blog.gopheracademy.com/advent-2017/parsing-with-antlr4-and-go/
	// Further inspiration was taken from the unofficial SQLite antlr grammar: https://github.com/antlr/grammars-v4/blob/master/sql/sqlite/SQLite.g4
	// Along with the official, but Lemon-based, SQLite grammar: https://github.com/sqlite/sqlite/blob/master/src/parse.y

	is := antlr.NewInputStream(query)

	// Create a lexer which can take arbitrary user-supplied strings and convert them
	// into tokens that we can produce a parse tree on.
	lexer := generated.NewSQLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)

	// Create a parser that can consume the list of tokens and produce a parse tree that we can walk:
	parser := generated.NewSQLParser(stream)

	// Create a listener that we will use to hook into antlr's runtime as it walks through
	// the parse tree.
	l := listener{
		program: []vm.Instruction{},
	}

	// Walk through the parse tree. This walk will invoke methods on the listener
	// which we can catch in order to produce our bytecode program.
	antlr.ParseTreeWalkerDefault.Walk(&l, parser.Start())

	return l.program, nil
}

type listener struct {
	*generated.BaseSQLListener

	program []vm.Instruction
}

var _ generated.SQLListener = &listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterStart is called when production start is entered.
func (s *listener) EnterStart(ctx *generated.StartContext) {
	// TODO: add the basic structure
}

// ExitStart is called when production start is exited.
func (s *listener) ExitStart(ctx *generated.StartContext) {}

// EnterExpression is called when production expression is entered.
func (s *listener) EnterExpression(ctx *generated.ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *listener) ExitExpression(ctx *generated.ExpressionContext) {}

// EnterSelectExpression is called when production selectExpression is entered.
func (s *listener) EnterSelectExpression(ctx *generated.SelectExpressionContext) {}

// ExitSelectExpression is called when production selectExpression is exited.
func (s *listener) ExitSelectExpression(ctx *generated.SelectExpressionContext) {}

// EnterArgs is called when production args is entered.
func (s *listener) EnterArgs(ctx *generated.ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *listener) ExitArgs(ctx *generated.ArgsContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *listener) EnterIdentifier(ctx *generated.IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *listener) ExitIdentifier(ctx *generated.IdentifierContext) {}

// EnterIdentifierEnd is called when production identifierEnd is entered.
func (s *listener) EnterIdentifierEnd(ctx *generated.IdentifierEndContext) {}

// ExitIdentifierEnd is called when production identifierEnd is exited.
func (s *listener) ExitIdentifierEnd(ctx *generated.IdentifierEndContext) {}

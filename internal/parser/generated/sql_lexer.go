// Code generated from SQL.g4 by ANTLR 4.8. DO NOT EDIT.

package generated

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 22, 122,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 6, 15, 99, 10, 15, 13, 15, 14, 15, 100,
	3, 16, 6, 16, 104, 10, 16, 13, 16, 14, 16, 105, 3, 17, 3, 17, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 6, 21, 117, 10, 21, 13, 21, 14,
	21, 118, 3, 21, 3, 21, 2, 2, 22, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8,
	15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 3, 2, 5, 3, 2, 50, 59, 4, 2, 67,
	92, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 124, 2, 3, 3, 2, 2, 2, 2,
	5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2,
	13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2,
	2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2,
	2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2,
	2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 3, 43, 3,
	2, 2, 2, 5, 50, 3, 2, 2, 2, 7, 55, 3, 2, 2, 2, 9, 61, 3, 2, 2, 2, 11, 67,
	3, 2, 2, 2, 13, 70, 3, 2, 2, 2, 15, 74, 3, 2, 2, 2, 17, 79, 3, 2, 2, 2,
	19, 85, 3, 2, 2, 2, 21, 87, 3, 2, 2, 2, 23, 89, 3, 2, 2, 2, 25, 93, 3,
	2, 2, 2, 27, 95, 3, 2, 2, 2, 29, 98, 3, 2, 2, 2, 31, 103, 3, 2, 2, 2, 33,
	107, 3, 2, 2, 2, 35, 109, 3, 2, 2, 2, 37, 111, 3, 2, 2, 2, 39, 113, 3,
	2, 2, 2, 41, 116, 3, 2, 2, 2, 43, 44, 7, 85, 2, 2, 44, 45, 7, 71, 2, 2,
	45, 46, 7, 78, 2, 2, 46, 47, 7, 71, 2, 2, 47, 48, 7, 69, 2, 2, 48, 49,
	7, 86, 2, 2, 49, 4, 3, 2, 2, 2, 50, 51, 7, 72, 2, 2, 51, 52, 7, 84, 2,
	2, 52, 53, 7, 81, 2, 2, 53, 54, 7, 79, 2, 2, 54, 6, 3, 2, 2, 2, 55, 56,
	7, 89, 2, 2, 56, 57, 7, 74, 2, 2, 57, 58, 7, 71, 2, 2, 58, 59, 7, 84, 2,
	2, 59, 60, 7, 71, 2, 2, 60, 8, 3, 2, 2, 2, 61, 62, 7, 81, 2, 2, 62, 63,
	7, 84, 2, 2, 63, 64, 7, 70, 2, 2, 64, 65, 7, 71, 2, 2, 65, 66, 7, 84, 2,
	2, 66, 10, 3, 2, 2, 2, 67, 68, 7, 68, 2, 2, 68, 69, 7, 91, 2, 2, 69, 12,
	3, 2, 2, 2, 70, 71, 7, 67, 2, 2, 71, 72, 7, 85, 2, 2, 72, 73, 7, 69, 2,
	2, 73, 14, 3, 2, 2, 2, 74, 75, 7, 70, 2, 2, 75, 76, 7, 71, 2, 2, 76, 77,
	7, 85, 2, 2, 77, 78, 7, 69, 2, 2, 78, 16, 3, 2, 2, 2, 79, 80, 7, 78, 2,
	2, 80, 81, 7, 75, 2, 2, 81, 82, 7, 79, 2, 2, 82, 83, 7, 75, 2, 2, 83, 84,
	7, 86, 2, 2, 84, 18, 3, 2, 2, 2, 85, 86, 7, 44, 2, 2, 86, 20, 3, 2, 2,
	2, 87, 88, 7, 65, 2, 2, 88, 22, 3, 2, 2, 2, 89, 90, 7, 67, 2, 2, 90, 91,
	7, 80, 2, 2, 91, 92, 7, 70, 2, 2, 92, 24, 3, 2, 2, 2, 93, 94, 7, 63, 2,
	2, 94, 26, 3, 2, 2, 2, 95, 96, 7, 64, 2, 2, 96, 28, 3, 2, 2, 2, 97, 99,
	9, 2, 2, 2, 98, 97, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 98, 3, 2, 2,
	2, 100, 101, 3, 2, 2, 2, 101, 30, 3, 2, 2, 2, 102, 104, 9, 3, 2, 2, 103,
	102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 105, 106,
	3, 2, 2, 2, 106, 32, 3, 2, 2, 2, 107, 108, 7, 46, 2, 2, 108, 34, 3, 2,
	2, 2, 109, 110, 7, 42, 2, 2, 110, 36, 3, 2, 2, 2, 111, 112, 7, 43, 2, 2,
	112, 38, 3, 2, 2, 2, 113, 114, 7, 61, 2, 2, 114, 40, 3, 2, 2, 2, 115, 117,
	9, 4, 2, 2, 116, 115, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 116, 3, 2,
	2, 2, 118, 119, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 121, 8, 21, 2, 2,
	121, 42, 3, 2, 2, 2, 6, 2, 100, 105, 118, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'SELECT'", "'FROM'", "'WHERE'", "'ORDER'", "'BY'", "'ASC'", "'DESC'",
	"'LIMIT'", "'*'", "'?'", "'AND'", "'='", "'>'", "", "", "','", "'('", "')'",
	"';'",
}

var lexerSymbolicNames = []string{
	"", "Select", "From", "Where", "Order", "By", "Asc", "Desc", "Limit", "Star",
	"Placeholder", "And", "Equal", "Greater", "Number", "Letter", "Comma",
	"LParen", "RParen", "Semicolon", "WHITESPACE",
}

var lexerRuleNames = []string{
	"Select", "From", "Where", "Order", "By", "Asc", "Desc", "Limit", "Star",
	"Placeholder", "And", "Equal", "Greater", "Number", "Letter", "Comma",
	"LParen", "RParen", "Semicolon", "WHITESPACE",
}

type SQLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewSQLLexer(input antlr.CharStream) *SQLLexer {

	l := new(SQLLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "SQL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SQLLexer tokens.
const (
	SQLLexerSelect      = 1
	SQLLexerFrom        = 2
	SQLLexerWhere       = 3
	SQLLexerOrder       = 4
	SQLLexerBy          = 5
	SQLLexerAsc         = 6
	SQLLexerDesc        = 7
	SQLLexerLimit       = 8
	SQLLexerStar        = 9
	SQLLexerPlaceholder = 10
	SQLLexerAnd         = 11
	SQLLexerEqual       = 12
	SQLLexerGreater     = 13
	SQLLexerNumber      = 14
	SQLLexerLetter      = 15
	SQLLexerComma       = 16
	SQLLexerLParen      = 17
	SQLLexerRParen      = 18
	SQLLexerSemicolon   = 19
	SQLLexerWHITESPACE  = 20
)

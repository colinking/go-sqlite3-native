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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 20, 109,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 2, 2, 20, 3, 3, 5, 4, 7, 5, 9, 6, 11,
	7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16,
	31, 17, 33, 18, 35, 19, 37, 20, 3, 2, 2, 2, 108, 2, 3, 3, 2, 2, 2, 2, 5,
	3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2,
	2, 2, 37, 3, 2, 2, 2, 3, 39, 3, 2, 2, 2, 5, 46, 3, 2, 2, 2, 7, 51, 3, 2,
	2, 2, 9, 57, 3, 2, 2, 2, 11, 63, 3, 2, 2, 2, 13, 66, 3, 2, 2, 2, 15, 70,
	3, 2, 2, 2, 17, 75, 3, 2, 2, 2, 19, 81, 3, 2, 2, 2, 21, 85, 3, 2, 2, 2,
	23, 87, 3, 2, 2, 2, 25, 89, 3, 2, 2, 2, 27, 94, 3, 2, 2, 2, 29, 96, 3,
	2, 2, 2, 31, 98, 3, 2, 2, 2, 33, 103, 3, 2, 2, 2, 35, 105, 3, 2, 2, 2,
	37, 107, 3, 2, 2, 2, 39, 40, 7, 85, 2, 2, 40, 41, 7, 71, 2, 2, 41, 42,
	7, 78, 2, 2, 42, 43, 7, 71, 2, 2, 43, 44, 7, 69, 2, 2, 44, 45, 7, 86, 2,
	2, 45, 4, 3, 2, 2, 2, 46, 47, 7, 72, 2, 2, 47, 48, 7, 84, 2, 2, 48, 49,
	7, 81, 2, 2, 49, 50, 7, 79, 2, 2, 50, 6, 3, 2, 2, 2, 51, 52, 7, 89, 2,
	2, 52, 53, 7, 74, 2, 2, 53, 54, 7, 71, 2, 2, 54, 55, 7, 84, 2, 2, 55, 56,
	7, 71, 2, 2, 56, 8, 3, 2, 2, 2, 57, 58, 7, 81, 2, 2, 58, 59, 7, 84, 2,
	2, 59, 60, 7, 70, 2, 2, 60, 61, 7, 71, 2, 2, 61, 62, 7, 84, 2, 2, 62, 10,
	3, 2, 2, 2, 63, 64, 7, 68, 2, 2, 64, 65, 7, 91, 2, 2, 65, 12, 3, 2, 2,
	2, 66, 67, 7, 67, 2, 2, 67, 68, 7, 85, 2, 2, 68, 69, 7, 69, 2, 2, 69, 14,
	3, 2, 2, 2, 70, 71, 7, 70, 2, 2, 71, 72, 7, 71, 2, 2, 72, 73, 7, 85, 2,
	2, 73, 74, 7, 69, 2, 2, 74, 16, 3, 2, 2, 2, 75, 76, 7, 78, 2, 2, 76, 77,
	7, 75, 2, 2, 77, 78, 7, 79, 2, 2, 78, 79, 7, 75, 2, 2, 79, 80, 7, 86, 2,
	2, 80, 18, 3, 2, 2, 2, 81, 82, 7, 67, 2, 2, 82, 83, 7, 80, 2, 2, 83, 84,
	7, 70, 2, 2, 84, 20, 3, 2, 2, 2, 85, 86, 7, 63, 2, 2, 86, 22, 3, 2, 2,
	2, 87, 88, 7, 64, 2, 2, 88, 24, 3, 2, 2, 2, 89, 90, 7, 118, 2, 2, 90, 91,
	7, 113, 2, 2, 91, 92, 7, 102, 2, 2, 92, 93, 7, 113, 2, 2, 93, 26, 3, 2,
	2, 2, 94, 95, 7, 65, 2, 2, 95, 28, 3, 2, 2, 2, 96, 97, 7, 50, 2, 2, 97,
	30, 3, 2, 2, 2, 98, 99, 7, 118, 2, 2, 99, 100, 7, 113, 2, 2, 100, 101,
	7, 102, 2, 2, 101, 102, 7, 113, 2, 2, 102, 32, 3, 2, 2, 2, 103, 104, 7,
	46, 2, 2, 104, 34, 3, 2, 2, 2, 105, 106, 7, 42, 2, 2, 106, 36, 3, 2, 2,
	2, 107, 108, 7, 43, 2, 2, 108, 38, 3, 2, 2, 2, 3, 2, 2,
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
	"'LIMIT'", "'AND'", "'='", "'>'", "", "'?'", "'0'", "", "','", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "SELECT", "FROM", "WHERE", "ORDER", "BY", "ASC", "DESC", "LIMIT", "AND",
	"EQUAL", "GREATER", "IDENTIFIER", "PLACEHOLDER", "INT_LITERAL", "STRING_LITERAL",
	"COMMA", "LPAREN", "RPAREN",
}

var lexerRuleNames = []string{
	"SELECT", "FROM", "WHERE", "ORDER", "BY", "ASC", "DESC", "LIMIT", "AND",
	"EQUAL", "GREATER", "IDENTIFIER", "PLACEHOLDER", "INT_LITERAL", "STRING_LITERAL",
	"COMMA", "LPAREN", "RPAREN",
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
	SQLLexerSELECT         = 1
	SQLLexerFROM           = 2
	SQLLexerWHERE          = 3
	SQLLexerORDER          = 4
	SQLLexerBY             = 5
	SQLLexerASC            = 6
	SQLLexerDESC           = 7
	SQLLexerLIMIT          = 8
	SQLLexerAND            = 9
	SQLLexerEQUAL          = 10
	SQLLexerGREATER        = 11
	SQLLexerIDENTIFIER     = 12
	SQLLexerPLACEHOLDER    = 13
	SQLLexerINT_LITERAL    = 14
	SQLLexerSTRING_LITERAL = 15
	SQLLexerCOMMA          = 16
	SQLLexerLPAREN         = 17
	SQLLexerRPAREN         = 18
)

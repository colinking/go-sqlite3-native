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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 24, 152,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 6, 15, 103,
	10, 15, 13, 15, 14, 15, 104, 3, 16, 6, 16, 108, 10, 16, 13, 16, 14, 16,
	109, 3, 17, 3, 17, 3, 17, 7, 17, 115, 10, 17, 12, 17, 14, 17, 118, 11,
	17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 6, 23, 147, 10, 23, 13, 23,
	14, 23, 148, 3, 23, 3, 23, 2, 2, 24, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 3, 2, 5, 3, 2,
	50, 59, 5, 2, 67, 92, 97, 97, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2,
	156, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2,
	2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3,
	2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25,
	3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2,
	33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2,
	2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 3, 47, 3, 2, 2,
	2, 5, 54, 3, 2, 2, 2, 7, 59, 3, 2, 2, 2, 9, 65, 3, 2, 2, 2, 11, 71, 3,
	2, 2, 2, 13, 74, 3, 2, 2, 2, 15, 78, 3, 2, 2, 2, 17, 83, 3, 2, 2, 2, 19,
	89, 3, 2, 2, 2, 21, 91, 3, 2, 2, 2, 23, 93, 3, 2, 2, 2, 25, 97, 3, 2, 2,
	2, 27, 99, 3, 2, 2, 2, 29, 102, 3, 2, 2, 2, 31, 107, 3, 2, 2, 2, 33, 111,
	3, 2, 2, 2, 35, 119, 3, 2, 2, 2, 37, 121, 3, 2, 2, 2, 39, 123, 3, 2, 2,
	2, 41, 125, 3, 2, 2, 2, 43, 127, 3, 2, 2, 2, 45, 146, 3, 2, 2, 2, 47, 48,
	7, 85, 2, 2, 48, 49, 7, 71, 2, 2, 49, 50, 7, 78, 2, 2, 50, 51, 7, 71, 2,
	2, 51, 52, 7, 69, 2, 2, 52, 53, 7, 86, 2, 2, 53, 4, 3, 2, 2, 2, 54, 55,
	7, 72, 2, 2, 55, 56, 7, 84, 2, 2, 56, 57, 7, 81, 2, 2, 57, 58, 7, 79, 2,
	2, 58, 6, 3, 2, 2, 2, 59, 60, 7, 89, 2, 2, 60, 61, 7, 74, 2, 2, 61, 62,
	7, 71, 2, 2, 62, 63, 7, 84, 2, 2, 63, 64, 7, 71, 2, 2, 64, 8, 3, 2, 2,
	2, 65, 66, 7, 81, 2, 2, 66, 67, 7, 84, 2, 2, 67, 68, 7, 70, 2, 2, 68, 69,
	7, 71, 2, 2, 69, 70, 7, 84, 2, 2, 70, 10, 3, 2, 2, 2, 71, 72, 7, 68, 2,
	2, 72, 73, 7, 91, 2, 2, 73, 12, 3, 2, 2, 2, 74, 75, 7, 67, 2, 2, 75, 76,
	7, 85, 2, 2, 76, 77, 7, 69, 2, 2, 77, 14, 3, 2, 2, 2, 78, 79, 7, 70, 2,
	2, 79, 80, 7, 71, 2, 2, 80, 81, 7, 85, 2, 2, 81, 82, 7, 69, 2, 2, 82, 16,
	3, 2, 2, 2, 83, 84, 7, 78, 2, 2, 84, 85, 7, 75, 2, 2, 85, 86, 7, 79, 2,
	2, 86, 87, 7, 75, 2, 2, 87, 88, 7, 86, 2, 2, 88, 18, 3, 2, 2, 2, 89, 90,
	7, 44, 2, 2, 90, 20, 3, 2, 2, 2, 91, 92, 7, 65, 2, 2, 92, 22, 3, 2, 2,
	2, 93, 94, 7, 67, 2, 2, 94, 95, 7, 80, 2, 2, 95, 96, 7, 70, 2, 2, 96, 24,
	3, 2, 2, 2, 97, 98, 7, 63, 2, 2, 98, 26, 3, 2, 2, 2, 99, 100, 7, 64, 2,
	2, 100, 28, 3, 2, 2, 2, 101, 103, 9, 2, 2, 2, 102, 101, 3, 2, 2, 2, 103,
	104, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 30, 3,
	2, 2, 2, 106, 108, 9, 3, 2, 2, 107, 106, 3, 2, 2, 2, 108, 109, 3, 2, 2,
	2, 109, 107, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 32, 3, 2, 2, 2, 111,
	116, 5, 31, 16, 2, 112, 115, 5, 31, 16, 2, 113, 115, 5, 29, 15, 2, 114,
	112, 3, 2, 2, 2, 114, 113, 3, 2, 2, 2, 115, 118, 3, 2, 2, 2, 116, 114,
	3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 34, 3, 2, 2, 2, 118, 116, 3, 2,
	2, 2, 119, 120, 7, 46, 2, 2, 120, 36, 3, 2, 2, 2, 121, 122, 7, 42, 2, 2,
	122, 38, 3, 2, 2, 2, 123, 124, 7, 43, 2, 2, 124, 40, 3, 2, 2, 2, 125, 126,
	7, 61, 2, 2, 126, 42, 3, 2, 2, 2, 127, 128, 7, 114, 2, 2, 128, 129, 7,
	116, 2, 2, 129, 130, 7, 99, 2, 2, 130, 131, 7, 105, 2, 2, 131, 132, 7,
	111, 2, 2, 132, 133, 7, 99, 2, 2, 133, 134, 7, 97, 2, 2, 134, 135, 7, 118,
	2, 2, 135, 136, 7, 99, 2, 2, 136, 137, 7, 100, 2, 2, 137, 138, 7, 110,
	2, 2, 138, 139, 7, 103, 2, 2, 139, 140, 7, 97, 2, 2, 140, 141, 7, 107,
	2, 2, 141, 142, 7, 112, 2, 2, 142, 143, 7, 104, 2, 2, 143, 144, 7, 113,
	2, 2, 144, 44, 3, 2, 2, 2, 145, 147, 9, 4, 2, 2, 146, 145, 3, 2, 2, 2,
	147, 148, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149,
	150, 3, 2, 2, 2, 150, 151, 8, 23, 2, 2, 151, 46, 3, 2, 2, 2, 8, 2, 104,
	109, 114, 116, 148, 3, 8, 2, 2,
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
	"'LIMIT'", "'*'", "'?'", "'AND'", "'='", "'>'", "", "", "", "','", "'('",
	"')'", "';'", "'pragma_table_info'",
}

var lexerSymbolicNames = []string{
	"", "Select", "From", "Where", "Order", "By", "Asc", "Desc", "Limit", "Star",
	"Placeholder", "And", "Equal", "Greater", "Number", "Letter", "Identifier",
	"Comma", "LParen", "RParen", "Semicolon", "PragmaTableInfo", "WHITESPACE",
}

var lexerRuleNames = []string{
	"Select", "From", "Where", "Order", "By", "Asc", "Desc", "Limit", "Star",
	"Placeholder", "And", "Equal", "Greater", "Number", "Letter", "Identifier",
	"Comma", "LParen", "RParen", "Semicolon", "PragmaTableInfo", "WHITESPACE",
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
	SQLLexerSelect          = 1
	SQLLexerFrom            = 2
	SQLLexerWhere           = 3
	SQLLexerOrder           = 4
	SQLLexerBy              = 5
	SQLLexerAsc             = 6
	SQLLexerDesc            = 7
	SQLLexerLimit           = 8
	SQLLexerStar            = 9
	SQLLexerPlaceholder     = 10
	SQLLexerAnd             = 11
	SQLLexerEqual           = 12
	SQLLexerGreater         = 13
	SQLLexerNumber          = 14
	SQLLexerLetter          = 15
	SQLLexerIdentifier      = 16
	SQLLexerComma           = 17
	SQLLexerLParen          = 18
	SQLLexerRParen          = 19
	SQLLexerSemicolon       = 20
	SQLLexerPragmaTableInfo = 21
	SQLLexerWHITESPACE      = 22
)

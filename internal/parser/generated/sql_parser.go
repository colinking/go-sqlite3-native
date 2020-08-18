// Code generated from SQL.g4 by ANTLR 4.8. DO NOT EDIT.

package generated // SQL
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 22, 37, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 5, 5, 25, 10,
	5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 35, 10, 7, 3,
	7, 2, 2, 8, 2, 4, 6, 8, 10, 12, 2, 2, 2, 33, 2, 14, 3, 2, 2, 2, 4, 17,
	3, 2, 2, 2, 6, 19, 3, 2, 2, 2, 8, 24, 3, 2, 2, 2, 10, 26, 3, 2, 2, 2, 12,
	34, 3, 2, 2, 2, 14, 15, 5, 4, 3, 2, 15, 16, 7, 2, 2, 3, 16, 3, 3, 2, 2,
	2, 17, 18, 5, 6, 4, 2, 18, 5, 3, 2, 2, 2, 19, 20, 7, 3, 2, 2, 20, 21, 5,
	8, 5, 2, 21, 7, 3, 2, 2, 2, 22, 25, 7, 11, 2, 2, 23, 25, 5, 10, 6, 2, 24,
	22, 3, 2, 2, 2, 24, 23, 3, 2, 2, 2, 25, 9, 3, 2, 2, 2, 26, 27, 7, 17, 2,
	2, 27, 28, 5, 12, 7, 2, 28, 11, 3, 2, 2, 2, 29, 35, 3, 2, 2, 2, 30, 31,
	7, 17, 2, 2, 31, 35, 5, 12, 7, 2, 32, 33, 7, 16, 2, 2, 33, 35, 5, 12, 7,
	2, 34, 29, 3, 2, 2, 2, 34, 30, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 35, 13,
	3, 2, 2, 2, 4, 24, 34,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'SELECT'", "'FROM'", "'WHERE'", "'ORDER'", "'BY'", "'ASC'", "'DESC'",
	"'LIMIT'", "'*'", "'?'", "'AND'", "'='", "'>'", "", "", "','", "'('", "')'",
	"';'",
}
var symbolicNames = []string{
	"", "Select", "From", "Where", "Order", "By", "Asc", "Desc", "Limit", "Star",
	"Placeholder", "And", "Equal", "Greater", "Number", "Letter", "Comma",
	"LParen", "RParen", "Semicolon", "WHITESPACE",
}

var ruleNames = []string{
	"start", "expression", "selectExpression", "args", "identifier", "identifierEnd",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SQLParser struct {
	*antlr.BaseParser
}

func NewSQLParser(input antlr.TokenStream) *SQLParser {
	this := new(SQLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "SQL.g4"

	return this
}

// SQLParser tokens.
const (
	SQLParserEOF         = antlr.TokenEOF
	SQLParserSelect      = 1
	SQLParserFrom        = 2
	SQLParserWhere       = 3
	SQLParserOrder       = 4
	SQLParserBy          = 5
	SQLParserAsc         = 6
	SQLParserDesc        = 7
	SQLParserLimit       = 8
	SQLParserStar        = 9
	SQLParserPlaceholder = 10
	SQLParserAnd         = 11
	SQLParserEqual       = 12
	SQLParserGreater     = 13
	SQLParserNumber      = 14
	SQLParserLetter      = 15
	SQLParserComma       = 16
	SQLParserLParen      = 17
	SQLParserRParen      = 18
	SQLParserSemicolon   = 19
	SQLParserWHITESPACE  = 20
)

// SQLParser rules.
const (
	SQLParserRULE_start            = 0
	SQLParserRULE_expression       = 1
	SQLParserRULE_selectExpression = 2
	SQLParserRULE_args             = 3
	SQLParserRULE_identifier       = 4
	SQLParserRULE_identifierEnd    = 5
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(SQLParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *SQLParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SQLParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)
		p.Expression()
	}
	{
		p.SetState(13)
		p.Match(SQLParserEOF)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) SelectExpression() ISelectExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SQLParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SQLParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(15)
		p.SelectExpression()
	}

	return localctx
}

// ISelectExpressionContext is an interface to support dynamic dispatch.
type ISelectExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectExpressionContext differentiates from other interfaces.
	IsSelectExpressionContext()
}

type SelectExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectExpressionContext() *SelectExpressionContext {
	var p = new(SelectExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_selectExpression
	return p
}

func (*SelectExpressionContext) IsSelectExpressionContext() {}

func NewSelectExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectExpressionContext {
	var p = new(SelectExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_selectExpression

	return p
}

func (s *SelectExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectExpressionContext) Select() antlr.TerminalNode {
	return s.GetToken(SQLParserSelect, 0)
}

func (s *SelectExpressionContext) Args() IArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *SelectExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterSelectExpression(s)
	}
}

func (s *SelectExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitSelectExpression(s)
	}
}

func (p *SQLParser) SelectExpression() (localctx ISelectExpressionContext) {
	localctx = NewSelectExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SQLParserRULE_selectExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(17)
		p.Match(SQLParserSelect)
	}
	{
		p.SetState(18)
		p.Args()
	}

	return localctx
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_args
	return p
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) Star() antlr.TerminalNode {
	return s.GetToken(SQLParserStar, 0)
}

func (s *ArgsContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (p *SQLParser) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SQLParserRULE_args)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(22)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLParserStar:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)
			p.Match(SQLParserStar)
		}

	case SQLParserLetter:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(21)
			p.Identifier()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) Letter() antlr.TerminalNode {
	return s.GetToken(SQLParserLetter, 0)
}

func (s *IdentifierContext) IdentifierEnd() IIdentifierEndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierEndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierEndContext)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *SQLParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SQLParserRULE_identifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Match(SQLParserLetter)
	}
	{
		p.SetState(25)
		p.IdentifierEnd()
	}

	return localctx
}

// IIdentifierEndContext is an interface to support dynamic dispatch.
type IIdentifierEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierEndContext differentiates from other interfaces.
	IsIdentifierEndContext()
}

type IdentifierEndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierEndContext() *IdentifierEndContext {
	var p = new(IdentifierEndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_identifierEnd
	return p
}

func (*IdentifierEndContext) IsIdentifierEndContext() {}

func NewIdentifierEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierEndContext {
	var p = new(IdentifierEndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_identifierEnd

	return p
}

func (s *IdentifierEndContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierEndContext) Letter() antlr.TerminalNode {
	return s.GetToken(SQLParserLetter, 0)
}

func (s *IdentifierEndContext) IdentifierEnd() IIdentifierEndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierEndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierEndContext)
}

func (s *IdentifierEndContext) Number() antlr.TerminalNode {
	return s.GetToken(SQLParserNumber, 0)
}

func (s *IdentifierEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierEndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterIdentifierEnd(s)
	}
}

func (s *IdentifierEndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitIdentifierEnd(s)
	}
}

func (p *SQLParser) IdentifierEnd() (localctx IIdentifierEndContext) {
	localctx = NewIdentifierEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SQLParserRULE_identifierEnd)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(32)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLParserEOF:
		p.EnterOuterAlt(localctx, 1)

	case SQLParserLetter:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(28)
			p.Match(SQLParserLetter)
		}
		{
			p.SetState(29)
			p.IdentifierEnd()
		}

	case SQLParserNumber:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(30)
			p.Match(SQLParserNumber)
		}
		{
			p.SetState(31)
			p.IdentifierEnd()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

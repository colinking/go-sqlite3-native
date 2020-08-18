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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 24, 83, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 33, 10, 4, 3, 4, 5, 4, 36, 10,
	4, 3, 4, 5, 4, 39, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 46, 10, 5,
	3, 6, 3, 6, 5, 6, 50, 10, 6, 3, 7, 3, 7, 3, 7, 5, 7, 55, 10, 7, 3, 8, 3,
	8, 3, 8, 3, 8, 7, 8, 61, 10, 8, 12, 8, 14, 8, 64, 11, 8, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 5, 9, 72, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10,
	78, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 2, 2, 12, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 2, 5, 3, 3, 22, 22, 4, 2, 12, 12, 16, 16, 3, 2, 8, 9, 2, 81,
	2, 22, 3, 2, 2, 2, 4, 25, 3, 2, 2, 2, 6, 27, 3, 2, 2, 2, 8, 45, 3, 2, 2,
	2, 10, 49, 3, 2, 2, 2, 12, 51, 3, 2, 2, 2, 14, 56, 3, 2, 2, 2, 16, 71,
	3, 2, 2, 2, 18, 73, 3, 2, 2, 2, 20, 79, 3, 2, 2, 2, 22, 23, 5, 4, 3, 2,
	23, 24, 9, 2, 2, 2, 24, 3, 3, 2, 2, 2, 25, 26, 5, 6, 4, 2, 26, 5, 3, 2,
	2, 2, 27, 28, 7, 3, 2, 2, 28, 29, 5, 10, 6, 2, 29, 30, 7, 4, 2, 2, 30,
	32, 5, 8, 5, 2, 31, 33, 5, 14, 8, 2, 32, 31, 3, 2, 2, 2, 32, 33, 3, 2,
	2, 2, 33, 35, 3, 2, 2, 2, 34, 36, 5, 18, 10, 2, 35, 34, 3, 2, 2, 2, 35,
	36, 3, 2, 2, 2, 36, 38, 3, 2, 2, 2, 37, 39, 5, 20, 11, 2, 38, 37, 3, 2,
	2, 2, 38, 39, 3, 2, 2, 2, 39, 7, 3, 2, 2, 2, 40, 46, 7, 18, 2, 2, 41, 42,
	7, 23, 2, 2, 42, 43, 7, 20, 2, 2, 43, 44, 7, 12, 2, 2, 44, 46, 7, 21, 2,
	2, 45, 40, 3, 2, 2, 2, 45, 41, 3, 2, 2, 2, 46, 9, 3, 2, 2, 2, 47, 50, 7,
	11, 2, 2, 48, 50, 5, 12, 7, 2, 49, 47, 3, 2, 2, 2, 49, 48, 3, 2, 2, 2,
	50, 11, 3, 2, 2, 2, 51, 54, 7, 18, 2, 2, 52, 53, 7, 19, 2, 2, 53, 55, 5,
	12, 7, 2, 54, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 13, 3, 2, 2, 2, 56,
	57, 7, 5, 2, 2, 57, 62, 5, 16, 9, 2, 58, 59, 7, 13, 2, 2, 59, 61, 5, 16,
	9, 2, 60, 58, 3, 2, 2, 2, 61, 64, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63,
	3, 2, 2, 2, 63, 15, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 65, 66, 7, 18, 2, 2,
	66, 67, 7, 14, 2, 2, 67, 72, 9, 3, 2, 2, 68, 69, 7, 18, 2, 2, 69, 70, 7,
	15, 2, 2, 70, 72, 9, 3, 2, 2, 71, 65, 3, 2, 2, 2, 71, 68, 3, 2, 2, 2, 72,
	17, 3, 2, 2, 2, 73, 74, 7, 6, 2, 2, 74, 75, 7, 7, 2, 2, 75, 77, 7, 18,
	2, 2, 76, 78, 9, 4, 2, 2, 77, 76, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 19,
	3, 2, 2, 2, 79, 80, 7, 10, 2, 2, 80, 81, 7, 16, 2, 2, 81, 21, 3, 2, 2,
	2, 11, 32, 35, 38, 45, 49, 54, 62, 71, 77,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'SELECT'", "'FROM'", "'WHERE'", "'ORDER'", "'BY'", "'ASC'", "'DESC'",
	"'LIMIT'", "'*'", "'?'", "'AND'", "'='", "'>'", "", "", "", "','", "'('",
	"')'", "';'", "'pragma_table_info'",
}
var symbolicNames = []string{
	"", "Select", "From", "Where", "Order", "By", "Asc", "Desc", "Limit", "Star",
	"Placeholder", "And", "Equal", "Greater", "Number", "Letter", "Identifier",
	"Comma", "LParen", "RParen", "Semicolon", "PragmaTableInfo", "WHITESPACE",
}

var ruleNames = []string{
	"start", "expression", "selectExpression", "table", "args", "columns",
	"where", "clause", "orderBy", "limit",
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
	SQLParserEOF             = antlr.TokenEOF
	SQLParserSelect          = 1
	SQLParserFrom            = 2
	SQLParserWhere           = 3
	SQLParserOrder           = 4
	SQLParserBy              = 5
	SQLParserAsc             = 6
	SQLParserDesc            = 7
	SQLParserLimit           = 8
	SQLParserStar            = 9
	SQLParserPlaceholder     = 10
	SQLParserAnd             = 11
	SQLParserEqual           = 12
	SQLParserGreater         = 13
	SQLParserNumber          = 14
	SQLParserLetter          = 15
	SQLParserIdentifier      = 16
	SQLParserComma           = 17
	SQLParserLParen          = 18
	SQLParserRParen          = 19
	SQLParserSemicolon       = 20
	SQLParserPragmaTableInfo = 21
	SQLParserWHITESPACE      = 22
)

// SQLParser rules.
const (
	SQLParserRULE_start            = 0
	SQLParserRULE_expression       = 1
	SQLParserRULE_selectExpression = 2
	SQLParserRULE_table            = 3
	SQLParserRULE_args             = 4
	SQLParserRULE_columns          = 5
	SQLParserRULE_where            = 6
	SQLParserRULE_clause           = 7
	SQLParserRULE_orderBy          = 8
	SQLParserRULE_limit            = 9
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

func (s *StartContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(SQLParserSemicolon, 0)
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
	var _la int

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
		p.SetState(20)
		p.Expression()
	}
	{
		p.SetState(21)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SQLParserEOF || _la == SQLParserSemicolon) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
		p.SetState(23)
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

func (s *SelectExpressionContext) From() antlr.TerminalNode {
	return s.GetToken(SQLParserFrom, 0)
}

func (s *SelectExpressionContext) Table() ITableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableContext)
}

func (s *SelectExpressionContext) Where() IWhereContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereContext)
}

func (s *SelectExpressionContext) OrderBy() IOrderByContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderByContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrderByContext)
}

func (s *SelectExpressionContext) Limit() ILimitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitContext)
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
	var _la int

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
		p.SetState(25)
		p.Match(SQLParserSelect)
	}
	{
		p.SetState(26)
		p.Args()
	}
	{
		p.SetState(27)
		p.Match(SQLParserFrom)
	}
	{
		p.SetState(28)
		p.Table()
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserWhere {
		{
			p.SetState(29)
			p.Where()
		}

	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserOrder {
		{
			p.SetState(32)
			p.OrderBy()
		}

	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserLimit {
		{
			p.SetState(35)
			p.Limit()
		}

	}

	return localctx
}

// ITableContext is an interface to support dynamic dispatch.
type ITableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableContext differentiates from other interfaces.
	IsTableContext()
}

type TableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableContext() *TableContext {
	var p = new(TableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_table
	return p
}

func (*TableContext) IsTableContext() {}

func NewTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableContext {
	var p = new(TableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_table

	return p
}

func (s *TableContext) GetParser() antlr.Parser { return s.parser }

func (s *TableContext) Identifier() antlr.TerminalNode {
	return s.GetToken(SQLParserIdentifier, 0)
}

func (s *TableContext) PragmaTableInfo() antlr.TerminalNode {
	return s.GetToken(SQLParserPragmaTableInfo, 0)
}

func (s *TableContext) LParen() antlr.TerminalNode {
	return s.GetToken(SQLParserLParen, 0)
}

func (s *TableContext) Placeholder() antlr.TerminalNode {
	return s.GetToken(SQLParserPlaceholder, 0)
}

func (s *TableContext) RParen() antlr.TerminalNode {
	return s.GetToken(SQLParserRParen, 0)
}

func (s *TableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterTable(s)
	}
}

func (s *TableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitTable(s)
	}
}

func (p *SQLParser) Table() (localctx ITableContext) {
	localctx = NewTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SQLParserRULE_table)

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

	p.SetState(43)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.Match(SQLParserIdentifier)
		}

	case SQLParserPragmaTableInfo:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.Match(SQLParserPragmaTableInfo)
		}
		{
			p.SetState(40)
			p.Match(SQLParserLParen)
		}
		{
			p.SetState(41)
			p.Match(SQLParserPlaceholder)
		}
		{
			p.SetState(42)
			p.Match(SQLParserRParen)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *ArgsContext) Columns() IColumnsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnsContext)
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
	p.EnterRule(localctx, 8, SQLParserRULE_args)

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

	p.SetState(47)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLParserStar:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)
			p.Match(SQLParserStar)
		}

	case SQLParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(46)
			p.Columns()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IColumnsContext is an interface to support dynamic dispatch.
type IColumnsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnsContext differentiates from other interfaces.
	IsColumnsContext()
}

type ColumnsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnsContext() *ColumnsContext {
	var p = new(ColumnsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_columns
	return p
}

func (*ColumnsContext) IsColumnsContext() {}

func NewColumnsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnsContext {
	var p = new(ColumnsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_columns

	return p
}

func (s *ColumnsContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnsContext) Identifier() antlr.TerminalNode {
	return s.GetToken(SQLParserIdentifier, 0)
}

func (s *ColumnsContext) Comma() antlr.TerminalNode {
	return s.GetToken(SQLParserComma, 0)
}

func (s *ColumnsContext) Columns() IColumnsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnsContext)
}

func (s *ColumnsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterColumns(s)
	}
}

func (s *ColumnsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitColumns(s)
	}
}

func (p *SQLParser) Columns() (localctx IColumnsContext) {
	localctx = NewColumnsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SQLParserRULE_columns)
	var _la int

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
		p.SetState(49)
		p.Match(SQLParserIdentifier)
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserComma {
		{
			p.SetState(50)
			p.Match(SQLParserComma)
		}
		{
			p.SetState(51)
			p.Columns()
		}

	}

	return localctx
}

// IWhereContext is an interface to support dynamic dispatch.
type IWhereContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhereContext differentiates from other interfaces.
	IsWhereContext()
}

type WhereContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereContext() *WhereContext {
	var p = new(WhereContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_where
	return p
}

func (*WhereContext) IsWhereContext() {}

func NewWhereContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereContext {
	var p = new(WhereContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_where

	return p
}

func (s *WhereContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereContext) Where() antlr.TerminalNode {
	return s.GetToken(SQLParserWhere, 0)
}

func (s *WhereContext) AllClause() []IClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClauseContext)(nil)).Elem())
	var tst = make([]IClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClauseContext)
		}
	}

	return tst
}

func (s *WhereContext) Clause(i int) IClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClauseContext)
}

func (s *WhereContext) AllAnd() []antlr.TerminalNode {
	return s.GetTokens(SQLParserAnd)
}

func (s *WhereContext) And(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserAnd, i)
}

func (s *WhereContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterWhere(s)
	}
}

func (s *WhereContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitWhere(s)
	}
}

func (p *SQLParser) Where() (localctx IWhereContext) {
	localctx = NewWhereContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SQLParserRULE_where)
	var _la int

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
		p.SetState(54)
		p.Match(SQLParserWhere)
	}
	{
		p.SetState(55)
		p.Clause()
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserAnd {
		{
			p.SetState(56)
			p.Match(SQLParserAnd)
		}
		{
			p.SetState(57)
			p.Clause()
		}

		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IClauseContext is an interface to support dynamic dispatch.
type IClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClauseContext differentiates from other interfaces.
	IsClauseContext()
}

type ClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClauseContext() *ClauseContext {
	var p = new(ClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_clause
	return p
}

func (*ClauseContext) IsClauseContext() {}

func NewClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClauseContext {
	var p = new(ClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_clause

	return p
}

func (s *ClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ClauseContext) Identifier() antlr.TerminalNode {
	return s.GetToken(SQLParserIdentifier, 0)
}

func (s *ClauseContext) Equal() antlr.TerminalNode {
	return s.GetToken(SQLParserEqual, 0)
}

func (s *ClauseContext) Number() antlr.TerminalNode {
	return s.GetToken(SQLParserNumber, 0)
}

func (s *ClauseContext) Placeholder() antlr.TerminalNode {
	return s.GetToken(SQLParserPlaceholder, 0)
}

func (s *ClauseContext) Greater() antlr.TerminalNode {
	return s.GetToken(SQLParserGreater, 0)
}

func (s *ClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterClause(s)
	}
}

func (s *ClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitClause(s)
	}
}

func (p *SQLParser) Clause() (localctx IClauseContext) {
	localctx = NewClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SQLParserRULE_clause)
	var _la int

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

	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.Match(SQLParserIdentifier)
		}
		{
			p.SetState(64)
			p.Match(SQLParserEqual)
		}
		{
			p.SetState(65)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SQLParserPlaceholder || _la == SQLParserNumber) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Match(SQLParserIdentifier)
		}
		{
			p.SetState(67)
			p.Match(SQLParserGreater)
		}
		{
			p.SetState(68)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SQLParserPlaceholder || _la == SQLParserNumber) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// IOrderByContext is an interface to support dynamic dispatch.
type IOrderByContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderByContext differentiates from other interfaces.
	IsOrderByContext()
}

type OrderByContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderByContext() *OrderByContext {
	var p = new(OrderByContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_orderBy
	return p
}

func (*OrderByContext) IsOrderByContext() {}

func NewOrderByContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByContext {
	var p = new(OrderByContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_orderBy

	return p
}

func (s *OrderByContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByContext) Order() antlr.TerminalNode {
	return s.GetToken(SQLParserOrder, 0)
}

func (s *OrderByContext) By() antlr.TerminalNode {
	return s.GetToken(SQLParserBy, 0)
}

func (s *OrderByContext) Identifier() antlr.TerminalNode {
	return s.GetToken(SQLParserIdentifier, 0)
}

func (s *OrderByContext) Asc() antlr.TerminalNode {
	return s.GetToken(SQLParserAsc, 0)
}

func (s *OrderByContext) Desc() antlr.TerminalNode {
	return s.GetToken(SQLParserDesc, 0)
}

func (s *OrderByContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderByContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterOrderBy(s)
	}
}

func (s *OrderByContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitOrderBy(s)
	}
}

func (p *SQLParser) OrderBy() (localctx IOrderByContext) {
	localctx = NewOrderByContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SQLParserRULE_orderBy)
	var _la int

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
		p.SetState(71)
		p.Match(SQLParserOrder)
	}
	{
		p.SetState(72)
		p.Match(SQLParserBy)
	}
	{
		p.SetState(73)
		p.Match(SQLParserIdentifier)
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserAsc || _la == SQLParserDesc {
		{
			p.SetState(74)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SQLParserAsc || _la == SQLParserDesc) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// ILimitContext is an interface to support dynamic dispatch.
type ILimitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimitContext differentiates from other interfaces.
	IsLimitContext()
}

type LimitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitContext() *LimitContext {
	var p = new(LimitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLParserRULE_limit
	return p
}

func (*LimitContext) IsLimitContext() {}

func NewLimitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitContext {
	var p = new(LimitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_limit

	return p
}

func (s *LimitContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitContext) Limit() antlr.TerminalNode {
	return s.GetToken(SQLParserLimit, 0)
}

func (s *LimitContext) Number() antlr.TerminalNode {
	return s.GetToken(SQLParserNumber, 0)
}

func (s *LimitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.EnterLimit(s)
	}
}

func (s *LimitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLListener); ok {
		listenerT.ExitLimit(s)
	}
}

func (p *SQLParser) Limit() (localctx ILimitContext) {
	localctx = NewLimitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SQLParserRULE_limit)

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
		p.SetState(77)
		p.Match(SQLParserLimit)
	}
	{
		p.SetState(78)
		p.Match(SQLParserNumber)
	}

	return localctx
}

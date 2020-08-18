// Code generated from SQL.g4 by ANTLR 4.8. DO NOT EDIT.

package generated // SQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSQLListener is a complete listener for a parse tree produced by SQLParser.
type BaseSQLListener struct{}

var _ SQLListener = &BaseSQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseSQLListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseSQLListener) ExitStart(ctx *StartContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSQLListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSQLListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSelectExpression is called when production selectExpression is entered.
func (s *BaseSQLListener) EnterSelectExpression(ctx *SelectExpressionContext) {}

// ExitSelectExpression is called when production selectExpression is exited.
func (s *BaseSQLListener) ExitSelectExpression(ctx *SelectExpressionContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseSQLListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseSQLListener) ExitArgs(ctx *ArgsContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSQLListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSQLListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterIdentifierEnd is called when production identifierEnd is entered.
func (s *BaseSQLListener) EnterIdentifierEnd(ctx *IdentifierEndContext) {}

// ExitIdentifierEnd is called when production identifierEnd is exited.
func (s *BaseSQLListener) ExitIdentifierEnd(ctx *IdentifierEndContext) {}

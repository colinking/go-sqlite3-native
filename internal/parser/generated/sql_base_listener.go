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

// EnterTable is called when production table is entered.
func (s *BaseSQLListener) EnterTable(ctx *TableContext) {}

// ExitTable is called when production table is exited.
func (s *BaseSQLListener) ExitTable(ctx *TableContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseSQLListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseSQLListener) ExitArgs(ctx *ArgsContext) {}

// EnterColumns is called when production columns is entered.
func (s *BaseSQLListener) EnterColumns(ctx *ColumnsContext) {}

// ExitColumns is called when production columns is exited.
func (s *BaseSQLListener) ExitColumns(ctx *ColumnsContext) {}

// EnterWhere is called when production where is entered.
func (s *BaseSQLListener) EnterWhere(ctx *WhereContext) {}

// ExitWhere is called when production where is exited.
func (s *BaseSQLListener) ExitWhere(ctx *WhereContext) {}

// EnterClause is called when production clause is entered.
func (s *BaseSQLListener) EnterClause(ctx *ClauseContext) {}

// ExitClause is called when production clause is exited.
func (s *BaseSQLListener) ExitClause(ctx *ClauseContext) {}

// EnterOrderBy is called when production orderBy is entered.
func (s *BaseSQLListener) EnterOrderBy(ctx *OrderByContext) {}

// ExitOrderBy is called when production orderBy is exited.
func (s *BaseSQLListener) ExitOrderBy(ctx *OrderByContext) {}

// EnterLimit is called when production limit is entered.
func (s *BaseSQLListener) EnterLimit(ctx *LimitContext) {}

// ExitLimit is called when production limit is exited.
func (s *BaseSQLListener) ExitLimit(ctx *LimitContext) {}

// Code generated from SQL.g4 by ANTLR 4.8. DO NOT EDIT.

package generated // SQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// SQLListener is a complete listener for a parse tree produced by SQLParser.
type SQLListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSelectExpression is called when entering the selectExpression production.
	EnterSelectExpression(c *SelectExpressionContext)

	// EnterTable is called when entering the table production.
	EnterTable(c *TableContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterColumns is called when entering the columns production.
	EnterColumns(c *ColumnsContext)

	// EnterWhere is called when entering the where production.
	EnterWhere(c *WhereContext)

	// EnterClause is called when entering the clause production.
	EnterClause(c *ClauseContext)

	// EnterOrderBy is called when entering the orderBy production.
	EnterOrderBy(c *OrderByContext)

	// EnterLimit is called when entering the limit production.
	EnterLimit(c *LimitContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSelectExpression is called when exiting the selectExpression production.
	ExitSelectExpression(c *SelectExpressionContext)

	// ExitTable is called when exiting the table production.
	ExitTable(c *TableContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitColumns is called when exiting the columns production.
	ExitColumns(c *ColumnsContext)

	// ExitWhere is called when exiting the where production.
	ExitWhere(c *WhereContext)

	// ExitClause is called when exiting the clause production.
	ExitClause(c *ClauseContext)

	// ExitOrderBy is called when exiting the orderBy production.
	ExitOrderBy(c *OrderByContext)

	// ExitLimit is called when exiting the limit production.
	ExitLimit(c *LimitContext)
}

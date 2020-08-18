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

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)
}

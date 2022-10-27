// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // listen_err

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// listen_errListener is a complete listener for a parse tree produced by listen_errParser.
type listen_errListener interface {
	antlr.ParseTreeListener

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}

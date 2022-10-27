// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // calculator

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// calculatorListener is a complete listener for a parse tree produced by calculatorParser.
type calculatorListener interface {
	antlr.ParseTreeListener

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterParens is called when entering the parens production.
	EnterParens(c *ParensContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterNum is called when entering the num production.
	EnterNum(c *NumContext)

	// EnterNegNum is called when entering the NegNum production.
	EnterNegNum(c *NegNumContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitParens is called when exiting the parens production.
	ExitParens(c *ParensContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitNum is called when exiting the num production.
	ExitNum(c *NumContext)

	// ExitNegNum is called when exiting the NegNum production.
	ExitNegNum(c *NegNumContext)
}

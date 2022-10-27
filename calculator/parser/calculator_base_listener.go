// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // calculator

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BasecalculatorListener is a complete listener for a parse tree produced by calculatorParser.
type BasecalculatorListener struct{}

var _ calculatorListener = &BasecalculatorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasecalculatorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasecalculatorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasecalculatorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasecalculatorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStat is called when production stat is entered.
func (s *BasecalculatorListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BasecalculatorListener) ExitStat(ctx *StatContext) {}

// EnterParens is called when production parens is entered.
func (s *BasecalculatorListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production parens is exited.
func (s *BasecalculatorListener) ExitParens(ctx *ParensContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BasecalculatorListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BasecalculatorListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BasecalculatorListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BasecalculatorListener) ExitAddSub(ctx *AddSubContext) {}

// EnterNum is called when production num is entered.
func (s *BasecalculatorListener) EnterNum(ctx *NumContext) {}

// ExitNum is called when production num is exited.
func (s *BasecalculatorListener) ExitNum(ctx *NumContext) {}

// EnterNegNum is called when production NegNum is entered.
func (s *BasecalculatorListener) EnterNegNum(ctx *NegNumContext) {}

// ExitNegNum is called when production NegNum is exited.
func (s *BasecalculatorListener) ExitNegNum(ctx *NegNumContext) {}

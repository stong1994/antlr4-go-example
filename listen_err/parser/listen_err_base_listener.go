// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // listen_err

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// Baselisten_errListener is a complete listener for a parse tree produced by listen_errParser.
type Baselisten_errListener struct{}

var _ listen_errListener = &Baselisten_errListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Baselisten_errListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Baselisten_errListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Baselisten_errListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Baselisten_errListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStat is called when production stat is entered.
func (s *Baselisten_errListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *Baselisten_errListener) ExitStat(ctx *StatContext) {}

// EnterAdd is called when production add is entered.
func (s *Baselisten_errListener) EnterAdd(ctx *AddContext) {}

// ExitAdd is called when production add is exited.
func (s *Baselisten_errListener) ExitAdd(ctx *AddContext) {}

// EnterNum is called when production num is entered.
func (s *Baselisten_errListener) EnterNum(ctx *NumContext) {}

// ExitNum is called when production num is exited.
func (s *Baselisten_errListener) ExitNum(ctx *NumContext) {}

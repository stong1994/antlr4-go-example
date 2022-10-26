// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // calculator

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by calculatorParser.
type calculatorVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by calculatorParser#stat.
	VisitStat(ctx *StatContext) interface{}

	// Visit a parse tree produced by calculatorParser#parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by calculatorParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by calculatorParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by calculatorParser#num.
	VisitNum(ctx *NumContext) interface{}
}

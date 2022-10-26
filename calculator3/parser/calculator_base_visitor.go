// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // calculator

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BasecalculatorVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasecalculatorVisitor) VisitStat(ctx *StatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecalculatorVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecalculatorVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecalculatorVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecalculatorVisitor) VisitNum(ctx *NumContext) interface{} {
	return v.VisitChildren(ctx)
}

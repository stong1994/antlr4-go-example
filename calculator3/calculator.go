package main

import (
	. "antlr4-go-example/calculator3/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"strconv"
)

type calculator struct {
	*BasecalculatorVisitor
}

func NewCalculator() *calculator {
	return &calculator{
		BasecalculatorVisitor: &BasecalculatorVisitor{},
	}
}

func (c *calculator) VisitStat(ctx *StatContext) interface{} {
	return c.VisitChildren(ctx.Expr())
}

func (c *calculator) VisitMulDiv(ctx *MulDivContext) interface{} {
	switch ctx.GetOp().GetText() {
	case "*":
		return c.VisitChildren(ctx.Expr(0)).(int) * c.VisitChildren(ctx.Expr(1)).(int)
	case "/":
		return c.VisitChildren(ctx.Expr(0)).(int) / c.VisitChildren(ctx.Expr(1)).(int)
	default:
		panic("unexpected op: " + ctx.GetOp().GetText())
	}
}

func (c *calculator) VisitAddSub(ctx *AddSubContext) interface{} {
	switch ctx.GetOp().GetText() {
	case "+":
		return c.VisitChildren(ctx.Expr(0)).(int) + c.VisitChildren(ctx.Expr(1)).(int)
	case "-":
		return c.VisitChildren(ctx.Expr(0)).(int) - c.VisitChildren(ctx.Expr(1)).(int)
	default:
		panic("unexpected op: " + ctx.GetOp().GetText())
	}
}

func (c *calculator) VisitNum(ctx *NumContext) interface{} {
	n, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(err)
	}
	return n
}

func (c *calculator) VisitChildren(node antlr.RuleNode) interface{} {
	return node.Accept(c)
}

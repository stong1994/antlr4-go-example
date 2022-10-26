package main

import (
	. "antlr4-go-example/calculator2/parser"
	"strconv"
)

type calcListener struct {
	*BasecalculatorListener
	result int
}

func NewCalcListener() *calcListener {
	return &calcListener{
		BasecalculatorListener: &BasecalculatorListener{},
	}
}

// ExitMulDiv is called when production MulDiv is exited.
func (c *calcListener) ExitMulDiv(ctx *MulDivContext) {
	switch ctx.GetOp().GetText() {
	case "*":
		ctx.SetValue(ctx.Expr(0).GetValue() * ctx.Expr(1).GetValue())
	case "/":
		ctx.SetValue(ctx.Expr(0).GetValue() / ctx.Expr(1).GetValue())
	default:
		panic("unexpected op: " + ctx.GetOp().GetText())
	}
}

// ExitAddSub is called when production AddSub is exited.
func (c *calcListener) ExitAddSub(ctx *AddSubContext) {
	switch ctx.GetOp().GetText() {
	case "+":
		ctx.SetValue(ctx.Expr(0).GetValue() + ctx.Expr(1).GetValue())
	case "-":
		ctx.SetValue(ctx.Expr(0).GetValue() - ctx.Expr(1).GetValue())
	default:
		panic("unexpected op: " + ctx.GetOp().GetText())
	}
}

// ExitId is called when production id is exited.
func (c *calcListener) ExitNum(ctx *NumContext) {
	n, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(err)
	}
	ctx.SetValue(n)
}

func (c *calcListener) ExitStat(ctx *StatContext) {
	c.result = ctx.Expr().GetValue()
}

func (c *calcListener) Result() int {
	return c.result
}

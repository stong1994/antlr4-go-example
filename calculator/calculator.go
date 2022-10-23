package main

import (
	. "antlr4-go-example/calculator/parser"
	"strconv"
)

type calcListener struct {
	*BasecalculatorListener
	stack []int
}

func NewCalcListener() *calcListener {
	return &calcListener{
		BasecalculatorListener: &BasecalculatorListener{},
	}
}

func (c *calcListener) push(i int) {
	c.stack = append(c.stack, i)
}

func (c *calcListener) pop() int {
	if len(c.stack) == 0 {
		panic("stack is empty, unable to pop")
	}
	rst := c.stack[len(c.stack)-1]
	c.stack = c.stack[:len(c.stack)-1]
	return rst
}

// ExitMulDiv is called when production MulDiv is exited.
func (c *calcListener) ExitMulDiv(ctx *MulDivContext) {
	right, left := c.pop(), c.pop()
	switch ctx.GetOp().GetText() {
	case "*":
		c.push(left * right)
	case "/":
		c.push(left / right)
	default:
		panic("unexpected op: " + ctx.GetOp().GetText())
	}
}

// ExitAddSub is called when production AddSub is exited.
func (c *calcListener) ExitAddSub(ctx *AddSubContext) {
	right, left := c.pop(), c.pop()
	switch ctx.GetOp().GetText() {
	case "+":
		c.push(left + right)
	case "-":
		c.push(left - right)
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
	c.push(n)
}

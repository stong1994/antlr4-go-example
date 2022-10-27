package main

import (
	. "antlr4-go-example/calculator/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"testing"
)

func TestNegativeNum(t *testing.T) {
	var tests = []struct {
		Input string
		Want  int
	}{
		{
			Input: "-1",
			Want:  -1,
		},
		{
			Input: "1-3",
			Want:  -2,
		},
		{
			Input: "1--3",
			Want:  4,
		},
		{
			Input: "1-2*3",
			Want:  -5,
		},
		{
			Input: "1-2*3",
			Want:  -5,
		},
		{
			Input: "-1+3",
			Want:  2,
		},
		{
			Input: "1---3",
			Want:  -2,
		},
		{
			Input: "-1-(2+3)",
			Want:  -6,
		},
	}
	for _, v := range tests {
		t.Run(v.Input, func(t *testing.T) {
			input := antlr.NewInputStream(v.Input)
			lexer := NewcalculatorLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
			parser := NewcalculatorParser(stream)

			calculator := NewCalcListener()
			antlr.ParseTreeWalkerDefault.Walk(calculator, parser.Stat())
			result := calculator.pop()
			if result != v.Want {
				t.Errorf("want %d got %d", v.Want, result)
			}
		})
	}
}

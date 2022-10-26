package main

import (
	. "antlr4-go-example/calculator3/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func main() {
	input := antlr.NewInputStream("2--4)")
	lexer := NewcalculatorLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewcalculatorParser(stream)

	calculator := NewCalculator()
	result := parser.Stat().Accept(calculator)
	fmt.Println(result)
}

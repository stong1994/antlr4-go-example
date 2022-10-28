package main

import (
	. "antlr4-go-example/calculator4/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func main() {
	input := antlr.NewInputStream("2-4)")
	lexer := NewcalculatorLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewcalculatorParser(stream)

	listener := BasecalculatorListener{}
	antlr.ParseTreeWalkerDefault.Walk(&listener, parser.Expr())
}

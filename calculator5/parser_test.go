package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"testing"
)

func TestParser(t *testing.T) {
	input := antlr.NewInputStream("2-4<should be -2> 100+10 <should be 110>")
	lexer := Newcal_lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := Newcal_parser(stream)

	antlr.ParseTreeWalkerDefault.Walk(&antlr.BaseParseTreeListener{}, parser.Stat())
}

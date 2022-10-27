package main

import (
	. "antlr4-go-example/listen_err/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func main() {
	input := antlr.NewInputStream("2--4")
	lexer := Newlisten_errLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := Newlisten_errParser(stream)
	errListener := &ErrListener{}
	parser.RemoveErrorListeners() // 默认会使用ConsoleErrorListener，需要移除。
	parser.AddErrorListener(errListener)
	parser.GetInterpreter().SetPredictionMode(antlr.PredictionModeLLExactAmbigDetection)
	antlr.ParseTreeWalkerDefault.Walk(&Baselisten_errListener{}, parser.Stat())

	errListener.Print()
}

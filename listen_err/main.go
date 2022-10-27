package main

import (
	. "antlr4-go-example/listen_err/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func main() {
	input := antlr.NewInputStream("2+4")
	lexer := Newlisten_errLexer(input)
	errListener := &ErrListener{}
	lexer.RemoveErrorListeners() // 默认会使用ConsoleErrorListenerINSTANCE，需要移除。
	lexer.AddErrorListener(errListener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := Newlisten_errParser(stream)

	antlr.ParseTreeWalkerDefault.Walk(&Baselisten_errListener{}, parser.Stat())

	errListener.Print()
}

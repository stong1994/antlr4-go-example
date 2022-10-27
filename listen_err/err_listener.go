package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type ErrListener struct {
	antlr.DefaultErrorListener
	errList []string
}

func (el *ErrListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int,
	msg string, e antlr.RecognitionException) {
	el.errList = append(el.errList, fmt.Sprintf("pos: %d:%d, msg: %s", line, column, msg))
}

func (el *ErrListener) Print() {
	for _, err := range el.errList {
		fmt.Println(err)
	}
}

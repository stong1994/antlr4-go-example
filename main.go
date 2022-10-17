package main

import (
	"antlr4-go-example/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func main() {
	// Setup the input
	is := antlr.NewInputStream("1 + 2 * 3")

	// Create the Lexer
	lexer := parser.NewcalcLexer(is)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}

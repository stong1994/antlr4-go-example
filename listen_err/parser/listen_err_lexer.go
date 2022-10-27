// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type listen_errLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var listen_errlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func listen_errlexerLexerInit() {
	staticData := &listen_errlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'-'",
	}
	staticData.symbolicNames = []string{
		"", "SUB", "INT", "WS",
	}
	staticData.ruleNames = []string{
		"SUB", "INT", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 3, 29, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 1,
		1, 4, 1, 11, 8, 1, 11, 1, 12, 1, 12, 1, 1, 1, 1, 4, 1, 17, 8, 1, 11, 1,
		12, 1, 18, 3, 1, 21, 8, 1, 1, 2, 4, 2, 24, 8, 2, 11, 2, 12, 2, 25, 1, 2,
		1, 2, 0, 0, 3, 1, 1, 3, 2, 5, 3, 1, 0, 2, 1, 0, 48, 57, 3, 0, 9, 10, 13,
		13, 32, 32, 32, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 1,
		7, 1, 0, 0, 0, 3, 20, 1, 0, 0, 0, 5, 23, 1, 0, 0, 0, 7, 8, 5, 45, 0, 0,
		8, 2, 1, 0, 0, 0, 9, 11, 7, 0, 0, 0, 10, 9, 1, 0, 0, 0, 11, 12, 1, 0, 0,
		0, 12, 10, 1, 0, 0, 0, 12, 13, 1, 0, 0, 0, 13, 21, 1, 0, 0, 0, 14, 16,
		5, 45, 0, 0, 15, 17, 7, 0, 0, 0, 16, 15, 1, 0, 0, 0, 17, 18, 1, 0, 0, 0,
		18, 16, 1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19, 21, 1, 0, 0, 0, 20, 10, 1,
		0, 0, 0, 20, 14, 1, 0, 0, 0, 21, 4, 1, 0, 0, 0, 22, 24, 7, 1, 0, 0, 23,
		22, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0,
		0, 26, 27, 1, 0, 0, 0, 27, 28, 6, 2, 0, 0, 28, 6, 1, 0, 0, 0, 5, 0, 12,
		18, 20, 25, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// listen_errLexerInit initializes any static state used to implement listen_errLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// Newlisten_errLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Listen_errLexerInit() {
	staticData := &listen_errlexerLexerStaticData
	staticData.once.Do(listen_errlexerLexerInit)
}

// Newlisten_errLexer produces a new lexer instance for the optional input antlr.CharStream.
func Newlisten_errLexer(input antlr.CharStream) *listen_errLexer {
	Listen_errLexerInit()
	l := new(listen_errLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &listen_errlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "listen_err.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// listen_errLexer tokens.
const (
	listen_errLexerSUB = 1
	listen_errLexerINT = 2
	listen_errLexerWS  = 3
)

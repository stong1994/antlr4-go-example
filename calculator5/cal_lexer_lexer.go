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

type cal_lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var cal_lexerLexerStaticData struct {
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

func cal_lexerLexerInit() {
	staticData := &cal_lexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE", "MARK",
	}
	staticData.literalNames = []string{
		"", "'<'", "'*'", "'/'", "'+'", "'-'", "", "", "'>'",
	}
	staticData.symbolicNames = []string{
		"", "OPEN", "MUL", "DIV", "ADD", "SUB", "INT", "WS", "CLOSE", "CONTENT",
	}
	staticData.ruleNames = []string{
		"OPEN", "MUL", "DIV", "ADD", "SUB", "INT", "WS", "CLOSE", "CONTENT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 9, 53, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7,
		3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 4,
		5, 34, 8, 5, 11, 5, 12, 5, 35, 1, 6, 4, 6, 39, 8, 6, 11, 6, 12, 6, 40,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 4, 8, 50, 8, 8, 11, 8, 12, 8,
		51, 0, 0, 9, 2, 1, 4, 2, 6, 3, 8, 4, 10, 5, 12, 6, 14, 7, 16, 8, 18, 9,
		2, 0, 1, 3, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 62, 62, 54,
		0, 2, 1, 0, 0, 0, 0, 4, 1, 0, 0, 0, 0, 6, 1, 0, 0, 0, 0, 8, 1, 0, 0, 0,
		0, 10, 1, 0, 0, 0, 0, 12, 1, 0, 0, 0, 0, 14, 1, 0, 0, 0, 1, 16, 1, 0, 0,
		0, 1, 18, 1, 0, 0, 0, 2, 20, 1, 0, 0, 0, 4, 24, 1, 0, 0, 0, 6, 26, 1, 0,
		0, 0, 8, 28, 1, 0, 0, 0, 10, 30, 1, 0, 0, 0, 12, 33, 1, 0, 0, 0, 14, 38,
		1, 0, 0, 0, 16, 44, 1, 0, 0, 0, 18, 49, 1, 0, 0, 0, 20, 21, 5, 60, 0, 0,
		21, 22, 1, 0, 0, 0, 22, 23, 6, 0, 0, 0, 23, 3, 1, 0, 0, 0, 24, 25, 5, 42,
		0, 0, 25, 5, 1, 0, 0, 0, 26, 27, 5, 47, 0, 0, 27, 7, 1, 0, 0, 0, 28, 29,
		5, 43, 0, 0, 29, 9, 1, 0, 0, 0, 30, 31, 5, 45, 0, 0, 31, 11, 1, 0, 0, 0,
		32, 34, 7, 0, 0, 0, 33, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 33, 1,
		0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 13, 1, 0, 0, 0, 37, 39, 7, 1, 0, 0, 38,
		37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0,
		0, 41, 42, 1, 0, 0, 0, 42, 43, 6, 6, 1, 0, 43, 15, 1, 0, 0, 0, 44, 45,
		5, 62, 0, 0, 45, 46, 1, 0, 0, 0, 46, 47, 6, 7, 2, 0, 47, 17, 1, 0, 0, 0,
		48, 50, 8, 2, 0, 0, 49, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 49, 1,
		0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 19, 1, 0, 0, 0, 5, 0, 1, 35, 40, 51, 3,
		2, 1, 0, 6, 0, 0, 2, 0, 0,
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

// cal_lexerInit initializes any static state used to implement cal_lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// Newcal_lexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Cal_lexerInit() {
	staticData := &cal_lexerLexerStaticData
	staticData.once.Do(cal_lexerLexerInit)
}

// Newcal_lexer produces a new lexer instance for the optional input antlr.CharStream.
func Newcal_lexer(input antlr.CharStream) *cal_lexer {
	Cal_lexerInit()
	l := new(cal_lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &cal_lexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "cal_lexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// cal_lexer tokens.
const (
	cal_lexerOPEN    = 1
	cal_lexerMUL     = 2
	cal_lexerDIV     = 3
	cal_lexerADD     = 4
	cal_lexerSUB     = 5
	cal_lexerINT     = 6
	cal_lexerWS      = 7
	cal_lexerCLOSE   = 8
	cal_lexerCONTENT = 9
)

// cal_lexerMARK is the cal_lexer mode.
const cal_lexerMARK = 1

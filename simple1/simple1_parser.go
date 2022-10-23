// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // simple1

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type simple1Parser struct {
	*antlr.BaseParser
}

var simple1ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func simple1ParserInit() {
	staticData := &simple1ParserStaticData
	staticData.literalNames = []string{
		"", "'+'",
	}
	staticData.symbolicNames = []string{
		"", "ADD", "Number", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"add",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 3, 7, 2, 0, 7, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0,
		5, 0, 2, 1, 0, 0, 0, 2, 3, 5, 2, 0, 0, 3, 4, 5, 1, 0, 0, 4, 5, 5, 2, 0,
		0, 5, 1, 1, 0, 0, 0, 0,
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

// simple1ParserInit initializes any static state used to implement simple1Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// Newsimple1Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Simple1ParserInit() {
	staticData := &simple1ParserStaticData
	staticData.once.Do(simple1ParserInit)
}

// Newsimple1Parser produces a new parser instance for the optional input antlr.TokenStream.
func Newsimple1Parser(input antlr.TokenStream) *simple1Parser {
	Simple1ParserInit()
	this := new(simple1Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &simple1ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// simple1Parser tokens.
const (
	simple1ParserEOF        = antlr.TokenEOF
	simple1ParserADD        = 1
	simple1ParserNumber     = 2
	simple1ParserWHITESPACE = 3
)

// simple1ParserRULE_add is the simple1Parser rule.
const simple1ParserRULE_add = 0

// IAddContext is an interface to support dynamic dispatch.
type IAddContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddContext differentiates from other interfaces.
	IsAddContext()
}

type AddContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddContext() *AddContext {
	var p = new(AddContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = simple1ParserRULE_add
	return p
}

func (*AddContext) IsAddContext() {}

func NewAddContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddContext {
	var p = new(AddContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = simple1ParserRULE_add

	return p
}

func (s *AddContext) GetParser() antlr.Parser { return s.parser }

func (s *AddContext) AllNumber() []antlr.TerminalNode {
	return s.GetTokens(simple1ParserNumber)
}

func (s *AddContext) Number(i int) antlr.TerminalNode {
	return s.GetToken(simple1ParserNumber, i)
}

func (s *AddContext) ADD() antlr.TerminalNode {
	return s.GetToken(simple1ParserADD, 0)
}

func (s *AddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simple1Listener); ok {
		listenerT.EnterAdd(s)
	}
}

func (s *AddContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(simple1Listener); ok {
		listenerT.ExitAdd(s)
	}
}

func (p *simple1Parser) Add() (localctx IAddContext) {
	this := p
	_ = this

	localctx = NewAddContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, simple1ParserRULE_add)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(2)
		p.Match(simple1ParserNumber)
	}
	{
		p.SetState(3)
		p.Match(simple1ParserADD)
	}
	{
		p.SetState(4)
		p.Match(simple1ParserNumber)
	}

	return localctx
}

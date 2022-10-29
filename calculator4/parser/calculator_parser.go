// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // calculator

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

type calculatorParser struct {
	*antlr.BaseParser
}

var calculatorParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calculatorParserInit() {
	staticData := &calculatorParserStaticData
	staticData.literalNames = []string{
		"", "'('", "')'", "'*'", "'/'", "'+'", "'-'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "MUL", "DIV", "ADD", "SUB", "INT", "WS",
	}
	staticData.ruleNames = []string{
		"stat", "expr",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 8, 32, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 15, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 27, 8, 1, 10, 1, 12, 1, 30, 9, 1, 1, 1,
		0, 1, 2, 2, 0, 2, 0, 2, 1, 0, 3, 4, 1, 0, 5, 6, 32, 0, 4, 1, 0, 0, 0, 2,
		14, 1, 0, 0, 0, 4, 5, 3, 2, 1, 0, 5, 1, 1, 0, 0, 0, 6, 7, 6, 1, -1, 0,
		7, 8, 5, 1, 0, 0, 8, 9, 3, 2, 1, 0, 9, 10, 5, 2, 0, 0, 10, 11, 6, 1, -1,
		0, 11, 15, 1, 0, 0, 0, 12, 13, 5, 7, 0, 0, 13, 15, 6, 1, -1, 0, 14, 6,
		1, 0, 0, 0, 14, 12, 1, 0, 0, 0, 15, 28, 1, 0, 0, 0, 16, 17, 10, 4, 0, 0,
		17, 18, 7, 0, 0, 0, 18, 19, 3, 2, 1, 5, 19, 20, 6, 1, -1, 0, 20, 27, 1,
		0, 0, 0, 21, 22, 10, 3, 0, 0, 22, 23, 7, 1, 0, 0, 23, 24, 3, 2, 1, 4, 24,
		25, 6, 1, -1, 0, 25, 27, 1, 0, 0, 0, 26, 16, 1, 0, 0, 0, 26, 21, 1, 0,
		0, 0, 27, 30, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 3,
		1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 3, 14, 26, 28,
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

// calculatorParserInit initializes any static state used to implement calculatorParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewcalculatorParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalculatorParserInit() {
	staticData := &calculatorParserStaticData
	staticData.once.Do(calculatorParserInit)
}

// NewcalculatorParser produces a new parser instance for the optional input antlr.TokenStream.
func NewcalculatorParser(input antlr.TokenStream) *calculatorParser {
	CalculatorParserInit()
	this := new(calculatorParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &calculatorParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

func handleExpr(op, left, right int) int {
	switch op {
	case calculatorParserADD:
		return left + right
	case calculatorParserSUB:
		return left - right
	case calculatorParserMUL:
		return left * right
	case calculatorParserDIV:
		return left / right
	default:
		return 0
	}
}

// calculatorParser tokens.
const (
	calculatorParserEOF  = antlr.TokenEOF
	calculatorParserT__0 = 1
	calculatorParserT__1 = 2
	calculatorParserMUL  = 3
	calculatorParserDIV  = 4
	calculatorParserADD  = 5
	calculatorParserSUB  = 6
	calculatorParserINT  = 7
	calculatorParserWS   = 8
)

// calculatorParser rules.
const (
	calculatorParserRULE_stat = 0
	calculatorParserRULE_expr = 1
)

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = calculatorParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = calculatorParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.ExitStat(s)
	}
}

func (p *calculatorParser) Stat() (localctx IStatContext) {
	this := p
	_ = this

	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, calculatorParserRULE_stat)

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
		p.SetState(4)
		p.expr(0)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value attribute.
	GetValue() int

	// SetValue sets the value attribute.
	SetValue(int)

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  int
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = calculatorParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = calculatorParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetValue() int { return s.value }

func (s *ExprContext) SetValue(v int) { s.value = v }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
	s.value = ctx.value
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParensContext struct {
	*ExprContext
	_expr IExprContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ParensContext) Get_expr() IExprContext { return s._expr }

func (s *ParensContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.ExitParens(s)
	}
}

type MulDivContext struct {
	*ExprContext
	a     IExprContext
	op    antlr.Token
	b     IExprContext
	_expr IExprContext
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetA() IExprContext { return s.a }

func (s *MulDivContext) GetB() IExprContext { return s.b }

func (s *MulDivContext) Get_expr() IExprContext { return s._expr }

func (s *MulDivContext) SetA(v IExprContext) { s.a = v }

func (s *MulDivContext) SetB(v IExprContext) { s.b = v }

func (s *MulDivContext) Set_expr(v IExprContext) { s._expr = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MulDivContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulDivContext) MUL() antlr.TerminalNode {
	return s.GetToken(calculatorParserMUL, 0)
}

func (s *MulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(calculatorParserDIV, 0)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

type AddSubContext struct {
	*ExprContext
	a     IExprContext
	op    antlr.Token
	b     IExprContext
	_expr IExprContext
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetA() IExprContext { return s.a }

func (s *AddSubContext) GetB() IExprContext { return s.b }

func (s *AddSubContext) Get_expr() IExprContext { return s._expr }

func (s *AddSubContext) SetA(v IExprContext) { s.a = v }

func (s *AddSubContext) SetB(v IExprContext) { s.b = v }

func (s *AddSubContext) Set_expr(v IExprContext) { s._expr = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AddSubContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddSubContext) ADD() antlr.TerminalNode {
	return s.GetToken(calculatorParserADD, 0)
}

func (s *AddSubContext) SUB() antlr.TerminalNode {
	return s.GetToken(calculatorParserSUB, 0)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.ExitAddSub(s)
	}
}

type NumContext struct {
	*ExprContext
	_INT antlr.Token
}

func NewNumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumContext {
	var p = new(NumContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NumContext) Get_INT() antlr.Token { return s._INT }

func (s *NumContext) Set_INT(v antlr.Token) { s._INT = v }

func (s *NumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumContext) INT() antlr.TerminalNode {
	return s.GetToken(calculatorParserINT, 0)
}

func (s *NumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.EnterNum(s)
	}
}

func (s *NumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.ExitNum(s)
	}
}

func (p *calculatorParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *calculatorParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, calculatorParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(14)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case calculatorParserT__0:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(7)
			p.Match(calculatorParserT__0)
		}
		{
			p.SetState(8)

			var _x = p.expr(0)

			localctx.(*ParensContext)._expr = _x
		}
		{
			p.SetState(9)
			p.Match(calculatorParserT__1)
		}

		localctx.(*ParensContext).value = localctx.(*ParensContext).Get_expr().GetValue()

	case calculatorParserINT:
		localctx = NewNumContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(12)

			var _m = p.Match(calculatorParserINT)

			localctx.(*NumContext)._INT = _m
		}

		localctx.(*NumContext).value = (func() int {
			if localctx.(*NumContext).Get_INT() == nil {
				return 0
			} else {
				i, _ := strconv.Atoi(localctx.(*NumContext).Get_INT().GetText())
				return i
			}
		}())
		fmt.Println("got", localctx.(*NumContext).value)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(26)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MulDivContext).a = _prevctx

				p.PushNewRecursionContext(localctx, _startState, calculatorParserRULE_expr)
				p.SetState(16)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(17)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == calculatorParserMUL || _la == calculatorParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(18)

					var _x = p.expr(5)

					localctx.(*MulDivContext).b = _x
					localctx.(*MulDivContext)._expr = _x
				}

				localctx.(*MulDivContext).value = handleExpr((func() int {
					if localctx.(*MulDivContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*MulDivContext).GetOp().GetTokenType()
					}
				}()), localctx.(*MulDivContext).GetA().GetValue(), localctx.(*MulDivContext).GetB().GetValue())
				fmt.Printf("%d %s %d = %d\n", localctx.(*MulDivContext).GetA().GetValue(), (func() string {
					if localctx.(*MulDivContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*MulDivContext).GetOp().GetText()
					}
				}()), localctx.(*MulDivContext).GetB().GetValue(), localctx.(*MulDivContext).value)

			case 2:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AddSubContext).a = _prevctx

				p.PushNewRecursionContext(localctx, _startState, calculatorParserRULE_expr)
				p.SetState(21)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(22)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == calculatorParserADD || _la == calculatorParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(23)

					var _x = p.expr(4)

					localctx.(*AddSubContext).b = _x
					localctx.(*AddSubContext)._expr = _x
				}

				localctx.(*AddSubContext).value = handleExpr((func() int {
					if localctx.(*AddSubContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*AddSubContext).GetOp().GetTokenType()
					}
				}()), localctx.(*AddSubContext).GetA().GetValue(), localctx.(*AddSubContext).GetB().GetValue())
				fmt.Printf("got %s\n", (func() string {
					if localctx.(*AddSubContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*AddSubContext).GetOp().GetText()
					}
				}()))
				fmt.Printf("calculating:\t%d %s %d = %d\n", localctx.(*AddSubContext).GetA().GetValue(), (func() string {
					if localctx.(*AddSubContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*AddSubContext).GetOp().GetText()
					}
				}()), localctx.(*AddSubContext).GetB().GetValue(), localctx.(*AddSubContext).value)

			}

		}
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

func (p *calculatorParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *calculatorParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // cal_parser

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

type cal_parser struct {
	*antlr.BaseParser
}

var cal_parserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cal_parserParserInit() {
	staticData := &cal_parserParserStaticData
	staticData.literalNames = []string{
		"", "'<'", "'*'", "'/'", "'+'", "'-'", "", "", "'>'",
	}
	staticData.symbolicNames = []string{
		"", "OPEN", "MUL", "DIV", "ADD", "SUB", "INT", "WS", "CLOSE", "CONTENT",
	}
	staticData.ruleNames = []string{
		"stat", "expr", "mark",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 9, 37, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 4, 0, 9, 8,
		0, 11, 0, 12, 0, 10, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 27, 8, 1, 10, 1, 12, 1, 30, 9, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 0, 1, 2, 3, 0, 2, 4, 0, 2, 1, 0, 2,
		3, 1, 0, 4, 5, 37, 0, 8, 1, 0, 0, 0, 2, 12, 1, 0, 0, 0, 4, 31, 1, 0, 0,
		0, 6, 9, 3, 2, 1, 0, 7, 9, 3, 4, 2, 0, 8, 6, 1, 0, 0, 0, 8, 7, 1, 0, 0,
		0, 9, 10, 1, 0, 0, 0, 10, 8, 1, 0, 0, 0, 10, 11, 1, 0, 0, 0, 11, 1, 1,
		0, 0, 0, 12, 13, 6, 1, -1, 0, 13, 14, 5, 6, 0, 0, 14, 15, 6, 1, -1, 0,
		15, 28, 1, 0, 0, 0, 16, 17, 10, 3, 0, 0, 17, 18, 7, 0, 0, 0, 18, 19, 3,
		2, 1, 4, 19, 20, 6, 1, -1, 0, 20, 27, 1, 0, 0, 0, 21, 22, 10, 2, 0, 0,
		22, 23, 7, 1, 0, 0, 23, 24, 3, 2, 1, 3, 24, 25, 6, 1, -1, 0, 25, 27, 1,
		0, 0, 0, 26, 16, 1, 0, 0, 0, 26, 21, 1, 0, 0, 0, 27, 30, 1, 0, 0, 0, 28,
		26, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 3, 1, 0, 0, 0, 30, 28, 1, 0, 0,
		0, 31, 32, 5, 1, 0, 0, 32, 33, 5, 9, 0, 0, 33, 34, 5, 8, 0, 0, 34, 35,
		6, 2, -1, 0, 35, 5, 1, 0, 0, 0, 4, 8, 10, 26, 28,
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

// cal_parserInit initializes any static state used to implement cal_parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// Newcal_parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Cal_parserInit() {
	staticData := &cal_parserParserStaticData
	staticData.once.Do(cal_parserParserInit)
}

// Newcal_parser produces a new parser instance for the optional input antlr.TokenStream.
func Newcal_parser(input antlr.TokenStream) *cal_parser {
	Cal_parserInit()
	this := new(cal_parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &cal_parserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

func handleExpr(op, left, right int) int {
	switch op {
	case cal_lexerADD:
		return left + right
	case cal_lexerSUB:
		return left - right
	case cal_lexerMUL:
		return left * right
	case cal_lexerDIV:
		return left / right
	default:
		return 0
	}
}

// cal_parser tokens.
const (
	cal_parserEOF     = antlr.TokenEOF
	cal_parserOPEN    = 1
	cal_parserMUL     = 2
	cal_parserDIV     = 3
	cal_parserADD     = 4
	cal_parserSUB     = 5
	cal_parserINT     = 6
	cal_parserWS      = 7
	cal_parserCLOSE   = 8
	cal_parserCONTENT = 9
)

// cal_parser rules.
const (
	cal_parserRULE_stat = 0
	cal_parserRULE_expr = 1
	cal_parserRULE_mark = 2
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
	p.RuleIndex = cal_parserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cal_parserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) AllExpr() []IExprContext {
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

func (s *StatContext) Expr(i int) IExprContext {
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

func (s *StatContext) AllMark() []IMarkContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMarkContext); ok {
			len++
		}
	}

	tst := make([]IMarkContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMarkContext); ok {
			tst[i] = t.(IMarkContext)
			i++
		}
	}

	return tst
}

func (s *StatContext) Mark(i int) IMarkContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMarkContext); ok {
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

	return t.(IMarkContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *cal_parser) Stat() (localctx IStatContext) {
	this := p
	_ = this

	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, cal_parserRULE_stat)
	var _la int

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
	p.SetState(8)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == cal_parserOPEN || _la == cal_parserINT {
		p.SetState(8)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case cal_parserINT:
			{
				p.SetState(6)
				p.expr(0)
			}

		case cal_parserOPEN:
			{
				p.SetState(7)
				p.Mark()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(10)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = cal_parserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cal_parserRULE_expr

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

type MulDivContext struct {
	*ExprContext
	a  IExprContext
	op antlr.Token
	b  IExprContext
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

func (s *MulDivContext) SetA(v IExprContext) { s.a = v }

func (s *MulDivContext) SetB(v IExprContext) { s.b = v }

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
	return s.GetToken(cal_parserMUL, 0)
}

func (s *MulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(cal_parserDIV, 0)
}

type AddSubContext struct {
	*ExprContext
	a  IExprContext
	op antlr.Token
	b  IExprContext
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

func (s *AddSubContext) SetA(v IExprContext) { s.a = v }

func (s *AddSubContext) SetB(v IExprContext) { s.b = v }

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
	return s.GetToken(cal_parserADD, 0)
}

func (s *AddSubContext) SUB() antlr.TerminalNode {
	return s.GetToken(cal_parserSUB, 0)
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
	return s.GetToken(cal_parserINT, 0)
}

func (p *cal_parser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *cal_parser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, cal_parserRULE_expr, _p)
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
	localctx = NewNumContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(13)

		var _m = p.Match(cal_parserINT)

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

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(26)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MulDivContext).a = _prevctx

				p.PushNewRecursionContext(localctx, _startState, cal_parserRULE_expr)
				p.SetState(16)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(17)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == cal_parserMUL || _la == cal_parserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(18)

					var _x = p.expr(4)

					localctx.(*MulDivContext).b = _x
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

				p.PushNewRecursionContext(localctx, _startState, cal_parserRULE_expr)
				p.SetState(21)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(22)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == cal_parserADD || _la == cal_parserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(23)

					var _x = p.expr(3)

					localctx.(*AddSubContext).b = _x
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
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IMarkContext is an interface to support dynamic dispatch.
type IMarkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CONTENT returns the _CONTENT token.
	Get_CONTENT() antlr.Token

	// Set_CONTENT sets the _CONTENT token.
	Set_CONTENT(antlr.Token)

	// IsMarkContext differentiates from other interfaces.
	IsMarkContext()
}

type MarkContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	_CONTENT antlr.Token
}

func NewEmptyMarkContext() *MarkContext {
	var p = new(MarkContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cal_parserRULE_mark
	return p
}

func (*MarkContext) IsMarkContext() {}

func NewMarkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MarkContext {
	var p = new(MarkContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cal_parserRULE_mark

	return p
}

func (s *MarkContext) GetParser() antlr.Parser { return s.parser }

func (s *MarkContext) Get_CONTENT() antlr.Token { return s._CONTENT }

func (s *MarkContext) Set_CONTENT(v antlr.Token) { s._CONTENT = v }

func (s *MarkContext) OPEN() antlr.TerminalNode {
	return s.GetToken(cal_parserOPEN, 0)
}

func (s *MarkContext) CONTENT() antlr.TerminalNode {
	return s.GetToken(cal_parserCONTENT, 0)
}

func (s *MarkContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(cal_parserCLOSE, 0)
}

func (s *MarkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MarkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *cal_parser) Mark() (localctx IMarkContext) {
	this := p
	_ = this

	localctx = NewMarkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, cal_parserRULE_mark)

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
		p.SetState(31)
		p.Match(cal_parserOPEN)
	}
	{
		p.SetState(32)

		var _m = p.Match(cal_parserCONTENT)

		localctx.(*MarkContext)._CONTENT = _m
	}
	{
		p.SetState(33)
		p.Match(cal_parserCLOSE)
	}
	fmt.Println("comment: ", (func() string {
		if localctx.(*MarkContext).Get_CONTENT() == nil {
			return ""
		} else {
			return localctx.(*MarkContext).Get_CONTENT().GetText()
		}
	}()))

	return localctx
}

func (p *cal_parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
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

func (p *cal_parser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

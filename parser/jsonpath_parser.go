// Code generated from ./parser/JsonPath.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // JsonPath

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type JsonPathParser struct {
	*antlr.BaseParser
}

var jsonpathParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jsonpathParserInit() {
	staticData := &jsonpathParserStaticData
	staticData.literalNames = []string{
		"", "'.'", "'[*]'", "'['", "']'", "'[?('", "')]'", "'[@.length-'", "'[:]'",
		"','", "", "", "", "", "':'", "'$.'", "'@.'", "'>'", "'>='", "'<'",
		"'<='", "'=='", "'!='", "'&&'", "'||'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "IDENTIFIER", "STRING", "INT",
		"WS", "COLON", "ROOT_IDENTIFIER", "LOCAL_IDENTIFIER", "GREATER", "GREATER_EQ",
		"LESS", "LESS_EQ", "EQ", "NEQ", "AND", "OR",
	}
	staticData.ruleNames = []string{
		"jsonpath", "root", "fieldAccess", "localFieldAccess", "arrayQualifier",
		"arrayRange", "intSequence", "queryFieldValue", "filterOperation", "queryExpr",
		"dbl",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 24, 172, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 32, 8,
		1, 11, 1, 12, 1, 33, 5, 1, 36, 8, 1, 10, 1, 12, 1, 39, 9, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 47, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 4, 3, 56, 8, 3, 11, 3, 12, 3, 57, 5, 3, 60, 8, 3, 10, 3, 12, 3,
		63, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 3, 4, 78, 8, 4, 1, 5, 1, 5, 1, 5, 4, 5, 83, 8, 5, 11, 5,
		12, 5, 84, 1, 5, 1, 5, 4, 5, 89, 8, 5, 11, 5, 12, 5, 90, 1, 5, 1, 5, 1,
		5, 1, 5, 4, 5, 97, 8, 5, 11, 5, 12, 5, 98, 1, 5, 1, 5, 1, 5, 4, 5, 104,
		8, 5, 11, 5, 12, 5, 105, 1, 5, 1, 5, 3, 5, 110, 8, 5, 1, 6, 1, 6, 1, 6,
		5, 6, 115, 8, 6, 10, 6, 12, 6, 118, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7,
		124, 8, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 134, 8,
		9, 1, 9, 1, 9, 1, 9, 4, 9, 139, 8, 9, 11, 9, 12, 9, 140, 5, 9, 143, 8,
		9, 10, 9, 12, 9, 146, 9, 9, 1, 10, 4, 10, 149, 8, 10, 11, 10, 12, 10, 150,
		1, 10, 1, 10, 4, 10, 155, 8, 10, 11, 10, 12, 10, 156, 1, 10, 1, 10, 4,
		10, 161, 8, 10, 11, 10, 12, 10, 162, 1, 10, 4, 10, 166, 8, 10, 11, 10,
		12, 10, 167, 3, 10, 170, 8, 10, 1, 10, 0, 3, 2, 6, 18, 11, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 0, 2, 1, 0, 17, 22, 1, 0, 23, 24, 190, 0, 22,
		1, 0, 0, 0, 2, 24, 1, 0, 0, 0, 4, 46, 1, 0, 0, 0, 6, 48, 1, 0, 0, 0, 8,
		77, 1, 0, 0, 0, 10, 109, 1, 0, 0, 0, 12, 111, 1, 0, 0, 0, 14, 123, 1, 0,
		0, 0, 16, 125, 1, 0, 0, 0, 18, 133, 1, 0, 0, 0, 20, 169, 1, 0, 0, 0, 22,
		23, 3, 2, 1, 0, 23, 1, 1, 0, 0, 0, 24, 25, 6, 1, -1, 0, 25, 26, 5, 15,
		0, 0, 26, 27, 3, 4, 2, 0, 27, 37, 1, 0, 0, 0, 28, 31, 10, 2, 0, 0, 29,
		30, 5, 1, 0, 0, 30, 32, 3, 4, 2, 0, 31, 29, 1, 0, 0, 0, 32, 33, 1, 0, 0,
		0, 33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 36, 1, 0, 0, 0, 35, 28,
		1, 0, 0, 0, 36, 39, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0,
		38, 3, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 40, 41, 5, 10, 0, 0, 41, 47, 3,
		8, 4, 0, 42, 43, 5, 10, 0, 0, 43, 44, 5, 1, 0, 0, 44, 47, 3, 8, 4, 0, 45,
		47, 5, 10, 0, 0, 46, 40, 1, 0, 0, 0, 46, 42, 1, 0, 0, 0, 46, 45, 1, 0,
		0, 0, 47, 5, 1, 0, 0, 0, 48, 49, 6, 3, -1, 0, 49, 50, 5, 16, 0, 0, 50,
		51, 3, 4, 2, 0, 51, 61, 1, 0, 0, 0, 52, 55, 10, 1, 0, 0, 53, 54, 5, 1,
		0, 0, 54, 56, 3, 4, 2, 0, 55, 53, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 55,
		1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 60, 1, 0, 0, 0, 59, 52, 1, 0, 0, 0,
		60, 63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 7, 1, 0,
		0, 0, 63, 61, 1, 0, 0, 0, 64, 78, 5, 2, 0, 0, 65, 66, 5, 3, 0, 0, 66, 67,
		3, 12, 6, 0, 67, 68, 5, 4, 0, 0, 68, 78, 1, 0, 0, 0, 69, 78, 3, 10, 5,
		0, 70, 71, 5, 5, 0, 0, 71, 72, 3, 16, 8, 0, 72, 73, 5, 6, 0, 0, 73, 78,
		1, 0, 0, 0, 74, 75, 5, 7, 0, 0, 75, 76, 5, 12, 0, 0, 76, 78, 5, 4, 0, 0,
		77, 64, 1, 0, 0, 0, 77, 65, 1, 0, 0, 0, 77, 69, 1, 0, 0, 0, 77, 70, 1,
		0, 0, 0, 77, 74, 1, 0, 0, 0, 78, 9, 1, 0, 0, 0, 79, 110, 5, 8, 0, 0, 80,
		82, 5, 3, 0, 0, 81, 83, 5, 12, 0, 0, 82, 81, 1, 0, 0, 0, 83, 84, 1, 0,
		0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 88,
		5, 14, 0, 0, 87, 89, 5, 12, 0, 0, 88, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0,
		0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 110,
		5, 4, 0, 0, 93, 94, 5, 3, 0, 0, 94, 96, 5, 14, 0, 0, 95, 97, 5, 12, 0,
		0, 96, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99,
		1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 110, 5, 4, 0, 0, 101, 103, 5, 3,
		0, 0, 102, 104, 5, 12, 0, 0, 103, 102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0,
		105, 103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107,
		108, 5, 14, 0, 0, 108, 110, 5, 4, 0, 0, 109, 79, 1, 0, 0, 0, 109, 80, 1,
		0, 0, 0, 109, 93, 1, 0, 0, 0, 109, 101, 1, 0, 0, 0, 110, 11, 1, 0, 0, 0,
		111, 116, 5, 12, 0, 0, 112, 113, 5, 9, 0, 0, 113, 115, 5, 12, 0, 0, 114,
		112, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116, 117,
		1, 0, 0, 0, 117, 13, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 119, 124, 3, 6,
		3, 0, 120, 124, 3, 2, 1, 0, 121, 124, 3, 20, 10, 0, 122, 124, 5, 11, 0,
		0, 123, 119, 1, 0, 0, 0, 123, 120, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123,
		122, 1, 0, 0, 0, 124, 15, 1, 0, 0, 0, 125, 126, 3, 18, 9, 0, 126, 17, 1,
		0, 0, 0, 127, 128, 6, 9, -1, 0, 128, 134, 3, 14, 7, 0, 129, 130, 3, 14,
		7, 0, 130, 131, 7, 0, 0, 0, 131, 132, 3, 14, 7, 0, 132, 134, 1, 0, 0, 0,
		133, 127, 1, 0, 0, 0, 133, 129, 1, 0, 0, 0, 134, 144, 1, 0, 0, 0, 135,
		138, 10, 3, 0, 0, 136, 137, 7, 1, 0, 0, 137, 139, 3, 18, 9, 0, 138, 136,
		1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 140, 141, 1, 0,
		0, 0, 141, 143, 1, 0, 0, 0, 142, 135, 1, 0, 0, 0, 143, 146, 1, 0, 0, 0,
		144, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 19, 1, 0, 0, 0, 146, 144,
		1, 0, 0, 0, 147, 149, 5, 12, 0, 0, 148, 147, 1, 0, 0, 0, 149, 150, 1, 0,
		0, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0,
		152, 154, 5, 1, 0, 0, 153, 155, 5, 12, 0, 0, 154, 153, 1, 0, 0, 0, 155,
		156, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 170,
		1, 0, 0, 0, 158, 160, 5, 1, 0, 0, 159, 161, 5, 12, 0, 0, 160, 159, 1, 0,
		0, 0, 161, 162, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0,
		163, 170, 1, 0, 0, 0, 164, 166, 5, 12, 0, 0, 165, 164, 1, 0, 0, 0, 166,
		167, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 170,
		1, 0, 0, 0, 169, 148, 1, 0, 0, 0, 169, 158, 1, 0, 0, 0, 169, 165, 1, 0,
		0, 0, 170, 21, 1, 0, 0, 0, 21, 33, 37, 46, 57, 61, 77, 84, 90, 98, 105,
		109, 116, 123, 133, 140, 144, 150, 156, 162, 167, 169,
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

// JsonPathParserInit initializes any static state used to implement JsonPathParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJsonPathParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JsonPathParserInit() {
	staticData := &jsonpathParserStaticData
	staticData.once.Do(jsonpathParserInit)
}

// NewJsonPathParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJsonPathParser(input antlr.TokenStream) *JsonPathParser {
	JsonPathParserInit()
	this := new(JsonPathParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &jsonpathParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "JsonPath.g4"

	return this
}

// JsonPathParser tokens.
const (
	JsonPathParserEOF              = antlr.TokenEOF
	JsonPathParserT__0             = 1
	JsonPathParserT__1             = 2
	JsonPathParserT__2             = 3
	JsonPathParserT__3             = 4
	JsonPathParserT__4             = 5
	JsonPathParserT__5             = 6
	JsonPathParserT__6             = 7
	JsonPathParserT__7             = 8
	JsonPathParserT__8             = 9
	JsonPathParserIDENTIFIER       = 10
	JsonPathParserSTRING           = 11
	JsonPathParserINT              = 12
	JsonPathParserWS               = 13
	JsonPathParserCOLON            = 14
	JsonPathParserROOT_IDENTIFIER  = 15
	JsonPathParserLOCAL_IDENTIFIER = 16
	JsonPathParserGREATER          = 17
	JsonPathParserGREATER_EQ       = 18
	JsonPathParserLESS             = 19
	JsonPathParserLESS_EQ          = 20
	JsonPathParserEQ               = 21
	JsonPathParserNEQ              = 22
	JsonPathParserAND              = 23
	JsonPathParserOR               = 24
)

// JsonPathParser rules.
const (
	JsonPathParserRULE_jsonpath         = 0
	JsonPathParserRULE_root             = 1
	JsonPathParserRULE_fieldAccess      = 2
	JsonPathParserRULE_localFieldAccess = 3
	JsonPathParserRULE_arrayQualifier   = 4
	JsonPathParserRULE_arrayRange       = 5
	JsonPathParserRULE_intSequence      = 6
	JsonPathParserRULE_queryFieldValue  = 7
	JsonPathParserRULE_filterOperation  = 8
	JsonPathParserRULE_queryExpr        = 9
	JsonPathParserRULE_dbl              = 10
)

// IJsonpathContext is an interface to support dynamic dispatch.
type IJsonpathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJsonpathContext differentiates from other interfaces.
	IsJsonpathContext()
}

type JsonpathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonpathContext() *JsonpathContext {
	var p = new(JsonpathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonPathParserRULE_jsonpath
	return p
}

func (*JsonpathContext) IsJsonpathContext() {}

func NewJsonpathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonpathContext {
	var p = new(JsonpathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonPathParserRULE_jsonpath

	return p
}

func (s *JsonpathContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonpathContext) Root() IRootContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRootContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRootContext)
}

func (s *JsonpathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonpathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonpathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterJsonpath(s)
	}
}

func (s *JsonpathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitJsonpath(s)
	}
}

func (p *JsonPathParser) Jsonpath() (localctx IJsonpathContext) {
	this := p
	_ = this

	localctx = NewJsonpathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JsonPathParserRULE_jsonpath)

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
		p.SetState(22)
		p.root(0)
	}

	return localctx
}

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonPathParserRULE_root
	return p
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonPathParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) ROOT_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(JsonPathParserROOT_IDENTIFIER, 0)
}

func (s *RootContext) AllFieldAccess() []IFieldAccessContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldAccessContext); ok {
			len++
		}
	}

	tst := make([]IFieldAccessContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldAccessContext); ok {
			tst[i] = t.(IFieldAccessContext)
			i++
		}
	}

	return tst
}

func (s *RootContext) FieldAccess(i int) IFieldAccessContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldAccessContext); ok {
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

	return t.(IFieldAccessContext)
}

func (s *RootContext) Root() IRootContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRootContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRootContext)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (p *JsonPathParser) Root() (localctx IRootContext) {
	return p.root(0)
}

func (p *JsonPathParser) root(_p int) (localctx IRootContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewRootContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRootContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, JsonPathParserRULE_root, _p)

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
	{
		p.SetState(25)
		p.Match(JsonPathParserROOT_IDENTIFIER)
	}
	{
		p.SetState(26)
		p.FieldAccess()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewRootContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, JsonPathParserRULE_root)
			p.SetState(28)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			p.SetState(31)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(29)
						p.Match(JsonPathParserT__0)
					}
					{
						p.SetState(30)
						p.FieldAccess()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(33)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
			}

		}
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IFieldAccessContext is an interface to support dynamic dispatch.
type IFieldAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldAccessContext differentiates from other interfaces.
	IsFieldAccessContext()
}

type FieldAccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldAccessContext() *FieldAccessContext {
	var p = new(FieldAccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonPathParserRULE_fieldAccess
	return p
}

func (*FieldAccessContext) IsFieldAccessContext() {}

func NewFieldAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldAccessContext {
	var p = new(FieldAccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonPathParserRULE_fieldAccess

	return p
}

func (s *FieldAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldAccessContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(JsonPathParserIDENTIFIER, 0)
}

func (s *FieldAccessContext) ArrayQualifier() IArrayQualifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayQualifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayQualifierContext)
}

func (s *FieldAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterFieldAccess(s)
	}
}

func (s *FieldAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitFieldAccess(s)
	}
}

func (p *JsonPathParser) FieldAccess() (localctx IFieldAccessContext) {
	this := p
	_ = this

	localctx = NewFieldAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JsonPathParserRULE_fieldAccess)

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

	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(40)
			p.Match(JsonPathParserIDENTIFIER)
		}
		{
			p.SetState(41)
			p.ArrayQualifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
			p.Match(JsonPathParserIDENTIFIER)
		}
		{
			p.SetState(43)
			p.Match(JsonPathParserT__0)
		}
		{
			p.SetState(44)
			p.ArrayQualifier()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(45)
			p.Match(JsonPathParserIDENTIFIER)
		}

	}

	return localctx
}

// ILocalFieldAccessContext is an interface to support dynamic dispatch.
type ILocalFieldAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalFieldAccessContext differentiates from other interfaces.
	IsLocalFieldAccessContext()
}

type LocalFieldAccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalFieldAccessContext() *LocalFieldAccessContext {
	var p = new(LocalFieldAccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonPathParserRULE_localFieldAccess
	return p
}

func (*LocalFieldAccessContext) IsLocalFieldAccessContext() {}

func NewLocalFieldAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalFieldAccessContext {
	var p = new(LocalFieldAccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonPathParserRULE_localFieldAccess

	return p
}

func (s *LocalFieldAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalFieldAccessContext) LOCAL_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(JsonPathParserLOCAL_IDENTIFIER, 0)
}

func (s *LocalFieldAccessContext) AllFieldAccess() []IFieldAccessContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldAccessContext); ok {
			len++
		}
	}

	tst := make([]IFieldAccessContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldAccessContext); ok {
			tst[i] = t.(IFieldAccessContext)
			i++
		}
	}

	return tst
}

func (s *LocalFieldAccessContext) FieldAccess(i int) IFieldAccessContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldAccessContext); ok {
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

	return t.(IFieldAccessContext)
}

func (s *LocalFieldAccessContext) LocalFieldAccess() ILocalFieldAccessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocalFieldAccessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocalFieldAccessContext)
}

func (s *LocalFieldAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalFieldAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalFieldAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterLocalFieldAccess(s)
	}
}

func (s *LocalFieldAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitLocalFieldAccess(s)
	}
}

func (p *JsonPathParser) LocalFieldAccess() (localctx ILocalFieldAccessContext) {
	return p.localFieldAccess(0)
}

func (p *JsonPathParser) localFieldAccess(_p int) (localctx ILocalFieldAccessContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewLocalFieldAccessContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILocalFieldAccessContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, JsonPathParserRULE_localFieldAccess, _p)

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
	{
		p.SetState(49)
		p.Match(JsonPathParserLOCAL_IDENTIFIER)
	}
	{
		p.SetState(50)
		p.FieldAccess()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLocalFieldAccessContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, JsonPathParserRULE_localFieldAccess)
			p.SetState(52)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			p.SetState(55)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(53)
						p.Match(JsonPathParserT__0)
					}
					{
						p.SetState(54)
						p.FieldAccess()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(57)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
			}

		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IArrayQualifierContext is an interface to support dynamic dispatch.
type IArrayQualifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayQualifierContext differentiates from other interfaces.
	IsArrayQualifierContext()
}

type ArrayQualifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayQualifierContext() *ArrayQualifierContext {
	var p = new(ArrayQualifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonPathParserRULE_arrayQualifier
	return p
}

func (*ArrayQualifierContext) IsArrayQualifierContext() {}

func NewArrayQualifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayQualifierContext {
	var p = new(ArrayQualifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonPathParserRULE_arrayQualifier

	return p
}

func (s *ArrayQualifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayQualifierContext) CopyFrom(ctx *ArrayQualifierContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ArrayQualifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayQualifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GetByQueryContext struct {
	*ArrayQualifierContext
}

func NewGetByQueryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetByQueryContext {
	var p = new(GetByQueryContext)

	p.ArrayQualifierContext = NewEmptyArrayQualifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArrayQualifierContext))

	return p
}

func (s *GetByQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetByQueryContext) FilterOperation() IFilterOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterOperationContext)
}

func (s *GetByQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterGetByQuery(s)
	}
}

func (s *GetByQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitGetByQuery(s)
	}
}

type GetAllContext struct {
	*ArrayQualifierContext
}

func NewGetAllContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetAllContext {
	var p = new(GetAllContext)

	p.ArrayQualifierContext = NewEmptyArrayQualifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArrayQualifierContext))

	return p
}

func (s *GetAllContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetAllContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterGetAll(s)
	}
}

func (s *GetAllContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitGetAll(s)
	}
}

type GetByIndexBackwardContext struct {
	*ArrayQualifierContext
}

func NewGetByIndexBackwardContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetByIndexBackwardContext {
	var p = new(GetByIndexBackwardContext)

	p.ArrayQualifierContext = NewEmptyArrayQualifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArrayQualifierContext))

	return p
}

func (s *GetByIndexBackwardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetByIndexBackwardContext) INT() antlr.TerminalNode {
	return s.GetToken(JsonPathParserINT, 0)
}

func (s *GetByIndexBackwardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterGetByIndexBackward(s)
	}
}

func (s *GetByIndexBackwardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitGetByIndexBackward(s)
	}
}

type GetByIndicesContext struct {
	*ArrayQualifierContext
}

func NewGetByIndicesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetByIndicesContext {
	var p = new(GetByIndicesContext)

	p.ArrayQualifierContext = NewEmptyArrayQualifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArrayQualifierContext))

	return p
}

func (s *GetByIndicesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetByIndicesContext) IntSequence() IIntSequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntSequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntSequenceContext)
}

func (s *GetByIndicesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterGetByIndices(s)
	}
}

func (s *GetByIndicesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitGetByIndices(s)
	}
}

type GetByRangeContext struct {
	*ArrayQualifierContext
}

func NewGetByRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetByRangeContext {
	var p = new(GetByRangeContext)

	p.ArrayQualifierContext = NewEmptyArrayQualifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArrayQualifierContext))

	return p
}

func (s *GetByRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetByRangeContext) ArrayRange() IArrayRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayRangeContext)
}

func (s *GetByRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterGetByRange(s)
	}
}

func (s *GetByRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitGetByRange(s)
	}
}

func (p *JsonPathParser) ArrayQualifier() (localctx IArrayQualifierContext) {
	this := p
	_ = this

	localctx = NewArrayQualifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JsonPathParserRULE_arrayQualifier)

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

	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewGetAllContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.Match(JsonPathParserT__1)
		}

	case 2:
		localctx = NewGetByIndicesContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Match(JsonPathParserT__2)
		}
		{
			p.SetState(66)
			p.IntSequence()
		}
		{
			p.SetState(67)
			p.Match(JsonPathParserT__3)
		}

	case 3:
		localctx = NewGetByRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(69)
			p.ArrayRange()
		}

	case 4:
		localctx = NewGetByQueryContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(70)
			p.Match(JsonPathParserT__4)
		}
		{
			p.SetState(71)
			p.FilterOperation()
		}
		{
			p.SetState(72)
			p.Match(JsonPathParserT__5)
		}

	case 5:
		localctx = NewGetByIndexBackwardContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(74)
			p.Match(JsonPathParserT__6)
		}
		{
			p.SetState(75)
			p.Match(JsonPathParserINT)
		}
		{
			p.SetState(76)
			p.Match(JsonPathParserT__3)
		}

	}

	return localctx
}

// IArrayRangeContext is an interface to support dynamic dispatch.
type IArrayRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayRangeContext differentiates from other interfaces.
	IsArrayRangeContext()
}

type ArrayRangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayRangeContext() *ArrayRangeContext {
	var p = new(ArrayRangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonPathParserRULE_arrayRange
	return p
}

func (*ArrayRangeContext) IsArrayRangeContext() {}

func NewArrayRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayRangeContext {
	var p = new(ArrayRangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonPathParserRULE_arrayRange

	return p
}

func (s *ArrayRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayRangeContext) CopyFrom(ctx *ArrayRangeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ArrayRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GivenStartEndIdxRangeContext struct {
	*ArrayRangeContext
}

func NewGivenStartEndIdxRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GivenStartEndIdxRangeContext {
	var p = new(GivenStartEndIdxRangeContext)

	p.ArrayRangeContext = NewEmptyArrayRangeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArrayRangeContext))

	return p
}

func (s *GivenStartEndIdxRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GivenStartEndIdxRangeContext) COLON() antlr.TerminalNode {
	return s.GetToken(JsonPathParserCOLON, 0)
}

func (s *GivenStartEndIdxRangeContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(JsonPathParserINT)
}

func (s *GivenStartEndIdxRangeContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(JsonPathParserINT, i)
}

func (s *GivenStartEndIdxRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterGivenStartEndIdxRange(s)
	}
}

func (s *GivenStartEndIdxRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitGivenStartEndIdxRange(s)
	}
}

type AllRangeContext struct {
	*ArrayRangeContext
}

func NewAllRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllRangeContext {
	var p = new(AllRangeContext)

	p.ArrayRangeContext = NewEmptyArrayRangeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArrayRangeContext))

	return p
}

func (s *AllRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterAllRange(s)
	}
}

func (s *AllRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitAllRange(s)
	}
}

type GivenOnlyEndIdxRangeContext struct {
	*ArrayRangeContext
}

func NewGivenOnlyEndIdxRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GivenOnlyEndIdxRangeContext {
	var p = new(GivenOnlyEndIdxRangeContext)

	p.ArrayRangeContext = NewEmptyArrayRangeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArrayRangeContext))

	return p
}

func (s *GivenOnlyEndIdxRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GivenOnlyEndIdxRangeContext) COLON() antlr.TerminalNode {
	return s.GetToken(JsonPathParserCOLON, 0)
}

func (s *GivenOnlyEndIdxRangeContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(JsonPathParserINT)
}

func (s *GivenOnlyEndIdxRangeContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(JsonPathParserINT, i)
}

func (s *GivenOnlyEndIdxRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterGivenOnlyEndIdxRange(s)
	}
}

func (s *GivenOnlyEndIdxRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitGivenOnlyEndIdxRange(s)
	}
}

type GivenOnlyStartIdxRangeContext struct {
	*ArrayRangeContext
}

func NewGivenOnlyStartIdxRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GivenOnlyStartIdxRangeContext {
	var p = new(GivenOnlyStartIdxRangeContext)

	p.ArrayRangeContext = NewEmptyArrayRangeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArrayRangeContext))

	return p
}

func (s *GivenOnlyStartIdxRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GivenOnlyStartIdxRangeContext) COLON() antlr.TerminalNode {
	return s.GetToken(JsonPathParserCOLON, 0)
}

func (s *GivenOnlyStartIdxRangeContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(JsonPathParserINT)
}

func (s *GivenOnlyStartIdxRangeContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(JsonPathParserINT, i)
}

func (s *GivenOnlyStartIdxRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterGivenOnlyStartIdxRange(s)
	}
}

func (s *GivenOnlyStartIdxRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitGivenOnlyStartIdxRange(s)
	}
}

func (p *JsonPathParser) ArrayRange() (localctx IArrayRangeContext) {
	this := p
	_ = this

	localctx = NewArrayRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JsonPathParserRULE_arrayRange)
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

	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAllRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(79)
			p.Match(JsonPathParserT__7)
		}

	case 2:
		localctx = NewGivenStartEndIdxRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(80)
			p.Match(JsonPathParserT__2)
		}
		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == JsonPathParserINT {
			{
				p.SetState(81)
				p.Match(JsonPathParserINT)
			}

			p.SetState(84)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(86)
			p.Match(JsonPathParserCOLON)
		}
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == JsonPathParserINT {
			{
				p.SetState(87)
				p.Match(JsonPathParserINT)
			}

			p.SetState(90)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(92)
			p.Match(JsonPathParserT__3)
		}

	case 3:
		localctx = NewGivenOnlyEndIdxRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(93)
			p.Match(JsonPathParserT__2)
		}
		{
			p.SetState(94)
			p.Match(JsonPathParserCOLON)
		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == JsonPathParserINT {
			{
				p.SetState(95)
				p.Match(JsonPathParserINT)
			}

			p.SetState(98)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(100)
			p.Match(JsonPathParserT__3)
		}

	case 4:
		localctx = NewGivenOnlyStartIdxRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(101)
			p.Match(JsonPathParserT__2)
		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == JsonPathParserINT {
			{
				p.SetState(102)
				p.Match(JsonPathParserINT)
			}

			p.SetState(105)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(107)
			p.Match(JsonPathParserCOLON)
		}
		{
			p.SetState(108)
			p.Match(JsonPathParserT__3)
		}

	}

	return localctx
}

// IIntSequenceContext is an interface to support dynamic dispatch.
type IIntSequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntSequenceContext differentiates from other interfaces.
	IsIntSequenceContext()
}

type IntSequenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntSequenceContext() *IntSequenceContext {
	var p = new(IntSequenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonPathParserRULE_intSequence
	return p
}

func (*IntSequenceContext) IsIntSequenceContext() {}

func NewIntSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntSequenceContext {
	var p = new(IntSequenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonPathParserRULE_intSequence

	return p
}

func (s *IntSequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *IntSequenceContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(JsonPathParserINT)
}

func (s *IntSequenceContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(JsonPathParserINT, i)
}

func (s *IntSequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntSequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntSequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterIntSequence(s)
	}
}

func (s *IntSequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitIntSequence(s)
	}
}

func (p *JsonPathParser) IntSequence() (localctx IIntSequenceContext) {
	this := p
	_ = this

	localctx = NewIntSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JsonPathParserRULE_intSequence)
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
	{
		p.SetState(111)
		p.Match(JsonPathParserINT)
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JsonPathParserT__8 {
		{
			p.SetState(112)
			p.Match(JsonPathParserT__8)
		}
		{
			p.SetState(113)
			p.Match(JsonPathParserINT)
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IQueryFieldValueContext is an interface to support dynamic dispatch.
type IQueryFieldValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryFieldValueContext differentiates from other interfaces.
	IsQueryFieldValueContext()
}

type QueryFieldValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryFieldValueContext() *QueryFieldValueContext {
	var p = new(QueryFieldValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonPathParserRULE_queryFieldValue
	return p
}

func (*QueryFieldValueContext) IsQueryFieldValueContext() {}

func NewQueryFieldValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryFieldValueContext {
	var p = new(QueryFieldValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonPathParserRULE_queryFieldValue

	return p
}

func (s *QueryFieldValueContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryFieldValueContext) CopyFrom(ctx *QueryFieldValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *QueryFieldValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryFieldValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type QueryFieldStringValueContext struct {
	*QueryFieldValueContext
}

func NewQueryFieldStringValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QueryFieldStringValueContext {
	var p = new(QueryFieldStringValueContext)

	p.QueryFieldValueContext = NewEmptyQueryFieldValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryFieldValueContext))

	return p
}

func (s *QueryFieldStringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryFieldStringValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(JsonPathParserSTRING, 0)
}

func (s *QueryFieldStringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterQueryFieldStringValue(s)
	}
}

func (s *QueryFieldStringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitQueryFieldStringValue(s)
	}
}

type QueryFieldValueLocalAccessContext struct {
	*QueryFieldValueContext
}

func NewQueryFieldValueLocalAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QueryFieldValueLocalAccessContext {
	var p = new(QueryFieldValueLocalAccessContext)

	p.QueryFieldValueContext = NewEmptyQueryFieldValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryFieldValueContext))

	return p
}

func (s *QueryFieldValueLocalAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryFieldValueLocalAccessContext) LocalFieldAccess() ILocalFieldAccessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocalFieldAccessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocalFieldAccessContext)
}

func (s *QueryFieldValueLocalAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterQueryFieldValueLocalAccess(s)
	}
}

func (s *QueryFieldValueLocalAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitQueryFieldValueLocalAccess(s)
	}
}

type QueryFieldValueRootAccessContext struct {
	*QueryFieldValueContext
}

func NewQueryFieldValueRootAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QueryFieldValueRootAccessContext {
	var p = new(QueryFieldValueRootAccessContext)

	p.QueryFieldValueContext = NewEmptyQueryFieldValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryFieldValueContext))

	return p
}

func (s *QueryFieldValueRootAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryFieldValueRootAccessContext) Root() IRootContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRootContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRootContext)
}

func (s *QueryFieldValueRootAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterQueryFieldValueRootAccess(s)
	}
}

func (s *QueryFieldValueRootAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitQueryFieldValueRootAccess(s)
	}
}

type QueryFieldDoubleValueContext struct {
	*QueryFieldValueContext
}

func NewQueryFieldDoubleValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QueryFieldDoubleValueContext {
	var p = new(QueryFieldDoubleValueContext)

	p.QueryFieldValueContext = NewEmptyQueryFieldValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryFieldValueContext))

	return p
}

func (s *QueryFieldDoubleValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryFieldDoubleValueContext) Dbl() IDblContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDblContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDblContext)
}

func (s *QueryFieldDoubleValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterQueryFieldDoubleValue(s)
	}
}

func (s *QueryFieldDoubleValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitQueryFieldDoubleValue(s)
	}
}

func (p *JsonPathParser) QueryFieldValue() (localctx IQueryFieldValueContext) {
	this := p
	_ = this

	localctx = NewQueryFieldValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JsonPathParserRULE_queryFieldValue)

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

	p.SetState(123)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonPathParserLOCAL_IDENTIFIER:
		localctx = NewQueryFieldValueLocalAccessContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.localFieldAccess(0)
		}

	case JsonPathParserROOT_IDENTIFIER:
		localctx = NewQueryFieldValueRootAccessContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.root(0)
		}

	case JsonPathParserT__0, JsonPathParserINT:
		localctx = NewQueryFieldDoubleValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(121)
			p.Dbl()
		}

	case JsonPathParserSTRING:
		localctx = NewQueryFieldStringValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(122)
			p.Match(JsonPathParserSTRING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFilterOperationContext is an interface to support dynamic dispatch.
type IFilterOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterOperationContext differentiates from other interfaces.
	IsFilterOperationContext()
}

type FilterOperationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterOperationContext() *FilterOperationContext {
	var p = new(FilterOperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonPathParserRULE_filterOperation
	return p
}

func (*FilterOperationContext) IsFilterOperationContext() {}

func NewFilterOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterOperationContext {
	var p = new(FilterOperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonPathParserRULE_filterOperation

	return p
}

func (s *FilterOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterOperationContext) QueryExpr() IQueryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQueryExprContext)
}

func (s *FilterOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterFilterOperation(s)
	}
}

func (s *FilterOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitFilterOperation(s)
	}
}

func (p *JsonPathParser) FilterOperation() (localctx IFilterOperationContext) {
	this := p
	_ = this

	localctx = NewFilterOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JsonPathParserRULE_filterOperation)

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
		p.SetState(125)
		p.queryExpr(0)
	}

	return localctx
}

// IQueryExprContext is an interface to support dynamic dispatch.
type IQueryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryExprContext differentiates from other interfaces.
	IsQueryExprContext()
}

type QueryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryExprContext() *QueryExprContext {
	var p = new(QueryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonPathParserRULE_queryExpr
	return p
}

func (*QueryExprContext) IsQueryExprContext() {}

func NewQueryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryExprContext {
	var p = new(QueryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonPathParserRULE_queryExpr

	return p
}

func (s *QueryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryExprContext) CopyFrom(ctx *QueryExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *QueryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type QueryWithLogicalOpContext struct {
	*QueryExprContext
	logical_op antlr.Token
}

func NewQueryWithLogicalOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QueryWithLogicalOpContext {
	var p = new(QueryWithLogicalOpContext)

	p.QueryExprContext = NewEmptyQueryExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryExprContext))

	return p
}

func (s *QueryWithLogicalOpContext) GetLogical_op() antlr.Token { return s.logical_op }

func (s *QueryWithLogicalOpContext) SetLogical_op(v antlr.Token) { s.logical_op = v }

func (s *QueryWithLogicalOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryWithLogicalOpContext) AllQueryExpr() []IQueryExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IQueryExprContext); ok {
			len++
		}
	}

	tst := make([]IQueryExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IQueryExprContext); ok {
			tst[i] = t.(IQueryExprContext)
			i++
		}
	}

	return tst
}

func (s *QueryWithLogicalOpContext) QueryExpr(i int) IQueryExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryExprContext); ok {
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

	return t.(IQueryExprContext)
}

func (s *QueryWithLogicalOpContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(JsonPathParserAND)
}

func (s *QueryWithLogicalOpContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(JsonPathParserAND, i)
}

func (s *QueryWithLogicalOpContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(JsonPathParserOR)
}

func (s *QueryWithLogicalOpContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(JsonPathParserOR, i)
}

func (s *QueryWithLogicalOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterQueryWithLogicalOp(s)
	}
}

func (s *QueryWithLogicalOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitQueryWithLogicalOp(s)
	}
}

type QueryExistenceContext struct {
	*QueryExprContext
}

func NewQueryExistenceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QueryExistenceContext {
	var p = new(QueryExistenceContext)

	p.QueryExprContext = NewEmptyQueryExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryExprContext))

	return p
}

func (s *QueryExistenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryExistenceContext) QueryFieldValue() IQueryFieldValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryFieldValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQueryFieldValueContext)
}

func (s *QueryExistenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterQueryExistence(s)
	}
}

func (s *QueryExistenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitQueryExistence(s)
	}
}

type QueryWithComparatorContext struct {
	*QueryExprContext
	comparator_op antlr.Token
}

func NewQueryWithComparatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QueryWithComparatorContext {
	var p = new(QueryWithComparatorContext)

	p.QueryExprContext = NewEmptyQueryExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryExprContext))

	return p
}

func (s *QueryWithComparatorContext) GetComparator_op() antlr.Token { return s.comparator_op }

func (s *QueryWithComparatorContext) SetComparator_op(v antlr.Token) { s.comparator_op = v }

func (s *QueryWithComparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryWithComparatorContext) AllQueryFieldValue() []IQueryFieldValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IQueryFieldValueContext); ok {
			len++
		}
	}

	tst := make([]IQueryFieldValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IQueryFieldValueContext); ok {
			tst[i] = t.(IQueryFieldValueContext)
			i++
		}
	}

	return tst
}

func (s *QueryWithComparatorContext) QueryFieldValue(i int) IQueryFieldValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryFieldValueContext); ok {
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

	return t.(IQueryFieldValueContext)
}

func (s *QueryWithComparatorContext) GREATER() antlr.TerminalNode {
	return s.GetToken(JsonPathParserGREATER, 0)
}

func (s *QueryWithComparatorContext) GREATER_EQ() antlr.TerminalNode {
	return s.GetToken(JsonPathParserGREATER_EQ, 0)
}

func (s *QueryWithComparatorContext) LESS() antlr.TerminalNode {
	return s.GetToken(JsonPathParserLESS, 0)
}

func (s *QueryWithComparatorContext) LESS_EQ() antlr.TerminalNode {
	return s.GetToken(JsonPathParserLESS_EQ, 0)
}

func (s *QueryWithComparatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(JsonPathParserEQ, 0)
}

func (s *QueryWithComparatorContext) NEQ() antlr.TerminalNode {
	return s.GetToken(JsonPathParserNEQ, 0)
}

func (s *QueryWithComparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterQueryWithComparator(s)
	}
}

func (s *QueryWithComparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitQueryWithComparator(s)
	}
}

func (p *JsonPathParser) QueryExpr() (localctx IQueryExprContext) {
	return p.queryExpr(0)
}

func (p *JsonPathParser) queryExpr(_p int) (localctx IQueryExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewQueryExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IQueryExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, JsonPathParserRULE_queryExpr, _p)
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
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		localctx = NewQueryExistenceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(128)
			p.QueryFieldValue()
		}

	case 2:
		localctx = NewQueryWithComparatorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(129)
			p.QueryFieldValue()
		}
		{
			p.SetState(130)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*QueryWithComparatorContext).comparator_op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonPathParserGREATER)|(1<<JsonPathParserGREATER_EQ)|(1<<JsonPathParserLESS)|(1<<JsonPathParserLESS_EQ)|(1<<JsonPathParserEQ)|(1<<JsonPathParserNEQ))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*QueryWithComparatorContext).comparator_op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(131)
			p.QueryFieldValue()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewQueryWithLogicalOpContext(p, NewQueryExprContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, JsonPathParserRULE_queryExpr)
			p.SetState(135)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			p.SetState(138)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(136)

						var _lt = p.GetTokenStream().LT(1)

						localctx.(*QueryWithLogicalOpContext).logical_op = _lt

						_la = p.GetTokenStream().LA(1)

						if !(_la == JsonPathParserAND || _la == JsonPathParserOR) {
							var _ri = p.GetErrorHandler().RecoverInline(p)

							localctx.(*QueryWithLogicalOpContext).logical_op = _ri
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}
					{
						p.SetState(137)
						p.queryExpr(0)
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(140)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
			}

		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}

	return localctx
}

// IDblContext is an interface to support dynamic dispatch.
type IDblContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDblContext differentiates from other interfaces.
	IsDblContext()
}

type DblContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDblContext() *DblContext {
	var p = new(DblContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonPathParserRULE_dbl
	return p
}

func (*DblContext) IsDblContext() {}

func NewDblContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DblContext {
	var p = new(DblContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonPathParserRULE_dbl

	return p
}

func (s *DblContext) GetParser() antlr.Parser { return s.parser }

func (s *DblContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(JsonPathParserINT)
}

func (s *DblContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(JsonPathParserINT, i)
}

func (s *DblContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DblContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DblContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.EnterDbl(s)
	}
}

func (s *DblContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonPathListener); ok {
		listenerT.ExitDbl(s)
	}
}

func (p *JsonPathParser) Dbl() (localctx IDblContext) {
	this := p
	_ = this

	localctx = NewDblContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JsonPathParserRULE_dbl)
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

	var _alt int

	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == JsonPathParserINT {
			{
				p.SetState(147)
				p.Match(JsonPathParserINT)
			}

			p.SetState(150)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(152)
			p.Match(JsonPathParserT__0)
		}
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(153)
					p.Match(JsonPathParserINT)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(156)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(158)
			p.Match(JsonPathParserT__0)
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(159)
					p.Match(JsonPathParserINT)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(162)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(164)
					p.Match(JsonPathParserINT)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(167)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
		}

	}

	return localctx
}

func (p *JsonPathParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *RootContext = nil
		if localctx != nil {
			t = localctx.(*RootContext)
		}
		return p.Root_Sempred(t, predIndex)

	case 3:
		var t *LocalFieldAccessContext = nil
		if localctx != nil {
			t = localctx.(*LocalFieldAccessContext)
		}
		return p.LocalFieldAccess_Sempred(t, predIndex)

	case 9:
		var t *QueryExprContext = nil
		if localctx != nil {
			t = localctx.(*QueryExprContext)
		}
		return p.QueryExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *JsonPathParser) Root_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JsonPathParser) LocalFieldAccess_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JsonPathParser) QueryExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

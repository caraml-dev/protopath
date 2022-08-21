// Code generated from ./parser/JsonPath.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type JsonPathLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var jsonpathlexerLexerStaticData struct {
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

func jsonpathlexerLexerInit() {
	staticData := &jsonpathlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"IDENTIFIER", "STRING", "INT", "WS", "COLON", "ROOT_IDENTIFIER", "LOCAL_IDENTIFIER",
		"GREATER", "GREATER_EQ", "LESS", "LESS_EQ", "EQ", "NEQ", "AND", "OR",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 24, 141, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 5, 9, 86, 8, 9, 10, 9, 12,
		9, 89, 9, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 5, 11, 98,
		8, 11, 10, 11, 12, 11, 101, 9, 11, 3, 11, 103, 8, 11, 1, 12, 4, 12, 106,
		8, 12, 11, 12, 12, 12, 107, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 23, 0, 0, 24, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29,
		15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47,
		24, 1, 0, 5, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95,
		97, 122, 1, 0, 49, 57, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 144,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 1, 49, 1, 0, 0, 0, 3, 51, 1, 0, 0, 0, 5, 55, 1, 0, 0, 0,
		7, 57, 1, 0, 0, 0, 9, 59, 1, 0, 0, 0, 11, 63, 1, 0, 0, 0, 13, 66, 1, 0,
		0, 0, 15, 77, 1, 0, 0, 0, 17, 81, 1, 0, 0, 0, 19, 83, 1, 0, 0, 0, 21, 90,
		1, 0, 0, 0, 23, 102, 1, 0, 0, 0, 25, 105, 1, 0, 0, 0, 27, 111, 1, 0, 0,
		0, 29, 113, 1, 0, 0, 0, 31, 116, 1, 0, 0, 0, 33, 119, 1, 0, 0, 0, 35, 121,
		1, 0, 0, 0, 37, 124, 1, 0, 0, 0, 39, 126, 1, 0, 0, 0, 41, 129, 1, 0, 0,
		0, 43, 132, 1, 0, 0, 0, 45, 135, 1, 0, 0, 0, 47, 138, 1, 0, 0, 0, 49, 50,
		5, 46, 0, 0, 50, 2, 1, 0, 0, 0, 51, 52, 5, 91, 0, 0, 52, 53, 5, 42, 0,
		0, 53, 54, 5, 93, 0, 0, 54, 4, 1, 0, 0, 0, 55, 56, 5, 91, 0, 0, 56, 6,
		1, 0, 0, 0, 57, 58, 5, 93, 0, 0, 58, 8, 1, 0, 0, 0, 59, 60, 5, 91, 0, 0,
		60, 61, 5, 63, 0, 0, 61, 62, 5, 40, 0, 0, 62, 10, 1, 0, 0, 0, 63, 64, 5,
		41, 0, 0, 64, 65, 5, 93, 0, 0, 65, 12, 1, 0, 0, 0, 66, 67, 5, 91, 0, 0,
		67, 68, 5, 64, 0, 0, 68, 69, 5, 46, 0, 0, 69, 70, 5, 108, 0, 0, 70, 71,
		5, 101, 0, 0, 71, 72, 5, 110, 0, 0, 72, 73, 5, 103, 0, 0, 73, 74, 5, 116,
		0, 0, 74, 75, 5, 104, 0, 0, 75, 76, 5, 45, 0, 0, 76, 14, 1, 0, 0, 0, 77,
		78, 5, 91, 0, 0, 78, 79, 5, 58, 0, 0, 79, 80, 5, 93, 0, 0, 80, 16, 1, 0,
		0, 0, 81, 82, 5, 44, 0, 0, 82, 18, 1, 0, 0, 0, 83, 87, 7, 0, 0, 0, 84,
		86, 7, 1, 0, 0, 85, 84, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85, 1, 0, 0,
		0, 87, 88, 1, 0, 0, 0, 88, 20, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 91,
		5, 39, 0, 0, 91, 92, 3, 19, 9, 0, 92, 93, 5, 39, 0, 0, 93, 22, 1, 0, 0,
		0, 94, 103, 5, 48, 0, 0, 95, 99, 7, 2, 0, 0, 96, 98, 7, 3, 0, 0, 97, 96,
		1, 0, 0, 0, 98, 101, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0,
		0, 100, 103, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 102, 94, 1, 0, 0, 0, 102,
		95, 1, 0, 0, 0, 103, 24, 1, 0, 0, 0, 104, 106, 7, 4, 0, 0, 105, 104, 1,
		0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0,
		0, 108, 109, 1, 0, 0, 0, 109, 110, 6, 12, 0, 0, 110, 26, 1, 0, 0, 0, 111,
		112, 5, 58, 0, 0, 112, 28, 1, 0, 0, 0, 113, 114, 5, 36, 0, 0, 114, 115,
		5, 46, 0, 0, 115, 30, 1, 0, 0, 0, 116, 117, 5, 64, 0, 0, 117, 118, 5, 46,
		0, 0, 118, 32, 1, 0, 0, 0, 119, 120, 5, 62, 0, 0, 120, 34, 1, 0, 0, 0,
		121, 122, 5, 62, 0, 0, 122, 123, 5, 61, 0, 0, 123, 36, 1, 0, 0, 0, 124,
		125, 5, 60, 0, 0, 125, 38, 1, 0, 0, 0, 126, 127, 5, 60, 0, 0, 127, 128,
		5, 61, 0, 0, 128, 40, 1, 0, 0, 0, 129, 130, 5, 61, 0, 0, 130, 131, 5, 61,
		0, 0, 131, 42, 1, 0, 0, 0, 132, 133, 5, 33, 0, 0, 133, 134, 5, 61, 0, 0,
		134, 44, 1, 0, 0, 0, 135, 136, 5, 38, 0, 0, 136, 137, 5, 38, 0, 0, 137,
		46, 1, 0, 0, 0, 138, 139, 5, 124, 0, 0, 139, 140, 5, 124, 0, 0, 140, 48,
		1, 0, 0, 0, 5, 0, 87, 99, 102, 107, 1, 6, 0, 0,
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

// JsonPathLexerInit initializes any static state used to implement JsonPathLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewJsonPathLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func JsonPathLexerInit() {
	staticData := &jsonpathlexerLexerStaticData
	staticData.once.Do(jsonpathlexerLexerInit)
}

// NewJsonPathLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewJsonPathLexer(input antlr.CharStream) *JsonPathLexer {
	JsonPathLexerInit()
	l := new(JsonPathLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &jsonpathlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "JsonPath.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JsonPathLexer tokens.
const (
	JsonPathLexerT__0             = 1
	JsonPathLexerT__1             = 2
	JsonPathLexerT__2             = 3
	JsonPathLexerT__3             = 4
	JsonPathLexerT__4             = 5
	JsonPathLexerT__5             = 6
	JsonPathLexerT__6             = 7
	JsonPathLexerT__7             = 8
	JsonPathLexerT__8             = 9
	JsonPathLexerIDENTIFIER       = 10
	JsonPathLexerSTRING           = 11
	JsonPathLexerINT              = 12
	JsonPathLexerWS               = 13
	JsonPathLexerCOLON            = 14
	JsonPathLexerROOT_IDENTIFIER  = 15
	JsonPathLexerLOCAL_IDENTIFIER = 16
	JsonPathLexerGREATER          = 17
	JsonPathLexerGREATER_EQ       = 18
	JsonPathLexerLESS             = 19
	JsonPathLexerLESS_EQ          = 20
	JsonPathLexerEQ               = 21
	JsonPathLexerNEQ              = 22
	JsonPathLexerAND              = 23
	JsonPathLexerOR               = 24
)

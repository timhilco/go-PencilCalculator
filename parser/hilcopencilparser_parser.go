// Code generated from HilcoPencilParser.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // HilcoPencilParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 51, 179,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5,
	3, 47, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 7, 5, 57,
	10, 5, 12, 5, 14, 5, 60, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 78, 10,
	8, 3, 9, 3, 9, 3, 9, 7, 9, 83, 10, 9, 12, 9, 14, 9, 86, 11, 9, 3, 9, 3,
	9, 5, 9, 90, 10, 9, 3, 10, 3, 10, 3, 10, 7, 10, 95, 10, 10, 12, 10, 14,
	10, 98, 11, 10, 3, 11, 3, 11, 3, 11, 7, 11, 103, 10, 11, 12, 11, 14, 11,
	106, 11, 11, 3, 12, 3, 12, 3, 12, 7, 12, 111, 10, 12, 12, 12, 14, 12, 114,
	11, 12, 3, 13, 3, 13, 3, 13, 5, 13, 119, 10, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14,
	134, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 140, 10, 15, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 6, 17, 149, 10, 17, 13, 17, 14,
	17, 150, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18,
	161, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 167, 10, 19, 7, 19, 169,
	10, 19, 12, 19, 14, 19, 172, 11, 19, 3, 19, 5, 19, 175, 10, 19, 3, 20,
	3, 20, 3, 20, 2, 2, 21, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
	28, 30, 32, 34, 36, 38, 2, 8, 3, 2, 17, 18, 4, 2, 34, 34, 42, 46, 3, 2,
	47, 48, 3, 2, 49, 50, 4, 2, 21, 21, 26, 26, 3, 2, 4, 8, 2, 185, 2, 40,
	3, 2, 2, 2, 4, 46, 3, 2, 2, 2, 6, 48, 3, 2, 2, 2, 8, 54, 3, 2, 2, 2, 10,
	61, 3, 2, 2, 2, 12, 66, 3, 2, 2, 2, 14, 70, 3, 2, 2, 2, 16, 89, 3, 2, 2,
	2, 18, 91, 3, 2, 2, 2, 20, 99, 3, 2, 2, 2, 22, 107, 3, 2, 2, 2, 24, 118,
	3, 2, 2, 2, 26, 133, 3, 2, 2, 2, 28, 139, 3, 2, 2, 2, 30, 141, 3, 2, 2,
	2, 32, 146, 3, 2, 2, 2, 34, 160, 3, 2, 2, 2, 36, 174, 3, 2, 2, 2, 38, 176,
	3, 2, 2, 2, 40, 41, 5, 4, 3, 2, 41, 3, 3, 2, 2, 2, 42, 47, 5, 16, 9, 2,
	43, 47, 5, 14, 8, 2, 44, 47, 5, 6, 4, 2, 45, 47, 5, 12, 7, 2, 46, 42, 3,
	2, 2, 2, 46, 43, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 45, 3, 2, 2, 2, 47,
	5, 3, 2, 2, 2, 48, 49, 7, 12, 2, 2, 49, 50, 5, 4, 3, 2, 50, 51, 7, 13,
	2, 2, 51, 52, 5, 8, 5, 2, 52, 53, 7, 14, 2, 2, 53, 7, 3, 2, 2, 2, 54, 58,
	5, 10, 6, 2, 55, 57, 5, 10, 6, 2, 56, 55, 3, 2, 2, 2, 57, 60, 3, 2, 2,
	2, 58, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 9, 3, 2, 2, 2, 60, 58, 3,
	2, 2, 2, 61, 62, 5, 26, 14, 2, 62, 63, 7, 3, 2, 2, 63, 64, 5, 4, 3, 2,
	64, 65, 7, 32, 2, 2, 65, 11, 3, 2, 2, 2, 66, 67, 7, 15, 2, 2, 67, 68, 5,
	8, 5, 2, 68, 69, 7, 16, 2, 2, 69, 13, 3, 2, 2, 2, 70, 71, 7, 9, 2, 2, 71,
	72, 5, 4, 3, 2, 72, 73, 7, 10, 2, 2, 73, 77, 5, 4, 3, 2, 74, 75, 7, 11,
	2, 2, 75, 78, 5, 4, 3, 2, 76, 78, 3, 2, 2, 2, 77, 74, 3, 2, 2, 2, 77, 76,
	3, 2, 2, 2, 78, 15, 3, 2, 2, 2, 79, 84, 5, 18, 10, 2, 80, 81, 9, 2, 2,
	2, 81, 83, 5, 18, 10, 2, 82, 80, 3, 2, 2, 2, 83, 86, 3, 2, 2, 2, 84, 82,
	3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 90, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2,
	87, 88, 7, 19, 2, 2, 88, 90, 5, 18, 10, 2, 89, 79, 3, 2, 2, 2, 89, 87,
	3, 2, 2, 2, 90, 17, 3, 2, 2, 2, 91, 96, 5, 20, 11, 2, 92, 93, 9, 3, 2,
	2, 93, 95, 5, 20, 11, 2, 94, 92, 3, 2, 2, 2, 95, 98, 3, 2, 2, 2, 96, 94,
	3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 19, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2,
	99, 104, 5, 22, 12, 2, 100, 101, 9, 4, 2, 2, 101, 103, 5, 22, 12, 2, 102,
	100, 3, 2, 2, 2, 103, 106, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 104, 105,
	3, 2, 2, 2, 105, 21, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 107, 112, 5, 24,
	13, 2, 108, 109, 9, 5, 2, 2, 109, 111, 5, 24, 13, 2, 110, 108, 3, 2, 2,
	2, 111, 114, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113,
	23, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 115, 116, 7, 48, 2, 2, 116, 119,
	5, 24, 13, 2, 117, 119, 5, 26, 14, 2, 118, 115, 3, 2, 2, 2, 118, 117, 3,
	2, 2, 2, 119, 25, 3, 2, 2, 2, 120, 121, 7, 40, 2, 2, 121, 122, 5, 4, 3,
	2, 122, 123, 7, 41, 2, 2, 123, 134, 3, 2, 2, 2, 124, 134, 5, 28, 15, 2,
	125, 134, 7, 27, 2, 2, 126, 134, 7, 20, 2, 2, 127, 134, 7, 23, 2, 2, 128,
	134, 7, 22, 2, 2, 129, 134, 7, 24, 2, 2, 130, 134, 5, 30, 16, 2, 131, 134,
	5, 32, 17, 2, 132, 134, 5, 38, 20, 2, 133, 120, 3, 2, 2, 2, 133, 124, 3,
	2, 2, 2, 133, 125, 3, 2, 2, 2, 133, 126, 3, 2, 2, 2, 133, 127, 3, 2, 2,
	2, 133, 128, 3, 2, 2, 2, 133, 129, 3, 2, 2, 2, 133, 130, 3, 2, 2, 2, 133,
	131, 3, 2, 2, 2, 133, 132, 3, 2, 2, 2, 134, 27, 3, 2, 2, 2, 135, 140, 7,
	21, 2, 2, 136, 137, 7, 26, 2, 2, 137, 138, 7, 31, 2, 2, 138, 140, 7, 21,
	2, 2, 139, 135, 3, 2, 2, 2, 139, 136, 3, 2, 2, 2, 140, 29, 3, 2, 2, 2,
	141, 142, 7, 25, 2, 2, 142, 143, 7, 40, 2, 2, 143, 144, 5, 36, 19, 2, 144,
	145, 7, 41, 2, 2, 145, 31, 3, 2, 2, 2, 146, 148, 7, 26, 2, 2, 147, 149,
	5, 34, 18, 2, 148, 147, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 148, 3,
	2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 33, 3, 2, 2, 2, 152, 153, 7, 30, 2,
	2, 153, 161, 9, 6, 2, 2, 154, 155, 7, 30, 2, 2, 155, 156, 9, 6, 2, 2, 156,
	157, 7, 40, 2, 2, 157, 158, 5, 36, 19, 2, 158, 159, 7, 41, 2, 2, 159, 161,
	3, 2, 2, 2, 160, 152, 3, 2, 2, 2, 160, 154, 3, 2, 2, 2, 161, 35, 3, 2,
	2, 2, 162, 170, 5, 4, 3, 2, 163, 166, 7, 33, 2, 2, 164, 167, 5, 4, 3, 2,
	165, 167, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 166, 165, 3, 2, 2, 2, 167,
	169, 3, 2, 2, 2, 168, 163, 3, 2, 2, 2, 169, 172, 3, 2, 2, 2, 170, 168,
	3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 175, 3, 2, 2, 2, 172, 170, 3, 2,
	2, 2, 173, 175, 3, 2, 2, 2, 174, 162, 3, 2, 2, 2, 174, 173, 3, 2, 2, 2,
	175, 37, 3, 2, 2, 2, 176, 177, 9, 7, 2, 2, 177, 39, 3, 2, 2, 2, 18, 46,
	58, 77, 84, 89, 96, 104, 112, 118, 133, 139, 150, 160, 166, 170, 174,
}
var literalNames = []string{
	"", "'::'", "'today'", "'true'", "'false'", "'nil'", "'default'", "'if'",
	"'then'", "'else'", "'case'", "'is'", "'endcase'", "'switch'", "'endswitch'",
	"'AND'", "'OR'", "'NOT'", "", "", "", "", "", "", "", "", "", "", "'.'",
	"':'", "';'", "','", "'='", "'['", "']'", "'{'", "'}'", "'..'", "'('",
	"')'", "'~='", "'<'", "'<='", "'>'", "'>='", "'+'", "'-'", "'*'", "'/'",
	"'^'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "IF", "THEN", "ELSE", "CASE", "CASE_IS", "CASE_END",
	"SWITCH", "SWITCH_END", "AND", "OR", "NOT", "INT", "ID", "STRING", "FLOAT",
	"DATE", "ATFUNCTION", "CLASSNAME", "PERCENT", "NEWLINE", "WS", "DOT", "COLON",
	"SEMI", "COMMA", "EQUALS", "LBRACKET", "RBRACKET", "CURLYLBRACKET", "CURLYRBRACKET",
	"DOTDOT", "LPAREN", "RPAREN", "NOT_EQUALS", "LT", "LTE", "GT", "GTE", "PLUS",
	"MINUS", "TIMES", "DIV", "EXPONENTIATE",
}

var ruleNames = []string{
	"prog", "statement", "caseStatement", "caseList", "caseItem", "switchStatement",
	"ifStatement", "logical", "relationalExpression", "addingExpression", "multiplyingExpression",
	"factor", "base", "name", "atFunction", "dataAccessor", "fieldName", "argList",
	"specialKeyword",
}

type HilcoPencilParserParser struct {
	*antlr.BaseParser
}

// NewHilcoPencilParserParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *HilcoPencilParserParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewHilcoPencilParserParser(input antlr.TokenStream) *HilcoPencilParserParser {
	this := new(HilcoPencilParserParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "HilcoPencilParser.g4"

	return this
}

// HilcoPencilParserParser tokens.
const (
	HilcoPencilParserParserEOF           = antlr.TokenEOF
	HilcoPencilParserParserT__0          = 1
	HilcoPencilParserParserT__1          = 2
	HilcoPencilParserParserT__2          = 3
	HilcoPencilParserParserT__3          = 4
	HilcoPencilParserParserT__4          = 5
	HilcoPencilParserParserT__5          = 6
	HilcoPencilParserParserIF            = 7
	HilcoPencilParserParserTHEN          = 8
	HilcoPencilParserParserELSE          = 9
	HilcoPencilParserParserCASE          = 10
	HilcoPencilParserParserCASE_IS       = 11
	HilcoPencilParserParserCASE_END      = 12
	HilcoPencilParserParserSWITCH        = 13
	HilcoPencilParserParserSWITCH_END    = 14
	HilcoPencilParserParserAND           = 15
	HilcoPencilParserParserOR            = 16
	HilcoPencilParserParserNOT           = 17
	HilcoPencilParserParserINT           = 18
	HilcoPencilParserParserID            = 19
	HilcoPencilParserParserSTRING        = 20
	HilcoPencilParserParserFLOAT         = 21
	HilcoPencilParserParserDATE          = 22
	HilcoPencilParserParserATFUNCTION    = 23
	HilcoPencilParserParserCLASSNAME     = 24
	HilcoPencilParserParserPERCENT       = 25
	HilcoPencilParserParserNEWLINE       = 26
	HilcoPencilParserParserWS            = 27
	HilcoPencilParserParserDOT           = 28
	HilcoPencilParserParserCOLON         = 29
	HilcoPencilParserParserSEMI          = 30
	HilcoPencilParserParserCOMMA         = 31
	HilcoPencilParserParserEQUALS        = 32
	HilcoPencilParserParserLBRACKET      = 33
	HilcoPencilParserParserRBRACKET      = 34
	HilcoPencilParserParserCURLYLBRACKET = 35
	HilcoPencilParserParserCURLYRBRACKET = 36
	HilcoPencilParserParserDOTDOT        = 37
	HilcoPencilParserParserLPAREN        = 38
	HilcoPencilParserParserRPAREN        = 39
	HilcoPencilParserParserNOT_EQUALS    = 40
	HilcoPencilParserParserLT            = 41
	HilcoPencilParserParserLTE           = 42
	HilcoPencilParserParserGT            = 43
	HilcoPencilParserParserGTE           = 44
	HilcoPencilParserParserPLUS          = 45
	HilcoPencilParserParserMINUS         = 46
	HilcoPencilParserParserTIMES         = 47
	HilcoPencilParserParserDIV           = 48
	HilcoPencilParserParserEXPONENTIATE  = 49
)

// HilcoPencilParserParser rules.
const (
	HilcoPencilParserParserRULE_prog                  = 0
	HilcoPencilParserParserRULE_statement             = 1
	HilcoPencilParserParserRULE_caseStatement         = 2
	HilcoPencilParserParserRULE_caseList              = 3
	HilcoPencilParserParserRULE_caseItem              = 4
	HilcoPencilParserParserRULE_switchStatement       = 5
	HilcoPencilParserParserRULE_ifStatement           = 6
	HilcoPencilParserParserRULE_logical               = 7
	HilcoPencilParserParserRULE_relationalExpression  = 8
	HilcoPencilParserParserRULE_addingExpression      = 9
	HilcoPencilParserParserRULE_multiplyingExpression = 10
	HilcoPencilParserParserRULE_factor                = 11
	HilcoPencilParserParserRULE_base                  = 12
	HilcoPencilParserParserRULE_name                  = 13
	HilcoPencilParserParserRULE_atFunction            = 14
	HilcoPencilParserParserRULE_dataAccessor          = 15
	HilcoPencilParserParserRULE_fieldName             = 16
	HilcoPencilParserParserRULE_argList               = 17
	HilcoPencilParserParserRULE_specialKeyword        = 18
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilParserParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilParserParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *HilcoPencilParserParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HilcoPencilParserParserRULE_prog)

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
		p.SetState(38)
		p.Statement()
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilParserParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilParserParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Logical() ILogicalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicalContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) CaseStatement() ICaseStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICaseStatementContext)
}

func (s *StatementContext) SwitchStatement() ISwitchStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISwitchStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *HilcoPencilParserParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HilcoPencilParserParserRULE_statement)

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

	p.SetState(44)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case HilcoPencilParserParserT__1, HilcoPencilParserParserT__2, HilcoPencilParserParserT__3, HilcoPencilParserParserT__4, HilcoPencilParserParserT__5, HilcoPencilParserParserNOT, HilcoPencilParserParserINT, HilcoPencilParserParserID, HilcoPencilParserParserSTRING, HilcoPencilParserParserFLOAT, HilcoPencilParserParserDATE, HilcoPencilParserParserATFUNCTION, HilcoPencilParserParserCLASSNAME, HilcoPencilParserParserPERCENT, HilcoPencilParserParserLPAREN, HilcoPencilParserParserMINUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(40)
			p.Logical()
		}

	case HilcoPencilParserParserIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(41)
			p.IfStatement()
		}

	case HilcoPencilParserParserCASE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(42)
			p.CaseStatement()
		}

	case HilcoPencilParserParserSWITCH:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(43)
			p.SwitchStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICaseStatementContext is an interface to support dynamic dispatch.
type ICaseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCaseStatementContext differentiates from other interfaces.
	IsCaseStatementContext()
}

type CaseStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseStatementContext() *CaseStatementContext {
	var p = new(CaseStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilParserParserRULE_caseStatement
	return p
}

func (*CaseStatementContext) IsCaseStatementContext() {}

func NewCaseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseStatementContext {
	var p = new(CaseStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilParserParserRULE_caseStatement

	return p
}

func (s *CaseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseStatementContext) CASE() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserCASE, 0)
}

func (s *CaseStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *CaseStatementContext) CASE_IS() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserCASE_IS, 0)
}

func (s *CaseStatementContext) CaseList() ICaseListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICaseListContext)
}

func (s *CaseStatementContext) CASE_END() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserCASE_END, 0)
}

func (s *CaseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.EnterCaseStatement(s)
	}
}

func (s *CaseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.ExitCaseStatement(s)
	}
}

func (p *HilcoPencilParserParser) CaseStatement() (localctx ICaseStatementContext) {
	localctx = NewCaseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, HilcoPencilParserParserRULE_caseStatement)

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
		p.SetState(46)
		p.Match(HilcoPencilParserParserCASE)
	}
	{
		p.SetState(47)
		p.Statement()
	}
	{
		p.SetState(48)
		p.Match(HilcoPencilParserParserCASE_IS)
	}
	{
		p.SetState(49)
		p.CaseList()
	}
	{
		p.SetState(50)
		p.Match(HilcoPencilParserParserCASE_END)
	}

	return localctx
}

// ICaseListContext is an interface to support dynamic dispatch.
type ICaseListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCaseListContext differentiates from other interfaces.
	IsCaseListContext()
}

type CaseListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseListContext() *CaseListContext {
	var p = new(CaseListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilParserParserRULE_caseList
	return p
}

func (*CaseListContext) IsCaseListContext() {}

func NewCaseListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseListContext {
	var p = new(CaseListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilParserParserRULE_caseList

	return p
}

func (s *CaseListContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseListContext) AllCaseItem() []ICaseItemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICaseItemContext)(nil)).Elem())
	var tst = make([]ICaseItemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICaseItemContext)
		}
	}

	return tst
}

func (s *CaseListContext) CaseItem(i int) ICaseItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseItemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICaseItemContext)
}

func (s *CaseListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.EnterCaseList(s)
	}
}

func (s *CaseListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.ExitCaseList(s)
	}
}

func (p *HilcoPencilParserParser) CaseList() (localctx ICaseListContext) {
	localctx = NewCaseListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, HilcoPencilParserParserRULE_caseList)
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
		p.SetState(52)
		p.CaseItem()
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<HilcoPencilParserParserT__1)|(1<<HilcoPencilParserParserT__2)|(1<<HilcoPencilParserParserT__3)|(1<<HilcoPencilParserParserT__4)|(1<<HilcoPencilParserParserT__5)|(1<<HilcoPencilParserParserINT)|(1<<HilcoPencilParserParserID)|(1<<HilcoPencilParserParserSTRING)|(1<<HilcoPencilParserParserFLOAT)|(1<<HilcoPencilParserParserDATE)|(1<<HilcoPencilParserParserATFUNCTION)|(1<<HilcoPencilParserParserCLASSNAME)|(1<<HilcoPencilParserParserPERCENT))) != 0) || _la == HilcoPencilParserParserLPAREN {
		{
			p.SetState(53)
			p.CaseItem()
		}

		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICaseItemContext is an interface to support dynamic dispatch.
type ICaseItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCaseItemContext differentiates from other interfaces.
	IsCaseItemContext()
}

type CaseItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseItemContext() *CaseItemContext {
	var p = new(CaseItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilParserParserRULE_caseItem
	return p
}

func (*CaseItemContext) IsCaseItemContext() {}

func NewCaseItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseItemContext {
	var p = new(CaseItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilParserParserRULE_caseItem

	return p
}

func (s *CaseItemContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseItemContext) Base() IBaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBaseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBaseContext)
}

func (s *CaseItemContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *CaseItemContext) SEMI() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserSEMI, 0)
}

func (s *CaseItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.EnterCaseItem(s)
	}
}

func (s *CaseItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.ExitCaseItem(s)
	}
}

func (p *HilcoPencilParserParser) CaseItem() (localctx ICaseItemContext) {
	localctx = NewCaseItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, HilcoPencilParserParserRULE_caseItem)

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
		p.SetState(59)
		p.Base()
	}
	{
		p.SetState(60)
		p.Match(HilcoPencilParserParserT__0)
	}
	{
		p.SetState(61)
		p.Statement()
	}
	{
		p.SetState(62)
		p.Match(HilcoPencilParserParserSEMI)
	}

	return localctx
}

// ISwitchStatementContext is an interface to support dynamic dispatch.
type ISwitchStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitchStatementContext differentiates from other interfaces.
	IsSwitchStatementContext()
}

type SwitchStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStatementContext() *SwitchStatementContext {
	var p = new(SwitchStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilParserParserRULE_switchStatement
	return p
}

func (*SwitchStatementContext) IsSwitchStatementContext() {}

func NewSwitchStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilParserParserRULE_switchStatement

	return p
}

func (s *SwitchStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStatementContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserSWITCH, 0)
}

func (s *SwitchStatementContext) CaseList() ICaseListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICaseListContext)
}

func (s *SwitchStatementContext) SWITCH_END() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserSWITCH_END, 0)
}

func (s *SwitchStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.EnterSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.ExitSwitchStatement(s)
	}
}

func (p *HilcoPencilParserParser) SwitchStatement() (localctx ISwitchStatementContext) {
	localctx = NewSwitchStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, HilcoPencilParserParserRULE_switchStatement)

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
		p.SetState(64)
		p.Match(HilcoPencilParserParserSWITCH)
	}
	{
		p.SetState(65)
		p.CaseList()
	}
	{
		p.SetState(66)
		p.Match(HilcoPencilParserParserSWITCH_END)
	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilParserParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilParserParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserIF, 0)
}

func (s *IfStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *IfStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfStatementContext) THEN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserTHEN, 0)
}

func (s *IfStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserELSE, 0)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (p *HilcoPencilParserParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, HilcoPencilParserParserRULE_ifStatement)

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
		p.SetState(68)
		p.Match(HilcoPencilParserParserIF)
	}
	{
		p.SetState(69)
		p.Statement()
	}
	{
		p.SetState(70)
		p.Match(HilcoPencilParserParserTHEN)
	}
	{
		p.SetState(71)
		p.Statement()
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(72)
			p.Match(HilcoPencilParserParserELSE)
		}

		{
			p.SetState(73)
			p.Statement()
		}

	case 2:

	}

	return localctx
}

// ILogicalContext is an interface to support dynamic dispatch.
type ILogicalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogicalContext differentiates from other interfaces.
	IsLogicalContext()
}

type LogicalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalContext() *LogicalContext {
	var p = new(LogicalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilParserParserRULE_logical
	return p
}

func (*LogicalContext) IsLogicalContext() {}

func NewLogicalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalContext {
	var p = new(LogicalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilParserParserRULE_logical

	return p
}

func (s *LogicalContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalContext) AllRelationalExpression() []IRelationalExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRelationalExpressionContext)(nil)).Elem())
	var tst = make([]IRelationalExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRelationalExpressionContext)
		}
	}

	return tst
}

func (s *LogicalContext) RelationalExpression(i int) IRelationalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationalExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionContext)
}

func (s *LogicalContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(HilcoPencilParserParserAND)
}

func (s *LogicalContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserAND, i)
}

func (s *LogicalContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(HilcoPencilParserParserOR)
}

func (s *LogicalContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserOR, i)
}

func (s *LogicalContext) NOT() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserNOT, 0)
}

func (s *LogicalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.EnterLogical(s)
	}
}

func (s *LogicalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.ExitLogical(s)
	}
}

func (p *HilcoPencilParserParser) Logical() (localctx ILogicalContext) {
	localctx = NewLogicalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, HilcoPencilParserParserRULE_logical)
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

	p.SetState(87)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case HilcoPencilParserParserT__1, HilcoPencilParserParserT__2, HilcoPencilParserParserT__3, HilcoPencilParserParserT__4, HilcoPencilParserParserT__5, HilcoPencilParserParserINT, HilcoPencilParserParserID, HilcoPencilParserParserSTRING, HilcoPencilParserParserFLOAT, HilcoPencilParserParserDATE, HilcoPencilParserParserATFUNCTION, HilcoPencilParserParserCLASSNAME, HilcoPencilParserParserPERCENT, HilcoPencilParserParserLPAREN, HilcoPencilParserParserMINUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(77)
			p.RelationalExpression()
		}
		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == HilcoPencilParserParserAND || _la == HilcoPencilParserParserOR {
			{
				p.SetState(78)
				_la = p.GetTokenStream().LA(1)

				if !(_la == HilcoPencilParserParserAND || _la == HilcoPencilParserParserOR) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(79)
				p.RelationalExpression()
			}

			p.SetState(84)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case HilcoPencilParserParserNOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.Match(HilcoPencilParserParserNOT)
		}
		{
			p.SetState(86)
			p.RelationalExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRelationalExpressionContext is an interface to support dynamic dispatch.
type IRelationalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationalExpressionContext differentiates from other interfaces.
	IsRelationalExpressionContext()
}

type RelationalExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalExpressionContext() *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilParserParserRULE_relationalExpression
	return p
}

func (*RelationalExpressionContext) IsRelationalExpressionContext() {}

func NewRelationalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilParserParserRULE_relationalExpression

	return p
}

func (s *RelationalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalExpressionContext) AllAddingExpression() []IAddingExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAddingExpressionContext)(nil)).Elem())
	var tst = make([]IAddingExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAddingExpressionContext)
		}
	}

	return tst
}

func (s *RelationalExpressionContext) AddingExpression(i int) IAddingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddingExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAddingExpressionContext)
}

func (s *RelationalExpressionContext) AllEQUALS() []antlr.TerminalNode {
	return s.GetTokens(HilcoPencilParserParserEQUALS)
}

func (s *RelationalExpressionContext) EQUALS(i int) antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserEQUALS, i)
}

func (s *RelationalExpressionContext) AllNOT_EQUALS() []antlr.TerminalNode {
	return s.GetTokens(HilcoPencilParserParserNOT_EQUALS)
}

func (s *RelationalExpressionContext) NOT_EQUALS(i int) antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserNOT_EQUALS, i)
}

func (s *RelationalExpressionContext) AllGT() []antlr.TerminalNode {
	return s.GetTokens(HilcoPencilParserParserGT)
}

func (s *RelationalExpressionContext) GT(i int) antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserGT, i)
}

func (s *RelationalExpressionContext) AllGTE() []antlr.TerminalNode {
	return s.GetTokens(HilcoPencilParserParserGTE)
}

func (s *RelationalExpressionContext) GTE(i int) antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserGTE, i)
}

func (s *RelationalExpressionContext) AllLT() []antlr.TerminalNode {
	return s.GetTokens(HilcoPencilParserParserLT)
}

func (s *RelationalExpressionContext) LT(i int) antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserLT, i)
}

func (s *RelationalExpressionContext) AllLTE() []antlr.TerminalNode {
	return s.GetTokens(HilcoPencilParserParserLTE)
}

func (s *RelationalExpressionContext) LTE(i int) antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserLTE, i)
}

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.EnterRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.ExitRelationalExpression(s)
	}
}

func (p *HilcoPencilParserParser) RelationalExpression() (localctx IRelationalExpressionContext) {
	localctx = NewRelationalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, HilcoPencilParserParserRULE_relationalExpression)
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
		p.SetState(89)
		p.AddingExpression()
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(HilcoPencilParserParserEQUALS-32))|(1<<(HilcoPencilParserParserNOT_EQUALS-32))|(1<<(HilcoPencilParserParserLT-32))|(1<<(HilcoPencilParserParserLTE-32))|(1<<(HilcoPencilParserParserGT-32))|(1<<(HilcoPencilParserParserGTE-32)))) != 0 {
		{
			p.SetState(90)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(HilcoPencilParserParserEQUALS-32))|(1<<(HilcoPencilParserParserNOT_EQUALS-32))|(1<<(HilcoPencilParserParserLT-32))|(1<<(HilcoPencilParserParserLTE-32))|(1<<(HilcoPencilParserParserGT-32))|(1<<(HilcoPencilParserParserGTE-32)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(91)
			p.AddingExpression()
		}

		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAddingExpressionContext is an interface to support dynamic dispatch.
type IAddingExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddingExpressionContext differentiates from other interfaces.
	IsAddingExpressionContext()
}

type AddingExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddingExpressionContext() *AddingExpressionContext {
	var p = new(AddingExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilParserParserRULE_addingExpression
	return p
}

func (*AddingExpressionContext) IsAddingExpressionContext() {}

func NewAddingExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddingExpressionContext {
	var p = new(AddingExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilParserParserRULE_addingExpression

	return p
}

func (s *AddingExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AddingExpressionContext) AllMultiplyingExpression() []IMultiplyingExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultiplyingExpressionContext)(nil)).Elem())
	var tst = make([]IMultiplyingExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultiplyingExpressionContext)
		}
	}

	return tst
}

func (s *AddingExpressionContext) MultiplyingExpression(i int) IMultiplyingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplyingExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultiplyingExpressionContext)
}

func (s *AddingExpressionContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(HilcoPencilParserParserPLUS)
}

func (s *AddingExpressionContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserPLUS, i)
}

func (s *AddingExpressionContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(HilcoPencilParserParserMINUS)
}

func (s *AddingExpressionContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserMINUS, i)
}

func (s *AddingExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddingExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddingExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.EnterAddingExpression(s)
	}
}

func (s *AddingExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.ExitAddingExpression(s)
	}
}

func (p *HilcoPencilParserParser) AddingExpression() (localctx IAddingExpressionContext) {
	localctx = NewAddingExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, HilcoPencilParserParserRULE_addingExpression)
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
		p.SetState(97)
		p.MultiplyingExpression()
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == HilcoPencilParserParserPLUS || _la == HilcoPencilParserParserMINUS {
		{
			p.SetState(98)
			_la = p.GetTokenStream().LA(1)

			if !(_la == HilcoPencilParserParserPLUS || _la == HilcoPencilParserParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(99)
			p.MultiplyingExpression()
		}

		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMultiplyingExpressionContext is an interface to support dynamic dispatch.
type IMultiplyingExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplyingExpressionContext differentiates from other interfaces.
	IsMultiplyingExpressionContext()
}

type MultiplyingExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplyingExpressionContext() *MultiplyingExpressionContext {
	var p = new(MultiplyingExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilParserParserRULE_multiplyingExpression
	return p
}

func (*MultiplyingExpressionContext) IsMultiplyingExpressionContext() {}

func NewMultiplyingExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplyingExpressionContext {
	var p = new(MultiplyingExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilParserParserRULE_multiplyingExpression

	return p
}

func (s *MultiplyingExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplyingExpressionContext) AllFactor() []IFactorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFactorContext)(nil)).Elem())
	var tst = make([]IFactorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFactorContext)
		}
	}

	return tst
}

func (s *MultiplyingExpressionContext) Factor(i int) IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *MultiplyingExpressionContext) AllTIMES() []antlr.TerminalNode {
	return s.GetTokens(HilcoPencilParserParserTIMES)
}

func (s *MultiplyingExpressionContext) TIMES(i int) antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserTIMES, i)
}

func (s *MultiplyingExpressionContext) AllDIV() []antlr.TerminalNode {
	return s.GetTokens(HilcoPencilParserParserDIV)
}

func (s *MultiplyingExpressionContext) DIV(i int) antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserDIV, i)
}

func (s *MultiplyingExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplyingExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplyingExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.EnterMultiplyingExpression(s)
	}
}

func (s *MultiplyingExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.ExitMultiplyingExpression(s)
	}
}

func (p *HilcoPencilParserParser) MultiplyingExpression() (localctx IMultiplyingExpressionContext) {
	localctx = NewMultiplyingExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, HilcoPencilParserParserRULE_multiplyingExpression)
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
		p.SetState(105)
		p.Factor()
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == HilcoPencilParserParserTIMES || _la == HilcoPencilParserParserDIV {
		{
			p.SetState(106)
			_la = p.GetTokenStream().LA(1)

			if !(_la == HilcoPencilParserParserTIMES || _la == HilcoPencilParserParserDIV) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(107)
			p.Factor()
		}

		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilParserParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilParserParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserMINUS, 0)
}

func (s *FactorContext) Factor() IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *FactorContext) Base() IBaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBaseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBaseContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *HilcoPencilParserParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, HilcoPencilParserParserRULE_factor)

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

	p.SetState(116)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case HilcoPencilParserParserMINUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(113)
			p.Match(HilcoPencilParserParserMINUS)
		}
		{
			p.SetState(114)
			p.Factor()
		}

	case HilcoPencilParserParserT__1, HilcoPencilParserParserT__2, HilcoPencilParserParserT__3, HilcoPencilParserParserT__4, HilcoPencilParserParserT__5, HilcoPencilParserParserINT, HilcoPencilParserParserID, HilcoPencilParserParserSTRING, HilcoPencilParserParserFLOAT, HilcoPencilParserParserDATE, HilcoPencilParserParserATFUNCTION, HilcoPencilParserParserCLASSNAME, HilcoPencilParserParserPERCENT, HilcoPencilParserParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(115)
			p.Base()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBaseContext is an interface to support dynamic dispatch.
type IBaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBaseContext differentiates from other interfaces.
	IsBaseContext()
}

type BaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaseContext() *BaseContext {
	var p = new(BaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilParserParserRULE_base
	return p
}

func (*BaseContext) IsBaseContext() {}

func NewBaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseContext {
	var p = new(BaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilParserParserRULE_base

	return p
}

func (s *BaseContext) GetParser() antlr.Parser { return s.parser }

func (s *BaseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserLPAREN, 0)
}

func (s *BaseContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BaseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserRPAREN, 0)
}

func (s *BaseContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *BaseContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserPERCENT, 0)
}

func (s *BaseContext) INT() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserINT, 0)
}

func (s *BaseContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserFLOAT, 0)
}

func (s *BaseContext) STRING() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserSTRING, 0)
}

func (s *BaseContext) DATE() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserDATE, 0)
}

func (s *BaseContext) AtFunction() IAtFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtFunctionContext)
}

func (s *BaseContext) DataAccessor() IDataAccessorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataAccessorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataAccessorContext)
}

func (s *BaseContext) SpecialKeyword() ISpecialKeywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecialKeywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecialKeywordContext)
}

func (s *BaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.EnterBase(s)
	}
}

func (s *BaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.ExitBase(s)
	}
}

func (p *HilcoPencilParserParser) Base() (localctx IBaseContext) {
	localctx = NewBaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, HilcoPencilParserParserRULE_base)

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

	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(118)
			p.Match(HilcoPencilParserParserLPAREN)
		}
		{
			p.SetState(119)
			p.Statement()
		}
		{
			p.SetState(120)
			p.Match(HilcoPencilParserParserRPAREN)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.Name()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(123)
			p.Match(HilcoPencilParserParserPERCENT)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(124)
			p.Match(HilcoPencilParserParserINT)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(125)
			p.Match(HilcoPencilParserParserFLOAT)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(126)
			p.Match(HilcoPencilParserParserSTRING)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(127)
			p.Match(HilcoPencilParserParserDATE)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(128)
			p.AtFunction()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(129)
			p.DataAccessor()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(130)
			p.SpecialKeyword()
		}

	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilParserParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilParserParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) ID() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserID, 0)
}

func (s *NameContext) CLASSNAME() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserCLASSNAME, 0)
}

func (s *NameContext) COLON() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserCOLON, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *HilcoPencilParserParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, HilcoPencilParserParserRULE_name)

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

	p.SetState(137)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case HilcoPencilParserParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(133)
			p.Match(HilcoPencilParserParserID)
		}

	case HilcoPencilParserParserCLASSNAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(134)
			p.Match(HilcoPencilParserParserCLASSNAME)
		}
		{
			p.SetState(135)
			p.Match(HilcoPencilParserParserCOLON)
		}
		{
			p.SetState(136)
			p.Match(HilcoPencilParserParserID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAtFunctionContext is an interface to support dynamic dispatch.
type IAtFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtFunctionContext differentiates from other interfaces.
	IsAtFunctionContext()
}

type AtFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtFunctionContext() *AtFunctionContext {
	var p = new(AtFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilParserParserRULE_atFunction
	return p
}

func (*AtFunctionContext) IsAtFunctionContext() {}

func NewAtFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtFunctionContext {
	var p = new(AtFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilParserParserRULE_atFunction

	return p
}

func (s *AtFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *AtFunctionContext) ATFUNCTION() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserATFUNCTION, 0)
}

func (s *AtFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserLPAREN, 0)
}

func (s *AtFunctionContext) ArgList() IArgListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *AtFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserRPAREN, 0)
}

func (s *AtFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.EnterAtFunction(s)
	}
}

func (s *AtFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.ExitAtFunction(s)
	}
}

func (p *HilcoPencilParserParser) AtFunction() (localctx IAtFunctionContext) {
	localctx = NewAtFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, HilcoPencilParserParserRULE_atFunction)

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
		p.SetState(139)
		p.Match(HilcoPencilParserParserATFUNCTION)
	}
	{
		p.SetState(140)
		p.Match(HilcoPencilParserParserLPAREN)
	}
	{
		p.SetState(141)
		p.ArgList()
	}
	{
		p.SetState(142)
		p.Match(HilcoPencilParserParserRPAREN)
	}

	return localctx
}

// IDataAccessorContext is an interface to support dynamic dispatch.
type IDataAccessorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataAccessorContext differentiates from other interfaces.
	IsDataAccessorContext()
}

type DataAccessorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataAccessorContext() *DataAccessorContext {
	var p = new(DataAccessorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilParserParserRULE_dataAccessor
	return p
}

func (*DataAccessorContext) IsDataAccessorContext() {}

func NewDataAccessorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataAccessorContext {
	var p = new(DataAccessorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilParserParserRULE_dataAccessor

	return p
}

func (s *DataAccessorContext) GetParser() antlr.Parser { return s.parser }

func (s *DataAccessorContext) CLASSNAME() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserCLASSNAME, 0)
}

func (s *DataAccessorContext) AllFieldName() []IFieldNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldNameContext)(nil)).Elem())
	var tst = make([]IFieldNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldNameContext)
		}
	}

	return tst
}

func (s *DataAccessorContext) FieldName(i int) IFieldNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *DataAccessorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataAccessorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataAccessorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.EnterDataAccessor(s)
	}
}

func (s *DataAccessorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.ExitDataAccessor(s)
	}
}

func (p *HilcoPencilParserParser) DataAccessor() (localctx IDataAccessorContext) {
	localctx = NewDataAccessorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, HilcoPencilParserParserRULE_dataAccessor)
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
		p.SetState(144)
		p.Match(HilcoPencilParserParserCLASSNAME)
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == HilcoPencilParserParserDOT {
		{
			p.SetState(145)
			p.FieldName()
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilParserParserRULE_fieldName
	return p
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilParserParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) DOT() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserDOT, 0)
}

func (s *FieldNameContext) ID() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserID, 0)
}

func (s *FieldNameContext) CLASSNAME() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserCLASSNAME, 0)
}

func (s *FieldNameContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserLPAREN, 0)
}

func (s *FieldNameContext) ArgList() IArgListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *FieldNameContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserRPAREN, 0)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (p *HilcoPencilParserParser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, HilcoPencilParserParserRULE_fieldName)
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

	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(150)
			p.Match(HilcoPencilParserParserDOT)
		}
		{
			p.SetState(151)
			_la = p.GetTokenStream().LA(1)

			if !(_la == HilcoPencilParserParserID || _la == HilcoPencilParserParserCLASSNAME) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(152)
			p.Match(HilcoPencilParserParserDOT)
		}
		{
			p.SetState(153)
			_la = p.GetTokenStream().LA(1)

			if !(_la == HilcoPencilParserParserID || _la == HilcoPencilParserParserCLASSNAME) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(154)
			p.Match(HilcoPencilParserParserLPAREN)
		}
		{
			p.SetState(155)
			p.ArgList()
		}
		{
			p.SetState(156)
			p.Match(HilcoPencilParserParserRPAREN)
		}

	}

	return localctx
}

// IArgListContext is an interface to support dynamic dispatch.
type IArgListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgListContext differentiates from other interfaces.
	IsArgListContext()
}

type ArgListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgListContext() *ArgListContext {
	var p = new(ArgListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilParserParserRULE_argList
	return p
}

func (*ArgListContext) IsArgListContext() {}

func NewArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgListContext {
	var p = new(ArgListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilParserParserRULE_argList

	return p
}

func (s *ArgListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgListContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ArgListContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ArgListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(HilcoPencilParserParserCOMMA)
}

func (s *ArgListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(HilcoPencilParserParserCOMMA, i)
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.ExitArgList(s)
	}
}

func (p *HilcoPencilParserParser) ArgList() (localctx IArgListContext) {
	localctx = NewArgListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, HilcoPencilParserParserRULE_argList)
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

	p.SetState(172)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case HilcoPencilParserParserT__1, HilcoPencilParserParserT__2, HilcoPencilParserParserT__3, HilcoPencilParserParserT__4, HilcoPencilParserParserT__5, HilcoPencilParserParserIF, HilcoPencilParserParserCASE, HilcoPencilParserParserSWITCH, HilcoPencilParserParserNOT, HilcoPencilParserParserINT, HilcoPencilParserParserID, HilcoPencilParserParserSTRING, HilcoPencilParserParserFLOAT, HilcoPencilParserParserDATE, HilcoPencilParserParserATFUNCTION, HilcoPencilParserParserCLASSNAME, HilcoPencilParserParserPERCENT, HilcoPencilParserParserLPAREN, HilcoPencilParserParserMINUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.Statement()
		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == HilcoPencilParserParserCOMMA {
			{
				p.SetState(161)
				p.Match(HilcoPencilParserParserCOMMA)
			}
			p.SetState(164)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case HilcoPencilParserParserT__1, HilcoPencilParserParserT__2, HilcoPencilParserParserT__3, HilcoPencilParserParserT__4, HilcoPencilParserParserT__5, HilcoPencilParserParserIF, HilcoPencilParserParserCASE, HilcoPencilParserParserSWITCH, HilcoPencilParserParserNOT, HilcoPencilParserParserINT, HilcoPencilParserParserID, HilcoPencilParserParserSTRING, HilcoPencilParserParserFLOAT, HilcoPencilParserParserDATE, HilcoPencilParserParserATFUNCTION, HilcoPencilParserParserCLASSNAME, HilcoPencilParserParserPERCENT, HilcoPencilParserParserLPAREN, HilcoPencilParserParserMINUS:
				{
					p.SetState(162)
					p.Statement()
				}

			case HilcoPencilParserParserCOMMA, HilcoPencilParserParserRPAREN:

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(170)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case HilcoPencilParserParserRPAREN:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISpecialKeywordContext is an interface to support dynamic dispatch.
type ISpecialKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecialKeywordContext differentiates from other interfaces.
	IsSpecialKeywordContext()
}

type SpecialKeywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecialKeywordContext() *SpecialKeywordContext {
	var p = new(SpecialKeywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilParserParserRULE_specialKeyword
	return p
}

func (*SpecialKeywordContext) IsSpecialKeywordContext() {}

func NewSpecialKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecialKeywordContext {
	var p = new(SpecialKeywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilParserParserRULE_specialKeyword

	return p
}

func (s *SpecialKeywordContext) GetParser() antlr.Parser { return s.parser }
func (s *SpecialKeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecialKeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecialKeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.EnterSpecialKeyword(s)
	}
}

func (s *SpecialKeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilParserListener); ok {
		listenerT.ExitSpecialKeyword(s)
	}
}

func (p *HilcoPencilParserParser) SpecialKeyword() (localctx ISpecialKeywordContext) {
	localctx = NewSpecialKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, HilcoPencilParserParserRULE_specialKeyword)
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
		p.SetState(174)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<HilcoPencilParserParserT__1)|(1<<HilcoPencilParserParserT__2)|(1<<HilcoPencilParserParserT__3)|(1<<HilcoPencilParserParserT__4)|(1<<HilcoPencilParserParserT__5))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

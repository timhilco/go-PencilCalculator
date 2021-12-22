// Code generated from HilcoPencilGrammarParser.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // HilcoPencilGrammarParser

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 51, 174,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 54, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3,
	95, 10, 3, 12, 3, 14, 3, 98, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 114, 10, 6, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 127,
	10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 133, 10, 9, 3, 10, 3, 10, 3, 10, 5,
	10, 138, 10, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 5, 11, 145, 10, 11,
	7, 11, 147, 10, 11, 12, 11, 14, 11, 150, 11, 11, 3, 12, 3, 12, 7, 12, 154,
	10, 12, 12, 12, 14, 12, 157, 11, 12, 3, 13, 3, 13, 3, 13, 5, 13, 162, 10,
	13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 170, 10, 14, 3, 15,
	3, 15, 3, 15, 2, 3, 4, 16, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 2, 3, 3, 2, 39, 41, 2, 195, 2, 30, 3, 2, 2, 2, 4, 53, 3, 2, 2,
	2, 6, 99, 3, 2, 2, 2, 8, 105, 3, 2, 2, 2, 10, 113, 3, 2, 2, 2, 12, 115,
	3, 2, 2, 2, 14, 120, 3, 2, 2, 2, 16, 132, 3, 2, 2, 2, 18, 134, 3, 2, 2,
	2, 20, 141, 3, 2, 2, 2, 22, 151, 3, 2, 2, 2, 24, 158, 3, 2, 2, 2, 26, 169,
	3, 2, 2, 2, 28, 171, 3, 2, 2, 2, 30, 31, 5, 4, 3, 2, 31, 3, 3, 2, 2, 2,
	32, 33, 8, 3, 1, 2, 33, 54, 5, 6, 4, 2, 34, 54, 5, 8, 5, 2, 35, 54, 5,
	14, 8, 2, 36, 37, 7, 32, 2, 2, 37, 38, 5, 4, 3, 2, 38, 39, 7, 33, 2, 2,
	39, 54, 3, 2, 2, 2, 40, 41, 7, 18, 2, 2, 41, 54, 5, 4, 3, 26, 42, 43, 7,
	13, 2, 2, 43, 54, 5, 4, 3, 25, 44, 54, 5, 16, 9, 2, 45, 54, 5, 18, 10,
	2, 46, 54, 5, 22, 12, 2, 47, 54, 5, 28, 15, 2, 48, 54, 7, 48, 2, 2, 49,
	54, 7, 45, 2, 2, 50, 54, 7, 47, 2, 2, 51, 54, 7, 46, 2, 2, 52, 54, 7, 49,
	2, 2, 53, 32, 3, 2, 2, 2, 53, 34, 3, 2, 2, 2, 53, 35, 3, 2, 2, 2, 53, 36,
	3, 2, 2, 2, 53, 40, 3, 2, 2, 2, 53, 42, 3, 2, 2, 2, 53, 44, 3, 2, 2, 2,
	53, 45, 3, 2, 2, 2, 53, 46, 3, 2, 2, 2, 53, 47, 3, 2, 2, 2, 53, 48, 3,
	2, 2, 2, 53, 49, 3, 2, 2, 2, 53, 50, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53,
	52, 3, 2, 2, 2, 54, 96, 3, 2, 2, 2, 55, 56, 12, 24, 2, 2, 56, 57, 7, 14,
	2, 2, 57, 95, 5, 4, 3, 24, 58, 59, 12, 23, 2, 2, 59, 60, 7, 15, 2, 2, 60,
	95, 5, 4, 3, 24, 61, 62, 12, 22, 2, 2, 62, 63, 7, 16, 2, 2, 63, 95, 5,
	4, 3, 23, 64, 65, 12, 21, 2, 2, 65, 66, 7, 17, 2, 2, 66, 95, 5, 4, 3, 22,
	67, 68, 12, 20, 2, 2, 68, 69, 7, 18, 2, 2, 69, 95, 5, 4, 3, 21, 70, 71,
	12, 19, 2, 2, 71, 72, 7, 19, 2, 2, 72, 95, 5, 4, 3, 20, 73, 74, 12, 18,
	2, 2, 74, 75, 7, 20, 2, 2, 75, 95, 5, 4, 3, 19, 76, 77, 12, 17, 2, 2, 77,
	78, 7, 21, 2, 2, 78, 95, 5, 4, 3, 18, 79, 80, 12, 16, 2, 2, 80, 81, 7,
	22, 2, 2, 81, 95, 5, 4, 3, 17, 82, 83, 12, 15, 2, 2, 83, 84, 7, 23, 2,
	2, 84, 95, 5, 4, 3, 16, 85, 86, 12, 14, 2, 2, 86, 87, 7, 24, 2, 2, 87,
	95, 5, 4, 3, 15, 88, 89, 12, 13, 2, 2, 89, 90, 7, 25, 2, 2, 90, 95, 5,
	4, 3, 14, 91, 92, 12, 12, 2, 2, 92, 93, 7, 26, 2, 2, 93, 95, 5, 4, 3, 13,
	94, 55, 3, 2, 2, 2, 94, 58, 3, 2, 2, 2, 94, 61, 3, 2, 2, 2, 94, 64, 3,
	2, 2, 2, 94, 67, 3, 2, 2, 2, 94, 70, 3, 2, 2, 2, 94, 73, 3, 2, 2, 2, 94,
	76, 3, 2, 2, 2, 94, 79, 3, 2, 2, 2, 94, 82, 3, 2, 2, 2, 94, 85, 3, 2, 2,
	2, 94, 88, 3, 2, 2, 2, 94, 91, 3, 2, 2, 2, 95, 98, 3, 2, 2, 2, 96, 94,
	3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 5, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2,
	99, 100, 7, 3, 2, 2, 100, 101, 5, 4, 3, 2, 101, 102, 7, 5, 2, 2, 102, 103,
	5, 10, 6, 2, 103, 104, 7, 4, 2, 2, 104, 7, 3, 2, 2, 2, 105, 106, 7, 6,
	2, 2, 106, 107, 5, 10, 6, 2, 107, 108, 7, 7, 2, 2, 108, 9, 3, 2, 2, 2,
	109, 110, 5, 12, 7, 2, 110, 111, 5, 10, 6, 2, 111, 114, 3, 2, 2, 2, 112,
	114, 3, 2, 2, 2, 113, 109, 3, 2, 2, 2, 113, 112, 3, 2, 2, 2, 114, 11, 3,
	2, 2, 2, 115, 116, 5, 4, 3, 2, 116, 117, 7, 11, 2, 2, 117, 118, 5, 4, 3,
	2, 118, 119, 7, 12, 2, 2, 119, 13, 3, 2, 2, 2, 120, 121, 7, 8, 2, 2, 121,
	122, 5, 4, 3, 2, 122, 123, 7, 9, 2, 2, 123, 126, 5, 4, 3, 2, 124, 125,
	7, 10, 2, 2, 125, 127, 5, 4, 3, 2, 126, 124, 3, 2, 2, 2, 126, 127, 3, 2,
	2, 2, 127, 15, 3, 2, 2, 2, 128, 133, 7, 43, 2, 2, 129, 130, 7, 42, 2, 2,
	130, 131, 7, 27, 2, 2, 131, 133, 7, 43, 2, 2, 132, 128, 3, 2, 2, 2, 132,
	129, 3, 2, 2, 2, 133, 17, 3, 2, 2, 2, 134, 135, 7, 44, 2, 2, 135, 137,
	7, 32, 2, 2, 136, 138, 5, 20, 11, 2, 137, 136, 3, 2, 2, 2, 137, 138, 3,
	2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 140, 7, 33, 2, 2, 140, 19, 3, 2, 2,
	2, 141, 148, 5, 4, 3, 2, 142, 144, 7, 37, 2, 2, 143, 145, 5, 4, 3, 2, 144,
	143, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 147, 3, 2, 2, 2, 146, 142,
	3, 2, 2, 2, 147, 150, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 148, 149, 3, 2,
	2, 2, 149, 21, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 151, 155, 7, 42, 2, 2,
	152, 154, 5, 24, 13, 2, 153, 152, 3, 2, 2, 2, 154, 157, 3, 2, 2, 2, 155,
	153, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 23, 3, 2, 2, 2, 157, 155, 3,
	2, 2, 2, 158, 161, 7, 38, 2, 2, 159, 162, 5, 26, 14, 2, 160, 162, 7, 42,
	2, 2, 161, 159, 3, 2, 2, 2, 161, 160, 3, 2, 2, 2, 162, 25, 3, 2, 2, 2,
	163, 170, 7, 43, 2, 2, 164, 165, 7, 43, 2, 2, 165, 166, 7, 32, 2, 2, 166,
	167, 5, 20, 11, 2, 167, 168, 7, 33, 2, 2, 168, 170, 3, 2, 2, 2, 169, 163,
	3, 2, 2, 2, 169, 164, 3, 2, 2, 2, 170, 27, 3, 2, 2, 2, 171, 172, 9, 2,
	2, 2, 172, 29, 3, 2, 2, 2, 14, 53, 94, 96, 113, 126, 132, 137, 144, 148,
	155, 161, 169,
}
var literalNames = []string{
	"", "'case'", "'endcase'", "'is'", "'switch'", "'endswitch'", "'if'", "'then'",
	"'else'", "'::'", "';'", "'NOT'", "'^'", "'*'", "'/'", "'+'", "'-'", "'>'",
	"'>='", "'<'", "'<='", "'='", "'~='", "'AND'", "'OR'", "':'", "'['", "']'",
	"'{'", "'}'", "'('", "')'", "'_'", "'%'", "'@'", "','", "'.'", "'true'",
	"'false'", "'nil'",
}
var symbolicNames = []string{
	"", "CASE", "END_CASE", "IS", "SWITCH", "END_SWITCH", "IF", "THEN", "ELSE",
	"DOUBLE_COLON", "SEMI_COLON", "NOT", "EXPONENTIAL", "MULTIPLY", "DIVIDE",
	"ADD", "MINUS", "GREATER_THAN", "GREATER_THAN_EQUAL", "LESS_THAN", "LESS_THAN_EQUAL",
	"EQUAL", "NOT_EQUAL", "AND", "OR", "COLON", "LBRACKET", "RBRACKET", "CURLYLBRACKET",
	"CURLYRBRACKET", "LPAREN", "RPAREN", "UNDERSCORE", "PERCENT_SIGN", "AT_SIGN",
	"COMMA", "DOT", "KEYWORD_TRUE", "KEYWORD_FALSE", "KEYWORD_NIL", "CLASSNAME",
	"ID", "ATFUNCTION", "INT", "STRING", "FLOAT", "PERCENT", "DATE", "NEWLINE",
	"WS",
}

var ruleNames = []string{
	"program", "expression", "caseStatement", "switchStatement", "caseList",
	"caseItem", "ifStatement", "name", "atFunction", "argList", "dataAccessor",
	"accessList", "accessorMessage", "specialKeyword",
}

type HilcoPencilGrammarParser struct {
	*antlr.BaseParser
}

// NewHilcoPencilGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *HilcoPencilGrammarParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewHilcoPencilGrammarParser(input antlr.TokenStream) *HilcoPencilGrammarParser {
	this := new(HilcoPencilGrammarParser)
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
	this.GrammarFileName = "HilcoPencilGrammarParser.g4"

	return this
}

// HilcoPencilGrammarParser tokens.
const (
	HilcoPencilGrammarParserEOF                = antlr.TokenEOF
	HilcoPencilGrammarParserCASE               = 1
	HilcoPencilGrammarParserEND_CASE           = 2
	HilcoPencilGrammarParserIS                 = 3
	HilcoPencilGrammarParserSWITCH             = 4
	HilcoPencilGrammarParserEND_SWITCH         = 5
	HilcoPencilGrammarParserIF                 = 6
	HilcoPencilGrammarParserTHEN               = 7
	HilcoPencilGrammarParserELSE               = 8
	HilcoPencilGrammarParserDOUBLE_COLON       = 9
	HilcoPencilGrammarParserSEMI_COLON         = 10
	HilcoPencilGrammarParserNOT                = 11
	HilcoPencilGrammarParserEXPONENTIAL        = 12
	HilcoPencilGrammarParserMULTIPLY           = 13
	HilcoPencilGrammarParserDIVIDE             = 14
	HilcoPencilGrammarParserADD                = 15
	HilcoPencilGrammarParserMINUS              = 16
	HilcoPencilGrammarParserGREATER_THAN       = 17
	HilcoPencilGrammarParserGREATER_THAN_EQUAL = 18
	HilcoPencilGrammarParserLESS_THAN          = 19
	HilcoPencilGrammarParserLESS_THAN_EQUAL    = 20
	HilcoPencilGrammarParserEQUAL              = 21
	HilcoPencilGrammarParserNOT_EQUAL          = 22
	HilcoPencilGrammarParserAND                = 23
	HilcoPencilGrammarParserOR                 = 24
	HilcoPencilGrammarParserCOLON              = 25
	HilcoPencilGrammarParserLBRACKET           = 26
	HilcoPencilGrammarParserRBRACKET           = 27
	HilcoPencilGrammarParserCURLYLBRACKET      = 28
	HilcoPencilGrammarParserCURLYRBRACKET      = 29
	HilcoPencilGrammarParserLPAREN             = 30
	HilcoPencilGrammarParserRPAREN             = 31
	HilcoPencilGrammarParserUNDERSCORE         = 32
	HilcoPencilGrammarParserPERCENT_SIGN       = 33
	HilcoPencilGrammarParserAT_SIGN            = 34
	HilcoPencilGrammarParserCOMMA              = 35
	HilcoPencilGrammarParserDOT                = 36
	HilcoPencilGrammarParserKEYWORD_TRUE       = 37
	HilcoPencilGrammarParserKEYWORD_FALSE      = 38
	HilcoPencilGrammarParserKEYWORD_NIL        = 39
	HilcoPencilGrammarParserCLASSNAME          = 40
	HilcoPencilGrammarParserID                 = 41
	HilcoPencilGrammarParserATFUNCTION         = 42
	HilcoPencilGrammarParserINT                = 43
	HilcoPencilGrammarParserSTRING             = 44
	HilcoPencilGrammarParserFLOAT              = 45
	HilcoPencilGrammarParserPERCENT            = 46
	HilcoPencilGrammarParserDATE               = 47
	HilcoPencilGrammarParserNEWLINE            = 48
	HilcoPencilGrammarParserWS                 = 49
)

// HilcoPencilGrammarParser rules.
const (
	HilcoPencilGrammarParserRULE_program         = 0
	HilcoPencilGrammarParserRULE_expression      = 1
	HilcoPencilGrammarParserRULE_caseStatement   = 2
	HilcoPencilGrammarParserRULE_switchStatement = 3
	HilcoPencilGrammarParserRULE_caseList        = 4
	HilcoPencilGrammarParserRULE_caseItem        = 5
	HilcoPencilGrammarParserRULE_ifStatement     = 6
	HilcoPencilGrammarParserRULE_name            = 7
	HilcoPencilGrammarParserRULE_atFunction      = 8
	HilcoPencilGrammarParserRULE_argList         = 9
	HilcoPencilGrammarParserRULE_dataAccessor    = 10
	HilcoPencilGrammarParserRULE_accessList      = 11
	HilcoPencilGrammarParserRULE_accessorMessage = 12
	HilcoPencilGrammarParserRULE_specialKeyword  = 13
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *HilcoPencilGrammarParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HilcoPencilGrammarParserRULE_program)

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
		p.SetState(28)
		p.expression(0)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NameCalculatorContext struct {
	*ExpressionContext
}

func NewNameCalculatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NameCalculatorContext {
	var p = new(NameCalculatorContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NameCalculatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameCalculatorContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *NameCalculatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterNameCalculator(s)
	}
}

func (s *NameCalculatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitNameCalculator(s)
	}
}

type UnaryNotCalculatorContext struct {
	*ExpressionContext
}

func NewUnaryNotCalculatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryNotCalculatorContext {
	var p = new(UnaryNotCalculatorContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryNotCalculatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryNotCalculatorContext) NOT() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserNOT, 0)
}

func (s *UnaryNotCalculatorContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryNotCalculatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterUnaryNotCalculator(s)
	}
}

func (s *UnaryNotCalculatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitUnaryNotCalculator(s)
	}
}

type BinaryArthmeticCalculatorContext struct {
	*ExpressionContext
}

func NewBinaryArthmeticCalculatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryArthmeticCalculatorContext {
	var p = new(BinaryArthmeticCalculatorContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryArthmeticCalculatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryArthmeticCalculatorContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BinaryArthmeticCalculatorContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinaryArthmeticCalculatorContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserMULTIPLY, 0)
}

func (s *BinaryArthmeticCalculatorContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserDIVIDE, 0)
}

func (s *BinaryArthmeticCalculatorContext) ADD() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserADD, 0)
}

func (s *BinaryArthmeticCalculatorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserMINUS, 0)
}

func (s *BinaryArthmeticCalculatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterBinaryArthmeticCalculator(s)
	}
}

func (s *BinaryArthmeticCalculatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitBinaryArthmeticCalculator(s)
	}
}

type PercentContext struct {
	*ExpressionContext
}

func NewPercentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PercentContext {
	var p = new(PercentContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *PercentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PercentContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserPERCENT, 0)
}

func (s *PercentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterPercent(s)
	}
}

func (s *PercentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitPercent(s)
	}
}

type ParensContext struct {
	*ExpressionContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserLPAREN, 0)
}

func (s *ParensContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParensContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserRPAREN, 0)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitParens(s)
	}
}

type BinaryRelationalCalculatorContext struct {
	*ExpressionContext
}

func NewBinaryRelationalCalculatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryRelationalCalculatorContext {
	var p = new(BinaryRelationalCalculatorContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryRelationalCalculatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryRelationalCalculatorContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BinaryRelationalCalculatorContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinaryRelationalCalculatorContext) GREATER_THAN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserGREATER_THAN, 0)
}

func (s *BinaryRelationalCalculatorContext) GREATER_THAN_EQUAL() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserGREATER_THAN_EQUAL, 0)
}

func (s *BinaryRelationalCalculatorContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserLESS_THAN, 0)
}

func (s *BinaryRelationalCalculatorContext) LESS_THAN_EQUAL() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserLESS_THAN_EQUAL, 0)
}

func (s *BinaryRelationalCalculatorContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserEQUAL, 0)
}

func (s *BinaryRelationalCalculatorContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserNOT_EQUAL, 0)
}

func (s *BinaryRelationalCalculatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterBinaryRelationalCalculator(s)
	}
}

func (s *BinaryRelationalCalculatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitBinaryRelationalCalculator(s)
	}
}

type StringContext struct {
	*ExpressionContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitString(s)
	}
}

type ExpressionAtFunctionContext struct {
	*ExpressionContext
}

func NewExpressionAtFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionAtFunctionContext {
	var p = new(ExpressionAtFunctionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionAtFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAtFunctionContext) AtFunction() IAtFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtFunctionContext)
}

func (s *ExpressionAtFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterExpressionAtFunction(s)
	}
}

func (s *ExpressionAtFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitExpressionAtFunction(s)
	}
}

type DateContext struct {
	*ExpressionContext
}

func NewDateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateContext {
	var p = new(DateContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *DateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateContext) DATE() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserDATE, 0)
}

func (s *DateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterDate(s)
	}
}

func (s *DateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitDate(s)
	}
}

type CaseContext struct {
	*ExpressionContext
}

func NewCaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CaseContext {
	var p = new(CaseContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseContext) CaseStatement() ICaseStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICaseStatementContext)
}

func (s *CaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterCase(s)
	}
}

func (s *CaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitCase(s)
	}
}

type IntegerContext struct {
	*ExpressionContext
}

func NewIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerContext {
	var p = new(IntegerContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) INT() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserINT, 0)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitInteger(s)
	}
}

type BinaryLogicalCalculatorContext struct {
	*ExpressionContext
}

func NewBinaryLogicalCalculatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryLogicalCalculatorContext {
	var p = new(BinaryLogicalCalculatorContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryLogicalCalculatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryLogicalCalculatorContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BinaryLogicalCalculatorContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinaryLogicalCalculatorContext) AND() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserAND, 0)
}

func (s *BinaryLogicalCalculatorContext) OR() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserOR, 0)
}

func (s *BinaryLogicalCalculatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterBinaryLogicalCalculator(s)
	}
}

func (s *BinaryLogicalCalculatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitBinaryLogicalCalculator(s)
	}
}

type ExpressionKeywordContext struct {
	*ExpressionContext
}

func NewExpressionKeywordContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionKeywordContext {
	var p = new(ExpressionKeywordContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionKeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionKeywordContext) SpecialKeyword() ISpecialKeywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecialKeywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecialKeywordContext)
}

func (s *ExpressionKeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterExpressionKeyword(s)
	}
}

func (s *ExpressionKeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitExpressionKeyword(s)
	}
}

type FloatContext struct {
	*ExpressionContext
}

func NewFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatContext {
	var p = new(FloatContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserFLOAT, 0)
}

func (s *FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterFloat(s)
	}
}

func (s *FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitFloat(s)
	}
}

type ExpressionSwitchContext struct {
	*ExpressionContext
}

func NewExpressionSwitchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionSwitchContext {
	var p = new(ExpressionSwitchContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionSwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSwitchContext) SwitchStatement() ISwitchStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISwitchStatementContext)
}

func (s *ExpressionSwitchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterExpressionSwitch(s)
	}
}

func (s *ExpressionSwitchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitExpressionSwitch(s)
	}
}

type UnaryMinusCalculatorContext struct {
	*ExpressionContext
}

func NewUnaryMinusCalculatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryMinusCalculatorContext {
	var p = new(UnaryMinusCalculatorContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryMinusCalculatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMinusCalculatorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserMINUS, 0)
}

func (s *UnaryMinusCalculatorContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryMinusCalculatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterUnaryMinusCalculator(s)
	}
}

func (s *UnaryMinusCalculatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitUnaryMinusCalculator(s)
	}
}

type BinaryExponentialCalculatorContext struct {
	*ExpressionContext
}

func NewBinaryExponentialCalculatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExponentialCalculatorContext {
	var p = new(BinaryExponentialCalculatorContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryExponentialCalculatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExponentialCalculatorContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BinaryExponentialCalculatorContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinaryExponentialCalculatorContext) EXPONENTIAL() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserEXPONENTIAL, 0)
}

func (s *BinaryExponentialCalculatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterBinaryExponentialCalculator(s)
	}
}

func (s *BinaryExponentialCalculatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitBinaryExponentialCalculator(s)
	}
}

type ExpressionDataAccessContext struct {
	*ExpressionContext
}

func NewExpressionDataAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionDataAccessContext {
	var p = new(ExpressionDataAccessContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionDataAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionDataAccessContext) DataAccessor() IDataAccessorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataAccessorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataAccessorContext)
}

func (s *ExpressionDataAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterExpressionDataAccess(s)
	}
}

func (s *ExpressionDataAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitExpressionDataAccess(s)
	}
}

type IfContext struct {
	*ExpressionContext
}

func NewIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfContext {
	var p = new(IfContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterIf(s)
	}
}

func (s *IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitIf(s)
	}
}

func (p *HilcoPencilGrammarParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *HilcoPencilGrammarParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, HilcoPencilGrammarParserRULE_expression, _p)

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
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCaseContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(31)
			p.CaseStatement()
		}

	case 2:
		localctx = NewExpressionSwitchContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(32)
			p.SwitchStatement()
		}

	case 3:
		localctx = NewIfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(33)
			p.IfStatement()
		}

	case 4:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(34)
			p.Match(HilcoPencilGrammarParserLPAREN)
		}
		{
			p.SetState(35)
			p.expression(0)
		}
		{
			p.SetState(36)
			p.Match(HilcoPencilGrammarParserRPAREN)
		}

	case 5:
		localctx = NewUnaryMinusCalculatorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(38)
			p.Match(HilcoPencilGrammarParserMINUS)
		}
		{
			p.SetState(39)
			p.expression(24)
		}

	case 6:
		localctx = NewUnaryNotCalculatorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(40)
			p.Match(HilcoPencilGrammarParserNOT)
		}
		{
			p.SetState(41)
			p.expression(23)
		}

	case 7:
		localctx = NewNameCalculatorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(42)
			p.Name()
		}

	case 8:
		localctx = NewExpressionAtFunctionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(43)
			p.AtFunction()
		}

	case 9:
		localctx = NewExpressionDataAccessContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(44)
			p.DataAccessor()
		}

	case 10:
		localctx = NewExpressionKeywordContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(45)
			p.SpecialKeyword()
		}

	case 11:
		localctx = NewPercentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(46)
			p.Match(HilcoPencilGrammarParserPERCENT)
		}

	case 12:
		localctx = NewIntegerContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(47)
			p.Match(HilcoPencilGrammarParserINT)
		}

	case 13:
		localctx = NewFloatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(48)
			p.Match(HilcoPencilGrammarParserFLOAT)
		}

	case 14:
		localctx = NewStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(49)
			p.Match(HilcoPencilGrammarParserSTRING)
		}

	case 15:
		localctx = NewDateContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(50)
			p.Match(HilcoPencilGrammarParserDATE)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(92)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExponentialCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(53)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
				}
				{
					p.SetState(54)
					p.Match(HilcoPencilGrammarParserEXPONENTIAL)
				}
				{
					p.SetState(55)
					p.expression(22)
				}

			case 2:
				localctx = NewBinaryArthmeticCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(56)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
				}
				{
					p.SetState(57)
					p.Match(HilcoPencilGrammarParserMULTIPLY)
				}
				{
					p.SetState(58)
					p.expression(22)
				}

			case 3:
				localctx = NewBinaryArthmeticCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(59)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
				}
				{
					p.SetState(60)
					p.Match(HilcoPencilGrammarParserDIVIDE)
				}
				{
					p.SetState(61)
					p.expression(21)
				}

			case 4:
				localctx = NewBinaryArthmeticCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(62)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(63)
					p.Match(HilcoPencilGrammarParserADD)
				}
				{
					p.SetState(64)
					p.expression(20)
				}

			case 5:
				localctx = NewBinaryArthmeticCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(65)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(66)
					p.Match(HilcoPencilGrammarParserMINUS)
				}
				{
					p.SetState(67)
					p.expression(19)
				}

			case 6:
				localctx = NewBinaryRelationalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(68)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(69)
					p.Match(HilcoPencilGrammarParserGREATER_THAN)
				}
				{
					p.SetState(70)
					p.expression(18)
				}

			case 7:
				localctx = NewBinaryRelationalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(71)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(72)
					p.Match(HilcoPencilGrammarParserGREATER_THAN_EQUAL)
				}
				{
					p.SetState(73)
					p.expression(17)
				}

			case 8:
				localctx = NewBinaryRelationalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(74)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(75)
					p.Match(HilcoPencilGrammarParserLESS_THAN)
				}
				{
					p.SetState(76)
					p.expression(16)
				}

			case 9:
				localctx = NewBinaryRelationalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(77)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(78)
					p.Match(HilcoPencilGrammarParserLESS_THAN_EQUAL)
				}
				{
					p.SetState(79)
					p.expression(15)
				}

			case 10:
				localctx = NewBinaryRelationalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(80)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(81)
					p.Match(HilcoPencilGrammarParserEQUAL)
				}
				{
					p.SetState(82)
					p.expression(14)
				}

			case 11:
				localctx = NewBinaryRelationalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(83)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(84)
					p.Match(HilcoPencilGrammarParserNOT_EQUAL)
				}
				{
					p.SetState(85)
					p.expression(13)
				}

			case 12:
				localctx = NewBinaryLogicalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(86)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(87)
					p.Match(HilcoPencilGrammarParserAND)
				}
				{
					p.SetState(88)
					p.expression(12)
				}

			case 13:
				localctx = NewBinaryLogicalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(89)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(90)
					p.Match(HilcoPencilGrammarParserOR)
				}
				{
					p.SetState(91)
					p.expression(11)
				}

			}

		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
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
	p.RuleIndex = HilcoPencilGrammarParserRULE_caseStatement
	return p
}

func (*CaseStatementContext) IsCaseStatementContext() {}

func NewCaseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseStatementContext {
	var p = new(CaseStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_caseStatement

	return p
}

func (s *CaseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseStatementContext) CASE() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserCASE, 0)
}

func (s *CaseStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CaseStatementContext) IS() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserIS, 0)
}

func (s *CaseStatementContext) CaseList() ICaseListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICaseListContext)
}

func (s *CaseStatementContext) END_CASE() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserEND_CASE, 0)
}

func (s *CaseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterCaseStatement(s)
	}
}

func (s *CaseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitCaseStatement(s)
	}
}

func (p *HilcoPencilGrammarParser) CaseStatement() (localctx ICaseStatementContext) {
	localctx = NewCaseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, HilcoPencilGrammarParserRULE_caseStatement)

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
		p.Match(HilcoPencilGrammarParserCASE)
	}
	{
		p.SetState(98)
		p.expression(0)
	}
	{
		p.SetState(99)
		p.Match(HilcoPencilGrammarParserIS)
	}
	{
		p.SetState(100)
		p.CaseList()
	}
	{
		p.SetState(101)
		p.Match(HilcoPencilGrammarParserEND_CASE)
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
	p.RuleIndex = HilcoPencilGrammarParserRULE_switchStatement
	return p
}

func (*SwitchStatementContext) IsSwitchStatementContext() {}

func NewSwitchStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_switchStatement

	return p
}

func (s *SwitchStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStatementContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserSWITCH, 0)
}

func (s *SwitchStatementContext) CaseList() ICaseListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICaseListContext)
}

func (s *SwitchStatementContext) END_SWITCH() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserEND_SWITCH, 0)
}

func (s *SwitchStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitSwitchStatement(s)
	}
}

func (p *HilcoPencilGrammarParser) SwitchStatement() (localctx ISwitchStatementContext) {
	localctx = NewSwitchStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, HilcoPencilGrammarParserRULE_switchStatement)

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
		p.SetState(103)
		p.Match(HilcoPencilGrammarParserSWITCH)
	}
	{
		p.SetState(104)
		p.CaseList()
	}
	{
		p.SetState(105)
		p.Match(HilcoPencilGrammarParserEND_SWITCH)
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
	p.RuleIndex = HilcoPencilGrammarParserRULE_caseList
	return p
}

func (*CaseListContext) IsCaseListContext() {}

func NewCaseListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseListContext {
	var p = new(CaseListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_caseList

	return p
}

func (s *CaseListContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseListContext) CaseItem() ICaseItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseItemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICaseItemContext)
}

func (s *CaseListContext) CaseList() ICaseListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICaseListContext)
}

func (s *CaseListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterCaseList(s)
	}
}

func (s *CaseListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitCaseList(s)
	}
}

func (p *HilcoPencilGrammarParser) CaseList() (localctx ICaseListContext) {
	localctx = NewCaseListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, HilcoPencilGrammarParserRULE_caseList)

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

	p.SetState(111)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case HilcoPencilGrammarParserCASE, HilcoPencilGrammarParserSWITCH, HilcoPencilGrammarParserIF, HilcoPencilGrammarParserNOT, HilcoPencilGrammarParserMINUS, HilcoPencilGrammarParserLPAREN, HilcoPencilGrammarParserKEYWORD_TRUE, HilcoPencilGrammarParserKEYWORD_FALSE, HilcoPencilGrammarParserKEYWORD_NIL, HilcoPencilGrammarParserCLASSNAME, HilcoPencilGrammarParserID, HilcoPencilGrammarParserATFUNCTION, HilcoPencilGrammarParserINT, HilcoPencilGrammarParserSTRING, HilcoPencilGrammarParserFLOAT, HilcoPencilGrammarParserPERCENT, HilcoPencilGrammarParserDATE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.CaseItem()
		}
		{
			p.SetState(108)
			p.CaseList()
		}

	case HilcoPencilGrammarParserEND_CASE, HilcoPencilGrammarParserEND_SWITCH:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = HilcoPencilGrammarParserRULE_caseItem
	return p
}

func (*CaseItemContext) IsCaseItemContext() {}

func NewCaseItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseItemContext {
	var p = new(CaseItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_caseItem

	return p
}

func (s *CaseItemContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseItemContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *CaseItemContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CaseItemContext) DOUBLE_COLON() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserDOUBLE_COLON, 0)
}

func (s *CaseItemContext) SEMI_COLON() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserSEMI_COLON, 0)
}

func (s *CaseItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterCaseItem(s)
	}
}

func (s *CaseItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitCaseItem(s)
	}
}

func (p *HilcoPencilGrammarParser) CaseItem() (localctx ICaseItemContext) {
	localctx = NewCaseItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, HilcoPencilGrammarParserRULE_caseItem)

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
		p.SetState(113)
		p.expression(0)
	}
	{
		p.SetState(114)
		p.Match(HilcoPencilGrammarParserDOUBLE_COLON)
	}
	{
		p.SetState(115)
		p.expression(0)
	}
	{
		p.SetState(116)
		p.Match(HilcoPencilGrammarParserSEMI_COLON)
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
	p.RuleIndex = HilcoPencilGrammarParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserIF, 0)
}

func (s *IfStatementContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *IfStatementContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) THEN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserTHEN, 0)
}

func (s *IfStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserELSE, 0)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (p *HilcoPencilGrammarParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, HilcoPencilGrammarParserRULE_ifStatement)

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
		p.SetState(118)
		p.Match(HilcoPencilGrammarParserIF)
	}
	{
		p.SetState(119)
		p.expression(0)
	}
	{
		p.SetState(120)
		p.Match(HilcoPencilGrammarParserTHEN)
	}
	{
		p.SetState(121)
		p.expression(0)
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(122)
			p.Match(HilcoPencilGrammarParserELSE)
		}
		{
			p.SetState(123)
			p.expression(0)
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
	p.RuleIndex = HilcoPencilGrammarParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) ID() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserID, 0)
}

func (s *NameContext) CLASSNAME() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserCLASSNAME, 0)
}

func (s *NameContext) COLON() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserCOLON, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *HilcoPencilGrammarParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, HilcoPencilGrammarParserRULE_name)

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

	p.SetState(130)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case HilcoPencilGrammarParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(126)
			p.Match(HilcoPencilGrammarParserID)
		}

	case HilcoPencilGrammarParserCLASSNAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(127)
			p.Match(HilcoPencilGrammarParserCLASSNAME)
		}
		{
			p.SetState(128)
			p.Match(HilcoPencilGrammarParserCOLON)
		}
		{
			p.SetState(129)
			p.Match(HilcoPencilGrammarParserID)
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
	p.RuleIndex = HilcoPencilGrammarParserRULE_atFunction
	return p
}

func (*AtFunctionContext) IsAtFunctionContext() {}

func NewAtFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtFunctionContext {
	var p = new(AtFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_atFunction

	return p
}

func (s *AtFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *AtFunctionContext) ATFUNCTION() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserATFUNCTION, 0)
}

func (s *AtFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserLPAREN, 0)
}

func (s *AtFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserRPAREN, 0)
}

func (s *AtFunctionContext) ArgList() IArgListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *AtFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterAtFunction(s)
	}
}

func (s *AtFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitAtFunction(s)
	}
}

func (p *HilcoPencilGrammarParser) AtFunction() (localctx IAtFunctionContext) {
	localctx = NewAtFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, HilcoPencilGrammarParserRULE_atFunction)
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
		p.SetState(132)
		p.Match(HilcoPencilGrammarParserATFUNCTION)
	}
	{
		p.SetState(133)
		p.Match(HilcoPencilGrammarParserLPAREN)
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<HilcoPencilGrammarParserCASE)|(1<<HilcoPencilGrammarParserSWITCH)|(1<<HilcoPencilGrammarParserIF)|(1<<HilcoPencilGrammarParserNOT)|(1<<HilcoPencilGrammarParserMINUS)|(1<<HilcoPencilGrammarParserLPAREN))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(HilcoPencilGrammarParserKEYWORD_TRUE-37))|(1<<(HilcoPencilGrammarParserKEYWORD_FALSE-37))|(1<<(HilcoPencilGrammarParserKEYWORD_NIL-37))|(1<<(HilcoPencilGrammarParserCLASSNAME-37))|(1<<(HilcoPencilGrammarParserID-37))|(1<<(HilcoPencilGrammarParserATFUNCTION-37))|(1<<(HilcoPencilGrammarParserINT-37))|(1<<(HilcoPencilGrammarParserSTRING-37))|(1<<(HilcoPencilGrammarParserFLOAT-37))|(1<<(HilcoPencilGrammarParserPERCENT-37))|(1<<(HilcoPencilGrammarParserDATE-37)))) != 0) {
		{
			p.SetState(134)
			p.ArgList()
		}

	}
	{
		p.SetState(137)
		p.Match(HilcoPencilGrammarParserRPAREN)
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
	p.RuleIndex = HilcoPencilGrammarParserRULE_argList
	return p
}

func (*ArgListContext) IsArgListContext() {}

func NewArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgListContext {
	var p = new(ArgListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_argList

	return p
}

func (s *ArgListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ArgListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(HilcoPencilGrammarParserCOMMA)
}

func (s *ArgListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserCOMMA, i)
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitArgList(s)
	}
}

func (p *HilcoPencilGrammarParser) ArgList() (localctx IArgListContext) {
	localctx = NewArgListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, HilcoPencilGrammarParserRULE_argList)
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
		p.SetState(139)
		p.expression(0)
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == HilcoPencilGrammarParserCOMMA {
		{
			p.SetState(140)
			p.Match(HilcoPencilGrammarParserCOMMA)
		}
		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<HilcoPencilGrammarParserCASE)|(1<<HilcoPencilGrammarParserSWITCH)|(1<<HilcoPencilGrammarParserIF)|(1<<HilcoPencilGrammarParserNOT)|(1<<HilcoPencilGrammarParserMINUS)|(1<<HilcoPencilGrammarParserLPAREN))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(HilcoPencilGrammarParserKEYWORD_TRUE-37))|(1<<(HilcoPencilGrammarParserKEYWORD_FALSE-37))|(1<<(HilcoPencilGrammarParserKEYWORD_NIL-37))|(1<<(HilcoPencilGrammarParserCLASSNAME-37))|(1<<(HilcoPencilGrammarParserID-37))|(1<<(HilcoPencilGrammarParserATFUNCTION-37))|(1<<(HilcoPencilGrammarParserINT-37))|(1<<(HilcoPencilGrammarParserSTRING-37))|(1<<(HilcoPencilGrammarParserFLOAT-37))|(1<<(HilcoPencilGrammarParserPERCENT-37))|(1<<(HilcoPencilGrammarParserDATE-37)))) != 0) {
			{
				p.SetState(141)
				p.expression(0)
			}

		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = HilcoPencilGrammarParserRULE_dataAccessor
	return p
}

func (*DataAccessorContext) IsDataAccessorContext() {}

func NewDataAccessorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataAccessorContext {
	var p = new(DataAccessorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_dataAccessor

	return p
}

func (s *DataAccessorContext) GetParser() antlr.Parser { return s.parser }

func (s *DataAccessorContext) CLASSNAME() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserCLASSNAME, 0)
}

func (s *DataAccessorContext) AllAccessList() []IAccessListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAccessListContext)(nil)).Elem())
	var tst = make([]IAccessListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAccessListContext)
		}
	}

	return tst
}

func (s *DataAccessorContext) AccessList(i int) IAccessListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccessListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAccessListContext)
}

func (s *DataAccessorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataAccessorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataAccessorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterDataAccessor(s)
	}
}

func (s *DataAccessorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitDataAccessor(s)
	}
}

func (p *HilcoPencilGrammarParser) DataAccessor() (localctx IDataAccessorContext) {
	localctx = NewDataAccessorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, HilcoPencilGrammarParserRULE_dataAccessor)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(HilcoPencilGrammarParserCLASSNAME)
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(150)
				p.AccessList()
			}

		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IAccessListContext is an interface to support dynamic dispatch.
type IAccessListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccessListContext differentiates from other interfaces.
	IsAccessListContext()
}

type AccessListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessListContext() *AccessListContext {
	var p = new(AccessListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessList
	return p
}

func (*AccessListContext) IsAccessListContext() {}

func NewAccessListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessListContext {
	var p = new(AccessListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessList

	return p
}

func (s *AccessListContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessListContext) DOT() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserDOT, 0)
}

func (s *AccessListContext) AccessorMessage() IAccessorMessageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccessorMessageContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccessorMessageContext)
}

func (s *AccessListContext) CLASSNAME() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserCLASSNAME, 0)
}

func (s *AccessListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterAccessList(s)
	}
}

func (s *AccessListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitAccessList(s)
	}
}

func (p *HilcoPencilGrammarParser) AccessList() (localctx IAccessListContext) {
	localctx = NewAccessListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, HilcoPencilGrammarParserRULE_accessList)

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
		p.SetState(156)
		p.Match(HilcoPencilGrammarParserDOT)
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case HilcoPencilGrammarParserID:
		{
			p.SetState(157)
			p.AccessorMessage()
		}

	case HilcoPencilGrammarParserCLASSNAME:
		{
			p.SetState(158)
			p.Match(HilcoPencilGrammarParserCLASSNAME)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAccessorMessageContext is an interface to support dynamic dispatch.
type IAccessorMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccessorMessageContext differentiates from other interfaces.
	IsAccessorMessageContext()
}

type AccessorMessageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessorMessageContext() *AccessorMessageContext {
	var p = new(AccessorMessageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorMessage
	return p
}

func (*AccessorMessageContext) IsAccessorMessageContext() {}

func NewAccessorMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessorMessageContext {
	var p = new(AccessorMessageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorMessage

	return p
}

func (s *AccessorMessageContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessorMessageContext) ID() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserID, 0)
}

func (s *AccessorMessageContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserLPAREN, 0)
}

func (s *AccessorMessageContext) ArgList() IArgListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *AccessorMessageContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserRPAREN, 0)
}

func (s *AccessorMessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessorMessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessorMessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterAccessorMessage(s)
	}
}

func (s *AccessorMessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitAccessorMessage(s)
	}
}

func (p *HilcoPencilGrammarParser) AccessorMessage() (localctx IAccessorMessageContext) {
	localctx = NewAccessorMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, HilcoPencilGrammarParserRULE_accessorMessage)

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

	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			p.Match(HilcoPencilGrammarParserID)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.Match(HilcoPencilGrammarParserID)
		}
		{
			p.SetState(163)
			p.Match(HilcoPencilGrammarParserLPAREN)
		}
		{
			p.SetState(164)
			p.ArgList()
		}
		{
			p.SetState(165)
			p.Match(HilcoPencilGrammarParserRPAREN)
		}

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
	p.RuleIndex = HilcoPencilGrammarParserRULE_specialKeyword
	return p
}

func (*SpecialKeywordContext) IsSpecialKeywordContext() {}

func NewSpecialKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecialKeywordContext {
	var p = new(SpecialKeywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_specialKeyword

	return p
}

func (s *SpecialKeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecialKeywordContext) KEYWORD_TRUE() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserKEYWORD_TRUE, 0)
}

func (s *SpecialKeywordContext) KEYWORD_FALSE() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserKEYWORD_FALSE, 0)
}

func (s *SpecialKeywordContext) KEYWORD_NIL() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserKEYWORD_NIL, 0)
}

func (s *SpecialKeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecialKeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecialKeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterSpecialKeyword(s)
	}
}

func (s *SpecialKeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitSpecialKeyword(s)
	}
}

func (p *HilcoPencilGrammarParser) SpecialKeyword() (localctx ISpecialKeywordContext) {
	localctx = NewSpecialKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, HilcoPencilGrammarParserRULE_specialKeyword)
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
		p.SetState(169)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(HilcoPencilGrammarParserKEYWORD_TRUE-37))|(1<<(HilcoPencilGrammarParserKEYWORD_FALSE-37))|(1<<(HilcoPencilGrammarParserKEYWORD_NIL-37)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *HilcoPencilGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *HilcoPencilGrammarParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

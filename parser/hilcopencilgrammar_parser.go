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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 52, 189,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 3, 2, 3, 2, 7, 2, 39, 10, 2, 12, 2, 14, 2, 42, 11, 2, 3, 2, 5, 2, 45,
	10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 5, 4, 73, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 114, 10, 4,
	12, 4, 14, 4, 117, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 5, 6, 129, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 142, 10, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 5, 11, 153, 10, 11, 3, 11, 3, 11, 3,
	12, 3, 12, 3, 12, 5, 12, 160, 10, 12, 7, 12, 162, 10, 12, 12, 12, 14, 12,
	165, 11, 12, 3, 13, 3, 13, 6, 13, 169, 10, 13, 13, 13, 14, 13, 170, 3,
	14, 3, 14, 3, 14, 3, 15, 3, 15, 5, 15, 178, 10, 15, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 2, 3, 6, 19, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 2, 3, 3, 2, 40,
	42, 2, 208, 2, 44, 3, 2, 2, 2, 4, 46, 3, 2, 2, 2, 6, 72, 3, 2, 2, 2, 8,
	118, 3, 2, 2, 2, 10, 128, 3, 2, 2, 2, 12, 130, 3, 2, 2, 2, 14, 135, 3,
	2, 2, 2, 16, 143, 3, 2, 2, 2, 18, 145, 3, 2, 2, 2, 20, 149, 3, 2, 2, 2,
	22, 156, 3, 2, 2, 2, 24, 166, 3, 2, 2, 2, 26, 172, 3, 2, 2, 2, 28, 177,
	3, 2, 2, 2, 30, 179, 3, 2, 2, 2, 32, 181, 3, 2, 2, 2, 34, 186, 3, 2, 2,
	2, 36, 45, 5, 6, 4, 2, 37, 39, 5, 4, 3, 2, 38, 37, 3, 2, 2, 2, 39, 42,
	3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 45, 3, 2, 2, 2,
	42, 40, 3, 2, 2, 2, 43, 45, 3, 2, 2, 2, 44, 36, 3, 2, 2, 2, 44, 40, 3,
	2, 2, 2, 44, 43, 3, 2, 2, 2, 45, 3, 3, 2, 2, 2, 46, 47, 5, 6, 4, 2, 47,
	48, 7, 3, 2, 2, 48, 49, 5, 6, 4, 2, 49, 50, 7, 13, 2, 2, 50, 5, 3, 2, 2,
	2, 51, 52, 8, 4, 1, 2, 52, 73, 5, 8, 5, 2, 53, 73, 5, 14, 8, 2, 54, 55,
	7, 33, 2, 2, 55, 56, 5, 6, 4, 2, 56, 57, 7, 34, 2, 2, 57, 73, 3, 2, 2,
	2, 58, 59, 7, 19, 2, 2, 59, 73, 5, 6, 4, 27, 60, 61, 7, 14, 2, 2, 61, 73,
	5, 6, 4, 26, 62, 73, 5, 16, 9, 2, 63, 73, 5, 18, 10, 2, 64, 73, 5, 20,
	11, 2, 65, 73, 5, 24, 13, 2, 66, 73, 5, 34, 18, 2, 67, 73, 7, 49, 2, 2,
	68, 73, 7, 46, 2, 2, 69, 73, 7, 48, 2, 2, 70, 73, 7, 47, 2, 2, 71, 73,
	7, 50, 2, 2, 72, 51, 3, 2, 2, 2, 72, 53, 3, 2, 2, 2, 72, 54, 3, 2, 2, 2,
	72, 58, 3, 2, 2, 2, 72, 60, 3, 2, 2, 2, 72, 62, 3, 2, 2, 2, 72, 63, 3,
	2, 2, 2, 72, 64, 3, 2, 2, 2, 72, 65, 3, 2, 2, 2, 72, 66, 3, 2, 2, 2, 72,
	67, 3, 2, 2, 2, 72, 68, 3, 2, 2, 2, 72, 69, 3, 2, 2, 2, 72, 70, 3, 2, 2,
	2, 72, 71, 3, 2, 2, 2, 73, 115, 3, 2, 2, 2, 74, 75, 12, 25, 2, 2, 75, 76,
	7, 15, 2, 2, 76, 114, 5, 6, 4, 25, 77, 78, 12, 24, 2, 2, 78, 79, 7, 16,
	2, 2, 79, 114, 5, 6, 4, 25, 80, 81, 12, 23, 2, 2, 81, 82, 7, 17, 2, 2,
	82, 114, 5, 6, 4, 24, 83, 84, 12, 22, 2, 2, 84, 85, 7, 18, 2, 2, 85, 114,
	5, 6, 4, 23, 86, 87, 12, 21, 2, 2, 87, 88, 7, 19, 2, 2, 88, 114, 5, 6,
	4, 22, 89, 90, 12, 20, 2, 2, 90, 91, 7, 20, 2, 2, 91, 114, 5, 6, 4, 21,
	92, 93, 12, 19, 2, 2, 93, 94, 7, 21, 2, 2, 94, 114, 5, 6, 4, 20, 95, 96,
	12, 18, 2, 2, 96, 97, 7, 22, 2, 2, 97, 114, 5, 6, 4, 19, 98, 99, 12, 17,
	2, 2, 99, 100, 7, 23, 2, 2, 100, 114, 5, 6, 4, 18, 101, 102, 12, 16, 2,
	2, 102, 103, 7, 24, 2, 2, 103, 114, 5, 6, 4, 17, 104, 105, 12, 15, 2, 2,
	105, 106, 7, 25, 2, 2, 106, 114, 5, 6, 4, 16, 107, 108, 12, 14, 2, 2, 108,
	109, 7, 26, 2, 2, 109, 114, 5, 6, 4, 15, 110, 111, 12, 13, 2, 2, 111, 112,
	7, 27, 2, 2, 112, 114, 5, 6, 4, 14, 113, 74, 3, 2, 2, 2, 113, 77, 3, 2,
	2, 2, 113, 80, 3, 2, 2, 2, 113, 83, 3, 2, 2, 2, 113, 86, 3, 2, 2, 2, 113,
	89, 3, 2, 2, 2, 113, 92, 3, 2, 2, 2, 113, 95, 3, 2, 2, 2, 113, 98, 3, 2,
	2, 2, 113, 101, 3, 2, 2, 2, 113, 104, 3, 2, 2, 2, 113, 107, 3, 2, 2, 2,
	113, 110, 3, 2, 2, 2, 114, 117, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 115,
	116, 3, 2, 2, 2, 116, 7, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 118, 119, 7,
	4, 2, 2, 119, 120, 5, 6, 4, 2, 120, 121, 7, 6, 2, 2, 121, 122, 5, 10, 6,
	2, 122, 123, 7, 5, 2, 2, 123, 9, 3, 2, 2, 2, 124, 125, 5, 12, 7, 2, 125,
	126, 5, 10, 6, 2, 126, 129, 3, 2, 2, 2, 127, 129, 3, 2, 2, 2, 128, 124,
	3, 2, 2, 2, 128, 127, 3, 2, 2, 2, 129, 11, 3, 2, 2, 2, 130, 131, 5, 6,
	4, 2, 131, 132, 7, 12, 2, 2, 132, 133, 5, 6, 4, 2, 133, 134, 7, 13, 2,
	2, 134, 13, 3, 2, 2, 2, 135, 136, 7, 9, 2, 2, 136, 137, 5, 6, 4, 2, 137,
	138, 7, 10, 2, 2, 138, 141, 5, 6, 4, 2, 139, 140, 7, 11, 2, 2, 140, 142,
	5, 6, 4, 2, 141, 139, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 15, 3, 2,
	2, 2, 143, 144, 7, 44, 2, 2, 144, 17, 3, 2, 2, 2, 145, 146, 7, 43, 2, 2,
	146, 147, 7, 28, 2, 2, 147, 148, 7, 44, 2, 2, 148, 19, 3, 2, 2, 2, 149,
	150, 7, 45, 2, 2, 150, 152, 7, 33, 2, 2, 151, 153, 5, 22, 12, 2, 152, 151,
	3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 155, 7, 34,
	2, 2, 155, 21, 3, 2, 2, 2, 156, 163, 5, 6, 4, 2, 157, 159, 7, 38, 2, 2,
	158, 160, 5, 6, 4, 2, 159, 158, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160,
	162, 3, 2, 2, 2, 161, 157, 3, 2, 2, 2, 162, 165, 3, 2, 2, 2, 163, 161,
	3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 23, 3, 2, 2, 2, 165, 163, 3, 2,
	2, 2, 166, 168, 7, 43, 2, 2, 167, 169, 5, 26, 14, 2, 168, 167, 3, 2, 2,
	2, 169, 170, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171,
	25, 3, 2, 2, 2, 172, 173, 7, 39, 2, 2, 173, 174, 5, 28, 15, 2, 174, 27,
	3, 2, 2, 2, 175, 178, 5, 30, 16, 2, 176, 178, 5, 32, 17, 2, 177, 175, 3,
	2, 2, 2, 177, 176, 3, 2, 2, 2, 178, 29, 3, 2, 2, 2, 179, 180, 7, 44, 2,
	2, 180, 31, 3, 2, 2, 2, 181, 182, 7, 44, 2, 2, 182, 183, 7, 33, 2, 2, 183,
	184, 5, 22, 12, 2, 184, 185, 7, 34, 2, 2, 185, 33, 3, 2, 2, 2, 186, 187,
	9, 2, 2, 2, 187, 35, 3, 2, 2, 2, 14, 40, 44, 72, 113, 115, 128, 141, 152,
	159, 163, 170, 177,
}
var literalNames = []string{
	"", "':='", "'case'", "'endcase'", "'is'", "'switch'", "'endswitch'", "'if'",
	"'then'", "'else'", "'::'", "';'", "'NOT'", "'^'", "'*'", "'/'", "'+'",
	"'-'", "'>'", "'>='", "'<'", "'<='", "'='", "'~='", "'AND'", "'OR'", "':'",
	"'['", "']'", "'{'", "'}'", "'('", "')'", "'_'", "'%'", "'@'", "','", "'.'",
	"'true'", "'false'", "'nil'",
}
var symbolicNames = []string{
	"", "ASSIGNMENT", "CASE", "END_CASE", "IS", "SWITCH", "END_SWITCH", "IF",
	"THEN", "ELSE", "DOUBLE_COLON", "SEMI_COLON", "NOT", "EXPONENTIAL", "MULTIPLY",
	"DIVIDE", "ADD", "MINUS", "GREATER_THAN", "GREATER_THAN_EQUAL", "LESS_THAN",
	"LESS_THAN_EQUAL", "EQUAL", "NOT_EQUAL", "AND", "OR", "COLON", "LBRACKET",
	"RBRACKET", "CURLYLBRACKET", "CURLYRBRACKET", "LPAREN", "RPAREN", "UNDERSCORE",
	"PERCENT_SIGN", "AT_SIGN", "COMMA", "DOT", "KEYWORD_TRUE", "KEYWORD_FALSE",
	"KEYWORD_NIL", "CLASSNAME", "ID", "ATFUNCTION", "INT", "STRING", "FLOAT",
	"PERCENT", "DATE", "NEWLINE", "WS",
}

var ruleNames = []string{
	"program", "statement", "expression", "caseStatement", "caseList", "caseItem",
	"ifStatement", "name", "worksheetVariable", "atFunction", "argList", "dataAccessor",
	"accessorList", "accessorObjectOrArray", "accessorObject", "accessorArray",
	"specialKeyword",
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
	HilcoPencilGrammarParserASSIGNMENT         = 1
	HilcoPencilGrammarParserCASE               = 2
	HilcoPencilGrammarParserEND_CASE           = 3
	HilcoPencilGrammarParserIS                 = 4
	HilcoPencilGrammarParserSWITCH             = 5
	HilcoPencilGrammarParserEND_SWITCH         = 6
	HilcoPencilGrammarParserIF                 = 7
	HilcoPencilGrammarParserTHEN               = 8
	HilcoPencilGrammarParserELSE               = 9
	HilcoPencilGrammarParserDOUBLE_COLON       = 10
	HilcoPencilGrammarParserSEMI_COLON         = 11
	HilcoPencilGrammarParserNOT                = 12
	HilcoPencilGrammarParserEXPONENTIAL        = 13
	HilcoPencilGrammarParserMULTIPLY           = 14
	HilcoPencilGrammarParserDIVIDE             = 15
	HilcoPencilGrammarParserADD                = 16
	HilcoPencilGrammarParserMINUS              = 17
	HilcoPencilGrammarParserGREATER_THAN       = 18
	HilcoPencilGrammarParserGREATER_THAN_EQUAL = 19
	HilcoPencilGrammarParserLESS_THAN          = 20
	HilcoPencilGrammarParserLESS_THAN_EQUAL    = 21
	HilcoPencilGrammarParserEQUAL              = 22
	HilcoPencilGrammarParserNOT_EQUAL          = 23
	HilcoPencilGrammarParserAND                = 24
	HilcoPencilGrammarParserOR                 = 25
	HilcoPencilGrammarParserCOLON              = 26
	HilcoPencilGrammarParserLBRACKET           = 27
	HilcoPencilGrammarParserRBRACKET           = 28
	HilcoPencilGrammarParserCURLYLBRACKET      = 29
	HilcoPencilGrammarParserCURLYRBRACKET      = 30
	HilcoPencilGrammarParserLPAREN             = 31
	HilcoPencilGrammarParserRPAREN             = 32
	HilcoPencilGrammarParserUNDERSCORE         = 33
	HilcoPencilGrammarParserPERCENT_SIGN       = 34
	HilcoPencilGrammarParserAT_SIGN            = 35
	HilcoPencilGrammarParserCOMMA              = 36
	HilcoPencilGrammarParserDOT                = 37
	HilcoPencilGrammarParserKEYWORD_TRUE       = 38
	HilcoPencilGrammarParserKEYWORD_FALSE      = 39
	HilcoPencilGrammarParserKEYWORD_NIL        = 40
	HilcoPencilGrammarParserCLASSNAME          = 41
	HilcoPencilGrammarParserID                 = 42
	HilcoPencilGrammarParserATFUNCTION         = 43
	HilcoPencilGrammarParserINT                = 44
	HilcoPencilGrammarParserSTRING             = 45
	HilcoPencilGrammarParserFLOAT              = 46
	HilcoPencilGrammarParserPERCENT            = 47
	HilcoPencilGrammarParserDATE               = 48
	HilcoPencilGrammarParserNEWLINE            = 49
	HilcoPencilGrammarParserWS                 = 50
)

// HilcoPencilGrammarParser rules.
const (
	HilcoPencilGrammarParserRULE_program               = 0
	HilcoPencilGrammarParserRULE_statement             = 1
	HilcoPencilGrammarParserRULE_expression            = 2
	HilcoPencilGrammarParserRULE_caseStatement         = 3
	HilcoPencilGrammarParserRULE_caseList              = 4
	HilcoPencilGrammarParserRULE_caseItem              = 5
	HilcoPencilGrammarParserRULE_ifStatement           = 6
	HilcoPencilGrammarParserRULE_name                  = 7
	HilcoPencilGrammarParserRULE_worksheetVariable     = 8
	HilcoPencilGrammarParserRULE_atFunction            = 9
	HilcoPencilGrammarParserRULE_argList               = 10
	HilcoPencilGrammarParserRULE_dataAccessor          = 11
	HilcoPencilGrammarParserRULE_accessorList          = 12
	HilcoPencilGrammarParserRULE_accessorObjectOrArray = 13
	HilcoPencilGrammarParserRULE_accessorObject        = 14
	HilcoPencilGrammarParserRULE_accessorArray         = 15
	HilcoPencilGrammarParserRULE_specialKeyword        = 16
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

func (s *ProgramContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
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

	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<HilcoPencilGrammarParserCASE)|(1<<HilcoPencilGrammarParserIF)|(1<<HilcoPencilGrammarParserNOT)|(1<<HilcoPencilGrammarParserMINUS)|(1<<HilcoPencilGrammarParserLPAREN))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(HilcoPencilGrammarParserKEYWORD_TRUE-38))|(1<<(HilcoPencilGrammarParserKEYWORD_FALSE-38))|(1<<(HilcoPencilGrammarParserKEYWORD_NIL-38))|(1<<(HilcoPencilGrammarParserCLASSNAME-38))|(1<<(HilcoPencilGrammarParserID-38))|(1<<(HilcoPencilGrammarParserATFUNCTION-38))|(1<<(HilcoPencilGrammarParserINT-38))|(1<<(HilcoPencilGrammarParserSTRING-38))|(1<<(HilcoPencilGrammarParserFLOAT-38))|(1<<(HilcoPencilGrammarParserPERCENT-38))|(1<<(HilcoPencilGrammarParserDATE-38)))) != 0) {
			{
				p.SetState(35)
				p.Statement()
			}

			p.SetState(40)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)

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
	p.RuleIndex = HilcoPencilGrammarParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *StatementContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserASSIGNMENT, 0)
}

func (s *StatementContext) SEMI_COLON() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserSEMI_COLON, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *HilcoPencilGrammarParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HilcoPencilGrammarParserRULE_statement)

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
		p.SetState(44)
		p.expression(0)
	}
	{
		p.SetState(45)
		p.Match(HilcoPencilGrammarParserASSIGNMENT)
	}
	{
		p.SetState(46)
		p.expression(0)
	}
	{
		p.SetState(47)
		p.Match(HilcoPencilGrammarParserSEMI_COLON)
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

type WorkSheetVariableCalculatorContext struct {
	*ExpressionContext
}

func NewWorkSheetVariableCalculatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WorkSheetVariableCalculatorContext {
	var p = new(WorkSheetVariableCalculatorContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *WorkSheetVariableCalculatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WorkSheetVariableCalculatorContext) WorksheetVariable() IWorksheetVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWorksheetVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWorksheetVariableContext)
}

func (s *WorkSheetVariableCalculatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterWorkSheetVariableCalculator(s)
	}
}

func (s *WorkSheetVariableCalculatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitWorkSheetVariableCalculator(s)
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
	_startState := 4
	p.EnterRecursionRule(localctx, 4, HilcoPencilGrammarParserRULE_expression, _p)

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
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCaseContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(50)
			p.CaseStatement()
		}

	case 2:
		localctx = NewIfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(51)
			p.IfStatement()
		}

	case 3:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(52)
			p.Match(HilcoPencilGrammarParserLPAREN)
		}
		{
			p.SetState(53)
			p.expression(0)
		}
		{
			p.SetState(54)
			p.Match(HilcoPencilGrammarParserRPAREN)
		}

	case 4:
		localctx = NewUnaryMinusCalculatorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(56)
			p.Match(HilcoPencilGrammarParserMINUS)
		}
		{
			p.SetState(57)
			p.expression(25)
		}

	case 5:
		localctx = NewUnaryNotCalculatorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(58)
			p.Match(HilcoPencilGrammarParserNOT)
		}
		{
			p.SetState(59)
			p.expression(24)
		}

	case 6:
		localctx = NewNameCalculatorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(60)
			p.Name()
		}

	case 7:
		localctx = NewWorkSheetVariableCalculatorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(61)
			p.WorksheetVariable()
		}

	case 8:
		localctx = NewExpressionAtFunctionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(62)
			p.AtFunction()
		}

	case 9:
		localctx = NewExpressionDataAccessContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(63)
			p.DataAccessor()
		}

	case 10:
		localctx = NewExpressionKeywordContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(64)
			p.SpecialKeyword()
		}

	case 11:
		localctx = NewPercentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(65)
			p.Match(HilcoPencilGrammarParserPERCENT)
		}

	case 12:
		localctx = NewIntegerContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(66)
			p.Match(HilcoPencilGrammarParserINT)
		}

	case 13:
		localctx = NewFloatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(67)
			p.Match(HilcoPencilGrammarParserFLOAT)
		}

	case 14:
		localctx = NewStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(68)
			p.Match(HilcoPencilGrammarParserSTRING)
		}

	case 15:
		localctx = NewDateContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(69)
			p.Match(HilcoPencilGrammarParserDATE)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExponentialCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(72)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
				}
				{
					p.SetState(73)
					p.Match(HilcoPencilGrammarParserEXPONENTIAL)
				}
				{
					p.SetState(74)
					p.expression(23)
				}

			case 2:
				localctx = NewBinaryArthmeticCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(75)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
				}
				{
					p.SetState(76)
					p.Match(HilcoPencilGrammarParserMULTIPLY)
				}
				{
					p.SetState(77)
					p.expression(23)
				}

			case 3:
				localctx = NewBinaryArthmeticCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(78)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
				}
				{
					p.SetState(79)
					p.Match(HilcoPencilGrammarParserDIVIDE)
				}
				{
					p.SetState(80)
					p.expression(22)
				}

			case 4:
				localctx = NewBinaryArthmeticCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(81)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
				}
				{
					p.SetState(82)
					p.Match(HilcoPencilGrammarParserADD)
				}
				{
					p.SetState(83)
					p.expression(21)
				}

			case 5:
				localctx = NewBinaryArthmeticCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(84)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(85)
					p.Match(HilcoPencilGrammarParserMINUS)
				}
				{
					p.SetState(86)
					p.expression(20)
				}

			case 6:
				localctx = NewBinaryRelationalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(88)
					p.Match(HilcoPencilGrammarParserGREATER_THAN)
				}
				{
					p.SetState(89)
					p.expression(19)
				}

			case 7:
				localctx = NewBinaryRelationalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(91)
					p.Match(HilcoPencilGrammarParserGREATER_THAN_EQUAL)
				}
				{
					p.SetState(92)
					p.expression(18)
				}

			case 8:
				localctx = NewBinaryRelationalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(93)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(94)
					p.Match(HilcoPencilGrammarParserLESS_THAN)
				}
				{
					p.SetState(95)
					p.expression(17)
				}

			case 9:
				localctx = NewBinaryRelationalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(96)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(97)
					p.Match(HilcoPencilGrammarParserLESS_THAN_EQUAL)
				}
				{
					p.SetState(98)
					p.expression(16)
				}

			case 10:
				localctx = NewBinaryRelationalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(99)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(100)
					p.Match(HilcoPencilGrammarParserEQUAL)
				}
				{
					p.SetState(101)
					p.expression(15)
				}

			case 11:
				localctx = NewBinaryRelationalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(102)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(103)
					p.Match(HilcoPencilGrammarParserNOT_EQUAL)
				}
				{
					p.SetState(104)
					p.expression(14)
				}

			case 12:
				localctx = NewBinaryLogicalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(105)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(106)
					p.Match(HilcoPencilGrammarParserAND)
				}
				{
					p.SetState(107)
					p.expression(13)
				}

			case 13:
				localctx = NewBinaryLogicalCalculatorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, HilcoPencilGrammarParserRULE_expression)
				p.SetState(108)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(109)
					p.Match(HilcoPencilGrammarParserOR)
				}
				{
					p.SetState(110)
					p.expression(12)
				}

			}

		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 6, HilcoPencilGrammarParserRULE_caseStatement)

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
		p.SetState(116)
		p.Match(HilcoPencilGrammarParserCASE)
	}
	{
		p.SetState(117)
		p.expression(0)
	}
	{
		p.SetState(118)
		p.Match(HilcoPencilGrammarParserIS)
	}
	{
		p.SetState(119)
		p.CaseList()
	}
	{
		p.SetState(120)
		p.Match(HilcoPencilGrammarParserEND_CASE)
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

	p.SetState(126)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case HilcoPencilGrammarParserCASE, HilcoPencilGrammarParserIF, HilcoPencilGrammarParserNOT, HilcoPencilGrammarParserMINUS, HilcoPencilGrammarParserLPAREN, HilcoPencilGrammarParserKEYWORD_TRUE, HilcoPencilGrammarParserKEYWORD_FALSE, HilcoPencilGrammarParserKEYWORD_NIL, HilcoPencilGrammarParserCLASSNAME, HilcoPencilGrammarParserID, HilcoPencilGrammarParserATFUNCTION, HilcoPencilGrammarParserINT, HilcoPencilGrammarParserSTRING, HilcoPencilGrammarParserFLOAT, HilcoPencilGrammarParserPERCENT, HilcoPencilGrammarParserDATE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.CaseItem()
		}
		{
			p.SetState(123)
			p.CaseList()
		}

	case HilcoPencilGrammarParserEND_CASE:
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
		p.SetState(128)
		p.expression(0)
	}
	{
		p.SetState(129)
		p.Match(HilcoPencilGrammarParserDOUBLE_COLON)
	}
	{
		p.SetState(130)
		p.expression(0)
	}
	{
		p.SetState(131)
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
		p.SetState(133)
		p.Match(HilcoPencilGrammarParserIF)
	}
	{
		p.SetState(134)
		p.expression(0)
	}
	{
		p.SetState(135)
		p.Match(HilcoPencilGrammarParserTHEN)
	}
	{
		p.SetState(136)
		p.expression(0)
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(137)
			p.Match(HilcoPencilGrammarParserELSE)
		}
		{
			p.SetState(138)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(HilcoPencilGrammarParserID)
	}

	return localctx
}

// IWorksheetVariableContext is an interface to support dynamic dispatch.
type IWorksheetVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWorksheetVariableContext differentiates from other interfaces.
	IsWorksheetVariableContext()
}

type WorksheetVariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWorksheetVariableContext() *WorksheetVariableContext {
	var p = new(WorksheetVariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_worksheetVariable
	return p
}

func (*WorksheetVariableContext) IsWorksheetVariableContext() {}

func NewWorksheetVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WorksheetVariableContext {
	var p = new(WorksheetVariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_worksheetVariable

	return p
}

func (s *WorksheetVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *WorksheetVariableContext) CLASSNAME() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserCLASSNAME, 0)
}

func (s *WorksheetVariableContext) COLON() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserCOLON, 0)
}

func (s *WorksheetVariableContext) ID() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserID, 0)
}

func (s *WorksheetVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WorksheetVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WorksheetVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterWorksheetVariable(s)
	}
}

func (s *WorksheetVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitWorksheetVariable(s)
	}
}

func (p *HilcoPencilGrammarParser) WorksheetVariable() (localctx IWorksheetVariableContext) {
	localctx = NewWorksheetVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, HilcoPencilGrammarParserRULE_worksheetVariable)

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
		p.SetState(143)
		p.Match(HilcoPencilGrammarParserCLASSNAME)
	}
	{
		p.SetState(144)
		p.Match(HilcoPencilGrammarParserCOLON)
	}
	{
		p.SetState(145)
		p.Match(HilcoPencilGrammarParserID)
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
	p.EnterRule(localctx, 18, HilcoPencilGrammarParserRULE_atFunction)
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
		p.SetState(147)
		p.Match(HilcoPencilGrammarParserATFUNCTION)
	}
	{
		p.SetState(148)
		p.Match(HilcoPencilGrammarParserLPAREN)
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<HilcoPencilGrammarParserCASE)|(1<<HilcoPencilGrammarParserIF)|(1<<HilcoPencilGrammarParserNOT)|(1<<HilcoPencilGrammarParserMINUS)|(1<<HilcoPencilGrammarParserLPAREN))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(HilcoPencilGrammarParserKEYWORD_TRUE-38))|(1<<(HilcoPencilGrammarParserKEYWORD_FALSE-38))|(1<<(HilcoPencilGrammarParserKEYWORD_NIL-38))|(1<<(HilcoPencilGrammarParserCLASSNAME-38))|(1<<(HilcoPencilGrammarParserID-38))|(1<<(HilcoPencilGrammarParserATFUNCTION-38))|(1<<(HilcoPencilGrammarParserINT-38))|(1<<(HilcoPencilGrammarParserSTRING-38))|(1<<(HilcoPencilGrammarParserFLOAT-38))|(1<<(HilcoPencilGrammarParserPERCENT-38))|(1<<(HilcoPencilGrammarParserDATE-38)))) != 0) {
		{
			p.SetState(149)
			p.ArgList()
		}

	}
	{
		p.SetState(152)
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
	p.EnterRule(localctx, 20, HilcoPencilGrammarParserRULE_argList)
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
		p.SetState(154)
		p.expression(0)
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == HilcoPencilGrammarParserCOMMA {
		{
			p.SetState(155)
			p.Match(HilcoPencilGrammarParserCOMMA)
		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<HilcoPencilGrammarParserCASE)|(1<<HilcoPencilGrammarParserIF)|(1<<HilcoPencilGrammarParserNOT)|(1<<HilcoPencilGrammarParserMINUS)|(1<<HilcoPencilGrammarParserLPAREN))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(HilcoPencilGrammarParserKEYWORD_TRUE-38))|(1<<(HilcoPencilGrammarParserKEYWORD_FALSE-38))|(1<<(HilcoPencilGrammarParserKEYWORD_NIL-38))|(1<<(HilcoPencilGrammarParserCLASSNAME-38))|(1<<(HilcoPencilGrammarParserID-38))|(1<<(HilcoPencilGrammarParserATFUNCTION-38))|(1<<(HilcoPencilGrammarParserINT-38))|(1<<(HilcoPencilGrammarParserSTRING-38))|(1<<(HilcoPencilGrammarParserFLOAT-38))|(1<<(HilcoPencilGrammarParserPERCENT-38))|(1<<(HilcoPencilGrammarParserDATE-38)))) != 0) {
			{
				p.SetState(156)
				p.expression(0)
			}

		}

		p.SetState(163)
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

func (s *DataAccessorContext) AllAccessorList() []IAccessorListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAccessorListContext)(nil)).Elem())
	var tst = make([]IAccessorListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAccessorListContext)
		}
	}

	return tst
}

func (s *DataAccessorContext) AccessorList(i int) IAccessorListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccessorListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAccessorListContext)
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
	p.EnterRule(localctx, 22, HilcoPencilGrammarParserRULE_dataAccessor)

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
		p.SetState(164)
		p.Match(HilcoPencilGrammarParserCLASSNAME)
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(165)
				p.AccessorList()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IAccessorListContext is an interface to support dynamic dispatch.
type IAccessorListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccessorListContext differentiates from other interfaces.
	IsAccessorListContext()
}

type AccessorListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessorListContext() *AccessorListContext {
	var p = new(AccessorListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorList
	return p
}

func (*AccessorListContext) IsAccessorListContext() {}

func NewAccessorListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessorListContext {
	var p = new(AccessorListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorList

	return p
}

func (s *AccessorListContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessorListContext) DOT() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserDOT, 0)
}

func (s *AccessorListContext) AccessorObjectOrArray() IAccessorObjectOrArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccessorObjectOrArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccessorObjectOrArrayContext)
}

func (s *AccessorListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessorListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessorListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterAccessorList(s)
	}
}

func (s *AccessorListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitAccessorList(s)
	}
}

func (p *HilcoPencilGrammarParser) AccessorList() (localctx IAccessorListContext) {
	localctx = NewAccessorListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, HilcoPencilGrammarParserRULE_accessorList)

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
		p.SetState(170)
		p.Match(HilcoPencilGrammarParserDOT)
	}

	{
		p.SetState(171)
		p.AccessorObjectOrArray()
	}

	return localctx
}

// IAccessorObjectOrArrayContext is an interface to support dynamic dispatch.
type IAccessorObjectOrArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccessorObjectOrArrayContext differentiates from other interfaces.
	IsAccessorObjectOrArrayContext()
}

type AccessorObjectOrArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessorObjectOrArrayContext() *AccessorObjectOrArrayContext {
	var p = new(AccessorObjectOrArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorObjectOrArray
	return p
}

func (*AccessorObjectOrArrayContext) IsAccessorObjectOrArrayContext() {}

func NewAccessorObjectOrArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessorObjectOrArrayContext {
	var p = new(AccessorObjectOrArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorObjectOrArray

	return p
}

func (s *AccessorObjectOrArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessorObjectOrArrayContext) AccessorObject() IAccessorObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccessorObjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccessorObjectContext)
}

func (s *AccessorObjectOrArrayContext) AccessorArray() IAccessorArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccessorArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccessorArrayContext)
}

func (s *AccessorObjectOrArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessorObjectOrArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessorObjectOrArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterAccessorObjectOrArray(s)
	}
}

func (s *AccessorObjectOrArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitAccessorObjectOrArray(s)
	}
}

func (p *HilcoPencilGrammarParser) AccessorObjectOrArray() (localctx IAccessorObjectOrArrayContext) {
	localctx = NewAccessorObjectOrArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, HilcoPencilGrammarParserRULE_accessorObjectOrArray)

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

	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)
			p.AccessorObject()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
			p.AccessorArray()
		}

	}

	return localctx
}

// IAccessorObjectContext is an interface to support dynamic dispatch.
type IAccessorObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccessorObjectContext differentiates from other interfaces.
	IsAccessorObjectContext()
}

type AccessorObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessorObjectContext() *AccessorObjectContext {
	var p = new(AccessorObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorObject
	return p
}

func (*AccessorObjectContext) IsAccessorObjectContext() {}

func NewAccessorObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessorObjectContext {
	var p = new(AccessorObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorObject

	return p
}

func (s *AccessorObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessorObjectContext) ID() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserID, 0)
}

func (s *AccessorObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessorObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessorObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterAccessorObject(s)
	}
}

func (s *AccessorObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitAccessorObject(s)
	}
}

func (p *HilcoPencilGrammarParser) AccessorObject() (localctx IAccessorObjectContext) {
	localctx = NewAccessorObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, HilcoPencilGrammarParserRULE_accessorObject)

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
		p.SetState(177)
		p.Match(HilcoPencilGrammarParserID)
	}

	return localctx
}

// IAccessorArrayContext is an interface to support dynamic dispatch.
type IAccessorArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccessorArrayContext differentiates from other interfaces.
	IsAccessorArrayContext()
}

type AccessorArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessorArrayContext() *AccessorArrayContext {
	var p = new(AccessorArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorArray
	return p
}

func (*AccessorArrayContext) IsAccessorArrayContext() {}

func NewAccessorArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessorArrayContext {
	var p = new(AccessorArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HilcoPencilGrammarParserRULE_accessorArray

	return p
}

func (s *AccessorArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessorArrayContext) ID() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserID, 0)
}

func (s *AccessorArrayContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserLPAREN, 0)
}

func (s *AccessorArrayContext) ArgList() IArgListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *AccessorArrayContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(HilcoPencilGrammarParserRPAREN, 0)
}

func (s *AccessorArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessorArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessorArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.EnterAccessorArray(s)
	}
}

func (s *AccessorArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HilcoPencilGrammarParserListener); ok {
		listenerT.ExitAccessorArray(s)
	}
}

func (p *HilcoPencilGrammarParser) AccessorArray() (localctx IAccessorArrayContext) {
	localctx = NewAccessorArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, HilcoPencilGrammarParserRULE_accessorArray)

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
		p.SetState(179)
		p.Match(HilcoPencilGrammarParserID)
	}
	{
		p.SetState(180)
		p.Match(HilcoPencilGrammarParserLPAREN)
	}
	{
		p.SetState(181)
		p.ArgList()
	}
	{
		p.SetState(182)
		p.Match(HilcoPencilGrammarParserRPAREN)
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
	p.EnterRule(localctx, 32, HilcoPencilGrammarParserRULE_specialKeyword)
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
		p.SetState(184)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(HilcoPencilGrammarParserKEYWORD_TRUE-38))|(1<<(HilcoPencilGrammarParserKEYWORD_FALSE-38))|(1<<(HilcoPencilGrammarParserKEYWORD_NIL-38)))) != 0) {
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
	case 2:
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
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 11)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

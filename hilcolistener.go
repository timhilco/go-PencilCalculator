package pencilCalculator

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/timhilco/go-PencilCalculator/parser"
	"github.com/timhilco/go-PencilCalculator/util/logger"
)

type HilcoPencilGrammarParserListener struct {
	*parser.BaseHilcoPencilGrammarParserListener
	stack           []PencilResult
	logger          *logger.HilcoLogger
	lexer           *parser.HilcoPencilGrammarLexer
	currentOperator string
}

func (p *HilcoPencilGrammarParserListener) SetLexer(lexer *parser.HilcoPencilGrammarLexer) {
	p.lexer = lexer
}
func (p *HilcoPencilGrammarParserListener) Result() PencilResult {
	return p.stack[0]
}
func (l *HilcoPencilGrammarParserListener) push(i PencilResult) {
	l.stack = append(l.stack, i)
}

func (l *HilcoPencilGrammarParserListener) pop() PencilResult {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}
func (p *HilcoPencilGrammarParserListener) logStack() {
	for i, pr := range p.stack {
		text := fmt.Sprintf("%d:%s", i, pr.String())
		p.logger.Info(text)
	}

}
func doBinaryArithmatic(left PencilResult, right PencilResult, operator string) (PencilResult, error) {
	var leftNumberINT int64
	var rightNumberINT int64
	var leftNumberFLOAT float64
	var rightNumberFLOAT float64
	binaryType, resultType := determineBinaryOperationType(left, right)

	switch binaryType {
	case LeftIntRightInt:
		leftNumberINT = left.Value.(int64)
		rightNumberINT = right.Value.(int64)
	case LeftIntRightFloat:
		leftNumberFLOAT = float64(left.Value.(int64))
		rightNumberFLOAT = right.Value.(float64)
	case LeftFloatRightInt:
		rightNumberFLOAT = float64(right.Value.(int64))
		leftNumberFLOAT = left.Value.(float64)
	case LeftFloatRightFloat:
		leftNumberFLOAT = left.Value.(float64)
		rightNumberFLOAT = right.Value.(float64)
	case LeftIntRightIntFloat:
		leftNumberFLOAT = float64(left.Value.(int64))
		fi := right.Value.(floatIntegerNumber)
		rightInteger := float64(fi.IntegerValue)
		divisor := math.Pow10(int(fi.Precision))
		rightNumberFLOAT = rightInteger / divisor
	case InvalidTypes:
		return PencilResult{}, fmt.Errorf("invaild binary operation element types")
	default:
		leftNumberINT = 0
		rightNumberINT = 0
	}
	var result interface{}
	returnType := PencilTypeInteger
	switch operator {
	case "ADD":
		switch resultType {
		case "INT":
			result = leftNumberINT + rightNumberINT
		case "FLOAT":
			result = leftNumberFLOAT + rightNumberFLOAT
			returnType = PencilTypeFloat
		case "BOOLEAN":
		case "STRING":
		default:
			panic(fmt.Sprintf("unexpected resultType: %s", operator))

		}
	case "GREATER_THAN":
		switch resultType {
		case "INT":
			result = leftNumberINT > rightNumberINT
			returnType = PencilTypeBoolean
		case "FLOAT":
			result = leftNumberFLOAT > rightNumberFLOAT
			returnType = PencilTypeBoolean
		case "BOOLEAN":
		case "STRING":
		default:
			panic(fmt.Sprintf("unexpected resultType: %s", operator))

		}
	case "EQUAL":
		switch resultType {
		case "INT":
			result = leftNumberINT == rightNumberINT
			returnType = PencilTypeBoolean
		case "FLOAT":
			result = leftNumberFLOAT == rightNumberFLOAT
			returnType = PencilTypeBoolean
		case "BOOLEAN":
		case "STRING":
		default:
			panic(fmt.Sprintf("unexpected resultType: %s", operator))

		}
	default:
		panic(fmt.Sprintf("unexpected operator: %s", operator))
	}
	pr := PencilResult{
		Type:  returnType,
		Value: result}
	return pr, nil
}

type floatIntegerNumber struct {
	IntegerValue int64
	Precision    int64
}

func convertFloatStringToFloatIntegerNumber(floatString string) floatIntegerNumber {

	//negative := strings.ContainsAny(floatString, "-")
	parts := strings.Split(floatString, ".")
	precision := 0
	var decimal int64 = 0
	if len(parts) == 2 {
		precision = len(parts[1])
		decimal, _ = strconv.ParseInt(parts[0], 10, 64)
	}
	val, _ := strconv.ParseInt(parts[0], 10, 64)
	val = (val * 100) + decimal

	return floatIntegerNumber{
		IntegerValue: val,
		Precision:    int64(precision),
	}
}

func Call(f reflect.Value, params []interface{}) (result interface{}, err error) {
	if len(params) != f.Type().NumIn() {
		err = errors.New("the number of params are out of index")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	var res []reflect.Value
	res = f.Call(in)
	result = res[0].Interface()
	return
}

/*
func doBooleanLogic(left PencilResult, right PencilResult, operator string) PencilResult {

	lNum := left.Value.(int64)
	rNum := right.Value.(int64)
	var result bool
	switch operator {
	case "GT":
		result = lNum > rNum

	default:
		panic(fmt.Sprintf("unexpected operator: %s", operator))
	}
	pr := PencilResult{
		Type:  PencilTypeInteger,
		Value: result}
	return pr
}
*/
//
// New Listeners start here
//
func (p *HilcoPencilGrammarParserListener) EnterProgram(ctx *parser.ProgramContext) {
	p.logger = logger.NewMultiWithFile(false)
	p.logger.Info("EnterProgram")
	p.stack = make([]PencilResult, 0)
}

// ExitInteger is called when production Integer is exited.
func (s *HilcoPencilGrammarParserListener) ExitInteger(ctx *parser.IntegerContext) {

	value := ctx.GetText()
	number, _ := strconv.ParseInt(value, 10, 64)
	pr := PencilResult{
		Type:  PencilTypeInteger,
		Value: number,
	}
	s.push(pr)
	s.logger.Info("ExitInteger: " + value)
	s.logStack()
}
func (s *HilcoPencilGrammarParserListener) ExitBinaryArthmeticCalculator(ctx *parser.BinaryArthmeticCalculatorContext) {

	s.logger.Info("ExitBinaryExponentialCalculator")
	right, left := s.pop(), s.pop()
	pr, _ := doBinaryArithmatic(left, right, s.currentOperator)
	s.push(pr)
	s.logStack()
}

// ExitBinaryRelationalCalculator is called when production BinaryRelationalCalculator is exited.
func (s *HilcoPencilGrammarParserListener) ExitBinaryRelationalCalculator(ctx *parser.BinaryRelationalCalculatorContext) {
	s.logger.Info("ExitBinaryRelationalCalculator")
	right, left := s.pop(), s.pop()
	pr, _ := doBinaryArithmatic(left, right, s.currentOperator)
	s.push(pr)
	s.logStack()
}

// ExitIf is called when production If is exited.
func (s *HilcoPencilGrammarParserListener) ExitIf(ctx *parser.IfContext) {
	falseExpression := s.pop()
	trueExpression := s.pop()
	conditionPR := s.pop()
	condition := conditionPR.Value.(bool)
	if condition {
		s.push(trueExpression)
	} else {
		s.push(falseExpression)
	}
	s.logger.Info("ExitIf: ")
	s.logStack()
}

// VisitTerminal is called when a terminal node is visited.

func (p *HilcoPencilGrammarParserListener) VisitTerminal(node antlr.TerminalNode) {
	token := node.GetSymbol()
	symbol := token.GetTokenType()
	name := p.lexer.SymbolicNames[token.GetTokenType()]
	value := node.GetText()
	text := fmt.Sprintf("%s(%d) -> %s", name, symbol, value)
	switch name {

	case "ADD",
		"GREATER_THAN",
		"EQUAL":
		p.currentOperator = name
		p.logger.Info("VisitTermnal <> Binary Operator: " + text)

	default:
		p.logger.Info("VisitTermnal: " + text)
	}
	p.logStack()

}

// EnterAtFunction is called when production atFunction is entered.

func (p *HilcoPencilGrammarParserListener) EnterAtFunction(ctx *parser.AtFunctionContext) {
	p.logger.Info("EnterAtFunction: " + ctx.GetText())
	p.logStack()

}

// ExitAtFunction is called when production atFunction is exited.
func (p *HilcoPencilGrammarParserListener) ExitAtFunction(ctx *parser.AtFunctionContext) {
	p.logger.Info("Enter ExitAtFunction: ")
	p.logStack()
	children := ctx.GetChildren()
	name := children[0].(antlr.TerminalNode).GetText()
	name = name[1:]
	i := children[2].GetChildCount()
	numberOfArgs := (i / 2) + 1
	text := fmt.Sprintf("Calling Function: %s, Arfs= %d", name, numberOfArgs)
	p.logger.Info("ExitAtFunction: " + text)
	f := functionMap[name]
	frv := reflect.ValueOf(f)
	params := make([]interface{}, numberOfArgs)
	j := numberOfArgs - 1
	for i := 0; i < numberOfArgs; i++ {
		r := p.pop()
		params[j] = r
		j = j - 1
	}
	result, _ := Call(frv, params)
	p.push(result.(PencilResult))
	text = result.(PencilResult).String()
	p.logger.Info("ExitAtFunction: " + text)
	p.logger.Info("Exit ExitAtFunction: ")
	p.logStack()
}

package pencilCalculator

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/rs/zerolog"
	"github.com/timhilco/go-PencilCalculator/parser"
	"github.com/timhilco/go-PencilCalculator/util/logger"
)

type dataAccessorElement struct {
	className             string
	accessorFieldElements []accessorElement
}
type accessorElement interface {
	GetDataAccesorValue(level int, structure interface{}, elements []accessorElement) interface{}
}
type accessorObjectElement struct {
	fieldName string
	//elementType string
}
type accessorArrayElement struct {
	fieldName       string
	accessorArgList []accessorArrayElementExpression
}
type accessorArrayElementExpression struct {
	right     PencilResult
	left      PencilResult
	operation PencilResult
}
type HilcoPencilGrammarParserListener struct {
	*parser.BaseHilcoPencilGrammarParserListener
	stack           []PencilResult
	logger          zerolog.Logger
	lexer           *parser.HilcoPencilGrammarLexer
	currentOperator string
	//inputDataObject   map[string]interface{}
	//dataAccessorStack []dataAccessorElement
	inDataAccessor bool
	//inArgList         bool
	inCaseItem        bool
	processingContext context.Context
	inputDataStore    map[string]interface{}
}

func (p *HilcoPencilGrammarParserListener) SetLexer(lexer *parser.HilcoPencilGrammarLexer) {
	p.lexer = lexer
}
func (p *HilcoPencilGrammarParserListener) SetContext(ctx context.Context) {
	p.processingContext = ctx
	inputData := ctx.Value(InputDataContextKey{}).(map[string]string)
	inputDataStore := make(map[string]interface{})
	// Unmarshal or Decode the JSON to the interface.
	for key, value := range inputData {
		var result map[string]interface{}
		jsonObject := value
		json.Unmarshal([]byte(jsonObject), &result)
		inputDataStore[key] = result
	}
	p.inputDataStore = inputDataStore
}

/*
func (p *HilcoPencilGrammarParserListener) SetInputData(jsonObject []byte) {
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal(jsonObject, &result)
	p.inputDataObject = result
}
*/
func (p *HilcoPencilGrammarParserListener) Result() PencilResult {
	return p.stack[0]
}
func (l *HilcoPencilGrammarParserListener) push(i PencilResult) {
	l.stack = append(l.stack, i)
}

/*
func (p *HilcoPencilGrammarParserListener) getTop() PencilResult {
	return p.stack[(len(p.stack) - 1)]
}
*/

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
	if len(p.stack) == 0 {
		p.logger.Info().Msg("Stack is empty")
		return
	}
	for i, pr := range p.stack {
		text := fmt.Sprintf("%d:%s", i, pr.String())
		p.logger.Info().Msg(text)
	}

}
func DoBinaryArithmatic(left PencilResult, right PencilResult, operator string) (PencilResult, error) {
	var leftNumberINT int64
	var rightNumberINT int64
	var leftNumberFLOAT float64
	var rightNumberFLOAT float64
	var leftString string
	var rightString string
	var leftIntFloat floatIntegerNumber
	var rightIntFloat floatIntegerNumber
	binaryType, resultType := determineBinaryOperationType(left, right)

	switch binaryType {
	case LeftIntRightInt:
		leftNumberINT = left.PrValue.(int64)
		rightNumberINT = right.PrValue.(int64)
	case LeftIntRightFloat:
		leftNumberFLOAT = float64(left.PrValue.(int64))
		rightNumberFLOAT = right.PrValue.(float64)
	case LeftFloatRightInt:
		rightNumberFLOAT = float64(right.PrValue.(int64))
		leftNumberFLOAT = left.PrValue.(float64)
	case LeftFloatRightFloat:
		leftNumberFLOAT = left.PrValue.(float64)
		rightNumberFLOAT = right.PrValue.(float64)
	case LeftIntRightIntFloat:
		leftNumberFLOAT = float64(left.PrValue.(int64))
		fi := right.PrValue.(floatIntegerNumber)
		rightInteger := float64(fi.IntegerValue)
		divisor := math.Pow10(int(fi.Precision))
		rightNumberFLOAT = rightInteger / divisor
	case LeftFloatIntRightIntFloat:
		leftIntFloat = left.PrValue.(floatIntegerNumber)
		rightIntFloat = right.PrValue.(floatIntegerNumber)
	case LeftStringRightString:
		leftString = left.PrValue.(string)
		rightString = right.PrValue.(string)
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
		case "INT_FLOAT":
			result = leftIntFloat.Add(rightIntFloat)
			returnType = PencilTypeIntegerFloat
		case "BOOLEAN":
			panic(fmt.Sprintf("Can not add Booleans - unexpected resultType: %s", operator))
		case "STRING":
			result = leftString + rightString
			returnType = PencilTypeString
		default:
			panic(fmt.Sprintf("unexpected resultType: %s", operator))

		}
	case "MINUS":
		switch resultType {
		case "INT":
			result = leftNumberINT - rightNumberINT
		case "FLOAT":
			result = leftNumberFLOAT - rightNumberFLOAT
			returnType = PencilTypeFloat
		case "INT_FLOAT":
			result = leftIntFloat.Subtract(rightIntFloat)
			returnType = PencilTypeIntegerFloat
		case "BOOLEAN",
			"STRING":
			panic(fmt.Sprintf("unexpected resultType: %s", operator))
		default:
			panic(fmt.Sprintf("unexpected resultType: %s", operator))

		}
	case "MULTIPLY":
		switch resultType {
		case "INT":
			result = leftNumberINT * rightNumberINT
		case "FLOAT":
			result = leftNumberFLOAT * rightNumberFLOAT
			returnType = PencilTypeFloat
		case "INT_FLOAT":
			result = leftIntFloat.Multiply(rightIntFloat)
			returnType = PencilTypeIntegerFloat
		case "BOOLEAN",
			"STRING":
			panic(fmt.Sprintf("unexpected resultType: %s", operator))
		default:
			panic(fmt.Sprintf("unexpected resultType: %s", operator))

		}
	case "DIVIDE":
		switch resultType {
		case "INT":
			result = leftNumberINT / rightNumberINT
		case "FLOAT":
			result = leftNumberFLOAT / rightNumberFLOAT
			returnType = PencilTypeFloat
		case "INT_FLOAT":
			result = leftIntFloat.Divide(rightIntFloat)
			returnType = PencilTypeIntegerFloat
		case "BOOLEAN",
			"STRING":
			panic(fmt.Sprintf("unexpected resultType: %s", operator))
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
		case "INT_FLOAT":
			result = leftIntFloat.GreaterThan(rightIntFloat)
			returnType = PencilTypeIntegerFloat
		default:
			panic(fmt.Sprintf("unexpected resultType: %s", operator))

		}
	case "LESS_THAN":
		switch resultType {
		case "INT":
			result = leftNumberINT < rightNumberINT
			returnType = PencilTypeBoolean
		case "FLOAT":
			result = leftNumberFLOAT < rightNumberFLOAT
			returnType = PencilTypeBoolean
		case "BOOLEAN":
		case "STRING":
		case "INT_FLOAT":
			result = leftIntFloat.LessThan(rightIntFloat)
			returnType = PencilTypeIntegerFloat
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
			result = left.PrValue == right.PrValue
			returnType = PencilTypeBoolean
		case "STRING":
			result = leftString == rightString
			returnType = PencilTypeBoolean
		default:
			panic(fmt.Sprintf("unexpected resultType: %s", operator))

		}
	default:
		panic(fmt.Sprintf("unexpected operator: %s", operator))
	}
	pr := PencilResult{
		Type:    returnType,
		PrValue: result}
	return pr, nil
}

func ConvertFloatStringToFloatIntegerNumber(floatString string) (floatIntegerNumber, float64) {

	//negative := strings.ContainsAny(floatString, "-")
	parts := strings.Split(floatString, ".")
	precision := 0
	var decimal int64 = 0
	if len(parts) == 2 {
		precision = len(parts[1])
		sDecimal := parts[1]
		var rounder int64 = 0
		if precision > 6 {
			sDecimal = sDecimal[0:6]
			precision = 6
			rounder, _ = strconv.ParseInt(parts[1][6:7], 10, 64)
			if rounder > 5 {
				rounder = 1
			} else {
				rounder = 0
			}
		}
		decimal, _ = strconv.ParseInt(sDecimal, 10, 64)
		decimal = decimal + rounder
	}
	val, _ := strconv.ParseInt(parts[0], 10, 64)
	power := (int64)(math.Pow10(int(precision)))
	val = (val * power) + decimal
	sDecimal := strconv.Itoa(int(decimal))
	difference := precision - len(sDecimal)
	prefix := ""
	for i := 0; i < difference; i++ {
		prefix = prefix + "0"
	}
	//sDecimal = prefix + sDecimal
	//s := parts[0] + "." + sDecimal
	f64, _ := strconv.ParseFloat(floatString, 64)
	//fmt.Printf("Before Rounding: s=%s -> %.10f\n", s, f64)
	//f64 = math.Round(f64*100) / 100
	//fmt.Printf("After  Rounding: s=%s -> %.10f\n", s, f64)
	return floatIntegerNumber{
		IntegerValue: val,
		Precision:    int64(precision),
	}, f64
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
	//var res []reflect.Value
	res := f.Call(in)
	result = res[0].Interface()
	return
}

/*
func DoBooleanLogic(left PencilResult, right PencilResult, operator string) PencilResult {

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
	p.logger = logger.NewMultiWithFile(false, false, "log.txt")
	loggingLevel := p.processingContext.Value(LoggingLevelContextKey{})
	if loggingLevel != nil {
		level := loggingLevel.(zerolog.Level)
		zerolog.SetGlobalLevel(level)
	}

	p.logger.Info().Msg("EnterProgram")
	p.stack = make([]PencilResult, 0)
}

// ExitInteger is called when production Integer is exited.
func (s *HilcoPencilGrammarParserListener) ExitInteger(ctx *parser.IntegerContext) {
	s.logger.Info().Msg("Enter ExitInteger: ")

	value := ctx.GetText()
	number, _ := strconv.ParseInt(value, 10, 64)
	pr := PencilResult{
		Type:    PencilTypeInteger,
		PrValue: number,
	}
	s.push(pr)
	s.logStack()
	s.logger.Info().Msg("Exit ExitInteger: " + value)
}

// ExitInteger is called when production Integer is exited.
func (s *HilcoPencilGrammarParserListener) ExitString(ctx *parser.StringContext) {
	s.logger.Info().Msg(" Enter ExitString: ")

	value := ctx.GetText()
	value = strings.ReplaceAll(value, "'", "")

	pr := PencilResult{
		Type:    PencilTypeString,
		PrValue: value,
	}
	s.push(pr)
	s.logStack()
	s.logger.Info().Msg(" Exit ExitString: " + value)
}

// ExitDate is called when production Date is exited.
func (s *HilcoPencilGrammarParserListener) ExitDate(ctx *parser.DateContext) {
	s.logger.Info().Msg(" Enter ExitDate: ")

	value := ctx.GetText()
	value = value[1 : len(value)-1]
	parts := strings.Split(value, "-")
	sDate := parts[2] + "-" + parts[1] + "-" + parts[0]

	date, _ := time.Parse("2006-01-02", sDate)
	//fmt.Printf("%s -> %v\n", date, err)
	pr := PencilResult{
		Type:    PencilTypeDateTime,
		PrValue: date,
	}
	s.push(pr)
	s.logStack()
	s.logger.Info().Msg(" Exit ExitDate: " + value)
}
func (s *HilcoPencilGrammarParserListener) ExitFloat(ctx *parser.FloatContext) {
	s.logger.Info().Msg(" Enter ExitFloat: ")

	value := ctx.GetText()
	_, f := ConvertFloatStringToFloatIntegerNumber(value)
	pr := PencilResult{
		Type:    PencilTypeFloat,
		PrValue: f,
	}
	s.push(pr)
	s.logStack()
	s.logger.Info().Msg(" Exit ExitFloat: " + value)
}

// EnterNameCalculator is called when production NameCalculator is entered.
//func (s *BaseHilcoPencilGrammarParserListener) EnterNameCalculator(ctx *NameCalculatorContext) {}

// ExitNameCalculator is called when production NameCalculator is exited.
func (s *HilcoPencilGrammarParserListener) ExitNameCalculator(ctx *parser.NameCalculatorContext) {
	s.logger.Info().Msg(" Enter ExitNameCalculator: ")

	value := ctx.GetText()
	/*
		pr := PencilResult{
			Type:  PencilTypeNameCalculator,
			PrValue: value,
		}
		s.push(pr)
		s.logStack()
	*/
	s.logger.Info().Msg(" Exit ExitNameCalculator: " + value)
}

func (s *HilcoPencilGrammarParserListener) ExitBinaryArthmeticCalculator(ctx *parser.BinaryArthmeticCalculatorContext) {

	s.logger.Info().Msg(" Enter ExitBinaryArthmeticCalculator")
	right, left := s.pop(), s.pop()
	pr, _ := DoBinaryArithmatic(left, right, s.currentOperator)
	s.push(pr)
	s.logStack()
	s.logger.Info().Msg(" Exit ExitBinaryArthmeticCalculator")
}

// ExitBinaryRelationalCalculator is called when production BinaryRelationalCalculator is exited.
func (s *HilcoPencilGrammarParserListener) ExitBinaryRelationalCalculator(ctx *parser.BinaryRelationalCalculatorContext) {
	s.logger.Info().Msg("Enter ExitBinaryRelationalCalculator")
	if s.inDataAccessor {
		pr := PencilResult{
			Type:    PencilTypeOperation,
			PrValue: s.currentOperator,
		}
		s.push(pr)
		s.logStack()
		s.logger.Info().Msg("Exit ExitBinaryRelationalCalculator: Exit because in DataAccessor")
		return

	}
	right, left := s.pop(), s.pop()
	pr, _ := DoBinaryArithmatic(left, right, s.currentOperator)
	s.push(pr)
	s.logStack()
	s.logger.Info().Msg("Exit ExitBinaryRelationalCalculator")
}

// EnterBinaryLogicalCalculator is called when production BinaryLogicalCalculator is entered.
//func ( *HilcoPencilGrammarParserListener) EnterBinaryLogicalCalculator(ctx *parser.BinaryLogicalCalculatorContext) {
//}

// ExitBinaryLogicalCalculator is called when production BinaryLogicalCalculator is exited.
func (p *HilcoPencilGrammarParserListener) ExitBinaryLogicalCalculator(ctx *parser.BinaryLogicalCalculatorContext) {
	p.logger.Info().Msg("Enter ExitBinaryLogicalCalculator")
	right, operation, left := p.pop(), p.pop(), p.pop()
	binaryType, _ := determineBinaryOperationType(left, right)
	pr := PencilResult{}
	switch binaryType {
	case LeftBooleanRightBoolean:

		leftBoolean := left.PrValue.(bool)
		rightBoolean := right.PrValue.(bool)
		var resultValue bool
		switch operation.PrValue {
		case "AND":
			resultValue = leftBoolean && rightBoolean
		case "OR":
			resultValue = leftBoolean || rightBoolean
		default:
		}
		pr.Type = PencilTypeBoolean
		pr.PrValue = resultValue
	default:
	}
	p.push(pr)
	p.logStack()
	p.logger.Info().Msg("Exit ExitBinaryLogicalCalculator")
}

// ExitIf is called when production If is exited.
func (s *HilcoPencilGrammarParserListener) ExitIf(ctx *parser.IfContext) {
	falseExpression := s.pop()
	trueExpression := s.pop()
	conditionPR := s.pop()
	condition := conditionPR.PrValue.(bool)
	if condition {
		s.push(trueExpression)
	} else {
		s.push(falseExpression)
	}
	s.logger.Info().Msg("ExitIf: ")
	s.logStack()
}

// EnterCaseStatement is called when production caseStatement is entered.
func (p *HilcoPencilGrammarParserListener) EnterCaseStatement(ctx *parser.CaseStatementContext) {
	p.logger.Info().Msg("Enter EnterCaseStatement ")
	/*
		pr := PencilResult {
			Type: PencilTypeCaseStatement
			PrValue: "CASE STATEMENT"
		}
		p.logStack()
	*/
	p.logger.Info().Msg("Exit EnterCaseStatement ")

}
func (cs CaseStatement) executeCaseStatement() PencilResult {
	var hitMatchingCondition bool = false
	pr := PencilResult{}
	for _, ci := range cs.caseItems {
		c1 := cs.expressionPencilResultValue
		c2 := ci.matchValue
		br, _ := DoBinaryArithmatic(c1, c2, "EQUAL")
		b := br.PrValue.(bool)
		if b {
			hitMatchingCondition = true
			pr = ci.resultValue
			break
		}

	}
	if !hitMatchingCondition {
		pr = cs.defaultCaseItem.resultValue
	}
	return pr
}

// ExitCaseStatement is called when production caseStatement is exited.
func (p *HilcoPencilGrammarParserListener) ExitCaseStatement(ctx *parser.CaseStatementContext) {
	p.logger.Info().Msg("Enter ExitCaseStatement ")
	// Build Case Statement and determine value
	caseStatement := CaseStatement{}
	hitCaseExpression := false
	var pr PencilResult
	for !hitCaseExpression {
		pr = p.pop()
		switch pr.Type {
		case PencilTypeCaseItem:
			ci := pr.PrValue.(CaseItem)
			mv := ci.matchValue.PrValue.(string)
			if mv == "default" {
				caseStatement.defaultCaseItem = ci
			} else {
				caseStatement.caseItems = append(caseStatement.caseItems, ci)
			}
		default:
			caseStatement.expressionPencilResultValue = pr
			hitCaseExpression = true

		}

	}
	newResult := caseStatement.executeCaseStatement()
	p.push(newResult)
	p.logStack()
	p.logger.Info().Msg("Exit ExitCaseStatement ")

}

// EnterCaseItem is called when production caseItem is entered.
func (p *HilcoPencilGrammarParserListener) EnterCaseItem(ctx *parser.CaseItemContext) {
	p.logger.Info().Msg("Enter EnterCaseItem ")
	p.inCaseItem = true
	p.logStack()
	p.logger.Info().Msg("Exit EnterCaseItem ")

}

// ExitCaseItem is called when production caseItem is exited.
func (p *HilcoPencilGrammarParserListener) ExitCaseItem(ctx *parser.CaseItemContext) {
	p.logger.Info().Msg("Enter ExitCaseItem ")
	resultValue, matchValue := p.pop(), p.pop()
	pr := PencilResult{
		Type: PencilTypeCaseItem,
		PrValue: CaseItem{
			matchValue:  matchValue,
			resultValue: resultValue,
		},
	}
	p.push(pr)
	p.inCaseItem = false
	p.logStack()
	p.logger.Info().Msg("Exit ExitCaseItem ")

}

// EnterAtFunction is called when production atFunction is entered.

func (p *HilcoPencilGrammarParserListener) EnterAtFunction(ctx *parser.AtFunctionContext) {
	p.logger.Info().Msg("EnterAtFunction: " + ctx.GetText())
	p.logStack()

}

// ExitAtFunction is called when production atFunction is exited.
func (p *HilcoPencilGrammarParserListener) ExitAtFunction(ctx *parser.AtFunctionContext) {
	p.logger.Info().Msg("Enter ExitAtFunction: ")
	p.logStack()
	children := ctx.GetChildren()
	name := children[0].(antlr.TerminalNode).GetText()
	name = name[1:]
	args := make([]PencilResult, 0)
	foundLeftParen := false
	for !foundLeftParen {

		foundComma := false
		for !foundComma {
			pr := p.pop()
			switch pr.Type {
			case PencilTypeLeftParen:
				foundLeftParen = true
				foundComma = true
			case PencilTypeComma:
				foundComma = true
			default:
				args = append(args, pr)

			}
		}

	}
	numberOfArgs := len(args)
	text := fmt.Sprintf("Calling Function: %s, Arfs= %d", name, numberOfArgs)
	p.logger.Info().Msg("ExitAtFunction: " + text)
	f := functionMap[name]
	frv := reflect.ValueOf(f)
	params := make([]interface{}, numberOfArgs)
	j := numberOfArgs - 1
	for i := 0; i < numberOfArgs; i++ {
		params[j] = args[i]
		j = j - 1
	}
	result, _ := Call(frv, params)
	p.push(result.(PencilResult))
	text = result.(PencilResult).String()
	p.logger.Info().Msg("ExitAtFunction: " + text)
	p.logStack()
	p.logger.Info().Msg("Exit ExitAtFunction: ")
}

// EnterDataAccessor is called when production dataAccessor is entered.
func (p *HilcoPencilGrammarParserListener) EnterDataAccessor(ctx *parser.DataAccessorContext) {
	p.logger.Info().Msg("Enter EnterDataAccessor: ")
	da := dataAccessorElement{
		className:             "No Name set",
		accessorFieldElements: make([]accessorElement, 0),
	}

	p.inDataAccessor = true
	pr := PencilResult{
		Type:    PencilTypeDataAccessor,
		PrValue: da,
	}
	p.push(pr)
	p.logStack()

	p.logger.Info().Msg("Exit EnterDataAccessor: ")
}

// ExitDataAccessor is called when production dataAccessor is exited.
func (p *HilcoPencilGrammarParserListener) ExitDataAccessor(ctx *parser.DataAccessorContext) {
	p.logger.Info().Msg("Enter ExitDataAccessor: ")
	dataAccessorElements := make([]accessorElement, 0)
	foundDataAccessor := false
	var da dataAccessorElement
	for !foundDataAccessor {
		pr := p.pop()
		switch pr.Type {
		case PencilTypeDataAccessor:
			da = pr.PrValue.(dataAccessorElement)
			foundDataAccessor = true

		case PencilTypeDataAccessorElement:
			dae := pr.PrValue.(accessorElement)
			dataAccessorElements = append(dataAccessorElements, dae)
		default:
		}

	}
	lenx := len(dataAccessorElements)               // lenx holds the original array length
	reversed_array := make([]accessorElement, lenx) // creates a slice that refer to a new array of length lenx

	for i := 0; i < lenx; i++ {
		j := lenx - (i + 1) // j initially holds (lenx - 1) and decreases to 0 while i initially holds 0 and increase to (lenx - 1)
		reversed_array[i] = dataAccessorElements[j]
	}
	da.accessorFieldElements = reversed_array

	key := da.className
	dataInputObjectMap := p.inputDataStore
	dataInputObject := dataInputObjectMap[key]

	afe := da.accessorFieldElements[0]
	value := afe.GetDataAccesorValue(0, dataInputObject, da.accessorFieldElements)

	rv := reflect.ValueOf(value).Kind()
	var resultType PencilType
	switch rv {
	case reflect.String:
		resultType = PencilTypeString
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		resultType = PencilTypeInteger
	case reflect.Float32, reflect.Float64:
		resultType = PencilTypeFloat

	default:
		fmt.Printf("unhandled kind %s", rv)
	}

	pr := PencilResult{
		Type:    resultType,
		PrValue: value,
	}
	p.push(pr)
	p.inDataAccessor = false
	p.logStack()

	p.logger.Info().Msg("Exit ExitDataAccessor: ")

}
func (a accessorObjectElement) GetDataAccesorValue(level int, structure interface{}, elements []accessorElement) interface{} {

	var result interface{}
	//text := fmt.Sprintf("aStrintf:%d: %s -> %s", level, a, elements[level])
	//fmt.Println(text)
	name := a.fieldName
	value := structure.(map[string]interface{})[name]
	nextLevel := level + 1
	if nextLevel == len(elements) {
		result = value
	} else {
		e := elements[nextLevel]
		result = e.GetDataAccesorValue(nextLevel, value, elements)
	}
	return result
}
func (a accessorArrayElement) GetDataAccesorValue(level int, structure interface{}, elements []accessorElement) interface{} {

	var result interface{}
	//text := fmt.Sprintf("aStrintf:%d: %s -> %s", level, a, elements[level])
	//fmt.Println(text)
	name := a.fieldName
	value := structure.(map[string]interface{})[name]
	anArray := value.([]interface{})
	var item interface{}
	for _, v := range anArray {
		item = v
	}

	nextLevel := level + 1
	if nextLevel == len(elements) {
		result = item
	} else {
		e := elements[nextLevel]
		result = e.GetDataAccesorValue(nextLevel, item, elements)
	}

	return result

}

/*
// EnterAccessorObjectOrArray is called when production accessorObjectOrArray is entered.
func (p *HilcoPencilGrammarParserListener) EnterAccessorObjectOrArray(ctx *parser.AccessorObjectOrArrayContext) {
	p.logger.Info().Msg("Enter EnterAccessorObjectOrArray: ")
	p.logger.Info().Msg("Exit EnterAccessorObjectOrArray: ")
}
*/
/*
// ExitAccessorObjectOrArray is called when production accessorObjectOrArray is exited.
func (p *HilcoPencilGrammarParserListener) ExitAccessorObjectOrArray(ctx *parser.AccessorObjectOrArrayContext) {
	p.logger.Info().Msg("Enter ExitAccessorObjectOrArray: ")
	p.logger.Info().Msg("Exit ExitAccessorObjectOrArray: ")
}

// EnterAccessorObject is called when production accessorObject is entered.
func (p *HilcoPencilGrammarParserListener) EnterAccessorObject(ctx *parser.AccessorObjectContext) {
	p.logger.Info().Msg("Enter EnterAccessorObjec: ")
	p.logger.Info().Msg("Exit EnterAccessorObject: ")

}
*/
// ExitAccessorObject is called when production accessorObject is exited.
func (p *HilcoPencilGrammarParserListener) ExitAccessorObject(ctx *parser.AccessorObjectContext) {
	p.logger.Info().Msg("Enter ExitAccessorObject: ")
	nc := p.pop() // Should be Name Calculator
	fieldName := nc.PrValue.(string)
	ao := accessorObjectElement{
		fieldName: fieldName,
	}
	pr := PencilResult{
		Type:    PencilTypeDataAccessorElement,
		PrValue: ao,
	}
	p.push(pr)
	p.logStack()
	p.logger.Info().Msg("Exit ExitAccessorObject: ")

}

/*
// EnterAccessorArray is called when production accessorArray is entered.
func (p *HilcoPencilGrammarParserListener) EnterAccessorArray(ctx *parser.AccessorArrayContext) {
	p.logger.Info().Msg("Enter EnterAccessArray: ")
	p.logger.Info().Msg("Exit EnterAccessorArray: ")

}
*/
// ExitAccessorArray is called when production accessorArray is exited.
func (p *HilcoPencilGrammarParserListener) ExitAccessorArray(ctx *parser.AccessorArrayContext) {
	p.logger.Info().Msg("Enter ExitAccessArray: ")
	p.logStack()

	foundLeftParen := false
	expressionArray := make([]accessorArrayElementExpression, 0)
	for !foundLeftParen {
		counter := 0
		expression := accessorArrayElementExpression{}
		foundComma := false
		for !foundComma {
			pr := p.pop()
			switch pr.Type {
			case PencilTypeLeftParen:
				foundLeftParen = true
				foundComma = true
			case PencilTypeComma:
				foundComma = true
			default:
				switch counter {
				case 0:
					expression.operation = pr
				case 1:
					expression.right = pr
				case 2:
					expression.left = pr
				}
				counter++

			}
		}
		expressionArray = append(expressionArray, expression)

	}
	npr := p.pop()
	a := accessorArrayElement{
		fieldName:       npr.PrValue.(string),
		accessorArgList: expressionArray,
	}
	npr = PencilResult{
		Type:    PencilTypeDataAccessorElement,
		PrValue: a,
	}
	p.push(npr)
	p.logStack()
	p.logger.Info().Msg("Exit ExitAccessorArray: ")

}

/* This is old stuff from the V1 version of the parser
// EnterAccessorMessage is called when production accessorMessage is entered.
func (p *HilcoPencilGrammarParserListener) EnterAccessorMessage(ctx *parser.AccessorMessageContext) {
	p.logger.Info().Msg("Enter EnterAccessorMessage: ")
	p.inAccessorMessage = true
	pr := PencilResult{
		Type:  PencilTypeAccessorMessage,
		PrValue: nil,
	}
	p.push(pr)
	p.logStack()
	p.logger.Info().Msg("Exit EnterAccessorMessage: ")
}

// ExitAccessorMessage is called when production accessorMessage is exited.
func (p *HilcoPencilGrammarParserListener) ExitAccessorMessage(ctx *parser.AccessorMessageContext) {
	p.logger.Info().Msg("Enter ExitAccessorMessage: ")
	p.inAccessorMessage = false

	p.logger.Info().Msg("Exit ExitAccessorMessage: ")

}
*/

// EnterArgList is called when production argList is entered.
/*
func (p *HilcoPencilGrammarParserListener) EnterArgList(ctx *parser.ArgListContext) {
	p.logger.Info().Msg("Enter EnterArgList: ")
	p.logger.Info().Msg("Exit EnterArgList: ")

}
*/
// ExitArgList is called when production argList is exited.
/*
func (p *HilcoPencilGrammarParserListener) ExitArgList(ctx *parser.ArgListContext) {
	p.logger.Info().Msg("Enter ExitArgList: ")
	p.logger.Info().Msg("Exit ExitArgList: ")

}
*/
// VisitTerminal is called when a terminal node is visited.
func (p *HilcoPencilGrammarParserListener) VisitTerminal(node antlr.TerminalNode) {
	p.logger.Info().Msg("Enter VisitTerminal: " + node.GetText())
	token := node.GetSymbol()
	symbol := token.GetTokenType()
	name := p.lexer.SymbolicNames[token.GetTokenType()]
	value := node.GetText()
	text := fmt.Sprintf("%s(%d) -> %s", name, symbol, value)
	switch name {

	case "AND",
		"OR":
		p.logger.Info().Msg("VisitTermnal <> Binary Operator: " + text)
		pr := PencilResult{
			Type:    PencilTypeOperation,
			PrValue: name,
		}
		p.push(pr)
	case "ADD",
		"MINUS",
		"MULTIPLY",
		"DIVIDE",
		"LESS_THAN",
		"GREATER_THAN",
		"EQUAL":

		p.currentOperator = name
		p.logger.Info().Msg("VisitTerminal <> Binary Operator: " + text)
	case "CLASSNAME":
		if p.inDataAccessor {
			pr := p.pop()
			dae := pr.PrValue.(dataAccessorElement)
			dae.className = value
			pr = PencilResult{
				Type:    PencilTypeDataAccessor,
				PrValue: dae,
			}
			p.push(pr)
		}
	case "ID":
		if p.inDataAccessor {

			pr := PencilResult{
				Type:    PencilTypeNameCalculator,
				PrValue: value,
			}
			p.push(pr)

		}
	case "COMMA":
		pr := PencilResult{
			Type:    PencilTypeComma,
			PrValue: name,
		}
		p.push(pr)
	case "LPAREN":
		pr := PencilResult{
			Type:    PencilTypeLeftParen,
			PrValue: name,
		}
		p.push(pr)
	default:
		p.logger.Info().Msg("VisitTerminal <> Ignoring Terminal:  " + name)
	}
	p.logStack()
	p.logger.Info().Msg("Exit VisitTerminal: " + text)

}
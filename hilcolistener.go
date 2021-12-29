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
	elementName string
	elementType string
}
type HilcoPencilGrammarParserListener struct {
	*parser.BaseHilcoPencilGrammarParserListener
	stack           []PencilResult
	logger          *logger.HilcoLogger
	lexer           *parser.HilcoPencilGrammarLexer
	currentOperator string
	//inputDataObject   map[string]interface{}
	dataAccessorStack []dataAccessorElement
	inDataAccessor    bool
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
		p.logger.Info("Stack is empty")
		return
	}
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
	var leftString string
	var rightString string
	var leftIntFloat floatIntegerNumber
	var rightIntFloat floatIntegerNumber
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
	case LeftFloatIntRightIntFloat:
		leftIntFloat = left.Value.(floatIntegerNumber)
		rightIntFloat = right.Value.(floatIntegerNumber)
	case LeftStringRightString:
		leftString = left.Value.(string)
		rightString = right.Value.(string)
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
			result = left.Value == right.Value
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
		Type:  returnType,
		Value: result}
	return pr, nil
}

func convertFloatStringToFloatIntegerNumber(floatString string) (floatIntegerNumber, float64) {

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
	sDecimal = prefix + sDecimal
	s := parts[0] + "." + sDecimal
	f64, _ := strconv.ParseFloat(s, 64)
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
	loggingLevel := p.processingContext.Value(LoggingLevelContextKey{})
	if loggingLevel != nil {
		level := loggingLevel.(zerolog.Level)
		zerolog.SetGlobalLevel(level)
	}

	p.logger.Info("EnterProgram")
	p.stack = make([]PencilResult, 0)
}

// ExitInteger is called when production Integer is exited.
func (s *HilcoPencilGrammarParserListener) ExitInteger(ctx *parser.IntegerContext) {
	s.logger.Info("Enter ExitInteger: ")

	value := ctx.GetText()
	number, _ := strconv.ParseInt(value, 10, 64)
	pr := PencilResult{
		Type:  PencilTypeInteger,
		Value: number,
	}
	s.push(pr)
	s.logStack()
	s.logger.Info("Exit ExitInteger: " + value)
}

// ExitInteger is called when production Integer is exited.
func (s *HilcoPencilGrammarParserListener) ExitString(ctx *parser.StringContext) {
	s.logger.Info(" Enter ExitString: ")

	value := ctx.GetText()
	value = strings.ReplaceAll(value, "'", "")

	pr := PencilResult{
		Type:  PencilTypeString,
		Value: value,
	}
	s.push(pr)
	s.logStack()
	s.logger.Info(" Exit ExitString: " + value)
}

// ExitDate is called when production Date is exited.
func (s *HilcoPencilGrammarParserListener) ExitDate(ctx *parser.DateContext) {
	s.logger.Info(" Enter ExitDate: ")

	value := ctx.GetText()
	value = value[1 : len(value)-1]
	parts := strings.Split(value, "-")
	sDate := parts[2] + "-" + parts[1] + "-" + parts[0]

	date, err := time.Parse("2006-01-02", sDate)
	fmt.Printf("%s -> %v", date, err)
	pr := PencilResult{
		Type:  PencilTypeDateTime,
		Value: date,
	}
	s.push(pr)
	s.logStack()
	s.logger.Info(" Exit ExitDate: " + value)
}
func (s *HilcoPencilGrammarParserListener) ExitFloat(ctx *parser.FloatContext) {
	s.logger.Info(" Enter ExitFloat: ")

	value := ctx.GetText()
	f, _ := convertFloatStringToFloatIntegerNumber(value)
	pr := PencilResult{
		Type:  PencilTypeIntegerFloat,
		Value: f,
	}
	s.push(pr)
	s.logStack()
	s.logger.Info(" Exit ExitFloat: " + value)
}
func (s *HilcoPencilGrammarParserListener) ExitBinaryArthmeticCalculator(ctx *parser.BinaryArthmeticCalculatorContext) {

	s.logger.Info(" Enter ExitBinaryArthmeticCalculator")
	right, left := s.pop(), s.pop()
	pr, _ := doBinaryArithmatic(left, right, s.currentOperator)
	s.push(pr)
	s.logStack()
	s.logger.Info(" Exit ExitBinaryArthmeticCalculator")
}

// ExitBinaryRelationalCalculator is called when production BinaryRelationalCalculator is exited.
func (s *HilcoPencilGrammarParserListener) ExitBinaryRelationalCalculator(ctx *parser.BinaryRelationalCalculatorContext) {
	s.logger.Info("ExitBinaryRelationalCalculator")
	right, left := s.pop(), s.pop()
	pr, _ := doBinaryArithmatic(left, right, s.currentOperator)
	s.push(pr)
	s.logStack()
}

// EnterBinaryLogicalCalculator is called when production BinaryLogicalCalculator is entered.
//func ( *HilcoPencilGrammarParserListener) EnterBinaryLogicalCalculator(ctx *parser.BinaryLogicalCalculatorContext) {
//}

// ExitBinaryLogicalCalculator is called when production BinaryLogicalCalculator is exited.
func (p *HilcoPencilGrammarParserListener) ExitBinaryLogicalCalculator(ctx *parser.BinaryLogicalCalculatorContext) {
	p.logger.Info("Enter ExitBinaryLogicalCalculator")
	right, operation, left := p.pop(), p.pop(), p.pop()
	binaryType, _ := determineBinaryOperationType(left, right)
	pr := PencilResult{}
	switch binaryType {
	case LeftBooleanRightBoolean:

		leftBoolean := left.Value.(bool)
		rightBoolean := right.Value.(bool)
		var resultValue bool
		switch operation.Value {
		case "AND":
			resultValue = leftBoolean && rightBoolean
		case "OR":
			resultValue = leftBoolean || rightBoolean
		default:
		}
		pr.Type = PencilTypeBoolean
		pr.Value = resultValue
	default:
	}
	p.push(pr)
	p.logStack()
	p.logger.Info("Exit ExitBinaryLogicalCalculator")
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

// EnterCaseStatement is called when production caseStatement is entered.
func (p *HilcoPencilGrammarParserListener) EnterCaseStatement(ctx *parser.CaseStatementContext) {
	p.logger.Info("Enter EnterCaseStatement ")
	/*
		pr := PencilResult {
			Type: PencilTypeCaseStatement
			Value: "CASE STATEMENT"
		}
		p.logStack()
	*/
	p.logger.Info("Exit EnterCaseStatement ")

}
func (cs CaseStatement) executeCaseStatement() PencilResult {
	var hitMatchingCondition bool = false
	pr := PencilResult{}
	for _, ci := range cs.caseItems {
		c1 := cs.expressionPencilResultValue
		c2 := ci.matchValue
		br, _ := doBinaryArithmatic(c1, c2, "EQUAL")
		b := br.Value.(bool)
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
	p.logger.Info("Enter ExitCaseStatement ")
	// Build Case Statement and determine value
	caseStatement := CaseStatement{}
	hitCaseExpression := false
	var pr PencilResult
	for !hitCaseExpression {
		pr = p.pop()
		switch pr.Type {
		case PencilTypeCaseItem:
			ci := pr.Value.(CaseItem)
			mv := ci.matchValue.Value.(string)
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
	p.logger.Info("Exit ExitCaseStatement ")

}

// EnterCaseItem is called when production caseItem is entered.
func (p *HilcoPencilGrammarParserListener) EnterCaseItem(ctx *parser.CaseItemContext) {
	p.logger.Info("Enter EnterCaseItem ")
	p.inCaseItem = true
	p.logStack()
	p.logger.Info("Exit EnterCaseItem ")

}

// ExitCaseItem is called when production caseItem is exited.
func (p *HilcoPencilGrammarParserListener) ExitCaseItem(ctx *parser.CaseItemContext) {
	p.logger.Info("Enter ExitCaseItem ")
	resultValue, matchValue := p.pop(), p.pop()
	pr := PencilResult{
		Type: PencilTypeCaseItem,
		Value: CaseItem{
			matchValue:  matchValue,
			resultValue: resultValue,
		},
	}
	p.push(pr)
	p.inCaseItem = false
	p.logStack()
	p.logger.Info("Exit ExitCaseItem ")

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
	p.logStack()
	p.logger.Info("Exit ExitAtFunction: ")
}

// EnterDataAccessor is called when production dataAccessor is entered.
func (p *HilcoPencilGrammarParserListener) EnterDataAccessor(ctx *parser.DataAccessorContext) {
	p.logger.Info("Enter EnterDataAccessor: ")
	p.dataAccessorStack = make([]dataAccessorElement, 0)
	p.inDataAccessor = true
	p.logStack()
	p.logger.Info("Exit EnterDataAccessor: ")
}

// ExitDataAccessor is called when production dataAccessor is exited.
func (p *HilcoPencilGrammarParserListener) ExitDataAccessor(ctx *parser.DataAccessorContext) {
	p.logger.Info("Enter ExitDataAccessor: ")
	key := p.dataAccessorStack[0].elementName
	dataInputObject := p.inputDataStore[key]
	value := GetDataAccesorValue(1, dataInputObject, p.dataAccessorStack)
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
	//fmt.Println(rv)
	pr := PencilResult{
		Type:  resultType,
		Value: value,
	}
	p.push(pr)
	p.inDataAccessor = false
	p.logStack()
	p.logger.Info("Exit ExitDataAccessor: ")

}
func GetDataAccesorValue(level int, structure interface{}, elements []dataAccessorElement) interface{} {
	var result interface{}
	if level == len(elements) {
		result = structure
	} else {
		//text := fmt.Sprintf("aStrint:%d -> %s", level, elements[level])
		key := elements[level].elementName
		value := structure.(map[string]interface{})[key]
		result = GetDataAccesorValue(level+1, value, elements)
	}
	return result
}

// VisitTerminal is called when a terminal node is visited.
func (p *HilcoPencilGrammarParserListener) VisitTerminal(node antlr.TerminalNode) {
	p.logger.Info("Enter VisitTerminal: " + node.GetText())
	token := node.GetSymbol()
	symbol := token.GetTokenType()
	name := p.lexer.SymbolicNames[token.GetTokenType()]
	value := node.GetText()
	text := fmt.Sprintf("%s(%d) -> %s", name, symbol, value)
	switch name {

	case "AND",
		"OR":
		p.logger.Info("VisitTermnal <> Binary Operator: " + text)
		pr := PencilResult{
			Type:  PencilTypeOperation,
			Value: name,
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
		p.logger.Info("VisitTerminal <> Binary Operator: " + text)
	case "CLASSNAME":
		if p.inDataAccessor {
			dae := dataAccessorElement{
				elementName: value,
				elementType: "CLASSNAME",
			}
			p.dataAccessorStack = append(p.dataAccessorStack, dae)
		}
	case "ID":
		if p.inDataAccessor {
			dae := dataAccessorElement{
				elementName: value,
				elementType: "ID",
			}
			p.dataAccessorStack = append(p.dataAccessorStack, dae)
		}
	default:
		p.logger.Info("VisitTerminal <> Ignoring Terminal:  " + name)
	}
	p.logStack()
	p.logger.Info("Exit VisitTerminal: " + text)

}

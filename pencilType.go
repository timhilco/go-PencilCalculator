package pencilCalculator

import (
	"fmt"
	"math"
)

type PencilType int
type InputDataContextKey struct{}

const (
	PencilTypeString PencilType = iota
	PencilTypeInteger
	PencilTypeFloat
	PencilTypeIntegerFloat
	PencilTypeBoolean
	PencilTypeDateTime
	PencilTypeSlice
	PencilTypeArray
	PencilTypeMap
	PencilTypeStruct
	PencilTypeOperation
	PencilTypeCaseStatement
	PencilTypeCaseItem
)

type BinaryElementType int

const (
	LeftIntRightInt BinaryElementType = iota
	LeftIntRightFloat
	LeftIntRightIntFloat
	LeftFloatRightFloat
	LeftFloatRightInt
	LeftFloatRightIntFloat
	LeftFloatIntRightFloat
	LeftFloatIntRightInt
	LeftFloatIntRightIntFloat
	LeftBooleanRightBoolean
	LeftStringRightString
	InvalidTypes
)

type PencilResult struct {
	Type  PencilType //
	Value interface{}
}
type CaseItem struct {
	matchValue  PencilResult
	resultValue PencilResult
}
type floatIntegerNumber struct {
	IntegerValue int64
	Precision    int64
}

func convertTwoFloatIntNumberToCommonPrecision(a floatIntegerNumber, b floatIntegerNumber) (floatIntegerNumber, floatIntegerNumber) {
	xPrecision := a.Precision
	yPrecision := b.Precision
	xInteger := a.IntegerValue
	yInteger := b.IntegerValue
	var difference int64 = 0
	var changeX bool = false
	if xPrecision > yPrecision {
		difference = xPrecision - yPrecision
		yPrecision = xPrecision
	} else {
		if xPrecision < yPrecision {
			difference = yPrecision - xPrecision
			xPrecision = yPrecision
			changeX = true
		}
	}
	power := (int64)(math.Pow10(int(difference)))
	if changeX {
		xInteger = xInteger * power
	} else {
		yInteger = yInteger * power

	}
	x := floatIntegerNumber{
		IntegerValue: xInteger,
		Precision:    xPrecision,
	}
	y := floatIntegerNumber{
		IntegerValue: yInteger,
		Precision:    yPrecision,
	}
	return x, y
}

func (fin floatIntegerNumber) Add(input floatIntegerNumber) floatIntegerNumber {
	left, right := convertTwoFloatIntNumberToCommonPrecision(fin, input)

	newInteger := (left.IntegerValue + right.IntegerValue)
	return floatIntegerNumber{
		IntegerValue: newInteger,
		Precision:    left.Precision,
	}
}
func (fin floatIntegerNumber) Subtract(input floatIntegerNumber) floatIntegerNumber {
	left, right := convertTwoFloatIntNumberToCommonPrecision(fin, input)

	newInteger := (left.IntegerValue - right.IntegerValue)
	return floatIntegerNumber{
		IntegerValue: newInteger,
		Precision:    left.Precision,
	}
}

func (fin floatIntegerNumber) Multiply(input floatIntegerNumber) floatIntegerNumber {
	precision := fin.Precision + input.Precision
	newInteger := fin.IntegerValue * input.IntegerValue
	return floatIntegerNumber{
		IntegerValue: newInteger,
		Precision:    precision,
	}
}
func (fin floatIntegerNumber) Divide(input floatIntegerNumber) floatIntegerNumber {
	//TODO Fix the precision
	precision := fin.Precision + input.Precision
	newInteger := fin.IntegerValue / input.IntegerValue
	return floatIntegerNumber{
		IntegerValue: newInteger,
		Precision:    precision,
	}
}
func (fin floatIntegerNumber) Equal(i floatIntegerNumber) bool {
	b := ((fin.IntegerValue == i.IntegerValue) &&
		(fin.Precision == i.Precision))
	return b
}
func (fin floatIntegerNumber) convertToFloat6Decimal() float64 {
	numerator := float64(fin.IntegerValue)
	divisor := (float64)(math.Pow10(int(fin.Precision)))
	n := numerator / float64(divisor)
	s := fmt.Sprintf("%.7f", n)
	_, f64 := convertFloatStringToFloatIntegerNumber(s)
	return f64
}
func (fin floatIntegerNumber) String() string {

	return fmt.Sprintf("%d (%d)", fin.IntegerValue, fin.Precision)
}
func (c CaseItem) String() string {
	t := fmt.Sprintf("Match: %s <> Result: %s", c.matchValue, c.resultValue)
	return t
}

type CaseStatement struct {
	expressionPencilResultValue PencilResult
	caseItems                   []CaseItem
	defaultCaseItem             CaseItem
}

func (p PencilResult) String() string {
	v := fmt.Sprintf("%v", p.Value)
	var t string
	switch p.Type {
	case PencilTypeString:
		t = "String"
	case PencilTypeInteger:
		t = "Integer"
	case PencilTypeFloat:
		t = "Float"
	case PencilTypeIntegerFloat:
		t = "InternalFloat (Integer)"
	case PencilTypeBoolean:
		t = "Boolean"
	case PencilTypeDateTime:
		t = "DateTime"
	case PencilTypeCaseItem:
		t = "CaseItem"
		v = p.Value.(CaseItem).String()
	case PencilTypeOperation:
		t = "Operation"
	default:
		t = "Unknown"
	}
	return t + " --> " + v
}
func determineBinaryOperationType(left PencilResult, right PencilResult) (BinaryElementType, string) {
	leftType := left.Type
	rightType := right.Type

	switch leftType {
	case PencilTypeInteger:
		switch rightType {
		case PencilTypeInteger:
			return LeftIntRightInt, "INT"
		case PencilTypeString, PencilTypeBoolean:
			return InvalidTypes, "INT"
		case PencilTypeFloat:
			return LeftIntRightFloat, "FLOAT"
		case PencilTypeIntegerFloat:
			return LeftIntRightIntFloat, "INT_FLOAT"
		default:
			return InvalidTypes, "INT"
		}
	case PencilTypeString:
		switch rightType {
		case PencilTypeString:
			return LeftStringRightString, "STRING"
		case PencilTypeInteger, PencilTypeFloat, PencilTypeIntegerFloat, PencilTypeBoolean:
			return InvalidTypes, "INT"
		default:
			return InvalidTypes, "INT"
		}
	case PencilTypeFloat:
		switch rightType {
		case PencilTypeFloat:
			return LeftFloatRightFloat, "FLOAT"
		case PencilTypeInteger:
			return LeftFloatRightInt, "FLOAT"
		case PencilTypeString, PencilTypeBoolean:
			return InvalidTypes, "INT"
		case PencilTypeIntegerFloat:
			return LeftFloatRightIntFloat, "FLOAT"
		default:
			return InvalidTypes, "INT"
		}
	case PencilTypeIntegerFloat:
		switch rightType {
		case PencilTypeFloat:
			return LeftFloatIntRightIntFloat, "INT"
		case PencilTypeInteger:
			return LeftFloatIntRightInt, "INT"
		case PencilTypeString, PencilTypeBoolean:
			return InvalidTypes, "INT"
		case PencilTypeIntegerFloat:
			return LeftFloatIntRightIntFloat, "INT_FLOAT"
		default:
			return InvalidTypes, "INT"
		}
	case PencilTypeBoolean:
		switch rightType {
		case PencilTypeBoolean:
			return LeftBooleanRightBoolean, "BOOLEAN"
		case PencilTypeInteger, PencilTypeFloat, PencilTypeIntegerFloat, PencilTypeString:
			return InvalidTypes, "INT"
		default:
			return InvalidTypes, "INT"
		}

	default:
		return InvalidTypes, "INT"
	}

}

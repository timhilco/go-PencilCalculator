package pencilCalculator

import (
	"fmt"
	"math"
)

type PencilType int

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

func (fin floatIntegerNumber) Equal(in floatIntegerNumber) bool {
	b := ((fin.IntegerValue == in.IntegerValue) &&
		(fin.Precision == in.Precision))
	return b
}
func (fin floatIntegerNumber) convertToFloat6Decimal() float64 {
	numerator := float64(fin.IntegerValue)
	divisor := (float64)(math.Pow10(int(fin.Precision)))
	return numerator / float64(divisor)
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
			return LeftIntRightIntFloat, "FLOAT"
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
			return LeftFloatIntRightIntFloat, "INT"
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

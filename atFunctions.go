package pencilCalculator

import (
	"math"
	"time"
)

var functionMap = map[string]interface{}{

	"Max": Max,
	"Now": Now,
}

func convertBinaryElementsToFloat(a PencilResult, b PencilResult) (float64, float64) {
	binaryType, _ := determineBinaryOperationType(a, b)
	var leftNumberFLOAT float64 = 0
	var rightNumberFLOAT float64 = 0
	switch binaryType {
	case LeftIntRightInt:
		leftNumberFLOAT = float64(a.Value.(int64))
		rightNumberFLOAT = float64(b.Value.(int64))
	case LeftIntRightFloat:
		leftNumberFLOAT = float64(a.Value.(int64))
		rightNumberFLOAT = b.Value.(float64)
	case LeftFloatRightInt:
		rightNumberFLOAT = float64(b.Value.(int64))
		leftNumberFLOAT = a.Value.(float64)
	case LeftFloatRightFloat:
		leftNumberFLOAT = a.Value.(float64)
		rightNumberFLOAT = b.Value.(float64)
	case LeftIntRightIntFloat:
		leftNumberFLOAT = float64(a.Value.(int64))
		fi := b.Value.(floatIntegerNumber)
		rightInteger := float64(fi.IntegerValue)
		divisor := math.Pow10(int(fi.Precision))
		rightNumberFLOAT = rightInteger / divisor
	case InvalidTypes:
		return 0.0, 0.0
	default:

	}
	return leftNumberFLOAT, rightNumberFLOAT
}
func Max(a, b interface{}) PencilResult {
	left := a.(PencilResult)
	right := b.(PencilResult)
	leftNumberFLOAT, rightNumberFLOAT := convertBinaryElementsToFloat(left, right)
	result := leftNumberFLOAT
	if rightNumberFLOAT > leftNumberFLOAT {
		result = rightNumberFLOAT
	}
	return PencilResult{
		Type:  PencilTypeFloat,
		Value: result,
	}
}
func Now() PencilResult {
	result := time.Now()
	return PencilResult{
		Type:  PencilTypeDateTime,
		Value: result,
	}
}

package pencilCalculator

import "fmt"

type PencilType int

const (
	PencilTypeString PencilType = iota
	PencilTypeInteger
	PencilTypeStruct
)

type PencilResult struct {
	Type  PencilType //
	Value interface{}
}

func (p PencilResult) String() string {
	var t string
	switch p.Type {
	case PencilTypeString:
		t = "String"
	case PencilTypeInteger:
		t = "Integer"
	default:
		t = "Unknown"
	}
	v := fmt.Sprintf("%v", p.Value)
	return t + " --> " + v
}

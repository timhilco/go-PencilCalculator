package pencilCalculator

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	statment := "if 1>2 then 3 else 4"
	result := Evaluate(statment)
	fmt.Println(result)
}
func TestTypeConversations(t *testing.T) {

	fi1 := PencilResult{
		PencilTypeIntegerFloat,
		convertFloatStringToFloatIntegerNumber("123.45"),
	}
	var tests = []struct {
		left      PencilResult
		right     PencilResult
		operation string
		want      interface{}
	}{
		{PencilResult{PencilTypeInteger, int64(10)}, fi1, "ADD", float64(133.45)},
		{PencilResult{PencilTypeFloat, 10.0}, PencilResult{PencilTypeFloat, 20.0}, "ADD", float64(30.0)},
		{PencilResult{PencilTypeInteger, int64(10)}, PencilResult{PencilTypeFloat, 20.0}, "ADD", float64(30.0)},
		{PencilResult{PencilTypeFloat, 10.0}, PencilResult{PencilTypeInteger, int64(20)}, "ADD", float64(30.0)},
		{PencilResult{PencilTypeInteger, int64(10)}, PencilResult{PencilTypeInteger, int64(20)}, "ADD", int64(30)},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.left, tt.right)
		t.Run(testname, func(t *testing.T) {
			ans, _ := doBinaryArithmatic(tt.left, tt.right, tt.operation)
			if ans.Value != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}
func TestExpresions(t *testing.T) {

	var tests = []struct {
		expression string
		want       interface{}
	}{
		{"100+200+300", int64(600)},
		{"@Max(1,2)", float64(2.0)},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.expression)
		t.Run(testname, func(t *testing.T) {
			ans := Evaluate(tt.expression)
			if ans.Value != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}

/*
type stubMapping map[string]interface{}

var StubStorage = stubMapping{}

func TestCall(t *testing.T) {
	StubStorage = map[string]interface{}{
		"funcA": funcA,
		"funcB": funcB,
		"Max":   Max,
	}

		resA, _ := Call("funcA", "value", "keyword1")
		var prntA string
		prntA = resA.(string)
		fmt.Println(prntA)

		resB, _ := Call("funcB", 10)
		var prntB int
		prntB = resB.(int)
		fmt.Println(prntB)

	f := reflect.ValueOf(Max)
	resB, _ := Call(f, 10, 20)

	fmt.Println(resB)
}
func Call(f reflect.Value, params ...interface{}) (result interface{}, err error) {
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
func funcA(arg0 string, arg1 string) string {
	return fmt.Sprintf("Function one ! %v %v", arg0, arg1)
}
func funcB(args0 int) int {
	return args0
}
*/

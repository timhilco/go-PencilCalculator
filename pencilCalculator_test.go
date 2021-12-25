package pencilCalculator

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	empJson := `{
	"id" : 11,
	"name" : "Irshad",
	"department" : "IT",
	"designation" : "Product Manager",
	"salary": 50000
}`
	context := context.Background()
	statment := "{01-01-2021}"
	result := Evaluate(context, statment, []byte(empJson))
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
	context := context.Background()
	empJson := `{
		"id" : 11,
		"name" : "Irshad",
		"department" : "IT",
		"designation" : "Product Manager",
		"salary": 50000
	}`
	var tests = []struct {
		expression string
		want       interface{}
	}{
		{"100+200+300", int64(600)},
		{"@Max(1,2)", float64(2.0)},
		{"Employee.salary", float64(50000.0)},
		{"Employee.salary > 10", true},
		{"Employee.salary > 10 AND 10 > 1", true},
		{"case Employee.department is 'IT' :: 1; 'default' :: 2 ; endcase", int64(1)},
		{"case Employee.department is 'ITS' :: 1; 'default' :: 2 ; endcase", int64(2)},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("Expression: %s", tt.expression)
		t.Run(testname, func(t *testing.T) {
			ans := Evaluate(context, tt.expression, []byte(empJson))
			if ans.Value != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}
func TestFloat(t *testing.T) {

	var tests = []struct {
		expression string
		want       float64
	}{
		{"123.45", float64(123.45)},
		{"123.1234567", float64(123.123457)},
		{"123.1234565", float64(123.123456)},
		{"123.1234564", float64(123.123456)},
		{"123.123456", float64(123.123456)},
		{"123.12345", float64(123.12345)},
		{"123.1234", float64(123.1234)},
		{"123.123", float64(123.123)},
		{"123.12", float64(123.12)},
		{"123.1", float64(123.1)},
		{"123", float64(123)},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("Expression: %s", tt.expression)
		t.Run(testname, func(t *testing.T) {
			fi := convertFloatStringToFloatIntegerNumber(tt.expression)
			ans := fi.convertToFloat6Decimal()
			if ans != tt.want {
				t.Errorf("got %f, want %f", ans, tt.want)
			}
		})
	}

}
func TestJSON(t *testing.T) {
	//Simple Employee JSON object which we will parse
	empJson := `{
        "id" : 11,
        "name" : "Irshad",
        "department" : "IT",
        "designation" : "Product Manager",
		"salary": 50000
	}`

	// Declared an empty interface
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(empJson), &result)

	//Reading each value by its key
	fmt.Println("Id :", result["id"],
		"\nName :", result["name"],
		"\nDepartment :", result["department"],
		"\nSalary :", result["salary"])
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

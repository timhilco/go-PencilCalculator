package pencilCalculator

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/rs/zerolog"
)

func getTwoRandonIntegers(min int, max int) (int64, int64) {
	var i1, i2 int64
	rand.Seed(time.Now().UnixNano())

	i1 = int64(rand.Intn(max-min+1) + min)
	i2 = int64(rand.Intn(max-min+1) + min)

	return i1, i2
}
func TestFloatRounding(t *testing.T) {
	f, _ := getTwoRandonFloats(10000000.0, 20000000.0)
	f2 := math.RoundToEven(f * 1000000)
	fmt.Printf("%.10f\n", f2)
	f3 := f2 / 1000000
	fmt.Printf("%.10f\n", f)
	fmt.Printf("%.10f\n", f3)
}

func getTwoRandonFloats(min float64, max float64) (float64, float64) {

	rand.Seed(time.Now().UnixNano())

	f1 := (rand.Float64() * (max - min)) + min
	f2 := (rand.Float64() * (max - min)) + min

	return f1, f2
}

/*
type tests struct {
	expression string
	want       interface{}
}
*/

func TestFloatPerformance(t *testing.T) {
	ctx := context.Background()
	empJson := `{
		"id": 11,
		"name": "Irshad",
		"department": "IT",
		"designation": "Product Manager",
		"salary": 50000,
		"payHistory": [{
			"effectiveDate": "01/01/2021",
			"amount": 40000
		}, {
			"effectiveDate": "01/01/2022",
			"amount": 50000
		}]
	
	}`
	inputData := make(map[string]string)
	inputData["Employee"] = empJson

	ctx = context.WithValue(ctx, InputDataContextKey{}, inputData)
	ctx = context.WithValue(ctx, LoggingLevelContextKey{}, zerolog.Disabled)
	var expression string
	var want float64
	for i := 1; i < 1000; i++ {
		i1, i2 := getTwoRandonFloats(0.0, 1000000.0)
		//fmt.Printf("i1=%.10f", i1)
		i1 = convertFloat64ToFloat64with6DecimalPlaces(i1)
		fmt.Printf("i1=%.10f\n", i1)
		//fmt.Printf("i2=%.10f", i2)
		i2 = convertFloat64ToFloat64with6DecimalPlaces(i2)
		fmt.Printf("i2=%.10f\n", i2)
		for j := 1; j < 5; j++ {
			switch j {
			case 1:
				expression = fmt.Sprintf("%.6f+%.6f", i1, i2)
				want = i1 + i2
			case 2:
				expression = fmt.Sprintf("%.6f-%.6f", i1, i2)
				want = i1 - i2
			case 3:
				expression = fmt.Sprintf("%.6f*%.6f", i1, i2)
				want = i1 * i2
			case 4:
				expression = fmt.Sprintf("%.6f/%.6f", i1, i2)
				want = i1 / i2
			}

			fmt.Printf("Before: WANT=%.10f\n", want)
			want = convertFloat64ToFloat64with6DecimalPlaces(want)
			fmt.Printf("After:  WANT=%.10f\n", want)
			testname := fmt.Sprintf("Expression: %s", expression)
			t.Run(testname, func(t *testing.T) {
				ans := Evaluate(ctx, expression)
				v := ans.Value.(float64)
				sV := fmt.Sprintf("%.10f", v)
				sWant := fmt.Sprintf("%.10f", want)
				fmt.Printf("v=%.10f\n", v)
				if sV != sWant {
					//t.Errorf("got %d, want %d", ans, want)
					diff := v - want
					t.Logf("got %s, want %s -> difference=%.10f", sV, sWant, diff)
					if diff > 0.00009 {
						t.Fatalf("i1 = %.10f, i2 = %.10f-> difference=%.7f", i1, i2, diff)
					}

				}
			})
		}
	}

}
func TestIntegerPerformance(t *testing.T) {
	ctx := context.Background()
	empJson := `{
		"id": 11,
		"name": "Irshad",
		"department": "IT",
		"designation": "Product Manager",
		"salary": 50000,
		"payHistory": [{
			"effectiveDate": "01/01/2021",
			"amount": 40000
		}, {
			"effectiveDate": "01/01/2022",
			"amount": 50000
		}]
	
	}`
	inputData := make(map[string]string)
	inputData["Employee"] = empJson

	ctx = context.WithValue(ctx, InputDataContextKey{}, inputData)
	ctx = context.WithValue(ctx, LoggingLevelContextKey{}, zerolog.Disabled)
	var expression string
	var want int64
	for i := 1; i < 100; i++ {
		i1, i2 := getTwoRandonIntegers(1, 1000000)
		for j := 1; j < 5; j++ {
			switch j {
			case 1:
				expression = fmt.Sprintf("%d+%d", i1, i2)
				want = i1 + i2
			case 2:
				expression = fmt.Sprintf("%d-%d", i1, i2)
				want = i1 - i2
			case 3:
				expression = fmt.Sprintf("%d*%d", i1, i2)
				want = i1 * i2
			case 4:
				expression = fmt.Sprintf("%d/%d", i1, i2)
				want = i1 / i2
			}
			testname := fmt.Sprintf("Expression: %s", expression)
			t.Run(testname, func(t *testing.T) {
				ans := Evaluate(ctx, expression)
				v := ans.Value
				if v != want {
					//t.Errorf("got %d, want %d", ans, want)
					t.Fatalf("got %d, want %d", ans, want)
				}
			})
		}
	}

}

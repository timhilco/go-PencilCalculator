package pencilCalculator

import (
	"context"
	"fmt"
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

/*
type tests struct {
	expression string
	want       interface{}
}
*/
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

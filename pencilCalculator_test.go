package pencilCalculator

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	statment := "100+200+300"
	result := Evalute(statment)
	fmt.Println(result)
}

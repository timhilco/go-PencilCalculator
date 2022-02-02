package tests

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/timhilco/go-PencilCalculator/pencilCalculator"
)

func TestProgram1(t *testing.T) {
	ctx := context.Background()
	empJson := `{
		"id": 11,
		"name": { "first": "John", "last" : "Sample"},
		"department": "IT",
		"designation": "Product Manager",
		"salary": 50000,
		"payHistory": [{
			"effectiveDate": "01/01/2021",
			"amount": 40000,
			"units" : 100
			}, {
				"effectiveDate": "01/01/2022",
				"amount": 50000,
				"units" : 110
				}]
				
				}`
	inputData := make(map[string]string)
	inputData["Employee"] = empJson

	ctx = context.WithValue(ctx, pencilCalculator.InputDataContextKey{}, inputData)

	b, err := os.ReadFile("../data/data.txt")
	program := string(b)
	if err != nil {
		fmt.Println(err)
	}
	statements := pencilCalculator.EvaluateProgram(ctx, program)
	fmt.Println("---------------Program Results-------------------------------------")
	for _, statement := range statements {
		fmt.Printf("Statement: %s = %v\n", statement.StatementKey, statement.Result)
	}

}
func buildRandomProgram() string {
	variables := make([]string, 0)
	var sb strings.Builder
	for i := 0; i < 16; i++ {
		variable := fmt.Sprintf("WE:e%d", i)
		variables = append(variables, variable)
		sb.WriteString(buildElementalStatement(i))
	}
	for level := 0; level < 50; level++ {
		variable := fmt.Sprintf("W%d:a%d", level, level)
		elements := make([]string, 0)
		for i := 0; i < level; i++ {
			count := len(variables)
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(count)
			field := variables[n]
			elements = append(elements, field)
		}
		sb.WriteString(buildStatement(variable, elements))
		variables = append(variables, variable)
	}
	s := sb.String()
	b := []byte(s)
	err := os.WriteFile("../data/data.txt", b, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return s

}
func TestProgram2(t *testing.T) {
	program := buildRandomProgram()
	fmt.Println(program)
}
func buildElementalStatement(number int) string {
	var sb strings.Builder
	s := fmt.Sprintf("WE:e%d", number)
	sb.WriteString(s)
	sb.WriteString(" := ")
	value := ""
	if number == 0 {
		value = fmt.Sprintf("@Long (%d, %d)", 60, 10)

	} else if number%2 == 1 {
		value = fmt.Sprintf("%d", number)
	} else {
		value = fmt.Sprintf("@Long (%d, %d)", number*2, number)

	}
	sb.WriteString(value)
	sb.WriteString(" ;\n")
	return sb.String()
}
func buildStatement(variable string, components []string) string {
	var sb strings.Builder
	sb.WriteString(variable)
	sb.WriteString(" := ")
	for _, element := range components {
		sb.WriteString(element + " + ")
	}
	sb.WriteString("WE:e0;\n")
	return sb.String()
}
func buildProgram1() string {
	var sb strings.Builder
	sb.WriteString("W1:a :=  20 ; ")
	sb.WriteString("W1:b :=  10 ; ")
	sb.WriteString("W1:c :=  W1:a + W1:b ; ")
	sb.WriteString("W1:d :=  W1:a - W1:b ; ")
	sb.WriteString("W1:e :=  W1:a * W1:b ; ")
	sb.WriteString("W1:f :=  W1:a - W1:b ; ")
	return sb.String()
}
func TestChannel(t *testing.T) {
	cell1 := Cell{
		formula:            10,
		SubscriberChannels: make([]chan int, 0),
		//postingChannel:      make(chan chan int),
	}
	for i := 0; i < 200; i++ {
		cj := make(chan int)
		fmt.Printf("GoRoutine %d Posting\n", i)
		cell1.AddSubcriberChannel(cj)
		go func(j int) {
			fmt.Printf("GoRoutine %d Waiting\n", j)
			value := <-cj * j
			fmt.Printf("J: %d = %d\n", j, value)
		}(i)
	}
	cell1.Evaluate()
	time.Sleep(10 * time.Second)
	fmt.Println("Program Ending")

}

type Cell struct {
	formula            int
	SubscriberChannels []chan int
	//postingChannel      chan chan int
}

func (c *Cell) AddSubcriberChannel(ch chan int) {

	fmt.Printf("Cell adding subscriber\n")
	c.SubscriberChannels = append(c.SubscriberChannels, ch)
	fmt.Printf("Cell adding subscriber length: %d\n", len(c.SubscriberChannels))

}

func (c *Cell) Evaluate() {
	fmt.Printf("--> Evaluate \n")
	value := c.formula
	for _, channel := range c.SubscriberChannels {
		fmt.Printf("Evaluate posting to channel\n")
		//time.Sleep(3 * time.Second)
		channel <- value
	}
}

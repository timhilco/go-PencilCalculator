package pencilCalculator

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/timhilco/go-PencilCalculator/parser"
	"github.com/timhilco/go-PencilCalculator/util/logger"
)

func Evaluate(ctx context.Context, input string) PencilResult {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewHilcoPencilGrammarLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewHilcoPencilGrammarParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener HilcoPencilGrammarParserListener

	listener.SetLexer(lexer)
	//listener.SetInputData(jsonObject)
	listener.SetContext(ctx)
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Program())
	value := listener.Result()
	if value.Type == PencilTypeIntegerFloat {
		v := value.PrValue.(FloatIntegerNumber)
		value = PencilResult{
			Type:    PencilTypeFloat,
			PrValue: v.ConvertFloatIntToFloatInputPlaces(v.Precision),
		}
	}
	return value

}
func EvaluateProgram(ctx context.Context, program string) map[string]Statement {
	// Setup the input
	is := antlr.NewInputStream(program) //
	logger := logger.NewMultiWithFile(false, false, "../logs/evaluateProgram.txt")

	// Create the Lexer
	lexer := parser.NewHilcoPencilGrammarLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewHilcoPencilGrammarParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener HilcoPencilStatementListener

	listener.SetLexer(lexer)
	//listener.SetInputData(jsonObject)
	listener.SetContext(ctx)
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Program())
	statements := listener.GetStatementMap()
	for sKey := range statements {
		statement := statements[sKey]
		variables := statement.Variables
		for _, variable := range variables {
			key := variable.Name
			componentStatement := statements[key]
			// Make Statement Changes - Components
			components := statement.ComponmentChannels
			components[key] = componentStatement.StatementChannel
			statement.ComponmentChannels = components
			statements[sKey] = statement
			// Subscriber Channels
			subscribers := componentStatement.SubscriberChannels
			subscribers[sKey] = statement.StatementChannel
			componentStatement.SubscriberChannels = subscribers
			statements[key] = componentStatement
		}

	}
	ctx = context.WithValue(ctx, StatementMapContextKey{}, statements)
	ctx = context.WithValue(ctx, ProcessingPassContextKey{}, "Elemental")
	elementalStatements := make([]Statement, 0)
	dependentStatements := make([]Statement, 0)
	for _, statement := range statements {
		logger.Info().Msg(statement.String())
		if len(statement.ComponmentChannels) != 0 {
			dependentStatements = append(dependentStatements, statement)
		} else {
			elementalStatements = append(elementalStatements, statement)
		}
	}

	var wg sync.WaitGroup

	for _, statement := range elementalStatements {
		wg.Add(1)
		go func(s Statement) {
			defer wg.Done()
			startTime := time.Now()
			sf := fmt.Sprintf("** Evaluating elemental statement: %s\n", s.StatementKey)
			logger.Info().Msg(sf)
			result := evaluateStatement(ctx, s)
			s.Result = result
			key := s.StatementKey
			statements[key] = s
			sf = fmt.Sprintf("-- Evaluated elemental statement: %v\n", s.StatementKey)
			logger.Info().Msg(sf)
			for key, subscriber := range s.SubscriberChannels {
				sf = fmt.Sprintf(" Calling Elemental Subscribers: %s->%s:%v - Result:%v\n", s.StatementKey, key, subscriber, result)
				logger.Info().Msg(sf)
				subscriber <- s.StatementKey
				sf = fmt.Sprintf(" Called  Elemental Subscribers: %s->%s:%v \n", s.StatementKey, key, subscriber)
				logger.Info().Msg(sf)
			}
			endTime := time.Now()
			d := endTime.Sub(startTime)
			sf = fmt.Sprintf("** Elemental statement processing complete: %s\n Time: %v", s.StatementKey, d)
			logger.Info().Msg(sf)
		}(statement)

	}
	ctx2 := context.Background()
	ctx2 = context.WithValue(ctx2, ProcessingPassContextKey{}, "Dependent")
	ctx2 = context.WithValue(ctx2, StatementMapContextKey{}, statements)
	inputData := ctx.Value(InputDataContextKey{}).(map[string]string)
	ctx2 = context.WithValue(ctx2, InputDataContextKey{}, inputData)

	for _, statement := range dependentStatements {
		wg.Add(1)
		go func(s Statement) {
			defer wg.Done()
			startTime := time.Now()
			count := len(s.ComponmentChannels)
			a := float64(rand.Intn(10)) + 5.00
			n := math.Abs(a)
			if n == 0.0 {
				n = 1.0
			}
			d := time.Duration(n) * time.Second
			ticker := time.NewTicker(d)
			run := true
			i := 0
			c := s.StatementChannel
			sf := fmt.Sprintf("-> Going to wait for: %s:%v\n", s.StatementKey, c)
			logger.Info().Msg(sf)
			for run {
				select {
				case <-ticker.C:
					t := time.Now()
					fTime := t.Format(time.RFC3339)
					sf := fmt.Sprintf("-> Hit ticker %s for %s\n", fTime, s.StatementKey)
					logger.Info().Msg(sf)
					//run = false
				case <-c:
					sf := fmt.Sprintf("-> Subscriber published for statement: %s", s.StatementKey)
					logger.Info().Msg(sf)
					i++
					if i == count {
						run = false
					}
				default:
					//sf := fmt.Sprintf("-> Hit default: %s", s.StatementKey)
					//logger.Info().Msg(sf)
					d := time.Duration(1) * time.Second
					time.Sleep(d)

				}
			}

			sf = fmt.Sprintf("** Evaluating dependent statement: %s\n", s.StatementKey)
			logger.Info().Msg(sf)
			result := evaluateStatement(ctx2, s)
			s.Result = result
			key := s.StatementKey
			statements[key] = s
			sf = fmt.Sprintf("-- Evaluated dependent statement: %s = %v\n", s.StatementKey, result)
			logger.Info().Msg(sf)
			for key, subscriber := range s.SubscriberChannels {
				sf := fmt.Sprintf("Calling Dependent Subscriber: %s->%s:%v - Result:%v\n", s.StatementKey, key, subscriber, result)
				logger.Info().Msg(sf)
				subscriber <- s.StatementKey
				sf = fmt.Sprintf("Called   Dependent Subscriber: %s->%s%v\n", s.StatementKey, key, subscriber)
				logger.Info().Msg(sf)
			}
			endTime := time.Now()
			dt := endTime.Sub(startTime)
			sf = fmt.Sprintf("** Dependent statement processing complete: %s Time: %v\n", s.StatementKey, dt)
			logger.Info().Msg(sf)
		}(statement)
		//sf:= fmt.Sprintf("Ignoreing statement: %v\n", statement.VariableName)
	}

	wg.Wait()

	return statements
}
func evaluateStatement(ctx context.Context, s Statement) PencilResult {
	expression := s.Expression
	//sf:= fmt.Sprintf("Evaluating expression: %s\n", expression)
	result := Evaluate(ctx, expression)
	//sf:= fmt.Sprintf("%s=%v\n", expression, result)
	return result
}

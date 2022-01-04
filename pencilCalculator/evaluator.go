package pencilCalculator

import (
	"context"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/timhilco/go-PencilCalculator/parser"
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
		v := value.PrValue.(floatIntegerNumber)
		value = PencilResult{
			Type:    PencilTypeFloat,
			PrValue: v.ConvertFloatIntToFloat6Decimal(),
		}
	}
	return value

}

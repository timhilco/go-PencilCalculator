package pencilCalculator

import (
	"context"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/timhilco/go-PencilCalculator/parser"
)

func Evaluate(ctx context.Context, input string, jsonObject []byte) PencilResult {
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
	listener.SetInputData(jsonObject)
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Program())
	value := listener.Result()
	return value

}

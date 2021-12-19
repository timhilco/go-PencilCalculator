package pencilCalculator

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/timhilco/go-PencilCalculator/parser"
)

func Evalute(input string) PencilResult {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewHilcoPencilParserLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewHilcoPencilParserParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener PencilListener
	listener.SetLexer(lexer)
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Prog())
	value := listener.Result()
	return value

}

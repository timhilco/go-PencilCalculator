package pencilCalculator

import (
	"context"
	"fmt"
	"strings"

	"github.com/rs/zerolog"
	"github.com/timhilco/go-PencilCalculator/parser"
	"github.com/timhilco/go-PencilCalculator/util/logger"
)

type HilcoPencilStatementListener struct {
	*parser.BaseHilcoPencilGrammarParserListener
	logger            zerolog.Logger
	lexer             *parser.HilcoPencilGrammarLexer
	processingContext context.Context
	statmentMap       map[string]Statement
	worksheetVariable []WorksheetVariable
}
type WorksheetVariable struct {
	Name string
}
type Statement struct {
	StatementKey       string
	Expression         string
	SubscriberChannels map[string]chan string
	ComponmentChannels map[string]chan string
	StatementChannel   chan string
	Result             PencilResult
	Variables          []WorksheetVariable
	//Callers            []Statement
}

func (p *HilcoPencilStatementListener) GetStatementMap() map[string]Statement {
	return p.statmentMap
}

func (p *HilcoPencilStatementListener) SetLexer(lexer *parser.HilcoPencilGrammarLexer) {
	p.lexer = lexer
}
func (p *HilcoPencilStatementListener) SetContext(ctx context.Context) {
	p.processingContext = ctx

}
func (p *HilcoPencilStatementListener) Result() PencilResult {
	return PencilResult{}
}
func (p *HilcoPencilStatementListener) EnterProgram(ctx *parser.ProgramContext) {
	p.logger = logger.NewMultiWithFile(false, false, "../logs/log.txt")
	loggingLevel := p.processingContext.Value(LoggingLevelContextKey{})
	if loggingLevel != nil {
		level := loggingLevel.(zerolog.Level)
		zerolog.SetGlobalLevel(level)
	}

	p.logger.Info().Msg("EnterProgram")
	p.statmentMap = make(map[string]Statement, 0)
}
func (p *HilcoPencilStatementListener) ExitProgram(ctx *parser.ProgramContext) {

	p.logger.Info().Msg("------------------ Exit Program Listener -------------------------")

}

// EnterStatement is called when production statement is entered.
func (s *HilcoPencilStatementListener) EnterStatement(ctx *parser.StatementContext) {
	s.logger.Info().Msg("<<Enter EnterStatement: ")
	s.worksheetVariable = make([]WorksheetVariable, 0)

}

// ExitStatement is called when production statement is exited.
func (s *HilcoPencilStatementListener) ExitStatement(ctx *parser.StatementContext) {
	s.logger.Info().Msg(">>Enter ExitStatement: ")
	s0 := ctx.GetText()
	value := strings.Split(s0, ":=")
	s1 := value[1]
	v := strings.Trim(s1, ";")
	children := ctx.GetChildren()
	var variables []WorksheetVariable
	key := children[0].(*parser.WorkSheetVariableCalculatorContext).GetText()
	cCH := make(map[string]chan string)
	for i := 1; i < len(s.worksheetVariable); i++ {
		variables = append(variables, s.worksheetVariable[i])
	}
	r := PencilResult{
		Type:    PencilTypeNil,
		PrValue: nil,
	}

	subCH := make(map[string]chan string)
	sCH := make(chan string)

	statement := Statement{
		StatementKey:       key,
		Expression:         v,
		Variables:          variables,
		StatementChannel:   sCH,
		SubscriberChannels: subCH,
		ComponmentChannels: cCH,
		Result:             r,
	}
	s.statmentMap[key] = statement
	//s.logger.Info().Msg(">>Exit ExitStatement: " + statement.String())

}
func (s Statement) String() string {
	var sb strings.Builder
	s1 := fmt.Sprintf("- Statement: %s --\n", s.StatementKey)
	sb.WriteString(s1)
	s1 = fmt.Sprintf("Expression: %s \n", s.Expression)
	sb.WriteString(s1)
	s1 = fmt.Sprintf("Statement(my) Channel: %T -> %p \n", s.StatementChannel, s.StatementChannel)
	sb.WriteString(s1)

	sb.WriteString("Component Channels ->\n")
	for key, entry := range s.ComponmentChannels {
		s1 = fmt.Sprintf("%s -> %T:%p \n", key, entry, entry)
		sb.WriteString(s1)
	}
	sb.WriteString("Subscriber Channels ->\n")
	for key, c := range s.SubscriberChannels {
		s1 = fmt.Sprintf("%s -> %T:%p \n", key, c, c)
		sb.WriteString(s1)
	}
	sb.WriteString("Varibles ->\n")
	for _, c := range s.Variables {
		s1 = fmt.Sprintf("%v\n", c)
		sb.WriteString(s1)
	}
	s1 = fmt.Sprintf("Result: = %v \n", s.Result)
	sb.WriteString(s1)

	return sb.String()
}
func (s *HilcoPencilStatementListener) EnterWorksheetVariable(ctx *parser.WorksheetVariableContext) {
	s.logger.Info().Msg("Enter EnterWorksheetVariable: ")
}

// ExitWorksheetVariable is called when production worksheetVariable is exited.
func (s *HilcoPencilStatementListener) ExitWorksheetVariable(ctx *parser.WorksheetVariableContext) {
	s.logger.Info().Msg("Exit  EnterWorksheetVariable: ")
	t := ctx.GetText()
	variable := WorksheetVariable{
		Name: t,
	}
	s.worksheetVariable = append(s.worksheetVariable, variable)

}

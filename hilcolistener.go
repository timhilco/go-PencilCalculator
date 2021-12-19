package pencilCalculator

import (
	"fmt"
	"strconv"
	"timhilco/go-PencilCalculator/parser"
	"timhilco/go-PencilCalculator/util/logger"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type PencilListener struct {
	*parser.BaseHilcoPencilParserListener
	stack           []PencilResult
	logger          *logger.HilcoLogger
	lexer           *parser.HilcoPencilParserLexer
	currentOperator string
}

func (p *PencilListener) SetLexer(lexer *parser.HilcoPencilParserLexer) {
	p.lexer = lexer
}
func (p *PencilListener) Result() PencilResult {
	return p.stack[0]
}
func (l *PencilListener) push(i PencilResult) {
	l.stack = append(l.stack, i)
}

func (l *PencilListener) pop() PencilResult {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}
func (p *PencilListener) logStack() {
	for i, pr := range p.stack {
		text := fmt.Sprintf("%d:%s", i, pr.String())
		p.logger.Info(text)
	}
}

// EnterProg is called when production prog is entered.
func (p *PencilListener) EnterProg(ctx *parser.ProgContext) {
	p.logger = logger.NewMultiWithFile(false)
	p.logger.Info("EnterProg")
	p.stack = make([]PencilResult, 0)
}
func (p *PencilListener) ExitAddingExpression(ctx *parser.AddingExpressionContext) {
	p.logger.Info("ExitAddingExpression")
	right, left := p.pop(), p.pop()

	switch p.currentOperator {
	case "PLUS":
		var lNum int64
		lNum = left.Value.(int64)
		var rNum int64
		rNum = right.Value.(int64)
		result := lNum + rNum
		pr := PencilResult{
			Type:  PencilTypeInteger,
			Value: result,
		}
		p.push(pr)
	case "MINUS":

	default:
		panic(fmt.Sprintf("unexpected op: %s", p.currentOperator))
	}
	p.logStack()
}

// VisitTerminal is called when a terminal node is visited.
func (p *PencilListener) VisitTerminal(node antlr.TerminalNode) {
	token := node.GetSymbol()
	symbol := token.GetTokenType()
	name := p.lexer.SymbolicNames[token.GetTokenType()]
	value := node.GetText()
	text := fmt.Sprintf("%s(%d) -> %s", name, symbol, value)
	switch name {
	case "INT": // INT
		number, _ := strconv.ParseInt(value, 10, 64)
		pr := PencilResult{
			Type:  PencilTypeInteger,
			Value: number,
		}
		p.push(pr)
	case "PLUS":
		p.currentOperator = name
		p.logger.Info("VisitTermnal <> Binary Operator: " + text)

	default:
		p.logger.Info("VisitTermnal <> Unknown Terminal: " + text)

	}
	p.logger.Info("VisitTermnal: " + text)
	p.logStack()

}

// VisitErrorNode is called when an error node is visited.
func (s *PencilListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *PencilListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *PencilListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStatement is called when production statement is entered.
func (s *PencilListener) EnterStatement(ctx *parser.StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *PencilListener) ExitStatement(ctx *parser.StatementContext) {}

// EnterCaseStatement is called when production caseStatement is entered.
func (s *PencilListener) EnterCaseStatement(ctx *parser.CaseStatementContext) {}

// ExitCaseStatement is called when production caseStatement is exited.
func (s *PencilListener) ExitCaseStatement(ctx *parser.CaseStatementContext) {}

// EnterCaseList is called when production caseList is entered.
func (s *PencilListener) EnterCaseList(ctx *parser.CaseListContext) {}

// ExitCaseList is called when production caseList is exited.
func (s *PencilListener) ExitCaseList(ctx *parser.CaseListContext) {}

// EnterCaseItem is called when production caseItem is entered.
func (s *PencilListener) EnterCaseItem(ctx *parser.CaseItemContext) {}

// ExitCaseItem is called when production caseItem is exited.
func (s *PencilListener) ExitCaseItem(ctx *parser.CaseItemContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *PencilListener) EnterSwitchStatement(ctx *parser.SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *PencilListener) ExitSwitchStatement(ctx *parser.SwitchStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *PencilListener) EnterIfStatement(ctx *parser.IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *PencilListener) ExitIfStatement(ctx *parser.IfStatementContext) {}

// EnterLogical is called when production logical is entered.
func (s *PencilListener) EnterLogical(ctx *parser.LogicalContext) {}

// ExitLogical is called when production logical is exited.
func (s *PencilListener) ExitLogical(ctx *parser.LogicalContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *PencilListener) EnterRelationalExpression(ctx *parser.RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *PencilListener) ExitRelationalExpression(ctx *parser.RelationalExpressionContext) {}

// EnterAddingExpression is called when production addingExpression is entered.
func (s *PencilListener) EnterAddingExpression(ctx *parser.AddingExpressionContext) {}

// EnterMultiplyingExpression is called when production multiplyingExpression is entered.
func (s *PencilListener) EnterMultiplyingExpression(ctx *parser.MultiplyingExpressionContext) {
}

// ExitMultiplyingExpression is called when production multiplyingExpression is exited.
func (s *PencilListener) ExitMultiplyingExpression(ctx *parser.MultiplyingExpressionContext) {
}

// EnterFactor is called when production factor is entered.
func (s *PencilListener) EnterFactor(ctx *parser.FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *PencilListener) ExitFactor(ctx *parser.FactorContext) {}

// EnterBase is called when production base is entered.
func (s *PencilListener) EnterBase(ctx *parser.BaseContext) {}

// ExitBase is called when production base is exited.
func (s *PencilListener) ExitBase(ctx *parser.BaseContext) {}

// EnterName is called when production name is entered.
func (s *PencilListener) EnterName(ctx *parser.NameContext) {}

// ExitName is called when production name is exited.
func (s *PencilListener) ExitName(ctx *parser.NameContext) {}

// EnterAtFunction is called when production atFunction is entered.
func (s *PencilListener) EnterAtFunction(ctx *parser.AtFunctionContext) {}

// ExitAtFunction is called when production atFunction is exited.
func (s *PencilListener) ExitAtFunction(ctx *parser.AtFunctionContext) {}

// EnterDataAccessor is called when production dataAccessor is entered.
func (s *PencilListener) EnterDataAccessor(ctx *parser.DataAccessorContext) {}

// ExitDataAccessor is called when production dataAccessor is exited.
func (s *PencilListener) ExitDataAccessor(ctx *parser.DataAccessorContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *PencilListener) EnterFieldName(ctx *parser.FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *PencilListener) ExitFieldName(ctx *parser.FieldNameContext) {}

// EnterArgList is called when production argList is entered.
func (s *PencilListener) EnterArgList(ctx *parser.ArgListContext) {}

// ExitArgList is called when production argList is exited.
func (s *PencilListener) ExitArgList(ctx *parser.ArgListContext) {}

// EnterSpecialKeyword is called when production specialKeyword is entered.
func (s *PencilListener) EnterSpecialKeyword(ctx *parser.SpecialKeywordContext) {}

// ExitSpecialKeyword is called when production specialKeyword is exited.
func (s *PencilListener) ExitSpecialKeyword(ctx *parser.SpecialKeywordContext) {}

// Code generated from HilcoPencilParser.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // HilcoPencilParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHilcoPencilParserListener is a complete listener for a parse tree produced by HilcoPencilParserParser.
type BaseHilcoPencilParserListener struct{}

var _ HilcoPencilParserListener = &BaseHilcoPencilParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHilcoPencilParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHilcoPencilParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHilcoPencilParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHilcoPencilParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseHilcoPencilParserListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseHilcoPencilParserListener) ExitProg(ctx *ProgContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseHilcoPencilParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseHilcoPencilParserListener) ExitStatement(ctx *StatementContext) {}

// EnterCaseStatement is called when production caseStatement is entered.
func (s *BaseHilcoPencilParserListener) EnterCaseStatement(ctx *CaseStatementContext) {}

// ExitCaseStatement is called when production caseStatement is exited.
func (s *BaseHilcoPencilParserListener) ExitCaseStatement(ctx *CaseStatementContext) {}

// EnterCaseList is called when production caseList is entered.
func (s *BaseHilcoPencilParserListener) EnterCaseList(ctx *CaseListContext) {}

// ExitCaseList is called when production caseList is exited.
func (s *BaseHilcoPencilParserListener) ExitCaseList(ctx *CaseListContext) {}

// EnterCaseItem is called when production caseItem is entered.
func (s *BaseHilcoPencilParserListener) EnterCaseItem(ctx *CaseItemContext) {}

// ExitCaseItem is called when production caseItem is exited.
func (s *BaseHilcoPencilParserListener) ExitCaseItem(ctx *CaseItemContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseHilcoPencilParserListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseHilcoPencilParserListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseHilcoPencilParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseHilcoPencilParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterLogical is called when production logical is entered.
func (s *BaseHilcoPencilParserListener) EnterLogical(ctx *LogicalContext) {}

// ExitLogical is called when production logical is exited.
func (s *BaseHilcoPencilParserListener) ExitLogical(ctx *LogicalContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *BaseHilcoPencilParserListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *BaseHilcoPencilParserListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterAddingExpression is called when production addingExpression is entered.
func (s *BaseHilcoPencilParserListener) EnterAddingExpression(ctx *AddingExpressionContext) {}

// ExitAddingExpression is called when production addingExpression is exited.
func (s *BaseHilcoPencilParserListener) ExitAddingExpression(ctx *AddingExpressionContext) {}

// EnterMultiplyingExpression is called when production multiplyingExpression is entered.
func (s *BaseHilcoPencilParserListener) EnterMultiplyingExpression(ctx *MultiplyingExpressionContext) {
}

// ExitMultiplyingExpression is called when production multiplyingExpression is exited.
func (s *BaseHilcoPencilParserListener) ExitMultiplyingExpression(ctx *MultiplyingExpressionContext) {
}

// EnterFactor is called when production factor is entered.
func (s *BaseHilcoPencilParserListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseHilcoPencilParserListener) ExitFactor(ctx *FactorContext) {}

// EnterBase is called when production base is entered.
func (s *BaseHilcoPencilParserListener) EnterBase(ctx *BaseContext) {}

// ExitBase is called when production base is exited.
func (s *BaseHilcoPencilParserListener) ExitBase(ctx *BaseContext) {}

// EnterName is called when production name is entered.
func (s *BaseHilcoPencilParserListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseHilcoPencilParserListener) ExitName(ctx *NameContext) {}

// EnterAtFunction is called when production atFunction is entered.
func (s *BaseHilcoPencilParserListener) EnterAtFunction(ctx *AtFunctionContext) {}

// ExitAtFunction is called when production atFunction is exited.
func (s *BaseHilcoPencilParserListener) ExitAtFunction(ctx *AtFunctionContext) {}

// EnterDataAccessor is called when production dataAccessor is entered.
func (s *BaseHilcoPencilParserListener) EnterDataAccessor(ctx *DataAccessorContext) {}

// ExitDataAccessor is called when production dataAccessor is exited.
func (s *BaseHilcoPencilParserListener) ExitDataAccessor(ctx *DataAccessorContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseHilcoPencilParserListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseHilcoPencilParserListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterArgList is called when production argList is entered.
func (s *BaseHilcoPencilParserListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production argList is exited.
func (s *BaseHilcoPencilParserListener) ExitArgList(ctx *ArgListContext) {}

// EnterSpecialKeyword is called when production specialKeyword is entered.
func (s *BaseHilcoPencilParserListener) EnterSpecialKeyword(ctx *SpecialKeywordContext) {}

// ExitSpecialKeyword is called when production specialKeyword is exited.
func (s *BaseHilcoPencilParserListener) ExitSpecialKeyword(ctx *SpecialKeywordContext) {}

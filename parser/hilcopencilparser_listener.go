// Code generated from HilcoPencilParser.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // HilcoPencilParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HilcoPencilParserListener is a complete listener for a parse tree produced by HilcoPencilParserParser.
type HilcoPencilParserListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterCaseStatement is called when entering the caseStatement production.
	EnterCaseStatement(c *CaseStatementContext)

	// EnterCaseList is called when entering the caseList production.
	EnterCaseList(c *CaseListContext)

	// EnterCaseItem is called when entering the caseItem production.
	EnterCaseItem(c *CaseItemContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterLogical is called when entering the logical production.
	EnterLogical(c *LogicalContext)

	// EnterRelationalExpression is called when entering the relationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterAddingExpression is called when entering the addingExpression production.
	EnterAddingExpression(c *AddingExpressionContext)

	// EnterMultiplyingExpression is called when entering the multiplyingExpression production.
	EnterMultiplyingExpression(c *MultiplyingExpressionContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterBase is called when entering the base production.
	EnterBase(c *BaseContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterAtFunction is called when entering the atFunction production.
	EnterAtFunction(c *AtFunctionContext)

	// EnterDataAccessor is called when entering the dataAccessor production.
	EnterDataAccessor(c *DataAccessorContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterArgList is called when entering the argList production.
	EnterArgList(c *ArgListContext)

	// EnterSpecialKeyword is called when entering the specialKeyword production.
	EnterSpecialKeyword(c *SpecialKeywordContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitCaseStatement is called when exiting the caseStatement production.
	ExitCaseStatement(c *CaseStatementContext)

	// ExitCaseList is called when exiting the caseList production.
	ExitCaseList(c *CaseListContext)

	// ExitCaseItem is called when exiting the caseItem production.
	ExitCaseItem(c *CaseItemContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitLogical is called when exiting the logical production.
	ExitLogical(c *LogicalContext)

	// ExitRelationalExpression is called when exiting the relationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitAddingExpression is called when exiting the addingExpression production.
	ExitAddingExpression(c *AddingExpressionContext)

	// ExitMultiplyingExpression is called when exiting the multiplyingExpression production.
	ExitMultiplyingExpression(c *MultiplyingExpressionContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitBase is called when exiting the base production.
	ExitBase(c *BaseContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitAtFunction is called when exiting the atFunction production.
	ExitAtFunction(c *AtFunctionContext)

	// ExitDataAccessor is called when exiting the dataAccessor production.
	ExitDataAccessor(c *DataAccessorContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitArgList is called when exiting the argList production.
	ExitArgList(c *ArgListContext)

	// ExitSpecialKeyword is called when exiting the specialKeyword production.
	ExitSpecialKeyword(c *SpecialKeywordContext)
}

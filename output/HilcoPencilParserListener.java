// Generated from /Users/timothyhilgenberg/go/src/github.com/timhilco/go-PencilCalculator/HilcoPencilParser.g4 by ANTLR 4.8
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link HilcoPencilParserParser}.
 */
public interface HilcoPencilParserListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link HilcoPencilParserParser#prog}.
	 * @param ctx the parse tree
	 */
	void enterProg(HilcoPencilParserParser.ProgContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilParserParser#prog}.
	 * @param ctx the parse tree
	 */
	void exitProg(HilcoPencilParserParser.ProgContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilParserParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatement(HilcoPencilParserParser.StatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilParserParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatement(HilcoPencilParserParser.StatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilParserParser#caseStatement}.
	 * @param ctx the parse tree
	 */
	void enterCaseStatement(HilcoPencilParserParser.CaseStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilParserParser#caseStatement}.
	 * @param ctx the parse tree
	 */
	void exitCaseStatement(HilcoPencilParserParser.CaseStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilParserParser#caseList}.
	 * @param ctx the parse tree
	 */
	void enterCaseList(HilcoPencilParserParser.CaseListContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilParserParser#caseList}.
	 * @param ctx the parse tree
	 */
	void exitCaseList(HilcoPencilParserParser.CaseListContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilParserParser#caseItem}.
	 * @param ctx the parse tree
	 */
	void enterCaseItem(HilcoPencilParserParser.CaseItemContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilParserParser#caseItem}.
	 * @param ctx the parse tree
	 */
	void exitCaseItem(HilcoPencilParserParser.CaseItemContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilParserParser#switchStatement}.
	 * @param ctx the parse tree
	 */
	void enterSwitchStatement(HilcoPencilParserParser.SwitchStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilParserParser#switchStatement}.
	 * @param ctx the parse tree
	 */
	void exitSwitchStatement(HilcoPencilParserParser.SwitchStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilParserParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIfStatement(HilcoPencilParserParser.IfStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilParserParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIfStatement(HilcoPencilParserParser.IfStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilParserParser#logical}.
	 * @param ctx the parse tree
	 */
	void enterLogical(HilcoPencilParserParser.LogicalContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilParserParser#logical}.
	 * @param ctx the parse tree
	 */
	void exitLogical(HilcoPencilParserParser.LogicalContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilParserParser#relationalExpression}.
	 * @param ctx the parse tree
	 */
	void enterRelationalExpression(HilcoPencilParserParser.RelationalExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilParserParser#relationalExpression}.
	 * @param ctx the parse tree
	 */
	void exitRelationalExpression(HilcoPencilParserParser.RelationalExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilParserParser#addingExpression}.
	 * @param ctx the parse tree
	 */
	void enterAddingExpression(HilcoPencilParserParser.AddingExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilParserParser#addingExpression}.
	 * @param ctx the parse tree
	 */
	void exitAddingExpression(HilcoPencilParserParser.AddingExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilParserParser#multiplyingExpression}.
	 * @param ctx the parse tree
	 */
	void enterMultiplyingExpression(HilcoPencilParserParser.MultiplyingExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilParserParser#multiplyingExpression}.
	 * @param ctx the parse tree
	 */
	void exitMultiplyingExpression(HilcoPencilParserParser.MultiplyingExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilParserParser#factor}.
	 * @param ctx the parse tree
	 */
	void enterFactor(HilcoPencilParserParser.FactorContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilParserParser#factor}.
	 * @param ctx the parse tree
	 */
	void exitFactor(HilcoPencilParserParser.FactorContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilParserParser#base}.
	 * @param ctx the parse tree
	 */
	void enterBase(HilcoPencilParserParser.BaseContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilParserParser#base}.
	 * @param ctx the parse tree
	 */
	void exitBase(HilcoPencilParserParser.BaseContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilParserParser#name}.
	 * @param ctx the parse tree
	 */
	void enterName(HilcoPencilParserParser.NameContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilParserParser#name}.
	 * @param ctx the parse tree
	 */
	void exitName(HilcoPencilParserParser.NameContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilParserParser#atFunction}.
	 * @param ctx the parse tree
	 */
	void enterAtFunction(HilcoPencilParserParser.AtFunctionContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilParserParser#atFunction}.
	 * @param ctx the parse tree
	 */
	void exitAtFunction(HilcoPencilParserParser.AtFunctionContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilParserParser#dataAccessor}.
	 * @param ctx the parse tree
	 */
	void enterDataAccessor(HilcoPencilParserParser.DataAccessorContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilParserParser#dataAccessor}.
	 * @param ctx the parse tree
	 */
	void exitDataAccessor(HilcoPencilParserParser.DataAccessorContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilParserParser#fieldName}.
	 * @param ctx the parse tree
	 */
	void enterFieldName(HilcoPencilParserParser.FieldNameContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilParserParser#fieldName}.
	 * @param ctx the parse tree
	 */
	void exitFieldName(HilcoPencilParserParser.FieldNameContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilParserParser#argList}.
	 * @param ctx the parse tree
	 */
	void enterArgList(HilcoPencilParserParser.ArgListContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilParserParser#argList}.
	 * @param ctx the parse tree
	 */
	void exitArgList(HilcoPencilParserParser.ArgListContext ctx);
	/**
	 * Enter a parse tree produced by {@link HilcoPencilParserParser#specialKeyword}.
	 * @param ctx the parse tree
	 */
	void enterSpecialKeyword(HilcoPencilParserParser.SpecialKeywordContext ctx);
	/**
	 * Exit a parse tree produced by {@link HilcoPencilParserParser#specialKeyword}.
	 * @param ctx the parse tree
	 */
	void exitSpecialKeyword(HilcoPencilParserParser.SpecialKeywordContext ctx);
}
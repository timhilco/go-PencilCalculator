// Generated from /Users/timothyhilgenberg/go/src/github.com/timhilco/go-PencilCalculator/HilcoPencilParser.g4 by ANTLR 4.8
import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link HilcoPencilParserParser}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface HilcoPencilParserVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link HilcoPencilParserParser#prog}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitProg(HilcoPencilParserParser.ProgContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilParserParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStatement(HilcoPencilParserParser.StatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilParserParser#caseStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCaseStatement(HilcoPencilParserParser.CaseStatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilParserParser#caseList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCaseList(HilcoPencilParserParser.CaseListContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilParserParser#caseItem}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCaseItem(HilcoPencilParserParser.CaseItemContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilParserParser#switchStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSwitchStatement(HilcoPencilParserParser.SwitchStatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilParserParser#ifStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIfStatement(HilcoPencilParserParser.IfStatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilParserParser#logical}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLogical(HilcoPencilParserParser.LogicalContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilParserParser#relationalExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRelationalExpression(HilcoPencilParserParser.RelationalExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilParserParser#addingExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAddingExpression(HilcoPencilParserParser.AddingExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilParserParser#multiplyingExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMultiplyingExpression(HilcoPencilParserParser.MultiplyingExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilParserParser#factor}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFactor(HilcoPencilParserParser.FactorContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilParserParser#base}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBase(HilcoPencilParserParser.BaseContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilParserParser#name}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitName(HilcoPencilParserParser.NameContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilParserParser#atFunction}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAtFunction(HilcoPencilParserParser.AtFunctionContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilParserParser#dataAccessor}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDataAccessor(HilcoPencilParserParser.DataAccessorContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilParserParser#fieldName}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFieldName(HilcoPencilParserParser.FieldNameContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilParserParser#argList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArgList(HilcoPencilParserParser.ArgListContext ctx);
	/**
	 * Visit a parse tree produced by {@link HilcoPencilParserParser#specialKeyword}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSpecialKeyword(HilcoPencilParserParser.SpecialKeywordContext ctx);
}
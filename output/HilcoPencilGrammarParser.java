// Generated from /Users/timothyhilgenberg/go/src/github.com/timhilco/go-PencilCalculator/HilcoPencilGrammarParser.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class HilcoPencilGrammarParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		CASE=1, END_CASE=2, IS=3, SWITCH=4, END_SWITCH=5, IF=6, THEN=7, ELSE=8, 
		DOUBLE_COLON=9, SEMI_COLON=10, NOT=11, EXPONENTIAL=12, MULTIPLY=13, DIVIDE=14, 
		ADD=15, MINUS=16, GREATER_THAN=17, GREATER_THAN_EQUAL=18, LESS_THAN=19, 
		LESS_THAN_EQUAL=20, EQUAL=21, NOT_EQUAL=22, AND=23, OR=24, COLON=25, LBRACKET=26, 
		RBRACKET=27, CURLYLBRACKET=28, CURLYRBRACKET=29, LPAREN=30, RPAREN=31, 
		UNDERSCORE=32, PERCENT_SIGN=33, AT_SIGN=34, COMMA=35, DOT=36, KEYWORD_TRUE=37, 
		KEYWORD_FALSE=38, KEYWORD_NIL=39, CLASSNAME=40, ID=41, ATFUNCTION=42, 
		INT=43, STRING=44, FLOAT=45, PERCENT=46, DATE=47, NEWLINE=48, WS=49;
	public static final int
		RULE_program = 0, RULE_expression = 1, RULE_caseStatement = 2, RULE_caseList = 3, 
		RULE_caseItem = 4, RULE_ifStatement = 5, RULE_name = 6, RULE_atFunction = 7, 
		RULE_argList = 8, RULE_dataAccessor = 9, RULE_accessList = 10, RULE_accessorMessage = 11, 
		RULE_specialKeyword = 12;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "expression", "caseStatement", "caseList", "caseItem", "ifStatement", 
			"name", "atFunction", "argList", "dataAccessor", "accessList", "accessorMessage", 
			"specialKeyword"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'case'", "'endcase'", "'is'", "'switch'", "'endswitch'", "'if'", 
			"'then'", "'else'", "'::'", "';'", "'NOT'", "'^'", "'*'", "'/'", "'+'", 
			"'-'", "'>'", "'>='", "'<'", "'<='", "'='", "'~='", "'AND'", "'OR'", 
			"':'", "'['", "']'", "'{'", "'}'", "'('", "')'", "'_'", "'%'", "'@'", 
			"','", "'.'", "'true'", "'false'", "'nil'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "CASE", "END_CASE", "IS", "SWITCH", "END_SWITCH", "IF", "THEN", 
			"ELSE", "DOUBLE_COLON", "SEMI_COLON", "NOT", "EXPONENTIAL", "MULTIPLY", 
			"DIVIDE", "ADD", "MINUS", "GREATER_THAN", "GREATER_THAN_EQUAL", "LESS_THAN", 
			"LESS_THAN_EQUAL", "EQUAL", "NOT_EQUAL", "AND", "OR", "COLON", "LBRACKET", 
			"RBRACKET", "CURLYLBRACKET", "CURLYRBRACKET", "LPAREN", "RPAREN", "UNDERSCORE", 
			"PERCENT_SIGN", "AT_SIGN", "COMMA", "DOT", "KEYWORD_TRUE", "KEYWORD_FALSE", 
			"KEYWORD_NIL", "CLASSNAME", "ID", "ATFUNCTION", "INT", "STRING", "FLOAT", 
			"PERCENT", "DATE", "NEWLINE", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "HilcoPencilGrammarParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public HilcoPencilGrammarParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ProgramContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterProgram(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitProgram(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitProgram(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(26);
			expression(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionContext extends ParserRuleContext {
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	 
		public ExpressionContext() { }
		public void copyFrom(ExpressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class NameCalculatorContext extends ExpressionContext {
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public NameCalculatorContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterNameCalculator(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitNameCalculator(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitNameCalculator(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class UnaryNotCalculatorContext extends ExpressionContext {
		public TerminalNode NOT() { return getToken(HilcoPencilGrammarParser.NOT, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public UnaryNotCalculatorContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterUnaryNotCalculator(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitUnaryNotCalculator(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitUnaryNotCalculator(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class BinaryArthmeticCalculatorContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode MULTIPLY() { return getToken(HilcoPencilGrammarParser.MULTIPLY, 0); }
		public TerminalNode DIVIDE() { return getToken(HilcoPencilGrammarParser.DIVIDE, 0); }
		public TerminalNode ADD() { return getToken(HilcoPencilGrammarParser.ADD, 0); }
		public TerminalNode MINUS() { return getToken(HilcoPencilGrammarParser.MINUS, 0); }
		public BinaryArthmeticCalculatorContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterBinaryArthmeticCalculator(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitBinaryArthmeticCalculator(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitBinaryArthmeticCalculator(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class PercentContext extends ExpressionContext {
		public TerminalNode PERCENT() { return getToken(HilcoPencilGrammarParser.PERCENT, 0); }
		public PercentContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterPercent(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitPercent(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitPercent(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class ParensContext extends ExpressionContext {
		public TerminalNode LPAREN() { return getToken(HilcoPencilGrammarParser.LPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(HilcoPencilGrammarParser.RPAREN, 0); }
		public ParensContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterParens(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitParens(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitParens(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class BinaryRelationalCalculatorContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode GREATER_THAN() { return getToken(HilcoPencilGrammarParser.GREATER_THAN, 0); }
		public TerminalNode GREATER_THAN_EQUAL() { return getToken(HilcoPencilGrammarParser.GREATER_THAN_EQUAL, 0); }
		public TerminalNode LESS_THAN() { return getToken(HilcoPencilGrammarParser.LESS_THAN, 0); }
		public TerminalNode LESS_THAN_EQUAL() { return getToken(HilcoPencilGrammarParser.LESS_THAN_EQUAL, 0); }
		public TerminalNode EQUAL() { return getToken(HilcoPencilGrammarParser.EQUAL, 0); }
		public TerminalNode NOT_EQUAL() { return getToken(HilcoPencilGrammarParser.NOT_EQUAL, 0); }
		public BinaryRelationalCalculatorContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterBinaryRelationalCalculator(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitBinaryRelationalCalculator(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitBinaryRelationalCalculator(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class StringContext extends ExpressionContext {
		public TerminalNode STRING() { return getToken(HilcoPencilGrammarParser.STRING, 0); }
		public StringContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterString(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitString(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitString(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class ExpressionAtFunctionContext extends ExpressionContext {
		public AtFunctionContext atFunction() {
			return getRuleContext(AtFunctionContext.class,0);
		}
		public ExpressionAtFunctionContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterExpressionAtFunction(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitExpressionAtFunction(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitExpressionAtFunction(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class DateContext extends ExpressionContext {
		public TerminalNode DATE() { return getToken(HilcoPencilGrammarParser.DATE, 0); }
		public DateContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterDate(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitDate(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitDate(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class CaseContext extends ExpressionContext {
		public CaseStatementContext caseStatement() {
			return getRuleContext(CaseStatementContext.class,0);
		}
		public CaseContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterCase(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitCase(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitCase(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class IntegerContext extends ExpressionContext {
		public TerminalNode INT() { return getToken(HilcoPencilGrammarParser.INT, 0); }
		public IntegerContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterInteger(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitInteger(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitInteger(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class BinaryLogicalCalculatorContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode AND() { return getToken(HilcoPencilGrammarParser.AND, 0); }
		public TerminalNode OR() { return getToken(HilcoPencilGrammarParser.OR, 0); }
		public BinaryLogicalCalculatorContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterBinaryLogicalCalculator(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitBinaryLogicalCalculator(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitBinaryLogicalCalculator(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class ExpressionKeywordContext extends ExpressionContext {
		public SpecialKeywordContext specialKeyword() {
			return getRuleContext(SpecialKeywordContext.class,0);
		}
		public ExpressionKeywordContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterExpressionKeyword(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitExpressionKeyword(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitExpressionKeyword(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class FloatContext extends ExpressionContext {
		public TerminalNode FLOAT() { return getToken(HilcoPencilGrammarParser.FLOAT, 0); }
		public FloatContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterFloat(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitFloat(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitFloat(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class UnaryMinusCalculatorContext extends ExpressionContext {
		public TerminalNode MINUS() { return getToken(HilcoPencilGrammarParser.MINUS, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public UnaryMinusCalculatorContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterUnaryMinusCalculator(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitUnaryMinusCalculator(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitUnaryMinusCalculator(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class BinaryExponentialCalculatorContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode EXPONENTIAL() { return getToken(HilcoPencilGrammarParser.EXPONENTIAL, 0); }
		public BinaryExponentialCalculatorContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterBinaryExponentialCalculator(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitBinaryExponentialCalculator(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitBinaryExponentialCalculator(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class ExpressionDataAccessContext extends ExpressionContext {
		public DataAccessorContext dataAccessor() {
			return getRuleContext(DataAccessorContext.class,0);
		}
		public ExpressionDataAccessContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterExpressionDataAccess(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitExpressionDataAccess(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitExpressionDataAccess(this);
			else return visitor.visitChildren(this);
		}
	}
	public static class IfContext extends ExpressionContext {
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public IfContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterIf(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitIf(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitIf(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 2;
		enterRecursionRule(_localctx, 2, RULE_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(48);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
			case 1:
				{
				_localctx = new CaseContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(29);
				caseStatement();
				}
				break;
			case 2:
				{
				_localctx = new IfContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(30);
				ifStatement();
				}
				break;
			case 3:
				{
				_localctx = new ParensContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(31);
				match(LPAREN);
				setState(32);
				expression(0);
				setState(33);
				match(RPAREN);
				}
				break;
			case 4:
				{
				_localctx = new UnaryMinusCalculatorContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(35);
				match(MINUS);
				setState(36);
				expression(24);
				}
				break;
			case 5:
				{
				_localctx = new UnaryNotCalculatorContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(37);
				match(NOT);
				setState(38);
				expression(23);
				}
				break;
			case 6:
				{
				_localctx = new NameCalculatorContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(39);
				name();
				}
				break;
			case 7:
				{
				_localctx = new ExpressionAtFunctionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(40);
				atFunction();
				}
				break;
			case 8:
				{
				_localctx = new ExpressionDataAccessContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(41);
				dataAccessor();
				}
				break;
			case 9:
				{
				_localctx = new ExpressionKeywordContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(42);
				specialKeyword();
				}
				break;
			case 10:
				{
				_localctx = new PercentContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(43);
				match(PERCENT);
				}
				break;
			case 11:
				{
				_localctx = new IntegerContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(44);
				match(INT);
				}
				break;
			case 12:
				{
				_localctx = new FloatContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(45);
				match(FLOAT);
				}
				break;
			case 13:
				{
				_localctx = new StringContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(46);
				match(STRING);
				}
				break;
			case 14:
				{
				_localctx = new DateContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(47);
				match(DATE);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(91);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(89);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
					case 1:
						{
						_localctx = new BinaryExponentialCalculatorContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(50);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(51);
						match(EXPONENTIAL);
						setState(52);
						expression(22);
						}
						break;
					case 2:
						{
						_localctx = new BinaryArthmeticCalculatorContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(53);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(54);
						match(MULTIPLY);
						setState(55);
						expression(22);
						}
						break;
					case 3:
						{
						_localctx = new BinaryArthmeticCalculatorContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(56);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(57);
						match(DIVIDE);
						setState(58);
						expression(21);
						}
						break;
					case 4:
						{
						_localctx = new BinaryArthmeticCalculatorContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(59);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(60);
						match(ADD);
						setState(61);
						expression(20);
						}
						break;
					case 5:
						{
						_localctx = new BinaryArthmeticCalculatorContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(62);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(63);
						match(MINUS);
						setState(64);
						expression(19);
						}
						break;
					case 6:
						{
						_localctx = new BinaryRelationalCalculatorContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(65);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(66);
						match(GREATER_THAN);
						setState(67);
						expression(18);
						}
						break;
					case 7:
						{
						_localctx = new BinaryRelationalCalculatorContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(68);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(69);
						match(GREATER_THAN_EQUAL);
						setState(70);
						expression(17);
						}
						break;
					case 8:
						{
						_localctx = new BinaryRelationalCalculatorContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(71);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(72);
						match(LESS_THAN);
						setState(73);
						expression(16);
						}
						break;
					case 9:
						{
						_localctx = new BinaryRelationalCalculatorContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(74);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(75);
						match(LESS_THAN_EQUAL);
						setState(76);
						expression(15);
						}
						break;
					case 10:
						{
						_localctx = new BinaryRelationalCalculatorContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(77);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(78);
						match(EQUAL);
						setState(79);
						expression(14);
						}
						break;
					case 11:
						{
						_localctx = new BinaryRelationalCalculatorContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(80);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(81);
						match(NOT_EQUAL);
						setState(82);
						expression(13);
						}
						break;
					case 12:
						{
						_localctx = new BinaryLogicalCalculatorContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(83);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(84);
						match(AND);
						setState(85);
						expression(12);
						}
						break;
					case 13:
						{
						_localctx = new BinaryLogicalCalculatorContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(86);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(87);
						match(OR);
						setState(88);
						expression(11);
						}
						break;
					}
					} 
				}
				setState(93);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class CaseStatementContext extends ParserRuleContext {
		public TerminalNode CASE() { return getToken(HilcoPencilGrammarParser.CASE, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode IS() { return getToken(HilcoPencilGrammarParser.IS, 0); }
		public CaseListContext caseList() {
			return getRuleContext(CaseListContext.class,0);
		}
		public TerminalNode END_CASE() { return getToken(HilcoPencilGrammarParser.END_CASE, 0); }
		public CaseStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_caseStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterCaseStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitCaseStatement(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitCaseStatement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final CaseStatementContext caseStatement() throws RecognitionException {
		CaseStatementContext _localctx = new CaseStatementContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_caseStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(94);
			match(CASE);
			setState(95);
			expression(0);
			setState(96);
			match(IS);
			setState(97);
			caseList();
			setState(98);
			match(END_CASE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CaseListContext extends ParserRuleContext {
		public CaseItemContext caseItem() {
			return getRuleContext(CaseItemContext.class,0);
		}
		public CaseListContext caseList() {
			return getRuleContext(CaseListContext.class,0);
		}
		public CaseListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_caseList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterCaseList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitCaseList(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitCaseList(this);
			else return visitor.visitChildren(this);
		}
	}

	public final CaseListContext caseList() throws RecognitionException {
		CaseListContext _localctx = new CaseListContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_caseList);
		try {
			setState(104);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CASE:
			case IF:
			case NOT:
			case MINUS:
			case LPAREN:
			case KEYWORD_TRUE:
			case KEYWORD_FALSE:
			case KEYWORD_NIL:
			case CLASSNAME:
			case ID:
			case ATFUNCTION:
			case INT:
			case STRING:
			case FLOAT:
			case PERCENT:
			case DATE:
				enterOuterAlt(_localctx, 1);
				{
				setState(100);
				caseItem();
				setState(101);
				caseList();
				}
				break;
			case END_CASE:
				enterOuterAlt(_localctx, 2);
				{
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CaseItemContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode DOUBLE_COLON() { return getToken(HilcoPencilGrammarParser.DOUBLE_COLON, 0); }
		public TerminalNode SEMI_COLON() { return getToken(HilcoPencilGrammarParser.SEMI_COLON, 0); }
		public CaseItemContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_caseItem; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterCaseItem(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitCaseItem(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitCaseItem(this);
			else return visitor.visitChildren(this);
		}
	}

	public final CaseItemContext caseItem() throws RecognitionException {
		CaseItemContext _localctx = new CaseItemContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_caseItem);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(106);
			expression(0);
			setState(107);
			match(DOUBLE_COLON);
			setState(108);
			expression(0);
			setState(109);
			match(SEMI_COLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IfStatementContext extends ParserRuleContext {
		public TerminalNode IF() { return getToken(HilcoPencilGrammarParser.IF, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode THEN() { return getToken(HilcoPencilGrammarParser.THEN, 0); }
		public TerminalNode ELSE() { return getToken(HilcoPencilGrammarParser.ELSE, 0); }
		public IfStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterIfStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitIfStatement(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitIfStatement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final IfStatementContext ifStatement() throws RecognitionException {
		IfStatementContext _localctx = new IfStatementContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_ifStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(111);
			match(IF);
			setState(112);
			expression(0);
			setState(113);
			match(THEN);
			setState(114);
			expression(0);
			setState(117);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				{
				setState(115);
				match(ELSE);
				setState(116);
				expression(0);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NameContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(HilcoPencilGrammarParser.ID, 0); }
		public TerminalNode CLASSNAME() { return getToken(HilcoPencilGrammarParser.CLASSNAME, 0); }
		public TerminalNode COLON() { return getToken(HilcoPencilGrammarParser.COLON, 0); }
		public NameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_name; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterName(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitName(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitName(this);
			else return visitor.visitChildren(this);
		}
	}

	public final NameContext name() throws RecognitionException {
		NameContext _localctx = new NameContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_name);
		try {
			setState(123);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ID:
				enterOuterAlt(_localctx, 1);
				{
				setState(119);
				match(ID);
				}
				break;
			case CLASSNAME:
				enterOuterAlt(_localctx, 2);
				{
				setState(120);
				match(CLASSNAME);
				setState(121);
				match(COLON);
				setState(122);
				match(ID);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AtFunctionContext extends ParserRuleContext {
		public TerminalNode ATFUNCTION() { return getToken(HilcoPencilGrammarParser.ATFUNCTION, 0); }
		public TerminalNode LPAREN() { return getToken(HilcoPencilGrammarParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(HilcoPencilGrammarParser.RPAREN, 0); }
		public ArgListContext argList() {
			return getRuleContext(ArgListContext.class,0);
		}
		public AtFunctionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_atFunction; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterAtFunction(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitAtFunction(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitAtFunction(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AtFunctionContext atFunction() throws RecognitionException {
		AtFunctionContext _localctx = new AtFunctionContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_atFunction);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(125);
			match(ATFUNCTION);
			setState(126);
			match(LPAREN);
			setState(128);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << CASE) | (1L << IF) | (1L << NOT) | (1L << MINUS) | (1L << LPAREN) | (1L << KEYWORD_TRUE) | (1L << KEYWORD_FALSE) | (1L << KEYWORD_NIL) | (1L << CLASSNAME) | (1L << ID) | (1L << ATFUNCTION) | (1L << INT) | (1L << STRING) | (1L << FLOAT) | (1L << PERCENT) | (1L << DATE))) != 0)) {
				{
				setState(127);
				argList();
				}
			}

			setState(130);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArgListContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(HilcoPencilGrammarParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(HilcoPencilGrammarParser.COMMA, i);
		}
		public ArgListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterArgList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitArgList(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitArgList(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ArgListContext argList() throws RecognitionException {
		ArgListContext _localctx = new ArgListContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_argList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(132);
			expression(0);
			setState(139);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(133);
				match(COMMA);
				setState(135);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << CASE) | (1L << IF) | (1L << NOT) | (1L << MINUS) | (1L << LPAREN) | (1L << KEYWORD_TRUE) | (1L << KEYWORD_FALSE) | (1L << KEYWORD_NIL) | (1L << CLASSNAME) | (1L << ID) | (1L << ATFUNCTION) | (1L << INT) | (1L << STRING) | (1L << FLOAT) | (1L << PERCENT) | (1L << DATE))) != 0)) {
					{
					setState(134);
					expression(0);
					}
				}

				}
				}
				setState(141);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DataAccessorContext extends ParserRuleContext {
		public TerminalNode CLASSNAME() { return getToken(HilcoPencilGrammarParser.CLASSNAME, 0); }
		public List<AccessListContext> accessList() {
			return getRuleContexts(AccessListContext.class);
		}
		public AccessListContext accessList(int i) {
			return getRuleContext(AccessListContext.class,i);
		}
		public DataAccessorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dataAccessor; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterDataAccessor(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitDataAccessor(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitDataAccessor(this);
			else return visitor.visitChildren(this);
		}
	}

	public final DataAccessorContext dataAccessor() throws RecognitionException {
		DataAccessorContext _localctx = new DataAccessorContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_dataAccessor);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(142);
			match(CLASSNAME);
			setState(146);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(143);
					accessList();
					}
					} 
				}
				setState(148);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AccessListContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(HilcoPencilGrammarParser.DOT, 0); }
		public AccessorMessageContext accessorMessage() {
			return getRuleContext(AccessorMessageContext.class,0);
		}
		public TerminalNode CLASSNAME() { return getToken(HilcoPencilGrammarParser.CLASSNAME, 0); }
		public AccessListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accessList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterAccessList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitAccessList(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitAccessList(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AccessListContext accessList() throws RecognitionException {
		AccessListContext _localctx = new AccessListContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_accessList);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(149);
			match(DOT);
			setState(152);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ID:
				{
				setState(150);
				accessorMessage();
				}
				break;
			case CLASSNAME:
				{
				setState(151);
				match(CLASSNAME);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AccessorMessageContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(HilcoPencilGrammarParser.ID, 0); }
		public TerminalNode LPAREN() { return getToken(HilcoPencilGrammarParser.LPAREN, 0); }
		public ArgListContext argList() {
			return getRuleContext(ArgListContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(HilcoPencilGrammarParser.RPAREN, 0); }
		public AccessorMessageContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accessorMessage; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterAccessorMessage(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitAccessorMessage(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitAccessorMessage(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AccessorMessageContext accessorMessage() throws RecognitionException {
		AccessorMessageContext _localctx = new AccessorMessageContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_accessorMessage);
		try {
			setState(160);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(154);
				match(ID);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(155);
				match(ID);
				setState(156);
				match(LPAREN);
				setState(157);
				argList();
				setState(158);
				match(RPAREN);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SpecialKeywordContext extends ParserRuleContext {
		public TerminalNode KEYWORD_TRUE() { return getToken(HilcoPencilGrammarParser.KEYWORD_TRUE, 0); }
		public TerminalNode KEYWORD_FALSE() { return getToken(HilcoPencilGrammarParser.KEYWORD_FALSE, 0); }
		public TerminalNode KEYWORD_NIL() { return getToken(HilcoPencilGrammarParser.KEYWORD_NIL, 0); }
		public SpecialKeywordContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_specialKeyword; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).enterSpecialKeyword(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilGrammarParserListener ) ((HilcoPencilGrammarParserListener)listener).exitSpecialKeyword(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilGrammarParserVisitor ) return ((HilcoPencilGrammarParserVisitor<? extends T>)visitor).visitSpecialKeyword(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SpecialKeywordContext specialKeyword() throws RecognitionException {
		SpecialKeywordContext _localctx = new SpecialKeywordContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_specialKeyword);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(162);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << KEYWORD_TRUE) | (1L << KEYWORD_FALSE) | (1L << KEYWORD_NIL))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 1:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 22);
		case 1:
			return precpred(_ctx, 21);
		case 2:
			return precpred(_ctx, 20);
		case 3:
			return precpred(_ctx, 19);
		case 4:
			return precpred(_ctx, 18);
		case 5:
			return precpred(_ctx, 17);
		case 6:
			return precpred(_ctx, 16);
		case 7:
			return precpred(_ctx, 15);
		case 8:
			return precpred(_ctx, 14);
		case 9:
			return precpred(_ctx, 13);
		case 10:
			return precpred(_ctx, 12);
		case 11:
			return precpred(_ctx, 11);
		case 12:
			return precpred(_ctx, 10);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\63\u00a7\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\5\3\63\n\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\3\7\3\\\n\3\f\3\16\3_\13\3\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5"+
		"\3\5\5\5k\n\5\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\5\7x\n\7\3\b"+
		"\3\b\3\b\3\b\5\b~\n\b\3\t\3\t\3\t\5\t\u0083\n\t\3\t\3\t\3\n\3\n\3\n\5"+
		"\n\u008a\n\n\7\n\u008c\n\n\f\n\16\n\u008f\13\n\3\13\3\13\7\13\u0093\n"+
		"\13\f\13\16\13\u0096\13\13\3\f\3\f\3\f\5\f\u009b\n\f\3\r\3\r\3\r\3\r\3"+
		"\r\3\r\5\r\u00a3\n\r\3\16\3\16\3\16\2\3\4\17\2\4\6\b\n\f\16\20\22\24\26"+
		"\30\32\2\3\3\2\')\2\u00bc\2\34\3\2\2\2\4\62\3\2\2\2\6`\3\2\2\2\bj\3\2"+
		"\2\2\nl\3\2\2\2\fq\3\2\2\2\16}\3\2\2\2\20\177\3\2\2\2\22\u0086\3\2\2\2"+
		"\24\u0090\3\2\2\2\26\u0097\3\2\2\2\30\u00a2\3\2\2\2\32\u00a4\3\2\2\2\34"+
		"\35\5\4\3\2\35\3\3\2\2\2\36\37\b\3\1\2\37\63\5\6\4\2 \63\5\f\7\2!\"\7"+
		" \2\2\"#\5\4\3\2#$\7!\2\2$\63\3\2\2\2%&\7\22\2\2&\63\5\4\3\32\'(\7\r\2"+
		"\2(\63\5\4\3\31)\63\5\16\b\2*\63\5\20\t\2+\63\5\24\13\2,\63\5\32\16\2"+
		"-\63\7\60\2\2.\63\7-\2\2/\63\7/\2\2\60\63\7.\2\2\61\63\7\61\2\2\62\36"+
		"\3\2\2\2\62 \3\2\2\2\62!\3\2\2\2\62%\3\2\2\2\62\'\3\2\2\2\62)\3\2\2\2"+
		"\62*\3\2\2\2\62+\3\2\2\2\62,\3\2\2\2\62-\3\2\2\2\62.\3\2\2\2\62/\3\2\2"+
		"\2\62\60\3\2\2\2\62\61\3\2\2\2\63]\3\2\2\2\64\65\f\30\2\2\65\66\7\16\2"+
		"\2\66\\\5\4\3\30\678\f\27\2\289\7\17\2\29\\\5\4\3\30:;\f\26\2\2;<\7\20"+
		"\2\2<\\\5\4\3\27=>\f\25\2\2>?\7\21\2\2?\\\5\4\3\26@A\f\24\2\2AB\7\22\2"+
		"\2B\\\5\4\3\25CD\f\23\2\2DE\7\23\2\2E\\\5\4\3\24FG\f\22\2\2GH\7\24\2\2"+
		"H\\\5\4\3\23IJ\f\21\2\2JK\7\25\2\2K\\\5\4\3\22LM\f\20\2\2MN\7\26\2\2N"+
		"\\\5\4\3\21OP\f\17\2\2PQ\7\27\2\2Q\\\5\4\3\20RS\f\16\2\2ST\7\30\2\2T\\"+
		"\5\4\3\17UV\f\r\2\2VW\7\31\2\2W\\\5\4\3\16XY\f\f\2\2YZ\7\32\2\2Z\\\5\4"+
		"\3\r[\64\3\2\2\2[\67\3\2\2\2[:\3\2\2\2[=\3\2\2\2[@\3\2\2\2[C\3\2\2\2["+
		"F\3\2\2\2[I\3\2\2\2[L\3\2\2\2[O\3\2\2\2[R\3\2\2\2[U\3\2\2\2[X\3\2\2\2"+
		"\\_\3\2\2\2][\3\2\2\2]^\3\2\2\2^\5\3\2\2\2_]\3\2\2\2`a\7\3\2\2ab\5\4\3"+
		"\2bc\7\5\2\2cd\5\b\5\2de\7\4\2\2e\7\3\2\2\2fg\5\n\6\2gh\5\b\5\2hk\3\2"+
		"\2\2ik\3\2\2\2jf\3\2\2\2ji\3\2\2\2k\t\3\2\2\2lm\5\4\3\2mn\7\13\2\2no\5"+
		"\4\3\2op\7\f\2\2p\13\3\2\2\2qr\7\b\2\2rs\5\4\3\2st\7\t\2\2tw\5\4\3\2u"+
		"v\7\n\2\2vx\5\4\3\2wu\3\2\2\2wx\3\2\2\2x\r\3\2\2\2y~\7+\2\2z{\7*\2\2{"+
		"|\7\33\2\2|~\7+\2\2}y\3\2\2\2}z\3\2\2\2~\17\3\2\2\2\177\u0080\7,\2\2\u0080"+
		"\u0082\7 \2\2\u0081\u0083\5\22\n\2\u0082\u0081\3\2\2\2\u0082\u0083\3\2"+
		"\2\2\u0083\u0084\3\2\2\2\u0084\u0085\7!\2\2\u0085\21\3\2\2\2\u0086\u008d"+
		"\5\4\3\2\u0087\u0089\7%\2\2\u0088\u008a\5\4\3\2\u0089\u0088\3\2\2\2\u0089"+
		"\u008a\3\2\2\2\u008a\u008c\3\2\2\2\u008b\u0087\3\2\2\2\u008c\u008f\3\2"+
		"\2\2\u008d\u008b\3\2\2\2\u008d\u008e\3\2\2\2\u008e\23\3\2\2\2\u008f\u008d"+
		"\3\2\2\2\u0090\u0094\7*\2\2\u0091\u0093\5\26\f\2\u0092\u0091\3\2\2\2\u0093"+
		"\u0096\3\2\2\2\u0094\u0092\3\2\2\2\u0094\u0095\3\2\2\2\u0095\25\3\2\2"+
		"\2\u0096\u0094\3\2\2\2\u0097\u009a\7&\2\2\u0098\u009b\5\30\r\2\u0099\u009b"+
		"\7*\2\2\u009a\u0098\3\2\2\2\u009a\u0099\3\2\2\2\u009b\27\3\2\2\2\u009c"+
		"\u00a3\7+\2\2\u009d\u009e\7+\2\2\u009e\u009f\7 \2\2\u009f\u00a0\5\22\n"+
		"\2\u00a0\u00a1\7!\2\2\u00a1\u00a3\3\2\2\2\u00a2\u009c\3\2\2\2\u00a2\u009d"+
		"\3\2\2\2\u00a3\31\3\2\2\2\u00a4\u00a5\t\2\2\2\u00a5\33\3\2\2\2\16\62["+
		"]jw}\u0082\u0089\u008d\u0094\u009a\u00a2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
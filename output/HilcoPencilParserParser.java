// Generated from /Users/timothyhilgenberg/go/src/github.com/timhilco/go-PencilCalculator/HilcoPencilParser.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class HilcoPencilParserParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, IF=7, THEN=8, ELSE=9, 
		CASE=10, CASE_IS=11, CASE_END=12, SWITCH=13, SWITCH_END=14, AND=15, OR=16, 
		NOT=17, INT=18, ID=19, STRING=20, FLOAT=21, DATE=22, ATFUNCTION=23, CLASSNAME=24, 
		PERCENT=25, NEWLINE=26, WS=27, DOT=28, COLON=29, SEMI=30, COMMA=31, EQUALS=32, 
		LBRACKET=33, RBRACKET=34, CURLYLBRACKET=35, CURLYRBRACKET=36, DOTDOT=37, 
		LPAREN=38, RPAREN=39, NOT_EQUALS=40, LT=41, LTE=42, GT=43, GTE=44, PLUS=45, 
		MINUS=46, TIMES=47, DIV=48, EXPONENTIATE=49;
	public static final int
		RULE_prog = 0, RULE_statement = 1, RULE_caseStatement = 2, RULE_caseList = 3, 
		RULE_caseItem = 4, RULE_switchStatement = 5, RULE_ifStatement = 6, RULE_logical = 7, 
		RULE_relationalExpression = 8, RULE_addingExpression = 9, RULE_multiplyingExpression = 10, 
		RULE_factor = 11, RULE_base = 12, RULE_name = 13, RULE_atFunction = 14, 
		RULE_dataAccessor = 15, RULE_fieldName = 16, RULE_argList = 17, RULE_specialKeyword = 18;
	private static String[] makeRuleNames() {
		return new String[] {
			"prog", "statement", "caseStatement", "caseList", "caseItem", "switchStatement", 
			"ifStatement", "logical", "relationalExpression", "addingExpression", 
			"multiplyingExpression", "factor", "base", "name", "atFunction", "dataAccessor", 
			"fieldName", "argList", "specialKeyword"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'::'", "'today'", "'true'", "'false'", "'nil'", "'default'", "'if'", 
			"'then'", "'else'", "'case'", "'is'", "'endcase'", "'switch'", "'endswitch'", 
			"'AND'", "'OR'", "'NOT'", null, null, null, null, null, null, null, null, 
			null, null, "'.'", "':'", "';'", "','", "'='", "'['", "']'", "'{'", "'}'", 
			"'..'", "'('", "')'", "'~='", "'<'", "'<='", "'>'", "'>='", "'+'", "'-'", 
			"'*'", "'/'", "'^'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, "IF", "THEN", "ELSE", "CASE", 
			"CASE_IS", "CASE_END", "SWITCH", "SWITCH_END", "AND", "OR", "NOT", "INT", 
			"ID", "STRING", "FLOAT", "DATE", "ATFUNCTION", "CLASSNAME", "PERCENT", 
			"NEWLINE", "WS", "DOT", "COLON", "SEMI", "COMMA", "EQUALS", "LBRACKET", 
			"RBRACKET", "CURLYLBRACKET", "CURLYRBRACKET", "DOTDOT", "LPAREN", "RPAREN", 
			"NOT_EQUALS", "LT", "LTE", "GT", "GTE", "PLUS", "MINUS", "TIMES", "DIV", 
			"EXPONENTIATE"
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
	public String getGrammarFileName() { return "HilcoPencilParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public HilcoPencilParserParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ProgContext extends ParserRuleContext {
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public ProgContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prog; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).enterProg(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).exitProg(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilParserVisitor ) return ((HilcoPencilParserVisitor<? extends T>)visitor).visitProg(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ProgContext prog() throws RecognitionException {
		ProgContext _localctx = new ProgContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_prog);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(38);
			statement();
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

	public static class StatementContext extends ParserRuleContext {
		public LogicalContext logical() {
			return getRuleContext(LogicalContext.class,0);
		}
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public CaseStatementContext caseStatement() {
			return getRuleContext(CaseStatementContext.class,0);
		}
		public SwitchStatementContext switchStatement() {
			return getRuleContext(SwitchStatementContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).enterStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).exitStatement(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilParserVisitor ) return ((HilcoPencilParserVisitor<? extends T>)visitor).visitStatement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_statement);
		try {
			setState(44);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__1:
			case T__2:
			case T__3:
			case T__4:
			case T__5:
			case NOT:
			case INT:
			case ID:
			case STRING:
			case FLOAT:
			case DATE:
			case ATFUNCTION:
			case CLASSNAME:
			case PERCENT:
			case LPAREN:
			case MINUS:
				enterOuterAlt(_localctx, 1);
				{
				setState(40);
				logical();
				}
				break;
			case IF:
				enterOuterAlt(_localctx, 2);
				{
				setState(41);
				ifStatement();
				}
				break;
			case CASE:
				enterOuterAlt(_localctx, 3);
				{
				setState(42);
				caseStatement();
				}
				break;
			case SWITCH:
				enterOuterAlt(_localctx, 4);
				{
				setState(43);
				switchStatement();
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

	public static class CaseStatementContext extends ParserRuleContext {
		public TerminalNode CASE() { return getToken(HilcoPencilParserParser.CASE, 0); }
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public TerminalNode CASE_IS() { return getToken(HilcoPencilParserParser.CASE_IS, 0); }
		public CaseListContext caseList() {
			return getRuleContext(CaseListContext.class,0);
		}
		public TerminalNode CASE_END() { return getToken(HilcoPencilParserParser.CASE_END, 0); }
		public CaseStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_caseStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).enterCaseStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).exitCaseStatement(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilParserVisitor ) return ((HilcoPencilParserVisitor<? extends T>)visitor).visitCaseStatement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final CaseStatementContext caseStatement() throws RecognitionException {
		CaseStatementContext _localctx = new CaseStatementContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_caseStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(46);
			match(CASE);
			setState(47);
			statement();
			setState(48);
			match(CASE_IS);
			setState(49);
			caseList();
			setState(50);
			match(CASE_END);
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
		public List<CaseItemContext> caseItem() {
			return getRuleContexts(CaseItemContext.class);
		}
		public CaseItemContext caseItem(int i) {
			return getRuleContext(CaseItemContext.class,i);
		}
		public CaseListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_caseList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).enterCaseList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).exitCaseList(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilParserVisitor ) return ((HilcoPencilParserVisitor<? extends T>)visitor).visitCaseList(this);
			else return visitor.visitChildren(this);
		}
	}

	public final CaseListContext caseList() throws RecognitionException {
		CaseListContext _localctx = new CaseListContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_caseList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(52);
			caseItem();
			setState(56);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__1) | (1L << T__2) | (1L << T__3) | (1L << T__4) | (1L << T__5) | (1L << INT) | (1L << ID) | (1L << STRING) | (1L << FLOAT) | (1L << DATE) | (1L << ATFUNCTION) | (1L << CLASSNAME) | (1L << PERCENT) | (1L << LPAREN))) != 0)) {
				{
				{
				setState(53);
				caseItem();
				}
				}
				setState(58);
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

	public static class CaseItemContext extends ParserRuleContext {
		public BaseContext base() {
			return getRuleContext(BaseContext.class,0);
		}
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(HilcoPencilParserParser.SEMI, 0); }
		public CaseItemContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_caseItem; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).enterCaseItem(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).exitCaseItem(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilParserVisitor ) return ((HilcoPencilParserVisitor<? extends T>)visitor).visitCaseItem(this);
			else return visitor.visitChildren(this);
		}
	}

	public final CaseItemContext caseItem() throws RecognitionException {
		CaseItemContext _localctx = new CaseItemContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_caseItem);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(59);
			base();
			setState(60);
			match(T__0);
			setState(61);
			statement();
			setState(62);
			match(SEMI);
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

	public static class SwitchStatementContext extends ParserRuleContext {
		public TerminalNode SWITCH() { return getToken(HilcoPencilParserParser.SWITCH, 0); }
		public CaseListContext caseList() {
			return getRuleContext(CaseListContext.class,0);
		}
		public TerminalNode SWITCH_END() { return getToken(HilcoPencilParserParser.SWITCH_END, 0); }
		public SwitchStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switchStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).enterSwitchStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).exitSwitchStatement(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilParserVisitor ) return ((HilcoPencilParserVisitor<? extends T>)visitor).visitSwitchStatement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SwitchStatementContext switchStatement() throws RecognitionException {
		SwitchStatementContext _localctx = new SwitchStatementContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_switchStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(64);
			match(SWITCH);
			setState(65);
			caseList();
			setState(66);
			match(SWITCH_END);
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
		public TerminalNode IF() { return getToken(HilcoPencilParserParser.IF, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public TerminalNode THEN() { return getToken(HilcoPencilParserParser.THEN, 0); }
		public TerminalNode ELSE() { return getToken(HilcoPencilParserParser.ELSE, 0); }
		public IfStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).enterIfStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).exitIfStatement(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilParserVisitor ) return ((HilcoPencilParserVisitor<? extends T>)visitor).visitIfStatement(this);
			else return visitor.visitChildren(this);
		}
	}

	public final IfStatementContext ifStatement() throws RecognitionException {
		IfStatementContext _localctx = new IfStatementContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_ifStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(68);
			match(IF);
			setState(69);
			statement();
			setState(70);
			match(THEN);
			setState(71);
			statement();
			setState(75);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				{
				{
				setState(72);
				match(ELSE);
				}
				setState(73);
				statement();
				}
				break;
			case 2:
				{
				{
				}
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

	public static class LogicalContext extends ParserRuleContext {
		public List<RelationalExpressionContext> relationalExpression() {
			return getRuleContexts(RelationalExpressionContext.class);
		}
		public RelationalExpressionContext relationalExpression(int i) {
			return getRuleContext(RelationalExpressionContext.class,i);
		}
		public List<TerminalNode> AND() { return getTokens(HilcoPencilParserParser.AND); }
		public TerminalNode AND(int i) {
			return getToken(HilcoPencilParserParser.AND, i);
		}
		public List<TerminalNode> OR() { return getTokens(HilcoPencilParserParser.OR); }
		public TerminalNode OR(int i) {
			return getToken(HilcoPencilParserParser.OR, i);
		}
		public TerminalNode NOT() { return getToken(HilcoPencilParserParser.NOT, 0); }
		public LogicalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_logical; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).enterLogical(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).exitLogical(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilParserVisitor ) return ((HilcoPencilParserVisitor<? extends T>)visitor).visitLogical(this);
			else return visitor.visitChildren(this);
		}
	}

	public final LogicalContext logical() throws RecognitionException {
		LogicalContext _localctx = new LogicalContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_logical);
		int _la;
		try {
			setState(87);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__1:
			case T__2:
			case T__3:
			case T__4:
			case T__5:
			case INT:
			case ID:
			case STRING:
			case FLOAT:
			case DATE:
			case ATFUNCTION:
			case CLASSNAME:
			case PERCENT:
			case LPAREN:
			case MINUS:
				enterOuterAlt(_localctx, 1);
				{
				setState(77);
				relationalExpression();
				setState(82);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==AND || _la==OR) {
					{
					{
					setState(78);
					_la = _input.LA(1);
					if ( !(_la==AND || _la==OR) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					setState(79);
					relationalExpression();
					}
					}
					setState(84);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case NOT:
				enterOuterAlt(_localctx, 2);
				{
				setState(85);
				match(NOT);
				setState(86);
				relationalExpression();
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

	public static class RelationalExpressionContext extends ParserRuleContext {
		public List<AddingExpressionContext> addingExpression() {
			return getRuleContexts(AddingExpressionContext.class);
		}
		public AddingExpressionContext addingExpression(int i) {
			return getRuleContext(AddingExpressionContext.class,i);
		}
		public List<TerminalNode> EQUALS() { return getTokens(HilcoPencilParserParser.EQUALS); }
		public TerminalNode EQUALS(int i) {
			return getToken(HilcoPencilParserParser.EQUALS, i);
		}
		public List<TerminalNode> NOT_EQUALS() { return getTokens(HilcoPencilParserParser.NOT_EQUALS); }
		public TerminalNode NOT_EQUALS(int i) {
			return getToken(HilcoPencilParserParser.NOT_EQUALS, i);
		}
		public List<TerminalNode> GT() { return getTokens(HilcoPencilParserParser.GT); }
		public TerminalNode GT(int i) {
			return getToken(HilcoPencilParserParser.GT, i);
		}
		public List<TerminalNode> GTE() { return getTokens(HilcoPencilParserParser.GTE); }
		public TerminalNode GTE(int i) {
			return getToken(HilcoPencilParserParser.GTE, i);
		}
		public List<TerminalNode> LT() { return getTokens(HilcoPencilParserParser.LT); }
		public TerminalNode LT(int i) {
			return getToken(HilcoPencilParserParser.LT, i);
		}
		public List<TerminalNode> LTE() { return getTokens(HilcoPencilParserParser.LTE); }
		public TerminalNode LTE(int i) {
			return getToken(HilcoPencilParserParser.LTE, i);
		}
		public RelationalExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationalExpression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).enterRelationalExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).exitRelationalExpression(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilParserVisitor ) return ((HilcoPencilParserVisitor<? extends T>)visitor).visitRelationalExpression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final RelationalExpressionContext relationalExpression() throws RecognitionException {
		RelationalExpressionContext _localctx = new RelationalExpressionContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_relationalExpression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(89);
			addingExpression();
			setState(94);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << EQUALS) | (1L << NOT_EQUALS) | (1L << LT) | (1L << LTE) | (1L << GT) | (1L << GTE))) != 0)) {
				{
				{
				setState(90);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << EQUALS) | (1L << NOT_EQUALS) | (1L << LT) | (1L << LTE) | (1L << GT) | (1L << GTE))) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(91);
				addingExpression();
				}
				}
				setState(96);
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

	public static class AddingExpressionContext extends ParserRuleContext {
		public List<MultiplyingExpressionContext> multiplyingExpression() {
			return getRuleContexts(MultiplyingExpressionContext.class);
		}
		public MultiplyingExpressionContext multiplyingExpression(int i) {
			return getRuleContext(MultiplyingExpressionContext.class,i);
		}
		public List<TerminalNode> PLUS() { return getTokens(HilcoPencilParserParser.PLUS); }
		public TerminalNode PLUS(int i) {
			return getToken(HilcoPencilParserParser.PLUS, i);
		}
		public List<TerminalNode> MINUS() { return getTokens(HilcoPencilParserParser.MINUS); }
		public TerminalNode MINUS(int i) {
			return getToken(HilcoPencilParserParser.MINUS, i);
		}
		public AddingExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_addingExpression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).enterAddingExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).exitAddingExpression(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilParserVisitor ) return ((HilcoPencilParserVisitor<? extends T>)visitor).visitAddingExpression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AddingExpressionContext addingExpression() throws RecognitionException {
		AddingExpressionContext _localctx = new AddingExpressionContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_addingExpression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(97);
			multiplyingExpression();
			setState(102);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==PLUS || _la==MINUS) {
				{
				{
				setState(98);
				_la = _input.LA(1);
				if ( !(_la==PLUS || _la==MINUS) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(99);
				multiplyingExpression();
				}
				}
				setState(104);
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

	public static class MultiplyingExpressionContext extends ParserRuleContext {
		public List<FactorContext> factor() {
			return getRuleContexts(FactorContext.class);
		}
		public FactorContext factor(int i) {
			return getRuleContext(FactorContext.class,i);
		}
		public List<TerminalNode> TIMES() { return getTokens(HilcoPencilParserParser.TIMES); }
		public TerminalNode TIMES(int i) {
			return getToken(HilcoPencilParserParser.TIMES, i);
		}
		public List<TerminalNode> DIV() { return getTokens(HilcoPencilParserParser.DIV); }
		public TerminalNode DIV(int i) {
			return getToken(HilcoPencilParserParser.DIV, i);
		}
		public MultiplyingExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_multiplyingExpression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).enterMultiplyingExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).exitMultiplyingExpression(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilParserVisitor ) return ((HilcoPencilParserVisitor<? extends T>)visitor).visitMultiplyingExpression(this);
			else return visitor.visitChildren(this);
		}
	}

	public final MultiplyingExpressionContext multiplyingExpression() throws RecognitionException {
		MultiplyingExpressionContext _localctx = new MultiplyingExpressionContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_multiplyingExpression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(105);
			factor();
			setState(110);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==TIMES || _la==DIV) {
				{
				{
				setState(106);
				_la = _input.LA(1);
				if ( !(_la==TIMES || _la==DIV) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(107);
				factor();
				}
				}
				setState(112);
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

	public static class FactorContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(HilcoPencilParserParser.MINUS, 0); }
		public FactorContext factor() {
			return getRuleContext(FactorContext.class,0);
		}
		public BaseContext base() {
			return getRuleContext(BaseContext.class,0);
		}
		public FactorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_factor; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).enterFactor(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).exitFactor(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilParserVisitor ) return ((HilcoPencilParserVisitor<? extends T>)visitor).visitFactor(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FactorContext factor() throws RecognitionException {
		FactorContext _localctx = new FactorContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_factor);
		try {
			setState(116);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MINUS:
				enterOuterAlt(_localctx, 1);
				{
				setState(113);
				match(MINUS);
				setState(114);
				factor();
				}
				break;
			case T__1:
			case T__2:
			case T__3:
			case T__4:
			case T__5:
			case INT:
			case ID:
			case STRING:
			case FLOAT:
			case DATE:
			case ATFUNCTION:
			case CLASSNAME:
			case PERCENT:
			case LPAREN:
				enterOuterAlt(_localctx, 2);
				{
				setState(115);
				base();
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

	public static class BaseContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(HilcoPencilParserParser.LPAREN, 0); }
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(HilcoPencilParserParser.RPAREN, 0); }
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public TerminalNode PERCENT() { return getToken(HilcoPencilParserParser.PERCENT, 0); }
		public TerminalNode INT() { return getToken(HilcoPencilParserParser.INT, 0); }
		public TerminalNode FLOAT() { return getToken(HilcoPencilParserParser.FLOAT, 0); }
		public TerminalNode STRING() { return getToken(HilcoPencilParserParser.STRING, 0); }
		public TerminalNode DATE() { return getToken(HilcoPencilParserParser.DATE, 0); }
		public AtFunctionContext atFunction() {
			return getRuleContext(AtFunctionContext.class,0);
		}
		public DataAccessorContext dataAccessor() {
			return getRuleContext(DataAccessorContext.class,0);
		}
		public SpecialKeywordContext specialKeyword() {
			return getRuleContext(SpecialKeywordContext.class,0);
		}
		public BaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_base; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).enterBase(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).exitBase(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilParserVisitor ) return ((HilcoPencilParserVisitor<? extends T>)visitor).visitBase(this);
			else return visitor.visitChildren(this);
		}
	}

	public final BaseContext base() throws RecognitionException {
		BaseContext _localctx = new BaseContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_base);
		try {
			setState(131);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(118);
				match(LPAREN);
				setState(119);
				statement();
				setState(120);
				match(RPAREN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(122);
				name();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(123);
				match(PERCENT);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(124);
				match(INT);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(125);
				match(FLOAT);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(126);
				match(STRING);
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(127);
				match(DATE);
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(128);
				atFunction();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(129);
				dataAccessor();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(130);
				specialKeyword();
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

	public static class NameContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(HilcoPencilParserParser.ID, 0); }
		public TerminalNode CLASSNAME() { return getToken(HilcoPencilParserParser.CLASSNAME, 0); }
		public TerminalNode COLON() { return getToken(HilcoPencilParserParser.COLON, 0); }
		public NameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_name; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).enterName(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).exitName(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilParserVisitor ) return ((HilcoPencilParserVisitor<? extends T>)visitor).visitName(this);
			else return visitor.visitChildren(this);
		}
	}

	public final NameContext name() throws RecognitionException {
		NameContext _localctx = new NameContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_name);
		try {
			setState(137);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ID:
				enterOuterAlt(_localctx, 1);
				{
				setState(133);
				match(ID);
				}
				break;
			case CLASSNAME:
				enterOuterAlt(_localctx, 2);
				{
				setState(134);
				match(CLASSNAME);
				setState(135);
				match(COLON);
				setState(136);
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
		public TerminalNode ATFUNCTION() { return getToken(HilcoPencilParserParser.ATFUNCTION, 0); }
		public TerminalNode LPAREN() { return getToken(HilcoPencilParserParser.LPAREN, 0); }
		public ArgListContext argList() {
			return getRuleContext(ArgListContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(HilcoPencilParserParser.RPAREN, 0); }
		public AtFunctionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_atFunction; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).enterAtFunction(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).exitAtFunction(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilParserVisitor ) return ((HilcoPencilParserVisitor<? extends T>)visitor).visitAtFunction(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AtFunctionContext atFunction() throws RecognitionException {
		AtFunctionContext _localctx = new AtFunctionContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_atFunction);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(139);
			match(ATFUNCTION);
			setState(140);
			match(LPAREN);
			setState(141);
			argList();
			setState(142);
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

	public static class DataAccessorContext extends ParserRuleContext {
		public TerminalNode CLASSNAME() { return getToken(HilcoPencilParserParser.CLASSNAME, 0); }
		public List<FieldNameContext> fieldName() {
			return getRuleContexts(FieldNameContext.class);
		}
		public FieldNameContext fieldName(int i) {
			return getRuleContext(FieldNameContext.class,i);
		}
		public DataAccessorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dataAccessor; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).enterDataAccessor(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).exitDataAccessor(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilParserVisitor ) return ((HilcoPencilParserVisitor<? extends T>)visitor).visitDataAccessor(this);
			else return visitor.visitChildren(this);
		}
	}

	public final DataAccessorContext dataAccessor() throws RecognitionException {
		DataAccessorContext _localctx = new DataAccessorContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_dataAccessor);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(144);
			match(CLASSNAME);
			setState(146); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(145);
				fieldName();
				}
				}
				setState(148); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==DOT );
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

	public static class FieldNameContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(HilcoPencilParserParser.DOT, 0); }
		public TerminalNode ID() { return getToken(HilcoPencilParserParser.ID, 0); }
		public TerminalNode CLASSNAME() { return getToken(HilcoPencilParserParser.CLASSNAME, 0); }
		public TerminalNode LPAREN() { return getToken(HilcoPencilParserParser.LPAREN, 0); }
		public ArgListContext argList() {
			return getRuleContext(ArgListContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(HilcoPencilParserParser.RPAREN, 0); }
		public FieldNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fieldName; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).enterFieldName(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).exitFieldName(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilParserVisitor ) return ((HilcoPencilParserVisitor<? extends T>)visitor).visitFieldName(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FieldNameContext fieldName() throws RecognitionException {
		FieldNameContext _localctx = new FieldNameContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_fieldName);
		int _la;
		try {
			setState(158);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(150);
				match(DOT);
				setState(151);
				_la = _input.LA(1);
				if ( !(_la==ID || _la==CLASSNAME) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(152);
				match(DOT);
				setState(153);
				_la = _input.LA(1);
				if ( !(_la==ID || _la==CLASSNAME) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(154);
				match(LPAREN);
				setState(155);
				argList();
				setState(156);
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

	public static class ArgListContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(HilcoPencilParserParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(HilcoPencilParserParser.COMMA, i);
		}
		public ArgListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).enterArgList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).exitArgList(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilParserVisitor ) return ((HilcoPencilParserVisitor<? extends T>)visitor).visitArgList(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ArgListContext argList() throws RecognitionException {
		ArgListContext _localctx = new ArgListContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_argList);
		int _la;
		try {
			setState(172);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__1:
			case T__2:
			case T__3:
			case T__4:
			case T__5:
			case IF:
			case CASE:
			case SWITCH:
			case NOT:
			case INT:
			case ID:
			case STRING:
			case FLOAT:
			case DATE:
			case ATFUNCTION:
			case CLASSNAME:
			case PERCENT:
			case LPAREN:
			case MINUS:
				enterOuterAlt(_localctx, 1);
				{
				setState(160);
				statement();
				setState(168);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(161);
					match(COMMA);
					setState(164);
					_errHandler.sync(this);
					switch (_input.LA(1)) {
					case T__1:
					case T__2:
					case T__3:
					case T__4:
					case T__5:
					case IF:
					case CASE:
					case SWITCH:
					case NOT:
					case INT:
					case ID:
					case STRING:
					case FLOAT:
					case DATE:
					case ATFUNCTION:
					case CLASSNAME:
					case PERCENT:
					case LPAREN:
					case MINUS:
						{
						setState(162);
						statement();
						}
						break;
					case COMMA:
					case RPAREN:
						{
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					}
					}
					setState(170);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case RPAREN:
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

	public static class SpecialKeywordContext extends ParserRuleContext {
		public SpecialKeywordContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_specialKeyword; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).enterSpecialKeyword(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HilcoPencilParserListener ) ((HilcoPencilParserListener)listener).exitSpecialKeyword(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HilcoPencilParserVisitor ) return ((HilcoPencilParserVisitor<? extends T>)visitor).visitSpecialKeyword(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SpecialKeywordContext specialKeyword() throws RecognitionException {
		SpecialKeywordContext _localctx = new SpecialKeywordContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_specialKeyword);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(174);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__1) | (1L << T__2) | (1L << T__3) | (1L << T__4) | (1L << T__5))) != 0)) ) {
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

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\63\u00b3\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\3\2\3\2\3\3\3\3\3\3\3\3\5\3/\n\3\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\5\3\5\7\59\n\5\f\5\16\5<\13\5\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3"+
		"\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\5\bN\n\b\3\t\3\t\3\t\7\tS\n\t\f\t\16"+
		"\tV\13\t\3\t\3\t\5\tZ\n\t\3\n\3\n\3\n\7\n_\n\n\f\n\16\nb\13\n\3\13\3\13"+
		"\3\13\7\13g\n\13\f\13\16\13j\13\13\3\f\3\f\3\f\7\fo\n\f\f\f\16\fr\13\f"+
		"\3\r\3\r\3\r\5\rw\n\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\5\16\u0086\n\16\3\17\3\17\3\17\3\17\5\17\u008c\n\17\3"+
		"\20\3\20\3\20\3\20\3\20\3\21\3\21\6\21\u0095\n\21\r\21\16\21\u0096\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u00a1\n\22\3\23\3\23\3\23\3\23"+
		"\5\23\u00a7\n\23\7\23\u00a9\n\23\f\23\16\23\u00ac\13\23\3\23\5\23\u00af"+
		"\n\23\3\24\3\24\3\24\2\2\25\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \""+
		"$&\2\b\3\2\21\22\4\2\"\"*.\3\2/\60\3\2\61\62\4\2\25\25\32\32\3\2\4\b\2"+
		"\u00b9\2(\3\2\2\2\4.\3\2\2\2\6\60\3\2\2\2\b\66\3\2\2\2\n=\3\2\2\2\fB\3"+
		"\2\2\2\16F\3\2\2\2\20Y\3\2\2\2\22[\3\2\2\2\24c\3\2\2\2\26k\3\2\2\2\30"+
		"v\3\2\2\2\32\u0085\3\2\2\2\34\u008b\3\2\2\2\36\u008d\3\2\2\2 \u0092\3"+
		"\2\2\2\"\u00a0\3\2\2\2$\u00ae\3\2\2\2&\u00b0\3\2\2\2()\5\4\3\2)\3\3\2"+
		"\2\2*/\5\20\t\2+/\5\16\b\2,/\5\6\4\2-/\5\f\7\2.*\3\2\2\2.+\3\2\2\2.,\3"+
		"\2\2\2.-\3\2\2\2/\5\3\2\2\2\60\61\7\f\2\2\61\62\5\4\3\2\62\63\7\r\2\2"+
		"\63\64\5\b\5\2\64\65\7\16\2\2\65\7\3\2\2\2\66:\5\n\6\2\679\5\n\6\28\67"+
		"\3\2\2\29<\3\2\2\2:8\3\2\2\2:;\3\2\2\2;\t\3\2\2\2<:\3\2\2\2=>\5\32\16"+
		"\2>?\7\3\2\2?@\5\4\3\2@A\7 \2\2A\13\3\2\2\2BC\7\17\2\2CD\5\b\5\2DE\7\20"+
		"\2\2E\r\3\2\2\2FG\7\t\2\2GH\5\4\3\2HI\7\n\2\2IM\5\4\3\2JK\7\13\2\2KN\5"+
		"\4\3\2LN\3\2\2\2MJ\3\2\2\2ML\3\2\2\2N\17\3\2\2\2OT\5\22\n\2PQ\t\2\2\2"+
		"QS\5\22\n\2RP\3\2\2\2SV\3\2\2\2TR\3\2\2\2TU\3\2\2\2UZ\3\2\2\2VT\3\2\2"+
		"\2WX\7\23\2\2XZ\5\22\n\2YO\3\2\2\2YW\3\2\2\2Z\21\3\2\2\2[`\5\24\13\2\\"+
		"]\t\3\2\2]_\5\24\13\2^\\\3\2\2\2_b\3\2\2\2`^\3\2\2\2`a\3\2\2\2a\23\3\2"+
		"\2\2b`\3\2\2\2ch\5\26\f\2de\t\4\2\2eg\5\26\f\2fd\3\2\2\2gj\3\2\2\2hf\3"+
		"\2\2\2hi\3\2\2\2i\25\3\2\2\2jh\3\2\2\2kp\5\30\r\2lm\t\5\2\2mo\5\30\r\2"+
		"nl\3\2\2\2or\3\2\2\2pn\3\2\2\2pq\3\2\2\2q\27\3\2\2\2rp\3\2\2\2st\7\60"+
		"\2\2tw\5\30\r\2uw\5\32\16\2vs\3\2\2\2vu\3\2\2\2w\31\3\2\2\2xy\7(\2\2y"+
		"z\5\4\3\2z{\7)\2\2{\u0086\3\2\2\2|\u0086\5\34\17\2}\u0086\7\33\2\2~\u0086"+
		"\7\24\2\2\177\u0086\7\27\2\2\u0080\u0086\7\26\2\2\u0081\u0086\7\30\2\2"+
		"\u0082\u0086\5\36\20\2\u0083\u0086\5 \21\2\u0084\u0086\5&\24\2\u0085x"+
		"\3\2\2\2\u0085|\3\2\2\2\u0085}\3\2\2\2\u0085~\3\2\2\2\u0085\177\3\2\2"+
		"\2\u0085\u0080\3\2\2\2\u0085\u0081\3\2\2\2\u0085\u0082\3\2\2\2\u0085\u0083"+
		"\3\2\2\2\u0085\u0084\3\2\2\2\u0086\33\3\2\2\2\u0087\u008c\7\25\2\2\u0088"+
		"\u0089\7\32\2\2\u0089\u008a\7\37\2\2\u008a\u008c\7\25\2\2\u008b\u0087"+
		"\3\2\2\2\u008b\u0088\3\2\2\2\u008c\35\3\2\2\2\u008d\u008e\7\31\2\2\u008e"+
		"\u008f\7(\2\2\u008f\u0090\5$\23\2\u0090\u0091\7)\2\2\u0091\37\3\2\2\2"+
		"\u0092\u0094\7\32\2\2\u0093\u0095\5\"\22\2\u0094\u0093\3\2\2\2\u0095\u0096"+
		"\3\2\2\2\u0096\u0094\3\2\2\2\u0096\u0097\3\2\2\2\u0097!\3\2\2\2\u0098"+
		"\u0099\7\36\2\2\u0099\u00a1\t\6\2\2\u009a\u009b\7\36\2\2\u009b\u009c\t"+
		"\6\2\2\u009c\u009d\7(\2\2\u009d\u009e\5$\23\2\u009e\u009f\7)\2\2\u009f"+
		"\u00a1\3\2\2\2\u00a0\u0098\3\2\2\2\u00a0\u009a\3\2\2\2\u00a1#\3\2\2\2"+
		"\u00a2\u00aa\5\4\3\2\u00a3\u00a6\7!\2\2\u00a4\u00a7\5\4\3\2\u00a5\u00a7"+
		"\3\2\2\2\u00a6\u00a4\3\2\2\2\u00a6\u00a5\3\2\2\2\u00a7\u00a9\3\2\2\2\u00a8"+
		"\u00a3\3\2\2\2\u00a9\u00ac\3\2\2\2\u00aa\u00a8\3\2\2\2\u00aa\u00ab\3\2"+
		"\2\2\u00ab\u00af\3\2\2\2\u00ac\u00aa\3\2\2\2\u00ad\u00af\3\2\2\2\u00ae"+
		"\u00a2\3\2\2\2\u00ae\u00ad\3\2\2\2\u00af%\3\2\2\2\u00b0\u00b1\t\7\2\2"+
		"\u00b1\'\3\2\2\2\22.:MTY`hpv\u0085\u008b\u0096\u00a0\u00a6\u00aa\u00ae";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
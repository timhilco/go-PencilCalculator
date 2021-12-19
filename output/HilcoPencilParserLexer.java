// Generated from /Users/timothyhilgenberg/go/src/github.com/timhilco/go-PencilCalculator/HilcoPencilParser.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class HilcoPencilParserLexer extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "IF", "THEN", "ELSE", 
			"CASE", "CASE_IS", "CASE_END", "SWITCH", "SWITCH_END", "AND", "OR", "NOT", 
			"DIGIT", "LOWERCASELETTERS", "UPPERCASELETTERS", "INT", "ID", "STRING", 
			"EXPONENT", "HEX_DIGIT", "ESC_SEQ", "OCTAL_ESC", "UNICODE_ESC", "FLOAT", 
			"DATE", "ATFUNCTION", "CLASSNAME", "PERCENT", "NEWLINE", "WS", "DOT", 
			"COLON", "SEMI", "COMMA", "EQUALS", "LBRACKET", "RBRACKET", "CURLYLBRACKET", 
			"CURLYRBRACKET", "DOTDOT", "LPAREN", "RPAREN", "NOT_EQUALS", "LT", "LTE", 
			"GT", "GTE", "PLUS", "MINUS", "TIMES", "DIV", "EXPONENTIATE"
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


	public HilcoPencilParserLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "HilcoPencilParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\63\u0199\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64"+
		"\t\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\3\2\3\2\3\2\3\3"+
		"\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3"+
		"\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\t\3\t\3\t\3\t"+
		"\3\t\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\r\3\r"+
		"\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\21\3\21"+
		"\3\21\3\22\3\22\3\22\3\22\3\23\3\23\3\24\3\24\3\25\3\25\3\26\6\26\u00d6"+
		"\n\26\r\26\16\26\u00d7\3\27\3\27\3\27\3\27\3\27\3\27\3\27\7\27\u00e1\n"+
		"\27\f\27\16\27\u00e4\13\27\3\30\3\30\3\30\7\30\u00e9\n\30\f\30\16\30\u00ec"+
		"\13\30\3\30\3\30\3\31\3\31\5\31\u00f2\n\31\3\31\6\31\u00f5\n\31\r\31\16"+
		"\31\u00f6\3\32\3\32\3\33\3\33\3\33\3\33\5\33\u00ff\n\33\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\5\34\u010a\n\34\3\35\3\35\3\35\3\35\3\35"+
		"\3\35\3\35\3\36\6\36\u0114\n\36\r\36\16\36\u0115\3\36\3\36\7\36\u011a"+
		"\n\36\f\36\16\36\u011d\13\36\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3"+
		"\37\5\37\u0128\n\37\3\37\3\37\5\37\u012c\n\37\3\37\3\37\5\37\u0130\n\37"+
		"\5\37\u0132\n\37\3\37\3\37\3\37\3\37\3\37\3\37\5\37\u013a\n\37\3\37\3"+
		"\37\3 \3 \3 \5 \u0141\n \3 \3 \3 \7 \u0146\n \f \16 \u0149\13 \3!\3!\3"+
		"!\3!\3!\7!\u0150\n!\f!\16!\u0153\13!\3\"\3\"\5\"\u0157\n\"\3\"\3\"\3#"+
		"\5#\u015c\n#\3#\3#\3#\3#\3$\3$\3$\3$\5$\u0166\n$\3$\3$\3%\3%\3&\3&\3\'"+
		"\3\'\3(\3(\3)\3)\3*\3*\3+\3+\3,\3,\3-\3-\3.\3.\3.\3/\3/\3\60\3\60\3\61"+
		"\3\61\3\61\3\62\3\62\3\63\3\63\3\63\3\64\3\64\3\65\3\65\3\65\3\66\3\66"+
		"\3\67\3\67\38\38\39\39\3:\3:\2\2;\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23"+
		"\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\2\'\2)\2+\24-\25/\26\61"+
		"\2\63\2\65\2\67\29\2;\27=\30?\31A\32C\33E\34G\35I\36K\37M O!Q\"S#U$W%"+
		"Y&[\'](_)a*c+e,g-i.k/m\60o\61q\62s\63\3\2\t\4\2))^^\4\2GGgg\4\2--//\5"+
		"\2\62;CHch\n\2$$))^^ddhhppttvv\4\2//\61\61\5\2\13\13\16\16\"\"\2\u01b2"+
		"\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2"+
		"\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2"+
		"\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2"+
		"\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2"+
		"\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M"+
		"\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2"+
		"\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2"+
		"\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s"+
		"\3\2\2\2\3u\3\2\2\2\5x\3\2\2\2\7~\3\2\2\2\t\u0083\3\2\2\2\13\u0089\3\2"+
		"\2\2\r\u008d\3\2\2\2\17\u0095\3\2\2\2\21\u0098\3\2\2\2\23\u009d\3\2\2"+
		"\2\25\u00a2\3\2\2\2\27\u00a7\3\2\2\2\31\u00aa\3\2\2\2\33\u00b2\3\2\2\2"+
		"\35\u00b9\3\2\2\2\37\u00c3\3\2\2\2!\u00c7\3\2\2\2#\u00ca\3\2\2\2%\u00ce"+
		"\3\2\2\2\'\u00d0\3\2\2\2)\u00d2\3\2\2\2+\u00d5\3\2\2\2-\u00d9\3\2\2\2"+
		"/\u00e5\3\2\2\2\61\u00ef\3\2\2\2\63\u00f8\3\2\2\2\65\u00fe\3\2\2\2\67"+
		"\u0109\3\2\2\29\u010b\3\2\2\2;\u0113\3\2\2\2=\u011e\3\2\2\2?\u013d\3\2"+
		"\2\2A\u014a\3\2\2\2C\u0156\3\2\2\2E\u015b\3\2\2\2G\u0165\3\2\2\2I\u0169"+
		"\3\2\2\2K\u016b\3\2\2\2M\u016d\3\2\2\2O\u016f\3\2\2\2Q\u0171\3\2\2\2S"+
		"\u0173\3\2\2\2U\u0175\3\2\2\2W\u0177\3\2\2\2Y\u0179\3\2\2\2[\u017b\3\2"+
		"\2\2]\u017e\3\2\2\2_\u0180\3\2\2\2a\u0182\3\2\2\2c\u0185\3\2\2\2e\u0187"+
		"\3\2\2\2g\u018a\3\2\2\2i\u018c\3\2\2\2k\u018f\3\2\2\2m\u0191\3\2\2\2o"+
		"\u0193\3\2\2\2q\u0195\3\2\2\2s\u0197\3\2\2\2uv\7<\2\2vw\7<\2\2w\4\3\2"+
		"\2\2xy\7v\2\2yz\7q\2\2z{\7f\2\2{|\7c\2\2|}\7{\2\2}\6\3\2\2\2~\177\7v\2"+
		"\2\177\u0080\7t\2\2\u0080\u0081\7w\2\2\u0081\u0082\7g\2\2\u0082\b\3\2"+
		"\2\2\u0083\u0084\7h\2\2\u0084\u0085\7c\2\2\u0085\u0086\7n\2\2\u0086\u0087"+
		"\7u\2\2\u0087\u0088\7g\2\2\u0088\n\3\2\2\2\u0089\u008a\7p\2\2\u008a\u008b"+
		"\7k\2\2\u008b\u008c\7n\2\2\u008c\f\3\2\2\2\u008d\u008e\7f\2\2\u008e\u008f"+
		"\7g\2\2\u008f\u0090\7h\2\2\u0090\u0091\7c\2\2\u0091\u0092\7w\2\2\u0092"+
		"\u0093\7n\2\2\u0093\u0094\7v\2\2\u0094\16\3\2\2\2\u0095\u0096\7k\2\2\u0096"+
		"\u0097\7h\2\2\u0097\20\3\2\2\2\u0098\u0099\7v\2\2\u0099\u009a\7j\2\2\u009a"+
		"\u009b\7g\2\2\u009b\u009c\7p\2\2\u009c\22\3\2\2\2\u009d\u009e\7g\2\2\u009e"+
		"\u009f\7n\2\2\u009f\u00a0\7u\2\2\u00a0\u00a1\7g\2\2\u00a1\24\3\2\2\2\u00a2"+
		"\u00a3\7e\2\2\u00a3\u00a4\7c\2\2\u00a4\u00a5\7u\2\2\u00a5\u00a6\7g\2\2"+
		"\u00a6\26\3\2\2\2\u00a7\u00a8\7k\2\2\u00a8\u00a9\7u\2\2\u00a9\30\3\2\2"+
		"\2\u00aa\u00ab\7g\2\2\u00ab\u00ac\7p\2\2\u00ac\u00ad\7f\2\2\u00ad\u00ae"+
		"\7e\2\2\u00ae\u00af\7c\2\2\u00af\u00b0\7u\2\2\u00b0\u00b1\7g\2\2\u00b1"+
		"\32\3\2\2\2\u00b2\u00b3\7u\2\2\u00b3\u00b4\7y\2\2\u00b4\u00b5\7k\2\2\u00b5"+
		"\u00b6\7v\2\2\u00b6\u00b7\7e\2\2\u00b7\u00b8\7j\2\2\u00b8\34\3\2\2\2\u00b9"+
		"\u00ba\7g\2\2\u00ba\u00bb\7p\2\2\u00bb\u00bc\7f\2\2\u00bc\u00bd\7u\2\2"+
		"\u00bd\u00be\7y\2\2\u00be\u00bf\7k\2\2\u00bf\u00c0\7v\2\2\u00c0\u00c1"+
		"\7e\2\2\u00c1\u00c2\7j\2\2\u00c2\36\3\2\2\2\u00c3\u00c4\7C\2\2\u00c4\u00c5"+
		"\7P\2\2\u00c5\u00c6\7F\2\2\u00c6 \3\2\2\2\u00c7\u00c8\7Q\2\2\u00c8\u00c9"+
		"\7T\2\2\u00c9\"\3\2\2\2\u00ca\u00cb\7P\2\2\u00cb\u00cc\7Q\2\2\u00cc\u00cd"+
		"\7V\2\2\u00cd$\3\2\2\2\u00ce\u00cf\4\62;\2\u00cf&\3\2\2\2\u00d0\u00d1"+
		"\4c|\2\u00d1(\3\2\2\2\u00d2\u00d3\4C\\\2\u00d3*\3\2\2\2\u00d4\u00d6\5"+
		"%\23\2\u00d5\u00d4\3\2\2\2\u00d6\u00d7\3\2\2\2\u00d7\u00d5\3\2\2\2\u00d7"+
		"\u00d8\3\2\2\2\u00d8,\3\2\2\2\u00d9\u00e2\5\'\24\2\u00da\u00e1\5\'\24"+
		"\2\u00db\u00e1\5)\25\2\u00dc\u00e1\5%\23\2\u00dd\u00e1\7a\2\2\u00de\u00e1"+
		"\5K&\2\u00df\u00e1\7^\2\2\u00e0\u00da\3\2\2\2\u00e0\u00db\3\2\2\2\u00e0"+
		"\u00dc\3\2\2\2\u00e0\u00dd\3\2\2\2\u00e0\u00de\3\2\2\2\u00e0\u00df\3\2"+
		"\2\2\u00e1\u00e4\3\2\2\2\u00e2\u00e0\3\2\2\2\u00e2\u00e3\3\2\2\2\u00e3"+
		".\3\2\2\2\u00e4\u00e2\3\2\2\2\u00e5\u00ea\7)\2\2\u00e6\u00e9\5\65\33\2"+
		"\u00e7\u00e9\n\2\2\2\u00e8\u00e6\3\2\2\2\u00e8\u00e7\3\2\2\2\u00e9\u00ec"+
		"\3\2\2\2\u00ea\u00e8\3\2\2\2\u00ea\u00eb\3\2\2\2\u00eb\u00ed\3\2\2\2\u00ec"+
		"\u00ea\3\2\2\2\u00ed\u00ee\7)\2\2\u00ee\60\3\2\2\2\u00ef\u00f1\t\3\2\2"+
		"\u00f0\u00f2\t\4\2\2\u00f1\u00f0\3\2\2\2\u00f1\u00f2\3\2\2\2\u00f2\u00f4"+
		"\3\2\2\2\u00f3\u00f5\4\62;\2\u00f4\u00f3\3\2\2\2\u00f5\u00f6\3\2\2\2\u00f6"+
		"\u00f4\3\2\2\2\u00f6\u00f7\3\2\2\2\u00f7\62\3\2\2\2\u00f8\u00f9\t\5\2"+
		"\2\u00f9\64\3\2\2\2\u00fa\u00fb\7^\2\2\u00fb\u00ff\t\6\2\2\u00fc\u00ff"+
		"\59\35\2\u00fd\u00ff\5\67\34\2\u00fe\u00fa\3\2\2\2\u00fe\u00fc\3\2\2\2"+
		"\u00fe\u00fd\3\2\2\2\u00ff\66\3\2\2\2\u0100\u0101\7^\2\2\u0101\u0102\4"+
		"\62\65\2\u0102\u0103\4\629\2\u0103\u010a\4\629\2\u0104\u0105\7^\2\2\u0105"+
		"\u0106\4\629\2\u0106\u010a\4\629\2\u0107\u0108\7^\2\2\u0108\u010a\4\62"+
		"9\2\u0109\u0100\3\2\2\2\u0109\u0104\3\2\2\2\u0109\u0107\3\2\2\2\u010a"+
		"8\3\2\2\2\u010b\u010c\7^\2\2\u010c\u010d\7w\2\2\u010d\u010e\5\63\32\2"+
		"\u010e\u010f\5\63\32\2\u010f\u0110\5\63\32\2\u0110\u0111\5\63\32\2\u0111"+
		":\3\2\2\2\u0112\u0114\5%\23\2\u0113\u0112\3\2\2\2\u0114\u0115\3\2\2\2"+
		"\u0115\u0113\3\2\2\2\u0115\u0116\3\2\2\2\u0116\u0117\3\2\2\2\u0117\u011b"+
		"\7\60\2\2\u0118\u011a\5%\23\2\u0119\u0118\3\2\2\2\u011a\u011d\3\2\2\2"+
		"\u011b\u0119\3\2\2\2\u011b\u011c\3\2\2\2\u011c<\3\2\2\2\u011d\u011b\3"+
		"\2\2\2\u011e\u011f\5W,\2\u011f\u0120\5%\23\2\u0120\u0121\5%\23\2\u0121"+
		"\u0131\t\7\2\2\u0122\u0123\5%\23\2\u0123\u0124\5%\23\2\u0124\u0132\3\2"+
		"\2\2\u0125\u0128\5\'\24\2\u0126\u0128\5)\25\2\u0127\u0125\3\2\2\2\u0127"+
		"\u0126\3\2\2\2\u0128\u012b\3\2\2\2\u0129\u012c\5\'\24\2\u012a\u012c\5"+
		")\25\2\u012b\u0129\3\2\2\2\u012b\u012a\3\2\2\2\u012c\u012f\3\2\2\2\u012d"+
		"\u0130\5\'\24\2\u012e\u0130\5)\25\2\u012f\u012d\3\2\2\2\u012f\u012e\3"+
		"\2\2\2\u0130\u0132\3\2\2\2\u0131\u0122\3\2\2\2\u0131\u0127\3\2\2\2\u0132"+
		"\u0133\3\2\2\2\u0133\u0134\t\7\2\2\u0134\u0135\5%\23\2\u0135\u0139\5%"+
		"\23\2\u0136\u0137\5%\23\2\u0137\u0138\5%\23\2\u0138\u013a\3\2\2\2\u0139"+
		"\u0136\3\2\2\2\u0139\u013a\3\2\2\2\u013a\u013b\3\2\2\2\u013b\u013c\5Y"+
		"-\2\u013c>\3\2\2\2\u013d\u0140\7B\2\2\u013e\u0141\5\'\24\2\u013f\u0141"+
		"\5)\25\2\u0140\u013e\3\2\2\2\u0140\u013f\3\2\2\2\u0141\u0147\3\2\2\2\u0142"+
		"\u0146\5\'\24\2\u0143\u0146\5)\25\2\u0144\u0146\5%\23\2\u0145\u0142\3"+
		"\2\2\2\u0145\u0143\3\2\2\2\u0145\u0144\3\2\2\2\u0146\u0149\3\2\2\2\u0147"+
		"\u0145\3\2\2\2\u0147\u0148\3\2\2\2\u0148@\3\2\2\2\u0149\u0147\3\2\2\2"+
		"\u014a\u0151\5)\25\2\u014b\u0150\5\'\24\2\u014c\u0150\5)\25\2\u014d\u0150"+
		"\5%\23\2\u014e\u0150\7a\2\2\u014f\u014b\3\2\2\2\u014f\u014c\3\2\2\2\u014f"+
		"\u014d\3\2\2\2\u014f\u014e\3\2\2\2\u0150\u0153\3\2\2\2\u0151\u014f\3\2"+
		"\2\2\u0151\u0152\3\2\2\2\u0152B\3\2\2\2\u0153\u0151\3\2\2\2\u0154\u0157"+
		"\5+\26\2\u0155\u0157\5;\36\2\u0156\u0154\3\2\2\2\u0156\u0155\3\2\2\2\u0157"+
		"\u0158\3\2\2\2\u0158\u0159\7\'\2\2\u0159D\3\2\2\2\u015a\u015c\7\17\2\2"+
		"\u015b\u015a\3\2\2\2\u015b\u015c\3\2\2\2\u015c\u015d\3\2\2\2\u015d\u015e"+
		"\7\f\2\2\u015e\u015f\3\2\2\2\u015f\u0160\b#\2\2\u0160F\3\2\2\2\u0161\u0166"+
		"\t\b\2\2\u0162\u0163\7\17\2\2\u0163\u0166\7\f\2\2\u0164\u0166\7\5\2\2"+
		"\u0165\u0161\3\2\2\2\u0165\u0162\3\2\2\2\u0165\u0164\3\2\2\2\u0166\u0167"+
		"\3\2\2\2\u0167\u0168\b$\2\2\u0168H\3\2\2\2\u0169\u016a\7\60\2\2\u016a"+
		"J\3\2\2\2\u016b\u016c\7<\2\2\u016cL\3\2\2\2\u016d\u016e\7=\2\2\u016eN"+
		"\3\2\2\2\u016f\u0170\7.\2\2\u0170P\3\2\2\2\u0171\u0172\7?\2\2\u0172R\3"+
		"\2\2\2\u0173\u0174\7]\2\2\u0174T\3\2\2\2\u0175\u0176\7_\2\2\u0176V\3\2"+
		"\2\2\u0177\u0178\7}\2\2\u0178X\3\2\2\2\u0179\u017a\7\177\2\2\u017aZ\3"+
		"\2\2\2\u017b\u017c\7\60\2\2\u017c\u017d\7\60\2\2\u017d\\\3\2\2\2\u017e"+
		"\u017f\7*\2\2\u017f^\3\2\2\2\u0180\u0181\7+\2\2\u0181`\3\2\2\2\u0182\u0183"+
		"\7\u0080\2\2\u0183\u0184\7?\2\2\u0184b\3\2\2\2\u0185\u0186\7>\2\2\u0186"+
		"d\3\2\2\2\u0187\u0188\7>\2\2\u0188\u0189\7?\2\2\u0189f\3\2\2\2\u018a\u018b"+
		"\7@\2\2\u018bh\3\2\2\2\u018c\u018d\7@\2\2\u018d\u018e\7?\2\2\u018ej\3"+
		"\2\2\2\u018f\u0190\7-\2\2\u0190l\3\2\2\2\u0191\u0192\7/\2\2\u0192n\3\2"+
		"\2\2\u0193\u0194\7,\2\2\u0194p\3\2\2\2\u0195\u0196\7\61\2\2\u0196r\3\2"+
		"\2\2\u0197\u0198\7`\2\2\u0198t\3\2\2\2\33\2\u00d7\u00e0\u00e2\u00e8\u00ea"+
		"\u00f1\u00f6\u00fe\u0109\u0115\u011b\u0127\u012b\u012f\u0131\u0139\u0140"+
		"\u0145\u0147\u014f\u0151\u0156\u015b\u0165\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
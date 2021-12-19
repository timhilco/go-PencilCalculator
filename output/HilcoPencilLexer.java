// Generated from /Users/timothyhilgenberg/go/src/github.com/timhilco/go-PencilCalculator/HilcoPencilLexer.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class HilcoPencilLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		IF=1, THEN=2, ELSE=3, CASE=4, CASE_IS=5, CASE_END=6, SWITCH=7, SWITCH_END=8, 
		AND=9, OR=10, NOT=11, INT=12, ID=13, STRING=14, FLOAT=15, DATE=16, ATFUNCTION=17, 
		CLASSNAME=18, PERCENT=19, NEWLINE=20, WS=21, DOT=22, COLON=23, SEMI=24, 
		COMMA=25, EQUALS=26, LBRACKET=27, RBRACKET=28, CURLYLBRACKET=29, CURLYRBRACKET=30, 
		DOTDOT=31, LPAREN=32, RPAREN=33, NOT_EQUALS=34, LT=35, LTE=36, GT=37, 
		GTE=38, PLUS=39, MINUS=40, TIMES=41, DIV=42, EXPONENTIATE=43;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"IF", "THEN", "ELSE", "CASE", "CASE_IS", "CASE_END", "SWITCH", "SWITCH_END", 
			"AND", "OR", "NOT", "DIGIT", "LOWERCASELETTERS", "UPPERCASELETTERS", 
			"INT", "ID", "STRING", "EXPONENT", "HEX_DIGIT", "ESC_SEQ", "OCTAL_ESC", 
			"UNICODE_ESC", "FLOAT", "DATE", "ATFUNCTION", "CLASSNAME", "PERCENT", 
			"NEWLINE", "WS", "DOT", "COLON", "SEMI", "COMMA", "EQUALS", "LBRACKET", 
			"RBRACKET", "CURLYLBRACKET", "CURLYRBRACKET", "DOTDOT", "LPAREN", "RPAREN", 
			"NOT_EQUALS", "LT", "LTE", "GT", "GTE", "PLUS", "MINUS", "TIMES", "DIV", 
			"EXPONENTIATE"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'if'", "'then'", "'else'", "'case'", "'is'", "'endcase'", "'switch'", 
			"'endswitch'", "'AND'", "'OR'", "'NOT'", null, null, null, null, null, 
			null, null, null, null, null, "'.'", "':'", "';'", "','", "'='", "'['", 
			"']'", "'{'", "'}'", "'..'", "'('", "')'", "'~='", "'<'", "'<='", "'>'", 
			"'>='", "'+'", "'-'", "'*'", "'/'", "'^'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "IF", "THEN", "ELSE", "CASE", "CASE_IS", "CASE_END", "SWITCH", 
			"SWITCH_END", "AND", "OR", "NOT", "INT", "ID", "STRING", "FLOAT", "DATE", 
			"ATFUNCTION", "CLASSNAME", "PERCENT", "NEWLINE", "WS", "DOT", "COLON", 
			"SEMI", "COMMA", "EQUALS", "LBRACKET", "RBRACKET", "CURLYLBRACKET", "CURLYRBRACKET", 
			"DOTDOT", "LPAREN", "RPAREN", "NOT_EQUALS", "LT", "LTE", "GT", "GTE", 
			"PLUS", "MINUS", "TIMES", "DIV", "EXPONENTIATE"
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


	public HilcoPencilLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "HilcoPencilLexer.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2-\u016d\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3"+
		"\5\3\5\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\13"+
		"\3\13\3\13\3\f\3\f\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3\20\6\20\u00aa"+
		"\n\20\r\20\16\20\u00ab\3\21\3\21\3\21\3\21\3\21\3\21\3\21\7\21\u00b5\n"+
		"\21\f\21\16\21\u00b8\13\21\3\22\3\22\3\22\7\22\u00bd\n\22\f\22\16\22\u00c0"+
		"\13\22\3\22\3\22\3\23\3\23\5\23\u00c6\n\23\3\23\6\23\u00c9\n\23\r\23\16"+
		"\23\u00ca\3\24\3\24\3\25\3\25\3\25\3\25\5\25\u00d3\n\25\3\26\3\26\3\26"+
		"\3\26\3\26\3\26\3\26\3\26\3\26\5\26\u00de\n\26\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\30\6\30\u00e8\n\30\r\30\16\30\u00e9\3\30\3\30\7\30\u00ee"+
		"\n\30\f\30\16\30\u00f1\13\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3"+
		"\31\5\31\u00fc\n\31\3\31\3\31\5\31\u0100\n\31\3\31\3\31\5\31\u0104\n\31"+
		"\5\31\u0106\n\31\3\31\3\31\3\31\3\31\3\31\3\31\5\31\u010e\n\31\3\31\3"+
		"\31\3\32\3\32\3\32\5\32\u0115\n\32\3\32\3\32\3\32\7\32\u011a\n\32\f\32"+
		"\16\32\u011d\13\32\3\33\3\33\3\33\3\33\3\33\7\33\u0124\n\33\f\33\16\33"+
		"\u0127\13\33\3\34\3\34\5\34\u012b\n\34\3\34\3\34\3\35\5\35\u0130\n\35"+
		"\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\5\36\u013a\n\36\3\36\3\36\3\37"+
		"\3\37\3 \3 \3!\3!\3\"\3\"\3#\3#\3$\3$\3%\3%\3&\3&\3\'\3\'\3(\3(\3(\3)"+
		"\3)\3*\3*\3+\3+\3+\3,\3,\3-\3-\3-\3.\3.\3/\3/\3/\3\60\3\60\3\61\3\61\3"+
		"\62\3\62\3\63\3\63\3\64\3\64\2\2\65\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n"+
		"\23\13\25\f\27\r\31\2\33\2\35\2\37\16!\17#\20%\2\'\2)\2+\2-\2/\21\61\22"+
		"\63\23\65\24\67\259\26;\27=\30?\31A\32C\33E\34G\35I\36K\37M O!Q\"S#U$"+
		"W%Y&[\'](_)a*c+e,g-\3\2\t\4\2))^^\4\2GGgg\4\2--//\5\2\62;CHch\n\2$$))"+
		"^^ddhhppttvv\4\2//\61\61\5\2\13\13\16\16\"\"\2\u0186\2\3\3\2\2\2\2\5\3"+
		"\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2"+
		"\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\37\3\2\2\2\2!\3\2"+
		"\2\2\2#\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67"+
		"\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2"+
		"\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2"+
		"\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]"+
		"\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\3i\3\2"+
		"\2\2\5l\3\2\2\2\7q\3\2\2\2\tv\3\2\2\2\13{\3\2\2\2\r~\3\2\2\2\17\u0086"+
		"\3\2\2\2\21\u008d\3\2\2\2\23\u0097\3\2\2\2\25\u009b\3\2\2\2\27\u009e\3"+
		"\2\2\2\31\u00a2\3\2\2\2\33\u00a4\3\2\2\2\35\u00a6\3\2\2\2\37\u00a9\3\2"+
		"\2\2!\u00ad\3\2\2\2#\u00b9\3\2\2\2%\u00c3\3\2\2\2\'\u00cc\3\2\2\2)\u00d2"+
		"\3\2\2\2+\u00dd\3\2\2\2-\u00df\3\2\2\2/\u00e7\3\2\2\2\61\u00f2\3\2\2\2"+
		"\63\u0111\3\2\2\2\65\u011e\3\2\2\2\67\u012a\3\2\2\29\u012f\3\2\2\2;\u0139"+
		"\3\2\2\2=\u013d\3\2\2\2?\u013f\3\2\2\2A\u0141\3\2\2\2C\u0143\3\2\2\2E"+
		"\u0145\3\2\2\2G\u0147\3\2\2\2I\u0149\3\2\2\2K\u014b\3\2\2\2M\u014d\3\2"+
		"\2\2O\u014f\3\2\2\2Q\u0152\3\2\2\2S\u0154\3\2\2\2U\u0156\3\2\2\2W\u0159"+
		"\3\2\2\2Y\u015b\3\2\2\2[\u015e\3\2\2\2]\u0160\3\2\2\2_\u0163\3\2\2\2a"+
		"\u0165\3\2\2\2c\u0167\3\2\2\2e\u0169\3\2\2\2g\u016b\3\2\2\2ij\7k\2\2j"+
		"k\7h\2\2k\4\3\2\2\2lm\7v\2\2mn\7j\2\2no\7g\2\2op\7p\2\2p\6\3\2\2\2qr\7"+
		"g\2\2rs\7n\2\2st\7u\2\2tu\7g\2\2u\b\3\2\2\2vw\7e\2\2wx\7c\2\2xy\7u\2\2"+
		"yz\7g\2\2z\n\3\2\2\2{|\7k\2\2|}\7u\2\2}\f\3\2\2\2~\177\7g\2\2\177\u0080"+
		"\7p\2\2\u0080\u0081\7f\2\2\u0081\u0082\7e\2\2\u0082\u0083\7c\2\2\u0083"+
		"\u0084\7u\2\2\u0084\u0085\7g\2\2\u0085\16\3\2\2\2\u0086\u0087\7u\2\2\u0087"+
		"\u0088\7y\2\2\u0088\u0089\7k\2\2\u0089\u008a\7v\2\2\u008a\u008b\7e\2\2"+
		"\u008b\u008c\7j\2\2\u008c\20\3\2\2\2\u008d\u008e\7g\2\2\u008e\u008f\7"+
		"p\2\2\u008f\u0090\7f\2\2\u0090\u0091\7u\2\2\u0091\u0092\7y\2\2\u0092\u0093"+
		"\7k\2\2\u0093\u0094\7v\2\2\u0094\u0095\7e\2\2\u0095\u0096\7j\2\2\u0096"+
		"\22\3\2\2\2\u0097\u0098\7C\2\2\u0098\u0099\7P\2\2\u0099\u009a\7F\2\2\u009a"+
		"\24\3\2\2\2\u009b\u009c\7Q\2\2\u009c\u009d\7T\2\2\u009d\26\3\2\2\2\u009e"+
		"\u009f\7P\2\2\u009f\u00a0\7Q\2\2\u00a0\u00a1\7V\2\2\u00a1\30\3\2\2\2\u00a2"+
		"\u00a3\4\62;\2\u00a3\32\3\2\2\2\u00a4\u00a5\4c|\2\u00a5\34\3\2\2\2\u00a6"+
		"\u00a7\4C\\\2\u00a7\36\3\2\2\2\u00a8\u00aa\5\31\r\2\u00a9\u00a8\3\2\2"+
		"\2\u00aa\u00ab\3\2\2\2\u00ab\u00a9\3\2\2\2\u00ab\u00ac\3\2\2\2\u00ac "+
		"\3\2\2\2\u00ad\u00b6\5\33\16\2\u00ae\u00b5\5\33\16\2\u00af\u00b5\5\35"+
		"\17\2\u00b0\u00b5\5\31\r\2\u00b1\u00b5\7a\2\2\u00b2\u00b5\5? \2\u00b3"+
		"\u00b5\7^\2\2\u00b4\u00ae\3\2\2\2\u00b4\u00af\3\2\2\2\u00b4\u00b0\3\2"+
		"\2\2\u00b4\u00b1\3\2\2\2\u00b4\u00b2\3\2\2\2\u00b4\u00b3\3\2\2\2\u00b5"+
		"\u00b8\3\2\2\2\u00b6\u00b4\3\2\2\2\u00b6\u00b7\3\2\2\2\u00b7\"\3\2\2\2"+
		"\u00b8\u00b6\3\2\2\2\u00b9\u00be\7)\2\2\u00ba\u00bd\5)\25\2\u00bb\u00bd"+
		"\n\2\2\2\u00bc\u00ba\3\2\2\2\u00bc\u00bb\3\2\2\2\u00bd\u00c0\3\2\2\2\u00be"+
		"\u00bc\3\2\2\2\u00be\u00bf\3\2\2\2\u00bf\u00c1\3\2\2\2\u00c0\u00be\3\2"+
		"\2\2\u00c1\u00c2\7)\2\2\u00c2$\3\2\2\2\u00c3\u00c5\t\3\2\2\u00c4\u00c6"+
		"\t\4\2\2\u00c5\u00c4\3\2\2\2\u00c5\u00c6\3\2\2\2\u00c6\u00c8\3\2\2\2\u00c7"+
		"\u00c9\4\62;\2\u00c8\u00c7\3\2\2\2\u00c9\u00ca\3\2\2\2\u00ca\u00c8\3\2"+
		"\2\2\u00ca\u00cb\3\2\2\2\u00cb&\3\2\2\2\u00cc\u00cd\t\5\2\2\u00cd(\3\2"+
		"\2\2\u00ce\u00cf\7^\2\2\u00cf\u00d3\t\6\2\2\u00d0\u00d3\5-\27\2\u00d1"+
		"\u00d3\5+\26\2\u00d2\u00ce\3\2\2\2\u00d2\u00d0\3\2\2\2\u00d2\u00d1\3\2"+
		"\2\2\u00d3*\3\2\2\2\u00d4\u00d5\7^\2\2\u00d5\u00d6\4\62\65\2\u00d6\u00d7"+
		"\4\629\2\u00d7\u00de\4\629\2\u00d8\u00d9\7^\2\2\u00d9\u00da\4\629\2\u00da"+
		"\u00de\4\629\2\u00db\u00dc\7^\2\2\u00dc\u00de\4\629\2\u00dd\u00d4\3\2"+
		"\2\2\u00dd\u00d8\3\2\2\2\u00dd\u00db\3\2\2\2\u00de,\3\2\2\2\u00df\u00e0"+
		"\7^\2\2\u00e0\u00e1\7w\2\2\u00e1\u00e2\5\'\24\2\u00e2\u00e3\5\'\24\2\u00e3"+
		"\u00e4\5\'\24\2\u00e4\u00e5\5\'\24\2\u00e5.\3\2\2\2\u00e6\u00e8\5\31\r"+
		"\2\u00e7\u00e6\3\2\2\2\u00e8\u00e9\3\2\2\2\u00e9\u00e7\3\2\2\2\u00e9\u00ea"+
		"\3\2\2\2\u00ea\u00eb\3\2\2\2\u00eb\u00ef\7\60\2\2\u00ec\u00ee\5\31\r\2"+
		"\u00ed\u00ec\3\2\2\2\u00ee\u00f1\3\2\2\2\u00ef\u00ed\3\2\2\2\u00ef\u00f0"+
		"\3\2\2\2\u00f0\60\3\2\2\2\u00f1\u00ef\3\2\2\2\u00f2\u00f3\5K&\2\u00f3"+
		"\u00f4\5\31\r\2\u00f4\u00f5\5\31\r\2\u00f5\u0105\t\7\2\2\u00f6\u00f7\5"+
		"\31\r\2\u00f7\u00f8\5\31\r\2\u00f8\u0106\3\2\2\2\u00f9\u00fc\5\33\16\2"+
		"\u00fa\u00fc\5\35\17\2\u00fb\u00f9\3\2\2\2\u00fb\u00fa\3\2\2\2\u00fc\u00ff"+
		"\3\2\2\2\u00fd\u0100\5\33\16\2\u00fe\u0100\5\35\17\2\u00ff\u00fd\3\2\2"+
		"\2\u00ff\u00fe\3\2\2\2\u0100\u0103\3\2\2\2\u0101\u0104\5\33\16\2\u0102"+
		"\u0104\5\35\17\2\u0103\u0101\3\2\2\2\u0103\u0102\3\2\2\2\u0104\u0106\3"+
		"\2\2\2\u0105\u00f6\3\2\2\2\u0105\u00fb\3\2\2\2\u0106\u0107\3\2\2\2\u0107"+
		"\u0108\t\7\2\2\u0108\u0109\5\31\r\2\u0109\u010d\5\31\r\2\u010a\u010b\5"+
		"\31\r\2\u010b\u010c\5\31\r\2\u010c\u010e\3\2\2\2\u010d\u010a\3\2\2\2\u010d"+
		"\u010e\3\2\2\2\u010e\u010f\3\2\2\2\u010f\u0110\5M\'\2\u0110\62\3\2\2\2"+
		"\u0111\u0114\7B\2\2\u0112\u0115\5\33\16\2\u0113\u0115\5\35\17\2\u0114"+
		"\u0112\3\2\2\2\u0114\u0113\3\2\2\2\u0115\u011b\3\2\2\2\u0116\u011a\5\33"+
		"\16\2\u0117\u011a\5\35\17\2\u0118\u011a\5\31\r\2\u0119\u0116\3\2\2\2\u0119"+
		"\u0117\3\2\2\2\u0119\u0118\3\2\2\2\u011a\u011d\3\2\2\2\u011b\u0119\3\2"+
		"\2\2\u011b\u011c\3\2\2\2\u011c\64\3\2\2\2\u011d\u011b\3\2\2\2\u011e\u0125"+
		"\5\35\17\2\u011f\u0124\5\33\16\2\u0120\u0124\5\35\17\2\u0121\u0124\5\31"+
		"\r\2\u0122\u0124\7a\2\2\u0123\u011f\3\2\2\2\u0123\u0120\3\2\2\2\u0123"+
		"\u0121\3\2\2\2\u0123\u0122\3\2\2\2\u0124\u0127\3\2\2\2\u0125\u0123\3\2"+
		"\2\2\u0125\u0126\3\2\2\2\u0126\66\3\2\2\2\u0127\u0125\3\2\2\2\u0128\u012b"+
		"\5\37\20\2\u0129\u012b\5/\30\2\u012a\u0128\3\2\2\2\u012a\u0129\3\2\2\2"+
		"\u012b\u012c\3\2\2\2\u012c\u012d\7\'\2\2\u012d8\3\2\2\2\u012e\u0130\7"+
		"\17\2\2\u012f\u012e\3\2\2\2\u012f\u0130\3\2\2\2\u0130\u0131\3\2\2\2\u0131"+
		"\u0132\7\f\2\2\u0132\u0133\3\2\2\2\u0133\u0134\b\35\2\2\u0134:\3\2\2\2"+
		"\u0135\u013a\t\b\2\2\u0136\u0137\7\17\2\2\u0137\u013a\7\f\2\2\u0138\u013a"+
		"\7\5\2\2\u0139\u0135\3\2\2\2\u0139\u0136\3\2\2\2\u0139\u0138\3\2\2\2\u013a"+
		"\u013b\3\2\2\2\u013b\u013c\b\36\2\2\u013c<\3\2\2\2\u013d\u013e\7\60\2"+
		"\2\u013e>\3\2\2\2\u013f\u0140\7<\2\2\u0140@\3\2\2\2\u0141\u0142\7=\2\2"+
		"\u0142B\3\2\2\2\u0143\u0144\7.\2\2\u0144D\3\2\2\2\u0145\u0146\7?\2\2\u0146"+
		"F\3\2\2\2\u0147\u0148\7]\2\2\u0148H\3\2\2\2\u0149\u014a\7_\2\2\u014aJ"+
		"\3\2\2\2\u014b\u014c\7}\2\2\u014cL\3\2\2\2\u014d\u014e\7\177\2\2\u014e"+
		"N\3\2\2\2\u014f\u0150\7\60\2\2\u0150\u0151\7\60\2\2\u0151P\3\2\2\2\u0152"+
		"\u0153\7*\2\2\u0153R\3\2\2\2\u0154\u0155\7+\2\2\u0155T\3\2\2\2\u0156\u0157"+
		"\7\u0080\2\2\u0157\u0158\7?\2\2\u0158V\3\2\2\2\u0159\u015a\7>\2\2\u015a"+
		"X\3\2\2\2\u015b\u015c\7>\2\2\u015c\u015d\7?\2\2\u015dZ\3\2\2\2\u015e\u015f"+
		"\7@\2\2\u015f\\\3\2\2\2\u0160\u0161\7@\2\2\u0161\u0162\7?\2\2\u0162^\3"+
		"\2\2\2\u0163\u0164\7-\2\2\u0164`\3\2\2\2\u0165\u0166\7/\2\2\u0166b\3\2"+
		"\2\2\u0167\u0168\7,\2\2\u0168d\3\2\2\2\u0169\u016a\7\61\2\2\u016af\3\2"+
		"\2\2\u016b\u016c\7`\2\2\u016ch\3\2\2\2\33\2\u00ab\u00b4\u00b6\u00bc\u00be"+
		"\u00c5\u00ca\u00d2\u00dd\u00e9\u00ef\u00fb\u00ff\u0103\u0105\u010d\u0114"+
		"\u0119\u011b\u0123\u0125\u012a\u012f\u0139\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
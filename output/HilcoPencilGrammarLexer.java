// Generated from /Users/timothyhilgenberg/go/src/github.com/timhilco/go-PencilCalculator/HilcoPencilGrammarLexer.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class HilcoPencilGrammarLexer extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"CASE", "END_CASE", "IS", "SWITCH", "END_SWITCH", "IF", "THEN", "ELSE", 
			"DOUBLE_COLON", "SEMI_COLON", "NOT", "EXPONENTIAL", "MULTIPLY", "DIVIDE", 
			"ADD", "MINUS", "GREATER_THAN", "GREATER_THAN_EQUAL", "LESS_THAN", "LESS_THAN_EQUAL", 
			"EQUAL", "NOT_EQUAL", "AND", "OR", "COLON", "LBRACKET", "RBRACKET", "CURLYLBRACKET", 
			"CURLYRBRACKET", "LPAREN", "RPAREN", "UNDERSCORE", "PERCENT_SIGN", "AT_SIGN", 
			"COMMA", "DOT", "KEYWORD_TRUE", "KEYWORD_FALSE", "KEYWORD_NIL", "DIGIT", 
			"LOWERCASELETTERS", "UPPERCASELETTERS", "CLASSNAME", "ID", "ATFUNCTION", 
			"INT", "STRING", "FLOAT", "PERCENT", "DATE", "EXPONENT", "HEX_DIGIT", 
			"ESC_SEQ", "OCTAL_ESC", "UNICODE_ESC", "NEWLINE", "WS"
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


	public HilcoPencilGrammarLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "HilcoPencilGrammarLexer.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\63\u018e\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64"+
		"\t\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\3\2\3\2\3\2\3\2"+
		"\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\b\3\b\3\b"+
		"\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\13\3\13\3\f\3\f\3\f\3\f\3\r"+
		"\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3\23"+
		"\3\24\3\24\3\25\3\25\3\25\3\26\3\26\3\27\3\27\3\27\3\30\3\30\3\30\3\30"+
		"\3\31\3\31\3\31\3\32\3\32\3\33\3\33\3\34\3\34\3\35\3\35\3\36\3\36\3\37"+
		"\3\37\3 \3 \3!\3!\3\"\3\"\3#\3#\3$\3$\3%\3%\3&\3&\3&\3&\3&\3\'\3\'\3\'"+
		"\3\'\3\'\3\'\3(\3(\3(\3(\3)\3)\3*\3*\3+\3+\3,\3,\3,\3,\3,\7,\u00ff\n,"+
		"\f,\16,\u0102\13,\3-\3-\3-\3-\3-\3-\3-\7-\u010b\n-\f-\16-\u010e\13-\3"+
		".\3.\3.\5.\u0113\n.\3.\3.\3.\7.\u0118\n.\f.\16.\u011b\13.\3/\6/\u011e"+
		"\n/\r/\16/\u011f\3\60\3\60\3\60\7\60\u0125\n\60\f\60\16\60\u0128\13\60"+
		"\3\60\3\60\3\61\6\61\u012d\n\61\r\61\16\61\u012e\3\61\3\61\7\61\u0133"+
		"\n\61\f\61\16\61\u0136\13\61\3\62\3\62\5\62\u013a\n\62\3\62\3\62\3\63"+
		"\3\63\3\63\3\63\3\63\3\63\3\63\3\63\3\63\5\63\u0147\n\63\3\63\3\63\5\63"+
		"\u014b\n\63\3\63\3\63\5\63\u014f\n\63\5\63\u0151\n\63\3\63\3\63\3\63\3"+
		"\63\3\63\3\63\5\63\u0159\n\63\3\63\3\63\3\64\3\64\5\64\u015f\n\64\3\64"+
		"\6\64\u0162\n\64\r\64\16\64\u0163\3\65\3\65\3\66\3\66\3\66\3\66\5\66\u016c"+
		"\n\66\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\5\67\u0177\n\67\38"+
		"\38\38\38\38\38\38\39\59\u0181\n9\39\39\39\39\3:\3:\3:\3:\5:\u018b\n:"+
		"\3:\3:\2\2;\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33"+
		"\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67"+
		"\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q\2S\2U\2W*Y+[,]-_.a/c\60e\61g\2i\2k"+
		"\2m\2o\2q\62s\63\3\2\t\4\2))^^\4\2//\61\61\4\2GGgg\4\2--//\5\2\62;CHc"+
		"h\n\2$$))^^ddhhppttvv\5\2\13\13\16\16\"\"\2\u01a7\2\3\3\2\2\2\2\5\3\2"+
		"\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21"+
		"\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2"+
		"\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3"+
		"\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3"+
		"\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3"+
		"\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2"+
		"\2\2M\3\2\2\2\2O\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2"+
		"_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\3u\3"+
		"\2\2\2\5z\3\2\2\2\7\u0082\3\2\2\2\t\u0085\3\2\2\2\13\u008c\3\2\2\2\r\u0096"+
		"\3\2\2\2\17\u0099\3\2\2\2\21\u009e\3\2\2\2\23\u00a3\3\2\2\2\25\u00a6\3"+
		"\2\2\2\27\u00a8\3\2\2\2\31\u00ac\3\2\2\2\33\u00ae\3\2\2\2\35\u00b0\3\2"+
		"\2\2\37\u00b2\3\2\2\2!\u00b4\3\2\2\2#\u00b6\3\2\2\2%\u00b8\3\2\2\2\'\u00bb"+
		"\3\2\2\2)\u00bd\3\2\2\2+\u00c0\3\2\2\2-\u00c2\3\2\2\2/\u00c5\3\2\2\2\61"+
		"\u00c9\3\2\2\2\63\u00cc\3\2\2\2\65\u00ce\3\2\2\2\67\u00d0\3\2\2\29\u00d2"+
		"\3\2\2\2;\u00d4\3\2\2\2=\u00d6\3\2\2\2?\u00d8\3\2\2\2A\u00da\3\2\2\2C"+
		"\u00dc\3\2\2\2E\u00de\3\2\2\2G\u00e0\3\2\2\2I\u00e2\3\2\2\2K\u00e4\3\2"+
		"\2\2M\u00e9\3\2\2\2O\u00ef\3\2\2\2Q\u00f3\3\2\2\2S\u00f5\3\2\2\2U\u00f7"+
		"\3\2\2\2W\u00f9\3\2\2\2Y\u0103\3\2\2\2[\u010f\3\2\2\2]\u011d\3\2\2\2_"+
		"\u0121\3\2\2\2a\u012c\3\2\2\2c\u0139\3\2\2\2e\u013d\3\2\2\2g\u015c\3\2"+
		"\2\2i\u0165\3\2\2\2k\u016b\3\2\2\2m\u0176\3\2\2\2o\u0178\3\2\2\2q\u0180"+
		"\3\2\2\2s\u018a\3\2\2\2uv\7e\2\2vw\7c\2\2wx\7u\2\2xy\7g\2\2y\4\3\2\2\2"+
		"z{\7g\2\2{|\7p\2\2|}\7f\2\2}~\7e\2\2~\177\7c\2\2\177\u0080\7u\2\2\u0080"+
		"\u0081\7g\2\2\u0081\6\3\2\2\2\u0082\u0083\7k\2\2\u0083\u0084\7u\2\2\u0084"+
		"\b\3\2\2\2\u0085\u0086\7u\2\2\u0086\u0087\7y\2\2\u0087\u0088\7k\2\2\u0088"+
		"\u0089\7v\2\2\u0089\u008a\7e\2\2\u008a\u008b\7j\2\2\u008b\n\3\2\2\2\u008c"+
		"\u008d\7g\2\2\u008d\u008e\7p\2\2\u008e\u008f\7f\2\2\u008f\u0090\7u\2\2"+
		"\u0090\u0091\7y\2\2\u0091\u0092\7k\2\2\u0092\u0093\7v\2\2\u0093\u0094"+
		"\7e\2\2\u0094\u0095\7j\2\2\u0095\f\3\2\2\2\u0096\u0097\7k\2\2\u0097\u0098"+
		"\7h\2\2\u0098\16\3\2\2\2\u0099\u009a\7v\2\2\u009a\u009b\7j\2\2\u009b\u009c"+
		"\7g\2\2\u009c\u009d\7p\2\2\u009d\20\3\2\2\2\u009e\u009f\7g\2\2\u009f\u00a0"+
		"\7n\2\2\u00a0\u00a1\7u\2\2\u00a1\u00a2\7g\2\2\u00a2\22\3\2\2\2\u00a3\u00a4"+
		"\7<\2\2\u00a4\u00a5\7<\2\2\u00a5\24\3\2\2\2\u00a6\u00a7\7=\2\2\u00a7\26"+
		"\3\2\2\2\u00a8\u00a9\7P\2\2\u00a9\u00aa\7Q\2\2\u00aa\u00ab\7V\2\2\u00ab"+
		"\30\3\2\2\2\u00ac\u00ad\7`\2\2\u00ad\32\3\2\2\2\u00ae\u00af\7,\2\2\u00af"+
		"\34\3\2\2\2\u00b0\u00b1\7\61\2\2\u00b1\36\3\2\2\2\u00b2\u00b3\7-\2\2\u00b3"+
		" \3\2\2\2\u00b4\u00b5\7/\2\2\u00b5\"\3\2\2\2\u00b6\u00b7\7@\2\2\u00b7"+
		"$\3\2\2\2\u00b8\u00b9\7@\2\2\u00b9\u00ba\7?\2\2\u00ba&\3\2\2\2\u00bb\u00bc"+
		"\7>\2\2\u00bc(\3\2\2\2\u00bd\u00be\7>\2\2\u00be\u00bf\7?\2\2\u00bf*\3"+
		"\2\2\2\u00c0\u00c1\7?\2\2\u00c1,\3\2\2\2\u00c2\u00c3\7\u0080\2\2\u00c3"+
		"\u00c4\7?\2\2\u00c4.\3\2\2\2\u00c5\u00c6\7C\2\2\u00c6\u00c7\7P\2\2\u00c7"+
		"\u00c8\7F\2\2\u00c8\60\3\2\2\2\u00c9\u00ca\7Q\2\2\u00ca\u00cb\7T\2\2\u00cb"+
		"\62\3\2\2\2\u00cc\u00cd\7<\2\2\u00cd\64\3\2\2\2\u00ce\u00cf\7]\2\2\u00cf"+
		"\66\3\2\2\2\u00d0\u00d1\7_\2\2\u00d18\3\2\2\2\u00d2\u00d3\7}\2\2\u00d3"+
		":\3\2\2\2\u00d4\u00d5\7\177\2\2\u00d5<\3\2\2\2\u00d6\u00d7\7*\2\2\u00d7"+
		">\3\2\2\2\u00d8\u00d9\7+\2\2\u00d9@\3\2\2\2\u00da\u00db\7a\2\2\u00dbB"+
		"\3\2\2\2\u00dc\u00dd\7\'\2\2\u00ddD\3\2\2\2\u00de\u00df\7B\2\2\u00dfF"+
		"\3\2\2\2\u00e0\u00e1\7.\2\2\u00e1H\3\2\2\2\u00e2\u00e3\7\60\2\2\u00e3"+
		"J\3\2\2\2\u00e4\u00e5\7v\2\2\u00e5\u00e6\7t\2\2\u00e6\u00e7\7w\2\2\u00e7"+
		"\u00e8\7g\2\2\u00e8L\3\2\2\2\u00e9\u00ea\7h\2\2\u00ea\u00eb\7c\2\2\u00eb"+
		"\u00ec\7n\2\2\u00ec\u00ed\7u\2\2\u00ed\u00ee\7g\2\2\u00eeN\3\2\2\2\u00ef"+
		"\u00f0\7p\2\2\u00f0\u00f1\7k\2\2\u00f1\u00f2\7n\2\2\u00f2P\3\2\2\2\u00f3"+
		"\u00f4\4\62;\2\u00f4R\3\2\2\2\u00f5\u00f6\4c|\2\u00f6T\3\2\2\2\u00f7\u00f8"+
		"\4C\\\2\u00f8V\3\2\2\2\u00f9\u0100\5U+\2\u00fa\u00ff\5S*\2\u00fb\u00ff"+
		"\5U+\2\u00fc\u00ff\5Q)\2\u00fd\u00ff\5A!\2\u00fe\u00fa\3\2\2\2\u00fe\u00fb"+
		"\3\2\2\2\u00fe\u00fc\3\2\2\2\u00fe\u00fd\3\2\2\2\u00ff\u0102\3\2\2\2\u0100"+
		"\u00fe\3\2\2\2\u0100\u0101\3\2\2\2\u0101X\3\2\2\2\u0102\u0100\3\2\2\2"+
		"\u0103\u010c\5S*\2\u0104\u010b\5S*\2\u0105\u010b\5U+\2\u0106\u010b\5Q"+
		")\2\u0107\u010b\5A!\2\u0108\u010b\5\63\32\2\u0109\u010b\7^\2\2\u010a\u0104"+
		"\3\2\2\2\u010a\u0105\3\2\2\2\u010a\u0106\3\2\2\2\u010a\u0107\3\2\2\2\u010a"+
		"\u0108\3\2\2\2\u010a\u0109\3\2\2\2\u010b\u010e\3\2\2\2\u010c\u010a\3\2"+
		"\2\2\u010c\u010d\3\2\2\2\u010dZ\3\2\2\2\u010e\u010c\3\2\2\2\u010f\u0112"+
		"\5E#\2\u0110\u0113\5S*\2\u0111\u0113\5U+\2\u0112\u0110\3\2\2\2\u0112\u0111"+
		"\3\2\2\2\u0113\u0119\3\2\2\2\u0114\u0118\5S*\2\u0115\u0118\5U+\2\u0116"+
		"\u0118\5Q)\2\u0117\u0114\3\2\2\2\u0117\u0115\3\2\2\2\u0117\u0116\3\2\2"+
		"\2\u0118\u011b\3\2\2\2\u0119\u0117\3\2\2\2\u0119\u011a\3\2\2\2\u011a\\"+
		"\3\2\2\2\u011b\u0119\3\2\2\2\u011c\u011e\5Q)\2\u011d\u011c\3\2\2\2\u011e"+
		"\u011f\3\2\2\2\u011f\u011d\3\2\2\2\u011f\u0120\3\2\2\2\u0120^\3\2\2\2"+
		"\u0121\u0126\7)\2\2\u0122\u0125\5k\66\2\u0123\u0125\n\2\2\2\u0124\u0122"+
		"\3\2\2\2\u0124\u0123\3\2\2\2\u0125\u0128\3\2\2\2\u0126\u0124\3\2\2\2\u0126"+
		"\u0127\3\2\2\2\u0127\u0129\3\2\2\2\u0128\u0126\3\2\2\2\u0129\u012a\7)"+
		"\2\2\u012a`\3\2\2\2\u012b\u012d\5Q)\2\u012c\u012b\3\2\2\2\u012d\u012e"+
		"\3\2\2\2\u012e\u012c\3\2\2\2\u012e\u012f\3\2\2\2\u012f\u0130\3\2\2\2\u0130"+
		"\u0134\7\60\2\2\u0131\u0133\5Q)\2\u0132\u0131\3\2\2\2\u0133\u0136\3\2"+
		"\2\2\u0134\u0132\3\2\2\2\u0134\u0135\3\2\2\2\u0135b\3\2\2\2\u0136\u0134"+
		"\3\2\2\2\u0137\u013a\5]/\2\u0138\u013a\5a\61\2\u0139\u0137\3\2\2\2\u0139"+
		"\u0138\3\2\2\2\u013a\u013b\3\2\2\2\u013b\u013c\5C\"\2\u013cd\3\2\2\2\u013d"+
		"\u013e\59\35\2\u013e\u013f\5Q)\2\u013f\u0140\5Q)\2\u0140\u0150\t\3\2\2"+
		"\u0141\u0142\5Q)\2\u0142\u0143\5Q)\2\u0143\u0151\3\2\2\2\u0144\u0147\5"+
		"S*\2\u0145\u0147\5U+\2\u0146\u0144\3\2\2\2\u0146\u0145\3\2\2\2\u0147\u014a"+
		"\3\2\2\2\u0148\u014b\5S*\2\u0149\u014b\5U+\2\u014a\u0148\3\2\2\2\u014a"+
		"\u0149\3\2\2\2\u014b\u014e\3\2\2\2\u014c\u014f\5S*\2\u014d\u014f\5U+\2"+
		"\u014e\u014c\3\2\2\2\u014e\u014d\3\2\2\2\u014f\u0151\3\2\2\2\u0150\u0141"+
		"\3\2\2\2\u0150\u0146\3\2\2\2\u0151\u0152\3\2\2\2\u0152\u0153\t\3\2\2\u0153"+
		"\u0154\5Q)\2\u0154\u0158\5Q)\2\u0155\u0156\5Q)\2\u0156\u0157\5Q)\2\u0157"+
		"\u0159\3\2\2\2\u0158\u0155\3\2\2\2\u0158\u0159\3\2\2\2\u0159\u015a\3\2"+
		"\2\2\u015a\u015b\5;\36\2\u015bf\3\2\2\2\u015c\u015e\t\4\2\2\u015d\u015f"+
		"\t\5\2\2\u015e\u015d\3\2\2\2\u015e\u015f\3\2\2\2\u015f\u0161\3\2\2\2\u0160"+
		"\u0162\4\62;\2\u0161\u0160\3\2\2\2\u0162\u0163\3\2\2\2\u0163\u0161\3\2"+
		"\2\2\u0163\u0164\3\2\2\2\u0164h\3\2\2\2\u0165\u0166\t\6\2\2\u0166j\3\2"+
		"\2\2\u0167\u0168\7^\2\2\u0168\u016c\t\7\2\2\u0169\u016c\5o8\2\u016a\u016c"+
		"\5m\67\2\u016b\u0167\3\2\2\2\u016b\u0169\3\2\2\2\u016b\u016a\3\2\2\2\u016c"+
		"l\3\2\2\2\u016d\u016e\7^\2\2\u016e\u016f\4\62\65\2\u016f\u0170\4\629\2"+
		"\u0170\u0177\4\629\2\u0171\u0172\7^\2\2\u0172\u0173\4\629\2\u0173\u0177"+
		"\4\629\2\u0174\u0175\7^\2\2\u0175\u0177\4\629\2\u0176\u016d\3\2\2\2\u0176"+
		"\u0171\3\2\2\2\u0176\u0174\3\2\2\2\u0177n\3\2\2\2\u0178\u0179\7^\2\2\u0179"+
		"\u017a\7w\2\2\u017a\u017b\5i\65\2\u017b\u017c\5i\65\2\u017c\u017d\5i\65"+
		"\2\u017d\u017e\5i\65\2\u017ep\3\2\2\2\u017f\u0181\7\17\2\2\u0180\u017f"+
		"\3\2\2\2\u0180\u0181\3\2\2\2\u0181\u0182\3\2\2\2\u0182\u0183\7\f\2\2\u0183"+
		"\u0184\3\2\2\2\u0184\u0185\b9\2\2\u0185r\3\2\2\2\u0186\u018b\t\b\2\2\u0187"+
		"\u0188\7\17\2\2\u0188\u018b\7\f\2\2\u0189\u018b\7\5\2\2\u018a\u0186\3"+
		"\2\2\2\u018a\u0187\3\2\2\2\u018a\u0189\3\2\2\2\u018b\u018c\3\2\2\2\u018c"+
		"\u018d\b:\2\2\u018dt\3\2\2\2\33\2\u00fe\u0100\u010a\u010c\u0112\u0117"+
		"\u0119\u011f\u0124\u0126\u012e\u0134\u0139\u0146\u014a\u014e\u0150\u0158"+
		"\u015e\u0163\u016b\u0176\u0180\u018a\3\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
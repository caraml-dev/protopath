// Generated from /Users/tio.pramayudi/Documents/dsp/jsonpath/parser/JsonPath.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class JsonPathLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		IDENTIFIER=10, STRING=11, INT=12, WS=13, COLON=14, ROOT_IDENTIFIER=15, 
		LOCAL_IDENTIFIER=16, GREATER=17, GREATER_EQ=18, LESS=19, LESS_EQ=20, EQ=21, 
		NEQ=22, AND=23, OR=24;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
			"IDENTIFIER", "STRING", "INT", "WS", "COLON", "ROOT_IDENTIFIER", "LOCAL_IDENTIFIER", 
			"GREATER", "GREATER_EQ", "LESS", "LESS_EQ", "EQ", "NEQ", "AND", "OR"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'.'", "'[*]'", "'['", "']'", "'[?('", "')]'", "'[@.length-'", 
			"'[:]'", "','", null, null, null, null, "':'", "'$.'", "'@.'", "'>'", 
			"'>='", "'<'", "'<='", "'=='", "'!='", "'&&'", "'||'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, "IDENTIFIER", 
			"STRING", "INT", "WS", "COLON", "ROOT_IDENTIFIER", "LOCAL_IDENTIFIER", 
			"GREATER", "GREATER_EQ", "LESS", "LESS_EQ", "EQ", "NEQ", "AND", "OR"
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


	public JsonPathLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "JsonPath.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\32\u008f\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\3\2\3\2\3\3\3\3\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\6\3\6\3\7\3\7"+
		"\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\n\3"+
		"\n\3\13\3\13\7\13X\n\13\f\13\16\13[\13\13\3\f\3\f\3\f\3\f\3\r\3\r\3\r"+
		"\7\rd\n\r\f\r\16\rg\13\r\5\ri\n\r\3\16\6\16l\n\16\r\16\16\16m\3\16\3\16"+
		"\3\17\3\17\3\20\3\20\3\20\3\21\3\21\3\21\3\22\3\22\3\23\3\23\3\23\3\24"+
		"\3\24\3\25\3\25\3\25\3\26\3\26\3\26\3\27\3\27\3\27\3\30\3\30\3\30\3\31"+
		"\3\31\3\31\2\2\32\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31"+
		"\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\3\2\7\5\2"+
		"C\\aac|\6\2\62;C\\aac|\3\2\63;\3\2\62;\5\2\13\f\17\17\"\"\2\u0092\2\3"+
		"\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2"+
		"\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31"+
		"\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2"+
		"\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2"+
		"\61\3\2\2\2\3\63\3\2\2\2\5\65\3\2\2\2\79\3\2\2\2\t;\3\2\2\2\13=\3\2\2"+
		"\2\rA\3\2\2\2\17D\3\2\2\2\21O\3\2\2\2\23S\3\2\2\2\25U\3\2\2\2\27\\\3\2"+
		"\2\2\31h\3\2\2\2\33k\3\2\2\2\35q\3\2\2\2\37s\3\2\2\2!v\3\2\2\2#y\3\2\2"+
		"\2%{\3\2\2\2\'~\3\2\2\2)\u0080\3\2\2\2+\u0083\3\2\2\2-\u0086\3\2\2\2/"+
		"\u0089\3\2\2\2\61\u008c\3\2\2\2\63\64\7\60\2\2\64\4\3\2\2\2\65\66\7]\2"+
		"\2\66\67\7,\2\2\678\7_\2\28\6\3\2\2\29:\7]\2\2:\b\3\2\2\2;<\7_\2\2<\n"+
		"\3\2\2\2=>\7]\2\2>?\7A\2\2?@\7*\2\2@\f\3\2\2\2AB\7+\2\2BC\7_\2\2C\16\3"+
		"\2\2\2DE\7]\2\2EF\7B\2\2FG\7\60\2\2GH\7n\2\2HI\7g\2\2IJ\7p\2\2JK\7i\2"+
		"\2KL\7v\2\2LM\7j\2\2MN\7/\2\2N\20\3\2\2\2OP\7]\2\2PQ\7<\2\2QR\7_\2\2R"+
		"\22\3\2\2\2ST\7.\2\2T\24\3\2\2\2UY\t\2\2\2VX\t\3\2\2WV\3\2\2\2X[\3\2\2"+
		"\2YW\3\2\2\2YZ\3\2\2\2Z\26\3\2\2\2[Y\3\2\2\2\\]\7)\2\2]^\5\25\13\2^_\7"+
		")\2\2_\30\3\2\2\2`i\7\62\2\2ae\t\4\2\2bd\t\5\2\2cb\3\2\2\2dg\3\2\2\2e"+
		"c\3\2\2\2ef\3\2\2\2fi\3\2\2\2ge\3\2\2\2h`\3\2\2\2ha\3\2\2\2i\32\3\2\2"+
		"\2jl\t\6\2\2kj\3\2\2\2lm\3\2\2\2mk\3\2\2\2mn\3\2\2\2no\3\2\2\2op\b\16"+
		"\2\2p\34\3\2\2\2qr\7<\2\2r\36\3\2\2\2st\7&\2\2tu\7\60\2\2u \3\2\2\2vw"+
		"\7B\2\2wx\7\60\2\2x\"\3\2\2\2yz\7@\2\2z$\3\2\2\2{|\7@\2\2|}\7?\2\2}&\3"+
		"\2\2\2~\177\7>\2\2\177(\3\2\2\2\u0080\u0081\7>\2\2\u0081\u0082\7?\2\2"+
		"\u0082*\3\2\2\2\u0083\u0084\7?\2\2\u0084\u0085\7?\2\2\u0085,\3\2\2\2\u0086"+
		"\u0087\7#\2\2\u0087\u0088\7?\2\2\u0088.\3\2\2\2\u0089\u008a\7(\2\2\u008a"+
		"\u008b\7(\2\2\u008b\60\3\2\2\2\u008c\u008d\7~\2\2\u008d\u008e\7~\2\2\u008e"+
		"\62\3\2\2\2\7\2Yehm\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
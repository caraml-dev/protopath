// Generated from /Users/tio.pramayudi/Documents/dsp/protopath/parser/JsonPath.g4 by ANTLR 4.9.2
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
		T__9=10, T__10=11, IDENTIFIER=12, STRING=13, INT=14, WS=15, COLON=16, 
		ROOT_IDENTIFIER=17, LOCAL_IDENTIFIER=18, GREATER=19, GREATER_EQ=20, LESS=21, 
		LESS_EQ=22, EQ=23, NEQ=24, AND=25, OR=26;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
			"T__9", "T__10", "IDENTIFIER", "STRING", "INT", "WS", "COLON", "ROOT_IDENTIFIER", 
			"LOCAL_IDENTIFIER", "GREATER", "GREATER_EQ", "LESS", "LESS_EQ", "EQ", 
			"NEQ", "AND", "OR"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'.'", "'[*]'", "'['", "']'", "'[?('", "')]'", "'[@.length-'", 
			"'[:]'", "','", "'true'", "'false'", null, null, null, null, "':'", "'$.'", 
			"'@.'", "'>'", "'>='", "'<'", "'<='", "'=='", "'!='", "'&&'", "'||'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			"IDENTIFIER", "STRING", "INT", "WS", "COLON", "ROOT_IDENTIFIER", "LOCAL_IDENTIFIER", 
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\34\u009e\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\3\2\3\2\3\3\3\3\3\3\3\3\3\4\3\4\3\5\3\5\3\6"+
		"\3\6\3\6\3\6\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\t\3\t\3\t\3\t\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\r\3\r\7\rg\n\r\f\r\16\rj\13\r\3\16\3\16\3\16\3\16\3\17\3\17\3\17"+
		"\7\17s\n\17\f\17\16\17v\13\17\5\17x\n\17\3\20\6\20{\n\20\r\20\16\20|\3"+
		"\20\3\20\3\21\3\21\3\22\3\22\3\22\3\23\3\23\3\23\3\24\3\24\3\25\3\25\3"+
		"\25\3\26\3\26\3\27\3\27\3\27\3\30\3\30\3\30\3\31\3\31\3\31\3\32\3\32\3"+
		"\32\3\33\3\33\3\33\2\2\34\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25"+
		"\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32"+
		"\63\33\65\34\3\2\7\5\2C\\aac|\6\2\62;C\\aac|\3\2\63;\3\2\62;\5\2\13\f"+
		"\17\17\"\"\2\u00a1\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13"+
		"\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2"+
		"\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2"+
		"!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3"+
		"\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\3\67\3\2\2\2"+
		"\59\3\2\2\2\7=\3\2\2\2\t?\3\2\2\2\13A\3\2\2\2\rE\3\2\2\2\17H\3\2\2\2\21"+
		"S\3\2\2\2\23W\3\2\2\2\25Y\3\2\2\2\27^\3\2\2\2\31d\3\2\2\2\33k\3\2\2\2"+
		"\35w\3\2\2\2\37z\3\2\2\2!\u0080\3\2\2\2#\u0082\3\2\2\2%\u0085\3\2\2\2"+
		"\'\u0088\3\2\2\2)\u008a\3\2\2\2+\u008d\3\2\2\2-\u008f\3\2\2\2/\u0092\3"+
		"\2\2\2\61\u0095\3\2\2\2\63\u0098\3\2\2\2\65\u009b\3\2\2\2\678\7\60\2\2"+
		"8\4\3\2\2\29:\7]\2\2:;\7,\2\2;<\7_\2\2<\6\3\2\2\2=>\7]\2\2>\b\3\2\2\2"+
		"?@\7_\2\2@\n\3\2\2\2AB\7]\2\2BC\7A\2\2CD\7*\2\2D\f\3\2\2\2EF\7+\2\2FG"+
		"\7_\2\2G\16\3\2\2\2HI\7]\2\2IJ\7B\2\2JK\7\60\2\2KL\7n\2\2LM\7g\2\2MN\7"+
		"p\2\2NO\7i\2\2OP\7v\2\2PQ\7j\2\2QR\7/\2\2R\20\3\2\2\2ST\7]\2\2TU\7<\2"+
		"\2UV\7_\2\2V\22\3\2\2\2WX\7.\2\2X\24\3\2\2\2YZ\7v\2\2Z[\7t\2\2[\\\7w\2"+
		"\2\\]\7g\2\2]\26\3\2\2\2^_\7h\2\2_`\7c\2\2`a\7n\2\2ab\7u\2\2bc\7g\2\2"+
		"c\30\3\2\2\2dh\t\2\2\2eg\t\3\2\2fe\3\2\2\2gj\3\2\2\2hf\3\2\2\2hi\3\2\2"+
		"\2i\32\3\2\2\2jh\3\2\2\2kl\7)\2\2lm\5\31\r\2mn\7)\2\2n\34\3\2\2\2ox\7"+
		"\62\2\2pt\t\4\2\2qs\t\5\2\2rq\3\2\2\2sv\3\2\2\2tr\3\2\2\2tu\3\2\2\2ux"+
		"\3\2\2\2vt\3\2\2\2wo\3\2\2\2wp\3\2\2\2x\36\3\2\2\2y{\t\6\2\2zy\3\2\2\2"+
		"{|\3\2\2\2|z\3\2\2\2|}\3\2\2\2}~\3\2\2\2~\177\b\20\2\2\177 \3\2\2\2\u0080"+
		"\u0081\7<\2\2\u0081\"\3\2\2\2\u0082\u0083\7&\2\2\u0083\u0084\7\60\2\2"+
		"\u0084$\3\2\2\2\u0085\u0086\7B\2\2\u0086\u0087\7\60\2\2\u0087&\3\2\2\2"+
		"\u0088\u0089\7@\2\2\u0089(\3\2\2\2\u008a\u008b\7@\2\2\u008b\u008c\7?\2"+
		"\2\u008c*\3\2\2\2\u008d\u008e\7>\2\2\u008e,\3\2\2\2\u008f\u0090\7>\2\2"+
		"\u0090\u0091\7?\2\2\u0091.\3\2\2\2\u0092\u0093\7?\2\2\u0093\u0094\7?\2"+
		"\2\u0094\60\3\2\2\2\u0095\u0096\7#\2\2\u0096\u0097\7?\2\2\u0097\62\3\2"+
		"\2\2\u0098\u0099\7(\2\2\u0099\u009a\7(\2\2\u009a\64\3\2\2\2\u009b\u009c"+
		"\7~\2\2\u009c\u009d\7~\2\2\u009d\66\3\2\2\2\7\2htw|\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
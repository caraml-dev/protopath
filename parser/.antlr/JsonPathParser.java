// Generated from /Users/tio.pramayudi/Documents/dsp/protopath/parser/JsonPath.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class JsonPathParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, IDENTIFIER=12, STRING=13, INT=14, WS=15, COLON=16, 
		ROOT_IDENTIFIER=17, LOCAL_IDENTIFIER=18, GREATER=19, GREATER_EQ=20, LESS=21, 
		LESS_EQ=22, EQ=23, NEQ=24, AND=25, OR=26;
	public static final int
		RULE_jsonpath = 0, RULE_root = 1, RULE_fieldAccess = 2, RULE_localFieldAccess = 3, 
		RULE_arrayQualifier = 4, RULE_arrayRange = 5, RULE_intSequence = 6, RULE_queryFieldValue = 7, 
		RULE_filterOperation = 8, RULE_queryExpr = 9, RULE_dbl = 10, RULE_bool_type = 11;
	private static String[] makeRuleNames() {
		return new String[] {
			"jsonpath", "root", "fieldAccess", "localFieldAccess", "arrayQualifier", 
			"arrayRange", "intSequence", "queryFieldValue", "filterOperation", "queryExpr", 
			"dbl", "bool_type"
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

	@Override
	public String getGrammarFileName() { return "JsonPath.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public JsonPathParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class JsonpathContext extends ParserRuleContext {
		public RootContext root() {
			return getRuleContext(RootContext.class,0);
		}
		public JsonpathContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_jsonpath; }
	}

	public final JsonpathContext jsonpath() throws RecognitionException {
		JsonpathContext _localctx = new JsonpathContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_jsonpath);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(24);
			root(0);
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

	public static class RootContext extends ParserRuleContext {
		public TerminalNode ROOT_IDENTIFIER() { return getToken(JsonPathParser.ROOT_IDENTIFIER, 0); }
		public List<FieldAccessContext> fieldAccess() {
			return getRuleContexts(FieldAccessContext.class);
		}
		public FieldAccessContext fieldAccess(int i) {
			return getRuleContext(FieldAccessContext.class,i);
		}
		public RootContext root() {
			return getRuleContext(RootContext.class,0);
		}
		public RootContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_root; }
	}

	public final RootContext root() throws RecognitionException {
		return root(0);
	}

	private RootContext root(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		RootContext _localctx = new RootContext(_ctx, _parentState);
		RootContext _prevctx = _localctx;
		int _startState = 2;
		enterRecursionRule(_localctx, 2, RULE_root, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(27);
			match(ROOT_IDENTIFIER);
			setState(28);
			fieldAccess();
			}
			_ctx.stop = _input.LT(-1);
			setState(39);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new RootContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_root);
					setState(30);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(33); 
					_errHandler.sync(this);
					_alt = 1;
					do {
						switch (_alt) {
						case 1:
							{
							{
							setState(31);
							match(T__0);
							setState(32);
							fieldAccess();
							}
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						setState(35); 
						_errHandler.sync(this);
						_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
					} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
					}
					} 
				}
				setState(41);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
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

	public static class FieldAccessContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(JsonPathParser.IDENTIFIER, 0); }
		public ArrayQualifierContext arrayQualifier() {
			return getRuleContext(ArrayQualifierContext.class,0);
		}
		public FieldAccessContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fieldAccess; }
	}

	public final FieldAccessContext fieldAccess() throws RecognitionException {
		FieldAccessContext _localctx = new FieldAccessContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_fieldAccess);
		try {
			setState(48);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(42);
				match(IDENTIFIER);
				setState(43);
				arrayQualifier();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(44);
				match(IDENTIFIER);
				setState(45);
				match(T__0);
				setState(46);
				arrayQualifier();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(47);
				match(IDENTIFIER);
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

	public static class LocalFieldAccessContext extends ParserRuleContext {
		public TerminalNode LOCAL_IDENTIFIER() { return getToken(JsonPathParser.LOCAL_IDENTIFIER, 0); }
		public List<FieldAccessContext> fieldAccess() {
			return getRuleContexts(FieldAccessContext.class);
		}
		public FieldAccessContext fieldAccess(int i) {
			return getRuleContext(FieldAccessContext.class,i);
		}
		public LocalFieldAccessContext localFieldAccess() {
			return getRuleContext(LocalFieldAccessContext.class,0);
		}
		public LocalFieldAccessContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_localFieldAccess; }
	}

	public final LocalFieldAccessContext localFieldAccess() throws RecognitionException {
		return localFieldAccess(0);
	}

	private LocalFieldAccessContext localFieldAccess(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		LocalFieldAccessContext _localctx = new LocalFieldAccessContext(_ctx, _parentState);
		LocalFieldAccessContext _prevctx = _localctx;
		int _startState = 6;
		enterRecursionRule(_localctx, 6, RULE_localFieldAccess, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(51);
			match(LOCAL_IDENTIFIER);
			setState(52);
			fieldAccess();
			}
			_ctx.stop = _input.LT(-1);
			setState(63);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new LocalFieldAccessContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_localFieldAccess);
					setState(54);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(57); 
					_errHandler.sync(this);
					_alt = 1;
					do {
						switch (_alt) {
						case 1:
							{
							{
							setState(55);
							match(T__0);
							setState(56);
							fieldAccess();
							}
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						setState(59); 
						_errHandler.sync(this);
						_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
					} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
					}
					} 
				}
				setState(65);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
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

	public static class ArrayQualifierContext extends ParserRuleContext {
		public ArrayQualifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrayQualifier; }
	 
		public ArrayQualifierContext() { }
		public void copyFrom(ArrayQualifierContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class GetByQueryContext extends ArrayQualifierContext {
		public FilterOperationContext filterOperation() {
			return getRuleContext(FilterOperationContext.class,0);
		}
		public GetByQueryContext(ArrayQualifierContext ctx) { copyFrom(ctx); }
	}
	public static class GetAllContext extends ArrayQualifierContext {
		public GetAllContext(ArrayQualifierContext ctx) { copyFrom(ctx); }
	}
	public static class GetByIndexBackwardContext extends ArrayQualifierContext {
		public TerminalNode INT() { return getToken(JsonPathParser.INT, 0); }
		public GetByIndexBackwardContext(ArrayQualifierContext ctx) { copyFrom(ctx); }
	}
	public static class GetByIndicesContext extends ArrayQualifierContext {
		public IntSequenceContext intSequence() {
			return getRuleContext(IntSequenceContext.class,0);
		}
		public GetByIndicesContext(ArrayQualifierContext ctx) { copyFrom(ctx); }
	}
	public static class GetByRangeContext extends ArrayQualifierContext {
		public ArrayRangeContext arrayRange() {
			return getRuleContext(ArrayRangeContext.class,0);
		}
		public GetByRangeContext(ArrayQualifierContext ctx) { copyFrom(ctx); }
	}

	public final ArrayQualifierContext arrayQualifier() throws RecognitionException {
		ArrayQualifierContext _localctx = new ArrayQualifierContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_arrayQualifier);
		try {
			setState(79);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				_localctx = new GetAllContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(66);
				match(T__1);
				}
				break;
			case 2:
				_localctx = new GetByIndicesContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(67);
				match(T__2);
				setState(68);
				intSequence();
				setState(69);
				match(T__3);
				}
				break;
			case 3:
				_localctx = new GetByRangeContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(71);
				arrayRange();
				}
				break;
			case 4:
				_localctx = new GetByQueryContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(72);
				match(T__4);
				setState(73);
				filterOperation();
				setState(74);
				match(T__5);
				}
				break;
			case 5:
				_localctx = new GetByIndexBackwardContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(76);
				match(T__6);
				setState(77);
				match(INT);
				setState(78);
				match(T__3);
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

	public static class ArrayRangeContext extends ParserRuleContext {
		public ArrayRangeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrayRange; }
	 
		public ArrayRangeContext() { }
		public void copyFrom(ArrayRangeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class GivenStartEndIdxRangeContext extends ArrayRangeContext {
		public TerminalNode COLON() { return getToken(JsonPathParser.COLON, 0); }
		public List<TerminalNode> INT() { return getTokens(JsonPathParser.INT); }
		public TerminalNode INT(int i) {
			return getToken(JsonPathParser.INT, i);
		}
		public GivenStartEndIdxRangeContext(ArrayRangeContext ctx) { copyFrom(ctx); }
	}
	public static class AllRangeContext extends ArrayRangeContext {
		public AllRangeContext(ArrayRangeContext ctx) { copyFrom(ctx); }
	}
	public static class GivenOnlyEndIdxRangeContext extends ArrayRangeContext {
		public TerminalNode COLON() { return getToken(JsonPathParser.COLON, 0); }
		public List<TerminalNode> INT() { return getTokens(JsonPathParser.INT); }
		public TerminalNode INT(int i) {
			return getToken(JsonPathParser.INT, i);
		}
		public GivenOnlyEndIdxRangeContext(ArrayRangeContext ctx) { copyFrom(ctx); }
	}
	public static class GivenOnlyStartIdxRangeContext extends ArrayRangeContext {
		public TerminalNode COLON() { return getToken(JsonPathParser.COLON, 0); }
		public List<TerminalNode> INT() { return getTokens(JsonPathParser.INT); }
		public TerminalNode INT(int i) {
			return getToken(JsonPathParser.INT, i);
		}
		public GivenOnlyStartIdxRangeContext(ArrayRangeContext ctx) { copyFrom(ctx); }
	}

	public final ArrayRangeContext arrayRange() throws RecognitionException {
		ArrayRangeContext _localctx = new ArrayRangeContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_arrayRange);
		int _la;
		try {
			setState(111);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				_localctx = new AllRangeContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(81);
				match(T__7);
				}
				break;
			case 2:
				_localctx = new GivenStartEndIdxRangeContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(82);
				match(T__2);
				setState(84); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(83);
					match(INT);
					}
					}
					setState(86); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==INT );
				setState(88);
				match(COLON);
				setState(90); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(89);
					match(INT);
					}
					}
					setState(92); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==INT );
				setState(94);
				match(T__3);
				}
				break;
			case 3:
				_localctx = new GivenOnlyEndIdxRangeContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(95);
				match(T__2);
				setState(96);
				match(COLON);
				setState(98); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(97);
					match(INT);
					}
					}
					setState(100); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==INT );
				setState(102);
				match(T__3);
				}
				break;
			case 4:
				_localctx = new GivenOnlyStartIdxRangeContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(103);
				match(T__2);
				setState(105); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(104);
					match(INT);
					}
					}
					setState(107); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==INT );
				setState(109);
				match(COLON);
				setState(110);
				match(T__3);
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

	public static class IntSequenceContext extends ParserRuleContext {
		public List<TerminalNode> INT() { return getTokens(JsonPathParser.INT); }
		public TerminalNode INT(int i) {
			return getToken(JsonPathParser.INT, i);
		}
		public IntSequenceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_intSequence; }
	}

	public final IntSequenceContext intSequence() throws RecognitionException {
		IntSequenceContext _localctx = new IntSequenceContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_intSequence);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(113);
			match(INT);
			setState(118);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__8) {
				{
				{
				setState(114);
				match(T__8);
				setState(115);
				match(INT);
				}
				}
				setState(120);
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

	public static class QueryFieldValueContext extends ParserRuleContext {
		public QueryFieldValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_queryFieldValue; }
	 
		public QueryFieldValueContext() { }
		public void copyFrom(QueryFieldValueContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class QueryFieldStringValueContext extends QueryFieldValueContext {
		public TerminalNode STRING() { return getToken(JsonPathParser.STRING, 0); }
		public QueryFieldStringValueContext(QueryFieldValueContext ctx) { copyFrom(ctx); }
	}
	public static class QueryFieldValueLocalAccessContext extends QueryFieldValueContext {
		public LocalFieldAccessContext localFieldAccess() {
			return getRuleContext(LocalFieldAccessContext.class,0);
		}
		public QueryFieldValueLocalAccessContext(QueryFieldValueContext ctx) { copyFrom(ctx); }
	}
	public static class QueryFieldValueRootAccessContext extends QueryFieldValueContext {
		public RootContext root() {
			return getRuleContext(RootContext.class,0);
		}
		public QueryFieldValueRootAccessContext(QueryFieldValueContext ctx) { copyFrom(ctx); }
	}
	public static class QueryFieldDoubleValueContext extends QueryFieldValueContext {
		public DblContext dbl() {
			return getRuleContext(DblContext.class,0);
		}
		public QueryFieldDoubleValueContext(QueryFieldValueContext ctx) { copyFrom(ctx); }
	}
	public static class QueryFieldBoolValueContext extends QueryFieldValueContext {
		public Bool_typeContext bool_type() {
			return getRuleContext(Bool_typeContext.class,0);
		}
		public QueryFieldBoolValueContext(QueryFieldValueContext ctx) { copyFrom(ctx); }
	}

	public final QueryFieldValueContext queryFieldValue() throws RecognitionException {
		QueryFieldValueContext _localctx = new QueryFieldValueContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_queryFieldValue);
		try {
			setState(126);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LOCAL_IDENTIFIER:
				_localctx = new QueryFieldValueLocalAccessContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(121);
				localFieldAccess(0);
				}
				break;
			case ROOT_IDENTIFIER:
				_localctx = new QueryFieldValueRootAccessContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(122);
				root(0);
				}
				break;
			case T__0:
			case INT:
				_localctx = new QueryFieldDoubleValueContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(123);
				dbl();
				}
				break;
			case STRING:
				_localctx = new QueryFieldStringValueContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(124);
				match(STRING);
				}
				break;
			case T__9:
			case T__10:
				_localctx = new QueryFieldBoolValueContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(125);
				bool_type();
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

	public static class FilterOperationContext extends ParserRuleContext {
		public QueryExprContext queryExpr() {
			return getRuleContext(QueryExprContext.class,0);
		}
		public FilterOperationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_filterOperation; }
	}

	public final FilterOperationContext filterOperation() throws RecognitionException {
		FilterOperationContext _localctx = new FilterOperationContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_filterOperation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(128);
			queryExpr(0);
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

	public static class QueryExprContext extends ParserRuleContext {
		public QueryExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_queryExpr; }
	 
		public QueryExprContext() { }
		public void copyFrom(QueryExprContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class QueryWithLogicalOpContext extends QueryExprContext {
		public Token logical_op;
		public List<QueryExprContext> queryExpr() {
			return getRuleContexts(QueryExprContext.class);
		}
		public QueryExprContext queryExpr(int i) {
			return getRuleContext(QueryExprContext.class,i);
		}
		public List<TerminalNode> AND() { return getTokens(JsonPathParser.AND); }
		public TerminalNode AND(int i) {
			return getToken(JsonPathParser.AND, i);
		}
		public List<TerminalNode> OR() { return getTokens(JsonPathParser.OR); }
		public TerminalNode OR(int i) {
			return getToken(JsonPathParser.OR, i);
		}
		public QueryWithLogicalOpContext(QueryExprContext ctx) { copyFrom(ctx); }
	}
	public static class QueryExistenceContext extends QueryExprContext {
		public QueryFieldValueContext queryFieldValue() {
			return getRuleContext(QueryFieldValueContext.class,0);
		}
		public QueryExistenceContext(QueryExprContext ctx) { copyFrom(ctx); }
	}
	public static class QueryWithComparatorContext extends QueryExprContext {
		public Token comparator_op;
		public List<QueryFieldValueContext> queryFieldValue() {
			return getRuleContexts(QueryFieldValueContext.class);
		}
		public QueryFieldValueContext queryFieldValue(int i) {
			return getRuleContext(QueryFieldValueContext.class,i);
		}
		public TerminalNode GREATER() { return getToken(JsonPathParser.GREATER, 0); }
		public TerminalNode GREATER_EQ() { return getToken(JsonPathParser.GREATER_EQ, 0); }
		public TerminalNode LESS() { return getToken(JsonPathParser.LESS, 0); }
		public TerminalNode LESS_EQ() { return getToken(JsonPathParser.LESS_EQ, 0); }
		public TerminalNode EQ() { return getToken(JsonPathParser.EQ, 0); }
		public TerminalNode NEQ() { return getToken(JsonPathParser.NEQ, 0); }
		public QueryWithComparatorContext(QueryExprContext ctx) { copyFrom(ctx); }
	}

	public final QueryExprContext queryExpr() throws RecognitionException {
		return queryExpr(0);
	}

	private QueryExprContext queryExpr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		QueryExprContext _localctx = new QueryExprContext(_ctx, _parentState);
		QueryExprContext _prevctx = _localctx;
		int _startState = 18;
		enterRecursionRule(_localctx, 18, RULE_queryExpr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(136);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				{
				_localctx = new QueryExistenceContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(131);
				queryFieldValue();
				}
				break;
			case 2:
				{
				_localctx = new QueryWithComparatorContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(132);
				queryFieldValue();
				setState(133);
				((QueryWithComparatorContext)_localctx).comparator_op = _input.LT(1);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << GREATER) | (1L << GREATER_EQ) | (1L << LESS) | (1L << LESS_EQ) | (1L << EQ) | (1L << NEQ))) != 0)) ) {
					((QueryWithComparatorContext)_localctx).comparator_op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(134);
				queryFieldValue();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(147);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new QueryWithLogicalOpContext(new QueryExprContext(_parentctx, _parentState));
					pushNewRecursionContext(_localctx, _startState, RULE_queryExpr);
					setState(138);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(141); 
					_errHandler.sync(this);
					_alt = 1;
					do {
						switch (_alt) {
						case 1:
							{
							{
							setState(139);
							((QueryWithLogicalOpContext)_localctx).logical_op = _input.LT(1);
							_la = _input.LA(1);
							if ( !(_la==AND || _la==OR) ) {
								((QueryWithLogicalOpContext)_localctx).logical_op = (Token)_errHandler.recoverInline(this);
							}
							else {
								if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
								_errHandler.reportMatch(this);
								consume();
							}
							setState(140);
							queryExpr(0);
							}
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						setState(143); 
						_errHandler.sync(this);
						_alt = getInterpreter().adaptivePredict(_input,14,_ctx);
					} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
					}
					} 
				}
				setState(149);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
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

	public static class DblContext extends ParserRuleContext {
		public List<TerminalNode> INT() { return getTokens(JsonPathParser.INT); }
		public TerminalNode INT(int i) {
			return getToken(JsonPathParser.INT, i);
		}
		public DblContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dbl; }
	}

	public final DblContext dbl() throws RecognitionException {
		DblContext _localctx = new DblContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_dbl);
		int _la;
		try {
			int _alt;
			setState(172);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(151); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(150);
					match(INT);
					}
					}
					setState(153); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==INT );
				setState(155);
				match(T__0);
				setState(157); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(156);
						match(INT);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(159); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(161);
				match(T__0);
				setState(163); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(162);
						match(INT);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(165); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(168); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(167);
						match(INT);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(170); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,19,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
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

	public static class Bool_typeContext extends ParserRuleContext {
		public Bool_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bool_type; }
	}

	public final Bool_typeContext bool_type() throws RecognitionException {
		Bool_typeContext _localctx = new Bool_typeContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_bool_type);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(174);
			_la = _input.LA(1);
			if ( !(_la==T__9 || _la==T__10) ) {
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
			return root_sempred((RootContext)_localctx, predIndex);
		case 3:
			return localFieldAccess_sempred((LocalFieldAccessContext)_localctx, predIndex);
		case 9:
			return queryExpr_sempred((QueryExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean root_sempred(RootContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean localFieldAccess_sempred(LocalFieldAccessContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean queryExpr_sempred(QueryExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 3);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\34\u00b3\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\6\3$\n\3\r\3"+
		"\16\3%\7\3(\n\3\f\3\16\3+\13\3\3\4\3\4\3\4\3\4\3\4\3\4\5\4\63\n\4\3\5"+
		"\3\5\3\5\3\5\3\5\3\5\3\5\6\5<\n\5\r\5\16\5=\7\5@\n\5\f\5\16\5C\13\5\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6R\n\6\3\7\3\7\3"+
		"\7\6\7W\n\7\r\7\16\7X\3\7\3\7\6\7]\n\7\r\7\16\7^\3\7\3\7\3\7\3\7\6\7e"+
		"\n\7\r\7\16\7f\3\7\3\7\3\7\6\7l\n\7\r\7\16\7m\3\7\3\7\5\7r\n\7\3\b\3\b"+
		"\3\b\7\bw\n\b\f\b\16\bz\13\b\3\t\3\t\3\t\3\t\3\t\5\t\u0081\n\t\3\n\3\n"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\5\13\u008b\n\13\3\13\3\13\3\13\6\13\u0090"+
		"\n\13\r\13\16\13\u0091\7\13\u0094\n\13\f\13\16\13\u0097\13\13\3\f\6\f"+
		"\u009a\n\f\r\f\16\f\u009b\3\f\3\f\6\f\u00a0\n\f\r\f\16\f\u00a1\3\f\3\f"+
		"\6\f\u00a6\n\f\r\f\16\f\u00a7\3\f\6\f\u00ab\n\f\r\f\16\f\u00ac\5\f\u00af"+
		"\n\f\3\r\3\r\3\r\2\5\4\b\24\16\2\4\6\b\n\f\16\20\22\24\26\30\2\5\3\2\25"+
		"\32\3\2\33\34\3\2\f\r\2\u00c5\2\32\3\2\2\2\4\34\3\2\2\2\6\62\3\2\2\2\b"+
		"\64\3\2\2\2\nQ\3\2\2\2\fq\3\2\2\2\16s\3\2\2\2\20\u0080\3\2\2\2\22\u0082"+
		"\3\2\2\2\24\u008a\3\2\2\2\26\u00ae\3\2\2\2\30\u00b0\3\2\2\2\32\33\5\4"+
		"\3\2\33\3\3\2\2\2\34\35\b\3\1\2\35\36\7\23\2\2\36\37\5\6\4\2\37)\3\2\2"+
		"\2 #\f\4\2\2!\"\7\3\2\2\"$\5\6\4\2#!\3\2\2\2$%\3\2\2\2%#\3\2\2\2%&\3\2"+
		"\2\2&(\3\2\2\2\' \3\2\2\2(+\3\2\2\2)\'\3\2\2\2)*\3\2\2\2*\5\3\2\2\2+)"+
		"\3\2\2\2,-\7\16\2\2-\63\5\n\6\2./\7\16\2\2/\60\7\3\2\2\60\63\5\n\6\2\61"+
		"\63\7\16\2\2\62,\3\2\2\2\62.\3\2\2\2\62\61\3\2\2\2\63\7\3\2\2\2\64\65"+
		"\b\5\1\2\65\66\7\24\2\2\66\67\5\6\4\2\67A\3\2\2\28;\f\3\2\29:\7\3\2\2"+
		":<\5\6\4\2;9\3\2\2\2<=\3\2\2\2=;\3\2\2\2=>\3\2\2\2>@\3\2\2\2?8\3\2\2\2"+
		"@C\3\2\2\2A?\3\2\2\2AB\3\2\2\2B\t\3\2\2\2CA\3\2\2\2DR\7\4\2\2EF\7\5\2"+
		"\2FG\5\16\b\2GH\7\6\2\2HR\3\2\2\2IR\5\f\7\2JK\7\7\2\2KL\5\22\n\2LM\7\b"+
		"\2\2MR\3\2\2\2NO\7\t\2\2OP\7\20\2\2PR\7\6\2\2QD\3\2\2\2QE\3\2\2\2QI\3"+
		"\2\2\2QJ\3\2\2\2QN\3\2\2\2R\13\3\2\2\2Sr\7\n\2\2TV\7\5\2\2UW\7\20\2\2"+
		"VU\3\2\2\2WX\3\2\2\2XV\3\2\2\2XY\3\2\2\2YZ\3\2\2\2Z\\\7\22\2\2[]\7\20"+
		"\2\2\\[\3\2\2\2]^\3\2\2\2^\\\3\2\2\2^_\3\2\2\2_`\3\2\2\2`r\7\6\2\2ab\7"+
		"\5\2\2bd\7\22\2\2ce\7\20\2\2dc\3\2\2\2ef\3\2\2\2fd\3\2\2\2fg\3\2\2\2g"+
		"h\3\2\2\2hr\7\6\2\2ik\7\5\2\2jl\7\20\2\2kj\3\2\2\2lm\3\2\2\2mk\3\2\2\2"+
		"mn\3\2\2\2no\3\2\2\2op\7\22\2\2pr\7\6\2\2qS\3\2\2\2qT\3\2\2\2qa\3\2\2"+
		"\2qi\3\2\2\2r\r\3\2\2\2sx\7\20\2\2tu\7\13\2\2uw\7\20\2\2vt\3\2\2\2wz\3"+
		"\2\2\2xv\3\2\2\2xy\3\2\2\2y\17\3\2\2\2zx\3\2\2\2{\u0081\5\b\5\2|\u0081"+
		"\5\4\3\2}\u0081\5\26\f\2~\u0081\7\17\2\2\177\u0081\5\30\r\2\u0080{\3\2"+
		"\2\2\u0080|\3\2\2\2\u0080}\3\2\2\2\u0080~\3\2\2\2\u0080\177\3\2\2\2\u0081"+
		"\21\3\2\2\2\u0082\u0083\5\24\13\2\u0083\23\3\2\2\2\u0084\u0085\b\13\1"+
		"\2\u0085\u008b\5\20\t\2\u0086\u0087\5\20\t\2\u0087\u0088\t\2\2\2\u0088"+
		"\u0089\5\20\t\2\u0089\u008b\3\2\2\2\u008a\u0084\3\2\2\2\u008a\u0086\3"+
		"\2\2\2\u008b\u0095\3\2\2\2\u008c\u008f\f\5\2\2\u008d\u008e\t\3\2\2\u008e"+
		"\u0090\5\24\13\2\u008f\u008d\3\2\2\2\u0090\u0091\3\2\2\2\u0091\u008f\3"+
		"\2\2\2\u0091\u0092\3\2\2\2\u0092\u0094\3\2\2\2\u0093\u008c\3\2\2\2\u0094"+
		"\u0097\3\2\2\2\u0095\u0093\3\2\2\2\u0095\u0096\3\2\2\2\u0096\25\3\2\2"+
		"\2\u0097\u0095\3\2\2\2\u0098\u009a\7\20\2\2\u0099\u0098\3\2\2\2\u009a"+
		"\u009b\3\2\2\2\u009b\u0099\3\2\2\2\u009b\u009c\3\2\2\2\u009c\u009d\3\2"+
		"\2\2\u009d\u009f\7\3\2\2\u009e\u00a0\7\20\2\2\u009f\u009e\3\2\2\2\u00a0"+
		"\u00a1\3\2\2\2\u00a1\u009f\3\2\2\2\u00a1\u00a2\3\2\2\2\u00a2\u00af\3\2"+
		"\2\2\u00a3\u00a5\7\3\2\2\u00a4\u00a6\7\20\2\2\u00a5\u00a4\3\2\2\2\u00a6"+
		"\u00a7\3\2\2\2\u00a7\u00a5\3\2\2\2\u00a7\u00a8\3\2\2\2\u00a8\u00af\3\2"+
		"\2\2\u00a9\u00ab\7\20\2\2\u00aa\u00a9\3\2\2\2\u00ab\u00ac\3\2\2\2\u00ac"+
		"\u00aa\3\2\2\2\u00ac\u00ad\3\2\2\2\u00ad\u00af\3\2\2\2\u00ae\u0099\3\2"+
		"\2\2\u00ae\u00a3\3\2\2\2\u00ae\u00aa\3\2\2\2\u00af\27\3\2\2\2\u00b0\u00b1"+
		"\t\4\2\2\u00b1\31\3\2\2\2\27%)\62=AQX^fmqx\u0080\u008a\u0091\u0095\u009b"+
		"\u00a1\u00a7\u00ac\u00ae";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
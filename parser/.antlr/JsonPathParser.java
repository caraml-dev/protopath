// Generated from /Users/tio.pramayudi/Documents/dsp/jsonpath/parser/JsonPath.g4 by ANTLR 4.9.2
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
		IDENTIFIER=10, STRING=11, INT=12, WS=13, COLON=14, ROOT_IDENTIFIER=15, 
		LOCAL_IDENTIFIER=16, GREATER=17, GREATER_EQ=18, LESS=19, LESS_EQ=20, EQ=21, 
		NEQ=22, AND=23, OR=24;
	public static final int
		RULE_jsonpath = 0, RULE_root = 1, RULE_fieldAccess = 2, RULE_localFieldAccess = 3, 
		RULE_arrayQualifier = 4, RULE_arrayRange = 5, RULE_intSequence = 6, RULE_queryFieldValue = 7, 
		RULE_filterOperation = 8, RULE_queryExpr = 9, RULE_dbl = 10;
	private static String[] makeRuleNames() {
		return new String[] {
			"jsonpath", "root", "fieldAccess", "localFieldAccess", "arrayQualifier", 
			"arrayRange", "intSequence", "queryFieldValue", "filterOperation", "queryExpr", 
			"dbl"
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
			setState(22);
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
			setState(25);
			match(ROOT_IDENTIFIER);
			setState(26);
			fieldAccess();
			}
			_ctx.stop = _input.LT(-1);
			setState(37);
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
					setState(28);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(31); 
					_errHandler.sync(this);
					_alt = 1;
					do {
						switch (_alt) {
						case 1:
							{
							{
							setState(29);
							match(T__0);
							setState(30);
							fieldAccess();
							}
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						setState(33); 
						_errHandler.sync(this);
						_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
					} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
					}
					} 
				}
				setState(39);
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
			setState(46);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(40);
				match(IDENTIFIER);
				setState(41);
				arrayQualifier();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(42);
				match(IDENTIFIER);
				setState(43);
				match(T__0);
				setState(44);
				arrayQualifier();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(45);
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
			setState(49);
			match(LOCAL_IDENTIFIER);
			setState(50);
			fieldAccess();
			}
			_ctx.stop = _input.LT(-1);
			setState(61);
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
					setState(52);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(55); 
					_errHandler.sync(this);
					_alt = 1;
					do {
						switch (_alt) {
						case 1:
							{
							{
							setState(53);
							match(T__0);
							setState(54);
							fieldAccess();
							}
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						setState(57); 
						_errHandler.sync(this);
						_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
					} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
					}
					} 
				}
				setState(63);
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
			setState(77);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				_localctx = new GetAllContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(64);
				match(T__1);
				}
				break;
			case 2:
				_localctx = new GetByIndicesContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(65);
				match(T__2);
				setState(66);
				intSequence();
				setState(67);
				match(T__3);
				}
				break;
			case 3:
				_localctx = new GetByRangeContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(69);
				arrayRange();
				}
				break;
			case 4:
				_localctx = new GetByQueryContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(70);
				match(T__4);
				setState(71);
				filterOperation();
				setState(72);
				match(T__5);
				}
				break;
			case 5:
				_localctx = new GetByIndexBackwardContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(74);
				match(T__6);
				setState(75);
				match(INT);
				setState(76);
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
			setState(109);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				_localctx = new AllRangeContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(79);
				match(T__7);
				}
				break;
			case 2:
				_localctx = new GivenStartEndIdxRangeContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(80);
				match(T__2);
				setState(82); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(81);
					match(INT);
					}
					}
					setState(84); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==INT );
				setState(86);
				match(COLON);
				setState(88); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(87);
					match(INT);
					}
					}
					setState(90); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==INT );
				setState(92);
				match(T__3);
				}
				break;
			case 3:
				_localctx = new GivenOnlyEndIdxRangeContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(93);
				match(T__2);
				setState(94);
				match(COLON);
				setState(96); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(95);
					match(INT);
					}
					}
					setState(98); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==INT );
				setState(100);
				match(T__3);
				}
				break;
			case 4:
				_localctx = new GivenOnlyStartIdxRangeContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(101);
				match(T__2);
				setState(103); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(102);
					match(INT);
					}
					}
					setState(105); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==INT );
				setState(107);
				match(COLON);
				setState(108);
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
			setState(111);
			match(INT);
			setState(116);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__8) {
				{
				{
				setState(112);
				match(T__8);
				setState(113);
				match(INT);
				}
				}
				setState(118);
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

	public final QueryFieldValueContext queryFieldValue() throws RecognitionException {
		QueryFieldValueContext _localctx = new QueryFieldValueContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_queryFieldValue);
		try {
			setState(123);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LOCAL_IDENTIFIER:
				_localctx = new QueryFieldValueLocalAccessContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(119);
				localFieldAccess(0);
				}
				break;
			case ROOT_IDENTIFIER:
				_localctx = new QueryFieldValueRootAccessContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(120);
				root(0);
				}
				break;
			case T__0:
			case INT:
				_localctx = new QueryFieldDoubleValueContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(121);
				dbl();
				}
				break;
			case STRING:
				_localctx = new QueryFieldStringValueContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(122);
				match(STRING);
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
			setState(125);
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
			setState(133);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				{
				_localctx = new QueryExistenceContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(128);
				queryFieldValue();
				}
				break;
			case 2:
				{
				_localctx = new QueryWithComparatorContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(129);
				queryFieldValue();
				setState(130);
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
				setState(131);
				queryFieldValue();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(144);
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
					setState(135);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(138); 
					_errHandler.sync(this);
					_alt = 1;
					do {
						switch (_alt) {
						case 1:
							{
							{
							setState(136);
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
							setState(137);
							queryExpr(0);
							}
							}
							break;
						default:
							throw new NoViableAltException(this);
						}
						setState(140); 
						_errHandler.sync(this);
						_alt = getInterpreter().adaptivePredict(_input,14,_ctx);
					} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
					}
					} 
				}
				setState(146);
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
			setState(169);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(148); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(147);
					match(INT);
					}
					}
					setState(150); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==INT );
				setState(152);
				match(T__0);
				setState(154); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(153);
						match(INT);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(156); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(158);
				match(T__0);
				setState(160); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(159);
						match(INT);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(162); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(165); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(164);
						match(INT);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(167); 
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\32\u00ae\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\6\3\"\n\3\r\3\16\3#"+
		"\7\3&\n\3\f\3\16\3)\13\3\3\4\3\4\3\4\3\4\3\4\3\4\5\4\61\n\4\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\6\5:\n\5\r\5\16\5;\7\5>\n\5\f\5\16\5A\13\5\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6P\n\6\3\7\3\7\3\7\6\7"+
		"U\n\7\r\7\16\7V\3\7\3\7\6\7[\n\7\r\7\16\7\\\3\7\3\7\3\7\3\7\6\7c\n\7\r"+
		"\7\16\7d\3\7\3\7\3\7\6\7j\n\7\r\7\16\7k\3\7\3\7\5\7p\n\7\3\b\3\b\3\b\7"+
		"\bu\n\b\f\b\16\bx\13\b\3\t\3\t\3\t\3\t\5\t~\n\t\3\n\3\n\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\5\13\u0088\n\13\3\13\3\13\3\13\6\13\u008d\n\13\r\13\16"+
		"\13\u008e\7\13\u0091\n\13\f\13\16\13\u0094\13\13\3\f\6\f\u0097\n\f\r\f"+
		"\16\f\u0098\3\f\3\f\6\f\u009d\n\f\r\f\16\f\u009e\3\f\3\f\6\f\u00a3\n\f"+
		"\r\f\16\f\u00a4\3\f\6\f\u00a8\n\f\r\f\16\f\u00a9\5\f\u00ac\n\f\3\f\2\5"+
		"\4\b\24\r\2\4\6\b\n\f\16\20\22\24\26\2\4\3\2\23\30\3\2\31\32\2\u00c0\2"+
		"\30\3\2\2\2\4\32\3\2\2\2\6\60\3\2\2\2\b\62\3\2\2\2\nO\3\2\2\2\fo\3\2\2"+
		"\2\16q\3\2\2\2\20}\3\2\2\2\22\177\3\2\2\2\24\u0087\3\2\2\2\26\u00ab\3"+
		"\2\2\2\30\31\5\4\3\2\31\3\3\2\2\2\32\33\b\3\1\2\33\34\7\21\2\2\34\35\5"+
		"\6\4\2\35\'\3\2\2\2\36!\f\4\2\2\37 \7\3\2\2 \"\5\6\4\2!\37\3\2\2\2\"#"+
		"\3\2\2\2#!\3\2\2\2#$\3\2\2\2$&\3\2\2\2%\36\3\2\2\2&)\3\2\2\2\'%\3\2\2"+
		"\2\'(\3\2\2\2(\5\3\2\2\2)\'\3\2\2\2*+\7\f\2\2+\61\5\n\6\2,-\7\f\2\2-."+
		"\7\3\2\2.\61\5\n\6\2/\61\7\f\2\2\60*\3\2\2\2\60,\3\2\2\2\60/\3\2\2\2\61"+
		"\7\3\2\2\2\62\63\b\5\1\2\63\64\7\22\2\2\64\65\5\6\4\2\65?\3\2\2\2\669"+
		"\f\3\2\2\678\7\3\2\28:\5\6\4\29\67\3\2\2\2:;\3\2\2\2;9\3\2\2\2;<\3\2\2"+
		"\2<>\3\2\2\2=\66\3\2\2\2>A\3\2\2\2?=\3\2\2\2?@\3\2\2\2@\t\3\2\2\2A?\3"+
		"\2\2\2BP\7\4\2\2CD\7\5\2\2DE\5\16\b\2EF\7\6\2\2FP\3\2\2\2GP\5\f\7\2HI"+
		"\7\7\2\2IJ\5\22\n\2JK\7\b\2\2KP\3\2\2\2LM\7\t\2\2MN\7\16\2\2NP\7\6\2\2"+
		"OB\3\2\2\2OC\3\2\2\2OG\3\2\2\2OH\3\2\2\2OL\3\2\2\2P\13\3\2\2\2Qp\7\n\2"+
		"\2RT\7\5\2\2SU\7\16\2\2TS\3\2\2\2UV\3\2\2\2VT\3\2\2\2VW\3\2\2\2WX\3\2"+
		"\2\2XZ\7\20\2\2Y[\7\16\2\2ZY\3\2\2\2[\\\3\2\2\2\\Z\3\2\2\2\\]\3\2\2\2"+
		"]^\3\2\2\2^p\7\6\2\2_`\7\5\2\2`b\7\20\2\2ac\7\16\2\2ba\3\2\2\2cd\3\2\2"+
		"\2db\3\2\2\2de\3\2\2\2ef\3\2\2\2fp\7\6\2\2gi\7\5\2\2hj\7\16\2\2ih\3\2"+
		"\2\2jk\3\2\2\2ki\3\2\2\2kl\3\2\2\2lm\3\2\2\2mn\7\20\2\2np\7\6\2\2oQ\3"+
		"\2\2\2oR\3\2\2\2o_\3\2\2\2og\3\2\2\2p\r\3\2\2\2qv\7\16\2\2rs\7\13\2\2"+
		"su\7\16\2\2tr\3\2\2\2ux\3\2\2\2vt\3\2\2\2vw\3\2\2\2w\17\3\2\2\2xv\3\2"+
		"\2\2y~\5\b\5\2z~\5\4\3\2{~\5\26\f\2|~\7\r\2\2}y\3\2\2\2}z\3\2\2\2}{\3"+
		"\2\2\2}|\3\2\2\2~\21\3\2\2\2\177\u0080\5\24\13\2\u0080\23\3\2\2\2\u0081"+
		"\u0082\b\13\1\2\u0082\u0088\5\20\t\2\u0083\u0084\5\20\t\2\u0084\u0085"+
		"\t\2\2\2\u0085\u0086\5\20\t\2\u0086\u0088\3\2\2\2\u0087\u0081\3\2\2\2"+
		"\u0087\u0083\3\2\2\2\u0088\u0092\3\2\2\2\u0089\u008c\f\5\2\2\u008a\u008b"+
		"\t\3\2\2\u008b\u008d\5\24\13\2\u008c\u008a\3\2\2\2\u008d\u008e\3\2\2\2"+
		"\u008e\u008c\3\2\2\2\u008e\u008f\3\2\2\2\u008f\u0091\3\2\2\2\u0090\u0089"+
		"\3\2\2\2\u0091\u0094\3\2\2\2\u0092\u0090\3\2\2\2\u0092\u0093\3\2\2\2\u0093"+
		"\25\3\2\2\2\u0094\u0092\3\2\2\2\u0095\u0097\7\16\2\2\u0096\u0095\3\2\2"+
		"\2\u0097\u0098\3\2\2\2\u0098\u0096\3\2\2\2\u0098\u0099\3\2\2\2\u0099\u009a"+
		"\3\2\2\2\u009a\u009c\7\3\2\2\u009b\u009d\7\16\2\2\u009c\u009b\3\2\2\2"+
		"\u009d\u009e\3\2\2\2\u009e\u009c\3\2\2\2\u009e\u009f\3\2\2\2\u009f\u00ac"+
		"\3\2\2\2\u00a0\u00a2\7\3\2\2\u00a1\u00a3\7\16\2\2\u00a2\u00a1\3\2\2\2"+
		"\u00a3\u00a4\3\2\2\2\u00a4\u00a2\3\2\2\2\u00a4\u00a5\3\2\2\2\u00a5\u00ac"+
		"\3\2\2\2\u00a6\u00a8\7\16\2\2\u00a7\u00a6\3\2\2\2\u00a8\u00a9\3\2\2\2"+
		"\u00a9\u00a7\3\2\2\2\u00a9\u00aa\3\2\2\2\u00aa\u00ac\3\2\2\2\u00ab\u0096"+
		"\3\2\2\2\u00ab\u00a0\3\2\2\2\u00ab\u00a7\3\2\2\2\u00ac\27\3\2\2\2\27#"+
		"\'\60;?OV\\dkov}\u0087\u008e\u0092\u0098\u009e\u00a4\u00a9\u00ab";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
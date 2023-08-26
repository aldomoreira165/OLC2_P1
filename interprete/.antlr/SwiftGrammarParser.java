// Generated from c:\Users\aldov\Desktop\Cursos Segundo Semestre 2023\Laboratorio Organización de Lenguajes y Compiladores 2 - copia\Proyectos\Proyecto 1\interprete\SwiftGrammar.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SwiftGrammarParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		INT=1, FLOAT=2, BOOL=3, CHARACTER=4, PSTRING=5, NIL=6, TRU=7, FAL=8, PRINT=9, 
		IF=10, ELSE=11, WHILE=12, SWITCH=13, CASE=14, DEFAULT=15, VAR=16, BREAK=17, 
		RETURN=18, FUNC=19, COUNT=20, ISEMPTY=21, APPEND=22, NUMBER=23, STRING=24, 
		ID=25, DIF=26, IG_IG=27, NOT=28, OR=29, AND=30, IG=31, MAY_IG=32, MEN_IG=33, 
		MAYOR=34, MENOR=35, MUL=36, DIV=37, ADD=38, SUB=39, MOD=40, PARIZQ=41, 
		PARDER=42, LLAVEIZQ=43, LLAVEDER=44, CORCHIZQ=45, CORCHDER=46, DOSPUNTOS=47, 
		COMA=48, PTCOMA=49, INTERROGACION=50, PUNTO=51, WHITESPACE=52, COMMENT=53, 
		LINE_COMMENT=54;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_stmt = 2, RULE_declvectorstmt = 3, RULE_defvectorstmt = 4, 
		RULE_listaexpresiones = 5, RULE_accesovectorstmt = 6, RULE_appendvectorstmt = 7, 
		RULE_countvectorstmt = 8, RULE_isemptyvectorstmt = 9, RULE_returnstmt = 10, 
		RULE_printstmt = 11, RULE_intstmt = 12, RULE_floatstmt = 13, RULE_stringstmt = 14, 
		RULE_funcdclstmt = 15, RULE_accfuncstm = 16, RULE_parametros = 17, RULE_parametroscall = 18, 
		RULE_breakstmt = 19, RULE_ifstmt = 20, RULE_elseifstmt = 21, RULE_switchstmt = 22, 
		RULE_caseStmt = 23, RULE_defaultCase = 24, RULE_typedDeclstmt = 25, RULE_untypedDeclstmt = 26, 
		RULE_optionalTypedDeclstmt = 27, RULE_asignstmt = 28, RULE_whilestmt = 29, 
		RULE_opasignstmt = 30, RULE_expr = 31, RULE_tipo = 32;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "stmt", "declvectorstmt", "defvectorstmt", "listaexpresiones", 
			"accesovectorstmt", "appendvectorstmt", "countvectorstmt", "isemptyvectorstmt", 
			"returnstmt", "printstmt", "intstmt", "floatstmt", "stringstmt", "funcdclstmt", 
			"accfuncstm", "parametros", "parametroscall", "breakstmt", "ifstmt", 
			"elseifstmt", "switchstmt", "caseStmt", "defaultCase", "typedDeclstmt", 
			"untypedDeclstmt", "optionalTypedDeclstmt", "asignstmt", "whilestmt", 
			"opasignstmt", "expr", "tipo"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'int'", "'float'", "'bool'", "'character'", "'String'", "'nil'", 
			"'true'", "'false'", "'print'", "'if'", "'else'", "'while'", "'switch'", 
			"'case'", "'default'", "'var'", "'break'", "'return'", "'func'", "'count'", 
			"'IsEmpty'", "'append'", null, null, null, "'!='", "'=='", "'!'", "'||'", 
			"'&&'", "'='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", 
			"'%'", "'('", "')'", "'{'", "'}'", "'['", "']'", "':'", "','", "';'", 
			"'?'", "'.'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "NIL", "TRU", "FAL", 
			"PRINT", "IF", "ELSE", "WHILE", "SWITCH", "CASE", "DEFAULT", "VAR", "BREAK", 
			"RETURN", "FUNC", "COUNT", "ISEMPTY", "APPEND", "NUMBER", "STRING", "ID", 
			"DIF", "IG_IG", "NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", 
			"MENOR", "MUL", "DIV", "ADD", "SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", 
			"LLAVEDER", "CORCHIZQ", "CORCHDER", "DOSPUNTOS", "COMA", "PTCOMA", "INTERROGACION", 
			"PUNTO", "WHITESPACE", "COMMENT", "LINE_COMMENT"
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
	public String getGrammarFileName() { return "SwiftGrammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public SwiftGrammarParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class SContext extends ParserRuleContext {
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode EOF() { return getToken(SwiftGrammarParser.EOF, 0); }
		public SContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_s; }
	}

	public final SContext s() throws RecognitionException {
		SContext _localctx = new SContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_s);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(66);
			block();
			setState(67);
			match(EOF);
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

	public static class BlockContext extends ParserRuleContext {
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(72);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PRINT) | (1L << IF) | (1L << WHILE) | (1L << SWITCH) | (1L << VAR) | (1L << BREAK) | (1L << RETURN) | (1L << FUNC) | (1L << ID))) != 0)) {
				{
				{
				setState(69);
				stmt();
				}
				}
				setState(74);
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

	public static class StmtContext extends ParserRuleContext {
		public PrintstmtContext printstmt() {
			return getRuleContext(PrintstmtContext.class,0);
		}
		public TypedDeclstmtContext typedDeclstmt() {
			return getRuleContext(TypedDeclstmtContext.class,0);
		}
		public UntypedDeclstmtContext untypedDeclstmt() {
			return getRuleContext(UntypedDeclstmtContext.class,0);
		}
		public OptionalTypedDeclstmtContext optionalTypedDeclstmt() {
			return getRuleContext(OptionalTypedDeclstmtContext.class,0);
		}
		public AsignstmtContext asignstmt() {
			return getRuleContext(AsignstmtContext.class,0);
		}
		public IfstmtContext ifstmt() {
			return getRuleContext(IfstmtContext.class,0);
		}
		public SwitchstmtContext switchstmt() {
			return getRuleContext(SwitchstmtContext.class,0);
		}
		public WhilestmtContext whilestmt() {
			return getRuleContext(WhilestmtContext.class,0);
		}
		public OpasignstmtContext opasignstmt() {
			return getRuleContext(OpasignstmtContext.class,0);
		}
		public BreakstmtContext breakstmt() {
			return getRuleContext(BreakstmtContext.class,0);
		}
		public FuncdclstmtContext funcdclstmt() {
			return getRuleContext(FuncdclstmtContext.class,0);
		}
		public AccfuncstmContext accfuncstm() {
			return getRuleContext(AccfuncstmContext.class,0);
		}
		public ReturnstmtContext returnstmt() {
			return getRuleContext(ReturnstmtContext.class,0);
		}
		public DeclvectorstmtContext declvectorstmt() {
			return getRuleContext(DeclvectorstmtContext.class,0);
		}
		public AccesovectorstmtContext accesovectorstmt() {
			return getRuleContext(AccesovectorstmtContext.class,0);
		}
		public AppendvectorstmtContext appendvectorstmt() {
			return getRuleContext(AppendvectorstmtContext.class,0);
		}
		public StmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt; }
	}

	public final StmtContext stmt() throws RecognitionException {
		StmtContext _localctx = new StmtContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_stmt);
		try {
			setState(91);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(75);
				printstmt();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(76);
				typedDeclstmt();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(77);
				untypedDeclstmt();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(78);
				optionalTypedDeclstmt();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(79);
				asignstmt();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(80);
				ifstmt();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(81);
				switchstmt();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(82);
				whilestmt();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(83);
				opasignstmt();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(84);
				breakstmt();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(85);
				funcdclstmt();
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(86);
				accfuncstm();
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(87);
				returnstmt();
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(88);
				declvectorstmt();
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(89);
				accesovectorstmt();
				}
				break;
			case 16:
				enterOuterAlt(_localctx, 16);
				{
				setState(90);
				appendvectorstmt();
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

	public static class DeclvectorstmtContext extends ParserRuleContext {
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode DOSPUNTOS() { return getToken(SwiftGrammarParser.DOSPUNTOS, 0); }
		public TerminalNode CORCHIZQ() { return getToken(SwiftGrammarParser.CORCHIZQ, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode CORCHDER() { return getToken(SwiftGrammarParser.CORCHDER, 0); }
		public DefvectorstmtContext defvectorstmt() {
			return getRuleContext(DefvectorstmtContext.class,0);
		}
		public DeclvectorstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declvectorstmt; }
	}

	public final DeclvectorstmtContext declvectorstmt() throws RecognitionException {
		DeclvectorstmtContext _localctx = new DeclvectorstmtContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_declvectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(93);
			match(VAR);
			setState(94);
			match(ID);
			setState(95);
			match(DOSPUNTOS);
			setState(96);
			match(CORCHIZQ);
			setState(97);
			tipo();
			setState(98);
			match(CORCHDER);
			setState(99);
			defvectorstmt();
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

	public static class DefvectorstmtContext extends ParserRuleContext {
		public DefvectorstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_defvectorstmt; }
	 
		public DefvectorstmtContext() { }
		public void copyFrom(DefvectorstmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class DefVectorContext extends DefvectorstmtContext {
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public TerminalNode CORCHIZQ() { return getToken(SwiftGrammarParser.CORCHIZQ, 0); }
		public TerminalNode CORCHDER() { return getToken(SwiftGrammarParser.CORCHDER, 0); }
		public ListaexpresionesContext listaexpresiones() {
			return getRuleContext(ListaexpresionesContext.class,0);
		}
		public DefVectorContext(DefvectorstmtContext ctx) { copyFrom(ctx); }
	}
	public static class DefVectorIDContext extends DefvectorstmtContext {
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public DefVectorIDContext(DefvectorstmtContext ctx) { copyFrom(ctx); }
	}

	public final DefvectorstmtContext defvectorstmt() throws RecognitionException {
		DefvectorstmtContext _localctx = new DefvectorstmtContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_defvectorstmt);
		int _la;
		try {
			setState(109);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				_localctx = new DefVectorContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(101);
				match(IG);
				setState(102);
				match(CORCHIZQ);
				setState(104);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << INT) | (1L << FLOAT) | (1L << PSTRING) | (1L << NIL) | (1L << TRU) | (1L << FAL) | (1L << NUMBER) | (1L << STRING) | (1L << ID) | (1L << SUB) | (1L << PARIZQ))) != 0)) {
					{
					setState(103);
					listaexpresiones();
					}
				}

				setState(106);
				match(CORCHDER);
				}
				break;
			case 2:
				_localctx = new DefVectorIDContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(107);
				match(IG);
				setState(108);
				match(ID);
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

	public static class ListaexpresionesContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> COMA() { return getTokens(SwiftGrammarParser.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(SwiftGrammarParser.COMA, i);
		}
		public ListaexpresionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listaexpresiones; }
	}

	public final ListaexpresionesContext listaexpresiones() throws RecognitionException {
		ListaexpresionesContext _localctx = new ListaexpresionesContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_listaexpresiones);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(111);
			expr(0);
			setState(116);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(112);
				match(COMA);
				setState(113);
				expr(0);
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

	public static class AccesovectorstmtContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode CORCHIZQ() { return getToken(SwiftGrammarParser.CORCHIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode CORCHDER() { return getToken(SwiftGrammarParser.CORCHDER, 0); }
		public AccesovectorstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accesovectorstmt; }
	}

	public final AccesovectorstmtContext accesovectorstmt() throws RecognitionException {
		AccesovectorstmtContext _localctx = new AccesovectorstmtContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_accesovectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(119);
			match(ID);
			setState(120);
			match(CORCHIZQ);
			setState(121);
			expr(0);
			setState(122);
			match(CORCHDER);
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

	public static class AppendvectorstmtContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
		public TerminalNode APPEND() { return getToken(SwiftGrammarParser.APPEND, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public AppendvectorstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_appendvectorstmt; }
	}

	public final AppendvectorstmtContext appendvectorstmt() throws RecognitionException {
		AppendvectorstmtContext _localctx = new AppendvectorstmtContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_appendvectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(124);
			match(ID);
			setState(125);
			match(PUNTO);
			setState(126);
			match(APPEND);
			setState(127);
			match(PARIZQ);
			setState(128);
			expr(0);
			setState(129);
			match(PARDER);
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

	public static class CountvectorstmtContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
		public TerminalNode COUNT() { return getToken(SwiftGrammarParser.COUNT, 0); }
		public CountvectorstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_countvectorstmt; }
	}

	public final CountvectorstmtContext countvectorstmt() throws RecognitionException {
		CountvectorstmtContext _localctx = new CountvectorstmtContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_countvectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(131);
			match(ID);
			setState(132);
			match(PUNTO);
			setState(133);
			match(COUNT);
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

	public static class IsemptyvectorstmtContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
		public TerminalNode ISEMPTY() { return getToken(SwiftGrammarParser.ISEMPTY, 0); }
		public IsemptyvectorstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_isemptyvectorstmt; }
	}

	public final IsemptyvectorstmtContext isemptyvectorstmt() throws RecognitionException {
		IsemptyvectorstmtContext _localctx = new IsemptyvectorstmtContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_isemptyvectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(135);
			match(ID);
			setState(136);
			match(PUNTO);
			setState(137);
			match(ISEMPTY);
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

	public static class ReturnstmtContext extends ParserRuleContext {
		public TerminalNode RETURN() { return getToken(SwiftGrammarParser.RETURN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(SwiftGrammarParser.PTCOMA, 0); }
		public ReturnstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returnstmt; }
	}

	public final ReturnstmtContext returnstmt() throws RecognitionException {
		ReturnstmtContext _localctx = new ReturnstmtContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_returnstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(139);
			match(RETURN);
			setState(140);
			expr(0);
			setState(141);
			match(PTCOMA);
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

	public static class PrintstmtContext extends ParserRuleContext {
		public TerminalNode PRINT() { return getToken(SwiftGrammarParser.PRINT, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public PrintstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printstmt; }
	}

	public final PrintstmtContext printstmt() throws RecognitionException {
		PrintstmtContext _localctx = new PrintstmtContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_printstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(143);
			match(PRINT);
			setState(144);
			match(PARIZQ);
			setState(145);
			expr(0);
			setState(146);
			match(PARDER);
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

	public static class IntstmtContext extends ParserRuleContext {
		public TerminalNode INT() { return getToken(SwiftGrammarParser.INT, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public IntstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_intstmt; }
	}

	public final IntstmtContext intstmt() throws RecognitionException {
		IntstmtContext _localctx = new IntstmtContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_intstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(148);
			match(INT);
			setState(149);
			match(PARIZQ);
			setState(150);
			expr(0);
			setState(151);
			match(PARDER);
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

	public static class FloatstmtContext extends ParserRuleContext {
		public TerminalNode FLOAT() { return getToken(SwiftGrammarParser.FLOAT, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public FloatstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_floatstmt; }
	}

	public final FloatstmtContext floatstmt() throws RecognitionException {
		FloatstmtContext _localctx = new FloatstmtContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_floatstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(153);
			match(FLOAT);
			setState(154);
			match(PARIZQ);
			setState(155);
			expr(0);
			setState(156);
			match(PARDER);
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

	public static class StringstmtContext extends ParserRuleContext {
		public TerminalNode PSTRING() { return getToken(SwiftGrammarParser.PSTRING, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public StringstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stringstmt; }
	}

	public final StringstmtContext stringstmt() throws RecognitionException {
		StringstmtContext _localctx = new StringstmtContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_stringstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(158);
			match(PSTRING);
			setState(159);
			match(PARIZQ);
			setState(160);
			expr(0);
			setState(161);
			match(PARDER);
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

	public static class FuncdclstmtContext extends ParserRuleContext {
		public FuncdclstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcdclstmt; }
	 
		public FuncdclstmtContext() { }
		public void copyFrom(FuncdclstmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class FuncionNormalContext extends FuncdclstmtContext {
		public TerminalNode FUNC() { return getToken(SwiftGrammarParser.FUNC, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public ParametrosContext parametros() {
			return getRuleContext(ParametrosContext.class,0);
		}
		public FuncionNormalContext(FuncdclstmtContext ctx) { copyFrom(ctx); }
	}
	public static class FuncionRetornoContext extends FuncdclstmtContext {
		public TerminalNode FUNC() { return getToken(SwiftGrammarParser.FUNC, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public TerminalNode SUB() { return getToken(SwiftGrammarParser.SUB, 0); }
		public TerminalNode MAYOR() { return getToken(SwiftGrammarParser.MAYOR, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public ParametrosContext parametros() {
			return getRuleContext(ParametrosContext.class,0);
		}
		public FuncionRetornoContext(FuncdclstmtContext ctx) { copyFrom(ctx); }
	}

	public final FuncdclstmtContext funcdclstmt() throws RecognitionException {
		FuncdclstmtContext _localctx = new FuncdclstmtContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_funcdclstmt);
		int _la;
		try {
			setState(188);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				_localctx = new FuncionNormalContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(163);
				match(FUNC);
				setState(164);
				match(ID);
				setState(165);
				match(PARIZQ);
				setState(167);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ID) {
					{
					setState(166);
					parametros();
					}
				}

				setState(169);
				match(PARDER);
				setState(170);
				match(LLAVEIZQ);
				setState(171);
				block();
				setState(172);
				match(LLAVEDER);
				}
				break;
			case 2:
				_localctx = new FuncionRetornoContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(174);
				match(FUNC);
				setState(175);
				match(ID);
				setState(176);
				match(PARIZQ);
				setState(178);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ID) {
					{
					setState(177);
					parametros();
					}
				}

				setState(180);
				match(PARDER);
				setState(181);
				match(SUB);
				setState(182);
				match(MAYOR);
				setState(183);
				tipo();
				setState(184);
				match(LLAVEIZQ);
				setState(185);
				block();
				setState(186);
				match(LLAVEDER);
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

	public static class AccfuncstmContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public ParametroscallContext parametroscall() {
			return getRuleContext(ParametroscallContext.class,0);
		}
		public AccfuncstmContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accfuncstm; }
	}

	public final AccfuncstmContext accfuncstm() throws RecognitionException {
		AccfuncstmContext _localctx = new AccfuncstmContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_accfuncstm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(190);
			match(ID);
			setState(191);
			match(PARIZQ);
			setState(193);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(192);
				parametroscall();
				}
			}

			setState(195);
			match(PARDER);
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

	public static class ParametrosContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(SwiftGrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(SwiftGrammarParser.ID, i);
		}
		public List<TerminalNode> DOSPUNTOS() { return getTokens(SwiftGrammarParser.DOSPUNTOS); }
		public TerminalNode DOSPUNTOS(int i) {
			return getToken(SwiftGrammarParser.DOSPUNTOS, i);
		}
		public List<TipoContext> tipo() {
			return getRuleContexts(TipoContext.class);
		}
		public TipoContext tipo(int i) {
			return getRuleContext(TipoContext.class,i);
		}
		public List<TerminalNode> COMA() { return getTokens(SwiftGrammarParser.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(SwiftGrammarParser.COMA, i);
		}
		public ParametrosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametros; }
	}

	public final ParametrosContext parametros() throws RecognitionException {
		ParametrosContext _localctx = new ParametrosContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_parametros);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(197);
			match(ID);
			setState(198);
			match(ID);
			setState(199);
			match(DOSPUNTOS);
			setState(200);
			tipo();
			setState(208);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(201);
				match(COMA);
				setState(202);
				match(ID);
				setState(203);
				match(ID);
				setState(204);
				match(DOSPUNTOS);
				setState(205);
				tipo();
				}
				}
				setState(210);
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

	public static class ParametroscallContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(SwiftGrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(SwiftGrammarParser.ID, i);
		}
		public List<TerminalNode> DOSPUNTOS() { return getTokens(SwiftGrammarParser.DOSPUNTOS); }
		public TerminalNode DOSPUNTOS(int i) {
			return getToken(SwiftGrammarParser.DOSPUNTOS, i);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> COMA() { return getTokens(SwiftGrammarParser.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(SwiftGrammarParser.COMA, i);
		}
		public ParametroscallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametroscall; }
	}

	public final ParametroscallContext parametroscall() throws RecognitionException {
		ParametroscallContext _localctx = new ParametroscallContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_parametroscall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(211);
			match(ID);
			setState(212);
			match(DOSPUNTOS);
			setState(213);
			expr(0);
			setState(220);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(214);
				match(COMA);
				setState(215);
				match(ID);
				setState(216);
				match(DOSPUNTOS);
				setState(217);
				expr(0);
				}
				}
				setState(222);
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

	public static class BreakstmtContext extends ParserRuleContext {
		public TerminalNode BREAK() { return getToken(SwiftGrammarParser.BREAK, 0); }
		public BreakstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_breakstmt; }
	}

	public final BreakstmtContext breakstmt() throws RecognitionException {
		BreakstmtContext _localctx = new BreakstmtContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_breakstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(223);
			match(BREAK);
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

	public static class IfstmtContext extends ParserRuleContext {
		public TerminalNode IF() { return getToken(SwiftGrammarParser.IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public List<TerminalNode> LLAVEIZQ() { return getTokens(SwiftGrammarParser.LLAVEIZQ); }
		public TerminalNode LLAVEIZQ(int i) {
			return getToken(SwiftGrammarParser.LLAVEIZQ, i);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public List<TerminalNode> LLAVEDER() { return getTokens(SwiftGrammarParser.LLAVEDER); }
		public TerminalNode LLAVEDER(int i) {
			return getToken(SwiftGrammarParser.LLAVEDER, i);
		}
		public List<ElseifstmtContext> elseifstmt() {
			return getRuleContexts(ElseifstmtContext.class);
		}
		public ElseifstmtContext elseifstmt(int i) {
			return getRuleContext(ElseifstmtContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(SwiftGrammarParser.ELSE, 0); }
		public IfstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifstmt; }
	}

	public final IfstmtContext ifstmt() throws RecognitionException {
		IfstmtContext _localctx = new IfstmtContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_ifstmt);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(225);
			match(IF);
			setState(226);
			expr(0);
			setState(227);
			match(LLAVEIZQ);
			setState(228);
			block();
			setState(229);
			match(LLAVEDER);
			setState(233);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(230);
					elseifstmt();
					}
					} 
				}
				setState(235);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			}
			setState(241);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(236);
				match(ELSE);
				setState(237);
				match(LLAVEIZQ);
				setState(238);
				block();
				setState(239);
				match(LLAVEDER);
				}
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

	public static class ElseifstmtContext extends ParserRuleContext {
		public TerminalNode ELSE() { return getToken(SwiftGrammarParser.ELSE, 0); }
		public TerminalNode IF() { return getToken(SwiftGrammarParser.IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public ElseifstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elseifstmt; }
	}

	public final ElseifstmtContext elseifstmt() throws RecognitionException {
		ElseifstmtContext _localctx = new ElseifstmtContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_elseifstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(243);
			match(ELSE);
			setState(244);
			match(IF);
			setState(245);
			expr(0);
			setState(246);
			match(LLAVEIZQ);
			setState(247);
			block();
			setState(248);
			match(LLAVEDER);
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

	public static class SwitchstmtContext extends ParserRuleContext {
		public TerminalNode SWITCH() { return getToken(SwiftGrammarParser.SWITCH, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public List<CaseStmtContext> caseStmt() {
			return getRuleContexts(CaseStmtContext.class);
		}
		public CaseStmtContext caseStmt(int i) {
			return getRuleContext(CaseStmtContext.class,i);
		}
		public DefaultCaseContext defaultCase() {
			return getRuleContext(DefaultCaseContext.class,0);
		}
		public SwitchstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switchstmt; }
	}

	public final SwitchstmtContext switchstmt() throws RecognitionException {
		SwitchstmtContext _localctx = new SwitchstmtContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_switchstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(250);
			match(SWITCH);
			setState(251);
			expr(0);
			setState(252);
			match(LLAVEIZQ);
			setState(254); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(253);
				caseStmt();
				}
				}
				setState(256); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==CASE );
			setState(259);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(258);
				defaultCase();
				}
			}

			setState(261);
			match(LLAVEDER);
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

	public static class CaseStmtContext extends ParserRuleContext {
		public TerminalNode CASE() { return getToken(SwiftGrammarParser.CASE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode DOSPUNTOS() { return getToken(SwiftGrammarParser.DOSPUNTOS, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public CaseStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_caseStmt; }
	}

	public final CaseStmtContext caseStmt() throws RecognitionException {
		CaseStmtContext _localctx = new CaseStmtContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_caseStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(263);
			match(CASE);
			setState(264);
			expr(0);
			setState(265);
			match(DOSPUNTOS);
			setState(266);
			block();
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

	public static class DefaultCaseContext extends ParserRuleContext {
		public TerminalNode DEFAULT() { return getToken(SwiftGrammarParser.DEFAULT, 0); }
		public TerminalNode DOSPUNTOS() { return getToken(SwiftGrammarParser.DOSPUNTOS, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public DefaultCaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_defaultCase; }
	}

	public final DefaultCaseContext defaultCase() throws RecognitionException {
		DefaultCaseContext _localctx = new DefaultCaseContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_defaultCase);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(268);
			match(DEFAULT);
			setState(269);
			match(DOSPUNTOS);
			setState(270);
			block();
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

	public static class TypedDeclstmtContext extends ParserRuleContext {
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode DOSPUNTOS() { return getToken(SwiftGrammarParser.DOSPUNTOS, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TypedDeclstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typedDeclstmt; }
	}

	public final TypedDeclstmtContext typedDeclstmt() throws RecognitionException {
		TypedDeclstmtContext _localctx = new TypedDeclstmtContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_typedDeclstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(272);
			match(VAR);
			setState(273);
			match(ID);
			setState(274);
			match(DOSPUNTOS);
			setState(275);
			tipo();
			setState(276);
			match(IG);
			setState(277);
			expr(0);
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

	public static class UntypedDeclstmtContext extends ParserRuleContext {
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public UntypedDeclstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_untypedDeclstmt; }
	}

	public final UntypedDeclstmtContext untypedDeclstmt() throws RecognitionException {
		UntypedDeclstmtContext _localctx = new UntypedDeclstmtContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_untypedDeclstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(279);
			match(VAR);
			setState(280);
			match(ID);
			setState(281);
			match(IG);
			setState(282);
			expr(0);
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

	public static class OptionalTypedDeclstmtContext extends ParserRuleContext {
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode DOSPUNTOS() { return getToken(SwiftGrammarParser.DOSPUNTOS, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode INTERROGACION() { return getToken(SwiftGrammarParser.INTERROGACION, 0); }
		public OptionalTypedDeclstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_optionalTypedDeclstmt; }
	}

	public final OptionalTypedDeclstmtContext optionalTypedDeclstmt() throws RecognitionException {
		OptionalTypedDeclstmtContext _localctx = new OptionalTypedDeclstmtContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_optionalTypedDeclstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(284);
			match(VAR);
			setState(285);
			match(ID);
			setState(286);
			match(DOSPUNTOS);
			setState(287);
			tipo();
			setState(288);
			match(INTERROGACION);
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

	public static class AsignstmtContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public AsignstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignstmt; }
	}

	public final AsignstmtContext asignstmt() throws RecognitionException {
		AsignstmtContext _localctx = new AsignstmtContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_asignstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(290);
			match(ID);
			setState(291);
			match(IG);
			setState(292);
			expr(0);
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

	public static class WhilestmtContext extends ParserRuleContext {
		public TerminalNode WHILE() { return getToken(SwiftGrammarParser.WHILE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public WhilestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whilestmt; }
	}

	public final WhilestmtContext whilestmt() throws RecognitionException {
		WhilestmtContext _localctx = new WhilestmtContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_whilestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(294);
			match(WHILE);
			setState(295);
			expr(0);
			setState(296);
			match(LLAVEIZQ);
			setState(297);
			block();
			setState(298);
			match(LLAVEDER);
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

	public static class OpasignstmtContext extends ParserRuleContext {
		public OpasignstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_opasignstmt; }
	 
		public OpasignstmtContext() { }
		public void copyFrom(OpasignstmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class DecrementoContext extends OpasignstmtContext {
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode SUB() { return getToken(SwiftGrammarParser.SUB, 0); }
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public DecrementoContext(OpasignstmtContext ctx) { copyFrom(ctx); }
	}
	public static class IncrementoContext extends OpasignstmtContext {
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode ADD() { return getToken(SwiftGrammarParser.ADD, 0); }
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public IncrementoContext(OpasignstmtContext ctx) { copyFrom(ctx); }
	}

	public final OpasignstmtContext opasignstmt() throws RecognitionException {
		OpasignstmtContext _localctx = new OpasignstmtContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_opasignstmt);
		try {
			setState(308);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				_localctx = new IncrementoContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(300);
				match(ID);
				setState(301);
				match(ADD);
				setState(302);
				match(IG);
				setState(303);
				expr(0);
				}
				break;
			case 2:
				_localctx = new DecrementoContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(304);
				match(ID);
				setState(305);
				match(SUB);
				setState(306);
				match(IG);
				setState(307);
				expr(0);
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

	public static class ExprContext extends ParserRuleContext {
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	 
		public ExprContext() { }
		public void copyFrom(ExprContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class BoolExprContext extends ExprContext {
		public TerminalNode TRU() { return getToken(SwiftGrammarParser.TRU, 0); }
		public TerminalNode FAL() { return getToken(SwiftGrammarParser.FAL, 0); }
		public BoolExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class StringExprContext extends ExprContext {
		public StringstmtContext stringstmt() {
			return getRuleContext(StringstmtContext.class,0);
		}
		public StringExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class NilExprContext extends ExprContext {
		public TerminalNode NIL() { return getToken(SwiftGrammarParser.NIL, 0); }
		public NilExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class FloatExprContext extends ExprContext {
		public FloatstmtContext floatstmt() {
			return getRuleContext(FloatstmtContext.class,0);
		}
		public FloatExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class IsEmptyVectorExprContext extends ExprContext {
		public IsemptyvectorstmtContext isemptyvectorstmt() {
			return getRuleContext(IsemptyvectorstmtContext.class,0);
		}
		public IsEmptyVectorExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class IdExprContext extends ExprContext {
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public IdExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class OpExprContext extends ExprContext {
		public ExprContext left;
		public Token op;
		public ExprContext right;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode MUL() { return getToken(SwiftGrammarParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(SwiftGrammarParser.DIV, 0); }
		public TerminalNode MOD() { return getToken(SwiftGrammarParser.MOD, 0); }
		public TerminalNode ADD() { return getToken(SwiftGrammarParser.ADD, 0); }
		public TerminalNode SUB() { return getToken(SwiftGrammarParser.SUB, 0); }
		public TerminalNode MAY_IG() { return getToken(SwiftGrammarParser.MAY_IG, 0); }
		public TerminalNode MAYOR() { return getToken(SwiftGrammarParser.MAYOR, 0); }
		public TerminalNode MEN_IG() { return getToken(SwiftGrammarParser.MEN_IG, 0); }
		public TerminalNode MENOR() { return getToken(SwiftGrammarParser.MENOR, 0); }
		public TerminalNode IG_IG() { return getToken(SwiftGrammarParser.IG_IG, 0); }
		public TerminalNode DIF() { return getToken(SwiftGrammarParser.DIF, 0); }
		public TerminalNode AND() { return getToken(SwiftGrammarParser.AND, 0); }
		public TerminalNode OR() { return getToken(SwiftGrammarParser.OR, 0); }
		public OpExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class AccesoVectorExprContext extends ExprContext {
		public AccesovectorstmtContext accesovectorstmt() {
			return getRuleContext(AccesovectorstmtContext.class,0);
		}
		public AccesoVectorExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class UnariaExprContext extends ExprContext {
		public TerminalNode SUB() { return getToken(SwiftGrammarParser.SUB, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public UnariaExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class CountVectorExprContext extends ExprContext {
		public CountvectorstmtContext countvectorstmt() {
			return getRuleContext(CountvectorstmtContext.class,0);
		}
		public CountVectorExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class NumExprContext extends ExprContext {
		public TerminalNode NUMBER() { return getToken(SwiftGrammarParser.NUMBER, 0); }
		public NumExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class ParExprContext extends ExprContext {
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public ParExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class StrExprContext extends ExprContext {
		public TerminalNode STRING() { return getToken(SwiftGrammarParser.STRING, 0); }
		public StrExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class AccFuncExprContext extends ExprContext {
		public AccfuncstmContext accfuncstm() {
			return getRuleContext(AccfuncstmContext.class,0);
		}
		public AccFuncExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class IntExprContext extends ExprContext {
		public IntstmtContext intstmt() {
			return getRuleContext(IntstmtContext.class,0);
		}
		public IntExprContext(ExprContext ctx) { copyFrom(ctx); }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 62;
		enterRecursionRule(_localctx, 62, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(329);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				{
				_localctx = new ParExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(311);
				match(PARIZQ);
				setState(312);
				expr(0);
				setState(313);
				match(PARDER);
				}
				break;
			case 2:
				{
				_localctx = new UnariaExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(315);
				match(SUB);
				setState(316);
				expr(20);
				}
				break;
			case 3:
				{
				_localctx = new NumExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(317);
				match(NUMBER);
				}
				break;
			case 4:
				{
				_localctx = new IdExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(318);
				match(ID);
				}
				break;
			case 5:
				{
				_localctx = new StrExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(319);
				match(STRING);
				}
				break;
			case 6:
				{
				_localctx = new BoolExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(320);
				_la = _input.LA(1);
				if ( !(_la==TRU || _la==FAL) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case 7:
				{
				_localctx = new NilExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(321);
				match(NIL);
				}
				break;
			case 8:
				{
				_localctx = new AccFuncExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(322);
				accfuncstm();
				}
				break;
			case 9:
				{
				_localctx = new IntExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(323);
				intstmt();
				}
				break;
			case 10:
				{
				_localctx = new FloatExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(324);
				floatstmt();
				}
				break;
			case 11:
				{
				_localctx = new StringExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(325);
				stringstmt();
				}
				break;
			case 12:
				{
				_localctx = new AccesoVectorExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(326);
				accesovectorstmt();
				}
				break;
			case 13:
				{
				_localctx = new CountVectorExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(327);
				countvectorstmt();
				}
				break;
			case 14:
				{
				_localctx = new IsEmptyVectorExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(328);
				isemptyvectorstmt();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(354);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(352);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
					case 1:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(331);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(332);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << MUL) | (1L << DIV) | (1L << MOD))) != 0)) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(333);
						((OpExprContext)_localctx).right = expr(20);
						}
						break;
					case 2:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(334);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(335);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==ADD || _la==SUB) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(336);
						((OpExprContext)_localctx).right = expr(19);
						}
						break;
					case 3:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(337);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(338);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MAY_IG || _la==MAYOR) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(339);
						((OpExprContext)_localctx).right = expr(18);
						}
						break;
					case 4:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(340);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(341);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MEN_IG || _la==MENOR) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(342);
						((OpExprContext)_localctx).right = expr(17);
						}
						break;
					case 5:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(343);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(344);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==DIF || _la==IG_IG) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(345);
						((OpExprContext)_localctx).right = expr(16);
						}
						break;
					case 6:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(346);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(347);
						((OpExprContext)_localctx).op = match(AND);
						setState(348);
						((OpExprContext)_localctx).right = expr(15);
						}
						break;
					case 7:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(349);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(350);
						((OpExprContext)_localctx).op = match(OR);
						setState(351);
						((OpExprContext)_localctx).right = expr(14);
						}
						break;
					}
					} 
				}
				setState(356);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
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

	public static class TipoContext extends ParserRuleContext {
		public TerminalNode INT() { return getToken(SwiftGrammarParser.INT, 0); }
		public TerminalNode FLOAT() { return getToken(SwiftGrammarParser.FLOAT, 0); }
		public TerminalNode BOOL() { return getToken(SwiftGrammarParser.BOOL, 0); }
		public TerminalNode CHARACTER() { return getToken(SwiftGrammarParser.CHARACTER, 0); }
		public TerminalNode PSTRING() { return getToken(SwiftGrammarParser.PSTRING, 0); }
		public TipoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo; }
	}

	public final TipoContext tipo() throws RecognitionException {
		TipoContext _localctx = new TipoContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_tipo);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(357);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << INT) | (1L << FLOAT) | (1L << BOOL) | (1L << CHARACTER) | (1L << PSTRING))) != 0)) ) {
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
		case 31:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 19);
		case 1:
			return precpred(_ctx, 18);
		case 2:
			return precpred(_ctx, 17);
		case 3:
			return precpred(_ctx, 16);
		case 4:
			return precpred(_ctx, 15);
		case 5:
			return precpred(_ctx, 14);
		case 6:
			return precpred(_ctx, 13);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\38\u016a\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\3\2\3\2\3\2\3\3\7\3I\n\3\f\3\16\3L\13\3\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4^\n\4\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\5\6k\n\6\3\6\3\6\3\6\5\6p\n\6\3\7\3\7\3"+
		"\7\7\7u\n\7\f\7\16\7x\13\7\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3"+
		"\r\3\r\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20"+
		"\3\20\3\20\3\21\3\21\3\21\3\21\5\21\u00aa\n\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\5\21\u00b5\n\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\5\21\u00bf\n\21\3\22\3\22\3\22\5\22\u00c4\n\22\3\22\3\22\3\23\3"+
		"\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\7\23\u00d1\n\23\f\23\16\23\u00d4"+
		"\13\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\7\24\u00dd\n\24\f\24\16\24\u00e0"+
		"\13\24\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\26\7\26\u00ea\n\26\f\26\16"+
		"\26\u00ed\13\26\3\26\3\26\3\26\3\26\3\26\5\26\u00f4\n\26\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\6\30\u0101\n\30\r\30\16\30\u0102"+
		"\3\30\5\30\u0106\n\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\32\3\32\3\32"+
		"\3\32\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\35"+
		"\3\35\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37"+
		"\3\37\3 \3 \3 \3 \3 \3 \3 \3 \5 \u0137\n \3!\3!\3!\3!\3!\3!\3!\3!\3!\3"+
		"!\3!\3!\3!\3!\3!\3!\3!\3!\3!\5!\u014c\n!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3"+
		"!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\7!\u0163\n!\f!\16!\u0166\13!\3\"\3"+
		"\"\3\"\2\3@#\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64"+
		"\668:<>@B\2\t\3\2\t\n\4\2&\'**\3\2()\4\2\"\"$$\4\2##%%\3\2\34\35\3\2\3"+
		"\7\2\u017a\2D\3\2\2\2\4J\3\2\2\2\6]\3\2\2\2\b_\3\2\2\2\no\3\2\2\2\fq\3"+
		"\2\2\2\16y\3\2\2\2\20~\3\2\2\2\22\u0085\3\2\2\2\24\u0089\3\2\2\2\26\u008d"+
		"\3\2\2\2\30\u0091\3\2\2\2\32\u0096\3\2\2\2\34\u009b\3\2\2\2\36\u00a0\3"+
		"\2\2\2 \u00be\3\2\2\2\"\u00c0\3\2\2\2$\u00c7\3\2\2\2&\u00d5\3\2\2\2(\u00e1"+
		"\3\2\2\2*\u00e3\3\2\2\2,\u00f5\3\2\2\2.\u00fc\3\2\2\2\60\u0109\3\2\2\2"+
		"\62\u010e\3\2\2\2\64\u0112\3\2\2\2\66\u0119\3\2\2\28\u011e\3\2\2\2:\u0124"+
		"\3\2\2\2<\u0128\3\2\2\2>\u0136\3\2\2\2@\u014b\3\2\2\2B\u0167\3\2\2\2D"+
		"E\5\4\3\2EF\7\2\2\3F\3\3\2\2\2GI\5\6\4\2HG\3\2\2\2IL\3\2\2\2JH\3\2\2\2"+
		"JK\3\2\2\2K\5\3\2\2\2LJ\3\2\2\2M^\5\30\r\2N^\5\64\33\2O^\5\66\34\2P^\5"+
		"8\35\2Q^\5:\36\2R^\5*\26\2S^\5.\30\2T^\5<\37\2U^\5> \2V^\5(\25\2W^\5 "+
		"\21\2X^\5\"\22\2Y^\5\26\f\2Z^\5\b\5\2[^\5\16\b\2\\^\5\20\t\2]M\3\2\2\2"+
		"]N\3\2\2\2]O\3\2\2\2]P\3\2\2\2]Q\3\2\2\2]R\3\2\2\2]S\3\2\2\2]T\3\2\2\2"+
		"]U\3\2\2\2]V\3\2\2\2]W\3\2\2\2]X\3\2\2\2]Y\3\2\2\2]Z\3\2\2\2][\3\2\2\2"+
		"]\\\3\2\2\2^\7\3\2\2\2_`\7\22\2\2`a\7\33\2\2ab\7\61\2\2bc\7/\2\2cd\5B"+
		"\"\2de\7\60\2\2ef\5\n\6\2f\t\3\2\2\2gh\7!\2\2hj\7/\2\2ik\5\f\7\2ji\3\2"+
		"\2\2jk\3\2\2\2kl\3\2\2\2lp\7\60\2\2mn\7!\2\2np\7\33\2\2og\3\2\2\2om\3"+
		"\2\2\2p\13\3\2\2\2qv\5@!\2rs\7\62\2\2su\5@!\2tr\3\2\2\2ux\3\2\2\2vt\3"+
		"\2\2\2vw\3\2\2\2w\r\3\2\2\2xv\3\2\2\2yz\7\33\2\2z{\7/\2\2{|\5@!\2|}\7"+
		"\60\2\2}\17\3\2\2\2~\177\7\33\2\2\177\u0080\7\65\2\2\u0080\u0081\7\30"+
		"\2\2\u0081\u0082\7+\2\2\u0082\u0083\5@!\2\u0083\u0084\7,\2\2\u0084\21"+
		"\3\2\2\2\u0085\u0086\7\33\2\2\u0086\u0087\7\65\2\2\u0087\u0088\7\26\2"+
		"\2\u0088\23\3\2\2\2\u0089\u008a\7\33\2\2\u008a\u008b\7\65\2\2\u008b\u008c"+
		"\7\27\2\2\u008c\25\3\2\2\2\u008d\u008e\7\24\2\2\u008e\u008f\5@!\2\u008f"+
		"\u0090\7\63\2\2\u0090\27\3\2\2\2\u0091\u0092\7\13\2\2\u0092\u0093\7+\2"+
		"\2\u0093\u0094\5@!\2\u0094\u0095\7,\2\2\u0095\31\3\2\2\2\u0096\u0097\7"+
		"\3\2\2\u0097\u0098\7+\2\2\u0098\u0099\5@!\2\u0099\u009a\7,\2\2\u009a\33"+
		"\3\2\2\2\u009b\u009c\7\4\2\2\u009c\u009d\7+\2\2\u009d\u009e\5@!\2\u009e"+
		"\u009f\7,\2\2\u009f\35\3\2\2\2\u00a0\u00a1\7\7\2\2\u00a1\u00a2\7+\2\2"+
		"\u00a2\u00a3\5@!\2\u00a3\u00a4\7,\2\2\u00a4\37\3\2\2\2\u00a5\u00a6\7\25"+
		"\2\2\u00a6\u00a7\7\33\2\2\u00a7\u00a9\7+\2\2\u00a8\u00aa\5$\23\2\u00a9"+
		"\u00a8\3\2\2\2\u00a9\u00aa\3\2\2\2\u00aa\u00ab\3\2\2\2\u00ab\u00ac\7,"+
		"\2\2\u00ac\u00ad\7-\2\2\u00ad\u00ae\5\4\3\2\u00ae\u00af\7.\2\2\u00af\u00bf"+
		"\3\2\2\2\u00b0\u00b1\7\25\2\2\u00b1\u00b2\7\33\2\2\u00b2\u00b4\7+\2\2"+
		"\u00b3\u00b5\5$\23\2\u00b4\u00b3\3\2\2\2\u00b4\u00b5\3\2\2\2\u00b5\u00b6"+
		"\3\2\2\2\u00b6\u00b7\7,\2\2\u00b7\u00b8\7)\2\2\u00b8\u00b9\7$\2\2\u00b9"+
		"\u00ba\5B\"\2\u00ba\u00bb\7-\2\2\u00bb\u00bc\5\4\3\2\u00bc\u00bd\7.\2"+
		"\2\u00bd\u00bf\3\2\2\2\u00be\u00a5\3\2\2\2\u00be\u00b0\3\2\2\2\u00bf!"+
		"\3\2\2\2\u00c0\u00c1\7\33\2\2\u00c1\u00c3\7+\2\2\u00c2\u00c4\5&\24\2\u00c3"+
		"\u00c2\3\2\2\2\u00c3\u00c4\3\2\2\2\u00c4\u00c5\3\2\2\2\u00c5\u00c6\7,"+
		"\2\2\u00c6#\3\2\2\2\u00c7\u00c8\7\33\2\2\u00c8\u00c9\7\33\2\2\u00c9\u00ca"+
		"\7\61\2\2\u00ca\u00d2\5B\"\2\u00cb\u00cc\7\62\2\2\u00cc\u00cd\7\33\2\2"+
		"\u00cd\u00ce\7\33\2\2\u00ce\u00cf\7\61\2\2\u00cf\u00d1\5B\"\2\u00d0\u00cb"+
		"\3\2\2\2\u00d1\u00d4\3\2\2\2\u00d2\u00d0\3\2\2\2\u00d2\u00d3\3\2\2\2\u00d3"+
		"%\3\2\2\2\u00d4\u00d2\3\2\2\2\u00d5\u00d6\7\33\2\2\u00d6\u00d7\7\61\2"+
		"\2\u00d7\u00de\5@!\2\u00d8\u00d9\7\62\2\2\u00d9\u00da\7\33\2\2\u00da\u00db"+
		"\7\61\2\2\u00db\u00dd\5@!\2\u00dc\u00d8\3\2\2\2\u00dd\u00e0\3\2\2\2\u00de"+
		"\u00dc\3\2\2\2\u00de\u00df\3\2\2\2\u00df\'\3\2\2\2\u00e0\u00de\3\2\2\2"+
		"\u00e1\u00e2\7\23\2\2\u00e2)\3\2\2\2\u00e3\u00e4\7\f\2\2\u00e4\u00e5\5"+
		"@!\2\u00e5\u00e6\7-\2\2\u00e6\u00e7\5\4\3\2\u00e7\u00eb\7.\2\2\u00e8\u00ea"+
		"\5,\27\2\u00e9\u00e8\3\2\2\2\u00ea\u00ed\3\2\2\2\u00eb\u00e9\3\2\2\2\u00eb"+
		"\u00ec\3\2\2\2\u00ec\u00f3\3\2\2\2\u00ed\u00eb\3\2\2\2\u00ee\u00ef\7\r"+
		"\2\2\u00ef\u00f0\7-\2\2\u00f0\u00f1\5\4\3\2\u00f1\u00f2\7.\2\2\u00f2\u00f4"+
		"\3\2\2\2\u00f3\u00ee\3\2\2\2\u00f3\u00f4\3\2\2\2\u00f4+\3\2\2\2\u00f5"+
		"\u00f6\7\r\2\2\u00f6\u00f7\7\f\2\2\u00f7\u00f8\5@!\2\u00f8\u00f9\7-\2"+
		"\2\u00f9\u00fa\5\4\3\2\u00fa\u00fb\7.\2\2\u00fb-\3\2\2\2\u00fc\u00fd\7"+
		"\17\2\2\u00fd\u00fe\5@!\2\u00fe\u0100\7-\2\2\u00ff\u0101\5\60\31\2\u0100"+
		"\u00ff\3\2\2\2\u0101\u0102\3\2\2\2\u0102\u0100\3\2\2\2\u0102\u0103\3\2"+
		"\2\2\u0103\u0105\3\2\2\2\u0104\u0106\5\62\32\2\u0105\u0104\3\2\2\2\u0105"+
		"\u0106\3\2\2\2\u0106\u0107\3\2\2\2\u0107\u0108\7.\2\2\u0108/\3\2\2\2\u0109"+
		"\u010a\7\20\2\2\u010a\u010b\5@!\2\u010b\u010c\7\61\2\2\u010c\u010d\5\4"+
		"\3\2\u010d\61\3\2\2\2\u010e\u010f\7\21\2\2\u010f\u0110\7\61\2\2\u0110"+
		"\u0111\5\4\3\2\u0111\63\3\2\2\2\u0112\u0113\7\22\2\2\u0113\u0114\7\33"+
		"\2\2\u0114\u0115\7\61\2\2\u0115\u0116\5B\"\2\u0116\u0117\7!\2\2\u0117"+
		"\u0118\5@!\2\u0118\65\3\2\2\2\u0119\u011a\7\22\2\2\u011a\u011b\7\33\2"+
		"\2\u011b\u011c\7!\2\2\u011c\u011d\5@!\2\u011d\67\3\2\2\2\u011e\u011f\7"+
		"\22\2\2\u011f\u0120\7\33\2\2\u0120\u0121\7\61\2\2\u0121\u0122\5B\"\2\u0122"+
		"\u0123\7\64\2\2\u01239\3\2\2\2\u0124\u0125\7\33\2\2\u0125\u0126\7!\2\2"+
		"\u0126\u0127\5@!\2\u0127;\3\2\2\2\u0128\u0129\7\16\2\2\u0129\u012a\5@"+
		"!\2\u012a\u012b\7-\2\2\u012b\u012c\5\4\3\2\u012c\u012d\7.\2\2\u012d=\3"+
		"\2\2\2\u012e\u012f\7\33\2\2\u012f\u0130\7(\2\2\u0130\u0131\7!\2\2\u0131"+
		"\u0137\5@!\2\u0132\u0133\7\33\2\2\u0133\u0134\7)\2\2\u0134\u0135\7!\2"+
		"\2\u0135\u0137\5@!\2\u0136\u012e\3\2\2\2\u0136\u0132\3\2\2\2\u0137?\3"+
		"\2\2\2\u0138\u0139\b!\1\2\u0139\u013a\7+\2\2\u013a\u013b\5@!\2\u013b\u013c"+
		"\7,\2\2\u013c\u014c\3\2\2\2\u013d\u013e\7)\2\2\u013e\u014c\5@!\26\u013f"+
		"\u014c\7\31\2\2\u0140\u014c\7\33\2\2\u0141\u014c\7\32\2\2\u0142\u014c"+
		"\t\2\2\2\u0143\u014c\7\b\2\2\u0144\u014c\5\"\22\2\u0145\u014c\5\32\16"+
		"\2\u0146\u014c\5\34\17\2\u0147\u014c\5\36\20\2\u0148\u014c\5\16\b\2\u0149"+
		"\u014c\5\22\n\2\u014a\u014c\5\24\13\2\u014b\u0138\3\2\2\2\u014b\u013d"+
		"\3\2\2\2\u014b\u013f\3\2\2\2\u014b\u0140\3\2\2\2\u014b\u0141\3\2\2\2\u014b"+
		"\u0142\3\2\2\2\u014b\u0143\3\2\2\2\u014b\u0144\3\2\2\2\u014b\u0145\3\2"+
		"\2\2\u014b\u0146\3\2\2\2\u014b\u0147\3\2\2\2\u014b\u0148\3\2\2\2\u014b"+
		"\u0149\3\2\2\2\u014b\u014a\3\2\2\2\u014c\u0164\3\2\2\2\u014d\u014e\f\25"+
		"\2\2\u014e\u014f\t\3\2\2\u014f\u0163\5@!\26\u0150\u0151\f\24\2\2\u0151"+
		"\u0152\t\4\2\2\u0152\u0163\5@!\25\u0153\u0154\f\23\2\2\u0154\u0155\t\5"+
		"\2\2\u0155\u0163\5@!\24\u0156\u0157\f\22\2\2\u0157\u0158\t\6\2\2\u0158"+
		"\u0163\5@!\23\u0159\u015a\f\21\2\2\u015a\u015b\t\7\2\2\u015b\u0163\5@"+
		"!\22\u015c\u015d\f\20\2\2\u015d\u015e\7 \2\2\u015e\u0163\5@!\21\u015f"+
		"\u0160\f\17\2\2\u0160\u0161\7\37\2\2\u0161\u0163\5@!\20\u0162\u014d\3"+
		"\2\2\2\u0162\u0150\3\2\2\2\u0162\u0153\3\2\2\2\u0162\u0156\3\2\2\2\u0162"+
		"\u0159\3\2\2\2\u0162\u015c\3\2\2\2\u0162\u015f\3\2\2\2\u0163\u0166\3\2"+
		"\2\2\u0164\u0162\3\2\2\2\u0164\u0165\3\2\2\2\u0165A\3\2\2\2\u0166\u0164"+
		"\3\2\2\2\u0167\u0168\t\b\2\2\u0168C\3\2\2\2\25J]jov\u00a9\u00b4\u00be"+
		"\u00c3\u00d2\u00de\u00eb\u00f3\u0102\u0105\u0136\u014b\u0162\u0164";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
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
		IF=10, ELSE=11, WHILE=12, FOR=13, IN=14, SWITCH=15, CASE=16, DEFAULT=17, 
		VAR=18, BREAK=19, RETURN=20, FUNC=21, COUNT=22, ISEMPTY=23, APPEND=24, 
		REMOVE_LAST=25, REMOVE=26, AT=27, NUMBER=28, STRING=29, ID=30, DIF=31, 
		IG_IG=32, NOT=33, OR=34, AND=35, IG=36, MAY_IG=37, MEN_IG=38, MAYOR=39, 
		MENOR=40, MUL=41, DIV=42, ADD=43, SUB=44, MOD=45, PARIZQ=46, PARDER=47, 
		LLAVEIZQ=48, LLAVEDER=49, CORCHIZQ=50, CORCHDER=51, DOSPUNTOS=52, COMA=53, 
		PTCOMA=54, INTERROGACION=55, PUNTO=56, WHITESPACE=57, COMMENT=58, LINE_COMMENT=59;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_stmt = 2, RULE_declvectorstmt = 3, RULE_defvectorstmt = 4, 
		RULE_listaexpresiones = 5, RULE_accesovectorstmt = 6, RULE_asignvectorstmt = 7, 
		RULE_appendvectorstmt = 8, RULE_removeatvectorstmt = 9, RULE_removelastvectorstmt = 10, 
		RULE_countvectorstmt = 11, RULE_isemptyvectorstmt = 12, RULE_declmatrizstmt = 13, 
		RULE_tipomatriz = 14, RULE_listavaloresmatriz = 15, RULE_listavaloresmatriz3 = 16, 
		RULE_accesomatriz = 17, RULE_asignmatrizstmt = 18, RULE_returnstmt = 19, 
		RULE_printstmt = 20, RULE_intstmt = 21, RULE_floatstmt = 22, RULE_stringstmt = 23, 
		RULE_funcdclstmt = 24, RULE_accfuncstm = 25, RULE_parametros = 26, RULE_parametroscall = 27, 
		RULE_breakstmt = 28, RULE_ifstmt = 29, RULE_elseifstmt = 30, RULE_switchstmt = 31, 
		RULE_caseStmt = 32, RULE_defaultCase = 33, RULE_typedDeclstmt = 34, RULE_untypedDeclstmt = 35, 
		RULE_optionalTypedDeclstmt = 36, RULE_asignstmt = 37, RULE_whilestmt = 38, 
		RULE_forstmt = 39, RULE_rangostmt = 40, RULE_opasignstmt = 41, RULE_expr = 42, 
		RULE_tipo = 43;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "stmt", "declvectorstmt", "defvectorstmt", "listaexpresiones", 
			"accesovectorstmt", "asignvectorstmt", "appendvectorstmt", "removeatvectorstmt", 
			"removelastvectorstmt", "countvectorstmt", "isemptyvectorstmt", "declmatrizstmt", 
			"tipomatriz", "listavaloresmatriz", "listavaloresmatriz3", "accesomatriz", 
			"asignmatrizstmt", "returnstmt", "printstmt", "intstmt", "floatstmt", 
			"stringstmt", "funcdclstmt", "accfuncstm", "parametros", "parametroscall", 
			"breakstmt", "ifstmt", "elseifstmt", "switchstmt", "caseStmt", "defaultCase", 
			"typedDeclstmt", "untypedDeclstmt", "optionalTypedDeclstmt", "asignstmt", 
			"whilestmt", "forstmt", "rangostmt", "opasignstmt", "expr", "tipo"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'int'", "'float'", "'bool'", "'character'", "'String'", "'nil'", 
			"'true'", "'false'", "'print'", "'if'", "'else'", "'while'", "'for'", 
			"'in'", "'switch'", "'case'", "'default'", "'var'", "'break'", "'return'", 
			"'func'", "'count'", "'IsEmpty'", "'append'", "'removeLast'", "'remove'", 
			"'at'", null, null, null, "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", 
			"'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'%'", "'('", 
			"')'", "'{'", "'}'", "'['", "']'", "':'", "','", "';'", "'?'", "'.'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "NIL", "TRU", "FAL", 
			"PRINT", "IF", "ELSE", "WHILE", "FOR", "IN", "SWITCH", "CASE", "DEFAULT", 
			"VAR", "BREAK", "RETURN", "FUNC", "COUNT", "ISEMPTY", "APPEND", "REMOVE_LAST", 
			"REMOVE", "AT", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", 
			"AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", 
			"SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "CORCHIZQ", 
			"CORCHDER", "DOSPUNTOS", "COMA", "PTCOMA", "INTERROGACION", "PUNTO", 
			"WHITESPACE", "COMMENT", "LINE_COMMENT"
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
			setState(88);
			block();
			setState(89);
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
		public List<TerminalNode> PTCOMA() { return getTokens(SwiftGrammarParser.PTCOMA); }
		public TerminalNode PTCOMA(int i) {
			return getToken(SwiftGrammarParser.PTCOMA, i);
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
			setState(97);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PRINT) | (1L << IF) | (1L << WHILE) | (1L << FOR) | (1L << SWITCH) | (1L << VAR) | (1L << BREAK) | (1L << RETURN) | (1L << FUNC) | (1L << ID))) != 0)) {
				{
				{
				setState(91);
				stmt();
				setState(93);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(92);
					match(PTCOMA);
					}
				}

				}
				}
				setState(99);
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
		public ForstmtContext forstmt() {
			return getRuleContext(ForstmtContext.class,0);
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
		public RemovelastvectorstmtContext removelastvectorstmt() {
			return getRuleContext(RemovelastvectorstmtContext.class,0);
		}
		public RemoveatvectorstmtContext removeatvectorstmt() {
			return getRuleContext(RemoveatvectorstmtContext.class,0);
		}
		public AsignvectorstmtContext asignvectorstmt() {
			return getRuleContext(AsignvectorstmtContext.class,0);
		}
		public DeclmatrizstmtContext declmatrizstmt() {
			return getRuleContext(DeclmatrizstmtContext.class,0);
		}
		public AsignmatrizstmtContext asignmatrizstmt() {
			return getRuleContext(AsignmatrizstmtContext.class,0);
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
			setState(122);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(100);
				printstmt();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(101);
				typedDeclstmt();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(102);
				untypedDeclstmt();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(103);
				optionalTypedDeclstmt();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(104);
				asignstmt();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(105);
				ifstmt();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(106);
				switchstmt();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(107);
				whilestmt();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(108);
				forstmt();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(109);
				opasignstmt();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(110);
				breakstmt();
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(111);
				funcdclstmt();
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(112);
				accfuncstm();
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(113);
				returnstmt();
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(114);
				declvectorstmt();
				}
				break;
			case 16:
				enterOuterAlt(_localctx, 16);
				{
				setState(115);
				accesovectorstmt();
				}
				break;
			case 17:
				enterOuterAlt(_localctx, 17);
				{
				setState(116);
				appendvectorstmt();
				}
				break;
			case 18:
				enterOuterAlt(_localctx, 18);
				{
				setState(117);
				removelastvectorstmt();
				}
				break;
			case 19:
				enterOuterAlt(_localctx, 19);
				{
				setState(118);
				removeatvectorstmt();
				}
				break;
			case 20:
				enterOuterAlt(_localctx, 20);
				{
				setState(119);
				asignvectorstmt();
				}
				break;
			case 21:
				enterOuterAlt(_localctx, 21);
				{
				setState(120);
				declmatrizstmt();
				}
				break;
			case 22:
				enterOuterAlt(_localctx, 22);
				{
				setState(121);
				asignmatrizstmt();
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
			setState(124);
			match(VAR);
			setState(125);
			match(ID);
			setState(126);
			match(DOSPUNTOS);
			setState(127);
			match(CORCHIZQ);
			setState(128);
			tipo();
			setState(129);
			match(CORCHDER);
			setState(130);
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
			setState(140);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				_localctx = new DefVectorContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(132);
				match(IG);
				setState(133);
				match(CORCHIZQ);
				setState(135);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << INT) | (1L << FLOAT) | (1L << PSTRING) | (1L << NIL) | (1L << TRU) | (1L << FAL) | (1L << NUMBER) | (1L << STRING) | (1L << ID) | (1L << SUB) | (1L << PARIZQ))) != 0)) {
					{
					setState(134);
					listaexpresiones();
					}
				}

				setState(137);
				match(CORCHDER);
				}
				break;
			case 2:
				_localctx = new DefVectorIDContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(138);
				match(IG);
				setState(139);
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
			setState(142);
			expr(0);
			setState(147);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(143);
				match(COMA);
				setState(144);
				expr(0);
				}
				}
				setState(149);
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
			setState(150);
			match(ID);
			setState(151);
			match(CORCHIZQ);
			setState(152);
			expr(0);
			setState(153);
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

	public static class AsignvectorstmtContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode CORCHIZQ() { return getToken(SwiftGrammarParser.CORCHIZQ, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode CORCHDER() { return getToken(SwiftGrammarParser.CORCHDER, 0); }
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public AsignvectorstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignvectorstmt; }
	}

	public final AsignvectorstmtContext asignvectorstmt() throws RecognitionException {
		AsignvectorstmtContext _localctx = new AsignvectorstmtContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_asignvectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(155);
			match(ID);
			setState(156);
			match(CORCHIZQ);
			setState(157);
			expr(0);
			setState(158);
			match(CORCHDER);
			setState(159);
			match(IG);
			setState(160);
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
		enterRule(_localctx, 16, RULE_appendvectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(162);
			match(ID);
			setState(163);
			match(PUNTO);
			setState(164);
			match(APPEND);
			setState(165);
			match(PARIZQ);
			setState(166);
			expr(0);
			setState(167);
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

	public static class RemoveatvectorstmtContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
		public TerminalNode REMOVE() { return getToken(SwiftGrammarParser.REMOVE, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public TerminalNode AT() { return getToken(SwiftGrammarParser.AT, 0); }
		public TerminalNode DOSPUNTOS() { return getToken(SwiftGrammarParser.DOSPUNTOS, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public RemoveatvectorstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_removeatvectorstmt; }
	}

	public final RemoveatvectorstmtContext removeatvectorstmt() throws RecognitionException {
		RemoveatvectorstmtContext _localctx = new RemoveatvectorstmtContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_removeatvectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(169);
			match(ID);
			setState(170);
			match(PUNTO);
			setState(171);
			match(REMOVE);
			setState(172);
			match(PARIZQ);
			setState(173);
			match(AT);
			setState(174);
			match(DOSPUNTOS);
			setState(175);
			expr(0);
			setState(176);
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

	public static class RemovelastvectorstmtContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
		public TerminalNode REMOVE_LAST() { return getToken(SwiftGrammarParser.REMOVE_LAST, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public RemovelastvectorstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_removelastvectorstmt; }
	}

	public final RemovelastvectorstmtContext removelastvectorstmt() throws RecognitionException {
		RemovelastvectorstmtContext _localctx = new RemovelastvectorstmtContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_removelastvectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(178);
			match(ID);
			setState(179);
			match(PUNTO);
			setState(180);
			match(REMOVE_LAST);
			setState(181);
			match(PARIZQ);
			setState(182);
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
		enterRule(_localctx, 22, RULE_countvectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(184);
			match(ID);
			setState(185);
			match(PUNTO);
			setState(186);
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
		enterRule(_localctx, 24, RULE_isemptyvectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(188);
			match(ID);
			setState(189);
			match(PUNTO);
			setState(190);
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

	public static class DeclmatrizstmtContext extends ParserRuleContext {
		public DeclmatrizstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declmatrizstmt; }
	 
		public DeclmatrizstmtContext() { }
		public void copyFrom(DeclmatrizstmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Declmatrizstmt2Context extends DeclmatrizstmtContext {
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ListavaloresmatrizContext listavaloresmatriz() {
			return getRuleContext(ListavaloresmatrizContext.class,0);
		}
		public TerminalNode DOSPUNTOS() { return getToken(SwiftGrammarParser.DOSPUNTOS, 0); }
		public TipomatrizContext tipomatriz() {
			return getRuleContext(TipomatrizContext.class,0);
		}
		public Declmatrizstmt2Context(DeclmatrizstmtContext ctx) { copyFrom(ctx); }
	}
	public static class Declmatrizstmt3Context extends DeclmatrizstmtContext {
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public Listavaloresmatriz3Context listavaloresmatriz3() {
			return getRuleContext(Listavaloresmatriz3Context.class,0);
		}
		public TerminalNode DOSPUNTOS() { return getToken(SwiftGrammarParser.DOSPUNTOS, 0); }
		public TipomatrizContext tipomatriz() {
			return getRuleContext(TipomatrizContext.class,0);
		}
		public Declmatrizstmt3Context(DeclmatrizstmtContext ctx) { copyFrom(ctx); }
	}

	public final DeclmatrizstmtContext declmatrizstmt() throws RecognitionException {
		DeclmatrizstmtContext _localctx = new DeclmatrizstmtContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_declmatrizstmt);
		int _la;
		try {
			setState(208);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				_localctx = new Declmatrizstmt2Context(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(192);
				match(VAR);
				setState(193);
				match(ID);
				setState(196);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DOSPUNTOS) {
					{
					setState(194);
					match(DOSPUNTOS);
					setState(195);
					tipomatriz();
					}
				}

				setState(198);
				match(IG);
				setState(199);
				listavaloresmatriz();
				}
				break;
			case 2:
				_localctx = new Declmatrizstmt3Context(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(200);
				match(VAR);
				setState(201);
				match(ID);
				setState(204);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DOSPUNTOS) {
					{
					setState(202);
					match(DOSPUNTOS);
					setState(203);
					tipomatriz();
					}
				}

				setState(206);
				match(IG);
				setState(207);
				listavaloresmatriz3();
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

	public static class TipomatrizContext extends ParserRuleContext {
		public TipomatrizContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipomatriz; }
	 
		public TipomatrizContext() { }
		public void copyFrom(TipomatrizContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Tipomatriz2Context extends TipomatrizContext {
		public List<TerminalNode> CORCHIZQ() { return getTokens(SwiftGrammarParser.CORCHIZQ); }
		public TerminalNode CORCHIZQ(int i) {
			return getToken(SwiftGrammarParser.CORCHIZQ, i);
		}
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public List<TerminalNode> CORCHDER() { return getTokens(SwiftGrammarParser.CORCHDER); }
		public TerminalNode CORCHDER(int i) {
			return getToken(SwiftGrammarParser.CORCHDER, i);
		}
		public Tipomatriz2Context(TipomatrizContext ctx) { copyFrom(ctx); }
	}
	public static class Tipomatriz3Context extends TipomatrizContext {
		public List<TerminalNode> CORCHIZQ() { return getTokens(SwiftGrammarParser.CORCHIZQ); }
		public TerminalNode CORCHIZQ(int i) {
			return getToken(SwiftGrammarParser.CORCHIZQ, i);
		}
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public List<TerminalNode> CORCHDER() { return getTokens(SwiftGrammarParser.CORCHDER); }
		public TerminalNode CORCHDER(int i) {
			return getToken(SwiftGrammarParser.CORCHDER, i);
		}
		public Tipomatriz3Context(TipomatrizContext ctx) { copyFrom(ctx); }
	}

	public final TipomatrizContext tipomatriz() throws RecognitionException {
		TipomatrizContext _localctx = new TipomatrizContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_tipomatriz);
		try {
			setState(224);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				_localctx = new Tipomatriz2Context(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(210);
				match(CORCHIZQ);
				setState(211);
				match(CORCHIZQ);
				setState(212);
				tipo();
				setState(213);
				match(CORCHDER);
				setState(214);
				match(CORCHDER);
				}
				break;
			case 2:
				_localctx = new Tipomatriz3Context(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(216);
				match(CORCHIZQ);
				setState(217);
				match(CORCHIZQ);
				setState(218);
				match(CORCHIZQ);
				setState(219);
				tipo();
				setState(220);
				match(CORCHDER);
				setState(221);
				match(CORCHDER);
				setState(222);
				match(CORCHDER);
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

	public static class ListavaloresmatrizContext extends ParserRuleContext {
		public ListavaloresmatrizContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listavaloresmatriz; }
	 
		public ListavaloresmatrizContext() { }
		public void copyFrom(ListavaloresmatrizContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Listavaloresmatriz2Context extends ListavaloresmatrizContext {
		public List<TerminalNode> CORCHIZQ() { return getTokens(SwiftGrammarParser.CORCHIZQ); }
		public TerminalNode CORCHIZQ(int i) {
			return getToken(SwiftGrammarParser.CORCHIZQ, i);
		}
		public List<ListaexpresionesContext> listaexpresiones() {
			return getRuleContexts(ListaexpresionesContext.class);
		}
		public ListaexpresionesContext listaexpresiones(int i) {
			return getRuleContext(ListaexpresionesContext.class,i);
		}
		public List<TerminalNode> CORCHDER() { return getTokens(SwiftGrammarParser.CORCHDER); }
		public TerminalNode CORCHDER(int i) {
			return getToken(SwiftGrammarParser.CORCHDER, i);
		}
		public List<TerminalNode> COMA() { return getTokens(SwiftGrammarParser.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(SwiftGrammarParser.COMA, i);
		}
		public Listavaloresmatriz2Context(ListavaloresmatrizContext ctx) { copyFrom(ctx); }
	}

	public final ListavaloresmatrizContext listavaloresmatriz() throws RecognitionException {
		ListavaloresmatrizContext _localctx = new ListavaloresmatrizContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_listavaloresmatriz);
		int _la;
		try {
			_localctx = new Listavaloresmatriz2Context(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(226);
			match(CORCHIZQ);
			setState(227);
			match(CORCHIZQ);
			setState(228);
			listaexpresiones();
			setState(229);
			match(CORCHDER);
			setState(235); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(230);
				match(COMA);
				setState(231);
				match(CORCHIZQ);
				setState(232);
				listaexpresiones();
				setState(233);
				match(CORCHDER);
				}
				}
				setState(237); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==COMA );
			setState(239);
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

	public static class Listavaloresmatriz3Context extends ParserRuleContext {
		public TerminalNode CORCHIZQ() { return getToken(SwiftGrammarParser.CORCHIZQ, 0); }
		public List<ListavaloresmatrizContext> listavaloresmatriz() {
			return getRuleContexts(ListavaloresmatrizContext.class);
		}
		public ListavaloresmatrizContext listavaloresmatriz(int i) {
			return getRuleContext(ListavaloresmatrizContext.class,i);
		}
		public TerminalNode CORCHDER() { return getToken(SwiftGrammarParser.CORCHDER, 0); }
		public List<TerminalNode> COMA() { return getTokens(SwiftGrammarParser.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(SwiftGrammarParser.COMA, i);
		}
		public Listavaloresmatriz3Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listavaloresmatriz3; }
	}

	public final Listavaloresmatriz3Context listavaloresmatriz3() throws RecognitionException {
		Listavaloresmatriz3Context _localctx = new Listavaloresmatriz3Context(_ctx, getState());
		enterRule(_localctx, 32, RULE_listavaloresmatriz3);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(241);
			match(CORCHIZQ);
			setState(242);
			listavaloresmatriz();
			setState(247);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(243);
				match(COMA);
				setState(244);
				listavaloresmatriz();
				}
				}
				setState(249);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(250);
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

	public static class AccesomatrizContext extends ParserRuleContext {
		public AccesomatrizContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accesomatriz; }
	 
		public AccesomatrizContext() { }
		public void copyFrom(AccesomatrizContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Accesomatriz2Context extends AccesomatrizContext {
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public List<TerminalNode> CORCHIZQ() { return getTokens(SwiftGrammarParser.CORCHIZQ); }
		public TerminalNode CORCHIZQ(int i) {
			return getToken(SwiftGrammarParser.CORCHIZQ, i);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> CORCHDER() { return getTokens(SwiftGrammarParser.CORCHDER); }
		public TerminalNode CORCHDER(int i) {
			return getToken(SwiftGrammarParser.CORCHDER, i);
		}
		public Accesomatriz2Context(AccesomatrizContext ctx) { copyFrom(ctx); }
	}
	public static class Accesomatriz3Context extends AccesomatrizContext {
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public List<TerminalNode> CORCHIZQ() { return getTokens(SwiftGrammarParser.CORCHIZQ); }
		public TerminalNode CORCHIZQ(int i) {
			return getToken(SwiftGrammarParser.CORCHIZQ, i);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> CORCHDER() { return getTokens(SwiftGrammarParser.CORCHDER); }
		public TerminalNode CORCHDER(int i) {
			return getToken(SwiftGrammarParser.CORCHDER, i);
		}
		public Accesomatriz3Context(AccesomatrizContext ctx) { copyFrom(ctx); }
	}

	public final AccesomatrizContext accesomatriz() throws RecognitionException {
		AccesomatrizContext _localctx = new AccesomatrizContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_accesomatriz);
		try {
			setState(271);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				_localctx = new Accesomatriz2Context(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(252);
				match(ID);
				setState(253);
				match(CORCHIZQ);
				setState(254);
				expr(0);
				setState(255);
				match(CORCHDER);
				setState(256);
				match(CORCHIZQ);
				setState(257);
				expr(0);
				setState(258);
				match(CORCHDER);
				}
				break;
			case 2:
				_localctx = new Accesomatriz3Context(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(260);
				match(ID);
				setState(261);
				match(CORCHIZQ);
				setState(262);
				expr(0);
				setState(263);
				match(CORCHDER);
				setState(264);
				match(CORCHIZQ);
				setState(265);
				expr(0);
				setState(266);
				match(CORCHDER);
				setState(267);
				match(CORCHIZQ);
				setState(268);
				expr(0);
				setState(269);
				match(CORCHDER);
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

	public static class AsignmatrizstmtContext extends ParserRuleContext {
		public AsignmatrizstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignmatrizstmt; }
	 
		public AsignmatrizstmtContext() { }
		public void copyFrom(AsignmatrizstmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Asignmatrizstmt3Context extends AsignmatrizstmtContext {
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public List<TerminalNode> CORCHIZQ() { return getTokens(SwiftGrammarParser.CORCHIZQ); }
		public TerminalNode CORCHIZQ(int i) {
			return getToken(SwiftGrammarParser.CORCHIZQ, i);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> CORCHDER() { return getTokens(SwiftGrammarParser.CORCHDER); }
		public TerminalNode CORCHDER(int i) {
			return getToken(SwiftGrammarParser.CORCHDER, i);
		}
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public Asignmatrizstmt3Context(AsignmatrizstmtContext ctx) { copyFrom(ctx); }
	}
	public static class Asignmatrizstmt2Context extends AsignmatrizstmtContext {
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public List<TerminalNode> CORCHIZQ() { return getTokens(SwiftGrammarParser.CORCHIZQ); }
		public TerminalNode CORCHIZQ(int i) {
			return getToken(SwiftGrammarParser.CORCHIZQ, i);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> CORCHDER() { return getTokens(SwiftGrammarParser.CORCHDER); }
		public TerminalNode CORCHDER(int i) {
			return getToken(SwiftGrammarParser.CORCHDER, i);
		}
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public Asignmatrizstmt2Context(AsignmatrizstmtContext ctx) { copyFrom(ctx); }
	}

	public final AsignmatrizstmtContext asignmatrizstmt() throws RecognitionException {
		AsignmatrizstmtContext _localctx = new AsignmatrizstmtContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_asignmatrizstmt);
		try {
			setState(296);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				_localctx = new Asignmatrizstmt2Context(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(273);
				match(ID);
				setState(274);
				match(CORCHIZQ);
				setState(275);
				expr(0);
				setState(276);
				match(CORCHDER);
				setState(277);
				match(CORCHIZQ);
				setState(278);
				expr(0);
				setState(279);
				match(CORCHDER);
				setState(280);
				match(IG);
				setState(281);
				expr(0);
				}
				break;
			case 2:
				_localctx = new Asignmatrizstmt3Context(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(283);
				match(ID);
				setState(284);
				match(CORCHIZQ);
				setState(285);
				expr(0);
				setState(286);
				match(CORCHDER);
				setState(287);
				match(CORCHIZQ);
				setState(288);
				expr(0);
				setState(289);
				match(CORCHDER);
				setState(290);
				match(CORCHIZQ);
				setState(291);
				expr(0);
				setState(292);
				match(CORCHDER);
				setState(293);
				match(IG);
				setState(294);
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
		enterRule(_localctx, 38, RULE_returnstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(298);
			match(RETURN);
			setState(299);
			expr(0);
			setState(300);
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
		enterRule(_localctx, 40, RULE_printstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(302);
			match(PRINT);
			setState(303);
			match(PARIZQ);
			setState(304);
			expr(0);
			setState(305);
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
		enterRule(_localctx, 42, RULE_intstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(307);
			match(INT);
			setState(308);
			match(PARIZQ);
			setState(309);
			expr(0);
			setState(310);
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
		enterRule(_localctx, 44, RULE_floatstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(312);
			match(FLOAT);
			setState(313);
			match(PARIZQ);
			setState(314);
			expr(0);
			setState(315);
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
		enterRule(_localctx, 46, RULE_stringstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(317);
			match(PSTRING);
			setState(318);
			match(PARIZQ);
			setState(319);
			expr(0);
			setState(320);
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
		enterRule(_localctx, 48, RULE_funcdclstmt);
		int _la;
		try {
			setState(347);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				_localctx = new FuncionNormalContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(322);
				match(FUNC);
				setState(323);
				match(ID);
				setState(324);
				match(PARIZQ);
				setState(326);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ID) {
					{
					setState(325);
					parametros();
					}
				}

				setState(328);
				match(PARDER);
				setState(329);
				match(LLAVEIZQ);
				setState(330);
				block();
				setState(331);
				match(LLAVEDER);
				}
				break;
			case 2:
				_localctx = new FuncionRetornoContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(333);
				match(FUNC);
				setState(334);
				match(ID);
				setState(335);
				match(PARIZQ);
				setState(337);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ID) {
					{
					setState(336);
					parametros();
					}
				}

				setState(339);
				match(PARDER);
				setState(340);
				match(SUB);
				setState(341);
				match(MAYOR);
				setState(342);
				tipo();
				setState(343);
				match(LLAVEIZQ);
				setState(344);
				block();
				setState(345);
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
		enterRule(_localctx, 50, RULE_accfuncstm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(349);
			match(ID);
			setState(350);
			match(PARIZQ);
			setState(352);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(351);
				parametroscall();
				}
			}

			setState(354);
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
		enterRule(_localctx, 52, RULE_parametros);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(356);
			match(ID);
			setState(357);
			match(ID);
			setState(358);
			match(DOSPUNTOS);
			setState(359);
			tipo();
			setState(367);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(360);
				match(COMA);
				setState(361);
				match(ID);
				setState(362);
				match(ID);
				setState(363);
				match(DOSPUNTOS);
				setState(364);
				tipo();
				}
				}
				setState(369);
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
		enterRule(_localctx, 54, RULE_parametroscall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(370);
			match(ID);
			setState(371);
			match(DOSPUNTOS);
			setState(372);
			expr(0);
			setState(379);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(373);
				match(COMA);
				setState(374);
				match(ID);
				setState(375);
				match(DOSPUNTOS);
				setState(376);
				expr(0);
				}
				}
				setState(381);
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
		enterRule(_localctx, 56, RULE_breakstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(382);
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
		enterRule(_localctx, 58, RULE_ifstmt);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(384);
			match(IF);
			setState(385);
			expr(0);
			setState(386);
			match(LLAVEIZQ);
			setState(387);
			block();
			setState(388);
			match(LLAVEDER);
			setState(392);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,20,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(389);
					elseifstmt();
					}
					} 
				}
				setState(394);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,20,_ctx);
			}
			setState(400);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(395);
				match(ELSE);
				setState(396);
				match(LLAVEIZQ);
				setState(397);
				block();
				setState(398);
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
		enterRule(_localctx, 60, RULE_elseifstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(402);
			match(ELSE);
			setState(403);
			match(IF);
			setState(404);
			expr(0);
			setState(405);
			match(LLAVEIZQ);
			setState(406);
			block();
			setState(407);
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
		enterRule(_localctx, 62, RULE_switchstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(409);
			match(SWITCH);
			setState(410);
			expr(0);
			setState(411);
			match(LLAVEIZQ);
			setState(413); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(412);
				caseStmt();
				}
				}
				setState(415); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==CASE );
			setState(418);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(417);
				defaultCase();
				}
			}

			setState(420);
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
		enterRule(_localctx, 64, RULE_caseStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(422);
			match(CASE);
			setState(423);
			expr(0);
			setState(424);
			match(DOSPUNTOS);
			setState(425);
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
		enterRule(_localctx, 66, RULE_defaultCase);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(427);
			match(DEFAULT);
			setState(428);
			match(DOSPUNTOS);
			setState(429);
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
		enterRule(_localctx, 68, RULE_typedDeclstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(431);
			match(VAR);
			setState(432);
			match(ID);
			setState(433);
			match(DOSPUNTOS);
			setState(434);
			tipo();
			setState(435);
			match(IG);
			setState(436);
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
		enterRule(_localctx, 70, RULE_untypedDeclstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(438);
			match(VAR);
			setState(439);
			match(ID);
			setState(440);
			match(IG);
			setState(441);
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
		enterRule(_localctx, 72, RULE_optionalTypedDeclstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(443);
			match(VAR);
			setState(444);
			match(ID);
			setState(445);
			match(DOSPUNTOS);
			setState(446);
			tipo();
			setState(447);
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
		enterRule(_localctx, 74, RULE_asignstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(449);
			match(ID);
			setState(450);
			match(IG);
			setState(451);
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
		enterRule(_localctx, 76, RULE_whilestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(453);
			match(WHILE);
			setState(454);
			expr(0);
			setState(455);
			match(LLAVEIZQ);
			setState(456);
			block();
			setState(457);
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

	public static class ForstmtContext extends ParserRuleContext {
		public TerminalNode FOR() { return getToken(SwiftGrammarParser.FOR, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode IN() { return getToken(SwiftGrammarParser.IN, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public RangostmtContext rangostmt() {
			return getRuleContext(RangostmtContext.class,0);
		}
		public ForstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forstmt; }
	}

	public final ForstmtContext forstmt() throws RecognitionException {
		ForstmtContext _localctx = new ForstmtContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_forstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(459);
			match(FOR);
			setState(460);
			match(ID);
			setState(461);
			match(IN);
			setState(464);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				{
				setState(462);
				expr(0);
				}
				break;
			case 2:
				{
				setState(463);
				rangostmt();
				}
				break;
			}
			setState(466);
			match(LLAVEIZQ);
			setState(467);
			block();
			setState(468);
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

	public static class RangostmtContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> PUNTO() { return getTokens(SwiftGrammarParser.PUNTO); }
		public TerminalNode PUNTO(int i) {
			return getToken(SwiftGrammarParser.PUNTO, i);
		}
		public RangostmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rangostmt; }
	}

	public final RangostmtContext rangostmt() throws RecognitionException {
		RangostmtContext _localctx = new RangostmtContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_rangostmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(470);
			expr(0);
			setState(471);
			match(PUNTO);
			setState(472);
			match(PUNTO);
			setState(473);
			match(PUNTO);
			setState(474);
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
		enterRule(_localctx, 82, RULE_opasignstmt);
		try {
			setState(484);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				_localctx = new IncrementoContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(476);
				match(ID);
				setState(477);
				match(ADD);
				setState(478);
				match(IG);
				setState(479);
				expr(0);
				}
				break;
			case 2:
				_localctx = new DecrementoContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(480);
				match(ID);
				setState(481);
				match(SUB);
				setState(482);
				match(IG);
				setState(483);
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
	public static class AccesoMatrizExprContext extends ExprContext {
		public AccesomatrizContext accesomatriz() {
			return getRuleContext(AccesomatrizContext.class,0);
		}
		public AccesoMatrizExprContext(ExprContext ctx) { copyFrom(ctx); }
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
		int _startState = 84;
		enterRecursionRule(_localctx, 84, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(506);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
			case 1:
				{
				_localctx = new ParExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(487);
				match(PARIZQ);
				setState(488);
				expr(0);
				setState(489);
				match(PARDER);
				}
				break;
			case 2:
				{
				_localctx = new UnariaExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(491);
				match(SUB);
				setState(492);
				expr(21);
				}
				break;
			case 3:
				{
				_localctx = new NumExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(493);
				match(NUMBER);
				}
				break;
			case 4:
				{
				_localctx = new IdExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(494);
				match(ID);
				}
				break;
			case 5:
				{
				_localctx = new StrExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(495);
				match(STRING);
				}
				break;
			case 6:
				{
				_localctx = new BoolExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(496);
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
				setState(497);
				match(NIL);
				}
				break;
			case 8:
				{
				_localctx = new AccFuncExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(498);
				accfuncstm();
				}
				break;
			case 9:
				{
				_localctx = new IntExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(499);
				intstmt();
				}
				break;
			case 10:
				{
				_localctx = new FloatExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(500);
				floatstmt();
				}
				break;
			case 11:
				{
				_localctx = new StringExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(501);
				stringstmt();
				}
				break;
			case 12:
				{
				_localctx = new AccesoVectorExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(502);
				accesovectorstmt();
				}
				break;
			case 13:
				{
				_localctx = new CountVectorExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(503);
				countvectorstmt();
				}
				break;
			case 14:
				{
				_localctx = new IsEmptyVectorExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(504);
				isemptyvectorstmt();
				}
				break;
			case 15:
				{
				_localctx = new AccesoMatrizExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(505);
				accesomatriz();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(531);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(529);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
					case 1:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(508);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(509);
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
						setState(510);
						((OpExprContext)_localctx).right = expr(21);
						}
						break;
					case 2:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(511);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(512);
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
						setState(513);
						((OpExprContext)_localctx).right = expr(20);
						}
						break;
					case 3:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(514);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(515);
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
						setState(516);
						((OpExprContext)_localctx).right = expr(19);
						}
						break;
					case 4:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(517);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(518);
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
						setState(519);
						((OpExprContext)_localctx).right = expr(18);
						}
						break;
					case 5:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(520);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(521);
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
						setState(522);
						((OpExprContext)_localctx).right = expr(17);
						}
						break;
					case 6:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(523);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(524);
						((OpExprContext)_localctx).op = match(AND);
						setState(525);
						((OpExprContext)_localctx).right = expr(16);
						}
						break;
					case 7:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(526);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(527);
						((OpExprContext)_localctx).op = match(OR);
						setState(528);
						((OpExprContext)_localctx).right = expr(15);
						}
						break;
					}
					} 
				}
				setState(533);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
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
		enterRule(_localctx, 86, RULE_tipo);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(534);
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
		case 42:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 20);
		case 1:
			return precpred(_ctx, 19);
		case 2:
			return precpred(_ctx, 18);
		case 3:
			return precpred(_ctx, 17);
		case 4:
			return precpred(_ctx, 16);
		case 5:
			return precpred(_ctx, 15);
		case 6:
			return precpred(_ctx, 14);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3=\u021b\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\3\2\3\2\3\2\3\3\3\3\5\3`\n\3\7\3b\n\3\f\3\16\3e\13\3\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\5\4}\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\5\6"+
		"\u008a\n\6\3\6\3\6\3\6\5\6\u008f\n\6\3\7\3\7\3\7\7\7\u0094\n\7\f\7\16"+
		"\7\u0097\13\7\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\17\3\17\3\17"+
		"\3\17\5\17\u00c7\n\17\3\17\3\17\3\17\3\17\3\17\3\17\5\17\u00cf\n\17\3"+
		"\17\3\17\5\17\u00d3\n\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\5\20\u00e3\n\20\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\6\21\u00ee\n\21\r\21\16\21\u00ef\3\21\3\21\3\22\3\22\3"+
		"\22\3\22\7\22\u00f8\n\22\f\22\16\22\u00fb\13\22\3\22\3\22\3\23\3\23\3"+
		"\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3"+
		"\23\3\23\3\23\5\23\u0112\n\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\5\24\u012b\n\24\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\27"+
		"\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31"+
		"\3\32\3\32\3\32\3\32\5\32\u0149\n\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32"+
		"\3\32\3\32\5\32\u0154\n\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\5\32"+
		"\u015e\n\32\3\33\3\33\3\33\5\33\u0163\n\33\3\33\3\33\3\34\3\34\3\34\3"+
		"\34\3\34\3\34\3\34\3\34\3\34\7\34\u0170\n\34\f\34\16\34\u0173\13\34\3"+
		"\35\3\35\3\35\3\35\3\35\3\35\3\35\7\35\u017c\n\35\f\35\16\35\u017f\13"+
		"\35\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3\37\7\37\u0189\n\37\f\37\16\37"+
		"\u018c\13\37\3\37\3\37\3\37\3\37\3\37\5\37\u0193\n\37\3 \3 \3 \3 \3 \3"+
		" \3 \3!\3!\3!\3!\6!\u01a0\n!\r!\16!\u01a1\3!\5!\u01a5\n!\3!\3!\3\"\3\""+
		"\3\"\3\"\3\"\3#\3#\3#\3#\3$\3$\3$\3$\3$\3$\3$\3%\3%\3%\3%\3%\3&\3&\3&"+
		"\3&\3&\3&\3\'\3\'\3\'\3\'\3(\3(\3(\3(\3(\3(\3)\3)\3)\3)\3)\5)\u01d3\n"+
		")\3)\3)\3)\3)\3*\3*\3*\3*\3*\3*\3+\3+\3+\3+\3+\3+\3+\3+\5+\u01e7\n+\3"+
		",\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\5,\u01fd\n"+
		",\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\7,\u0214"+
		"\n,\f,\16,\u0217\13,\3-\3-\3-\2\3V.\2\4\6\b\n\f\16\20\22\24\26\30\32\34"+
		"\36 \"$&(*,.\60\62\64\668:<>@BDFHJLNPRTVX\2\t\3\2\t\n\4\2+,//\3\2-.\4"+
		"\2\'\'))\4\2((**\3\2!\"\3\2\3\7\2\u0231\2Z\3\2\2\2\4c\3\2\2\2\6|\3\2\2"+
		"\2\b~\3\2\2\2\n\u008e\3\2\2\2\f\u0090\3\2\2\2\16\u0098\3\2\2\2\20\u009d"+
		"\3\2\2\2\22\u00a4\3\2\2\2\24\u00ab\3\2\2\2\26\u00b4\3\2\2\2\30\u00ba\3"+
		"\2\2\2\32\u00be\3\2\2\2\34\u00d2\3\2\2\2\36\u00e2\3\2\2\2 \u00e4\3\2\2"+
		"\2\"\u00f3\3\2\2\2$\u0111\3\2\2\2&\u012a\3\2\2\2(\u012c\3\2\2\2*\u0130"+
		"\3\2\2\2,\u0135\3\2\2\2.\u013a\3\2\2\2\60\u013f\3\2\2\2\62\u015d\3\2\2"+
		"\2\64\u015f\3\2\2\2\66\u0166\3\2\2\28\u0174\3\2\2\2:\u0180\3\2\2\2<\u0182"+
		"\3\2\2\2>\u0194\3\2\2\2@\u019b\3\2\2\2B\u01a8\3\2\2\2D\u01ad\3\2\2\2F"+
		"\u01b1\3\2\2\2H\u01b8\3\2\2\2J\u01bd\3\2\2\2L\u01c3\3\2\2\2N\u01c7\3\2"+
		"\2\2P\u01cd\3\2\2\2R\u01d8\3\2\2\2T\u01e6\3\2\2\2V\u01fc\3\2\2\2X\u0218"+
		"\3\2\2\2Z[\5\4\3\2[\\\7\2\2\3\\\3\3\2\2\2]_\5\6\4\2^`\78\2\2_^\3\2\2\2"+
		"_`\3\2\2\2`b\3\2\2\2a]\3\2\2\2be\3\2\2\2ca\3\2\2\2cd\3\2\2\2d\5\3\2\2"+
		"\2ec\3\2\2\2f}\5*\26\2g}\5F$\2h}\5H%\2i}\5J&\2j}\5L\'\2k}\5<\37\2l}\5"+
		"@!\2m}\5N(\2n}\5P)\2o}\5T+\2p}\5:\36\2q}\5\62\32\2r}\5\64\33\2s}\5(\25"+
		"\2t}\5\b\5\2u}\5\16\b\2v}\5\22\n\2w}\5\26\f\2x}\5\24\13\2y}\5\20\t\2z"+
		"}\5\34\17\2{}\5&\24\2|f\3\2\2\2|g\3\2\2\2|h\3\2\2\2|i\3\2\2\2|j\3\2\2"+
		"\2|k\3\2\2\2|l\3\2\2\2|m\3\2\2\2|n\3\2\2\2|o\3\2\2\2|p\3\2\2\2|q\3\2\2"+
		"\2|r\3\2\2\2|s\3\2\2\2|t\3\2\2\2|u\3\2\2\2|v\3\2\2\2|w\3\2\2\2|x\3\2\2"+
		"\2|y\3\2\2\2|z\3\2\2\2|{\3\2\2\2}\7\3\2\2\2~\177\7\24\2\2\177\u0080\7"+
		" \2\2\u0080\u0081\7\66\2\2\u0081\u0082\7\64\2\2\u0082\u0083\5X-\2\u0083"+
		"\u0084\7\65\2\2\u0084\u0085\5\n\6\2\u0085\t\3\2\2\2\u0086\u0087\7&\2\2"+
		"\u0087\u0089\7\64\2\2\u0088\u008a\5\f\7\2\u0089\u0088\3\2\2\2\u0089\u008a"+
		"\3\2\2\2\u008a\u008b\3\2\2\2\u008b\u008f\7\65\2\2\u008c\u008d\7&\2\2\u008d"+
		"\u008f\7 \2\2\u008e\u0086\3\2\2\2\u008e\u008c\3\2\2\2\u008f\13\3\2\2\2"+
		"\u0090\u0095\5V,\2\u0091\u0092\7\67\2\2\u0092\u0094\5V,\2\u0093\u0091"+
		"\3\2\2\2\u0094\u0097\3\2\2\2\u0095\u0093\3\2\2\2\u0095\u0096\3\2\2\2\u0096"+
		"\r\3\2\2\2\u0097\u0095\3\2\2\2\u0098\u0099\7 \2\2\u0099\u009a\7\64\2\2"+
		"\u009a\u009b\5V,\2\u009b\u009c\7\65\2\2\u009c\17\3\2\2\2\u009d\u009e\7"+
		" \2\2\u009e\u009f\7\64\2\2\u009f\u00a0\5V,\2\u00a0\u00a1\7\65\2\2\u00a1"+
		"\u00a2\7&\2\2\u00a2\u00a3\5V,\2\u00a3\21\3\2\2\2\u00a4\u00a5\7 \2\2\u00a5"+
		"\u00a6\7:\2\2\u00a6\u00a7\7\32\2\2\u00a7\u00a8\7\60\2\2\u00a8\u00a9\5"+
		"V,\2\u00a9\u00aa\7\61\2\2\u00aa\23\3\2\2\2\u00ab\u00ac\7 \2\2\u00ac\u00ad"+
		"\7:\2\2\u00ad\u00ae\7\34\2\2\u00ae\u00af\7\60\2\2\u00af\u00b0\7\35\2\2"+
		"\u00b0\u00b1\7\66\2\2\u00b1\u00b2\5V,\2\u00b2\u00b3\7\61\2\2\u00b3\25"+
		"\3\2\2\2\u00b4\u00b5\7 \2\2\u00b5\u00b6\7:\2\2\u00b6\u00b7\7\33\2\2\u00b7"+
		"\u00b8\7\60\2\2\u00b8\u00b9\7\61\2\2\u00b9\27\3\2\2\2\u00ba\u00bb\7 \2"+
		"\2\u00bb\u00bc\7:\2\2\u00bc\u00bd\7\30\2\2\u00bd\31\3\2\2\2\u00be\u00bf"+
		"\7 \2\2\u00bf\u00c0\7:\2\2\u00c0\u00c1\7\31\2\2\u00c1\33\3\2\2\2\u00c2"+
		"\u00c3\7\24\2\2\u00c3\u00c6\7 \2\2\u00c4\u00c5\7\66\2\2\u00c5\u00c7\5"+
		"\36\20\2\u00c6\u00c4\3\2\2\2\u00c6\u00c7\3\2\2\2\u00c7\u00c8\3\2\2\2\u00c8"+
		"\u00c9\7&\2\2\u00c9\u00d3\5 \21\2\u00ca\u00cb\7\24\2\2\u00cb\u00ce\7 "+
		"\2\2\u00cc\u00cd\7\66\2\2\u00cd\u00cf\5\36\20\2\u00ce\u00cc\3\2\2\2\u00ce"+
		"\u00cf\3\2\2\2\u00cf\u00d0\3\2\2\2\u00d0\u00d1\7&\2\2\u00d1\u00d3\5\""+
		"\22\2\u00d2\u00c2\3\2\2\2\u00d2\u00ca\3\2\2\2\u00d3\35\3\2\2\2\u00d4\u00d5"+
		"\7\64\2\2\u00d5\u00d6\7\64\2\2\u00d6\u00d7\5X-\2\u00d7\u00d8\7\65\2\2"+
		"\u00d8\u00d9\7\65\2\2\u00d9\u00e3\3\2\2\2\u00da\u00db\7\64\2\2\u00db\u00dc"+
		"\7\64\2\2\u00dc\u00dd\7\64\2\2\u00dd\u00de\5X-\2\u00de\u00df\7\65\2\2"+
		"\u00df\u00e0\7\65\2\2\u00e0\u00e1\7\65\2\2\u00e1\u00e3\3\2\2\2\u00e2\u00d4"+
		"\3\2\2\2\u00e2\u00da\3\2\2\2\u00e3\37\3\2\2\2\u00e4\u00e5\7\64\2\2\u00e5"+
		"\u00e6\7\64\2\2\u00e6\u00e7\5\f\7\2\u00e7\u00ed\7\65\2\2\u00e8\u00e9\7"+
		"\67\2\2\u00e9\u00ea\7\64\2\2\u00ea\u00eb\5\f\7\2\u00eb\u00ec\7\65\2\2"+
		"\u00ec\u00ee\3\2\2\2\u00ed\u00e8\3\2\2\2\u00ee\u00ef\3\2\2\2\u00ef\u00ed"+
		"\3\2\2\2\u00ef\u00f0\3\2\2\2\u00f0\u00f1\3\2\2\2\u00f1\u00f2\7\65\2\2"+
		"\u00f2!\3\2\2\2\u00f3\u00f4\7\64\2\2\u00f4\u00f9\5 \21\2\u00f5\u00f6\7"+
		"\67\2\2\u00f6\u00f8\5 \21\2\u00f7\u00f5\3\2\2\2\u00f8\u00fb\3\2\2\2\u00f9"+
		"\u00f7\3\2\2\2\u00f9\u00fa\3\2\2\2\u00fa\u00fc\3\2\2\2\u00fb\u00f9\3\2"+
		"\2\2\u00fc\u00fd\7\65\2\2\u00fd#\3\2\2\2\u00fe\u00ff\7 \2\2\u00ff\u0100"+
		"\7\64\2\2\u0100\u0101\5V,\2\u0101\u0102\7\65\2\2\u0102\u0103\7\64\2\2"+
		"\u0103\u0104\5V,\2\u0104\u0105\7\65\2\2\u0105\u0112\3\2\2\2\u0106\u0107"+
		"\7 \2\2\u0107\u0108\7\64\2\2\u0108\u0109\5V,\2\u0109\u010a\7\65\2\2\u010a"+
		"\u010b\7\64\2\2\u010b\u010c\5V,\2\u010c\u010d\7\65\2\2\u010d\u010e\7\64"+
		"\2\2\u010e\u010f\5V,\2\u010f\u0110\7\65\2\2\u0110\u0112\3\2\2\2\u0111"+
		"\u00fe\3\2\2\2\u0111\u0106\3\2\2\2\u0112%\3\2\2\2\u0113\u0114\7 \2\2\u0114"+
		"\u0115\7\64\2\2\u0115\u0116\5V,\2\u0116\u0117\7\65\2\2\u0117\u0118\7\64"+
		"\2\2\u0118\u0119\5V,\2\u0119\u011a\7\65\2\2\u011a\u011b\7&\2\2\u011b\u011c"+
		"\5V,\2\u011c\u012b\3\2\2\2\u011d\u011e\7 \2\2\u011e\u011f\7\64\2\2\u011f"+
		"\u0120\5V,\2\u0120\u0121\7\65\2\2\u0121\u0122\7\64\2\2\u0122\u0123\5V"+
		",\2\u0123\u0124\7\65\2\2\u0124\u0125\7\64\2\2\u0125\u0126\5V,\2\u0126"+
		"\u0127\7\65\2\2\u0127\u0128\7&\2\2\u0128\u0129\5V,\2\u0129\u012b\3\2\2"+
		"\2\u012a\u0113\3\2\2\2\u012a\u011d\3\2\2\2\u012b\'\3\2\2\2\u012c\u012d"+
		"\7\26\2\2\u012d\u012e\5V,\2\u012e\u012f\78\2\2\u012f)\3\2\2\2\u0130\u0131"+
		"\7\13\2\2\u0131\u0132\7\60\2\2\u0132\u0133\5V,\2\u0133\u0134\7\61\2\2"+
		"\u0134+\3\2\2\2\u0135\u0136\7\3\2\2\u0136\u0137\7\60\2\2\u0137\u0138\5"+
		"V,\2\u0138\u0139\7\61\2\2\u0139-\3\2\2\2\u013a\u013b\7\4\2\2\u013b\u013c"+
		"\7\60\2\2\u013c\u013d\5V,\2\u013d\u013e\7\61\2\2\u013e/\3\2\2\2\u013f"+
		"\u0140\7\7\2\2\u0140\u0141\7\60\2\2\u0141\u0142\5V,\2\u0142\u0143\7\61"+
		"\2\2\u0143\61\3\2\2\2\u0144\u0145\7\27\2\2\u0145\u0146\7 \2\2\u0146\u0148"+
		"\7\60\2\2\u0147\u0149\5\66\34\2\u0148\u0147\3\2\2\2\u0148\u0149\3\2\2"+
		"\2\u0149\u014a\3\2\2\2\u014a\u014b\7\61\2\2\u014b\u014c\7\62\2\2\u014c"+
		"\u014d\5\4\3\2\u014d\u014e\7\63\2\2\u014e\u015e\3\2\2\2\u014f\u0150\7"+
		"\27\2\2\u0150\u0151\7 \2\2\u0151\u0153\7\60\2\2\u0152\u0154\5\66\34\2"+
		"\u0153\u0152\3\2\2\2\u0153\u0154\3\2\2\2\u0154\u0155\3\2\2\2\u0155\u0156"+
		"\7\61\2\2\u0156\u0157\7.\2\2\u0157\u0158\7)\2\2\u0158\u0159\5X-\2\u0159"+
		"\u015a\7\62\2\2\u015a\u015b\5\4\3\2\u015b\u015c\7\63\2\2\u015c\u015e\3"+
		"\2\2\2\u015d\u0144\3\2\2\2\u015d\u014f\3\2\2\2\u015e\63\3\2\2\2\u015f"+
		"\u0160\7 \2\2\u0160\u0162\7\60\2\2\u0161\u0163\58\35\2\u0162\u0161\3\2"+
		"\2\2\u0162\u0163\3\2\2\2\u0163\u0164\3\2\2\2\u0164\u0165\7\61\2\2\u0165"+
		"\65\3\2\2\2\u0166\u0167\7 \2\2\u0167\u0168\7 \2\2\u0168\u0169\7\66\2\2"+
		"\u0169\u0171\5X-\2\u016a\u016b\7\67\2\2\u016b\u016c\7 \2\2\u016c\u016d"+
		"\7 \2\2\u016d\u016e\7\66\2\2\u016e\u0170\5X-\2\u016f\u016a\3\2\2\2\u0170"+
		"\u0173\3\2\2\2\u0171\u016f\3\2\2\2\u0171\u0172\3\2\2\2\u0172\67\3\2\2"+
		"\2\u0173\u0171\3\2\2\2\u0174\u0175\7 \2\2\u0175\u0176\7\66\2\2\u0176\u017d"+
		"\5V,\2\u0177\u0178\7\67\2\2\u0178\u0179\7 \2\2\u0179\u017a\7\66\2\2\u017a"+
		"\u017c\5V,\2\u017b\u0177\3\2\2\2\u017c\u017f\3\2\2\2\u017d\u017b\3\2\2"+
		"\2\u017d\u017e\3\2\2\2\u017e9\3\2\2\2\u017f\u017d\3\2\2\2\u0180\u0181"+
		"\7\25\2\2\u0181;\3\2\2\2\u0182\u0183\7\f\2\2\u0183\u0184\5V,\2\u0184\u0185"+
		"\7\62\2\2\u0185\u0186\5\4\3\2\u0186\u018a\7\63\2\2\u0187\u0189\5> \2\u0188"+
		"\u0187\3\2\2\2\u0189\u018c\3\2\2\2\u018a\u0188\3\2\2\2\u018a\u018b\3\2"+
		"\2\2\u018b\u0192\3\2\2\2\u018c\u018a\3\2\2\2\u018d\u018e\7\r\2\2\u018e"+
		"\u018f\7\62\2\2\u018f\u0190\5\4\3\2\u0190\u0191\7\63\2\2\u0191\u0193\3"+
		"\2\2\2\u0192\u018d\3\2\2\2\u0192\u0193\3\2\2\2\u0193=\3\2\2\2\u0194\u0195"+
		"\7\r\2\2\u0195\u0196\7\f\2\2\u0196\u0197\5V,\2\u0197\u0198\7\62\2\2\u0198"+
		"\u0199\5\4\3\2\u0199\u019a\7\63\2\2\u019a?\3\2\2\2\u019b\u019c\7\21\2"+
		"\2\u019c\u019d\5V,\2\u019d\u019f\7\62\2\2\u019e\u01a0\5B\"\2\u019f\u019e"+
		"\3\2\2\2\u01a0\u01a1\3\2\2\2\u01a1\u019f\3\2\2\2\u01a1\u01a2\3\2\2\2\u01a2"+
		"\u01a4\3\2\2\2\u01a3\u01a5\5D#\2\u01a4\u01a3\3\2\2\2\u01a4\u01a5\3\2\2"+
		"\2\u01a5\u01a6\3\2\2\2\u01a6\u01a7\7\63\2\2\u01a7A\3\2\2\2\u01a8\u01a9"+
		"\7\22\2\2\u01a9\u01aa\5V,\2\u01aa\u01ab\7\66\2\2\u01ab\u01ac\5\4\3\2\u01ac"+
		"C\3\2\2\2\u01ad\u01ae\7\23\2\2\u01ae\u01af\7\66\2\2\u01af\u01b0\5\4\3"+
		"\2\u01b0E\3\2\2\2\u01b1\u01b2\7\24\2\2\u01b2\u01b3\7 \2\2\u01b3\u01b4"+
		"\7\66\2\2\u01b4\u01b5\5X-\2\u01b5\u01b6\7&\2\2\u01b6\u01b7\5V,\2\u01b7"+
		"G\3\2\2\2\u01b8\u01b9\7\24\2\2\u01b9\u01ba\7 \2\2\u01ba\u01bb\7&\2\2\u01bb"+
		"\u01bc\5V,\2\u01bcI\3\2\2\2\u01bd\u01be\7\24\2\2\u01be\u01bf\7 \2\2\u01bf"+
		"\u01c0\7\66\2\2\u01c0\u01c1\5X-\2\u01c1\u01c2\79\2\2\u01c2K\3\2\2\2\u01c3"+
		"\u01c4\7 \2\2\u01c4\u01c5\7&\2\2\u01c5\u01c6\5V,\2\u01c6M\3\2\2\2\u01c7"+
		"\u01c8\7\16\2\2\u01c8\u01c9\5V,\2\u01c9\u01ca\7\62\2\2\u01ca\u01cb\5\4"+
		"\3\2\u01cb\u01cc\7\63\2\2\u01ccO\3\2\2\2\u01cd\u01ce\7\17\2\2\u01ce\u01cf"+
		"\7 \2\2\u01cf\u01d2\7\20\2\2\u01d0\u01d3\5V,\2\u01d1\u01d3\5R*\2\u01d2"+
		"\u01d0\3\2\2\2\u01d2\u01d1\3\2\2\2\u01d3\u01d4\3\2\2\2\u01d4\u01d5\7\62"+
		"\2\2\u01d5\u01d6\5\4\3\2\u01d6\u01d7\7\63\2\2\u01d7Q\3\2\2\2\u01d8\u01d9"+
		"\5V,\2\u01d9\u01da\7:\2\2\u01da\u01db\7:\2\2\u01db\u01dc\7:\2\2\u01dc"+
		"\u01dd\5V,\2\u01ddS\3\2\2\2\u01de\u01df\7 \2\2\u01df\u01e0\7-\2\2\u01e0"+
		"\u01e1\7&\2\2\u01e1\u01e7\5V,\2\u01e2\u01e3\7 \2\2\u01e3\u01e4\7.\2\2"+
		"\u01e4\u01e5\7&\2\2\u01e5\u01e7\5V,\2\u01e6\u01de\3\2\2\2\u01e6\u01e2"+
		"\3\2\2\2\u01e7U\3\2\2\2\u01e8\u01e9\b,\1\2\u01e9\u01ea\7\60\2\2\u01ea"+
		"\u01eb\5V,\2\u01eb\u01ec\7\61\2\2\u01ec\u01fd\3\2\2\2\u01ed\u01ee\7.\2"+
		"\2\u01ee\u01fd\5V,\27\u01ef\u01fd\7\36\2\2\u01f0\u01fd\7 \2\2\u01f1\u01fd"+
		"\7\37\2\2\u01f2\u01fd\t\2\2\2\u01f3\u01fd\7\b\2\2\u01f4\u01fd\5\64\33"+
		"\2\u01f5\u01fd\5,\27\2\u01f6\u01fd\5.\30\2\u01f7\u01fd\5\60\31\2\u01f8"+
		"\u01fd\5\16\b\2\u01f9\u01fd\5\30\r\2\u01fa\u01fd\5\32\16\2\u01fb\u01fd"+
		"\5$\23\2\u01fc\u01e8\3\2\2\2\u01fc\u01ed\3\2\2\2\u01fc\u01ef\3\2\2\2\u01fc"+
		"\u01f0\3\2\2\2\u01fc\u01f1\3\2\2\2\u01fc\u01f2\3\2\2\2\u01fc\u01f3\3\2"+
		"\2\2\u01fc\u01f4\3\2\2\2\u01fc\u01f5\3\2\2\2\u01fc\u01f6\3\2\2\2\u01fc"+
		"\u01f7\3\2\2\2\u01fc\u01f8\3\2\2\2\u01fc\u01f9\3\2\2\2\u01fc\u01fa\3\2"+
		"\2\2\u01fc\u01fb\3\2\2\2\u01fd\u0215\3\2\2\2\u01fe\u01ff\f\26\2\2\u01ff"+
		"\u0200\t\3\2\2\u0200\u0214\5V,\27\u0201\u0202\f\25\2\2\u0202\u0203\t\4"+
		"\2\2\u0203\u0214\5V,\26\u0204\u0205\f\24\2\2\u0205\u0206\t\5\2\2\u0206"+
		"\u0214\5V,\25\u0207\u0208\f\23\2\2\u0208\u0209\t\6\2\2\u0209\u0214\5V"+
		",\24\u020a\u020b\f\22\2\2\u020b\u020c\t\7\2\2\u020c\u0214\5V,\23\u020d"+
		"\u020e\f\21\2\2\u020e\u020f\7%\2\2\u020f\u0214\5V,\22\u0210\u0211\f\20"+
		"\2\2\u0211\u0212\7$\2\2\u0212\u0214\5V,\21\u0213\u01fe\3\2\2\2\u0213\u0201"+
		"\3\2\2\2\u0213\u0204\3\2\2\2\u0213\u0207\3\2\2\2\u0213\u020a\3\2\2\2\u0213"+
		"\u020d\3\2\2\2\u0213\u0210\3\2\2\2\u0214\u0217\3\2\2\2\u0215\u0213\3\2"+
		"\2\2\u0215\u0216\3\2\2\2\u0216W\3\2\2\2\u0217\u0215\3\2\2\2\u0218\u0219"+
		"\t\b\2\2\u0219Y\3\2\2\2\37_c|\u0089\u008e\u0095\u00c6\u00ce\u00d2\u00e2"+
		"\u00ef\u00f9\u0111\u012a\u0148\u0153\u015d\u0162\u0171\u017d\u018a\u0192"+
		"\u01a1\u01a4\u01d2\u01e6\u01fc\u0213\u0215";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
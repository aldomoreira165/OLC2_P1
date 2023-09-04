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
		IF=10, ELSE=11, WHILE=12, FOR=13, GUARD=14, IN=15, SWITCH=16, CASE=17, 
		DEFAULT=18, VAR=19, LET=20, BREAK=21, RETURN=22, CONTINUE=23, FUNC=24, 
		COUNT=25, ISEMPTY=26, APPEND=27, REMOVE_LAST=28, REMOVE=29, STRUCT=30, 
		STRUCT_VAR=31, STRUCT_LET=32, INOUT=33, AT=34, ST=35, NUMBER=36, STRING=37, 
		ID=38, DIF=39, IG_IG=40, NOT=41, OR=42, AND=43, IG=44, MAY_IG=45, MEN_IG=46, 
		MAYOR=47, MENOR=48, MUL=49, DIV=50, ADD=51, SUB=52, MOD=53, PARIZQ=54, 
		PARDER=55, LLAVEIZQ=56, LLAVEDER=57, CORCHIZQ=58, CORCHDER=59, DOSPUNTOS=60, 
		COMA=61, PTCOMA=62, INTERROGACION=63, PUNTO=64, GUIONBAJO=65, AMPERSON=66, 
		WHITESPACE=67, COMMENT=68, LINE_COMMENT=69;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_stmt = 2, RULE_defstructstmt = 3, RULE_lista_atributos = 4, 
		RULE_struct_expr = 5, RULE_valor_struct_expr = 6, RULE_l_dupla = 7, RULE_accesostructstmt = 8, 
		RULE_declvectorstmt = 9, RULE_defvectorstmt = 10, RULE_listaexpresiones = 11, 
		RULE_accesovectorstmt = 12, RULE_asignvectorstmt = 13, RULE_appendvectorstmt = 14, 
		RULE_removeatvectorstmt = 15, RULE_removelastvectorstmt = 16, RULE_countvectorstmt = 17, 
		RULE_isemptyvectorstmt = 18, RULE_declmatrizstmt = 19, RULE_tipomatriz = 20, 
		RULE_listavaloresmatriz = 21, RULE_listavaloresmatriz3 = 22, RULE_accesomatriz = 23, 
		RULE_asignmatrizstmt = 24, RULE_returnstmt = 25, RULE_breakstmt = 26, 
		RULE_continuestmt = 27, RULE_printstmt = 28, RULE_intstmt = 29, RULE_floatstmt = 30, 
		RULE_stringstmt = 31, RULE_funcdclstmt = 32, RULE_accfuncstm = 33, RULE_parametros = 34, 
		RULE_parametroscall = 35, RULE_ifstmt = 36, RULE_elseifstmt = 37, RULE_switchstmt = 38, 
		RULE_caseStmt = 39, RULE_defaultCase = 40, RULE_typedDeclstmt = 41, RULE_untypedDeclstmt = 42, 
		RULE_optionalTypedDeclstmt = 43, RULE_asignstmt = 44, RULE_whilestmt = 45, 
		RULE_forstmt = 46, RULE_guardstmt = 47, RULE_rangostmt = 48, RULE_opasignstmt = 49, 
		RULE_expr = 50, RULE_tipo = 51, RULE_tipo_vector = 52, RULE_tipo_matriz2 = 53, 
		RULE_tipo_matriz3 = 54;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "stmt", "defstructstmt", "lista_atributos", "struct_expr", 
			"valor_struct_expr", "l_dupla", "accesostructstmt", "declvectorstmt", 
			"defvectorstmt", "listaexpresiones", "accesovectorstmt", "asignvectorstmt", 
			"appendvectorstmt", "removeatvectorstmt", "removelastvectorstmt", "countvectorstmt", 
			"isemptyvectorstmt", "declmatrizstmt", "tipomatriz", "listavaloresmatriz", 
			"listavaloresmatriz3", "accesomatriz", "asignmatrizstmt", "returnstmt", 
			"breakstmt", "continuestmt", "printstmt", "intstmt", "floatstmt", "stringstmt", 
			"funcdclstmt", "accfuncstm", "parametros", "parametroscall", "ifstmt", 
			"elseifstmt", "switchstmt", "caseStmt", "defaultCase", "typedDeclstmt", 
			"untypedDeclstmt", "optionalTypedDeclstmt", "asignstmt", "whilestmt", 
			"forstmt", "guardstmt", "rangostmt", "opasignstmt", "expr", "tipo", "tipo_vector", 
			"tipo_matriz2", "tipo_matriz3"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'int'", "'float'", "'bool'", "'character'", "'String'", "'nil'", 
			"'true'", "'false'", "'print'", "'if'", "'else'", "'while'", "'for'", 
			"'guard'", "'in'", "'switch'", "'case'", "'default'", "'var'", "'let'", 
			"'break'", "'return'", "'continue'", "'func'", "'count'", "'IsEmpty'", 
			"'append'", "'removeLast'", "'remove'", "'struct'", "'varst'", "'letst'", 
			"'inout'", "'at'", "'st'", null, null, null, "'!='", "'=='", "'!'", "'||'", 
			"'&&'", "'='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", 
			"'%'", "'('", "')'", "'{'", "'}'", "'['", "']'", "':'", "','", "';'", 
			"'?'", "'.'", "'_'", "'&'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "NIL", "TRU", "FAL", 
			"PRINT", "IF", "ELSE", "WHILE", "FOR", "GUARD", "IN", "SWITCH", "CASE", 
			"DEFAULT", "VAR", "LET", "BREAK", "RETURN", "CONTINUE", "FUNC", "COUNT", 
			"ISEMPTY", "APPEND", "REMOVE_LAST", "REMOVE", "STRUCT", "STRUCT_VAR", 
			"STRUCT_LET", "INOUT", "AT", "ST", "NUMBER", "STRING", "ID", "DIF", "IG_IG", 
			"NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", 
			"DIV", "ADD", "SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", 
			"CORCHIZQ", "CORCHDER", "DOSPUNTOS", "COMA", "PTCOMA", "INTERROGACION", 
			"PUNTO", "GUIONBAJO", "AMPERSON", "WHITESPACE", "COMMENT", "LINE_COMMENT"
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
			setState(110);
			block();
			setState(111);
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
			setState(119);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PRINT) | (1L << IF) | (1L << WHILE) | (1L << FOR) | (1L << GUARD) | (1L << SWITCH) | (1L << VAR) | (1L << LET) | (1L << BREAK) | (1L << RETURN) | (1L << CONTINUE) | (1L << FUNC) | (1L << STRUCT) | (1L << ST) | (1L << ID))) != 0)) {
				{
				{
				setState(113);
				stmt();
				setState(115);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(114);
					match(PTCOMA);
					}
				}

				}
				}
				setState(121);
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
		public GuardstmtContext guardstmt() {
			return getRuleContext(GuardstmtContext.class,0);
		}
		public OpasignstmtContext opasignstmt() {
			return getRuleContext(OpasignstmtContext.class,0);
		}
		public FuncdclstmtContext funcdclstmt() {
			return getRuleContext(FuncdclstmtContext.class,0);
		}
		public AccfuncstmContext accfuncstm() {
			return getRuleContext(AccfuncstmContext.class,0);
		}
		public BreakstmtContext breakstmt() {
			return getRuleContext(BreakstmtContext.class,0);
		}
		public ContinuestmtContext continuestmt() {
			return getRuleContext(ContinuestmtContext.class,0);
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
		public DefstructstmtContext defstructstmt() {
			return getRuleContext(DefstructstmtContext.class,0);
		}
		public Struct_exprContext struct_expr() {
			return getRuleContext(Struct_exprContext.class,0);
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
			setState(148);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(122);
				printstmt();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(123);
				typedDeclstmt();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(124);
				untypedDeclstmt();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(125);
				optionalTypedDeclstmt();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(126);
				asignstmt();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(127);
				ifstmt();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(128);
				switchstmt();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(129);
				whilestmt();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(130);
				forstmt();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(131);
				guardstmt();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(132);
				opasignstmt();
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(133);
				funcdclstmt();
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(134);
				accfuncstm();
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(135);
				breakstmt();
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(136);
				continuestmt();
				}
				break;
			case 16:
				enterOuterAlt(_localctx, 16);
				{
				setState(137);
				returnstmt();
				}
				break;
			case 17:
				enterOuterAlt(_localctx, 17);
				{
				setState(138);
				declvectorstmt();
				}
				break;
			case 18:
				enterOuterAlt(_localctx, 18);
				{
				setState(139);
				accesovectorstmt();
				}
				break;
			case 19:
				enterOuterAlt(_localctx, 19);
				{
				setState(140);
				appendvectorstmt();
				}
				break;
			case 20:
				enterOuterAlt(_localctx, 20);
				{
				setState(141);
				removelastvectorstmt();
				}
				break;
			case 21:
				enterOuterAlt(_localctx, 21);
				{
				setState(142);
				removeatvectorstmt();
				}
				break;
			case 22:
				enterOuterAlt(_localctx, 22);
				{
				setState(143);
				asignvectorstmt();
				}
				break;
			case 23:
				enterOuterAlt(_localctx, 23);
				{
				setState(144);
				declmatrizstmt();
				}
				break;
			case 24:
				enterOuterAlt(_localctx, 24);
				{
				setState(145);
				asignmatrizstmt();
				}
				break;
			case 25:
				enterOuterAlt(_localctx, 25);
				{
				setState(146);
				defstructstmt();
				}
				break;
			case 26:
				enterOuterAlt(_localctx, 26);
				{
				setState(147);
				struct_expr();
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

	public static class DefstructstmtContext extends ParserRuleContext {
		public DefstructstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_defstructstmt; }
	 
		public DefstructstmtContext() { }
		public void copyFrom(DefstructstmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class DefStructContext extends DefstructstmtContext {
		public TerminalNode STRUCT() { return getToken(SwiftGrammarParser.STRUCT, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public List<Lista_atributosContext> lista_atributos() {
			return getRuleContexts(Lista_atributosContext.class);
		}
		public Lista_atributosContext lista_atributos(int i) {
			return getRuleContext(Lista_atributosContext.class,i);
		}
		public DefStructContext(DefstructstmtContext ctx) { copyFrom(ctx); }
	}

	public final DefstructstmtContext defstructstmt() throws RecognitionException {
		DefstructstmtContext _localctx = new DefstructstmtContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_defstructstmt);
		int _la;
		try {
			_localctx = new DefStructContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(150);
			match(STRUCT);
			setState(151);
			match(ID);
			setState(152);
			match(LLAVEIZQ);
			setState(156);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR || _la==LET) {
				{
				{
				setState(153);
				lista_atributos();
				}
				}
				setState(158);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(159);
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

	public static class Lista_atributosContext extends ParserRuleContext {
		public Lista_atributosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lista_atributos; }
	 
		public Lista_atributosContext() { }
		public void copyFrom(Lista_atributosContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class AtributoStructContext extends Lista_atributosContext {
		public List<TerminalNode> ID() { return getTokens(SwiftGrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(SwiftGrammarParser.ID, i);
		}
		public TerminalNode LET() { return getToken(SwiftGrammarParser.LET, 0); }
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TerminalNode DOSPUNTOS() { return getToken(SwiftGrammarParser.DOSPUNTOS, 0); }
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public AtributoStructContext(Lista_atributosContext ctx) { copyFrom(ctx); }
	}

	public final Lista_atributosContext lista_atributos() throws RecognitionException {
		Lista_atributosContext _localctx = new Lista_atributosContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_lista_atributos);
		int _la;
		try {
			_localctx = new AtributoStructContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(161);
			_la = _input.LA(1);
			if ( !(_la==VAR || _la==LET) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(162);
			match(ID);
			setState(168);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DOSPUNTOS) {
				{
				setState(163);
				match(DOSPUNTOS);
				setState(166);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case INT:
				case FLOAT:
				case BOOL:
				case CHARACTER:
				case PSTRING:
				case CORCHIZQ:
					{
					setState(164);
					tipo();
					}
					break;
				case ID:
					{
					setState(165);
					match(ID);
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
			}

			setState(172);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IG) {
				{
				setState(170);
				match(IG);
				setState(171);
				expr(0);
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

	public static class Struct_exprContext extends ParserRuleContext {
		public Struct_exprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_expr; }
	 
		public Struct_exprContext() { }
		public void copyFrom(Struct_exprContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class StructExprContext extends Struct_exprContext {
		public TerminalNode ST() { return getToken(SwiftGrammarParser.ST, 0); }
		public List<TerminalNode> ID() { return getTokens(SwiftGrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(SwiftGrammarParser.ID, i);
		}
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public Valor_struct_exprContext valor_struct_expr() {
			return getRuleContext(Valor_struct_exprContext.class,0);
		}
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TerminalNode LET() { return getToken(SwiftGrammarParser.LET, 0); }
		public TerminalNode DOSPUNTOS() { return getToken(SwiftGrammarParser.DOSPUNTOS, 0); }
		public StructExprContext(Struct_exprContext ctx) { copyFrom(ctx); }
	}

	public final Struct_exprContext struct_expr() throws RecognitionException {
		Struct_exprContext _localctx = new Struct_exprContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_struct_expr);
		int _la;
		try {
			_localctx = new StructExprContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(174);
			match(ST);
			setState(175);
			_la = _input.LA(1);
			if ( !(_la==VAR || _la==LET) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(176);
			match(ID);
			setState(179);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DOSPUNTOS) {
				{
				setState(177);
				match(DOSPUNTOS);
				setState(178);
				match(ID);
				}
			}

			setState(181);
			match(IG);
			setState(182);
			valor_struct_expr();
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

	public static class Valor_struct_exprContext extends ParserRuleContext {
		public Valor_struct_exprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_valor_struct_expr; }
	 
		public Valor_struct_exprContext() { }
		public void copyFrom(Valor_struct_exprContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class ValorStructExprContext extends Valor_struct_exprContext {
		public TerminalNode ST() { return getToken(SwiftGrammarParser.ST, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public L_duplaContext l_dupla() {
			return getRuleContext(L_duplaContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public ValorStructExprContext(Valor_struct_exprContext ctx) { copyFrom(ctx); }
	}

	public final Valor_struct_exprContext valor_struct_expr() throws RecognitionException {
		Valor_struct_exprContext _localctx = new Valor_struct_exprContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_valor_struct_expr);
		try {
			_localctx = new ValorStructExprContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(184);
			match(ST);
			setState(185);
			match(ID);
			setState(190);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				{
				setState(186);
				match(PARIZQ);
				setState(187);
				l_dupla();
				setState(188);
				match(PARDER);
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

	public static class L_duplaContext extends ParserRuleContext {
		public L_duplaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_l_dupla; }
	 
		public L_duplaContext() { }
		public void copyFrom(L_duplaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class DuplastructContext extends L_duplaContext {
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
		public DuplastructContext(L_duplaContext ctx) { copyFrom(ctx); }
	}

	public final L_duplaContext l_dupla() throws RecognitionException {
		L_duplaContext _localctx = new L_duplaContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_l_dupla);
		int _la;
		try {
			_localctx = new DuplastructContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(192);
			match(ID);
			setState(193);
			match(DOSPUNTOS);
			setState(194);
			expr(0);
			setState(201);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(195);
				match(COMA);
				setState(196);
				match(ID);
				setState(197);
				match(DOSPUNTOS);
				setState(198);
				expr(0);
				}
				}
				setState(203);
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

	public static class AccesostructstmtContext extends ParserRuleContext {
		public AccesostructstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accesostructstmt; }
	 
		public AccesostructstmtContext() { }
		public void copyFrom(AccesostructstmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class AccesoStructContext extends AccesostructstmtContext {
		public List<TerminalNode> ID() { return getTokens(SwiftGrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(SwiftGrammarParser.ID, i);
		}
		public List<TerminalNode> PUNTO() { return getTokens(SwiftGrammarParser.PUNTO); }
		public TerminalNode PUNTO(int i) {
			return getToken(SwiftGrammarParser.PUNTO, i);
		}
		public AccesoStructContext(AccesostructstmtContext ctx) { copyFrom(ctx); }
	}

	public final AccesostructstmtContext accesostructstmt() throws RecognitionException {
		AccesostructstmtContext _localctx = new AccesostructstmtContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_accesostructstmt);
		try {
			int _alt;
			_localctx = new AccesoStructContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(204);
			match(ID);
			setState(207); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(205);
					match(PUNTO);
					setState(206);
					match(ID);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(209); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
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
		enterRule(_localctx, 18, RULE_declvectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(211);
			match(VAR);
			setState(212);
			match(ID);
			setState(213);
			match(DOSPUNTOS);
			setState(214);
			match(CORCHIZQ);
			setState(215);
			tipo();
			setState(216);
			match(CORCHDER);
			setState(217);
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
		enterRule(_localctx, 20, RULE_defvectorstmt);
		int _la;
		try {
			setState(227);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				_localctx = new DefVectorContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(219);
				match(IG);
				setState(220);
				match(CORCHIZQ);
				setState(222);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << INT) | (1L << FLOAT) | (1L << PSTRING) | (1L << NIL) | (1L << TRU) | (1L << FAL) | (1L << ST) | (1L << NUMBER) | (1L << STRING) | (1L << ID) | (1L << NOT) | (1L << SUB) | (1L << PARIZQ))) != 0)) {
					{
					setState(221);
					listaexpresiones();
					}
				}

				setState(224);
				match(CORCHDER);
				}
				break;
			case 2:
				_localctx = new DefVectorIDContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(225);
				match(IG);
				setState(226);
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
		enterRule(_localctx, 22, RULE_listaexpresiones);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(229);
			expr(0);
			setState(234);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(230);
				match(COMA);
				setState(231);
				expr(0);
				}
				}
				setState(236);
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
		enterRule(_localctx, 24, RULE_accesovectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(237);
			match(ID);
			setState(238);
			match(CORCHIZQ);
			setState(239);
			expr(0);
			setState(240);
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
		enterRule(_localctx, 26, RULE_asignvectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(242);
			match(ID);
			setState(243);
			match(CORCHIZQ);
			setState(244);
			expr(0);
			setState(245);
			match(CORCHDER);
			setState(246);
			match(IG);
			setState(247);
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
		enterRule(_localctx, 28, RULE_appendvectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(249);
			match(ID);
			setState(250);
			match(PUNTO);
			setState(251);
			match(APPEND);
			setState(252);
			match(PARIZQ);
			setState(253);
			expr(0);
			setState(254);
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
		enterRule(_localctx, 30, RULE_removeatvectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(256);
			match(ID);
			setState(257);
			match(PUNTO);
			setState(258);
			match(REMOVE);
			setState(259);
			match(PARIZQ);
			setState(260);
			match(AT);
			setState(261);
			match(DOSPUNTOS);
			setState(262);
			expr(0);
			setState(263);
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
		enterRule(_localctx, 32, RULE_removelastvectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(265);
			match(ID);
			setState(266);
			match(PUNTO);
			setState(267);
			match(REMOVE_LAST);
			setState(268);
			match(PARIZQ);
			setState(269);
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
		enterRule(_localctx, 34, RULE_countvectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(271);
			match(ID);
			setState(272);
			match(PUNTO);
			setState(273);
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
		enterRule(_localctx, 36, RULE_isemptyvectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(275);
			match(ID);
			setState(276);
			match(PUNTO);
			setState(277);
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
		enterRule(_localctx, 38, RULE_declmatrizstmt);
		int _la;
		try {
			setState(295);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				_localctx = new Declmatrizstmt2Context(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(279);
				match(VAR);
				setState(280);
				match(ID);
				setState(283);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DOSPUNTOS) {
					{
					setState(281);
					match(DOSPUNTOS);
					setState(282);
					tipomatriz();
					}
				}

				setState(285);
				match(IG);
				setState(286);
				listavaloresmatriz();
				}
				break;
			case 2:
				_localctx = new Declmatrizstmt3Context(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(287);
				match(VAR);
				setState(288);
				match(ID);
				setState(291);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DOSPUNTOS) {
					{
					setState(289);
					match(DOSPUNTOS);
					setState(290);
					tipomatriz();
					}
				}

				setState(293);
				match(IG);
				setState(294);
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
		enterRule(_localctx, 40, RULE_tipomatriz);
		try {
			setState(311);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				_localctx = new Tipomatriz3Context(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(297);
				match(CORCHIZQ);
				setState(298);
				match(CORCHIZQ);
				setState(299);
				match(CORCHIZQ);
				setState(300);
				tipo();
				setState(301);
				match(CORCHDER);
				setState(302);
				match(CORCHDER);
				setState(303);
				match(CORCHDER);
				}
				break;
			case 2:
				_localctx = new Tipomatriz2Context(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(305);
				match(CORCHIZQ);
				setState(306);
				match(CORCHIZQ);
				setState(307);
				tipo();
				setState(308);
				match(CORCHDER);
				setState(309);
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
		enterRule(_localctx, 42, RULE_listavaloresmatriz);
		int _la;
		try {
			_localctx = new Listavaloresmatriz2Context(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(313);
			match(CORCHIZQ);
			setState(314);
			match(CORCHIZQ);
			setState(315);
			listaexpresiones();
			setState(316);
			match(CORCHDER);
			setState(322); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(317);
				match(COMA);
				setState(318);
				match(CORCHIZQ);
				setState(319);
				listaexpresiones();
				setState(320);
				match(CORCHDER);
				}
				}
				setState(324); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==COMA );
			setState(326);
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
		enterRule(_localctx, 44, RULE_listavaloresmatriz3);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(328);
			match(CORCHIZQ);
			setState(329);
			listavaloresmatriz();
			setState(334);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(330);
				match(COMA);
				setState(331);
				listavaloresmatriz();
				}
				}
				setState(336);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(337);
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
		enterRule(_localctx, 46, RULE_accesomatriz);
		try {
			setState(358);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				_localctx = new Accesomatriz2Context(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(339);
				match(ID);
				setState(340);
				match(CORCHIZQ);
				setState(341);
				expr(0);
				setState(342);
				match(CORCHDER);
				setState(343);
				match(CORCHIZQ);
				setState(344);
				expr(0);
				setState(345);
				match(CORCHDER);
				}
				break;
			case 2:
				_localctx = new Accesomatriz3Context(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(347);
				match(ID);
				setState(348);
				match(CORCHIZQ);
				setState(349);
				expr(0);
				setState(350);
				match(CORCHDER);
				setState(351);
				match(CORCHIZQ);
				setState(352);
				expr(0);
				setState(353);
				match(CORCHDER);
				setState(354);
				match(CORCHIZQ);
				setState(355);
				expr(0);
				setState(356);
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
		enterRule(_localctx, 48, RULE_asignmatrizstmt);
		try {
			setState(383);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				_localctx = new Asignmatrizstmt2Context(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(360);
				match(ID);
				setState(361);
				match(CORCHIZQ);
				setState(362);
				expr(0);
				setState(363);
				match(CORCHDER);
				setState(364);
				match(CORCHIZQ);
				setState(365);
				expr(0);
				setState(366);
				match(CORCHDER);
				setState(367);
				match(IG);
				setState(368);
				expr(0);
				}
				break;
			case 2:
				_localctx = new Asignmatrizstmt3Context(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(370);
				match(ID);
				setState(371);
				match(CORCHIZQ);
				setState(372);
				expr(0);
				setState(373);
				match(CORCHDER);
				setState(374);
				match(CORCHIZQ);
				setState(375);
				expr(0);
				setState(376);
				match(CORCHDER);
				setState(377);
				match(CORCHIZQ);
				setState(378);
				expr(0);
				setState(379);
				match(CORCHDER);
				setState(380);
				match(IG);
				setState(381);
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
		enterRule(_localctx, 50, RULE_returnstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(385);
			match(RETURN);
			setState(386);
			expr(0);
			setState(387);
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

	public static class BreakstmtContext extends ParserRuleContext {
		public TerminalNode BREAK() { return getToken(SwiftGrammarParser.BREAK, 0); }
		public BreakstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_breakstmt; }
	}

	public final BreakstmtContext breakstmt() throws RecognitionException {
		BreakstmtContext _localctx = new BreakstmtContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_breakstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(389);
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

	public static class ContinuestmtContext extends ParserRuleContext {
		public TerminalNode CONTINUE() { return getToken(SwiftGrammarParser.CONTINUE, 0); }
		public ContinuestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_continuestmt; }
	}

	public final ContinuestmtContext continuestmt() throws RecognitionException {
		ContinuestmtContext _localctx = new ContinuestmtContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_continuestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(391);
			match(CONTINUE);
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
		enterRule(_localctx, 56, RULE_printstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(393);
			match(PRINT);
			setState(394);
			match(PARIZQ);
			setState(395);
			expr(0);
			setState(396);
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
		enterRule(_localctx, 58, RULE_intstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(398);
			match(INT);
			setState(399);
			match(PARIZQ);
			setState(400);
			expr(0);
			setState(401);
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
		enterRule(_localctx, 60, RULE_floatstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(403);
			match(FLOAT);
			setState(404);
			match(PARIZQ);
			setState(405);
			expr(0);
			setState(406);
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
		enterRule(_localctx, 62, RULE_stringstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(408);
			match(PSTRING);
			setState(409);
			match(PARIZQ);
			setState(410);
			expr(0);
			setState(411);
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
		enterRule(_localctx, 64, RULE_funcdclstmt);
		int _la;
		try {
			setState(438);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				_localctx = new FuncionNormalContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(413);
				match(FUNC);
				setState(414);
				match(ID);
				setState(415);
				match(PARIZQ);
				setState(417);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ID) {
					{
					setState(416);
					parametros();
					}
				}

				setState(419);
				match(PARDER);
				setState(420);
				match(LLAVEIZQ);
				setState(421);
				block();
				setState(422);
				match(LLAVEDER);
				}
				break;
			case 2:
				_localctx = new FuncionRetornoContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(424);
				match(FUNC);
				setState(425);
				match(ID);
				setState(426);
				match(PARIZQ);
				setState(428);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ID) {
					{
					setState(427);
					parametros();
					}
				}

				setState(430);
				match(PARDER);
				setState(431);
				match(SUB);
				setState(432);
				match(MAYOR);
				setState(433);
				tipo();
				setState(434);
				match(LLAVEIZQ);
				setState(435);
				block();
				setState(436);
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
		enterRule(_localctx, 66, RULE_accfuncstm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(440);
			match(ID);
			setState(441);
			match(PARIZQ);
			setState(443);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(442);
				parametroscall();
				}
			}

			setState(445);
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
		public List<TerminalNode> INOUT() { return getTokens(SwiftGrammarParser.INOUT); }
		public TerminalNode INOUT(int i) {
			return getToken(SwiftGrammarParser.INOUT, i);
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
		enterRule(_localctx, 68, RULE_parametros);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(447);
			match(ID);
			setState(448);
			match(ID);
			setState(449);
			match(DOSPUNTOS);
			setState(451);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==INOUT) {
				{
				setState(450);
				match(INOUT);
				}
			}

			setState(453);
			tipo();
			setState(464);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(454);
				match(COMA);
				setState(455);
				match(ID);
				setState(456);
				match(ID);
				setState(457);
				match(DOSPUNTOS);
				setState(459);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==INOUT) {
					{
					setState(458);
					match(INOUT);
					}
				}

				setState(461);
				tipo();
				}
				}
				setState(466);
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
		public List<TerminalNode> AMPERSON() { return getTokens(SwiftGrammarParser.AMPERSON); }
		public TerminalNode AMPERSON(int i) {
			return getToken(SwiftGrammarParser.AMPERSON, i);
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
		enterRule(_localctx, 70, RULE_parametroscall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(467);
			match(ID);
			setState(468);
			match(DOSPUNTOS);
			setState(470);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AMPERSON) {
				{
				setState(469);
				match(AMPERSON);
				}
			}

			setState(472);
			expr(0);
			setState(482);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(473);
				match(COMA);
				setState(474);
				match(ID);
				setState(475);
				match(DOSPUNTOS);
				setState(477);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AMPERSON) {
					{
					setState(476);
					match(AMPERSON);
					}
				}

				setState(479);
				expr(0);
				}
				}
				setState(484);
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
		enterRule(_localctx, 72, RULE_ifstmt);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(485);
			match(IF);
			setState(486);
			expr(0);
			setState(487);
			match(LLAVEIZQ);
			setState(488);
			block();
			setState(489);
			match(LLAVEDER);
			setState(493);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,32,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(490);
					elseifstmt();
					}
					} 
				}
				setState(495);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,32,_ctx);
			}
			setState(501);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(496);
				match(ELSE);
				setState(497);
				match(LLAVEIZQ);
				setState(498);
				block();
				setState(499);
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
		enterRule(_localctx, 74, RULE_elseifstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(503);
			match(ELSE);
			setState(504);
			match(IF);
			setState(505);
			expr(0);
			setState(506);
			match(LLAVEIZQ);
			setState(507);
			block();
			setState(508);
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
		enterRule(_localctx, 76, RULE_switchstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(510);
			match(SWITCH);
			setState(511);
			expr(0);
			setState(512);
			match(LLAVEIZQ);
			setState(514); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(513);
				caseStmt();
				}
				}
				setState(516); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==CASE );
			setState(519);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(518);
				defaultCase();
				}
			}

			setState(521);
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
		enterRule(_localctx, 78, RULE_caseStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(523);
			match(CASE);
			setState(524);
			expr(0);
			setState(525);
			match(DOSPUNTOS);
			setState(526);
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
		enterRule(_localctx, 80, RULE_defaultCase);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(528);
			match(DEFAULT);
			setState(529);
			match(DOSPUNTOS);
			setState(530);
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
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode DOSPUNTOS() { return getToken(SwiftGrammarParser.DOSPUNTOS, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TerminalNode LET() { return getToken(SwiftGrammarParser.LET, 0); }
		public TypedDeclstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typedDeclstmt; }
	}

	public final TypedDeclstmtContext typedDeclstmt() throws RecognitionException {
		TypedDeclstmtContext _localctx = new TypedDeclstmtContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_typedDeclstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(532);
			_la = _input.LA(1);
			if ( !(_la==VAR || _la==LET) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(533);
			match(ID);
			setState(534);
			match(DOSPUNTOS);
			setState(535);
			tipo();
			setState(536);
			match(IG);
			setState(537);
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
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TerminalNode LET() { return getToken(SwiftGrammarParser.LET, 0); }
		public UntypedDeclstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_untypedDeclstmt; }
	}

	public final UntypedDeclstmtContext untypedDeclstmt() throws RecognitionException {
		UntypedDeclstmtContext _localctx = new UntypedDeclstmtContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_untypedDeclstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(539);
			_la = _input.LA(1);
			if ( !(_la==VAR || _la==LET) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(540);
			match(ID);
			setState(541);
			match(IG);
			setState(542);
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
		enterRule(_localctx, 86, RULE_optionalTypedDeclstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(544);
			match(VAR);
			setState(545);
			match(ID);
			setState(546);
			match(DOSPUNTOS);
			setState(547);
			tipo();
			setState(548);
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
		enterRule(_localctx, 88, RULE_asignstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(550);
			match(ID);
			setState(551);
			match(IG);
			setState(552);
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
		enterRule(_localctx, 90, RULE_whilestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(554);
			match(WHILE);
			setState(555);
			expr(0);
			setState(556);
			match(LLAVEIZQ);
			setState(557);
			block();
			setState(558);
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
		enterRule(_localctx, 92, RULE_forstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(560);
			match(FOR);
			setState(561);
			match(ID);
			setState(562);
			match(IN);
			setState(565);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,36,_ctx) ) {
			case 1:
				{
				setState(563);
				expr(0);
				}
				break;
			case 2:
				{
				setState(564);
				rangostmt();
				}
				break;
			}
			setState(567);
			match(LLAVEIZQ);
			setState(568);
			block();
			setState(569);
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

	public static class GuardstmtContext extends ParserRuleContext {
		public TerminalNode GUARD() { return getToken(SwiftGrammarParser.GUARD, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(SwiftGrammarParser.ELSE, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public GuardstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_guardstmt; }
	}

	public final GuardstmtContext guardstmt() throws RecognitionException {
		GuardstmtContext _localctx = new GuardstmtContext(_ctx, getState());
		enterRule(_localctx, 94, RULE_guardstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(571);
			match(GUARD);
			setState(572);
			expr(0);
			setState(573);
			match(ELSE);
			setState(574);
			match(LLAVEIZQ);
			setState(575);
			block();
			setState(576);
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
		enterRule(_localctx, 96, RULE_rangostmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(578);
			expr(0);
			setState(579);
			match(PUNTO);
			setState(580);
			match(PUNTO);
			setState(581);
			match(PUNTO);
			setState(582);
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
		enterRule(_localctx, 98, RULE_opasignstmt);
		try {
			setState(592);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,37,_ctx) ) {
			case 1:
				_localctx = new IncrementoContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(584);
				match(ID);
				setState(585);
				match(ADD);
				setState(586);
				match(IG);
				setState(587);
				expr(0);
				}
				break;
			case 2:
				_localctx = new DecrementoContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(588);
				match(ID);
				setState(589);
				match(SUB);
				setState(590);
				match(IG);
				setState(591);
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
		public TerminalNode DIV() { return getToken(SwiftGrammarParser.DIV, 0); }
		public TerminalNode MOD() { return getToken(SwiftGrammarParser.MOD, 0); }
		public TerminalNode MUL() { return getToken(SwiftGrammarParser.MUL, 0); }
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
	public static class AccesoValorStructExprContext extends ExprContext {
		public Valor_struct_exprContext valor_struct_expr() {
			return getRuleContext(Valor_struct_exprContext.class,0);
		}
		public AccesoValorStructExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class NotExprContext extends ExprContext {
		public TerminalNode NOT() { return getToken(SwiftGrammarParser.NOT, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public NotExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class IntExprContext extends ExprContext {
		public IntstmtContext intstmt() {
			return getRuleContext(IntstmtContext.class,0);
		}
		public IntExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class AccesoStructExprContext extends ExprContext {
		public AccesostructstmtContext accesostructstmt() {
			return getRuleContext(AccesostructstmtContext.class,0);
		}
		public AccesoStructExprContext(ExprContext ctx) { copyFrom(ctx); }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 100;
		enterRecursionRule(_localctx, 100, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(618);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,38,_ctx) ) {
			case 1:
				{
				_localctx = new ParExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(595);
				match(PARIZQ);
				setState(596);
				expr(0);
				setState(597);
				match(PARDER);
				}
				break;
			case 2:
				{
				_localctx = new UnariaExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(599);
				match(SUB);
				setState(600);
				expr(24);
				}
				break;
			case 3:
				{
				_localctx = new NotExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(601);
				match(NOT);
				setState(602);
				expr(23);
				}
				break;
			case 4:
				{
				_localctx = new NumExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(603);
				match(NUMBER);
				}
				break;
			case 5:
				{
				_localctx = new IdExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(604);
				match(ID);
				}
				break;
			case 6:
				{
				_localctx = new StrExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(605);
				match(STRING);
				}
				break;
			case 7:
				{
				_localctx = new BoolExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(606);
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
			case 8:
				{
				_localctx = new NilExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(607);
				match(NIL);
				}
				break;
			case 9:
				{
				_localctx = new AccFuncExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(608);
				accfuncstm();
				}
				break;
			case 10:
				{
				_localctx = new IntExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(609);
				intstmt();
				}
				break;
			case 11:
				{
				_localctx = new FloatExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(610);
				floatstmt();
				}
				break;
			case 12:
				{
				_localctx = new StringExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(611);
				stringstmt();
				}
				break;
			case 13:
				{
				_localctx = new AccesoVectorExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(612);
				accesovectorstmt();
				}
				break;
			case 14:
				{
				_localctx = new CountVectorExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(613);
				countvectorstmt();
				}
				break;
			case 15:
				{
				_localctx = new IsEmptyVectorExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(614);
				isemptyvectorstmt();
				}
				break;
			case 16:
				{
				_localctx = new AccesoMatrizExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(615);
				accesomatriz();
				}
				break;
			case 17:
				{
				_localctx = new AccesoStructExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(616);
				accesostructstmt();
				}
				break;
			case 18:
				{
				_localctx = new AccesoValorStructExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(617);
				valor_struct_expr();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(643);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,40,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(641);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,39,_ctx) ) {
					case 1:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(620);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(621);
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
						setState(622);
						((OpExprContext)_localctx).right = expr(23);
						}
						break;
					case 2:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(623);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(624);
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
						setState(625);
						((OpExprContext)_localctx).right = expr(22);
						}
						break;
					case 3:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(626);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(627);
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
						setState(628);
						((OpExprContext)_localctx).right = expr(21);
						}
						break;
					case 4:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(629);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(630);
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
						setState(631);
						((OpExprContext)_localctx).right = expr(20);
						}
						break;
					case 5:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(632);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(633);
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
						setState(634);
						((OpExprContext)_localctx).right = expr(19);
						}
						break;
					case 6:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(635);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(636);
						((OpExprContext)_localctx).op = match(AND);
						setState(637);
						((OpExprContext)_localctx).right = expr(18);
						}
						break;
					case 7:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(638);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(639);
						((OpExprContext)_localctx).op = match(OR);
						setState(640);
						((OpExprContext)_localctx).right = expr(17);
						}
						break;
					}
					} 
				}
				setState(645);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,40,_ctx);
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
		public Tipo_vectorContext tipo_vector() {
			return getRuleContext(Tipo_vectorContext.class,0);
		}
		public Tipo_matriz2Context tipo_matriz2() {
			return getRuleContext(Tipo_matriz2Context.class,0);
		}
		public Tipo_matriz3Context tipo_matriz3() {
			return getRuleContext(Tipo_matriz3Context.class,0);
		}
		public TipoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo; }
	}

	public final TipoContext tipo() throws RecognitionException {
		TipoContext _localctx = new TipoContext(_ctx, getState());
		enterRule(_localctx, 102, RULE_tipo);
		try {
			setState(654);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,41,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(646);
				match(INT);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(647);
				match(FLOAT);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(648);
				match(BOOL);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(649);
				match(CHARACTER);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(650);
				match(PSTRING);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(651);
				tipo_vector();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(652);
				tipo_matriz2();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(653);
				tipo_matriz3();
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

	public static class Tipo_vectorContext extends ParserRuleContext {
		public TerminalNode CORCHIZQ() { return getToken(SwiftGrammarParser.CORCHIZQ, 0); }
		public TerminalNode INT() { return getToken(SwiftGrammarParser.INT, 0); }
		public TerminalNode CORCHDER() { return getToken(SwiftGrammarParser.CORCHDER, 0); }
		public TerminalNode FLOAT() { return getToken(SwiftGrammarParser.FLOAT, 0); }
		public TerminalNode BOOL() { return getToken(SwiftGrammarParser.BOOL, 0); }
		public TerminalNode CHARACTER() { return getToken(SwiftGrammarParser.CHARACTER, 0); }
		public TerminalNode PSTRING() { return getToken(SwiftGrammarParser.PSTRING, 0); }
		public Tipo_vectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo_vector; }
	}

	public final Tipo_vectorContext tipo_vector() throws RecognitionException {
		Tipo_vectorContext _localctx = new Tipo_vectorContext(_ctx, getState());
		enterRule(_localctx, 104, RULE_tipo_vector);
		try {
			setState(671);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,42,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(656);
				match(CORCHIZQ);
				setState(657);
				match(INT);
				setState(658);
				match(CORCHDER);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(659);
				match(CORCHIZQ);
				setState(660);
				match(FLOAT);
				setState(661);
				match(CORCHDER);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(662);
				match(CORCHIZQ);
				setState(663);
				match(BOOL);
				setState(664);
				match(CORCHDER);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(665);
				match(CORCHIZQ);
				setState(666);
				match(CHARACTER);
				setState(667);
				match(CORCHDER);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(668);
				match(CORCHIZQ);
				setState(669);
				match(PSTRING);
				setState(670);
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

	public static class Tipo_matriz2Context extends ParserRuleContext {
		public List<TerminalNode> CORCHIZQ() { return getTokens(SwiftGrammarParser.CORCHIZQ); }
		public TerminalNode CORCHIZQ(int i) {
			return getToken(SwiftGrammarParser.CORCHIZQ, i);
		}
		public TerminalNode INT() { return getToken(SwiftGrammarParser.INT, 0); }
		public List<TerminalNode> CORCHDER() { return getTokens(SwiftGrammarParser.CORCHDER); }
		public TerminalNode CORCHDER(int i) {
			return getToken(SwiftGrammarParser.CORCHDER, i);
		}
		public TerminalNode FLOAT() { return getToken(SwiftGrammarParser.FLOAT, 0); }
		public TerminalNode BOOL() { return getToken(SwiftGrammarParser.BOOL, 0); }
		public TerminalNode CHARACTER() { return getToken(SwiftGrammarParser.CHARACTER, 0); }
		public TerminalNode PSTRING() { return getToken(SwiftGrammarParser.PSTRING, 0); }
		public Tipo_matriz2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo_matriz2; }
	}

	public final Tipo_matriz2Context tipo_matriz2() throws RecognitionException {
		Tipo_matriz2Context _localctx = new Tipo_matriz2Context(_ctx, getState());
		enterRule(_localctx, 106, RULE_tipo_matriz2);
		try {
			setState(698);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,43,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(673);
				match(CORCHIZQ);
				setState(674);
				match(CORCHIZQ);
				setState(675);
				match(INT);
				setState(676);
				match(CORCHDER);
				setState(677);
				match(CORCHDER);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(678);
				match(CORCHIZQ);
				setState(679);
				match(CORCHIZQ);
				setState(680);
				match(FLOAT);
				setState(681);
				match(CORCHDER);
				setState(682);
				match(CORCHDER);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(683);
				match(CORCHIZQ);
				setState(684);
				match(CORCHIZQ);
				setState(685);
				match(BOOL);
				setState(686);
				match(CORCHDER);
				setState(687);
				match(CORCHDER);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(688);
				match(CORCHIZQ);
				setState(689);
				match(CORCHIZQ);
				setState(690);
				match(CHARACTER);
				setState(691);
				match(CORCHDER);
				setState(692);
				match(CORCHDER);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(693);
				match(CORCHIZQ);
				setState(694);
				match(CORCHIZQ);
				setState(695);
				match(PSTRING);
				setState(696);
				match(CORCHDER);
				setState(697);
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

	public static class Tipo_matriz3Context extends ParserRuleContext {
		public List<TerminalNode> CORCHIZQ() { return getTokens(SwiftGrammarParser.CORCHIZQ); }
		public TerminalNode CORCHIZQ(int i) {
			return getToken(SwiftGrammarParser.CORCHIZQ, i);
		}
		public TerminalNode INT() { return getToken(SwiftGrammarParser.INT, 0); }
		public List<TerminalNode> CORCHDER() { return getTokens(SwiftGrammarParser.CORCHDER); }
		public TerminalNode CORCHDER(int i) {
			return getToken(SwiftGrammarParser.CORCHDER, i);
		}
		public TerminalNode FLOAT() { return getToken(SwiftGrammarParser.FLOAT, 0); }
		public TerminalNode BOOL() { return getToken(SwiftGrammarParser.BOOL, 0); }
		public TerminalNode CHARACTER() { return getToken(SwiftGrammarParser.CHARACTER, 0); }
		public TerminalNode PSTRING() { return getToken(SwiftGrammarParser.PSTRING, 0); }
		public Tipo_matriz3Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo_matriz3; }
	}

	public final Tipo_matriz3Context tipo_matriz3() throws RecognitionException {
		Tipo_matriz3Context _localctx = new Tipo_matriz3Context(_ctx, getState());
		enterRule(_localctx, 108, RULE_tipo_matriz3);
		try {
			setState(735);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,44,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(700);
				match(CORCHIZQ);
				setState(701);
				match(CORCHIZQ);
				setState(702);
				match(CORCHIZQ);
				setState(703);
				match(INT);
				setState(704);
				match(CORCHDER);
				setState(705);
				match(CORCHDER);
				setState(706);
				match(CORCHDER);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(707);
				match(CORCHIZQ);
				setState(708);
				match(CORCHIZQ);
				setState(709);
				match(CORCHIZQ);
				setState(710);
				match(FLOAT);
				setState(711);
				match(CORCHDER);
				setState(712);
				match(CORCHDER);
				setState(713);
				match(CORCHDER);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(714);
				match(CORCHIZQ);
				setState(715);
				match(CORCHIZQ);
				setState(716);
				match(CORCHIZQ);
				setState(717);
				match(BOOL);
				setState(718);
				match(CORCHDER);
				setState(719);
				match(CORCHDER);
				setState(720);
				match(CORCHDER);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(721);
				match(CORCHIZQ);
				setState(722);
				match(CORCHIZQ);
				setState(723);
				match(CORCHIZQ);
				setState(724);
				match(CHARACTER);
				setState(725);
				match(CORCHDER);
				setState(726);
				match(CORCHDER);
				setState(727);
				match(CORCHDER);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(728);
				match(CORCHIZQ);
				setState(729);
				match(CORCHIZQ);
				setState(730);
				match(CORCHIZQ);
				setState(731);
				match(PSTRING);
				setState(732);
				match(CORCHDER);
				setState(733);
				match(CORCHDER);
				setState(734);
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 50:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 22);
		case 1:
			return precpred(_ctx, 21);
		case 2:
			return precpred(_ctx, 20);
		case 3:
			return precpred(_ctx, 19);
		case 4:
			return precpred(_ctx, 18);
		case 5:
			return precpred(_ctx, 17);
		case 6:
			return precpred(_ctx, 16);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3G\u02e4\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\3\2\3\2\3\2\3\3\3\3\5\3v\n\3\7"+
		"\3x\n\3\f\3\16\3{\13\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4\u0097\n\4"+
		"\3\5\3\5\3\5\3\5\7\5\u009d\n\5\f\5\16\5\u00a0\13\5\3\5\3\5\3\6\3\6\3\6"+
		"\3\6\3\6\5\6\u00a9\n\6\5\6\u00ab\n\6\3\6\3\6\5\6\u00af\n\6\3\7\3\7\3\7"+
		"\3\7\3\7\5\7\u00b6\n\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\5\b\u00c1\n"+
		"\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\7\t\u00ca\n\t\f\t\16\t\u00cd\13\t\3\n\3"+
		"\n\3\n\6\n\u00d2\n\n\r\n\16\n\u00d3\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\f\3\f\3\f\5\f\u00e1\n\f\3\f\3\f\3\f\5\f\u00e6\n\f\3\r\3\r\3\r"+
		"\7\r\u00eb\n\r\f\r\16\r\u00ee\13\r\3\16\3\16\3\16\3\16\3\16\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\23"+
		"\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\5\25\u011e\n\25"+
		"\3\25\3\25\3\25\3\25\3\25\3\25\5\25\u0126\n\25\3\25\3\25\5\25\u012a\n"+
		"\25\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3"+
		"\26\5\26\u013a\n\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\6\27"+
		"\u0145\n\27\r\27\16\27\u0146\3\27\3\27\3\30\3\30\3\30\3\30\7\30\u014f"+
		"\n\30\f\30\16\30\u0152\13\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3"+
		"\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\5\31\u0169"+
		"\n\31\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\5\32\u0182\n\32\3\33"+
		"\3\33\3\33\3\33\3\34\3\34\3\35\3\35\3\36\3\36\3\36\3\36\3\36\3\37\3\37"+
		"\3\37\3\37\3\37\3 \3 \3 \3 \3 \3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\5\"\u01a4"+
		"\n\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\5\"\u01af\n\"\3\"\3\"\3\"\3\""+
		"\3\"\3\"\3\"\3\"\5\"\u01b9\n\"\3#\3#\3#\5#\u01be\n#\3#\3#\3$\3$\3$\3$"+
		"\5$\u01c6\n$\3$\3$\3$\3$\3$\3$\5$\u01ce\n$\3$\7$\u01d1\n$\f$\16$\u01d4"+
		"\13$\3%\3%\3%\5%\u01d9\n%\3%\3%\3%\3%\3%\5%\u01e0\n%\3%\7%\u01e3\n%\f"+
		"%\16%\u01e6\13%\3&\3&\3&\3&\3&\3&\7&\u01ee\n&\f&\16&\u01f1\13&\3&\3&\3"+
		"&\3&\3&\5&\u01f8\n&\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3(\3(\3(\3(\6(\u0205\n"+
		"(\r(\16(\u0206\3(\5(\u020a\n(\3(\3(\3)\3)\3)\3)\3)\3*\3*\3*\3*\3+\3+\3"+
		"+\3+\3+\3+\3+\3,\3,\3,\3,\3,\3-\3-\3-\3-\3-\3-\3.\3.\3.\3.\3/\3/\3/\3"+
		"/\3/\3/\3\60\3\60\3\60\3\60\3\60\5\60\u0238\n\60\3\60\3\60\3\60\3\60\3"+
		"\61\3\61\3\61\3\61\3\61\3\61\3\61\3\62\3\62\3\62\3\62\3\62\3\62\3\63\3"+
		"\63\3\63\3\63\3\63\3\63\3\63\3\63\5\63\u0253\n\63\3\64\3\64\3\64\3\64"+
		"\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64"+
		"\3\64\3\64\3\64\3\64\3\64\3\64\5\64\u026d\n\64\3\64\3\64\3\64\3\64\3\64"+
		"\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64"+
		"\3\64\3\64\7\64\u0284\n\64\f\64\16\64\u0287\13\64\3\65\3\65\3\65\3\65"+
		"\3\65\3\65\3\65\3\65\5\65\u0291\n\65\3\66\3\66\3\66\3\66\3\66\3\66\3\66"+
		"\3\66\3\66\3\66\3\66\3\66\3\66\3\66\3\66\5\66\u02a2\n\66\3\67\3\67\3\67"+
		"\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67"+
		"\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\5\67\u02bd\n\67\38\38\38\38\3"+
		"8\38\38\38\38\38\38\38\38\38\38\38\38\38\38\38\38\38\38\38\38\38\38\3"+
		"8\38\38\38\38\38\38\38\58\u02e2\n8\38\2\3f9\2\4\6\b\n\f\16\20\22\24\26"+
		"\30\32\34\36 \"$&(*,.\60\62\64\668:<>@BDFHJLNPRTVXZ\\^`bdfhjln\2\t\3\2"+
		"\25\26\3\2\t\n\4\2\63\64\67\67\3\2\65\66\4\2//\61\61\4\2\60\60\62\62\3"+
		"\2)*\2\u0315\2p\3\2\2\2\4y\3\2\2\2\6\u0096\3\2\2\2\b\u0098\3\2\2\2\n\u00a3"+
		"\3\2\2\2\f\u00b0\3\2\2\2\16\u00ba\3\2\2\2\20\u00c2\3\2\2\2\22\u00ce\3"+
		"\2\2\2\24\u00d5\3\2\2\2\26\u00e5\3\2\2\2\30\u00e7\3\2\2\2\32\u00ef\3\2"+
		"\2\2\34\u00f4\3\2\2\2\36\u00fb\3\2\2\2 \u0102\3\2\2\2\"\u010b\3\2\2\2"+
		"$\u0111\3\2\2\2&\u0115\3\2\2\2(\u0129\3\2\2\2*\u0139\3\2\2\2,\u013b\3"+
		"\2\2\2.\u014a\3\2\2\2\60\u0168\3\2\2\2\62\u0181\3\2\2\2\64\u0183\3\2\2"+
		"\2\66\u0187\3\2\2\28\u0189\3\2\2\2:\u018b\3\2\2\2<\u0190\3\2\2\2>\u0195"+
		"\3\2\2\2@\u019a\3\2\2\2B\u01b8\3\2\2\2D\u01ba\3\2\2\2F\u01c1\3\2\2\2H"+
		"\u01d5\3\2\2\2J\u01e7\3\2\2\2L\u01f9\3\2\2\2N\u0200\3\2\2\2P\u020d\3\2"+
		"\2\2R\u0212\3\2\2\2T\u0216\3\2\2\2V\u021d\3\2\2\2X\u0222\3\2\2\2Z\u0228"+
		"\3\2\2\2\\\u022c\3\2\2\2^\u0232\3\2\2\2`\u023d\3\2\2\2b\u0244\3\2\2\2"+
		"d\u0252\3\2\2\2f\u026c\3\2\2\2h\u0290\3\2\2\2j\u02a1\3\2\2\2l\u02bc\3"+
		"\2\2\2n\u02e1\3\2\2\2pq\5\4\3\2qr\7\2\2\3r\3\3\2\2\2su\5\6\4\2tv\7@\2"+
		"\2ut\3\2\2\2uv\3\2\2\2vx\3\2\2\2ws\3\2\2\2x{\3\2\2\2yw\3\2\2\2yz\3\2\2"+
		"\2z\5\3\2\2\2{y\3\2\2\2|\u0097\5:\36\2}\u0097\5T+\2~\u0097\5V,\2\177\u0097"+
		"\5X-\2\u0080\u0097\5Z.\2\u0081\u0097\5J&\2\u0082\u0097\5N(\2\u0083\u0097"+
		"\5\\/\2\u0084\u0097\5^\60\2\u0085\u0097\5`\61\2\u0086\u0097\5d\63\2\u0087"+
		"\u0097\5B\"\2\u0088\u0097\5D#\2\u0089\u0097\5\66\34\2\u008a\u0097\58\35"+
		"\2\u008b\u0097\5\64\33\2\u008c\u0097\5\24\13\2\u008d\u0097\5\32\16\2\u008e"+
		"\u0097\5\36\20\2\u008f\u0097\5\"\22\2\u0090\u0097\5 \21\2\u0091\u0097"+
		"\5\34\17\2\u0092\u0097\5(\25\2\u0093\u0097\5\62\32\2\u0094\u0097\5\b\5"+
		"\2\u0095\u0097\5\f\7\2\u0096|\3\2\2\2\u0096}\3\2\2\2\u0096~\3\2\2\2\u0096"+
		"\177\3\2\2\2\u0096\u0080\3\2\2\2\u0096\u0081\3\2\2\2\u0096\u0082\3\2\2"+
		"\2\u0096\u0083\3\2\2\2\u0096\u0084\3\2\2\2\u0096\u0085\3\2\2\2\u0096\u0086"+
		"\3\2\2\2\u0096\u0087\3\2\2\2\u0096\u0088\3\2\2\2\u0096\u0089\3\2\2\2\u0096"+
		"\u008a\3\2\2\2\u0096\u008b\3\2\2\2\u0096\u008c\3\2\2\2\u0096\u008d\3\2"+
		"\2\2\u0096\u008e\3\2\2\2\u0096\u008f\3\2\2\2\u0096\u0090\3\2\2\2\u0096"+
		"\u0091\3\2\2\2\u0096\u0092\3\2\2\2\u0096\u0093\3\2\2\2\u0096\u0094\3\2"+
		"\2\2\u0096\u0095\3\2\2\2\u0097\7\3\2\2\2\u0098\u0099\7 \2\2\u0099\u009a"+
		"\7(\2\2\u009a\u009e\7:\2\2\u009b\u009d\5\n\6\2\u009c\u009b\3\2\2\2\u009d"+
		"\u00a0\3\2\2\2\u009e\u009c\3\2\2\2\u009e\u009f\3\2\2\2\u009f\u00a1\3\2"+
		"\2\2\u00a0\u009e\3\2\2\2\u00a1\u00a2\7;\2\2\u00a2\t\3\2\2\2\u00a3\u00a4"+
		"\t\2\2\2\u00a4\u00aa\7(\2\2\u00a5\u00a8\7>\2\2\u00a6\u00a9\5h\65\2\u00a7"+
		"\u00a9\7(\2\2\u00a8\u00a6\3\2\2\2\u00a8\u00a7\3\2\2\2\u00a9\u00ab\3\2"+
		"\2\2\u00aa\u00a5\3\2\2\2\u00aa\u00ab\3\2\2\2\u00ab\u00ae\3\2\2\2\u00ac"+
		"\u00ad\7.\2\2\u00ad\u00af\5f\64\2\u00ae\u00ac\3\2\2\2\u00ae\u00af\3\2"+
		"\2\2\u00af\13\3\2\2\2\u00b0\u00b1\7%\2\2\u00b1\u00b2\t\2\2\2\u00b2\u00b5"+
		"\7(\2\2\u00b3\u00b4\7>\2\2\u00b4\u00b6\7(\2\2\u00b5\u00b3\3\2\2\2\u00b5"+
		"\u00b6\3\2\2\2\u00b6\u00b7\3\2\2\2\u00b7\u00b8\7.\2\2\u00b8\u00b9\5\16"+
		"\b\2\u00b9\r\3\2\2\2\u00ba\u00bb\7%\2\2\u00bb\u00c0\7(\2\2\u00bc\u00bd"+
		"\78\2\2\u00bd\u00be\5\20\t\2\u00be\u00bf\79\2\2\u00bf\u00c1\3\2\2\2\u00c0"+
		"\u00bc\3\2\2\2\u00c0\u00c1\3\2\2\2\u00c1\17\3\2\2\2\u00c2\u00c3\7(\2\2"+
		"\u00c3\u00c4\7>\2\2\u00c4\u00cb\5f\64\2\u00c5\u00c6\7?\2\2\u00c6\u00c7"+
		"\7(\2\2\u00c7\u00c8\7>\2\2\u00c8\u00ca\5f\64\2\u00c9\u00c5\3\2\2\2\u00ca"+
		"\u00cd\3\2\2\2\u00cb\u00c9\3\2\2\2\u00cb\u00cc\3\2\2\2\u00cc\21\3\2\2"+
		"\2\u00cd\u00cb\3\2\2\2\u00ce\u00d1\7(\2\2\u00cf\u00d0\7B\2\2\u00d0\u00d2"+
		"\7(\2\2\u00d1\u00cf\3\2\2\2\u00d2\u00d3\3\2\2\2\u00d3\u00d1\3\2\2\2\u00d3"+
		"\u00d4\3\2\2\2\u00d4\23\3\2\2\2\u00d5\u00d6\7\25\2\2\u00d6\u00d7\7(\2"+
		"\2\u00d7\u00d8\7>\2\2\u00d8\u00d9\7<\2\2\u00d9\u00da\5h\65\2\u00da\u00db"+
		"\7=\2\2\u00db\u00dc\5\26\f\2\u00dc\25\3\2\2\2\u00dd\u00de\7.\2\2\u00de"+
		"\u00e0\7<\2\2\u00df\u00e1\5\30\r\2\u00e0\u00df\3\2\2\2\u00e0\u00e1\3\2"+
		"\2\2\u00e1\u00e2\3\2\2\2\u00e2\u00e6\7=\2\2\u00e3\u00e4\7.\2\2\u00e4\u00e6"+
		"\7(\2\2\u00e5\u00dd\3\2\2\2\u00e5\u00e3\3\2\2\2\u00e6\27\3\2\2\2\u00e7"+
		"\u00ec\5f\64\2\u00e8\u00e9\7?\2\2\u00e9\u00eb\5f\64\2\u00ea\u00e8\3\2"+
		"\2\2\u00eb\u00ee\3\2\2\2\u00ec\u00ea\3\2\2\2\u00ec\u00ed\3\2\2\2\u00ed"+
		"\31\3\2\2\2\u00ee\u00ec\3\2\2\2\u00ef\u00f0\7(\2\2\u00f0\u00f1\7<\2\2"+
		"\u00f1\u00f2\5f\64\2\u00f2\u00f3\7=\2\2\u00f3\33\3\2\2\2\u00f4\u00f5\7"+
		"(\2\2\u00f5\u00f6\7<\2\2\u00f6\u00f7\5f\64\2\u00f7\u00f8\7=\2\2\u00f8"+
		"\u00f9\7.\2\2\u00f9\u00fa\5f\64\2\u00fa\35\3\2\2\2\u00fb\u00fc\7(\2\2"+
		"\u00fc\u00fd\7B\2\2\u00fd\u00fe\7\35\2\2\u00fe\u00ff\78\2\2\u00ff\u0100"+
		"\5f\64\2\u0100\u0101\79\2\2\u0101\37\3\2\2\2\u0102\u0103\7(\2\2\u0103"+
		"\u0104\7B\2\2\u0104\u0105\7\37\2\2\u0105\u0106\78\2\2\u0106\u0107\7$\2"+
		"\2\u0107\u0108\7>\2\2\u0108\u0109\5f\64\2\u0109\u010a\79\2\2\u010a!\3"+
		"\2\2\2\u010b\u010c\7(\2\2\u010c\u010d\7B\2\2\u010d\u010e\7\36\2\2\u010e"+
		"\u010f\78\2\2\u010f\u0110\79\2\2\u0110#\3\2\2\2\u0111\u0112\7(\2\2\u0112"+
		"\u0113\7B\2\2\u0113\u0114\7\33\2\2\u0114%\3\2\2\2\u0115\u0116\7(\2\2\u0116"+
		"\u0117\7B\2\2\u0117\u0118\7\34\2\2\u0118\'\3\2\2\2\u0119\u011a\7\25\2"+
		"\2\u011a\u011d\7(\2\2\u011b\u011c\7>\2\2\u011c\u011e\5*\26\2\u011d\u011b"+
		"\3\2\2\2\u011d\u011e\3\2\2\2\u011e\u011f\3\2\2\2\u011f\u0120\7.\2\2\u0120"+
		"\u012a\5,\27\2\u0121\u0122\7\25\2\2\u0122\u0125\7(\2\2\u0123\u0124\7>"+
		"\2\2\u0124\u0126\5*\26\2\u0125\u0123\3\2\2\2\u0125\u0126\3\2\2\2\u0126"+
		"\u0127\3\2\2\2\u0127\u0128\7.\2\2\u0128\u012a\5.\30\2\u0129\u0119\3\2"+
		"\2\2\u0129\u0121\3\2\2\2\u012a)\3\2\2\2\u012b\u012c\7<\2\2\u012c\u012d"+
		"\7<\2\2\u012d\u012e\7<\2\2\u012e\u012f\5h\65\2\u012f\u0130\7=\2\2\u0130"+
		"\u0131\7=\2\2\u0131\u0132\7=\2\2\u0132\u013a\3\2\2\2\u0133\u0134\7<\2"+
		"\2\u0134\u0135\7<\2\2\u0135\u0136\5h\65\2\u0136\u0137\7=\2\2\u0137\u0138"+
		"\7=\2\2\u0138\u013a\3\2\2\2\u0139\u012b\3\2\2\2\u0139\u0133\3\2\2\2\u013a"+
		"+\3\2\2\2\u013b\u013c\7<\2\2\u013c\u013d\7<\2\2\u013d\u013e\5\30\r\2\u013e"+
		"\u0144\7=\2\2\u013f\u0140\7?\2\2\u0140\u0141\7<\2\2\u0141\u0142\5\30\r"+
		"\2\u0142\u0143\7=\2\2\u0143\u0145\3\2\2\2\u0144\u013f\3\2\2\2\u0145\u0146"+
		"\3\2\2\2\u0146\u0144\3\2\2\2\u0146\u0147\3\2\2\2\u0147\u0148\3\2\2\2\u0148"+
		"\u0149\7=\2\2\u0149-\3\2\2\2\u014a\u014b\7<\2\2\u014b\u0150\5,\27\2\u014c"+
		"\u014d\7?\2\2\u014d\u014f\5,\27\2\u014e\u014c\3\2\2\2\u014f\u0152\3\2"+
		"\2\2\u0150\u014e\3\2\2\2\u0150\u0151\3\2\2\2\u0151\u0153\3\2\2\2\u0152"+
		"\u0150\3\2\2\2\u0153\u0154\7=\2\2\u0154/\3\2\2\2\u0155\u0156\7(\2\2\u0156"+
		"\u0157\7<\2\2\u0157\u0158\5f\64\2\u0158\u0159\7=\2\2\u0159\u015a\7<\2"+
		"\2\u015a\u015b\5f\64\2\u015b\u015c\7=\2\2\u015c\u0169\3\2\2\2\u015d\u015e"+
		"\7(\2\2\u015e\u015f\7<\2\2\u015f\u0160\5f\64\2\u0160\u0161\7=\2\2\u0161"+
		"\u0162\7<\2\2\u0162\u0163\5f\64\2\u0163\u0164\7=\2\2\u0164\u0165\7<\2"+
		"\2\u0165\u0166\5f\64\2\u0166\u0167\7=\2\2\u0167\u0169\3\2\2\2\u0168\u0155"+
		"\3\2\2\2\u0168\u015d\3\2\2\2\u0169\61\3\2\2\2\u016a\u016b\7(\2\2\u016b"+
		"\u016c\7<\2\2\u016c\u016d\5f\64\2\u016d\u016e\7=\2\2\u016e\u016f\7<\2"+
		"\2\u016f\u0170\5f\64\2\u0170\u0171\7=\2\2\u0171\u0172\7.\2\2\u0172\u0173"+
		"\5f\64\2\u0173\u0182\3\2\2\2\u0174\u0175\7(\2\2\u0175\u0176\7<\2\2\u0176"+
		"\u0177\5f\64\2\u0177\u0178\7=\2\2\u0178\u0179\7<\2\2\u0179\u017a\5f\64"+
		"\2\u017a\u017b\7=\2\2\u017b\u017c\7<\2\2\u017c\u017d\5f\64\2\u017d\u017e"+
		"\7=\2\2\u017e\u017f\7.\2\2\u017f\u0180\5f\64\2\u0180\u0182\3\2\2\2\u0181"+
		"\u016a\3\2\2\2\u0181\u0174\3\2\2\2\u0182\63\3\2\2\2\u0183\u0184\7\30\2"+
		"\2\u0184\u0185\5f\64\2\u0185\u0186\7@\2\2\u0186\65\3\2\2\2\u0187\u0188"+
		"\7\27\2\2\u0188\67\3\2\2\2\u0189\u018a\7\31\2\2\u018a9\3\2\2\2\u018b\u018c"+
		"\7\13\2\2\u018c\u018d\78\2\2\u018d\u018e\5f\64\2\u018e\u018f\79\2\2\u018f"+
		";\3\2\2\2\u0190\u0191\7\3\2\2\u0191\u0192\78\2\2\u0192\u0193\5f\64\2\u0193"+
		"\u0194\79\2\2\u0194=\3\2\2\2\u0195\u0196\7\4\2\2\u0196\u0197\78\2\2\u0197"+
		"\u0198\5f\64\2\u0198\u0199\79\2\2\u0199?\3\2\2\2\u019a\u019b\7\7\2\2\u019b"+
		"\u019c\78\2\2\u019c\u019d\5f\64\2\u019d\u019e\79\2\2\u019eA\3\2\2\2\u019f"+
		"\u01a0\7\32\2\2\u01a0\u01a1\7(\2\2\u01a1\u01a3\78\2\2\u01a2\u01a4\5F$"+
		"\2\u01a3\u01a2\3\2\2\2\u01a3\u01a4\3\2\2\2\u01a4\u01a5\3\2\2\2\u01a5\u01a6"+
		"\79\2\2\u01a6\u01a7\7:\2\2\u01a7\u01a8\5\4\3\2\u01a8\u01a9\7;\2\2\u01a9"+
		"\u01b9\3\2\2\2\u01aa\u01ab\7\32\2\2\u01ab\u01ac\7(\2\2\u01ac\u01ae\78"+
		"\2\2\u01ad\u01af\5F$\2\u01ae\u01ad\3\2\2\2\u01ae\u01af\3\2\2\2\u01af\u01b0"+
		"\3\2\2\2\u01b0\u01b1\79\2\2\u01b1\u01b2\7\66\2\2\u01b2\u01b3\7\61\2\2"+
		"\u01b3\u01b4\5h\65\2\u01b4\u01b5\7:\2\2\u01b5\u01b6\5\4\3\2\u01b6\u01b7"+
		"\7;\2\2\u01b7\u01b9\3\2\2\2\u01b8\u019f\3\2\2\2\u01b8\u01aa\3\2\2\2\u01b9"+
		"C\3\2\2\2\u01ba\u01bb\7(\2\2\u01bb\u01bd\78\2\2\u01bc\u01be\5H%\2\u01bd"+
		"\u01bc\3\2\2\2\u01bd\u01be\3\2\2\2\u01be\u01bf\3\2\2\2\u01bf\u01c0\79"+
		"\2\2\u01c0E\3\2\2\2\u01c1\u01c2\7(\2\2\u01c2\u01c3\7(\2\2\u01c3\u01c5"+
		"\7>\2\2\u01c4\u01c6\7#\2\2\u01c5\u01c4\3\2\2\2\u01c5\u01c6\3\2\2\2\u01c6"+
		"\u01c7\3\2\2\2\u01c7\u01d2\5h\65\2\u01c8\u01c9\7?\2\2\u01c9\u01ca\7(\2"+
		"\2\u01ca\u01cb\7(\2\2\u01cb\u01cd\7>\2\2\u01cc\u01ce\7#\2\2\u01cd\u01cc"+
		"\3\2\2\2\u01cd\u01ce\3\2\2\2\u01ce\u01cf\3\2\2\2\u01cf\u01d1\5h\65\2\u01d0"+
		"\u01c8\3\2\2\2\u01d1\u01d4\3\2\2\2\u01d2\u01d0\3\2\2\2\u01d2\u01d3\3\2"+
		"\2\2\u01d3G\3\2\2\2\u01d4\u01d2\3\2\2\2\u01d5\u01d6\7(\2\2\u01d6\u01d8"+
		"\7>\2\2\u01d7\u01d9\7D\2\2\u01d8\u01d7\3\2\2\2\u01d8\u01d9\3\2\2\2\u01d9"+
		"\u01da\3\2\2\2\u01da\u01e4\5f\64\2\u01db\u01dc\7?\2\2\u01dc\u01dd\7(\2"+
		"\2\u01dd\u01df\7>\2\2\u01de\u01e0\7D\2\2\u01df\u01de\3\2\2\2\u01df\u01e0"+
		"\3\2\2\2\u01e0\u01e1\3\2\2\2\u01e1\u01e3\5f\64\2\u01e2\u01db\3\2\2\2\u01e3"+
		"\u01e6\3\2\2\2\u01e4\u01e2\3\2\2\2\u01e4\u01e5\3\2\2\2\u01e5I\3\2\2\2"+
		"\u01e6\u01e4\3\2\2\2\u01e7\u01e8\7\f\2\2\u01e8\u01e9\5f\64\2\u01e9\u01ea"+
		"\7:\2\2\u01ea\u01eb\5\4\3\2\u01eb\u01ef\7;\2\2\u01ec\u01ee\5L\'\2\u01ed"+
		"\u01ec\3\2\2\2\u01ee\u01f1\3\2\2\2\u01ef\u01ed\3\2\2\2\u01ef\u01f0\3\2"+
		"\2\2\u01f0\u01f7\3\2\2\2\u01f1\u01ef\3\2\2\2\u01f2\u01f3\7\r\2\2\u01f3"+
		"\u01f4\7:\2\2\u01f4\u01f5\5\4\3\2\u01f5\u01f6\7;\2\2\u01f6\u01f8\3\2\2"+
		"\2\u01f7\u01f2\3\2\2\2\u01f7\u01f8\3\2\2\2\u01f8K\3\2\2\2\u01f9\u01fa"+
		"\7\r\2\2\u01fa\u01fb\7\f\2\2\u01fb\u01fc\5f\64\2\u01fc\u01fd\7:\2\2\u01fd"+
		"\u01fe\5\4\3\2\u01fe\u01ff\7;\2\2\u01ffM\3\2\2\2\u0200\u0201\7\22\2\2"+
		"\u0201\u0202\5f\64\2\u0202\u0204\7:\2\2\u0203\u0205\5P)\2\u0204\u0203"+
		"\3\2\2\2\u0205\u0206\3\2\2\2\u0206\u0204\3\2\2\2\u0206\u0207\3\2\2\2\u0207"+
		"\u0209\3\2\2\2\u0208\u020a\5R*\2\u0209\u0208\3\2\2\2\u0209\u020a\3\2\2"+
		"\2\u020a\u020b\3\2\2\2\u020b\u020c\7;\2\2\u020cO\3\2\2\2\u020d\u020e\7"+
		"\23\2\2\u020e\u020f\5f\64\2\u020f\u0210\7>\2\2\u0210\u0211\5\4\3\2\u0211"+
		"Q\3\2\2\2\u0212\u0213\7\24\2\2\u0213\u0214\7>\2\2\u0214\u0215\5\4\3\2"+
		"\u0215S\3\2\2\2\u0216\u0217\t\2\2\2\u0217\u0218\7(\2\2\u0218\u0219\7>"+
		"\2\2\u0219\u021a\5h\65\2\u021a\u021b\7.\2\2\u021b\u021c\5f\64\2\u021c"+
		"U\3\2\2\2\u021d\u021e\t\2\2\2\u021e\u021f\7(\2\2\u021f\u0220\7.\2\2\u0220"+
		"\u0221\5f\64\2\u0221W\3\2\2\2\u0222\u0223\7\25\2\2\u0223\u0224\7(\2\2"+
		"\u0224\u0225\7>\2\2\u0225\u0226\5h\65\2\u0226\u0227\7A\2\2\u0227Y\3\2"+
		"\2\2\u0228\u0229\7(\2\2\u0229\u022a\7.\2\2\u022a\u022b\5f\64\2\u022b["+
		"\3\2\2\2\u022c\u022d\7\16\2\2\u022d\u022e\5f\64\2\u022e\u022f\7:\2\2\u022f"+
		"\u0230\5\4\3\2\u0230\u0231\7;\2\2\u0231]\3\2\2\2\u0232\u0233\7\17\2\2"+
		"\u0233\u0234\7(\2\2\u0234\u0237\7\21\2\2\u0235\u0238\5f\64\2\u0236\u0238"+
		"\5b\62\2\u0237\u0235\3\2\2\2\u0237\u0236\3\2\2\2\u0238\u0239\3\2\2\2\u0239"+
		"\u023a\7:\2\2\u023a\u023b\5\4\3\2\u023b\u023c\7;\2\2\u023c_\3\2\2\2\u023d"+
		"\u023e\7\20\2\2\u023e\u023f\5f\64\2\u023f\u0240\7\r\2\2\u0240\u0241\7"+
		":\2\2\u0241\u0242\5\4\3\2\u0242\u0243\7;\2\2\u0243a\3\2\2\2\u0244\u0245"+
		"\5f\64\2\u0245\u0246\7B\2\2\u0246\u0247\7B\2\2\u0247\u0248\7B\2\2\u0248"+
		"\u0249\5f\64\2\u0249c\3\2\2\2\u024a\u024b\7(\2\2\u024b\u024c\7\65\2\2"+
		"\u024c\u024d\7.\2\2\u024d\u0253\5f\64\2\u024e\u024f\7(\2\2\u024f\u0250"+
		"\7\66\2\2\u0250\u0251\7.\2\2\u0251\u0253\5f\64\2\u0252\u024a\3\2\2\2\u0252"+
		"\u024e\3\2\2\2\u0253e\3\2\2\2\u0254\u0255\b\64\1\2\u0255\u0256\78\2\2"+
		"\u0256\u0257\5f\64\2\u0257\u0258\79\2\2\u0258\u026d\3\2\2\2\u0259\u025a"+
		"\7\66\2\2\u025a\u026d\5f\64\32\u025b\u025c\7+\2\2\u025c\u026d\5f\64\31"+
		"\u025d\u026d\7&\2\2\u025e\u026d\7(\2\2\u025f\u026d\7\'\2\2\u0260\u026d"+
		"\t\3\2\2\u0261\u026d\7\b\2\2\u0262\u026d\5D#\2\u0263\u026d\5<\37\2\u0264"+
		"\u026d\5> \2\u0265\u026d\5@!\2\u0266\u026d\5\32\16\2\u0267\u026d\5$\23"+
		"\2\u0268\u026d\5&\24\2\u0269\u026d\5\60\31\2\u026a\u026d\5\22\n\2\u026b"+
		"\u026d\5\16\b\2\u026c\u0254\3\2\2\2\u026c\u0259\3\2\2\2\u026c\u025b\3"+
		"\2\2\2\u026c\u025d\3\2\2\2\u026c\u025e\3\2\2\2\u026c\u025f\3\2\2\2\u026c"+
		"\u0260\3\2\2\2\u026c\u0261\3\2\2\2\u026c\u0262\3\2\2\2\u026c\u0263\3\2"+
		"\2\2\u026c\u0264\3\2\2\2\u026c\u0265\3\2\2\2\u026c\u0266\3\2\2\2\u026c"+
		"\u0267\3\2\2\2\u026c\u0268\3\2\2\2\u026c\u0269\3\2\2\2\u026c\u026a\3\2"+
		"\2\2\u026c\u026b\3\2\2\2\u026d\u0285\3\2\2\2\u026e\u026f\f\30\2\2\u026f"+
		"\u0270\t\4\2\2\u0270\u0284\5f\64\31\u0271\u0272\f\27\2\2\u0272\u0273\t"+
		"\5\2\2\u0273\u0284\5f\64\30\u0274\u0275\f\26\2\2\u0275\u0276\t\6\2\2\u0276"+
		"\u0284\5f\64\27\u0277\u0278\f\25\2\2\u0278\u0279\t\7\2\2\u0279\u0284\5"+
		"f\64\26\u027a\u027b\f\24\2\2\u027b\u027c\t\b\2\2\u027c\u0284\5f\64\25"+
		"\u027d\u027e\f\23\2\2\u027e\u027f\7-\2\2\u027f\u0284\5f\64\24\u0280\u0281"+
		"\f\22\2\2\u0281\u0282\7,\2\2\u0282\u0284\5f\64\23\u0283\u026e\3\2\2\2"+
		"\u0283\u0271\3\2\2\2\u0283\u0274\3\2\2\2\u0283\u0277\3\2\2\2\u0283\u027a"+
		"\3\2\2\2\u0283\u027d\3\2\2\2\u0283\u0280\3\2\2\2\u0284\u0287\3\2\2\2\u0285"+
		"\u0283\3\2\2\2\u0285\u0286\3\2\2\2\u0286g\3\2\2\2\u0287\u0285\3\2\2\2"+
		"\u0288\u0291\7\3\2\2\u0289\u0291\7\4\2\2\u028a\u0291\7\5\2\2\u028b\u0291"+
		"\7\6\2\2\u028c\u0291\7\7\2\2\u028d\u0291\5j\66\2\u028e\u0291\5l\67\2\u028f"+
		"\u0291\5n8\2\u0290\u0288\3\2\2\2\u0290\u0289\3\2\2\2\u0290\u028a\3\2\2"+
		"\2\u0290\u028b\3\2\2\2\u0290\u028c\3\2\2\2\u0290\u028d\3\2\2\2\u0290\u028e"+
		"\3\2\2\2\u0290\u028f\3\2\2\2\u0291i\3\2\2\2\u0292\u0293\7<\2\2\u0293\u0294"+
		"\7\3\2\2\u0294\u02a2\7=\2\2\u0295\u0296\7<\2\2\u0296\u0297\7\4\2\2\u0297"+
		"\u02a2\7=\2\2\u0298\u0299\7<\2\2\u0299\u029a\7\5\2\2\u029a\u02a2\7=\2"+
		"\2\u029b\u029c\7<\2\2\u029c\u029d\7\6\2\2\u029d\u02a2\7=\2\2\u029e\u029f"+
		"\7<\2\2\u029f\u02a0\7\7\2\2\u02a0\u02a2\7=\2\2\u02a1\u0292\3\2\2\2\u02a1"+
		"\u0295\3\2\2\2\u02a1\u0298\3\2\2\2\u02a1\u029b\3\2\2\2\u02a1\u029e\3\2"+
		"\2\2\u02a2k\3\2\2\2\u02a3\u02a4\7<\2\2\u02a4\u02a5\7<\2\2\u02a5\u02a6"+
		"\7\3\2\2\u02a6\u02a7\7=\2\2\u02a7\u02bd\7=\2\2\u02a8\u02a9\7<\2\2\u02a9"+
		"\u02aa\7<\2\2\u02aa\u02ab\7\4\2\2\u02ab\u02ac\7=\2\2\u02ac\u02bd\7=\2"+
		"\2\u02ad\u02ae\7<\2\2\u02ae\u02af\7<\2\2\u02af\u02b0\7\5\2\2\u02b0\u02b1"+
		"\7=\2\2\u02b1\u02bd\7=\2\2\u02b2\u02b3\7<\2\2\u02b3\u02b4\7<\2\2\u02b4"+
		"\u02b5\7\6\2\2\u02b5\u02b6\7=\2\2\u02b6\u02bd\7=\2\2\u02b7\u02b8\7<\2"+
		"\2\u02b8\u02b9\7<\2\2\u02b9\u02ba\7\7\2\2\u02ba\u02bb\7=\2\2\u02bb\u02bd"+
		"\7=\2\2\u02bc\u02a3\3\2\2\2\u02bc\u02a8\3\2\2\2\u02bc\u02ad\3\2\2\2\u02bc"+
		"\u02b2\3\2\2\2\u02bc\u02b7\3\2\2\2\u02bdm\3\2\2\2\u02be\u02bf\7<\2\2\u02bf"+
		"\u02c0\7<\2\2\u02c0\u02c1\7<\2\2\u02c1\u02c2\7\3\2\2\u02c2\u02c3\7=\2"+
		"\2\u02c3\u02c4\7=\2\2\u02c4\u02e2\7=\2\2\u02c5\u02c6\7<\2\2\u02c6\u02c7"+
		"\7<\2\2\u02c7\u02c8\7<\2\2\u02c8\u02c9\7\4\2\2\u02c9\u02ca\7=\2\2\u02ca"+
		"\u02cb\7=\2\2\u02cb\u02e2\7=\2\2\u02cc\u02cd\7<\2\2\u02cd\u02ce\7<\2\2"+
		"\u02ce\u02cf\7<\2\2\u02cf\u02d0\7\5\2\2\u02d0\u02d1\7=\2\2\u02d1\u02d2"+
		"\7=\2\2\u02d2\u02e2\7=\2\2\u02d3\u02d4\7<\2\2\u02d4\u02d5\7<\2\2\u02d5"+
		"\u02d6\7<\2\2\u02d6\u02d7\7\6\2\2\u02d7\u02d8\7=\2\2\u02d8\u02d9\7=\2"+
		"\2\u02d9\u02e2\7=\2\2\u02da\u02db\7<\2\2\u02db\u02dc\7<\2\2\u02dc\u02dd"+
		"\7<\2\2\u02dd\u02de\7\7\2\2\u02de\u02df\7=\2\2\u02df\u02e0\7=\2\2\u02e0"+
		"\u02e2\7=\2\2\u02e1\u02be\3\2\2\2\u02e1\u02c5\3\2\2\2\u02e1\u02cc\3\2"+
		"\2\2\u02e1\u02d3\3\2\2\2\u02e1\u02da\3\2\2\2\u02e2o\3\2\2\2/uy\u0096\u009e"+
		"\u00a8\u00aa\u00ae\u00b5\u00c0\u00cb\u00d3\u00e0\u00e5\u00ec\u011d\u0125"+
		"\u0129\u0139\u0146\u0150\u0168\u0181\u01a3\u01ae\u01b8\u01bd\u01c5\u01cd"+
		"\u01d2\u01d8\u01df\u01e4\u01ef\u01f7\u0206\u0209\u0237\u0252\u026c\u0283"+
		"\u0285\u0290\u02a1\u02bc\u02e1";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
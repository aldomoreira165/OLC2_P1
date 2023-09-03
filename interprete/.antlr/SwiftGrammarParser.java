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
		STRUCT_VAR=31, STRUCT_LET=32, AT=33, ST=34, NUMBER=35, STRING=36, ID=37, 
		DIF=38, IG_IG=39, NOT=40, OR=41, AND=42, IG=43, MAY_IG=44, MEN_IG=45, 
		MAYOR=46, MENOR=47, MUL=48, DIV=49, ADD=50, SUB=51, MOD=52, PARIZQ=53, 
		PARDER=54, LLAVEIZQ=55, LLAVEDER=56, CORCHIZQ=57, CORCHDER=58, DOSPUNTOS=59, 
		COMA=60, PTCOMA=61, INTERROGACION=62, PUNTO=63, GUIONBAJO=64, WHITESPACE=65, 
		COMMENT=66, LINE_COMMENT=67;
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
		RULE_expr = 50, RULE_tipo = 51, RULE_tipo_vector = 52;
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
			"forstmt", "guardstmt", "rangostmt", "opasignstmt", "expr", "tipo", "tipo_vector"
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
			"'at'", "'st'", null, null, null, "'!='", "'=='", "'!'", "'||'", "'&&'", 
			"'='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'%'", 
			"'('", "')'", "'{'", "'}'", "'['", "']'", "':'", "','", "';'", "'?'", 
			"'.'", "'_'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "NIL", "TRU", "FAL", 
			"PRINT", "IF", "ELSE", "WHILE", "FOR", "GUARD", "IN", "SWITCH", "CASE", 
			"DEFAULT", "VAR", "LET", "BREAK", "RETURN", "CONTINUE", "FUNC", "COUNT", 
			"ISEMPTY", "APPEND", "REMOVE_LAST", "REMOVE", "STRUCT", "STRUCT_VAR", 
			"STRUCT_LET", "AT", "ST", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", 
			"OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", 
			"ADD", "SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "CORCHIZQ", 
			"CORCHDER", "DOSPUNTOS", "COMA", "PTCOMA", "INTERROGACION", "PUNTO", 
			"GUIONBAJO", "WHITESPACE", "COMMENT", "LINE_COMMENT"
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
			setState(106);
			block();
			setState(107);
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
			setState(115);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PRINT) | (1L << IF) | (1L << WHILE) | (1L << FOR) | (1L << GUARD) | (1L << SWITCH) | (1L << VAR) | (1L << LET) | (1L << BREAK) | (1L << RETURN) | (1L << CONTINUE) | (1L << FUNC) | (1L << STRUCT) | (1L << ST) | (1L << ID))) != 0)) {
				{
				{
				setState(109);
				stmt();
				setState(111);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(110);
					match(PTCOMA);
					}
				}

				}
				}
				setState(117);
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
			setState(144);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(118);
				printstmt();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(119);
				typedDeclstmt();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(120);
				untypedDeclstmt();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(121);
				optionalTypedDeclstmt();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(122);
				asignstmt();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(123);
				ifstmt();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(124);
				switchstmt();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(125);
				whilestmt();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(126);
				forstmt();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(127);
				guardstmt();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(128);
				opasignstmt();
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(129);
				funcdclstmt();
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(130);
				accfuncstm();
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(131);
				breakstmt();
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(132);
				continuestmt();
				}
				break;
			case 16:
				enterOuterAlt(_localctx, 16);
				{
				setState(133);
				returnstmt();
				}
				break;
			case 17:
				enterOuterAlt(_localctx, 17);
				{
				setState(134);
				declvectorstmt();
				}
				break;
			case 18:
				enterOuterAlt(_localctx, 18);
				{
				setState(135);
				accesovectorstmt();
				}
				break;
			case 19:
				enterOuterAlt(_localctx, 19);
				{
				setState(136);
				appendvectorstmt();
				}
				break;
			case 20:
				enterOuterAlt(_localctx, 20);
				{
				setState(137);
				removelastvectorstmt();
				}
				break;
			case 21:
				enterOuterAlt(_localctx, 21);
				{
				setState(138);
				removeatvectorstmt();
				}
				break;
			case 22:
				enterOuterAlt(_localctx, 22);
				{
				setState(139);
				asignvectorstmt();
				}
				break;
			case 23:
				enterOuterAlt(_localctx, 23);
				{
				setState(140);
				declmatrizstmt();
				}
				break;
			case 24:
				enterOuterAlt(_localctx, 24);
				{
				setState(141);
				asignmatrizstmt();
				}
				break;
			case 25:
				enterOuterAlt(_localctx, 25);
				{
				setState(142);
				defstructstmt();
				}
				break;
			case 26:
				enterOuterAlt(_localctx, 26);
				{
				setState(143);
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
			setState(146);
			match(STRUCT);
			setState(147);
			match(ID);
			setState(148);
			match(LLAVEIZQ);
			setState(152);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR || _la==LET) {
				{
				{
				setState(149);
				lista_atributos();
				}
				}
				setState(154);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(155);
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
			setState(157);
			_la = _input.LA(1);
			if ( !(_la==VAR || _la==LET) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(158);
			match(ID);
			setState(164);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DOSPUNTOS) {
				{
				setState(159);
				match(DOSPUNTOS);
				setState(162);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case INT:
				case FLOAT:
				case BOOL:
				case CHARACTER:
				case PSTRING:
				case CORCHIZQ:
					{
					setState(160);
					tipo();
					}
					break;
				case ID:
					{
					setState(161);
					match(ID);
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
			}

			setState(168);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IG) {
				{
				setState(166);
				match(IG);
				setState(167);
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
			setState(170);
			match(ST);
			setState(171);
			_la = _input.LA(1);
			if ( !(_la==VAR || _la==LET) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(172);
			match(ID);
			setState(175);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DOSPUNTOS) {
				{
				setState(173);
				match(DOSPUNTOS);
				setState(174);
				match(ID);
				}
			}

			setState(177);
			match(IG);
			setState(178);
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
			setState(180);
			match(ST);
			setState(181);
			match(ID);
			setState(186);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				{
				setState(182);
				match(PARIZQ);
				setState(183);
				l_dupla();
				setState(184);
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
			setState(188);
			match(ID);
			setState(189);
			match(DOSPUNTOS);
			setState(190);
			expr(0);
			setState(197);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(191);
				match(COMA);
				setState(192);
				match(ID);
				setState(193);
				match(DOSPUNTOS);
				setState(194);
				expr(0);
				}
				}
				setState(199);
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
			setState(200);
			match(ID);
			setState(203); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(201);
					match(PUNTO);
					setState(202);
					match(ID);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(205); 
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
			setState(207);
			match(VAR);
			setState(208);
			match(ID);
			setState(209);
			match(DOSPUNTOS);
			setState(210);
			match(CORCHIZQ);
			setState(211);
			tipo();
			setState(212);
			match(CORCHDER);
			setState(213);
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
			setState(223);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				_localctx = new DefVectorContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(215);
				match(IG);
				setState(216);
				match(CORCHIZQ);
				setState(218);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << INT) | (1L << FLOAT) | (1L << PSTRING) | (1L << NIL) | (1L << TRU) | (1L << FAL) | (1L << ST) | (1L << NUMBER) | (1L << STRING) | (1L << ID) | (1L << NOT) | (1L << SUB) | (1L << PARIZQ))) != 0)) {
					{
					setState(217);
					listaexpresiones();
					}
				}

				setState(220);
				match(CORCHDER);
				}
				break;
			case 2:
				_localctx = new DefVectorIDContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(221);
				match(IG);
				setState(222);
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
			setState(225);
			expr(0);
			setState(230);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(226);
				match(COMA);
				setState(227);
				expr(0);
				}
				}
				setState(232);
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
			setState(233);
			match(ID);
			setState(234);
			match(CORCHIZQ);
			setState(235);
			expr(0);
			setState(236);
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
			setState(238);
			match(ID);
			setState(239);
			match(CORCHIZQ);
			setState(240);
			expr(0);
			setState(241);
			match(CORCHDER);
			setState(242);
			match(IG);
			setState(243);
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
			setState(245);
			match(ID);
			setState(246);
			match(PUNTO);
			setState(247);
			match(APPEND);
			setState(248);
			match(PARIZQ);
			setState(249);
			expr(0);
			setState(250);
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
			setState(252);
			match(ID);
			setState(253);
			match(PUNTO);
			setState(254);
			match(REMOVE);
			setState(255);
			match(PARIZQ);
			setState(256);
			match(AT);
			setState(257);
			match(DOSPUNTOS);
			setState(258);
			expr(0);
			setState(259);
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
			setState(261);
			match(ID);
			setState(262);
			match(PUNTO);
			setState(263);
			match(REMOVE_LAST);
			setState(264);
			match(PARIZQ);
			setState(265);
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
			setState(267);
			match(ID);
			setState(268);
			match(PUNTO);
			setState(269);
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
			setState(271);
			match(ID);
			setState(272);
			match(PUNTO);
			setState(273);
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
			setState(291);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				_localctx = new Declmatrizstmt2Context(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(275);
				match(VAR);
				setState(276);
				match(ID);
				setState(279);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DOSPUNTOS) {
					{
					setState(277);
					match(DOSPUNTOS);
					setState(278);
					tipomatriz();
					}
				}

				setState(281);
				match(IG);
				setState(282);
				listavaloresmatriz();
				}
				break;
			case 2:
				_localctx = new Declmatrizstmt3Context(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(283);
				match(VAR);
				setState(284);
				match(ID);
				setState(287);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DOSPUNTOS) {
					{
					setState(285);
					match(DOSPUNTOS);
					setState(286);
					tipomatriz();
					}
				}

				setState(289);
				match(IG);
				setState(290);
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
			setState(307);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				_localctx = new Tipomatriz2Context(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(293);
				match(CORCHIZQ);
				setState(294);
				match(CORCHIZQ);
				setState(295);
				tipo();
				setState(296);
				match(CORCHDER);
				setState(297);
				match(CORCHDER);
				}
				break;
			case 2:
				_localctx = new Tipomatriz3Context(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(299);
				match(CORCHIZQ);
				setState(300);
				match(CORCHIZQ);
				setState(301);
				match(CORCHIZQ);
				setState(302);
				tipo();
				setState(303);
				match(CORCHDER);
				setState(304);
				match(CORCHDER);
				setState(305);
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
			setState(309);
			match(CORCHIZQ);
			setState(310);
			match(CORCHIZQ);
			setState(311);
			listaexpresiones();
			setState(312);
			match(CORCHDER);
			setState(318); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(313);
				match(COMA);
				setState(314);
				match(CORCHIZQ);
				setState(315);
				listaexpresiones();
				setState(316);
				match(CORCHDER);
				}
				}
				setState(320); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==COMA );
			setState(322);
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
			setState(324);
			match(CORCHIZQ);
			setState(325);
			listavaloresmatriz();
			setState(330);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(326);
				match(COMA);
				setState(327);
				listavaloresmatriz();
				}
				}
				setState(332);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(333);
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
			setState(354);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				_localctx = new Accesomatriz2Context(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(335);
				match(ID);
				setState(336);
				match(CORCHIZQ);
				setState(337);
				expr(0);
				setState(338);
				match(CORCHDER);
				setState(339);
				match(CORCHIZQ);
				setState(340);
				expr(0);
				setState(341);
				match(CORCHDER);
				}
				break;
			case 2:
				_localctx = new Accesomatriz3Context(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(343);
				match(ID);
				setState(344);
				match(CORCHIZQ);
				setState(345);
				expr(0);
				setState(346);
				match(CORCHDER);
				setState(347);
				match(CORCHIZQ);
				setState(348);
				expr(0);
				setState(349);
				match(CORCHDER);
				setState(350);
				match(CORCHIZQ);
				setState(351);
				expr(0);
				setState(352);
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
			setState(379);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				_localctx = new Asignmatrizstmt2Context(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(356);
				match(ID);
				setState(357);
				match(CORCHIZQ);
				setState(358);
				expr(0);
				setState(359);
				match(CORCHDER);
				setState(360);
				match(CORCHIZQ);
				setState(361);
				expr(0);
				setState(362);
				match(CORCHDER);
				setState(363);
				match(IG);
				setState(364);
				expr(0);
				}
				break;
			case 2:
				_localctx = new Asignmatrizstmt3Context(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(366);
				match(ID);
				setState(367);
				match(CORCHIZQ);
				setState(368);
				expr(0);
				setState(369);
				match(CORCHDER);
				setState(370);
				match(CORCHIZQ);
				setState(371);
				expr(0);
				setState(372);
				match(CORCHDER);
				setState(373);
				match(CORCHIZQ);
				setState(374);
				expr(0);
				setState(375);
				match(CORCHDER);
				setState(376);
				match(IG);
				setState(377);
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
			setState(381);
			match(RETURN);
			setState(382);
			expr(0);
			setState(383);
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
			setState(385);
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
			setState(387);
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
			setState(389);
			match(PRINT);
			setState(390);
			match(PARIZQ);
			setState(391);
			expr(0);
			setState(392);
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
			setState(394);
			match(INT);
			setState(395);
			match(PARIZQ);
			setState(396);
			expr(0);
			setState(397);
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
			setState(399);
			match(FLOAT);
			setState(400);
			match(PARIZQ);
			setState(401);
			expr(0);
			setState(402);
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
			setState(404);
			match(PSTRING);
			setState(405);
			match(PARIZQ);
			setState(406);
			expr(0);
			setState(407);
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
			setState(434);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				_localctx = new FuncionNormalContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(409);
				match(FUNC);
				setState(410);
				match(ID);
				setState(411);
				match(PARIZQ);
				setState(413);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ID) {
					{
					setState(412);
					parametros();
					}
				}

				setState(415);
				match(PARDER);
				setState(416);
				match(LLAVEIZQ);
				setState(417);
				block();
				setState(418);
				match(LLAVEDER);
				}
				break;
			case 2:
				_localctx = new FuncionRetornoContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(420);
				match(FUNC);
				setState(421);
				match(ID);
				setState(422);
				match(PARIZQ);
				setState(424);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ID) {
					{
					setState(423);
					parametros();
					}
				}

				setState(426);
				match(PARDER);
				setState(427);
				match(SUB);
				setState(428);
				match(MAYOR);
				setState(429);
				tipo();
				setState(430);
				match(LLAVEIZQ);
				setState(431);
				block();
				setState(432);
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
			setState(436);
			match(ID);
			setState(437);
			match(PARIZQ);
			setState(439);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(438);
				parametroscall();
				}
			}

			setState(441);
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
		enterRule(_localctx, 68, RULE_parametros);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(443);
			match(ID);
			setState(444);
			match(ID);
			setState(445);
			match(DOSPUNTOS);
			setState(446);
			tipo();
			setState(454);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(447);
				match(COMA);
				setState(448);
				match(ID);
				setState(449);
				match(ID);
				setState(450);
				match(DOSPUNTOS);
				setState(451);
				tipo();
				}
				}
				setState(456);
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
		enterRule(_localctx, 70, RULE_parametroscall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(457);
			match(ID);
			setState(458);
			match(DOSPUNTOS);
			setState(459);
			expr(0);
			setState(466);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(460);
				match(COMA);
				setState(461);
				match(ID);
				setState(462);
				match(DOSPUNTOS);
				setState(463);
				expr(0);
				}
				}
				setState(468);
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
			setState(469);
			match(IF);
			setState(470);
			expr(0);
			setState(471);
			match(LLAVEIZQ);
			setState(472);
			block();
			setState(473);
			match(LLAVEDER);
			setState(477);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(474);
					elseifstmt();
					}
					} 
				}
				setState(479);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
			}
			setState(485);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(480);
				match(ELSE);
				setState(481);
				match(LLAVEIZQ);
				setState(482);
				block();
				setState(483);
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
			setState(487);
			match(ELSE);
			setState(488);
			match(IF);
			setState(489);
			expr(0);
			setState(490);
			match(LLAVEIZQ);
			setState(491);
			block();
			setState(492);
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
			setState(494);
			match(SWITCH);
			setState(495);
			expr(0);
			setState(496);
			match(LLAVEIZQ);
			setState(498); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(497);
				caseStmt();
				}
				}
				setState(500); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==CASE );
			setState(503);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(502);
				defaultCase();
				}
			}

			setState(505);
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
			setState(507);
			match(CASE);
			setState(508);
			expr(0);
			setState(509);
			match(DOSPUNTOS);
			setState(510);
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
			setState(512);
			match(DEFAULT);
			setState(513);
			match(DOSPUNTOS);
			setState(514);
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
			setState(516);
			_la = _input.LA(1);
			if ( !(_la==VAR || _la==LET) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(517);
			match(ID);
			setState(518);
			match(DOSPUNTOS);
			setState(519);
			tipo();
			setState(520);
			match(IG);
			setState(521);
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
			setState(523);
			_la = _input.LA(1);
			if ( !(_la==VAR || _la==LET) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(524);
			match(ID);
			setState(525);
			match(IG);
			setState(526);
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
			setState(528);
			match(VAR);
			setState(529);
			match(ID);
			setState(530);
			match(DOSPUNTOS);
			setState(531);
			tipo();
			setState(532);
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
			setState(534);
			match(ID);
			setState(535);
			match(IG);
			setState(536);
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
			setState(538);
			match(WHILE);
			setState(539);
			expr(0);
			setState(540);
			match(LLAVEIZQ);
			setState(541);
			block();
			setState(542);
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
			setState(544);
			match(FOR);
			setState(545);
			match(ID);
			setState(546);
			match(IN);
			setState(549);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				{
				setState(547);
				expr(0);
				}
				break;
			case 2:
				{
				setState(548);
				rangostmt();
				}
				break;
			}
			setState(551);
			match(LLAVEIZQ);
			setState(552);
			block();
			setState(553);
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
			setState(555);
			match(GUARD);
			setState(556);
			expr(0);
			setState(557);
			match(ELSE);
			setState(558);
			match(LLAVEIZQ);
			setState(559);
			block();
			setState(560);
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
			setState(562);
			expr(0);
			setState(563);
			match(PUNTO);
			setState(564);
			match(PUNTO);
			setState(565);
			match(PUNTO);
			setState(566);
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
			setState(576);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				_localctx = new IncrementoContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(568);
				match(ID);
				setState(569);
				match(ADD);
				setState(570);
				match(IG);
				setState(571);
				expr(0);
				}
				break;
			case 2:
				_localctx = new DecrementoContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(572);
				match(ID);
				setState(573);
				match(SUB);
				setState(574);
				match(IG);
				setState(575);
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
			setState(602);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
			case 1:
				{
				_localctx = new ParExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(579);
				match(PARIZQ);
				setState(580);
				expr(0);
				setState(581);
				match(PARDER);
				}
				break;
			case 2:
				{
				_localctx = new UnariaExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(583);
				match(SUB);
				setState(584);
				expr(24);
				}
				break;
			case 3:
				{
				_localctx = new NotExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(585);
				match(NOT);
				setState(586);
				expr(23);
				}
				break;
			case 4:
				{
				_localctx = new NumExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(587);
				match(NUMBER);
				}
				break;
			case 5:
				{
				_localctx = new IdExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(588);
				match(ID);
				}
				break;
			case 6:
				{
				_localctx = new StrExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(589);
				match(STRING);
				}
				break;
			case 7:
				{
				_localctx = new BoolExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(590);
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
				setState(591);
				match(NIL);
				}
				break;
			case 9:
				{
				_localctx = new AccFuncExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(592);
				accfuncstm();
				}
				break;
			case 10:
				{
				_localctx = new IntExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(593);
				intstmt();
				}
				break;
			case 11:
				{
				_localctx = new FloatExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(594);
				floatstmt();
				}
				break;
			case 12:
				{
				_localctx = new StringExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(595);
				stringstmt();
				}
				break;
			case 13:
				{
				_localctx = new AccesoVectorExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(596);
				accesovectorstmt();
				}
				break;
			case 14:
				{
				_localctx = new CountVectorExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(597);
				countvectorstmt();
				}
				break;
			case 15:
				{
				_localctx = new IsEmptyVectorExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(598);
				isemptyvectorstmt();
				}
				break;
			case 16:
				{
				_localctx = new AccesoMatrizExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(599);
				accesomatriz();
				}
				break;
			case 17:
				{
				_localctx = new AccesoStructExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(600);
				accesostructstmt();
				}
				break;
			case 18:
				{
				_localctx = new AccesoValorStructExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(601);
				valor_struct_expr();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(627);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,36,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(625);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,35,_ctx) ) {
					case 1:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(604);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(605);
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
						setState(606);
						((OpExprContext)_localctx).right = expr(23);
						}
						break;
					case 2:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(607);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(608);
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
						setState(609);
						((OpExprContext)_localctx).right = expr(22);
						}
						break;
					case 3:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(610);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(611);
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
						setState(612);
						((OpExprContext)_localctx).right = expr(21);
						}
						break;
					case 4:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(613);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(614);
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
						setState(615);
						((OpExprContext)_localctx).right = expr(20);
						}
						break;
					case 5:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(616);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(617);
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
						setState(618);
						((OpExprContext)_localctx).right = expr(19);
						}
						break;
					case 6:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(619);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(620);
						((OpExprContext)_localctx).op = match(AND);
						setState(621);
						((OpExprContext)_localctx).right = expr(18);
						}
						break;
					case 7:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(622);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(623);
						((OpExprContext)_localctx).op = match(OR);
						setState(624);
						((OpExprContext)_localctx).right = expr(17);
						}
						break;
					}
					} 
				}
				setState(629);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,36,_ctx);
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
		public TipoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo; }
	}

	public final TipoContext tipo() throws RecognitionException {
		TipoContext _localctx = new TipoContext(_ctx, getState());
		enterRule(_localctx, 102, RULE_tipo);
		try {
			setState(636);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(630);
				match(INT);
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(631);
				match(FLOAT);
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 3);
				{
				setState(632);
				match(BOOL);
				}
				break;
			case CHARACTER:
				enterOuterAlt(_localctx, 4);
				{
				setState(633);
				match(CHARACTER);
				}
				break;
			case PSTRING:
				enterOuterAlt(_localctx, 5);
				{
				setState(634);
				match(PSTRING);
				}
				break;
			case CORCHIZQ:
				enterOuterAlt(_localctx, 6);
				{
				setState(635);
				tipo_vector();
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
			setState(653);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,38,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(638);
				match(CORCHIZQ);
				setState(639);
				match(INT);
				setState(640);
				match(CORCHDER);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(641);
				match(CORCHIZQ);
				setState(642);
				match(FLOAT);
				setState(643);
				match(CORCHDER);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(644);
				match(CORCHIZQ);
				setState(645);
				match(BOOL);
				setState(646);
				match(CORCHDER);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(647);
				match(CORCHIZQ);
				setState(648);
				match(CHARACTER);
				setState(649);
				match(CORCHDER);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(650);
				match(CORCHIZQ);
				setState(651);
				match(PSTRING);
				setState(652);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3E\u0292\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\3\2\3\2\3\2\3\3\3\3\5\3r\n\3\7\3t\n\3\f\3\16\3"+
		"w\13\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4\u0093\n\4\3\5\3\5\3\5\3\5"+
		"\7\5\u0099\n\5\f\5\16\5\u009c\13\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\5\6\u00a5"+
		"\n\6\5\6\u00a7\n\6\3\6\3\6\5\6\u00ab\n\6\3\7\3\7\3\7\3\7\3\7\5\7\u00b2"+
		"\n\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\5\b\u00bd\n\b\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\7\t\u00c6\n\t\f\t\16\t\u00c9\13\t\3\n\3\n\3\n\6\n\u00ce\n"+
		"\n\r\n\16\n\u00cf\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f"+
		"\5\f\u00dd\n\f\3\f\3\f\3\f\5\f\u00e2\n\f\3\r\3\r\3\r\7\r\u00e7\n\r\f\r"+
		"\16\r\u00ea\13\r\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\24"+
		"\3\24\3\24\3\24\3\25\3\25\3\25\3\25\5\25\u011a\n\25\3\25\3\25\3\25\3\25"+
		"\3\25\3\25\5\25\u0122\n\25\3\25\3\25\5\25\u0126\n\25\3\26\3\26\3\26\3"+
		"\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\5\26\u0136\n\26"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\6\27\u0141\n\27\r\27\16"+
		"\27\u0142\3\27\3\27\3\30\3\30\3\30\3\30\7\30\u014b\n\30\f\30\16\30\u014e"+
		"\13\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\5\31\u0165\n\31\3\32\3\32\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\5\32\u017e\n\32\3\33\3\33\3\33\3\33\3\34"+
		"\3\34\3\35\3\35\3\36\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3 \3"+
		" \3 \3 \3 \3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\5\"\u01a0\n\"\3\"\3\"\3\"\3"+
		"\"\3\"\3\"\3\"\3\"\3\"\5\"\u01ab\n\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\5"+
		"\"\u01b5\n\"\3#\3#\3#\5#\u01ba\n#\3#\3#\3$\3$\3$\3$\3$\3$\3$\3$\3$\7$"+
		"\u01c7\n$\f$\16$\u01ca\13$\3%\3%\3%\3%\3%\3%\3%\7%\u01d3\n%\f%\16%\u01d6"+
		"\13%\3&\3&\3&\3&\3&\3&\7&\u01de\n&\f&\16&\u01e1\13&\3&\3&\3&\3&\3&\5&"+
		"\u01e8\n&\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3(\3(\3(\3(\6(\u01f5\n(\r(\16(\u01f6"+
		"\3(\5(\u01fa\n(\3(\3(\3)\3)\3)\3)\3)\3*\3*\3*\3*\3+\3+\3+\3+\3+\3+\3+"+
		"\3,\3,\3,\3,\3,\3-\3-\3-\3-\3-\3-\3.\3.\3.\3.\3/\3/\3/\3/\3/\3/\3\60\3"+
		"\60\3\60\3\60\3\60\5\60\u0228\n\60\3\60\3\60\3\60\3\60\3\61\3\61\3\61"+
		"\3\61\3\61\3\61\3\61\3\62\3\62\3\62\3\62\3\62\3\62\3\63\3\63\3\63\3\63"+
		"\3\63\3\63\3\63\3\63\5\63\u0243\n\63\3\64\3\64\3\64\3\64\3\64\3\64\3\64"+
		"\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64"+
		"\3\64\3\64\3\64\5\64\u025d\n\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64"+
		"\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\7\64"+
		"\u0274\n\64\f\64\16\64\u0277\13\64\3\65\3\65\3\65\3\65\3\65\3\65\5\65"+
		"\u027f\n\65\3\66\3\66\3\66\3\66\3\66\3\66\3\66\3\66\3\66\3\66\3\66\3\66"+
		"\3\66\3\66\3\66\5\66\u0290\n\66\3\66\2\3f\67\2\4\6\b\n\f\16\20\22\24\26"+
		"\30\32\34\36 \"$&(*,.\60\62\64\668:<>@BDFHJLNPRTVXZ\\^`bdfhj\2\t\3\2\25"+
		"\26\3\2\t\n\4\2\62\63\66\66\3\2\64\65\4\2..\60\60\4\2//\61\61\3\2()\2"+
		"\u02b7\2l\3\2\2\2\4u\3\2\2\2\6\u0092\3\2\2\2\b\u0094\3\2\2\2\n\u009f\3"+
		"\2\2\2\f\u00ac\3\2\2\2\16\u00b6\3\2\2\2\20\u00be\3\2\2\2\22\u00ca\3\2"+
		"\2\2\24\u00d1\3\2\2\2\26\u00e1\3\2\2\2\30\u00e3\3\2\2\2\32\u00eb\3\2\2"+
		"\2\34\u00f0\3\2\2\2\36\u00f7\3\2\2\2 \u00fe\3\2\2\2\"\u0107\3\2\2\2$\u010d"+
		"\3\2\2\2&\u0111\3\2\2\2(\u0125\3\2\2\2*\u0135\3\2\2\2,\u0137\3\2\2\2."+
		"\u0146\3\2\2\2\60\u0164\3\2\2\2\62\u017d\3\2\2\2\64\u017f\3\2\2\2\66\u0183"+
		"\3\2\2\28\u0185\3\2\2\2:\u0187\3\2\2\2<\u018c\3\2\2\2>\u0191\3\2\2\2@"+
		"\u0196\3\2\2\2B\u01b4\3\2\2\2D\u01b6\3\2\2\2F\u01bd\3\2\2\2H\u01cb\3\2"+
		"\2\2J\u01d7\3\2\2\2L\u01e9\3\2\2\2N\u01f0\3\2\2\2P\u01fd\3\2\2\2R\u0202"+
		"\3\2\2\2T\u0206\3\2\2\2V\u020d\3\2\2\2X\u0212\3\2\2\2Z\u0218\3\2\2\2\\"+
		"\u021c\3\2\2\2^\u0222\3\2\2\2`\u022d\3\2\2\2b\u0234\3\2\2\2d\u0242\3\2"+
		"\2\2f\u025c\3\2\2\2h\u027e\3\2\2\2j\u028f\3\2\2\2lm\5\4\3\2mn\7\2\2\3"+
		"n\3\3\2\2\2oq\5\6\4\2pr\7?\2\2qp\3\2\2\2qr\3\2\2\2rt\3\2\2\2so\3\2\2\2"+
		"tw\3\2\2\2us\3\2\2\2uv\3\2\2\2v\5\3\2\2\2wu\3\2\2\2x\u0093\5:\36\2y\u0093"+
		"\5T+\2z\u0093\5V,\2{\u0093\5X-\2|\u0093\5Z.\2}\u0093\5J&\2~\u0093\5N("+
		"\2\177\u0093\5\\/\2\u0080\u0093\5^\60\2\u0081\u0093\5`\61\2\u0082\u0093"+
		"\5d\63\2\u0083\u0093\5B\"\2\u0084\u0093\5D#\2\u0085\u0093\5\66\34\2\u0086"+
		"\u0093\58\35\2\u0087\u0093\5\64\33\2\u0088\u0093\5\24\13\2\u0089\u0093"+
		"\5\32\16\2\u008a\u0093\5\36\20\2\u008b\u0093\5\"\22\2\u008c\u0093\5 \21"+
		"\2\u008d\u0093\5\34\17\2\u008e\u0093\5(\25\2\u008f\u0093\5\62\32\2\u0090"+
		"\u0093\5\b\5\2\u0091\u0093\5\f\7\2\u0092x\3\2\2\2\u0092y\3\2\2\2\u0092"+
		"z\3\2\2\2\u0092{\3\2\2\2\u0092|\3\2\2\2\u0092}\3\2\2\2\u0092~\3\2\2\2"+
		"\u0092\177\3\2\2\2\u0092\u0080\3\2\2\2\u0092\u0081\3\2\2\2\u0092\u0082"+
		"\3\2\2\2\u0092\u0083\3\2\2\2\u0092\u0084\3\2\2\2\u0092\u0085\3\2\2\2\u0092"+
		"\u0086\3\2\2\2\u0092\u0087\3\2\2\2\u0092\u0088\3\2\2\2\u0092\u0089\3\2"+
		"\2\2\u0092\u008a\3\2\2\2\u0092\u008b\3\2\2\2\u0092\u008c\3\2\2\2\u0092"+
		"\u008d\3\2\2\2\u0092\u008e\3\2\2\2\u0092\u008f\3\2\2\2\u0092\u0090\3\2"+
		"\2\2\u0092\u0091\3\2\2\2\u0093\7\3\2\2\2\u0094\u0095\7 \2\2\u0095\u0096"+
		"\7\'\2\2\u0096\u009a\79\2\2\u0097\u0099\5\n\6\2\u0098\u0097\3\2\2\2\u0099"+
		"\u009c\3\2\2\2\u009a\u0098\3\2\2\2\u009a\u009b\3\2\2\2\u009b\u009d\3\2"+
		"\2\2\u009c\u009a\3\2\2\2\u009d\u009e\7:\2\2\u009e\t\3\2\2\2\u009f\u00a0"+
		"\t\2\2\2\u00a0\u00a6\7\'\2\2\u00a1\u00a4\7=\2\2\u00a2\u00a5\5h\65\2\u00a3"+
		"\u00a5\7\'\2\2\u00a4\u00a2\3\2\2\2\u00a4\u00a3\3\2\2\2\u00a5\u00a7\3\2"+
		"\2\2\u00a6\u00a1\3\2\2\2\u00a6\u00a7\3\2\2\2\u00a7\u00aa\3\2\2\2\u00a8"+
		"\u00a9\7-\2\2\u00a9\u00ab\5f\64\2\u00aa\u00a8\3\2\2\2\u00aa\u00ab\3\2"+
		"\2\2\u00ab\13\3\2\2\2\u00ac\u00ad\7$\2\2\u00ad\u00ae\t\2\2\2\u00ae\u00b1"+
		"\7\'\2\2\u00af\u00b0\7=\2\2\u00b0\u00b2\7\'\2\2\u00b1\u00af\3\2\2\2\u00b1"+
		"\u00b2\3\2\2\2\u00b2\u00b3\3\2\2\2\u00b3\u00b4\7-\2\2\u00b4\u00b5\5\16"+
		"\b\2\u00b5\r\3\2\2\2\u00b6\u00b7\7$\2\2\u00b7\u00bc\7\'\2\2\u00b8\u00b9"+
		"\7\67\2\2\u00b9\u00ba\5\20\t\2\u00ba\u00bb\78\2\2\u00bb\u00bd\3\2\2\2"+
		"\u00bc\u00b8\3\2\2\2\u00bc\u00bd\3\2\2\2\u00bd\17\3\2\2\2\u00be\u00bf"+
		"\7\'\2\2\u00bf\u00c0\7=\2\2\u00c0\u00c7\5f\64\2\u00c1\u00c2\7>\2\2\u00c2"+
		"\u00c3\7\'\2\2\u00c3\u00c4\7=\2\2\u00c4\u00c6\5f\64\2\u00c5\u00c1\3\2"+
		"\2\2\u00c6\u00c9\3\2\2\2\u00c7\u00c5\3\2\2\2\u00c7\u00c8\3\2\2\2\u00c8"+
		"\21\3\2\2\2\u00c9\u00c7\3\2\2\2\u00ca\u00cd\7\'\2\2\u00cb\u00cc\7A\2\2"+
		"\u00cc\u00ce\7\'\2\2\u00cd\u00cb\3\2\2\2\u00ce\u00cf\3\2\2\2\u00cf\u00cd"+
		"\3\2\2\2\u00cf\u00d0\3\2\2\2\u00d0\23\3\2\2\2\u00d1\u00d2\7\25\2\2\u00d2"+
		"\u00d3\7\'\2\2\u00d3\u00d4\7=\2\2\u00d4\u00d5\7;\2\2\u00d5\u00d6\5h\65"+
		"\2\u00d6\u00d7\7<\2\2\u00d7\u00d8\5\26\f\2\u00d8\25\3\2\2\2\u00d9\u00da"+
		"\7-\2\2\u00da\u00dc\7;\2\2\u00db\u00dd\5\30\r\2\u00dc\u00db\3\2\2\2\u00dc"+
		"\u00dd\3\2\2\2\u00dd\u00de\3\2\2\2\u00de\u00e2\7<\2\2\u00df\u00e0\7-\2"+
		"\2\u00e0\u00e2\7\'\2\2\u00e1\u00d9\3\2\2\2\u00e1\u00df\3\2\2\2\u00e2\27"+
		"\3\2\2\2\u00e3\u00e8\5f\64\2\u00e4\u00e5\7>\2\2\u00e5\u00e7\5f\64\2\u00e6"+
		"\u00e4\3\2\2\2\u00e7\u00ea\3\2\2\2\u00e8\u00e6\3\2\2\2\u00e8\u00e9\3\2"+
		"\2\2\u00e9\31\3\2\2\2\u00ea\u00e8\3\2\2\2\u00eb\u00ec\7\'\2\2\u00ec\u00ed"+
		"\7;\2\2\u00ed\u00ee\5f\64\2\u00ee\u00ef\7<\2\2\u00ef\33\3\2\2\2\u00f0"+
		"\u00f1\7\'\2\2\u00f1\u00f2\7;\2\2\u00f2\u00f3\5f\64\2\u00f3\u00f4\7<\2"+
		"\2\u00f4\u00f5\7-\2\2\u00f5\u00f6\5f\64\2\u00f6\35\3\2\2\2\u00f7\u00f8"+
		"\7\'\2\2\u00f8\u00f9\7A\2\2\u00f9\u00fa\7\35\2\2\u00fa\u00fb\7\67\2\2"+
		"\u00fb\u00fc\5f\64\2\u00fc\u00fd\78\2\2\u00fd\37\3\2\2\2\u00fe\u00ff\7"+
		"\'\2\2\u00ff\u0100\7A\2\2\u0100\u0101\7\37\2\2\u0101\u0102\7\67\2\2\u0102"+
		"\u0103\7#\2\2\u0103\u0104\7=\2\2\u0104\u0105\5f\64\2\u0105\u0106\78\2"+
		"\2\u0106!\3\2\2\2\u0107\u0108\7\'\2\2\u0108\u0109\7A\2\2\u0109\u010a\7"+
		"\36\2\2\u010a\u010b\7\67\2\2\u010b\u010c\78\2\2\u010c#\3\2\2\2\u010d\u010e"+
		"\7\'\2\2\u010e\u010f\7A\2\2\u010f\u0110\7\33\2\2\u0110%\3\2\2\2\u0111"+
		"\u0112\7\'\2\2\u0112\u0113\7A\2\2\u0113\u0114\7\34\2\2\u0114\'\3\2\2\2"+
		"\u0115\u0116\7\25\2\2\u0116\u0119\7\'\2\2\u0117\u0118\7=\2\2\u0118\u011a"+
		"\5*\26\2\u0119\u0117\3\2\2\2\u0119\u011a\3\2\2\2\u011a\u011b\3\2\2\2\u011b"+
		"\u011c\7-\2\2\u011c\u0126\5,\27\2\u011d\u011e\7\25\2\2\u011e\u0121\7\'"+
		"\2\2\u011f\u0120\7=\2\2\u0120\u0122\5*\26\2\u0121\u011f\3\2\2\2\u0121"+
		"\u0122\3\2\2\2\u0122\u0123\3\2\2\2\u0123\u0124\7-\2\2\u0124\u0126\5.\30"+
		"\2\u0125\u0115\3\2\2\2\u0125\u011d\3\2\2\2\u0126)\3\2\2\2\u0127\u0128"+
		"\7;\2\2\u0128\u0129\7;\2\2\u0129\u012a\5h\65\2\u012a\u012b\7<\2\2\u012b"+
		"\u012c\7<\2\2\u012c\u0136\3\2\2\2\u012d\u012e\7;\2\2\u012e\u012f\7;\2"+
		"\2\u012f\u0130\7;\2\2\u0130\u0131\5h\65\2\u0131\u0132\7<\2\2\u0132\u0133"+
		"\7<\2\2\u0133\u0134\7<\2\2\u0134\u0136\3\2\2\2\u0135\u0127\3\2\2\2\u0135"+
		"\u012d\3\2\2\2\u0136+\3\2\2\2\u0137\u0138\7;\2\2\u0138\u0139\7;\2\2\u0139"+
		"\u013a\5\30\r\2\u013a\u0140\7<\2\2\u013b\u013c\7>\2\2\u013c\u013d\7;\2"+
		"\2\u013d\u013e\5\30\r\2\u013e\u013f\7<\2\2\u013f\u0141\3\2\2\2\u0140\u013b"+
		"\3\2\2\2\u0141\u0142\3\2\2\2\u0142\u0140\3\2\2\2\u0142\u0143\3\2\2\2\u0143"+
		"\u0144\3\2\2\2\u0144\u0145\7<\2\2\u0145-\3\2\2\2\u0146\u0147\7;\2\2\u0147"+
		"\u014c\5,\27\2\u0148\u0149\7>\2\2\u0149\u014b\5,\27\2\u014a\u0148\3\2"+
		"\2\2\u014b\u014e\3\2\2\2\u014c\u014a\3\2\2\2\u014c\u014d\3\2\2\2\u014d"+
		"\u014f\3\2\2\2\u014e\u014c\3\2\2\2\u014f\u0150\7<\2\2\u0150/\3\2\2\2\u0151"+
		"\u0152\7\'\2\2\u0152\u0153\7;\2\2\u0153\u0154\5f\64\2\u0154\u0155\7<\2"+
		"\2\u0155\u0156\7;\2\2\u0156\u0157\5f\64\2\u0157\u0158\7<\2\2\u0158\u0165"+
		"\3\2\2\2\u0159\u015a\7\'\2\2\u015a\u015b\7;\2\2\u015b\u015c\5f\64\2\u015c"+
		"\u015d\7<\2\2\u015d\u015e\7;\2\2\u015e\u015f\5f\64\2\u015f\u0160\7<\2"+
		"\2\u0160\u0161\7;\2\2\u0161\u0162\5f\64\2\u0162\u0163\7<\2\2\u0163\u0165"+
		"\3\2\2\2\u0164\u0151\3\2\2\2\u0164\u0159\3\2\2\2\u0165\61\3\2\2\2\u0166"+
		"\u0167\7\'\2\2\u0167\u0168\7;\2\2\u0168\u0169\5f\64\2\u0169\u016a\7<\2"+
		"\2\u016a\u016b\7;\2\2\u016b\u016c\5f\64\2\u016c\u016d\7<\2\2\u016d\u016e"+
		"\7-\2\2\u016e\u016f\5f\64\2\u016f\u017e\3\2\2\2\u0170\u0171\7\'\2\2\u0171"+
		"\u0172\7;\2\2\u0172\u0173\5f\64\2\u0173\u0174\7<\2\2\u0174\u0175\7;\2"+
		"\2\u0175\u0176\5f\64\2\u0176\u0177\7<\2\2\u0177\u0178\7;\2\2\u0178\u0179"+
		"\5f\64\2\u0179\u017a\7<\2\2\u017a\u017b\7-\2\2\u017b\u017c\5f\64\2\u017c"+
		"\u017e\3\2\2\2\u017d\u0166\3\2\2\2\u017d\u0170\3\2\2\2\u017e\63\3\2\2"+
		"\2\u017f\u0180\7\30\2\2\u0180\u0181\5f\64\2\u0181\u0182\7?\2\2\u0182\65"+
		"\3\2\2\2\u0183\u0184\7\27\2\2\u0184\67\3\2\2\2\u0185\u0186\7\31\2\2\u0186"+
		"9\3\2\2\2\u0187\u0188\7\13\2\2\u0188\u0189\7\67\2\2\u0189\u018a\5f\64"+
		"\2\u018a\u018b\78\2\2\u018b;\3\2\2\2\u018c\u018d\7\3\2\2\u018d\u018e\7"+
		"\67\2\2\u018e\u018f\5f\64\2\u018f\u0190\78\2\2\u0190=\3\2\2\2\u0191\u0192"+
		"\7\4\2\2\u0192\u0193\7\67\2\2\u0193\u0194\5f\64\2\u0194\u0195\78\2\2\u0195"+
		"?\3\2\2\2\u0196\u0197\7\7\2\2\u0197\u0198\7\67\2\2\u0198\u0199\5f\64\2"+
		"\u0199\u019a\78\2\2\u019aA\3\2\2\2\u019b\u019c\7\32\2\2\u019c\u019d\7"+
		"\'\2\2\u019d\u019f\7\67\2\2\u019e\u01a0\5F$\2\u019f\u019e\3\2\2\2\u019f"+
		"\u01a0\3\2\2\2\u01a0\u01a1\3\2\2\2\u01a1\u01a2\78\2\2\u01a2\u01a3\79\2"+
		"\2\u01a3\u01a4\5\4\3\2\u01a4\u01a5\7:\2\2\u01a5\u01b5\3\2\2\2\u01a6\u01a7"+
		"\7\32\2\2\u01a7\u01a8\7\'\2\2\u01a8\u01aa\7\67\2\2\u01a9\u01ab\5F$\2\u01aa"+
		"\u01a9\3\2\2\2\u01aa\u01ab\3\2\2\2\u01ab\u01ac\3\2\2\2\u01ac\u01ad\78"+
		"\2\2\u01ad\u01ae\7\65\2\2\u01ae\u01af\7\60\2\2\u01af\u01b0\5h\65\2\u01b0"+
		"\u01b1\79\2\2\u01b1\u01b2\5\4\3\2\u01b2\u01b3\7:\2\2\u01b3\u01b5\3\2\2"+
		"\2\u01b4\u019b\3\2\2\2\u01b4\u01a6\3\2\2\2\u01b5C\3\2\2\2\u01b6\u01b7"+
		"\7\'\2\2\u01b7\u01b9\7\67\2\2\u01b8\u01ba\5H%\2\u01b9\u01b8\3\2\2\2\u01b9"+
		"\u01ba\3\2\2\2\u01ba\u01bb\3\2\2\2\u01bb\u01bc\78\2\2\u01bcE\3\2\2\2\u01bd"+
		"\u01be\7\'\2\2\u01be\u01bf\7\'\2\2\u01bf\u01c0\7=\2\2\u01c0\u01c8\5h\65"+
		"\2\u01c1\u01c2\7>\2\2\u01c2\u01c3\7\'\2\2\u01c3\u01c4\7\'\2\2\u01c4\u01c5"+
		"\7=\2\2\u01c5\u01c7\5h\65\2\u01c6\u01c1\3\2\2\2\u01c7\u01ca\3\2\2\2\u01c8"+
		"\u01c6\3\2\2\2\u01c8\u01c9\3\2\2\2\u01c9G\3\2\2\2\u01ca\u01c8\3\2\2\2"+
		"\u01cb\u01cc\7\'\2\2\u01cc\u01cd\7=\2\2\u01cd\u01d4\5f\64\2\u01ce\u01cf"+
		"\7>\2\2\u01cf\u01d0\7\'\2\2\u01d0\u01d1\7=\2\2\u01d1\u01d3\5f\64\2\u01d2"+
		"\u01ce\3\2\2\2\u01d3\u01d6\3\2\2\2\u01d4\u01d2\3\2\2\2\u01d4\u01d5\3\2"+
		"\2\2\u01d5I\3\2\2\2\u01d6\u01d4\3\2\2\2\u01d7\u01d8\7\f\2\2\u01d8\u01d9"+
		"\5f\64\2\u01d9\u01da\79\2\2\u01da\u01db\5\4\3\2\u01db\u01df\7:\2\2\u01dc"+
		"\u01de\5L\'\2\u01dd\u01dc\3\2\2\2\u01de\u01e1\3\2\2\2\u01df\u01dd\3\2"+
		"\2\2\u01df\u01e0\3\2\2\2\u01e0\u01e7\3\2\2\2\u01e1\u01df\3\2\2\2\u01e2"+
		"\u01e3\7\r\2\2\u01e3\u01e4\79\2\2\u01e4\u01e5\5\4\3\2\u01e5\u01e6\7:\2"+
		"\2\u01e6\u01e8\3\2\2\2\u01e7\u01e2\3\2\2\2\u01e7\u01e8\3\2\2\2\u01e8K"+
		"\3\2\2\2\u01e9\u01ea\7\r\2\2\u01ea\u01eb\7\f\2\2\u01eb\u01ec\5f\64\2\u01ec"+
		"\u01ed\79\2\2\u01ed\u01ee\5\4\3\2\u01ee\u01ef\7:\2\2\u01efM\3\2\2\2\u01f0"+
		"\u01f1\7\22\2\2\u01f1\u01f2\5f\64\2\u01f2\u01f4\79\2\2\u01f3\u01f5\5P"+
		")\2\u01f4\u01f3\3\2\2\2\u01f5\u01f6\3\2\2\2\u01f6\u01f4\3\2\2\2\u01f6"+
		"\u01f7\3\2\2\2\u01f7\u01f9\3\2\2\2\u01f8\u01fa\5R*\2\u01f9\u01f8\3\2\2"+
		"\2\u01f9\u01fa\3\2\2\2\u01fa\u01fb\3\2\2\2\u01fb\u01fc\7:\2\2\u01fcO\3"+
		"\2\2\2\u01fd\u01fe\7\23\2\2\u01fe\u01ff\5f\64\2\u01ff\u0200\7=\2\2\u0200"+
		"\u0201\5\4\3\2\u0201Q\3\2\2\2\u0202\u0203\7\24\2\2\u0203\u0204\7=\2\2"+
		"\u0204\u0205\5\4\3\2\u0205S\3\2\2\2\u0206\u0207\t\2\2\2\u0207\u0208\7"+
		"\'\2\2\u0208\u0209\7=\2\2\u0209\u020a\5h\65\2\u020a\u020b\7-\2\2\u020b"+
		"\u020c\5f\64\2\u020cU\3\2\2\2\u020d\u020e\t\2\2\2\u020e\u020f\7\'\2\2"+
		"\u020f\u0210\7-\2\2\u0210\u0211\5f\64\2\u0211W\3\2\2\2\u0212\u0213\7\25"+
		"\2\2\u0213\u0214\7\'\2\2\u0214\u0215\7=\2\2\u0215\u0216\5h\65\2\u0216"+
		"\u0217\7@\2\2\u0217Y\3\2\2\2\u0218\u0219\7\'\2\2\u0219\u021a\7-\2\2\u021a"+
		"\u021b\5f\64\2\u021b[\3\2\2\2\u021c\u021d\7\16\2\2\u021d\u021e\5f\64\2"+
		"\u021e\u021f\79\2\2\u021f\u0220\5\4\3\2\u0220\u0221\7:\2\2\u0221]\3\2"+
		"\2\2\u0222\u0223\7\17\2\2\u0223\u0224\7\'\2\2\u0224\u0227\7\21\2\2\u0225"+
		"\u0228\5f\64\2\u0226\u0228\5b\62\2\u0227\u0225\3\2\2\2\u0227\u0226\3\2"+
		"\2\2\u0228\u0229\3\2\2\2\u0229\u022a\79\2\2\u022a\u022b\5\4\3\2\u022b"+
		"\u022c\7:\2\2\u022c_\3\2\2\2\u022d\u022e\7\20\2\2\u022e\u022f\5f\64\2"+
		"\u022f\u0230\7\r\2\2\u0230\u0231\79\2\2\u0231\u0232\5\4\3\2\u0232\u0233"+
		"\7:\2\2\u0233a\3\2\2\2\u0234\u0235\5f\64\2\u0235\u0236\7A\2\2\u0236\u0237"+
		"\7A\2\2\u0237\u0238\7A\2\2\u0238\u0239\5f\64\2\u0239c\3\2\2\2\u023a\u023b"+
		"\7\'\2\2\u023b\u023c\7\64\2\2\u023c\u023d\7-\2\2\u023d\u0243\5f\64\2\u023e"+
		"\u023f\7\'\2\2\u023f\u0240\7\65\2\2\u0240\u0241\7-\2\2\u0241\u0243\5f"+
		"\64\2\u0242\u023a\3\2\2\2\u0242\u023e\3\2\2\2\u0243e\3\2\2\2\u0244\u0245"+
		"\b\64\1\2\u0245\u0246\7\67\2\2\u0246\u0247\5f\64\2\u0247\u0248\78\2\2"+
		"\u0248\u025d\3\2\2\2\u0249\u024a\7\65\2\2\u024a\u025d\5f\64\32\u024b\u024c"+
		"\7*\2\2\u024c\u025d\5f\64\31\u024d\u025d\7%\2\2\u024e\u025d\7\'\2\2\u024f"+
		"\u025d\7&\2\2\u0250\u025d\t\3\2\2\u0251\u025d\7\b\2\2\u0252\u025d\5D#"+
		"\2\u0253\u025d\5<\37\2\u0254\u025d\5> \2\u0255\u025d\5@!\2\u0256\u025d"+
		"\5\32\16\2\u0257\u025d\5$\23\2\u0258\u025d\5&\24\2\u0259\u025d\5\60\31"+
		"\2\u025a\u025d\5\22\n\2\u025b\u025d\5\16\b\2\u025c\u0244\3\2\2\2\u025c"+
		"\u0249\3\2\2\2\u025c\u024b\3\2\2\2\u025c\u024d\3\2\2\2\u025c\u024e\3\2"+
		"\2\2\u025c\u024f\3\2\2\2\u025c\u0250\3\2\2\2\u025c\u0251\3\2\2\2\u025c"+
		"\u0252\3\2\2\2\u025c\u0253\3\2\2\2\u025c\u0254\3\2\2\2\u025c\u0255\3\2"+
		"\2\2\u025c\u0256\3\2\2\2\u025c\u0257\3\2\2\2\u025c\u0258\3\2\2\2\u025c"+
		"\u0259\3\2\2\2\u025c\u025a\3\2\2\2\u025c\u025b\3\2\2\2\u025d\u0275\3\2"+
		"\2\2\u025e\u025f\f\30\2\2\u025f\u0260\t\4\2\2\u0260\u0274\5f\64\31\u0261"+
		"\u0262\f\27\2\2\u0262\u0263\t\5\2\2\u0263\u0274\5f\64\30\u0264\u0265\f"+
		"\26\2\2\u0265\u0266\t\6\2\2\u0266\u0274\5f\64\27\u0267\u0268\f\25\2\2"+
		"\u0268\u0269\t\7\2\2\u0269\u0274\5f\64\26\u026a\u026b\f\24\2\2\u026b\u026c"+
		"\t\b\2\2\u026c\u0274\5f\64\25\u026d\u026e\f\23\2\2\u026e\u026f\7,\2\2"+
		"\u026f\u0274\5f\64\24\u0270\u0271\f\22\2\2\u0271\u0272\7+\2\2\u0272\u0274"+
		"\5f\64\23\u0273\u025e\3\2\2\2\u0273\u0261\3\2\2\2\u0273\u0264\3\2\2\2"+
		"\u0273\u0267\3\2\2\2\u0273\u026a\3\2\2\2\u0273\u026d\3\2\2\2\u0273\u0270"+
		"\3\2\2\2\u0274\u0277\3\2\2\2\u0275\u0273\3\2\2\2\u0275\u0276\3\2\2\2\u0276"+
		"g\3\2\2\2\u0277\u0275\3\2\2\2\u0278\u027f\7\3\2\2\u0279\u027f\7\4\2\2"+
		"\u027a\u027f\7\5\2\2\u027b\u027f\7\6\2\2\u027c\u027f\7\7\2\2\u027d\u027f"+
		"\5j\66\2\u027e\u0278\3\2\2\2\u027e\u0279\3\2\2\2\u027e\u027a\3\2\2\2\u027e"+
		"\u027b\3\2\2\2\u027e\u027c\3\2\2\2\u027e\u027d\3\2\2\2\u027fi\3\2\2\2"+
		"\u0280\u0281\7;\2\2\u0281\u0282\7\3\2\2\u0282\u0290\7<\2\2\u0283\u0284"+
		"\7;\2\2\u0284\u0285\7\4\2\2\u0285\u0290\7<\2\2\u0286\u0287\7;\2\2\u0287"+
		"\u0288\7\5\2\2\u0288\u0290\7<\2\2\u0289\u028a\7;\2\2\u028a\u028b\7\6\2"+
		"\2\u028b\u0290\7<\2\2\u028c\u028d\7;\2\2\u028d\u028e\7\7\2\2\u028e\u0290"+
		"\7<\2\2\u028f\u0280\3\2\2\2\u028f\u0283\3\2\2\2\u028f\u0286\3\2\2\2\u028f"+
		"\u0289\3\2\2\2\u028f\u028c\3\2\2\2\u0290k\3\2\2\2)qu\u0092\u009a\u00a4"+
		"\u00a6\u00aa\u00b1\u00bc\u00c7\u00cf\u00dc\u00e1\u00e8\u0119\u0121\u0125"+
		"\u0135\u0142\u014c\u0164\u017d\u019f\u01aa\u01b4\u01b9\u01c8\u01d4\u01df"+
		"\u01e7\u01f6\u01f9\u0227\u0242\u025c\u0273\u0275\u027e\u028f";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
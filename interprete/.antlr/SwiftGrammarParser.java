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
		NUMBER=18, STRING=19, ID=20, DIF=21, IG_IG=22, NOT=23, OR=24, AND=25, 
		IG=26, MAY_IG=27, MEN_IG=28, MAYOR=29, MENOR=30, MUL=31, DIV=32, ADD=33, 
		SUB=34, MOD=35, PARIZQ=36, PARDER=37, LLAVEIZQ=38, LLAVEDER=39, DOSPUNTOS=40, 
		INTERROGACION=41, WHITESPACE=42, COMMENT=43, LINE_COMMENT=44;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_stmt = 2, RULE_printstmt = 3, RULE_breakstmt = 4, 
		RULE_ifstmt = 5, RULE_elseifstmt = 6, RULE_switchstmt = 7, RULE_caseStmt = 8, 
		RULE_defaultCase = 9, RULE_typedDeclstmt = 10, RULE_untypedDeclstmt = 11, 
		RULE_optionalTypedDeclstmt = 12, RULE_asignstmt = 13, RULE_whilestmt = 14, 
		RULE_opasignstmt = 15, RULE_expr = 16, RULE_tipo = 17;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "stmt", "printstmt", "breakstmt", "ifstmt", "elseifstmt", 
			"switchstmt", "caseStmt", "defaultCase", "typedDeclstmt", "untypedDeclstmt", 
			"optionalTypedDeclstmt", "asignstmt", "whilestmt", "opasignstmt", "expr", 
			"tipo"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'int'", "'float'", "'bool'", "'character'", "'String'", "'nil'", 
			"'true'", "'false'", "'print'", "'if'", "'else'", "'while'", "'switch'", 
			"'case'", "'default'", "'var'", "'break'", null, null, null, "'!='", 
			"'=='", "'!'", "'||'", "'&&'", "'='", "'>='", "'<='", "'>'", "'<'", "'*'", 
			"'/'", "'+'", "'-'", "'%'", "'('", "')'", "'{'", "'}'", "':'", "'?'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "NIL", "TRU", "FAL", 
			"PRINT", "IF", "ELSE", "WHILE", "SWITCH", "CASE", "DEFAULT", "VAR", "BREAK", 
			"NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND", "IG", "MAY_IG", 
			"MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB", "MOD", "PARIZQ", 
			"PARDER", "LLAVEIZQ", "LLAVEDER", "DOSPUNTOS", "INTERROGACION", "WHITESPACE", 
			"COMMENT", "LINE_COMMENT"
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
			setState(36);
			block();
			setState(37);
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
			setState(42);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PRINT) | (1L << IF) | (1L << WHILE) | (1L << SWITCH) | (1L << VAR) | (1L << BREAK) | (1L << ID))) != 0)) {
				{
				{
				setState(39);
				stmt();
				}
				}
				setState(44);
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
		public StmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt; }
	}

	public final StmtContext stmt() throws RecognitionException {
		StmtContext _localctx = new StmtContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_stmt);
		try {
			setState(55);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(45);
				printstmt();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(46);
				typedDeclstmt();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(47);
				untypedDeclstmt();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(48);
				optionalTypedDeclstmt();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(49);
				asignstmt();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(50);
				ifstmt();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(51);
				switchstmt();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(52);
				whilestmt();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(53);
				opasignstmt();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(54);
				breakstmt();
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
		enterRule(_localctx, 6, RULE_printstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(57);
			match(PRINT);
			setState(58);
			match(PARIZQ);
			setState(59);
			expr(0);
			setState(60);
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

	public static class BreakstmtContext extends ParserRuleContext {
		public TerminalNode BREAK() { return getToken(SwiftGrammarParser.BREAK, 0); }
		public BreakstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_breakstmt; }
	}

	public final BreakstmtContext breakstmt() throws RecognitionException {
		BreakstmtContext _localctx = new BreakstmtContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_breakstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(62);
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
		enterRule(_localctx, 10, RULE_ifstmt);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(64);
			match(IF);
			setState(65);
			expr(0);
			setState(66);
			match(LLAVEIZQ);
			setState(67);
			block();
			setState(68);
			match(LLAVEDER);
			setState(72);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(69);
					elseifstmt();
					}
					} 
				}
				setState(74);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			}
			setState(80);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(75);
				match(ELSE);
				setState(76);
				match(LLAVEIZQ);
				setState(77);
				block();
				setState(78);
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
		enterRule(_localctx, 12, RULE_elseifstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(82);
			match(ELSE);
			setState(83);
			match(IF);
			setState(84);
			expr(0);
			setState(85);
			match(LLAVEIZQ);
			setState(86);
			block();
			setState(87);
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
		enterRule(_localctx, 14, RULE_switchstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(89);
			match(SWITCH);
			setState(90);
			expr(0);
			setState(91);
			match(LLAVEIZQ);
			setState(93); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(92);
				caseStmt();
				}
				}
				setState(95); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==CASE );
			setState(98);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(97);
				defaultCase();
				}
			}

			setState(100);
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
		enterRule(_localctx, 16, RULE_caseStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(102);
			match(CASE);
			setState(103);
			expr(0);
			setState(104);
			match(DOSPUNTOS);
			setState(105);
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
		enterRule(_localctx, 18, RULE_defaultCase);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(107);
			match(DEFAULT);
			setState(108);
			match(DOSPUNTOS);
			setState(109);
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
		enterRule(_localctx, 20, RULE_typedDeclstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(111);
			match(VAR);
			setState(112);
			match(ID);
			setState(113);
			match(DOSPUNTOS);
			setState(114);
			tipo();
			setState(115);
			match(IG);
			setState(116);
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
		enterRule(_localctx, 22, RULE_untypedDeclstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(118);
			match(VAR);
			setState(119);
			match(ID);
			setState(120);
			match(IG);
			setState(121);
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
		enterRule(_localctx, 24, RULE_optionalTypedDeclstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(123);
			match(VAR);
			setState(124);
			match(ID);
			setState(125);
			match(DOSPUNTOS);
			setState(126);
			tipo();
			setState(127);
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
		enterRule(_localctx, 26, RULE_asignstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(129);
			match(ID);
			setState(130);
			match(IG);
			setState(131);
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
		enterRule(_localctx, 28, RULE_whilestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(133);
			match(WHILE);
			setState(134);
			expr(0);
			setState(135);
			match(LLAVEIZQ);
			setState(136);
			block();
			setState(137);
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
		enterRule(_localctx, 30, RULE_opasignstmt);
		try {
			setState(147);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				_localctx = new IncrementoContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(139);
				match(ID);
				setState(140);
				match(ADD);
				setState(141);
				match(IG);
				setState(142);
				expr(0);
				}
				break;
			case 2:
				_localctx = new DecrementoContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(143);
				match(ID);
				setState(144);
				match(SUB);
				setState(145);
				match(IG);
				setState(146);
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
	public static class UnariaExprContext extends ExprContext {
		public TerminalNode SUB() { return getToken(SwiftGrammarParser.SUB, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public UnariaExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class BoolExprContext extends ExprContext {
		public TerminalNode TRU() { return getToken(SwiftGrammarParser.TRU, 0); }
		public TerminalNode FAL() { return getToken(SwiftGrammarParser.FAL, 0); }
		public BoolExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class NilExprContext extends ExprContext {
		public TerminalNode NIL() { return getToken(SwiftGrammarParser.NIL, 0); }
		public NilExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class IdExprContext extends ExprContext {
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public IdExprContext(ExprContext ctx) { copyFrom(ctx); }
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

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 32;
		enterRecursionRule(_localctx, 32, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(161);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PARIZQ:
				{
				_localctx = new ParExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(150);
				match(PARIZQ);
				setState(151);
				expr(0);
				setState(152);
				match(PARDER);
				}
				break;
			case SUB:
				{
				_localctx = new UnariaExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(154);
				match(SUB);
				setState(155);
				expr(13);
				}
				break;
			case NUMBER:
				{
				_localctx = new NumExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(156);
				match(NUMBER);
				}
				break;
			case ID:
				{
				_localctx = new IdExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(157);
				match(ID);
				}
				break;
			case STRING:
				{
				_localctx = new StrExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(158);
				match(STRING);
				}
				break;
			case TRU:
			case FAL:
				{
				_localctx = new BoolExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(159);
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
			case NIL:
				{
				_localctx = new NilExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(160);
				match(NIL);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(186);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(184);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
					case 1:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(163);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(164);
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
						setState(165);
						((OpExprContext)_localctx).right = expr(13);
						}
						break;
					case 2:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(166);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(167);
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
						setState(168);
						((OpExprContext)_localctx).right = expr(12);
						}
						break;
					case 3:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(169);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(170);
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
						setState(171);
						((OpExprContext)_localctx).right = expr(11);
						}
						break;
					case 4:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(172);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(173);
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
						setState(174);
						((OpExprContext)_localctx).right = expr(10);
						}
						break;
					case 5:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(175);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(176);
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
						setState(177);
						((OpExprContext)_localctx).right = expr(9);
						}
						break;
					case 6:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(178);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(179);
						((OpExprContext)_localctx).op = match(AND);
						setState(180);
						((OpExprContext)_localctx).right = expr(8);
						}
						break;
					case 7:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(181);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(182);
						((OpExprContext)_localctx).op = match(OR);
						setState(183);
						((OpExprContext)_localctx).right = expr(7);
						}
						break;
					}
					} 
				}
				setState(188);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
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
		enterRule(_localctx, 34, RULE_tipo);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(189);
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
		case 16:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 12);
		case 1:
			return precpred(_ctx, 11);
		case 2:
			return precpred(_ctx, 10);
		case 3:
			return precpred(_ctx, 9);
		case 4:
			return precpred(_ctx, 8);
		case 5:
			return precpred(_ctx, 7);
		case 6:
			return precpred(_ctx, 6);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3.\u00c2\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\3\2\3\2\3\2\3\3\7\3+\n\3\f\3\16\3.\13\3\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\5\4:\n\4\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\7\3\7\3\7"+
		"\3\7\3\7\3\7\7\7I\n\7\f\7\16\7L\13\7\3\7\3\7\3\7\3\7\3\7\5\7S\n\7\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\6\t`\n\t\r\t\16\ta\3\t\5\te\n"+
		"\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3"+
		"\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3"+
		"\21\3\21\5\21\u0096\n\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\5\22\u00a4\n\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\7\22"+
		"\u00bb\n\22\f\22\16\22\u00be\13\22\3\23\3\23\3\23\2\3\"\24\2\4\6\b\n\f"+
		"\16\20\22\24\26\30\32\34\36 \"$\2\t\3\2\t\n\4\2!\"%%\3\2#$\4\2\35\35\37"+
		"\37\4\2\36\36  \3\2\27\30\3\2\3\7\2\u00cb\2&\3\2\2\2\4,\3\2\2\2\69\3\2"+
		"\2\2\b;\3\2\2\2\n@\3\2\2\2\fB\3\2\2\2\16T\3\2\2\2\20[\3\2\2\2\22h\3\2"+
		"\2\2\24m\3\2\2\2\26q\3\2\2\2\30x\3\2\2\2\32}\3\2\2\2\34\u0083\3\2\2\2"+
		"\36\u0087\3\2\2\2 \u0095\3\2\2\2\"\u00a3\3\2\2\2$\u00bf\3\2\2\2&\'\5\4"+
		"\3\2\'(\7\2\2\3(\3\3\2\2\2)+\5\6\4\2*)\3\2\2\2+.\3\2\2\2,*\3\2\2\2,-\3"+
		"\2\2\2-\5\3\2\2\2.,\3\2\2\2/:\5\b\5\2\60:\5\26\f\2\61:\5\30\r\2\62:\5"+
		"\32\16\2\63:\5\34\17\2\64:\5\f\7\2\65:\5\20\t\2\66:\5\36\20\2\67:\5 \21"+
		"\28:\5\n\6\29/\3\2\2\29\60\3\2\2\29\61\3\2\2\29\62\3\2\2\29\63\3\2\2\2"+
		"9\64\3\2\2\29\65\3\2\2\29\66\3\2\2\29\67\3\2\2\298\3\2\2\2:\7\3\2\2\2"+
		";<\7\13\2\2<=\7&\2\2=>\5\"\22\2>?\7\'\2\2?\t\3\2\2\2@A\7\23\2\2A\13\3"+
		"\2\2\2BC\7\f\2\2CD\5\"\22\2DE\7(\2\2EF\5\4\3\2FJ\7)\2\2GI\5\16\b\2HG\3"+
		"\2\2\2IL\3\2\2\2JH\3\2\2\2JK\3\2\2\2KR\3\2\2\2LJ\3\2\2\2MN\7\r\2\2NO\7"+
		"(\2\2OP\5\4\3\2PQ\7)\2\2QS\3\2\2\2RM\3\2\2\2RS\3\2\2\2S\r\3\2\2\2TU\7"+
		"\r\2\2UV\7\f\2\2VW\5\"\22\2WX\7(\2\2XY\5\4\3\2YZ\7)\2\2Z\17\3\2\2\2[\\"+
		"\7\17\2\2\\]\5\"\22\2]_\7(\2\2^`\5\22\n\2_^\3\2\2\2`a\3\2\2\2a_\3\2\2"+
		"\2ab\3\2\2\2bd\3\2\2\2ce\5\24\13\2dc\3\2\2\2de\3\2\2\2ef\3\2\2\2fg\7)"+
		"\2\2g\21\3\2\2\2hi\7\20\2\2ij\5\"\22\2jk\7*\2\2kl\5\4\3\2l\23\3\2\2\2"+
		"mn\7\21\2\2no\7*\2\2op\5\4\3\2p\25\3\2\2\2qr\7\22\2\2rs\7\26\2\2st\7*"+
		"\2\2tu\5$\23\2uv\7\34\2\2vw\5\"\22\2w\27\3\2\2\2xy\7\22\2\2yz\7\26\2\2"+
		"z{\7\34\2\2{|\5\"\22\2|\31\3\2\2\2}~\7\22\2\2~\177\7\26\2\2\177\u0080"+
		"\7*\2\2\u0080\u0081\5$\23\2\u0081\u0082\7+\2\2\u0082\33\3\2\2\2\u0083"+
		"\u0084\7\26\2\2\u0084\u0085\7\34\2\2\u0085\u0086\5\"\22\2\u0086\35\3\2"+
		"\2\2\u0087\u0088\7\16\2\2\u0088\u0089\5\"\22\2\u0089\u008a\7(\2\2\u008a"+
		"\u008b\5\4\3\2\u008b\u008c\7)\2\2\u008c\37\3\2\2\2\u008d\u008e\7\26\2"+
		"\2\u008e\u008f\7#\2\2\u008f\u0090\7\34\2\2\u0090\u0096\5\"\22\2\u0091"+
		"\u0092\7\26\2\2\u0092\u0093\7$\2\2\u0093\u0094\7\34\2\2\u0094\u0096\5"+
		"\"\22\2\u0095\u008d\3\2\2\2\u0095\u0091\3\2\2\2\u0096!\3\2\2\2\u0097\u0098"+
		"\b\22\1\2\u0098\u0099\7&\2\2\u0099\u009a\5\"\22\2\u009a\u009b\7\'\2\2"+
		"\u009b\u00a4\3\2\2\2\u009c\u009d\7$\2\2\u009d\u00a4\5\"\22\17\u009e\u00a4"+
		"\7\24\2\2\u009f\u00a4\7\26\2\2\u00a0\u00a4\7\25\2\2\u00a1\u00a4\t\2\2"+
		"\2\u00a2\u00a4\7\b\2\2\u00a3\u0097\3\2\2\2\u00a3\u009c\3\2\2\2\u00a3\u009e"+
		"\3\2\2\2\u00a3\u009f\3\2\2\2\u00a3\u00a0\3\2\2\2\u00a3\u00a1\3\2\2\2\u00a3"+
		"\u00a2\3\2\2\2\u00a4\u00bc\3\2\2\2\u00a5\u00a6\f\16\2\2\u00a6\u00a7\t"+
		"\3\2\2\u00a7\u00bb\5\"\22\17\u00a8\u00a9\f\r\2\2\u00a9\u00aa\t\4\2\2\u00aa"+
		"\u00bb\5\"\22\16\u00ab\u00ac\f\f\2\2\u00ac\u00ad\t\5\2\2\u00ad\u00bb\5"+
		"\"\22\r\u00ae\u00af\f\13\2\2\u00af\u00b0\t\6\2\2\u00b0\u00bb\5\"\22\f"+
		"\u00b1\u00b2\f\n\2\2\u00b2\u00b3\t\7\2\2\u00b3\u00bb\5\"\22\13\u00b4\u00b5"+
		"\f\t\2\2\u00b5\u00b6\7\33\2\2\u00b6\u00bb\5\"\22\n\u00b7\u00b8\f\b\2\2"+
		"\u00b8\u00b9\7\32\2\2\u00b9\u00bb\5\"\22\t\u00ba\u00a5\3\2\2\2\u00ba\u00a8"+
		"\3\2\2\2\u00ba\u00ab\3\2\2\2\u00ba\u00ae\3\2\2\2\u00ba\u00b1\3\2\2\2\u00ba"+
		"\u00b4\3\2\2\2\u00ba\u00b7\3\2\2\2\u00bb\u00be\3\2\2\2\u00bc\u00ba\3\2"+
		"\2\2\u00bc\u00bd\3\2\2\2\u00bd#\3\2\2\2\u00be\u00bc\3\2\2\2\u00bf\u00c0"+
		"\t\b\2\2\u00c0%\3\2\2\2\f,9JRad\u0095\u00a3\u00ba\u00bc";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
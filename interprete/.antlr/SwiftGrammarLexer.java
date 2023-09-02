// Generated from c:\Users\aldov\Desktop\Cursos Segundo Semestre 2023\Laboratorio Organización de Lenguajes y Compiladores 2 - copia\Proyectos\Proyecto 1\interprete\SwiftGrammar.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SwiftGrammarLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		INT=1, FLOAT=2, BOOL=3, CHARACTER=4, PSTRING=5, NIL=6, TRU=7, FAL=8, PRINT=9, 
		IF=10, ELSE=11, WHILE=12, FOR=13, GUARD=14, IN=15, SWITCH=16, CASE=17, 
		DEFAULT=18, VAR=19, LET=20, BREAK=21, RETURN=22, CONTINUE=23, FUNC=24, 
		COUNT=25, ISEMPTY=26, APPEND=27, REMOVE_LAST=28, REMOVE=29, STRUCT=30, 
		STRUCT_VAR=31, STRUCT_LET=32, AT=33, NUMBER=34, STRING=35, ID=36, DIF=37, 
		IG_IG=38, NOT=39, OR=40, AND=41, IG=42, MAY_IG=43, MEN_IG=44, MAYOR=45, 
		MENOR=46, MUL=47, DIV=48, ADD=49, SUB=50, MOD=51, PARIZQ=52, PARDER=53, 
		LLAVEIZQ=54, LLAVEDER=55, CORCHIZQ=56, CORCHDER=57, DOSPUNTOS=58, COMA=59, 
		PTCOMA=60, INTERROGACION=61, PUNTO=62, WHITESPACE=63, COMMENT=64, LINE_COMMENT=65;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "NIL", "TRU", "FAL", 
			"PRINT", "IF", "ELSE", "WHILE", "FOR", "GUARD", "IN", "SWITCH", "CASE", 
			"DEFAULT", "VAR", "LET", "BREAK", "RETURN", "CONTINUE", "FUNC", "COUNT", 
			"ISEMPTY", "APPEND", "REMOVE_LAST", "REMOVE", "STRUCT", "STRUCT_VAR", 
			"STRUCT_LET", "AT", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", 
			"OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", 
			"ADD", "SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "CORCHIZQ", 
			"CORCHDER", "DOSPUNTOS", "COMA", "PTCOMA", "INTERROGACION", "PUNTO", 
			"WHITESPACE", "COMMENT", "LINE_COMMENT", "ESC_SEQ"
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
			"'at'", null, null, null, "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", 
			"'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'%'", "'('", 
			"')'", "'{'", "'}'", "'['", "']'", "':'", "','", "';'", "'?'", "'.'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "NIL", "TRU", "FAL", 
			"PRINT", "IF", "ELSE", "WHILE", "FOR", "GUARD", "IN", "SWITCH", "CASE", 
			"DEFAULT", "VAR", "LET", "BREAK", "RETURN", "CONTINUE", "FUNC", "COUNT", 
			"ISEMPTY", "APPEND", "REMOVE_LAST", "REMOVE", "STRUCT", "STRUCT_VAR", 
			"STRUCT_LET", "AT", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", 
			"OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", 
			"ADD", "SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "CORCHIZQ", 
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


	public SwiftGrammarLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "SwiftGrammar.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2C\u01c5\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\f\3\f\3\f\3\f"+
		"\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3"+
		"\17\3\17\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3"+
		"\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3"+
		"\24\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3"+
		"\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\31\3"+
		"\31\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3"+
		"\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3"+
		"\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\36\3\36\3"+
		"\36\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3 \3 \3!\3!\3!\3!\3"+
		"!\3!\3\"\3\"\3\"\3#\6#\u014d\n#\r#\16#\u014e\3#\3#\6#\u0153\n#\r#\16#"+
		"\u0154\5#\u0157\n#\3$\3$\7$\u015b\n$\f$\16$\u015e\13$\3$\3$\3%\3%\7%\u0164"+
		"\n%\f%\16%\u0167\13%\3&\3&\3&\3\'\3\'\3\'\3(\3(\3)\3)\3)\3*\3*\3*\3+\3"+
		"+\3,\3,\3,\3-\3-\3-\3.\3.\3/\3/\3\60\3\60\3\61\3\61\3\62\3\62\3\63\3\63"+
		"\3\64\3\64\3\65\3\65\3\66\3\66\3\67\3\67\38\38\39\39\3:\3:\3;\3;\3<\3"+
		"<\3=\3=\3>\3>\3?\3?\3@\6@\u01a4\n@\r@\16@\u01a5\3@\3@\3A\3A\3A\3A\7A\u01ae"+
		"\nA\fA\16A\u01b1\13A\3A\3A\3A\3A\3A\3B\3B\3B\3B\7B\u01bc\nB\fB\16B\u01bf"+
		"\13B\3B\3B\3C\3C\3C\3\u01af\2D\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13"+
		"\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61"+
		"\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61"+
		"a\62c\63e\64g\65i\66k\67m8o9q:s;u<w=y>{?}@\177A\u0081B\u0083C\u0085\2"+
		"\3\2\t\3\2\62;\3\2$$\5\2C\\aac|\6\2\62;C\\aac|\6\2\13\f\17\17\"\"^^\4"+
		"\2\f\f\17\17\t\2\"#%%--/\60<<BB]_\2\u01cb\2\3\3\2\2\2\2\5\3\2\2\2\2\7"+
		"\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2"+
		"\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2"+
		"\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2"+
		"\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2"+
		"\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2"+
		"\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M"+
		"\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2"+
		"\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2"+
		"\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s"+
		"\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177"+
		"\3\2\2\2\2\u0081\3\2\2\2\2\u0083\3\2\2\2\3\u0087\3\2\2\2\5\u008b\3\2\2"+
		"\2\7\u0091\3\2\2\2\t\u0096\3\2\2\2\13\u00a0\3\2\2\2\r\u00a7\3\2\2\2\17"+
		"\u00ab\3\2\2\2\21\u00b0\3\2\2\2\23\u00b6\3\2\2\2\25\u00bc\3\2\2\2\27\u00bf"+
		"\3\2\2\2\31\u00c4\3\2\2\2\33\u00ca\3\2\2\2\35\u00ce\3\2\2\2\37\u00d4\3"+
		"\2\2\2!\u00d7\3\2\2\2#\u00de\3\2\2\2%\u00e3\3\2\2\2\'\u00eb\3\2\2\2)\u00ef"+
		"\3\2\2\2+\u00f3\3\2\2\2-\u00f9\3\2\2\2/\u0100\3\2\2\2\61\u0109\3\2\2\2"+
		"\63\u010e\3\2\2\2\65\u0114\3\2\2\2\67\u011c\3\2\2\29\u0123\3\2\2\2;\u012e"+
		"\3\2\2\2=\u0135\3\2\2\2?\u013c\3\2\2\2A\u0142\3\2\2\2C\u0148\3\2\2\2E"+
		"\u014c\3\2\2\2G\u0158\3\2\2\2I\u0161\3\2\2\2K\u0168\3\2\2\2M\u016b\3\2"+
		"\2\2O\u016e\3\2\2\2Q\u0170\3\2\2\2S\u0173\3\2\2\2U\u0176\3\2\2\2W\u0178"+
		"\3\2\2\2Y\u017b\3\2\2\2[\u017e\3\2\2\2]\u0180\3\2\2\2_\u0182\3\2\2\2a"+
		"\u0184\3\2\2\2c\u0186\3\2\2\2e\u0188\3\2\2\2g\u018a\3\2\2\2i\u018c\3\2"+
		"\2\2k\u018e\3\2\2\2m\u0190\3\2\2\2o\u0192\3\2\2\2q\u0194\3\2\2\2s\u0196"+
		"\3\2\2\2u\u0198\3\2\2\2w\u019a\3\2\2\2y\u019c\3\2\2\2{\u019e\3\2\2\2}"+
		"\u01a0\3\2\2\2\177\u01a3\3\2\2\2\u0081\u01a9\3\2\2\2\u0083\u01b7\3\2\2"+
		"\2\u0085\u01c2\3\2\2\2\u0087\u0088\7k\2\2\u0088\u0089\7p\2\2\u0089\u008a"+
		"\7v\2\2\u008a\4\3\2\2\2\u008b\u008c\7h\2\2\u008c\u008d\7n\2\2\u008d\u008e"+
		"\7q\2\2\u008e\u008f\7c\2\2\u008f\u0090\7v\2\2\u0090\6\3\2\2\2\u0091\u0092"+
		"\7d\2\2\u0092\u0093\7q\2\2\u0093\u0094\7q\2\2\u0094\u0095\7n\2\2\u0095"+
		"\b\3\2\2\2\u0096\u0097\7e\2\2\u0097\u0098\7j\2\2\u0098\u0099\7c\2\2\u0099"+
		"\u009a\7t\2\2\u009a\u009b\7c\2\2\u009b\u009c\7e\2\2\u009c\u009d\7v\2\2"+
		"\u009d\u009e\7g\2\2\u009e\u009f\7t\2\2\u009f\n\3\2\2\2\u00a0\u00a1\7U"+
		"\2\2\u00a1\u00a2\7v\2\2\u00a2\u00a3\7t\2\2\u00a3\u00a4\7k\2\2\u00a4\u00a5"+
		"\7p\2\2\u00a5\u00a6\7i\2\2\u00a6\f\3\2\2\2\u00a7\u00a8\7p\2\2\u00a8\u00a9"+
		"\7k\2\2\u00a9\u00aa\7n\2\2\u00aa\16\3\2\2\2\u00ab\u00ac\7v\2\2\u00ac\u00ad"+
		"\7t\2\2\u00ad\u00ae\7w\2\2\u00ae\u00af\7g\2\2\u00af\20\3\2\2\2\u00b0\u00b1"+
		"\7h\2\2\u00b1\u00b2\7c\2\2\u00b2\u00b3\7n\2\2\u00b3\u00b4\7u\2\2\u00b4"+
		"\u00b5\7g\2\2\u00b5\22\3\2\2\2\u00b6\u00b7\7r\2\2\u00b7\u00b8\7t\2\2\u00b8"+
		"\u00b9\7k\2\2\u00b9\u00ba\7p\2\2\u00ba\u00bb\7v\2\2\u00bb\24\3\2\2\2\u00bc"+
		"\u00bd\7k\2\2\u00bd\u00be\7h\2\2\u00be\26\3\2\2\2\u00bf\u00c0\7g\2\2\u00c0"+
		"\u00c1\7n\2\2\u00c1\u00c2\7u\2\2\u00c2\u00c3\7g\2\2\u00c3\30\3\2\2\2\u00c4"+
		"\u00c5\7y\2\2\u00c5\u00c6\7j\2\2\u00c6\u00c7\7k\2\2\u00c7\u00c8\7n\2\2"+
		"\u00c8\u00c9\7g\2\2\u00c9\32\3\2\2\2\u00ca\u00cb\7h\2\2\u00cb\u00cc\7"+
		"q\2\2\u00cc\u00cd\7t\2\2\u00cd\34\3\2\2\2\u00ce\u00cf\7i\2\2\u00cf\u00d0"+
		"\7w\2\2\u00d0\u00d1\7c\2\2\u00d1\u00d2\7t\2\2\u00d2\u00d3\7f\2\2\u00d3"+
		"\36\3\2\2\2\u00d4\u00d5\7k\2\2\u00d5\u00d6\7p\2\2\u00d6 \3\2\2\2\u00d7"+
		"\u00d8\7u\2\2\u00d8\u00d9\7y\2\2\u00d9\u00da\7k\2\2\u00da\u00db\7v\2\2"+
		"\u00db\u00dc\7e\2\2\u00dc\u00dd\7j\2\2\u00dd\"\3\2\2\2\u00de\u00df\7e"+
		"\2\2\u00df\u00e0\7c\2\2\u00e0\u00e1\7u\2\2\u00e1\u00e2\7g\2\2\u00e2$\3"+
		"\2\2\2\u00e3\u00e4\7f\2\2\u00e4\u00e5\7g\2\2\u00e5\u00e6\7h\2\2\u00e6"+
		"\u00e7\7c\2\2\u00e7\u00e8\7w\2\2\u00e8\u00e9\7n\2\2\u00e9\u00ea\7v\2\2"+
		"\u00ea&\3\2\2\2\u00eb\u00ec\7x\2\2\u00ec\u00ed\7c\2\2\u00ed\u00ee\7t\2"+
		"\2\u00ee(\3\2\2\2\u00ef\u00f0\7n\2\2\u00f0\u00f1\7g\2\2\u00f1\u00f2\7"+
		"v\2\2\u00f2*\3\2\2\2\u00f3\u00f4\7d\2\2\u00f4\u00f5\7t\2\2\u00f5\u00f6"+
		"\7g\2\2\u00f6\u00f7\7c\2\2\u00f7\u00f8\7m\2\2\u00f8,\3\2\2\2\u00f9\u00fa"+
		"\7t\2\2\u00fa\u00fb\7g\2\2\u00fb\u00fc\7v\2\2\u00fc\u00fd\7w\2\2\u00fd"+
		"\u00fe\7t\2\2\u00fe\u00ff\7p\2\2\u00ff.\3\2\2\2\u0100\u0101\7e\2\2\u0101"+
		"\u0102\7q\2\2\u0102\u0103\7p\2\2\u0103\u0104\7v\2\2\u0104\u0105\7k\2\2"+
		"\u0105\u0106\7p\2\2\u0106\u0107\7w\2\2\u0107\u0108\7g\2\2\u0108\60\3\2"+
		"\2\2\u0109\u010a\7h\2\2\u010a\u010b\7w\2\2\u010b\u010c\7p\2\2\u010c\u010d"+
		"\7e\2\2\u010d\62\3\2\2\2\u010e\u010f\7e\2\2\u010f\u0110\7q\2\2\u0110\u0111"+
		"\7w\2\2\u0111\u0112\7p\2\2\u0112\u0113\7v\2\2\u0113\64\3\2\2\2\u0114\u0115"+
		"\7K\2\2\u0115\u0116\7u\2\2\u0116\u0117\7G\2\2\u0117\u0118\7o\2\2\u0118"+
		"\u0119\7r\2\2\u0119\u011a\7v\2\2\u011a\u011b\7{\2\2\u011b\66\3\2\2\2\u011c"+
		"\u011d\7c\2\2\u011d\u011e\7r\2\2\u011e\u011f\7r\2\2\u011f\u0120\7g\2\2"+
		"\u0120\u0121\7p\2\2\u0121\u0122\7f\2\2\u01228\3\2\2\2\u0123\u0124\7t\2"+
		"\2\u0124\u0125\7g\2\2\u0125\u0126\7o\2\2\u0126\u0127\7q\2\2\u0127\u0128"+
		"\7x\2\2\u0128\u0129\7g\2\2\u0129\u012a\7N\2\2\u012a\u012b\7c\2\2\u012b"+
		"\u012c\7u\2\2\u012c\u012d\7v\2\2\u012d:\3\2\2\2\u012e\u012f\7t\2\2\u012f"+
		"\u0130\7g\2\2\u0130\u0131\7o\2\2\u0131\u0132\7q\2\2\u0132\u0133\7x\2\2"+
		"\u0133\u0134\7g\2\2\u0134<\3\2\2\2\u0135\u0136\7u\2\2\u0136\u0137\7v\2"+
		"\2\u0137\u0138\7t\2\2\u0138\u0139\7w\2\2\u0139\u013a\7e\2\2\u013a\u013b"+
		"\7v\2\2\u013b>\3\2\2\2\u013c\u013d\7x\2\2\u013d\u013e\7c\2\2\u013e\u013f"+
		"\7t\2\2\u013f\u0140\7u\2\2\u0140\u0141\7v\2\2\u0141@\3\2\2\2\u0142\u0143"+
		"\7n\2\2\u0143\u0144\7g\2\2\u0144\u0145\7v\2\2\u0145\u0146\7u\2\2\u0146"+
		"\u0147\7v\2\2\u0147B\3\2\2\2\u0148\u0149\7c\2\2\u0149\u014a\7v\2\2\u014a"+
		"D\3\2\2\2\u014b\u014d\t\2\2\2\u014c\u014b\3\2\2\2\u014d\u014e\3\2\2\2"+
		"\u014e\u014c\3\2\2\2\u014e\u014f\3\2\2\2\u014f\u0156\3\2\2\2\u0150\u0152"+
		"\7\60\2\2\u0151\u0153\t\2\2\2\u0152\u0151\3\2\2\2\u0153\u0154\3\2\2\2"+
		"\u0154\u0152\3\2\2\2\u0154\u0155\3\2\2\2\u0155\u0157\3\2\2\2\u0156\u0150"+
		"\3\2\2\2\u0156\u0157\3\2\2\2\u0157F\3\2\2\2\u0158\u015c\7$\2\2\u0159\u015b"+
		"\n\3\2\2\u015a\u0159\3\2\2\2\u015b\u015e\3\2\2\2\u015c\u015a\3\2\2\2\u015c"+
		"\u015d\3\2\2\2\u015d\u015f\3\2\2\2\u015e\u015c\3\2\2\2\u015f\u0160\7$"+
		"\2\2\u0160H\3\2\2\2\u0161\u0165\t\4\2\2\u0162\u0164\t\5\2\2\u0163\u0162"+
		"\3\2\2\2\u0164\u0167\3\2\2\2\u0165\u0163\3\2\2\2\u0165\u0166\3\2\2\2\u0166"+
		"J\3\2\2\2\u0167\u0165\3\2\2\2\u0168\u0169\7#\2\2\u0169\u016a\7?\2\2\u016a"+
		"L\3\2\2\2\u016b\u016c\7?\2\2\u016c\u016d\7?\2\2\u016dN\3\2\2\2\u016e\u016f"+
		"\7#\2\2\u016fP\3\2\2\2\u0170\u0171\7~\2\2\u0171\u0172\7~\2\2\u0172R\3"+
		"\2\2\2\u0173\u0174\7(\2\2\u0174\u0175\7(\2\2\u0175T\3\2\2\2\u0176\u0177"+
		"\7?\2\2\u0177V\3\2\2\2\u0178\u0179\7@\2\2\u0179\u017a\7?\2\2\u017aX\3"+
		"\2\2\2\u017b\u017c\7>\2\2\u017c\u017d\7?\2\2\u017dZ\3\2\2\2\u017e\u017f"+
		"\7@\2\2\u017f\\\3\2\2\2\u0180\u0181\7>\2\2\u0181^\3\2\2\2\u0182\u0183"+
		"\7,\2\2\u0183`\3\2\2\2\u0184\u0185\7\61\2\2\u0185b\3\2\2\2\u0186\u0187"+
		"\7-\2\2\u0187d\3\2\2\2\u0188\u0189\7/\2\2\u0189f\3\2\2\2\u018a\u018b\7"+
		"\'\2\2\u018bh\3\2\2\2\u018c\u018d\7*\2\2\u018dj\3\2\2\2\u018e\u018f\7"+
		"+\2\2\u018fl\3\2\2\2\u0190\u0191\7}\2\2\u0191n\3\2\2\2\u0192\u0193\7\177"+
		"\2\2\u0193p\3\2\2\2\u0194\u0195\7]\2\2\u0195r\3\2\2\2\u0196\u0197\7_\2"+
		"\2\u0197t\3\2\2\2\u0198\u0199\7<\2\2\u0199v\3\2\2\2\u019a\u019b\7.\2\2"+
		"\u019bx\3\2\2\2\u019c\u019d\7=\2\2\u019dz\3\2\2\2\u019e\u019f\7A\2\2\u019f"+
		"|\3\2\2\2\u01a0\u01a1\7\60\2\2\u01a1~\3\2\2\2\u01a2\u01a4\t\6\2\2\u01a3"+
		"\u01a2\3\2\2\2\u01a4\u01a5\3\2\2\2\u01a5\u01a3\3\2\2\2\u01a5\u01a6\3\2"+
		"\2\2\u01a6\u01a7\3\2\2\2\u01a7\u01a8\b@\2\2\u01a8\u0080\3\2\2\2\u01a9"+
		"\u01aa\7\61\2\2\u01aa\u01ab\7,\2\2\u01ab\u01af\3\2\2\2\u01ac\u01ae\13"+
		"\2\2\2\u01ad\u01ac\3\2\2\2\u01ae\u01b1\3\2\2\2\u01af\u01b0\3\2\2\2\u01af"+
		"\u01ad\3\2\2\2\u01b0\u01b2\3\2\2\2\u01b1\u01af\3\2\2\2\u01b2\u01b3\7,"+
		"\2\2\u01b3\u01b4\7\61\2\2\u01b4\u01b5\3\2\2\2\u01b5\u01b6\bA\2\2\u01b6"+
		"\u0082\3\2\2\2\u01b7\u01b8\7\61\2\2\u01b8\u01b9\7\61\2\2\u01b9\u01bd\3"+
		"\2\2\2\u01ba\u01bc\n\7\2\2\u01bb\u01ba\3\2\2\2\u01bc\u01bf\3\2\2\2\u01bd"+
		"\u01bb\3\2\2\2\u01bd\u01be\3\2\2\2\u01be\u01c0\3\2\2\2\u01bf\u01bd\3\2"+
		"\2\2\u01c0\u01c1\bB\2\2\u01c1\u0084\3\2\2\2\u01c2\u01c3\7^\2\2\u01c3\u01c4"+
		"\t\b\2\2\u01c4\u0086\3\2\2\2\13\2\u014e\u0154\u0156\u015c\u0165\u01a5"+
		"\u01af\u01bd\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
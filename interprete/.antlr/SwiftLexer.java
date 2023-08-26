// Generated from c:\Users\aldov\Desktop\Cursos Segundo Semestre 2023\Laboratorio Organizaci�n de Lenguajes y Compiladores 2 - copia\Proyectos\Proyecto 1\interprete\SwiftLexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SwiftLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		INT=1, FLOAT=2, BOOL=3, CHARACTER=4, PSTRING=5, NIL=6, TRU=7, FAL=8, PRINT=9, 
		IF=10, ELSE=11, WHILE=12, SWITCH=13, CASE=14, DEFAULT=15, VAR=16, BREAK=17, 
		RETURN=18, FUNC=19, NUMBER=20, STRING=21, ID=22, DIF=23, IG_IG=24, NOT=25, 
		OR=26, AND=27, IG=28, MAY_IG=29, MEN_IG=30, MAYOR=31, MENOR=32, MUL=33, 
		DIV=34, ADD=35, SUB=36, MOD=37, PARIZQ=38, PARDER=39, LLAVEIZQ=40, LLAVEDER=41, 
		CORCHIZQ=42, CORCHDER=43, DOSPUNTOS=44, COMA=45, PTCOMA=46, INTERROGACION=47, 
		WHITESPACE=48, COMMENT=49, LINE_COMMENT=50;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "NIL", "TRU", "FAL", 
			"PRINT", "IF", "ELSE", "WHILE", "SWITCH", "CASE", "DEFAULT", "VAR", "BREAK", 
			"RETURN", "FUNC", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", 
			"AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", 
			"SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "CORCHIZQ", 
			"CORCHDER", "DOSPUNTOS", "COMA", "PTCOMA", "INTERROGACION", "WHITESPACE", 
			"COMMENT", "LINE_COMMENT", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'int'", "'float'", "'bool'", "'character'", "'String'", "'nil'", 
			"'true'", "'false'", "'print'", "'if'", "'else'", "'while'", "'switch'", 
			"'case'", "'default'", "'var'", "'break'", "'return'", "'func'", null, 
			null, null, "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", "'>='", "'<='", 
			"'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'%'", "'('", "')'", "'{'", 
			"'}'", "'['", "']'", "':'", "','", "';'", "'?'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "NIL", "TRU", "FAL", 
			"PRINT", "IF", "ELSE", "WHILE", "SWITCH", "CASE", "DEFAULT", "VAR", "BREAK", 
			"RETURN", "FUNC", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", 
			"AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", 
			"SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "CORCHIZQ", 
			"CORCHDER", "DOSPUNTOS", "COMA", "PTCOMA", "INTERROGACION", "WHITESPACE", 
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


	public SwiftLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "SwiftLexer.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\64\u014e\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64"+
		"\t\64\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5"+
		"\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3"+
		"\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n"+
		"\3\n\3\n\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3"+
		"\16\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3"+
		"\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3"+
		"\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3"+
		"\25\6\25\u00d8\n\25\r\25\16\25\u00d9\3\25\3\25\6\25\u00de\n\25\r\25\16"+
		"\25\u00df\5\25\u00e2\n\25\3\26\3\26\7\26\u00e6\n\26\f\26\16\26\u00e9\13"+
		"\26\3\26\3\26\3\27\3\27\7\27\u00ef\n\27\f\27\16\27\u00f2\13\27\3\30\3"+
		"\30\3\30\3\31\3\31\3\31\3\32\3\32\3\33\3\33\3\33\3\34\3\34\3\34\3\35\3"+
		"\35\3\36\3\36\3\36\3\37\3\37\3\37\3 \3 \3!\3!\3\"\3\"\3#\3#\3$\3$\3%\3"+
		"%\3&\3&\3\'\3\'\3(\3(\3)\3)\3*\3*\3+\3+\3,\3,\3-\3-\3.\3.\3/\3/\3\60\3"+
		"\60\3\61\6\61\u012d\n\61\r\61\16\61\u012e\3\61\3\61\3\62\3\62\3\62\3\62"+
		"\7\62\u0137\n\62\f\62\16\62\u013a\13\62\3\62\3\62\3\62\3\62\3\62\3\63"+
		"\3\63\3\63\3\63\7\63\u0145\n\63\f\63\16\63\u0148\13\63\3\63\3\63\3\64"+
		"\3\64\3\64\3\u0138\2\65\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f"+
		"\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63"+
		"\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62"+
		"c\63e\64g\2\3\2\t\3\2\62;\3\2$$\5\2C\\aac|\6\2\62;C\\aac|\6\2\13\f\17"+
		"\17\"\"^^\4\2\f\f\17\17\t\2\"#%%--/\60<<BB]_\2\u0154\2\3\3\2\2\2\2\5\3"+
		"\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2"+
		"\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3"+
		"\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'"+
		"\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63"+
		"\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2"+
		"?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3"+
		"\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2"+
		"\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2"+
		"e\3\2\2\2\3i\3\2\2\2\5m\3\2\2\2\7s\3\2\2\2\tx\3\2\2\2\13\u0082\3\2\2\2"+
		"\r\u0089\3\2\2\2\17\u008d\3\2\2\2\21\u0092\3\2\2\2\23\u0098\3\2\2\2\25"+
		"\u009e\3\2\2\2\27\u00a1\3\2\2\2\31\u00a6\3\2\2\2\33\u00ac\3\2\2\2\35\u00b3"+
		"\3\2\2\2\37\u00b8\3\2\2\2!\u00c0\3\2\2\2#\u00c4\3\2\2\2%\u00ca\3\2\2\2"+
		"\'\u00d1\3\2\2\2)\u00d7\3\2\2\2+\u00e3\3\2\2\2-\u00ec\3\2\2\2/\u00f3\3"+
		"\2\2\2\61\u00f6\3\2\2\2\63\u00f9\3\2\2\2\65\u00fb\3\2\2\2\67\u00fe\3\2"+
		"\2\29\u0101\3\2\2\2;\u0103\3\2\2\2=\u0106\3\2\2\2?\u0109\3\2\2\2A\u010b"+
		"\3\2\2\2C\u010d\3\2\2\2E\u010f\3\2\2\2G\u0111\3\2\2\2I\u0113\3\2\2\2K"+
		"\u0115\3\2\2\2M\u0117\3\2\2\2O\u0119\3\2\2\2Q\u011b\3\2\2\2S\u011d\3\2"+
		"\2\2U\u011f\3\2\2\2W\u0121\3\2\2\2Y\u0123\3\2\2\2[\u0125\3\2\2\2]\u0127"+
		"\3\2\2\2_\u0129\3\2\2\2a\u012c\3\2\2\2c\u0132\3\2\2\2e\u0140\3\2\2\2g"+
		"\u014b\3\2\2\2ij\7k\2\2jk\7p\2\2kl\7v\2\2l\4\3\2\2\2mn\7h\2\2no\7n\2\2"+
		"op\7q\2\2pq\7c\2\2qr\7v\2\2r\6\3\2\2\2st\7d\2\2tu\7q\2\2uv\7q\2\2vw\7"+
		"n\2\2w\b\3\2\2\2xy\7e\2\2yz\7j\2\2z{\7c\2\2{|\7t\2\2|}\7c\2\2}~\7e\2\2"+
		"~\177\7v\2\2\177\u0080\7g\2\2\u0080\u0081\7t\2\2\u0081\n\3\2\2\2\u0082"+
		"\u0083\7U\2\2\u0083\u0084\7v\2\2\u0084\u0085\7t\2\2\u0085\u0086\7k\2\2"+
		"\u0086\u0087\7p\2\2\u0087\u0088\7i\2\2\u0088\f\3\2\2\2\u0089\u008a\7p"+
		"\2\2\u008a\u008b\7k\2\2\u008b\u008c\7n\2\2\u008c\16\3\2\2\2\u008d\u008e"+
		"\7v\2\2\u008e\u008f\7t\2\2\u008f\u0090\7w\2\2\u0090\u0091\7g\2\2\u0091"+
		"\20\3\2\2\2\u0092\u0093\7h\2\2\u0093\u0094\7c\2\2\u0094\u0095\7n\2\2\u0095"+
		"\u0096\7u\2\2\u0096\u0097\7g\2\2\u0097\22\3\2\2\2\u0098\u0099\7r\2\2\u0099"+
		"\u009a\7t\2\2\u009a\u009b\7k\2\2\u009b\u009c\7p\2\2\u009c\u009d\7v\2\2"+
		"\u009d\24\3\2\2\2\u009e\u009f\7k\2\2\u009f\u00a0\7h\2\2\u00a0\26\3\2\2"+
		"\2\u00a1\u00a2\7g\2\2\u00a2\u00a3\7n\2\2\u00a3\u00a4\7u\2\2\u00a4\u00a5"+
		"\7g\2\2\u00a5\30\3\2\2\2\u00a6\u00a7\7y\2\2\u00a7\u00a8\7j\2\2\u00a8\u00a9"+
		"\7k\2\2\u00a9\u00aa\7n\2\2\u00aa\u00ab\7g\2\2\u00ab\32\3\2\2\2\u00ac\u00ad"+
		"\7u\2\2\u00ad\u00ae\7y\2\2\u00ae\u00af\7k\2\2\u00af\u00b0\7v\2\2\u00b0"+
		"\u00b1\7e\2\2\u00b1\u00b2\7j\2\2\u00b2\34\3\2\2\2\u00b3\u00b4\7e\2\2\u00b4"+
		"\u00b5\7c\2\2\u00b5\u00b6\7u\2\2\u00b6\u00b7\7g\2\2\u00b7\36\3\2\2\2\u00b8"+
		"\u00b9\7f\2\2\u00b9\u00ba\7g\2\2\u00ba\u00bb\7h\2\2\u00bb\u00bc\7c\2\2"+
		"\u00bc\u00bd\7w\2\2\u00bd\u00be\7n\2\2\u00be\u00bf\7v\2\2\u00bf \3\2\2"+
		"\2\u00c0\u00c1\7x\2\2\u00c1\u00c2\7c\2\2\u00c2\u00c3\7t\2\2\u00c3\"\3"+
		"\2\2\2\u00c4\u00c5\7d\2\2\u00c5\u00c6\7t\2\2\u00c6\u00c7\7g\2\2\u00c7"+
		"\u00c8\7c\2\2\u00c8\u00c9\7m\2\2\u00c9$\3\2\2\2\u00ca\u00cb\7t\2\2\u00cb"+
		"\u00cc\7g\2\2\u00cc\u00cd\7v\2\2\u00cd\u00ce\7w\2\2\u00ce\u00cf\7t\2\2"+
		"\u00cf\u00d0\7p\2\2\u00d0&\3\2\2\2\u00d1\u00d2\7h\2\2\u00d2\u00d3\7w\2"+
		"\2\u00d3\u00d4\7p\2\2\u00d4\u00d5\7e\2\2\u00d5(\3\2\2\2\u00d6\u00d8\t"+
		"\2\2\2\u00d7\u00d6\3\2\2\2\u00d8\u00d9\3\2\2\2\u00d9\u00d7\3\2\2\2\u00d9"+
		"\u00da\3\2\2\2\u00da\u00e1\3\2\2\2\u00db\u00dd\7\60\2\2\u00dc\u00de\t"+
		"\2\2\2\u00dd\u00dc\3\2\2\2\u00de\u00df\3\2\2\2\u00df\u00dd\3\2\2\2\u00df"+
		"\u00e0\3\2\2\2\u00e0\u00e2\3\2\2\2\u00e1\u00db\3\2\2\2\u00e1\u00e2\3\2"+
		"\2\2\u00e2*\3\2\2\2\u00e3\u00e7\7$\2\2\u00e4\u00e6\n\3\2\2\u00e5\u00e4"+
		"\3\2\2\2\u00e6\u00e9\3\2\2\2\u00e7\u00e5\3\2\2\2\u00e7\u00e8\3\2\2\2\u00e8"+
		"\u00ea\3\2\2\2\u00e9\u00e7\3\2\2\2\u00ea\u00eb\7$\2\2\u00eb,\3\2\2\2\u00ec"+
		"\u00f0\t\4\2\2\u00ed\u00ef\t\5\2\2\u00ee\u00ed\3\2\2\2\u00ef\u00f2\3\2"+
		"\2\2\u00f0\u00ee\3\2\2\2\u00f0\u00f1\3\2\2\2\u00f1.\3\2\2\2\u00f2\u00f0"+
		"\3\2\2\2\u00f3\u00f4\7#\2\2\u00f4\u00f5\7?\2\2\u00f5\60\3\2\2\2\u00f6"+
		"\u00f7\7?\2\2\u00f7\u00f8\7?\2\2\u00f8\62\3\2\2\2\u00f9\u00fa\7#\2\2\u00fa"+
		"\64\3\2\2\2\u00fb\u00fc\7~\2\2\u00fc\u00fd\7~\2\2\u00fd\66\3\2\2\2\u00fe"+
		"\u00ff\7(\2\2\u00ff\u0100\7(\2\2\u01008\3\2\2\2\u0101\u0102\7?\2\2\u0102"+
		":\3\2\2\2\u0103\u0104\7@\2\2\u0104\u0105\7?\2\2\u0105<\3\2\2\2\u0106\u0107"+
		"\7>\2\2\u0107\u0108\7?\2\2\u0108>\3\2\2\2\u0109\u010a\7@\2\2\u010a@\3"+
		"\2\2\2\u010b\u010c\7>\2\2\u010cB\3\2\2\2\u010d\u010e\7,\2\2\u010eD\3\2"+
		"\2\2\u010f\u0110\7\61\2\2\u0110F\3\2\2\2\u0111\u0112\7-\2\2\u0112H\3\2"+
		"\2\2\u0113\u0114\7/\2\2\u0114J\3\2\2\2\u0115\u0116\7\'\2\2\u0116L\3\2"+
		"\2\2\u0117\u0118\7*\2\2\u0118N\3\2\2\2\u0119\u011a\7+\2\2\u011aP\3\2\2"+
		"\2\u011b\u011c\7}\2\2\u011cR\3\2\2\2\u011d\u011e\7\177\2\2\u011eT\3\2"+
		"\2\2\u011f\u0120\7]\2\2\u0120V\3\2\2\2\u0121\u0122\7_\2\2\u0122X\3\2\2"+
		"\2\u0123\u0124\7<\2\2\u0124Z\3\2\2\2\u0125\u0126\7.\2\2\u0126\\\3\2\2"+
		"\2\u0127\u0128\7=\2\2\u0128^\3\2\2\2\u0129\u012a\7A\2\2\u012a`\3\2\2\2"+
		"\u012b\u012d\t\6\2\2\u012c\u012b\3\2\2\2\u012d\u012e\3\2\2\2\u012e\u012c"+
		"\3\2\2\2\u012e\u012f\3\2\2\2\u012f\u0130\3\2\2\2\u0130\u0131\b\61\2\2"+
		"\u0131b\3\2\2\2\u0132\u0133\7\61\2\2\u0133\u0134\7,\2\2\u0134\u0138\3"+
		"\2\2\2\u0135\u0137\13\2\2\2\u0136\u0135\3\2\2\2\u0137\u013a\3\2\2\2\u0138"+
		"\u0139\3\2\2\2\u0138\u0136\3\2\2\2\u0139\u013b\3\2\2\2\u013a\u0138\3\2"+
		"\2\2\u013b\u013c\7,\2\2\u013c\u013d\7\61\2\2\u013d\u013e\3\2\2\2\u013e"+
		"\u013f\b\62\2\2\u013fd\3\2\2\2\u0140\u0141\7\61\2\2\u0141\u0142\7\61\2"+
		"\2\u0142\u0146\3\2\2\2\u0143\u0145\n\7\2\2\u0144\u0143\3\2\2\2\u0145\u0148"+
		"\3\2\2\2\u0146\u0144\3\2\2\2\u0146\u0147\3\2\2\2\u0147\u0149\3\2\2\2\u0148"+
		"\u0146\3\2\2\2\u0149\u014a\b\63\2\2\u014af\3\2\2\2\u014b\u014c\7^\2\2"+
		"\u014c\u014d\t\b\2\2\u014dh\3\2\2\2\13\2\u00d9\u00df\u00e1\u00e7\u00f0"+
		"\u012e\u0138\u0146\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
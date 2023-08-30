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
		IF=10, ELSE=11, WHILE=12, FOR=13, IN=14, SWITCH=15, CASE=16, DEFAULT=17, 
		VAR=18, BREAK=19, RETURN=20, FUNC=21, COUNT=22, ISEMPTY=23, APPEND=24, 
		REMOVE_LAST=25, REMOVE=26, AT=27, NUMBER=28, STRING=29, ID=30, DIF=31, 
		IG_IG=32, NOT=33, OR=34, AND=35, IG=36, MAY_IG=37, MEN_IG=38, MAYOR=39, 
		MENOR=40, MUL=41, DIV=42, ADD=43, SUB=44, MOD=45, PARIZQ=46, PARDER=47, 
		LLAVEIZQ=48, LLAVEDER=49, CORCHIZQ=50, CORCHDER=51, DOSPUNTOS=52, COMA=53, 
		PTCOMA=54, INTERROGACION=55, PUNTO=56, WHITESPACE=57, COMMENT=58, LINE_COMMENT=59;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "NIL", "TRU", "FAL", 
			"PRINT", "IF", "ELSE", "WHILE", "FOR", "IN", "SWITCH", "CASE", "DEFAULT", 
			"VAR", "BREAK", "RETURN", "FUNC", "COUNT", "ISEMPTY", "APPEND", "REMOVE_LAST", 
			"REMOVE", "AT", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", 
			"AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", 
			"SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "CORCHIZQ", 
			"CORCHDER", "DOSPUNTOS", "COMA", "PTCOMA", "INTERROGACION", "PUNTO", 
			"WHITESPACE", "COMMENT", "LINE_COMMENT", "ESC_SEQ"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2=\u0193\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7"+
		"\3\7\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16"+
		"\3\16\3\16\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21"+
		"\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23"+
		"\3\23\3\24\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\25"+
		"\3\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30"+
		"\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\32\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\34\3\34\3\34\3\35\6\35\u011b\n\35\r\35\16\35\u011c\3\35\3"+
		"\35\6\35\u0121\n\35\r\35\16\35\u0122\5\35\u0125\n\35\3\36\3\36\7\36\u0129"+
		"\n\36\f\36\16\36\u012c\13\36\3\36\3\36\3\37\3\37\7\37\u0132\n\37\f\37"+
		"\16\37\u0135\13\37\3 \3 \3 \3!\3!\3!\3\"\3\"\3#\3#\3#\3$\3$\3$\3%\3%\3"+
		"&\3&\3&\3\'\3\'\3\'\3(\3(\3)\3)\3*\3*\3+\3+\3,\3,\3-\3-\3.\3.\3/\3/\3"+
		"\60\3\60\3\61\3\61\3\62\3\62\3\63\3\63\3\64\3\64\3\65\3\65\3\66\3\66\3"+
		"\67\3\67\38\38\39\39\3:\6:\u0172\n:\r:\16:\u0173\3:\3:\3;\3;\3;\3;\7;"+
		"\u017c\n;\f;\16;\u017f\13;\3;\3;\3;\3;\3;\3<\3<\3<\3<\7<\u018a\n<\f<\16"+
		"<\u018d\13<\3<\3<\3=\3=\3=\3\u017d\2>\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21"+
		"\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30"+
		"/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.["+
		"/]\60_\61a\62c\63e\64g\65i\66k\67m8o9q:s;u<w=y\2\3\2\t\3\2\62;\3\2$$\5"+
		"\2C\\aac|\6\2\62;C\\aac|\6\2\13\f\17\17\"\"^^\4\2\f\f\17\17\t\2\"#%%-"+
		"-/\60<<BB]_\2\u0199\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2"+
		"\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3"+
		"\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2"+
		"\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2"+
		"\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2"+
		"\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2"+
		"\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q"+
		"\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2"+
		"\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2"+
		"\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2w"+
		"\3\2\2\2\3{\3\2\2\2\5\177\3\2\2\2\7\u0085\3\2\2\2\t\u008a\3\2\2\2\13\u0094"+
		"\3\2\2\2\r\u009b\3\2\2\2\17\u009f\3\2\2\2\21\u00a4\3\2\2\2\23\u00aa\3"+
		"\2\2\2\25\u00b0\3\2\2\2\27\u00b3\3\2\2\2\31\u00b8\3\2\2\2\33\u00be\3\2"+
		"\2\2\35\u00c2\3\2\2\2\37\u00c5\3\2\2\2!\u00cc\3\2\2\2#\u00d1\3\2\2\2%"+
		"\u00d9\3\2\2\2\'\u00dd\3\2\2\2)\u00e3\3\2\2\2+\u00ea\3\2\2\2-\u00ef\3"+
		"\2\2\2/\u00f5\3\2\2\2\61\u00fd\3\2\2\2\63\u0104\3\2\2\2\65\u010f\3\2\2"+
		"\2\67\u0116\3\2\2\29\u011a\3\2\2\2;\u0126\3\2\2\2=\u012f\3\2\2\2?\u0136"+
		"\3\2\2\2A\u0139\3\2\2\2C\u013c\3\2\2\2E\u013e\3\2\2\2G\u0141\3\2\2\2I"+
		"\u0144\3\2\2\2K\u0146\3\2\2\2M\u0149\3\2\2\2O\u014c\3\2\2\2Q\u014e\3\2"+
		"\2\2S\u0150\3\2\2\2U\u0152\3\2\2\2W\u0154\3\2\2\2Y\u0156\3\2\2\2[\u0158"+
		"\3\2\2\2]\u015a\3\2\2\2_\u015c\3\2\2\2a\u015e\3\2\2\2c\u0160\3\2\2\2e"+
		"\u0162\3\2\2\2g\u0164\3\2\2\2i\u0166\3\2\2\2k\u0168\3\2\2\2m\u016a\3\2"+
		"\2\2o\u016c\3\2\2\2q\u016e\3\2\2\2s\u0171\3\2\2\2u\u0177\3\2\2\2w\u0185"+
		"\3\2\2\2y\u0190\3\2\2\2{|\7k\2\2|}\7p\2\2}~\7v\2\2~\4\3\2\2\2\177\u0080"+
		"\7h\2\2\u0080\u0081\7n\2\2\u0081\u0082\7q\2\2\u0082\u0083\7c\2\2\u0083"+
		"\u0084\7v\2\2\u0084\6\3\2\2\2\u0085\u0086\7d\2\2\u0086\u0087\7q\2\2\u0087"+
		"\u0088\7q\2\2\u0088\u0089\7n\2\2\u0089\b\3\2\2\2\u008a\u008b\7e\2\2\u008b"+
		"\u008c\7j\2\2\u008c\u008d\7c\2\2\u008d\u008e\7t\2\2\u008e\u008f\7c\2\2"+
		"\u008f\u0090\7e\2\2\u0090\u0091\7v\2\2\u0091\u0092\7g\2\2\u0092\u0093"+
		"\7t\2\2\u0093\n\3\2\2\2\u0094\u0095\7U\2\2\u0095\u0096\7v\2\2\u0096\u0097"+
		"\7t\2\2\u0097\u0098\7k\2\2\u0098\u0099\7p\2\2\u0099\u009a\7i\2\2\u009a"+
		"\f\3\2\2\2\u009b\u009c\7p\2\2\u009c\u009d\7k\2\2\u009d\u009e\7n\2\2\u009e"+
		"\16\3\2\2\2\u009f\u00a0\7v\2\2\u00a0\u00a1\7t\2\2\u00a1\u00a2\7w\2\2\u00a2"+
		"\u00a3\7g\2\2\u00a3\20\3\2\2\2\u00a4\u00a5\7h\2\2\u00a5\u00a6\7c\2\2\u00a6"+
		"\u00a7\7n\2\2\u00a7\u00a8\7u\2\2\u00a8\u00a9\7g\2\2\u00a9\22\3\2\2\2\u00aa"+
		"\u00ab\7r\2\2\u00ab\u00ac\7t\2\2\u00ac\u00ad\7k\2\2\u00ad\u00ae\7p\2\2"+
		"\u00ae\u00af\7v\2\2\u00af\24\3\2\2\2\u00b0\u00b1\7k\2\2\u00b1\u00b2\7"+
		"h\2\2\u00b2\26\3\2\2\2\u00b3\u00b4\7g\2\2\u00b4\u00b5\7n\2\2\u00b5\u00b6"+
		"\7u\2\2\u00b6\u00b7\7g\2\2\u00b7\30\3\2\2\2\u00b8\u00b9\7y\2\2\u00b9\u00ba"+
		"\7j\2\2\u00ba\u00bb\7k\2\2\u00bb\u00bc\7n\2\2\u00bc\u00bd\7g\2\2\u00bd"+
		"\32\3\2\2\2\u00be\u00bf\7h\2\2\u00bf\u00c0\7q\2\2\u00c0\u00c1\7t\2\2\u00c1"+
		"\34\3\2\2\2\u00c2\u00c3\7k\2\2\u00c3\u00c4\7p\2\2\u00c4\36\3\2\2\2\u00c5"+
		"\u00c6\7u\2\2\u00c6\u00c7\7y\2\2\u00c7\u00c8\7k\2\2\u00c8\u00c9\7v\2\2"+
		"\u00c9\u00ca\7e\2\2\u00ca\u00cb\7j\2\2\u00cb \3\2\2\2\u00cc\u00cd\7e\2"+
		"\2\u00cd\u00ce\7c\2\2\u00ce\u00cf\7u\2\2\u00cf\u00d0\7g\2\2\u00d0\"\3"+
		"\2\2\2\u00d1\u00d2\7f\2\2\u00d2\u00d3\7g\2\2\u00d3\u00d4\7h\2\2\u00d4"+
		"\u00d5\7c\2\2\u00d5\u00d6\7w\2\2\u00d6\u00d7\7n\2\2\u00d7\u00d8\7v\2\2"+
		"\u00d8$\3\2\2\2\u00d9\u00da\7x\2\2\u00da\u00db\7c\2\2\u00db\u00dc\7t\2"+
		"\2\u00dc&\3\2\2\2\u00dd\u00de\7d\2\2\u00de\u00df\7t\2\2\u00df\u00e0\7"+
		"g\2\2\u00e0\u00e1\7c\2\2\u00e1\u00e2\7m\2\2\u00e2(\3\2\2\2\u00e3\u00e4"+
		"\7t\2\2\u00e4\u00e5\7g\2\2\u00e5\u00e6\7v\2\2\u00e6\u00e7\7w\2\2\u00e7"+
		"\u00e8\7t\2\2\u00e8\u00e9\7p\2\2\u00e9*\3\2\2\2\u00ea\u00eb\7h\2\2\u00eb"+
		"\u00ec\7w\2\2\u00ec\u00ed\7p\2\2\u00ed\u00ee\7e\2\2\u00ee,\3\2\2\2\u00ef"+
		"\u00f0\7e\2\2\u00f0\u00f1\7q\2\2\u00f1\u00f2\7w\2\2\u00f2\u00f3\7p\2\2"+
		"\u00f3\u00f4\7v\2\2\u00f4.\3\2\2\2\u00f5\u00f6\7K\2\2\u00f6\u00f7\7u\2"+
		"\2\u00f7\u00f8\7G\2\2\u00f8\u00f9\7o\2\2\u00f9\u00fa\7r\2\2\u00fa\u00fb"+
		"\7v\2\2\u00fb\u00fc\7{\2\2\u00fc\60\3\2\2\2\u00fd\u00fe\7c\2\2\u00fe\u00ff"+
		"\7r\2\2\u00ff\u0100\7r\2\2\u0100\u0101\7g\2\2\u0101\u0102\7p\2\2\u0102"+
		"\u0103\7f\2\2\u0103\62\3\2\2\2\u0104\u0105\7t\2\2\u0105\u0106\7g\2\2\u0106"+
		"\u0107\7o\2\2\u0107\u0108\7q\2\2\u0108\u0109\7x\2\2\u0109\u010a\7g\2\2"+
		"\u010a\u010b\7N\2\2\u010b\u010c\7c\2\2\u010c\u010d\7u\2\2\u010d\u010e"+
		"\7v\2\2\u010e\64\3\2\2\2\u010f\u0110\7t\2\2\u0110\u0111\7g\2\2\u0111\u0112"+
		"\7o\2\2\u0112\u0113\7q\2\2\u0113\u0114\7x\2\2\u0114\u0115\7g\2\2\u0115"+
		"\66\3\2\2\2\u0116\u0117\7c\2\2\u0117\u0118\7v\2\2\u01188\3\2\2\2\u0119"+
		"\u011b\t\2\2\2\u011a\u0119\3\2\2\2\u011b\u011c\3\2\2\2\u011c\u011a\3\2"+
		"\2\2\u011c\u011d\3\2\2\2\u011d\u0124\3\2\2\2\u011e\u0120\7\60\2\2\u011f"+
		"\u0121\t\2\2\2\u0120\u011f\3\2\2\2\u0121\u0122\3\2\2\2\u0122\u0120\3\2"+
		"\2\2\u0122\u0123\3\2\2\2\u0123\u0125\3\2\2\2\u0124\u011e\3\2\2\2\u0124"+
		"\u0125\3\2\2\2\u0125:\3\2\2\2\u0126\u012a\7$\2\2\u0127\u0129\n\3\2\2\u0128"+
		"\u0127\3\2\2\2\u0129\u012c\3\2\2\2\u012a\u0128\3\2\2\2\u012a\u012b\3\2"+
		"\2\2\u012b\u012d\3\2\2\2\u012c\u012a\3\2\2\2\u012d\u012e\7$\2\2\u012e"+
		"<\3\2\2\2\u012f\u0133\t\4\2\2\u0130\u0132\t\5\2\2\u0131\u0130\3\2\2\2"+
		"\u0132\u0135\3\2\2\2\u0133\u0131\3\2\2\2\u0133\u0134\3\2\2\2\u0134>\3"+
		"\2\2\2\u0135\u0133\3\2\2\2\u0136\u0137\7#\2\2\u0137\u0138\7?\2\2\u0138"+
		"@\3\2\2\2\u0139\u013a\7?\2\2\u013a\u013b\7?\2\2\u013bB\3\2\2\2\u013c\u013d"+
		"\7#\2\2\u013dD\3\2\2\2\u013e\u013f\7~\2\2\u013f\u0140\7~\2\2\u0140F\3"+
		"\2\2\2\u0141\u0142\7(\2\2\u0142\u0143\7(\2\2\u0143H\3\2\2\2\u0144\u0145"+
		"\7?\2\2\u0145J\3\2\2\2\u0146\u0147\7@\2\2\u0147\u0148\7?\2\2\u0148L\3"+
		"\2\2\2\u0149\u014a\7>\2\2\u014a\u014b\7?\2\2\u014bN\3\2\2\2\u014c\u014d"+
		"\7@\2\2\u014dP\3\2\2\2\u014e\u014f\7>\2\2\u014fR\3\2\2\2\u0150\u0151\7"+
		",\2\2\u0151T\3\2\2\2\u0152\u0153\7\61\2\2\u0153V\3\2\2\2\u0154\u0155\7"+
		"-\2\2\u0155X\3\2\2\2\u0156\u0157\7/\2\2\u0157Z\3\2\2\2\u0158\u0159\7\'"+
		"\2\2\u0159\\\3\2\2\2\u015a\u015b\7*\2\2\u015b^\3\2\2\2\u015c\u015d\7+"+
		"\2\2\u015d`\3\2\2\2\u015e\u015f\7}\2\2\u015fb\3\2\2\2\u0160\u0161\7\177"+
		"\2\2\u0161d\3\2\2\2\u0162\u0163\7]\2\2\u0163f\3\2\2\2\u0164\u0165\7_\2"+
		"\2\u0165h\3\2\2\2\u0166\u0167\7<\2\2\u0167j\3\2\2\2\u0168\u0169\7.\2\2"+
		"\u0169l\3\2\2\2\u016a\u016b\7=\2\2\u016bn\3\2\2\2\u016c\u016d\7A\2\2\u016d"+
		"p\3\2\2\2\u016e\u016f\7\60\2\2\u016fr\3\2\2\2\u0170\u0172\t\6\2\2\u0171"+
		"\u0170\3\2\2\2\u0172\u0173\3\2\2\2\u0173\u0171\3\2\2\2\u0173\u0174\3\2"+
		"\2\2\u0174\u0175\3\2\2\2\u0175\u0176\b:\2\2\u0176t\3\2\2\2\u0177\u0178"+
		"\7\61\2\2\u0178\u0179\7,\2\2\u0179\u017d\3\2\2\2\u017a\u017c\13\2\2\2"+
		"\u017b\u017a\3\2\2\2\u017c\u017f\3\2\2\2\u017d\u017e\3\2\2\2\u017d\u017b"+
		"\3\2\2\2\u017e\u0180\3\2\2\2\u017f\u017d\3\2\2\2\u0180\u0181\7,\2\2\u0181"+
		"\u0182\7\61\2\2\u0182\u0183\3\2\2\2\u0183\u0184\b;\2\2\u0184v\3\2\2\2"+
		"\u0185\u0186\7\61\2\2\u0186\u0187\7\61\2\2\u0187\u018b\3\2\2\2\u0188\u018a"+
		"\n\7\2\2\u0189\u0188\3\2\2\2\u018a\u018d\3\2\2\2\u018b\u0189\3\2\2\2\u018b"+
		"\u018c\3\2\2\2\u018c\u018e\3\2\2\2\u018d\u018b\3\2\2\2\u018e\u018f\b<"+
		"\2\2\u018fx\3\2\2\2\u0190\u0191\7^\2\2\u0191\u0192\t\b\2\2\u0192z\3\2"+
		"\2\2\13\2\u011c\u0122\u0124\u012a\u0133\u0173\u017d\u018b\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
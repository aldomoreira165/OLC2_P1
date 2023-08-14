// Generated from c:\Users\aldov\Desktop\Cursos Segundo Semestre 2023\Laboratorio Organización de Lenguajes y Compiladores 2 - copia\Proyectos\Proyecto 1\interprete\SwiftLexer.g4 by ANTLR 4.9.2
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
		IF=10, ELSE=11, WHILE=12, VAR=13, NUMBER=14, STRING=15, ID=16, DIF=17, 
		IG_IG=18, NOT=19, OR=20, AND=21, IG=22, MAY_IG=23, MEN_IG=24, MAYOR=25, 
		MENOR=26, MUL=27, DIV=28, ADD=29, SUB=30, MOD=31, PARIZQ=32, PARDER=33, 
		LLAVEIZQ=34, LLAVEDER=35, DOSPUNTOS=36, INTERROGACION=37, WHITESPACE=38, 
		COMMENT=39, LINE_COMMENT=40;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "NIL", "TRU", "FAL", 
			"PRINT", "IF", "ELSE", "WHILE", "VAR", "NUMBER", "STRING", "ID", "DIF", 
			"IG_IG", "NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", 
			"MUL", "DIV", "ADD", "SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", 
			"DOSPUNTOS", "INTERROGACION", "WHITESPACE", "COMMENT", "LINE_COMMENT", 
			"ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'int'", "'float'", "'bool'", "'character'", "'String'", "'nil'", 
			"'true'", "'false'", "'print'", "'if'", "'else'", "'while'", "'var'", 
			null, null, null, "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", "'>='", 
			"'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'%'", "'('", "')'", 
			"'{'", "'}'", "':'", "'?'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "NIL", "TRU", "FAL", 
			"PRINT", "IF", "ELSE", "WHILE", "VAR", "NUMBER", "STRING", "ID", "DIF", 
			"IG_IG", "NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", 
			"MUL", "DIV", "ADD", "SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", 
			"DOSPUNTOS", "INTERROGACION", "WHITESPACE", "COMMENT", "LINE_COMMENT"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2*\u010c\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\3\2\3\2"+
		"\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\b"+
		"\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\13"+
		"\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16"+
		"\3\16\3\17\6\17\u009e\n\17\r\17\16\17\u009f\3\17\3\17\6\17\u00a4\n\17"+
		"\r\17\16\17\u00a5\5\17\u00a8\n\17\3\20\3\20\7\20\u00ac\n\20\f\20\16\20"+
		"\u00af\13\20\3\20\3\20\3\21\3\21\7\21\u00b5\n\21\f\21\16\21\u00b8\13\21"+
		"\3\22\3\22\3\22\3\23\3\23\3\23\3\24\3\24\3\25\3\25\3\25\3\26\3\26\3\26"+
		"\3\27\3\27\3\30\3\30\3\30\3\31\3\31\3\31\3\32\3\32\3\33\3\33\3\34\3\34"+
		"\3\35\3\35\3\36\3\36\3\37\3\37\3 \3 \3!\3!\3\"\3\"\3#\3#\3$\3$\3%\3%\3"+
		"&\3&\3\'\6\'\u00eb\n\'\r\'\16\'\u00ec\3\'\3\'\3(\3(\3(\3(\7(\u00f5\n("+
		"\f(\16(\u00f8\13(\3(\3(\3(\3(\3(\3)\3)\3)\3)\7)\u0103\n)\f)\16)\u0106"+
		"\13)\3)\3)\3*\3*\3*\3\u00f6\2+\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13"+
		"\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61"+
		"\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S\2\3\2\t\3\2\62"+
		";\3\2$$\5\2C\\aac|\6\2\62;C\\aac|\6\2\13\f\17\17\"\"^^\4\2\f\f\17\17\t"+
		"\2\"#%%--/\60<<BB]_\2\u0112\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3"+
		"\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2"+
		"\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37"+
		"\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3"+
		"\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2"+
		"\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C"+
		"\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2"+
		"\2\2\2Q\3\2\2\2\3U\3\2\2\2\5Y\3\2\2\2\7_\3\2\2\2\td\3\2\2\2\13n\3\2\2"+
		"\2\ru\3\2\2\2\17y\3\2\2\2\21~\3\2\2\2\23\u0084\3\2\2\2\25\u008a\3\2\2"+
		"\2\27\u008d\3\2\2\2\31\u0092\3\2\2\2\33\u0098\3\2\2\2\35\u009d\3\2\2\2"+
		"\37\u00a9\3\2\2\2!\u00b2\3\2\2\2#\u00b9\3\2\2\2%\u00bc\3\2\2\2\'\u00bf"+
		"\3\2\2\2)\u00c1\3\2\2\2+\u00c4\3\2\2\2-\u00c7\3\2\2\2/\u00c9\3\2\2\2\61"+
		"\u00cc\3\2\2\2\63\u00cf\3\2\2\2\65\u00d1\3\2\2\2\67\u00d3\3\2\2\29\u00d5"+
		"\3\2\2\2;\u00d7\3\2\2\2=\u00d9\3\2\2\2?\u00db\3\2\2\2A\u00dd\3\2\2\2C"+
		"\u00df\3\2\2\2E\u00e1\3\2\2\2G\u00e3\3\2\2\2I\u00e5\3\2\2\2K\u00e7\3\2"+
		"\2\2M\u00ea\3\2\2\2O\u00f0\3\2\2\2Q\u00fe\3\2\2\2S\u0109\3\2\2\2UV\7k"+
		"\2\2VW\7p\2\2WX\7v\2\2X\4\3\2\2\2YZ\7h\2\2Z[\7n\2\2[\\\7q\2\2\\]\7c\2"+
		"\2]^\7v\2\2^\6\3\2\2\2_`\7d\2\2`a\7q\2\2ab\7q\2\2bc\7n\2\2c\b\3\2\2\2"+
		"de\7e\2\2ef\7j\2\2fg\7c\2\2gh\7t\2\2hi\7c\2\2ij\7e\2\2jk\7v\2\2kl\7g\2"+
		"\2lm\7t\2\2m\n\3\2\2\2no\7U\2\2op\7v\2\2pq\7t\2\2qr\7k\2\2rs\7p\2\2st"+
		"\7i\2\2t\f\3\2\2\2uv\7p\2\2vw\7k\2\2wx\7n\2\2x\16\3\2\2\2yz\7v\2\2z{\7"+
		"t\2\2{|\7w\2\2|}\7g\2\2}\20\3\2\2\2~\177\7h\2\2\177\u0080\7c\2\2\u0080"+
		"\u0081\7n\2\2\u0081\u0082\7u\2\2\u0082\u0083\7g\2\2\u0083\22\3\2\2\2\u0084"+
		"\u0085\7r\2\2\u0085\u0086\7t\2\2\u0086\u0087\7k\2\2\u0087\u0088\7p\2\2"+
		"\u0088\u0089\7v\2\2\u0089\24\3\2\2\2\u008a\u008b\7k\2\2\u008b\u008c\7"+
		"h\2\2\u008c\26\3\2\2\2\u008d\u008e\7g\2\2\u008e\u008f\7n\2\2\u008f\u0090"+
		"\7u\2\2\u0090\u0091\7g\2\2\u0091\30\3\2\2\2\u0092\u0093\7y\2\2\u0093\u0094"+
		"\7j\2\2\u0094\u0095\7k\2\2\u0095\u0096\7n\2\2\u0096\u0097\7g\2\2\u0097"+
		"\32\3\2\2\2\u0098\u0099\7x\2\2\u0099\u009a\7c\2\2\u009a\u009b\7t\2\2\u009b"+
		"\34\3\2\2\2\u009c\u009e\t\2\2\2\u009d\u009c\3\2\2\2\u009e\u009f\3\2\2"+
		"\2\u009f\u009d\3\2\2\2\u009f\u00a0\3\2\2\2\u00a0\u00a7\3\2\2\2\u00a1\u00a3"+
		"\7\60\2\2\u00a2\u00a4\t\2\2\2\u00a3\u00a2\3\2\2\2\u00a4\u00a5\3\2\2\2"+
		"\u00a5\u00a3\3\2\2\2\u00a5\u00a6\3\2\2\2\u00a6\u00a8\3\2\2\2\u00a7\u00a1"+
		"\3\2\2\2\u00a7\u00a8\3\2\2\2\u00a8\36\3\2\2\2\u00a9\u00ad\7$\2\2\u00aa"+
		"\u00ac\n\3\2\2\u00ab\u00aa\3\2\2\2\u00ac\u00af\3\2\2\2\u00ad\u00ab\3\2"+
		"\2\2\u00ad\u00ae\3\2\2\2\u00ae\u00b0\3\2\2\2\u00af\u00ad\3\2\2\2\u00b0"+
		"\u00b1\7$\2\2\u00b1 \3\2\2\2\u00b2\u00b6\t\4\2\2\u00b3\u00b5\t\5\2\2\u00b4"+
		"\u00b3\3\2\2\2\u00b5\u00b8\3\2\2\2\u00b6\u00b4\3\2\2\2\u00b6\u00b7\3\2"+
		"\2\2\u00b7\"\3\2\2\2\u00b8\u00b6\3\2\2\2\u00b9\u00ba\7#\2\2\u00ba\u00bb"+
		"\7?\2\2\u00bb$\3\2\2\2\u00bc\u00bd\7?\2\2\u00bd\u00be\7?\2\2\u00be&\3"+
		"\2\2\2\u00bf\u00c0\7#\2\2\u00c0(\3\2\2\2\u00c1\u00c2\7~\2\2\u00c2\u00c3"+
		"\7~\2\2\u00c3*\3\2\2\2\u00c4\u00c5\7(\2\2\u00c5\u00c6\7(\2\2\u00c6,\3"+
		"\2\2\2\u00c7\u00c8\7?\2\2\u00c8.\3\2\2\2\u00c9\u00ca\7@\2\2\u00ca\u00cb"+
		"\7?\2\2\u00cb\60\3\2\2\2\u00cc\u00cd\7>\2\2\u00cd\u00ce\7?\2\2\u00ce\62"+
		"\3\2\2\2\u00cf\u00d0\7@\2\2\u00d0\64\3\2\2\2\u00d1\u00d2\7>\2\2\u00d2"+
		"\66\3\2\2\2\u00d3\u00d4\7,\2\2\u00d48\3\2\2\2\u00d5\u00d6\7\61\2\2\u00d6"+
		":\3\2\2\2\u00d7\u00d8\7-\2\2\u00d8<\3\2\2\2\u00d9\u00da\7/\2\2\u00da>"+
		"\3\2\2\2\u00db\u00dc\7\'\2\2\u00dc@\3\2\2\2\u00dd\u00de\7*\2\2\u00deB"+
		"\3\2\2\2\u00df\u00e0\7+\2\2\u00e0D\3\2\2\2\u00e1\u00e2\7}\2\2\u00e2F\3"+
		"\2\2\2\u00e3\u00e4\7\177\2\2\u00e4H\3\2\2\2\u00e5\u00e6\7<\2\2\u00e6J"+
		"\3\2\2\2\u00e7\u00e8\7A\2\2\u00e8L\3\2\2\2\u00e9\u00eb\t\6\2\2\u00ea\u00e9"+
		"\3\2\2\2\u00eb\u00ec\3\2\2\2\u00ec\u00ea\3\2\2\2\u00ec\u00ed\3\2\2\2\u00ed"+
		"\u00ee\3\2\2\2\u00ee\u00ef\b\'\2\2\u00efN\3\2\2\2\u00f0\u00f1\7\61\2\2"+
		"\u00f1\u00f2\7,\2\2\u00f2\u00f6\3\2\2\2\u00f3\u00f5\13\2\2\2\u00f4\u00f3"+
		"\3\2\2\2\u00f5\u00f8\3\2\2\2\u00f6\u00f7\3\2\2\2\u00f6\u00f4\3\2\2\2\u00f7"+
		"\u00f9\3\2\2\2\u00f8\u00f6\3\2\2\2\u00f9\u00fa\7,\2\2\u00fa\u00fb\7\61"+
		"\2\2\u00fb\u00fc\3\2\2\2\u00fc\u00fd\b(\2\2\u00fdP\3\2\2\2\u00fe\u00ff"+
		"\7\61\2\2\u00ff\u0100\7\61\2\2\u0100\u0104\3\2\2\2\u0101\u0103\n\7\2\2"+
		"\u0102\u0101\3\2\2\2\u0103\u0106\3\2\2\2\u0104\u0102\3\2\2\2\u0104\u0105"+
		"\3\2\2\2\u0105\u0107\3\2\2\2\u0106\u0104\3\2\2\2\u0107\u0108\b)\2\2\u0108"+
		"R\3\2\2\2\u0109\u010a\7^\2\2\u010a\u010b\t\b\2\2\u010bT\3\2\2\2\13\2\u009f"+
		"\u00a5\u00a7\u00ad\u00b6\u00ec\u00f6\u0104\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
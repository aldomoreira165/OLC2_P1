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
		INT=1, FLOAT=2, BOOL=3, CHARACTER=4, PSTRING=5, TRU=6, FAL=7, PRINT=8, 
		IF=9, ELSE=10, WHILE=11, VAR=12, NUMBER=13, STRING=14, ID=15, DIF=16, 
		IG_IG=17, NOT=18, OR=19, AND=20, IG=21, MAY_IG=22, MEN_IG=23, MAYOR=24, 
		MENOR=25, MUL=26, DIV=27, ADD=28, SUB=29, PARIZQ=30, PARDER=31, LLAVEIZQ=32, 
		LLAVEDER=33, DOSPUNTOS=34, INTERROGACION=35, WHITESPACE=36, COMMENT=37, 
		LINE_COMMENT=38;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "TRU", "FAL", "PRINT", 
			"IF", "ELSE", "WHILE", "VAR", "NUMBER", "STRING", "ID", "DIF", "IG_IG", 
			"NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", 
			"DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "DOSPUNTOS", 
			"INTERROGACION", "WHITESPACE", "COMMENT", "LINE_COMMENT", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'Bool'", "'Character'", "'String'", "'true'", 
			"'false'", "'print'", "'if'", "'else'", "'while'", "'var'", null, null, 
			null, "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", "'>='", "'<='", "'>'", 
			"'<'", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'{'", "'}'", "':'", 
			"'?'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "TRU", "FAL", "PRINT", 
			"IF", "ELSE", "WHILE", "VAR", "NUMBER", "STRING", "ID", "DIF", "IG_IG", 
			"NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", 
			"DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "DOSPUNTOS", 
			"INTERROGACION", "WHITESPACE", "COMMENT", "LINE_COMMENT"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2(\u0102\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\3\2\3\2\3\2\3\2\3\3"+
		"\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3"+
		"\13\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\16\6\16\u0096\n\16\r\16"+
		"\16\16\u0097\3\16\3\16\6\16\u009c\n\16\r\16\16\16\u009d\5\16\u00a0\n\16"+
		"\3\17\3\17\7\17\u00a4\n\17\f\17\16\17\u00a7\13\17\3\17\3\17\3\20\3\20"+
		"\7\20\u00ad\n\20\f\20\16\20\u00b0\13\20\3\21\3\21\3\21\3\22\3\22\3\22"+
		"\3\23\3\23\3\24\3\24\3\24\3\25\3\25\3\25\3\26\3\26\3\27\3\27\3\27\3\30"+
		"\3\30\3\30\3\31\3\31\3\32\3\32\3\33\3\33\3\34\3\34\3\35\3\35\3\36\3\36"+
		"\3\37\3\37\3 \3 \3!\3!\3\"\3\"\3#\3#\3$\3$\3%\6%\u00e1\n%\r%\16%\u00e2"+
		"\3%\3%\3&\3&\3&\3&\7&\u00eb\n&\f&\16&\u00ee\13&\3&\3&\3&\3&\3&\3\'\3\'"+
		"\3\'\3\'\7\'\u00f9\n\'\f\'\16\'\u00fc\13\'\3\'\3\'\3(\3(\3(\3\u00ec\2"+
		")\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20"+
		"\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37"+
		"= ?!A\"C#E$G%I&K\'M(O\2\3\2\t\3\2\62;\3\2$$\5\2C\\aac|\6\2\62;C\\aac|"+
		"\6\2\13\f\17\17\"\"^^\4\2\f\f\17\17\t\2\"#%%--/\60<<BB]_\2\u0108\2\3\3"+
		"\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2"+
		"\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3"+
		"\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2"+
		"%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61"+
		"\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2"+
		"\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I"+
		"\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\3Q\3\2\2\2\5U\3\2\2\2\7[\3\2\2\2\t`\3\2"+
		"\2\2\13j\3\2\2\2\rq\3\2\2\2\17v\3\2\2\2\21|\3\2\2\2\23\u0082\3\2\2\2\25"+
		"\u0085\3\2\2\2\27\u008a\3\2\2\2\31\u0090\3\2\2\2\33\u0095\3\2\2\2\35\u00a1"+
		"\3\2\2\2\37\u00aa\3\2\2\2!\u00b1\3\2\2\2#\u00b4\3\2\2\2%\u00b7\3\2\2\2"+
		"\'\u00b9\3\2\2\2)\u00bc\3\2\2\2+\u00bf\3\2\2\2-\u00c1\3\2\2\2/\u00c4\3"+
		"\2\2\2\61\u00c7\3\2\2\2\63\u00c9\3\2\2\2\65\u00cb\3\2\2\2\67\u00cd\3\2"+
		"\2\29\u00cf\3\2\2\2;\u00d1\3\2\2\2=\u00d3\3\2\2\2?\u00d5\3\2\2\2A\u00d7"+
		"\3\2\2\2C\u00d9\3\2\2\2E\u00db\3\2\2\2G\u00dd\3\2\2\2I\u00e0\3\2\2\2K"+
		"\u00e6\3\2\2\2M\u00f4\3\2\2\2O\u00ff\3\2\2\2QR\7K\2\2RS\7p\2\2ST\7v\2"+
		"\2T\4\3\2\2\2UV\7H\2\2VW\7n\2\2WX\7q\2\2XY\7c\2\2YZ\7v\2\2Z\6\3\2\2\2"+
		"[\\\7D\2\2\\]\7q\2\2]^\7q\2\2^_\7n\2\2_\b\3\2\2\2`a\7E\2\2ab\7j\2\2bc"+
		"\7c\2\2cd\7t\2\2de\7c\2\2ef\7e\2\2fg\7v\2\2gh\7g\2\2hi\7t\2\2i\n\3\2\2"+
		"\2jk\7U\2\2kl\7v\2\2lm\7t\2\2mn\7k\2\2no\7p\2\2op\7i\2\2p\f\3\2\2\2qr"+
		"\7v\2\2rs\7t\2\2st\7w\2\2tu\7g\2\2u\16\3\2\2\2vw\7h\2\2wx\7c\2\2xy\7n"+
		"\2\2yz\7u\2\2z{\7g\2\2{\20\3\2\2\2|}\7r\2\2}~\7t\2\2~\177\7k\2\2\177\u0080"+
		"\7p\2\2\u0080\u0081\7v\2\2\u0081\22\3\2\2\2\u0082\u0083\7k\2\2\u0083\u0084"+
		"\7h\2\2\u0084\24\3\2\2\2\u0085\u0086\7g\2\2\u0086\u0087\7n\2\2\u0087\u0088"+
		"\7u\2\2\u0088\u0089\7g\2\2\u0089\26\3\2\2\2\u008a\u008b\7y\2\2\u008b\u008c"+
		"\7j\2\2\u008c\u008d\7k\2\2\u008d\u008e\7n\2\2\u008e\u008f\7g\2\2\u008f"+
		"\30\3\2\2\2\u0090\u0091\7x\2\2\u0091\u0092\7c\2\2\u0092\u0093\7t\2\2\u0093"+
		"\32\3\2\2\2\u0094\u0096\t\2\2\2\u0095\u0094\3\2\2\2\u0096\u0097\3\2\2"+
		"\2\u0097\u0095\3\2\2\2\u0097\u0098\3\2\2\2\u0098\u009f\3\2\2\2\u0099\u009b"+
		"\7\60\2\2\u009a\u009c\t\2\2\2\u009b\u009a\3\2\2\2\u009c\u009d\3\2\2\2"+
		"\u009d\u009b\3\2\2\2\u009d\u009e\3\2\2\2\u009e\u00a0\3\2\2\2\u009f\u0099"+
		"\3\2\2\2\u009f\u00a0\3\2\2\2\u00a0\34\3\2\2\2\u00a1\u00a5\7$\2\2\u00a2"+
		"\u00a4\n\3\2\2\u00a3\u00a2\3\2\2\2\u00a4\u00a7\3\2\2\2\u00a5\u00a3\3\2"+
		"\2\2\u00a5\u00a6\3\2\2\2\u00a6\u00a8\3\2\2\2\u00a7\u00a5\3\2\2\2\u00a8"+
		"\u00a9\7$\2\2\u00a9\36\3\2\2\2\u00aa\u00ae\t\4\2\2\u00ab\u00ad\t\5\2\2"+
		"\u00ac\u00ab\3\2\2\2\u00ad\u00b0\3\2\2\2\u00ae\u00ac\3\2\2\2\u00ae\u00af"+
		"\3\2\2\2\u00af \3\2\2\2\u00b0\u00ae\3\2\2\2\u00b1\u00b2\7#\2\2\u00b2\u00b3"+
		"\7?\2\2\u00b3\"\3\2\2\2\u00b4\u00b5\7?\2\2\u00b5\u00b6\7?\2\2\u00b6$\3"+
		"\2\2\2\u00b7\u00b8\7#\2\2\u00b8&\3\2\2\2\u00b9\u00ba\7~\2\2\u00ba\u00bb"+
		"\7~\2\2\u00bb(\3\2\2\2\u00bc\u00bd\7(\2\2\u00bd\u00be\7(\2\2\u00be*\3"+
		"\2\2\2\u00bf\u00c0\7?\2\2\u00c0,\3\2\2\2\u00c1\u00c2\7@\2\2\u00c2\u00c3"+
		"\7?\2\2\u00c3.\3\2\2\2\u00c4\u00c5\7>\2\2\u00c5\u00c6\7?\2\2\u00c6\60"+
		"\3\2\2\2\u00c7\u00c8\7@\2\2\u00c8\62\3\2\2\2\u00c9\u00ca\7>\2\2\u00ca"+
		"\64\3\2\2\2\u00cb\u00cc\7,\2\2\u00cc\66\3\2\2\2\u00cd\u00ce\7\61\2\2\u00ce"+
		"8\3\2\2\2\u00cf\u00d0\7-\2\2\u00d0:\3\2\2\2\u00d1\u00d2\7/\2\2\u00d2<"+
		"\3\2\2\2\u00d3\u00d4\7*\2\2\u00d4>\3\2\2\2\u00d5\u00d6\7+\2\2\u00d6@\3"+
		"\2\2\2\u00d7\u00d8\7}\2\2\u00d8B\3\2\2\2\u00d9\u00da\7\177\2\2\u00daD"+
		"\3\2\2\2\u00db\u00dc\7<\2\2\u00dcF\3\2\2\2\u00dd\u00de\7A\2\2\u00deH\3"+
		"\2\2\2\u00df\u00e1\t\6\2\2\u00e0\u00df\3\2\2\2\u00e1\u00e2\3\2\2\2\u00e2"+
		"\u00e0\3\2\2\2\u00e2\u00e3\3\2\2\2\u00e3\u00e4\3\2\2\2\u00e4\u00e5\b%"+
		"\2\2\u00e5J\3\2\2\2\u00e6\u00e7\7\61\2\2\u00e7\u00e8\7,\2\2\u00e8\u00ec"+
		"\3\2\2\2\u00e9\u00eb\13\2\2\2\u00ea\u00e9\3\2\2\2\u00eb\u00ee\3\2\2\2"+
		"\u00ec\u00ed\3\2\2\2\u00ec\u00ea\3\2\2\2\u00ed\u00ef\3\2\2\2\u00ee\u00ec"+
		"\3\2\2\2\u00ef\u00f0\7,\2\2\u00f0\u00f1\7\61\2\2\u00f1\u00f2\3\2\2\2\u00f2"+
		"\u00f3\b&\2\2\u00f3L\3\2\2\2\u00f4\u00f5\7\61\2\2\u00f5\u00f6\7\61\2\2"+
		"\u00f6\u00fa\3\2\2\2\u00f7\u00f9\n\7\2\2\u00f8\u00f7\3\2\2\2\u00f9\u00fc"+
		"\3\2\2\2\u00fa\u00f8\3\2\2\2\u00fa\u00fb\3\2\2\2\u00fb\u00fd\3\2\2\2\u00fc"+
		"\u00fa\3\2\2\2\u00fd\u00fe\b\'\2\2\u00feN\3\2\2\2\u00ff\u0100\7^\2\2\u0100"+
		"\u0101\t\b\2\2\u0101P\3\2\2\2\13\2\u0097\u009d\u009f\u00a5\u00ae\u00e2"+
		"\u00ec\u00fa\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
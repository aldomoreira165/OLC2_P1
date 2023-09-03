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
		IF=10, ELSE=11, WHILE=12, FOR=13, GUARD=14, IN=15, SWITCH=16, CASE=17, 
		DEFAULT=18, VAR=19, LET=20, BREAK=21, RETURN=22, CONTINUE=23, FUNC=24, 
		COUNT=25, ISEMPTY=26, APPEND=27, REMOVE_LAST=28, REMOVE=29, STRUCT=30, 
		STRUCT_VAR=31, STRUCT_LET=32, AT=33, ST=34, NUMBER=35, STRING=36, ID=37, 
		DIF=38, IG_IG=39, NOT=40, OR=41, AND=42, IG=43, MAY_IG=44, MEN_IG=45, 
		MAYOR=46, MENOR=47, MUL=48, DIV=49, ADD=50, SUB=51, MOD=52, PARIZQ=53, 
		PARDER=54, LLAVEIZQ=55, LLAVEDER=56, CORCHIZQ=57, CORCHDER=58, DOSPUNTOS=59, 
		COMA=60, PTCOMA=61, INTERROGACION=62, PUNTO=63, GUIONBAJO=64, WHITESPACE=65, 
		COMMENT=66, LINE_COMMENT=67;
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
			"STRUCT_LET", "AT", "ST", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", 
			"OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", 
			"ADD", "SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "CORCHIZQ", 
			"CORCHDER", "DOSPUNTOS", "COMA", "PTCOMA", "INTERROGACION", "PUNTO", 
			"GUIONBAJO", "WHITESPACE", "COMMENT", "LINE_COMMENT", "ESC_SEQ"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2E\u01ce\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\3\2\3\2\3\2\3\2\3\3\3"+
		"\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5"+
		"\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3"+
		"\b\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\f"+
		"\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\24"+
		"\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\26\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30"+
		"\3\30\3\31\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\35"+
		"\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36"+
		"\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3 \3 \3"+
		"!\3!\3!\3!\3!\3!\3\"\3\"\3\"\3#\3#\3#\3$\6$\u0154\n$\r$\16$\u0155\3$\3"+
		"$\6$\u015a\n$\r$\16$\u015b\5$\u015e\n$\3%\3%\7%\u0162\n%\f%\16%\u0165"+
		"\13%\3%\3%\3&\3&\7&\u016b\n&\f&\16&\u016e\13&\3\'\3\'\3\'\3(\3(\3(\3)"+
		"\3)\3*\3*\3*\3+\3+\3+\3,\3,\3-\3-\3-\3.\3.\3.\3/\3/\3\60\3\60\3\61\3\61"+
		"\3\62\3\62\3\63\3\63\3\64\3\64\3\65\3\65\3\66\3\66\3\67\3\67\38\38\39"+
		"\39\3:\3:\3;\3;\3<\3<\3=\3=\3>\3>\3?\3?\3@\3@\3A\3A\3B\6B\u01ad\nB\rB"+
		"\16B\u01ae\3B\3B\3C\3C\3C\3C\7C\u01b7\nC\fC\16C\u01ba\13C\3C\3C\3C\3C"+
		"\3C\3D\3D\3D\3D\7D\u01c5\nD\fD\16D\u01c8\13D\3D\3D\3E\3E\3E\3\u01b8\2"+
		"F\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20"+
		"\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37"+
		"= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67m8o"+
		"9q:s;u<w=y>{?}@\177A\u0081B\u0083C\u0085D\u0087E\u0089\2\3\2\t\3\2\62"+
		";\3\2$$\5\2C\\aac|\6\2\62;C\\aac|\6\2\13\f\17\17\"\"^^\4\2\f\f\17\17\t"+
		"\2\"#%%--/\60<<BB]_\2\u01d4\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3"+
		"\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2"+
		"\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37"+
		"\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3"+
		"\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2"+
		"\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C"+
		"\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2"+
		"\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2"+
		"\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i"+
		"\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2"+
		"\2\2\2w\3\2\2\2\2y\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2\2\u0081"+
		"\3\2\2\2\2\u0083\3\2\2\2\2\u0085\3\2\2\2\2\u0087\3\2\2\2\3\u008b\3\2\2"+
		"\2\5\u008f\3\2\2\2\7\u0095\3\2\2\2\t\u009a\3\2\2\2\13\u00a4\3\2\2\2\r"+
		"\u00ab\3\2\2\2\17\u00af\3\2\2\2\21\u00b4\3\2\2\2\23\u00ba\3\2\2\2\25\u00c0"+
		"\3\2\2\2\27\u00c3\3\2\2\2\31\u00c8\3\2\2\2\33\u00ce\3\2\2\2\35\u00d2\3"+
		"\2\2\2\37\u00d8\3\2\2\2!\u00db\3\2\2\2#\u00e2\3\2\2\2%\u00e7\3\2\2\2\'"+
		"\u00ef\3\2\2\2)\u00f3\3\2\2\2+\u00f7\3\2\2\2-\u00fd\3\2\2\2/\u0104\3\2"+
		"\2\2\61\u010d\3\2\2\2\63\u0112\3\2\2\2\65\u0118\3\2\2\2\67\u0120\3\2\2"+
		"\29\u0127\3\2\2\2;\u0132\3\2\2\2=\u0139\3\2\2\2?\u0140\3\2\2\2A\u0146"+
		"\3\2\2\2C\u014c\3\2\2\2E\u014f\3\2\2\2G\u0153\3\2\2\2I\u015f\3\2\2\2K"+
		"\u0168\3\2\2\2M\u016f\3\2\2\2O\u0172\3\2\2\2Q\u0175\3\2\2\2S\u0177\3\2"+
		"\2\2U\u017a\3\2\2\2W\u017d\3\2\2\2Y\u017f\3\2\2\2[\u0182\3\2\2\2]\u0185"+
		"\3\2\2\2_\u0187\3\2\2\2a\u0189\3\2\2\2c\u018b\3\2\2\2e\u018d\3\2\2\2g"+
		"\u018f\3\2\2\2i\u0191\3\2\2\2k\u0193\3\2\2\2m\u0195\3\2\2\2o\u0197\3\2"+
		"\2\2q\u0199\3\2\2\2s\u019b\3\2\2\2u\u019d\3\2\2\2w\u019f\3\2\2\2y\u01a1"+
		"\3\2\2\2{\u01a3\3\2\2\2}\u01a5\3\2\2\2\177\u01a7\3\2\2\2\u0081\u01a9\3"+
		"\2\2\2\u0083\u01ac\3\2\2\2\u0085\u01b2\3\2\2\2\u0087\u01c0\3\2\2\2\u0089"+
		"\u01cb\3\2\2\2\u008b\u008c\7k\2\2\u008c\u008d\7p\2\2\u008d\u008e\7v\2"+
		"\2\u008e\4\3\2\2\2\u008f\u0090\7h\2\2\u0090\u0091\7n\2\2\u0091\u0092\7"+
		"q\2\2\u0092\u0093\7c\2\2\u0093\u0094\7v\2\2\u0094\6\3\2\2\2\u0095\u0096"+
		"\7d\2\2\u0096\u0097\7q\2\2\u0097\u0098\7q\2\2\u0098\u0099\7n\2\2\u0099"+
		"\b\3\2\2\2\u009a\u009b\7e\2\2\u009b\u009c\7j\2\2\u009c\u009d\7c\2\2\u009d"+
		"\u009e\7t\2\2\u009e\u009f\7c\2\2\u009f\u00a0\7e\2\2\u00a0\u00a1\7v\2\2"+
		"\u00a1\u00a2\7g\2\2\u00a2\u00a3\7t\2\2\u00a3\n\3\2\2\2\u00a4\u00a5\7U"+
		"\2\2\u00a5\u00a6\7v\2\2\u00a6\u00a7\7t\2\2\u00a7\u00a8\7k\2\2\u00a8\u00a9"+
		"\7p\2\2\u00a9\u00aa\7i\2\2\u00aa\f\3\2\2\2\u00ab\u00ac\7p\2\2\u00ac\u00ad"+
		"\7k\2\2\u00ad\u00ae\7n\2\2\u00ae\16\3\2\2\2\u00af\u00b0\7v\2\2\u00b0\u00b1"+
		"\7t\2\2\u00b1\u00b2\7w\2\2\u00b2\u00b3\7g\2\2\u00b3\20\3\2\2\2\u00b4\u00b5"+
		"\7h\2\2\u00b5\u00b6\7c\2\2\u00b6\u00b7\7n\2\2\u00b7\u00b8\7u\2\2\u00b8"+
		"\u00b9\7g\2\2\u00b9\22\3\2\2\2\u00ba\u00bb\7r\2\2\u00bb\u00bc\7t\2\2\u00bc"+
		"\u00bd\7k\2\2\u00bd\u00be\7p\2\2\u00be\u00bf\7v\2\2\u00bf\24\3\2\2\2\u00c0"+
		"\u00c1\7k\2\2\u00c1\u00c2\7h\2\2\u00c2\26\3\2\2\2\u00c3\u00c4\7g\2\2\u00c4"+
		"\u00c5\7n\2\2\u00c5\u00c6\7u\2\2\u00c6\u00c7\7g\2\2\u00c7\30\3\2\2\2\u00c8"+
		"\u00c9\7y\2\2\u00c9\u00ca\7j\2\2\u00ca\u00cb\7k\2\2\u00cb\u00cc\7n\2\2"+
		"\u00cc\u00cd\7g\2\2\u00cd\32\3\2\2\2\u00ce\u00cf\7h\2\2\u00cf\u00d0\7"+
		"q\2\2\u00d0\u00d1\7t\2\2\u00d1\34\3\2\2\2\u00d2\u00d3\7i\2\2\u00d3\u00d4"+
		"\7w\2\2\u00d4\u00d5\7c\2\2\u00d5\u00d6\7t\2\2\u00d6\u00d7\7f\2\2\u00d7"+
		"\36\3\2\2\2\u00d8\u00d9\7k\2\2\u00d9\u00da\7p\2\2\u00da \3\2\2\2\u00db"+
		"\u00dc\7u\2\2\u00dc\u00dd\7y\2\2\u00dd\u00de\7k\2\2\u00de\u00df\7v\2\2"+
		"\u00df\u00e0\7e\2\2\u00e0\u00e1\7j\2\2\u00e1\"\3\2\2\2\u00e2\u00e3\7e"+
		"\2\2\u00e3\u00e4\7c\2\2\u00e4\u00e5\7u\2\2\u00e5\u00e6\7g\2\2\u00e6$\3"+
		"\2\2\2\u00e7\u00e8\7f\2\2\u00e8\u00e9\7g\2\2\u00e9\u00ea\7h\2\2\u00ea"+
		"\u00eb\7c\2\2\u00eb\u00ec\7w\2\2\u00ec\u00ed\7n\2\2\u00ed\u00ee\7v\2\2"+
		"\u00ee&\3\2\2\2\u00ef\u00f0\7x\2\2\u00f0\u00f1\7c\2\2\u00f1\u00f2\7t\2"+
		"\2\u00f2(\3\2\2\2\u00f3\u00f4\7n\2\2\u00f4\u00f5\7g\2\2\u00f5\u00f6\7"+
		"v\2\2\u00f6*\3\2\2\2\u00f7\u00f8\7d\2\2\u00f8\u00f9\7t\2\2\u00f9\u00fa"+
		"\7g\2\2\u00fa\u00fb\7c\2\2\u00fb\u00fc\7m\2\2\u00fc,\3\2\2\2\u00fd\u00fe"+
		"\7t\2\2\u00fe\u00ff\7g\2\2\u00ff\u0100\7v\2\2\u0100\u0101\7w\2\2\u0101"+
		"\u0102\7t\2\2\u0102\u0103\7p\2\2\u0103.\3\2\2\2\u0104\u0105\7e\2\2\u0105"+
		"\u0106\7q\2\2\u0106\u0107\7p\2\2\u0107\u0108\7v\2\2\u0108\u0109\7k\2\2"+
		"\u0109\u010a\7p\2\2\u010a\u010b\7w\2\2\u010b\u010c\7g\2\2\u010c\60\3\2"+
		"\2\2\u010d\u010e\7h\2\2\u010e\u010f\7w\2\2\u010f\u0110\7p\2\2\u0110\u0111"+
		"\7e\2\2\u0111\62\3\2\2\2\u0112\u0113\7e\2\2\u0113\u0114\7q\2\2\u0114\u0115"+
		"\7w\2\2\u0115\u0116\7p\2\2\u0116\u0117\7v\2\2\u0117\64\3\2\2\2\u0118\u0119"+
		"\7K\2\2\u0119\u011a\7u\2\2\u011a\u011b\7G\2\2\u011b\u011c\7o\2\2\u011c"+
		"\u011d\7r\2\2\u011d\u011e\7v\2\2\u011e\u011f\7{\2\2\u011f\66\3\2\2\2\u0120"+
		"\u0121\7c\2\2\u0121\u0122\7r\2\2\u0122\u0123\7r\2\2\u0123\u0124\7g\2\2"+
		"\u0124\u0125\7p\2\2\u0125\u0126\7f\2\2\u01268\3\2\2\2\u0127\u0128\7t\2"+
		"\2\u0128\u0129\7g\2\2\u0129\u012a\7o\2\2\u012a\u012b\7q\2\2\u012b\u012c"+
		"\7x\2\2\u012c\u012d\7g\2\2\u012d\u012e\7N\2\2\u012e\u012f\7c\2\2\u012f"+
		"\u0130\7u\2\2\u0130\u0131\7v\2\2\u0131:\3\2\2\2\u0132\u0133\7t\2\2\u0133"+
		"\u0134\7g\2\2\u0134\u0135\7o\2\2\u0135\u0136\7q\2\2\u0136\u0137\7x\2\2"+
		"\u0137\u0138\7g\2\2\u0138<\3\2\2\2\u0139\u013a\7u\2\2\u013a\u013b\7v\2"+
		"\2\u013b\u013c\7t\2\2\u013c\u013d\7w\2\2\u013d\u013e\7e\2\2\u013e\u013f"+
		"\7v\2\2\u013f>\3\2\2\2\u0140\u0141\7x\2\2\u0141\u0142\7c\2\2\u0142\u0143"+
		"\7t\2\2\u0143\u0144\7u\2\2\u0144\u0145\7v\2\2\u0145@\3\2\2\2\u0146\u0147"+
		"\7n\2\2\u0147\u0148\7g\2\2\u0148\u0149\7v\2\2\u0149\u014a\7u\2\2\u014a"+
		"\u014b\7v\2\2\u014bB\3\2\2\2\u014c\u014d\7c\2\2\u014d\u014e\7v\2\2\u014e"+
		"D\3\2\2\2\u014f\u0150\7u\2\2\u0150\u0151\7v\2\2\u0151F\3\2\2\2\u0152\u0154"+
		"\t\2\2\2\u0153\u0152\3\2\2\2\u0154\u0155\3\2\2\2\u0155\u0153\3\2\2\2\u0155"+
		"\u0156\3\2\2\2\u0156\u015d\3\2\2\2\u0157\u0159\7\60\2\2\u0158\u015a\t"+
		"\2\2\2\u0159\u0158\3\2\2\2\u015a\u015b\3\2\2\2\u015b\u0159\3\2\2\2\u015b"+
		"\u015c\3\2\2\2\u015c\u015e\3\2\2\2\u015d\u0157\3\2\2\2\u015d\u015e\3\2"+
		"\2\2\u015eH\3\2\2\2\u015f\u0163\7$\2\2\u0160\u0162\n\3\2\2\u0161\u0160"+
		"\3\2\2\2\u0162\u0165\3\2\2\2\u0163\u0161\3\2\2\2\u0163\u0164\3\2\2\2\u0164"+
		"\u0166\3\2\2\2\u0165\u0163\3\2\2\2\u0166\u0167\7$\2\2\u0167J\3\2\2\2\u0168"+
		"\u016c\t\4\2\2\u0169\u016b\t\5\2\2\u016a\u0169\3\2\2\2\u016b\u016e\3\2"+
		"\2\2\u016c\u016a\3\2\2\2\u016c\u016d\3\2\2\2\u016dL\3\2\2\2\u016e\u016c"+
		"\3\2\2\2\u016f\u0170\7#\2\2\u0170\u0171\7?\2\2\u0171N\3\2\2\2\u0172\u0173"+
		"\7?\2\2\u0173\u0174\7?\2\2\u0174P\3\2\2\2\u0175\u0176\7#\2\2\u0176R\3"+
		"\2\2\2\u0177\u0178\7~\2\2\u0178\u0179\7~\2\2\u0179T\3\2\2\2\u017a\u017b"+
		"\7(\2\2\u017b\u017c\7(\2\2\u017cV\3\2\2\2\u017d\u017e\7?\2\2\u017eX\3"+
		"\2\2\2\u017f\u0180\7@\2\2\u0180\u0181\7?\2\2\u0181Z\3\2\2\2\u0182\u0183"+
		"\7>\2\2\u0183\u0184\7?\2\2\u0184\\\3\2\2\2\u0185\u0186\7@\2\2\u0186^\3"+
		"\2\2\2\u0187\u0188\7>\2\2\u0188`\3\2\2\2\u0189\u018a\7,\2\2\u018ab\3\2"+
		"\2\2\u018b\u018c\7\61\2\2\u018cd\3\2\2\2\u018d\u018e\7-\2\2\u018ef\3\2"+
		"\2\2\u018f\u0190\7/\2\2\u0190h\3\2\2\2\u0191\u0192\7\'\2\2\u0192j\3\2"+
		"\2\2\u0193\u0194\7*\2\2\u0194l\3\2\2\2\u0195\u0196\7+\2\2\u0196n\3\2\2"+
		"\2\u0197\u0198\7}\2\2\u0198p\3\2\2\2\u0199\u019a\7\177\2\2\u019ar\3\2"+
		"\2\2\u019b\u019c\7]\2\2\u019ct\3\2\2\2\u019d\u019e\7_\2\2\u019ev\3\2\2"+
		"\2\u019f\u01a0\7<\2\2\u01a0x\3\2\2\2\u01a1\u01a2\7.\2\2\u01a2z\3\2\2\2"+
		"\u01a3\u01a4\7=\2\2\u01a4|\3\2\2\2\u01a5\u01a6\7A\2\2\u01a6~\3\2\2\2\u01a7"+
		"\u01a8\7\60\2\2\u01a8\u0080\3\2\2\2\u01a9\u01aa\7a\2\2\u01aa\u0082\3\2"+
		"\2\2\u01ab\u01ad\t\6\2\2\u01ac\u01ab\3\2\2\2\u01ad\u01ae\3\2\2\2\u01ae"+
		"\u01ac\3\2\2\2\u01ae\u01af\3\2\2\2\u01af\u01b0\3\2\2\2\u01b0\u01b1\bB"+
		"\2\2\u01b1\u0084\3\2\2\2\u01b2\u01b3\7\61\2\2\u01b3\u01b4\7,\2\2\u01b4"+
		"\u01b8\3\2\2\2\u01b5\u01b7\13\2\2\2\u01b6\u01b5\3\2\2\2\u01b7\u01ba\3"+
		"\2\2\2\u01b8\u01b9\3\2\2\2\u01b8\u01b6\3\2\2\2\u01b9\u01bb\3\2\2\2\u01ba"+
		"\u01b8\3\2\2\2\u01bb\u01bc\7,\2\2\u01bc\u01bd\7\61\2\2\u01bd\u01be\3\2"+
		"\2\2\u01be\u01bf\bC\2\2\u01bf\u0086\3\2\2\2\u01c0\u01c1\7\61\2\2\u01c1"+
		"\u01c2\7\61\2\2\u01c2\u01c6\3\2\2\2\u01c3\u01c5\n\7\2\2\u01c4\u01c3\3"+
		"\2\2\2\u01c5\u01c8\3\2\2\2\u01c6\u01c4\3\2\2\2\u01c6\u01c7\3\2\2\2\u01c7"+
		"\u01c9\3\2\2\2\u01c8\u01c6\3\2\2\2\u01c9\u01ca\bD\2\2\u01ca\u0088\3\2"+
		"\2\2\u01cb\u01cc\7^\2\2\u01cc\u01cd\t\b\2\2\u01cd\u008a\3\2\2\2\13\2\u0155"+
		"\u015b\u015d\u0163\u016c\u01ae\u01b8\u01c6\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
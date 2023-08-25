// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SwiftGrammarParser struct {
	*antlr.BaseParser
}

var SwiftGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func swiftgrammarParserInit() {
	staticData := &SwiftGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'int'", "'float'", "'bool'", "'character'", "'String'", "'nil'",
		"'true'", "'false'", "'print'", "'if'", "'else'", "'while'", "'switch'",
		"'case'", "'default'", "'var'", "'break'", "'return'", "'func'", "",
		"", "", "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", "'>='", "'<='",
		"'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'%'", "'('", "')'", "'{'",
		"'}'", "':'", "';'", "'?'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "NIL", "TRU", "FAL",
		"PRINT", "IF", "ELSE", "WHILE", "SWITCH", "CASE", "DEFAULT", "VAR",
		"BREAK", "RETURN", "FUNC", "NUMBER", "STRING", "ID", "DIF", "IG_IG",
		"NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL",
		"DIV", "ADD", "SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER",
		"DOSPUNTOS", "PTCOMA", "INTERROGACION", "WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "stmt", "returnstmt", "funcdclstmt", "accfuncstm", "printstmt",
		"breakstmt", "ifstmt", "elseifstmt", "switchstmt", "caseStmt", "defaultCase",
		"typedDeclstmt", "untypedDeclstmt", "optionalTypedDeclstmt", "asignstmt",
		"whilestmt", "opasignstmt", "expr", "tipo",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 47, 231, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 1,
		0, 1, 0, 1, 0, 1, 1, 5, 1, 47, 8, 1, 10, 1, 12, 1, 50, 9, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		65, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 3, 4, 90, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 109, 8, 8, 10, 8,
		12, 8, 112, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 119, 8, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 4, 10,
		132, 8, 10, 11, 10, 12, 10, 133, 1, 10, 3, 10, 137, 8, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 3, 18, 186, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 201,
		8, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 5, 19, 224, 8, 19, 10, 19, 12, 19, 227, 9, 19, 1, 20, 1, 20, 1,
		20, 0, 1, 38, 21, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 38, 40, 0, 7, 1, 0, 7, 8, 2, 0, 33, 34, 37, 37, 1, 0, 35,
		36, 2, 0, 29, 29, 31, 31, 2, 0, 30, 30, 32, 32, 1, 0, 23, 24, 1, 0, 1,
		5, 242, 0, 42, 1, 0, 0, 0, 2, 48, 1, 0, 0, 0, 4, 64, 1, 0, 0, 0, 6, 66,
		1, 0, 0, 0, 8, 89, 1, 0, 0, 0, 10, 91, 1, 0, 0, 0, 12, 95, 1, 0, 0, 0,
		14, 100, 1, 0, 0, 0, 16, 102, 1, 0, 0, 0, 18, 120, 1, 0, 0, 0, 20, 127,
		1, 0, 0, 0, 22, 140, 1, 0, 0, 0, 24, 145, 1, 0, 0, 0, 26, 149, 1, 0, 0,
		0, 28, 156, 1, 0, 0, 0, 30, 161, 1, 0, 0, 0, 32, 167, 1, 0, 0, 0, 34, 171,
		1, 0, 0, 0, 36, 185, 1, 0, 0, 0, 38, 200, 1, 0, 0, 0, 40, 228, 1, 0, 0,
		0, 42, 43, 3, 2, 1, 0, 43, 44, 5, 0, 0, 1, 44, 1, 1, 0, 0, 0, 45, 47, 3,
		4, 2, 0, 46, 45, 1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 48,
		49, 1, 0, 0, 0, 49, 3, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 51, 65, 3, 12, 6,
		0, 52, 65, 3, 26, 13, 0, 53, 65, 3, 28, 14, 0, 54, 65, 3, 30, 15, 0, 55,
		65, 3, 32, 16, 0, 56, 65, 3, 16, 8, 0, 57, 65, 3, 20, 10, 0, 58, 65, 3,
		34, 17, 0, 59, 65, 3, 36, 18, 0, 60, 65, 3, 14, 7, 0, 61, 65, 3, 8, 4,
		0, 62, 65, 3, 10, 5, 0, 63, 65, 3, 6, 3, 0, 64, 51, 1, 0, 0, 0, 64, 52,
		1, 0, 0, 0, 64, 53, 1, 0, 0, 0, 64, 54, 1, 0, 0, 0, 64, 55, 1, 0, 0, 0,
		64, 56, 1, 0, 0, 0, 64, 57, 1, 0, 0, 0, 64, 58, 1, 0, 0, 0, 64, 59, 1,
		0, 0, 0, 64, 60, 1, 0, 0, 0, 64, 61, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64,
		63, 1, 0, 0, 0, 65, 5, 1, 0, 0, 0, 66, 67, 5, 18, 0, 0, 67, 68, 3, 38,
		19, 0, 68, 69, 5, 43, 0, 0, 69, 7, 1, 0, 0, 0, 70, 71, 5, 19, 0, 0, 71,
		72, 5, 22, 0, 0, 72, 73, 5, 38, 0, 0, 73, 74, 5, 39, 0, 0, 74, 75, 5, 40,
		0, 0, 75, 76, 3, 2, 1, 0, 76, 77, 5, 41, 0, 0, 77, 90, 1, 0, 0, 0, 78,
		79, 5, 19, 0, 0, 79, 80, 5, 22, 0, 0, 80, 81, 5, 38, 0, 0, 81, 82, 5, 39,
		0, 0, 82, 83, 5, 36, 0, 0, 83, 84, 5, 31, 0, 0, 84, 85, 3, 40, 20, 0, 85,
		86, 5, 40, 0, 0, 86, 87, 3, 2, 1, 0, 87, 88, 5, 41, 0, 0, 88, 90, 1, 0,
		0, 0, 89, 70, 1, 0, 0, 0, 89, 78, 1, 0, 0, 0, 90, 9, 1, 0, 0, 0, 91, 92,
		5, 22, 0, 0, 92, 93, 5, 38, 0, 0, 93, 94, 5, 39, 0, 0, 94, 11, 1, 0, 0,
		0, 95, 96, 5, 9, 0, 0, 96, 97, 5, 38, 0, 0, 97, 98, 3, 38, 19, 0, 98, 99,
		5, 39, 0, 0, 99, 13, 1, 0, 0, 0, 100, 101, 5, 17, 0, 0, 101, 15, 1, 0,
		0, 0, 102, 103, 5, 10, 0, 0, 103, 104, 3, 38, 19, 0, 104, 105, 5, 40, 0,
		0, 105, 106, 3, 2, 1, 0, 106, 110, 5, 41, 0, 0, 107, 109, 3, 18, 9, 0,
		108, 107, 1, 0, 0, 0, 109, 112, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 110,
		111, 1, 0, 0, 0, 111, 118, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 113, 114,
		5, 11, 0, 0, 114, 115, 5, 40, 0, 0, 115, 116, 3, 2, 1, 0, 116, 117, 5,
		41, 0, 0, 117, 119, 1, 0, 0, 0, 118, 113, 1, 0, 0, 0, 118, 119, 1, 0, 0,
		0, 119, 17, 1, 0, 0, 0, 120, 121, 5, 11, 0, 0, 121, 122, 5, 10, 0, 0, 122,
		123, 3, 38, 19, 0, 123, 124, 5, 40, 0, 0, 124, 125, 3, 2, 1, 0, 125, 126,
		5, 41, 0, 0, 126, 19, 1, 0, 0, 0, 127, 128, 5, 13, 0, 0, 128, 129, 3, 38,
		19, 0, 129, 131, 5, 40, 0, 0, 130, 132, 3, 22, 11, 0, 131, 130, 1, 0, 0,
		0, 132, 133, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134,
		136, 1, 0, 0, 0, 135, 137, 3, 24, 12, 0, 136, 135, 1, 0, 0, 0, 136, 137,
		1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 5, 41, 0, 0, 139, 21, 1, 0,
		0, 0, 140, 141, 5, 14, 0, 0, 141, 142, 3, 38, 19, 0, 142, 143, 5, 42, 0,
		0, 143, 144, 3, 2, 1, 0, 144, 23, 1, 0, 0, 0, 145, 146, 5, 15, 0, 0, 146,
		147, 5, 42, 0, 0, 147, 148, 3, 2, 1, 0, 148, 25, 1, 0, 0, 0, 149, 150,
		5, 16, 0, 0, 150, 151, 5, 22, 0, 0, 151, 152, 5, 42, 0, 0, 152, 153, 3,
		40, 20, 0, 153, 154, 5, 28, 0, 0, 154, 155, 3, 38, 19, 0, 155, 27, 1, 0,
		0, 0, 156, 157, 5, 16, 0, 0, 157, 158, 5, 22, 0, 0, 158, 159, 5, 28, 0,
		0, 159, 160, 3, 38, 19, 0, 160, 29, 1, 0, 0, 0, 161, 162, 5, 16, 0, 0,
		162, 163, 5, 22, 0, 0, 163, 164, 5, 42, 0, 0, 164, 165, 3, 40, 20, 0, 165,
		166, 5, 44, 0, 0, 166, 31, 1, 0, 0, 0, 167, 168, 5, 22, 0, 0, 168, 169,
		5, 28, 0, 0, 169, 170, 3, 38, 19, 0, 170, 33, 1, 0, 0, 0, 171, 172, 5,
		12, 0, 0, 172, 173, 3, 38, 19, 0, 173, 174, 5, 40, 0, 0, 174, 175, 3, 2,
		1, 0, 175, 176, 5, 41, 0, 0, 176, 35, 1, 0, 0, 0, 177, 178, 5, 22, 0, 0,
		178, 179, 5, 35, 0, 0, 179, 180, 5, 28, 0, 0, 180, 186, 3, 38, 19, 0, 181,
		182, 5, 22, 0, 0, 182, 183, 5, 36, 0, 0, 183, 184, 5, 28, 0, 0, 184, 186,
		3, 38, 19, 0, 185, 177, 1, 0, 0, 0, 185, 181, 1, 0, 0, 0, 186, 37, 1, 0,
		0, 0, 187, 188, 6, 19, -1, 0, 188, 189, 5, 38, 0, 0, 189, 190, 3, 38, 19,
		0, 190, 191, 5, 39, 0, 0, 191, 201, 1, 0, 0, 0, 192, 193, 5, 36, 0, 0,
		193, 201, 3, 38, 19, 14, 194, 201, 5, 20, 0, 0, 195, 201, 5, 22, 0, 0,
		196, 201, 5, 21, 0, 0, 197, 201, 7, 0, 0, 0, 198, 201, 5, 6, 0, 0, 199,
		201, 3, 10, 5, 0, 200, 187, 1, 0, 0, 0, 200, 192, 1, 0, 0, 0, 200, 194,
		1, 0, 0, 0, 200, 195, 1, 0, 0, 0, 200, 196, 1, 0, 0, 0, 200, 197, 1, 0,
		0, 0, 200, 198, 1, 0, 0, 0, 200, 199, 1, 0, 0, 0, 201, 225, 1, 0, 0, 0,
		202, 203, 10, 13, 0, 0, 203, 204, 7, 1, 0, 0, 204, 224, 3, 38, 19, 14,
		205, 206, 10, 12, 0, 0, 206, 207, 7, 2, 0, 0, 207, 224, 3, 38, 19, 13,
		208, 209, 10, 11, 0, 0, 209, 210, 7, 3, 0, 0, 210, 224, 3, 38, 19, 12,
		211, 212, 10, 10, 0, 0, 212, 213, 7, 4, 0, 0, 213, 224, 3, 38, 19, 11,
		214, 215, 10, 9, 0, 0, 215, 216, 7, 5, 0, 0, 216, 224, 3, 38, 19, 10, 217,
		218, 10, 8, 0, 0, 218, 219, 5, 27, 0, 0, 219, 224, 3, 38, 19, 9, 220, 221,
		10, 7, 0, 0, 221, 222, 5, 26, 0, 0, 222, 224, 3, 38, 19, 8, 223, 202, 1,
		0, 0, 0, 223, 205, 1, 0, 0, 0, 223, 208, 1, 0, 0, 0, 223, 211, 1, 0, 0,
		0, 223, 214, 1, 0, 0, 0, 223, 217, 1, 0, 0, 0, 223, 220, 1, 0, 0, 0, 224,
		227, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 39, 1,
		0, 0, 0, 227, 225, 1, 0, 0, 0, 228, 229, 7, 6, 0, 0, 229, 41, 1, 0, 0,
		0, 11, 48, 64, 89, 110, 118, 133, 136, 185, 200, 223, 225,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SwiftGrammarParserInit initializes any static state used to implement SwiftGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSwiftGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftGrammarParserInit() {
	staticData := &SwiftGrammarParserStaticData
	staticData.once.Do(swiftgrammarParserInit)
}

// NewSwiftGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSwiftGrammarParser(input antlr.TokenStream) *SwiftGrammarParser {
	SwiftGrammarParserInit()
	this := new(SwiftGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SwiftGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SwiftGrammar.g4"

	return this
}

// SwiftGrammarParser tokens.
const (
	SwiftGrammarParserEOF           = antlr.TokenEOF
	SwiftGrammarParserINT           = 1
	SwiftGrammarParserFLOAT         = 2
	SwiftGrammarParserBOOL          = 3
	SwiftGrammarParserCHARACTER     = 4
	SwiftGrammarParserPSTRING       = 5
	SwiftGrammarParserNIL           = 6
	SwiftGrammarParserTRU           = 7
	SwiftGrammarParserFAL           = 8
	SwiftGrammarParserPRINT         = 9
	SwiftGrammarParserIF            = 10
	SwiftGrammarParserELSE          = 11
	SwiftGrammarParserWHILE         = 12
	SwiftGrammarParserSWITCH        = 13
	SwiftGrammarParserCASE          = 14
	SwiftGrammarParserDEFAULT       = 15
	SwiftGrammarParserVAR           = 16
	SwiftGrammarParserBREAK         = 17
	SwiftGrammarParserRETURN        = 18
	SwiftGrammarParserFUNC          = 19
	SwiftGrammarParserNUMBER        = 20
	SwiftGrammarParserSTRING        = 21
	SwiftGrammarParserID            = 22
	SwiftGrammarParserDIF           = 23
	SwiftGrammarParserIG_IG         = 24
	SwiftGrammarParserNOT           = 25
	SwiftGrammarParserOR            = 26
	SwiftGrammarParserAND           = 27
	SwiftGrammarParserIG            = 28
	SwiftGrammarParserMAY_IG        = 29
	SwiftGrammarParserMEN_IG        = 30
	SwiftGrammarParserMAYOR         = 31
	SwiftGrammarParserMENOR         = 32
	SwiftGrammarParserMUL           = 33
	SwiftGrammarParserDIV           = 34
	SwiftGrammarParserADD           = 35
	SwiftGrammarParserSUB           = 36
	SwiftGrammarParserMOD           = 37
	SwiftGrammarParserPARIZQ        = 38
	SwiftGrammarParserPARDER        = 39
	SwiftGrammarParserLLAVEIZQ      = 40
	SwiftGrammarParserLLAVEDER      = 41
	SwiftGrammarParserDOSPUNTOS     = 42
	SwiftGrammarParserPTCOMA        = 43
	SwiftGrammarParserINTERROGACION = 44
	SwiftGrammarParserWHITESPACE    = 45
	SwiftGrammarParserCOMMENT       = 46
	SwiftGrammarParserLINE_COMMENT  = 47
)

// SwiftGrammarParser rules.
const (
	SwiftGrammarParserRULE_s                     = 0
	SwiftGrammarParserRULE_block                 = 1
	SwiftGrammarParserRULE_stmt                  = 2
	SwiftGrammarParserRULE_returnstmt            = 3
	SwiftGrammarParserRULE_funcdclstmt           = 4
	SwiftGrammarParserRULE_accfuncstm            = 5
	SwiftGrammarParserRULE_printstmt             = 6
	SwiftGrammarParserRULE_breakstmt             = 7
	SwiftGrammarParserRULE_ifstmt                = 8
	SwiftGrammarParserRULE_elseifstmt            = 9
	SwiftGrammarParserRULE_switchstmt            = 10
	SwiftGrammarParserRULE_caseStmt              = 11
	SwiftGrammarParserRULE_defaultCase           = 12
	SwiftGrammarParserRULE_typedDeclstmt         = 13
	SwiftGrammarParserRULE_untypedDeclstmt       = 14
	SwiftGrammarParserRULE_optionalTypedDeclstmt = 15
	SwiftGrammarParserRULE_asignstmt             = 16
	SwiftGrammarParserRULE_whilestmt             = 17
	SwiftGrammarParserRULE_opasignstmt           = 18
	SwiftGrammarParserRULE_expr                  = 19
	SwiftGrammarParserRULE_tipo                  = 20
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SContext) EOF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserEOF, 0)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterS(s)
	}
}

func (s *SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitS(s)
	}
}

func (s *SContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitS(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SwiftGrammarParserRULE_s)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Block()
	}
	{
		p.SetState(43)
		p.Match(SwiftGrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SwiftGrammarParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&5191168) != 0 {
		{
			p.SetState(45)
			p.Stmt()
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Printstmt() IPrintstmtContext
	TypedDeclstmt() ITypedDeclstmtContext
	UntypedDeclstmt() IUntypedDeclstmtContext
	OptionalTypedDeclstmt() IOptionalTypedDeclstmtContext
	Asignstmt() IAsignstmtContext
	Ifstmt() IIfstmtContext
	Switchstmt() ISwitchstmtContext
	Whilestmt() IWhilestmtContext
	Opasignstmt() IOpasignstmtContext
	Breakstmt() IBreakstmtContext
	Funcdclstmt() IFuncdclstmtContext
	Accfuncstm() IAccfuncstmContext
	Returnstmt() IReturnstmtContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Printstmt() IPrintstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintstmtContext)
}

func (s *StmtContext) TypedDeclstmt() ITypedDeclstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypedDeclstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypedDeclstmtContext)
}

func (s *StmtContext) UntypedDeclstmt() IUntypedDeclstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUntypedDeclstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUntypedDeclstmtContext)
}

func (s *StmtContext) OptionalTypedDeclstmt() IOptionalTypedDeclstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionalTypedDeclstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionalTypedDeclstmtContext)
}

func (s *StmtContext) Asignstmt() IAsignstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignstmtContext)
}

func (s *StmtContext) Ifstmt() IIfstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *StmtContext) Switchstmt() ISwitchstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchstmtContext)
}

func (s *StmtContext) Whilestmt() IWhilestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhilestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhilestmtContext)
}

func (s *StmtContext) Opasignstmt() IOpasignstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpasignstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpasignstmtContext)
}

func (s *StmtContext) Breakstmt() IBreakstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakstmtContext)
}

func (s *StmtContext) Funcdclstmt() IFuncdclstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncdclstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncdclstmtContext)
}

func (s *StmtContext) Accfuncstm() IAccfuncstmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccfuncstmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccfuncstmContext)
}

func (s *StmtContext) Returnstmt() IReturnstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnstmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SwiftGrammarParserRULE_stmt)
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(51)
			p.Printstmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)
			p.TypedDeclstmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(53)
			p.UntypedDeclstmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(54)
			p.OptionalTypedDeclstmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(55)
			p.Asignstmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(56)
			p.Ifstmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(57)
			p.Switchstmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(58)
			p.Whilestmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(59)
			p.Opasignstmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(60)
			p.Breakstmt()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(61)
			p.Funcdclstmt()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(62)
			p.Accfuncstm()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(63)
			p.Returnstmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnstmtContext is an interface to support dynamic dispatch.
type IReturnstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expr() IExprContext
	PTCOMA() antlr.TerminalNode

	// IsReturnstmtContext differentiates from other interfaces.
	IsReturnstmtContext()
}

type ReturnstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnstmtContext() *ReturnstmtContext {
	var p = new(ReturnstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_returnstmt
	return p
}

func InitEmptyReturnstmtContext(p *ReturnstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_returnstmt
}

func (*ReturnstmtContext) IsReturnstmtContext() {}

func NewReturnstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnstmtContext {
	var p = new(ReturnstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_returnstmt

	return p
}

func (s *ReturnstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnstmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRETURN, 0)
}

func (s *ReturnstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnstmtContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPTCOMA, 0)
}

func (s *ReturnstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterReturnstmt(s)
	}
}

func (s *ReturnstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitReturnstmt(s)
	}
}

func (s *ReturnstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitReturnstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Returnstmt() (localctx IReturnstmtContext) {
	localctx = NewReturnstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SwiftGrammarParserRULE_returnstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(SwiftGrammarParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.expr(0)
	}
	{
		p.SetState(68)
		p.Match(SwiftGrammarParserPTCOMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncdclstmtContext is an interface to support dynamic dispatch.
type IFuncdclstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFuncdclstmtContext differentiates from other interfaces.
	IsFuncdclstmtContext()
}

type FuncdclstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncdclstmtContext() *FuncdclstmtContext {
	var p = new(FuncdclstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_funcdclstmt
	return p
}

func InitEmptyFuncdclstmtContext(p *FuncdclstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_funcdclstmt
}

func (*FuncdclstmtContext) IsFuncdclstmtContext() {}

func NewFuncdclstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncdclstmtContext {
	var p = new(FuncdclstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_funcdclstmt

	return p
}

func (s *FuncdclstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncdclstmtContext) CopyAll(ctx *FuncdclstmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FuncdclstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncdclstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncionNormalContext struct {
	FuncdclstmtContext
}

func NewFuncionNormalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncionNormalContext {
	var p = new(FuncionNormalContext)

	InitEmptyFuncdclstmtContext(&p.FuncdclstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncdclstmtContext))

	return p
}

func (s *FuncionNormalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionNormalContext) FUNC() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFUNC, 0)
}

func (s *FuncionNormalContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *FuncionNormalContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *FuncionNormalContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *FuncionNormalContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *FuncionNormalContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncionNormalContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *FuncionNormalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterFuncionNormal(s)
	}
}

func (s *FuncionNormalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitFuncionNormal(s)
	}
}

func (s *FuncionNormalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitFuncionNormal(s)

	default:
		return t.VisitChildren(s)
	}
}

type FuncionRetornoContext struct {
	FuncdclstmtContext
}

func NewFuncionRetornoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncionRetornoContext {
	var p = new(FuncionRetornoContext)

	InitEmptyFuncdclstmtContext(&p.FuncdclstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncdclstmtContext))

	return p
}

func (s *FuncionRetornoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionRetornoContext) FUNC() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFUNC, 0)
}

func (s *FuncionRetornoContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *FuncionRetornoContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *FuncionRetornoContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *FuncionRetornoContext) SUB() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSUB, 0)
}

func (s *FuncionRetornoContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMAYOR, 0)
}

func (s *FuncionRetornoContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *FuncionRetornoContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *FuncionRetornoContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncionRetornoContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *FuncionRetornoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterFuncionRetorno(s)
	}
}

func (s *FuncionRetornoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitFuncionRetorno(s)
	}
}

func (s *FuncionRetornoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitFuncionRetorno(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Funcdclstmt() (localctx IFuncdclstmtContext) {
	localctx = NewFuncdclstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SwiftGrammarParserRULE_funcdclstmt)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFuncionNormalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Match(SwiftGrammarParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.Block()
		}
		{
			p.SetState(76)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFuncionRetornoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(78)
			p.Match(SwiftGrammarParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.Match(SwiftGrammarParserSUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Match(SwiftGrammarParserMAYOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.Tipo()
		}
		{
			p.SetState(85)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(86)
			p.Block()
		}
		{
			p.SetState(87)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAccfuncstmContext is an interface to support dynamic dispatch.
type IAccfuncstmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode

	// IsAccfuncstmContext differentiates from other interfaces.
	IsAccfuncstmContext()
}

type AccfuncstmContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccfuncstmContext() *AccfuncstmContext {
	var p = new(AccfuncstmContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_accfuncstm
	return p
}

func InitEmptyAccfuncstmContext(p *AccfuncstmContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_accfuncstm
}

func (*AccfuncstmContext) IsAccfuncstmContext() {}

func NewAccfuncstmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccfuncstmContext {
	var p = new(AccfuncstmContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_accfuncstm

	return p
}

func (s *AccfuncstmContext) GetParser() antlr.Parser { return s.parser }

func (s *AccfuncstmContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *AccfuncstmContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *AccfuncstmContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *AccfuncstmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccfuncstmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccfuncstmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAccfuncstm(s)
	}
}

func (s *AccfuncstmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAccfuncstm(s)
	}
}

func (s *AccfuncstmContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitAccfuncstm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Accfuncstm() (localctx IAccfuncstmContext) {
	localctx = NewAccfuncstmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SwiftGrammarParserRULE_accfuncstm)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintstmtContext is an interface to support dynamic dispatch.
type IPrintstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode

	// IsPrintstmtContext differentiates from other interfaces.
	IsPrintstmtContext()
}

type PrintstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintstmtContext() *PrintstmtContext {
	var p = new(PrintstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_printstmt
	return p
}

func InitEmptyPrintstmtContext(p *PrintstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_printstmt
}

func (*PrintstmtContext) IsPrintstmtContext() {}

func NewPrintstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintstmtContext {
	var p = new(PrintstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_printstmt

	return p
}

func (s *PrintstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintstmtContext) PRINT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPRINT, 0)
}

func (s *PrintstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *PrintstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *PrintstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterPrintstmt(s)
	}
}

func (s *PrintstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitPrintstmt(s)
	}
}

func (s *PrintstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitPrintstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Printstmt() (localctx IPrintstmtContext) {
	localctx = NewPrintstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SwiftGrammarParserRULE_printstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(SwiftGrammarParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(97)
		p.expr(0)
	}
	{
		p.SetState(98)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreakstmtContext is an interface to support dynamic dispatch.
type IBreakstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BREAK() antlr.TerminalNode

	// IsBreakstmtContext differentiates from other interfaces.
	IsBreakstmtContext()
}

type BreakstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakstmtContext() *BreakstmtContext {
	var p = new(BreakstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_breakstmt
	return p
}

func InitEmptyBreakstmtContext(p *BreakstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_breakstmt
}

func (*BreakstmtContext) IsBreakstmtContext() {}

func NewBreakstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakstmtContext {
	var p = new(BreakstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_breakstmt

	return p
}

func (s *BreakstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakstmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserBREAK, 0)
}

func (s *BreakstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterBreakstmt(s)
	}
}

func (s *BreakstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitBreakstmt(s)
	}
}

func (s *BreakstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitBreakstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Breakstmt() (localctx IBreakstmtContext) {
	localctx = NewBreakstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SwiftGrammarParserRULE_breakstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(SwiftGrammarParserBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfstmtContext is an interface to support dynamic dispatch.
type IIfstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expr() IExprContext
	AllLLAVEIZQ() []antlr.TerminalNode
	LLAVEIZQ(i int) antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	AllLLAVEDER() []antlr.TerminalNode
	LLAVEDER(i int) antlr.TerminalNode
	AllElseifstmt() []IElseifstmtContext
	Elseifstmt(i int) IElseifstmtContext
	ELSE() antlr.TerminalNode

	// IsIfstmtContext differentiates from other interfaces.
	IsIfstmtContext()
}

type IfstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfstmtContext() *IfstmtContext {
	var p = new(IfstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt
	return p
}

func InitEmptyIfstmtContext(p *IfstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt
}

func (*IfstmtContext) IsIfstmtContext() {}

func NewIfstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstmtContext {
	var p = new(IfstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt

	return p
}

func (s *IfstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstmtContext) IF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIF, 0)
}

func (s *IfstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfstmtContext) AllLLAVEIZQ() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserLLAVEIZQ)
}

func (s *IfstmtContext) LLAVEIZQ(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, i)
}

func (s *IfstmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfstmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfstmtContext) AllLLAVEDER() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserLLAVEDER)
}

func (s *IfstmtContext) LLAVEDER(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, i)
}

func (s *IfstmtContext) AllElseifstmt() []IElseifstmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseifstmtContext); ok {
			len++
		}
	}

	tst := make([]IElseifstmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseifstmtContext); ok {
			tst[i] = t.(IElseifstmtContext)
			i++
		}
	}

	return tst
}

func (s *IfstmtContext) Elseifstmt(i int) IElseifstmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseifstmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseifstmtContext)
}

func (s *IfstmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
}

func (s *IfstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterIfstmt(s)
	}
}

func (s *IfstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitIfstmt(s)
	}
}

func (s *IfstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitIfstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Ifstmt() (localctx IIfstmtContext) {
	localctx = NewIfstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SwiftGrammarParserRULE_ifstmt)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Match(SwiftGrammarParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.expr(0)
	}
	{
		p.SetState(104)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Block()
	}
	{
		p.SetState(106)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(107)
				p.Elseifstmt()
			}

		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftGrammarParserELSE {
		{
			p.SetState(113)
			p.Match(SwiftGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
			p.Block()
		}
		{
			p.SetState(116)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseifstmtContext is an interface to support dynamic dispatch.
type IElseifstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	IF() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsElseifstmtContext differentiates from other interfaces.
	IsElseifstmtContext()
}

type ElseifstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseifstmtContext() *ElseifstmtContext {
	var p = new(ElseifstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_elseifstmt
	return p
}

func InitEmptyElseifstmtContext(p *ElseifstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_elseifstmt
}

func (*ElseifstmtContext) IsElseifstmtContext() {}

func NewElseifstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseifstmtContext {
	var p = new(ElseifstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_elseifstmt

	return p
}

func (s *ElseifstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseifstmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
}

func (s *ElseifstmtContext) IF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIF, 0)
}

func (s *ElseifstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ElseifstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *ElseifstmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseifstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *ElseifstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseifstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseifstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterElseifstmt(s)
	}
}

func (s *ElseifstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitElseifstmt(s)
	}
}

func (s *ElseifstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitElseifstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Elseifstmt() (localctx IElseifstmtContext) {
	localctx = NewElseifstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SwiftGrammarParserRULE_elseifstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Match(SwiftGrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.Match(SwiftGrammarParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.expr(0)
	}
	{
		p.SetState(123)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Block()
	}
	{
		p.SetState(125)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchstmtContext is an interface to support dynamic dispatch.
type ISwitchstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	LLAVEDER() antlr.TerminalNode
	AllCaseStmt() []ICaseStmtContext
	CaseStmt(i int) ICaseStmtContext
	DefaultCase() IDefaultCaseContext

	// IsSwitchstmtContext differentiates from other interfaces.
	IsSwitchstmtContext()
}

type SwitchstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchstmtContext() *SwitchstmtContext {
	var p = new(SwitchstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_switchstmt
	return p
}

func InitEmptySwitchstmtContext(p *SwitchstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_switchstmt
}

func (*SwitchstmtContext) IsSwitchstmtContext() {}

func NewSwitchstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchstmtContext {
	var p = new(SwitchstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_switchstmt

	return p
}

func (s *SwitchstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchstmtContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSWITCH, 0)
}

func (s *SwitchstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *SwitchstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *SwitchstmtContext) AllCaseStmt() []ICaseStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICaseStmtContext); ok {
			len++
		}
	}

	tst := make([]ICaseStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICaseStmtContext); ok {
			tst[i] = t.(ICaseStmtContext)
			i++
		}
	}

	return tst
}

func (s *SwitchstmtContext) CaseStmt(i int) ICaseStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseStmtContext)
}

func (s *SwitchstmtContext) DefaultCase() IDefaultCaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultCaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultCaseContext)
}

func (s *SwitchstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterSwitchstmt(s)
	}
}

func (s *SwitchstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitSwitchstmt(s)
	}
}

func (s *SwitchstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitSwitchstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Switchstmt() (localctx ISwitchstmtContext) {
	localctx = NewSwitchstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SwiftGrammarParserRULE_switchstmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(SwiftGrammarParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.expr(0)
	}
	{
		p.SetState(129)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SwiftGrammarParserCASE {
		{
			p.SetState(130)
			p.CaseStmt()
		}

		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftGrammarParserDEFAULT {
		{
			p.SetState(135)
			p.DefaultCase()
		}

	}
	{
		p.SetState(138)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICaseStmtContext is an interface to support dynamic dispatch.
type ICaseStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	Expr() IExprContext
	DOSPUNTOS() antlr.TerminalNode
	Block() IBlockContext

	// IsCaseStmtContext differentiates from other interfaces.
	IsCaseStmtContext()
}

type CaseStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseStmtContext() *CaseStmtContext {
	var p = new(CaseStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_caseStmt
	return p
}

func InitEmptyCaseStmtContext(p *CaseStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_caseStmt
}

func (*CaseStmtContext) IsCaseStmtContext() {}

func NewCaseStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseStmtContext {
	var p = new(CaseStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_caseStmt

	return p
}

func (s *CaseStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseStmtContext) CASE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCASE, 0)
}

func (s *CaseStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CaseStmtContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPUNTOS, 0)
}

func (s *CaseStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CaseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterCaseStmt(s)
	}
}

func (s *CaseStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitCaseStmt(s)
	}
}

func (s *CaseStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitCaseStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) CaseStmt() (localctx ICaseStmtContext) {
	localctx = NewCaseStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SwiftGrammarParserRULE_caseStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(SwiftGrammarParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.expr(0)
	}
	{
		p.SetState(142)
		p.Match(SwiftGrammarParserDOSPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefaultCaseContext is an interface to support dynamic dispatch.
type IDefaultCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFAULT() antlr.TerminalNode
	DOSPUNTOS() antlr.TerminalNode
	Block() IBlockContext

	// IsDefaultCaseContext differentiates from other interfaces.
	IsDefaultCaseContext()
}

type DefaultCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultCaseContext() *DefaultCaseContext {
	var p = new(DefaultCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_defaultCase
	return p
}

func InitEmptyDefaultCaseContext(p *DefaultCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_defaultCase
}

func (*DefaultCaseContext) IsDefaultCaseContext() {}

func NewDefaultCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultCaseContext {
	var p = new(DefaultCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_defaultCase

	return p
}

func (s *DefaultCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultCaseContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDEFAULT, 0)
}

func (s *DefaultCaseContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPUNTOS, 0)
}

func (s *DefaultCaseContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DefaultCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDefaultCase(s)
	}
}

func (s *DefaultCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDefaultCase(s)
	}
}

func (s *DefaultCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitDefaultCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) DefaultCase() (localctx IDefaultCaseContext) {
	localctx = NewDefaultCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SwiftGrammarParserRULE_defaultCase)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(SwiftGrammarParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.Match(SwiftGrammarParserDOSPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypedDeclstmtContext is an interface to support dynamic dispatch.
type ITypedDeclstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	DOSPUNTOS() antlr.TerminalNode
	Tipo() ITipoContext
	IG() antlr.TerminalNode
	Expr() IExprContext

	// IsTypedDeclstmtContext differentiates from other interfaces.
	IsTypedDeclstmtContext()
}

type TypedDeclstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypedDeclstmtContext() *TypedDeclstmtContext {
	var p = new(TypedDeclstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_typedDeclstmt
	return p
}

func InitEmptyTypedDeclstmtContext(p *TypedDeclstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_typedDeclstmt
}

func (*TypedDeclstmtContext) IsTypedDeclstmtContext() {}

func NewTypedDeclstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedDeclstmtContext {
	var p = new(TypedDeclstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_typedDeclstmt

	return p
}

func (s *TypedDeclstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *TypedDeclstmtContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *TypedDeclstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *TypedDeclstmtContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPUNTOS, 0)
}

func (s *TypedDeclstmtContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *TypedDeclstmtContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *TypedDeclstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TypedDeclstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedDeclstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypedDeclstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterTypedDeclstmt(s)
	}
}

func (s *TypedDeclstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitTypedDeclstmt(s)
	}
}

func (s *TypedDeclstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitTypedDeclstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) TypedDeclstmt() (localctx ITypedDeclstmtContext) {
	localctx = NewTypedDeclstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SwiftGrammarParserRULE_typedDeclstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(SwiftGrammarParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Match(SwiftGrammarParserDOSPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Tipo()
	}
	{
		p.SetState(153)
		p.Match(SwiftGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUntypedDeclstmtContext is an interface to support dynamic dispatch.
type IUntypedDeclstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	IG() antlr.TerminalNode
	Expr() IExprContext

	// IsUntypedDeclstmtContext differentiates from other interfaces.
	IsUntypedDeclstmtContext()
}

type UntypedDeclstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUntypedDeclstmtContext() *UntypedDeclstmtContext {
	var p = new(UntypedDeclstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_untypedDeclstmt
	return p
}

func InitEmptyUntypedDeclstmtContext(p *UntypedDeclstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_untypedDeclstmt
}

func (*UntypedDeclstmtContext) IsUntypedDeclstmtContext() {}

func NewUntypedDeclstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UntypedDeclstmtContext {
	var p = new(UntypedDeclstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_untypedDeclstmt

	return p
}

func (s *UntypedDeclstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *UntypedDeclstmtContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *UntypedDeclstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *UntypedDeclstmtContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *UntypedDeclstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UntypedDeclstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UntypedDeclstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UntypedDeclstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterUntypedDeclstmt(s)
	}
}

func (s *UntypedDeclstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitUntypedDeclstmt(s)
	}
}

func (s *UntypedDeclstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitUntypedDeclstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) UntypedDeclstmt() (localctx IUntypedDeclstmtContext) {
	localctx = NewUntypedDeclstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SwiftGrammarParserRULE_untypedDeclstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(SwiftGrammarParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Match(SwiftGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionalTypedDeclstmtContext is an interface to support dynamic dispatch.
type IOptionalTypedDeclstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	DOSPUNTOS() antlr.TerminalNode
	Tipo() ITipoContext
	INTERROGACION() antlr.TerminalNode

	// IsOptionalTypedDeclstmtContext differentiates from other interfaces.
	IsOptionalTypedDeclstmtContext()
}

type OptionalTypedDeclstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionalTypedDeclstmtContext() *OptionalTypedDeclstmtContext {
	var p = new(OptionalTypedDeclstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_optionalTypedDeclstmt
	return p
}

func InitEmptyOptionalTypedDeclstmtContext(p *OptionalTypedDeclstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_optionalTypedDeclstmt
}

func (*OptionalTypedDeclstmtContext) IsOptionalTypedDeclstmtContext() {}

func NewOptionalTypedDeclstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionalTypedDeclstmtContext {
	var p = new(OptionalTypedDeclstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_optionalTypedDeclstmt

	return p
}

func (s *OptionalTypedDeclstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionalTypedDeclstmtContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *OptionalTypedDeclstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *OptionalTypedDeclstmtContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPUNTOS, 0)
}

func (s *OptionalTypedDeclstmtContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *OptionalTypedDeclstmtContext) INTERROGACION() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserINTERROGACION, 0)
}

func (s *OptionalTypedDeclstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionalTypedDeclstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionalTypedDeclstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterOptionalTypedDeclstmt(s)
	}
}

func (s *OptionalTypedDeclstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitOptionalTypedDeclstmt(s)
	}
}

func (s *OptionalTypedDeclstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitOptionalTypedDeclstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) OptionalTypedDeclstmt() (localctx IOptionalTypedDeclstmtContext) {
	localctx = NewOptionalTypedDeclstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SwiftGrammarParserRULE_optionalTypedDeclstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(SwiftGrammarParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.Match(SwiftGrammarParserDOSPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(164)
		p.Tipo()
	}
	{
		p.SetState(165)
		p.Match(SwiftGrammarParserINTERROGACION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignstmtContext is an interface to support dynamic dispatch.
type IAsignstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	IG() antlr.TerminalNode
	Expr() IExprContext

	// IsAsignstmtContext differentiates from other interfaces.
	IsAsignstmtContext()
}

type AsignstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignstmtContext() *AsignstmtContext {
	var p = new(AsignstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_asignstmt
	return p
}

func InitEmptyAsignstmtContext(p *AsignstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_asignstmt
}

func (*AsignstmtContext) IsAsignstmtContext() {}

func NewAsignstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignstmtContext {
	var p = new(AsignstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_asignstmt

	return p
}

func (s *AsignstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *AsignstmtContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *AsignstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAsignstmt(s)
	}
}

func (s *AsignstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAsignstmt(s)
	}
}

func (s *AsignstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitAsignstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Asignstmt() (localctx IAsignstmtContext) {
	localctx = NewAsignstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SwiftGrammarParserRULE_asignstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Match(SwiftGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhilestmtContext is an interface to support dynamic dispatch.
type IWhilestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsWhilestmtContext differentiates from other interfaces.
	IsWhilestmtContext()
}

type WhilestmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhilestmtContext() *WhilestmtContext {
	var p = new(WhilestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_whilestmt
	return p
}

func InitEmptyWhilestmtContext(p *WhilestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_whilestmt
}

func (*WhilestmtContext) IsWhilestmtContext() {}

func NewWhilestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhilestmtContext {
	var p = new(WhilestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_whilestmt

	return p
}

func (s *WhilestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhilestmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserWHILE, 0)
}

func (s *WhilestmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhilestmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *WhilestmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhilestmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *WhilestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhilestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhilestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterWhilestmt(s)
	}
}

func (s *WhilestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitWhilestmt(s)
	}
}

func (s *WhilestmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitWhilestmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Whilestmt() (localctx IWhilestmtContext) {
	localctx = NewWhilestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SwiftGrammarParserRULE_whilestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Match(SwiftGrammarParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		p.expr(0)
	}
	{
		p.SetState(173)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
		p.Block()
	}
	{
		p.SetState(175)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOpasignstmtContext is an interface to support dynamic dispatch.
type IOpasignstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOpasignstmtContext differentiates from other interfaces.
	IsOpasignstmtContext()
}

type OpasignstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpasignstmtContext() *OpasignstmtContext {
	var p = new(OpasignstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_opasignstmt
	return p
}

func InitEmptyOpasignstmtContext(p *OpasignstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_opasignstmt
}

func (*OpasignstmtContext) IsOpasignstmtContext() {}

func NewOpasignstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpasignstmtContext {
	var p = new(OpasignstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_opasignstmt

	return p
}

func (s *OpasignstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *OpasignstmtContext) CopyAll(ctx *OpasignstmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *OpasignstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpasignstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DecrementoContext struct {
	OpasignstmtContext
}

func NewDecrementoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecrementoContext {
	var p = new(DecrementoContext)

	InitEmptyOpasignstmtContext(&p.OpasignstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*OpasignstmtContext))

	return p
}

func (s *DecrementoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecrementoContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *DecrementoContext) SUB() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSUB, 0)
}

func (s *DecrementoContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *DecrementoContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DecrementoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDecremento(s)
	}
}

func (s *DecrementoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDecremento(s)
	}
}

func (s *DecrementoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitDecremento(s)

	default:
		return t.VisitChildren(s)
	}
}

type IncrementoContext struct {
	OpasignstmtContext
}

func NewIncrementoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IncrementoContext {
	var p = new(IncrementoContext)

	InitEmptyOpasignstmtContext(&p.OpasignstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*OpasignstmtContext))

	return p
}

func (s *IncrementoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncrementoContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *IncrementoContext) ADD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserADD, 0)
}

func (s *IncrementoContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *IncrementoContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IncrementoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterIncremento(s)
	}
}

func (s *IncrementoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitIncremento(s)
	}
}

func (s *IncrementoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitIncremento(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Opasignstmt() (localctx IOpasignstmtContext) {
	localctx = NewOpasignstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SwiftGrammarParserRULE_opasignstmt)
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIncrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(177)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.Match(SwiftGrammarParserADD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.expr(0)
		}

	case 2:
		localctx = NewDecrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.Match(SwiftGrammarParserSUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UnariaExprContext struct {
	ExprContext
}

func NewUnariaExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnariaExprContext {
	var p = new(UnariaExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *UnariaExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnariaExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSUB, 0)
}

func (s *UnariaExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnariaExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterUnariaExpr(s)
	}
}

func (s *UnariaExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitUnariaExpr(s)
	}
}

func (s *UnariaExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitUnariaExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolExprContext struct {
	ExprContext
}

func NewBoolExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolExprContext {
	var p = new(BoolExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BoolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExprContext) TRU() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserTRU, 0)
}

func (s *BoolExprContext) FAL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFAL, 0)
}

func (s *BoolExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterBoolExpr(s)
	}
}

func (s *BoolExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitBoolExpr(s)
	}
}

func (s *BoolExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitBoolExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NilExprContext struct {
	ExprContext
}

func NewNilExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilExprContext {
	var p = new(NilExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NilExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilExprContext) NIL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNIL, 0)
}

func (s *NilExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterNilExpr(s)
	}
}

func (s *NilExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitNilExpr(s)
	}
}

func (s *NilExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitNilExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdExprContext struct {
	ExprContext
}

func NewIdExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExprContext {
	var p = new(IdExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExprContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *IdExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterIdExpr(s)
	}
}

func (s *IdExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitIdExpr(s)
	}
}

func (s *IdExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitIdExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumExprContext struct {
	ExprContext
}

func NewNumExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumExprContext {
	var p = new(NumExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NumExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNUMBER, 0)
}

func (s *NumExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterNumExpr(s)
	}
}

func (s *NumExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitNumExpr(s)
	}
}

func (s *NumExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitNumExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParExprContext struct {
	ExprContext
}

func NewParExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParExprContext {
	var p = new(ParExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParExprContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *ParExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParExprContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *ParExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterParExpr(s)
	}
}

func (s *ParExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitParExpr(s)
	}
}

func (s *ParExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitParExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type StrExprContext struct {
	ExprContext
}

func NewStrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StrExprContext {
	var p = new(StrExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTRING, 0)
}

func (s *StrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterStrExpr(s)
	}
}

func (s *StrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitStrExpr(s)
	}
}

func (s *StrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitStrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AccFuncExprContext struct {
	ExprContext
}

func NewAccFuncExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccFuncExprContext {
	var p = new(AccFuncExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AccFuncExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccFuncExprContext) Accfuncstm() IAccfuncstmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccfuncstmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccfuncstmContext)
}

func (s *AccFuncExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAccFuncExpr(s)
	}
}

func (s *AccFuncExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAccFuncExpr(s)
	}
}

func (s *AccFuncExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitAccFuncExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OpExprContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewOpExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpExprContext {
	var p = new(OpExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *OpExprContext) GetOp() antlr.Token { return s.op }

func (s *OpExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *OpExprContext) GetLeft() IExprContext { return s.left }

func (s *OpExprContext) GetRight() IExprContext { return s.right }

func (s *OpExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *OpExprContext) SetRight(v IExprContext) { s.right = v }

func (s *OpExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *OpExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OpExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMUL, 0)
}

func (s *OpExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDIV, 0)
}

func (s *OpExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMOD, 0)
}

func (s *OpExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserADD, 0)
}

func (s *OpExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSUB, 0)
}

func (s *OpExprContext) MAY_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMAY_IG, 0)
}

func (s *OpExprContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMAYOR, 0)
}

func (s *OpExprContext) MEN_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMEN_IG, 0)
}

func (s *OpExprContext) MENOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMENOR, 0)
}

func (s *OpExprContext) IG_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG_IG, 0)
}

func (s *OpExprContext) DIF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDIF, 0)
}

func (s *OpExprContext) AND() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAND, 0)
}

func (s *OpExprContext) OR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserOR, 0)
}

func (s *OpExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterOpExpr(s)
	}
}

func (s *OpExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitOpExpr(s)
	}
}

func (s *OpExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitOpExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *SwiftGrammarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 38
	p.EnterRecursionRule(localctx, 38, SwiftGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(188)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.expr(0)
		}
		{
			p.SetState(190)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewUnariaExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(192)
			p.Match(SwiftGrammarParserSUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.expr(14)
		}

	case 3:
		localctx = NewNumExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(194)
			p.Match(SwiftGrammarParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(195)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewStrExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(196)
			p.Match(SwiftGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(197)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SwiftGrammarParserTRU || _la == SwiftGrammarParserFAL) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 7:
		localctx = NewNilExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(198)
			p.Match(SwiftGrammarParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewAccFuncExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(199)
			p.Accfuncstm()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(223)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(202)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(203)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&163208757248) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(204)

					var _x = p.expr(14)

					localctx.(*OpExprContext).right = _x
				}

			case 2:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(205)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(206)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserADD || _la == SwiftGrammarParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(207)

					var _x = p.expr(13)

					localctx.(*OpExprContext).right = _x
				}

			case 3:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(208)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(209)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserMAY_IG || _la == SwiftGrammarParserMAYOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(210)

					var _x = p.expr(12)

					localctx.(*OpExprContext).right = _x
				}

			case 4:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(211)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(212)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserMEN_IG || _la == SwiftGrammarParserMENOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(213)

					var _x = p.expr(11)

					localctx.(*OpExprContext).right = _x
				}

			case 5:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(214)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(215)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserDIF || _la == SwiftGrammarParserIG_IG) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(216)

					var _x = p.expr(10)

					localctx.(*OpExprContext).right = _x
				}

			case 6:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(217)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(218)

					var _m = p.Match(SwiftGrammarParserAND)

					localctx.(*OpExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(219)

					var _x = p.expr(9)

					localctx.(*OpExprContext).right = _x
				}

			case 7:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(220)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(221)

					var _m = p.Match(SwiftGrammarParserOR)

					localctx.(*OpExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(222)

					var _x = p.expr(8)

					localctx.(*OpExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	CHARACTER() antlr.TerminalNode
	PSTRING() antlr.TerminalNode

	// IsTipoContext differentiates from other interfaces.
	IsTipoContext()
}

type TipoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipoContext() *TipoContext {
	var p = new(TipoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_tipo
	return p
}

func InitEmptyTipoContext(p *TipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_tipo
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) INT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserINT, 0)
}

func (s *TipoContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFLOAT, 0)
}

func (s *TipoContext) BOOL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserBOOL, 0)
}

func (s *TipoContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCHARACTER, 0)
}

func (s *TipoContext) PSTRING() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPSTRING, 0)
}

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterTipo(s)
	}
}

func (s *TipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitTipo(s)
	}
}

func (s *TipoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitTipo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Tipo() (localctx ITipoContext) {
	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SwiftGrammarParserRULE_tipo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *SwiftGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 19:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SwiftGrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

// Code generated from SwiftLexer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SwiftLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SwiftLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func swiftlexerLexerInit() {
	staticData := &SwiftLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'Int'", "'Float'", "'Bool'", "'Character'", "'String'", "'true'",
		"'false'", "'print'", "'if'", "'else'", "'while'", "'var'", "", "",
		"", "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", "'>='", "'<='", "'>'",
		"'<'", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'{'", "'}'", "':'",
		"'?'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "TRU", "FAL", "PRINT",
		"IF", "ELSE", "WHILE", "VAR", "NUMBER", "STRING", "ID", "DIF", "IG_IG",
		"NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL",
		"DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "DOSPUNTOS",
		"INTERROGACION", "WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "TRU", "FAL", "PRINT",
		"IF", "ELSE", "WHILE", "VAR", "NUMBER", "STRING", "ID", "DIF", "IG_IG",
		"NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL",
		"DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "DOSPUNTOS",
		"INTERROGACION", "WHITESPACE", "COMMENT", "LINE_COMMENT", "ESC_SEQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 38, 256, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 12, 4, 12, 148, 8, 12, 11, 12, 12, 12, 149, 1, 12, 1,
		12, 4, 12, 154, 8, 12, 11, 12, 12, 12, 155, 3, 12, 158, 8, 12, 1, 13, 1,
		13, 5, 13, 162, 8, 13, 10, 13, 12, 13, 165, 9, 13, 1, 13, 1, 13, 1, 14,
		1, 14, 5, 14, 171, 8, 14, 10, 14, 12, 14, 174, 9, 14, 1, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1,
		23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28,
		1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1,
		34, 1, 34, 1, 35, 4, 35, 223, 8, 35, 11, 35, 12, 35, 224, 1, 35, 1, 35,
		1, 36, 1, 36, 1, 36, 1, 36, 5, 36, 233, 8, 36, 10, 36, 12, 36, 236, 9,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 5, 37,
		247, 8, 37, 10, 37, 12, 37, 250, 9, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1,
		38, 1, 234, 0, 39, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35,
		18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53,
		27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71,
		36, 73, 37, 75, 38, 77, 0, 1, 0, 7, 1, 0, 48, 57, 1, 0, 34, 34, 3, 0, 65,
		90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 4, 0, 9, 10,
		13, 13, 32, 32, 92, 92, 2, 0, 10, 10, 13, 13, 7, 0, 32, 33, 35, 35, 43,
		43, 45, 46, 58, 58, 64, 64, 91, 93, 262, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0,
		35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0,
		0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0,
		0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0,
		0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1,
		0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73,
		1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 1, 79, 1, 0, 0, 0, 3, 83, 1, 0, 0, 0, 5,
		89, 1, 0, 0, 0, 7, 94, 1, 0, 0, 0, 9, 104, 1, 0, 0, 0, 11, 111, 1, 0, 0,
		0, 13, 116, 1, 0, 0, 0, 15, 122, 1, 0, 0, 0, 17, 128, 1, 0, 0, 0, 19, 131,
		1, 0, 0, 0, 21, 136, 1, 0, 0, 0, 23, 142, 1, 0, 0, 0, 25, 147, 1, 0, 0,
		0, 27, 159, 1, 0, 0, 0, 29, 168, 1, 0, 0, 0, 31, 175, 1, 0, 0, 0, 33, 178,
		1, 0, 0, 0, 35, 181, 1, 0, 0, 0, 37, 183, 1, 0, 0, 0, 39, 186, 1, 0, 0,
		0, 41, 189, 1, 0, 0, 0, 43, 191, 1, 0, 0, 0, 45, 194, 1, 0, 0, 0, 47, 197,
		1, 0, 0, 0, 49, 199, 1, 0, 0, 0, 51, 201, 1, 0, 0, 0, 53, 203, 1, 0, 0,
		0, 55, 205, 1, 0, 0, 0, 57, 207, 1, 0, 0, 0, 59, 209, 1, 0, 0, 0, 61, 211,
		1, 0, 0, 0, 63, 213, 1, 0, 0, 0, 65, 215, 1, 0, 0, 0, 67, 217, 1, 0, 0,
		0, 69, 219, 1, 0, 0, 0, 71, 222, 1, 0, 0, 0, 73, 228, 1, 0, 0, 0, 75, 242,
		1, 0, 0, 0, 77, 253, 1, 0, 0, 0, 79, 80, 5, 73, 0, 0, 80, 81, 5, 110, 0,
		0, 81, 82, 5, 116, 0, 0, 82, 2, 1, 0, 0, 0, 83, 84, 5, 70, 0, 0, 84, 85,
		5, 108, 0, 0, 85, 86, 5, 111, 0, 0, 86, 87, 5, 97, 0, 0, 87, 88, 5, 116,
		0, 0, 88, 4, 1, 0, 0, 0, 89, 90, 5, 66, 0, 0, 90, 91, 5, 111, 0, 0, 91,
		92, 5, 111, 0, 0, 92, 93, 5, 108, 0, 0, 93, 6, 1, 0, 0, 0, 94, 95, 5, 67,
		0, 0, 95, 96, 5, 104, 0, 0, 96, 97, 5, 97, 0, 0, 97, 98, 5, 114, 0, 0,
		98, 99, 5, 97, 0, 0, 99, 100, 5, 99, 0, 0, 100, 101, 5, 116, 0, 0, 101,
		102, 5, 101, 0, 0, 102, 103, 5, 114, 0, 0, 103, 8, 1, 0, 0, 0, 104, 105,
		5, 83, 0, 0, 105, 106, 5, 116, 0, 0, 106, 107, 5, 114, 0, 0, 107, 108,
		5, 105, 0, 0, 108, 109, 5, 110, 0, 0, 109, 110, 5, 103, 0, 0, 110, 10,
		1, 0, 0, 0, 111, 112, 5, 116, 0, 0, 112, 113, 5, 114, 0, 0, 113, 114, 5,
		117, 0, 0, 114, 115, 5, 101, 0, 0, 115, 12, 1, 0, 0, 0, 116, 117, 5, 102,
		0, 0, 117, 118, 5, 97, 0, 0, 118, 119, 5, 108, 0, 0, 119, 120, 5, 115,
		0, 0, 120, 121, 5, 101, 0, 0, 121, 14, 1, 0, 0, 0, 122, 123, 5, 112, 0,
		0, 123, 124, 5, 114, 0, 0, 124, 125, 5, 105, 0, 0, 125, 126, 5, 110, 0,
		0, 126, 127, 5, 116, 0, 0, 127, 16, 1, 0, 0, 0, 128, 129, 5, 105, 0, 0,
		129, 130, 5, 102, 0, 0, 130, 18, 1, 0, 0, 0, 131, 132, 5, 101, 0, 0, 132,
		133, 5, 108, 0, 0, 133, 134, 5, 115, 0, 0, 134, 135, 5, 101, 0, 0, 135,
		20, 1, 0, 0, 0, 136, 137, 5, 119, 0, 0, 137, 138, 5, 104, 0, 0, 138, 139,
		5, 105, 0, 0, 139, 140, 5, 108, 0, 0, 140, 141, 5, 101, 0, 0, 141, 22,
		1, 0, 0, 0, 142, 143, 5, 118, 0, 0, 143, 144, 5, 97, 0, 0, 144, 145, 5,
		114, 0, 0, 145, 24, 1, 0, 0, 0, 146, 148, 7, 0, 0, 0, 147, 146, 1, 0, 0,
		0, 148, 149, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150,
		157, 1, 0, 0, 0, 151, 153, 5, 46, 0, 0, 152, 154, 7, 0, 0, 0, 153, 152,
		1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 155, 156, 1, 0,
		0, 0, 156, 158, 1, 0, 0, 0, 157, 151, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0,
		158, 26, 1, 0, 0, 0, 159, 163, 5, 34, 0, 0, 160, 162, 8, 1, 0, 0, 161,
		160, 1, 0, 0, 0, 162, 165, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164,
		1, 0, 0, 0, 164, 166, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 166, 167, 5, 34,
		0, 0, 167, 28, 1, 0, 0, 0, 168, 172, 7, 2, 0, 0, 169, 171, 7, 3, 0, 0,
		170, 169, 1, 0, 0, 0, 171, 174, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172,
		173, 1, 0, 0, 0, 173, 30, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 175, 176, 5,
		33, 0, 0, 176, 177, 5, 61, 0, 0, 177, 32, 1, 0, 0, 0, 178, 179, 5, 61,
		0, 0, 179, 180, 5, 61, 0, 0, 180, 34, 1, 0, 0, 0, 181, 182, 5, 33, 0, 0,
		182, 36, 1, 0, 0, 0, 183, 184, 5, 124, 0, 0, 184, 185, 5, 124, 0, 0, 185,
		38, 1, 0, 0, 0, 186, 187, 5, 38, 0, 0, 187, 188, 5, 38, 0, 0, 188, 40,
		1, 0, 0, 0, 189, 190, 5, 61, 0, 0, 190, 42, 1, 0, 0, 0, 191, 192, 5, 62,
		0, 0, 192, 193, 5, 61, 0, 0, 193, 44, 1, 0, 0, 0, 194, 195, 5, 60, 0, 0,
		195, 196, 5, 61, 0, 0, 196, 46, 1, 0, 0, 0, 197, 198, 5, 62, 0, 0, 198,
		48, 1, 0, 0, 0, 199, 200, 5, 60, 0, 0, 200, 50, 1, 0, 0, 0, 201, 202, 5,
		42, 0, 0, 202, 52, 1, 0, 0, 0, 203, 204, 5, 47, 0, 0, 204, 54, 1, 0, 0,
		0, 205, 206, 5, 43, 0, 0, 206, 56, 1, 0, 0, 0, 207, 208, 5, 45, 0, 0, 208,
		58, 1, 0, 0, 0, 209, 210, 5, 40, 0, 0, 210, 60, 1, 0, 0, 0, 211, 212, 5,
		41, 0, 0, 212, 62, 1, 0, 0, 0, 213, 214, 5, 123, 0, 0, 214, 64, 1, 0, 0,
		0, 215, 216, 5, 125, 0, 0, 216, 66, 1, 0, 0, 0, 217, 218, 5, 58, 0, 0,
		218, 68, 1, 0, 0, 0, 219, 220, 5, 63, 0, 0, 220, 70, 1, 0, 0, 0, 221, 223,
		7, 4, 0, 0, 222, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 222, 1, 0,
		0, 0, 224, 225, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 227, 6, 35, 0, 0,
		227, 72, 1, 0, 0, 0, 228, 229, 5, 47, 0, 0, 229, 230, 5, 42, 0, 0, 230,
		234, 1, 0, 0, 0, 231, 233, 9, 0, 0, 0, 232, 231, 1, 0, 0, 0, 233, 236,
		1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 235, 237, 1, 0,
		0, 0, 236, 234, 1, 0, 0, 0, 237, 238, 5, 42, 0, 0, 238, 239, 5, 47, 0,
		0, 239, 240, 1, 0, 0, 0, 240, 241, 6, 36, 0, 0, 241, 74, 1, 0, 0, 0, 242,
		243, 5, 47, 0, 0, 243, 244, 5, 47, 0, 0, 244, 248, 1, 0, 0, 0, 245, 247,
		8, 5, 0, 0, 246, 245, 1, 0, 0, 0, 247, 250, 1, 0, 0, 0, 248, 246, 1, 0,
		0, 0, 248, 249, 1, 0, 0, 0, 249, 251, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0,
		251, 252, 6, 37, 0, 0, 252, 76, 1, 0, 0, 0, 253, 254, 5, 92, 0, 0, 254,
		255, 7, 6, 0, 0, 255, 78, 1, 0, 0, 0, 9, 0, 149, 155, 157, 163, 172, 224,
		234, 248, 1, 6, 0, 0,
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

// SwiftLexerInit initializes any static state used to implement SwiftLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSwiftLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftLexerInit() {
	staticData := &SwiftLexerLexerStaticData
	staticData.once.Do(swiftlexerLexerInit)
}

// NewSwiftLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSwiftLexer(input antlr.CharStream) *SwiftLexer {
	SwiftLexerInit()
	l := new(SwiftLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SwiftLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "SwiftLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SwiftLexer tokens.
const (
	SwiftLexerINT           = 1
	SwiftLexerFLOAT         = 2
	SwiftLexerBOOL          = 3
	SwiftLexerCHARACTER     = 4
	SwiftLexerPSTRING       = 5
	SwiftLexerTRU           = 6
	SwiftLexerFAL           = 7
	SwiftLexerPRINT         = 8
	SwiftLexerIF            = 9
	SwiftLexerELSE          = 10
	SwiftLexerWHILE         = 11
	SwiftLexerVAR           = 12
	SwiftLexerNUMBER        = 13
	SwiftLexerSTRING        = 14
	SwiftLexerID            = 15
	SwiftLexerDIF           = 16
	SwiftLexerIG_IG         = 17
	SwiftLexerNOT           = 18
	SwiftLexerOR            = 19
	SwiftLexerAND           = 20
	SwiftLexerIG            = 21
	SwiftLexerMAY_IG        = 22
	SwiftLexerMEN_IG        = 23
	SwiftLexerMAYOR         = 24
	SwiftLexerMENOR         = 25
	SwiftLexerMUL           = 26
	SwiftLexerDIV           = 27
	SwiftLexerADD           = 28
	SwiftLexerSUB           = 29
	SwiftLexerPARIZQ        = 30
	SwiftLexerPARDER        = 31
	SwiftLexerLLAVEIZQ      = 32
	SwiftLexerLLAVEDER      = 33
	SwiftLexerDOSPUNTOS     = 34
	SwiftLexerINTERROGACION = 35
	SwiftLexerWHITESPACE    = 36
	SwiftLexerCOMMENT       = 37
	SwiftLexerLINE_COMMENT  = 38
)

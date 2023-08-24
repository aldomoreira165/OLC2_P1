// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

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

type SwiftGrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SwiftGrammarLexerLexerStaticData struct {
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

func swiftgrammarlexerLexerInit() {
	staticData := &SwiftGrammarLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "NIL", "TRU", "FAL",
		"PRINT", "IF", "ELSE", "WHILE", "SWITCH", "CASE", "DEFAULT", "VAR",
		"BREAK", "RETURN", "FUNC", "NUMBER", "STRING", "ID", "DIF", "IG_IG",
		"NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL",
		"DIV", "ADD", "SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER",
		"DOSPUNTOS", "PTCOMA", "INTERROGACION", "WHITESPACE", "COMMENT", "LINE_COMMENT",
		"ESC_SEQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 47, 320, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19,
		4, 19, 208, 8, 19, 11, 19, 12, 19, 209, 1, 19, 1, 19, 4, 19, 214, 8, 19,
		11, 19, 12, 19, 215, 3, 19, 218, 8, 19, 1, 20, 1, 20, 5, 20, 222, 8, 20,
		10, 20, 12, 20, 225, 9, 20, 1, 20, 1, 20, 1, 21, 1, 21, 5, 21, 231, 8,
		21, 10, 21, 12, 21, 234, 9, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1,
		28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32,
		1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1,
		37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42,
		1, 43, 1, 43, 1, 44, 4, 44, 287, 8, 44, 11, 44, 12, 44, 288, 1, 44, 1,
		44, 1, 45, 1, 45, 1, 45, 1, 45, 5, 45, 297, 8, 45, 10, 45, 12, 45, 300,
		9, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 5,
		46, 311, 8, 46, 10, 46, 12, 46, 314, 9, 46, 1, 46, 1, 46, 1, 47, 1, 47,
		1, 47, 1, 298, 0, 48, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8,
		17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35,
		71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44,
		89, 45, 91, 46, 93, 47, 95, 0, 1, 0, 7, 1, 0, 48, 57, 1, 0, 34, 34, 3,
		0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 4, 0,
		9, 10, 13, 13, 32, 32, 92, 92, 2, 0, 10, 10, 13, 13, 7, 0, 32, 33, 35,
		35, 43, 43, 45, 46, 58, 58, 64, 64, 91, 93, 326, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0,
		0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1,
		0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57,
		1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0,
		65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0,
		0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0,
		0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0,
		0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 1, 97, 1,
		0, 0, 0, 3, 101, 1, 0, 0, 0, 5, 107, 1, 0, 0, 0, 7, 112, 1, 0, 0, 0, 9,
		122, 1, 0, 0, 0, 11, 129, 1, 0, 0, 0, 13, 133, 1, 0, 0, 0, 15, 138, 1,
		0, 0, 0, 17, 144, 1, 0, 0, 0, 19, 150, 1, 0, 0, 0, 21, 153, 1, 0, 0, 0,
		23, 158, 1, 0, 0, 0, 25, 164, 1, 0, 0, 0, 27, 171, 1, 0, 0, 0, 29, 176,
		1, 0, 0, 0, 31, 184, 1, 0, 0, 0, 33, 188, 1, 0, 0, 0, 35, 194, 1, 0, 0,
		0, 37, 201, 1, 0, 0, 0, 39, 207, 1, 0, 0, 0, 41, 219, 1, 0, 0, 0, 43, 228,
		1, 0, 0, 0, 45, 235, 1, 0, 0, 0, 47, 238, 1, 0, 0, 0, 49, 241, 1, 0, 0,
		0, 51, 243, 1, 0, 0, 0, 53, 246, 1, 0, 0, 0, 55, 249, 1, 0, 0, 0, 57, 251,
		1, 0, 0, 0, 59, 254, 1, 0, 0, 0, 61, 257, 1, 0, 0, 0, 63, 259, 1, 0, 0,
		0, 65, 261, 1, 0, 0, 0, 67, 263, 1, 0, 0, 0, 69, 265, 1, 0, 0, 0, 71, 267,
		1, 0, 0, 0, 73, 269, 1, 0, 0, 0, 75, 271, 1, 0, 0, 0, 77, 273, 1, 0, 0,
		0, 79, 275, 1, 0, 0, 0, 81, 277, 1, 0, 0, 0, 83, 279, 1, 0, 0, 0, 85, 281,
		1, 0, 0, 0, 87, 283, 1, 0, 0, 0, 89, 286, 1, 0, 0, 0, 91, 292, 1, 0, 0,
		0, 93, 306, 1, 0, 0, 0, 95, 317, 1, 0, 0, 0, 97, 98, 5, 105, 0, 0, 98,
		99, 5, 110, 0, 0, 99, 100, 5, 116, 0, 0, 100, 2, 1, 0, 0, 0, 101, 102,
		5, 102, 0, 0, 102, 103, 5, 108, 0, 0, 103, 104, 5, 111, 0, 0, 104, 105,
		5, 97, 0, 0, 105, 106, 5, 116, 0, 0, 106, 4, 1, 0, 0, 0, 107, 108, 5, 98,
		0, 0, 108, 109, 5, 111, 0, 0, 109, 110, 5, 111, 0, 0, 110, 111, 5, 108,
		0, 0, 111, 6, 1, 0, 0, 0, 112, 113, 5, 99, 0, 0, 113, 114, 5, 104, 0, 0,
		114, 115, 5, 97, 0, 0, 115, 116, 5, 114, 0, 0, 116, 117, 5, 97, 0, 0, 117,
		118, 5, 99, 0, 0, 118, 119, 5, 116, 0, 0, 119, 120, 5, 101, 0, 0, 120,
		121, 5, 114, 0, 0, 121, 8, 1, 0, 0, 0, 122, 123, 5, 83, 0, 0, 123, 124,
		5, 116, 0, 0, 124, 125, 5, 114, 0, 0, 125, 126, 5, 105, 0, 0, 126, 127,
		5, 110, 0, 0, 127, 128, 5, 103, 0, 0, 128, 10, 1, 0, 0, 0, 129, 130, 5,
		110, 0, 0, 130, 131, 5, 105, 0, 0, 131, 132, 5, 108, 0, 0, 132, 12, 1,
		0, 0, 0, 133, 134, 5, 116, 0, 0, 134, 135, 5, 114, 0, 0, 135, 136, 5, 117,
		0, 0, 136, 137, 5, 101, 0, 0, 137, 14, 1, 0, 0, 0, 138, 139, 5, 102, 0,
		0, 139, 140, 5, 97, 0, 0, 140, 141, 5, 108, 0, 0, 141, 142, 5, 115, 0,
		0, 142, 143, 5, 101, 0, 0, 143, 16, 1, 0, 0, 0, 144, 145, 5, 112, 0, 0,
		145, 146, 5, 114, 0, 0, 146, 147, 5, 105, 0, 0, 147, 148, 5, 110, 0, 0,
		148, 149, 5, 116, 0, 0, 149, 18, 1, 0, 0, 0, 150, 151, 5, 105, 0, 0, 151,
		152, 5, 102, 0, 0, 152, 20, 1, 0, 0, 0, 153, 154, 5, 101, 0, 0, 154, 155,
		5, 108, 0, 0, 155, 156, 5, 115, 0, 0, 156, 157, 5, 101, 0, 0, 157, 22,
		1, 0, 0, 0, 158, 159, 5, 119, 0, 0, 159, 160, 5, 104, 0, 0, 160, 161, 5,
		105, 0, 0, 161, 162, 5, 108, 0, 0, 162, 163, 5, 101, 0, 0, 163, 24, 1,
		0, 0, 0, 164, 165, 5, 115, 0, 0, 165, 166, 5, 119, 0, 0, 166, 167, 5, 105,
		0, 0, 167, 168, 5, 116, 0, 0, 168, 169, 5, 99, 0, 0, 169, 170, 5, 104,
		0, 0, 170, 26, 1, 0, 0, 0, 171, 172, 5, 99, 0, 0, 172, 173, 5, 97, 0, 0,
		173, 174, 5, 115, 0, 0, 174, 175, 5, 101, 0, 0, 175, 28, 1, 0, 0, 0, 176,
		177, 5, 100, 0, 0, 177, 178, 5, 101, 0, 0, 178, 179, 5, 102, 0, 0, 179,
		180, 5, 97, 0, 0, 180, 181, 5, 117, 0, 0, 181, 182, 5, 108, 0, 0, 182,
		183, 5, 116, 0, 0, 183, 30, 1, 0, 0, 0, 184, 185, 5, 118, 0, 0, 185, 186,
		5, 97, 0, 0, 186, 187, 5, 114, 0, 0, 187, 32, 1, 0, 0, 0, 188, 189, 5,
		98, 0, 0, 189, 190, 5, 114, 0, 0, 190, 191, 5, 101, 0, 0, 191, 192, 5,
		97, 0, 0, 192, 193, 5, 107, 0, 0, 193, 34, 1, 0, 0, 0, 194, 195, 5, 114,
		0, 0, 195, 196, 5, 101, 0, 0, 196, 197, 5, 116, 0, 0, 197, 198, 5, 117,
		0, 0, 198, 199, 5, 114, 0, 0, 199, 200, 5, 110, 0, 0, 200, 36, 1, 0, 0,
		0, 201, 202, 5, 102, 0, 0, 202, 203, 5, 117, 0, 0, 203, 204, 5, 110, 0,
		0, 204, 205, 5, 99, 0, 0, 205, 38, 1, 0, 0, 0, 206, 208, 7, 0, 0, 0, 207,
		206, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 209, 210,
		1, 0, 0, 0, 210, 217, 1, 0, 0, 0, 211, 213, 5, 46, 0, 0, 212, 214, 7, 0,
		0, 0, 213, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0,
		215, 216, 1, 0, 0, 0, 216, 218, 1, 0, 0, 0, 217, 211, 1, 0, 0, 0, 217,
		218, 1, 0, 0, 0, 218, 40, 1, 0, 0, 0, 219, 223, 5, 34, 0, 0, 220, 222,
		8, 1, 0, 0, 221, 220, 1, 0, 0, 0, 222, 225, 1, 0, 0, 0, 223, 221, 1, 0,
		0, 0, 223, 224, 1, 0, 0, 0, 224, 226, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0,
		226, 227, 5, 34, 0, 0, 227, 42, 1, 0, 0, 0, 228, 232, 7, 2, 0, 0, 229,
		231, 7, 3, 0, 0, 230, 229, 1, 0, 0, 0, 231, 234, 1, 0, 0, 0, 232, 230,
		1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 44, 1, 0, 0, 0, 234, 232, 1, 0,
		0, 0, 235, 236, 5, 33, 0, 0, 236, 237, 5, 61, 0, 0, 237, 46, 1, 0, 0, 0,
		238, 239, 5, 61, 0, 0, 239, 240, 5, 61, 0, 0, 240, 48, 1, 0, 0, 0, 241,
		242, 5, 33, 0, 0, 242, 50, 1, 0, 0, 0, 243, 244, 5, 124, 0, 0, 244, 245,
		5, 124, 0, 0, 245, 52, 1, 0, 0, 0, 246, 247, 5, 38, 0, 0, 247, 248, 5,
		38, 0, 0, 248, 54, 1, 0, 0, 0, 249, 250, 5, 61, 0, 0, 250, 56, 1, 0, 0,
		0, 251, 252, 5, 62, 0, 0, 252, 253, 5, 61, 0, 0, 253, 58, 1, 0, 0, 0, 254,
		255, 5, 60, 0, 0, 255, 256, 5, 61, 0, 0, 256, 60, 1, 0, 0, 0, 257, 258,
		5, 62, 0, 0, 258, 62, 1, 0, 0, 0, 259, 260, 5, 60, 0, 0, 260, 64, 1, 0,
		0, 0, 261, 262, 5, 42, 0, 0, 262, 66, 1, 0, 0, 0, 263, 264, 5, 47, 0, 0,
		264, 68, 1, 0, 0, 0, 265, 266, 5, 43, 0, 0, 266, 70, 1, 0, 0, 0, 267, 268,
		5, 45, 0, 0, 268, 72, 1, 0, 0, 0, 269, 270, 5, 37, 0, 0, 270, 74, 1, 0,
		0, 0, 271, 272, 5, 40, 0, 0, 272, 76, 1, 0, 0, 0, 273, 274, 5, 41, 0, 0,
		274, 78, 1, 0, 0, 0, 275, 276, 5, 123, 0, 0, 276, 80, 1, 0, 0, 0, 277,
		278, 5, 125, 0, 0, 278, 82, 1, 0, 0, 0, 279, 280, 5, 58, 0, 0, 280, 84,
		1, 0, 0, 0, 281, 282, 5, 59, 0, 0, 282, 86, 1, 0, 0, 0, 283, 284, 5, 63,
		0, 0, 284, 88, 1, 0, 0, 0, 285, 287, 7, 4, 0, 0, 286, 285, 1, 0, 0, 0,
		287, 288, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289,
		290, 1, 0, 0, 0, 290, 291, 6, 44, 0, 0, 291, 90, 1, 0, 0, 0, 292, 293,
		5, 47, 0, 0, 293, 294, 5, 42, 0, 0, 294, 298, 1, 0, 0, 0, 295, 297, 9,
		0, 0, 0, 296, 295, 1, 0, 0, 0, 297, 300, 1, 0, 0, 0, 298, 299, 1, 0, 0,
		0, 298, 296, 1, 0, 0, 0, 299, 301, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0, 301,
		302, 5, 42, 0, 0, 302, 303, 5, 47, 0, 0, 303, 304, 1, 0, 0, 0, 304, 305,
		6, 45, 0, 0, 305, 92, 1, 0, 0, 0, 306, 307, 5, 47, 0, 0, 307, 308, 5, 47,
		0, 0, 308, 312, 1, 0, 0, 0, 309, 311, 8, 5, 0, 0, 310, 309, 1, 0, 0, 0,
		311, 314, 1, 0, 0, 0, 312, 310, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313,
		315, 1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 315, 316, 6, 46, 0, 0, 316, 94,
		1, 0, 0, 0, 317, 318, 5, 92, 0, 0, 318, 319, 7, 6, 0, 0, 319, 96, 1, 0,
		0, 0, 9, 0, 209, 215, 217, 223, 232, 288, 298, 312, 1, 6, 0, 0,
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

// SwiftGrammarLexerInit initializes any static state used to implement SwiftGrammarLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSwiftGrammarLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftGrammarLexerInit() {
	staticData := &SwiftGrammarLexerLexerStaticData
	staticData.once.Do(swiftgrammarlexerLexerInit)
}

// NewSwiftGrammarLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSwiftGrammarLexer(input antlr.CharStream) *SwiftGrammarLexer {
	SwiftGrammarLexerInit()
	l := new(SwiftGrammarLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SwiftGrammarLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "SwiftGrammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SwiftGrammarLexer tokens.
const (
	SwiftGrammarLexerINT           = 1
	SwiftGrammarLexerFLOAT         = 2
	SwiftGrammarLexerBOOL          = 3
	SwiftGrammarLexerCHARACTER     = 4
	SwiftGrammarLexerPSTRING       = 5
	SwiftGrammarLexerNIL           = 6
	SwiftGrammarLexerTRU           = 7
	SwiftGrammarLexerFAL           = 8
	SwiftGrammarLexerPRINT         = 9
	SwiftGrammarLexerIF            = 10
	SwiftGrammarLexerELSE          = 11
	SwiftGrammarLexerWHILE         = 12
	SwiftGrammarLexerSWITCH        = 13
	SwiftGrammarLexerCASE          = 14
	SwiftGrammarLexerDEFAULT       = 15
	SwiftGrammarLexerVAR           = 16
	SwiftGrammarLexerBREAK         = 17
	SwiftGrammarLexerRETURN        = 18
	SwiftGrammarLexerFUNC          = 19
	SwiftGrammarLexerNUMBER        = 20
	SwiftGrammarLexerSTRING        = 21
	SwiftGrammarLexerID            = 22
	SwiftGrammarLexerDIF           = 23
	SwiftGrammarLexerIG_IG         = 24
	SwiftGrammarLexerNOT           = 25
	SwiftGrammarLexerOR            = 26
	SwiftGrammarLexerAND           = 27
	SwiftGrammarLexerIG            = 28
	SwiftGrammarLexerMAY_IG        = 29
	SwiftGrammarLexerMEN_IG        = 30
	SwiftGrammarLexerMAYOR         = 31
	SwiftGrammarLexerMENOR         = 32
	SwiftGrammarLexerMUL           = 33
	SwiftGrammarLexerDIV           = 34
	SwiftGrammarLexerADD           = 35
	SwiftGrammarLexerSUB           = 36
	SwiftGrammarLexerMOD           = 37
	SwiftGrammarLexerPARIZQ        = 38
	SwiftGrammarLexerPARDER        = 39
	SwiftGrammarLexerLLAVEIZQ      = 40
	SwiftGrammarLexerLLAVEDER      = 41
	SwiftGrammarLexerDOSPUNTOS     = 42
	SwiftGrammarLexerPTCOMA        = 43
	SwiftGrammarLexerINTERROGACION = 44
	SwiftGrammarLexerWHITESPACE    = 45
	SwiftGrammarLexerCOMMENT       = 46
	SwiftGrammarLexerLINE_COMMENT  = 47
)

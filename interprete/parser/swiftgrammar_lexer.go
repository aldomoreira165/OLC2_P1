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
		"'case'", "'default'", "'var'", "'break'", "", "", "", "'!='", "'=='",
		"'!'", "'||'", "'&&'", "'='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'",
		"'+'", "'-'", "'%'", "'('", "')'", "'{'", "'}'", "':'", "'?'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "NIL", "TRU", "FAL",
		"PRINT", "IF", "ELSE", "WHILE", "SWITCH", "CASE", "DEFAULT", "VAR",
		"BREAK", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND",
		"IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB",
		"MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "DOSPUNTOS", "INTERROGACION",
		"WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "NIL", "TRU", "FAL",
		"PRINT", "IF", "ELSE", "WHILE", "SWITCH", "CASE", "DEFAULT", "VAR",
		"BREAK", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND",
		"IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB",
		"MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "DOSPUNTOS", "INTERROGACION",
		"WHITESPACE", "COMMENT", "LINE_COMMENT", "ESC_SEQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 44, 300, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 17, 4, 17, 190, 8, 17, 11, 17, 12, 17, 191, 1, 17, 1,
		17, 4, 17, 196, 8, 17, 11, 17, 12, 17, 197, 3, 17, 200, 8, 17, 1, 18, 1,
		18, 5, 18, 204, 8, 18, 10, 18, 12, 18, 207, 9, 18, 1, 18, 1, 18, 1, 19,
		1, 19, 5, 19, 213, 8, 19, 10, 19, 12, 19, 216, 9, 19, 1, 20, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1,
		28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33,
		1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1,
		39, 1, 39, 1, 40, 1, 40, 1, 41, 4, 41, 267, 8, 41, 11, 41, 12, 41, 268,
		1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 5, 42, 277, 8, 42, 10, 42, 12,
		42, 280, 9, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43,
		1, 43, 5, 43, 291, 8, 43, 10, 43, 12, 43, 294, 9, 43, 1, 43, 1, 43, 1,
		44, 1, 44, 1, 44, 1, 278, 0, 45, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25,
		51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34,
		69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43,
		87, 44, 89, 0, 1, 0, 7, 1, 0, 48, 57, 1, 0, 34, 34, 3, 0, 65, 90, 95, 95,
		97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 4, 0, 9, 10, 13, 13, 32,
		32, 92, 92, 2, 0, 10, 10, 13, 13, 7, 0, 32, 33, 35, 35, 43, 43, 45, 46,
		58, 58, 64, 64, 91, 93, 306, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1,
		0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59,
		1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0,
		67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0,
		0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0,
		0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 1, 91, 1, 0,
		0, 0, 3, 95, 1, 0, 0, 0, 5, 101, 1, 0, 0, 0, 7, 106, 1, 0, 0, 0, 9, 116,
		1, 0, 0, 0, 11, 123, 1, 0, 0, 0, 13, 127, 1, 0, 0, 0, 15, 132, 1, 0, 0,
		0, 17, 138, 1, 0, 0, 0, 19, 144, 1, 0, 0, 0, 21, 147, 1, 0, 0, 0, 23, 152,
		1, 0, 0, 0, 25, 158, 1, 0, 0, 0, 27, 165, 1, 0, 0, 0, 29, 170, 1, 0, 0,
		0, 31, 178, 1, 0, 0, 0, 33, 182, 1, 0, 0, 0, 35, 189, 1, 0, 0, 0, 37, 201,
		1, 0, 0, 0, 39, 210, 1, 0, 0, 0, 41, 217, 1, 0, 0, 0, 43, 220, 1, 0, 0,
		0, 45, 223, 1, 0, 0, 0, 47, 225, 1, 0, 0, 0, 49, 228, 1, 0, 0, 0, 51, 231,
		1, 0, 0, 0, 53, 233, 1, 0, 0, 0, 55, 236, 1, 0, 0, 0, 57, 239, 1, 0, 0,
		0, 59, 241, 1, 0, 0, 0, 61, 243, 1, 0, 0, 0, 63, 245, 1, 0, 0, 0, 65, 247,
		1, 0, 0, 0, 67, 249, 1, 0, 0, 0, 69, 251, 1, 0, 0, 0, 71, 253, 1, 0, 0,
		0, 73, 255, 1, 0, 0, 0, 75, 257, 1, 0, 0, 0, 77, 259, 1, 0, 0, 0, 79, 261,
		1, 0, 0, 0, 81, 263, 1, 0, 0, 0, 83, 266, 1, 0, 0, 0, 85, 272, 1, 0, 0,
		0, 87, 286, 1, 0, 0, 0, 89, 297, 1, 0, 0, 0, 91, 92, 5, 105, 0, 0, 92,
		93, 5, 110, 0, 0, 93, 94, 5, 116, 0, 0, 94, 2, 1, 0, 0, 0, 95, 96, 5, 102,
		0, 0, 96, 97, 5, 108, 0, 0, 97, 98, 5, 111, 0, 0, 98, 99, 5, 97, 0, 0,
		99, 100, 5, 116, 0, 0, 100, 4, 1, 0, 0, 0, 101, 102, 5, 98, 0, 0, 102,
		103, 5, 111, 0, 0, 103, 104, 5, 111, 0, 0, 104, 105, 5, 108, 0, 0, 105,
		6, 1, 0, 0, 0, 106, 107, 5, 99, 0, 0, 107, 108, 5, 104, 0, 0, 108, 109,
		5, 97, 0, 0, 109, 110, 5, 114, 0, 0, 110, 111, 5, 97, 0, 0, 111, 112, 5,
		99, 0, 0, 112, 113, 5, 116, 0, 0, 113, 114, 5, 101, 0, 0, 114, 115, 5,
		114, 0, 0, 115, 8, 1, 0, 0, 0, 116, 117, 5, 83, 0, 0, 117, 118, 5, 116,
		0, 0, 118, 119, 5, 114, 0, 0, 119, 120, 5, 105, 0, 0, 120, 121, 5, 110,
		0, 0, 121, 122, 5, 103, 0, 0, 122, 10, 1, 0, 0, 0, 123, 124, 5, 110, 0,
		0, 124, 125, 5, 105, 0, 0, 125, 126, 5, 108, 0, 0, 126, 12, 1, 0, 0, 0,
		127, 128, 5, 116, 0, 0, 128, 129, 5, 114, 0, 0, 129, 130, 5, 117, 0, 0,
		130, 131, 5, 101, 0, 0, 131, 14, 1, 0, 0, 0, 132, 133, 5, 102, 0, 0, 133,
		134, 5, 97, 0, 0, 134, 135, 5, 108, 0, 0, 135, 136, 5, 115, 0, 0, 136,
		137, 5, 101, 0, 0, 137, 16, 1, 0, 0, 0, 138, 139, 5, 112, 0, 0, 139, 140,
		5, 114, 0, 0, 140, 141, 5, 105, 0, 0, 141, 142, 5, 110, 0, 0, 142, 143,
		5, 116, 0, 0, 143, 18, 1, 0, 0, 0, 144, 145, 5, 105, 0, 0, 145, 146, 5,
		102, 0, 0, 146, 20, 1, 0, 0, 0, 147, 148, 5, 101, 0, 0, 148, 149, 5, 108,
		0, 0, 149, 150, 5, 115, 0, 0, 150, 151, 5, 101, 0, 0, 151, 22, 1, 0, 0,
		0, 152, 153, 5, 119, 0, 0, 153, 154, 5, 104, 0, 0, 154, 155, 5, 105, 0,
		0, 155, 156, 5, 108, 0, 0, 156, 157, 5, 101, 0, 0, 157, 24, 1, 0, 0, 0,
		158, 159, 5, 115, 0, 0, 159, 160, 5, 119, 0, 0, 160, 161, 5, 105, 0, 0,
		161, 162, 5, 116, 0, 0, 162, 163, 5, 99, 0, 0, 163, 164, 5, 104, 0, 0,
		164, 26, 1, 0, 0, 0, 165, 166, 5, 99, 0, 0, 166, 167, 5, 97, 0, 0, 167,
		168, 5, 115, 0, 0, 168, 169, 5, 101, 0, 0, 169, 28, 1, 0, 0, 0, 170, 171,
		5, 100, 0, 0, 171, 172, 5, 101, 0, 0, 172, 173, 5, 102, 0, 0, 173, 174,
		5, 97, 0, 0, 174, 175, 5, 117, 0, 0, 175, 176, 5, 108, 0, 0, 176, 177,
		5, 116, 0, 0, 177, 30, 1, 0, 0, 0, 178, 179, 5, 118, 0, 0, 179, 180, 5,
		97, 0, 0, 180, 181, 5, 114, 0, 0, 181, 32, 1, 0, 0, 0, 182, 183, 5, 98,
		0, 0, 183, 184, 5, 114, 0, 0, 184, 185, 5, 101, 0, 0, 185, 186, 5, 97,
		0, 0, 186, 187, 5, 107, 0, 0, 187, 34, 1, 0, 0, 0, 188, 190, 7, 0, 0, 0,
		189, 188, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 191,
		192, 1, 0, 0, 0, 192, 199, 1, 0, 0, 0, 193, 195, 5, 46, 0, 0, 194, 196,
		7, 0, 0, 0, 195, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 195, 1, 0,
		0, 0, 197, 198, 1, 0, 0, 0, 198, 200, 1, 0, 0, 0, 199, 193, 1, 0, 0, 0,
		199, 200, 1, 0, 0, 0, 200, 36, 1, 0, 0, 0, 201, 205, 5, 34, 0, 0, 202,
		204, 8, 1, 0, 0, 203, 202, 1, 0, 0, 0, 204, 207, 1, 0, 0, 0, 205, 203,
		1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 208, 1, 0, 0, 0, 207, 205, 1, 0,
		0, 0, 208, 209, 5, 34, 0, 0, 209, 38, 1, 0, 0, 0, 210, 214, 7, 2, 0, 0,
		211, 213, 7, 3, 0, 0, 212, 211, 1, 0, 0, 0, 213, 216, 1, 0, 0, 0, 214,
		212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 40, 1, 0, 0, 0, 216, 214, 1,
		0, 0, 0, 217, 218, 5, 33, 0, 0, 218, 219, 5, 61, 0, 0, 219, 42, 1, 0, 0,
		0, 220, 221, 5, 61, 0, 0, 221, 222, 5, 61, 0, 0, 222, 44, 1, 0, 0, 0, 223,
		224, 5, 33, 0, 0, 224, 46, 1, 0, 0, 0, 225, 226, 5, 124, 0, 0, 226, 227,
		5, 124, 0, 0, 227, 48, 1, 0, 0, 0, 228, 229, 5, 38, 0, 0, 229, 230, 5,
		38, 0, 0, 230, 50, 1, 0, 0, 0, 231, 232, 5, 61, 0, 0, 232, 52, 1, 0, 0,
		0, 233, 234, 5, 62, 0, 0, 234, 235, 5, 61, 0, 0, 235, 54, 1, 0, 0, 0, 236,
		237, 5, 60, 0, 0, 237, 238, 5, 61, 0, 0, 238, 56, 1, 0, 0, 0, 239, 240,
		5, 62, 0, 0, 240, 58, 1, 0, 0, 0, 241, 242, 5, 60, 0, 0, 242, 60, 1, 0,
		0, 0, 243, 244, 5, 42, 0, 0, 244, 62, 1, 0, 0, 0, 245, 246, 5, 47, 0, 0,
		246, 64, 1, 0, 0, 0, 247, 248, 5, 43, 0, 0, 248, 66, 1, 0, 0, 0, 249, 250,
		5, 45, 0, 0, 250, 68, 1, 0, 0, 0, 251, 252, 5, 37, 0, 0, 252, 70, 1, 0,
		0, 0, 253, 254, 5, 40, 0, 0, 254, 72, 1, 0, 0, 0, 255, 256, 5, 41, 0, 0,
		256, 74, 1, 0, 0, 0, 257, 258, 5, 123, 0, 0, 258, 76, 1, 0, 0, 0, 259,
		260, 5, 125, 0, 0, 260, 78, 1, 0, 0, 0, 261, 262, 5, 58, 0, 0, 262, 80,
		1, 0, 0, 0, 263, 264, 5, 63, 0, 0, 264, 82, 1, 0, 0, 0, 265, 267, 7, 4,
		0, 0, 266, 265, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0,
		268, 269, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 271, 6, 41, 0, 0, 271,
		84, 1, 0, 0, 0, 272, 273, 5, 47, 0, 0, 273, 274, 5, 42, 0, 0, 274, 278,
		1, 0, 0, 0, 275, 277, 9, 0, 0, 0, 276, 275, 1, 0, 0, 0, 277, 280, 1, 0,
		0, 0, 278, 279, 1, 0, 0, 0, 278, 276, 1, 0, 0, 0, 279, 281, 1, 0, 0, 0,
		280, 278, 1, 0, 0, 0, 281, 282, 5, 42, 0, 0, 282, 283, 5, 47, 0, 0, 283,
		284, 1, 0, 0, 0, 284, 285, 6, 42, 0, 0, 285, 86, 1, 0, 0, 0, 286, 287,
		5, 47, 0, 0, 287, 288, 5, 47, 0, 0, 288, 292, 1, 0, 0, 0, 289, 291, 8,
		5, 0, 0, 290, 289, 1, 0, 0, 0, 291, 294, 1, 0, 0, 0, 292, 290, 1, 0, 0,
		0, 292, 293, 1, 0, 0, 0, 293, 295, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 295,
		296, 6, 43, 0, 0, 296, 88, 1, 0, 0, 0, 297, 298, 5, 92, 0, 0, 298, 299,
		7, 6, 0, 0, 299, 90, 1, 0, 0, 0, 9, 0, 191, 197, 199, 205, 214, 268, 278,
		292, 1, 6, 0, 0,
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
	SwiftGrammarLexerNUMBER        = 18
	SwiftGrammarLexerSTRING        = 19
	SwiftGrammarLexerID            = 20
	SwiftGrammarLexerDIF           = 21
	SwiftGrammarLexerIG_IG         = 22
	SwiftGrammarLexerNOT           = 23
	SwiftGrammarLexerOR            = 24
	SwiftGrammarLexerAND           = 25
	SwiftGrammarLexerIG            = 26
	SwiftGrammarLexerMAY_IG        = 27
	SwiftGrammarLexerMEN_IG        = 28
	SwiftGrammarLexerMAYOR         = 29
	SwiftGrammarLexerMENOR         = 30
	SwiftGrammarLexerMUL           = 31
	SwiftGrammarLexerDIV           = 32
	SwiftGrammarLexerADD           = 33
	SwiftGrammarLexerSUB           = 34
	SwiftGrammarLexerMOD           = 35
	SwiftGrammarLexerPARIZQ        = 36
	SwiftGrammarLexerPARDER        = 37
	SwiftGrammarLexerLLAVEIZQ      = 38
	SwiftGrammarLexerLLAVEDER      = 39
	SwiftGrammarLexerDOSPUNTOS     = 40
	SwiftGrammarLexerINTERROGACION = 41
	SwiftGrammarLexerWHITESPACE    = 42
	SwiftGrammarLexerCOMMENT       = 43
	SwiftGrammarLexerLINE_COMMENT  = 44
)

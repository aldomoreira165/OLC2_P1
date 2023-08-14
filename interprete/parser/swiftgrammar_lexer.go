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
		"'case'", "'default'", "'var'", "", "", "", "'!='", "'=='", "'!'", "'||'",
		"'&&'", "'='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'",
		"'%'", "'('", "')'", "'{'", "'}'", "':'", "'?'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "NIL", "TRU", "FAL",
		"PRINT", "IF", "ELSE", "WHILE", "SWITCH", "CASE", "DEFAULT", "VAR",
		"NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND", "IG",
		"MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB", "MOD",
		"PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "DOSPUNTOS", "INTERROGACION",
		"WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "NIL", "TRU", "FAL",
		"PRINT", "IF", "ELSE", "WHILE", "SWITCH", "CASE", "DEFAULT", "VAR",
		"NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND", "IG",
		"MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB", "MOD",
		"PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "DOSPUNTOS", "INTERROGACION",
		"WHITESPACE", "COMMENT", "LINE_COMMENT", "ESC_SEQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 43, 292, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 4, 16, 182, 8, 16, 11, 16, 12,
		16, 183, 1, 16, 1, 16, 4, 16, 188, 8, 16, 11, 16, 12, 16, 189, 3, 16, 192,
		8, 16, 1, 17, 1, 17, 5, 17, 196, 8, 17, 10, 17, 12, 17, 199, 9, 17, 1,
		17, 1, 17, 1, 18, 1, 18, 5, 18, 205, 8, 18, 10, 18, 12, 18, 208, 9, 18,
		1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26,
		1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1,
		31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36,
		1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 4, 40, 259, 8, 40, 11,
		40, 12, 40, 260, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 5, 41, 269,
		8, 41, 10, 41, 12, 41, 272, 9, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		42, 1, 42, 1, 42, 1, 42, 5, 42, 283, 8, 42, 10, 42, 12, 42, 286, 9, 42,
		1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 270, 0, 44, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45,
		23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63,
		32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81,
		41, 83, 42, 85, 43, 87, 0, 1, 0, 7, 1, 0, 48, 57, 1, 0, 34, 34, 3, 0, 65,
		90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 4, 0, 9, 10,
		13, 13, 32, 32, 92, 92, 2, 0, 10, 10, 13, 13, 7, 0, 32, 33, 35, 35, 43,
		43, 45, 46, 58, 58, 64, 64, 91, 93, 298, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0,
		35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0,
		0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0,
		0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0,
		0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1,
		0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73,
		1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0,
		81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 1, 89, 1, 0, 0, 0,
		3, 93, 1, 0, 0, 0, 5, 99, 1, 0, 0, 0, 7, 104, 1, 0, 0, 0, 9, 114, 1, 0,
		0, 0, 11, 121, 1, 0, 0, 0, 13, 125, 1, 0, 0, 0, 15, 130, 1, 0, 0, 0, 17,
		136, 1, 0, 0, 0, 19, 142, 1, 0, 0, 0, 21, 145, 1, 0, 0, 0, 23, 150, 1,
		0, 0, 0, 25, 156, 1, 0, 0, 0, 27, 163, 1, 0, 0, 0, 29, 168, 1, 0, 0, 0,
		31, 176, 1, 0, 0, 0, 33, 181, 1, 0, 0, 0, 35, 193, 1, 0, 0, 0, 37, 202,
		1, 0, 0, 0, 39, 209, 1, 0, 0, 0, 41, 212, 1, 0, 0, 0, 43, 215, 1, 0, 0,
		0, 45, 217, 1, 0, 0, 0, 47, 220, 1, 0, 0, 0, 49, 223, 1, 0, 0, 0, 51, 225,
		1, 0, 0, 0, 53, 228, 1, 0, 0, 0, 55, 231, 1, 0, 0, 0, 57, 233, 1, 0, 0,
		0, 59, 235, 1, 0, 0, 0, 61, 237, 1, 0, 0, 0, 63, 239, 1, 0, 0, 0, 65, 241,
		1, 0, 0, 0, 67, 243, 1, 0, 0, 0, 69, 245, 1, 0, 0, 0, 71, 247, 1, 0, 0,
		0, 73, 249, 1, 0, 0, 0, 75, 251, 1, 0, 0, 0, 77, 253, 1, 0, 0, 0, 79, 255,
		1, 0, 0, 0, 81, 258, 1, 0, 0, 0, 83, 264, 1, 0, 0, 0, 85, 278, 1, 0, 0,
		0, 87, 289, 1, 0, 0, 0, 89, 90, 5, 105, 0, 0, 90, 91, 5, 110, 0, 0, 91,
		92, 5, 116, 0, 0, 92, 2, 1, 0, 0, 0, 93, 94, 5, 102, 0, 0, 94, 95, 5, 108,
		0, 0, 95, 96, 5, 111, 0, 0, 96, 97, 5, 97, 0, 0, 97, 98, 5, 116, 0, 0,
		98, 4, 1, 0, 0, 0, 99, 100, 5, 98, 0, 0, 100, 101, 5, 111, 0, 0, 101, 102,
		5, 111, 0, 0, 102, 103, 5, 108, 0, 0, 103, 6, 1, 0, 0, 0, 104, 105, 5,
		99, 0, 0, 105, 106, 5, 104, 0, 0, 106, 107, 5, 97, 0, 0, 107, 108, 5, 114,
		0, 0, 108, 109, 5, 97, 0, 0, 109, 110, 5, 99, 0, 0, 110, 111, 5, 116, 0,
		0, 111, 112, 5, 101, 0, 0, 112, 113, 5, 114, 0, 0, 113, 8, 1, 0, 0, 0,
		114, 115, 5, 83, 0, 0, 115, 116, 5, 116, 0, 0, 116, 117, 5, 114, 0, 0,
		117, 118, 5, 105, 0, 0, 118, 119, 5, 110, 0, 0, 119, 120, 5, 103, 0, 0,
		120, 10, 1, 0, 0, 0, 121, 122, 5, 110, 0, 0, 122, 123, 5, 105, 0, 0, 123,
		124, 5, 108, 0, 0, 124, 12, 1, 0, 0, 0, 125, 126, 5, 116, 0, 0, 126, 127,
		5, 114, 0, 0, 127, 128, 5, 117, 0, 0, 128, 129, 5, 101, 0, 0, 129, 14,
		1, 0, 0, 0, 130, 131, 5, 102, 0, 0, 131, 132, 5, 97, 0, 0, 132, 133, 5,
		108, 0, 0, 133, 134, 5, 115, 0, 0, 134, 135, 5, 101, 0, 0, 135, 16, 1,
		0, 0, 0, 136, 137, 5, 112, 0, 0, 137, 138, 5, 114, 0, 0, 138, 139, 5, 105,
		0, 0, 139, 140, 5, 110, 0, 0, 140, 141, 5, 116, 0, 0, 141, 18, 1, 0, 0,
		0, 142, 143, 5, 105, 0, 0, 143, 144, 5, 102, 0, 0, 144, 20, 1, 0, 0, 0,
		145, 146, 5, 101, 0, 0, 146, 147, 5, 108, 0, 0, 147, 148, 5, 115, 0, 0,
		148, 149, 5, 101, 0, 0, 149, 22, 1, 0, 0, 0, 150, 151, 5, 119, 0, 0, 151,
		152, 5, 104, 0, 0, 152, 153, 5, 105, 0, 0, 153, 154, 5, 108, 0, 0, 154,
		155, 5, 101, 0, 0, 155, 24, 1, 0, 0, 0, 156, 157, 5, 115, 0, 0, 157, 158,
		5, 119, 0, 0, 158, 159, 5, 105, 0, 0, 159, 160, 5, 116, 0, 0, 160, 161,
		5, 99, 0, 0, 161, 162, 5, 104, 0, 0, 162, 26, 1, 0, 0, 0, 163, 164, 5,
		99, 0, 0, 164, 165, 5, 97, 0, 0, 165, 166, 5, 115, 0, 0, 166, 167, 5, 101,
		0, 0, 167, 28, 1, 0, 0, 0, 168, 169, 5, 100, 0, 0, 169, 170, 5, 101, 0,
		0, 170, 171, 5, 102, 0, 0, 171, 172, 5, 97, 0, 0, 172, 173, 5, 117, 0,
		0, 173, 174, 5, 108, 0, 0, 174, 175, 5, 116, 0, 0, 175, 30, 1, 0, 0, 0,
		176, 177, 5, 118, 0, 0, 177, 178, 5, 97, 0, 0, 178, 179, 5, 114, 0, 0,
		179, 32, 1, 0, 0, 0, 180, 182, 7, 0, 0, 0, 181, 180, 1, 0, 0, 0, 182, 183,
		1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 191, 1, 0,
		0, 0, 185, 187, 5, 46, 0, 0, 186, 188, 7, 0, 0, 0, 187, 186, 1, 0, 0, 0,
		188, 189, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190,
		192, 1, 0, 0, 0, 191, 185, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 34, 1,
		0, 0, 0, 193, 197, 5, 34, 0, 0, 194, 196, 8, 1, 0, 0, 195, 194, 1, 0, 0,
		0, 196, 199, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198,
		200, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 200, 201, 5, 34, 0, 0, 201, 36,
		1, 0, 0, 0, 202, 206, 7, 2, 0, 0, 203, 205, 7, 3, 0, 0, 204, 203, 1, 0,
		0, 0, 205, 208, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0,
		207, 38, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 209, 210, 5, 33, 0, 0, 210,
		211, 5, 61, 0, 0, 211, 40, 1, 0, 0, 0, 212, 213, 5, 61, 0, 0, 213, 214,
		5, 61, 0, 0, 214, 42, 1, 0, 0, 0, 215, 216, 5, 33, 0, 0, 216, 44, 1, 0,
		0, 0, 217, 218, 5, 124, 0, 0, 218, 219, 5, 124, 0, 0, 219, 46, 1, 0, 0,
		0, 220, 221, 5, 38, 0, 0, 221, 222, 5, 38, 0, 0, 222, 48, 1, 0, 0, 0, 223,
		224, 5, 61, 0, 0, 224, 50, 1, 0, 0, 0, 225, 226, 5, 62, 0, 0, 226, 227,
		5, 61, 0, 0, 227, 52, 1, 0, 0, 0, 228, 229, 5, 60, 0, 0, 229, 230, 5, 61,
		0, 0, 230, 54, 1, 0, 0, 0, 231, 232, 5, 62, 0, 0, 232, 56, 1, 0, 0, 0,
		233, 234, 5, 60, 0, 0, 234, 58, 1, 0, 0, 0, 235, 236, 5, 42, 0, 0, 236,
		60, 1, 0, 0, 0, 237, 238, 5, 47, 0, 0, 238, 62, 1, 0, 0, 0, 239, 240, 5,
		43, 0, 0, 240, 64, 1, 0, 0, 0, 241, 242, 5, 45, 0, 0, 242, 66, 1, 0, 0,
		0, 243, 244, 5, 37, 0, 0, 244, 68, 1, 0, 0, 0, 245, 246, 5, 40, 0, 0, 246,
		70, 1, 0, 0, 0, 247, 248, 5, 41, 0, 0, 248, 72, 1, 0, 0, 0, 249, 250, 5,
		123, 0, 0, 250, 74, 1, 0, 0, 0, 251, 252, 5, 125, 0, 0, 252, 76, 1, 0,
		0, 0, 253, 254, 5, 58, 0, 0, 254, 78, 1, 0, 0, 0, 255, 256, 5, 63, 0, 0,
		256, 80, 1, 0, 0, 0, 257, 259, 7, 4, 0, 0, 258, 257, 1, 0, 0, 0, 259, 260,
		1, 0, 0, 0, 260, 258, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 262, 1, 0,
		0, 0, 262, 263, 6, 40, 0, 0, 263, 82, 1, 0, 0, 0, 264, 265, 5, 47, 0, 0,
		265, 266, 5, 42, 0, 0, 266, 270, 1, 0, 0, 0, 267, 269, 9, 0, 0, 0, 268,
		267, 1, 0, 0, 0, 269, 272, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 270, 268,
		1, 0, 0, 0, 271, 273, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 273, 274, 5, 42,
		0, 0, 274, 275, 5, 47, 0, 0, 275, 276, 1, 0, 0, 0, 276, 277, 6, 41, 0,
		0, 277, 84, 1, 0, 0, 0, 278, 279, 5, 47, 0, 0, 279, 280, 5, 47, 0, 0, 280,
		284, 1, 0, 0, 0, 281, 283, 8, 5, 0, 0, 282, 281, 1, 0, 0, 0, 283, 286,
		1, 0, 0, 0, 284, 282, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 287, 1, 0,
		0, 0, 286, 284, 1, 0, 0, 0, 287, 288, 6, 42, 0, 0, 288, 86, 1, 0, 0, 0,
		289, 290, 5, 92, 0, 0, 290, 291, 7, 6, 0, 0, 291, 88, 1, 0, 0, 0, 9, 0,
		183, 189, 191, 197, 206, 260, 270, 284, 1, 6, 0, 0,
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
	SwiftGrammarLexerNUMBER        = 17
	SwiftGrammarLexerSTRING        = 18
	SwiftGrammarLexerID            = 19
	SwiftGrammarLexerDIF           = 20
	SwiftGrammarLexerIG_IG         = 21
	SwiftGrammarLexerNOT           = 22
	SwiftGrammarLexerOR            = 23
	SwiftGrammarLexerAND           = 24
	SwiftGrammarLexerIG            = 25
	SwiftGrammarLexerMAY_IG        = 26
	SwiftGrammarLexerMEN_IG        = 27
	SwiftGrammarLexerMAYOR         = 28
	SwiftGrammarLexerMENOR         = 29
	SwiftGrammarLexerMUL           = 30
	SwiftGrammarLexerDIV           = 31
	SwiftGrammarLexerADD           = 32
	SwiftGrammarLexerSUB           = 33
	SwiftGrammarLexerMOD           = 34
	SwiftGrammarLexerPARIZQ        = 35
	SwiftGrammarLexerPARDER        = 36
	SwiftGrammarLexerLLAVEIZQ      = 37
	SwiftGrammarLexerLLAVEDER      = 38
	SwiftGrammarLexerDOSPUNTOS     = 39
	SwiftGrammarLexerINTERROGACION = 40
	SwiftGrammarLexerWHITESPACE    = 41
	SwiftGrammarLexerCOMMENT       = 42
	SwiftGrammarLexerLINE_COMMENT  = 43
)

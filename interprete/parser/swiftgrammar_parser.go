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
		"'true'", "'false'", "'print'", "'if'", "'else'", "'while'", "'for'",
		"'guard'", "'in'", "'switch'", "'case'", "'default'", "'var'", "'let'",
		"'break'", "'return'", "'continue'", "'func'", "'count'", "'IsEmpty'",
		"'append'", "'removeLast'", "'remove'", "'struct'", "'varst'", "'letst'",
		"'at'", "", "", "", "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", "'>='",
		"'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'%'", "'('", "')'",
		"'{'", "'}'", "'['", "']'", "':'", "','", "';'", "'?'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "BOOL", "CHARACTER", "PSTRING", "NIL", "TRU", "FAL",
		"PRINT", "IF", "ELSE", "WHILE", "FOR", "GUARD", "IN", "SWITCH", "CASE",
		"DEFAULT", "VAR", "LET", "BREAK", "RETURN", "CONTINUE", "FUNC", "COUNT",
		"ISEMPTY", "APPEND", "REMOVE_LAST", "REMOVE", "STRUCT", "STRUCT_VAR",
		"STRUCT_LET", "AT", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT",
		"OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV",
		"ADD", "SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "CORCHIZQ",
		"CORCHDER", "DOSPUNTOS", "COMA", "PTCOMA", "INTERROGACION", "PUNTO",
		"WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "stmt", "defstructstmt", "lista_atributos", "struct_expr",
		"valor_struct_expr", "l_dupla", "accesostructstmt", "declvectorstmt",
		"defvectorstmt", "listaexpresiones", "accesovectorstmt", "asignvectorstmt",
		"appendvectorstmt", "removeatvectorstmt", "removelastvectorstmt", "countvectorstmt",
		"isemptyvectorstmt", "declmatrizstmt", "tipomatriz", "listavaloresmatriz",
		"listavaloresmatriz3", "accesomatriz", "asignmatrizstmt", "returnstmt",
		"breakstmt", "continuestmt", "printstmt", "intstmt", "floatstmt", "stringstmt",
		"funcdclstmt", "accfuncstm", "parametros", "parametroscall", "ifstmt",
		"elseifstmt", "switchstmt", "caseStmt", "defaultCase", "typedDeclstmt",
		"untypedDeclstmt", "optionalTypedDeclstmt", "asignstmt", "whilestmt",
		"forstmt", "guardstmt", "rangostmt", "opasignstmt", "expr", "tipo",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 65, 628, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 3, 1, 110, 8, 1, 5, 1, 112, 8, 1, 10, 1, 12, 1, 115,
		9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 3, 2, 143, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 149, 8,
		3, 10, 3, 12, 3, 152, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3,
		4, 161, 8, 4, 3, 4, 163, 8, 4, 1, 4, 1, 4, 3, 4, 167, 8, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 3, 5, 173, 8, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 3, 6, 183, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7,
		192, 8, 7, 10, 7, 12, 7, 195, 9, 7, 1, 8, 1, 8, 1, 8, 4, 8, 200, 8, 8,
		11, 8, 12, 8, 201, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 3, 10, 215, 8, 10, 1, 10, 1, 10, 1, 10, 3, 10, 220, 8, 10,
		1, 11, 1, 11, 1, 11, 5, 11, 225, 8, 11, 10, 11, 12, 11, 228, 9, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 19, 1, 19, 3, 19, 276, 8, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 3, 19, 284, 8, 19, 1, 19, 1, 19, 3, 19, 288, 8, 19, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 3, 20, 304, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 4, 21, 315, 8, 21, 11, 21, 12, 21, 316, 1,
		21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 325, 8, 22, 10, 22, 12, 22,
		328, 9, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 3, 23, 351, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 376, 8, 24, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32,
		3, 32, 410, 8, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 3, 32, 421, 8, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 3, 32, 431, 8, 32, 1, 33, 1, 33, 1, 33, 3, 33, 436, 8, 33,
		1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 5, 34, 449, 8, 34, 10, 34, 12, 34, 452, 9, 34, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 5, 35, 461, 8, 35, 10, 35, 12, 35, 464, 9,
		35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 5, 36, 472, 8, 36, 10, 36,
		12, 36, 475, 9, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 482, 8, 36,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1,
		38, 4, 38, 495, 8, 38, 11, 38, 12, 38, 496, 1, 38, 3, 38, 500, 8, 38, 1,
		38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40,
		1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1,
		42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44,
		1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1,
		46, 1, 46, 3, 46, 546, 8, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47,
		1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1,
		48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 3, 49, 573,
		8, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1,
		50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50,
		1, 50, 1, 50, 1, 50, 3, 50, 598, 8, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1,
		50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50,
		1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 5, 50, 621, 8, 50, 10, 50, 12,
		50, 624, 9, 50, 1, 51, 1, 51, 1, 51, 0, 1, 100, 52, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
		48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82,
		84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 0, 9, 1, 0, 19, 20, 1, 0, 31,
		32, 1, 0, 7, 8, 2, 0, 47, 48, 51, 51, 1, 0, 49, 50, 2, 0, 43, 43, 45, 45,
		2, 0, 44, 44, 46, 46, 1, 0, 37, 38, 1, 0, 1, 5, 656, 0, 104, 1, 0, 0, 0,
		2, 113, 1, 0, 0, 0, 4, 142, 1, 0, 0, 0, 6, 144, 1, 0, 0, 0, 8, 155, 1,
		0, 0, 0, 10, 168, 1, 0, 0, 0, 12, 177, 1, 0, 0, 0, 14, 184, 1, 0, 0, 0,
		16, 196, 1, 0, 0, 0, 18, 203, 1, 0, 0, 0, 20, 219, 1, 0, 0, 0, 22, 221,
		1, 0, 0, 0, 24, 229, 1, 0, 0, 0, 26, 234, 1, 0, 0, 0, 28, 241, 1, 0, 0,
		0, 30, 248, 1, 0, 0, 0, 32, 257, 1, 0, 0, 0, 34, 263, 1, 0, 0, 0, 36, 267,
		1, 0, 0, 0, 38, 287, 1, 0, 0, 0, 40, 303, 1, 0, 0, 0, 42, 305, 1, 0, 0,
		0, 44, 320, 1, 0, 0, 0, 46, 350, 1, 0, 0, 0, 48, 375, 1, 0, 0, 0, 50, 377,
		1, 0, 0, 0, 52, 381, 1, 0, 0, 0, 54, 383, 1, 0, 0, 0, 56, 385, 1, 0, 0,
		0, 58, 390, 1, 0, 0, 0, 60, 395, 1, 0, 0, 0, 62, 400, 1, 0, 0, 0, 64, 430,
		1, 0, 0, 0, 66, 432, 1, 0, 0, 0, 68, 439, 1, 0, 0, 0, 70, 453, 1, 0, 0,
		0, 72, 465, 1, 0, 0, 0, 74, 483, 1, 0, 0, 0, 76, 490, 1, 0, 0, 0, 78, 503,
		1, 0, 0, 0, 80, 508, 1, 0, 0, 0, 82, 512, 1, 0, 0, 0, 84, 519, 1, 0, 0,
		0, 86, 524, 1, 0, 0, 0, 88, 530, 1, 0, 0, 0, 90, 534, 1, 0, 0, 0, 92, 540,
		1, 0, 0, 0, 94, 551, 1, 0, 0, 0, 96, 558, 1, 0, 0, 0, 98, 572, 1, 0, 0,
		0, 100, 597, 1, 0, 0, 0, 102, 625, 1, 0, 0, 0, 104, 105, 3, 2, 1, 0, 105,
		106, 5, 0, 0, 1, 106, 1, 1, 0, 0, 0, 107, 109, 3, 4, 2, 0, 108, 110, 5,
		60, 0, 0, 109, 108, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 112, 1, 0, 0,
		0, 111, 107, 1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113,
		114, 1, 0, 0, 0, 114, 3, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 143, 3,
		56, 28, 0, 117, 143, 3, 82, 41, 0, 118, 143, 3, 84, 42, 0, 119, 143, 3,
		86, 43, 0, 120, 143, 3, 88, 44, 0, 121, 143, 3, 72, 36, 0, 122, 143, 3,
		76, 38, 0, 123, 143, 3, 90, 45, 0, 124, 143, 3, 92, 46, 0, 125, 143, 3,
		94, 47, 0, 126, 143, 3, 98, 49, 0, 127, 143, 3, 64, 32, 0, 128, 143, 3,
		66, 33, 0, 129, 143, 3, 52, 26, 0, 130, 143, 3, 54, 27, 0, 131, 143, 3,
		50, 25, 0, 132, 143, 3, 18, 9, 0, 133, 143, 3, 24, 12, 0, 134, 143, 3,
		28, 14, 0, 135, 143, 3, 32, 16, 0, 136, 143, 3, 30, 15, 0, 137, 143, 3,
		26, 13, 0, 138, 143, 3, 38, 19, 0, 139, 143, 3, 48, 24, 0, 140, 143, 3,
		6, 3, 0, 141, 143, 3, 10, 5, 0, 142, 116, 1, 0, 0, 0, 142, 117, 1, 0, 0,
		0, 142, 118, 1, 0, 0, 0, 142, 119, 1, 0, 0, 0, 142, 120, 1, 0, 0, 0, 142,
		121, 1, 0, 0, 0, 142, 122, 1, 0, 0, 0, 142, 123, 1, 0, 0, 0, 142, 124,
		1, 0, 0, 0, 142, 125, 1, 0, 0, 0, 142, 126, 1, 0, 0, 0, 142, 127, 1, 0,
		0, 0, 142, 128, 1, 0, 0, 0, 142, 129, 1, 0, 0, 0, 142, 130, 1, 0, 0, 0,
		142, 131, 1, 0, 0, 0, 142, 132, 1, 0, 0, 0, 142, 133, 1, 0, 0, 0, 142,
		134, 1, 0, 0, 0, 142, 135, 1, 0, 0, 0, 142, 136, 1, 0, 0, 0, 142, 137,
		1, 0, 0, 0, 142, 138, 1, 0, 0, 0, 142, 139, 1, 0, 0, 0, 142, 140, 1, 0,
		0, 0, 142, 141, 1, 0, 0, 0, 143, 5, 1, 0, 0, 0, 144, 145, 5, 30, 0, 0,
		145, 146, 5, 36, 0, 0, 146, 150, 5, 54, 0, 0, 147, 149, 3, 8, 4, 0, 148,
		147, 1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151,
		1, 0, 0, 0, 151, 153, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 153, 154, 5, 55,
		0, 0, 154, 7, 1, 0, 0, 0, 155, 156, 7, 0, 0, 0, 156, 162, 5, 36, 0, 0,
		157, 160, 5, 58, 0, 0, 158, 161, 3, 102, 51, 0, 159, 161, 5, 36, 0, 0,
		160, 158, 1, 0, 0, 0, 160, 159, 1, 0, 0, 0, 161, 163, 1, 0, 0, 0, 162,
		157, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 166, 1, 0, 0, 0, 164, 165,
		5, 42, 0, 0, 165, 167, 3, 100, 50, 0, 166, 164, 1, 0, 0, 0, 166, 167, 1,
		0, 0, 0, 167, 9, 1, 0, 0, 0, 168, 169, 7, 1, 0, 0, 169, 172, 5, 36, 0,
		0, 170, 171, 5, 58, 0, 0, 171, 173, 5, 36, 0, 0, 172, 170, 1, 0, 0, 0,
		172, 173, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 175, 5, 42, 0, 0, 175,
		176, 3, 12, 6, 0, 176, 11, 1, 0, 0, 0, 177, 182, 5, 36, 0, 0, 178, 179,
		5, 52, 0, 0, 179, 180, 3, 14, 7, 0, 180, 181, 5, 53, 0, 0, 181, 183, 1,
		0, 0, 0, 182, 178, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 13, 1, 0, 0,
		0, 184, 185, 5, 36, 0, 0, 185, 186, 5, 58, 0, 0, 186, 193, 3, 100, 50,
		0, 187, 188, 5, 59, 0, 0, 188, 189, 5, 36, 0, 0, 189, 190, 5, 58, 0, 0,
		190, 192, 3, 100, 50, 0, 191, 187, 1, 0, 0, 0, 192, 195, 1, 0, 0, 0, 193,
		191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 15, 1, 0, 0, 0, 195, 193, 1,
		0, 0, 0, 196, 199, 5, 36, 0, 0, 197, 198, 5, 62, 0, 0, 198, 200, 5, 36,
		0, 0, 199, 197, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0,
		201, 202, 1, 0, 0, 0, 202, 17, 1, 0, 0, 0, 203, 204, 5, 19, 0, 0, 204,
		205, 5, 36, 0, 0, 205, 206, 5, 58, 0, 0, 206, 207, 5, 56, 0, 0, 207, 208,
		3, 102, 51, 0, 208, 209, 5, 57, 0, 0, 209, 210, 3, 20, 10, 0, 210, 19,
		1, 0, 0, 0, 211, 212, 5, 42, 0, 0, 212, 214, 5, 56, 0, 0, 213, 215, 3,
		22, 11, 0, 214, 213, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 216, 1, 0,
		0, 0, 216, 220, 5, 57, 0, 0, 217, 218, 5, 42, 0, 0, 218, 220, 5, 36, 0,
		0, 219, 211, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 220, 21, 1, 0, 0, 0, 221,
		226, 3, 100, 50, 0, 222, 223, 5, 59, 0, 0, 223, 225, 3, 100, 50, 0, 224,
		222, 1, 0, 0, 0, 225, 228, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 226, 227,
		1, 0, 0, 0, 227, 23, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 229, 230, 5, 36,
		0, 0, 230, 231, 5, 56, 0, 0, 231, 232, 3, 100, 50, 0, 232, 233, 5, 57,
		0, 0, 233, 25, 1, 0, 0, 0, 234, 235, 5, 36, 0, 0, 235, 236, 5, 56, 0, 0,
		236, 237, 3, 100, 50, 0, 237, 238, 5, 57, 0, 0, 238, 239, 5, 42, 0, 0,
		239, 240, 3, 100, 50, 0, 240, 27, 1, 0, 0, 0, 241, 242, 5, 36, 0, 0, 242,
		243, 5, 62, 0, 0, 243, 244, 5, 27, 0, 0, 244, 245, 5, 52, 0, 0, 245, 246,
		3, 100, 50, 0, 246, 247, 5, 53, 0, 0, 247, 29, 1, 0, 0, 0, 248, 249, 5,
		36, 0, 0, 249, 250, 5, 62, 0, 0, 250, 251, 5, 29, 0, 0, 251, 252, 5, 52,
		0, 0, 252, 253, 5, 33, 0, 0, 253, 254, 5, 58, 0, 0, 254, 255, 3, 100, 50,
		0, 255, 256, 5, 53, 0, 0, 256, 31, 1, 0, 0, 0, 257, 258, 5, 36, 0, 0, 258,
		259, 5, 62, 0, 0, 259, 260, 5, 28, 0, 0, 260, 261, 5, 52, 0, 0, 261, 262,
		5, 53, 0, 0, 262, 33, 1, 0, 0, 0, 263, 264, 5, 36, 0, 0, 264, 265, 5, 62,
		0, 0, 265, 266, 5, 25, 0, 0, 266, 35, 1, 0, 0, 0, 267, 268, 5, 36, 0, 0,
		268, 269, 5, 62, 0, 0, 269, 270, 5, 26, 0, 0, 270, 37, 1, 0, 0, 0, 271,
		272, 5, 19, 0, 0, 272, 275, 5, 36, 0, 0, 273, 274, 5, 58, 0, 0, 274, 276,
		3, 40, 20, 0, 275, 273, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 277, 1,
		0, 0, 0, 277, 278, 5, 42, 0, 0, 278, 288, 3, 42, 21, 0, 279, 280, 5, 19,
		0, 0, 280, 283, 5, 36, 0, 0, 281, 282, 5, 58, 0, 0, 282, 284, 3, 40, 20,
		0, 283, 281, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285,
		286, 5, 42, 0, 0, 286, 288, 3, 44, 22, 0, 287, 271, 1, 0, 0, 0, 287, 279,
		1, 0, 0, 0, 288, 39, 1, 0, 0, 0, 289, 290, 5, 56, 0, 0, 290, 291, 5, 56,
		0, 0, 291, 292, 3, 102, 51, 0, 292, 293, 5, 57, 0, 0, 293, 294, 5, 57,
		0, 0, 294, 304, 1, 0, 0, 0, 295, 296, 5, 56, 0, 0, 296, 297, 5, 56, 0,
		0, 297, 298, 5, 56, 0, 0, 298, 299, 3, 102, 51, 0, 299, 300, 5, 57, 0,
		0, 300, 301, 5, 57, 0, 0, 301, 302, 5, 57, 0, 0, 302, 304, 1, 0, 0, 0,
		303, 289, 1, 0, 0, 0, 303, 295, 1, 0, 0, 0, 304, 41, 1, 0, 0, 0, 305, 306,
		5, 56, 0, 0, 306, 307, 5, 56, 0, 0, 307, 308, 3, 22, 11, 0, 308, 314, 5,
		57, 0, 0, 309, 310, 5, 59, 0, 0, 310, 311, 5, 56, 0, 0, 311, 312, 3, 22,
		11, 0, 312, 313, 5, 57, 0, 0, 313, 315, 1, 0, 0, 0, 314, 309, 1, 0, 0,
		0, 315, 316, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 316, 317, 1, 0, 0, 0, 317,
		318, 1, 0, 0, 0, 318, 319, 5, 57, 0, 0, 319, 43, 1, 0, 0, 0, 320, 321,
		5, 56, 0, 0, 321, 326, 3, 42, 21, 0, 322, 323, 5, 59, 0, 0, 323, 325, 3,
		42, 21, 0, 324, 322, 1, 0, 0, 0, 325, 328, 1, 0, 0, 0, 326, 324, 1, 0,
		0, 0, 326, 327, 1, 0, 0, 0, 327, 329, 1, 0, 0, 0, 328, 326, 1, 0, 0, 0,
		329, 330, 5, 57, 0, 0, 330, 45, 1, 0, 0, 0, 331, 332, 5, 36, 0, 0, 332,
		333, 5, 56, 0, 0, 333, 334, 3, 100, 50, 0, 334, 335, 5, 57, 0, 0, 335,
		336, 5, 56, 0, 0, 336, 337, 3, 100, 50, 0, 337, 338, 5, 57, 0, 0, 338,
		351, 1, 0, 0, 0, 339, 340, 5, 36, 0, 0, 340, 341, 5, 56, 0, 0, 341, 342,
		3, 100, 50, 0, 342, 343, 5, 57, 0, 0, 343, 344, 5, 56, 0, 0, 344, 345,
		3, 100, 50, 0, 345, 346, 5, 57, 0, 0, 346, 347, 5, 56, 0, 0, 347, 348,
		3, 100, 50, 0, 348, 349, 5, 57, 0, 0, 349, 351, 1, 0, 0, 0, 350, 331, 1,
		0, 0, 0, 350, 339, 1, 0, 0, 0, 351, 47, 1, 0, 0, 0, 352, 353, 5, 36, 0,
		0, 353, 354, 5, 56, 0, 0, 354, 355, 3, 100, 50, 0, 355, 356, 5, 57, 0,
		0, 356, 357, 5, 56, 0, 0, 357, 358, 3, 100, 50, 0, 358, 359, 5, 57, 0,
		0, 359, 360, 5, 42, 0, 0, 360, 361, 3, 100, 50, 0, 361, 376, 1, 0, 0, 0,
		362, 363, 5, 36, 0, 0, 363, 364, 5, 56, 0, 0, 364, 365, 3, 100, 50, 0,
		365, 366, 5, 57, 0, 0, 366, 367, 5, 56, 0, 0, 367, 368, 3, 100, 50, 0,
		368, 369, 5, 57, 0, 0, 369, 370, 5, 56, 0, 0, 370, 371, 3, 100, 50, 0,
		371, 372, 5, 57, 0, 0, 372, 373, 5, 42, 0, 0, 373, 374, 3, 100, 50, 0,
		374, 376, 1, 0, 0, 0, 375, 352, 1, 0, 0, 0, 375, 362, 1, 0, 0, 0, 376,
		49, 1, 0, 0, 0, 377, 378, 5, 22, 0, 0, 378, 379, 3, 100, 50, 0, 379, 380,
		5, 60, 0, 0, 380, 51, 1, 0, 0, 0, 381, 382, 5, 21, 0, 0, 382, 53, 1, 0,
		0, 0, 383, 384, 5, 23, 0, 0, 384, 55, 1, 0, 0, 0, 385, 386, 5, 9, 0, 0,
		386, 387, 5, 52, 0, 0, 387, 388, 3, 100, 50, 0, 388, 389, 5, 53, 0, 0,
		389, 57, 1, 0, 0, 0, 390, 391, 5, 1, 0, 0, 391, 392, 5, 52, 0, 0, 392,
		393, 3, 100, 50, 0, 393, 394, 5, 53, 0, 0, 394, 59, 1, 0, 0, 0, 395, 396,
		5, 2, 0, 0, 396, 397, 5, 52, 0, 0, 397, 398, 3, 100, 50, 0, 398, 399, 5,
		53, 0, 0, 399, 61, 1, 0, 0, 0, 400, 401, 5, 5, 0, 0, 401, 402, 5, 52, 0,
		0, 402, 403, 3, 100, 50, 0, 403, 404, 5, 53, 0, 0, 404, 63, 1, 0, 0, 0,
		405, 406, 5, 24, 0, 0, 406, 407, 5, 36, 0, 0, 407, 409, 5, 52, 0, 0, 408,
		410, 3, 68, 34, 0, 409, 408, 1, 0, 0, 0, 409, 410, 1, 0, 0, 0, 410, 411,
		1, 0, 0, 0, 411, 412, 5, 53, 0, 0, 412, 413, 5, 54, 0, 0, 413, 414, 3,
		2, 1, 0, 414, 415, 5, 55, 0, 0, 415, 431, 1, 0, 0, 0, 416, 417, 5, 24,
		0, 0, 417, 418, 5, 36, 0, 0, 418, 420, 5, 52, 0, 0, 419, 421, 3, 68, 34,
		0, 420, 419, 1, 0, 0, 0, 420, 421, 1, 0, 0, 0, 421, 422, 1, 0, 0, 0, 422,
		423, 5, 53, 0, 0, 423, 424, 5, 50, 0, 0, 424, 425, 5, 45, 0, 0, 425, 426,
		3, 102, 51, 0, 426, 427, 5, 54, 0, 0, 427, 428, 3, 2, 1, 0, 428, 429, 5,
		55, 0, 0, 429, 431, 1, 0, 0, 0, 430, 405, 1, 0, 0, 0, 430, 416, 1, 0, 0,
		0, 431, 65, 1, 0, 0, 0, 432, 433, 5, 36, 0, 0, 433, 435, 5, 52, 0, 0, 434,
		436, 3, 70, 35, 0, 435, 434, 1, 0, 0, 0, 435, 436, 1, 0, 0, 0, 436, 437,
		1, 0, 0, 0, 437, 438, 5, 53, 0, 0, 438, 67, 1, 0, 0, 0, 439, 440, 5, 36,
		0, 0, 440, 441, 5, 36, 0, 0, 441, 442, 5, 58, 0, 0, 442, 450, 3, 102, 51,
		0, 443, 444, 5, 59, 0, 0, 444, 445, 5, 36, 0, 0, 445, 446, 5, 36, 0, 0,
		446, 447, 5, 58, 0, 0, 447, 449, 3, 102, 51, 0, 448, 443, 1, 0, 0, 0, 449,
		452, 1, 0, 0, 0, 450, 448, 1, 0, 0, 0, 450, 451, 1, 0, 0, 0, 451, 69, 1,
		0, 0, 0, 452, 450, 1, 0, 0, 0, 453, 454, 5, 36, 0, 0, 454, 455, 5, 58,
		0, 0, 455, 462, 3, 100, 50, 0, 456, 457, 5, 59, 0, 0, 457, 458, 5, 36,
		0, 0, 458, 459, 5, 58, 0, 0, 459, 461, 3, 100, 50, 0, 460, 456, 1, 0, 0,
		0, 461, 464, 1, 0, 0, 0, 462, 460, 1, 0, 0, 0, 462, 463, 1, 0, 0, 0, 463,
		71, 1, 0, 0, 0, 464, 462, 1, 0, 0, 0, 465, 466, 5, 10, 0, 0, 466, 467,
		3, 100, 50, 0, 467, 468, 5, 54, 0, 0, 468, 469, 3, 2, 1, 0, 469, 473, 5,
		55, 0, 0, 470, 472, 3, 74, 37, 0, 471, 470, 1, 0, 0, 0, 472, 475, 1, 0,
		0, 0, 473, 471, 1, 0, 0, 0, 473, 474, 1, 0, 0, 0, 474, 481, 1, 0, 0, 0,
		475, 473, 1, 0, 0, 0, 476, 477, 5, 11, 0, 0, 477, 478, 5, 54, 0, 0, 478,
		479, 3, 2, 1, 0, 479, 480, 5, 55, 0, 0, 480, 482, 1, 0, 0, 0, 481, 476,
		1, 0, 0, 0, 481, 482, 1, 0, 0, 0, 482, 73, 1, 0, 0, 0, 483, 484, 5, 11,
		0, 0, 484, 485, 5, 10, 0, 0, 485, 486, 3, 100, 50, 0, 486, 487, 5, 54,
		0, 0, 487, 488, 3, 2, 1, 0, 488, 489, 5, 55, 0, 0, 489, 75, 1, 0, 0, 0,
		490, 491, 5, 16, 0, 0, 491, 492, 3, 100, 50, 0, 492, 494, 5, 54, 0, 0,
		493, 495, 3, 78, 39, 0, 494, 493, 1, 0, 0, 0, 495, 496, 1, 0, 0, 0, 496,
		494, 1, 0, 0, 0, 496, 497, 1, 0, 0, 0, 497, 499, 1, 0, 0, 0, 498, 500,
		3, 80, 40, 0, 499, 498, 1, 0, 0, 0, 499, 500, 1, 0, 0, 0, 500, 501, 1,
		0, 0, 0, 501, 502, 5, 55, 0, 0, 502, 77, 1, 0, 0, 0, 503, 504, 5, 17, 0,
		0, 504, 505, 3, 100, 50, 0, 505, 506, 5, 58, 0, 0, 506, 507, 3, 2, 1, 0,
		507, 79, 1, 0, 0, 0, 508, 509, 5, 18, 0, 0, 509, 510, 5, 58, 0, 0, 510,
		511, 3, 2, 1, 0, 511, 81, 1, 0, 0, 0, 512, 513, 7, 0, 0, 0, 513, 514, 5,
		36, 0, 0, 514, 515, 5, 58, 0, 0, 515, 516, 3, 102, 51, 0, 516, 517, 5,
		42, 0, 0, 517, 518, 3, 100, 50, 0, 518, 83, 1, 0, 0, 0, 519, 520, 7, 0,
		0, 0, 520, 521, 5, 36, 0, 0, 521, 522, 5, 42, 0, 0, 522, 523, 3, 100, 50,
		0, 523, 85, 1, 0, 0, 0, 524, 525, 5, 19, 0, 0, 525, 526, 5, 36, 0, 0, 526,
		527, 5, 58, 0, 0, 527, 528, 3, 102, 51, 0, 528, 529, 5, 61, 0, 0, 529,
		87, 1, 0, 0, 0, 530, 531, 5, 36, 0, 0, 531, 532, 5, 42, 0, 0, 532, 533,
		3, 100, 50, 0, 533, 89, 1, 0, 0, 0, 534, 535, 5, 12, 0, 0, 535, 536, 3,
		100, 50, 0, 536, 537, 5, 54, 0, 0, 537, 538, 3, 2, 1, 0, 538, 539, 5, 55,
		0, 0, 539, 91, 1, 0, 0, 0, 540, 541, 5, 13, 0, 0, 541, 542, 5, 36, 0, 0,
		542, 545, 5, 15, 0, 0, 543, 546, 3, 100, 50, 0, 544, 546, 3, 96, 48, 0,
		545, 543, 1, 0, 0, 0, 545, 544, 1, 0, 0, 0, 546, 547, 1, 0, 0, 0, 547,
		548, 5, 54, 0, 0, 548, 549, 3, 2, 1, 0, 549, 550, 5, 55, 0, 0, 550, 93,
		1, 0, 0, 0, 551, 552, 5, 14, 0, 0, 552, 553, 3, 100, 50, 0, 553, 554, 5,
		11, 0, 0, 554, 555, 5, 54, 0, 0, 555, 556, 3, 2, 1, 0, 556, 557, 5, 55,
		0, 0, 557, 95, 1, 0, 0, 0, 558, 559, 3, 100, 50, 0, 559, 560, 5, 62, 0,
		0, 560, 561, 5, 62, 0, 0, 561, 562, 5, 62, 0, 0, 562, 563, 3, 100, 50,
		0, 563, 97, 1, 0, 0, 0, 564, 565, 5, 36, 0, 0, 565, 566, 5, 49, 0, 0, 566,
		567, 5, 42, 0, 0, 567, 573, 3, 100, 50, 0, 568, 569, 5, 36, 0, 0, 569,
		570, 5, 50, 0, 0, 570, 571, 5, 42, 0, 0, 571, 573, 3, 100, 50, 0, 572,
		564, 1, 0, 0, 0, 572, 568, 1, 0, 0, 0, 573, 99, 1, 0, 0, 0, 574, 575, 6,
		50, -1, 0, 575, 576, 5, 52, 0, 0, 576, 577, 3, 100, 50, 0, 577, 578, 5,
		53, 0, 0, 578, 598, 1, 0, 0, 0, 579, 580, 5, 50, 0, 0, 580, 598, 3, 100,
		50, 23, 581, 582, 5, 39, 0, 0, 582, 598, 3, 100, 50, 22, 583, 598, 5, 34,
		0, 0, 584, 598, 5, 36, 0, 0, 585, 598, 5, 35, 0, 0, 586, 598, 7, 2, 0,
		0, 587, 598, 5, 6, 0, 0, 588, 598, 3, 66, 33, 0, 589, 598, 3, 58, 29, 0,
		590, 598, 3, 60, 30, 0, 591, 598, 3, 62, 31, 0, 592, 598, 3, 24, 12, 0,
		593, 598, 3, 34, 17, 0, 594, 598, 3, 36, 18, 0, 595, 598, 3, 46, 23, 0,
		596, 598, 3, 16, 8, 0, 597, 574, 1, 0, 0, 0, 597, 579, 1, 0, 0, 0, 597,
		581, 1, 0, 0, 0, 597, 583, 1, 0, 0, 0, 597, 584, 1, 0, 0, 0, 597, 585,
		1, 0, 0, 0, 597, 586, 1, 0, 0, 0, 597, 587, 1, 0, 0, 0, 597, 588, 1, 0,
		0, 0, 597, 589, 1, 0, 0, 0, 597, 590, 1, 0, 0, 0, 597, 591, 1, 0, 0, 0,
		597, 592, 1, 0, 0, 0, 597, 593, 1, 0, 0, 0, 597, 594, 1, 0, 0, 0, 597,
		595, 1, 0, 0, 0, 597, 596, 1, 0, 0, 0, 598, 622, 1, 0, 0, 0, 599, 600,
		10, 21, 0, 0, 600, 601, 7, 3, 0, 0, 601, 621, 3, 100, 50, 22, 602, 603,
		10, 20, 0, 0, 603, 604, 7, 4, 0, 0, 604, 621, 3, 100, 50, 21, 605, 606,
		10, 19, 0, 0, 606, 607, 7, 5, 0, 0, 607, 621, 3, 100, 50, 20, 608, 609,
		10, 18, 0, 0, 609, 610, 7, 6, 0, 0, 610, 621, 3, 100, 50, 19, 611, 612,
		10, 17, 0, 0, 612, 613, 7, 7, 0, 0, 613, 621, 3, 100, 50, 18, 614, 615,
		10, 16, 0, 0, 615, 616, 5, 41, 0, 0, 616, 621, 3, 100, 50, 17, 617, 618,
		10, 15, 0, 0, 618, 619, 5, 40, 0, 0, 619, 621, 3, 100, 50, 16, 620, 599,
		1, 0, 0, 0, 620, 602, 1, 0, 0, 0, 620, 605, 1, 0, 0, 0, 620, 608, 1, 0,
		0, 0, 620, 611, 1, 0, 0, 0, 620, 614, 1, 0, 0, 0, 620, 617, 1, 0, 0, 0,
		621, 624, 1, 0, 0, 0, 622, 620, 1, 0, 0, 0, 622, 623, 1, 0, 0, 0, 623,
		101, 1, 0, 0, 0, 624, 622, 1, 0, 0, 0, 625, 626, 7, 8, 0, 0, 626, 103,
		1, 0, 0, 0, 37, 109, 113, 142, 150, 160, 162, 166, 172, 182, 193, 201,
		214, 219, 226, 275, 283, 287, 303, 316, 326, 350, 375, 409, 420, 430, 435,
		450, 462, 473, 481, 496, 499, 545, 572, 597, 620, 622,
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
	SwiftGrammarParserFOR           = 13
	SwiftGrammarParserGUARD         = 14
	SwiftGrammarParserIN            = 15
	SwiftGrammarParserSWITCH        = 16
	SwiftGrammarParserCASE          = 17
	SwiftGrammarParserDEFAULT       = 18
	SwiftGrammarParserVAR           = 19
	SwiftGrammarParserLET           = 20
	SwiftGrammarParserBREAK         = 21
	SwiftGrammarParserRETURN        = 22
	SwiftGrammarParserCONTINUE      = 23
	SwiftGrammarParserFUNC          = 24
	SwiftGrammarParserCOUNT         = 25
	SwiftGrammarParserISEMPTY       = 26
	SwiftGrammarParserAPPEND        = 27
	SwiftGrammarParserREMOVE_LAST   = 28
	SwiftGrammarParserREMOVE        = 29
	SwiftGrammarParserSTRUCT        = 30
	SwiftGrammarParserSTRUCT_VAR    = 31
	SwiftGrammarParserSTRUCT_LET    = 32
	SwiftGrammarParserAT            = 33
	SwiftGrammarParserNUMBER        = 34
	SwiftGrammarParserSTRING        = 35
	SwiftGrammarParserID            = 36
	SwiftGrammarParserDIF           = 37
	SwiftGrammarParserIG_IG         = 38
	SwiftGrammarParserNOT           = 39
	SwiftGrammarParserOR            = 40
	SwiftGrammarParserAND           = 41
	SwiftGrammarParserIG            = 42
	SwiftGrammarParserMAY_IG        = 43
	SwiftGrammarParserMEN_IG        = 44
	SwiftGrammarParserMAYOR         = 45
	SwiftGrammarParserMENOR         = 46
	SwiftGrammarParserMUL           = 47
	SwiftGrammarParserDIV           = 48
	SwiftGrammarParserADD           = 49
	SwiftGrammarParserSUB           = 50
	SwiftGrammarParserMOD           = 51
	SwiftGrammarParserPARIZQ        = 52
	SwiftGrammarParserPARDER        = 53
	SwiftGrammarParserLLAVEIZQ      = 54
	SwiftGrammarParserLLAVEDER      = 55
	SwiftGrammarParserCORCHIZQ      = 56
	SwiftGrammarParserCORCHDER      = 57
	SwiftGrammarParserDOSPUNTOS     = 58
	SwiftGrammarParserCOMA          = 59
	SwiftGrammarParserPTCOMA        = 60
	SwiftGrammarParserINTERROGACION = 61
	SwiftGrammarParserPUNTO         = 62
	SwiftGrammarParserWHITESPACE    = 63
	SwiftGrammarParserCOMMENT       = 64
	SwiftGrammarParserLINE_COMMENT  = 65
)

// SwiftGrammarParser rules.
const (
	SwiftGrammarParserRULE_s                     = 0
	SwiftGrammarParserRULE_block                 = 1
	SwiftGrammarParserRULE_stmt                  = 2
	SwiftGrammarParserRULE_defstructstmt         = 3
	SwiftGrammarParserRULE_lista_atributos       = 4
	SwiftGrammarParserRULE_struct_expr           = 5
	SwiftGrammarParserRULE_valor_struct_expr     = 6
	SwiftGrammarParserRULE_l_dupla               = 7
	SwiftGrammarParserRULE_accesostructstmt      = 8
	SwiftGrammarParserRULE_declvectorstmt        = 9
	SwiftGrammarParserRULE_defvectorstmt         = 10
	SwiftGrammarParserRULE_listaexpresiones      = 11
	SwiftGrammarParserRULE_accesovectorstmt      = 12
	SwiftGrammarParserRULE_asignvectorstmt       = 13
	SwiftGrammarParserRULE_appendvectorstmt      = 14
	SwiftGrammarParserRULE_removeatvectorstmt    = 15
	SwiftGrammarParserRULE_removelastvectorstmt  = 16
	SwiftGrammarParserRULE_countvectorstmt       = 17
	SwiftGrammarParserRULE_isemptyvectorstmt     = 18
	SwiftGrammarParserRULE_declmatrizstmt        = 19
	SwiftGrammarParserRULE_tipomatriz            = 20
	SwiftGrammarParserRULE_listavaloresmatriz    = 21
	SwiftGrammarParserRULE_listavaloresmatriz3   = 22
	SwiftGrammarParserRULE_accesomatriz          = 23
	SwiftGrammarParserRULE_asignmatrizstmt       = 24
	SwiftGrammarParserRULE_returnstmt            = 25
	SwiftGrammarParserRULE_breakstmt             = 26
	SwiftGrammarParserRULE_continuestmt          = 27
	SwiftGrammarParserRULE_printstmt             = 28
	SwiftGrammarParserRULE_intstmt               = 29
	SwiftGrammarParserRULE_floatstmt             = 30
	SwiftGrammarParserRULE_stringstmt            = 31
	SwiftGrammarParserRULE_funcdclstmt           = 32
	SwiftGrammarParserRULE_accfuncstm            = 33
	SwiftGrammarParserRULE_parametros            = 34
	SwiftGrammarParserRULE_parametroscall        = 35
	SwiftGrammarParserRULE_ifstmt                = 36
	SwiftGrammarParserRULE_elseifstmt            = 37
	SwiftGrammarParserRULE_switchstmt            = 38
	SwiftGrammarParserRULE_caseStmt              = 39
	SwiftGrammarParserRULE_defaultCase           = 40
	SwiftGrammarParserRULE_typedDeclstmt         = 41
	SwiftGrammarParserRULE_untypedDeclstmt       = 42
	SwiftGrammarParserRULE_optionalTypedDeclstmt = 43
	SwiftGrammarParserRULE_asignstmt             = 44
	SwiftGrammarParserRULE_whilestmt             = 45
	SwiftGrammarParserRULE_forstmt               = 46
	SwiftGrammarParserRULE_guardstmt             = 47
	SwiftGrammarParserRULE_rangostmt             = 48
	SwiftGrammarParserRULE_opasignstmt           = 49
	SwiftGrammarParserRULE_expr                  = 50
	SwiftGrammarParserRULE_tipo                  = 51
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
		p.SetState(104)
		p.Block()
	}
	{
		p.SetState(105)
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
	AllPTCOMA() []antlr.TerminalNode
	PTCOMA(i int) antlr.TerminalNode

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

func (s *BlockContext) AllPTCOMA() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserPTCOMA)
}

func (s *BlockContext) PTCOMA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPTCOMA, i)
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
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&76268795392) != 0 {
		{
			p.SetState(107)
			p.Stmt()
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTCOMA {
			{
				p.SetState(108)
				p.Match(SwiftGrammarParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

		p.SetState(115)
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
	Forstmt() IForstmtContext
	Guardstmt() IGuardstmtContext
	Opasignstmt() IOpasignstmtContext
	Funcdclstmt() IFuncdclstmtContext
	Accfuncstm() IAccfuncstmContext
	Breakstmt() IBreakstmtContext
	Continuestmt() IContinuestmtContext
	Returnstmt() IReturnstmtContext
	Declvectorstmt() IDeclvectorstmtContext
	Accesovectorstmt() IAccesovectorstmtContext
	Appendvectorstmt() IAppendvectorstmtContext
	Removelastvectorstmt() IRemovelastvectorstmtContext
	Removeatvectorstmt() IRemoveatvectorstmtContext
	Asignvectorstmt() IAsignvectorstmtContext
	Declmatrizstmt() IDeclmatrizstmtContext
	Asignmatrizstmt() IAsignmatrizstmtContext
	Defstructstmt() IDefstructstmtContext
	Struct_expr() IStruct_exprContext

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

func (s *StmtContext) Forstmt() IForstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForstmtContext)
}

func (s *StmtContext) Guardstmt() IGuardstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuardstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuardstmtContext)
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

func (s *StmtContext) Continuestmt() IContinuestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinuestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinuestmtContext)
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

func (s *StmtContext) Declvectorstmt() IDeclvectorstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclvectorstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclvectorstmtContext)
}

func (s *StmtContext) Accesovectorstmt() IAccesovectorstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccesovectorstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccesovectorstmtContext)
}

func (s *StmtContext) Appendvectorstmt() IAppendvectorstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAppendvectorstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAppendvectorstmtContext)
}

func (s *StmtContext) Removelastvectorstmt() IRemovelastvectorstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRemovelastvectorstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRemovelastvectorstmtContext)
}

func (s *StmtContext) Removeatvectorstmt() IRemoveatvectorstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRemoveatvectorstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRemoveatvectorstmtContext)
}

func (s *StmtContext) Asignvectorstmt() IAsignvectorstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignvectorstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignvectorstmtContext)
}

func (s *StmtContext) Declmatrizstmt() IDeclmatrizstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclmatrizstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclmatrizstmtContext)
}

func (s *StmtContext) Asignmatrizstmt() IAsignmatrizstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignmatrizstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignmatrizstmtContext)
}

func (s *StmtContext) Defstructstmt() IDefstructstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefstructstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefstructstmtContext)
}

func (s *StmtContext) Struct_expr() IStruct_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_exprContext)
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
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.Printstmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.TypedDeclstmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)
			p.UntypedDeclstmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(119)
			p.OptionalTypedDeclstmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(120)
			p.Asignstmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(121)
			p.Ifstmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(122)
			p.Switchstmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(123)
			p.Whilestmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(124)
			p.Forstmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(125)
			p.Guardstmt()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(126)
			p.Opasignstmt()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(127)
			p.Funcdclstmt()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(128)
			p.Accfuncstm()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(129)
			p.Breakstmt()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(130)
			p.Continuestmt()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(131)
			p.Returnstmt()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(132)
			p.Declvectorstmt()
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(133)
			p.Accesovectorstmt()
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(134)
			p.Appendvectorstmt()
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(135)
			p.Removelastvectorstmt()
		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(136)
			p.Removeatvectorstmt()
		}

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(137)
			p.Asignvectorstmt()
		}

	case 23:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(138)
			p.Declmatrizstmt()
		}

	case 24:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(139)
			p.Asignmatrizstmt()
		}

	case 25:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(140)
			p.Defstructstmt()
		}

	case 26:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(141)
			p.Struct_expr()
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

// IDefstructstmtContext is an interface to support dynamic dispatch.
type IDefstructstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDefstructstmtContext differentiates from other interfaces.
	IsDefstructstmtContext()
}

type DefstructstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefstructstmtContext() *DefstructstmtContext {
	var p = new(DefstructstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_defstructstmt
	return p
}

func InitEmptyDefstructstmtContext(p *DefstructstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_defstructstmt
}

func (*DefstructstmtContext) IsDefstructstmtContext() {}

func NewDefstructstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefstructstmtContext {
	var p = new(DefstructstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_defstructstmt

	return p
}

func (s *DefstructstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DefstructstmtContext) CopyAll(ctx *DefstructstmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DefstructstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefstructstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DefStructContext struct {
	DefstructstmtContext
}

func NewDefStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefStructContext {
	var p = new(DefStructContext)

	InitEmptyDefstructstmtContext(&p.DefstructstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*DefstructstmtContext))

	return p
}

func (s *DefStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefStructContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTRUCT, 0)
}

func (s *DefStructContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *DefStructContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *DefStructContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *DefStructContext) AllLista_atributos() []ILista_atributosContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILista_atributosContext); ok {
			len++
		}
	}

	tst := make([]ILista_atributosContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILista_atributosContext); ok {
			tst[i] = t.(ILista_atributosContext)
			i++
		}
	}

	return tst
}

func (s *DefStructContext) Lista_atributos(i int) ILista_atributosContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILista_atributosContext); ok {
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

	return t.(ILista_atributosContext)
}

func (s *DefStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDefStruct(s)
	}
}

func (s *DefStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDefStruct(s)
	}
}

func (s *DefStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitDefStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Defstructstmt() (localctx IDefstructstmtContext) {
	localctx = NewDefstructstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SwiftGrammarParserRULE_defstructstmt)
	var _la int

	localctx = NewDefStructContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(SwiftGrammarParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftGrammarParserVAR || _la == SwiftGrammarParserLET {
		{
			p.SetState(147)
			p.Lista_atributos()
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(153)
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

// ILista_atributosContext is an interface to support dynamic dispatch.
type ILista_atributosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLista_atributosContext differentiates from other interfaces.
	IsLista_atributosContext()
}

type Lista_atributosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLista_atributosContext() *Lista_atributosContext {
	var p = new(Lista_atributosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_lista_atributos
	return p
}

func InitEmptyLista_atributosContext(p *Lista_atributosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_lista_atributos
}

func (*Lista_atributosContext) IsLista_atributosContext() {}

func NewLista_atributosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lista_atributosContext {
	var p = new(Lista_atributosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_lista_atributos

	return p
}

func (s *Lista_atributosContext) GetParser() antlr.Parser { return s.parser }

func (s *Lista_atributosContext) CopyAll(ctx *Lista_atributosContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Lista_atributosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lista_atributosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AtributoStructContext struct {
	Lista_atributosContext
}

func NewAtributoStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtributoStructContext {
	var p = new(AtributoStructContext)

	InitEmptyLista_atributosContext(&p.Lista_atributosContext)
	p.parser = parser
	p.CopyAll(ctx.(*Lista_atributosContext))

	return p
}

func (s *AtributoStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtributoStructContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserID)
}

func (s *AtributoStructContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, i)
}

func (s *AtributoStructContext) LET() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLET, 0)
}

func (s *AtributoStructContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *AtributoStructContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPUNTOS, 0)
}

func (s *AtributoStructContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *AtributoStructContext) Expr() IExprContext {
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

func (s *AtributoStructContext) Tipo() ITipoContext {
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

func (s *AtributoStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAtributoStruct(s)
	}
}

func (s *AtributoStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAtributoStruct(s)
	}
}

func (s *AtributoStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitAtributoStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Lista_atributos() (localctx ILista_atributosContext) {
	localctx = NewLista_atributosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SwiftGrammarParserRULE_lista_atributos)
	var _la int

	localctx = NewAtributoStructContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SwiftGrammarParserVAR || _la == SwiftGrammarParserLET) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(156)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftGrammarParserDOSPUNTOS {
		{
			p.SetState(157)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SwiftGrammarParserINT, SwiftGrammarParserFLOAT, SwiftGrammarParserBOOL, SwiftGrammarParserCHARACTER, SwiftGrammarParserPSTRING:
			{
				p.SetState(158)
				p.Tipo()
			}

		case SwiftGrammarParserID:
			{
				p.SetState(159)
				p.Match(SwiftGrammarParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftGrammarParserIG {
		{
			p.SetState(164)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.expr(0)
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

// IStruct_exprContext is an interface to support dynamic dispatch.
type IStruct_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_exprContext differentiates from other interfaces.
	IsStruct_exprContext()
}

type Struct_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_exprContext() *Struct_exprContext {
	var p = new(Struct_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_struct_expr
	return p
}

func InitEmptyStruct_exprContext(p *Struct_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_struct_expr
}

func (*Struct_exprContext) IsStruct_exprContext() {}

func NewStruct_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_exprContext {
	var p = new(Struct_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_struct_expr

	return p
}

func (s *Struct_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_exprContext) CopyAll(ctx *Struct_exprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructExprContext struct {
	Struct_exprContext
}

func NewStructExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructExprContext {
	var p = new(StructExprContext)

	InitEmptyStruct_exprContext(&p.Struct_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_exprContext))

	return p
}

func (s *StructExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructExprContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserID)
}

func (s *StructExprContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, i)
}

func (s *StructExprContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *StructExprContext) Valor_struct_expr() IValor_struct_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValor_struct_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValor_struct_exprContext)
}

func (s *StructExprContext) STRUCT_VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTRUCT_VAR, 0)
}

func (s *StructExprContext) STRUCT_LET() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTRUCT_LET, 0)
}

func (s *StructExprContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPUNTOS, 0)
}

func (s *StructExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterStructExpr(s)
	}
}

func (s *StructExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitStructExpr(s)
	}
}

func (s *StructExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitStructExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Struct_expr() (localctx IStruct_exprContext) {
	localctx = NewStruct_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SwiftGrammarParserRULE_struct_expr)
	var _la int

	localctx = NewStructExprContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SwiftGrammarParserSTRUCT_VAR || _la == SwiftGrammarParserSTRUCT_LET) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(169)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftGrammarParserDOSPUNTOS {
		{
			p.SetState(170)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(174)
		p.Match(SwiftGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.Valor_struct_expr()
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

// IValor_struct_exprContext is an interface to support dynamic dispatch.
type IValor_struct_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValor_struct_exprContext differentiates from other interfaces.
	IsValor_struct_exprContext()
}

type Valor_struct_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValor_struct_exprContext() *Valor_struct_exprContext {
	var p = new(Valor_struct_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_valor_struct_expr
	return p
}

func InitEmptyValor_struct_exprContext(p *Valor_struct_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_valor_struct_expr
}

func (*Valor_struct_exprContext) IsValor_struct_exprContext() {}

func NewValor_struct_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Valor_struct_exprContext {
	var p = new(Valor_struct_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_valor_struct_expr

	return p
}

func (s *Valor_struct_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Valor_struct_exprContext) CopyAll(ctx *Valor_struct_exprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Valor_struct_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Valor_struct_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ValorStructExprContext struct {
	Valor_struct_exprContext
}

func NewValorStructExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValorStructExprContext {
	var p = new(ValorStructExprContext)

	InitEmptyValor_struct_exprContext(&p.Valor_struct_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Valor_struct_exprContext))

	return p
}

func (s *ValorStructExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorStructExprContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ValorStructExprContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *ValorStructExprContext) L_dupla() IL_duplaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_duplaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_duplaContext)
}

func (s *ValorStructExprContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *ValorStructExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterValorStructExpr(s)
	}
}

func (s *ValorStructExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitValorStructExpr(s)
	}
}

func (s *ValorStructExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitValorStructExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Valor_struct_expr() (localctx IValor_struct_exprContext) {
	localctx = NewValor_struct_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SwiftGrammarParserRULE_valor_struct_expr)
	var _la int

	localctx = NewValorStructExprContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftGrammarParserPARIZQ {
		{
			p.SetState(178)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)
			p.L_dupla()
		}
		{
			p.SetState(180)
			p.Match(SwiftGrammarParserPARDER)
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

// IL_duplaContext is an interface to support dynamic dispatch.
type IL_duplaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsL_duplaContext differentiates from other interfaces.
	IsL_duplaContext()
}

type L_duplaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyL_duplaContext() *L_duplaContext {
	var p = new(L_duplaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_l_dupla
	return p
}

func InitEmptyL_duplaContext(p *L_duplaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_l_dupla
}

func (*L_duplaContext) IsL_duplaContext() {}

func NewL_duplaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *L_duplaContext {
	var p = new(L_duplaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_l_dupla

	return p
}

func (s *L_duplaContext) GetParser() antlr.Parser { return s.parser }

func (s *L_duplaContext) CopyAll(ctx *L_duplaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *L_duplaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_duplaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DuplastructContext struct {
	L_duplaContext
}

func NewDuplastructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DuplastructContext {
	var p = new(DuplastructContext)

	InitEmptyL_duplaContext(&p.L_duplaContext)
	p.parser = parser
	p.CopyAll(ctx.(*L_duplaContext))

	return p
}

func (s *DuplastructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DuplastructContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserID)
}

func (s *DuplastructContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, i)
}

func (s *DuplastructContext) AllDOSPUNTOS() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserDOSPUNTOS)
}

func (s *DuplastructContext) DOSPUNTOS(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPUNTOS, i)
}

func (s *DuplastructContext) AllExpr() []IExprContext {
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

func (s *DuplastructContext) Expr(i int) IExprContext {
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

func (s *DuplastructContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCOMA)
}

func (s *DuplastructContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, i)
}

func (s *DuplastructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDuplastruct(s)
	}
}

func (s *DuplastructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDuplastruct(s)
	}
}

func (s *DuplastructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitDuplastruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) L_dupla() (localctx IL_duplaContext) {
	localctx = NewL_duplaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SwiftGrammarParserRULE_l_dupla)
	var _la int

	localctx = NewDuplastructContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
		p.Match(SwiftGrammarParserDOSPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.expr(0)
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftGrammarParserCOMA {
		{
			p.SetState(187)
			p.Match(SwiftGrammarParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.expr(0)
		}

		p.SetState(195)
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

// IAccesostructstmtContext is an interface to support dynamic dispatch.
type IAccesostructstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAccesostructstmtContext differentiates from other interfaces.
	IsAccesostructstmtContext()
}

type AccesostructstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccesostructstmtContext() *AccesostructstmtContext {
	var p = new(AccesostructstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_accesostructstmt
	return p
}

func InitEmptyAccesostructstmtContext(p *AccesostructstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_accesostructstmt
}

func (*AccesostructstmtContext) IsAccesostructstmtContext() {}

func NewAccesostructstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccesostructstmtContext {
	var p = new(AccesostructstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_accesostructstmt

	return p
}

func (s *AccesostructstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AccesostructstmtContext) CopyAll(ctx *AccesostructstmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AccesostructstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccesostructstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AccesoStructContext struct {
	AccesostructstmtContext
}

func NewAccesoStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccesoStructContext {
	var p = new(AccesoStructContext)

	InitEmptyAccesostructstmtContext(&p.AccesostructstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccesostructstmtContext))

	return p
}

func (s *AccesoStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccesoStructContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserID)
}

func (s *AccesoStructContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, i)
}

func (s *AccesoStructContext) AllPUNTO() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserPUNTO)
}

func (s *AccesoStructContext) PUNTO(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, i)
}

func (s *AccesoStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAccesoStruct(s)
	}
}

func (s *AccesoStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAccesoStruct(s)
	}
}

func (s *AccesoStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitAccesoStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Accesostructstmt() (localctx IAccesostructstmtContext) {
	localctx = NewAccesostructstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SwiftGrammarParserRULE_accesostructstmt)
	var _alt int

	localctx = NewAccesoStructContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(197)
				p.Match(SwiftGrammarParserPUNTO)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(198)
				p.Match(SwiftGrammarParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(201)
		p.GetErrorHandler().Sync(p)
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclvectorstmtContext is an interface to support dynamic dispatch.
type IDeclvectorstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	DOSPUNTOS() antlr.TerminalNode
	CORCHIZQ() antlr.TerminalNode
	Tipo() ITipoContext
	CORCHDER() antlr.TerminalNode
	Defvectorstmt() IDefvectorstmtContext

	// IsDeclvectorstmtContext differentiates from other interfaces.
	IsDeclvectorstmtContext()
}

type DeclvectorstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclvectorstmtContext() *DeclvectorstmtContext {
	var p = new(DeclvectorstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_declvectorstmt
	return p
}

func InitEmptyDeclvectorstmtContext(p *DeclvectorstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_declvectorstmt
}

func (*DeclvectorstmtContext) IsDeclvectorstmtContext() {}

func NewDeclvectorstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclvectorstmtContext {
	var p = new(DeclvectorstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_declvectorstmt

	return p
}

func (s *DeclvectorstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclvectorstmtContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *DeclvectorstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *DeclvectorstmtContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPUNTOS, 0)
}

func (s *DeclvectorstmtContext) CORCHIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHIZQ, 0)
}

func (s *DeclvectorstmtContext) Tipo() ITipoContext {
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

func (s *DeclvectorstmtContext) CORCHDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHDER, 0)
}

func (s *DeclvectorstmtContext) Defvectorstmt() IDefvectorstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefvectorstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefvectorstmtContext)
}

func (s *DeclvectorstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclvectorstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclvectorstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDeclvectorstmt(s)
	}
}

func (s *DeclvectorstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDeclvectorstmt(s)
	}
}

func (s *DeclvectorstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitDeclvectorstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Declvectorstmt() (localctx IDeclvectorstmtContext) {
	localctx = NewDeclvectorstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SwiftGrammarParserRULE_declvectorstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Match(SwiftGrammarParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Match(SwiftGrammarParserDOSPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Match(SwiftGrammarParserCORCHIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
		p.Tipo()
	}
	{
		p.SetState(208)
		p.Match(SwiftGrammarParserCORCHDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Defvectorstmt()
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

// IDefvectorstmtContext is an interface to support dynamic dispatch.
type IDefvectorstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDefvectorstmtContext differentiates from other interfaces.
	IsDefvectorstmtContext()
}

type DefvectorstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefvectorstmtContext() *DefvectorstmtContext {
	var p = new(DefvectorstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_defvectorstmt
	return p
}

func InitEmptyDefvectorstmtContext(p *DefvectorstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_defvectorstmt
}

func (*DefvectorstmtContext) IsDefvectorstmtContext() {}

func NewDefvectorstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefvectorstmtContext {
	var p = new(DefvectorstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_defvectorstmt

	return p
}

func (s *DefvectorstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DefvectorstmtContext) CopyAll(ctx *DefvectorstmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DefvectorstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefvectorstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DefVectorContext struct {
	DefvectorstmtContext
}

func NewDefVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefVectorContext {
	var p = new(DefVectorContext)

	InitEmptyDefvectorstmtContext(&p.DefvectorstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*DefvectorstmtContext))

	return p
}

func (s *DefVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefVectorContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *DefVectorContext) CORCHIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHIZQ, 0)
}

func (s *DefVectorContext) CORCHDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHDER, 0)
}

func (s *DefVectorContext) Listaexpresiones() IListaexpresionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaexpresionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaexpresionesContext)
}

func (s *DefVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDefVector(s)
	}
}

func (s *DefVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDefVector(s)
	}
}

func (s *DefVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitDefVector(s)

	default:
		return t.VisitChildren(s)
	}
}

type DefVectorIDContext struct {
	DefvectorstmtContext
}

func NewDefVectorIDContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefVectorIDContext {
	var p = new(DefVectorIDContext)

	InitEmptyDefvectorstmtContext(&p.DefvectorstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*DefvectorstmtContext))

	return p
}

func (s *DefVectorIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefVectorIDContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *DefVectorIDContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *DefVectorIDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDefVectorID(s)
	}
}

func (s *DefVectorIDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDefVectorID(s)
	}
}

func (s *DefVectorIDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitDefVectorID(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Defvectorstmt() (localctx IDefvectorstmtContext) {
	localctx = NewDefvectorstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SwiftGrammarParserRULE_defvectorstmt)
	var _la int

	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDefVectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(212)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&5630169549111782) != 0 {
			{
				p.SetState(213)
				p.Listaexpresiones()
			}

		}
		{
			p.SetState(216)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewDefVectorIDContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
			p.Match(SwiftGrammarParserID)
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

// IListaexpresionesContext is an interface to support dynamic dispatch.
type IListaexpresionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMA() []antlr.TerminalNode
	COMA(i int) antlr.TerminalNode

	// IsListaexpresionesContext differentiates from other interfaces.
	IsListaexpresionesContext()
}

type ListaexpresionesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaexpresionesContext() *ListaexpresionesContext {
	var p = new(ListaexpresionesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listaexpresiones
	return p
}

func InitEmptyListaexpresionesContext(p *ListaexpresionesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listaexpresiones
}

func (*ListaexpresionesContext) IsListaexpresionesContext() {}

func NewListaexpresionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaexpresionesContext {
	var p = new(ListaexpresionesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listaexpresiones

	return p
}

func (s *ListaexpresionesContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaexpresionesContext) AllExpr() []IExprContext {
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

func (s *ListaexpresionesContext) Expr(i int) IExprContext {
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

func (s *ListaexpresionesContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCOMA)
}

func (s *ListaexpresionesContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, i)
}

func (s *ListaexpresionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaexpresionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaexpresionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListaexpresiones(s)
	}
}

func (s *ListaexpresionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListaexpresiones(s)
	}
}

func (s *ListaexpresionesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitListaexpresiones(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Listaexpresiones() (localctx IListaexpresionesContext) {
	localctx = NewListaexpresionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SwiftGrammarParserRULE_listaexpresiones)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.expr(0)
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftGrammarParserCOMA {
		{
			p.SetState(222)
			p.Match(SwiftGrammarParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.expr(0)
		}

		p.SetState(228)
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

// IAccesovectorstmtContext is an interface to support dynamic dispatch.
type IAccesovectorstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	CORCHIZQ() antlr.TerminalNode
	Expr() IExprContext
	CORCHDER() antlr.TerminalNode

	// IsAccesovectorstmtContext differentiates from other interfaces.
	IsAccesovectorstmtContext()
}

type AccesovectorstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccesovectorstmtContext() *AccesovectorstmtContext {
	var p = new(AccesovectorstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_accesovectorstmt
	return p
}

func InitEmptyAccesovectorstmtContext(p *AccesovectorstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_accesovectorstmt
}

func (*AccesovectorstmtContext) IsAccesovectorstmtContext() {}

func NewAccesovectorstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccesovectorstmtContext {
	var p = new(AccesovectorstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_accesovectorstmt

	return p
}

func (s *AccesovectorstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AccesovectorstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *AccesovectorstmtContext) CORCHIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHIZQ, 0)
}

func (s *AccesovectorstmtContext) Expr() IExprContext {
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

func (s *AccesovectorstmtContext) CORCHDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHDER, 0)
}

func (s *AccesovectorstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccesovectorstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccesovectorstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAccesovectorstmt(s)
	}
}

func (s *AccesovectorstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAccesovectorstmt(s)
	}
}

func (s *AccesovectorstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitAccesovectorstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Accesovectorstmt() (localctx IAccesovectorstmtContext) {
	localctx = NewAccesovectorstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SwiftGrammarParserRULE_accesovectorstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.Match(SwiftGrammarParserCORCHIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)
		p.expr(0)
	}
	{
		p.SetState(232)
		p.Match(SwiftGrammarParserCORCHDER)
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

// IAsignvectorstmtContext is an interface to support dynamic dispatch.
type IAsignvectorstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	CORCHIZQ() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	CORCHDER() antlr.TerminalNode
	IG() antlr.TerminalNode

	// IsAsignvectorstmtContext differentiates from other interfaces.
	IsAsignvectorstmtContext()
}

type AsignvectorstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignvectorstmtContext() *AsignvectorstmtContext {
	var p = new(AsignvectorstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_asignvectorstmt
	return p
}

func InitEmptyAsignvectorstmtContext(p *AsignvectorstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_asignvectorstmt
}

func (*AsignvectorstmtContext) IsAsignvectorstmtContext() {}

func NewAsignvectorstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignvectorstmtContext {
	var p = new(AsignvectorstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_asignvectorstmt

	return p
}

func (s *AsignvectorstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignvectorstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *AsignvectorstmtContext) CORCHIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHIZQ, 0)
}

func (s *AsignvectorstmtContext) AllExpr() []IExprContext {
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

func (s *AsignvectorstmtContext) Expr(i int) IExprContext {
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

func (s *AsignvectorstmtContext) CORCHDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHDER, 0)
}

func (s *AsignvectorstmtContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *AsignvectorstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignvectorstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignvectorstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAsignvectorstmt(s)
	}
}

func (s *AsignvectorstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAsignvectorstmt(s)
	}
}

func (s *AsignvectorstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitAsignvectorstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Asignvectorstmt() (localctx IAsignvectorstmtContext) {
	localctx = NewAsignvectorstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SwiftGrammarParserRULE_asignvectorstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(235)
		p.Match(SwiftGrammarParserCORCHIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(236)
		p.expr(0)
	}
	{
		p.SetState(237)
		p.Match(SwiftGrammarParserCORCHDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(238)
		p.Match(SwiftGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(239)
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

// IAppendvectorstmtContext is an interface to support dynamic dispatch.
type IAppendvectorstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	PUNTO() antlr.TerminalNode
	APPEND() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode

	// IsAppendvectorstmtContext differentiates from other interfaces.
	IsAppendvectorstmtContext()
}

type AppendvectorstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAppendvectorstmtContext() *AppendvectorstmtContext {
	var p = new(AppendvectorstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_appendvectorstmt
	return p
}

func InitEmptyAppendvectorstmtContext(p *AppendvectorstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_appendvectorstmt
}

func (*AppendvectorstmtContext) IsAppendvectorstmtContext() {}

func NewAppendvectorstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AppendvectorstmtContext {
	var p = new(AppendvectorstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_appendvectorstmt

	return p
}

func (s *AppendvectorstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AppendvectorstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *AppendvectorstmtContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *AppendvectorstmtContext) APPEND() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAPPEND, 0)
}

func (s *AppendvectorstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *AppendvectorstmtContext) Expr() IExprContext {
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

func (s *AppendvectorstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *AppendvectorstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendvectorstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AppendvectorstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAppendvectorstmt(s)
	}
}

func (s *AppendvectorstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAppendvectorstmt(s)
	}
}

func (s *AppendvectorstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitAppendvectorstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Appendvectorstmt() (localctx IAppendvectorstmtContext) {
	localctx = NewAppendvectorstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SwiftGrammarParserRULE_appendvectorstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)
		p.Match(SwiftGrammarParserPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.Match(SwiftGrammarParserAPPEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.expr(0)
	}
	{
		p.SetState(246)
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

// IRemoveatvectorstmtContext is an interface to support dynamic dispatch.
type IRemoveatvectorstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	PUNTO() antlr.TerminalNode
	REMOVE() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	AT() antlr.TerminalNode
	DOSPUNTOS() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode

	// IsRemoveatvectorstmtContext differentiates from other interfaces.
	IsRemoveatvectorstmtContext()
}

type RemoveatvectorstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRemoveatvectorstmtContext() *RemoveatvectorstmtContext {
	var p = new(RemoveatvectorstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_removeatvectorstmt
	return p
}

func InitEmptyRemoveatvectorstmtContext(p *RemoveatvectorstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_removeatvectorstmt
}

func (*RemoveatvectorstmtContext) IsRemoveatvectorstmtContext() {}

func NewRemoveatvectorstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RemoveatvectorstmtContext {
	var p = new(RemoveatvectorstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_removeatvectorstmt

	return p
}

func (s *RemoveatvectorstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *RemoveatvectorstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *RemoveatvectorstmtContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *RemoveatvectorstmtContext) REMOVE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserREMOVE, 0)
}

func (s *RemoveatvectorstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *RemoveatvectorstmtContext) AT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAT, 0)
}

func (s *RemoveatvectorstmtContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPUNTOS, 0)
}

func (s *RemoveatvectorstmtContext) Expr() IExprContext {
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

func (s *RemoveatvectorstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *RemoveatvectorstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RemoveatvectorstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RemoveatvectorstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterRemoveatvectorstmt(s)
	}
}

func (s *RemoveatvectorstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitRemoveatvectorstmt(s)
	}
}

func (s *RemoveatvectorstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitRemoveatvectorstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Removeatvectorstmt() (localctx IRemoveatvectorstmtContext) {
	localctx = NewRemoveatvectorstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SwiftGrammarParserRULE_removeatvectorstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.Match(SwiftGrammarParserPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)
		p.Match(SwiftGrammarParserREMOVE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(251)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(252)
		p.Match(SwiftGrammarParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(253)
		p.Match(SwiftGrammarParserDOSPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(254)
		p.expr(0)
	}
	{
		p.SetState(255)
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

// IRemovelastvectorstmtContext is an interface to support dynamic dispatch.
type IRemovelastvectorstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	PUNTO() antlr.TerminalNode
	REMOVE_LAST() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode

	// IsRemovelastvectorstmtContext differentiates from other interfaces.
	IsRemovelastvectorstmtContext()
}

type RemovelastvectorstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRemovelastvectorstmtContext() *RemovelastvectorstmtContext {
	var p = new(RemovelastvectorstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_removelastvectorstmt
	return p
}

func InitEmptyRemovelastvectorstmtContext(p *RemovelastvectorstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_removelastvectorstmt
}

func (*RemovelastvectorstmtContext) IsRemovelastvectorstmtContext() {}

func NewRemovelastvectorstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RemovelastvectorstmtContext {
	var p = new(RemovelastvectorstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_removelastvectorstmt

	return p
}

func (s *RemovelastvectorstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *RemovelastvectorstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *RemovelastvectorstmtContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *RemovelastvectorstmtContext) REMOVE_LAST() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserREMOVE_LAST, 0)
}

func (s *RemovelastvectorstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *RemovelastvectorstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *RemovelastvectorstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RemovelastvectorstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RemovelastvectorstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterRemovelastvectorstmt(s)
	}
}

func (s *RemovelastvectorstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitRemovelastvectorstmt(s)
	}
}

func (s *RemovelastvectorstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitRemovelastvectorstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Removelastvectorstmt() (localctx IRemovelastvectorstmtContext) {
	localctx = NewRemovelastvectorstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SwiftGrammarParserRULE_removelastvectorstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.Match(SwiftGrammarParserPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(259)
		p.Match(SwiftGrammarParserREMOVE_LAST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
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

// ICountvectorstmtContext is an interface to support dynamic dispatch.
type ICountvectorstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	PUNTO() antlr.TerminalNode
	COUNT() antlr.TerminalNode

	// IsCountvectorstmtContext differentiates from other interfaces.
	IsCountvectorstmtContext()
}

type CountvectorstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCountvectorstmtContext() *CountvectorstmtContext {
	var p = new(CountvectorstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_countvectorstmt
	return p
}

func InitEmptyCountvectorstmtContext(p *CountvectorstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_countvectorstmt
}

func (*CountvectorstmtContext) IsCountvectorstmtContext() {}

func NewCountvectorstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CountvectorstmtContext {
	var p = new(CountvectorstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_countvectorstmt

	return p
}

func (s *CountvectorstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CountvectorstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *CountvectorstmtContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *CountvectorstmtContext) COUNT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOUNT, 0)
}

func (s *CountvectorstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountvectorstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CountvectorstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterCountvectorstmt(s)
	}
}

func (s *CountvectorstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitCountvectorstmt(s)
	}
}

func (s *CountvectorstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitCountvectorstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Countvectorstmt() (localctx ICountvectorstmtContext) {
	localctx = NewCountvectorstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SwiftGrammarParserRULE_countvectorstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(264)
		p.Match(SwiftGrammarParserPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(265)
		p.Match(SwiftGrammarParserCOUNT)
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

// IIsemptyvectorstmtContext is an interface to support dynamic dispatch.
type IIsemptyvectorstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	PUNTO() antlr.TerminalNode
	ISEMPTY() antlr.TerminalNode

	// IsIsemptyvectorstmtContext differentiates from other interfaces.
	IsIsemptyvectorstmtContext()
}

type IsemptyvectorstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIsemptyvectorstmtContext() *IsemptyvectorstmtContext {
	var p = new(IsemptyvectorstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_isemptyvectorstmt
	return p
}

func InitEmptyIsemptyvectorstmtContext(p *IsemptyvectorstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_isemptyvectorstmt
}

func (*IsemptyvectorstmtContext) IsIsemptyvectorstmtContext() {}

func NewIsemptyvectorstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsemptyvectorstmtContext {
	var p = new(IsemptyvectorstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_isemptyvectorstmt

	return p
}

func (s *IsemptyvectorstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IsemptyvectorstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *IsemptyvectorstmtContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *IsemptyvectorstmtContext) ISEMPTY() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserISEMPTY, 0)
}

func (s *IsemptyvectorstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsemptyvectorstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsemptyvectorstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterIsemptyvectorstmt(s)
	}
}

func (s *IsemptyvectorstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitIsemptyvectorstmt(s)
	}
}

func (s *IsemptyvectorstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitIsemptyvectorstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Isemptyvectorstmt() (localctx IIsemptyvectorstmtContext) {
	localctx = NewIsemptyvectorstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SwiftGrammarParserRULE_isemptyvectorstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
		p.Match(SwiftGrammarParserPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
		p.Match(SwiftGrammarParserISEMPTY)
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

// IDeclmatrizstmtContext is an interface to support dynamic dispatch.
type IDeclmatrizstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclmatrizstmtContext differentiates from other interfaces.
	IsDeclmatrizstmtContext()
}

type DeclmatrizstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclmatrizstmtContext() *DeclmatrizstmtContext {
	var p = new(DeclmatrizstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_declmatrizstmt
	return p
}

func InitEmptyDeclmatrizstmtContext(p *DeclmatrizstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_declmatrizstmt
}

func (*DeclmatrizstmtContext) IsDeclmatrizstmtContext() {}

func NewDeclmatrizstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclmatrizstmtContext {
	var p = new(DeclmatrizstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_declmatrizstmt

	return p
}

func (s *DeclmatrizstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclmatrizstmtContext) CopyAll(ctx *DeclmatrizstmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DeclmatrizstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclmatrizstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Declmatrizstmt2Context struct {
	DeclmatrizstmtContext
}

func NewDeclmatrizstmt2Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declmatrizstmt2Context {
	var p = new(Declmatrizstmt2Context)

	InitEmptyDeclmatrizstmtContext(&p.DeclmatrizstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclmatrizstmtContext))

	return p
}

func (s *Declmatrizstmt2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declmatrizstmt2Context) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *Declmatrizstmt2Context) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *Declmatrizstmt2Context) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *Declmatrizstmt2Context) Listavaloresmatriz() IListavaloresmatrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListavaloresmatrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListavaloresmatrizContext)
}

func (s *Declmatrizstmt2Context) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPUNTOS, 0)
}

func (s *Declmatrizstmt2Context) Tipomatriz() ITipomatrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipomatrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipomatrizContext)
}

func (s *Declmatrizstmt2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDeclmatrizstmt2(s)
	}
}

func (s *Declmatrizstmt2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDeclmatrizstmt2(s)
	}
}

func (s *Declmatrizstmt2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitDeclmatrizstmt2(s)

	default:
		return t.VisitChildren(s)
	}
}

type Declmatrizstmt3Context struct {
	DeclmatrizstmtContext
}

func NewDeclmatrizstmt3Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declmatrizstmt3Context {
	var p = new(Declmatrizstmt3Context)

	InitEmptyDeclmatrizstmtContext(&p.DeclmatrizstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclmatrizstmtContext))

	return p
}

func (s *Declmatrizstmt3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declmatrizstmt3Context) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *Declmatrizstmt3Context) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *Declmatrizstmt3Context) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *Declmatrizstmt3Context) Listavaloresmatriz3() IListavaloresmatriz3Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListavaloresmatriz3Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListavaloresmatriz3Context)
}

func (s *Declmatrizstmt3Context) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPUNTOS, 0)
}

func (s *Declmatrizstmt3Context) Tipomatriz() ITipomatrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipomatrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipomatrizContext)
}

func (s *Declmatrizstmt3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDeclmatrizstmt3(s)
	}
}

func (s *Declmatrizstmt3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDeclmatrizstmt3(s)
	}
}

func (s *Declmatrizstmt3Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitDeclmatrizstmt3(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Declmatrizstmt() (localctx IDeclmatrizstmtContext) {
	localctx = NewDeclmatrizstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SwiftGrammarParserRULE_declmatrizstmt)
	var _la int

	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclmatrizstmt2Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(271)
			p.Match(SwiftGrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(272)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserDOSPUNTOS {
			{
				p.SetState(273)
				p.Match(SwiftGrammarParserDOSPUNTOS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(274)
				p.Tipomatriz()
			}

		}
		{
			p.SetState(277)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(278)
			p.Listavaloresmatriz()
		}

	case 2:
		localctx = NewDeclmatrizstmt3Context(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(279)
			p.Match(SwiftGrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(280)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserDOSPUNTOS {
			{
				p.SetState(281)
				p.Match(SwiftGrammarParserDOSPUNTOS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(282)
				p.Tipomatriz()
			}

		}
		{
			p.SetState(285)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(286)
			p.Listavaloresmatriz3()
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

// ITipomatrizContext is an interface to support dynamic dispatch.
type ITipomatrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTipomatrizContext differentiates from other interfaces.
	IsTipomatrizContext()
}

type TipomatrizContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipomatrizContext() *TipomatrizContext {
	var p = new(TipomatrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_tipomatriz
	return p
}

func InitEmptyTipomatrizContext(p *TipomatrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_tipomatriz
}

func (*TipomatrizContext) IsTipomatrizContext() {}

func NewTipomatrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipomatrizContext {
	var p = new(TipomatrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_tipomatriz

	return p
}

func (s *TipomatrizContext) GetParser() antlr.Parser { return s.parser }

func (s *TipomatrizContext) CopyAll(ctx *TipomatrizContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TipomatrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipomatrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Tipomatriz2Context struct {
	TipomatrizContext
}

func NewTipomatriz2Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipomatriz2Context {
	var p = new(Tipomatriz2Context)

	InitEmptyTipomatrizContext(&p.TipomatrizContext)
	p.parser = parser
	p.CopyAll(ctx.(*TipomatrizContext))

	return p
}

func (s *Tipomatriz2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipomatriz2Context) AllCORCHIZQ() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORCHIZQ)
}

func (s *Tipomatriz2Context) CORCHIZQ(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHIZQ, i)
}

func (s *Tipomatriz2Context) Tipo() ITipoContext {
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

func (s *Tipomatriz2Context) AllCORCHDER() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORCHDER)
}

func (s *Tipomatriz2Context) CORCHDER(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHDER, i)
}

func (s *Tipomatriz2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterTipomatriz2(s)
	}
}

func (s *Tipomatriz2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitTipomatriz2(s)
	}
}

func (s *Tipomatriz2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitTipomatriz2(s)

	default:
		return t.VisitChildren(s)
	}
}

type Tipomatriz3Context struct {
	TipomatrizContext
}

func NewTipomatriz3Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipomatriz3Context {
	var p = new(Tipomatriz3Context)

	InitEmptyTipomatrizContext(&p.TipomatrizContext)
	p.parser = parser
	p.CopyAll(ctx.(*TipomatrizContext))

	return p
}

func (s *Tipomatriz3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipomatriz3Context) AllCORCHIZQ() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORCHIZQ)
}

func (s *Tipomatriz3Context) CORCHIZQ(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHIZQ, i)
}

func (s *Tipomatriz3Context) Tipo() ITipoContext {
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

func (s *Tipomatriz3Context) AllCORCHDER() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORCHDER)
}

func (s *Tipomatriz3Context) CORCHDER(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHDER, i)
}

func (s *Tipomatriz3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterTipomatriz3(s)
	}
}

func (s *Tipomatriz3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitTipomatriz3(s)
	}
}

func (s *Tipomatriz3Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitTipomatriz3(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Tipomatriz() (localctx ITipomatrizContext) {
	localctx = NewTipomatrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SwiftGrammarParserRULE_tipomatriz)
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTipomatriz2Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(289)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(290)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(291)
			p.Tipo()
		}
		{
			p.SetState(292)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(293)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewTipomatriz3Context(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(295)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(296)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(297)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(298)
			p.Tipo()
		}
		{
			p.SetState(299)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(300)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(301)
			p.Match(SwiftGrammarParserCORCHDER)
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

// IListavaloresmatrizContext is an interface to support dynamic dispatch.
type IListavaloresmatrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsListavaloresmatrizContext differentiates from other interfaces.
	IsListavaloresmatrizContext()
}

type ListavaloresmatrizContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListavaloresmatrizContext() *ListavaloresmatrizContext {
	var p = new(ListavaloresmatrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listavaloresmatriz
	return p
}

func InitEmptyListavaloresmatrizContext(p *ListavaloresmatrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listavaloresmatriz
}

func (*ListavaloresmatrizContext) IsListavaloresmatrizContext() {}

func NewListavaloresmatrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListavaloresmatrizContext {
	var p = new(ListavaloresmatrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listavaloresmatriz

	return p
}

func (s *ListavaloresmatrizContext) GetParser() antlr.Parser { return s.parser }

func (s *ListavaloresmatrizContext) CopyAll(ctx *ListavaloresmatrizContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ListavaloresmatrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListavaloresmatrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Listavaloresmatriz2Context struct {
	ListavaloresmatrizContext
}

func NewListavaloresmatriz2Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Listavaloresmatriz2Context {
	var p = new(Listavaloresmatriz2Context)

	InitEmptyListavaloresmatrizContext(&p.ListavaloresmatrizContext)
	p.parser = parser
	p.CopyAll(ctx.(*ListavaloresmatrizContext))

	return p
}

func (s *Listavaloresmatriz2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Listavaloresmatriz2Context) AllCORCHIZQ() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORCHIZQ)
}

func (s *Listavaloresmatriz2Context) CORCHIZQ(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHIZQ, i)
}

func (s *Listavaloresmatriz2Context) AllListaexpresiones() []IListaexpresionesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IListaexpresionesContext); ok {
			len++
		}
	}

	tst := make([]IListaexpresionesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IListaexpresionesContext); ok {
			tst[i] = t.(IListaexpresionesContext)
			i++
		}
	}

	return tst
}

func (s *Listavaloresmatriz2Context) Listaexpresiones(i int) IListaexpresionesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaexpresionesContext); ok {
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

	return t.(IListaexpresionesContext)
}

func (s *Listavaloresmatriz2Context) AllCORCHDER() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORCHDER)
}

func (s *Listavaloresmatriz2Context) CORCHDER(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHDER, i)
}

func (s *Listavaloresmatriz2Context) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCOMA)
}

func (s *Listavaloresmatriz2Context) COMA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, i)
}

func (s *Listavaloresmatriz2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListavaloresmatriz2(s)
	}
}

func (s *Listavaloresmatriz2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListavaloresmatriz2(s)
	}
}

func (s *Listavaloresmatriz2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitListavaloresmatriz2(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Listavaloresmatriz() (localctx IListavaloresmatrizContext) {
	localctx = NewListavaloresmatrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SwiftGrammarParserRULE_listavaloresmatriz)
	var _la int

	localctx = NewListavaloresmatriz2Context(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Match(SwiftGrammarParserCORCHIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
		p.Match(SwiftGrammarParserCORCHIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)
		p.Listaexpresiones()
	}
	{
		p.SetState(308)
		p.Match(SwiftGrammarParserCORCHDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SwiftGrammarParserCOMA {
		{
			p.SetState(309)
			p.Match(SwiftGrammarParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(310)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(311)
			p.Listaexpresiones()
		}
		{
			p.SetState(312)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(318)
		p.Match(SwiftGrammarParserCORCHDER)
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

// IListavaloresmatriz3Context is an interface to support dynamic dispatch.
type IListavaloresmatriz3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CORCHIZQ() antlr.TerminalNode
	AllListavaloresmatriz() []IListavaloresmatrizContext
	Listavaloresmatriz(i int) IListavaloresmatrizContext
	CORCHDER() antlr.TerminalNode
	AllCOMA() []antlr.TerminalNode
	COMA(i int) antlr.TerminalNode

	// IsListavaloresmatriz3Context differentiates from other interfaces.
	IsListavaloresmatriz3Context()
}

type Listavaloresmatriz3Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListavaloresmatriz3Context() *Listavaloresmatriz3Context {
	var p = new(Listavaloresmatriz3Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listavaloresmatriz3
	return p
}

func InitEmptyListavaloresmatriz3Context(p *Listavaloresmatriz3Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listavaloresmatriz3
}

func (*Listavaloresmatriz3Context) IsListavaloresmatriz3Context() {}

func NewListavaloresmatriz3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Listavaloresmatriz3Context {
	var p = new(Listavaloresmatriz3Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listavaloresmatriz3

	return p
}

func (s *Listavaloresmatriz3Context) GetParser() antlr.Parser { return s.parser }

func (s *Listavaloresmatriz3Context) CORCHIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHIZQ, 0)
}

func (s *Listavaloresmatriz3Context) AllListavaloresmatriz() []IListavaloresmatrizContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IListavaloresmatrizContext); ok {
			len++
		}
	}

	tst := make([]IListavaloresmatrizContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IListavaloresmatrizContext); ok {
			tst[i] = t.(IListavaloresmatrizContext)
			i++
		}
	}

	return tst
}

func (s *Listavaloresmatriz3Context) Listavaloresmatriz(i int) IListavaloresmatrizContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListavaloresmatrizContext); ok {
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

	return t.(IListavaloresmatrizContext)
}

func (s *Listavaloresmatriz3Context) CORCHDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHDER, 0)
}

func (s *Listavaloresmatriz3Context) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCOMA)
}

func (s *Listavaloresmatriz3Context) COMA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, i)
}

func (s *Listavaloresmatriz3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Listavaloresmatriz3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Listavaloresmatriz3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListavaloresmatriz3(s)
	}
}

func (s *Listavaloresmatriz3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListavaloresmatriz3(s)
	}
}

func (s *Listavaloresmatriz3Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitListavaloresmatriz3(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Listavaloresmatriz3() (localctx IListavaloresmatriz3Context) {
	localctx = NewListavaloresmatriz3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SwiftGrammarParserRULE_listavaloresmatriz3)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.Match(SwiftGrammarParserCORCHIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(321)
		p.Listavaloresmatriz()
	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftGrammarParserCOMA {
		{
			p.SetState(322)
			p.Match(SwiftGrammarParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
			p.Listavaloresmatriz()
		}

		p.SetState(328)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(329)
		p.Match(SwiftGrammarParserCORCHDER)
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

// IAccesomatrizContext is an interface to support dynamic dispatch.
type IAccesomatrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAccesomatrizContext differentiates from other interfaces.
	IsAccesomatrizContext()
}

type AccesomatrizContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccesomatrizContext() *AccesomatrizContext {
	var p = new(AccesomatrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_accesomatriz
	return p
}

func InitEmptyAccesomatrizContext(p *AccesomatrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_accesomatriz
}

func (*AccesomatrizContext) IsAccesomatrizContext() {}

func NewAccesomatrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccesomatrizContext {
	var p = new(AccesomatrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_accesomatriz

	return p
}

func (s *AccesomatrizContext) GetParser() antlr.Parser { return s.parser }

func (s *AccesomatrizContext) CopyAll(ctx *AccesomatrizContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AccesomatrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccesomatrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Accesomatriz2Context struct {
	AccesomatrizContext
}

func NewAccesomatriz2Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Accesomatriz2Context {
	var p = new(Accesomatriz2Context)

	InitEmptyAccesomatrizContext(&p.AccesomatrizContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccesomatrizContext))

	return p
}

func (s *Accesomatriz2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Accesomatriz2Context) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *Accesomatriz2Context) AllCORCHIZQ() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORCHIZQ)
}

func (s *Accesomatriz2Context) CORCHIZQ(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHIZQ, i)
}

func (s *Accesomatriz2Context) AllExpr() []IExprContext {
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

func (s *Accesomatriz2Context) Expr(i int) IExprContext {
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

func (s *Accesomatriz2Context) AllCORCHDER() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORCHDER)
}

func (s *Accesomatriz2Context) CORCHDER(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHDER, i)
}

func (s *Accesomatriz2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAccesomatriz2(s)
	}
}

func (s *Accesomatriz2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAccesomatriz2(s)
	}
}

func (s *Accesomatriz2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitAccesomatriz2(s)

	default:
		return t.VisitChildren(s)
	}
}

type Accesomatriz3Context struct {
	AccesomatrizContext
}

func NewAccesomatriz3Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Accesomatriz3Context {
	var p = new(Accesomatriz3Context)

	InitEmptyAccesomatrizContext(&p.AccesomatrizContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccesomatrizContext))

	return p
}

func (s *Accesomatriz3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Accesomatriz3Context) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *Accesomatriz3Context) AllCORCHIZQ() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORCHIZQ)
}

func (s *Accesomatriz3Context) CORCHIZQ(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHIZQ, i)
}

func (s *Accesomatriz3Context) AllExpr() []IExprContext {
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

func (s *Accesomatriz3Context) Expr(i int) IExprContext {
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

func (s *Accesomatriz3Context) AllCORCHDER() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORCHDER)
}

func (s *Accesomatriz3Context) CORCHDER(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHDER, i)
}

func (s *Accesomatriz3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAccesomatriz3(s)
	}
}

func (s *Accesomatriz3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAccesomatriz3(s)
	}
}

func (s *Accesomatriz3Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitAccesomatriz3(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Accesomatriz() (localctx IAccesomatrizContext) {
	localctx = NewAccesomatrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SwiftGrammarParserRULE_accesomatriz)
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAccesomatriz2Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(331)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.expr(0)
		}
		{
			p.SetState(334)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)
			p.expr(0)
		}
		{
			p.SetState(337)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewAccesomatriz3Context(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(339)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(340)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.expr(0)
		}
		{
			p.SetState(342)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(344)
			p.expr(0)
		}
		{
			p.SetState(345)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(346)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(347)
			p.expr(0)
		}
		{
			p.SetState(348)
			p.Match(SwiftGrammarParserCORCHDER)
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

// IAsignmatrizstmtContext is an interface to support dynamic dispatch.
type IAsignmatrizstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAsignmatrizstmtContext differentiates from other interfaces.
	IsAsignmatrizstmtContext()
}

type AsignmatrizstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignmatrizstmtContext() *AsignmatrizstmtContext {
	var p = new(AsignmatrizstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_asignmatrizstmt
	return p
}

func InitEmptyAsignmatrizstmtContext(p *AsignmatrizstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_asignmatrizstmt
}

func (*AsignmatrizstmtContext) IsAsignmatrizstmtContext() {}

func NewAsignmatrizstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignmatrizstmtContext {
	var p = new(AsignmatrizstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_asignmatrizstmt

	return p
}

func (s *AsignmatrizstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignmatrizstmtContext) CopyAll(ctx *AsignmatrizstmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AsignmatrizstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignmatrizstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Asignmatrizstmt3Context struct {
	AsignmatrizstmtContext
}

func NewAsignmatrizstmt3Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Asignmatrizstmt3Context {
	var p = new(Asignmatrizstmt3Context)

	InitEmptyAsignmatrizstmtContext(&p.AsignmatrizstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignmatrizstmtContext))

	return p
}

func (s *Asignmatrizstmt3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignmatrizstmt3Context) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *Asignmatrizstmt3Context) AllCORCHIZQ() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORCHIZQ)
}

func (s *Asignmatrizstmt3Context) CORCHIZQ(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHIZQ, i)
}

func (s *Asignmatrizstmt3Context) AllExpr() []IExprContext {
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

func (s *Asignmatrizstmt3Context) Expr(i int) IExprContext {
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

func (s *Asignmatrizstmt3Context) AllCORCHDER() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORCHDER)
}

func (s *Asignmatrizstmt3Context) CORCHDER(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHDER, i)
}

func (s *Asignmatrizstmt3Context) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *Asignmatrizstmt3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAsignmatrizstmt3(s)
	}
}

func (s *Asignmatrizstmt3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAsignmatrizstmt3(s)
	}
}

func (s *Asignmatrizstmt3Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitAsignmatrizstmt3(s)

	default:
		return t.VisitChildren(s)
	}
}

type Asignmatrizstmt2Context struct {
	AsignmatrizstmtContext
}

func NewAsignmatrizstmt2Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Asignmatrizstmt2Context {
	var p = new(Asignmatrizstmt2Context)

	InitEmptyAsignmatrizstmtContext(&p.AsignmatrizstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*AsignmatrizstmtContext))

	return p
}

func (s *Asignmatrizstmt2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignmatrizstmt2Context) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *Asignmatrizstmt2Context) AllCORCHIZQ() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORCHIZQ)
}

func (s *Asignmatrizstmt2Context) CORCHIZQ(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHIZQ, i)
}

func (s *Asignmatrizstmt2Context) AllExpr() []IExprContext {
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

func (s *Asignmatrizstmt2Context) Expr(i int) IExprContext {
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

func (s *Asignmatrizstmt2Context) AllCORCHDER() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORCHDER)
}

func (s *Asignmatrizstmt2Context) CORCHDER(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHDER, i)
}

func (s *Asignmatrizstmt2Context) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *Asignmatrizstmt2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAsignmatrizstmt2(s)
	}
}

func (s *Asignmatrizstmt2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAsignmatrizstmt2(s)
	}
}

func (s *Asignmatrizstmt2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitAsignmatrizstmt2(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Asignmatrizstmt() (localctx IAsignmatrizstmtContext) {
	localctx = NewAsignmatrizstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SwiftGrammarParserRULE_asignmatrizstmt)
	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAsignmatrizstmt2Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(352)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(353)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)
			p.expr(0)
		}
		{
			p.SetState(355)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.expr(0)
		}
		{
			p.SetState(358)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(359)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(360)
			p.expr(0)
		}

	case 2:
		localctx = NewAsignmatrizstmt3Context(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(362)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(363)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(364)
			p.expr(0)
		}
		{
			p.SetState(365)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(366)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(367)
			p.expr(0)
		}
		{
			p.SetState(368)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(369)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(370)
			p.expr(0)
		}
		{
			p.SetState(371)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(372)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)
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
	p.EnterRule(localctx, 50, SwiftGrammarParserRULE_returnstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(377)
		p.Match(SwiftGrammarParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(378)
		p.expr(0)
	}
	{
		p.SetState(379)
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
	p.EnterRule(localctx, 52, SwiftGrammarParserRULE_breakstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)
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

// IContinuestmtContext is an interface to support dynamic dispatch.
type IContinuestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsContinuestmtContext differentiates from other interfaces.
	IsContinuestmtContext()
}

type ContinuestmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinuestmtContext() *ContinuestmtContext {
	var p = new(ContinuestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_continuestmt
	return p
}

func InitEmptyContinuestmtContext(p *ContinuestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_continuestmt
}

func (*ContinuestmtContext) IsContinuestmtContext() {}

func NewContinuestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinuestmtContext {
	var p = new(ContinuestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_continuestmt

	return p
}

func (s *ContinuestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinuestmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCONTINUE, 0)
}

func (s *ContinuestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinuestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinuestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterContinuestmt(s)
	}
}

func (s *ContinuestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitContinuestmt(s)
	}
}

func (s *ContinuestmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitContinuestmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Continuestmt() (localctx IContinuestmtContext) {
	localctx = NewContinuestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SwiftGrammarParserRULE_continuestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(383)
		p.Match(SwiftGrammarParserCONTINUE)
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
	p.EnterRule(localctx, 56, SwiftGrammarParserRULE_printstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(385)
		p.Match(SwiftGrammarParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(386)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(387)
		p.expr(0)
	}
	{
		p.SetState(388)
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

// IIntstmtContext is an interface to support dynamic dispatch.
type IIntstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode

	// IsIntstmtContext differentiates from other interfaces.
	IsIntstmtContext()
}

type IntstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntstmtContext() *IntstmtContext {
	var p = new(IntstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_intstmt
	return p
}

func InitEmptyIntstmtContext(p *IntstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_intstmt
}

func (*IntstmtContext) IsIntstmtContext() {}

func NewIntstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntstmtContext {
	var p = new(IntstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_intstmt

	return p
}

func (s *IntstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IntstmtContext) INT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserINT, 0)
}

func (s *IntstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *IntstmtContext) Expr() IExprContext {
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

func (s *IntstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *IntstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterIntstmt(s)
	}
}

func (s *IntstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitIntstmt(s)
	}
}

func (s *IntstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitIntstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Intstmt() (localctx IIntstmtContext) {
	localctx = NewIntstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SwiftGrammarParserRULE_intstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.Match(SwiftGrammarParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(391)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(392)
		p.expr(0)
	}
	{
		p.SetState(393)
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

// IFloatstmtContext is an interface to support dynamic dispatch.
type IFloatstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FLOAT() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode

	// IsFloatstmtContext differentiates from other interfaces.
	IsFloatstmtContext()
}

type FloatstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatstmtContext() *FloatstmtContext {
	var p = new(FloatstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_floatstmt
	return p
}

func InitEmptyFloatstmtContext(p *FloatstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_floatstmt
}

func (*FloatstmtContext) IsFloatstmtContext() {}

func NewFloatstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatstmtContext {
	var p = new(FloatstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_floatstmt

	return p
}

func (s *FloatstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatstmtContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFLOAT, 0)
}

func (s *FloatstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *FloatstmtContext) Expr() IExprContext {
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

func (s *FloatstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *FloatstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterFloatstmt(s)
	}
}

func (s *FloatstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitFloatstmt(s)
	}
}

func (s *FloatstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitFloatstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Floatstmt() (localctx IFloatstmtContext) {
	localctx = NewFloatstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SwiftGrammarParserRULE_floatstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(395)
		p.Match(SwiftGrammarParserFLOAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(396)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(397)
		p.expr(0)
	}
	{
		p.SetState(398)
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

// IStringstmtContext is an interface to support dynamic dispatch.
type IStringstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PSTRING() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode

	// IsStringstmtContext differentiates from other interfaces.
	IsStringstmtContext()
}

type StringstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringstmtContext() *StringstmtContext {
	var p = new(StringstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_stringstmt
	return p
}

func InitEmptyStringstmtContext(p *StringstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_stringstmt
}

func (*StringstmtContext) IsStringstmtContext() {}

func NewStringstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringstmtContext {
	var p = new(StringstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_stringstmt

	return p
}

func (s *StringstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StringstmtContext) PSTRING() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPSTRING, 0)
}

func (s *StringstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *StringstmtContext) Expr() IExprContext {
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

func (s *StringstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *StringstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterStringstmt(s)
	}
}

func (s *StringstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitStringstmt(s)
	}
}

func (s *StringstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitStringstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Stringstmt() (localctx IStringstmtContext) {
	localctx = NewStringstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SwiftGrammarParserRULE_stringstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(400)
		p.Match(SwiftGrammarParserPSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(401)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(402)
		p.expr(0)
	}
	{
		p.SetState(403)
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

func (s *FuncionNormalContext) Parametros() IParametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametrosContext)
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

func (s *FuncionRetornoContext) Parametros() IParametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametrosContext)
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
	p.EnterRule(localctx, 64, SwiftGrammarParserRULE_funcdclstmt)
	var _la int

	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFuncionNormalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(405)
			p.Match(SwiftGrammarParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(406)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserID {
			{
				p.SetState(408)
				p.Parametros()
			}

		}
		{
			p.SetState(411)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(412)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(413)
			p.Block()
		}
		{
			p.SetState(414)
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
			p.SetState(416)
			p.Match(SwiftGrammarParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(417)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(418)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(420)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserID {
			{
				p.SetState(419)
				p.Parametros()
			}

		}
		{
			p.SetState(422)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(423)
			p.Match(SwiftGrammarParserSUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(424)
			p.Match(SwiftGrammarParserMAYOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(425)
			p.Tipo()
		}
		{
			p.SetState(426)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(427)
			p.Block()
		}
		{
			p.SetState(428)
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
	Parametroscall() IParametroscallContext

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

func (s *AccfuncstmContext) Parametroscall() IParametroscallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroscallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroscallContext)
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
	p.EnterRule(localctx, 66, SwiftGrammarParserRULE_accfuncstm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(432)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(433)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(435)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftGrammarParserID {
		{
			p.SetState(434)
			p.Parametroscall()
		}

	}
	{
		p.SetState(437)
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

// IParametrosContext is an interface to support dynamic dispatch.
type IParametrosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllDOSPUNTOS() []antlr.TerminalNode
	DOSPUNTOS(i int) antlr.TerminalNode
	AllTipo() []ITipoContext
	Tipo(i int) ITipoContext
	AllCOMA() []antlr.TerminalNode
	COMA(i int) antlr.TerminalNode

	// IsParametrosContext differentiates from other interfaces.
	IsParametrosContext()
}

type ParametrosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametrosContext() *ParametrosContext {
	var p = new(ParametrosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_parametros
	return p
}

func InitEmptyParametrosContext(p *ParametrosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_parametros
}

func (*ParametrosContext) IsParametrosContext() {}

func NewParametrosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosContext {
	var p = new(ParametrosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_parametros

	return p
}

func (s *ParametrosContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserID)
}

func (s *ParametrosContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, i)
}

func (s *ParametrosContext) AllDOSPUNTOS() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserDOSPUNTOS)
}

func (s *ParametrosContext) DOSPUNTOS(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPUNTOS, i)
}

func (s *ParametrosContext) AllTipo() []ITipoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITipoContext); ok {
			len++
		}
	}

	tst := make([]ITipoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITipoContext); ok {
			tst[i] = t.(ITipoContext)
			i++
		}
	}

	return tst
}

func (s *ParametrosContext) Tipo(i int) ITipoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
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

	return t.(ITipoContext)
}

func (s *ParametrosContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCOMA)
}

func (s *ParametrosContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, i)
}

func (s *ParametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterParametros(s)
	}
}

func (s *ParametrosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitParametros(s)
	}
}

func (s *ParametrosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitParametros(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Parametros() (localctx IParametrosContext) {
	localctx = NewParametrosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SwiftGrammarParserRULE_parametros)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(439)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(440)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(441)
		p.Match(SwiftGrammarParserDOSPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(442)
		p.Tipo()
	}
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftGrammarParserCOMA {
		{
			p.SetState(443)
			p.Match(SwiftGrammarParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(444)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(445)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(446)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(447)
			p.Tipo()
		}

		p.SetState(452)
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

// IParametroscallContext is an interface to support dynamic dispatch.
type IParametroscallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllDOSPUNTOS() []antlr.TerminalNode
	DOSPUNTOS(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMA() []antlr.TerminalNode
	COMA(i int) antlr.TerminalNode

	// IsParametroscallContext differentiates from other interfaces.
	IsParametroscallContext()
}

type ParametroscallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametroscallContext() *ParametroscallContext {
	var p = new(ParametroscallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_parametroscall
	return p
}

func InitEmptyParametroscallContext(p *ParametroscallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_parametroscall
}

func (*ParametroscallContext) IsParametroscallContext() {}

func NewParametroscallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametroscallContext {
	var p = new(ParametroscallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_parametroscall

	return p
}

func (s *ParametroscallContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametroscallContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserID)
}

func (s *ParametroscallContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, i)
}

func (s *ParametroscallContext) AllDOSPUNTOS() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserDOSPUNTOS)
}

func (s *ParametroscallContext) DOSPUNTOS(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPUNTOS, i)
}

func (s *ParametroscallContext) AllExpr() []IExprContext {
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

func (s *ParametroscallContext) Expr(i int) IExprContext {
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

func (s *ParametroscallContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCOMA)
}

func (s *ParametroscallContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, i)
}

func (s *ParametroscallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametroscallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametroscallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterParametroscall(s)
	}
}

func (s *ParametroscallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitParametroscall(s)
	}
}

func (s *ParametroscallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitParametroscall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Parametroscall() (localctx IParametroscallContext) {
	localctx = NewParametroscallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SwiftGrammarParserRULE_parametroscall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(453)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(454)
		p.Match(SwiftGrammarParserDOSPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(455)
		p.expr(0)
	}
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftGrammarParserCOMA {
		{
			p.SetState(456)
			p.Match(SwiftGrammarParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(457)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(458)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(459)
			p.expr(0)
		}

		p.SetState(464)
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
	p.EnterRule(localctx, 72, SwiftGrammarParserRULE_ifstmt)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(465)
		p.Match(SwiftGrammarParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(466)
		p.expr(0)
	}
	{
		p.SetState(467)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(468)
		p.Block()
	}
	{
		p.SetState(469)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(470)
				p.Elseifstmt()
			}

		}
		p.SetState(475)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftGrammarParserELSE {
		{
			p.SetState(476)
			p.Match(SwiftGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(477)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(478)
			p.Block()
		}
		{
			p.SetState(479)
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
	p.EnterRule(localctx, 74, SwiftGrammarParserRULE_elseifstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(483)
		p.Match(SwiftGrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(484)
		p.Match(SwiftGrammarParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(485)
		p.expr(0)
	}
	{
		p.SetState(486)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(487)
		p.Block()
	}
	{
		p.SetState(488)
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
	p.EnterRule(localctx, 76, SwiftGrammarParserRULE_switchstmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(490)
		p.Match(SwiftGrammarParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(491)
		p.expr(0)
	}
	{
		p.SetState(492)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(494)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SwiftGrammarParserCASE {
		{
			p.SetState(493)
			p.CaseStmt()
		}

		p.SetState(496)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(499)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftGrammarParserDEFAULT {
		{
			p.SetState(498)
			p.DefaultCase()
		}

	}
	{
		p.SetState(501)
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
	p.EnterRule(localctx, 78, SwiftGrammarParserRULE_caseStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(503)
		p.Match(SwiftGrammarParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(504)
		p.expr(0)
	}
	{
		p.SetState(505)
		p.Match(SwiftGrammarParserDOSPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(506)
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
	p.EnterRule(localctx, 80, SwiftGrammarParserRULE_defaultCase)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(508)
		p.Match(SwiftGrammarParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(509)
		p.Match(SwiftGrammarParserDOSPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(510)
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
	ID() antlr.TerminalNode
	DOSPUNTOS() antlr.TerminalNode
	Tipo() ITipoContext
	IG() antlr.TerminalNode
	Expr() IExprContext
	VAR() antlr.TerminalNode
	LET() antlr.TerminalNode

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

func (s *TypedDeclstmtContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *TypedDeclstmtContext) LET() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLET, 0)
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
	p.EnterRule(localctx, 82, SwiftGrammarParserRULE_typedDeclstmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(512)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SwiftGrammarParserVAR || _la == SwiftGrammarParserLET) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(513)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(514)
		p.Match(SwiftGrammarParserDOSPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(515)
		p.Tipo()
	}
	{
		p.SetState(516)
		p.Match(SwiftGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(517)
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
	ID() antlr.TerminalNode
	IG() antlr.TerminalNode
	Expr() IExprContext
	VAR() antlr.TerminalNode
	LET() antlr.TerminalNode

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

func (s *UntypedDeclstmtContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *UntypedDeclstmtContext) LET() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLET, 0)
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
	p.EnterRule(localctx, 84, SwiftGrammarParserRULE_untypedDeclstmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(519)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SwiftGrammarParserVAR || _la == SwiftGrammarParserLET) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(520)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(521)
		p.Match(SwiftGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(522)
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
	p.EnterRule(localctx, 86, SwiftGrammarParserRULE_optionalTypedDeclstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(524)
		p.Match(SwiftGrammarParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(525)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(526)
		p.Match(SwiftGrammarParserDOSPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(527)
		p.Tipo()
	}
	{
		p.SetState(528)
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
	p.EnterRule(localctx, 88, SwiftGrammarParserRULE_asignstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(530)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(531)
		p.Match(SwiftGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(532)
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
	p.EnterRule(localctx, 90, SwiftGrammarParserRULE_whilestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(534)
		p.Match(SwiftGrammarParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(535)
		p.expr(0)
	}
	{
		p.SetState(536)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(537)
		p.Block()
	}
	{
		p.SetState(538)
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

// IForstmtContext is an interface to support dynamic dispatch.
type IForstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	ID() antlr.TerminalNode
	IN() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode
	Expr() IExprContext
	Rangostmt() IRangostmtContext

	// IsForstmtContext differentiates from other interfaces.
	IsForstmtContext()
}

type ForstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForstmtContext() *ForstmtContext {
	var p = new(ForstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_forstmt
	return p
}

func InitEmptyForstmtContext(p *ForstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_forstmt
}

func (*ForstmtContext) IsForstmtContext() {}

func NewForstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForstmtContext {
	var p = new(ForstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_forstmt

	return p
}

func (s *ForstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForstmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFOR, 0)
}

func (s *ForstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ForstmtContext) IN() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIN, 0)
}

func (s *ForstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *ForstmtContext) Block() IBlockContext {
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

func (s *ForstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *ForstmtContext) Expr() IExprContext {
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

func (s *ForstmtContext) Rangostmt() IRangostmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangostmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangostmtContext)
}

func (s *ForstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterForstmt(s)
	}
}

func (s *ForstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitForstmt(s)
	}
}

func (s *ForstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitForstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Forstmt() (localctx IForstmtContext) {
	localctx = NewForstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, SwiftGrammarParserRULE_forstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(540)
		p.Match(SwiftGrammarParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(541)
		p.Match(SwiftGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(542)
		p.Match(SwiftGrammarParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(545)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(543)
			p.expr(0)
		}

	case 2:
		{
			p.SetState(544)
			p.Rangostmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(547)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(548)
		p.Block()
	}
	{
		p.SetState(549)
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

// IGuardstmtContext is an interface to support dynamic dispatch.
type IGuardstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GUARD() antlr.TerminalNode
	Expr() IExprContext
	ELSE() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsGuardstmtContext differentiates from other interfaces.
	IsGuardstmtContext()
}

type GuardstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGuardstmtContext() *GuardstmtContext {
	var p = new(GuardstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_guardstmt
	return p
}

func InitEmptyGuardstmtContext(p *GuardstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_guardstmt
}

func (*GuardstmtContext) IsGuardstmtContext() {}

func NewGuardstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GuardstmtContext {
	var p = new(GuardstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_guardstmt

	return p
}

func (s *GuardstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *GuardstmtContext) GUARD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserGUARD, 0)
}

func (s *GuardstmtContext) Expr() IExprContext {
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

func (s *GuardstmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
}

func (s *GuardstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *GuardstmtContext) Block() IBlockContext {
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

func (s *GuardstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *GuardstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GuardstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterGuardstmt(s)
	}
}

func (s *GuardstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitGuardstmt(s)
	}
}

func (s *GuardstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitGuardstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Guardstmt() (localctx IGuardstmtContext) {
	localctx = NewGuardstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, SwiftGrammarParserRULE_guardstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(551)
		p.Match(SwiftGrammarParserGUARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(552)
		p.expr(0)
	}
	{
		p.SetState(553)
		p.Match(SwiftGrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(554)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(555)
		p.Block()
	}
	{
		p.SetState(556)
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

// IRangostmtContext is an interface to support dynamic dispatch.
type IRangostmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllPUNTO() []antlr.TerminalNode
	PUNTO(i int) antlr.TerminalNode

	// IsRangostmtContext differentiates from other interfaces.
	IsRangostmtContext()
}

type RangostmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangostmtContext() *RangostmtContext {
	var p = new(RangostmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_rangostmt
	return p
}

func InitEmptyRangostmtContext(p *RangostmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_rangostmt
}

func (*RangostmtContext) IsRangostmtContext() {}

func NewRangostmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangostmtContext {
	var p = new(RangostmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_rangostmt

	return p
}

func (s *RangostmtContext) GetParser() antlr.Parser { return s.parser }

func (s *RangostmtContext) AllExpr() []IExprContext {
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

func (s *RangostmtContext) Expr(i int) IExprContext {
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

func (s *RangostmtContext) AllPUNTO() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserPUNTO)
}

func (s *RangostmtContext) PUNTO(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, i)
}

func (s *RangostmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangostmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangostmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterRangostmt(s)
	}
}

func (s *RangostmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitRangostmt(s)
	}
}

func (s *RangostmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitRangostmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftGrammarParser) Rangostmt() (localctx IRangostmtContext) {
	localctx = NewRangostmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, SwiftGrammarParserRULE_rangostmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(558)
		p.expr(0)
	}
	{
		p.SetState(559)
		p.Match(SwiftGrammarParserPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(560)
		p.Match(SwiftGrammarParserPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(561)
		p.Match(SwiftGrammarParserPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(562)
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
	p.EnterRule(localctx, 98, SwiftGrammarParserRULE_opasignstmt)
	p.SetState(572)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIncrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(564)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(565)
			p.Match(SwiftGrammarParserADD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(566)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(567)
			p.expr(0)
		}

	case 2:
		localctx = NewDecrementoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(568)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(569)
			p.Match(SwiftGrammarParserSUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(570)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(571)
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

type StringExprContext struct {
	ExprContext
}

func NewStringExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringExprContext {
	var p = new(StringExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StringExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringExprContext) Stringstmt() IStringstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringstmtContext)
}

func (s *StringExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterStringExpr(s)
	}
}

func (s *StringExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitStringExpr(s)
	}
}

func (s *StringExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitStringExpr(s)

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

type FloatExprContext struct {
	ExprContext
}

func NewFloatExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatExprContext {
	var p = new(FloatExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FloatExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatExprContext) Floatstmt() IFloatstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatstmtContext)
}

func (s *FloatExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterFloatExpr(s)
	}
}

func (s *FloatExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitFloatExpr(s)
	}
}

func (s *FloatExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitFloatExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsEmptyVectorExprContext struct {
	ExprContext
}

func NewIsEmptyVectorExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsEmptyVectorExprContext {
	var p = new(IsEmptyVectorExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IsEmptyVectorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsEmptyVectorExprContext) Isemptyvectorstmt() IIsemptyvectorstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIsemptyvectorstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIsemptyvectorstmtContext)
}

func (s *IsEmptyVectorExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterIsEmptyVectorExpr(s)
	}
}

func (s *IsEmptyVectorExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitIsEmptyVectorExpr(s)
	}
}

func (s *IsEmptyVectorExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitIsEmptyVectorExpr(s)

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

func (s *OpExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDIV, 0)
}

func (s *OpExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMOD, 0)
}

func (s *OpExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMUL, 0)
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

type AccesoVectorExprContext struct {
	ExprContext
}

func NewAccesoVectorExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccesoVectorExprContext {
	var p = new(AccesoVectorExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AccesoVectorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccesoVectorExprContext) Accesovectorstmt() IAccesovectorstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccesovectorstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccesovectorstmtContext)
}

func (s *AccesoVectorExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAccesoVectorExpr(s)
	}
}

func (s *AccesoVectorExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAccesoVectorExpr(s)
	}
}

func (s *AccesoVectorExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitAccesoVectorExpr(s)

	default:
		return t.VisitChildren(s)
	}
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

type AccesoMatrizExprContext struct {
	ExprContext
}

func NewAccesoMatrizExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccesoMatrizExprContext {
	var p = new(AccesoMatrizExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AccesoMatrizExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccesoMatrizExprContext) Accesomatriz() IAccesomatrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccesomatrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccesomatrizContext)
}

func (s *AccesoMatrizExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAccesoMatrizExpr(s)
	}
}

func (s *AccesoMatrizExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAccesoMatrizExpr(s)
	}
}

func (s *AccesoMatrizExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitAccesoMatrizExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CountVectorExprContext struct {
	ExprContext
}

func NewCountVectorExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CountVectorExprContext {
	var p = new(CountVectorExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CountVectorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountVectorExprContext) Countvectorstmt() ICountvectorstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICountvectorstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICountvectorstmtContext)
}

func (s *CountVectorExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterCountVectorExpr(s)
	}
}

func (s *CountVectorExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitCountVectorExpr(s)
	}
}

func (s *CountVectorExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitCountVectorExpr(s)

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

type NotExprContext struct {
	ExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNOT, 0)
}

func (s *NotExprContext) Expr() IExprContext {
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

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntExprContext struct {
	ExprContext
}

func NewIntExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntExprContext {
	var p = new(IntExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IntExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntExprContext) Intstmt() IIntstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntstmtContext)
}

func (s *IntExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterIntExpr(s)
	}
}

func (s *IntExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitIntExpr(s)
	}
}

func (s *IntExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitIntExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AccesoStructExprContext struct {
	ExprContext
}

func NewAccesoStructExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccesoStructExprContext {
	var p = new(AccesoStructExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AccesoStructExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccesoStructExprContext) Accesostructstmt() IAccesostructstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccesostructstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccesostructstmtContext)
}

func (s *AccesoStructExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAccesoStructExpr(s)
	}
}

func (s *AccesoStructExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAccesoStructExpr(s)
	}
}

func (s *AccesoStructExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftGrammarVisitor:
		return t.VisitAccesoStructExpr(s)

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
	_startState := 100
	p.EnterRecursionRule(localctx, 100, SwiftGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(597)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(575)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(576)
			p.expr(0)
		}
		{
			p.SetState(577)
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
			p.SetState(579)
			p.Match(SwiftGrammarParserSUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(580)
			p.expr(23)
		}

	case 3:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(581)
			p.Match(SwiftGrammarParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(582)
			p.expr(22)
		}

	case 4:
		localctx = NewNumExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(583)
			p.Match(SwiftGrammarParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(584)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewStrExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(585)
			p.Match(SwiftGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(586)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SwiftGrammarParserTRU || _la == SwiftGrammarParserFAL) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 8:
		localctx = NewNilExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(587)
			p.Match(SwiftGrammarParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewAccFuncExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(588)
			p.Accfuncstm()
		}

	case 10:
		localctx = NewIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(589)
			p.Intstmt()
		}

	case 11:
		localctx = NewFloatExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(590)
			p.Floatstmt()
		}

	case 12:
		localctx = NewStringExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(591)
			p.Stringstmt()
		}

	case 13:
		localctx = NewAccesoVectorExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(592)
			p.Accesovectorstmt()
		}

	case 14:
		localctx = NewCountVectorExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(593)
			p.Countvectorstmt()
		}

	case 15:
		localctx = NewIsEmptyVectorExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(594)
			p.Isemptyvectorstmt()
		}

	case 16:
		localctx = NewAccesoMatrizExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(595)
			p.Accesomatriz()
		}

	case 17:
		localctx = NewAccesoStructExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(596)
			p.Accesostructstmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(622)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(620)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
			case 1:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(599)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(600)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2674012278751232) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(601)

					var _x = p.expr(22)

					localctx.(*OpExprContext).right = _x
				}

			case 2:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(602)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(603)

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
					p.SetState(604)

					var _x = p.expr(21)

					localctx.(*OpExprContext).right = _x
				}

			case 3:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(605)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(606)

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
					p.SetState(607)

					var _x = p.expr(20)

					localctx.(*OpExprContext).right = _x
				}

			case 4:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(608)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(609)

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
					p.SetState(610)

					var _x = p.expr(19)

					localctx.(*OpExprContext).right = _x
				}

			case 5:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(611)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(612)

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
					p.SetState(613)

					var _x = p.expr(18)

					localctx.(*OpExprContext).right = _x
				}

			case 6:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(614)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(615)

					var _m = p.Match(SwiftGrammarParserAND)

					localctx.(*OpExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(616)

					var _x = p.expr(17)

					localctx.(*OpExprContext).right = _x
				}

			case 7:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(617)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(618)

					var _m = p.Match(SwiftGrammarParserOR)

					localctx.(*OpExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(619)

					var _x = p.expr(16)

					localctx.(*OpExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(624)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 102, SwiftGrammarParserRULE_tipo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(625)
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
	case 50:
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
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 15)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

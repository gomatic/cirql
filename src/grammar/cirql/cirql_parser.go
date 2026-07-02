// Code generated from cirql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package cirqlgrammar // cirql
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

type cirqlParser struct {
	*antlr.BaseParser
}

var CirqlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cirqlParserInit() {
	staticData := &CirqlParserStaticData
	staticData.LiteralNames = []string{
		"", "'query'", "'http'", "'file'", "'stdin'", "'map'", "'flatMap'",
		"'filter'", "'reduce'", "'sort'", "'limit'", "'uniq'", "'count'", "'sum'",
		"'min'", "'max'", "'avg'", "'first'", "'last'", "'group_by'", "'collect'",
		"'asc'", "'desc'", "'true'", "'false'", "'null'", "'|'", "'.'", "','",
		"':'", "'$'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'||'", "'&&'",
		"'!'", "'=='", "'!='", "'>='", "'<='", "'>'", "'<'", "'+'", "'-'", "'*'",
		"'/'", "'%'",
	}
	staticData.SymbolicNames = []string{
		"", "QUERY", "HTTP", "FILE", "STDIN", "MAP", "FLATMAP", "FILTER", "REDUCE",
		"SORT", "LIMIT", "UNIQ", "COUNT", "SUM", "MIN", "MAX", "AVG", "FIRST",
		"LAST", "GROUP_BY", "COLLECT", "ASC", "DESC", "TRUE", "FALSE", "NULL",
		"PIPE", "DOT", "COMMA", "COLON", "DOLLAR", "LPAREN", "RPAREN", "LBRACE",
		"RBRACE", "LBRACK", "RBRACK", "OR", "AND", "NOT", "EQ", "NE", "GE",
		"LE", "GT", "LT", "PLUS", "MINUS", "STAR", "SLASH", "PERCENT", "NAME",
		"FLOAT", "INT", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"pipeline", "stage", "sourceStage", "queryStage", "queryBody", "selectionSet",
		"field", "arguments", "argument", "httpStage", "fileStage", "stdinStage",
		"transformStage", "mapStage", "flatMapStage", "mapping", "filterStage",
		"reduceStage", "reduceOp", "sortStage", "limitStage", "uniqStage", "expr",
		"fieldAccess", "pathSeg", "funcCall", "variable", "argValue", "objectLit",
		"listLit", "literal",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 55, 301, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 0, 1,
		0, 1, 0, 5, 0, 66, 8, 0, 10, 0, 12, 0, 69, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		3, 1, 75, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 81, 8, 2, 1, 3, 1, 3, 1,
		3, 3, 3, 86, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 4, 5, 93, 8, 5, 11, 5,
		12, 5, 94, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 102, 8, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 3, 6, 108, 8, 6, 1, 7, 1, 7, 1, 7, 5, 7, 113, 8, 7, 10, 7, 12,
		7, 116, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 3, 9, 125, 8, 9,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 3, 12, 139, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13,
		146, 8, 13, 10, 13, 12, 13, 149, 9, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 5, 14, 158, 8, 14, 10, 14, 12, 14, 161, 9, 14, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 3, 17, 178, 8, 17, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 19, 3, 19, 185, 8, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 3, 21, 192,
		8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 3, 22, 205, 8, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 222,
		8, 22, 10, 22, 12, 22, 225, 9, 22, 1, 23, 1, 23, 1, 23, 5, 23, 230, 8,
		23, 10, 23, 12, 23, 233, 9, 23, 1, 23, 3, 23, 236, 8, 23, 1, 24, 3, 24,
		239, 8, 24, 1, 24, 1, 24, 1, 24, 3, 24, 244, 8, 24, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 5, 25, 251, 8, 25, 10, 25, 12, 25, 254, 9, 25, 3, 25, 256,
		8, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 3,
		27, 267, 8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		5, 28, 277, 8, 28, 10, 28, 12, 28, 280, 9, 28, 3, 28, 282, 8, 28, 1, 28,
		1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 290, 8, 29, 10, 29, 12, 29, 293,
		9, 29, 3, 29, 295, 8, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 0, 1, 44,
		31, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 0, 7, 1, 0, 12, 20,
		1, 0, 21, 22, 2, 0, 39, 39, 47, 47, 1, 0, 48, 50, 1, 0, 46, 47, 1, 0, 40,
		45, 2, 0, 23, 25, 52, 54, 314, 0, 62, 1, 0, 0, 0, 2, 74, 1, 0, 0, 0, 4,
		80, 1, 0, 0, 0, 6, 85, 1, 0, 0, 0, 8, 87, 1, 0, 0, 0, 10, 92, 1, 0, 0,
		0, 12, 96, 1, 0, 0, 0, 14, 109, 1, 0, 0, 0, 16, 117, 1, 0, 0, 0, 18, 121,
		1, 0, 0, 0, 20, 126, 1, 0, 0, 0, 22, 129, 1, 0, 0, 0, 24, 138, 1, 0, 0,
		0, 26, 140, 1, 0, 0, 0, 28, 152, 1, 0, 0, 0, 30, 164, 1, 0, 0, 0, 32, 168,
		1, 0, 0, 0, 34, 171, 1, 0, 0, 0, 36, 179, 1, 0, 0, 0, 38, 181, 1, 0, 0,
		0, 40, 186, 1, 0, 0, 0, 42, 189, 1, 0, 0, 0, 44, 204, 1, 0, 0, 0, 46, 235,
		1, 0, 0, 0, 48, 243, 1, 0, 0, 0, 50, 245, 1, 0, 0, 0, 52, 259, 1, 0, 0,
		0, 54, 266, 1, 0, 0, 0, 56, 268, 1, 0, 0, 0, 58, 285, 1, 0, 0, 0, 60, 298,
		1, 0, 0, 0, 62, 67, 3, 2, 1, 0, 63, 64, 5, 26, 0, 0, 64, 66, 3, 2, 1, 0,
		65, 63, 1, 0, 0, 0, 66, 69, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1,
		0, 0, 0, 68, 70, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 70, 71, 5, 0, 0, 1, 71,
		1, 1, 0, 0, 0, 72, 75, 3, 4, 2, 0, 73, 75, 3, 24, 12, 0, 74, 72, 1, 0,
		0, 0, 74, 73, 1, 0, 0, 0, 75, 3, 1, 0, 0, 0, 76, 81, 3, 6, 3, 0, 77, 81,
		3, 18, 9, 0, 78, 81, 3, 20, 10, 0, 79, 81, 3, 22, 11, 0, 80, 76, 1, 0,
		0, 0, 80, 77, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 79, 1, 0, 0, 0, 81, 5,
		1, 0, 0, 0, 82, 83, 5, 1, 0, 0, 83, 86, 3, 8, 4, 0, 84, 86, 3, 8, 4, 0,
		85, 82, 1, 0, 0, 0, 85, 84, 1, 0, 0, 0, 86, 7, 1, 0, 0, 0, 87, 88, 5, 33,
		0, 0, 88, 89, 3, 10, 5, 0, 89, 90, 5, 34, 0, 0, 90, 9, 1, 0, 0, 0, 91,
		93, 3, 12, 6, 0, 92, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 92, 1, 0,
		0, 0, 94, 95, 1, 0, 0, 0, 95, 11, 1, 0, 0, 0, 96, 101, 5, 51, 0, 0, 97,
		98, 5, 31, 0, 0, 98, 99, 3, 14, 7, 0, 99, 100, 5, 32, 0, 0, 100, 102, 1,
		0, 0, 0, 101, 97, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 107, 1, 0, 0,
		0, 103, 104, 5, 33, 0, 0, 104, 105, 3, 10, 5, 0, 105, 106, 5, 34, 0, 0,
		106, 108, 1, 0, 0, 0, 107, 103, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108,
		13, 1, 0, 0, 0, 109, 114, 3, 16, 8, 0, 110, 111, 5, 28, 0, 0, 111, 113,
		3, 16, 8, 0, 112, 110, 1, 0, 0, 0, 113, 116, 1, 0, 0, 0, 114, 112, 1, 0,
		0, 0, 114, 115, 1, 0, 0, 0, 115, 15, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0,
		117, 118, 5, 51, 0, 0, 118, 119, 5, 29, 0, 0, 119, 120, 3, 54, 27, 0, 120,
		17, 1, 0, 0, 0, 121, 124, 5, 2, 0, 0, 122, 125, 5, 54, 0, 0, 123, 125,
		3, 52, 26, 0, 124, 122, 1, 0, 0, 0, 124, 123, 1, 0, 0, 0, 125, 19, 1, 0,
		0, 0, 126, 127, 5, 3, 0, 0, 127, 128, 5, 54, 0, 0, 128, 21, 1, 0, 0, 0,
		129, 130, 5, 4, 0, 0, 130, 23, 1, 0, 0, 0, 131, 139, 3, 26, 13, 0, 132,
		139, 3, 32, 16, 0, 133, 139, 3, 34, 17, 0, 134, 139, 3, 38, 19, 0, 135,
		139, 3, 28, 14, 0, 136, 139, 3, 40, 20, 0, 137, 139, 3, 42, 21, 0, 138,
		131, 1, 0, 0, 0, 138, 132, 1, 0, 0, 0, 138, 133, 1, 0, 0, 0, 138, 134,
		1, 0, 0, 0, 138, 135, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 137, 1, 0,
		0, 0, 139, 25, 1, 0, 0, 0, 140, 141, 5, 5, 0, 0, 141, 142, 5, 33, 0, 0,
		142, 147, 3, 30, 15, 0, 143, 144, 5, 28, 0, 0, 144, 146, 3, 30, 15, 0,
		145, 143, 1, 0, 0, 0, 146, 149, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147,
		148, 1, 0, 0, 0, 148, 150, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 150, 151,
		5, 34, 0, 0, 151, 27, 1, 0, 0, 0, 152, 153, 5, 6, 0, 0, 153, 154, 5, 33,
		0, 0, 154, 159, 3, 30, 15, 0, 155, 156, 5, 28, 0, 0, 156, 158, 3, 30, 15,
		0, 157, 155, 1, 0, 0, 0, 158, 161, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159,
		160, 1, 0, 0, 0, 160, 162, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 162, 163,
		5, 34, 0, 0, 163, 29, 1, 0, 0, 0, 164, 165, 5, 51, 0, 0, 165, 166, 5, 29,
		0, 0, 166, 167, 3, 44, 22, 0, 167, 31, 1, 0, 0, 0, 168, 169, 5, 7, 0, 0,
		169, 170, 3, 44, 22, 0, 170, 33, 1, 0, 0, 0, 171, 172, 5, 8, 0, 0, 172,
		177, 3, 36, 18, 0, 173, 174, 5, 31, 0, 0, 174, 175, 3, 44, 22, 0, 175,
		176, 5, 32, 0, 0, 176, 178, 1, 0, 0, 0, 177, 173, 1, 0, 0, 0, 177, 178,
		1, 0, 0, 0, 178, 35, 1, 0, 0, 0, 179, 180, 7, 0, 0, 0, 180, 37, 1, 0, 0,
		0, 181, 182, 5, 9, 0, 0, 182, 184, 3, 44, 22, 0, 183, 185, 7, 1, 0, 0,
		184, 183, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 39, 1, 0, 0, 0, 186, 187,
		5, 10, 0, 0, 187, 188, 5, 53, 0, 0, 188, 41, 1, 0, 0, 0, 189, 191, 5, 11,
		0, 0, 190, 192, 3, 44, 22, 0, 191, 190, 1, 0, 0, 0, 191, 192, 1, 0, 0,
		0, 192, 43, 1, 0, 0, 0, 193, 194, 6, 22, -1, 0, 194, 195, 7, 2, 0, 0, 195,
		205, 3, 44, 22, 11, 196, 197, 5, 31, 0, 0, 197, 198, 3, 44, 22, 0, 198,
		199, 5, 32, 0, 0, 199, 205, 1, 0, 0, 0, 200, 205, 3, 50, 25, 0, 201, 205,
		3, 46, 23, 0, 202, 205, 3, 52, 26, 0, 203, 205, 3, 60, 30, 0, 204, 193,
		1, 0, 0, 0, 204, 196, 1, 0, 0, 0, 204, 200, 1, 0, 0, 0, 204, 201, 1, 0,
		0, 0, 204, 202, 1, 0, 0, 0, 204, 203, 1, 0, 0, 0, 205, 223, 1, 0, 0, 0,
		206, 207, 10, 10, 0, 0, 207, 208, 7, 3, 0, 0, 208, 222, 3, 44, 22, 11,
		209, 210, 10, 9, 0, 0, 210, 211, 7, 4, 0, 0, 211, 222, 3, 44, 22, 10, 212,
		213, 10, 8, 0, 0, 213, 214, 7, 5, 0, 0, 214, 222, 3, 44, 22, 9, 215, 216,
		10, 7, 0, 0, 216, 217, 5, 38, 0, 0, 217, 222, 3, 44, 22, 8, 218, 219, 10,
		6, 0, 0, 219, 220, 5, 37, 0, 0, 220, 222, 3, 44, 22, 7, 221, 206, 1, 0,
		0, 0, 221, 209, 1, 0, 0, 0, 221, 212, 1, 0, 0, 0, 221, 215, 1, 0, 0, 0,
		221, 218, 1, 0, 0, 0, 222, 225, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223,
		224, 1, 0, 0, 0, 224, 45, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 226, 227, 5,
		27, 0, 0, 227, 231, 3, 48, 24, 0, 228, 230, 3, 48, 24, 0, 229, 228, 1,
		0, 0, 0, 230, 233, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 231, 232, 1, 0, 0,
		0, 232, 236, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 234, 236, 5, 27, 0, 0, 235,
		226, 1, 0, 0, 0, 235, 234, 1, 0, 0, 0, 236, 47, 1, 0, 0, 0, 237, 239, 5,
		27, 0, 0, 238, 237, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 240, 1, 0, 0,
		0, 240, 244, 5, 51, 0, 0, 241, 242, 5, 35, 0, 0, 242, 244, 5, 36, 0, 0,
		243, 238, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 244, 49, 1, 0, 0, 0, 245, 246,
		5, 51, 0, 0, 246, 255, 5, 31, 0, 0, 247, 252, 3, 44, 22, 0, 248, 249, 5,
		28, 0, 0, 249, 251, 3, 44, 22, 0, 250, 248, 1, 0, 0, 0, 251, 254, 1, 0,
		0, 0, 252, 250, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 256, 1, 0, 0, 0,
		254, 252, 1, 0, 0, 0, 255, 247, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256,
		257, 1, 0, 0, 0, 257, 258, 5, 32, 0, 0, 258, 51, 1, 0, 0, 0, 259, 260,
		5, 30, 0, 0, 260, 261, 5, 51, 0, 0, 261, 53, 1, 0, 0, 0, 262, 267, 3, 52,
		26, 0, 263, 267, 3, 60, 30, 0, 264, 267, 3, 58, 29, 0, 265, 267, 3, 56,
		28, 0, 266, 262, 1, 0, 0, 0, 266, 263, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0,
		266, 265, 1, 0, 0, 0, 267, 55, 1, 0, 0, 0, 268, 281, 5, 33, 0, 0, 269,
		270, 5, 51, 0, 0, 270, 271, 5, 29, 0, 0, 271, 278, 3, 54, 27, 0, 272, 273,
		5, 28, 0, 0, 273, 274, 5, 51, 0, 0, 274, 275, 5, 29, 0, 0, 275, 277, 3,
		54, 27, 0, 276, 272, 1, 0, 0, 0, 277, 280, 1, 0, 0, 0, 278, 276, 1, 0,
		0, 0, 278, 279, 1, 0, 0, 0, 279, 282, 1, 0, 0, 0, 280, 278, 1, 0, 0, 0,
		281, 269, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283,
		284, 5, 34, 0, 0, 284, 57, 1, 0, 0, 0, 285, 294, 5, 35, 0, 0, 286, 291,
		3, 54, 27, 0, 287, 288, 5, 28, 0, 0, 288, 290, 3, 54, 27, 0, 289, 287,
		1, 0, 0, 0, 290, 293, 1, 0, 0, 0, 291, 289, 1, 0, 0, 0, 291, 292, 1, 0,
		0, 0, 292, 295, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 294, 286, 1, 0, 0, 0,
		294, 295, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 297, 5, 36, 0, 0, 297,
		59, 1, 0, 0, 0, 298, 299, 7, 6, 0, 0, 299, 61, 1, 0, 0, 0, 29, 67, 74,
		80, 85, 94, 101, 107, 114, 124, 138, 147, 159, 177, 184, 191, 204, 221,
		223, 231, 235, 238, 243, 252, 255, 266, 278, 281, 291, 294,
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

// cirqlParserInit initializes any static state used to implement cirqlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewcirqlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CirqlParserInit() {
	staticData := &CirqlParserStaticData
	staticData.once.Do(cirqlParserInit)
}

// NewcirqlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewcirqlParser(input antlr.TokenStream) *cirqlParser {
	CirqlParserInit()
	this := new(cirqlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CirqlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "cirql.g4"

	return this
}

// cirqlParser tokens.
const (
	cirqlParserEOF      = antlr.TokenEOF
	cirqlParserQUERY    = 1
	cirqlParserHTTP     = 2
	cirqlParserFILE     = 3
	cirqlParserSTDIN    = 4
	cirqlParserMAP      = 5
	cirqlParserFLATMAP  = 6
	cirqlParserFILTER   = 7
	cirqlParserREDUCE   = 8
	cirqlParserSORT     = 9
	cirqlParserLIMIT    = 10
	cirqlParserUNIQ     = 11
	cirqlParserCOUNT    = 12
	cirqlParserSUM      = 13
	cirqlParserMIN      = 14
	cirqlParserMAX      = 15
	cirqlParserAVG      = 16
	cirqlParserFIRST    = 17
	cirqlParserLAST     = 18
	cirqlParserGROUP_BY = 19
	cirqlParserCOLLECT  = 20
	cirqlParserASC      = 21
	cirqlParserDESC     = 22
	cirqlParserTRUE     = 23
	cirqlParserFALSE    = 24
	cirqlParserNULL     = 25
	cirqlParserPIPE     = 26
	cirqlParserDOT      = 27
	cirqlParserCOMMA    = 28
	cirqlParserCOLON    = 29
	cirqlParserDOLLAR   = 30
	cirqlParserLPAREN   = 31
	cirqlParserRPAREN   = 32
	cirqlParserLBRACE   = 33
	cirqlParserRBRACE   = 34
	cirqlParserLBRACK   = 35
	cirqlParserRBRACK   = 36
	cirqlParserOR       = 37
	cirqlParserAND      = 38
	cirqlParserNOT      = 39
	cirqlParserEQ       = 40
	cirqlParserNE       = 41
	cirqlParserGE       = 42
	cirqlParserLE       = 43
	cirqlParserGT       = 44
	cirqlParserLT       = 45
	cirqlParserPLUS     = 46
	cirqlParserMINUS    = 47
	cirqlParserSTAR     = 48
	cirqlParserSLASH    = 49
	cirqlParserPERCENT  = 50
	cirqlParserNAME     = 51
	cirqlParserFLOAT    = 52
	cirqlParserINT      = 53
	cirqlParserSTRING   = 54
	cirqlParserWS       = 55
)

// cirqlParser rules.
const (
	cirqlParserRULE_pipeline       = 0
	cirqlParserRULE_stage          = 1
	cirqlParserRULE_sourceStage    = 2
	cirqlParserRULE_queryStage     = 3
	cirqlParserRULE_queryBody      = 4
	cirqlParserRULE_selectionSet   = 5
	cirqlParserRULE_field          = 6
	cirqlParserRULE_arguments      = 7
	cirqlParserRULE_argument       = 8
	cirqlParserRULE_httpStage      = 9
	cirqlParserRULE_fileStage      = 10
	cirqlParserRULE_stdinStage     = 11
	cirqlParserRULE_transformStage = 12
	cirqlParserRULE_mapStage       = 13
	cirqlParserRULE_flatMapStage   = 14
	cirqlParserRULE_mapping        = 15
	cirqlParserRULE_filterStage    = 16
	cirqlParserRULE_reduceStage    = 17
	cirqlParserRULE_reduceOp       = 18
	cirqlParserRULE_sortStage      = 19
	cirqlParserRULE_limitStage     = 20
	cirqlParserRULE_uniqStage      = 21
	cirqlParserRULE_expr           = 22
	cirqlParserRULE_fieldAccess    = 23
	cirqlParserRULE_pathSeg        = 24
	cirqlParserRULE_funcCall       = 25
	cirqlParserRULE_variable       = 26
	cirqlParserRULE_argValue       = 27
	cirqlParserRULE_objectLit      = 28
	cirqlParserRULE_listLit        = 29
	cirqlParserRULE_literal        = 30
)

// IPipelineContext is an interface to support dynamic dispatch.
type IPipelineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStage() []IStageContext
	Stage(i int) IStageContext
	EOF() antlr.TerminalNode
	AllPIPE() []antlr.TerminalNode
	PIPE(i int) antlr.TerminalNode

	// IsPipelineContext differentiates from other interfaces.
	IsPipelineContext()
}

type PipelineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipelineContext() *PipelineContext {
	var p = new(PipelineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_pipeline
	return p
}

func InitEmptyPipelineContext(p *PipelineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_pipeline
}

func (PipelineContext) IsPipelineContext() {}

func NewPipelineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipelineContext {
	var p = new(PipelineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_pipeline

	return p
}

func (s PipelineContext) GetParser() antlr.Parser { return s.parser }

func (s *PipelineContext) AllStage() []IStageContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStageContext); ok {
			len++
		}
	}

	tst := make([]IStageContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStageContext); ok {
			tst[i] = t.(IStageContext)
			i++
		}
	}

	return tst
}

func (s *PipelineContext) Stage(i int) IStageContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStageContext); ok {
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

	return t.(IStageContext)
}

func (s *PipelineContext) EOF() antlr.TerminalNode {
	return s.GetToken(cirqlParserEOF, 0)
}

func (s *PipelineContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(cirqlParserPIPE)
}

func (s *PipelineContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(cirqlParserPIPE, i)
}

func (s *PipelineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipelineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipelineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterPipeline(s)
	}
}

func (s *PipelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitPipeline(s)
	}
}

func (p *cirqlParser) Pipeline() (localctx IPipelineContext) {
	localctx = NewPipelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, cirqlParserRULE_pipeline)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Stage()
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)
	for _la == cirqlParserPIPE {
		{
			p.SetState(63)
			p.Match(cirqlParserPIPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.Stage()
		}

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(70)
		p.Match(cirqlParserEOF)
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
}

// IStageContext is an interface to support dynamic dispatch.
type IStageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SourceStage() ISourceStageContext
	TransformStage() ITransformStageContext

	// IsStageContext differentiates from other interfaces.
	IsStageContext()
}

type StageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStageContext() *StageContext {
	var p = new(StageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_stage
	return p
}

func InitEmptyStageContext(p *StageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_stage
}

func (StageContext) IsStageContext() {}

func NewStageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StageContext {
	var p = new(StageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_stage

	return p
}

func (s StageContext) GetParser() antlr.Parser { return s.parser }

func (s *StageContext) SourceStage() ISourceStageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISourceStageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISourceStageContext)
}

func (s *StageContext) TransformStage() ITransformStageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransformStageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransformStageContext)
}

func (s *StageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterStage(s)
	}
}

func (s *StageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitStage(s)
	}
}

func (p *cirqlParser) Stage() (localctx IStageContext) {
	localctx = NewStageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, cirqlParserRULE_stage)
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case cirqlParserQUERY, cirqlParserHTTP, cirqlParserFILE, cirqlParserSTDIN, cirqlParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.SourceStage()
		}

	case cirqlParserMAP, cirqlParserFLATMAP, cirqlParserFILTER, cirqlParserREDUCE, cirqlParserSORT, cirqlParserLIMIT, cirqlParserUNIQ:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.TransformStage()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
}

// ISourceStageContext is an interface to support dynamic dispatch.
type ISourceStageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QueryStage() IQueryStageContext
	HttpStage() IHttpStageContext
	FileStage() IFileStageContext
	StdinStage() IStdinStageContext

	// IsSourceStageContext differentiates from other interfaces.
	IsSourceStageContext()
}

type SourceStageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceStageContext() *SourceStageContext {
	var p = new(SourceStageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_sourceStage
	return p
}

func InitEmptySourceStageContext(p *SourceStageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_sourceStage
}

func (SourceStageContext) IsSourceStageContext() {}

func NewSourceStageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceStageContext {
	var p = new(SourceStageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_sourceStage

	return p
}

func (s SourceStageContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceStageContext) QueryStage() IQueryStageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryStageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQueryStageContext)
}

func (s *SourceStageContext) HttpStage() IHttpStageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHttpStageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHttpStageContext)
}

func (s *SourceStageContext) FileStage() IFileStageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFileStageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFileStageContext)
}

func (s *SourceStageContext) StdinStage() IStdinStageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStdinStageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStdinStageContext)
}

func (s *SourceStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceStageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterSourceStage(s)
	}
}

func (s *SourceStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitSourceStage(s)
	}
}

func (p *cirqlParser) SourceStage() (localctx ISourceStageContext) {
	localctx = NewSourceStageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, cirqlParserRULE_sourceStage)
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case cirqlParserQUERY, cirqlParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.QueryStage()
		}

	case cirqlParserHTTP:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.HttpStage()
		}

	case cirqlParserFILE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(78)
			p.FileStage()
		}

	case cirqlParserSTDIN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(79)
			p.StdinStage()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
}

// IQueryStageContext is an interface to support dynamic dispatch.
type IQueryStageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QUERY() antlr.TerminalNode
	QueryBody() IQueryBodyContext

	// IsQueryStageContext differentiates from other interfaces.
	IsQueryStageContext()
}

type QueryStageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryStageContext() *QueryStageContext {
	var p = new(QueryStageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_queryStage
	return p
}

func InitEmptyQueryStageContext(p *QueryStageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_queryStage
}

func (QueryStageContext) IsQueryStageContext() {}

func NewQueryStageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryStageContext {
	var p = new(QueryStageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_queryStage

	return p
}

func (s QueryStageContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryStageContext) QUERY() antlr.TerminalNode {
	return s.GetToken(cirqlParserQUERY, 0)
}

func (s *QueryStageContext) QueryBody() IQueryBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQueryBodyContext)
}

func (s *QueryStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryStageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterQueryStage(s)
	}
}

func (s *QueryStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitQueryStage(s)
	}
}

func (p *cirqlParser) QueryStage() (localctx IQueryStageContext) {
	localctx = NewQueryStageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, cirqlParserRULE_queryStage)
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case cirqlParserQUERY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.Match(cirqlParserQUERY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.QueryBody()
		}

	case cirqlParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)
			p.QueryBody()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
}

// IQueryBodyContext is an interface to support dynamic dispatch.
type IQueryBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	SelectionSet() ISelectionSetContext
	RBRACE() antlr.TerminalNode

	// IsQueryBodyContext differentiates from other interfaces.
	IsQueryBodyContext()
}

type QueryBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryBodyContext() *QueryBodyContext {
	var p = new(QueryBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_queryBody
	return p
}

func InitEmptyQueryBodyContext(p *QueryBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_queryBody
}

func (QueryBodyContext) IsQueryBodyContext() {}

func NewQueryBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryBodyContext {
	var p = new(QueryBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_queryBody

	return p
}

func (s QueryBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryBodyContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(cirqlParserLBRACE, 0)
}

func (s *QueryBodyContext) SelectionSet() ISelectionSetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectionSetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *QueryBodyContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(cirqlParserRBRACE, 0)
}

func (s *QueryBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterQueryBody(s)
	}
}

func (s *QueryBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitQueryBody(s)
	}
}

func (p *cirqlParser) QueryBody() (localctx IQueryBodyContext) {
	localctx = NewQueryBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, cirqlParserRULE_queryBody)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.Match(cirqlParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)
		p.SelectionSet()
	}
	{
		p.SetState(89)
		p.Match(cirqlParserRBRACE)
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
}

// ISelectionSetContext is an interface to support dynamic dispatch.
type ISelectionSetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllField() []IFieldContext
	Field(i int) IFieldContext

	// IsSelectionSetContext differentiates from other interfaces.
	IsSelectionSetContext()
}

type SelectionSetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionSetContext() *SelectionSetContext {
	var p = new(SelectionSetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_selectionSet
	return p
}

func InitEmptySelectionSetContext(p *SelectionSetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_selectionSet
}

func (SelectionSetContext) IsSelectionSetContext() {}

func NewSelectionSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionSetContext {
	var p = new(SelectionSetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_selectionSet

	return p
}

func (s SelectionSetContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionSetContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *SelectionSetContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
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

	return t.(IFieldContext)
}

func (s *SelectionSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionSetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionSetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterSelectionSet(s)
	}
}

func (s *SelectionSetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitSelectionSet(s)
	}
}

func (p *cirqlParser) SelectionSet() (localctx ISelectionSetContext) {
	localctx = NewSelectionSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, cirqlParserRULE_selectionSet)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	for ok := true; ok; ok = _la == cirqlParserNAME {
		{
			p.SetState(91)
			p.Field()
		}

		p.SetState(94)
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
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Arguments() IArgumentsContext
	RPAREN() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	SelectionSet() ISelectionSetContext
	RBRACE() antlr.TerminalNode

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_field
}

func (FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_field

	return p
}

func (s FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) NAME() antlr.TerminalNode {
	return s.GetToken(cirqlParserNAME, 0)
}

func (s *FieldContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(cirqlParserLPAREN, 0)
}

func (s *FieldContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FieldContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(cirqlParserRPAREN, 0)
}

func (s *FieldContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(cirqlParserLBRACE, 0)
}

func (s *FieldContext) SelectionSet() ISelectionSetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectionSetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *FieldContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(cirqlParserRBRACE, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *cirqlParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, cirqlParserRULE_field)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(cirqlParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)
	if _la == cirqlParserLPAREN {
		{
			p.SetState(97)
			p.Match(cirqlParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.Arguments()
		}
		{
			p.SetState(99)
			p.Match(cirqlParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)
	if _la == cirqlParserLBRACE {
		{
			p.SetState(103)
			p.Match(cirqlParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.SelectionSet()
		}
		{
			p.SetState(105)
			p.Match(cirqlParserRBRACE)
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
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllArgument() []IArgumentContext
	Argument(i int) IArgumentContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_arguments
}

func (ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_arguments

	return p
}

func (s ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllArgument() []IArgumentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgumentContext); ok {
			len++
		}
	}

	tst := make([]IArgumentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgumentContext); ok {
			tst[i] = t.(IArgumentContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsContext) Argument(i int) IArgumentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
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

	return t.(IArgumentContext)
}

func (s *ArgumentsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(cirqlParserCOMMA)
}

func (s *ArgumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(cirqlParserCOMMA, i)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *cirqlParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, cirqlParserRULE_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Argument()
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)
	for _la == cirqlParserCOMMA {
		{
			p.SetState(110)
			p.Match(cirqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.Argument()
		}

		p.SetState(116)
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
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	COLON() antlr.TerminalNode
	ArgValue() IArgValueContext

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_argument
	return p
}

func InitEmptyArgumentContext(p *ArgumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_argument
}

func (ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_argument

	return p
}

func (s ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) NAME() antlr.TerminalNode {
	return s.GetToken(cirqlParserNAME, 0)
}

func (s *ArgumentContext) COLON() antlr.TerminalNode {
	return s.GetToken(cirqlParserCOLON, 0)
}

func (s *ArgumentContext) ArgValue() IArgValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgValueContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *cirqlParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, cirqlParserRULE_argument)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(cirqlParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(cirqlParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.ArgValue()
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
}

// IHttpStageContext is an interface to support dynamic dispatch.
type IHttpStageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HTTP() antlr.TerminalNode
	STRING() antlr.TerminalNode
	Variable() IVariableContext

	// IsHttpStageContext differentiates from other interfaces.
	IsHttpStageContext()
}

type HttpStageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHttpStageContext() *HttpStageContext {
	var p = new(HttpStageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_httpStage
	return p
}

func InitEmptyHttpStageContext(p *HttpStageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_httpStage
}

func (HttpStageContext) IsHttpStageContext() {}

func NewHttpStageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HttpStageContext {
	var p = new(HttpStageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_httpStage

	return p
}

func (s HttpStageContext) GetParser() antlr.Parser { return s.parser }

func (s *HttpStageContext) HTTP() antlr.TerminalNode {
	return s.GetToken(cirqlParserHTTP, 0)
}

func (s *HttpStageContext) STRING() antlr.TerminalNode {
	return s.GetToken(cirqlParserSTRING, 0)
}

func (s *HttpStageContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *HttpStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HttpStageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HttpStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterHttpStage(s)
	}
}

func (s *HttpStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitHttpStage(s)
	}
}

func (p *cirqlParser) HttpStage() (localctx IHttpStageContext) {
	localctx = NewHttpStageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, cirqlParserRULE_httpStage)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(cirqlParserHTTP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case cirqlParserSTRING:
		{
			p.SetState(122)
			p.Match(cirqlParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case cirqlParserDOLLAR:
		{
			p.SetState(123)
			p.Variable()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
}

// IFileStageContext is an interface to support dynamic dispatch.
type IFileStageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FILE() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsFileStageContext differentiates from other interfaces.
	IsFileStageContext()
}

type FileStageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileStageContext() *FileStageContext {
	var p = new(FileStageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_fileStage
	return p
}

func InitEmptyFileStageContext(p *FileStageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_fileStage
}

func (FileStageContext) IsFileStageContext() {}

func NewFileStageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileStageContext {
	var p = new(FileStageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_fileStage

	return p
}

func (s FileStageContext) GetParser() antlr.Parser { return s.parser }

func (s *FileStageContext) FILE() antlr.TerminalNode {
	return s.GetToken(cirqlParserFILE, 0)
}

func (s *FileStageContext) STRING() antlr.TerminalNode {
	return s.GetToken(cirqlParserSTRING, 0)
}

func (s *FileStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileStageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterFileStage(s)
	}
}

func (s *FileStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitFileStage(s)
	}
}

func (p *cirqlParser) FileStage() (localctx IFileStageContext) {
	localctx = NewFileStageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, cirqlParserRULE_fileStage)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(cirqlParserFILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.Match(cirqlParserSTRING)
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
}

// IStdinStageContext is an interface to support dynamic dispatch.
type IStdinStageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STDIN() antlr.TerminalNode

	// IsStdinStageContext differentiates from other interfaces.
	IsStdinStageContext()
}

type StdinStageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStdinStageContext() *StdinStageContext {
	var p = new(StdinStageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_stdinStage
	return p
}

func InitEmptyStdinStageContext(p *StdinStageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_stdinStage
}

func (StdinStageContext) IsStdinStageContext() {}

func NewStdinStageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StdinStageContext {
	var p = new(StdinStageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_stdinStage

	return p
}

func (s StdinStageContext) GetParser() antlr.Parser { return s.parser }

func (s *StdinStageContext) STDIN() antlr.TerminalNode {
	return s.GetToken(cirqlParserSTDIN, 0)
}

func (s *StdinStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StdinStageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StdinStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterStdinStage(s)
	}
}

func (s *StdinStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitStdinStage(s)
	}
}

func (p *cirqlParser) StdinStage() (localctx IStdinStageContext) {
	localctx = NewStdinStageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, cirqlParserRULE_stdinStage)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Match(cirqlParserSTDIN)
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
}

// ITransformStageContext is an interface to support dynamic dispatch.
type ITransformStageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MapStage() IMapStageContext
	FilterStage() IFilterStageContext
	ReduceStage() IReduceStageContext
	SortStage() ISortStageContext
	FlatMapStage() IFlatMapStageContext
	LimitStage() ILimitStageContext
	UniqStage() IUniqStageContext

	// IsTransformStageContext differentiates from other interfaces.
	IsTransformStageContext()
}

type TransformStageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransformStageContext() *TransformStageContext {
	var p = new(TransformStageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_transformStage
	return p
}

func InitEmptyTransformStageContext(p *TransformStageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_transformStage
}

func (TransformStageContext) IsTransformStageContext() {}

func NewTransformStageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransformStageContext {
	var p = new(TransformStageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_transformStage

	return p
}

func (s TransformStageContext) GetParser() antlr.Parser { return s.parser }

func (s *TransformStageContext) MapStage() IMapStageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapStageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapStageContext)
}

func (s *TransformStageContext) FilterStage() IFilterStageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterStageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterStageContext)
}

func (s *TransformStageContext) ReduceStage() IReduceStageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReduceStageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReduceStageContext)
}

func (s *TransformStageContext) SortStage() ISortStageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortStageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortStageContext)
}

func (s *TransformStageContext) FlatMapStage() IFlatMapStageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFlatMapStageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFlatMapStageContext)
}

func (s *TransformStageContext) LimitStage() ILimitStageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimitStageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimitStageContext)
}

func (s *TransformStageContext) UniqStage() IUniqStageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUniqStageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUniqStageContext)
}

func (s *TransformStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransformStageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransformStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterTransformStage(s)
	}
}

func (s *TransformStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitTransformStage(s)
	}
}

func (p *cirqlParser) TransformStage() (localctx ITransformStageContext) {
	localctx = NewTransformStageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, cirqlParserRULE_transformStage)
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case cirqlParserMAP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.MapStage()
		}

	case cirqlParserFILTER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)
			p.FilterStage()
		}

	case cirqlParserREDUCE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(133)
			p.ReduceStage()
		}

	case cirqlParserSORT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(134)
			p.SortStage()
		}

	case cirqlParserFLATMAP:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(135)
			p.FlatMapStage()
		}

	case cirqlParserLIMIT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(136)
			p.LimitStage()
		}

	case cirqlParserUNIQ:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(137)
			p.UniqStage()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
}

// IMapStageContext is an interface to support dynamic dispatch.
type IMapStageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MAP() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	AllMapping() []IMappingContext
	Mapping(i int) IMappingContext
	RBRACE() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsMapStageContext differentiates from other interfaces.
	IsMapStageContext()
}

type MapStageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapStageContext() *MapStageContext {
	var p = new(MapStageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_mapStage
	return p
}

func InitEmptyMapStageContext(p *MapStageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_mapStage
}

func (MapStageContext) IsMapStageContext() {}

func NewMapStageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapStageContext {
	var p = new(MapStageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_mapStage

	return p
}

func (s MapStageContext) GetParser() antlr.Parser { return s.parser }

func (s *MapStageContext) MAP() antlr.TerminalNode {
	return s.GetToken(cirqlParserMAP, 0)
}

func (s *MapStageContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(cirqlParserLBRACE, 0)
}

func (s *MapStageContext) AllMapping() []IMappingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMappingContext); ok {
			len++
		}
	}

	tst := make([]IMappingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMappingContext); ok {
			tst[i] = t.(IMappingContext)
			i++
		}
	}

	return tst
}

func (s *MapStageContext) Mapping(i int) IMappingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMappingContext); ok {
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

	return t.(IMappingContext)
}

func (s *MapStageContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(cirqlParserRBRACE, 0)
}

func (s *MapStageContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(cirqlParserCOMMA)
}

func (s *MapStageContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(cirqlParserCOMMA, i)
}

func (s *MapStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapStageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterMapStage(s)
	}
}

func (s *MapStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitMapStage(s)
	}
}

func (p *cirqlParser) MapStage() (localctx IMapStageContext) {
	localctx = NewMapStageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, cirqlParserRULE_mapStage)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(cirqlParserMAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Match(cirqlParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.Mapping()
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)
	for _la == cirqlParserCOMMA {
		{
			p.SetState(143)
			p.Match(cirqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.Mapping()
		}

		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(150)
		p.Match(cirqlParserRBRACE)
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
}

// IFlatMapStageContext is an interface to support dynamic dispatch.
type IFlatMapStageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FLATMAP() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	AllMapping() []IMappingContext
	Mapping(i int) IMappingContext
	RBRACE() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFlatMapStageContext differentiates from other interfaces.
	IsFlatMapStageContext()
}

type FlatMapStageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlatMapStageContext() *FlatMapStageContext {
	var p = new(FlatMapStageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_flatMapStage
	return p
}

func InitEmptyFlatMapStageContext(p *FlatMapStageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_flatMapStage
}

func (FlatMapStageContext) IsFlatMapStageContext() {}

func NewFlatMapStageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FlatMapStageContext {
	var p = new(FlatMapStageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_flatMapStage

	return p
}

func (s FlatMapStageContext) GetParser() antlr.Parser { return s.parser }

func (s *FlatMapStageContext) FLATMAP() antlr.TerminalNode {
	return s.GetToken(cirqlParserFLATMAP, 0)
}

func (s *FlatMapStageContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(cirqlParserLBRACE, 0)
}

func (s *FlatMapStageContext) AllMapping() []IMappingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMappingContext); ok {
			len++
		}
	}

	tst := make([]IMappingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMappingContext); ok {
			tst[i] = t.(IMappingContext)
			i++
		}
	}

	return tst
}

func (s *FlatMapStageContext) Mapping(i int) IMappingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMappingContext); ok {
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

	return t.(IMappingContext)
}

func (s *FlatMapStageContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(cirqlParserRBRACE, 0)
}

func (s *FlatMapStageContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(cirqlParserCOMMA)
}

func (s *FlatMapStageContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(cirqlParserCOMMA, i)
}

func (s *FlatMapStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlatMapStageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FlatMapStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterFlatMapStage(s)
	}
}

func (s *FlatMapStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitFlatMapStage(s)
	}
}

func (p *cirqlParser) FlatMapStage() (localctx IFlatMapStageContext) {
	localctx = NewFlatMapStageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, cirqlParserRULE_flatMapStage)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Match(cirqlParserFLATMAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(cirqlParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.Mapping()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)
	for _la == cirqlParserCOMMA {
		{
			p.SetState(155)
			p.Match(cirqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.Mapping()
		}

		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(162)
		p.Match(cirqlParserRBRACE)
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
}

// IMappingContext is an interface to support dynamic dispatch.
type IMappingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Expr() IExprContext

	// IsMappingContext differentiates from other interfaces.
	IsMappingContext()
}

type MappingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMappingContext() *MappingContext {
	var p = new(MappingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_mapping
	return p
}

func InitEmptyMappingContext(p *MappingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_mapping
}

func (MappingContext) IsMappingContext() {}

func NewMappingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MappingContext {
	var p = new(MappingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_mapping

	return p
}

func (s MappingContext) GetParser() antlr.Parser { return s.parser }

func (s *MappingContext) NAME() antlr.TerminalNode {
	return s.GetToken(cirqlParserNAME, 0)
}

func (s *MappingContext) COLON() antlr.TerminalNode {
	return s.GetToken(cirqlParserCOLON, 0)
}

func (s *MappingContext) Expr() IExprContext {
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

func (s *MappingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MappingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MappingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterMapping(s)
	}
}

func (s *MappingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitMapping(s)
	}
}

func (p *cirqlParser) Mapping() (localctx IMappingContext) {
	localctx = NewMappingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, cirqlParserRULE_mapping)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Match(cirqlParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)
		p.Match(cirqlParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
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
}

// IFilterStageContext is an interface to support dynamic dispatch.
type IFilterStageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FILTER() antlr.TerminalNode
	Expr() IExprContext

	// IsFilterStageContext differentiates from other interfaces.
	IsFilterStageContext()
}

type FilterStageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterStageContext() *FilterStageContext {
	var p = new(FilterStageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_filterStage
	return p
}

func InitEmptyFilterStageContext(p *FilterStageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_filterStage
}

func (FilterStageContext) IsFilterStageContext() {}

func NewFilterStageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterStageContext {
	var p = new(FilterStageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_filterStage

	return p
}

func (s FilterStageContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterStageContext) FILTER() antlr.TerminalNode {
	return s.GetToken(cirqlParserFILTER, 0)
}

func (s *FilterStageContext) Expr() IExprContext {
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

func (s *FilterStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterStageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterFilterStage(s)
	}
}

func (s *FilterStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitFilterStage(s)
	}
}

func (p *cirqlParser) FilterStage() (localctx IFilterStageContext) {
	localctx = NewFilterStageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, cirqlParserRULE_filterStage)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(cirqlParserFILTER)
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
}

// IReduceStageContext is an interface to support dynamic dispatch.
type IReduceStageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REDUCE() antlr.TerminalNode
	ReduceOp() IReduceOpContext
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode

	// IsReduceStageContext differentiates from other interfaces.
	IsReduceStageContext()
}

type ReduceStageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReduceStageContext() *ReduceStageContext {
	var p = new(ReduceStageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_reduceStage
	return p
}

func InitEmptyReduceStageContext(p *ReduceStageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_reduceStage
}

func (ReduceStageContext) IsReduceStageContext() {}

func NewReduceStageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReduceStageContext {
	var p = new(ReduceStageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_reduceStage

	return p
}

func (s ReduceStageContext) GetParser() antlr.Parser { return s.parser }

func (s *ReduceStageContext) REDUCE() antlr.TerminalNode {
	return s.GetToken(cirqlParserREDUCE, 0)
}

func (s *ReduceStageContext) ReduceOp() IReduceOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReduceOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReduceOpContext)
}

func (s *ReduceStageContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(cirqlParserLPAREN, 0)
}

func (s *ReduceStageContext) Expr() IExprContext {
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

func (s *ReduceStageContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(cirqlParserRPAREN, 0)
}

func (s *ReduceStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReduceStageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReduceStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterReduceStage(s)
	}
}

func (s *ReduceStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitReduceStage(s)
	}
}

func (p *cirqlParser) ReduceStage() (localctx IReduceStageContext) {
	localctx = NewReduceStageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, cirqlParserRULE_reduceStage)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Match(cirqlParserREDUCE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		p.ReduceOp()
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)
	if _la == cirqlParserLPAREN {
		{
			p.SetState(173)
			p.Match(cirqlParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.expr(0)
		}
		{
			p.SetState(175)
			p.Match(cirqlParserRPAREN)
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
}

// IReduceOpContext is an interface to support dynamic dispatch.
type IReduceOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COUNT() antlr.TerminalNode
	SUM() antlr.TerminalNode
	MIN() antlr.TerminalNode
	MAX() antlr.TerminalNode
	AVG() antlr.TerminalNode
	FIRST() antlr.TerminalNode
	LAST() antlr.TerminalNode
	GROUP_BY() antlr.TerminalNode
	COLLECT() antlr.TerminalNode

	// IsReduceOpContext differentiates from other interfaces.
	IsReduceOpContext()
}

type ReduceOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReduceOpContext() *ReduceOpContext {
	var p = new(ReduceOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_reduceOp
	return p
}

func InitEmptyReduceOpContext(p *ReduceOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_reduceOp
}

func (ReduceOpContext) IsReduceOpContext() {}

func NewReduceOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReduceOpContext {
	var p = new(ReduceOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_reduceOp

	return p
}

func (s ReduceOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ReduceOpContext) COUNT() antlr.TerminalNode {
	return s.GetToken(cirqlParserCOUNT, 0)
}

func (s *ReduceOpContext) SUM() antlr.TerminalNode {
	return s.GetToken(cirqlParserSUM, 0)
}

func (s *ReduceOpContext) MIN() antlr.TerminalNode {
	return s.GetToken(cirqlParserMIN, 0)
}

func (s *ReduceOpContext) MAX() antlr.TerminalNode {
	return s.GetToken(cirqlParserMAX, 0)
}

func (s *ReduceOpContext) AVG() antlr.TerminalNode {
	return s.GetToken(cirqlParserAVG, 0)
}

func (s *ReduceOpContext) FIRST() antlr.TerminalNode {
	return s.GetToken(cirqlParserFIRST, 0)
}

func (s *ReduceOpContext) LAST() antlr.TerminalNode {
	return s.GetToken(cirqlParserLAST, 0)
}

func (s *ReduceOpContext) GROUP_BY() antlr.TerminalNode {
	return s.GetToken(cirqlParserGROUP_BY, 0)
}

func (s *ReduceOpContext) COLLECT() antlr.TerminalNode {
	return s.GetToken(cirqlParserCOLLECT, 0)
}

func (s *ReduceOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReduceOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReduceOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterReduceOp(s)
	}
}

func (s *ReduceOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitReduceOp(s)
	}
}

func (p *cirqlParser) ReduceOp() (localctx IReduceOpContext) {
	localctx = NewReduceOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, cirqlParserRULE_reduceOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2093056) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
}

// ISortStageContext is an interface to support dynamic dispatch.
type ISortStageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SORT() antlr.TerminalNode
	Expr() IExprContext
	ASC() antlr.TerminalNode
	DESC() antlr.TerminalNode

	// IsSortStageContext differentiates from other interfaces.
	IsSortStageContext()
}

type SortStageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortStageContext() *SortStageContext {
	var p = new(SortStageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_sortStage
	return p
}

func InitEmptySortStageContext(p *SortStageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_sortStage
}

func (SortStageContext) IsSortStageContext() {}

func NewSortStageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortStageContext {
	var p = new(SortStageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_sortStage

	return p
}

func (s SortStageContext) GetParser() antlr.Parser { return s.parser }

func (s *SortStageContext) SORT() antlr.TerminalNode {
	return s.GetToken(cirqlParserSORT, 0)
}

func (s *SortStageContext) Expr() IExprContext {
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

func (s *SortStageContext) ASC() antlr.TerminalNode {
	return s.GetToken(cirqlParserASC, 0)
}

func (s *SortStageContext) DESC() antlr.TerminalNode {
	return s.GetToken(cirqlParserDESC, 0)
}

func (s *SortStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortStageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterSortStage(s)
	}
}

func (s *SortStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitSortStage(s)
	}
}

func (p *cirqlParser) SortStage() (localctx ISortStageContext) {
	localctx = NewSortStageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, cirqlParserRULE_sortStage)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(cirqlParserSORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.expr(0)
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)
	if _la == cirqlParserASC || _la == cirqlParserDESC {
		{
			p.SetState(183)
			_la = p.GetTokenStream().LA(1)

			if !(_la == cirqlParserASC || _la == cirqlParserDESC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
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
}

// ILimitStageContext is an interface to support dynamic dispatch.
type ILimitStageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LIMIT() antlr.TerminalNode
	INT() antlr.TerminalNode

	// IsLimitStageContext differentiates from other interfaces.
	IsLimitStageContext()
}

type LimitStageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitStageContext() *LimitStageContext {
	var p = new(LimitStageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_limitStage
	return p
}

func InitEmptyLimitStageContext(p *LimitStageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_limitStage
}

func (LimitStageContext) IsLimitStageContext() {}

func NewLimitStageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitStageContext {
	var p = new(LimitStageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_limitStage

	return p
}

func (s LimitStageContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitStageContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(cirqlParserLIMIT, 0)
}

func (s *LimitStageContext) INT() antlr.TerminalNode {
	return s.GetToken(cirqlParserINT, 0)
}

func (s *LimitStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitStageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterLimitStage(s)
	}
}

func (s *LimitStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitLimitStage(s)
	}
}

func (p *cirqlParser) LimitStage() (localctx ILimitStageContext) {
	localctx = NewLimitStageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, cirqlParserRULE_limitStage)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Match(cirqlParserLIMIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		p.Match(cirqlParserINT)
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
}

// IUniqStageContext is an interface to support dynamic dispatch.
type IUniqStageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UNIQ() antlr.TerminalNode
	Expr() IExprContext

	// IsUniqStageContext differentiates from other interfaces.
	IsUniqStageContext()
}

type UniqStageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUniqStageContext() *UniqStageContext {
	var p = new(UniqStageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_uniqStage
	return p
}

func InitEmptyUniqStageContext(p *UniqStageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_uniqStage
}

func (UniqStageContext) IsUniqStageContext() {}

func NewUniqStageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UniqStageContext {
	var p = new(UniqStageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_uniqStage

	return p
}

func (s UniqStageContext) GetParser() antlr.Parser { return s.parser }

func (s *UniqStageContext) UNIQ() antlr.TerminalNode {
	return s.GetToken(cirqlParserUNIQ, 0)
}

func (s *UniqStageContext) Expr() IExprContext {
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

func (s *UniqStageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UniqStageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UniqStageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterUniqStage(s)
	}
}

func (s *UniqStageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitUniqStage(s)
	}
}

func (p *cirqlParser) UniqStage() (localctx IUniqStageContext) {
	localctx = NewUniqStageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, cirqlParserRULE_uniqStage)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Match(cirqlParserUNIQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)
	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33918287863611392) != 0 {
		{
			p.SetState(190)
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
	p.RuleIndex = cirqlParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_expr
}

func (ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_expr

	return p
}

func (s ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MulExprContext struct {
	ExprContext
}

func NewMulExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulExprContext {
	var p = new(MulExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MulExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulExprContext) AllExpr() []IExprContext {
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

func (s *MulExprContext) Expr(i int) IExprContext {
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

func (s *MulExprContext) STAR() antlr.TerminalNode {
	return s.GetToken(cirqlParserSTAR, 0)
}

func (s *MulExprContext) SLASH() antlr.TerminalNode {
	return s.GetToken(cirqlParserSLASH, 0)
}

func (s *MulExprContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(cirqlParserPERCENT, 0)
}

func (s *MulExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterMulExpr(s)
	}
}

func (s *MulExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitMulExpr(s)
	}
}

type AndExprContext struct {
	ExprContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpr() []IExprContext {
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

func (s *AndExprContext) Expr(i int) IExprContext {
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

func (s *AndExprContext) AND() antlr.TerminalNode {
	return s.GetToken(cirqlParserAND, 0)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

type LitExprContext struct {
	ExprContext
}

func NewLitExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitExprContext {
	var p = new(LitExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LitExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitExprContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LitExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterLitExpr(s)
	}
}

func (s *LitExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitLitExpr(s)
	}
}

type CmpExprContext struct {
	ExprContext
}

func NewCmpExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CmpExprContext {
	var p = new(CmpExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CmpExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmpExprContext) AllExpr() []IExprContext {
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

func (s *CmpExprContext) Expr(i int) IExprContext {
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

func (s *CmpExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(cirqlParserEQ, 0)
}

func (s *CmpExprContext) NE() antlr.TerminalNode {
	return s.GetToken(cirqlParserNE, 0)
}

func (s *CmpExprContext) GE() antlr.TerminalNode {
	return s.GetToken(cirqlParserGE, 0)
}

func (s *CmpExprContext) LE() antlr.TerminalNode {
	return s.GetToken(cirqlParserLE, 0)
}

func (s *CmpExprContext) GT() antlr.TerminalNode {
	return s.GetToken(cirqlParserGT, 0)
}

func (s *CmpExprContext) LT() antlr.TerminalNode {
	return s.GetToken(cirqlParserLT, 0)
}

func (s *CmpExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterCmpExpr(s)
	}
}

func (s *CmpExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitCmpExpr(s)
	}
}

type VarExprContext struct {
	ExprContext
}

func NewVarExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarExprContext {
	var p = new(VarExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *VarExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarExprContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VarExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterVarExpr(s)
	}
}

func (s *VarExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitVarExpr(s)
	}
}

type CallExprContext struct {
	ExprContext
}

func NewCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExprContext {
	var p = new(CallExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExprContext) FuncCall() IFuncCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *CallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterCallExpr(s)
	}
}

func (s *CallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitCallExpr(s)
	}
}

type AddExprContext struct {
	ExprContext
}

func NewAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExprContext {
	var p = new(AddExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExprContext) AllExpr() []IExprContext {
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

func (s *AddExprContext) Expr(i int) IExprContext {
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

func (s *AddExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(cirqlParserPLUS, 0)
}

func (s *AddExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(cirqlParserMINUS, 0)
}

func (s *AddExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterAddExpr(s)
	}
}

func (s *AddExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitAddExpr(s)
	}
}

type FieldExprContext struct {
	ExprContext
}

func NewFieldExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldExprContext {
	var p = new(FieldExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FieldExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldExprContext) FieldAccess() IFieldAccessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldAccessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldAccessContext)
}

func (s *FieldExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterFieldExpr(s)
	}
}

func (s *FieldExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitFieldExpr(s)
	}
}

type ParenExprContext struct {
	ExprContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(cirqlParserLPAREN, 0)
}

func (s *ParenExprContext) Expr() IExprContext {
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

func (s *ParenExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(cirqlParserRPAREN, 0)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

type UnaryExprContext struct {
	ExprContext
}

func NewUnaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExprContext {
	var p = new(UnaryExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) Expr() IExprContext {
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

func (s *UnaryExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(cirqlParserNOT, 0)
}

func (s *UnaryExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(cirqlParserMINUS, 0)
}

func (s *UnaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterUnaryExpr(s)
	}
}

func (s *UnaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitUnaryExpr(s)
	}
}

type OrExprContext struct {
	ExprContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpr() []IExprContext {
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

func (s *OrExprContext) Expr(i int) IExprContext {
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

func (s *OrExprContext) OR() antlr.TerminalNode {
	return s.GetToken(cirqlParserOR, 0)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (p *cirqlParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *cirqlParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, cirqlParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case cirqlParserNOT, cirqlParserMINUS:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(194)
			_la = p.GetTokenStream().LA(1)

			if !(_la == cirqlParserNOT || _la == cirqlParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(195)
			p.expr(11)
		}

	case cirqlParserLPAREN:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(196)
			p.Match(cirqlParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.expr(0)
		}
		{
			p.SetState(198)
			p.Match(cirqlParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case cirqlParserNAME:
		localctx = NewCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(200)
			p.FuncCall()
		}

	case cirqlParserDOT:
		localctx = NewFieldExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(201)
			p.FieldAccess()
		}

	case cirqlParserDOLLAR:
		localctx = NewVarExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(202)
			p.Variable()
		}

	case cirqlParserTRUE, cirqlParserFALSE, cirqlParserNULL, cirqlParserFLOAT, cirqlParserINT, cirqlParserSTRING:
		localctx = NewLitExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(203)
			p.Literal()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(221)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, cirqlParserRULE_expr)
				p.SetState(206)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(207)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1970324836974592) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(208)
					p.expr(11)
				}

			case 2:
				localctx = NewAddExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, cirqlParserRULE_expr)
				p.SetState(209)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(210)
					_la = p.GetTokenStream().LA(1)

					if !(_la == cirqlParserPLUS || _la == cirqlParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(211)
					p.expr(10)
				}

			case 3:
				localctx = NewCmpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, cirqlParserRULE_expr)
				p.SetState(212)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(213)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&69269232549888) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(214)
					p.expr(9)
				}

			case 4:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, cirqlParserRULE_expr)
				p.SetState(215)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(216)
					p.Match(cirqlParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(217)
					p.expr(8)
				}

			case 5:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, cirqlParserRULE_expr)
				p.SetState(218)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(219)
					p.Match(cirqlParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(220)
					p.expr(7)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
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
}

// IFieldAccessContext is an interface to support dynamic dispatch.
type IFieldAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOT() antlr.TerminalNode
	AllPathSeg() []IPathSegContext
	PathSeg(i int) IPathSegContext

	// IsFieldAccessContext differentiates from other interfaces.
	IsFieldAccessContext()
}

type FieldAccessContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldAccessContext() *FieldAccessContext {
	var p = new(FieldAccessContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_fieldAccess
	return p
}

func InitEmptyFieldAccessContext(p *FieldAccessContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_fieldAccess
}

func (FieldAccessContext) IsFieldAccessContext() {}

func NewFieldAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldAccessContext {
	var p = new(FieldAccessContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_fieldAccess

	return p
}

func (s FieldAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldAccessContext) DOT() antlr.TerminalNode {
	return s.GetToken(cirqlParserDOT, 0)
}

func (s *FieldAccessContext) AllPathSeg() []IPathSegContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPathSegContext); ok {
			len++
		}
	}

	tst := make([]IPathSegContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPathSegContext); ok {
			tst[i] = t.(IPathSegContext)
			i++
		}
	}

	return tst
}

func (s *FieldAccessContext) PathSeg(i int) IPathSegContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathSegContext); ok {
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

	return t.(IPathSegContext)
}

func (s *FieldAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterFieldAccess(s)
	}
}

func (s *FieldAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitFieldAccess(s)
	}
}

func (p *cirqlParser) FieldAccess() (localctx IFieldAccessContext) {
	localctx = NewFieldAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, cirqlParserRULE_fieldAccess)
	var _alt int

	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(226)
			p.Match(cirqlParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)
			p.PathSeg()
		}
		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(228)
					p.PathSeg()
				}

			}
			p.SetState(233)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(234)
			p.Match(cirqlParserDOT)
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
}

// IPathSegContext is an interface to support dynamic dispatch.
type IPathSegContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	DOT() antlr.TerminalNode
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode

	// IsPathSegContext differentiates from other interfaces.
	IsPathSegContext()
}

type PathSegContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathSegContext() *PathSegContext {
	var p = new(PathSegContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_pathSeg
	return p
}

func InitEmptyPathSegContext(p *PathSegContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_pathSeg
}

func (PathSegContext) IsPathSegContext() {}

func NewPathSegContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathSegContext {
	var p = new(PathSegContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_pathSeg

	return p
}

func (s PathSegContext) GetParser() antlr.Parser { return s.parser }

func (s *PathSegContext) NAME() antlr.TerminalNode {
	return s.GetToken(cirqlParserNAME, 0)
}

func (s *PathSegContext) DOT() antlr.TerminalNode {
	return s.GetToken(cirqlParserDOT, 0)
}

func (s *PathSegContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(cirqlParserLBRACK, 0)
}

func (s *PathSegContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(cirqlParserRBRACK, 0)
}

func (s *PathSegContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathSegContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathSegContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterPathSeg(s)
	}
}

func (s *PathSegContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitPathSeg(s)
	}
}

func (p *cirqlParser) PathSeg() (localctx IPathSegContext) {
	localctx = NewPathSegContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, cirqlParserRULE_pathSeg)
	var _la int

	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case cirqlParserDOT, cirqlParserNAME:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == cirqlParserDOT {
			{
				p.SetState(237)
				p.Match(cirqlParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(240)
			p.Match(cirqlParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case cirqlParserLBRACK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(241)
			p.Match(cirqlParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.Match(cirqlParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
}

// IFuncCallContext is an interface to support dynamic dispatch.
type IFuncCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFuncCallContext differentiates from other interfaces.
	IsFuncCallContext()
}

type FuncCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallContext() *FuncCallContext {
	var p = new(FuncCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_funcCall
	return p
}

func InitEmptyFuncCallContext(p *FuncCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_funcCall
}

func (FuncCallContext) IsFuncCallContext() {}

func NewFuncCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallContext {
	var p = new(FuncCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_funcCall

	return p
}

func (s FuncCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallContext) NAME() antlr.TerminalNode {
	return s.GetToken(cirqlParserNAME, 0)
}

func (s *FuncCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(cirqlParserLPAREN, 0)
}

func (s *FuncCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(cirqlParserRPAREN, 0)
}

func (s *FuncCallContext) AllExpr() []IExprContext {
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

func (s *FuncCallContext) Expr(i int) IExprContext {
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

func (s *FuncCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(cirqlParserCOMMA)
}

func (s *FuncCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(cirqlParserCOMMA, i)
}

func (s *FuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterFuncCall(s)
	}
}

func (s *FuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitFuncCall(s)
	}
}

func (p *cirqlParser) FuncCall() (localctx IFuncCallContext) {
	localctx = NewFuncCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, cirqlParserRULE_funcCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.Match(cirqlParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(246)
		p.Match(cirqlParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)
	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33918287863611392) != 0 {
		{
			p.SetState(247)
			p.expr(0)
		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == cirqlParserCOMMA {
			{
				p.SetState(248)
				p.Match(cirqlParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(249)
				p.expr(0)
			}

			p.SetState(254)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(257)
		p.Match(cirqlParserRPAREN)
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
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOLLAR() antlr.TerminalNode
	NAME() antlr.TerminalNode

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_variable
	return p
}

func InitEmptyVariableContext(p *VariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_variable
}

func (VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_variable

	return p
}

func (s VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) DOLLAR() antlr.TerminalNode {
	return s.GetToken(cirqlParserDOLLAR, 0)
}

func (s *VariableContext) NAME() antlr.TerminalNode {
	return s.GetToken(cirqlParserNAME, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *cirqlParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, cirqlParserRULE_variable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(cirqlParserDOLLAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.Match(cirqlParserNAME)
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
}

// IArgValueContext is an interface to support dynamic dispatch.
type IArgValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable() IVariableContext
	Literal() ILiteralContext
	ListLit() IListLitContext
	ObjectLit() IObjectLitContext

	// IsArgValueContext differentiates from other interfaces.
	IsArgValueContext()
}

type ArgValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgValueContext() *ArgValueContext {
	var p = new(ArgValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_argValue
	return p
}

func InitEmptyArgValueContext(p *ArgValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_argValue
}

func (ArgValueContext) IsArgValueContext() {}

func NewArgValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgValueContext {
	var p = new(ArgValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_argValue

	return p
}

func (s ArgValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgValueContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ArgValueContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ArgValueContext) ListLit() IListLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListLitContext)
}

func (s *ArgValueContext) ObjectLit() IObjectLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectLitContext)
}

func (s *ArgValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterArgValue(s)
	}
}

func (s *ArgValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitArgValue(s)
	}
}

func (p *cirqlParser) ArgValue() (localctx IArgValueContext) {
	localctx = NewArgValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, cirqlParserRULE_argValue)
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case cirqlParserDOLLAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(262)
			p.Variable()
		}

	case cirqlParserTRUE, cirqlParserFALSE, cirqlParserNULL, cirqlParserFLOAT, cirqlParserINT, cirqlParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(263)
			p.Literal()
		}

	case cirqlParserLBRACK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(264)
			p.ListLit()
		}

	case cirqlParserLBRACE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(265)
			p.ObjectLit()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
}

// IObjectLitContext is an interface to support dynamic dispatch.
type IObjectLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllNAME() []antlr.TerminalNode
	NAME(i int) antlr.TerminalNode
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	AllArgValue() []IArgValueContext
	ArgValue(i int) IArgValueContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsObjectLitContext differentiates from other interfaces.
	IsObjectLitContext()
}

type ObjectLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectLitContext() *ObjectLitContext {
	var p = new(ObjectLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_objectLit
	return p
}

func InitEmptyObjectLitContext(p *ObjectLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_objectLit
}

func (ObjectLitContext) IsObjectLitContext() {}

func NewObjectLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectLitContext {
	var p = new(ObjectLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_objectLit

	return p
}

func (s ObjectLitContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectLitContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(cirqlParserLBRACE, 0)
}

func (s *ObjectLitContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(cirqlParserRBRACE, 0)
}

func (s *ObjectLitContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(cirqlParserNAME)
}

func (s *ObjectLitContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(cirqlParserNAME, i)
}

func (s *ObjectLitContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(cirqlParserCOLON)
}

func (s *ObjectLitContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(cirqlParserCOLON, i)
}

func (s *ObjectLitContext) AllArgValue() []IArgValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgValueContext); ok {
			len++
		}
	}

	tst := make([]IArgValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgValueContext); ok {
			tst[i] = t.(IArgValueContext)
			i++
		}
	}

	return tst
}

func (s *ObjectLitContext) ArgValue(i int) IArgValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgValueContext); ok {
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

	return t.(IArgValueContext)
}

func (s *ObjectLitContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(cirqlParserCOMMA)
}

func (s *ObjectLitContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(cirqlParserCOMMA, i)
}

func (s *ObjectLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterObjectLit(s)
	}
}

func (s *ObjectLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitObjectLit(s)
	}
}

func (p *cirqlParser) ObjectLit() (localctx IObjectLitContext) {
	localctx = NewObjectLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, cirqlParserRULE_objectLit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Match(cirqlParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)
	if _la == cirqlParserNAME {
		{
			p.SetState(269)
			p.Match(cirqlParserNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)
			p.Match(cirqlParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(271)
			p.ArgValue()
		}
		p.SetState(278)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == cirqlParserCOMMA {
			{
				p.SetState(272)
				p.Match(cirqlParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(273)
				p.Match(cirqlParserNAME)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(274)
				p.Match(cirqlParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(275)
				p.ArgValue()
			}

			p.SetState(280)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(283)
		p.Match(cirqlParserRBRACE)
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
}

// IListLitContext is an interface to support dynamic dispatch.
type IListLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	AllArgValue() []IArgValueContext
	ArgValue(i int) IArgValueContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsListLitContext differentiates from other interfaces.
	IsListLitContext()
}

type ListLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListLitContext() *ListLitContext {
	var p = new(ListLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_listLit
	return p
}

func InitEmptyListLitContext(p *ListLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_listLit
}

func (ListLitContext) IsListLitContext() {}

func NewListLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListLitContext {
	var p = new(ListLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_listLit

	return p
}

func (s ListLitContext) GetParser() antlr.Parser { return s.parser }

func (s *ListLitContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(cirqlParserLBRACK, 0)
}

func (s *ListLitContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(cirqlParserRBRACK, 0)
}

func (s *ListLitContext) AllArgValue() []IArgValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgValueContext); ok {
			len++
		}
	}

	tst := make([]IArgValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgValueContext); ok {
			tst[i] = t.(IArgValueContext)
			i++
		}
	}

	return tst
}

func (s *ListLitContext) ArgValue(i int) IArgValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgValueContext); ok {
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

	return t.(IArgValueContext)
}

func (s *ListLitContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(cirqlParserCOMMA)
}

func (s *ListLitContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(cirqlParserCOMMA, i)
}

func (s *ListLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterListLit(s)
	}
}

func (s *ListLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitListLit(s)
	}
}

func (p *cirqlParser) ListLit() (localctx IListLitContext) {
	localctx = NewListLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, cirqlParserRULE_listLit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		p.Match(cirqlParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)
	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31525241473728512) != 0 {
		{
			p.SetState(286)
			p.ArgValue()
		}
		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == cirqlParserCOMMA {
			{
				p.SetState(287)
				p.Match(cirqlParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(288)
				p.ArgValue()
			}

			p.SetState(293)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(296)
		p.Match(cirqlParserRBRACK)
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
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	INT() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	NULL() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = cirqlParserRULE_literal
}

func (LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = cirqlParserRULE_literal

	return p
}

func (s LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(cirqlParserSTRING, 0)
}

func (s *LiteralContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(cirqlParserFLOAT, 0)
}

func (s *LiteralContext) INT() antlr.TerminalNode {
	return s.GetToken(cirqlParserINT, 0)
}

func (s *LiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(cirqlParserTRUE, 0)
}

func (s *LiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(cirqlParserFALSE, 0)
}

func (s *LiteralContext) NULL() antlr.TerminalNode {
	return s.GetToken(cirqlParserNULL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cirqlListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *cirqlParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, cirqlParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31525197450313728) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
}

func (p *cirqlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 22:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *cirqlParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

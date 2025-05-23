// Code generated from grammar/BigC.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // BigC
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

type BigCParser struct {
	*antlr.BaseParser
}

var BigCParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bigcParserInit() {
	staticData := &BigCParserStaticData
	staticData.LiteralNames = []string{
		"", "'['", "']'", "'int'", "'float'", "'bool'", "'char'", "'void'",
		"'('", "')'", "';'", "','", "'{'", "'}'", "'if'", "'else'", "'while'",
		"'return'", "'='", "'||'", "'&&'", "'=='", "'!='", "'>'", "'<'", "'>='",
		"'<='", "'+'", "'-'", "'*'", "'/'", "'!'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "BooleanConstant",
		"CharConstant", "Identifier", "IntegerConstant", "FloatingConstant",
		"WS", "COMMENT", "MULTILINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"program", "declaration", "arrayNotation", "type", "declarationRemainder",
		"parameterList", "parameter", "block", "blockItem", "statement", "ifStatement",
		"elseClause", "nonIfStatement", "whileStatement", "returnStatement",
		"expression", "assignmentExpression", "assignmentRest", "variableInitializer",
		"logicalOrExpression", "logicalOrRest", "logicalAndExpression", "logicalAndRest",
		"equalityExpression", "equalityRest", "equalityOperator", "comparisonExpression",
		"comparisonRest", "comparisonOperator", "additionExpression", "additionExpressionRest",
		"addSubtractOperator", "multiplicationExpression", "multiplicationExpressionRest",
		"multDivModOperator", "unaryExpression", "unaryOperator", "postfixExpression",
		"arrayAccess", "functionCallArgs", "argList", "primaryExpression", "constant",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 39, 301, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 1, 0, 5, 0, 88, 8, 0, 10, 0, 12, 0, 91, 9, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 3, 1, 98, 8, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 4, 1, 4, 3, 4, 110, 8, 4, 1, 4, 1, 4, 1, 4, 3, 4, 115, 8, 4, 1,
		4, 3, 4, 118, 8, 4, 1, 5, 1, 5, 1, 5, 5, 5, 123, 8, 5, 10, 5, 12, 5, 126,
		9, 5, 1, 6, 1, 6, 1, 6, 3, 6, 131, 8, 6, 1, 7, 1, 7, 5, 7, 135, 8, 7, 10,
		7, 12, 7, 138, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 3, 8, 144, 8, 8, 1, 9, 1,
		9, 3, 9, 148, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 156,
		8, 10, 1, 11, 1, 11, 1, 11, 3, 11, 161, 8, 11, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 3, 12, 168, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 3, 16, 184, 8,
		16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 5, 19, 194,
		8, 19, 10, 19, 12, 19, 197, 9, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 5,
		21, 204, 8, 21, 10, 21, 12, 21, 207, 9, 21, 1, 22, 1, 22, 1, 22, 1, 23,
		1, 23, 5, 23, 214, 8, 23, 10, 23, 12, 23, 217, 9, 23, 1, 24, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 26, 1, 26, 5, 26, 226, 8, 26, 10, 26, 12, 26, 229,
		9, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 5, 29, 238, 8,
		29, 10, 29, 12, 29, 241, 9, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32,
		1, 32, 5, 32, 250, 8, 32, 10, 32, 12, 32, 253, 9, 32, 1, 33, 1, 33, 1,
		33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 264, 8, 35, 1, 36,
		1, 36, 1, 37, 1, 37, 1, 37, 3, 37, 271, 8, 37, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 39, 1, 39, 3, 39, 279, 8, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40,
		5, 40, 286, 8, 40, 10, 40, 12, 40, 289, 9, 40, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 41, 3, 41, 297, 8, 41, 1, 42, 1, 42, 1, 42, 0, 0, 43, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
		40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74,
		76, 78, 80, 82, 84, 0, 7, 1, 0, 3, 7, 1, 0, 21, 22, 1, 0, 23, 26, 1, 0,
		27, 28, 1, 0, 29, 30, 2, 0, 28, 28, 31, 31, 2, 0, 32, 33, 35, 36, 285,
		0, 89, 1, 0, 0, 0, 2, 94, 1, 0, 0, 0, 4, 101, 1, 0, 0, 0, 6, 105, 1, 0,
		0, 0, 8, 117, 1, 0, 0, 0, 10, 119, 1, 0, 0, 0, 12, 127, 1, 0, 0, 0, 14,
		132, 1, 0, 0, 0, 16, 143, 1, 0, 0, 0, 18, 147, 1, 0, 0, 0, 20, 149, 1,
		0, 0, 0, 22, 157, 1, 0, 0, 0, 24, 167, 1, 0, 0, 0, 26, 169, 1, 0, 0, 0,
		28, 175, 1, 0, 0, 0, 30, 179, 1, 0, 0, 0, 32, 181, 1, 0, 0, 0, 34, 185,
		1, 0, 0, 0, 36, 188, 1, 0, 0, 0, 38, 191, 1, 0, 0, 0, 40, 198, 1, 0, 0,
		0, 42, 201, 1, 0, 0, 0, 44, 208, 1, 0, 0, 0, 46, 211, 1, 0, 0, 0, 48, 218,
		1, 0, 0, 0, 50, 221, 1, 0, 0, 0, 52, 223, 1, 0, 0, 0, 54, 230, 1, 0, 0,
		0, 56, 233, 1, 0, 0, 0, 58, 235, 1, 0, 0, 0, 60, 242, 1, 0, 0, 0, 62, 245,
		1, 0, 0, 0, 64, 247, 1, 0, 0, 0, 66, 254, 1, 0, 0, 0, 68, 257, 1, 0, 0,
		0, 70, 263, 1, 0, 0, 0, 72, 265, 1, 0, 0, 0, 74, 267, 1, 0, 0, 0, 76, 272,
		1, 0, 0, 0, 78, 276, 1, 0, 0, 0, 80, 282, 1, 0, 0, 0, 82, 296, 1, 0, 0,
		0, 84, 298, 1, 0, 0, 0, 86, 88, 3, 2, 1, 0, 87, 86, 1, 0, 0, 0, 88, 91,
		1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 92, 1, 0, 0, 0,
		91, 89, 1, 0, 0, 0, 92, 93, 5, 0, 0, 1, 93, 1, 1, 0, 0, 0, 94, 95, 3, 6,
		3, 0, 95, 97, 5, 34, 0, 0, 96, 98, 3, 4, 2, 0, 97, 96, 1, 0, 0, 0, 97,
		98, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 100, 3, 8, 4, 0, 100, 3, 1, 0,
		0, 0, 101, 102, 5, 1, 0, 0, 102, 103, 3, 30, 15, 0, 103, 104, 5, 2, 0,
		0, 104, 5, 1, 0, 0, 0, 105, 106, 7, 0, 0, 0, 106, 7, 1, 0, 0, 0, 107, 109,
		5, 8, 0, 0, 108, 110, 3, 10, 5, 0, 109, 108, 1, 0, 0, 0, 109, 110, 1, 0,
		0, 0, 110, 111, 1, 0, 0, 0, 111, 112, 5, 9, 0, 0, 112, 118, 3, 14, 7, 0,
		113, 115, 3, 36, 18, 0, 114, 113, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115,
		116, 1, 0, 0, 0, 116, 118, 5, 10, 0, 0, 117, 107, 1, 0, 0, 0, 117, 114,
		1, 0, 0, 0, 118, 9, 1, 0, 0, 0, 119, 124, 3, 12, 6, 0, 120, 121, 5, 11,
		0, 0, 121, 123, 3, 12, 6, 0, 122, 120, 1, 0, 0, 0, 123, 126, 1, 0, 0, 0,
		124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 11, 1, 0, 0, 0, 126, 124,
		1, 0, 0, 0, 127, 128, 3, 6, 3, 0, 128, 130, 5, 34, 0, 0, 129, 131, 3, 4,
		2, 0, 130, 129, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 13, 1, 0, 0, 0,
		132, 136, 5, 12, 0, 0, 133, 135, 3, 16, 8, 0, 134, 133, 1, 0, 0, 0, 135,
		138, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 139,
		1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 139, 140, 5, 13, 0, 0, 140, 15, 1, 0,
		0, 0, 141, 144, 3, 2, 1, 0, 142, 144, 3, 18, 9, 0, 143, 141, 1, 0, 0, 0,
		143, 142, 1, 0, 0, 0, 144, 17, 1, 0, 0, 0, 145, 148, 3, 20, 10, 0, 146,
		148, 3, 24, 12, 0, 147, 145, 1, 0, 0, 0, 147, 146, 1, 0, 0, 0, 148, 19,
		1, 0, 0, 0, 149, 150, 5, 14, 0, 0, 150, 151, 5, 8, 0, 0, 151, 152, 3, 30,
		15, 0, 152, 153, 5, 9, 0, 0, 153, 155, 3, 14, 7, 0, 154, 156, 3, 22, 11,
		0, 155, 154, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 21, 1, 0, 0, 0, 157,
		160, 5, 15, 0, 0, 158, 161, 3, 14, 7, 0, 159, 161, 3, 20, 10, 0, 160, 158,
		1, 0, 0, 0, 160, 159, 1, 0, 0, 0, 161, 23, 1, 0, 0, 0, 162, 163, 3, 30,
		15, 0, 163, 164, 5, 10, 0, 0, 164, 168, 1, 0, 0, 0, 165, 168, 3, 26, 13,
		0, 166, 168, 3, 28, 14, 0, 167, 162, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0,
		167, 166, 1, 0, 0, 0, 168, 25, 1, 0, 0, 0, 169, 170, 5, 16, 0, 0, 170,
		171, 5, 8, 0, 0, 171, 172, 3, 30, 15, 0, 172, 173, 5, 9, 0, 0, 173, 174,
		3, 14, 7, 0, 174, 27, 1, 0, 0, 0, 175, 176, 5, 17, 0, 0, 176, 177, 3, 30,
		15, 0, 177, 178, 5, 10, 0, 0, 178, 29, 1, 0, 0, 0, 179, 180, 3, 32, 16,
		0, 180, 31, 1, 0, 0, 0, 181, 183, 3, 38, 19, 0, 182, 184, 3, 34, 17, 0,
		183, 182, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 33, 1, 0, 0, 0, 185, 186,
		5, 18, 0, 0, 186, 187, 3, 32, 16, 0, 187, 35, 1, 0, 0, 0, 188, 189, 5,
		18, 0, 0, 189, 190, 3, 30, 15, 0, 190, 37, 1, 0, 0, 0, 191, 195, 3, 42,
		21, 0, 192, 194, 3, 40, 20, 0, 193, 192, 1, 0, 0, 0, 194, 197, 1, 0, 0,
		0, 195, 193, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 39, 1, 0, 0, 0, 197,
		195, 1, 0, 0, 0, 198, 199, 5, 19, 0, 0, 199, 200, 3, 42, 21, 0, 200, 41,
		1, 0, 0, 0, 201, 205, 3, 46, 23, 0, 202, 204, 3, 44, 22, 0, 203, 202, 1,
		0, 0, 0, 204, 207, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 205, 206, 1, 0, 0,
		0, 206, 43, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 208, 209, 5, 20, 0, 0, 209,
		210, 3, 46, 23, 0, 210, 45, 1, 0, 0, 0, 211, 215, 3, 52, 26, 0, 212, 214,
		3, 48, 24, 0, 213, 212, 1, 0, 0, 0, 214, 217, 1, 0, 0, 0, 215, 213, 1,
		0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 47, 1, 0, 0, 0, 217, 215, 1, 0, 0,
		0, 218, 219, 3, 50, 25, 0, 219, 220, 3, 52, 26, 0, 220, 49, 1, 0, 0, 0,
		221, 222, 7, 1, 0, 0, 222, 51, 1, 0, 0, 0, 223, 227, 3, 58, 29, 0, 224,
		226, 3, 54, 27, 0, 225, 224, 1, 0, 0, 0, 226, 229, 1, 0, 0, 0, 227, 225,
		1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 53, 1, 0, 0, 0, 229, 227, 1, 0,
		0, 0, 230, 231, 3, 56, 28, 0, 231, 232, 3, 58, 29, 0, 232, 55, 1, 0, 0,
		0, 233, 234, 7, 2, 0, 0, 234, 57, 1, 0, 0, 0, 235, 239, 3, 64, 32, 0, 236,
		238, 3, 60, 30, 0, 237, 236, 1, 0, 0, 0, 238, 241, 1, 0, 0, 0, 239, 237,
		1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 59, 1, 0, 0, 0, 241, 239, 1, 0,
		0, 0, 242, 243, 3, 62, 31, 0, 243, 244, 3, 64, 32, 0, 244, 61, 1, 0, 0,
		0, 245, 246, 7, 3, 0, 0, 246, 63, 1, 0, 0, 0, 247, 251, 3, 70, 35, 0, 248,
		250, 3, 66, 33, 0, 249, 248, 1, 0, 0, 0, 250, 253, 1, 0, 0, 0, 251, 249,
		1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 65, 1, 0, 0, 0, 253, 251, 1, 0,
		0, 0, 254, 255, 3, 68, 34, 0, 255, 256, 3, 70, 35, 0, 256, 67, 1, 0, 0,
		0, 257, 258, 7, 4, 0, 0, 258, 69, 1, 0, 0, 0, 259, 264, 3, 74, 37, 0, 260,
		261, 3, 72, 36, 0, 261, 262, 3, 70, 35, 0, 262, 264, 1, 0, 0, 0, 263, 259,
		1, 0, 0, 0, 263, 260, 1, 0, 0, 0, 264, 71, 1, 0, 0, 0, 265, 266, 7, 5,
		0, 0, 266, 73, 1, 0, 0, 0, 267, 270, 3, 82, 41, 0, 268, 271, 3, 76, 38,
		0, 269, 271, 3, 78, 39, 0, 270, 268, 1, 0, 0, 0, 270, 269, 1, 0, 0, 0,
		270, 271, 1, 0, 0, 0, 271, 75, 1, 0, 0, 0, 272, 273, 5, 1, 0, 0, 273, 274,
		3, 30, 15, 0, 274, 275, 5, 2, 0, 0, 275, 77, 1, 0, 0, 0, 276, 278, 5, 8,
		0, 0, 277, 279, 3, 80, 40, 0, 278, 277, 1, 0, 0, 0, 278, 279, 1, 0, 0,
		0, 279, 280, 1, 0, 0, 0, 280, 281, 5, 9, 0, 0, 281, 79, 1, 0, 0, 0, 282,
		287, 3, 32, 16, 0, 283, 284, 5, 11, 0, 0, 284, 286, 3, 32, 16, 0, 285,
		283, 1, 0, 0, 0, 286, 289, 1, 0, 0, 0, 287, 285, 1, 0, 0, 0, 287, 288,
		1, 0, 0, 0, 288, 81, 1, 0, 0, 0, 289, 287, 1, 0, 0, 0, 290, 297, 5, 34,
		0, 0, 291, 297, 3, 84, 42, 0, 292, 293, 5, 8, 0, 0, 293, 294, 3, 30, 15,
		0, 294, 295, 5, 9, 0, 0, 295, 297, 1, 0, 0, 0, 296, 290, 1, 0, 0, 0, 296,
		291, 1, 0, 0, 0, 296, 292, 1, 0, 0, 0, 297, 83, 1, 0, 0, 0, 298, 299, 7,
		6, 0, 0, 299, 85, 1, 0, 0, 0, 25, 89, 97, 109, 114, 117, 124, 130, 136,
		143, 147, 155, 160, 167, 183, 195, 205, 215, 227, 239, 251, 263, 270, 278,
		287, 296,
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

// BigCParserInit initializes any static state used to implement BigCParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBigCParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BigCParserInit() {
	staticData := &BigCParserStaticData
	staticData.once.Do(bigcParserInit)
}

// NewBigCParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBigCParser(input antlr.TokenStream) *BigCParser {
	BigCParserInit()
	this := new(BigCParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &BigCParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "BigC.g4"

	return this
}

// BigCParser tokens.
const (
	BigCParserEOF               = antlr.TokenEOF
	BigCParserT__0              = 1
	BigCParserT__1              = 2
	BigCParserT__2              = 3
	BigCParserT__3              = 4
	BigCParserT__4              = 5
	BigCParserT__5              = 6
	BigCParserT__6              = 7
	BigCParserT__7              = 8
	BigCParserT__8              = 9
	BigCParserT__9              = 10
	BigCParserT__10             = 11
	BigCParserT__11             = 12
	BigCParserT__12             = 13
	BigCParserT__13             = 14
	BigCParserT__14             = 15
	BigCParserT__15             = 16
	BigCParserT__16             = 17
	BigCParserT__17             = 18
	BigCParserT__18             = 19
	BigCParserT__19             = 20
	BigCParserT__20             = 21
	BigCParserT__21             = 22
	BigCParserT__22             = 23
	BigCParserT__23             = 24
	BigCParserT__24             = 25
	BigCParserT__25             = 26
	BigCParserT__26             = 27
	BigCParserT__27             = 28
	BigCParserT__28             = 29
	BigCParserT__29             = 30
	BigCParserT__30             = 31
	BigCParserBooleanConstant   = 32
	BigCParserCharConstant      = 33
	BigCParserIdentifier        = 34
	BigCParserIntegerConstant   = 35
	BigCParserFloatingConstant  = 36
	BigCParserWS                = 37
	BigCParserCOMMENT           = 38
	BigCParserMULTILINE_COMMENT = 39
)

// BigCParser rules.
const (
	BigCParserRULE_program                      = 0
	BigCParserRULE_declaration                  = 1
	BigCParserRULE_arrayNotation                = 2
	BigCParserRULE_type                         = 3
	BigCParserRULE_declarationRemainder         = 4
	BigCParserRULE_parameterList                = 5
	BigCParserRULE_parameter                    = 6
	BigCParserRULE_block                        = 7
	BigCParserRULE_blockItem                    = 8
	BigCParserRULE_statement                    = 9
	BigCParserRULE_ifStatement                  = 10
	BigCParserRULE_elseClause                   = 11
	BigCParserRULE_nonIfStatement               = 12
	BigCParserRULE_whileStatement               = 13
	BigCParserRULE_returnStatement              = 14
	BigCParserRULE_expression                   = 15
	BigCParserRULE_assignmentExpression         = 16
	BigCParserRULE_assignmentRest               = 17
	BigCParserRULE_variableInitializer          = 18
	BigCParserRULE_logicalOrExpression          = 19
	BigCParserRULE_logicalOrRest                = 20
	BigCParserRULE_logicalAndExpression         = 21
	BigCParserRULE_logicalAndRest               = 22
	BigCParserRULE_equalityExpression           = 23
	BigCParserRULE_equalityRest                 = 24
	BigCParserRULE_equalityOperator             = 25
	BigCParserRULE_comparisonExpression         = 26
	BigCParserRULE_comparisonRest               = 27
	BigCParserRULE_comparisonOperator           = 28
	BigCParserRULE_additionExpression           = 29
	BigCParserRULE_additionExpressionRest       = 30
	BigCParserRULE_addSubtractOperator          = 31
	BigCParserRULE_multiplicationExpression     = 32
	BigCParserRULE_multiplicationExpressionRest = 33
	BigCParserRULE_multDivModOperator           = 34
	BigCParserRULE_unaryExpression              = 35
	BigCParserRULE_unaryOperator                = 36
	BigCParserRULE_postfixExpression            = 37
	BigCParserRULE_arrayAccess                  = 38
	BigCParserRULE_functionCallArgs             = 39
	BigCParserRULE_argList                      = 40
	BigCParserRULE_primaryExpression            = 41
	BigCParserRULE_constant                     = 42
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllDeclaration() []IDeclarationContext
	Declaration(i int) IDeclarationContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(BigCParserEOF, 0)
}

func (s *ProgramContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
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

	return t.(IDeclarationContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BigCParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&248) != 0 {
		{
			p.SetState(86)
			p.Declaration()
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(92)
		p.Match(BigCParserEOF)
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

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	Identifier() antlr.TerminalNode
	DeclarationRemainder() IDeclarationRemainderContext
	ArrayNotation() IArrayNotationContext

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BigCParserIdentifier, 0)
}

func (s *DeclarationContext) DeclarationRemainder() IDeclarationRemainderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationRemainderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationRemainderContext)
}

func (s *DeclarationContext) ArrayNotation() IArrayNotationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayNotationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayNotationContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BigCParserRULE_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Type_()
	}
	{
		p.SetState(95)
		p.Match(BigCParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BigCParserT__0 {
		{
			p.SetState(96)
			p.ArrayNotation()
		}

	}
	{
		p.SetState(99)
		p.DeclarationRemainder()
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

// IArrayNotationContext is an interface to support dynamic dispatch.
type IArrayNotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsArrayNotationContext differentiates from other interfaces.
	IsArrayNotationContext()
}

type ArrayNotationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayNotationContext() *ArrayNotationContext {
	var p = new(ArrayNotationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_arrayNotation
	return p
}

func InitEmptyArrayNotationContext(p *ArrayNotationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_arrayNotation
}

func (*ArrayNotationContext) IsArrayNotationContext() {}

func NewArrayNotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayNotationContext {
	var p = new(ArrayNotationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_arrayNotation

	return p
}

func (s *ArrayNotationContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayNotationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayNotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayNotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayNotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterArrayNotation(s)
	}
}

func (s *ArrayNotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitArrayNotation(s)
	}
}

func (s *ArrayNotationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitArrayNotation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) ArrayNotation() (localctx IArrayNotationContext) {
	localctx = NewArrayNotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BigCParserRULE_arrayNotation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(BigCParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Expression()
	}
	{
		p.SetState(103)
		p.Match(BigCParserT__1)
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BigCParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&248) != 0) {
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

// IDeclarationRemainderContext is an interface to support dynamic dispatch.
type IDeclarationRemainderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	ParameterList() IParameterListContext
	VariableInitializer() IVariableInitializerContext

	// IsDeclarationRemainderContext differentiates from other interfaces.
	IsDeclarationRemainderContext()
}

type DeclarationRemainderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationRemainderContext() *DeclarationRemainderContext {
	var p = new(DeclarationRemainderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_declarationRemainder
	return p
}

func InitEmptyDeclarationRemainderContext(p *DeclarationRemainderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_declarationRemainder
}

func (*DeclarationRemainderContext) IsDeclarationRemainderContext() {}

func NewDeclarationRemainderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationRemainderContext {
	var p = new(DeclarationRemainderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_declarationRemainder

	return p
}

func (s *DeclarationRemainderContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationRemainderContext) Block() IBlockContext {
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

func (s *DeclarationRemainderContext) ParameterList() IParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *DeclarationRemainderContext) VariableInitializer() IVariableInitializerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableInitializerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableInitializerContext)
}

func (s *DeclarationRemainderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationRemainderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationRemainderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterDeclarationRemainder(s)
	}
}

func (s *DeclarationRemainderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitDeclarationRemainder(s)
	}
}

func (s *DeclarationRemainderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitDeclarationRemainder(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) DeclarationRemainder() (localctx IDeclarationRemainderContext) {
	localctx = NewDeclarationRemainderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BigCParserRULE_declarationRemainder)
	var _la int

	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BigCParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.Match(BigCParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&248) != 0 {
			{
				p.SetState(108)
				p.ParameterList()
			}

		}
		{
			p.SetState(111)
			p.Match(BigCParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.Block()
		}

	case BigCParserT__9, BigCParserT__17:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BigCParserT__17 {
			{
				p.SetState(113)
				p.VariableInitializer()
			}

		}
		{
			p.SetState(116)
			p.Match(BigCParserT__9)
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_parameterList
	return p
}

func InitEmptyParameterListContext(p *ParameterListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_parameterList
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *ParameterListContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
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

	return t.(IParameterContext)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (s *ParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitParameterList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BigCParserRULE_parameterList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Parameter()
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BigCParserT__10 {
		{
			p.SetState(120)
			p.Match(BigCParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.Parameter()
		}

		p.SetState(126)
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

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	Identifier() antlr.TerminalNode
	ArrayNotation() IArrayNotationContext

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_parameter
	return p
}

func InitEmptyParameterContext(p *ParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_parameter
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ParameterContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BigCParserIdentifier, 0)
}

func (s *ParameterContext) ArrayNotation() IArrayNotationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayNotationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayNotationContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BigCParserRULE_parameter)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Type_()
	}
	{
		p.SetState(128)
		p.Match(BigCParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BigCParserT__0 {
		{
			p.SetState(129)
			p.ArrayNotation()
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
	AllBlockItem() []IBlockItemContext
	BlockItem(i int) IBlockItemContext

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
	p.RuleIndex = BigCParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllBlockItem() []IBlockItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockItemContext); ok {
			len++
		}
	}

	tst := make([]IBlockItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockItemContext); ok {
			tst[i] = t.(IBlockItemContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) BlockItem(i int) IBlockItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockItemContext); ok {
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

	return t.(IBlockItemContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BigCParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(BigCParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135560118776) != 0 {
		{
			p.SetState(133)
			p.BlockItem()
		}

		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(139)
		p.Match(BigCParserT__12)
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

// IBlockItemContext is an interface to support dynamic dispatch.
type IBlockItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declaration() IDeclarationContext
	Statement() IStatementContext

	// IsBlockItemContext differentiates from other interfaces.
	IsBlockItemContext()
}

type BlockItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockItemContext() *BlockItemContext {
	var p = new(BlockItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_blockItem
	return p
}

func InitEmptyBlockItemContext(p *BlockItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_blockItem
}

func (*BlockItemContext) IsBlockItemContext() {}

func NewBlockItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockItemContext {
	var p = new(BlockItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_blockItem

	return p
}

func (s *BlockItemContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockItemContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *BlockItemContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterBlockItem(s)
	}
}

func (s *BlockItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitBlockItem(s)
	}
}

func (s *BlockItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitBlockItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) BlockItem() (localctx IBlockItemContext) {
	localctx = NewBlockItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BigCParserRULE_blockItem)
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BigCParserT__2, BigCParserT__3, BigCParserT__4, BigCParserT__5, BigCParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.Declaration()
		}

	case BigCParserT__7, BigCParserT__13, BigCParserT__15, BigCParserT__16, BigCParserT__27, BigCParserT__30, BigCParserBooleanConstant, BigCParserCharConstant, BigCParserIdentifier, BigCParserIntegerConstant, BigCParserFloatingConstant:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(142)
			p.Statement()
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfStatement() IIfStatementContext
	NonIfStatement() INonIfStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) NonIfStatement() INonIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INonIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INonIfStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BigCParserRULE_statement)
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BigCParserT__13:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(145)
			p.IfStatement()
		}

	case BigCParserT__7, BigCParserT__15, BigCParserT__16, BigCParserT__27, BigCParserT__30, BigCParserBooleanConstant, BigCParserCharConstant, BigCParserIdentifier, BigCParserIntegerConstant, BigCParserFloatingConstant:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)
			p.NonIfStatement()
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	Block() IBlockContext
	ElseClause() IElseClauseContext

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) Block() IBlockContext {
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

func (s *IfStatementContext) ElseClause() IElseClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseClauseContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BigCParserRULE_ifStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(BigCParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Match(BigCParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Expression()
	}
	{
		p.SetState(152)
		p.Match(BigCParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Block()
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BigCParserT__14 {
		{
			p.SetState(154)
			p.ElseClause()
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

// IElseClauseContext is an interface to support dynamic dispatch.
type IElseClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	IfStatement() IIfStatementContext

	// IsElseClauseContext differentiates from other interfaces.
	IsElseClauseContext()
}

type ElseClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseClauseContext() *ElseClauseContext {
	var p = new(ElseClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_elseClause
	return p
}

func InitEmptyElseClauseContext(p *ElseClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_elseClause
}

func (*ElseClauseContext) IsElseClauseContext() {}

func NewElseClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseClauseContext {
	var p = new(ElseClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_elseClause

	return p
}

func (s *ElseClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseClauseContext) Block() IBlockContext {
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

func (s *ElseClauseContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *ElseClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterElseClause(s)
	}
}

func (s *ElseClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitElseClause(s)
	}
}

func (s *ElseClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitElseClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) ElseClause() (localctx IElseClauseContext) {
	localctx = NewElseClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BigCParserRULE_elseClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(BigCParserT__14)
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
	case BigCParserT__11:
		{
			p.SetState(158)
			p.Block()
		}

	case BigCParserT__13:
		{
			p.SetState(159)
			p.IfStatement()
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INonIfStatementContext is an interface to support dynamic dispatch.
type INonIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	WhileStatement() IWhileStatementContext
	ReturnStatement() IReturnStatementContext

	// IsNonIfStatementContext differentiates from other interfaces.
	IsNonIfStatementContext()
}

type NonIfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNonIfStatementContext() *NonIfStatementContext {
	var p = new(NonIfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_nonIfStatement
	return p
}

func InitEmptyNonIfStatementContext(p *NonIfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_nonIfStatement
}

func (*NonIfStatementContext) IsNonIfStatementContext() {}

func NewNonIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NonIfStatementContext {
	var p = new(NonIfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_nonIfStatement

	return p
}

func (s *NonIfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *NonIfStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NonIfStatementContext) WhileStatement() IWhileStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *NonIfStatementContext) ReturnStatement() IReturnStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *NonIfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonIfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NonIfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterNonIfStatement(s)
	}
}

func (s *NonIfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitNonIfStatement(s)
	}
}

func (s *NonIfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitNonIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) NonIfStatement() (localctx INonIfStatementContext) {
	localctx = NewNonIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BigCParserRULE_nonIfStatement)
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BigCParserT__7, BigCParserT__27, BigCParserT__30, BigCParserBooleanConstant, BigCParserCharConstant, BigCParserIdentifier, BigCParserIntegerConstant, BigCParserFloatingConstant:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.Expression()
		}
		{
			p.SetState(163)
			p.Match(BigCParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BigCParserT__15:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.WhileStatement()
		}

	case BigCParserT__16:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(166)
			p.ReturnStatement()
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	Block() IBlockContext

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_whileStatement
	return p
}

func InitEmptyWhileStatementContext(p *WhileStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_whileStatement
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) Block() IBlockContext {
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

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BigCParserRULE_whileStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(BigCParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(BigCParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Expression()
	}
	{
		p.SetState(172)
		p.Match(BigCParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
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

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_returnStatement
	return p
}

func InitEmptyReturnStatementContext(p *ReturnStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_returnStatement
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BigCParserRULE_returnStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(BigCParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Expression()
	}
	{
		p.SetState(177)
		p.Match(BigCParserT__9)
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AssignmentExpression() IAssignmentExpressionContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AssignmentExpression() IAssignmentExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BigCParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.AssignmentExpression()
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

// IAssignmentExpressionContext is an interface to support dynamic dispatch.
type IAssignmentExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LogicalOrExpression() ILogicalOrExpressionContext
	AssignmentRest() IAssignmentRestContext

	// IsAssignmentExpressionContext differentiates from other interfaces.
	IsAssignmentExpressionContext()
}

type AssignmentExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentExpressionContext() *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_assignmentExpression
	return p
}

func InitEmptyAssignmentExpressionContext(p *AssignmentExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_assignmentExpression
}

func (*AssignmentExpressionContext) IsAssignmentExpressionContext() {}

func NewAssignmentExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_assignmentExpression

	return p
}

func (s *AssignmentExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentExpressionContext) LogicalOrExpression() ILogicalOrExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOrExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalOrExpressionContext)
}

func (s *AssignmentExpressionContext) AssignmentRest() IAssignmentRestContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentRestContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentRestContext)
}

func (s *AssignmentExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterAssignmentExpression(s)
	}
}

func (s *AssignmentExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitAssignmentExpression(s)
	}
}

func (s *AssignmentExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitAssignmentExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) AssignmentExpression() (localctx IAssignmentExpressionContext) {
	localctx = NewAssignmentExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BigCParserRULE_assignmentExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.LogicalOrExpression()
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BigCParserT__17 {
		{
			p.SetState(182)
			p.AssignmentRest()
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

// IAssignmentRestContext is an interface to support dynamic dispatch.
type IAssignmentRestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AssignmentExpression() IAssignmentExpressionContext

	// IsAssignmentRestContext differentiates from other interfaces.
	IsAssignmentRestContext()
}

type AssignmentRestContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentRestContext() *AssignmentRestContext {
	var p = new(AssignmentRestContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_assignmentRest
	return p
}

func InitEmptyAssignmentRestContext(p *AssignmentRestContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_assignmentRest
}

func (*AssignmentRestContext) IsAssignmentRestContext() {}

func NewAssignmentRestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentRestContext {
	var p = new(AssignmentRestContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_assignmentRest

	return p
}

func (s *AssignmentRestContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentRestContext) AssignmentExpression() IAssignmentExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentExpressionContext)
}

func (s *AssignmentRestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentRestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentRestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterAssignmentRest(s)
	}
}

func (s *AssignmentRestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitAssignmentRest(s)
	}
}

func (s *AssignmentRestContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitAssignmentRest(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) AssignmentRest() (localctx IAssignmentRestContext) {
	localctx = NewAssignmentRestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BigCParserRULE_assignmentRest)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Match(BigCParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.AssignmentExpression()
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

// IVariableInitializerContext is an interface to support dynamic dispatch.
type IVariableInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsVariableInitializerContext differentiates from other interfaces.
	IsVariableInitializerContext()
}

type VariableInitializerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableInitializerContext() *VariableInitializerContext {
	var p = new(VariableInitializerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_variableInitializer
	return p
}

func InitEmptyVariableInitializerContext(p *VariableInitializerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_variableInitializer
}

func (*VariableInitializerContext) IsVariableInitializerContext() {}

func NewVariableInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableInitializerContext {
	var p = new(VariableInitializerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_variableInitializer

	return p
}

func (s *VariableInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableInitializerContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableInitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterVariableInitializer(s)
	}
}

func (s *VariableInitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitVariableInitializer(s)
	}
}

func (s *VariableInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitVariableInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) VariableInitializer() (localctx IVariableInitializerContext) {
	localctx = NewVariableInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BigCParserRULE_variableInitializer)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(BigCParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(189)
		p.Expression()
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

// ILogicalOrExpressionContext is an interface to support dynamic dispatch.
type ILogicalOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LogicalAndExpression() ILogicalAndExpressionContext
	AllLogicalOrRest() []ILogicalOrRestContext
	LogicalOrRest(i int) ILogicalOrRestContext

	// IsLogicalOrExpressionContext differentiates from other interfaces.
	IsLogicalOrExpressionContext()
}

type LogicalOrExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOrExpressionContext() *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_logicalOrExpression
	return p
}

func InitEmptyLogicalOrExpressionContext(p *LogicalOrExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_logicalOrExpression
}

func (*LogicalOrExpressionContext) IsLogicalOrExpressionContext() {}

func NewLogicalOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_logicalOrExpression

	return p
}

func (s *LogicalOrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOrExpressionContext) LogicalAndExpression() ILogicalAndExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalAndExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalAndExpressionContext)
}

func (s *LogicalOrExpressionContext) AllLogicalOrRest() []ILogicalOrRestContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogicalOrRestContext); ok {
			len++
		}
	}

	tst := make([]ILogicalOrRestContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogicalOrRestContext); ok {
			tst[i] = t.(ILogicalOrRestContext)
			i++
		}
	}

	return tst
}

func (s *LogicalOrExpressionContext) LogicalOrRest(i int) ILogicalOrRestContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOrRestContext); ok {
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

	return t.(ILogicalOrRestContext)
}

func (s *LogicalOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterLogicalOrExpression(s)
	}
}

func (s *LogicalOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitLogicalOrExpression(s)
	}
}

func (s *LogicalOrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitLogicalOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) LogicalOrExpression() (localctx ILogicalOrExpressionContext) {
	localctx = NewLogicalOrExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BigCParserRULE_logicalOrExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.LogicalAndExpression()
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BigCParserT__18 {
		{
			p.SetState(192)
			p.LogicalOrRest()
		}

		p.SetState(197)
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

// ILogicalOrRestContext is an interface to support dynamic dispatch.
type ILogicalOrRestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LogicalAndExpression() ILogicalAndExpressionContext

	// IsLogicalOrRestContext differentiates from other interfaces.
	IsLogicalOrRestContext()
}

type LogicalOrRestContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOrRestContext() *LogicalOrRestContext {
	var p = new(LogicalOrRestContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_logicalOrRest
	return p
}

func InitEmptyLogicalOrRestContext(p *LogicalOrRestContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_logicalOrRest
}

func (*LogicalOrRestContext) IsLogicalOrRestContext() {}

func NewLogicalOrRestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOrRestContext {
	var p = new(LogicalOrRestContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_logicalOrRest

	return p
}

func (s *LogicalOrRestContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOrRestContext) LogicalAndExpression() ILogicalAndExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalAndExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalAndExpressionContext)
}

func (s *LogicalOrRestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrRestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOrRestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterLogicalOrRest(s)
	}
}

func (s *LogicalOrRestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitLogicalOrRest(s)
	}
}

func (s *LogicalOrRestContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitLogicalOrRest(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) LogicalOrRest() (localctx ILogicalOrRestContext) {
	localctx = NewLogicalOrRestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, BigCParserRULE_logicalOrRest)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(BigCParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.LogicalAndExpression()
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

// ILogicalAndExpressionContext is an interface to support dynamic dispatch.
type ILogicalAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EqualityExpression() IEqualityExpressionContext
	AllLogicalAndRest() []ILogicalAndRestContext
	LogicalAndRest(i int) ILogicalAndRestContext

	// IsLogicalAndExpressionContext differentiates from other interfaces.
	IsLogicalAndExpressionContext()
}

type LogicalAndExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalAndExpressionContext() *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_logicalAndExpression
	return p
}

func InitEmptyLogicalAndExpressionContext(p *LogicalAndExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_logicalAndExpression
}

func (*LogicalAndExpressionContext) IsLogicalAndExpressionContext() {}

func NewLogicalAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_logicalAndExpression

	return p
}

func (s *LogicalAndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalAndExpressionContext) EqualityExpression() IEqualityExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionContext)
}

func (s *LogicalAndExpressionContext) AllLogicalAndRest() []ILogicalAndRestContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogicalAndRestContext); ok {
			len++
		}
	}

	tst := make([]ILogicalAndRestContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogicalAndRestContext); ok {
			tst[i] = t.(ILogicalAndRestContext)
			i++
		}
	}

	return tst
}

func (s *LogicalAndExpressionContext) LogicalAndRest(i int) ILogicalAndRestContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalAndRestContext); ok {
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

	return t.(ILogicalAndRestContext)
}

func (s *LogicalAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterLogicalAndExpression(s)
	}
}

func (s *LogicalAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitLogicalAndExpression(s)
	}
}

func (s *LogicalAndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitLogicalAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) LogicalAndExpression() (localctx ILogicalAndExpressionContext) {
	localctx = NewLogicalAndExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BigCParserRULE_logicalAndExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.EqualityExpression()
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BigCParserT__19 {
		{
			p.SetState(202)
			p.LogicalAndRest()
		}

		p.SetState(207)
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

// ILogicalAndRestContext is an interface to support dynamic dispatch.
type ILogicalAndRestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EqualityExpression() IEqualityExpressionContext

	// IsLogicalAndRestContext differentiates from other interfaces.
	IsLogicalAndRestContext()
}

type LogicalAndRestContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalAndRestContext() *LogicalAndRestContext {
	var p = new(LogicalAndRestContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_logicalAndRest
	return p
}

func InitEmptyLogicalAndRestContext(p *LogicalAndRestContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_logicalAndRest
}

func (*LogicalAndRestContext) IsLogicalAndRestContext() {}

func NewLogicalAndRestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalAndRestContext {
	var p = new(LogicalAndRestContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_logicalAndRest

	return p
}

func (s *LogicalAndRestContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalAndRestContext) EqualityExpression() IEqualityExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionContext)
}

func (s *LogicalAndRestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndRestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalAndRestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterLogicalAndRest(s)
	}
}

func (s *LogicalAndRestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitLogicalAndRest(s)
	}
}

func (s *LogicalAndRestContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitLogicalAndRest(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) LogicalAndRest() (localctx ILogicalAndRestContext) {
	localctx = NewLogicalAndRestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, BigCParserRULE_logicalAndRest)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(BigCParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.EqualityExpression()
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

// IEqualityExpressionContext is an interface to support dynamic dispatch.
type IEqualityExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ComparisonExpression() IComparisonExpressionContext
	AllEqualityRest() []IEqualityRestContext
	EqualityRest(i int) IEqualityRestContext

	// IsEqualityExpressionContext differentiates from other interfaces.
	IsEqualityExpressionContext()
}

type EqualityExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityExpressionContext() *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_equalityExpression
	return p
}

func InitEmptyEqualityExpressionContext(p *EqualityExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_equalityExpression
}

func (*EqualityExpressionContext) IsEqualityExpressionContext() {}

func NewEqualityExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_equalityExpression

	return p
}

func (s *EqualityExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityExpressionContext) ComparisonExpression() IComparisonExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonExpressionContext)
}

func (s *EqualityExpressionContext) AllEqualityRest() []IEqualityRestContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEqualityRestContext); ok {
			len++
		}
	}

	tst := make([]IEqualityRestContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEqualityRestContext); ok {
			tst[i] = t.(IEqualityRestContext)
			i++
		}
	}

	return tst
}

func (s *EqualityExpressionContext) EqualityRest(i int) IEqualityRestContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityRestContext); ok {
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

	return t.(IEqualityRestContext)
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitEqualityExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) EqualityExpression() (localctx IEqualityExpressionContext) {
	localctx = NewEqualityExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, BigCParserRULE_equalityExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.ComparisonExpression()
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BigCParserT__20 || _la == BigCParserT__21 {
		{
			p.SetState(212)
			p.EqualityRest()
		}

		p.SetState(217)
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

// IEqualityRestContext is an interface to support dynamic dispatch.
type IEqualityRestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EqualityOperator() IEqualityOperatorContext
	ComparisonExpression() IComparisonExpressionContext

	// IsEqualityRestContext differentiates from other interfaces.
	IsEqualityRestContext()
}

type EqualityRestContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityRestContext() *EqualityRestContext {
	var p = new(EqualityRestContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_equalityRest
	return p
}

func InitEmptyEqualityRestContext(p *EqualityRestContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_equalityRest
}

func (*EqualityRestContext) IsEqualityRestContext() {}

func NewEqualityRestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityRestContext {
	var p = new(EqualityRestContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_equalityRest

	return p
}

func (s *EqualityRestContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityRestContext) EqualityOperator() IEqualityOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualityOperatorContext)
}

func (s *EqualityRestContext) ComparisonExpression() IComparisonExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonExpressionContext)
}

func (s *EqualityRestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityRestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityRestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterEqualityRest(s)
	}
}

func (s *EqualityRestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitEqualityRest(s)
	}
}

func (s *EqualityRestContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitEqualityRest(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) EqualityRest() (localctx IEqualityRestContext) {
	localctx = NewEqualityRestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, BigCParserRULE_equalityRest)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.EqualityOperator()
	}
	{
		p.SetState(219)
		p.ComparisonExpression()
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

// IEqualityOperatorContext is an interface to support dynamic dispatch.
type IEqualityOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsEqualityOperatorContext differentiates from other interfaces.
	IsEqualityOperatorContext()
}

type EqualityOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityOperatorContext() *EqualityOperatorContext {
	var p = new(EqualityOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_equalityOperator
	return p
}

func InitEmptyEqualityOperatorContext(p *EqualityOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_equalityOperator
}

func (*EqualityOperatorContext) IsEqualityOperatorContext() {}

func NewEqualityOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityOperatorContext {
	var p = new(EqualityOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_equalityOperator

	return p
}

func (s *EqualityOperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *EqualityOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterEqualityOperator(s)
	}
}

func (s *EqualityOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitEqualityOperator(s)
	}
}

func (s *EqualityOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitEqualityOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) EqualityOperator() (localctx IEqualityOperatorContext) {
	localctx = NewEqualityOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, BigCParserRULE_equalityOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BigCParserT__20 || _la == BigCParserT__21) {
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

// IComparisonExpressionContext is an interface to support dynamic dispatch.
type IComparisonExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AdditionExpression() IAdditionExpressionContext
	AllComparisonRest() []IComparisonRestContext
	ComparisonRest(i int) IComparisonRestContext

	// IsComparisonExpressionContext differentiates from other interfaces.
	IsComparisonExpressionContext()
}

type ComparisonExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonExpressionContext() *ComparisonExpressionContext {
	var p = new(ComparisonExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_comparisonExpression
	return p
}

func InitEmptyComparisonExpressionContext(p *ComparisonExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_comparisonExpression
}

func (*ComparisonExpressionContext) IsComparisonExpressionContext() {}

func NewComparisonExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonExpressionContext {
	var p = new(ComparisonExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_comparisonExpression

	return p
}

func (s *ComparisonExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonExpressionContext) AdditionExpression() IAdditionExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditionExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditionExpressionContext)
}

func (s *ComparisonExpressionContext) AllComparisonRest() []IComparisonRestContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComparisonRestContext); ok {
			len++
		}
	}

	tst := make([]IComparisonRestContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComparisonRestContext); ok {
			tst[i] = t.(IComparisonRestContext)
			i++
		}
	}

	return tst
}

func (s *ComparisonExpressionContext) ComparisonRest(i int) IComparisonRestContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonRestContext); ok {
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

	return t.(IComparisonRestContext)
}

func (s *ComparisonExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterComparisonExpression(s)
	}
}

func (s *ComparisonExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitComparisonExpression(s)
	}
}

func (s *ComparisonExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitComparisonExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) ComparisonExpression() (localctx IComparisonExpressionContext) {
	localctx = NewComparisonExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, BigCParserRULE_comparisonExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.AdditionExpression()
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&125829120) != 0 {
		{
			p.SetState(224)
			p.ComparisonRest()
		}

		p.SetState(229)
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

// IComparisonRestContext is an interface to support dynamic dispatch.
type IComparisonRestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ComparisonOperator() IComparisonOperatorContext
	AdditionExpression() IAdditionExpressionContext

	// IsComparisonRestContext differentiates from other interfaces.
	IsComparisonRestContext()
}

type ComparisonRestContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonRestContext() *ComparisonRestContext {
	var p = new(ComparisonRestContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_comparisonRest
	return p
}

func InitEmptyComparisonRestContext(p *ComparisonRestContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_comparisonRest
}

func (*ComparisonRestContext) IsComparisonRestContext() {}

func NewComparisonRestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonRestContext {
	var p = new(ComparisonRestContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_comparisonRest

	return p
}

func (s *ComparisonRestContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonRestContext) ComparisonOperator() IComparisonOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *ComparisonRestContext) AdditionExpression() IAdditionExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditionExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditionExpressionContext)
}

func (s *ComparisonRestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonRestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonRestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterComparisonRest(s)
	}
}

func (s *ComparisonRestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitComparisonRest(s)
	}
}

func (s *ComparisonRestContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitComparisonRest(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) ComparisonRest() (localctx IComparisonRestContext) {
	localctx = NewComparisonRestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, BigCParserRULE_comparisonRest)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.ComparisonOperator()
	}
	{
		p.SetState(231)
		p.AdditionExpression()
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

// IComparisonOperatorContext is an interface to support dynamic dispatch.
type IComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsComparisonOperatorContext differentiates from other interfaces.
	IsComparisonOperatorContext()
}

type ComparisonOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_comparisonOperator
	return p
}

func InitEmptyComparisonOperatorContext(p *ComparisonOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_comparisonOperator
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitComparisonOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, BigCParserRULE_comparisonOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&125829120) != 0) {
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

// IAdditionExpressionContext is an interface to support dynamic dispatch.
type IAdditionExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MultiplicationExpression() IMultiplicationExpressionContext
	AllAdditionExpressionRest() []IAdditionExpressionRestContext
	AdditionExpressionRest(i int) IAdditionExpressionRestContext

	// IsAdditionExpressionContext differentiates from other interfaces.
	IsAdditionExpressionContext()
}

type AdditionExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditionExpressionContext() *AdditionExpressionContext {
	var p = new(AdditionExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_additionExpression
	return p
}

func InitEmptyAdditionExpressionContext(p *AdditionExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_additionExpression
}

func (*AdditionExpressionContext) IsAdditionExpressionContext() {}

func NewAdditionExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditionExpressionContext {
	var p = new(AdditionExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_additionExpression

	return p
}

func (s *AdditionExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditionExpressionContext) MultiplicationExpression() IMultiplicationExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicationExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicationExpressionContext)
}

func (s *AdditionExpressionContext) AllAdditionExpressionRest() []IAdditionExpressionRestContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAdditionExpressionRestContext); ok {
			len++
		}
	}

	tst := make([]IAdditionExpressionRestContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAdditionExpressionRestContext); ok {
			tst[i] = t.(IAdditionExpressionRestContext)
			i++
		}
	}

	return tst
}

func (s *AdditionExpressionContext) AdditionExpressionRest(i int) IAdditionExpressionRestContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditionExpressionRestContext); ok {
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

	return t.(IAdditionExpressionRestContext)
}

func (s *AdditionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditionExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterAdditionExpression(s)
	}
}

func (s *AdditionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitAdditionExpression(s)
	}
}

func (s *AdditionExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitAdditionExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) AdditionExpression() (localctx IAdditionExpressionContext) {
	localctx = NewAdditionExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, BigCParserRULE_additionExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.MultiplicationExpression()
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BigCParserT__26 || _la == BigCParserT__27 {
		{
			p.SetState(236)
			p.AdditionExpressionRest()
		}

		p.SetState(241)
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

// IAdditionExpressionRestContext is an interface to support dynamic dispatch.
type IAdditionExpressionRestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AddSubtractOperator() IAddSubtractOperatorContext
	MultiplicationExpression() IMultiplicationExpressionContext

	// IsAdditionExpressionRestContext differentiates from other interfaces.
	IsAdditionExpressionRestContext()
}

type AdditionExpressionRestContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditionExpressionRestContext() *AdditionExpressionRestContext {
	var p = new(AdditionExpressionRestContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_additionExpressionRest
	return p
}

func InitEmptyAdditionExpressionRestContext(p *AdditionExpressionRestContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_additionExpressionRest
}

func (*AdditionExpressionRestContext) IsAdditionExpressionRestContext() {}

func NewAdditionExpressionRestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditionExpressionRestContext {
	var p = new(AdditionExpressionRestContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_additionExpressionRest

	return p
}

func (s *AdditionExpressionRestContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditionExpressionRestContext) AddSubtractOperator() IAddSubtractOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddSubtractOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddSubtractOperatorContext)
}

func (s *AdditionExpressionRestContext) MultiplicationExpression() IMultiplicationExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicationExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicationExpressionContext)
}

func (s *AdditionExpressionRestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditionExpressionRestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditionExpressionRestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterAdditionExpressionRest(s)
	}
}

func (s *AdditionExpressionRestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitAdditionExpressionRest(s)
	}
}

func (s *AdditionExpressionRestContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitAdditionExpressionRest(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) AdditionExpressionRest() (localctx IAdditionExpressionRestContext) {
	localctx = NewAdditionExpressionRestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, BigCParserRULE_additionExpressionRest)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.AddSubtractOperator()
	}
	{
		p.SetState(243)
		p.MultiplicationExpression()
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

// IAddSubtractOperatorContext is an interface to support dynamic dispatch.
type IAddSubtractOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAddSubtractOperatorContext differentiates from other interfaces.
	IsAddSubtractOperatorContext()
}

type AddSubtractOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddSubtractOperatorContext() *AddSubtractOperatorContext {
	var p = new(AddSubtractOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_addSubtractOperator
	return p
}

func InitEmptyAddSubtractOperatorContext(p *AddSubtractOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_addSubtractOperator
}

func (*AddSubtractOperatorContext) IsAddSubtractOperatorContext() {}

func NewAddSubtractOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddSubtractOperatorContext {
	var p = new(AddSubtractOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_addSubtractOperator

	return p
}

func (s *AddSubtractOperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *AddSubtractOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubtractOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddSubtractOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterAddSubtractOperator(s)
	}
}

func (s *AddSubtractOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitAddSubtractOperator(s)
	}
}

func (s *AddSubtractOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitAddSubtractOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) AddSubtractOperator() (localctx IAddSubtractOperatorContext) {
	localctx = NewAddSubtractOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, BigCParserRULE_addSubtractOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BigCParserT__26 || _la == BigCParserT__27) {
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

// IMultiplicationExpressionContext is an interface to support dynamic dispatch.
type IMultiplicationExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryExpression() IUnaryExpressionContext
	AllMultiplicationExpressionRest() []IMultiplicationExpressionRestContext
	MultiplicationExpressionRest(i int) IMultiplicationExpressionRestContext

	// IsMultiplicationExpressionContext differentiates from other interfaces.
	IsMultiplicationExpressionContext()
}

type MultiplicationExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicationExpressionContext() *MultiplicationExpressionContext {
	var p = new(MultiplicationExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_multiplicationExpression
	return p
}

func InitEmptyMultiplicationExpressionContext(p *MultiplicationExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_multiplicationExpression
}

func (*MultiplicationExpressionContext) IsMultiplicationExpressionContext() {}

func NewMultiplicationExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicationExpressionContext {
	var p = new(MultiplicationExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_multiplicationExpression

	return p
}

func (s *MultiplicationExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicationExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *MultiplicationExpressionContext) AllMultiplicationExpressionRest() []IMultiplicationExpressionRestContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicationExpressionRestContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicationExpressionRestContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicationExpressionRestContext); ok {
			tst[i] = t.(IMultiplicationExpressionRestContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicationExpressionContext) MultiplicationExpressionRest(i int) IMultiplicationExpressionRestContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicationExpressionRestContext); ok {
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

	return t.(IMultiplicationExpressionRestContext)
}

func (s *MultiplicationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicationExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicationExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterMultiplicationExpression(s)
	}
}

func (s *MultiplicationExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitMultiplicationExpression(s)
	}
}

func (s *MultiplicationExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitMultiplicationExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) MultiplicationExpression() (localctx IMultiplicationExpressionContext) {
	localctx = NewMultiplicationExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, BigCParserRULE_multiplicationExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.UnaryExpression()
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BigCParserT__28 || _la == BigCParserT__29 {
		{
			p.SetState(248)
			p.MultiplicationExpressionRest()
		}

		p.SetState(253)
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

// IMultiplicationExpressionRestContext is an interface to support dynamic dispatch.
type IMultiplicationExpressionRestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MultDivModOperator() IMultDivModOperatorContext
	UnaryExpression() IUnaryExpressionContext

	// IsMultiplicationExpressionRestContext differentiates from other interfaces.
	IsMultiplicationExpressionRestContext()
}

type MultiplicationExpressionRestContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicationExpressionRestContext() *MultiplicationExpressionRestContext {
	var p = new(MultiplicationExpressionRestContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_multiplicationExpressionRest
	return p
}

func InitEmptyMultiplicationExpressionRestContext(p *MultiplicationExpressionRestContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_multiplicationExpressionRest
}

func (*MultiplicationExpressionRestContext) IsMultiplicationExpressionRestContext() {}

func NewMultiplicationExpressionRestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicationExpressionRestContext {
	var p = new(MultiplicationExpressionRestContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_multiplicationExpressionRest

	return p
}

func (s *MultiplicationExpressionRestContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicationExpressionRestContext) MultDivModOperator() IMultDivModOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultDivModOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultDivModOperatorContext)
}

func (s *MultiplicationExpressionRestContext) UnaryExpression() IUnaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *MultiplicationExpressionRestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicationExpressionRestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicationExpressionRestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterMultiplicationExpressionRest(s)
	}
}

func (s *MultiplicationExpressionRestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitMultiplicationExpressionRest(s)
	}
}

func (s *MultiplicationExpressionRestContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitMultiplicationExpressionRest(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) MultiplicationExpressionRest() (localctx IMultiplicationExpressionRestContext) {
	localctx = NewMultiplicationExpressionRestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, BigCParserRULE_multiplicationExpressionRest)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.MultDivModOperator()
	}
	{
		p.SetState(255)
		p.UnaryExpression()
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

// IMultDivModOperatorContext is an interface to support dynamic dispatch.
type IMultDivModOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMultDivModOperatorContext differentiates from other interfaces.
	IsMultDivModOperatorContext()
}

type MultDivModOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultDivModOperatorContext() *MultDivModOperatorContext {
	var p = new(MultDivModOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_multDivModOperator
	return p
}

func InitEmptyMultDivModOperatorContext(p *MultDivModOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_multDivModOperator
}

func (*MultDivModOperatorContext) IsMultDivModOperatorContext() {}

func NewMultDivModOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultDivModOperatorContext {
	var p = new(MultDivModOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_multDivModOperator

	return p
}

func (s *MultDivModOperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *MultDivModOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultDivModOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultDivModOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterMultDivModOperator(s)
	}
}

func (s *MultDivModOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitMultDivModOperator(s)
	}
}

func (s *MultDivModOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitMultDivModOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) MultDivModOperator() (localctx IMultDivModOperatorContext) {
	localctx = NewMultDivModOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, BigCParserRULE_multDivModOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BigCParserT__28 || _la == BigCParserT__29) {
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

// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PostfixExpression() IPostfixExpressionContext
	UnaryOperator() IUnaryOperatorContext
	UnaryExpression() IUnaryExpressionContext

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}

type UnaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_unaryExpression
	return p
}

func InitEmptyUnaryExpressionContext(p *UnaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_unaryExpression
}

func (*UnaryExpressionContext) IsUnaryExpressionContext() {}

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_unaryExpression

	return p
}

func (s *UnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionContext) PostfixExpression() IPostfixExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfixExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfixExpressionContext)
}

func (s *UnaryExpressionContext) UnaryOperator() IUnaryOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryOperatorContext)
}

func (s *UnaryExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitUnaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) UnaryExpression() (localctx IUnaryExpressionContext) {
	localctx = NewUnaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, BigCParserRULE_unaryExpression)
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BigCParserT__7, BigCParserBooleanConstant, BigCParserCharConstant, BigCParserIdentifier, BigCParserIntegerConstant, BigCParserFloatingConstant:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(259)
			p.PostfixExpression()
		}

	case BigCParserT__27, BigCParserT__30:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(260)
			p.UnaryOperator()
		}
		{
			p.SetState(261)
			p.UnaryExpression()
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryOperatorContext is an interface to support dynamic dispatch.
type IUnaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsUnaryOperatorContext differentiates from other interfaces.
	IsUnaryOperatorContext()
}

type UnaryOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOperatorContext() *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_unaryOperator
	return p
}

func InitEmptyUnaryOperatorContext(p *UnaryOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_unaryOperator
}

func (*UnaryOperatorContext) IsUnaryOperatorContext() {}

func NewUnaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_unaryOperator

	return p
}

func (s *UnaryOperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *UnaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterUnaryOperator(s)
	}
}

func (s *UnaryOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitUnaryOperator(s)
	}
}

func (s *UnaryOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitUnaryOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) UnaryOperator() (localctx IUnaryOperatorContext) {
	localctx = NewUnaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, BigCParserRULE_unaryOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BigCParserT__27 || _la == BigCParserT__30) {
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

// IPostfixExpressionContext is an interface to support dynamic dispatch.
type IPostfixExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryExpression() IPrimaryExpressionContext
	ArrayAccess() IArrayAccessContext
	FunctionCallArgs() IFunctionCallArgsContext

	// IsPostfixExpressionContext differentiates from other interfaces.
	IsPostfixExpressionContext()
}

type PostfixExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostfixExpressionContext() *PostfixExpressionContext {
	var p = new(PostfixExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_postfixExpression
	return p
}

func InitEmptyPostfixExpressionContext(p *PostfixExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_postfixExpression
}

func (*PostfixExpressionContext) IsPostfixExpressionContext() {}

func NewPostfixExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostfixExpressionContext {
	var p = new(PostfixExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_postfixExpression

	return p
}

func (s *PostfixExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PostfixExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PostfixExpressionContext) ArrayAccess() IArrayAccessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayAccessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayAccessContext)
}

func (s *PostfixExpressionContext) FunctionCallArgs() IFunctionCallArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallArgsContext)
}

func (s *PostfixExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostfixExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterPostfixExpression(s)
	}
}

func (s *PostfixExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitPostfixExpression(s)
	}
}

func (s *PostfixExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitPostfixExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) PostfixExpression() (localctx IPostfixExpressionContext) {
	localctx = NewPostfixExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, BigCParserRULE_postfixExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		p.PrimaryExpression()
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case BigCParserT__0:
		{
			p.SetState(268)
			p.ArrayAccess()
		}

	case BigCParserT__7:
		{
			p.SetState(269)
			p.FunctionCallArgs()
		}

	case BigCParserT__1, BigCParserT__8, BigCParserT__9, BigCParserT__10, BigCParserT__17, BigCParserT__18, BigCParserT__19, BigCParserT__20, BigCParserT__21, BigCParserT__22, BigCParserT__23, BigCParserT__24, BigCParserT__25, BigCParserT__26, BigCParserT__27, BigCParserT__28, BigCParserT__29:

	default:
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

// IArrayAccessContext is an interface to support dynamic dispatch.
type IArrayAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsArrayAccessContext differentiates from other interfaces.
	IsArrayAccessContext()
}

type ArrayAccessContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayAccessContext() *ArrayAccessContext {
	var p = new(ArrayAccessContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_arrayAccess
	return p
}

func InitEmptyArrayAccessContext(p *ArrayAccessContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_arrayAccess
}

func (*ArrayAccessContext) IsArrayAccessContext() {}

func NewArrayAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayAccessContext {
	var p = new(ArrayAccessContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_arrayAccess

	return p
}

func (s *ArrayAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayAccessContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterArrayAccess(s)
	}
}

func (s *ArrayAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitArrayAccess(s)
	}
}

func (s *ArrayAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitArrayAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) ArrayAccess() (localctx IArrayAccessContext) {
	localctx = NewArrayAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, BigCParserRULE_arrayAccess)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.Match(BigCParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(273)
		p.Expression()
	}
	{
		p.SetState(274)
		p.Match(BigCParserT__1)
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

// IFunctionCallArgsContext is an interface to support dynamic dispatch.
type IFunctionCallArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ArgList() IArgListContext

	// IsFunctionCallArgsContext differentiates from other interfaces.
	IsFunctionCallArgsContext()
}

type FunctionCallArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallArgsContext() *FunctionCallArgsContext {
	var p = new(FunctionCallArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_functionCallArgs
	return p
}

func InitEmptyFunctionCallArgsContext(p *FunctionCallArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_functionCallArgs
}

func (*FunctionCallArgsContext) IsFunctionCallArgsContext() {}

func NewFunctionCallArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallArgsContext {
	var p = new(FunctionCallArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_functionCallArgs

	return p
}

func (s *FunctionCallArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallArgsContext) ArgList() IArgListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *FunctionCallArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterFunctionCallArgs(s)
	}
}

func (s *FunctionCallArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitFunctionCallArgs(s)
	}
}

func (s *FunctionCallArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitFunctionCallArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) FunctionCallArgs() (localctx IFunctionCallArgsContext) {
	localctx = NewFunctionCallArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, BigCParserRULE_functionCallArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Match(BigCParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135559905536) != 0 {
		{
			p.SetState(277)
			p.ArgList()
		}

	}
	{
		p.SetState(280)
		p.Match(BigCParserT__8)
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

// IArgListContext is an interface to support dynamic dispatch.
type IArgListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAssignmentExpression() []IAssignmentExpressionContext
	AssignmentExpression(i int) IAssignmentExpressionContext

	// IsArgListContext differentiates from other interfaces.
	IsArgListContext()
}

type ArgListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgListContext() *ArgListContext {
	var p = new(ArgListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_argList
	return p
}

func InitEmptyArgListContext(p *ArgListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_argList
}

func (*ArgListContext) IsArgListContext() {}

func NewArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgListContext {
	var p = new(ArgListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_argList

	return p
}

func (s *ArgListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgListContext) AllAssignmentExpression() []IAssignmentExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentExpressionContext); ok {
			tst[i] = t.(IAssignmentExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgListContext) AssignmentExpression(i int) IAssignmentExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
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

	return t.(IAssignmentExpressionContext)
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitArgList(s)
	}
}

func (s *ArgListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitArgList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) ArgList() (localctx IArgListContext) {
	localctx = NewArgListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, BigCParserRULE_argList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.AssignmentExpression()
	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BigCParserT__10 {
		{
			p.SetState(283)
			p.Match(BigCParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
			p.AssignmentExpression()
		}

		p.SetState(289)
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

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Constant() IConstantContext
	Expression() IExpressionContext

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(BigCParserIdentifier, 0)
}

func (s *PrimaryExpressionContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *PrimaryExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitPrimaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, BigCParserRULE_primaryExpression)
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BigCParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(290)
			p.Match(BigCParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BigCParserBooleanConstant, BigCParserCharConstant, BigCParserIntegerConstant, BigCParserFloatingConstant:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(291)
			p.Constant()
		}

	case BigCParserT__7:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(292)
			p.Match(BigCParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(293)
			p.Expression()
		}
		{
			p.SetState(294)
			p.Match(BigCParserT__8)
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IntegerConstant() antlr.TerminalNode
	FloatingConstant() antlr.TerminalNode
	BooleanConstant() antlr.TerminalNode
	CharConstant() antlr.TerminalNode

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BigCParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BigCParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) IntegerConstant() antlr.TerminalNode {
	return s.GetToken(BigCParserIntegerConstant, 0)
}

func (s *ConstantContext) FloatingConstant() antlr.TerminalNode {
	return s.GetToken(BigCParserFloatingConstant, 0)
}

func (s *ConstantContext) BooleanConstant() antlr.TerminalNode {
	return s.GetToken(BigCParserBooleanConstant, 0)
}

func (s *ConstantContext) CharConstant() antlr.TerminalNode {
	return s.GetToken(BigCParserCharConstant, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BigCListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BigCVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BigCParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, BigCParserRULE_constant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&115964116992) != 0) {
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

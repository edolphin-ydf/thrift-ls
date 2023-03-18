// Code generated from Thrift.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Thrift

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ThriftParser struct {
	*antlr.BaseParser
}

var thriftParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func thriftParserInit() {
	staticData := &thriftParserStaticData
	staticData.literalNames = []string{
		"", "'include'", "'namespace'", "'*'", "'cpp_namespace'", "'php_namespace'",
		"'cpp_include'", "'const'", "'='", "'typedef'", "'enum'", "'{'", "'}'",
		"'senum'", "'struct'", "'union'", "'exception'", "'service'", "'extends'",
		"':'", "'required'", "'optional'", "'('", "')'", "'oneway'", "'async'",
		"'void'", "'throws'", "'map'", "'<'", "'>'", "'set'", "'list'", "'cpp_type'",
		"'['", "']'", "';'", "", "", "", "'bool'", "'byte'", "'i16'", "'i32'",
		"'i64'", "'double'", "'string'", "'binary'", "", "", "','",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "INTEGER", "HEX_INTEGER", "DOUBLE", "TYPE_BOOL", "TYPE_BYTE",
		"TYPE_I16", "TYPE_I32", "TYPE_I64", "TYPE_DOUBLE", "TYPE_STRING", "TYPE_BINARY",
		"LITERAL", "IDENTIFIER", "COMMA", "WS", "SL_COMMENT", "ML_COMMENT",
	}
	staticData.ruleNames = []string{
		"document", "header", "include_", "namespace_", "cpp_include", "definition",
		"const_rule", "typedef_", "enum_rule", "enum_field", "senum", "struct_",
		"union_", "exception", "service", "field", "field_id", "field_req",
		"function_", "oneway", "function_type", "throws_list", "type_annotations",
		"type_annotation", "annotation_value", "field_type", "base_type", "container_type",
		"map_type", "set_type", "list_type", "cpp_type", "const_value", "integer",
		"const_list", "const_map_entry", "const_map", "list_separator", "real_base_type",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 53, 409, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 1, 0, 5, 0, 80, 8, 0, 10, 0, 12, 0, 83, 9,
		0, 1, 0, 5, 0, 86, 8, 0, 10, 0, 12, 0, 89, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 3, 1, 96, 8, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 3, 3, 108, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 114, 8, 3, 1, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 127,
		8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 134, 8, 6, 1, 6, 3, 6, 137, 8,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 143, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 5,
		8, 149, 8, 8, 10, 8, 12, 8, 152, 9, 8, 1, 8, 1, 8, 3, 8, 156, 8, 8, 1,
		9, 1, 9, 1, 9, 3, 9, 161, 8, 9, 1, 9, 3, 9, 164, 8, 9, 1, 9, 3, 9, 167,
		8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 174, 8, 10, 5, 10, 176,
		8, 10, 10, 10, 12, 10, 179, 9, 10, 1, 10, 1, 10, 3, 10, 183, 8, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 5, 11, 189, 8, 11, 10, 11, 12, 11, 192, 9, 11,
		1, 11, 1, 11, 3, 11, 196, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 202,
		8, 12, 10, 12, 12, 12, 205, 9, 12, 1, 12, 1, 12, 3, 12, 209, 8, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 5, 13, 215, 8, 13, 10, 13, 12, 13, 218, 9, 13,
		1, 13, 1, 13, 3, 13, 222, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 228,
		8, 14, 1, 14, 1, 14, 5, 14, 232, 8, 14, 10, 14, 12, 14, 235, 9, 14, 1,
		14, 1, 14, 3, 14, 239, 8, 14, 1, 15, 3, 15, 242, 8, 15, 1, 15, 3, 15, 245,
		8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 251, 8, 15, 1, 15, 3, 15, 254,
		8, 15, 1, 15, 3, 15, 257, 8, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		18, 3, 18, 265, 8, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 271, 8, 18, 10,
		18, 12, 18, 274, 9, 18, 1, 18, 1, 18, 3, 18, 278, 8, 18, 1, 18, 3, 18,
		281, 8, 18, 1, 18, 3, 18, 284, 8, 18, 1, 19, 1, 19, 1, 20, 1, 20, 3, 20,
		290, 8, 20, 1, 21, 1, 21, 1, 21, 5, 21, 295, 8, 21, 10, 21, 12, 21, 298,
		9, 21, 1, 21, 1, 21, 1, 22, 1, 22, 5, 22, 304, 8, 22, 10, 22, 12, 22, 307,
		9, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 3, 23, 314, 8, 23, 1, 23, 3,
		23, 317, 8, 23, 1, 24, 1, 24, 3, 24, 321, 8, 24, 1, 25, 1, 25, 1, 25, 3,
		25, 326, 8, 25, 1, 26, 1, 26, 3, 26, 330, 8, 26, 1, 27, 1, 27, 1, 27, 3,
		27, 335, 8, 27, 1, 27, 3, 27, 338, 8, 27, 1, 28, 1, 28, 3, 28, 342, 8,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 3, 29, 352,
		8, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3,
		30, 363, 8, 30, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 3, 32, 374, 8, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 3, 34, 381,
		8, 34, 5, 34, 383, 8, 34, 10, 34, 12, 34, 386, 9, 34, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 35, 1, 35, 3, 35, 394, 8, 35, 1, 36, 1, 36, 5, 36, 398, 8,
		36, 10, 36, 12, 36, 401, 9, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38,
		1, 38, 0, 0, 39, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
		66, 68, 70, 72, 74, 76, 0, 6, 1, 0, 48, 49, 1, 0, 20, 21, 1, 0, 24, 25,
		1, 0, 37, 38, 2, 0, 36, 36, 50, 50, 1, 0, 40, 47, 438, 0, 81, 1, 0, 0,
		0, 2, 95, 1, 0, 0, 0, 4, 97, 1, 0, 0, 0, 6, 113, 1, 0, 0, 0, 8, 115, 1,
		0, 0, 0, 10, 126, 1, 0, 0, 0, 12, 128, 1, 0, 0, 0, 14, 138, 1, 0, 0, 0,
		16, 144, 1, 0, 0, 0, 18, 157, 1, 0, 0, 0, 20, 168, 1, 0, 0, 0, 22, 184,
		1, 0, 0, 0, 24, 197, 1, 0, 0, 0, 26, 210, 1, 0, 0, 0, 28, 223, 1, 0, 0,
		0, 30, 241, 1, 0, 0, 0, 32, 258, 1, 0, 0, 0, 34, 261, 1, 0, 0, 0, 36, 264,
		1, 0, 0, 0, 38, 285, 1, 0, 0, 0, 40, 289, 1, 0, 0, 0, 42, 291, 1, 0, 0,
		0, 44, 301, 1, 0, 0, 0, 46, 310, 1, 0, 0, 0, 48, 320, 1, 0, 0, 0, 50, 325,
		1, 0, 0, 0, 52, 327, 1, 0, 0, 0, 54, 334, 1, 0, 0, 0, 56, 339, 1, 0, 0,
		0, 58, 349, 1, 0, 0, 0, 60, 357, 1, 0, 0, 0, 62, 364, 1, 0, 0, 0, 64, 373,
		1, 0, 0, 0, 66, 375, 1, 0, 0, 0, 68, 377, 1, 0, 0, 0, 70, 389, 1, 0, 0,
		0, 72, 395, 1, 0, 0, 0, 74, 404, 1, 0, 0, 0, 76, 406, 1, 0, 0, 0, 78, 80,
		3, 2, 1, 0, 79, 78, 1, 0, 0, 0, 80, 83, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0,
		81, 82, 1, 0, 0, 0, 82, 87, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 84, 86, 3,
		10, 5, 0, 85, 84, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87,
		88, 1, 0, 0, 0, 88, 90, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 91, 5, 0, 0,
		1, 91, 1, 1, 0, 0, 0, 92, 96, 3, 4, 2, 0, 93, 96, 3, 6, 3, 0, 94, 96, 3,
		8, 4, 0, 95, 92, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 94, 1, 0, 0, 0, 96,
		3, 1, 0, 0, 0, 97, 98, 5, 1, 0, 0, 98, 99, 5, 48, 0, 0, 99, 5, 1, 0, 0,
		0, 100, 101, 5, 2, 0, 0, 101, 102, 5, 3, 0, 0, 102, 114, 7, 0, 0, 0, 103,
		104, 5, 2, 0, 0, 104, 105, 5, 49, 0, 0, 105, 107, 7, 0, 0, 0, 106, 108,
		3, 44, 22, 0, 107, 106, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 114, 1,
		0, 0, 0, 109, 110, 5, 4, 0, 0, 110, 114, 5, 49, 0, 0, 111, 112, 5, 5, 0,
		0, 112, 114, 5, 49, 0, 0, 113, 100, 1, 0, 0, 0, 113, 103, 1, 0, 0, 0, 113,
		109, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 114, 7, 1, 0, 0, 0, 115, 116, 5,
		6, 0, 0, 116, 117, 5, 48, 0, 0, 117, 9, 1, 0, 0, 0, 118, 127, 3, 12, 6,
		0, 119, 127, 3, 14, 7, 0, 120, 127, 3, 16, 8, 0, 121, 127, 3, 20, 10, 0,
		122, 127, 3, 22, 11, 0, 123, 127, 3, 24, 12, 0, 124, 127, 3, 26, 13, 0,
		125, 127, 3, 28, 14, 0, 126, 118, 1, 0, 0, 0, 126, 119, 1, 0, 0, 0, 126,
		120, 1, 0, 0, 0, 126, 121, 1, 0, 0, 0, 126, 122, 1, 0, 0, 0, 126, 123,
		1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 125, 1, 0, 0, 0, 127, 11, 1, 0,
		0, 0, 128, 129, 5, 7, 0, 0, 129, 130, 3, 50, 25, 0, 130, 133, 5, 49, 0,
		0, 131, 132, 5, 8, 0, 0, 132, 134, 3, 64, 32, 0, 133, 131, 1, 0, 0, 0,
		133, 134, 1, 0, 0, 0, 134, 136, 1, 0, 0, 0, 135, 137, 3, 74, 37, 0, 136,
		135, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 13, 1, 0, 0, 0, 138, 139, 5,
		9, 0, 0, 139, 140, 3, 50, 25, 0, 140, 142, 5, 49, 0, 0, 141, 143, 3, 44,
		22, 0, 142, 141, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 15, 1, 0, 0, 0,
		144, 145, 5, 10, 0, 0, 145, 146, 5, 49, 0, 0, 146, 150, 5, 11, 0, 0, 147,
		149, 3, 18, 9, 0, 148, 147, 1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148,
		1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 153, 1, 0, 0, 0, 152, 150, 1, 0,
		0, 0, 153, 155, 5, 12, 0, 0, 154, 156, 3, 44, 22, 0, 155, 154, 1, 0, 0,
		0, 155, 156, 1, 0, 0, 0, 156, 17, 1, 0, 0, 0, 157, 160, 5, 49, 0, 0, 158,
		159, 5, 8, 0, 0, 159, 161, 3, 66, 33, 0, 160, 158, 1, 0, 0, 0, 160, 161,
		1, 0, 0, 0, 161, 163, 1, 0, 0, 0, 162, 164, 3, 44, 22, 0, 163, 162, 1,
		0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 166, 1, 0, 0, 0, 165, 167, 3, 74, 37,
		0, 166, 165, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 19, 1, 0, 0, 0, 168,
		169, 5, 13, 0, 0, 169, 170, 5, 49, 0, 0, 170, 177, 5, 11, 0, 0, 171, 173,
		5, 48, 0, 0, 172, 174, 3, 74, 37, 0, 173, 172, 1, 0, 0, 0, 173, 174, 1,
		0, 0, 0, 174, 176, 1, 0, 0, 0, 175, 171, 1, 0, 0, 0, 176, 179, 1, 0, 0,
		0, 177, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 180, 1, 0, 0, 0, 179,
		177, 1, 0, 0, 0, 180, 182, 5, 12, 0, 0, 181, 183, 3, 44, 22, 0, 182, 181,
		1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 21, 1, 0, 0, 0, 184, 185, 5, 14,
		0, 0, 185, 186, 5, 49, 0, 0, 186, 190, 5, 11, 0, 0, 187, 189, 3, 30, 15,
		0, 188, 187, 1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190,
		191, 1, 0, 0, 0, 191, 193, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 193, 195,
		5, 12, 0, 0, 194, 196, 3, 44, 22, 0, 195, 194, 1, 0, 0, 0, 195, 196, 1,
		0, 0, 0, 196, 23, 1, 0, 0, 0, 197, 198, 5, 15, 0, 0, 198, 199, 5, 49, 0,
		0, 199, 203, 5, 11, 0, 0, 200, 202, 3, 30, 15, 0, 201, 200, 1, 0, 0, 0,
		202, 205, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204,
		206, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 206, 208, 5, 12, 0, 0, 207, 209,
		3, 44, 22, 0, 208, 207, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 25, 1, 0,
		0, 0, 210, 211, 5, 16, 0, 0, 211, 212, 5, 49, 0, 0, 212, 216, 5, 11, 0,
		0, 213, 215, 3, 30, 15, 0, 214, 213, 1, 0, 0, 0, 215, 218, 1, 0, 0, 0,
		216, 214, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 219, 1, 0, 0, 0, 218,
		216, 1, 0, 0, 0, 219, 221, 5, 12, 0, 0, 220, 222, 3, 44, 22, 0, 221, 220,
		1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 27, 1, 0, 0, 0, 223, 224, 5, 17,
		0, 0, 224, 227, 5, 49, 0, 0, 225, 226, 5, 18, 0, 0, 226, 228, 5, 49, 0,
		0, 227, 225, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229,
		233, 5, 11, 0, 0, 230, 232, 3, 36, 18, 0, 231, 230, 1, 0, 0, 0, 232, 235,
		1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 236, 1, 0,
		0, 0, 235, 233, 1, 0, 0, 0, 236, 238, 5, 12, 0, 0, 237, 239, 3, 44, 22,
		0, 238, 237, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 29, 1, 0, 0, 0, 240,
		242, 3, 32, 16, 0, 241, 240, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 244,
		1, 0, 0, 0, 243, 245, 3, 34, 17, 0, 244, 243, 1, 0, 0, 0, 244, 245, 1,
		0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 247, 3, 50, 25, 0, 247, 250, 5, 49,
		0, 0, 248, 249, 5, 8, 0, 0, 249, 251, 3, 64, 32, 0, 250, 248, 1, 0, 0,
		0, 250, 251, 1, 0, 0, 0, 251, 253, 1, 0, 0, 0, 252, 254, 3, 44, 22, 0,
		253, 252, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 256, 1, 0, 0, 0, 255,
		257, 3, 74, 37, 0, 256, 255, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 31,
		1, 0, 0, 0, 258, 259, 3, 66, 33, 0, 259, 260, 5, 19, 0, 0, 260, 33, 1,
		0, 0, 0, 261, 262, 7, 1, 0, 0, 262, 35, 1, 0, 0, 0, 263, 265, 3, 38, 19,
		0, 264, 263, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266,
		267, 3, 40, 20, 0, 267, 268, 5, 49, 0, 0, 268, 272, 5, 22, 0, 0, 269, 271,
		3, 30, 15, 0, 270, 269, 1, 0, 0, 0, 271, 274, 1, 0, 0, 0, 272, 270, 1,
		0, 0, 0, 272, 273, 1, 0, 0, 0, 273, 275, 1, 0, 0, 0, 274, 272, 1, 0, 0,
		0, 275, 277, 5, 23, 0, 0, 276, 278, 3, 42, 21, 0, 277, 276, 1, 0, 0, 0,
		277, 278, 1, 0, 0, 0, 278, 280, 1, 0, 0, 0, 279, 281, 3, 44, 22, 0, 280,
		279, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 283, 1, 0, 0, 0, 282, 284,
		3, 74, 37, 0, 283, 282, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284, 37, 1, 0,
		0, 0, 285, 286, 7, 2, 0, 0, 286, 39, 1, 0, 0, 0, 287, 290, 3, 50, 25, 0,
		288, 290, 5, 26, 0, 0, 289, 287, 1, 0, 0, 0, 289, 288, 1, 0, 0, 0, 290,
		41, 1, 0, 0, 0, 291, 292, 5, 27, 0, 0, 292, 296, 5, 22, 0, 0, 293, 295,
		3, 30, 15, 0, 294, 293, 1, 0, 0, 0, 295, 298, 1, 0, 0, 0, 296, 294, 1,
		0, 0, 0, 296, 297, 1, 0, 0, 0, 297, 299, 1, 0, 0, 0, 298, 296, 1, 0, 0,
		0, 299, 300, 5, 23, 0, 0, 300, 43, 1, 0, 0, 0, 301, 305, 5, 22, 0, 0, 302,
		304, 3, 46, 23, 0, 303, 302, 1, 0, 0, 0, 304, 307, 1, 0, 0, 0, 305, 303,
		1, 0, 0, 0, 305, 306, 1, 0, 0, 0, 306, 308, 1, 0, 0, 0, 307, 305, 1, 0,
		0, 0, 308, 309, 5, 23, 0, 0, 309, 45, 1, 0, 0, 0, 310, 313, 5, 49, 0, 0,
		311, 312, 5, 8, 0, 0, 312, 314, 3, 48, 24, 0, 313, 311, 1, 0, 0, 0, 313,
		314, 1, 0, 0, 0, 314, 316, 1, 0, 0, 0, 315, 317, 3, 74, 37, 0, 316, 315,
		1, 0, 0, 0, 316, 317, 1, 0, 0, 0, 317, 47, 1, 0, 0, 0, 318, 321, 3, 66,
		33, 0, 319, 321, 5, 48, 0, 0, 320, 318, 1, 0, 0, 0, 320, 319, 1, 0, 0,
		0, 321, 49, 1, 0, 0, 0, 322, 326, 3, 52, 26, 0, 323, 326, 5, 49, 0, 0,
		324, 326, 3, 54, 27, 0, 325, 322, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 325,
		324, 1, 0, 0, 0, 326, 51, 1, 0, 0, 0, 327, 329, 3, 76, 38, 0, 328, 330,
		3, 44, 22, 0, 329, 328, 1, 0, 0, 0, 329, 330, 1, 0, 0, 0, 330, 53, 1, 0,
		0, 0, 331, 335, 3, 56, 28, 0, 332, 335, 3, 58, 29, 0, 333, 335, 3, 60,
		30, 0, 334, 331, 1, 0, 0, 0, 334, 332, 1, 0, 0, 0, 334, 333, 1, 0, 0, 0,
		335, 337, 1, 0, 0, 0, 336, 338, 3, 44, 22, 0, 337, 336, 1, 0, 0, 0, 337,
		338, 1, 0, 0, 0, 338, 55, 1, 0, 0, 0, 339, 341, 5, 28, 0, 0, 340, 342,
		3, 62, 31, 0, 341, 340, 1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342, 343, 1,
		0, 0, 0, 343, 344, 5, 29, 0, 0, 344, 345, 3, 50, 25, 0, 345, 346, 5, 50,
		0, 0, 346, 347, 3, 50, 25, 0, 347, 348, 5, 30, 0, 0, 348, 57, 1, 0, 0,
		0, 349, 351, 5, 31, 0, 0, 350, 352, 3, 62, 31, 0, 351, 350, 1, 0, 0, 0,
		351, 352, 1, 0, 0, 0, 352, 353, 1, 0, 0, 0, 353, 354, 5, 29, 0, 0, 354,
		355, 3, 50, 25, 0, 355, 356, 5, 30, 0, 0, 356, 59, 1, 0, 0, 0, 357, 358,
		5, 32, 0, 0, 358, 359, 5, 29, 0, 0, 359, 360, 3, 50, 25, 0, 360, 362, 5,
		30, 0, 0, 361, 363, 3, 62, 31, 0, 362, 361, 1, 0, 0, 0, 362, 363, 1, 0,
		0, 0, 363, 61, 1, 0, 0, 0, 364, 365, 5, 33, 0, 0, 365, 366, 5, 48, 0, 0,
		366, 63, 1, 0, 0, 0, 367, 374, 3, 66, 33, 0, 368, 374, 5, 39, 0, 0, 369,
		374, 5, 48, 0, 0, 370, 374, 5, 49, 0, 0, 371, 374, 3, 68, 34, 0, 372, 374,
		3, 72, 36, 0, 373, 367, 1, 0, 0, 0, 373, 368, 1, 0, 0, 0, 373, 369, 1,
		0, 0, 0, 373, 370, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 373, 372, 1, 0, 0,
		0, 374, 65, 1, 0, 0, 0, 375, 376, 7, 3, 0, 0, 376, 67, 1, 0, 0, 0, 377,
		384, 5, 34, 0, 0, 378, 380, 3, 64, 32, 0, 379, 381, 3, 74, 37, 0, 380,
		379, 1, 0, 0, 0, 380, 381, 1, 0, 0, 0, 381, 383, 1, 0, 0, 0, 382, 378,
		1, 0, 0, 0, 383, 386, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 384, 385, 1, 0,
		0, 0, 385, 387, 1, 0, 0, 0, 386, 384, 1, 0, 0, 0, 387, 388, 5, 35, 0, 0,
		388, 69, 1, 0, 0, 0, 389, 390, 3, 64, 32, 0, 390, 391, 5, 19, 0, 0, 391,
		393, 3, 64, 32, 0, 392, 394, 3, 74, 37, 0, 393, 392, 1, 0, 0, 0, 393, 394,
		1, 0, 0, 0, 394, 71, 1, 0, 0, 0, 395, 399, 5, 11, 0, 0, 396, 398, 3, 70,
		35, 0, 397, 396, 1, 0, 0, 0, 398, 401, 1, 0, 0, 0, 399, 397, 1, 0, 0, 0,
		399, 400, 1, 0, 0, 0, 400, 402, 1, 0, 0, 0, 401, 399, 1, 0, 0, 0, 402,
		403, 5, 12, 0, 0, 403, 73, 1, 0, 0, 0, 404, 405, 7, 4, 0, 0, 405, 75, 1,
		0, 0, 0, 406, 407, 7, 5, 0, 0, 407, 77, 1, 0, 0, 0, 54, 81, 87, 95, 107,
		113, 126, 133, 136, 142, 150, 155, 160, 163, 166, 173, 177, 182, 190, 195,
		203, 208, 216, 221, 227, 233, 238, 241, 244, 250, 253, 256, 264, 272, 277,
		280, 283, 289, 296, 305, 313, 316, 320, 325, 329, 334, 337, 341, 351, 362,
		373, 380, 384, 393, 399,
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

// ThriftParserInit initializes any static state used to implement ThriftParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewThriftParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ThriftParserInit() {
	staticData := &thriftParserStaticData
	staticData.once.Do(thriftParserInit)
}

// NewThriftParser produces a new parser instance for the optional input antlr.TokenStream.
func NewThriftParser(input antlr.TokenStream) *ThriftParser {
	ThriftParserInit()
	this := new(ThriftParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &thriftParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Thrift.g4"

	return this
}

// ThriftParser tokens.
const (
	ThriftParserEOF         = antlr.TokenEOF
	ThriftParserT__0        = 1
	ThriftParserT__1        = 2
	ThriftParserT__2        = 3
	ThriftParserT__3        = 4
	ThriftParserT__4        = 5
	ThriftParserT__5        = 6
	ThriftParserT__6        = 7
	ThriftParserT__7        = 8
	ThriftParserT__8        = 9
	ThriftParserT__9        = 10
	ThriftParserT__10       = 11
	ThriftParserT__11       = 12
	ThriftParserT__12       = 13
	ThriftParserT__13       = 14
	ThriftParserT__14       = 15
	ThriftParserT__15       = 16
	ThriftParserT__16       = 17
	ThriftParserT__17       = 18
	ThriftParserT__18       = 19
	ThriftParserT__19       = 20
	ThriftParserT__20       = 21
	ThriftParserT__21       = 22
	ThriftParserT__22       = 23
	ThriftParserT__23       = 24
	ThriftParserT__24       = 25
	ThriftParserT__25       = 26
	ThriftParserT__26       = 27
	ThriftParserT__27       = 28
	ThriftParserT__28       = 29
	ThriftParserT__29       = 30
	ThriftParserT__30       = 31
	ThriftParserT__31       = 32
	ThriftParserT__32       = 33
	ThriftParserT__33       = 34
	ThriftParserT__34       = 35
	ThriftParserT__35       = 36
	ThriftParserINTEGER     = 37
	ThriftParserHEX_INTEGER = 38
	ThriftParserDOUBLE      = 39
	ThriftParserTYPE_BOOL   = 40
	ThriftParserTYPE_BYTE   = 41
	ThriftParserTYPE_I16    = 42
	ThriftParserTYPE_I32    = 43
	ThriftParserTYPE_I64    = 44
	ThriftParserTYPE_DOUBLE = 45
	ThriftParserTYPE_STRING = 46
	ThriftParserTYPE_BINARY = 47
	ThriftParserLITERAL     = 48
	ThriftParserIDENTIFIER  = 49
	ThriftParserCOMMA       = 50
	ThriftParserWS          = 51
	ThriftParserSL_COMMENT  = 52
	ThriftParserML_COMMENT  = 53
)

// ThriftParser rules.
const (
	ThriftParserRULE_document         = 0
	ThriftParserRULE_header           = 1
	ThriftParserRULE_include_         = 2
	ThriftParserRULE_namespace_       = 3
	ThriftParserRULE_cpp_include      = 4
	ThriftParserRULE_definition       = 5
	ThriftParserRULE_const_rule       = 6
	ThriftParserRULE_typedef_         = 7
	ThriftParserRULE_enum_rule        = 8
	ThriftParserRULE_enum_field       = 9
	ThriftParserRULE_senum            = 10
	ThriftParserRULE_struct_          = 11
	ThriftParserRULE_union_           = 12
	ThriftParserRULE_exception        = 13
	ThriftParserRULE_service          = 14
	ThriftParserRULE_field            = 15
	ThriftParserRULE_field_id         = 16
	ThriftParserRULE_field_req        = 17
	ThriftParserRULE_function_        = 18
	ThriftParserRULE_oneway           = 19
	ThriftParserRULE_function_type    = 20
	ThriftParserRULE_throws_list      = 21
	ThriftParserRULE_type_annotations = 22
	ThriftParserRULE_type_annotation  = 23
	ThriftParserRULE_annotation_value = 24
	ThriftParserRULE_field_type       = 25
	ThriftParserRULE_base_type        = 26
	ThriftParserRULE_container_type   = 27
	ThriftParserRULE_map_type         = 28
	ThriftParserRULE_set_type         = 29
	ThriftParserRULE_list_type        = 30
	ThriftParserRULE_cpp_type         = 31
	ThriftParserRULE_const_value      = 32
	ThriftParserRULE_integer          = 33
	ThriftParserRULE_const_list       = 34
	ThriftParserRULE_const_map_entry  = 35
	ThriftParserRULE_const_map        = 36
	ThriftParserRULE_list_separator   = 37
	ThriftParserRULE_real_base_type   = 38
)

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllHeader() []IHeaderContext
	Header(i int) IHeaderContext
	AllDefinition() []IDefinitionContext
	Definition(i int) IDefinitionContext

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) EOF() antlr.TerminalNode {
	return s.GetToken(ThriftParserEOF, 0)
}

func (s *DocumentContext) AllHeader() []IHeaderContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHeaderContext); ok {
			len++
		}
	}

	tst := make([]IHeaderContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHeaderContext); ok {
			tst[i] = t.(IHeaderContext)
			i++
		}
	}

	return tst
}

func (s *DocumentContext) Header(i int) IHeaderContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHeaderContext); ok {
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

	return t.(IHeaderContext)
}

func (s *DocumentContext) AllDefinition() []IDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDefinitionContext); ok {
			tst[i] = t.(IDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *DocumentContext) Definition(i int) IDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionContext); ok {
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

	return t.(IDefinitionContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *ThriftParser) Document() (localctx IDocumentContext) {
	this := p
	_ = this

	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ThriftParserRULE_document)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&118) != 0 {
		{
			p.SetState(78)
			p.Header()
		}

		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&255616) != 0 {
		{
			p.SetState(84)
			p.Definition()
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
		p.Match(ThriftParserEOF)
	}

	return localctx
}

// IHeaderContext is an interface to support dynamic dispatch.
type IHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Include_() IInclude_Context
	Namespace_() INamespace_Context
	Cpp_include() ICpp_includeContext

	// IsHeaderContext differentiates from other interfaces.
	IsHeaderContext()
}

type HeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderContext() *HeaderContext {
	var p = new(HeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_header
	return p
}

func (*HeaderContext) IsHeaderContext() {}

func NewHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderContext {
	var p = new(HeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_header

	return p
}

func (s *HeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderContext) Include_() IInclude_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInclude_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInclude_Context)
}

func (s *HeaderContext) Namespace_() INamespace_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamespace_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamespace_Context)
}

func (s *HeaderContext) Cpp_include() ICpp_includeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICpp_includeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICpp_includeContext)
}

func (s *HeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterHeader(s)
	}
}

func (s *HeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitHeader(s)
	}
}

func (p *ThriftParser) Header() (localctx IHeaderContext) {
	this := p
	_ = this

	localctx = NewHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ThriftParserRULE_header)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(95)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThriftParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Include_()
		}

	case ThriftParserT__1, ThriftParserT__3, ThriftParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Namespace_()
		}

	case ThriftParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(94)
			p.Cpp_include()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInclude_Context is an interface to support dynamic dispatch.
type IInclude_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LITERAL() antlr.TerminalNode

	// IsInclude_Context differentiates from other interfaces.
	IsInclude_Context()
}

type Include_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInclude_Context() *Include_Context {
	var p = new(Include_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_include_
	return p
}

func (*Include_Context) IsInclude_Context() {}

func NewInclude_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Include_Context {
	var p = new(Include_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_include_

	return p
}

func (s *Include_Context) GetParser() antlr.Parser { return s.parser }

func (s *Include_Context) LITERAL() antlr.TerminalNode {
	return s.GetToken(ThriftParserLITERAL, 0)
}

func (s *Include_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Include_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Include_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterInclude_(s)
	}
}

func (s *Include_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitInclude_(s)
	}
}

func (p *ThriftParser) Include_() (localctx IInclude_Context) {
	this := p
	_ = this

	localctx = NewInclude_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ThriftParserRULE_include_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(ThriftParserT__0)
	}
	{
		p.SetState(98)
		p.Match(ThriftParserLITERAL)
	}

	return localctx
}

// INamespace_Context is an interface to support dynamic dispatch.
type INamespace_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	LITERAL() antlr.TerminalNode
	Type_annotations() IType_annotationsContext

	// IsNamespace_Context differentiates from other interfaces.
	IsNamespace_Context()
}

type Namespace_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespace_Context() *Namespace_Context {
	var p = new(Namespace_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_namespace_
	return p
}

func (*Namespace_Context) IsNamespace_Context() {}

func NewNamespace_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Namespace_Context {
	var p = new(Namespace_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_namespace_

	return p
}

func (s *Namespace_Context) GetParser() antlr.Parser { return s.parser }

func (s *Namespace_Context) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ThriftParserIDENTIFIER)
}

func (s *Namespace_Context) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, i)
}

func (s *Namespace_Context) LITERAL() antlr.TerminalNode {
	return s.GetToken(ThriftParserLITERAL, 0)
}

func (s *Namespace_Context) Type_annotations() IType_annotationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_annotationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *Namespace_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Namespace_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Namespace_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterNamespace_(s)
	}
}

func (s *Namespace_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitNamespace_(s)
	}
}

func (p *ThriftParser) Namespace_() (localctx INamespace_Context) {
	this := p
	_ = this

	localctx = NewNamespace_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ThriftParserRULE_namespace_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.Match(ThriftParserT__1)
		}
		{
			p.SetState(101)
			p.Match(ThriftParserT__2)
		}
		{
			p.SetState(102)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ThriftParserLITERAL || _la == ThriftParserIDENTIFIER) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Match(ThriftParserT__1)
		}
		{
			p.SetState(104)
			p.Match(ThriftParserIDENTIFIER)
		}
		{
			p.SetState(105)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ThriftParserLITERAL || _la == ThriftParserIDENTIFIER) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ThriftParserT__21 {
			{
				p.SetState(106)
				p.Type_annotations()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(109)
			p.Match(ThriftParserT__3)
		}
		{
			p.SetState(110)
			p.Match(ThriftParserIDENTIFIER)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(111)
			p.Match(ThriftParserT__4)
		}
		{
			p.SetState(112)
			p.Match(ThriftParserIDENTIFIER)
		}

	}

	return localctx
}

// ICpp_includeContext is an interface to support dynamic dispatch.
type ICpp_includeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LITERAL() antlr.TerminalNode

	// IsCpp_includeContext differentiates from other interfaces.
	IsCpp_includeContext()
}

type Cpp_includeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCpp_includeContext() *Cpp_includeContext {
	var p = new(Cpp_includeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_cpp_include
	return p
}

func (*Cpp_includeContext) IsCpp_includeContext() {}

func NewCpp_includeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cpp_includeContext {
	var p = new(Cpp_includeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_cpp_include

	return p
}

func (s *Cpp_includeContext) GetParser() antlr.Parser { return s.parser }

func (s *Cpp_includeContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(ThriftParserLITERAL, 0)
}

func (s *Cpp_includeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cpp_includeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cpp_includeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterCpp_include(s)
	}
}

func (s *Cpp_includeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitCpp_include(s)
	}
}

func (p *ThriftParser) Cpp_include() (localctx ICpp_includeContext) {
	this := p
	_ = this

	localctx = NewCpp_includeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ThriftParserRULE_cpp_include)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(ThriftParserT__5)
	}
	{
		p.SetState(116)
		p.Match(ThriftParserLITERAL)
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Const_rule() IConst_ruleContext
	Typedef_() ITypedef_Context
	Enum_rule() IEnum_ruleContext
	Senum() ISenumContext
	Struct_() IStruct_Context
	Union_() IUnion_Context
	Exception() IExceptionContext
	Service() IServiceContext

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) Const_rule() IConst_ruleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_ruleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_ruleContext)
}

func (s *DefinitionContext) Typedef_() ITypedef_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypedef_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypedef_Context)
}

func (s *DefinitionContext) Enum_rule() IEnum_ruleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnum_ruleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnum_ruleContext)
}

func (s *DefinitionContext) Senum() ISenumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISenumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISenumContext)
}

func (s *DefinitionContext) Struct_() IStruct_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_Context)
}

func (s *DefinitionContext) Union_() IUnion_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnion_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnion_Context)
}

func (s *DefinitionContext) Exception() IExceptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExceptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExceptionContext)
}

func (s *DefinitionContext) Service() IServiceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *ThriftParser) Definition() (localctx IDefinitionContext) {
	this := p
	_ = this

	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ThriftParserRULE_definition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(126)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThriftParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(118)
			p.Const_rule()
		}

	case ThriftParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(119)
			p.Typedef_()
		}

	case ThriftParserT__9:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(120)
			p.Enum_rule()
		}

	case ThriftParserT__12:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(121)
			p.Senum()
		}

	case ThriftParserT__13:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(122)
			p.Struct_()
		}

	case ThriftParserT__14:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(123)
			p.Union_()
		}

	case ThriftParserT__15:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(124)
			p.Exception()
		}

	case ThriftParserT__16:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(125)
			p.Service()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConst_ruleContext is an interface to support dynamic dispatch.
type IConst_ruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field_type() IField_typeContext
	IDENTIFIER() antlr.TerminalNode
	Const_value() IConst_valueContext
	List_separator() IList_separatorContext

	// IsConst_ruleContext differentiates from other interfaces.
	IsConst_ruleContext()
}

type Const_ruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_ruleContext() *Const_ruleContext {
	var p = new(Const_ruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_const_rule
	return p
}

func (*Const_ruleContext) IsConst_ruleContext() {}

func NewConst_ruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_ruleContext {
	var p = new(Const_ruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_const_rule

	return p
}

func (s *Const_ruleContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_ruleContext) Field_type() IField_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_typeContext)
}

func (s *Const_ruleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *Const_ruleContext) Const_value() IConst_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_valueContext)
}

func (s *Const_ruleContext) List_separator() IList_separatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_separatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_separatorContext)
}

func (s *Const_ruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_ruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_ruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterConst_rule(s)
	}
}

func (s *Const_ruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitConst_rule(s)
	}
}

func (p *ThriftParser) Const_rule() (localctx IConst_ruleContext) {
	this := p
	_ = this

	localctx = NewConst_ruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ThriftParserRULE_const_rule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(ThriftParserT__6)
	}
	{
		p.SetState(129)
		p.Field_type()
	}
	{
		p.SetState(130)
		p.Match(ThriftParserIDENTIFIER)
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__7 {
		{
			p.SetState(131)
			p.Match(ThriftParserT__7)
		}
		{
			p.SetState(132)
			p.Const_value()
		}

	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__35 || _la == ThriftParserCOMMA {
		{
			p.SetState(135)
			p.List_separator()
		}

	}

	return localctx
}

// ITypedef_Context is an interface to support dynamic dispatch.
type ITypedef_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field_type() IField_typeContext
	IDENTIFIER() antlr.TerminalNode
	Type_annotations() IType_annotationsContext

	// IsTypedef_Context differentiates from other interfaces.
	IsTypedef_Context()
}

type Typedef_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypedef_Context() *Typedef_Context {
	var p = new(Typedef_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_typedef_
	return p
}

func (*Typedef_Context) IsTypedef_Context() {}

func NewTypedef_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Typedef_Context {
	var p = new(Typedef_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_typedef_

	return p
}

func (s *Typedef_Context) GetParser() antlr.Parser { return s.parser }

func (s *Typedef_Context) Field_type() IField_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_typeContext)
}

func (s *Typedef_Context) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *Typedef_Context) Type_annotations() IType_annotationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_annotationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *Typedef_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Typedef_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Typedef_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterTypedef_(s)
	}
}

func (s *Typedef_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitTypedef_(s)
	}
}

func (p *ThriftParser) Typedef_() (localctx ITypedef_Context) {
	this := p
	_ = this

	localctx = NewTypedef_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ThriftParserRULE_typedef_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(ThriftParserT__8)
	}
	{
		p.SetState(139)
		p.Field_type()
	}
	{
		p.SetState(140)
		p.Match(ThriftParserIDENTIFIER)
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(141)
			p.Type_annotations()
		}

	}

	return localctx
}

// IEnum_ruleContext is an interface to support dynamic dispatch.
type IEnum_ruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllEnum_field() []IEnum_fieldContext
	Enum_field(i int) IEnum_fieldContext
	Type_annotations() IType_annotationsContext

	// IsEnum_ruleContext differentiates from other interfaces.
	IsEnum_ruleContext()
}

type Enum_ruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnum_ruleContext() *Enum_ruleContext {
	var p = new(Enum_ruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_enum_rule
	return p
}

func (*Enum_ruleContext) IsEnum_ruleContext() {}

func NewEnum_ruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Enum_ruleContext {
	var p = new(Enum_ruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_enum_rule

	return p
}

func (s *Enum_ruleContext) GetParser() antlr.Parser { return s.parser }

func (s *Enum_ruleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *Enum_ruleContext) AllEnum_field() []IEnum_fieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnum_fieldContext); ok {
			len++
		}
	}

	tst := make([]IEnum_fieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnum_fieldContext); ok {
			tst[i] = t.(IEnum_fieldContext)
			i++
		}
	}

	return tst
}

func (s *Enum_ruleContext) Enum_field(i int) IEnum_fieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnum_fieldContext); ok {
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

	return t.(IEnum_fieldContext)
}

func (s *Enum_ruleContext) Type_annotations() IType_annotationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_annotationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *Enum_ruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Enum_ruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Enum_ruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterEnum_rule(s)
	}
}

func (s *Enum_ruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitEnum_rule(s)
	}
}

func (p *ThriftParser) Enum_rule() (localctx IEnum_ruleContext) {
	this := p
	_ = this

	localctx = NewEnum_ruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ThriftParserRULE_enum_rule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(ThriftParserT__9)
	}
	{
		p.SetState(145)
		p.Match(ThriftParserIDENTIFIER)
	}
	{
		p.SetState(146)
		p.Match(ThriftParserT__10)
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ThriftParserIDENTIFIER {
		{
			p.SetState(147)
			p.Enum_field()
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(153)
		p.Match(ThriftParserT__11)
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(154)
			p.Type_annotations()
		}

	}

	return localctx
}

// IEnum_fieldContext is an interface to support dynamic dispatch.
type IEnum_fieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Integer() IIntegerContext
	Type_annotations() IType_annotationsContext
	List_separator() IList_separatorContext

	// IsEnum_fieldContext differentiates from other interfaces.
	IsEnum_fieldContext()
}

type Enum_fieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnum_fieldContext() *Enum_fieldContext {
	var p = new(Enum_fieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_enum_field
	return p
}

func (*Enum_fieldContext) IsEnum_fieldContext() {}

func NewEnum_fieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Enum_fieldContext {
	var p = new(Enum_fieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_enum_field

	return p
}

func (s *Enum_fieldContext) GetParser() antlr.Parser { return s.parser }

func (s *Enum_fieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *Enum_fieldContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *Enum_fieldContext) Type_annotations() IType_annotationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_annotationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *Enum_fieldContext) List_separator() IList_separatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_separatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_separatorContext)
}

func (s *Enum_fieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Enum_fieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Enum_fieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterEnum_field(s)
	}
}

func (s *Enum_fieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitEnum_field(s)
	}
}

func (p *ThriftParser) Enum_field() (localctx IEnum_fieldContext) {
	this := p
	_ = this

	localctx = NewEnum_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ThriftParserRULE_enum_field)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(ThriftParserIDENTIFIER)
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__7 {
		{
			p.SetState(158)
			p.Match(ThriftParserT__7)
		}
		{
			p.SetState(159)
			p.Integer()
		}

	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(162)
			p.Type_annotations()
		}

	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__35 || _la == ThriftParserCOMMA {
		{
			p.SetState(165)
			p.List_separator()
		}

	}

	return localctx
}

// ISenumContext is an interface to support dynamic dispatch.
type ISenumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllLITERAL() []antlr.TerminalNode
	LITERAL(i int) antlr.TerminalNode
	Type_annotations() IType_annotationsContext
	AllList_separator() []IList_separatorContext
	List_separator(i int) IList_separatorContext

	// IsSenumContext differentiates from other interfaces.
	IsSenumContext()
}

type SenumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySenumContext() *SenumContext {
	var p = new(SenumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_senum
	return p
}

func (*SenumContext) IsSenumContext() {}

func NewSenumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SenumContext {
	var p = new(SenumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_senum

	return p
}

func (s *SenumContext) GetParser() antlr.Parser { return s.parser }

func (s *SenumContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *SenumContext) AllLITERAL() []antlr.TerminalNode {
	return s.GetTokens(ThriftParserLITERAL)
}

func (s *SenumContext) LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(ThriftParserLITERAL, i)
}

func (s *SenumContext) Type_annotations() IType_annotationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_annotationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *SenumContext) AllList_separator() []IList_separatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IList_separatorContext); ok {
			len++
		}
	}

	tst := make([]IList_separatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IList_separatorContext); ok {
			tst[i] = t.(IList_separatorContext)
			i++
		}
	}

	return tst
}

func (s *SenumContext) List_separator(i int) IList_separatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_separatorContext); ok {
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

	return t.(IList_separatorContext)
}

func (s *SenumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SenumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SenumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterSenum(s)
	}
}

func (s *SenumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitSenum(s)
	}
}

func (p *ThriftParser) Senum() (localctx ISenumContext) {
	this := p
	_ = this

	localctx = NewSenumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ThriftParserRULE_senum)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(ThriftParserT__12)
	}
	{
		p.SetState(169)
		p.Match(ThriftParserIDENTIFIER)
	}
	{
		p.SetState(170)
		p.Match(ThriftParserT__10)
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ThriftParserLITERAL {
		{
			p.SetState(171)
			p.Match(ThriftParserLITERAL)
		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ThriftParserT__35 || _la == ThriftParserCOMMA {
			{
				p.SetState(172)
				p.List_separator()
			}

		}

		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(180)
		p.Match(ThriftParserT__11)
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(181)
			p.Type_annotations()
		}

	}

	return localctx
}

// IStruct_Context is an interface to support dynamic dispatch.
type IStruct_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllField() []IFieldContext
	Field(i int) IFieldContext
	Type_annotations() IType_annotationsContext

	// IsStruct_Context differentiates from other interfaces.
	IsStruct_Context()
}

type Struct_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_Context() *Struct_Context {
	var p = new(Struct_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_struct_
	return p
}

func (*Struct_Context) IsStruct_Context() {}

func NewStruct_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_Context {
	var p = new(Struct_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_struct_

	return p
}

func (s *Struct_Context) GetParser() antlr.Parser { return s.parser }

func (s *Struct_Context) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *Struct_Context) AllField() []IFieldContext {
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

func (s *Struct_Context) Field(i int) IFieldContext {
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

func (s *Struct_Context) Type_annotations() IType_annotationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_annotationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *Struct_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Struct_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterStruct_(s)
	}
}

func (s *Struct_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitStruct_(s)
	}
}

func (p *ThriftParser) Struct_() (localctx IStruct_Context) {
	this := p
	_ = this

	localctx = NewStruct_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ThriftParserRULE_struct_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(ThriftParserT__13)
	}
	{
		p.SetState(185)
		p.Match(ThriftParserIDENTIFIER)
	}
	{
		p.SetState(186)
		p.Match(ThriftParserT__10)
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&843744449396736) != 0 {
		{
			p.SetState(187)
			p.Field()
		}

		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(193)
		p.Match(ThriftParserT__11)
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(194)
			p.Type_annotations()
		}

	}

	return localctx
}

// IUnion_Context is an interface to support dynamic dispatch.
type IUnion_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllField() []IFieldContext
	Field(i int) IFieldContext
	Type_annotations() IType_annotationsContext

	// IsUnion_Context differentiates from other interfaces.
	IsUnion_Context()
}

type Union_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnion_Context() *Union_Context {
	var p = new(Union_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_union_
	return p
}

func (*Union_Context) IsUnion_Context() {}

func NewUnion_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Union_Context {
	var p = new(Union_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_union_

	return p
}

func (s *Union_Context) GetParser() antlr.Parser { return s.parser }

func (s *Union_Context) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *Union_Context) AllField() []IFieldContext {
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

func (s *Union_Context) Field(i int) IFieldContext {
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

func (s *Union_Context) Type_annotations() IType_annotationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_annotationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *Union_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Union_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Union_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterUnion_(s)
	}
}

func (s *Union_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitUnion_(s)
	}
}

func (p *ThriftParser) Union_() (localctx IUnion_Context) {
	this := p
	_ = this

	localctx = NewUnion_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ThriftParserRULE_union_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(ThriftParserT__14)
	}
	{
		p.SetState(198)
		p.Match(ThriftParserIDENTIFIER)
	}
	{
		p.SetState(199)
		p.Match(ThriftParserT__10)
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&843744449396736) != 0 {
		{
			p.SetState(200)
			p.Field()
		}

		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(206)
		p.Match(ThriftParserT__11)
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(207)
			p.Type_annotations()
		}

	}

	return localctx
}

// IExceptionContext is an interface to support dynamic dispatch.
type IExceptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllField() []IFieldContext
	Field(i int) IFieldContext
	Type_annotations() IType_annotationsContext

	// IsExceptionContext differentiates from other interfaces.
	IsExceptionContext()
}

type ExceptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExceptionContext() *ExceptionContext {
	var p = new(ExceptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_exception
	return p
}

func (*ExceptionContext) IsExceptionContext() {}

func NewExceptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExceptionContext {
	var p = new(ExceptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_exception

	return p
}

func (s *ExceptionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExceptionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *ExceptionContext) AllField() []IFieldContext {
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

func (s *ExceptionContext) Field(i int) IFieldContext {
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

func (s *ExceptionContext) Type_annotations() IType_annotationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_annotationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *ExceptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExceptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExceptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterException(s)
	}
}

func (s *ExceptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitException(s)
	}
}

func (p *ThriftParser) Exception() (localctx IExceptionContext) {
	this := p
	_ = this

	localctx = NewExceptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ThriftParserRULE_exception)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Match(ThriftParserT__15)
	}
	{
		p.SetState(211)
		p.Match(ThriftParserIDENTIFIER)
	}
	{
		p.SetState(212)
		p.Match(ThriftParserT__10)
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&843744449396736) != 0 {
		{
			p.SetState(213)
			p.Field()
		}

		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(219)
		p.Match(ThriftParserT__11)
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(220)
			p.Type_annotations()
		}

	}

	return localctx
}

// IServiceContext is an interface to support dynamic dispatch.
type IServiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllFunction_() []IFunction_Context
	Function_(i int) IFunction_Context
	Type_annotations() IType_annotationsContext

	// IsServiceContext differentiates from other interfaces.
	IsServiceContext()
}

type ServiceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceContext() *ServiceContext {
	var p = new(ServiceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_service
	return p
}

func (*ServiceContext) IsServiceContext() {}

func NewServiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceContext {
	var p = new(ServiceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_service

	return p
}

func (s *ServiceContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ThriftParserIDENTIFIER)
}

func (s *ServiceContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, i)
}

func (s *ServiceContext) AllFunction_() []IFunction_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunction_Context); ok {
			len++
		}
	}

	tst := make([]IFunction_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunction_Context); ok {
			tst[i] = t.(IFunction_Context)
			i++
		}
	}

	return tst
}

func (s *ServiceContext) Function_(i int) IFunction_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_Context); ok {
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

	return t.(IFunction_Context)
}

func (s *ServiceContext) Type_annotations() IType_annotationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_annotationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *ServiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterService(s)
	}
}

func (s *ServiceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitService(s)
	}
}

func (p *ThriftParser) Service() (localctx IServiceContext) {
	this := p
	_ = this

	localctx = NewServiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ThriftParserRULE_service)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(ThriftParserT__16)
	}
	{
		p.SetState(224)
		p.Match(ThriftParserIDENTIFIER)
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__17 {
		{
			p.SetState(225)
			p.Match(ThriftParserT__17)
		}
		{
			p.SetState(226)
			p.Match(ThriftParserIDENTIFIER)
		}

	}
	{
		p.SetState(229)
		p.Match(ThriftParserT__10)
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&843332246831104) != 0 {
		{
			p.SetState(230)
			p.Function_()
		}

		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(236)
		p.Match(ThriftParserT__11)
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(237)
			p.Type_annotations()
		}

	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field_type() IField_typeContext
	IDENTIFIER() antlr.TerminalNode
	Field_id() IField_idContext
	Field_req() IField_reqContext
	Const_value() IConst_valueContext
	Type_annotations() IType_annotationsContext
	List_separator() IList_separatorContext

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Field_type() IField_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_typeContext)
}

func (s *FieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *FieldContext) Field_id() IField_idContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_idContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_idContext)
}

func (s *FieldContext) Field_req() IField_reqContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_reqContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_reqContext)
}

func (s *FieldContext) Const_value() IConst_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_valueContext)
}

func (s *FieldContext) Type_annotations() IType_annotationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_annotationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *FieldContext) List_separator() IList_separatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_separatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_separatorContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *ThriftParser) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ThriftParserRULE_field)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserINTEGER || _la == ThriftParserHEX_INTEGER {
		{
			p.SetState(240)
			p.Field_id()
		}

	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__19 || _la == ThriftParserT__20 {
		{
			p.SetState(243)
			p.Field_req()
		}

	}
	{
		p.SetState(246)
		p.Field_type()
	}
	{
		p.SetState(247)
		p.Match(ThriftParserIDENTIFIER)
	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__7 {
		{
			p.SetState(248)
			p.Match(ThriftParserT__7)
		}
		{
			p.SetState(249)
			p.Const_value()
		}

	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(252)
			p.Type_annotations()
		}

	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__35 || _la == ThriftParserCOMMA {
		{
			p.SetState(255)
			p.List_separator()
		}

	}

	return localctx
}

// IField_idContext is an interface to support dynamic dispatch.
type IField_idContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Integer() IIntegerContext

	// IsField_idContext differentiates from other interfaces.
	IsField_idContext()
}

type Field_idContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_idContext() *Field_idContext {
	var p = new(Field_idContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_field_id
	return p
}

func (*Field_idContext) IsField_idContext() {}

func NewField_idContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_idContext {
	var p = new(Field_idContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_field_id

	return p
}

func (s *Field_idContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_idContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *Field_idContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_idContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_idContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterField_id(s)
	}
}

func (s *Field_idContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitField_id(s)
	}
}

func (p *ThriftParser) Field_id() (localctx IField_idContext) {
	this := p
	_ = this

	localctx = NewField_idContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ThriftParserRULE_field_id)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		p.Integer()
	}
	{
		p.SetState(259)
		p.Match(ThriftParserT__18)
	}

	return localctx
}

// IField_reqContext is an interface to support dynamic dispatch.
type IField_reqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsField_reqContext differentiates from other interfaces.
	IsField_reqContext()
}

type Field_reqContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_reqContext() *Field_reqContext {
	var p = new(Field_reqContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_field_req
	return p
}

func (*Field_reqContext) IsField_reqContext() {}

func NewField_reqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_reqContext {
	var p = new(Field_reqContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_field_req

	return p
}

func (s *Field_reqContext) GetParser() antlr.Parser { return s.parser }
func (s *Field_reqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_reqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_reqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterField_req(s)
	}
}

func (s *Field_reqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitField_req(s)
	}
}

func (p *ThriftParser) Field_req() (localctx IField_reqContext) {
	this := p
	_ = this

	localctx = NewField_reqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ThriftParserRULE_field_req)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ThriftParserT__19 || _la == ThriftParserT__20) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunction_Context is an interface to support dynamic dispatch.
type IFunction_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Function_type() IFunction_typeContext
	IDENTIFIER() antlr.TerminalNode
	Oneway() IOnewayContext
	AllField() []IFieldContext
	Field(i int) IFieldContext
	Throws_list() IThrows_listContext
	Type_annotations() IType_annotationsContext
	List_separator() IList_separatorContext

	// IsFunction_Context differentiates from other interfaces.
	IsFunction_Context()
}

type Function_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_Context() *Function_Context {
	var p = new(Function_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_function_
	return p
}

func (*Function_Context) IsFunction_Context() {}

func NewFunction_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_Context {
	var p = new(Function_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_function_

	return p
}

func (s *Function_Context) GetParser() antlr.Parser { return s.parser }

func (s *Function_Context) Function_type() IFunction_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_typeContext)
}

func (s *Function_Context) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *Function_Context) Oneway() IOnewayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOnewayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOnewayContext)
}

func (s *Function_Context) AllField() []IFieldContext {
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

func (s *Function_Context) Field(i int) IFieldContext {
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

func (s *Function_Context) Throws_list() IThrows_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThrows_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IThrows_listContext)
}

func (s *Function_Context) Type_annotations() IType_annotationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_annotationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *Function_Context) List_separator() IList_separatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_separatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_separatorContext)
}

func (s *Function_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterFunction_(s)
	}
}

func (s *Function_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitFunction_(s)
	}
}

func (p *ThriftParser) Function_() (localctx IFunction_Context) {
	this := p
	_ = this

	localctx = NewFunction_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ThriftParserRULE_function_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__23 || _la == ThriftParserT__24 {
		{
			p.SetState(263)
			p.Oneway()
		}

	}
	{
		p.SetState(266)
		p.Function_type()
	}
	{
		p.SetState(267)
		p.Match(ThriftParserIDENTIFIER)
	}
	{
		p.SetState(268)
		p.Match(ThriftParserT__21)
	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&843744449396736) != 0 {
		{
			p.SetState(269)
			p.Field()
		}

		p.SetState(274)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(275)
		p.Match(ThriftParserT__22)
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__26 {
		{
			p.SetState(276)
			p.Throws_list()
		}

	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(279)
			p.Type_annotations()
		}

	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__35 || _la == ThriftParserCOMMA {
		{
			p.SetState(282)
			p.List_separator()
		}

	}

	return localctx
}

// IOnewayContext is an interface to support dynamic dispatch.
type IOnewayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOnewayContext differentiates from other interfaces.
	IsOnewayContext()
}

type OnewayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOnewayContext() *OnewayContext {
	var p = new(OnewayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_oneway
	return p
}

func (*OnewayContext) IsOnewayContext() {}

func NewOnewayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OnewayContext {
	var p = new(OnewayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_oneway

	return p
}

func (s *OnewayContext) GetParser() antlr.Parser { return s.parser }
func (s *OnewayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OnewayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OnewayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterOneway(s)
	}
}

func (s *OnewayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitOneway(s)
	}
}

func (p *ThriftParser) Oneway() (localctx IOnewayContext) {
	this := p
	_ = this

	localctx = NewOnewayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ThriftParserRULE_oneway)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ThriftParserT__23 || _la == ThriftParserT__24) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunction_typeContext is an interface to support dynamic dispatch.
type IFunction_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field_type() IField_typeContext

	// IsFunction_typeContext differentiates from other interfaces.
	IsFunction_typeContext()
}

type Function_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_typeContext() *Function_typeContext {
	var p = new(Function_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_function_type
	return p
}

func (*Function_typeContext) IsFunction_typeContext() {}

func NewFunction_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_typeContext {
	var p = new(Function_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_function_type

	return p
}

func (s *Function_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_typeContext) Field_type() IField_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_typeContext)
}

func (s *Function_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterFunction_type(s)
	}
}

func (s *Function_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitFunction_type(s)
	}
}

func (p *ThriftParser) Function_type() (localctx IFunction_typeContext) {
	this := p
	_ = this

	localctx = NewFunction_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ThriftParserRULE_function_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(289)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThriftParserT__27, ThriftParserT__30, ThriftParserT__31, ThriftParserTYPE_BOOL, ThriftParserTYPE_BYTE, ThriftParserTYPE_I16, ThriftParserTYPE_I32, ThriftParserTYPE_I64, ThriftParserTYPE_DOUBLE, ThriftParserTYPE_STRING, ThriftParserTYPE_BINARY, ThriftParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(287)
			p.Field_type()
		}

	case ThriftParserT__25:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(288)
			p.Match(ThriftParserT__25)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IThrows_listContext is an interface to support dynamic dispatch.
type IThrows_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllField() []IFieldContext
	Field(i int) IFieldContext

	// IsThrows_listContext differentiates from other interfaces.
	IsThrows_listContext()
}

type Throws_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThrows_listContext() *Throws_listContext {
	var p = new(Throws_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_throws_list
	return p
}

func (*Throws_listContext) IsThrows_listContext() {}

func NewThrows_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Throws_listContext {
	var p = new(Throws_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_throws_list

	return p
}

func (s *Throws_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Throws_listContext) AllField() []IFieldContext {
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

func (s *Throws_listContext) Field(i int) IFieldContext {
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

func (s *Throws_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Throws_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Throws_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterThrows_list(s)
	}
}

func (s *Throws_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitThrows_list(s)
	}
}

func (p *ThriftParser) Throws_list() (localctx IThrows_listContext) {
	this := p
	_ = this

	localctx = NewThrows_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ThriftParserRULE_throws_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.Match(ThriftParserT__26)
	}
	{
		p.SetState(292)
		p.Match(ThriftParserT__21)
	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&843744449396736) != 0 {
		{
			p.SetState(293)
			p.Field()
		}

		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(299)
		p.Match(ThriftParserT__22)
	}

	return localctx
}

// IType_annotationsContext is an interface to support dynamic dispatch.
type IType_annotationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllType_annotation() []IType_annotationContext
	Type_annotation(i int) IType_annotationContext

	// IsType_annotationsContext differentiates from other interfaces.
	IsType_annotationsContext()
}

type Type_annotationsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_annotationsContext() *Type_annotationsContext {
	var p = new(Type_annotationsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_type_annotations
	return p
}

func (*Type_annotationsContext) IsType_annotationsContext() {}

func NewType_annotationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_annotationsContext {
	var p = new(Type_annotationsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_type_annotations

	return p
}

func (s *Type_annotationsContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_annotationsContext) AllType_annotation() []IType_annotationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType_annotationContext); ok {
			len++
		}
	}

	tst := make([]IType_annotationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType_annotationContext); ok {
			tst[i] = t.(IType_annotationContext)
			i++
		}
	}

	return tst
}

func (s *Type_annotationsContext) Type_annotation(i int) IType_annotationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_annotationContext); ok {
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

	return t.(IType_annotationContext)
}

func (s *Type_annotationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_annotationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_annotationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterType_annotations(s)
	}
}

func (s *Type_annotationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitType_annotations(s)
	}
}

func (p *ThriftParser) Type_annotations() (localctx IType_annotationsContext) {
	this := p
	_ = this

	localctx = NewType_annotationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ThriftParserRULE_type_annotations)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.Match(ThriftParserT__21)
	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ThriftParserIDENTIFIER {
		{
			p.SetState(302)
			p.Type_annotation()
		}

		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(308)
		p.Match(ThriftParserT__22)
	}

	return localctx
}

// IType_annotationContext is an interface to support dynamic dispatch.
type IType_annotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Annotation_value() IAnnotation_valueContext
	List_separator() IList_separatorContext

	// IsType_annotationContext differentiates from other interfaces.
	IsType_annotationContext()
}

type Type_annotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_annotationContext() *Type_annotationContext {
	var p = new(Type_annotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_type_annotation
	return p
}

func (*Type_annotationContext) IsType_annotationContext() {}

func NewType_annotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_annotationContext {
	var p = new(Type_annotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_type_annotation

	return p
}

func (s *Type_annotationContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_annotationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *Type_annotationContext) Annotation_value() IAnnotation_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnnotation_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnnotation_valueContext)
}

func (s *Type_annotationContext) List_separator() IList_separatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_separatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_separatorContext)
}

func (s *Type_annotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_annotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_annotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterType_annotation(s)
	}
}

func (s *Type_annotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitType_annotation(s)
	}
}

func (p *ThriftParser) Type_annotation() (localctx IType_annotationContext) {
	this := p
	_ = this

	localctx = NewType_annotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ThriftParserRULE_type_annotation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(ThriftParserIDENTIFIER)
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__7 {
		{
			p.SetState(311)
			p.Match(ThriftParserT__7)
		}
		{
			p.SetState(312)
			p.Annotation_value()
		}

	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__35 || _la == ThriftParserCOMMA {
		{
			p.SetState(315)
			p.List_separator()
		}

	}

	return localctx
}

// IAnnotation_valueContext is an interface to support dynamic dispatch.
type IAnnotation_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Integer() IIntegerContext
	LITERAL() antlr.TerminalNode

	// IsAnnotation_valueContext differentiates from other interfaces.
	IsAnnotation_valueContext()
}

type Annotation_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotation_valueContext() *Annotation_valueContext {
	var p = new(Annotation_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_annotation_value
	return p
}

func (*Annotation_valueContext) IsAnnotation_valueContext() {}

func NewAnnotation_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Annotation_valueContext {
	var p = new(Annotation_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_annotation_value

	return p
}

func (s *Annotation_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Annotation_valueContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *Annotation_valueContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(ThriftParserLITERAL, 0)
}

func (s *Annotation_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Annotation_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Annotation_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterAnnotation_value(s)
	}
}

func (s *Annotation_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitAnnotation_value(s)
	}
}

func (p *ThriftParser) Annotation_value() (localctx IAnnotation_valueContext) {
	this := p
	_ = this

	localctx = NewAnnotation_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ThriftParserRULE_annotation_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(320)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThriftParserINTEGER, ThriftParserHEX_INTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(318)
			p.Integer()
		}

	case ThriftParserLITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(319)
			p.Match(ThriftParserLITERAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IField_typeContext is an interface to support dynamic dispatch.
type IField_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Base_type() IBase_typeContext
	IDENTIFIER() antlr.TerminalNode
	Container_type() IContainer_typeContext

	// IsField_typeContext differentiates from other interfaces.
	IsField_typeContext()
}

type Field_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_typeContext() *Field_typeContext {
	var p = new(Field_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_field_type
	return p
}

func (*Field_typeContext) IsField_typeContext() {}

func NewField_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_typeContext {
	var p = new(Field_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_field_type

	return p
}

func (s *Field_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_typeContext) Base_type() IBase_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBase_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBase_typeContext)
}

func (s *Field_typeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *Field_typeContext) Container_type() IContainer_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContainer_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContainer_typeContext)
}

func (s *Field_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterField_type(s)
	}
}

func (s *Field_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitField_type(s)
	}
}

func (p *ThriftParser) Field_type() (localctx IField_typeContext) {
	this := p
	_ = this

	localctx = NewField_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ThriftParserRULE_field_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(325)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThriftParserTYPE_BOOL, ThriftParserTYPE_BYTE, ThriftParserTYPE_I16, ThriftParserTYPE_I32, ThriftParserTYPE_I64, ThriftParserTYPE_DOUBLE, ThriftParserTYPE_STRING, ThriftParserTYPE_BINARY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(322)
			p.Base_type()
		}

	case ThriftParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(323)
			p.Match(ThriftParserIDENTIFIER)
		}

	case ThriftParserT__27, ThriftParserT__30, ThriftParserT__31:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(324)
			p.Container_type()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBase_typeContext is an interface to support dynamic dispatch.
type IBase_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Real_base_type() IReal_base_typeContext
	Type_annotations() IType_annotationsContext

	// IsBase_typeContext differentiates from other interfaces.
	IsBase_typeContext()
}

type Base_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBase_typeContext() *Base_typeContext {
	var p = new(Base_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_base_type
	return p
}

func (*Base_typeContext) IsBase_typeContext() {}

func NewBase_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Base_typeContext {
	var p = new(Base_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_base_type

	return p
}

func (s *Base_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Base_typeContext) Real_base_type() IReal_base_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReal_base_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReal_base_typeContext)
}

func (s *Base_typeContext) Type_annotations() IType_annotationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_annotationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *Base_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Base_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Base_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterBase_type(s)
	}
}

func (s *Base_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitBase_type(s)
	}
}

func (p *ThriftParser) Base_type() (localctx IBase_typeContext) {
	this := p
	_ = this

	localctx = NewBase_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ThriftParserRULE_base_type)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(327)
		p.Real_base_type()
	}
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(328)
			p.Type_annotations()
		}

	}

	return localctx
}

// IContainer_typeContext is an interface to support dynamic dispatch.
type IContainer_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Map_type() IMap_typeContext
	Set_type() ISet_typeContext
	List_type() IList_typeContext
	Type_annotations() IType_annotationsContext

	// IsContainer_typeContext differentiates from other interfaces.
	IsContainer_typeContext()
}

type Container_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContainer_typeContext() *Container_typeContext {
	var p = new(Container_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_container_type
	return p
}

func (*Container_typeContext) IsContainer_typeContext() {}

func NewContainer_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Container_typeContext {
	var p = new(Container_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_container_type

	return p
}

func (s *Container_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Container_typeContext) Map_type() IMap_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMap_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMap_typeContext)
}

func (s *Container_typeContext) Set_type() ISet_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISet_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISet_typeContext)
}

func (s *Container_typeContext) List_type() IList_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_typeContext)
}

func (s *Container_typeContext) Type_annotations() IType_annotationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_annotationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *Container_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Container_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Container_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterContainer_type(s)
	}
}

func (s *Container_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitContainer_type(s)
	}
}

func (p *ThriftParser) Container_type() (localctx IContainer_typeContext) {
	this := p
	_ = this

	localctx = NewContainer_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ThriftParserRULE_container_type)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(334)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThriftParserT__27:
		{
			p.SetState(331)
			p.Map_type()
		}

	case ThriftParserT__30:
		{
			p.SetState(332)
			p.Set_type()
		}

	case ThriftParserT__31:
		{
			p.SetState(333)
			p.List_type()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(336)
			p.Type_annotations()
		}

	}

	return localctx
}

// IMap_typeContext is an interface to support dynamic dispatch.
type IMap_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllField_type() []IField_typeContext
	Field_type(i int) IField_typeContext
	COMMA() antlr.TerminalNode
	Cpp_type() ICpp_typeContext

	// IsMap_typeContext differentiates from other interfaces.
	IsMap_typeContext()
}

type Map_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMap_typeContext() *Map_typeContext {
	var p = new(Map_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_map_type
	return p
}

func (*Map_typeContext) IsMap_typeContext() {}

func NewMap_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Map_typeContext {
	var p = new(Map_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_map_type

	return p
}

func (s *Map_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Map_typeContext) AllField_type() []IField_typeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IField_typeContext); ok {
			len++
		}
	}

	tst := make([]IField_typeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IField_typeContext); ok {
			tst[i] = t.(IField_typeContext)
			i++
		}
	}

	return tst
}

func (s *Map_typeContext) Field_type(i int) IField_typeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_typeContext); ok {
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

	return t.(IField_typeContext)
}

func (s *Map_typeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ThriftParserCOMMA, 0)
}

func (s *Map_typeContext) Cpp_type() ICpp_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICpp_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICpp_typeContext)
}

func (s *Map_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Map_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Map_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterMap_type(s)
	}
}

func (s *Map_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitMap_type(s)
	}
}

func (p *ThriftParser) Map_type() (localctx IMap_typeContext) {
	this := p
	_ = this

	localctx = NewMap_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ThriftParserRULE_map_type)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Match(ThriftParserT__27)
	}
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__32 {
		{
			p.SetState(340)
			p.Cpp_type()
		}

	}
	{
		p.SetState(343)
		p.Match(ThriftParserT__28)
	}
	{
		p.SetState(344)
		p.Field_type()
	}
	{
		p.SetState(345)
		p.Match(ThriftParserCOMMA)
	}
	{
		p.SetState(346)
		p.Field_type()
	}
	{
		p.SetState(347)
		p.Match(ThriftParserT__29)
	}

	return localctx
}

// ISet_typeContext is an interface to support dynamic dispatch.
type ISet_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field_type() IField_typeContext
	Cpp_type() ICpp_typeContext

	// IsSet_typeContext differentiates from other interfaces.
	IsSet_typeContext()
}

type Set_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_typeContext() *Set_typeContext {
	var p = new(Set_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_set_type
	return p
}

func (*Set_typeContext) IsSet_typeContext() {}

func NewSet_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_typeContext {
	var p = new(Set_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_set_type

	return p
}

func (s *Set_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_typeContext) Field_type() IField_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_typeContext)
}

func (s *Set_typeContext) Cpp_type() ICpp_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICpp_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICpp_typeContext)
}

func (s *Set_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterSet_type(s)
	}
}

func (s *Set_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitSet_type(s)
	}
}

func (p *ThriftParser) Set_type() (localctx ISet_typeContext) {
	this := p
	_ = this

	localctx = NewSet_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ThriftParserRULE_set_type)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(349)
		p.Match(ThriftParserT__30)
	}
	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__32 {
		{
			p.SetState(350)
			p.Cpp_type()
		}

	}
	{
		p.SetState(353)
		p.Match(ThriftParserT__28)
	}
	{
		p.SetState(354)
		p.Field_type()
	}
	{
		p.SetState(355)
		p.Match(ThriftParserT__29)
	}

	return localctx
}

// IList_typeContext is an interface to support dynamic dispatch.
type IList_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field_type() IField_typeContext
	Cpp_type() ICpp_typeContext

	// IsList_typeContext differentiates from other interfaces.
	IsList_typeContext()
}

type List_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_typeContext() *List_typeContext {
	var p = new(List_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_list_type
	return p
}

func (*List_typeContext) IsList_typeContext() {}

func NewList_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_typeContext {
	var p = new(List_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_list_type

	return p
}

func (s *List_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *List_typeContext) Field_type() IField_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_typeContext)
}

func (s *List_typeContext) Cpp_type() ICpp_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICpp_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICpp_typeContext)
}

func (s *List_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterList_type(s)
	}
}

func (s *List_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitList_type(s)
	}
}

func (p *ThriftParser) List_type() (localctx IList_typeContext) {
	this := p
	_ = this

	localctx = NewList_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ThriftParserRULE_list_type)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.Match(ThriftParserT__31)
	}
	{
		p.SetState(358)
		p.Match(ThriftParserT__28)
	}
	{
		p.SetState(359)
		p.Field_type()
	}
	{
		p.SetState(360)
		p.Match(ThriftParserT__29)
	}
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__32 {
		{
			p.SetState(361)
			p.Cpp_type()
		}

	}

	return localctx
}

// ICpp_typeContext is an interface to support dynamic dispatch.
type ICpp_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LITERAL() antlr.TerminalNode

	// IsCpp_typeContext differentiates from other interfaces.
	IsCpp_typeContext()
}

type Cpp_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCpp_typeContext() *Cpp_typeContext {
	var p = new(Cpp_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_cpp_type
	return p
}

func (*Cpp_typeContext) IsCpp_typeContext() {}

func NewCpp_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cpp_typeContext {
	var p = new(Cpp_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_cpp_type

	return p
}

func (s *Cpp_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Cpp_typeContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(ThriftParserLITERAL, 0)
}

func (s *Cpp_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cpp_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cpp_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterCpp_type(s)
	}
}

func (s *Cpp_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitCpp_type(s)
	}
}

func (p *ThriftParser) Cpp_type() (localctx ICpp_typeContext) {
	this := p
	_ = this

	localctx = NewCpp_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ThriftParserRULE_cpp_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(364)
		p.Match(ThriftParserT__32)
	}
	{
		p.SetState(365)
		p.Match(ThriftParserLITERAL)
	}

	return localctx
}

// IConst_valueContext is an interface to support dynamic dispatch.
type IConst_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Integer() IIntegerContext
	DOUBLE() antlr.TerminalNode
	LITERAL() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	Const_list() IConst_listContext
	Const_map() IConst_mapContext

	// IsConst_valueContext differentiates from other interfaces.
	IsConst_valueContext()
}

type Const_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_valueContext() *Const_valueContext {
	var p = new(Const_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_const_value
	return p
}

func (*Const_valueContext) IsConst_valueContext() {}

func NewConst_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_valueContext {
	var p = new(Const_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_const_value

	return p
}

func (s *Const_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_valueContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *Const_valueContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(ThriftParserDOUBLE, 0)
}

func (s *Const_valueContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(ThriftParserLITERAL, 0)
}

func (s *Const_valueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *Const_valueContext) Const_list() IConst_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_listContext)
}

func (s *Const_valueContext) Const_map() IConst_mapContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_mapContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_mapContext)
}

func (s *Const_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterConst_value(s)
	}
}

func (s *Const_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitConst_value(s)
	}
}

func (p *ThriftParser) Const_value() (localctx IConst_valueContext) {
	this := p
	_ = this

	localctx = NewConst_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ThriftParserRULE_const_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(373)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThriftParserINTEGER, ThriftParserHEX_INTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(367)
			p.Integer()
		}

	case ThriftParserDOUBLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(368)
			p.Match(ThriftParserDOUBLE)
		}

	case ThriftParserLITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(369)
			p.Match(ThriftParserLITERAL)
		}

	case ThriftParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(370)
			p.Match(ThriftParserIDENTIFIER)
		}

	case ThriftParserT__33:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(371)
			p.Const_list()
		}

	case ThriftParserT__10:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(372)
			p.Const_map()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER() antlr.TerminalNode
	HEX_INTEGER() antlr.TerminalNode

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(ThriftParserINTEGER, 0)
}

func (s *IntegerContext) HEX_INTEGER() antlr.TerminalNode {
	return s.GetToken(ThriftParserHEX_INTEGER, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (p *ThriftParser) Integer() (localctx IIntegerContext) {
	this := p
	_ = this

	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ThriftParserRULE_integer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(375)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ThriftParserINTEGER || _la == ThriftParserHEX_INTEGER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IConst_listContext is an interface to support dynamic dispatch.
type IConst_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllConst_value() []IConst_valueContext
	Const_value(i int) IConst_valueContext
	AllList_separator() []IList_separatorContext
	List_separator(i int) IList_separatorContext

	// IsConst_listContext differentiates from other interfaces.
	IsConst_listContext()
}

type Const_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_listContext() *Const_listContext {
	var p = new(Const_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_const_list
	return p
}

func (*Const_listContext) IsConst_listContext() {}

func NewConst_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_listContext {
	var p = new(Const_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_const_list

	return p
}

func (s *Const_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_listContext) AllConst_value() []IConst_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConst_valueContext); ok {
			len++
		}
	}

	tst := make([]IConst_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConst_valueContext); ok {
			tst[i] = t.(IConst_valueContext)
			i++
		}
	}

	return tst
}

func (s *Const_listContext) Const_value(i int) IConst_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_valueContext); ok {
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

	return t.(IConst_valueContext)
}

func (s *Const_listContext) AllList_separator() []IList_separatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IList_separatorContext); ok {
			len++
		}
	}

	tst := make([]IList_separatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IList_separatorContext); ok {
			tst[i] = t.(IList_separatorContext)
			i++
		}
	}

	return tst
}

func (s *Const_listContext) List_separator(i int) IList_separatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_separatorContext); ok {
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

	return t.(IList_separatorContext)
}

func (s *Const_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterConst_list(s)
	}
}

func (s *Const_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitConst_list(s)
	}
}

func (p *ThriftParser) Const_list() (localctx IConst_listContext) {
	this := p
	_ = this

	localctx = NewConst_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ThriftParserRULE_const_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(377)
		p.Match(ThriftParserT__33)
	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&845404182677504) != 0 {
		{
			p.SetState(378)
			p.Const_value()
		}
		p.SetState(380)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ThriftParserT__35 || _la == ThriftParserCOMMA {
			{
				p.SetState(379)
				p.List_separator()
			}

		}

		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(387)
		p.Match(ThriftParserT__34)
	}

	return localctx
}

// IConst_map_entryContext is an interface to support dynamic dispatch.
type IConst_map_entryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllConst_value() []IConst_valueContext
	Const_value(i int) IConst_valueContext
	List_separator() IList_separatorContext

	// IsConst_map_entryContext differentiates from other interfaces.
	IsConst_map_entryContext()
}

type Const_map_entryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_map_entryContext() *Const_map_entryContext {
	var p = new(Const_map_entryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_const_map_entry
	return p
}

func (*Const_map_entryContext) IsConst_map_entryContext() {}

func NewConst_map_entryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_map_entryContext {
	var p = new(Const_map_entryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_const_map_entry

	return p
}

func (s *Const_map_entryContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_map_entryContext) AllConst_value() []IConst_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConst_valueContext); ok {
			len++
		}
	}

	tst := make([]IConst_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConst_valueContext); ok {
			tst[i] = t.(IConst_valueContext)
			i++
		}
	}

	return tst
}

func (s *Const_map_entryContext) Const_value(i int) IConst_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_valueContext); ok {
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

	return t.(IConst_valueContext)
}

func (s *Const_map_entryContext) List_separator() IList_separatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_separatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_separatorContext)
}

func (s *Const_map_entryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_map_entryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_map_entryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterConst_map_entry(s)
	}
}

func (s *Const_map_entryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitConst_map_entry(s)
	}
}

func (p *ThriftParser) Const_map_entry() (localctx IConst_map_entryContext) {
	this := p
	_ = this

	localctx = NewConst_map_entryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ThriftParserRULE_const_map_entry)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(389)
		p.Const_value()
	}
	{
		p.SetState(390)
		p.Match(ThriftParserT__18)
	}
	{
		p.SetState(391)
		p.Const_value()
	}
	p.SetState(393)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__35 || _la == ThriftParserCOMMA {
		{
			p.SetState(392)
			p.List_separator()
		}

	}

	return localctx
}

// IConst_mapContext is an interface to support dynamic dispatch.
type IConst_mapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllConst_map_entry() []IConst_map_entryContext
	Const_map_entry(i int) IConst_map_entryContext

	// IsConst_mapContext differentiates from other interfaces.
	IsConst_mapContext()
}

type Const_mapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_mapContext() *Const_mapContext {
	var p = new(Const_mapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_const_map
	return p
}

func (*Const_mapContext) IsConst_mapContext() {}

func NewConst_mapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_mapContext {
	var p = new(Const_mapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_const_map

	return p
}

func (s *Const_mapContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_mapContext) AllConst_map_entry() []IConst_map_entryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConst_map_entryContext); ok {
			len++
		}
	}

	tst := make([]IConst_map_entryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConst_map_entryContext); ok {
			tst[i] = t.(IConst_map_entryContext)
			i++
		}
	}

	return tst
}

func (s *Const_mapContext) Const_map_entry(i int) IConst_map_entryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_map_entryContext); ok {
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

	return t.(IConst_map_entryContext)
}

func (s *Const_mapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_mapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_mapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterConst_map(s)
	}
}

func (s *Const_mapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitConst_map(s)
	}
}

func (p *ThriftParser) Const_map() (localctx IConst_mapContext) {
	this := p
	_ = this

	localctx = NewConst_mapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ThriftParserRULE_const_map)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(395)
		p.Match(ThriftParserT__10)
	}
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&845404182677504) != 0 {
		{
			p.SetState(396)
			p.Const_map_entry()
		}

		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(402)
		p.Match(ThriftParserT__11)
	}

	return localctx
}

// IList_separatorContext is an interface to support dynamic dispatch.
type IList_separatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMA() antlr.TerminalNode

	// IsList_separatorContext differentiates from other interfaces.
	IsList_separatorContext()
}

type List_separatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_separatorContext() *List_separatorContext {
	var p = new(List_separatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_list_separator
	return p
}

func (*List_separatorContext) IsList_separatorContext() {}

func NewList_separatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_separatorContext {
	var p = new(List_separatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_list_separator

	return p
}

func (s *List_separatorContext) GetParser() antlr.Parser { return s.parser }

func (s *List_separatorContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ThriftParserCOMMA, 0)
}

func (s *List_separatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_separatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_separatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterList_separator(s)
	}
}

func (s *List_separatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitList_separator(s)
	}
}

func (p *ThriftParser) List_separator() (localctx IList_separatorContext) {
	this := p
	_ = this

	localctx = NewList_separatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ThriftParserRULE_list_separator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(404)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ThriftParserT__35 || _la == ThriftParserCOMMA) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IReal_base_typeContext is an interface to support dynamic dispatch.
type IReal_base_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE_BOOL() antlr.TerminalNode
	TYPE_BYTE() antlr.TerminalNode
	TYPE_I16() antlr.TerminalNode
	TYPE_I32() antlr.TerminalNode
	TYPE_I64() antlr.TerminalNode
	TYPE_DOUBLE() antlr.TerminalNode
	TYPE_STRING() antlr.TerminalNode
	TYPE_BINARY() antlr.TerminalNode

	// IsReal_base_typeContext differentiates from other interfaces.
	IsReal_base_typeContext()
}

type Real_base_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReal_base_typeContext() *Real_base_typeContext {
	var p = new(Real_base_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_real_base_type
	return p
}

func (*Real_base_typeContext) IsReal_base_typeContext() {}

func NewReal_base_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Real_base_typeContext {
	var p = new(Real_base_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_real_base_type

	return p
}

func (s *Real_base_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Real_base_typeContext) TYPE_BOOL() antlr.TerminalNode {
	return s.GetToken(ThriftParserTYPE_BOOL, 0)
}

func (s *Real_base_typeContext) TYPE_BYTE() antlr.TerminalNode {
	return s.GetToken(ThriftParserTYPE_BYTE, 0)
}

func (s *Real_base_typeContext) TYPE_I16() antlr.TerminalNode {
	return s.GetToken(ThriftParserTYPE_I16, 0)
}

func (s *Real_base_typeContext) TYPE_I32() antlr.TerminalNode {
	return s.GetToken(ThriftParserTYPE_I32, 0)
}

func (s *Real_base_typeContext) TYPE_I64() antlr.TerminalNode {
	return s.GetToken(ThriftParserTYPE_I64, 0)
}

func (s *Real_base_typeContext) TYPE_DOUBLE() antlr.TerminalNode {
	return s.GetToken(ThriftParserTYPE_DOUBLE, 0)
}

func (s *Real_base_typeContext) TYPE_STRING() antlr.TerminalNode {
	return s.GetToken(ThriftParserTYPE_STRING, 0)
}

func (s *Real_base_typeContext) TYPE_BINARY() antlr.TerminalNode {
	return s.GetToken(ThriftParserTYPE_BINARY, 0)
}

func (s *Real_base_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Real_base_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Real_base_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterReal_base_type(s)
	}
}

func (s *Real_base_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitReal_base_type(s)
	}
}

func (p *ThriftParser) Real_base_type() (localctx IReal_base_typeContext) {
	this := p
	_ = this

	localctx = NewReal_base_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ThriftParserRULE_real_base_type)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(406)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&280375465082880) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

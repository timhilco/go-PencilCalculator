// Code generated from HilcoPencilParser.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 51, 409,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 3, 2, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 6, 22, 214, 10, 22, 13, 22, 14, 22,
	215, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 7, 23, 225, 10, 23,
	12, 23, 14, 23, 228, 11, 23, 3, 24, 3, 24, 3, 24, 7, 24, 233, 10, 24, 12,
	24, 14, 24, 236, 11, 24, 3, 24, 3, 24, 3, 25, 3, 25, 5, 25, 242, 10, 25,
	3, 25, 6, 25, 245, 10, 25, 13, 25, 14, 25, 246, 3, 26, 3, 26, 3, 27, 3,
	27, 3, 27, 3, 27, 5, 27, 255, 10, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 266, 10, 28, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 30, 6, 30, 276, 10, 30, 13, 30, 14, 30, 277,
	3, 30, 3, 30, 7, 30, 282, 10, 30, 12, 30, 14, 30, 285, 11, 30, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 296, 10, 31,
	3, 31, 3, 31, 5, 31, 300, 10, 31, 3, 31, 3, 31, 5, 31, 304, 10, 31, 5,
	31, 306, 10, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 314,
	10, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 5, 32, 321, 10, 32, 3, 32, 3,
	32, 3, 32, 7, 32, 326, 10, 32, 12, 32, 14, 32, 329, 11, 32, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 7, 33, 336, 10, 33, 12, 33, 14, 33, 339, 11, 33, 3,
	34, 3, 34, 5, 34, 343, 10, 34, 3, 34, 3, 34, 3, 35, 5, 35, 348, 10, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 5, 36, 358, 10,
	36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40,
	3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3,
	46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 50,
	3, 50, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 54, 3,
	54, 3, 55, 3, 55, 3, 56, 3, 56, 3, 57, 3, 57, 3, 58, 3, 58, 2, 2, 59, 3,
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 2, 39, 2, 41, 2, 43,
	20, 45, 21, 47, 22, 49, 2, 51, 2, 53, 2, 55, 2, 57, 2, 59, 23, 61, 24,
	63, 25, 65, 26, 67, 27, 69, 28, 71, 29, 73, 30, 75, 31, 77, 32, 79, 33,
	81, 34, 83, 35, 85, 36, 87, 37, 89, 38, 91, 39, 93, 40, 95, 41, 97, 42,
	99, 43, 101, 44, 103, 45, 105, 46, 107, 47, 109, 48, 111, 49, 113, 50,
	115, 51, 3, 2, 9, 4, 2, 41, 41, 94, 94, 4, 2, 71, 71, 103, 103, 4, 2, 45,
	45, 47, 47, 5, 2, 50, 59, 67, 72, 99, 104, 10, 2, 36, 36, 41, 41, 94, 94,
	100, 100, 104, 104, 112, 112, 116, 116, 118, 118, 4, 2, 47, 47, 49, 49,
	5, 2, 11, 11, 14, 14, 34, 34, 2, 434, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2,
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2,
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3,
	2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29,
	3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2,
	43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2,
	2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2,
	2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2,
	2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3,
	2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91,
	3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2,
	99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2, 2,
	2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113,
	3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 3, 117, 3, 2, 2, 2, 5, 120, 3, 2, 2, 2,
	7, 126, 3, 2, 2, 2, 9, 131, 3, 2, 2, 2, 11, 137, 3, 2, 2, 2, 13, 141, 3,
	2, 2, 2, 15, 149, 3, 2, 2, 2, 17, 152, 3, 2, 2, 2, 19, 157, 3, 2, 2, 2,
	21, 162, 3, 2, 2, 2, 23, 167, 3, 2, 2, 2, 25, 170, 3, 2, 2, 2, 27, 178,
	3, 2, 2, 2, 29, 185, 3, 2, 2, 2, 31, 195, 3, 2, 2, 2, 33, 199, 3, 2, 2,
	2, 35, 202, 3, 2, 2, 2, 37, 206, 3, 2, 2, 2, 39, 208, 3, 2, 2, 2, 41, 210,
	3, 2, 2, 2, 43, 213, 3, 2, 2, 2, 45, 217, 3, 2, 2, 2, 47, 229, 3, 2, 2,
	2, 49, 239, 3, 2, 2, 2, 51, 248, 3, 2, 2, 2, 53, 254, 3, 2, 2, 2, 55, 265,
	3, 2, 2, 2, 57, 267, 3, 2, 2, 2, 59, 275, 3, 2, 2, 2, 61, 286, 3, 2, 2,
	2, 63, 317, 3, 2, 2, 2, 65, 330, 3, 2, 2, 2, 67, 342, 3, 2, 2, 2, 69, 347,
	3, 2, 2, 2, 71, 357, 3, 2, 2, 2, 73, 361, 3, 2, 2, 2, 75, 363, 3, 2, 2,
	2, 77, 365, 3, 2, 2, 2, 79, 367, 3, 2, 2, 2, 81, 369, 3, 2, 2, 2, 83, 371,
	3, 2, 2, 2, 85, 373, 3, 2, 2, 2, 87, 375, 3, 2, 2, 2, 89, 377, 3, 2, 2,
	2, 91, 379, 3, 2, 2, 2, 93, 382, 3, 2, 2, 2, 95, 384, 3, 2, 2, 2, 97, 386,
	3, 2, 2, 2, 99, 389, 3, 2, 2, 2, 101, 391, 3, 2, 2, 2, 103, 394, 3, 2,
	2, 2, 105, 396, 3, 2, 2, 2, 107, 399, 3, 2, 2, 2, 109, 401, 3, 2, 2, 2,
	111, 403, 3, 2, 2, 2, 113, 405, 3, 2, 2, 2, 115, 407, 3, 2, 2, 2, 117,
	118, 7, 60, 2, 2, 118, 119, 7, 60, 2, 2, 119, 4, 3, 2, 2, 2, 120, 121,
	7, 118, 2, 2, 121, 122, 7, 113, 2, 2, 122, 123, 7, 102, 2, 2, 123, 124,
	7, 99, 2, 2, 124, 125, 7, 123, 2, 2, 125, 6, 3, 2, 2, 2, 126, 127, 7, 118,
	2, 2, 127, 128, 7, 116, 2, 2, 128, 129, 7, 119, 2, 2, 129, 130, 7, 103,
	2, 2, 130, 8, 3, 2, 2, 2, 131, 132, 7, 104, 2, 2, 132, 133, 7, 99, 2, 2,
	133, 134, 7, 110, 2, 2, 134, 135, 7, 117, 2, 2, 135, 136, 7, 103, 2, 2,
	136, 10, 3, 2, 2, 2, 137, 138, 7, 112, 2, 2, 138, 139, 7, 107, 2, 2, 139,
	140, 7, 110, 2, 2, 140, 12, 3, 2, 2, 2, 141, 142, 7, 102, 2, 2, 142, 143,
	7, 103, 2, 2, 143, 144, 7, 104, 2, 2, 144, 145, 7, 99, 2, 2, 145, 146,
	7, 119, 2, 2, 146, 147, 7, 110, 2, 2, 147, 148, 7, 118, 2, 2, 148, 14,
	3, 2, 2, 2, 149, 150, 7, 107, 2, 2, 150, 151, 7, 104, 2, 2, 151, 16, 3,
	2, 2, 2, 152, 153, 7, 118, 2, 2, 153, 154, 7, 106, 2, 2, 154, 155, 7, 103,
	2, 2, 155, 156, 7, 112, 2, 2, 156, 18, 3, 2, 2, 2, 157, 158, 7, 103, 2,
	2, 158, 159, 7, 110, 2, 2, 159, 160, 7, 117, 2, 2, 160, 161, 7, 103, 2,
	2, 161, 20, 3, 2, 2, 2, 162, 163, 7, 101, 2, 2, 163, 164, 7, 99, 2, 2,
	164, 165, 7, 117, 2, 2, 165, 166, 7, 103, 2, 2, 166, 22, 3, 2, 2, 2, 167,
	168, 7, 107, 2, 2, 168, 169, 7, 117, 2, 2, 169, 24, 3, 2, 2, 2, 170, 171,
	7, 103, 2, 2, 171, 172, 7, 112, 2, 2, 172, 173, 7, 102, 2, 2, 173, 174,
	7, 101, 2, 2, 174, 175, 7, 99, 2, 2, 175, 176, 7, 117, 2, 2, 176, 177,
	7, 103, 2, 2, 177, 26, 3, 2, 2, 2, 178, 179, 7, 117, 2, 2, 179, 180, 7,
	121, 2, 2, 180, 181, 7, 107, 2, 2, 181, 182, 7, 118, 2, 2, 182, 183, 7,
	101, 2, 2, 183, 184, 7, 106, 2, 2, 184, 28, 3, 2, 2, 2, 185, 186, 7, 103,
	2, 2, 186, 187, 7, 112, 2, 2, 187, 188, 7, 102, 2, 2, 188, 189, 7, 117,
	2, 2, 189, 190, 7, 121, 2, 2, 190, 191, 7, 107, 2, 2, 191, 192, 7, 118,
	2, 2, 192, 193, 7, 101, 2, 2, 193, 194, 7, 106, 2, 2, 194, 30, 3, 2, 2,
	2, 195, 196, 7, 67, 2, 2, 196, 197, 7, 80, 2, 2, 197, 198, 7, 70, 2, 2,
	198, 32, 3, 2, 2, 2, 199, 200, 7, 81, 2, 2, 200, 201, 7, 84, 2, 2, 201,
	34, 3, 2, 2, 2, 202, 203, 7, 80, 2, 2, 203, 204, 7, 81, 2, 2, 204, 205,
	7, 86, 2, 2, 205, 36, 3, 2, 2, 2, 206, 207, 4, 50, 59, 2, 207, 38, 3, 2,
	2, 2, 208, 209, 4, 99, 124, 2, 209, 40, 3, 2, 2, 2, 210, 211, 4, 67, 92,
	2, 211, 42, 3, 2, 2, 2, 212, 214, 5, 37, 19, 2, 213, 212, 3, 2, 2, 2, 214,
	215, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 44, 3,
	2, 2, 2, 217, 226, 5, 39, 20, 2, 218, 225, 5, 39, 20, 2, 219, 225, 5, 41,
	21, 2, 220, 225, 5, 37, 19, 2, 221, 225, 7, 97, 2, 2, 222, 225, 5, 75,
	38, 2, 223, 225, 7, 94, 2, 2, 224, 218, 3, 2, 2, 2, 224, 219, 3, 2, 2,
	2, 224, 220, 3, 2, 2, 2, 224, 221, 3, 2, 2, 2, 224, 222, 3, 2, 2, 2, 224,
	223, 3, 2, 2, 2, 225, 228, 3, 2, 2, 2, 226, 224, 3, 2, 2, 2, 226, 227,
	3, 2, 2, 2, 227, 46, 3, 2, 2, 2, 228, 226, 3, 2, 2, 2, 229, 234, 7, 41,
	2, 2, 230, 233, 5, 53, 27, 2, 231, 233, 10, 2, 2, 2, 232, 230, 3, 2, 2,
	2, 232, 231, 3, 2, 2, 2, 233, 236, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2, 234,
	235, 3, 2, 2, 2, 235, 237, 3, 2, 2, 2, 236, 234, 3, 2, 2, 2, 237, 238,
	7, 41, 2, 2, 238, 48, 3, 2, 2, 2, 239, 241, 9, 3, 2, 2, 240, 242, 9, 4,
	2, 2, 241, 240, 3, 2, 2, 2, 241, 242, 3, 2, 2, 2, 242, 244, 3, 2, 2, 2,
	243, 245, 4, 50, 59, 2, 244, 243, 3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246,
	244, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247, 50, 3, 2, 2, 2, 248, 249, 9,
	5, 2, 2, 249, 52, 3, 2, 2, 2, 250, 251, 7, 94, 2, 2, 251, 255, 9, 6, 2,
	2, 252, 255, 5, 57, 29, 2, 253, 255, 5, 55, 28, 2, 254, 250, 3, 2, 2, 2,
	254, 252, 3, 2, 2, 2, 254, 253, 3, 2, 2, 2, 255, 54, 3, 2, 2, 2, 256, 257,
	7, 94, 2, 2, 257, 258, 4, 50, 53, 2, 258, 259, 4, 50, 57, 2, 259, 266,
	4, 50, 57, 2, 260, 261, 7, 94, 2, 2, 261, 262, 4, 50, 57, 2, 262, 266,
	4, 50, 57, 2, 263, 264, 7, 94, 2, 2, 264, 266, 4, 50, 57, 2, 265, 256,
	3, 2, 2, 2, 265, 260, 3, 2, 2, 2, 265, 263, 3, 2, 2, 2, 266, 56, 3, 2,
	2, 2, 267, 268, 7, 94, 2, 2, 268, 269, 7, 119, 2, 2, 269, 270, 5, 51, 26,
	2, 270, 271, 5, 51, 26, 2, 271, 272, 5, 51, 26, 2, 272, 273, 5, 51, 26,
	2, 273, 58, 3, 2, 2, 2, 274, 276, 5, 37, 19, 2, 275, 274, 3, 2, 2, 2, 276,
	277, 3, 2, 2, 2, 277, 275, 3, 2, 2, 2, 277, 278, 3, 2, 2, 2, 278, 279,
	3, 2, 2, 2, 279, 283, 7, 48, 2, 2, 280, 282, 5, 37, 19, 2, 281, 280, 3,
	2, 2, 2, 282, 285, 3, 2, 2, 2, 283, 281, 3, 2, 2, 2, 283, 284, 3, 2, 2,
	2, 284, 60, 3, 2, 2, 2, 285, 283, 3, 2, 2, 2, 286, 287, 5, 87, 44, 2, 287,
	288, 5, 37, 19, 2, 288, 289, 5, 37, 19, 2, 289, 305, 9, 7, 2, 2, 290, 291,
	5, 37, 19, 2, 291, 292, 5, 37, 19, 2, 292, 306, 3, 2, 2, 2, 293, 296, 5,
	39, 20, 2, 294, 296, 5, 41, 21, 2, 295, 293, 3, 2, 2, 2, 295, 294, 3, 2,
	2, 2, 296, 299, 3, 2, 2, 2, 297, 300, 5, 39, 20, 2, 298, 300, 5, 41, 21,
	2, 299, 297, 3, 2, 2, 2, 299, 298, 3, 2, 2, 2, 300, 303, 3, 2, 2, 2, 301,
	304, 5, 39, 20, 2, 302, 304, 5, 41, 21, 2, 303, 301, 3, 2, 2, 2, 303, 302,
	3, 2, 2, 2, 304, 306, 3, 2, 2, 2, 305, 290, 3, 2, 2, 2, 305, 295, 3, 2,
	2, 2, 306, 307, 3, 2, 2, 2, 307, 308, 9, 7, 2, 2, 308, 309, 5, 37, 19,
	2, 309, 313, 5, 37, 19, 2, 310, 311, 5, 37, 19, 2, 311, 312, 5, 37, 19,
	2, 312, 314, 3, 2, 2, 2, 313, 310, 3, 2, 2, 2, 313, 314, 3, 2, 2, 2, 314,
	315, 3, 2, 2, 2, 315, 316, 5, 89, 45, 2, 316, 62, 3, 2, 2, 2, 317, 320,
	7, 66, 2, 2, 318, 321, 5, 39, 20, 2, 319, 321, 5, 41, 21, 2, 320, 318,
	3, 2, 2, 2, 320, 319, 3, 2, 2, 2, 321, 327, 3, 2, 2, 2, 322, 326, 5, 39,
	20, 2, 323, 326, 5, 41, 21, 2, 324, 326, 5, 37, 19, 2, 325, 322, 3, 2,
	2, 2, 325, 323, 3, 2, 2, 2, 325, 324, 3, 2, 2, 2, 326, 329, 3, 2, 2, 2,
	327, 325, 3, 2, 2, 2, 327, 328, 3, 2, 2, 2, 328, 64, 3, 2, 2, 2, 329, 327,
	3, 2, 2, 2, 330, 337, 5, 41, 21, 2, 331, 336, 5, 39, 20, 2, 332, 336, 5,
	41, 21, 2, 333, 336, 5, 37, 19, 2, 334, 336, 7, 97, 2, 2, 335, 331, 3,
	2, 2, 2, 335, 332, 3, 2, 2, 2, 335, 333, 3, 2, 2, 2, 335, 334, 3, 2, 2,
	2, 336, 339, 3, 2, 2, 2, 337, 335, 3, 2, 2, 2, 337, 338, 3, 2, 2, 2, 338,
	66, 3, 2, 2, 2, 339, 337, 3, 2, 2, 2, 340, 343, 5, 43, 22, 2, 341, 343,
	5, 59, 30, 2, 342, 340, 3, 2, 2, 2, 342, 341, 3, 2, 2, 2, 343, 344, 3,
	2, 2, 2, 344, 345, 7, 39, 2, 2, 345, 68, 3, 2, 2, 2, 346, 348, 7, 15, 2,
	2, 347, 346, 3, 2, 2, 2, 347, 348, 3, 2, 2, 2, 348, 349, 3, 2, 2, 2, 349,
	350, 7, 12, 2, 2, 350, 351, 3, 2, 2, 2, 351, 352, 8, 35, 2, 2, 352, 70,
	3, 2, 2, 2, 353, 358, 9, 8, 2, 2, 354, 355, 7, 15, 2, 2, 355, 358, 7, 12,
	2, 2, 356, 358, 7, 5, 2, 2, 357, 353, 3, 2, 2, 2, 357, 354, 3, 2, 2, 2,
	357, 356, 3, 2, 2, 2, 358, 359, 3, 2, 2, 2, 359, 360, 8, 36, 2, 2, 360,
	72, 3, 2, 2, 2, 361, 362, 7, 48, 2, 2, 362, 74, 3, 2, 2, 2, 363, 364, 7,
	60, 2, 2, 364, 76, 3, 2, 2, 2, 365, 366, 7, 61, 2, 2, 366, 78, 3, 2, 2,
	2, 367, 368, 7, 46, 2, 2, 368, 80, 3, 2, 2, 2, 369, 370, 7, 63, 2, 2, 370,
	82, 3, 2, 2, 2, 371, 372, 7, 93, 2, 2, 372, 84, 3, 2, 2, 2, 373, 374, 7,
	95, 2, 2, 374, 86, 3, 2, 2, 2, 375, 376, 7, 125, 2, 2, 376, 88, 3, 2, 2,
	2, 377, 378, 7, 127, 2, 2, 378, 90, 3, 2, 2, 2, 379, 380, 7, 48, 2, 2,
	380, 381, 7, 48, 2, 2, 381, 92, 3, 2, 2, 2, 382, 383, 7, 42, 2, 2, 383,
	94, 3, 2, 2, 2, 384, 385, 7, 43, 2, 2, 385, 96, 3, 2, 2, 2, 386, 387, 7,
	128, 2, 2, 387, 388, 7, 63, 2, 2, 388, 98, 3, 2, 2, 2, 389, 390, 7, 62,
	2, 2, 390, 100, 3, 2, 2, 2, 391, 392, 7, 62, 2, 2, 392, 393, 7, 63, 2,
	2, 393, 102, 3, 2, 2, 2, 394, 395, 7, 64, 2, 2, 395, 104, 3, 2, 2, 2, 396,
	397, 7, 64, 2, 2, 397, 398, 7, 63, 2, 2, 398, 106, 3, 2, 2, 2, 399, 400,
	7, 45, 2, 2, 400, 108, 3, 2, 2, 2, 401, 402, 7, 47, 2, 2, 402, 110, 3,
	2, 2, 2, 403, 404, 7, 44, 2, 2, 404, 112, 3, 2, 2, 2, 405, 406, 7, 49,
	2, 2, 406, 114, 3, 2, 2, 2, 407, 408, 7, 96, 2, 2, 408, 116, 3, 2, 2, 2,
	27, 2, 215, 224, 226, 232, 234, 241, 246, 254, 265, 277, 283, 295, 299,
	303, 305, 313, 320, 325, 327, 335, 337, 342, 347, 357, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'::'", "'today'", "'true'", "'false'", "'nil'", "'default'", "'if'",
	"'then'", "'else'", "'case'", "'is'", "'endcase'", "'switch'", "'endswitch'",
	"'AND'", "'OR'", "'NOT'", "", "", "", "", "", "", "", "", "", "", "'.'",
	"':'", "';'", "','", "'='", "'['", "']'", "'{'", "'}'", "'..'", "'('",
	"')'", "'~='", "'<'", "'<='", "'>'", "'>='", "'+'", "'-'", "'*'", "'/'",
	"'^'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "IF", "THEN", "ELSE", "CASE", "CASE_IS", "CASE_END",
	"SWITCH", "SWITCH_END", "AND", "OR", "NOT", "INT", "ID", "STRING", "FLOAT",
	"DATE", "ATFUNCTION", "CLASSNAME", "PERCENT", "NEWLINE", "WS", "DOT", "COLON",
	"SEMI", "COMMA", "EQUALS", "LBRACKET", "RBRACKET", "CURLYLBRACKET", "CURLYRBRACKET",
	"DOTDOT", "LPAREN", "RPAREN", "NOT_EQUALS", "LT", "LTE", "GT", "GTE", "PLUS",
	"MINUS", "TIMES", "DIV", "EXPONENTIATE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "IF", "THEN", "ELSE", "CASE",
	"CASE_IS", "CASE_END", "SWITCH", "SWITCH_END", "AND", "OR", "NOT", "DIGIT",
	"LOWERCASELETTERS", "UPPERCASELETTERS", "INT", "ID", "STRING", "EXPONENT",
	"HEX_DIGIT", "ESC_SEQ", "OCTAL_ESC", "UNICODE_ESC", "FLOAT", "DATE", "ATFUNCTION",
	"CLASSNAME", "PERCENT", "NEWLINE", "WS", "DOT", "COLON", "SEMI", "COMMA",
	"EQUALS", "LBRACKET", "RBRACKET", "CURLYLBRACKET", "CURLYRBRACKET", "DOTDOT",
	"LPAREN", "RPAREN", "NOT_EQUALS", "LT", "LTE", "GT", "GTE", "PLUS", "MINUS",
	"TIMES", "DIV", "EXPONENTIATE",
}

type HilcoPencilParserLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewHilcoPencilParserLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *HilcoPencilParserLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewHilcoPencilParserLexer(input antlr.CharStream) *HilcoPencilParserLexer {
	l := new(HilcoPencilParserLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "HilcoPencilParser.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// HilcoPencilParserLexer tokens.
const (
	HilcoPencilParserLexerT__0          = 1
	HilcoPencilParserLexerT__1          = 2
	HilcoPencilParserLexerT__2          = 3
	HilcoPencilParserLexerT__3          = 4
	HilcoPencilParserLexerT__4          = 5
	HilcoPencilParserLexerT__5          = 6
	HilcoPencilParserLexerIF            = 7
	HilcoPencilParserLexerTHEN          = 8
	HilcoPencilParserLexerELSE          = 9
	HilcoPencilParserLexerCASE          = 10
	HilcoPencilParserLexerCASE_IS       = 11
	HilcoPencilParserLexerCASE_END      = 12
	HilcoPencilParserLexerSWITCH        = 13
	HilcoPencilParserLexerSWITCH_END    = 14
	HilcoPencilParserLexerAND           = 15
	HilcoPencilParserLexerOR            = 16
	HilcoPencilParserLexerNOT           = 17
	HilcoPencilParserLexerINT           = 18
	HilcoPencilParserLexerID            = 19
	HilcoPencilParserLexerSTRING        = 20
	HilcoPencilParserLexerFLOAT         = 21
	HilcoPencilParserLexerDATE          = 22
	HilcoPencilParserLexerATFUNCTION    = 23
	HilcoPencilParserLexerCLASSNAME     = 24
	HilcoPencilParserLexerPERCENT       = 25
	HilcoPencilParserLexerNEWLINE       = 26
	HilcoPencilParserLexerWS            = 27
	HilcoPencilParserLexerDOT           = 28
	HilcoPencilParserLexerCOLON         = 29
	HilcoPencilParserLexerSEMI          = 30
	HilcoPencilParserLexerCOMMA         = 31
	HilcoPencilParserLexerEQUALS        = 32
	HilcoPencilParserLexerLBRACKET      = 33
	HilcoPencilParserLexerRBRACKET      = 34
	HilcoPencilParserLexerCURLYLBRACKET = 35
	HilcoPencilParserLexerCURLYRBRACKET = 36
	HilcoPencilParserLexerDOTDOT        = 37
	HilcoPencilParserLexerLPAREN        = 38
	HilcoPencilParserLexerRPAREN        = 39
	HilcoPencilParserLexerNOT_EQUALS    = 40
	HilcoPencilParserLexerLT            = 41
	HilcoPencilParserLexerLTE           = 42
	HilcoPencilParserLexerGT            = 43
	HilcoPencilParserLexerGTE           = 44
	HilcoPencilParserLexerPLUS          = 45
	HilcoPencilParserLexerMINUS         = 46
	HilcoPencilParserLexerTIMES         = 47
	HilcoPencilParserLexerDIV           = 48
	HilcoPencilParserLexerEXPONENTIATE  = 49
)

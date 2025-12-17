package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(GoLang)
}

// GoLang defines syntax highlighting rules for Go.
var GoLang = &syntax.Language{
	Name:       "Go",
	Extensions: []string{".go"},
	Rules: []syntax.Rule{
		// Comments (must come first to take precedence)
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},

		// Strings
		{Type: syntax.TokenString, Pattern: regexp.MustCompile("`[^`]*`")},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)*'`)},

		// Keywords
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(break|case|chan|const|continue|default|defer|else|fallthrough|for|func|go|goto|if|import|interface|map|package|range|return|select|struct|switch|type|var)\b`)},

		// Types
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\b(bool|byte|complex64|complex128|error|float32|float64|int|int8|int16|int32|int64|rune|string|uint|uint8|uint16|uint32|uint64|uintptr)\b`)},

		// Built-in functions
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\b(append|cap|close|complex|copy|delete|imag|len|make|new|panic|print|println|real|recover)\b`)},

		// Constants
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|nil|iota)\b`)},

		// Numbers
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[oO][0-7_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[bB][01_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9][0-9_]*\.?[0-9_]*([eE][+-]?[0-9_]+)?\b`)},

		// Function definitions
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`\bfunc\s+(\w+)`)},

		// Operators
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!:]+`)},
	},
}

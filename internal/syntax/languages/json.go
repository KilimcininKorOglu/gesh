package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(JSONLang)
}

// JSONLang defines syntax highlighting rules for JSON.
var JSONLang = &syntax.Language{
	Name:       "JSON",
	Extensions: []string{".json", ".jsonc"},
	Rules: []syntax.Rule{
		// Strings (keys and values)
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},

		// Constants
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|null)\b`)},

		// Numbers
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`-?[0-9]+\.?[0-9]*([eE][+-]?[0-9]+)?`)},

		// Operators (colons, commas, brackets)
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[:\[\]{}]`)},
	},
}

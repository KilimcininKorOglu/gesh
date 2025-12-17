package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(YAMLLang)
}

// YAMLLang defines syntax highlighting rules for YAML.
var YAMLLang = &syntax.Language{
	Name:       "YAML",
	Extensions: []string{".yaml", ".yml"},
	Rules: []syntax.Rule{
		// Comments
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`#.*$`)},

		// Keys (word followed by colon)
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`^[\s]*[\w-]+:`)},

		// Strings
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)*'`)},

		// Constants
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|yes|no|on|off|null|~)\b`)},

		// Numbers
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\.?[0-9]*([eE][+-]?[0-9]+)?\b`)},

		// Anchors and aliases
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`[&*]\w+`)},

		// Tags
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`!\w+`)},
	},
}

package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(HTMLLang)
}

// HTMLLang defines syntax highlighting rules for HTML.
var HTMLLang = &syntax.Language{
	Name:       "HTML",
	Extensions: []string{".html", ".htm", ".xhtml"},
	Rules: []syntax.Rule{
		// Comments
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`<!--[\s\S]*?-->`)},

		// DOCTYPE
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`<!DOCTYPE[^>]*>`)},

		// Tags
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`</?\w+`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`/?>`)},

		// Attributes (word followed by equals)
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`\b[\w-]+=`)},

		// Strings (attribute values)
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"[^"]*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'[^']*'`)},

		// Entities
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`&\w+;`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`&#\d+;`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`&#x[0-9a-fA-F]+;`)},
	},
}

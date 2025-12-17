package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(MarkdownLang)
}

// MarkdownLang defines syntax highlighting rules for Markdown.
var MarkdownLang = &syntax.Language{
	Name:       "Markdown",
	Extensions: []string{".md", ".markdown", ".mkd"},
	Rules: []syntax.Rule{
		// Code blocks (inline)
		{Type: syntax.TokenString, Pattern: regexp.MustCompile("`[^`]+`")},

		// Headers
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`^#{1,6}\s.*$`)},

		// Bold
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\*\*[^*]+\*\*`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`__[^_]+__`)},

		// Italic
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\*[^*]+\*`)},
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`_[^_]+_`)},

		// Links
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`\[[^\]]+\]\([^)]+\)`)},

		// Images
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`!\[[^\]]*\]\([^)]+\)`)},

		// Block quotes
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`^>\s.*$`)},

		// List items
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`^[\s]*[-*+]\s`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`^[\s]*[0-9]+\.\s`)},

		// Horizontal rules
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`^[-*_]{3,}$`)},
	},
}

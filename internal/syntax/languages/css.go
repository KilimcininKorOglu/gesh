package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(CSSLang)
	syntax.RegisterLanguage(SCSSLang)
}

// CSSLang defines syntax highlighting rules for CSS.
var CSSLang = &syntax.Language{
	Name:       "CSS",
	Extensions: []string{".css"},
	Rules: []syntax.Rule{
		// Comments
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},

		// Strings
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"[^"]*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'[^']*'`)},

		// At-rules
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`@\w+`)},

		// Selectors (classes, IDs)
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`\.[a-zA-Z_][\w-]*`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`#[a-zA-Z_][\w-]*`)},

		// Pseudo-classes and pseudo-elements
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`::?[\w-]+`)},

		// Properties (word followed by colon)
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`[\w-]+\s*:`)},

		// Units and values
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`-?[0-9]+\.?[0-9]*(px|em|rem|%|vh|vw|vmin|vmax|ch|ex|cm|mm|in|pt|pc|deg|rad|turn|s|ms)?`)},

		// Colors (hex)
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`#[0-9a-fA-F]{3,8}`)},

		// Functions
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`\b(url|rgb|rgba|hsl|hsla|calc|var|attr|counter|linear-gradient|radial-gradient|repeating-linear-gradient|repeating-radial-gradient)\b`)},

		// Important
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`!important`)},
	},
}

// SCSSLang defines syntax highlighting rules for SCSS/Sass.
var SCSSLang = &syntax.Language{
	Name:       "SCSS",
	Extensions: []string{".scss", ".sass"},
	Rules: []syntax.Rule{
		// Comments
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},

		// Strings
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"[^"]*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'[^']*'`)},

		// Variables
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`\$[\w-]+`)},

		// At-rules and directives
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`@(import|include|extend|mixin|function|if|else|for|each|while|return|content|at-root|debug|warn|error)\b`)},

		// Selectors
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`\.[a-zA-Z_][\w-]*`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`#[a-zA-Z_][\w-]*`)},

		// Pseudo-classes
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`::?[\w-]+`)},

		// Numbers with units
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`-?[0-9]+\.?[0-9]*(px|em|rem|%|vh|vw|deg|s|ms)?`)},

		// Colors
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`#[0-9a-fA-F]{3,8}`)},

		// Important
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`!important`)},
	},
}

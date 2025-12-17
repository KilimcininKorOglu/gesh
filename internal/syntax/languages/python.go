package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(PythonLang)
}

// PythonLang defines syntax highlighting rules for Python.
var PythonLang = &syntax.Language{
	Name:       "Python",
	Extensions: []string{".py", ".pyw", ".pyi"},
	Rules: []syntax.Rule{
		// Comments
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`#.*$`)},

		// Strings (triple-quoted first)
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"""[\s\S]*?"""`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'''[\s\S]*?'''`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)*'`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`[fFrRbBuU]"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`[fFrRbBuU]'(?:[^'\\]|\\.)*'`)},

		// Keywords
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(and|as|assert|async|await|break|class|continue|def|del|elif|else|except|finally|for|from|global|if|import|in|is|lambda|nonlocal|not|or|pass|raise|return|try|while|with|yield)\b`)},

		// Built-in functions
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\b(abs|all|any|ascii|bin|bool|breakpoint|bytearray|bytes|callable|chr|classmethod|compile|complex|delattr|dict|dir|divmod|enumerate|eval|exec|filter|float|format|frozenset|getattr|globals|hasattr|hash|help|hex|id|input|int|isinstance|issubclass|iter|len|list|locals|map|max|memoryview|min|next|object|oct|open|ord|pow|print|property|range|repr|reversed|round|set|setattr|slice|sorted|staticmethod|str|sum|super|tuple|type|vars|zip)\b`)},

		// Constants
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(True|False|None|Ellipsis|NotImplemented)\b`)},

		// Decorators
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`@\w+`)},

		// Numbers
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[oO][0-7_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[bB][01_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9][0-9_]*\.?[0-9_]*([eE][+-]?[0-9_]+)?[jJ]?\b`)},

		// Operators
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!@:]+`)},
	},
}

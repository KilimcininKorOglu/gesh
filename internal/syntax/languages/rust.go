package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(RustLang)
}

// RustLang defines syntax highlighting rules for Rust.
var RustLang = &syntax.Language{
	Name:       "Rust",
	Extensions: []string{".rs"},
	Rules: []syntax.Rule{
		// Comments
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},

		// Strings
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)'`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`r#*"[^"]*"#*`)},

		// Keywords
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(as|async|await|break|const|continue|crate|dyn|else|enum|extern|false|fn|for|if|impl|in|let|loop|match|mod|move|mut|pub|ref|return|self|Self|static|struct|super|trait|true|type|unsafe|use|where|while)\b`)},

		// Types
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\b(bool|char|f32|f64|i8|i16|i32|i64|i128|isize|str|u8|u16|u32|u64|u128|usize|String|Vec|Option|Result|Box|Rc|Arc|Cell|RefCell|HashMap|HashSet|BTreeMap|BTreeSet)\b`)},

		// Macros
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\b\w+!`)},

		// Attributes
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`#\[[\w(,="'\s)]+\]`)},
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`#!\[[\w(,="'\s)]+\]`)},

		// Lifetimes
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`'\w+`)},

		// Numbers
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[oO][0-7_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[bB][01_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9][0-9_]*\.?[0-9_]*([eE][+-]?[0-9_]+)?(_?[iu](8|16|32|64|128|size)|_?f(32|64))?\b`)},

		// Operators
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!?:@]+`)},
	},
}

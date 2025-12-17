package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(CLang)
	syntax.RegisterLanguage(CppLang)
}

// CLang defines syntax highlighting rules for C.
var CLang = &syntax.Language{
	Name:       "C",
	Extensions: []string{".c", ".h"},
	Rules: []syntax.Rule{
		// Comments
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},

		// Preprocessor directives
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`^#\s*(include|define|undef|ifdef|ifndef|if|else|elif|endif|error|pragma|line)\b.*$`)},

		// Strings
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)*'`)},

		// Keywords
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(auto|break|case|char|const|continue|default|do|double|else|enum|extern|float|for|goto|if|inline|int|long|register|restrict|return|short|signed|sizeof|static|struct|switch|typedef|union|unsigned|void|volatile|while|_Alignas|_Alignof|_Atomic|_Bool|_Complex|_Generic|_Imaginary|_Noreturn|_Static_assert|_Thread_local)\b`)},

		// Types
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\b(size_t|ptrdiff_t|intptr_t|uintptr_t|int8_t|int16_t|int32_t|int64_t|uint8_t|uint16_t|uint32_t|uint64_t|FILE|NULL)\b`)},

		// Numbers
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F]+[uUlL]*\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[0-7]+[uUlL]*\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\.?[0-9]*([eE][+-]?[0-9]+)?[fFlL]?\b`)},

		// Operators
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!~?:]+`)},
	},
}

// CppLang defines syntax highlighting rules for C++.
var CppLang = &syntax.Language{
	Name:       "C++",
	Extensions: []string{".cpp", ".cc", ".cxx", ".hpp", ".hh", ".hxx"},
	Rules: []syntax.Rule{
		// Comments
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},

		// Preprocessor directives
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`^#\s*(include|define|undef|ifdef|ifndef|if|else|elif|endif|error|pragma|line)\b.*$`)},

		// Strings
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)*'`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`R"[^(]*\([^)]*\)[^"]*"`)},

		// Keywords (C + C++ specific)
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(alignas|alignof|and|and_eq|asm|auto|bitand|bitor|bool|break|case|catch|char|char8_t|char16_t|char32_t|class|compl|concept|const|consteval|constexpr|constinit|const_cast|continue|co_await|co_return|co_yield|decltype|default|delete|do|double|dynamic_cast|else|enum|explicit|export|extern|false|float|for|friend|goto|if|inline|int|long|mutable|namespace|new|noexcept|not|not_eq|nullptr|operator|or|or_eq|private|protected|public|register|reinterpret_cast|requires|return|short|signed|sizeof|static|static_assert|static_cast|struct|switch|template|this|thread_local|throw|true|try|typedef|typeid|typename|union|unsigned|using|virtual|void|volatile|wchar_t|while|xor|xor_eq)\b`)},

		// Standard library types
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\b(string|vector|map|set|list|deque|queue|stack|pair|tuple|array|unique_ptr|shared_ptr|weak_ptr|optional|variant|any|span|string_view)\b`)},

		// Numbers
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F']+[uUlL]*\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[bB][01']+[uUlL]*\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9][0-9']*\.?[0-9']*([eE][+-]?[0-9']+)?[fFlL]?\b`)},

		// Operators
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!~?:]+`)},
	},
}

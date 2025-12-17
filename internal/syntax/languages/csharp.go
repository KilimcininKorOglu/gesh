package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(CSharpLang)
	syntax.RegisterLanguage(FSharpLang)
}

// CSharpLang defines syntax highlighting rules for C#.
var CSharpLang = &syntax.Language{
	Name:       "C#",
	Extensions: []string{".cs", ".csx"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`@"[^"]*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`\$"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)'`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(abstract|as|base|break|case|catch|checked|class|const|continue|default|delegate|do|else|enum|event|explicit|extern|finally|fixed|for|foreach|goto|if|implicit|in|interface|internal|is|lock|namespace|new|operator|out|override|params|private|protected|public|readonly|record|ref|return|sealed|sizeof|stackalloc|static|struct|switch|this|throw|try|typeof|unchecked|unsafe|using|virtual|void|volatile|while|add|alias|ascending|async|await|by|descending|dynamic|equals|from|get|global|group|init|into|join|let|nameof|not|on|or|orderby|partial|remove|required|scoped|select|set|unmanaged|value|var|when|where|with|yield)\b`)},
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\b(bool|byte|char|decimal|double|float|int|long|object|sbyte|short|string|uint|ulong|ushort|nint|nuint|String|Int32|Int64|Boolean|Object|Exception|List|Dictionary|Array|Task|Action|Func|IEnumerable|IList|ICollection)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|null|this|base)\b`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\[\w+\]`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F_]+[uUlL]*\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[bB][01_]+[uUlL]*\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9][0-9_]*\.?[0-9_]*([eE][+-]?[0-9_]+)?[fFdDmMuUlL]*\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!~?:]+`)},
	},
}

// FSharpLang defines syntax highlighting rules for F#.
var FSharpLang = &syntax.Language{
	Name:       "F#",
	Extensions: []string{".fs", ".fsx", ".fsi"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`\(\*[\s\S]*?\*\)`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"""[\s\S]*?"""`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(abstract|and|as|assert|base|begin|class|default|delegate|do|done|downcast|downto|elif|else|end|exception|extern|finally|fixed|for|fun|function|global|if|in|inherit|inline|interface|internal|lazy|let|let!|match|match!|member|module|mutable|namespace|new|not|null|of|open|or|override|private|public|rec|return|return!|select|static|struct|then|to|try|type|upcast|use|use!|val|void|when|while|with|yield|yield!)\b`)},
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\b(bool|byte|char|decimal|double|float|float32|int|int16|int32|int64|nativeint|sbyte|single|string|uint|uint16|uint32|uint64|unativeint|unit|bigint|option|list|array|seq|async|Result|Option|List|Array|Seq|Map|Set)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|null|None|Some)\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F]+[uUlLnN]*\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[bB][01]+[uUlLnN]*\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\.?[0-9]*([eE][+-]?[0-9]+)?[fFmMuUlLnN]*\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!~?:@]+`)},
	},
}

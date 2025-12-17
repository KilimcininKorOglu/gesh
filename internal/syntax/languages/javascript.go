package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(JavaScriptLang)
	syntax.RegisterLanguage(TypeScriptLang)
}

// JavaScriptLang defines syntax highlighting rules for JavaScript.
var JavaScriptLang = &syntax.Language{
	Name:       "JavaScript",
	Extensions: []string{".js", ".jsx", ".mjs", ".cjs"},
	Rules: []syntax.Rule{
		// Comments
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},

		// Strings
		{Type: syntax.TokenString, Pattern: regexp.MustCompile("`[^`]*`")},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)*'`)},

		// Keywords
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(async|await|break|case|catch|class|const|continue|debugger|default|delete|do|else|export|extends|finally|for|from|function|if|import|in|instanceof|let|new|of|return|static|super|switch|this|throw|try|typeof|var|void|while|with|yield)\b`)},

		// Built-in objects
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\b(Array|Boolean|Date|Error|Function|JSON|Map|Math|Number|Object|Promise|Proxy|RegExp|Set|String|Symbol|WeakMap|WeakSet|console|document|window|globalThis|process|require|module|exports)\b`)},

		// Constants
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|null|undefined|NaN|Infinity)\b`)},

		// Numbers
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F_]+n?\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[oO][0-7_]+n?\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[bB][01_]+n?\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9][0-9_]*\.?[0-9_]*([eE][+-]?[0-9_]+)?n?\b`)},

		// Operators
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!?:]+`)},
	},
}

// TypeScriptLang defines syntax highlighting rules for TypeScript.
var TypeScriptLang = &syntax.Language{
	Name:       "TypeScript",
	Extensions: []string{".ts", ".tsx"},
	Rules: []syntax.Rule{
		// Comments
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},

		// Strings
		{Type: syntax.TokenString, Pattern: regexp.MustCompile("`[^`]*`")},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)*'`)},

		// Keywords (JS + TS specific)
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(abstract|any|as|async|await|boolean|break|case|catch|class|const|constructor|continue|debugger|declare|default|delete|do|else|enum|export|extends|finally|for|from|function|get|if|implements|import|in|infer|instanceof|interface|is|keyof|let|module|namespace|never|new|null|number|object|of|package|private|protected|public|readonly|return|set|static|string|super|switch|symbol|this|throw|try|type|typeof|undefined|unique|unknown|var|void|while|with|yield)\b`)},

		// Built-in objects
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\b(Array|Boolean|Date|Error|Function|JSON|Map|Math|Number|Object|Promise|Proxy|RegExp|Set|String|Symbol|WeakMap|WeakSet|console|document|window)\b`)},

		// Constants
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|null|undefined|NaN|Infinity)\b`)},

		// Numbers
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F_]+n?\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[oO][0-7_]+n?\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[bB][01_]+n?\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9][0-9_]*\.?[0-9_]*([eE][+-]?[0-9_]+)?n?\b`)},

		// Operators
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!?:]+`)},
	},
}

package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(PHPLang)
}

// PHPLang defines syntax highlighting rules for PHP.
var PHPLang = &syntax.Language{
	Name:       "PHP",
	Extensions: []string{".php", ".phtml", ".php3", ".php4", ".php5", ".php7", ".phps"},
	Rules: []syntax.Rule{
		// Comments
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`#.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},

		// PHP tags
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`<\?php`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`<\?=`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`<\?`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\?>`)},

		// Strings
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)*'`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile("<<<['\"]?\\w+['\"]?[\\s\\S]*?^\\w+;")},

		// Variables
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`\$\w+`)},

		// Keywords
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(abstract|and|array|as|break|callable|case|catch|class|clone|const|continue|declare|default|die|do|echo|else|elseif|empty|enddeclare|endfor|endforeach|endif|endswitch|endwhile|eval|exit|extends|final|finally|fn|for|foreach|function|global|goto|if|implements|include|include_once|instanceof|insteadof|interface|isset|list|match|namespace|new|or|print|private|protected|public|readonly|require|require_once|return|static|switch|throw|trait|try|unset|use|var|while|xor|yield|yield from)\b`)},

		// Types
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\b(int|float|bool|string|array|object|callable|iterable|void|mixed|never|null|false|true|self|parent|static)\b`)},

		// Built-in functions (common ones)
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\b(echo|print|isset|unset|empty|var_dump|print_r|die|exit|array|list|count|strlen|strpos|substr|str_replace|explode|implode|trim|strtolower|strtoupper|array_push|array_pop|array_shift|array_merge|array_keys|array_values|in_array|array_search|sort|rsort|asort|ksort|file_get_contents|file_put_contents|fopen|fclose|fread|fwrite|json_encode|json_decode|preg_match|preg_replace|date|time|strtotime)\b`)},

		// Constants
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(TRUE|FALSE|NULL|__CLASS__|__DIR__|__FILE__|__FUNCTION__|__LINE__|__METHOD__|__NAMESPACE__|__TRAIT__)\b`)},

		// Numbers
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[bB][01]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\.?[0-9]*([eE][+-]?[0-9]+)?\b`)},

		// Operators
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!.?:]+`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`=>`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`->`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`::`)},
	},
}

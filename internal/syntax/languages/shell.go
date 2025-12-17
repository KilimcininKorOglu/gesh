package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(ShellLang)
}

// ShellLang defines syntax highlighting rules for Shell/Bash.
var ShellLang = &syntax.Language{
	Name:       "Shell",
	Extensions: []string{".sh", ".bash", ".zsh", ".fish"},
	Rules: []syntax.Rule{
		// Comments
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`#.*$`)},

		// Strings
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'[^']*'`)},

		// Keywords
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(if|then|else|elif|fi|for|while|until|do|done|case|esac|in|function|select|time|coproc)\b`)},

		// Built-in commands
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\b(alias|bg|bind|break|builtin|caller|cd|command|compgen|complete|compopt|continue|declare|dirs|disown|echo|enable|eval|exec|exit|export|false|fc|fg|getopts|hash|help|history|jobs|kill|let|local|logout|mapfile|popd|printf|pushd|pwd|read|readarray|readonly|return|set|shift|shopt|source|suspend|test|times|trap|true|type|typeset|ulimit|umask|unalias|unset|wait)\b`)},

		// Variables
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`\$\{?[a-zA-Z_][a-zA-Z0-9_]*\}?`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`\$[0-9@#?$!*-]`)},

		// Numbers
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\b`)},

		// Operators
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[|&;<>()]+`)},
	},
}

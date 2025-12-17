package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(CoffeeScriptLang)
	syntax.RegisterLanguage(VueLang)
	syntax.RegisterLanguage(SvelteLang)
	syntax.RegisterLanguage(LessLang)
	syntax.RegisterLanguage(StylusLang)
}

// CoffeeScriptLang defines syntax highlighting rules for CoffeeScript.
var CoffeeScriptLang = &syntax.Language{
	Name:       "CoffeeScript",
	Extensions: []string{".coffee", ".cson", ".litcoffee"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`###[\s\S]*?###`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`#.*$`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"""[\s\S]*?"""`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'''[\s\S]*?'''`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)*'`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(and|break|by|catch|class|continue|debugger|delete|do|else|extends|finally|for|if|in|instanceof|is|isnt|loop|new|no|not|of|off|on|or|return|super|switch|then|this|throw|try|typeof|undefined|unless|until|when|while|yes)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|null|undefined)\b`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`@\w+`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\.?[0-9]*([eE][+-]?[0-9]+)?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%<>=!&|^~?:]+`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`->|=>`)},
	},
}

// VueLang defines syntax highlighting rules for Vue SFC.
var VueLang = &syntax.Language{
	Name:       "Vue",
	Extensions: []string{".vue"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`<!--[\s\S]*?-->`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`</?(template|script|style)\b[^>]*>`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`</?\w+`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`/?>`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`v-[\w:-]+`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`@[\w:-]+`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`:[\w:-]+`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`#[\w:-]+`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"[^"]*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'[^']*'`)},
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`\{\{[^}]*\}\}`)},
	},
}

// SvelteLang defines syntax highlighting rules for Svelte.
var SvelteLang = &syntax.Language{
	Name:       "Svelte",
	Extensions: []string{".svelte"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`<!--[\s\S]*?-->`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`</?(script|style)\b[^>]*>`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`</?\w+`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`/?>`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\{#(if|each|await|key)\b[^}]*\}`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\{:(else|then|catch)\b[^}]*\}`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\{/(if|each|await|key)\}`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`on:[\w]+`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`bind:[\w]+`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`class:[\w]+`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`use:[\w]+`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"[^"]*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'[^']*'`)},
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`\{[^}]*\}`)},
	},
}

// LessLang defines syntax highlighting rules for Less CSS.
var LessLang = &syntax.Language{
	Name:       "Less",
	Extensions: []string{".less"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"[^"]*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'[^']*'`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`@[\w-]+`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`&`)},
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`\.[a-zA-Z_][\w-]*`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`#[a-zA-Z_][\w-]*`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`::?[\w-]+`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`-?[0-9]+\.?[0-9]*(px|em|rem|%|vh|vw|deg|s|ms)?`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`#[0-9a-fA-F]{3,8}`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`!important`)},
	},
}

// StylusLang defines syntax highlighting rules for Stylus CSS.
var StylusLang = &syntax.Language{
	Name:       "Stylus",
	Extensions: []string{".styl", ".stylus"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"[^"]*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'[^']*'`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`\$[\w-]+`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`&`)},
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`\.[a-zA-Z_][\w-]*`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`#[a-zA-Z_][\w-]*`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`::?[\w-]+`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`-?[0-9]+\.?[0-9]*(px|em|rem|%|vh|vw|deg|s|ms)?`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`#[0-9a-fA-F]{3,8}`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`!important`)},
	},
}

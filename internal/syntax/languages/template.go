package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(JinjaLang)
	syntax.RegisterLanguage(HandleBarsLang)
	syntax.RegisterLanguage(EJSLang)
}

// JinjaLang defines syntax highlighting rules for Jinja2/Django templates.
var JinjaLang = &syntax.Language{
	Name:       "Jinja",
	Extensions: []string{".jinja", ".jinja2", ".j2", ".html.j2", ".django"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`\{#[\s\S]*?#\}`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\{%-?\s*(if|elif|else|endif|for|endfor|block|endblock|extends|include|import|from|macro|endmacro|call|endcall|filter|endfilter|set|endset|raw|endraw|autoescape|endautoescape|with|endwith|trans|endtrans|pluralize)\b[^%]*-?%\}`)},
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`\{\{-?[^}]*-?\}\}`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\{%-?[^%]*-?%\}`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`<!--[\s\S]*?-->`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`</?\w+`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`/?>`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"[^"]*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'[^']*'`)},
	},
}

// HandleBarsLang defines syntax highlighting rules for Handlebars/Mustache templates.
var HandleBarsLang = &syntax.Language{
	Name:       "Handlebars",
	Extensions: []string{".hbs", ".handlebars", ".mustache"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`\{\{!--[\s\S]*?--\}\}`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`\{\{![^}]*\}\}`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\{\{#[\w/][^}]*\}\}`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\{\{/[\w][^}]*\}\}`)},
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`\{\{[\^&>][^}]*\}\}`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`\{\{[^}]*\}\}`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`<!--[\s\S]*?-->`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`</?\w+`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`/?>`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"[^"]*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'[^']*'`)},
	},
}

// EJSLang defines syntax highlighting rules for EJS templates.
var EJSLang = &syntax.Language{
	Name:       "EJS",
	Extensions: []string{".ejs"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`<%#[\s\S]*?%>`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`<%-[\s\S]*?%>`)},
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`<%=[\s\S]*?%>`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`<%[\s\S]*?%>`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`<!--[\s\S]*?-->`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`</?\w+`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`/?>`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"[^"]*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'[^']*'`)},
	},
}

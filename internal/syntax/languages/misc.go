package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(LaTeXLang)
	syntax.RegisterLanguage(DiffLang)
	syntax.RegisterLanguage(RegexLang)
	syntax.RegisterLanguage(TerraformLang)
	syntax.RegisterLanguage(ProtobufLang)
}

// LaTeXLang defines syntax highlighting rules for LaTeX.
var LaTeXLang = &syntax.Language{
	Name:       "LaTeX",
	Extensions: []string{".tex", ".latex", ".ltx", ".sty", ".cls"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`%.*$`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\\[a-zA-Z@]+\*?`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`\{[^{}]*\}`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`\[[^\]]*\]`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\\begin\{[^}]+\}`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\\end\{[^}]+\}`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\$[^$]+\$`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\$\$[^$]+\$\$`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\.?[0-9]*(pt|em|ex|cm|mm|in|pc|bp|dd|cc|sp)?\b`)},
	},
}

// DiffLang defines syntax highlighting rules for Diff/Patch files.
var DiffLang = &syntax.Language{
	Name:       "Diff",
	Extensions: []string{".diff", ".patch"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`^@@.*@@.*$`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`^(diff|index|---|\+\+\+).*$`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`^\+.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`^-.*$`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`^!.*$`)},
	},
}

// RegexLang defines syntax highlighting rules for Regular Expressions.
var RegexLang = &syntax.Language{
	Name:       "Regex",
	Extensions: []string{".regex", ".regexp"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`[.^$]`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[*+?|]`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`\[[^\]]*\]`)},
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`\([^)]*\)`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\\[wWdDsSntrfvbB0-9]`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\{\d+(,\d*)?\}`)},
	},
}

// TerraformLang defines syntax highlighting rules for Terraform (HCL).
var TerraformLang = &syntax.Language{
	Name:       "Terraform",
	Extensions: []string{".tf", ".tfvars", ".hcl"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`#.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(resource|data|variable|output|locals|module|terraform|provider|provisioner|connection|lifecycle|dynamic|for_each|count|depends_on|source|version)\b`)},
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\b(string|number|bool|list|map|set|object|tuple|any)\b`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`<<-?EOF[\s\S]*?EOF`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`\$\{[^}]+\}`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`var\.\w+`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`local\.\w+`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\b(abs|ceil|floor|log|max|min|pow|signum|chomp|format|formatlist|indent|join|lower|upper|replace|split|substr|title|trim|trimprefix|trimsuffix|trimspace|can|try|tobool|tolist|tomap|tonumber|toset|tostring|concat|contains|distinct|element|flatten|index|keys|length|lookup|merge|range|reverse|setintersection|setproduct|setsubtract|setunion|slice|sort|sum|transpose|values|zipmap|base64decode|base64encode|base64gzip|csvdecode|jsondecode|jsonencode|urlencode|yamldecode|yamlencode|abspath|basename|dirname|pathexpand|file|fileexists|fileset|filebase64|templatefile|cidrhost|cidrnetmask|cidrsubnet|coalesce|coalescelist|compact|matchkeys|one|alltrue|anytrue|sensitive|nonsensitive|regex|regexall|uuid|uuidv5|bcrypt|md5|rsadecrypt|sha1|sha256|sha512|timestamp|timeadd|formatdate|parseint)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|null)\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\.?[0-9]*([eE][+-]?[0-9]+)?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%<>=!&|?:]+`)},
	},
}

// ProtobufLang defines syntax highlighting rules for Protocol Buffers.
var ProtobufLang = &syntax.Language{
	Name:       "Protobuf",
	Extensions: []string{".proto"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)*'`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(syntax|package|import|public|weak|option|message|enum|service|rpc|returns|stream|extend|extensions|reserved|to|max|oneof|map|repeated|optional|required|group)\b`)},
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\b(double|float|int32|int64|uint32|uint64|sint32|sint64|fixed32|fixed64|sfixed32|sfixed64|bool|string|bytes)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false)\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\.?[0-9]*([eE][+-]?[0-9]+)?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[=;{}\[\]<>]+`)},
	},
}

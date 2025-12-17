package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(ZigLang)
	syntax.RegisterLanguage(NimLang)
	syntax.RegisterLanguage(DLang)
	syntax.RegisterLanguage(AdaLang)
	syntax.RegisterLanguage(FortranLang)
}

// ZigLang defines syntax highlighting rules for Zig.
var ZigLang = &syntax.Language{
	Name:       "Zig",
	Extensions: []string{".zig"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)'`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(addrspace|align|allowzero|and|anyframe|anytype|asm|async|await|break|callconv|catch|comptime|const|continue|defer|else|enum|errdefer|error|export|extern|fn|for|if|inline|linksection|noalias|noinline|nosuspend|opaque|or|orelse|packed|pub|resume|return|struct|suspend|switch|test|threadlocal|try|union|unreachable|usingnamespace|var|volatile|while)\b`)},
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\b(bool|c_int|c_long|c_longlong|c_longdouble|c_short|c_uint|c_ulong|c_ulonglong|c_ushort|c_void|comptime_float|comptime_int|f16|f32|f64|f80|f128|i8|i16|i32|i64|i128|isize|noreturn|null|type|u8|u16|u32|u64|u128|undefined|usize|void)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|null|undefined)\b`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`@\w+`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[oO][0-7_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[bB][01_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9][0-9_]*\.?[0-9_]*([eE][+-]?[0-9_]+)?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!~?:]+`)},
	},
}

// NimLang defines syntax highlighting rules for Nim.
var NimLang = &syntax.Language{
	Name:       "Nim",
	Extensions: []string{".nim", ".nims", ".nimble"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`#.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`#\[[\s\S]*?\]#`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"""[\s\S]*?"""`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(addr|and|as|asm|bind|block|break|case|cast|concept|const|continue|converter|defer|discard|distinct|div|do|elif|else|end|enum|except|export|finally|for|from|func|if|import|in|include|interface|is|isnot|iterator|let|macro|method|mixin|mod|nil|not|notin|object|of|or|out|proc|ptr|raise|ref|return|shl|shr|static|template|try|tuple|type|using|var|when|while|xor|yield)\b`)},
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\b(int|int8|int16|int32|int64|uint|uint8|uint16|uint32|uint64|float|float32|float64|bool|char|string|cstring|pointer|typedesc|void|auto|any|untyped|typed|range|array|openArray|varargs|seq|set|tuple|object|ref|ptr|enum)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|nil)\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F_]+('?[iIuUfF](8|16|32|64))?\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[oO][0-7_]+('?[iIuUfF](8|16|32|64))?\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[bB][01_]+('?[iIuUfF](8|16|32|64))?\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9][0-9_]*\.?[0-9_]*([eE][+-]?[0-9_]+)?('?[iIuUfFdD](8|16|32|64))?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!~@$?:]+`)},
	},
}

// DLang defines syntax highlighting rules for D.
var DLang = &syntax.Language{
	Name:       "D",
	Extensions: []string{".d", ".di"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\+[\s\S]*?\+/`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile("`[^`]*`")},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`q"\{[\s\S]*?\}"`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(abstract|alias|align|asm|assert|auto|body|break|case|cast|catch|class|const|continue|debug|default|delegate|delete|deprecated|do|else|enum|export|extern|final|finally|for|foreach|foreach_reverse|function|goto|if|immutable|import|in|inout|interface|invariant|is|lazy|macro|mixin|module|new|nothrow|null|out|override|package|pragma|private|protected|public|pure|ref|return|scope|shared|static|struct|super|switch|synchronized|template|this|throw|try|typedef|typeid|typeof|union|unittest|version|void|while|with|__FILE__|__LINE__|__gshared|__traits|__vector|__parameters)\b`)},
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\b(bool|byte|cdouble|cent|cfloat|char|creal|dchar|double|float|idouble|ifloat|int|ireal|long|real|short|ubyte|ucent|uint|ulong|ushort|void|wchar|string|wstring|dstring|size_t|ptrdiff_t)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|null)\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F_]+[uUL]*\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[bB][01_]+[uUL]*\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9][0-9_]*\.?[0-9_]*([eE][+-]?[0-9_]+)?[fFLi]*\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!~?:@$]+`)},
	},
}

// AdaLang defines syntax highlighting rules for Ada.
var AdaLang = &syntax.Language{
	Name:       "Ada",
	Extensions: []string{".adb", ".ads", ".ada"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`--.*$`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'.'`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`(?i)\b(abort|abs|abstract|accept|access|aliased|all|and|array|at|begin|body|case|constant|declare|delay|delta|digits|do|else|elsif|end|entry|exception|exit|for|function|generic|goto|if|in|interface|is|limited|loop|mod|new|not|null|of|or|others|out|overriding|package|pragma|private|procedure|protected|raise|range|record|rem|renames|requeue|return|reverse|select|separate|some|subtype|synchronized|tagged|task|terminate|then|type|until|use|when|while|with|xor)\b`)},
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`(?i)\b(Boolean|Character|Integer|Natural|Positive|Float|Duration|String|Wide_String|Wide_Wide_String)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`(?i)\b(True|False)\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9][0-9_]*#[0-9a-fA-F_]+#\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9][0-9_]*\.?[0-9_]*([eE][+-]?[0-9_]+)?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%<>=&|:]+`)},
	},
}

// FortranLang defines syntax highlighting rules for Fortran.
var FortranLang = &syntax.Language{
	Name:       "Fortran",
	Extensions: []string{".f", ".for", ".f90", ".f95", ".f03", ".f08"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`!.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`(?i)^[c*].*$`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)*'`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`(?i)\b(allocatable|allocate|assign|associate|asynchronous|backspace|block|call|case|class|close|codimension|common|concurrent|contains|contiguous|continue|critical|cycle|data|deallocate|default|dimension|do|else|elseif|elsewhere|end|endfile|endif|entry|enum|enumerator|equivalence|error|exit|extends|external|final|flush|forall|format|function|generic|goto|if|images|implicit|import|include|inquire|intent|interface|intrinsic|lock|module|namelist|non_overridable|nopass|nullify|only|open|operator|optional|parameter|pass|pause|pointer|print|private|procedure|program|protected|public|pure|read|recursive|result|return|rewind|rewrite|save|select|sequence|stop|submodule|subroutine|sync|target|then|type|unlock|use|value|volatile|wait|where|while|write)\b`)},
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`(?i)\b(integer|real|double|precision|complex|character|logical|type)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`(?i)\b(\.true\.|\.false\.)\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`(?i)\b[0-9]+\.?[0-9]*([edq][+-]?[0-9]+)?(_\w+)?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`(?i)(\.[a-z]+\.)|[+\-*/%<>=:]+`)},
	},
}

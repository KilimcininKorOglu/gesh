package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(JavaLang)
	syntax.RegisterLanguage(KotlinLang)
	syntax.RegisterLanguage(ScalaLang)
	syntax.RegisterLanguage(GroovyLang)
}

// JavaLang defines syntax highlighting rules for Java.
var JavaLang = &syntax.Language{
	Name:       "Java",
	Extensions: []string{".java"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)'`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(abstract|assert|break|case|catch|class|const|continue|default|do|else|enum|extends|final|finally|for|goto|if|implements|import|instanceof|interface|native|new|package|private|protected|public|return|static|strictfp|super|switch|synchronized|this|throw|throws|transient|try|volatile|while|var|yield|record|sealed|permits|non-sealed)\b`)},
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\b(boolean|byte|char|double|float|int|long|short|void|String|Integer|Long|Double|Float|Boolean|Character|Object|Class|System|Exception|Error|Throwable|Thread|Runnable)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|null)\b`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`@\w+`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F_]+[lL]?\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[bB][01_]+[lL]?\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9][0-9_]*\.?[0-9_]*([eE][+-]?[0-9_]+)?[fFdDlL]?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!~?:]+`)},
	},
}

// KotlinLang defines syntax highlighting rules for Kotlin.
var KotlinLang = &syntax.Language{
	Name:       "Kotlin",
	Extensions: []string{".kt", ".kts"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"""[\s\S]*?"""`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)'`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(abstract|actual|annotation|as|break|by|catch|class|companion|const|constructor|continue|crossinline|data|do|else|enum|expect|external|final|finally|for|fun|get|if|import|in|infix|init|inline|inner|interface|internal|is|lateinit|noinline|object|open|operator|out|override|package|private|protected|public|reified|return|sealed|set|super|suspend|tailrec|this|throw|try|typealias|typeof|val|var|vararg|when|where|while)\b`)},
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\b(Boolean|Byte|Char|Double|Float|Int|Long|Short|String|Unit|Any|Nothing|Array|List|Map|Set|Pair|Triple)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|null)\b`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`@\w+`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F_]+[lL]?\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9][0-9_]*\.?[0-9_]*([eE][+-]?[0-9_]+)?[fFdDlL]?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!~?:]+`)},
	},
}

// ScalaLang defines syntax highlighting rules for Scala.
var ScalaLang = &syntax.Language{
	Name:       "Scala",
	Extensions: []string{".scala", ".sc"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"""[\s\S]*?"""`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(abstract|case|catch|class|def|do|else|extends|final|finally|for|forSome|if|implicit|import|lazy|match|new|object|override|package|private|protected|return|sealed|super|this|throw|trait|try|type|val|var|while|with|yield|given|using|enum|export|then)\b`)},
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\b(Boolean|Byte|Char|Double|Float|Int|Long|Short|String|Unit|Any|AnyRef|AnyVal|Nothing|Null|Option|Some|None|List|Map|Set|Seq|Vector)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|null)\b`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`@\w+`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F]+[lL]?\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\.?[0-9]*([eE][+-]?[0-9]+)?[fFdDlL]?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!~?:]+`)},
	},
}

// GroovyLang defines syntax highlighting rules for Groovy.
var GroovyLang = &syntax.Language{
	Name:       "Groovy",
	Extensions: []string{".groovy", ".gradle"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`/\*[\s\S]*?\*/`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"""[\s\S]*?"""`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'''[\s\S]*?'''`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)*'`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(abstract|as|assert|break|case|catch|class|const|continue|def|default|do|else|enum|extends|final|finally|for|goto|if|implements|import|in|instanceof|interface|native|new|package|private|protected|public|return|static|strictfp|super|switch|synchronized|this|throw|throws|trait|transient|try|volatile|while)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|null)\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\.?[0-9]*([eE][+-]?[0-9]+)?[gGlLfFdD]?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!~?:]+`)},
	},
}

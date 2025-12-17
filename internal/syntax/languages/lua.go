package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(LuaLang)
	syntax.RegisterLanguage(RLang)
}

// LuaLang defines syntax highlighting rules for Lua.
var LuaLang = &syntax.Language{
	Name:       "Lua",
	Extensions: []string{".lua"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`--\[\[[\s\S]*?\]\]`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`--.*$`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`\[\[[\s\S]*?\]\]`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)*'`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(and|break|do|else|elseif|end|for|function|goto|if|in|local|not|or|repeat|return|then|until|while)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|nil)\b`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\b(assert|collectgarbage|dofile|error|getmetatable|ipairs|load|loadfile|next|pairs|pcall|print|rawequal|rawget|rawlen|rawset|require|select|setmetatable|tonumber|tostring|type|xpcall|coroutine|debug|io|math|os|package|string|table|utf8)\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\.?[0-9]*([eE][+-]?[0-9]+)?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%^#<>=~]+`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`\.\.\.?`)},
	},
}

// RLang defines syntax highlighting rules for R.
var RLang = &syntax.Language{
	Name:       "R",
	Extensions: []string{".r", ".R", ".rmd", ".Rmd"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`#.*$`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)*'`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(if|else|repeat|while|function|for|in|next|break|TRUE|FALSE|NULL|Inf|NaN|NA|NA_integer_|NA_real_|NA_complex_|NA_character_)\b`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\b(abs|acos|acosh|all|any|anyNA|Arg|as|asin|asinh|atan|atan2|atanh|attr|attributes|baseenv|basename|besselI|besselJ|besselK|besselY|beta|bindingIsActive|bindingIsLocked|bitwAnd|bitwNot|bitwOr|bitwShiftL|bitwShiftR|bitwXor|body|bquote|browser|browserCondition|browserSetDebug|browserText|c|call|capabilities|casefold|cat|cbind|ceiling|char|chartr|charmatch|chol|choose|class|close|col|colMeans|colnames|colSums|commandArgs|comment|complex|conflicts|Conj|cos|cosh|cospi|crossprod|Cstack_info|cumall|cumany|cummax|cummin|cumprod|cumsum|cut|data|date|debug|default|delayedAssign|det|diag|diff|difftime|digamma|dim|dimnames|dir|dirname|do|double|drop|dump|duplicated|dyn|eigen|emptyenv|enc2native|enc2utf8|encodeString|Encoding|endsWith|enquote|env|environment|environmentIsLocked|environmentName|eval|evalq|exists|exp|expm1|expression|F|factor|factorial|fifo|file|Filter|Find|floor|flush|for|Force|formals|format|formatC|formatDL|forwardsolve|function|gamma|gc|get|getConnection|getElement|getExportedValue|getNativeSymbolInfo|getOption|getRversion|getSrcLines|getTaskCallbackNames|gettext|gettextf|getwd|gl|globalenv|gregexpr|grep|grepl|grepRaw|gsub|head|iconv|iconvlist|icuGetCollate|icuSetCollate|identical|identity|if|ifelse|Im|import|in|inherits|integer|interaction|interactive|invisible|is|isatty|isBaseNamespace|isdebugged|isIncomplete|isNamespace|isNamespaceLoaded|isOpen|isRestart|isS4|isSeekable|isSymmetric|isTRUE|jitter|julian|kappa|kronecker|l10n_info|labels|lapply|lazyLoad|lazyLoadDBexec|lazyLoadDBfetch|lbeta|lchoose|length|lengths|letters|LETTERS|levels|lfactorial|lgamma|library|licence|license|list|list2env|load|loadedNamespaces|loadingNamespaceInfo|loadNamespace|local|lockBinding|lockEnvironment|log|log10|log1p|log2|logb|logical|lower|ls|make|makeActiveBinding|Map|mapply|margin|match|Math|mat|matrix|max|mean|mem|memCompress|memDecompress|memory|merge|message|mget|min|missing|Mod|mode|month|names|nargs|nchar|ncol|NCOL|new|next|NextMethod|ngettext|nlevels|noquote|norm|normalizePath|nrow|NROW|numeric|numeric_version|nzchar|objects|oldClass|on|open|options|order|ordered|outer|package_version|packageEvent|packageHasNamespace|packageStartupMessage|pairlist|parent|parse|paste|paste0|path|pipe|pmatch|pmax|pmin|polyroot|pos|Position|pretty|print|prmatrix|proc|prod|prop|provideDimnames|psigamma|pushBack|pushBackLength|q|qr|quit|quote|R_system_version|range|rank|rapply|raw|rawConnection|rawConnectionValue|rawShift|rawToBits|rawToChar|rbind|rcond|Re|read|readBin|readChar|readline|readLines|readRDS|readRenviron|Recall|Reduce|reg|regexec|regexpr|regmatches|remove|rep|repeat|replace|replicate|require|requireNamespace|restartDescription|restartFormals|retracemem|return|rev|rle|rm|RNGkind|RNGversion|round|row|rowMeans|rownames|rowsum|rowSums|sample|sapply|save|saveRDS|scale|scan|search|searchpaths|seek|seq|seq_along|seq_len|sequence|serialize|set|setdiff|setequal|setHook|setNamespaceInfo|setSessionTimeLimit|setTimeLimit|setwd|showConnections|shQuote|sign|signalCondition|signif|simpleCondition|simpleError|simpleMessage|simpleWarning|sin|single|sinh|sinpi|sink|slice|socketConnection|socketSelect|solve|sort|source|split|sprintf|sqrt|sQuote|srcfile|srcfilealias|srcfilecopy|srcref|standardGeneric|startsWith|stderr|stdin|stdout|stop|stopifnot|storage|strftime|strptime|strsplit|strtoi|strtrim|structure|strwrap|sub|subset|substitute|substr|substring|sum|summary|suppressMessages|suppressPackageStartupMessages|suppressWarnings|svd|sweep|switch|sys|Sys|system|system2|T|t|table|tabulate|tail|tan|tanh|tanpi|tapply|taskCallbackManager|tcrossprod|tempdir|tempfile|testPlatformEquivalence|textConnection|textConnectionValue|tolower|topenv|toString|toupper|trace|traceback|tracemem|tracingState|transform|trigamma|trunc|truncate|try|tryCatch|type|typeof|unclass|undebug|union|unique|units|unix|unlink|unlist|unloadNamespace|unlockBinding|unname|unserialize|unsplit|untrace|untracemem|unz|upper|url|UseMethod|utf8ToInt|vapply|vector|Vectorize|warning|warnings|weekdays|which|while|with|withAutoprint|withCallingHandlers|within|withRestarts|withVisible|write|writeBin|writeChar|writeLines|xor|xpdrows|xtfrm|xzfile|zapsmall)\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\.?[0-9]*([eE][+-]?[0-9]+)?[iL]?\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F]+[L]?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%^<>=!&|~$@:]+`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`<-|->|<<-|->>`)},
	},
}

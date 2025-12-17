package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(RubyLang)
	syntax.RegisterLanguage(PerlLang)
}

// RubyLang defines syntax highlighting rules for Ruby.
var RubyLang = &syntax.Language{
	Name:       "Ruby",
	Extensions: []string{".rb", ".rake", ".gemspec", ".ru", ".erb"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`#.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`=begin[\s\S]*?=end`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)*'`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`%[qQwWiIxsr]?[\[\]{}()<>|/!].*?[\[\]{}()<>|/!]`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(BEGIN|END|alias|and|begin|break|case|class|def|defined\?|do|else|elsif|end|ensure|for|if|in|module|next|nil|not|or|redo|rescue|retry|return|self|super|then|undef|unless|until|when|while|yield|__FILE__|__LINE__|__ENCODING__)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|nil)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b[A-Z][A-Z0-9_]*\b`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`@{1,2}\w+`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`\$\w+`)},
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`:\w+`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[bB][01_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[oO]?[0-7_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9][0-9_]*\.?[0-9_]*([eE][+-]?[0-9_]+)?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!~?:]+`)},
	},
}

// PerlLang defines syntax highlighting rules for Perl.
var PerlLang = &syntax.Language{
	Name:       "Perl",
	Extensions: []string{".pl", ".pm", ".pod", ".t", ".psgi"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`#.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`=\w+[\s\S]*?=cut`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)*'`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`q[qwx]?[\[\]{}()<>|/!].*?[\[\]{}()<>|/!]`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(BEGIN|END|CHECK|INIT|UNITCHECK|AUTOLOAD|DESTROY|abs|accept|alarm|atan2|bind|binmode|bless|break|caller|chdir|chmod|chomp|chop|chown|chr|chroot|close|closedir|connect|continue|cos|crypt|dbmclose|dbmopen|default|defined|delete|die|do|dump|each|else|elsif|endgrent|endhostent|endnetent|endprotoent|endpwent|endservent|eof|eval|exec|exists|exit|exp|fcntl|fileno|flock|for|foreach|fork|format|formline|getc|getgrent|getgrgid|getgrnam|gethostbyaddr|gethostbyname|gethostent|getlogin|getnetbyaddr|getnetbyname|getnetent|getpeername|getpgrp|getppid|getpriority|getprotobyname|getprotobynumber|getprotoent|getpwent|getpwnam|getpwuid|getservbyname|getservbyport|getservent|getsockname|getsockopt|given|glob|gmtime|goto|grep|hex|if|import|index|int|ioctl|join|keys|kill|last|lc|lcfirst|length|link|listen|local|localtime|lock|log|lstat|map|mkdir|msgctl|msgget|msgrcv|msgsnd|my|new|next|no|oct|open|opendir|ord|our|pack|package|pipe|pop|pos|print|printf|prototype|push|quotemeta|rand|read|readdir|readline|readlink|readpipe|recv|redo|ref|rename|require|reset|return|reverse|rewinddir|rindex|rmdir|say|scalar|seek|seekdir|select|semctl|semget|semop|send|setgrent|sethostent|setnetent|setpgrp|setpriority|setprotoent|setpwent|setservent|setsockopt|shift|shmctl|shmget|shmread|shmwrite|shutdown|sin|sleep|socket|socketpair|sort|splice|split|sprintf|sqrt|srand|stat|state|study|sub|substr|symlink|syscall|sysopen|sysread|sysseek|system|syswrite|tell|telldir|tie|tied|time|times|truncate|uc|ucfirst|umask|undef|unless|unlink|unpack|unshift|untie|until|use|utime|values|vec|wait|waitpid|wantarray|warn|when|while|write)\b`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`[\$@%]\w+`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`\$[_&`+"`"+`'+./:;<=>?@\\^|~-]`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[bB][01_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9][0-9_]*\.?[0-9_]*([eE][+-]?[0-9_]+)?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%&|^<>=!~?:]+`)},
	},
}

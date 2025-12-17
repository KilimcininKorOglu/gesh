package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(TOMLLang)
	syntax.RegisterLanguage(INILang)
	syntax.RegisterLanguage(XMLLang)
	syntax.RegisterLanguage(DockerfileLang)
	syntax.RegisterLanguage(MakefileLang)
	syntax.RegisterLanguage(EnvLang)
	syntax.RegisterLanguage(NginxLang)
}

// TOMLLang defines syntax highlighting rules for TOML.
var TOMLLang = &syntax.Language{
	Name:       "TOML",
	Extensions: []string{".toml"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`#.*$`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`^\s*\[[^\]]+\]`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`^\s*[\w.-]+\s*=`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"""[\s\S]*?"""`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'''[\s\S]*?'''`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'[^']*'`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false)\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[+-]?[0-9][0-9_]*\.?[0-9_]*([eE][+-]?[0-9_]+)?\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[oO][0-7_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[bB][01_]+\b`)},
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\d{4}-\d{2}-\d{2}(T\d{2}:\d{2}:\d{2})?`)},
	},
}

// INILang defines syntax highlighting rules for INI files.
var INILang = &syntax.Language{
	Name:       "INI",
	Extensions: []string{".ini", ".cfg", ".conf", ".properties"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`[;#].*$`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`^\s*\[[^\]]+\]`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`^\s*[\w.-]+\s*=`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"[^"]*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'[^']*'`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\.?[0-9]*\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|yes|no|on|off)\b`)},
	},
}

// XMLLang defines syntax highlighting rules for XML.
var XMLLang = &syntax.Language{
	Name:       "XML",
	Extensions: []string{".xml", ".xsl", ".xslt", ".xsd", ".svg", ".plist", ".rss", ".atom"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`<!--[\s\S]*?-->`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`<\?[\s\S]*?\?>`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`<!DOCTYPE[^>]*>`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`<!\[CDATA\[[\s\S]*?\]\]>`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`</?\w+`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`/?>`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`\b[\w:-]+=`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"[^"]*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'[^']*'`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`&\w+;`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`&#\d+;`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`&#x[0-9a-fA-F]+;`)},
	},
}

// DockerfileLang defines syntax highlighting rules for Dockerfile.
var DockerfileLang = &syntax.Language{
	Name:       "Dockerfile",
	Extensions: []string{".dockerfile", "Dockerfile"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`#.*$`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`(?i)^(FROM|RUN|CMD|LABEL|MAINTAINER|EXPOSE|ENV|ADD|COPY|ENTRYPOINT|VOLUME|USER|WORKDIR|ARG|ONBUILD|STOPSIGNAL|HEALTHCHECK|SHELL)\b`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)*'`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`\$\{?\w+\}?`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\b`)},
	},
}

// MakefileLang defines syntax highlighting rules for Makefile.
var MakefileLang = &syntax.Language{
	Name:       "Makefile",
	Extensions: []string{".mk", "Makefile", "makefile", "GNUmakefile"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`#.*$`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`^\.\w+:`)},
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`^[\w.-]+\s*:`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`\$[\(\{][\w.-]+[\)\}]`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`\$[@<^?*%]`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`^\s*[\w.-]+\s*[:+?]?=`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`\b(ifeq|ifneq|ifdef|ifndef|else|endif|include|-include|sinclude|override|export|unexport|define|endef|undefine)\b`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)*'`)},
	},
}

// EnvLang defines syntax highlighting rules for .env files.
var EnvLang = &syntax.Language{
	Name:       "Env",
	Extensions: []string{".env", ".env.local", ".env.development", ".env.production", ".env.test"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`#.*$`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`^[\w]+\s*=`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'[^']*'`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false)\b`)},
	},
}

// NginxLang defines syntax highlighting rules for Nginx config.
var NginxLang = &syntax.Language{
	Name:       "Nginx",
	Extensions: []string{".nginx", "nginx.conf"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`#.*$`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(server|location|upstream|http|events|stream|mail|types|map|geo|split_clients|if|set|rewrite|return|break|include|root|index|alias|try_files|error_page|access_log|error_log|proxy_pass|fastcgi_pass|uwsgi_pass|scgi_pass|memcached_pass|listen|server_name|ssl_certificate|ssl_certificate_key|ssl_protocols|ssl_ciphers|gzip|gzip_types|add_header|expires|deny|allow|auth_basic|auth_basic_user_file|limit_req|limit_conn|worker_processes|worker_connections|use|multi_accept|sendfile|tcp_nopush|tcp_nodelay|keepalive_timeout|client_max_body_size|proxy_set_header|proxy_read_timeout|proxy_connect_timeout|proxy_buffer_size|proxy_buffers|fastcgi_param|uwsgi_param)\b`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'[^']*'`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`\$\w+`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+[kmgKMG]?\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(on|off)\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[~^=]+`)},
	},
}

package languages

import (
	"regexp"

	"github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
	syntax.RegisterLanguage(HaskellLang)
	syntax.RegisterLanguage(ElixirLang)
	syntax.RegisterLanguage(ErlangLang)
	syntax.RegisterLanguage(ClojureLang)
	syntax.RegisterLanguage(OCamlLang)
}

// HaskellLang defines syntax highlighting rules for Haskell.
var HaskellLang = &syntax.Language{
	Name:       "Haskell",
	Extensions: []string{".hs", ".lhs"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`--.*$`)},
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`\{-[\s\S]*?-\}`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)'`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(as|case|class|data|default|deriving|do|else|family|forall|foreign|hiding|if|import|in|infix|infixl|infixr|instance|let|mdo|module|newtype|of|proc|qualified|rec|then|type|where)\b`)},
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\b[A-Z][a-zA-Z0-9_']*\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(True|False|Nothing|Just|Left|Right|LT|EQ|GT)\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[oO][0-7]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\.?[0-9]*([eE][+-]?[0-9]+)?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%<>=!&|.@$^~:\\]+`)},
	},
}

// ElixirLang defines syntax highlighting rules for Elixir.
var ElixirLang = &syntax.Language{
	Name:       "Elixir",
	Extensions: []string{".ex", ".exs"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`#.*$`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"""[\s\S]*?"""`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)*'`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`~[a-zA-Z][\[\]{}()<>|/'"]+.*?[\[\]{}()<>|/'"]+[a-zA-Z]*`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(after|alias|and|case|catch|cond|def|defcallback|defdelegate|defexception|defguard|defguardp|defimpl|defmacro|defmacrop|defmodule|defoverridable|defp|defprotocol|defstruct|do|else|end|fn|for|if|import|in|not|or|quote|raise|receive|require|rescue|try|unless|unquote|unquote_splicing|use|when|with)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|nil)\b`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`@\w+`)},
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`:\w+`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[bB][01_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[oO][0-7_]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9][0-9_]*\.?[0-9_]*([eE][+-]?[0-9_]+)?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%<>=!&|^~@.:]+`)},
	},
}

// ErlangLang defines syntax highlighting rules for Erlang.
var ErlangLang = &syntax.Language{
	Name:       "Erlang",
	Extensions: []string{".erl", ".hrl"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`%.*$`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(after|and|andalso|band|begin|bnot|bor|bsl|bsr|bxor|case|catch|cond|div|end|fun|if|let|not|of|or|orelse|receive|rem|try|when|xor)\b`)},
		{Type: syntax.TokenBuiltin, Pattern: regexp.MustCompile(`-\w+`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|undefined)\b`)},
		{Type: syntax.TokenVariable, Pattern: regexp.MustCompile(`\b[A-Z_][a-zA-Z0-9_]*\b`)},
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`\b[a-z][a-zA-Z0-9_]*\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+#[0-9a-zA-Z]+\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\.?[0-9]*([eE][+-]?[0-9]+)?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%<>=!|:]+`)},
	},
}

// ClojureLang defines syntax highlighting rules for Clojure.
var ClojureLang = &syntax.Language{
	Name:       "Clojure",
	Extensions: []string{".clj", ".cljs", ".cljc", ".edn"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`;.*$`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(def|defn|defn-|defmacro|defmethod|defmulti|defonce|defprotocol|defrecord|defstruct|deftype|fn|if|if-let|if-not|if-some|when|when-let|when-not|when-some|when-first|cond|condp|case|do|doseq|dotimes|doto|for|let|letfn|loop|recur|throw|try|catch|finally|monitor-enter|monitor-exit|new|quote|var|set!|import|require|use|ns|in-ns|refer|refer-clojure|alias|comment|declare|binding|locking|proxy|reify|extend|extend-protocol|extend-type|gen-class|gen-interface)\b`)},
		{Type: syntax.TokenFunction, Pattern: regexp.MustCompile(`:\w[\w-]*`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|nil)\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F]+N?\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\.?[0-9]*([eE][+-]?[0-9]+)?[MN]?\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+/[0-9]+\b`)},
	},
}

// OCamlLang defines syntax highlighting rules for OCaml.
var OCamlLang = &syntax.Language{
	Name:       "OCaml",
	Extensions: []string{".ml", ".mli"},
	Rules: []syntax.Rule{
		{Type: syntax.TokenComment, Pattern: regexp.MustCompile(`\(\*[\s\S]*?\*\)`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)},
		{Type: syntax.TokenString, Pattern: regexp.MustCompile(`'(?:[^'\\]|\\.)'`)},
		{Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(and|as|assert|asr|begin|class|constraint|do|done|downto|else|end|exception|external|for|fun|function|functor|if|in|include|inherit|initializer|land|lazy|let|lor|lsl|lsr|lxor|match|method|mod|module|mutable|new|nonrec|object|of|open|or|private|rec|sig|struct|then|to|try|type|val|virtual|when|while|with)\b`)},
		{Type: syntax.TokenType_, Pattern: regexp.MustCompile(`\b(int|float|bool|char|string|unit|list|array|option|ref|exn|format|lazy_t|nativeint|int32|int64|bytes)\b`)},
		{Type: syntax.TokenConstant, Pattern: regexp.MustCompile(`\b(true|false|None|Some)\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[xX][0-9a-fA-F_]+[lLn]?\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[oO][0-7_]+[lLn]?\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b0[bB][01_]+[lLn]?\b`)},
		{Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9][0-9_]*\.?[0-9_]*([eE][+-]?[0-9_]+)?[lLn]?\b`)},
		{Type: syntax.TokenOperator, Pattern: regexp.MustCompile(`[+\-*/%<>=!&|@^~:;.#]+`)},
	},
}

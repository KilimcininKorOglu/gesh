// Package syntax provides syntax highlighting support.
package syntax

import (
	"regexp"
	"strings"
)

// TokenType represents the type of a syntax token.
type TokenType int

const (
	TokenNormal TokenType = iota
	TokenKeyword
	TokenType_
	TokenString
	TokenNumber
	TokenComment
	TokenOperator
	TokenFunction
	TokenVariable
	TokenConstant
	TokenBuiltin
)

// Token represents a highlighted token in a line.
type Token struct {
	Type  TokenType
	Start int
	End   int
	Text  string
}

// Rule represents a syntax highlighting rule.
type Rule struct {
	Type    TokenType
	Pattern *regexp.Regexp
}

// Language represents a programming language syntax definition.
type Language struct {
	Name       string
	Extensions []string
	Rules      []Rule
}

// Highlighter provides syntax highlighting for source code.
type Highlighter struct {
	language *Language
	enabled  bool

	// Token cache for performance
	cache      map[int][]Token // line number -> tokens
	cacheValid map[int]bool    // line number -> is valid
}

// New creates a new highlighter for the given language.
func New(lang *Language) *Highlighter {
	return &Highlighter{
		language:   lang,
		enabled:    true,
		cache:      make(map[int][]Token),
		cacheValid: make(map[int]bool),
	}
}

// SetEnabled enables or disables syntax highlighting.
func (h *Highlighter) SetEnabled(enabled bool) {
	h.enabled = enabled
}

// IsEnabled returns whether syntax highlighting is enabled.
func (h *Highlighter) IsEnabled() bool {
	return h.enabled
}

// Language returns the current language.
func (h *Highlighter) Language() *Language {
	return h.language
}

// SetLanguage sets the current language.
func (h *Highlighter) SetLanguage(lang *Language) {
	h.language = lang
	h.ClearCache() // Clear cache when language changes
}

// ClearCache clears the entire token cache.
func (h *Highlighter) ClearCache() {
	h.cache = make(map[int][]Token)
	h.cacheValid = make(map[int]bool)
}

// InvalidateLine marks a specific line as needing re-highlighting.
func (h *Highlighter) InvalidateLine(line int) {
	h.cacheValid[line] = false
}

// InvalidateLineRange marks a range of lines as needing re-highlighting.
func (h *Highlighter) InvalidateLineRange(startLine, endLine int) {
	for i := startLine; i <= endLine; i++ {
		h.cacheValid[i] = false
	}
}

// InvalidateFromLine marks all lines from startLine onwards as invalid.
func (h *Highlighter) InvalidateFromLine(startLine int) {
	for line := range h.cacheValid {
		if line >= startLine {
			h.cacheValid[line] = false
		}
	}
}

// Highlight tokenizes a line of source code.
func (h *Highlighter) Highlight(line string) []Token {
	if !h.enabled || h.language == nil || len(h.language.Rules) == 0 {
		return []Token{{Type: TokenNormal, Start: 0, End: len(line), Text: line}}
	}

	// Track which positions are already highlighted
	highlighted := make([]bool, len(line))
	tokens := []Token{}

	// Apply rules in order (first match wins for each position)
	for _, rule := range h.language.Rules {
		matches := rule.Pattern.FindAllStringIndex(line, -1)
		for _, match := range matches {
			start, end := match[0], match[1]

			// Check if any part of this match is already highlighted
			alreadyHighlighted := false
			for i := start; i < end; i++ {
				if highlighted[i] {
					alreadyHighlighted = true
					break
				}
			}

			if !alreadyHighlighted {
				// Mark positions as highlighted
				for i := start; i < end; i++ {
					highlighted[i] = i < len(highlighted)
					if i < len(highlighted) {
						highlighted[i] = true
					}
				}
				tokens = append(tokens, Token{
					Type:  rule.Type,
					Start: start,
					End:   end,
					Text:  line[start:end],
				})
			}
		}
	}

	// Sort tokens by start position
	sortTokens(tokens)

	// Fill gaps with normal tokens
	result := []Token{}
	pos := 0
	for _, t := range tokens {
		if t.Start > pos {
			result = append(result, Token{
				Type:  TokenNormal,
				Start: pos,
				End:   t.Start,
				Text:  line[pos:t.Start],
			})
		}
		result = append(result, t)
		pos = t.End
	}
	if pos < len(line) {
		result = append(result, Token{
			Type:  TokenNormal,
			Start: pos,
			End:   len(line),
			Text:  line[pos:],
		})
	}

	return result
}

// HighlightLine tokenizes a line with caching support.
// lineNum is used as the cache key.
func (h *Highlighter) HighlightLine(lineNum int, line string) []Token {
	// Check cache first
	if h.cacheValid[lineNum] {
		if cached, ok := h.cache[lineNum]; ok {
			return cached
		}
	}

	// Highlight and cache
	tokens := h.Highlight(line)
	h.cache[lineNum] = tokens
	h.cacheValid[lineNum] = true

	return tokens
}

// GetCachedTokens returns cached tokens if available.
func (h *Highlighter) GetCachedTokens(lineNum int) ([]Token, bool) {
	if !h.cacheValid[lineNum] {
		return nil, false
	}
	tokens, ok := h.cache[lineNum]
	return tokens, ok
}

// CacheStats returns cache statistics.
func (h *Highlighter) CacheStats() (cached, total int) {
	total = len(h.cache)
	for _, valid := range h.cacheValid {
		if valid {
			cached++
		}
	}
	return cached, total
}

// sortTokens sorts tokens by start position (simple insertion sort).
func sortTokens(tokens []Token) {
	for i := 1; i < len(tokens); i++ {
		j := i
		for j > 0 && tokens[j].Start < tokens[j-1].Start {
			tokens[j], tokens[j-1] = tokens[j-1], tokens[j]
			j--
		}
	}
}

// DetectLanguage detects the language from a filename.
func DetectLanguage(filename string) *Language {
	ext := ""
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			ext = strings.ToLower(filename[i:])
			break
		}
	}

	for _, lang := range Languages {
		for _, langExt := range lang.Extensions {
			if ext == langExt {
				return lang
			}
		}
	}

	return nil
}

// Languages is a map of all supported languages.
var Languages = map[string]*Language{}

// RegisterLanguage registers a language for syntax highlighting.
func RegisterLanguage(lang *Language) {
	Languages[lang.Name] = lang
	for _, ext := range lang.Extensions {
		Languages[ext] = lang
	}
}

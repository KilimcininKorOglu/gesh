// Package file provides file I/O operations for the editor.
package file

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

// LineEnding represents the type of line ending used in a file.
type LineEnding string

const (
	LineEndingLF   LineEnding = "LF"   // Unix: \n
	LineEndingCRLF LineEnding = "CRLF" // Windows: \r\n
	LineEndingCR   LineEnding = "CR"   // Old Mac: \r
)

// Encoding represents the character encoding of a file.
type Encoding string

const (
	EncodingUTF8    Encoding = "UTF-8"
	EncodingUTF8BOM Encoding = "UTF-8 BOM"
	EncodingLatin1  Encoding = "Latin-1"
	EncodingUnknown Encoding = "Unknown"
)

// FileInfo contains metadata about a loaded file.
type FileInfo struct {
	Content    string
	Encoding   Encoding
	LineEnding LineEnding
	HasBOM     bool
}

// SaveOptions contains options for saving files.
type SaveOptions struct {
	TrimTrailingSpaces bool
	FinalNewline       bool
	CreateBackup       bool
}

// DefaultSaveOptions returns default save options.
func DefaultSaveOptions() SaveOptions {
	return SaveOptions{
		TrimTrailingSpaces: false,
		FinalNewline:       false,
		CreateBackup:       false,
	}
}

// Save writes content to a file.
// Creates the file if it doesn't exist, overwrites if it does.
func Save(path string, content string) error {
	return SaveWithOptions(path, content, DefaultSaveOptions())
}

// SaveWithOptions writes content to a file with specified options.
func SaveWithOptions(path string, content string, opts SaveOptions) error {
	// Ensure directory exists
	dir := filepath.Dir(path)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	// Create backup if requested
	if opts.CreateBackup {
		if Exists(path) {
			backupPath := path + ".bak"
			if data, err := os.ReadFile(path); err == nil {
				_ = os.WriteFile(backupPath, data, 0644)
			}
		}
	}

	// Process content
	processedContent := content

	// Trim trailing whitespace from each line
	if opts.TrimTrailingSpaces {
		processedContent = trimTrailingWhitespace(processedContent)
	}

	// Ensure final newline
	if opts.FinalNewline {
		if len(processedContent) > 0 && !strings.HasSuffix(processedContent, "\n") {
			processedContent += "\n"
		}
	}

	return os.WriteFile(path, []byte(processedContent), 0644)
}

// trimTrailingWhitespace removes trailing spaces/tabs from each line.
func trimTrailingWhitespace(content string) string {
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " \t")
	}
	return strings.Join(lines, "\n")
}

// Load reads content from a file.
func Load(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// LoadWithInfo reads content from a file and returns metadata.
func LoadWithInfo(path string) (*FileInfo, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	info := &FileInfo{}

	// Detect encoding
	info.Encoding, info.HasBOM = detectEncoding(data)

	// Remove BOM if present
	content := data
	if info.HasBOM && len(data) >= 3 {
		content = data[3:] // Skip UTF-8 BOM
	}

	// Detect line ending
	info.LineEnding = detectLineEnding(content)

	// Normalize line endings to LF for internal use
	info.Content = normalizeLineEndings(string(content))

	return info, nil
}

// detectEncoding detects the character encoding of the data.
func detectEncoding(data []byte) (Encoding, bool) {
	// Check for UTF-8 BOM
	if len(data) >= 3 && data[0] == 0xEF && data[1] == 0xBB && data[2] == 0xBF {
		return EncodingUTF8BOM, true
	}

	// Check if valid UTF-8
	if utf8.Valid(data) {
		return EncodingUTF8, false
	}

	// Assume Latin-1 for non-UTF-8
	return EncodingLatin1, false
}

// detectLineEnding detects the predominant line ending in the content.
func detectLineEnding(data []byte) LineEnding {
	crlfCount := bytes.Count(data, []byte("\r\n"))
	crCount := bytes.Count(data, []byte("\r")) - crlfCount // CR not followed by LF
	lfCount := bytes.Count(data, []byte("\n")) - crlfCount // LF not preceded by CR

	// Return the most common line ending
	if crlfCount >= lfCount && crlfCount >= crCount {
		if crlfCount > 0 {
			return LineEndingCRLF
		}
	}
	if crCount >= lfCount {
		if crCount > 0 {
			return LineEndingCR
		}
	}

	// Default to LF (Unix style)
	return LineEndingLF
}

// normalizeLineEndings converts all line endings to LF.
func normalizeLineEndings(content string) string {
	// Replace CRLF first, then CR
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.ReplaceAll(content, "\r", "\n")
	return content
}

// ConvertLineEndings converts content to the specified line ending format.
func ConvertLineEndings(content string, lineEnding LineEnding) string {
	// First normalize to LF
	normalized := normalizeLineEndings(content)

	switch lineEnding {
	case LineEndingCRLF:
		return strings.ReplaceAll(normalized, "\n", "\r\n")
	case LineEndingCR:
		return strings.ReplaceAll(normalized, "\n", "\r")
	default:
		return normalized
	}
}

// Exists checks if a file exists.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Filename extracts the filename from a path.
func Filename(path string) string {
	return filepath.Base(path)
}

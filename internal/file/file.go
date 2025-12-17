// Package file provides file I/O operations for the editor.
package file

import (
	"os"
	"path/filepath"
	"strings"
)

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

// Exists checks if a file exists.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Filename extracts the filename from a path.
func Filename(path string) string {
	return filepath.Base(path)
}

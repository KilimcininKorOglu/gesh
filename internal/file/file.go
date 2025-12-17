// Package file provides file I/O operations for the editor.
package file

import (
	"os"
	"path/filepath"
)

// Save writes content to a file.
// Creates the file if it doesn't exist, overwrites if it does.
func Save(path string, content string) error {
	// Ensure directory exists
	dir := filepath.Dir(path)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	// Write file with newline at end if not present
	data := []byte(content)
	if len(data) > 0 && data[len(data)-1] != '\n' {
		data = append(data, '\n')
	}

	return os.WriteFile(path, data, 0644)
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

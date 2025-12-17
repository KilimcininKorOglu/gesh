package file

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveAndLoad(t *testing.T) {
	// Create temp directory
	tmpDir, err := os.MkdirTemp("", "gesh-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	testPath := filepath.Join(tmpDir, "test.txt")
	content := "Hello World\nLine 2"

	// Test Save (default: no final newline)
	err = Save(testPath, content)
	if err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	// Test Load
	loaded, err := Load(testPath)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	// Content should be unchanged (no auto newline by default)
	if loaded != content {
		t.Errorf("Load() = %q, want %q", loaded, content)
	}
}

func TestSaveWithOptions(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gesh-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	t.Run("FinalNewline", func(t *testing.T) {
		testPath := filepath.Join(tmpDir, "final_newline.txt")
		content := "No newline at end"

		opts := SaveOptions{FinalNewline: true}
		err = SaveWithOptions(testPath, content, opts)
		if err != nil {
			t.Fatalf("SaveWithOptions() error: %v", err)
		}

		loaded, err := Load(testPath)
		if err != nil {
			t.Fatalf("Load() error: %v", err)
		}

		expected := content + "\n"
		if loaded != expected {
			t.Errorf("Load() = %q, want %q", loaded, expected)
		}
	})

	t.Run("TrimTrailingSpaces", func(t *testing.T) {
		testPath := filepath.Join(tmpDir, "trim_spaces.txt")
		content := "Line with spaces   \nAnother line\t\t\n"

		opts := SaveOptions{TrimTrailingSpaces: true}
		err = SaveWithOptions(testPath, content, opts)
		if err != nil {
			t.Fatalf("SaveWithOptions() error: %v", err)
		}

		loaded, err := Load(testPath)
		if err != nil {
			t.Fatalf("Load() error: %v", err)
		}

		expected := "Line with spaces\nAnother line\n"
		if loaded != expected {
			t.Errorf("Load() = %q, want %q", loaded, expected)
		}
	})

	t.Run("CreateBackup", func(t *testing.T) {
		testPath := filepath.Join(tmpDir, "backup_test.txt")
		originalContent := "Original content"

		// Create original file
		err = Save(testPath, originalContent)
		if err != nil {
			t.Fatalf("Save() error: %v", err)
		}

		// Save with backup
		newContent := "New content"
		opts := SaveOptions{CreateBackup: true}
		err = SaveWithOptions(testPath, newContent, opts)
		if err != nil {
			t.Fatalf("SaveWithOptions() error: %v", err)
		}

		// Check backup exists
		backupPath := testPath + ".bak"
		if !Exists(backupPath) {
			t.Error("Backup file should exist")
		}

		// Check backup content
		backupContent, err := Load(backupPath)
		if err != nil {
			t.Fatalf("Load() backup error: %v", err)
		}
		if backupContent != originalContent {
			t.Errorf("Backup content = %q, want %q", backupContent, originalContent)
		}
	})

	t.Run("CombinedOptions", func(t *testing.T) {
		testPath := filepath.Join(tmpDir, "combined.txt")
		content := "Line with spaces   \nNo final newline"

		opts := SaveOptions{
			TrimTrailingSpaces: true,
			FinalNewline:       true,
		}
		err = SaveWithOptions(testPath, content, opts)
		if err != nil {
			t.Fatalf("SaveWithOptions() error: %v", err)
		}

		loaded, err := Load(testPath)
		if err != nil {
			t.Fatalf("Load() error: %v", err)
		}

		expected := "Line with spaces\nNo final newline\n"
		if loaded != expected {
			t.Errorf("Load() = %q, want %q", loaded, expected)
		}
	})
}

func TestSaveCreatesDirectory(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gesh-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Path with nested directory that doesn't exist
	testPath := filepath.Join(tmpDir, "subdir", "nested", "test.txt")

	err = Save(testPath, "content")
	if err != nil {
		t.Fatalf("Save() should create directories, got error: %v", err)
	}

	if !Exists(testPath) {
		t.Error("File should exist after Save()")
	}
}

func TestExists(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gesh-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	testPath := filepath.Join(tmpDir, "test.txt")

	// Should not exist yet
	if Exists(testPath) {
		t.Error("Exists() should return false for non-existent file")
	}

	// Create the file
	err = Save(testPath, "content")
	if err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	// Should exist now
	if !Exists(testPath) {
		t.Error("Exists() should return true after file is created")
	}
}

func TestFilename(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{"/path/to/file.txt", "file.txt"},
		{"file.txt", "file.txt"},
		{"/path/to/dir/", "dir"},
		{"", "."},
		{"C:\\Users\\test\\file.go", "file.go"},
	}

	for _, tt := range tests {
		got := Filename(tt.path)
		if got != tt.want {
			t.Errorf("Filename(%q) = %q, want %q", tt.path, got, tt.want)
		}
	}
}

func TestLoadNonExistent(t *testing.T) {
	_, err := Load("/nonexistent/path/file.txt")
	if err == nil {
		t.Error("Load() should return error for non-existent file")
	}
}

func TestLoadWithInfo(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gesh-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	t.Run("UTF-8 with LF", func(t *testing.T) {
		testPath := filepath.Join(tmpDir, "utf8_lf.txt")
		content := "Hello\nWorld\n"
		err := os.WriteFile(testPath, []byte(content), 0644)
		if err != nil {
			t.Fatalf("Failed to write test file: %v", err)
		}

		info, err := LoadWithInfo(testPath)
		if err != nil {
			t.Fatalf("LoadWithInfo() error: %v", err)
		}

		if info.Encoding != EncodingUTF8 {
			t.Errorf("Encoding = %q, want %q", info.Encoding, EncodingUTF8)
		}
		if info.LineEnding != LineEndingLF {
			t.Errorf("LineEnding = %q, want %q", info.LineEnding, LineEndingLF)
		}
		if info.HasBOM {
			t.Error("HasBOM should be false")
		}
	})

	t.Run("UTF-8 with CRLF", func(t *testing.T) {
		testPath := filepath.Join(tmpDir, "utf8_crlf.txt")
		content := "Hello\r\nWorld\r\n"
		err := os.WriteFile(testPath, []byte(content), 0644)
		if err != nil {
			t.Fatalf("Failed to write test file: %v", err)
		}

		info, err := LoadWithInfo(testPath)
		if err != nil {
			t.Fatalf("LoadWithInfo() error: %v", err)
		}

		if info.Encoding != EncodingUTF8 {
			t.Errorf("Encoding = %q, want %q", info.Encoding, EncodingUTF8)
		}
		if info.LineEnding != LineEndingCRLF {
			t.Errorf("LineEnding = %q, want %q", info.LineEnding, LineEndingCRLF)
		}
		// Content should be normalized to LF
		if info.Content != "Hello\nWorld\n" {
			t.Errorf("Content = %q, want normalized LF", info.Content)
		}
	})

	t.Run("UTF-8 with BOM", func(t *testing.T) {
		testPath := filepath.Join(tmpDir, "utf8_bom.txt")
		bom := []byte{0xEF, 0xBB, 0xBF}
		content := append(bom, []byte("Hello\nWorld")...)
		err := os.WriteFile(testPath, content, 0644)
		if err != nil {
			t.Fatalf("Failed to write test file: %v", err)
		}

		info, err := LoadWithInfo(testPath)
		if err != nil {
			t.Fatalf("LoadWithInfo() error: %v", err)
		}

		if info.Encoding != EncodingUTF8BOM {
			t.Errorf("Encoding = %q, want %q", info.Encoding, EncodingUTF8BOM)
		}
		if info.HasBOM != true {
			t.Error("HasBOM should be true")
		}
		// Content should not include BOM
		if info.Content != "Hello\nWorld" {
			t.Errorf("Content = %q, should not include BOM", info.Content)
		}
	})
}

func TestDetectLineEnding(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected LineEnding
	}{
		{"empty", "", LineEndingLF},
		{"no newlines", "hello world", LineEndingLF},
		{"LF only", "hello\nworld\n", LineEndingLF},
		{"CRLF only", "hello\r\nworld\r\n", LineEndingCRLF},
		{"mixed prefer CRLF", "hello\r\nworld\r\ntest\n", LineEndingCRLF},
		{"mixed prefer LF", "hello\nworld\ntest\r\n", LineEndingLF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := detectLineEnding([]byte(tt.content))
			if got != tt.expected {
				t.Errorf("detectLineEnding() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestConvertLineEndings(t *testing.T) {
	tests := []struct {
		name       string
		content    string
		lineEnding LineEnding
		expected   string
	}{
		{"LF to LF", "hello\nworld", LineEndingLF, "hello\nworld"},
		{"LF to CRLF", "hello\nworld", LineEndingCRLF, "hello\r\nworld"},
		{"CRLF to LF", "hello\r\nworld", LineEndingLF, "hello\nworld"},
		{"CRLF to CRLF", "hello\r\nworld", LineEndingCRLF, "hello\r\nworld"},
		{"mixed to LF", "hello\r\nworld\n", LineEndingLF, "hello\nworld\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertLineEndings(tt.content, tt.lineEnding)
			if got != tt.expected {
				t.Errorf("ConvertLineEndings() = %q, want %q", got, tt.expected)
			}
		})
	}
}

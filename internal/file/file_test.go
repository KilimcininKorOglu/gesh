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

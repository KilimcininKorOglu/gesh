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

	// Test Save
	err = Save(testPath, content)
	if err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	// Test Load
	loaded, err := Load(testPath)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	// Content should have newline appended
	expected := content + "\n"
	if loaded != expected {
		t.Errorf("Load() = %q, want %q", loaded, expected)
	}
}

func TestSaveWithNewline(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gesh-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	testPath := filepath.Join(tmpDir, "test.txt")
	content := "Already has newline\n"

	err = Save(testPath, content)
	if err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	loaded, err := Load(testPath)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	// Should not double the newline
	if loaded != content {
		t.Errorf("Load() = %q, want %q", loaded, content)
	}
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

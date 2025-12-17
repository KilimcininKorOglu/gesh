package app

import (
	"testing"
)

func TestNew(t *testing.T) {
	m := New()

	if m == nil {
		t.Fatal("New() returned nil")
	}

	if m.buffer == nil {
		t.Error("buffer is nil")
	}

	if m.filename != "[New File]" {
		t.Errorf("filename = %q, want %q", m.filename, "[New File]")
	}

	if m.mode != ModeNormal {
		t.Errorf("mode = %d, want ModeNormal", m.mode)
	}

	if m.modified {
		t.Error("new model should not be modified")
	}
}

func TestNewWithContent(t *testing.T) {
	content := "Hello World"
	m := NewWithContent(content)

	if m.buffer.String() != content {
		t.Errorf("buffer content = %q, want %q", m.buffer.String(), content)
	}
}

func TestNewFromFile(t *testing.T) {
	m := NewFromFile("/path/to/file.txt", "file.txt", "content")

	if m.filepath != "/path/to/file.txt" {
		t.Errorf("filepath = %q, want %q", m.filepath, "/path/to/file.txt")
	}

	if m.filename != "file.txt" {
		t.Errorf("filename = %q, want %q", m.filename, "file.txt")
	}

	if m.buffer.String() != "content" {
		t.Errorf("buffer content = %q, want %q", m.buffer.String(), "content")
	}
}

func TestBuffer(t *testing.T) {
	m := NewWithContent("test")

	buf := m.Buffer()
	if buf == nil {
		t.Error("Buffer() returned nil")
	}

	if buf.String() != "test" {
		t.Errorf("Buffer().String() = %q, want %q", buf.String(), "test")
	}
}

func TestFilename(t *testing.T) {
	m := NewFromFile("/path/to/test.go", "test.go", "")

	if m.Filename() != "test.go" {
		t.Errorf("Filename() = %q, want %q", m.Filename(), "test.go")
	}
}

func TestModified(t *testing.T) {
	m := New()

	if m.IsModified() {
		t.Error("IsModified() should be false initially")
	}

	m.SetModified(true)
	if !m.IsModified() {
		t.Error("IsModified() should be true after SetModified(true)")
	}

	m.SetModified(false)
	if m.IsModified() {
		t.Error("IsModified() should be false after SetModified(false)")
	}
}

func TestMode(t *testing.T) {
	m := New()

	if m.Mode() != ModeNormal {
		t.Errorf("Mode() = %d, want ModeNormal", m.Mode())
	}

	m.SetMode(ModeSearch)
	if m.Mode() != ModeSearch {
		t.Errorf("Mode() = %d, want ModeSearch", m.Mode())
	}

	m.SetMode(ModeGoto)
	if m.Mode() != ModeGoto {
		t.Errorf("Mode() = %d, want ModeGoto", m.Mode())
	}
}

func TestStatusMessage(t *testing.T) {
	m := New()

	if m.StatusMessage() != "" {
		t.Errorf("StatusMessage() = %q, want empty", m.StatusMessage())
	}

	m.SetStatusMessage("File saved")
	if m.StatusMessage() != "File saved" {
		t.Errorf("StatusMessage() = %q, want %q", m.StatusMessage(), "File saved")
	}
}

func TestSize(t *testing.T) {
	m := New()

	if m.Width() != 0 || m.Height() != 0 {
		t.Errorf("Initial size = (%d, %d), want (0, 0)", m.Width(), m.Height())
	}

	m.SetSize(80, 24)

	if m.Width() != 80 {
		t.Errorf("Width() = %d, want 80", m.Width())
	}

	if m.Height() != 24 {
		t.Errorf("Height() = %d, want 24", m.Height())
	}
}

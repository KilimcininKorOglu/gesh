// Package app provides the main application logic using Bubble Tea.
package app

import (
	"github.com/KilimcininKorOglu/gesh/internal/buffer"
)

// Mode represents the current editor mode.
type Mode int

const (
	// ModeNormal is the default editing mode.
	ModeNormal Mode = iota
	// ModeSearch is the search mode.
	ModeSearch
	// ModeGoto is the "go to line" mode.
	ModeGoto
	// ModeQuit is the quit confirmation mode.
	ModeQuit
	// ModeSaveAs is the "save as" mode for entering filename.
	ModeSaveAs
	// ModeReplace is the "find and replace" mode.
	ModeReplace
	// ModeReplaceConfirm is the replace confirmation mode.
	ModeReplaceConfirm
	// ModeOpen is the "open file" mode.
	ModeOpen
)

// Model is the main Bubble Tea model for the editor.
type Model struct {
	// Buffer holds the text content.
	buffer  *buffer.GapBuffer
	history *buffer.History

	// File information
	filename string
	filepath string
	modified bool
	readonly bool

	// Terminal dimensions
	width  int
	height int

	// Viewport (visible area)
	viewportTopLine    int
	viewportLeftColumn int

	// Editor mode
	mode Mode

	// Input buffer for prompts (search, goto, etc.)
	inputBuffer string
	inputPrompt string

	// Search state
	searchQuery   string
	searchMatches []int // positions of matches
	searchIndex   int   // current match index

	// Selection state
	selecting      bool
	selectionStart int
	selectionEnd   int

	// Clipboard
	clipboard string

	// Replace state
	replaceText string

	// Status message
	statusMessage string

	// Quit flag
	quitting bool
}

// New creates a new editor model with an empty buffer.
func New() *Model {
	return &Model{
		buffer:   buffer.New(),
		history:  buffer.NewHistory(),
		filename: "[New File]",
		mode:     ModeNormal,
	}
}

// NewWithContent creates a new editor model with initial content.
func NewWithContent(content string) *Model {
	return &Model{
		buffer:   buffer.NewFromString(content),
		history:  buffer.NewHistory(),
		filename: "[New File]",
		mode:     ModeNormal,
	}
}

// NewFromFile creates a new editor model for a specific file.
func NewFromFile(filepath, filename, content string) *Model {
	return &Model{
		buffer:   buffer.NewFromString(content),
		history:  buffer.NewHistory(),
		filename: filename,
		filepath: filepath,
		mode:     ModeNormal,
	}
}

// Buffer returns the underlying gap buffer.
func (m *Model) Buffer() *buffer.GapBuffer {
	return m.buffer
}

// Filename returns the current filename.
func (m *Model) Filename() string {
	return m.filename
}

// IsModified returns whether the buffer has been modified.
func (m *Model) IsModified() bool {
	return m.modified
}

// SetModified sets the modified flag.
func (m *Model) SetModified(modified bool) {
	m.modified = modified
}

// Mode returns the current editor mode.
func (m *Model) Mode() Mode {
	return m.mode
}

// SetMode changes the editor mode.
func (m *Model) SetMode(mode Mode) {
	m.mode = mode
}

// SetStatusMessage sets a status message to display.
func (m *Model) SetStatusMessage(msg string) {
	m.statusMessage = msg
}

// StatusMessage returns the current status message.
func (m *Model) StatusMessage() string {
	return m.statusMessage
}

// Width returns the terminal width.
func (m *Model) Width() int {
	return m.width
}

// Height returns the terminal height.
func (m *Model) Height() int {
	return m.height
}

// SetSize sets the terminal dimensions.
func (m *Model) SetSize(width, height int) {
	m.width = width
	m.height = height
}

// Filepath returns the current file path.
func (m *Model) Filepath() string {
	return m.filepath
}

// SetFilepath sets the file path and updates filename.
func (m *Model) SetFilepath(path string) {
	m.filepath = path
	if path != "" {
		// Extract filename from path
		for i := len(path) - 1; i >= 0; i-- {
			if path[i] == '/' || path[i] == '\\' {
				m.filename = path[i+1:]
				return
			}
		}
		m.filename = path
	}
}

// Content returns the buffer content as string.
func (m *Model) Content() string {
	return m.buffer.String()
}

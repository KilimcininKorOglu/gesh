// Package app provides tab/buffer management for multiple file editing.
package app

import (
	"github.com/KilimcininKorOglu/gesh/internal/buffer"
)

// Tab represents a single buffer/file in the editor.
type Tab struct {
	// Buffer holds the text content.
	buffer  *buffer.GapBuffer
	history *buffer.History

	// File information
	filename   string
	filepath   string
	modified   bool
	readonly   bool
	encoding   string
	lineEnding string

	// Viewport state (preserved when switching tabs)
	viewportTopLine    int
	viewportLeftColumn int

	// Cursor position (preserved when switching tabs)
	cursorPos int

	// Selection state
	selecting      bool
	selectionStart int
	selectionEnd   int

	// Search state
	searchQuery   string
	searchMatches []int
	searchIndex   int
}

// TabManager manages multiple tabs/buffers.
type TabManager struct {
	tabs        []*Tab
	activeIndex int
}

// NewTabManager creates a new tab manager with an empty tab.
func NewTabManager() *TabManager {
	return &TabManager{
		tabs:        []*Tab{newEmptyTab()},
		activeIndex: 0,
	}
}

// newEmptyTab creates a new empty tab.
func newEmptyTab() *Tab {
	return &Tab{
		buffer:     buffer.New(),
		history:    buffer.NewHistory(),
		filename:   "[New File]",
		encoding:   "UTF-8",
		lineEnding: "LF",
	}
}

// NewTabFromFile creates a tab from a file.
func NewTabFromFile(filepath, filename, content, encoding, lineEnding string) *Tab {
	return &Tab{
		buffer:     buffer.NewFromString(content),
		history:    buffer.NewHistory(),
		filename:   filename,
		filepath:   filepath,
		encoding:   encoding,
		lineEnding: lineEnding,
	}
}

// ActiveTab returns the currently active tab.
func (tm *TabManager) ActiveTab() *Tab {
	if tm.activeIndex >= 0 && tm.activeIndex < len(tm.tabs) {
		return tm.tabs[tm.activeIndex]
	}
	return nil
}

// ActiveIndex returns the index of the active tab.
func (tm *TabManager) ActiveIndex() int {
	return tm.activeIndex
}

// Count returns the number of tabs.
func (tm *TabManager) Count() int {
	return len(tm.tabs)
}

// Tabs returns all tabs.
func (tm *TabManager) Tabs() []*Tab {
	return tm.tabs
}

// AddTab adds a new tab and makes it active.
func (tm *TabManager) AddTab(tab *Tab) {
	tm.tabs = append(tm.tabs, tab)
	tm.activeIndex = len(tm.tabs) - 1
}

// AddEmptyTab adds a new empty tab and makes it active.
func (tm *TabManager) AddEmptyTab() {
	tm.AddTab(newEmptyTab())
}

// CloseActiveTab closes the currently active tab.
// Returns false if it's the last tab and cannot be closed.
func (tm *TabManager) CloseActiveTab() bool {
	if len(tm.tabs) <= 1 {
		return false
	}

	// Remove the active tab
	tm.tabs = append(tm.tabs[:tm.activeIndex], tm.tabs[tm.activeIndex+1:]...)

	// Adjust active index
	if tm.activeIndex >= len(tm.tabs) {
		tm.activeIndex = len(tm.tabs) - 1
	}

	return true
}

// NextTab switches to the next tab.
func (tm *TabManager) NextTab() {
	if len(tm.tabs) > 1 {
		tm.activeIndex = (tm.activeIndex + 1) % len(tm.tabs)
	}
}

// PrevTab switches to the previous tab.
func (tm *TabManager) PrevTab() {
	if len(tm.tabs) > 1 {
		tm.activeIndex = (tm.activeIndex - 1 + len(tm.tabs)) % len(tm.tabs)
	}
}

// SelectTab switches to a specific tab by index.
func (tm *TabManager) SelectTab(index int) bool {
	if index >= 0 && index < len(tm.tabs) {
		tm.activeIndex = index
		return true
	}
	return false
}

// HasUnsavedChanges returns true if any tab has unsaved changes.
func (tm *TabManager) HasUnsavedChanges() bool {
	for _, tab := range tm.tabs {
		if tab.modified {
			return true
		}
	}
	return false
}

// TabWithUnsavedChanges returns the first tab with unsaved changes, or nil.
func (tm *TabManager) TabWithUnsavedChanges() *Tab {
	for _, tab := range tm.tabs {
		if tab.modified {
			return tab
		}
	}
	return nil
}

// SaveCursorPosition saves the current cursor position to the active tab.
func (tm *TabManager) SaveCursorPosition(pos int) {
	if tab := tm.ActiveTab(); tab != nil {
		tab.cursorPos = pos
	}
}

// SaveViewport saves the current viewport to the active tab.
func (tm *TabManager) SaveViewport(topLine, leftCol int) {
	if tab := tm.ActiveTab(); tab != nil {
		tab.viewportTopLine = topLine
		tab.viewportLeftColumn = leftCol
	}
}

// SaveSelection saves the current selection state to the active tab.
func (tm *TabManager) SaveSelection(selecting bool, start, end int) {
	if tab := tm.ActiveTab(); tab != nil {
		tab.selecting = selecting
		tab.selectionStart = start
		tab.selectionEnd = end
	}
}

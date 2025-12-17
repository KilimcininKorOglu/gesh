// Package app provides the main application logic using Bubble Tea.
package app

import (
	"time"

	"github.com/KilimcininKorOglu/gesh/internal/buffer"
	"github.com/KilimcininKorOglu/gesh/internal/syntax"
	"github.com/KilimcininKorOglu/gesh/internal/ui/styles"
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
	// ModeReplace is the "find and replace one" mode.
	ModeReplace
	// ModeReplaceConfirm is the replace confirmation mode (for single replace).
	ModeReplaceConfirm
	// ModeReplaceAll is the "replace all" mode.
	ModeReplaceAll
	// ModeReplaceAllConfirm is the replace all confirmation mode.
	ModeReplaceAllConfirm
	// ModeOpen is the "open file" mode.
	ModeOpen
	// ModeSaveMacro is the "save macro" mode.
	ModeSaveMacro
	// ModeLoadMacro is the "load macro" mode.
	ModeLoadMacro
)

// Model is the main Bubble Tea model for the editor.
type Model struct {
	// Tab management (multi-buffer support)
	tabs *TabManager

	// Split view management
	split *SplitManager

	// Buffer holds the text content (shortcut to active tab's buffer)
	buffer  *buffer.GapBuffer
	history *buffer.History

	// File information (shortcut to active tab)
	filename   string
	filepath   string
	modified   bool
	readonly   bool
	encoding   string
	lineEnding string

	// Display options
	showLineNumbers    bool
	wordWrap           bool
	syntaxHighlighting bool
	showTabs           bool // show tab bar

	// Edit mode
	overwriteMode bool // false = insert, true = overwrite

	// Save options
	trimTrailingSpaces bool
	finalNewline       bool
	createBackup       bool

	// Terminal dimensions
	width  int
	height int

	// Viewport (visible area)
	viewportTopLine    int
	viewportLeftColumn int

	// Smooth scroll state
	targetTopLine   int  // target line for smooth scroll
	scrollAnimating bool // whether scroll animation is in progress

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

	// Double Ctrl+A detection
	lastCtrlATime int64

	// Clipboard
	clipboard string

	// Replace state
	replaceText string

	// Macro recorder
	macro *MacroRecorder

	// Auto-save
	autoSaveInterval int // seconds, 0 = disabled
	lastSaveTime     int64

	// File watcher
	fileChanged bool // external change detected

	// Render cache for incremental rendering
	cachedLines       map[int]string // line number -> rendered content
	lastRenderVersion int            // buffer version at last render
	dirtyLines        map[int]bool   // lines that need re-render

	// Syntax highlighter (cached per model)
	highlighter *syntax.Highlighter

	// Status message
	statusMessage string

	// Quit flag
	quitting bool
}

// New creates a new editor model with an empty buffer.
func New() *Model {
	tabs := NewTabManager()
	tab := tabs.ActiveTab()
	return &Model{
		tabs:               tabs,
		split:              NewSplitManager(),
		buffer:             tab.buffer,
		history:            tab.history,
		filename:           tab.filename,
		encoding:           tab.encoding,
		lineEnding:         tab.lineEnding,
		mode:               ModeNormal,
		modified:           false, // New file, not modified yet
		showLineNumbers:    true,
		syntaxHighlighting: true,
		showTabs:           true,
		macro:              NewMacroRecorder(),
		cachedLines:        make(map[int]string),
		dirtyLines:         make(map[int]bool),
	}
}

// NewWithContent creates a new editor model with initial content.
func NewWithContent(content string) *Model {
	tabs := NewTabManager()
	tab := tabs.ActiveTab()
	tab.buffer = buffer.NewFromString(content)
	return &Model{
		tabs:               tabs,
		split:              NewSplitManager(),
		buffer:             tab.buffer,
		history:            tab.history,
		filename:           tab.filename,
		encoding:           tab.encoding,
		lineEnding:         tab.lineEnding,
		mode:               ModeNormal,
		modified:           false, // Content loaded, not modified yet
		showLineNumbers:    true,
		syntaxHighlighting: true,
		showTabs:           true,
		macro:              NewMacroRecorder(),
		cachedLines:        make(map[int]string),
		dirtyLines:         make(map[int]bool),
	}
}

// NewFromFile creates a new editor model for a specific file.
func NewFromFile(filepath, filename, content string) *Model {
	tabs := &TabManager{
		tabs:        []*Tab{NewTabFromFile(filepath, filename, content, "UTF-8", "LF")},
		activeIndex: 0,
	}
	tab := tabs.ActiveTab()
	return &Model{
		tabs:               tabs,
		split:              NewSplitManager(),
		buffer:             tab.buffer,
		history:            tab.history,
		filename:           tab.filename,
		filepath:           tab.filepath,
		encoding:           tab.encoding,
		lineEnding:         tab.lineEnding,
		mode:               ModeNormal,
		modified:           false, // Explicitly set - file just loaded
		showLineNumbers:    true,
		syntaxHighlighting: true,
		showTabs:           true,
		macro:              NewMacroRecorder(),
		cachedLines:        make(map[int]string),
		dirtyLines:         make(map[int]bool),
	}
}

// NewFromFileWithInfo creates a new editor model with file metadata.
func NewFromFileWithInfo(filepath, filename, content, encoding, lineEnding string) *Model {
	tabs := &TabManager{
		tabs:        []*Tab{NewTabFromFile(filepath, filename, content, encoding, lineEnding)},
		activeIndex: 0,
	}
	tab := tabs.ActiveTab()
	return &Model{
		tabs:               tabs,
		split:              NewSplitManager(),
		buffer:             tab.buffer,
		history:            tab.history,
		filename:           tab.filename,
		filepath:           tab.filepath,
		encoding:           tab.encoding,
		lineEnding:         tab.lineEnding,
		mode:               ModeNormal,
		modified:           false, // Explicitly set - file just loaded
		showLineNumbers:    true,
		syntaxHighlighting: true,
		showTabs:           true,
		macro:              NewMacroRecorder(),
		cachedLines:        make(map[int]string),
		dirtyLines:         make(map[int]bool),
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

// Encoding returns the file encoding.
func (m *Model) Encoding() string {
	return m.encoding
}

// SetEncoding sets the file encoding.
func (m *Model) SetEncoding(encoding string) {
	m.encoding = encoding
}

// LineEnding returns the line ending type.
func (m *Model) LineEnding() string {
	return m.lineEnding
}

// SetLineEnding sets the line ending type.
func (m *Model) SetLineEnding(lineEnding string) {
	m.lineEnding = lineEnding
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
				m.updateHighlighter()
				return
			}
		}
		m.filename = path
		m.updateHighlighter()
	}
}

// Content returns the buffer content as string.
func (m *Model) Content() string {
	return m.buffer.String()
}

// SetReadonly sets the readonly mode.
func (m *Model) SetReadonly(readonly bool) {
	m.readonly = readonly
}

// GotoLine moves cursor to specified line and column.
func (m *Model) GotoLine(line, col int) {
	// Convert to 0-indexed
	line--
	col--
	if line < 0 {
		line = 0
	}
	if col < 0 {
		col = 0
	}

	maxLine := m.buffer.LineCount() - 1
	if line > maxLine {
		line = maxLine
	}

	lineStart := m.buffer.LineStart(line)
	lineEnd := m.buffer.LineEnd(line)
	lineLen := lineEnd - lineStart

	if col > lineLen {
		col = lineLen
	}

	m.buffer.MoveTo(lineStart + col)
}

// SetTheme sets the editor theme by name.
func SetTheme(name string) {
	theme := styles.GetTheme(name)
	applyTheme(theme)
}

// GetCurrentTheme returns the name of the current theme.
func GetCurrentTheme() string {
	return currentTheme.Name
}

// SetShowLineNumbers sets whether line numbers are shown.
func (m *Model) SetShowLineNumbers(show bool) {
	m.showLineNumbers = show
}

// ToggleLineNumbers toggles line number display.
func (m *Model) ToggleLineNumbers() {
	m.showLineNumbers = !m.showLineNumbers
}

// SetWordWrap sets whether word wrap is enabled.
func (m *Model) SetWordWrap(wrap bool) {
	m.wordWrap = wrap
}

// ToggleWordWrap toggles word wrap.
func (m *Model) ToggleWordWrap() {
	m.wordWrap = !m.wordWrap
}

// SetSyntaxHighlighting sets whether syntax highlighting is enabled.
func (m *Model) SetSyntaxHighlighting(enabled bool) {
	m.syntaxHighlighting = enabled
}

// ToggleSyntaxHighlighting toggles syntax highlighting.
func (m *Model) ToggleSyntaxHighlighting() {
	m.syntaxHighlighting = !m.syntaxHighlighting
}

// SetTrimTrailingSpaces sets whether to trim trailing whitespace on save.
func (m *Model) SetTrimTrailingSpaces(trim bool) {
	m.trimTrailingSpaces = trim
}

// SetFinalNewline sets whether to ensure file ends with newline.
func (m *Model) SetFinalNewline(add bool) {
	m.finalNewline = add
}

// SetCreateBackup sets whether to create backup files on save.
func (m *Model) SetCreateBackup(backup bool) {
	m.createBackup = backup
}

// TabCount returns the number of open tabs.
func (m *Model) TabCount() int {
	return m.tabs.Count()
}

// ActiveTabIndex returns the index of the active tab.
func (m *Model) ActiveTabIndex() int {
	return m.tabs.ActiveIndex()
}

// syncFromActiveTab updates model fields from the active tab.
func (m *Model) syncFromActiveTab() {
	tab := m.tabs.ActiveTab()
	if tab == nil {
		return
	}
	m.buffer = tab.buffer
	m.history = tab.history
	m.filename = tab.filename
	m.filepath = tab.filepath
	m.encoding = tab.encoding
	m.lineEnding = tab.lineEnding
	m.modified = tab.modified
	m.readonly = tab.readonly
	m.viewportTopLine = tab.viewportTopLine
	m.viewportLeftColumn = tab.viewportLeftColumn
	m.selecting = tab.selecting
	m.selectionStart = tab.selectionStart
	m.selectionEnd = tab.selectionEnd
	m.searchQuery = tab.searchQuery
	m.searchMatches = tab.searchMatches
	m.searchIndex = tab.searchIndex
	m.updateHighlighter() // Update highlighter for new tab

	// Restore cursor position
	if tab.cursorPos > 0 && tab.cursorPos <= m.buffer.Len() {
		m.buffer.MoveTo(tab.cursorPos)
	}
}

// syncToActiveTab saves model fields to the active tab.
func (m *Model) syncToActiveTab() {
	tab := m.tabs.ActiveTab()
	if tab == nil {
		return
	}
	tab.buffer = m.buffer
	tab.history = m.history
	tab.filename = m.filename
	tab.filepath = m.filepath
	tab.encoding = m.encoding
	tab.lineEnding = m.lineEnding
	tab.modified = m.modified
	tab.readonly = m.readonly
	tab.viewportTopLine = m.viewportTopLine
	tab.viewportLeftColumn = m.viewportLeftColumn
	tab.cursorPos = m.buffer.CursorPos()
	tab.selecting = m.selecting
	tab.selectionStart = m.selectionStart
	tab.selectionEnd = m.selectionEnd
	tab.searchQuery = m.searchQuery
	tab.searchMatches = m.searchMatches
	tab.searchIndex = m.searchIndex
}

// NextTab switches to the next tab.
func (m *Model) NextTab() {
	if m.tabs.Count() <= 1 {
		return
	}
	m.syncToActiveTab()
	m.tabs.NextTab()
	m.syncFromActiveTab()
}

// PrevTab switches to the previous tab.
func (m *Model) PrevTab() {
	if m.tabs.Count() <= 1 {
		return
	}
	m.syncToActiveTab()
	m.tabs.PrevTab()
	m.syncFromActiveTab()
}

// SelectTab switches to a specific tab.
func (m *Model) SelectTab(index int) {
	if index == m.tabs.ActiveIndex() {
		return
	}
	m.syncToActiveTab()
	if m.tabs.SelectTab(index) {
		m.syncFromActiveTab()
	}
}

// NewTab creates a new empty tab.
func (m *Model) NewTab() {
	m.syncToActiveTab()
	m.tabs.AddEmptyTab()
	m.syncFromActiveTab()
	m.SetStatusMessage("New tab created")
}

// OpenFileInNewTab opens a file in a new tab.
func (m *Model) OpenFileInNewTab(filepath, filename, content, encoding, lineEnding string) {
	m.syncToActiveTab()
	tab := NewTabFromFile(filepath, filename, content, encoding, lineEnding)
	m.tabs.AddTab(tab)
	m.syncFromActiveTab()
}

// CloseTab closes the current tab.
// Returns true if the tab was closed, false if it's the last tab.
func (m *Model) CloseTab() bool {
	if m.tabs.Count() <= 1 {
		return false
	}
	if m.tabs.CloseActiveTab() {
		m.syncFromActiveTab()
		return true
	}
	return false
}

// HasUnsavedTabs returns true if any tab has unsaved changes.
func (m *Model) HasUnsavedTabs() bool {
	m.syncToActiveTab()
	return m.tabs.HasUnsavedChanges()
}

// GetTabNames returns the filenames of all tabs.
func (m *Model) GetTabNames() []string {
	names := make([]string, m.tabs.Count())
	for i, tab := range m.tabs.Tabs() {
		names[i] = tab.filename
		if tab.modified {
			names[i] += " *"
		}
	}
	return names
}

// SetShowTabs sets whether to show the tab bar.
func (m *Model) SetShowTabs(show bool) {
	m.showTabs = show
}

// ToggleShowTabs toggles the tab bar visibility.
func (m *Model) ToggleShowTabs() {
	m.showTabs = !m.showTabs
}

// IsOverwriteMode returns true if in overwrite mode.
func (m *Model) IsOverwriteMode() bool {
	return m.overwriteMode
}

// SetOverwriteMode sets the overwrite mode.
func (m *Model) SetOverwriteMode(overwrite bool) {
	m.overwriteMode = overwrite
}

// ToggleOverwriteMode toggles between insert and overwrite mode.
func (m *Model) ToggleOverwriteMode() {
	m.overwriteMode = !m.overwriteMode
	if m.overwriteMode {
		m.SetStatusMessage("Overwrite mode")
	} else {
		m.SetStatusMessage("Insert mode")
	}
}

// SetAutoSaveInterval sets the auto-save interval in seconds.
func (m *Model) SetAutoSaveInterval(seconds int) {
	m.autoSaveInterval = seconds
}

// GetAutoSaveInterval returns the auto-save interval in seconds.
func (m *Model) GetAutoSaveInterval() int {
	return m.autoSaveInterval
}

// ShouldAutoSave checks if auto-save should trigger now.
func (m *Model) ShouldAutoSave() bool {
	if m.autoSaveInterval <= 0 || !m.modified || m.filepath == "" {
		return false
	}
	now := time.Now().Unix()
	return now-m.lastSaveTime >= int64(m.autoSaveInterval)
}

// UpdateLastSaveTime updates the last save timestamp.
func (m *Model) UpdateLastSaveTime() {
	m.lastSaveTime = time.Now().Unix()
}

// SetFileChanged sets the file changed flag.
func (m *Model) SetFileChanged(changed bool) {
	m.fileChanged = changed
}

// IsFileChanged returns true if file was changed externally.
func (m *Model) IsFileChanged() bool {
	return m.fileChanged
}

// InvalidateCache marks all cached lines as dirty.
func (m *Model) InvalidateCache() {
	m.cachedLines = make(map[int]string)
	m.dirtyLines = make(map[int]bool)
}

// InvalidateLine marks a specific line as dirty.
func (m *Model) InvalidateLine(line int) {
	if m.dirtyLines == nil {
		m.dirtyLines = make(map[int]bool)
	}
	m.dirtyLines[line] = true
	delete(m.cachedLines, line)
}

// InvalidateLineRange marks a range of lines as dirty.
func (m *Model) InvalidateLineRange(startLine, endLine int) {
	for i := startLine; i <= endLine; i++ {
		m.InvalidateLine(i)
	}
}

// GetCachedLine returns a cached line if available.
func (m *Model) GetCachedLine(line int) (string, bool) {
	if m.cachedLines == nil {
		return "", false
	}
	content, ok := m.cachedLines[line]
	return content, ok
}

// SetCachedLine caches a rendered line.
func (m *Model) SetCachedLine(line int, content string) {
	if m.cachedLines == nil {
		m.cachedLines = make(map[int]string)
	}
	m.cachedLines[line] = content
	delete(m.dirtyLines, line)
}

// IsLineDirty checks if a line needs re-rendering.
func (m *Model) IsLineDirty(line int) bool {
	if m.dirtyLines == nil {
		return true
	}
	_, dirty := m.dirtyLines[line]
	_, cached := m.cachedLines[line]
	return dirty || !cached
}

// StartSmoothScroll initiates smooth scrolling to a target line.
func (m *Model) StartSmoothScroll(targetLine int) {
	m.targetTopLine = targetLine
	m.scrollAnimating = true
}

// UpdateSmoothScroll advances the smooth scroll animation.
// Returns true if animation is still in progress.
func (m *Model) UpdateSmoothScroll() bool {
	if !m.scrollAnimating {
		return false
	}

	// Calculate scroll step (ease-out effect)
	diff := m.targetTopLine - m.viewportTopLine
	if diff == 0 {
		m.scrollAnimating = false
		return false
	}

	// Move by a fraction of the remaining distance (smooth easing)
	step := diff / 3
	if step == 0 {
		if diff > 0 {
			step = 1
		} else {
			step = -1
		}
	}

	m.viewportTopLine += step

	// Check if we've reached the target
	if m.viewportTopLine == m.targetTopLine {
		m.scrollAnimating = false
		return false
	}

	return true
}

// IsScrollAnimating returns true if smooth scroll is in progress.
func (m *Model) IsScrollAnimating() bool {
	return m.scrollAnimating
}

// StopSmoothScroll immediately stops the scroll animation.
func (m *Model) StopSmoothScroll() {
	m.scrollAnimating = false
	m.viewportTopLine = m.targetTopLine
}

// IsSplit returns true if the view is split.
func (m *Model) IsSplit() bool {
	return m.split.IsSplit()
}

// SplitHorizontal creates a horizontal split (left/right).
func (m *Model) SplitHorizontal() {
	if m.split.IsSplit() {
		m.SetStatusMessage("Already split - close first (Ctrl+Shift+\\)")
		return
	}
	// Split showing the same tab in both panes
	m.split.SplitHorizontal(m.tabs.ActiveIndex())
	m.SetStatusMessage("Horizontal split created")
}

// SplitVertical creates a vertical split (top/bottom).
func (m *Model) SplitVertical() {
	if m.split.IsSplit() {
		m.SetStatusMessage("Already split - close first (Ctrl+Shift+\\)")
		return
	}
	// Split showing the same tab in both panes
	m.split.SplitVertical(m.tabs.ActiveIndex())
	m.SetStatusMessage("Vertical split created")
}

// CloseSplit closes the split view.
func (m *Model) CloseSplit() {
	if !m.split.IsSplit() {
		m.SetStatusMessage("No split to close")
		return
	}
	m.split.CloseSplit()
	m.SetStatusMessage("Split closed")
}

// NextPane switches to the next pane.
func (m *Model) NextPane() {
	if !m.split.IsSplit() {
		return
	}
	// Save current pane state
	m.split.SavePaneState(
		m.split.ActivePaneIndex(),
		m.viewportTopLine,
		m.viewportLeftColumn,
		m.buffer.CursorPos(),
	)

	m.split.NextPane()

	// Restore new pane state
	topLine, leftCol, cursorPos := m.split.RestorePaneState(m.split.ActivePaneIndex())
	m.viewportTopLine = topLine
	m.viewportLeftColumn = leftCol

	// Switch to the tab this pane is showing
	paneTabIndex := m.split.GetActiveTabIndex()
	if paneTabIndex != m.tabs.ActiveIndex() {
		m.syncToActiveTab()
		m.tabs.SelectTab(paneTabIndex)
		m.syncFromActiveTab()
	}

	// Restore cursor position
	if cursorPos > 0 && cursorPos <= m.buffer.Len() {
		m.buffer.MoveTo(cursorPos)
	}
}

// PrevPane switches to the previous pane.
func (m *Model) PrevPane() {
	if !m.split.IsSplit() {
		return
	}
	// Save current pane state
	m.split.SavePaneState(
		m.split.ActivePaneIndex(),
		m.viewportTopLine,
		m.viewportLeftColumn,
		m.buffer.CursorPos(),
	)

	m.split.PrevPane()

	// Restore new pane state
	topLine, leftCol, cursorPos := m.split.RestorePaneState(m.split.ActivePaneIndex())
	m.viewportTopLine = topLine
	m.viewportLeftColumn = leftCol

	// Switch to the tab this pane is showing
	paneTabIndex := m.split.GetActiveTabIndex()
	if paneTabIndex != m.tabs.ActiveIndex() {
		m.syncToActiveTab()
		m.tabs.SelectTab(paneTabIndex)
		m.syncFromActiveTab()
	}

	// Restore cursor position
	if cursorPos > 0 && cursorPos <= m.buffer.Len() {
		m.buffer.MoveTo(cursorPos)
	}
}

// SetPaneTab sets which tab the active pane shows.
func (m *Model) SetPaneTab(tabIndex int) {
	if tabIndex >= 0 && tabIndex < m.tabs.Count() {
		m.split.SetPaneTab(m.split.ActivePaneIndex(), tabIndex)
		m.syncToActiveTab()
		m.tabs.SelectTab(tabIndex)
		m.syncFromActiveTab()
	}
}

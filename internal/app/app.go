package app

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/KilimcininKorOglu/gesh/internal/buffer"
	"github.com/KilimcininKorOglu/gesh/internal/file"
	"github.com/KilimcininKorOglu/gesh/internal/syntax"
	_ "github.com/KilimcininKorOglu/gesh/internal/syntax/languages" // Register languages
	"github.com/KilimcininKorOglu/gesh/internal/ui/styles"
)

// Styles for the UI components (initialized with default theme)
var (
	currentTheme = styles.DarkTheme

	headerStyle      lipgloss.Style
	statusStyle      lipgloss.Style
	helpStyle        lipgloss.Style
	helpKeyStyle     lipgloss.Style
	lineNumberStyle  lipgloss.Style
	editorStyle      lipgloss.Style
	selectionStyle   lipgloss.Style
	searchMatchStyle lipgloss.Style

	// Syntax highlighting styles
	syntaxKeywordStyle  lipgloss.Style
	syntaxTypeStyle     lipgloss.Style
	syntaxStringStyle   lipgloss.Style
	syntaxNumberStyle   lipgloss.Style
	syntaxCommentStyle  lipgloss.Style
	syntaxOperatorStyle lipgloss.Style
	syntaxFunctionStyle lipgloss.Style
	syntaxVariableStyle lipgloss.Style
	syntaxConstantStyle lipgloss.Style
	syntaxBuiltinStyle  lipgloss.Style
)

func init() {
	applyTheme(currentTheme)
}

// applyTheme updates all styles based on the given theme.
func applyTheme(theme styles.Theme) {
	currentTheme = theme

	headerStyle = lipgloss.NewStyle().
		Background(theme.HeaderBg).
		Foreground(theme.HeaderFg).
		Bold(true)

	statusStyle = lipgloss.NewStyle().
		Background(theme.StatusBg).
		Foreground(theme.StatusFg)

	helpStyle = lipgloss.NewStyle().
		Background(theme.HelpBg).
		Foreground(theme.HelpFg)

	helpKeyStyle = lipgloss.NewStyle().
		Background(theme.HelpBg).
		Foreground(theme.HelpKeyFg).
		Bold(true)

	lineNumberStyle = lipgloss.NewStyle().
		Foreground(theme.LineNumberFg).
		Width(4).
		Align(lipgloss.Right)

	editorStyle = lipgloss.NewStyle().
		Foreground(theme.EditorFg)

	selectionStyle = lipgloss.NewStyle().
		Background(theme.SelectionBg).
		Foreground(theme.SelectionFg)

	searchMatchStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("#ffff00")).
		Foreground(lipgloss.Color("#000000"))

	// Syntax highlighting colors (theme-aware)
	syntaxKeywordStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff79c6"))
	syntaxTypeStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#8be9fd"))
	syntaxStringStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#f1fa8c"))
	syntaxNumberStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#bd93f9"))
	syntaxCommentStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#6272a4"))
	syntaxOperatorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff79c6"))
	syntaxFunctionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#50fa7b"))
	syntaxVariableStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffb86c"))
	syntaxConstantStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#bd93f9"))
	syntaxBuiltinStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#8be9fd"))

	// Update tab styles
	styles.UpdateTabStyles(theme)
}

// autoSaveTickMsg is sent periodically to check for auto-save and file changes.
type autoSaveTickMsg struct{}

// scrollTickMsg is sent during smooth scroll animation.
type scrollTickMsg struct{}

// autoSaveTick returns a command that sends an autoSaveTickMsg after a delay.
func autoSaveTick() tea.Cmd {
	return tea.Tick(time.Second*5, func(t time.Time) tea.Msg {
		return autoSaveTickMsg{}
	})
}

// scrollTick returns a command for smooth scroll animation (16ms = ~60fps).
func scrollTick() tea.Cmd {
	return tea.Tick(time.Millisecond*16, func(t time.Time) tea.Msg {
		return scrollTickMsg{}
	})
}

// Init initializes the model.
func (m *Model) Init() tea.Cmd {
	return autoSaveTick()
}

// Update handles messages and updates the model.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyMsg(msg)
	case tea.MouseMsg:
		return m.handleMouseMsg(msg)
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
		return m, nil
	case autoSaveTickMsg:
		// Check if auto-save should trigger
		if m.ShouldAutoSave() {
			m.autoSave()
		}
		// Check for external file changes
		if m.filepath != "" && !m.fileChanged {
			m.checkFileChanged()
		}
		// Continue ticking
		return m, autoSaveTick()
	case scrollTickMsg:
		// Update smooth scroll animation
		if m.UpdateSmoothScroll() {
			// Continue animation
			return m, scrollTick()
		}
		// Animation complete
		return m, nil
	}
	return m, nil
}

// handleMouseMsg processes mouse input.
func (m *Model) handleMouseMsg(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	// Only handle in normal mode
	if m.mode != ModeNormal {
		return m, nil
	}

	switch msg.Button {
	case tea.MouseButtonLeft:
		// User clicked - reset mouse scrolling mode
		m.mouseScrolling = false

		// Calculate which line was clicked (account for header)
		clickedLine := m.viewportTopLine + msg.Y - 1 // -1 for header

		// Check bounds
		if clickedLine < 0 {
			clickedLine = 0
		}
		if clickedLine >= m.buffer.LineCount() {
			clickedLine = m.buffer.LineCount() - 1
		}

		// Calculate column (account for line numbers: "→123 │ " = ~7 chars)
		lineNumWidth := 7
		clickedCol := msg.X - lineNumWidth
		if clickedCol < 0 {
			clickedCol = 0
		}

		// Get line content and clamp column
		lineContent := m.buffer.Line(clickedLine)
		lineLen := len([]rune(lineContent))
		if clickedCol > lineLen {
			clickedCol = lineLen
		}

		// Calculate target position
		lineStart := m.buffer.LineStart(clickedLine)
		targetPos := lineStart + clickedCol

		if msg.Action == tea.MouseActionPress {
			// Start selection on mouse down
			m.buffer.MoveTo(targetPos)
			m.selectionStart = targetPos
			m.selectionEnd = targetPos
			m.selecting = true
		} else if msg.Action == tea.MouseActionMotion && m.selecting {
			// Extend selection while dragging
			m.buffer.MoveTo(targetPos)
			m.selectionEnd = targetPos
		} else if msg.Action == tea.MouseActionRelease {
			// Finish selection on mouse up
			m.buffer.MoveTo(targetPos)
			m.selectionEnd = targetPos
			// If start == end, cancel selection (just a click)
			if m.selectionStart == m.selectionEnd {
				m.selecting = false
			}
		}

	case tea.MouseButtonRight:
		// Right click - only handle on press, not release
		if msg.Action != tea.MouseActionPress {
			return m, nil
		}
		// Copy if selection exists, paste if not
		if m.selecting && m.selectionStart != m.selectionEnd {
			m.copySelection()
		} else if !m.readonly {
			m.paste()
		}

	case tea.MouseButtonWheelUp:
		// Scroll up - user controls viewport
		m.mouseScrolling = true
		m.viewportTopLine -= 3
		if m.viewportTopLine < 0 {
			m.viewportTopLine = 0
		}

	case tea.MouseButtonWheelDown:
		// Scroll down - user controls viewport
		m.mouseScrolling = true
		m.viewportTopLine += 3
		// Max scroll = show at least 1 line at bottom
		visibleLines := m.height - 4 // header, status, 2 help lines
		maxTop := m.buffer.LineCount() - visibleLines
		if maxTop < 0 {
			maxTop = 0
		}
		if m.viewportTopLine > maxTop {
			m.viewportTopLine = maxTop
		}
	}

	return m, nil
}

// handleKeyMsg processes keyboard input.
func (m *Model) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Reset mouse scrolling mode on any key press
	m.mouseScrolling = false

	// Record key for macro (if recording)
	if m.macro != nil {
		m.macro.RecordKey(msg)
	}

	// Handle quit confirmation mode
	if m.mode == ModeQuit {
		switch msg.String() {
		case "y", "Y":
			// Save and quit
			m.saveFile()
			m.quitting = true
			return m, tea.Quit
		case "n", "N":
			// Quit without saving
			m.quitting = true
			return m, tea.Quit
		case "c", "C", "esc":
			// Cancel - go back to editing
			m.mode = ModeNormal
			m.SetStatusMessage("")
			return m, nil
		}
		return m, nil
	}

	// Handle save-as mode
	if m.mode == ModeSaveAs {
		return m.handleSaveAsInput(msg)
	}

	// Handle goto mode
	if m.mode == ModeGoto {
		return m.handleGotoInput(msg)
	}

	// Handle search mode
	if m.mode == ModeSearch {
		return m.handleSearchInput(msg)
	}

	// Handle replace mode
	if m.mode == ModeReplace {
		return m.handleReplaceInput(msg)
	}

	// Handle replace confirm mode
	if m.mode == ModeReplaceConfirm {
		return m.handleReplaceConfirm(msg)
	}

	// Handle replace all mode
	if m.mode == ModeReplaceAll {
		return m.handleReplaceAllInput(msg)
	}

	// Handle replace all confirm mode
	if m.mode == ModeReplaceAllConfirm {
		return m.handleReplaceAllConfirm(msg)
	}

	// Handle open file mode
	if m.mode == ModeOpen {
		return m.handleOpenInput(msg)
	}

	// Handle save macro mode
	if m.mode == ModeSaveMacro {
		return m.handleSaveMacroInput(msg)
	}

	// Handle load macro mode
	if m.mode == ModeLoadMacro {
		return m.handleLoadMacroInput(msg)
	}

	// Normal mode key handling - NANO COMPATIBLE
	switch msg.String() {

	// ==================== NANO FILE OPERATIONS ====================

	case "ctrl+x":
		// Nano: Exit (ask to save if modified)
		if m.modified {
			m.mode = ModeQuit
			m.SetStatusMessage("Save modified buffer? (Y)es, (N)o, (C)ancel")
			return m, nil
		}
		m.quitting = true
		return m, tea.Quit

	case "ctrl+o":
		// Nano: Write Out (Save)
		if m.filepath == "" {
			m.mode = ModeSaveAs
			m.inputBuffer = ""
			m.inputPrompt = "File Name to Write: "
			return m, nil
		}
		return m.saveFile()

	case "ctrl+r":
		// Nano: Read File (Insert file at cursor)
		m.mode = ModeOpen
		m.inputBuffer = ""
		m.inputPrompt = "File to insert: "
		return m, nil

	case "ctrl+g":
		// Nano: Display help text (show version/about info)
		m.SetStatusMessage("GESH 1.0.0 - Nano-compatible text editor | ^X Exit | ^O Save | ^W Search")
		return m, nil

	// ==================== NANO SEARCH & REPLACE ====================

	case "ctrl+w":
		// Nano: Where Is (Search)
		m.mode = ModeSearch
		m.inputBuffer = m.searchQuery
		m.inputPrompt = "Search: "
		return m, nil

	case "alt+w":
		// Nano: Next search result
		m.nextMatch()
		return m, nil

	case "ctrl+q":
		// Nano: Search backward (Where Was)
		m.prevMatch()
		return m, nil

	case "ctrl+\\", "alt+r":
		// Nano: Replace
		m.mode = ModeReplace
		m.inputBuffer = m.searchQuery
		m.inputPrompt = "Search (to replace): "
		return m, nil

	// ==================== NANO EDITING ====================

	case "ctrl+k":
		// Nano: Cut line (or selection)
		if m.readonly {
			m.SetStatusMessage("File is read-only")
			return m, nil
		}
		if m.selecting {
			m.cutSelection()
		} else {
			m.cutLine()
		}
		return m, nil

	case "ctrl+u":
		// Nano: Uncut/Paste
		if m.readonly {
			m.SetStatusMessage("File is read-only")
			return m, nil
		}
		m.paste()
		return m, nil

	case "alt+6":
		// Nano: Copy line (or selection)
		if m.selecting {
			m.copySelection()
		} else {
			m.copyLine()
		}
		return m, nil

	case "alt+u":
		// Nano: Undo
		m.undo()
		return m, nil

	case "alt+e":
		// Nano: Redo
		m.redo()
		return m, nil

	case "ctrl+j":
		// Nano: Justify (not implemented, show message)
		m.SetStatusMessage("Justify not implemented")
		return m, nil

	case "ctrl+t":
		// Nano: Execute command / Spell check (we use for new tab)
		m.NewTab()
		return m, nil

	// ==================== NANO NAVIGATION ====================

	case "ctrl+y":
		// Nano: Page Up
		m.clearSelection()
		m.pageUp()
		return m, scrollTick()

	case "ctrl+v":
		// Nano: Page Down
		m.clearSelection()
		m.pageDown()
		return m, scrollTick()

	case "alt+\\":
		// Nano: Go to beginning of file
		m.clearSelection()
		m.buffer.MoveToStart()
		m.ensureCursorVisible()
		return m, nil

	case "alt+/":
		// Nano: Go to end of file
		m.clearSelection()
		m.buffer.MoveToEnd()
		m.ensureCursorVisible()
		return m, nil

	case "ctrl+_", "alt+g":
		// Nano: Go to line
		m.mode = ModeGoto
		m.inputBuffer = ""
		m.inputPrompt = fmt.Sprintf("Enter line number [1-%d]: ", m.buffer.LineCount())
		return m, nil

	case "ctrl+c":
		// Nano: Show cursor position
		line := m.buffer.CurrentLine() + 1
		col := m.buffer.CurrentColumn() + 1
		total := m.buffer.LineCount()
		pos := m.buffer.CursorPos()
		length := m.buffer.Len()
		percent := 0
		if length > 0 {
			percent = (pos * 100) / length
		}
		m.SetStatusMessage(fmt.Sprintf("line %d/%d (%d%%), col %d, char %d/%d", line, total, percent, col, pos, length))
		return m, nil

	case "ctrl+a":
		// Nano: Go to beginning of line
		m.clearSelection()
		m.moveToLineStart()
		return m, nil

	case "ctrl+e":
		// Nano: Go to end of line
		m.clearSelection()
		m.moveToLineEnd()
		return m, nil

	case "ctrl+p", "up":
		// Nano: Previous line
		m.clearSelection()
		m.moveCursorUp()
		return m, nil

	case "ctrl+n", "down":
		// Nano: Next line
		m.clearSelection()
		m.moveCursorDown()
		return m, nil

	case "ctrl+b", "left":
		// Nano: Back one character
		m.clearSelection()
		m.buffer.MoveLeft()
		return m, nil

	case "ctrl+f", "right":
		// Nano: Forward one character
		m.clearSelection()
		m.buffer.MoveRight()
		return m, nil

	case "alt+space", "ctrl+left":
		// Nano: Back one word
		m.clearSelection()
		m.moveWordLeft()
		return m, nil

	case "ctrl+space", "ctrl+right":
		// Nano: Forward one word
		m.clearSelection()
		m.moveWordRight()
		return m, nil

	case "pgup":
		m.clearSelection()
		m.pageUp()
		return m, scrollTick()

	case "pgdown":
		m.clearSelection()
		m.pageDown()
		return m, scrollTick()

	case "home":
		m.clearSelection()
		m.moveToLineStart()
		return m, nil

	case "end":
		m.clearSelection()
		m.moveToLineEnd()
		return m, nil

	case "ctrl+home":
		m.clearSelection()
		m.buffer.MoveToStart()
		m.ensureCursorVisible()
		return m, nil

	case "ctrl+end":
		m.clearSelection()
		m.buffer.MoveToEnd()
		m.ensureCursorVisible()
		return m, nil

	// ==================== NANO SELECTION (Mark) ====================

	case "alt+a", "ctrl+6":
		// Nano: Set mark / Start selection
		m.toggleSelection()
		return m, nil

	// Shift+arrow selection (modern extension)
	case "shift+up":
		m.startSelection()
		m.moveCursorUp()
		m.updateSelection()
		return m, nil
	case "shift+down":
		m.startSelection()
		m.moveCursorDown()
		m.updateSelection()
		return m, nil
	case "shift+left":
		m.startSelection()
		m.buffer.MoveLeft()
		m.updateSelection()
		return m, nil
	case "shift+right":
		m.startSelection()
		m.buffer.MoveRight()
		m.updateSelection()
		return m, nil

	// ==================== NANO DELETION ====================

	case "ctrl+h", "backspace":
		// Nano: Delete character before cursor
		if m.readonly {
			m.SetStatusMessage("File is read-only")
			return m, nil
		}
		pos := m.buffer.CursorPos()
		if r := m.buffer.Delete(); r != 0 {
			m.history.Push(buffer.EditOperation{
				Type:     buffer.OpDelete,
				Position: pos - 1,
				Text:     string(r),
			})
			m.setModified()
		}
		return m, nil

	case "ctrl+d", "delete":
		// Nano: Delete character under cursor
		if m.readonly {
			m.SetStatusMessage("File is read-only")
			return m, nil
		}
		pos := m.buffer.CursorPos()
		if r := m.buffer.DeleteForward(); r != 0 {
			m.history.Push(buffer.EditOperation{
				Type:     buffer.OpDelete,
				Position: pos,
				Text:     string(r),
			})
			m.setModified()
		}
		return m, nil

	case "alt+backspace":
		// Nano: Delete word to the left
		if m.readonly {
			m.SetStatusMessage("File is read-only")
			return m, nil
		}
		m.deleteWordLeft()
		return m, nil

	case "ctrl+delete":
		// Nano: Delete word to the right
		if m.readonly {
			m.SetStatusMessage("File is read-only")
			return m, nil
		}
		m.deleteWordRight()
		return m, nil

	// ==================== NANO OTHER ====================

	case "alt+n":
		// Nano: Toggle line numbers
		m.showLineNumbers = !m.showLineNumbers
		return m, nil

	case "alt+p":
		// Nano: Toggle whitespace display (not implemented)
		m.SetStatusMessage("Whitespace display not implemented")
		return m, nil

	case "alt+x":
		// Nano: Toggle expert mode (we just show a message)
		m.SetStatusMessage("Expert mode not available")
		return m, nil

	case "ctrl+l":
		// Nano: Refresh screen
		return m, nil

	case "ctrl+m", "enter":
		// Nano: Insert newline (Enter)
		if m.readonly {
			m.SetStatusMessage("File is read-only")
			return m, nil
		}
		currentLine := m.buffer.CurrentLine()
		lineContent := m.buffer.Line(currentLine)
		indent := getIndent(lineContent)

		pos := m.buffer.CursorPos()
		insertText := "\n" + indent
		m.buffer.Insert('\n')
		if indent != "" {
			m.buffer.InsertString(indent)
		}
		m.history.Push(buffer.EditOperation{
			Type:     buffer.OpInsert,
			Position: pos,
			Text:     insertText,
		})
		m.setModified()
		return m, nil

	case "ctrl+i", "tab":
		// Nano: Insert tab
		if m.readonly {
			m.SetStatusMessage("File is read-only")
			return m, nil
		}
		pos := m.buffer.CursorPos()
		m.buffer.InsertString("    ")
		m.history.Push(buffer.EditOperation{
			Type:     buffer.OpInsert,
			Position: pos,
			Text:     "    ",
		})
		m.setModified()
		return m, nil

	case "insert":
		// Toggle insert/overwrite mode
		m.ToggleOverwriteMode()
		return m, nil

	// ==================== TAB MANAGEMENT (Extension) ====================

	case "ctrl+tab", "ctrl+pgdn":
		m.NextTab()
		return m, nil

	case "ctrl+shift+tab", "ctrl+pgup":
		m.PrevTab()
		return m, nil

	// ==================== SPLIT VIEW (Extension) ====================

	case "alt+\\\\":
		m.SplitHorizontal()
		return m, nil

	case "alt+-":
		m.SplitVertical()
		return m, nil

	case "alt+c":
		m.CloseSplit()
		return m, nil

	case "alt+left", "alt+h":
		if m.IsSplit() && m.split.ActivePaneIndex() == 1 {
			m.PrevPane()
		}
		return m, nil

	case "alt+right", "alt+l":
		if m.IsSplit() && m.split.ActivePaneIndex() == 0 {
			m.NextPane()
		}
		return m, nil

	// ==================== MACRO (Extension) ====================

	case "f4":
		// Toggle macro recording
		if m.macro.ToggleRecording() {
			m.SetStatusMessage("Recording macro...")
		} else {
			m.SetStatusMessage(fmt.Sprintf("Macro recorded (%d keys)", m.macro.KeyCount()))
		}
		return m, nil

	case "f5":
		// Play macro
		if m.macro.IsRecording() {
			m.SetStatusMessage("Stop recording first (F4)")
			return m, nil
		}
		if m.macro.KeyCount() == 0 {
			m.SetStatusMessage("No macro recorded")
			return m, nil
		}
		return m.playMacro()

	case "f3":
		// Find next (also nano compatible)
		m.nextMatch()
		return m, nil

	case "shift+f3":
		m.prevMatch()
		return m, nil

	default:
		// Insert printable characters only
		// Ignore: modifier keys alone, control characters, non-printable
		// Only accept KeyRunes type with actual printable runes
		if msg.Type == tea.KeyRunes && len(msg.Runes) > 0 && !msg.Alt {
			// Additional check: ensure runes are printable (>= space)
			printable := true
			for _, r := range msg.Runes {
				if r < 32 { // Control characters
					printable = false
					break
				}
			}
			if !printable {
				return m, nil
			}
			if m.readonly {
				m.SetStatusMessage("File is read-only")
				return m, nil
			}
			pos := m.buffer.CursorPos()
			text := string(msg.Runes)

			if m.overwriteMode {
				// Overwrite mode: replace characters
				for _, r := range msg.Runes {
					// Don't overwrite past end of line
					if m.buffer.CursorPos() < m.buffer.Len() {
						currentRune := m.buffer.RuneAt(m.buffer.CursorPos())
						if currentRune != '\n' {
							m.buffer.DeleteForward()
						}
					}
					m.buffer.Insert(r)
				}
			} else {
				// Insert mode: normal insert
				for _, r := range msg.Runes {
					m.buffer.Insert(r)
				}
			}

			m.history.Push(buffer.EditOperation{
				Type:     buffer.OpInsert,
				Position: pos,
				Text:     text,
			})
			m.setModified()
		}
	}

	return m, nil
}

// undo reverses the last edit operation.
func (m *Model) undo() {
	op := m.history.Undo()
	if op == nil {
		m.SetStatusMessage("Nothing to undo")
		return
	}

	// Reverse the operation
	if op.Type == buffer.OpInsert {
		// Undo insert: delete the text
		m.buffer.MoveTo(op.Position)
		for range []rune(op.Text) {
			m.buffer.DeleteForward()
		}
	} else {
		// Undo delete: insert the text
		m.buffer.MoveTo(op.Position)
		m.buffer.InsertString(op.Text)
	}

	m.SetStatusMessage("Undo")
}

// redo re-applies the last undone operation.
func (m *Model) redo() {
	op := m.history.Redo()
	if op == nil {
		m.SetStatusMessage("Nothing to redo")
		return
	}

	// Re-apply the operation
	if op.Type == buffer.OpInsert {
		// Redo insert: insert the text
		m.buffer.MoveTo(op.Position)
		m.buffer.InsertString(op.Text)
	} else {
		// Redo delete: delete the text
		m.buffer.MoveTo(op.Position)
		for range []rune(op.Text) {
			m.buffer.DeleteForward()
		}
	}

	m.SetStatusMessage("Redo")
}

// moveCursorUp moves the cursor up one line.
func (m *Model) moveCursorUp() {
	currentLine := m.buffer.CurrentLine()
	if currentLine == 0 {
		return
	}

	currentCol := m.buffer.CurrentColumn()
	targetLineStart := m.buffer.LineStart(currentLine - 1)
	targetLineEnd := m.buffer.LineEnd(currentLine - 1)
	targetLineLen := targetLineEnd - targetLineStart

	// Calculate target position
	targetPos := targetLineStart + currentCol
	if currentCol > targetLineLen {
		targetPos = targetLineEnd
	}

	m.buffer.MoveTo(targetPos)
}

// moveCursorDown moves the cursor down one line.
func (m *Model) moveCursorDown() {
	currentLine := m.buffer.CurrentLine()
	if currentLine >= m.buffer.LineCount()-1 {
		return
	}

	currentCol := m.buffer.CurrentColumn()
	targetLineStart := m.buffer.LineStart(currentLine + 1)
	targetLineEnd := m.buffer.LineEnd(currentLine + 1)
	targetLineLen := targetLineEnd - targetLineStart

	// Calculate target position
	targetPos := targetLineStart + currentCol
	if currentCol > targetLineLen {
		targetPos = targetLineEnd
	}

	m.buffer.MoveTo(targetPos)
}

// moveToLineStart moves cursor to the start of the current line.
func (m *Model) moveToLineStart() {
	currentLine := m.buffer.CurrentLine()
	lineStart := m.buffer.LineStart(currentLine)
	m.buffer.MoveTo(lineStart)
}

// moveToLineEnd moves cursor to the end of the current line.
func (m *Model) moveToLineEnd() {
	currentLine := m.buffer.CurrentLine()
	lineEnd := m.buffer.LineEnd(currentLine)
	m.buffer.MoveTo(lineEnd)
}

// pageUp moves the cursor up by a page.
func (m *Model) pageUp() {
	visibleLines := m.height - 3
	if visibleLines < 1 {
		visibleLines = 1
	}

	currentLine := m.buffer.CurrentLine()
	targetLine := currentLine - visibleLines
	if targetLine < 0 {
		targetLine = 0
	}

	lineStart := m.buffer.LineStart(targetLine)
	if lineStart >= 0 {
		m.buffer.MoveTo(lineStart)
	}

	// Start smooth scroll
	newTopLine := m.viewportTopLine - visibleLines
	if newTopLine < 0 {
		newTopLine = 0
	}
	m.StartSmoothScroll(newTopLine)
}

// pageDown moves the cursor down by a page.
func (m *Model) pageDown() {
	visibleLines := m.height - 3
	if visibleLines < 1 {
		visibleLines = 1
	}

	currentLine := m.buffer.CurrentLine()
	maxLine := m.buffer.LineCount() - 1
	targetLine := currentLine + visibleLines
	if targetLine > maxLine {
		targetLine = maxLine
	}

	lineStart := m.buffer.LineStart(targetLine)
	if lineStart >= 0 {
		m.buffer.MoveTo(lineStart)
	}

	// Start smooth scroll
	maxTopLine := maxLine - visibleLines + 1
	if maxTopLine < 0 {
		maxTopLine = 0
	}
	newTopLine := m.viewportTopLine + visibleLines
	if newTopLine > maxTopLine {
		newTopLine = maxTopLine
	}
	m.StartSmoothScroll(newTopLine)
}

// deleteLine deletes the current line.
func (m *Model) deleteLine() {
	currentLine := m.buffer.CurrentLine()
	lineStart := m.buffer.LineStart(currentLine)
	lineEnd := m.buffer.LineEnd(currentLine)

	if lineStart < 0 || lineEnd < 0 {
		return
	}

	// Include newline if not last line
	deleteEnd := lineEnd
	if currentLine < m.buffer.LineCount()-1 {
		deleteEnd++ // Include newline
	} else if lineStart > 0 {
		// Last line: delete preceding newline instead
		lineStart--
	}

	// Get deleted text for undo
	deletedText := m.buffer.Slice(lineStart, deleteEnd)

	// Delete the line
	m.buffer.MoveTo(lineStart)
	for i := lineStart; i < deleteEnd; i++ {
		m.buffer.DeleteForward()
	}

	// Record for undo
	m.history.Push(buffer.EditOperation{
		Type:     buffer.OpDelete,
		Position: lineStart,
		Text:     deletedText,
	})

	m.setModified()
	m.SetStatusMessage("Line deleted")
}

// checkFileChanged checks if the file was modified externally.
func (m *Model) checkFileChanged() {
	if m.filepath == "" {
		return
	}

	info, err := os.Stat(m.filepath)
	if err != nil {
		return
	}

	// Compare modification time (simple check)
	modTime := info.ModTime().Unix()
	if modTime > m.lastSaveTime && m.lastSaveTime > 0 {
		m.fileChanged = true
		m.SetStatusMessage("⚠ File changed externally! Press Ctrl+O to reload")
	}
}

// autoSave performs an automatic save.
func (m *Model) autoSave() {
	if m.filepath == "" || !m.modified || m.readonly {
		return
	}

	opts := file.SaveOptions{
		TrimTrailingSpaces: m.trimTrailingSpaces,
		FinalNewline:       m.finalNewline,
		CreateBackup:       m.createBackup,
	}
	err := file.SaveWithOptions(m.filepath, m.Content(), opts)
	if err != nil {
		m.SetStatusMessage("Auto-save failed: " + err.Error())
		return
	}

	m.modified = false
	m.UpdateLastSaveTime()
	m.SetStatusMessage("Auto-saved")
}

// playMacro plays back the recorded macro.
func (m *Model) playMacro() (tea.Model, tea.Cmd) {
	if !m.macro.Play() {
		m.SetStatusMessage("No macro to play")
		return m, nil
	}

	// Play all keys in sequence
	keysPlayed := 0
	for {
		key := m.macro.NextKey()
		if key == nil {
			break
		}
		// Recursively handle each key (without recording)
		m.handleKeyMsg(*key)
		keysPlayed++
	}

	m.SetStatusMessage(fmt.Sprintf("Macro played (%d keys)", keysPlayed))
	return m, nil
}

// saveFile saves the buffer to file.
func (m *Model) saveFile() (tea.Model, tea.Cmd) {
	// If no filepath, enter save-as mode
	if m.filepath == "" {
		m.mode = ModeSaveAs
		m.inputBuffer = ""
		m.inputPrompt = "Save as: "
		return m, nil
	}

	// Save to existing filepath with options
	opts := file.SaveOptions{
		TrimTrailingSpaces: m.trimTrailingSpaces,
		FinalNewline:       m.finalNewline,
		CreateBackup:       m.createBackup,
	}
	err := file.SaveWithOptions(m.filepath, m.Content(), opts)
	if err != nil {
		m.SetStatusMessage("Error: " + err.Error())
		return m, nil
	}

	m.modified = false
	m.UpdateLastSaveTime()
	m.SetStatusMessage("Saved: " + m.filename)
	return m, nil
}

// handleSaveAsInput handles input in save-as mode.
func (m *Model) handleSaveAsInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		if m.inputBuffer != "" {
			m.SetFilepath(m.inputBuffer)
			m.mode = ModeNormal
			m.inputBuffer = ""
			return m.saveFile()
		}
		return m, nil

	case "esc":
		m.mode = ModeNormal
		m.inputBuffer = ""
		m.SetStatusMessage("")
		return m, nil

	case "backspace":
		if len(m.inputBuffer) > 0 {
			m.inputBuffer = m.inputBuffer[:len(m.inputBuffer)-1]
		}
		return m, nil

	default:
		if len(msg.Runes) > 0 {
			m.inputBuffer += string(msg.Runes)
		}
		return m, nil
	}
}

// handleGotoInput handles input in goto mode.
func (m *Model) handleGotoInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		if m.inputBuffer != "" {
			// Parse line number
			var lineNum int
			_, err := fmt.Sscanf(m.inputBuffer, "%d", &lineNum)
			if err != nil || lineNum < 1 {
				m.SetStatusMessage("Invalid line number")
				m.mode = ModeNormal
				m.inputBuffer = ""
				return m, nil
			}

			// Convert to 0-indexed and clamp
			lineNum--
			maxLine := m.buffer.LineCount() - 1
			if lineNum > maxLine {
				lineNum = maxLine
			}

			// Move to line start
			lineStart := m.buffer.LineStart(lineNum)
			if lineStart >= 0 {
				m.buffer.MoveTo(lineStart)
			}

			m.mode = ModeNormal
			m.inputBuffer = ""
			m.SetStatusMessage(fmt.Sprintf("Line %d", lineNum+1))
		}
		return m, nil

	case "esc":
		m.mode = ModeNormal
		m.inputBuffer = ""
		m.SetStatusMessage("")
		return m, nil

	case "backspace":
		if len(m.inputBuffer) > 0 {
			m.inputBuffer = m.inputBuffer[:len(m.inputBuffer)-1]
		}
		return m, nil

	default:
		// Only accept digits
		for _, r := range msg.Runes {
			if r >= '0' && r <= '9' {
				m.inputBuffer += string(r)
			}
		}
		return m, nil
	}
}

// handleSearchInput handles input in search mode.
func (m *Model) handleSearchInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		if m.inputBuffer != "" {
			m.searchQuery = m.inputBuffer
			m.findMatches()
			if len(m.searchMatches) > 0 {
				m.searchIndex = 0
				m.goToMatch(0)
				m.SetStatusMessage(fmt.Sprintf("Match %d of %d", m.searchIndex+1, len(m.searchMatches)))
			} else {
				m.SetStatusMessage("No matches found")
			}
		}
		m.mode = ModeNormal
		m.inputBuffer = ""
		return m, nil

	case "esc":
		m.mode = ModeNormal
		m.inputBuffer = ""
		m.SetStatusMessage("")
		return m, nil

	case "backspace":
		if len(m.inputBuffer) > 0 {
			m.inputBuffer = m.inputBuffer[:len(m.inputBuffer)-1]
		}
		return m, nil

	default:
		if len(msg.Runes) > 0 {
			m.inputBuffer += string(msg.Runes)
		}
		return m, nil
	}
}

// findMatches finds all occurrences of the search query in the buffer.
func (m *Model) findMatches() {
	m.searchMatches = nil
	if m.searchQuery == "" {
		return
	}

	content := m.buffer.String()
	query := m.searchQuery
	pos := 0

	for {
		idx := strings.Index(content[pos:], query)
		if idx == -1 {
			break
		}
		m.searchMatches = append(m.searchMatches, pos+idx)
		pos += idx + 1
	}
}

// goToMatch moves the cursor to the specified match index.
func (m *Model) goToMatch(index int) {
	if index < 0 || index >= len(m.searchMatches) {
		return
	}
	m.buffer.MoveTo(m.searchMatches[index])
}

// nextMatch moves to the next search match.
func (m *Model) nextMatch() {
	if len(m.searchMatches) == 0 {
		return
	}
	m.searchIndex = (m.searchIndex + 1) % len(m.searchMatches)
	m.goToMatch(m.searchIndex)
	m.SetStatusMessage(fmt.Sprintf("Match %d of %d", m.searchIndex+1, len(m.searchMatches)))
}

// prevMatch moves to the previous search match.
func (m *Model) prevMatch() {
	if len(m.searchMatches) == 0 {
		return
	}
	m.searchIndex--
	if m.searchIndex < 0 {
		m.searchIndex = len(m.searchMatches) - 1
	}
	m.goToMatch(m.searchIndex)
	m.SetStatusMessage(fmt.Sprintf("Match %d of %d", m.searchIndex+1, len(m.searchMatches)))
}

// handleReplaceInput handles input in replace mode.
func (m *Model) handleReplaceInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		if m.inputBuffer != "" {
			m.searchQuery = m.inputBuffer
			m.inputBuffer = m.replaceText
			m.inputPrompt = "Replace with: "
			m.mode = ModeReplaceConfirm
		}
		return m, nil

	case "esc":
		m.mode = ModeNormal
		m.inputBuffer = ""
		m.SetStatusMessage("")
		return m, nil

	case "backspace":
		if len(m.inputBuffer) > 0 {
			m.inputBuffer = m.inputBuffer[:len(m.inputBuffer)-1]
		}
		return m, nil

	default:
		if len(msg.Runes) > 0 {
			m.inputBuffer += string(msg.Runes)
		}
		return m, nil
	}
}

// handleReplaceConfirm handles input in replace confirm mode (single replace).
func (m *Model) handleReplaceConfirm(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		m.replaceText = m.inputBuffer
		m.findMatches()
		if len(m.searchMatches) > 0 {
			m.replaceOne()
			m.SetStatusMessage("Replaced. Press F3 for next, Ctrl+R to replace again")
		} else {
			m.SetStatusMessage("No matches found")
		}
		m.mode = ModeNormal
		m.inputBuffer = ""
		return m, nil

	case "esc":
		m.mode = ModeNormal
		m.inputBuffer = ""
		m.SetStatusMessage("")
		return m, nil

	case "backspace":
		if len(m.inputBuffer) > 0 {
			m.inputBuffer = m.inputBuffer[:len(m.inputBuffer)-1]
		}
		return m, nil

	default:
		if len(msg.Runes) > 0 {
			m.inputBuffer += string(msg.Runes)
		}
		return m, nil
	}
}

// replaceOne replaces the current match only.
func (m *Model) replaceOne() {
	if m.searchQuery == "" || len(m.searchMatches) == 0 {
		return
	}

	// Get current match position
	pos := m.searchMatches[m.searchIndex]
	queryLen := len([]rune(m.searchQuery))

	// Move to position and delete the match
	m.buffer.MoveTo(pos + queryLen)
	for i := 0; i < queryLen; i++ {
		m.buffer.Delete()
	}

	// Insert replacement
	m.buffer.InsertString(m.replaceText)

	// Record for undo
	m.history.Push(buffer.EditOperation{
		Type:     buffer.OpDelete,
		Position: pos,
		Text:     m.searchQuery,
	})
	m.history.Push(buffer.EditOperation{
		Type:     buffer.OpInsert,
		Position: pos,
		Text:     m.replaceText,
	})

	m.setModified()

	// Re-find matches and go to next
	m.findMatches()
	if len(m.searchMatches) > 0 {
		if m.searchIndex >= len(m.searchMatches) {
			m.searchIndex = 0
		}
		m.goToMatch(m.searchIndex)
	}
}

// replaceAll replaces all occurrences of searchQuery with replaceText.
func (m *Model) replaceAll() {
	if m.searchQuery == "" {
		return
	}

	content := m.buffer.String()
	newContent := strings.ReplaceAll(content, m.searchQuery, m.replaceText)

	if content != newContent {
		// Record for undo
		m.history.Push(buffer.EditOperation{
			Type:     buffer.OpDelete,
			Position: 0,
			Text:     content,
		})

		// Replace buffer content
		m.buffer = buffer.NewFromString(newContent)

		m.history.Push(buffer.EditOperation{
			Type:     buffer.OpInsert,
			Position: 0,
			Text:     newContent,
		})

		m.setModified()
	}
}

// handleReplaceAllInput handles input in replace all mode.
func (m *Model) handleReplaceAllInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		if m.inputBuffer != "" {
			m.searchQuery = m.inputBuffer
			m.inputBuffer = m.replaceText
			m.inputPrompt = "Replace all with: "
			m.mode = ModeReplaceAllConfirm
		}
		return m, nil

	case "esc":
		m.mode = ModeNormal
		m.inputBuffer = ""
		m.SetStatusMessage("")
		return m, nil

	case "backspace":
		if len(m.inputBuffer) > 0 {
			m.inputBuffer = m.inputBuffer[:len(m.inputBuffer)-1]
		}
		return m, nil

	default:
		if len(msg.Runes) > 0 {
			m.inputBuffer += string(msg.Runes)
		}
		return m, nil
	}
}

// handleReplaceAllConfirm handles input in replace all confirm mode.
func (m *Model) handleReplaceAllConfirm(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		m.replaceText = m.inputBuffer
		m.findMatches()
		count := len(m.searchMatches)
		if count > 0 {
			m.replaceAll()
			m.SetStatusMessage(fmt.Sprintf("Replaced %d occurrences", count))
		} else {
			m.SetStatusMessage("No matches found")
		}
		m.mode = ModeNormal
		m.inputBuffer = ""
		return m, nil

	case "esc":
		m.mode = ModeNormal
		m.inputBuffer = ""
		m.SetStatusMessage("")
		return m, nil

	case "backspace":
		if len(m.inputBuffer) > 0 {
			m.inputBuffer = m.inputBuffer[:len(m.inputBuffer)-1]
		}
		return m, nil

	default:
		if len(msg.Runes) > 0 {
			m.inputBuffer += string(msg.Runes)
		}
		return m, nil
	}
}

// handleOpenInput handles input in open file mode.
func (m *Model) handleOpenInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		if m.inputBuffer != "" {
			content, err := file.Load(m.inputBuffer)
			if err != nil {
				m.SetStatusMessage("Error: " + err.Error())
			} else {
				m.buffer = buffer.NewFromString(content)
				m.history = buffer.NewHistory()
				m.SetFilepath(m.inputBuffer)
				m.modified = false
				m.fileChanged = false // Reset external change flag
				m.UpdateLastSaveTime()
				m.SetStatusMessage("Opened: " + m.filename)
			}
		}
		m.mode = ModeNormal
		m.inputBuffer = ""
		return m, nil

	case "esc":
		m.mode = ModeNormal
		m.inputBuffer = ""
		m.SetStatusMessage("")
		return m, nil

	case "backspace":
		if len(m.inputBuffer) > 0 {
			m.inputBuffer = m.inputBuffer[:len(m.inputBuffer)-1]
		}
		return m, nil

	default:
		if len(msg.Runes) > 0 {
			m.inputBuffer += string(msg.Runes)
		}
		return m, nil
	}
}

// handleSaveMacroInput handles input in save macro mode.
func (m *Model) handleSaveMacroInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		if m.inputBuffer != "" {
			if err := m.macro.SaveMacro(m.inputBuffer); err != nil {
				m.SetStatusMessage("Error saving macro: " + err.Error())
			} else {
				m.SetStatusMessage("Macro saved as: " + m.inputBuffer)
			}
		}
		m.mode = ModeNormal
		m.inputBuffer = ""
		return m, nil

	case "esc":
		m.mode = ModeNormal
		m.inputBuffer = ""
		m.SetStatusMessage("")
		return m, nil

	case "backspace":
		if len(m.inputBuffer) > 0 {
			m.inputBuffer = m.inputBuffer[:len(m.inputBuffer)-1]
		}
		return m, nil

	default:
		if len(msg.Runes) > 0 {
			m.inputBuffer += string(msg.Runes)
		}
		return m, nil
	}
}

// handleLoadMacroInput handles input in load macro mode.
func (m *Model) handleLoadMacroInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		if m.inputBuffer != "" {
			if err := m.macro.LoadMacro(m.inputBuffer); err != nil {
				m.SetStatusMessage("Error loading macro: " + err.Error())
			} else {
				m.SetStatusMessage(fmt.Sprintf("Loaded macro: %s (%d keys)", m.inputBuffer, m.macro.KeyCount()))
			}
		}
		m.mode = ModeNormal
		m.inputBuffer = ""
		return m, nil

	case "esc":
		m.mode = ModeNormal
		m.inputBuffer = ""
		m.SetStatusMessage("")
		return m, nil

	case "backspace":
		if len(m.inputBuffer) > 0 {
			m.inputBuffer = m.inputBuffer[:len(m.inputBuffer)-1]
		}
		return m, nil

	case "tab":
		// Tab completion for macro names
		if names, err := ListSavedMacros(); err == nil {
			for _, name := range names {
				if strings.HasPrefix(name, m.inputBuffer) {
					m.inputBuffer = name
					break
				}
			}
		}
		return m, nil

	default:
		if len(msg.Runes) > 0 {
			m.inputBuffer += string(msg.Runes)
		}
		return m, nil
	}
}

// cutLine cuts the current line to clipboard.
func (m *Model) cutLine() {
	currentLine := m.buffer.CurrentLine()
	lineStart := m.buffer.LineStart(currentLine)
	lineEnd := m.buffer.LineEnd(currentLine)

	if lineStart < 0 || lineEnd < 0 {
		return
	}

	// Get line content
	m.clipboard = m.buffer.Line(currentLine)

	// Include newline if not last line
	deleteEnd := lineEnd
	if currentLine < m.buffer.LineCount()-1 {
		deleteEnd++
		m.clipboard += "\n"
	} else if lineStart > 0 {
		lineStart--
	}

	// Delete the line
	deletedText := m.buffer.Slice(lineStart, deleteEnd)
	m.buffer.MoveTo(lineStart)
	for i := lineStart; i < deleteEnd; i++ {
		m.buffer.DeleteForward()
	}

	m.history.Push(buffer.EditOperation{
		Type:     buffer.OpDelete,
		Position: lineStart,
		Text:     deletedText,
	})

	m.setModified()
	m.SetStatusMessage("Line cut to clipboard")
}

// copyLine copies the current line to clipboard (nano Alt+6 style).
func (m *Model) copyLine() {
	currentLine := m.buffer.CurrentLine()
	m.clipboard = m.buffer.Line(currentLine)
	// Include newline for consistency with cut
	if currentLine < m.buffer.LineCount()-1 {
		m.clipboard += "\n"
	}
	m.SetStatusMessage("Copied 1 line")
}

// deleteWordLeft deletes word to the left of cursor.
func (m *Model) deleteWordLeft() {
	if m.buffer.CursorPos() == 0 {
		return
	}

	startPos := m.buffer.CursorPos()

	// Move to start of word
	m.moveWordLeft()
	endPos := m.buffer.CursorPos()

	if startPos == endPos {
		return
	}

	// Delete from current position to original position
	deletedText := m.buffer.Slice(endPos, startPos)
	for i := endPos; i < startPos; i++ {
		m.buffer.DeleteForward()
	}

	m.history.Push(buffer.EditOperation{
		Type:     buffer.OpDelete,
		Position: endPos,
		Text:     deletedText,
	})
	m.setModified()
}

// deleteWordRight deletes word to the right of cursor.
func (m *Model) deleteWordRight() {
	if m.buffer.CursorPos() >= m.buffer.Len() {
		return
	}

	startPos := m.buffer.CursorPos()

	// Calculate end position (word right)
	pos := startPos
	length := m.buffer.Len()

	// Skip current word
	for pos < length {
		r := m.buffer.RuneAt(pos)
		if r == ' ' || r == '\n' || r == '\t' {
			break
		}
		pos++
	}
	// Skip whitespace
	for pos < length {
		r := m.buffer.RuneAt(pos)
		if r != ' ' && r != '\n' && r != '\t' {
			break
		}
		pos++
	}

	if pos == startPos {
		return
	}

	// Delete from start to calculated end
	deletedText := m.buffer.Slice(startPos, pos)
	for i := startPos; i < pos; i++ {
		m.buffer.DeleteForward()
	}

	m.history.Push(buffer.EditOperation{
		Type:     buffer.OpDelete,
		Position: startPos,
		Text:     deletedText,
	})
	m.setModified()
}

// ensureCursorVisible ensures the cursor is visible in the viewport.
func (m *Model) ensureCursorVisible() {
	currentLine := m.buffer.CurrentLine()
	editorHeight := m.height - 3 // header + status + help

	if currentLine < m.viewportTopLine {
		m.viewportTopLine = currentLine
	} else if currentLine >= m.viewportTopLine+editorHeight {
		m.viewportTopLine = currentLine - editorHeight + 1
	}
}

// paste pastes content from clipboard.
func (m *Model) paste() {
	if m.clipboard == "" {
		m.SetStatusMessage("Clipboard is empty")
		return
	}

	pos := m.buffer.CursorPos()
	m.buffer.InsertString(m.clipboard)

	m.history.Push(buffer.EditOperation{
		Type:     buffer.OpInsert,
		Position: pos,
		Text:     m.clipboard,
	})

	m.setModified()
	m.SetStatusMessage("Pasted")
}

// moveWordLeft moves cursor to the beginning of the previous word.
func (m *Model) moveWordLeft() {
	pos := m.buffer.CursorPos()
	if pos == 0 {
		return
	}

	// Skip any spaces before current position
	for pos > 0 && isSpace(m.buffer.RuneAt(pos-1)) {
		pos--
	}

	// Move to start of word
	for pos > 0 && !isSpace(m.buffer.RuneAt(pos-1)) {
		pos--
	}

	m.buffer.MoveTo(pos)
}

// moveWordRight moves cursor to the beginning of the next word.
func (m *Model) moveWordRight() {
	pos := m.buffer.CursorPos()
	length := m.buffer.Len()

	// Skip current word
	for pos < length && !isSpace(m.buffer.RuneAt(pos)) {
		pos++
	}

	// Skip spaces
	for pos < length && isSpace(m.buffer.RuneAt(pos)) {
		pos++
	}

	m.buffer.MoveTo(pos)
}

// isSpace checks if a rune is a whitespace character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

// renderLineWithSyntax renders a line with syntax highlighting.
// lineNum is used for token caching.
func (m *Model) renderLineWithSyntax(lineNum int, line string) string {
	// Initialize highlighter if needed
	if m.highlighter == nil {
		lang := syntax.DetectLanguage(m.filename)
		if lang == nil {
			return editorStyle.Render(line)
		}
		m.highlighter = syntax.New(lang)
	}

	// Use cached highlighting
	tokens := m.highlighter.HighlightLine(lineNum, line)

	var result strings.Builder
	for _, token := range tokens {
		style := getSyntaxStyle(token.Type)
		result.WriteString(style.Render(token.Text))
	}
	return result.String()
}

// updateHighlighter updates the highlighter when filename changes.
func (m *Model) updateHighlighter() {
	lang := syntax.DetectLanguage(m.filename)
	if lang != nil {
		m.highlighter = syntax.New(lang)
	} else {
		m.highlighter = nil
	}
}

// invalidateSyntaxCache invalidates syntax cache for modified lines.
func (m *Model) invalidateSyntaxCache(lineNum int) {
	if m.highlighter != nil {
		m.highlighter.InvalidateFromLine(lineNum)
	}
}

// setModified sets the modified flag and invalidates syntax cache.
func (m *Model) setModified() {
	m.modified = true
	currentLine := m.buffer.CurrentLine()
	m.invalidateSyntaxCache(currentLine)
}

// getSyntaxStyle returns the lipgloss style for a token type.
func getSyntaxStyle(tokenType syntax.TokenType) lipgloss.Style {
	switch tokenType {
	case syntax.TokenKeyword:
		return syntaxKeywordStyle
	case syntax.TokenType_:
		return syntaxTypeStyle
	case syntax.TokenString:
		return syntaxStringStyle
	case syntax.TokenNumber:
		return syntaxNumberStyle
	case syntax.TokenComment:
		return syntaxCommentStyle
	case syntax.TokenOperator:
		return syntaxOperatorStyle
	case syntax.TokenFunction:
		return syntaxFunctionStyle
	case syntax.TokenVariable:
		return syntaxVariableStyle
	case syntax.TokenConstant:
		return syntaxConstantStyle
	case syntax.TokenBuiltin:
		return syntaxBuiltinStyle
	default:
		return editorStyle
	}
}

// renderLineWithSearchMatches renders a line with search matches highlighted.
func (m *Model) renderLineWithSearchMatches(line, query string) string {
	var result strings.Builder
	pos := 0
	for {
		idx := strings.Index(line[pos:], query)
		if idx == -1 {
			result.WriteString(editorStyle.Render(line[pos:]))
			break
		}
		// Text before match
		if idx > 0 {
			result.WriteString(editorStyle.Render(line[pos : pos+idx]))
		}
		// Highlighted match
		result.WriteString(searchMatchStyle.Render(query))
		pos = pos + idx + len(query)
	}
	return result.String()
}

// wrapLine wraps a line to fit within the given width.
func wrapLine(line string, width int) string {
	runes := []rune(line)
	if len(runes) <= width {
		return line
	}
	return string(runes[:width])
}

// getIndent extracts leading whitespace from a line.
func getIndent(line string) string {
	var indent strings.Builder
	for _, r := range line {
		if r == ' ' || r == '\t' {
			indent.WriteRune(r)
		} else {
			break
		}
	}
	return indent.String()
}

// formatSize formats byte size to human readable string.
func formatSize(bytes int) string {
	if bytes < 1024 {
		return fmt.Sprintf("%d B", bytes)
	} else if bytes < 1024*1024 {
		return fmt.Sprintf("%.1f KB", float64(bytes)/1024)
	} else {
		return fmt.Sprintf("%.1f MB", float64(bytes)/(1024*1024))
	}
}

// detectLanguage detects programming language from filename.
func detectLanguage(filename string) string {
	// Extract extension
	ext := ""
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			ext = filename[i:]
			break
		}
	}

	switch ext {
	case ".go":
		return "Go"
	case ".py":
		return "Python"
	case ".js":
		return "JavaScript"
	case ".ts":
		return "TypeScript"
	case ".json":
		return "JSON"
	case ".yaml", ".yml":
		return "YAML"
	case ".toml":
		return "TOML"
	case ".md":
		return "Markdown"
	case ".html", ".htm":
		return "HTML"
	case ".css":
		return "CSS"
	case ".c", ".h":
		return "C"
	case ".cpp", ".hpp", ".cc":
		return "C++"
	case ".rs":
		return "Rust"
	case ".java":
		return "Java"
	case ".rb":
		return "Ruby"
	case ".php":
		return "PHP"
	case ".sh", ".bash":
		return "Shell"
	case ".sql":
		return "SQL"
	case ".xml":
		return "XML"
	case ".txt":
		return "Text"
	default:
		return ""
	}
}

// toggleSelection toggles selection mode.
func (m *Model) toggleSelection() {
	if m.selecting {
		m.clearSelection()
		m.SetStatusMessage("Selection ended")
	} else {
		m.selecting = true
		m.selectionStart = m.buffer.CursorPos()
		m.selectionEnd = m.selectionStart
		m.SetStatusMessage("Selection started")
	}
}

// startSelection starts selection if not already selecting.
func (m *Model) startSelection() {
	if !m.selecting {
		m.selecting = true
		m.selectionStart = m.buffer.CursorPos()
		m.selectionEnd = m.selectionStart
	}
}

// updateSelection updates the selection end to current cursor position.
func (m *Model) updateSelection() {
	if m.selecting {
		m.selectionEnd = m.buffer.CursorPos()
	}
}

// clearSelection clears the current selection.
func (m *Model) clearSelection() {
	m.selecting = false
	m.selectionStart = 0
	m.selectionEnd = 0
}

// renderLineWithSelection renders a line with selection highlighting.
func (m *Model) renderLineWithSelection(runes []rune, lineStart, lineEnd, selStart, selEnd, cursorCol int) string {
	var result strings.Builder

	for i, r := range runes {
		charPos := lineStart + i
		isSelected := charPos >= selStart && charPos < selEnd
		isCursor := i == cursorCol

		if isCursor {
			result.WriteString(lipgloss.NewStyle().Reverse(true).Render(string(r)))
		} else if isSelected {
			result.WriteString(selectionStyle.Render(string(r)))
		} else {
			result.WriteString(editorStyle.Render(string(r)))
		}
	}

	// Add cursor at end if needed
	if cursorCol >= len(runes) && cursorCol >= 0 {
		result.WriteString("█")
	}

	return result.String()
}

// selectAll selects all text in the buffer.
func (m *Model) selectAll() {
	m.selecting = true
	m.selectionStart = 0
	m.selectionEnd = m.buffer.Len()
	m.buffer.MoveToEnd()
	m.SetStatusMessage("All text selected")
}

// getSelectionBounds returns the start and end of selection in order.
func (m *Model) getSelectionBounds() (int, int) {
	start, end := m.selectionStart, m.selectionEnd
	if start > end {
		start, end = end, start
	}
	return start, end
}

// copySelection copies selected text to clipboard.
func (m *Model) copySelection() {
	if !m.selecting {
		return
	}
	start, end := m.getSelectionBounds()
	if start == end {
		return
	}
	m.clipboard = m.buffer.Slice(start, end)
	m.clearSelection()
	m.SetStatusMessage("Copied to clipboard")
}

// cutSelection cuts selected text to clipboard.
func (m *Model) cutSelection() {
	if !m.selecting || m.readonly {
		if m.readonly {
			m.SetStatusMessage("File is read-only")
		}
		return
	}
	start, end := m.getSelectionBounds()
	if start == end {
		return
	}

	// Copy to clipboard
	m.clipboard = m.buffer.Slice(start, end)

	// Delete selection
	m.buffer.MoveTo(end)
	for i := start; i < end; i++ {
		m.buffer.Delete()
	}

	// Record for undo
	m.history.Push(buffer.EditOperation{
		Type:     buffer.OpDelete,
		Position: start,
		Text:     m.clipboard,
	})

	m.clearSelection()
	m.setModified()
	m.SetStatusMessage("Cut to clipboard")
}

// View renders the UI.
func (m *Model) View() string {
	if m.quitting {
		return ""
	}

	if m.width == 0 || m.height == 0 {
		return "Initializing..."
	}

	var b strings.Builder

	// Tab bar (if multiple tabs)
	if m.showTabs && m.TabCount() > 1 {
		b.WriteString(m.renderTabBar())
		b.WriteString("\n")
	}

	// Header
	b.WriteString(m.renderHeader())
	b.WriteString("\n")

	// Editor area (with split support)
	if m.IsSplit() {
		b.WriteString(m.renderSplitEditor())
	} else {
		b.WriteString(m.renderEditor())
	}

	// Status bar
	b.WriteString(m.renderStatusBar())

	// Help bar
	b.WriteString(m.renderHelpBar())

	return b.String()
}

// renderHeader renders the top header bar.
func (m *Model) renderHeader() string {
	// Logo
	logo := " 𒄑 GESH "

	// Filename with modified/readonly indicator
	filename := m.filename
	if m.readonly {
		filename += " [RO]"
	} else if m.modified {
		filename += " *"
	}

	// Encoding and line ending info
	rightInfo := fmt.Sprintf("%s %s", m.encoding, m.lineEnding)

	// Calculate padding
	contentWidth := len(logo) + 3 + len(filename) + len(rightInfo)
	padding := m.width - contentWidth
	if padding < 0 {
		padding = 0
	}

	line := headerStyle.Render(logo) +
		headerStyle.Render(" │ ") +
		headerStyle.Render(filename) +
		headerStyle.Render(strings.Repeat(" ", padding)) +
		headerStyle.Render(rightInfo)

	return line
}

// renderTabBar renders the tab bar showing all open tabs.
func (m *Model) renderTabBar() string {
	var b strings.Builder

	activeIdx := m.ActiveTabIndex()
	tabs := m.tabs.Tabs()

	for i, tab := range tabs {
		// Tab number (1-based)
		tabNum := fmt.Sprintf(" %d ", i+1)

		// Filename (truncate if needed)
		name := tab.filename
		if len(name) > 15 {
			name = name[:12] + "..."
		}
		if tab.modified {
			name += "*"
		}
		name = " " + name + " "

		// Style based on active/inactive
		if i == activeIdx {
			// Active tab: highlighted
			b.WriteString(styles.TabActiveStyle.Render(tabNum + name))
		} else {
			// Inactive tab
			b.WriteString(styles.TabInactiveStyle.Render(tabNum + name))
		}

		// Separator
		if i < len(tabs)-1 {
			b.WriteString(" ")
		}
	}

	// Fill remaining space
	content := b.String()
	// Remove ANSI codes for length calculation is complex; just pad to width
	padding := m.width - len(content)/2 // Rough estimate
	if padding > 0 {
		b.WriteString(strings.Repeat(" ", padding))
	}

	return b.String()
}

// renderSplitEditor renders the editor area with split panes.
func (m *Model) renderSplitEditor() string {
	var b strings.Builder

	// Calculate editor area height
	editorHeight := m.height - 3
	if m.showTabs && m.TabCount() > 1 {
		editorHeight--
	}
	if editorHeight < 1 {
		editorHeight = 1
	}

	// Calculate pane dimensions
	m.split.CalculatePaneDimensions(m.width, editorHeight)

	switch m.split.Direction() {
	case SplitHorizontal:
		// Side by side - render line by line
		panes := m.split.Panes()
		leftPane := panes[0]
		rightPane := panes[1]

		// Get content for each pane
		leftLines := m.renderPaneLines(leftPane, leftPane.height)
		rightLines := m.renderPaneLines(rightPane, rightPane.height)

		// Combine line by line
		for i := 0; i < editorHeight; i++ {
			leftLine := ""
			rightLine := ""

			if i < len(leftLines) {
				leftLine = leftLines[i]
			}
			if i < len(rightLines) {
				rightLine = rightLines[i]
			}

			// Pad left pane to width
			leftLine = padOrTruncate(leftLine, leftPane.width)

			// Add separator (│) with highlighting for active pane
			separator := "│"
			if m.split.ActivePaneIndex() == 0 {
				separator = styles.TabActiveStyle.Render("│")
			} else {
				separator = styles.TabInactiveStyle.Render("│")
			}

			b.WriteString(leftLine)
			b.WriteString(separator)
			b.WriteString(rightLine)
			b.WriteString("\n")
		}

	case SplitVertical:
		// Stacked - render top pane, separator, bottom pane
		panes := m.split.Panes()
		topPane := panes[0]
		bottomPane := panes[1]

		// Render top pane
		topLines := m.renderPaneLines(topPane, topPane.height)
		for _, line := range topLines {
			b.WriteString(padOrTruncate(line, m.width))
			b.WriteString("\n")
		}

		// Separator line
		separator := strings.Repeat("─", m.width)
		if m.split.ActivePaneIndex() == 0 {
			b.WriteString(styles.TabActiveStyle.Render(separator))
		} else {
			b.WriteString(styles.TabInactiveStyle.Render(separator))
		}
		b.WriteString("\n")

		// Render bottom pane
		bottomLines := m.renderPaneLines(bottomPane, bottomPane.height)
		for _, line := range bottomLines {
			b.WriteString(padOrTruncate(line, m.width))
			b.WriteString("\n")
		}
	}

	return b.String()
}

// renderPaneLines renders content lines for a pane.
func (m *Model) renderPaneLines(pane *Pane, height int) []string {
	lines := make([]string, 0, height)

	// Get the tab this pane is showing
	if pane.tabIndex >= m.tabs.Count() {
		pane.tabIndex = 0
	}
	tab := m.tabs.Tabs()[pane.tabIndex]

	// Use pane's viewport
	topLine := pane.viewportTopLine
	lineCount := tab.buffer.LineCount()

	for i := 0; i < height; i++ {
		lineNum := topLine + i
		if lineNum >= lineCount {
			// Empty line after file content
			lines = append(lines, "")
			continue
		}

		var lineBuilder strings.Builder

		// Line number
		if m.showLineNumbers {
			// Check if this is the current line (cursor is here)
			isCurrentLine := (pane.tabIndex == m.tabs.ActiveIndex() &&
				pane == m.split.ActivePane() &&
				lineNum == tab.buffer.CurrentLine())

			if isCurrentLine {
				lineBuilder.WriteString(lineNumberStyle.Render(fmt.Sprintf("→%3d ", lineNum+1)))
			} else {
				lineBuilder.WriteString(lineNumberStyle.Render(fmt.Sprintf(" %3d ", lineNum+1)))
			}
		}

		// Line content
		lineContent := tab.buffer.Line(lineNum)
		lineBuilder.WriteString(lineContent)

		lines = append(lines, lineBuilder.String())
	}

	return lines
}

// padOrTruncate ensures a string is exactly the given width.
func padOrTruncate(s string, width int) string {
	// Simple rune-based length (ignoring ANSI codes for now)
	runes := []rune(s)
	if len(runes) >= width {
		return string(runes[:width])
	}
	return s + strings.Repeat(" ", width-len(runes))
}

// renderEditor renders the main editor area.
func (m *Model) renderEditor() string {
	var b strings.Builder

	// Calculate visible lines (height minus header, status, help, and optionally tab bar)
	visibleLines := m.height - 3
	if m.showTabs && m.TabCount() > 1 {
		visibleLines-- // Account for tab bar
	}
	if visibleLines < 1 {
		visibleLines = 1
	}

	lineCount := m.buffer.LineCount()
	cursorLine := m.buffer.CurrentLine()
	cursorCol := m.buffer.CurrentColumn()

	// Adjust viewport to keep cursor visible with scroll padding
	// But NOT when user is scrolling with mouse - let them scroll freely
	if !m.mouseScrolling {
		scrollPadding := 5
		if scrollPadding >= visibleLines/2 {
			scrollPadding = visibleLines / 3
		}

		if cursorLine < m.viewportTopLine+scrollPadding {
			m.viewportTopLine = cursorLine - scrollPadding
			if m.viewportTopLine < 0 {
				m.viewportTopLine = 0
			}
		}
		if cursorLine >= m.viewportTopLine+visibleLines-scrollPadding {
			m.viewportTopLine = cursorLine - visibleLines + scrollPadding + 1
			if m.viewportTopLine > lineCount-visibleLines {
				m.viewportTopLine = lineCount - visibleLines
			}
			if m.viewportTopLine < 0 {
				m.viewportTopLine = 0
			}
		}
	}

	for i := 0; i < visibleLines; i++ {
		lineNum := m.viewportTopLine + i

		if lineNum < lineCount {
			// Line number with current line marker (if enabled)
			if m.showLineNumbers {
				var lineNumStr string
				// Calculate width needed for line numbers
				maxLineNum := lineCount
				numWidth := len(fmt.Sprintf("%d", maxLineNum))
				if numWidth < 3 {
					numWidth = 3
				}
				
				if lineNum == cursorLine {
					// Current line with arrow marker
					lineNumStr = lineNumberStyle.Render(fmt.Sprintf("→%*d", numWidth, lineNum+1))
				} else {
					lineNumStr = lineNumberStyle.Render(fmt.Sprintf(" %*d", numWidth, lineNum+1))
				}
				b.WriteString(lineNumStr)
				b.WriteString(" │ ")
			}

			// Line content
			lineContent := m.buffer.Line(lineNum)

			// Calculate available width for text
			textWidth := m.width
			if m.showLineNumbers {
				textWidth -= 7 // "→123 │ " = 7 chars
			}
			if textWidth < 10 {
				textWidth = 10
			}

			// Word wrap if enabled
			if m.wordWrap && len(lineContent) > textWidth {
				lineContent = wrapLine(lineContent, textWidth)
			}

			// Calculate line start position in buffer
			lineStart := m.buffer.LineStart(lineNum)
			lineEnd := lineStart + len([]rune(m.buffer.Line(lineNum)))

			// Get selection bounds if selecting
			selStart, selEnd := 0, 0
			hasSelection := false
			if m.selecting {
				selStart, selEnd = m.getSelectionBounds()
				hasSelection = selStart != selEnd
			}

			// Render line with selection and cursor
			runes := []rune(lineContent)
			if lineNum == cursorLine {
				// Cursor line - render with cursor
				if hasSelection {
					b.WriteString(m.renderLineWithSelection(runes, lineStart, lineEnd, selStart, selEnd, cursorCol))
				} else if cursorCol >= len(runes) {
					b.WriteString(editorStyle.Render(lineContent))
					b.WriteString("█")
				} else {
					before := string(runes[:cursorCol])
					cursor := string(runes[cursorCol])
					after := string(runes[cursorCol+1:])
					b.WriteString(editorStyle.Render(before))
					b.WriteString(lipgloss.NewStyle().Reverse(true).Render(cursor))
					b.WriteString(editorStyle.Render(after))
				}
			} else if hasSelection && lineEnd > selStart && lineStart < selEnd {
				// Line has selection
				b.WriteString(m.renderLineWithSelection(runes, lineStart, lineEnd, selStart, selEnd, -1))
			} else if m.searchQuery != "" && strings.Contains(lineContent, m.searchQuery) {
				// Line has search matches
				b.WriteString(m.renderLineWithSearchMatches(lineContent, m.searchQuery))
			} else if m.syntaxHighlighting {
				// Syntax highlighting with cache
				b.WriteString(m.renderLineWithSyntax(lineNum, lineContent))
			} else {
				b.WriteString(editorStyle.Render(lineContent))
			}
		} else {
			// Empty line indicator (after end of file)
			if m.showLineNumbers {
				// Calculate width for alignment
				maxLineNum := lineCount
				numWidth := len(fmt.Sprintf("%d", maxLineNum))
				if numWidth < 3 {
					numWidth = 3
				}
				lineNumStr := lineNumberStyle.Render(fmt.Sprintf(" %*s", numWidth, "~"))
				b.WriteString(lineNumStr)
				b.WriteString(" │ ")
			}
		}

		b.WriteString("\n")
	}

	return b.String()
}

// renderStatusBar renders the status bar.
func (m *Model) renderStatusBar() string {
	// Position info
	line := m.buffer.CurrentLine() + 1
	col := m.buffer.CurrentColumn() + 1
	posInfo := fmt.Sprintf(" Ln %d, Col %d", line, col)

	// File info with size
	lineCount := m.buffer.LineCount()
	size := m.buffer.Len()
	sizeStr := formatSize(size)
	fileInfo := fmt.Sprintf(" │ %d lines │ %s", lineCount, sizeStr)

	// Language detection
	lang := detectLanguage(m.filename)
	if lang != "" {
		fileInfo += " │ " + lang
	}

	// Status message or mode indicator
	var modeInfo string
	if m.statusMessage != "" {
		modeInfo = " │ " + m.statusMessage
	}

	// Right-aligned mode indicator
	rightInfo := "INS"
	if m.overwriteMode {
		rightInfo = "OVR"
	}
	if m.macro != nil && m.macro.IsRecording() {
		rightInfo = "REC"
	}

	// Calculate content - use rune count for proper width calculation
	leftContent := posInfo + fileInfo + modeInfo
	leftLen := len([]rune(leftContent))
	rightLen := len([]rune(rightInfo))
	padding := m.width - leftLen - rightLen - 2
	if padding < 0 {
		padding = 0
	}

	return statusStyle.Render(leftContent + strings.Repeat(" ", padding) + rightInfo + " ") + "\n"
}

// renderHelpBar renders the bottom help bar (nano style).
func (m *Model) renderHelpBar() string {
	switch m.mode {
	case ModeQuit:
		content := " [Y] Save & Exit  [N] Discard & Exit  [C] Cancel"
		padding := m.width - len(content)
		if padding < 0 {
			padding = 0
		}
		return helpStyle.Render(content + strings.Repeat(" ", padding))

	case ModeSaveAs, ModeGoto, ModeSearch, ModeReplace, ModeReplaceConfirm, ModeReplaceAll, ModeReplaceAllConfirm, ModeOpen, ModeSaveMacro, ModeLoadMacro:
		// Show input prompt
		prompt := " " + m.inputPrompt + m.inputBuffer + "█"
		padding := m.width - len([]rune(prompt))
		if padding < 0 {
			padding = 0
		}
		return helpStyle.Render(prompt + strings.Repeat(" ", padding))

	default:
		// Nano style help - always visible, two lines
		line1 := "^G Help  ^O WriteOut  ^W Where Is  ^K Cut      ^C Cur Pos  ^X Exit"
		line2 := "^R Read File  ^\\ Replace  ^U Paste  ^Y Prev Pg  ^V Next Pg  M-U Undo"

		// Pad lines to width
		pad1 := m.width - len(line1)
		pad2 := m.width - len(line2)
		if pad1 < 0 {
			pad1 = 0
		}
		if pad2 < 0 {
			pad2 = 0
		}

		return helpStyle.Render(line1+strings.Repeat(" ", pad1)) + "\n" +
			helpStyle.Render(line2+strings.Repeat(" ", pad2))
	}
}

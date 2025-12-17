package app

import (
	"fmt"
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
}

// Init initializes the model.
func (m *Model) Init() tea.Cmd {
	return nil
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
		if msg.Action == tea.MouseActionPress {
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

			// Move cursor
			lineStart := m.buffer.LineStart(clickedLine)
			m.buffer.MoveTo(lineStart + clickedCol)
			m.clearSelection()
		}

	case tea.MouseButtonWheelUp:
		// Scroll up
		m.viewportTopLine -= 3
		if m.viewportTopLine < 0 {
			m.viewportTopLine = 0
		}

	case tea.MouseButtonWheelDown:
		// Scroll down
		maxTop := m.buffer.LineCount() - (m.height - 3)
		if maxTop < 0 {
			maxTop = 0
		}
		m.viewportTopLine += 3
		if m.viewportTopLine > maxTop {
			m.viewportTopLine = maxTop
		}
	}

	return m, nil
}

// handleKeyMsg processes keyboard input.
func (m *Model) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Handle quit confirmation mode
	if m.mode == ModeQuit {
		switch msg.String() {
		case "y", "Y":
			m.quitting = true
			return m, tea.Quit
		case "n", "N", "esc":
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

	// Normal mode key handling
	switch msg.String() {
	case "ctrl+x":
		// If selecting, cut selection
		if m.selecting {
			m.cutSelection()
			return m, nil
		}
		// Otherwise exit
		if m.modified {
			m.mode = ModeQuit
			m.SetStatusMessage("Save changes? (Y)es, (N)o, (Esc)Cancel")
			return m, nil
		}
		m.quitting = true
		return m, tea.Quit

	case "ctrl+c":
		if m.selecting {
			m.copySelection()
			return m, nil
		}
		m.quitting = true
		return m, tea.Quit

	case "ctrl+alt+n":
		if m.modified {
			m.SetStatusMessage("Save first or force quit with Ctrl+C")
			return m, nil
		}
		m.buffer = buffer.New()
		m.history = buffer.NewHistory()
		m.filename = "[New File]"
		m.filepath = ""
		m.modified = false
		m.SetStatusMessage("New file")
		return m, nil

	case "ctrl+s":
		return m.saveFile()

	case "ctrl+shift+s":
		// Save As - always prompt for filename
		m.mode = ModeSaveAs
		m.inputBuffer = m.filepath
		m.inputPrompt = "Save as: "
		return m, nil

	case "ctrl+g":
		m.mode = ModeGoto
		m.inputBuffer = ""
		m.inputPrompt = fmt.Sprintf("Go to line [1-%d]: ", m.buffer.LineCount())
		return m, nil

	case "ctrl+w":
		m.mode = ModeSearch
		m.inputBuffer = m.searchQuery // Pre-fill with last search
		m.inputPrompt = "Search: "
		return m, nil

	case "f3":
		m.nextMatch()
		return m, nil

	case "shift+f3":
		m.prevMatch()
		return m, nil

	case "ctrl+r":
		m.mode = ModeReplace
		m.inputBuffer = m.searchQuery
		m.inputPrompt = "Replace: "
		return m, nil

	case "ctrl+shift+r":
		m.mode = ModeReplaceAll
		m.inputBuffer = m.searchQuery
		m.inputPrompt = "Replace all: "
		return m, nil

	case "ctrl+o":
		m.mode = ModeOpen
		m.inputBuffer = ""
		m.inputPrompt = "Open file: "
		return m, nil

	case "ctrl+u":
		m.cutLine()
		return m, nil

	case "ctrl+v":
		m.paste()
		return m, nil

	case "ctrl+left":
		m.moveWordLeft()
		return m, nil

	case "ctrl+right":
		m.moveWordRight()
		return m, nil

	case "ctrl+ ":
		m.toggleSelection()
		return m, nil

	// Selection with shift
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

	// Navigation
	case "up", "ctrl+p":
		m.clearSelection()
		m.moveCursorUp()
	case "down", "ctrl+n":
		m.clearSelection()
		m.moveCursorDown()
	case "left", "ctrl+b":
		m.clearSelection()
		m.buffer.MoveLeft()
	case "right", "ctrl+f":
		m.clearSelection()
		m.buffer.MoveRight()
	case "home":
		m.clearSelection()
		m.moveToLineStart()
	case "ctrl+a":
		// Double Ctrl+A = select all
		now := time.Now().UnixMilli()
		if now-m.lastCtrlATime < 500 {
			m.selectAll()
			m.lastCtrlATime = 0
		} else {
			m.clearSelection()
			m.moveToLineStart()
			m.lastCtrlATime = now
		}
	case "end", "ctrl+e":
		m.clearSelection()
		m.moveToLineEnd()
	case "ctrl+home":
		m.clearSelection()
		m.buffer.MoveToStart()
	case "ctrl+end":
		m.clearSelection()
		m.buffer.MoveToEnd()
	case "pgup":
		m.clearSelection()
		m.pageUp()
	case "pgdown":
		m.clearSelection()
		m.pageDown()

	// Editing (check readonly)
	case "backspace":
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
			m.modified = true
		}
	case "delete":
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
			m.modified = true
		}
	case "enter":
		if m.readonly {
			m.SetStatusMessage("File is read-only")
			return m, nil
		}
		// Get current line's indentation
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
		m.modified = true
	case "tab":
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
		m.modified = true

	case "ctrl+z":
		m.undo()
		return m, nil

	case "ctrl+y":
		m.redo()
		return m, nil

	case "ctrl+k":
		if m.readonly {
			m.SetStatusMessage("File is read-only")
			return m, nil
		}
		m.deleteLine()
		return m, nil

	default:
		// Insert printable characters
		if len(msg.Runes) > 0 {
			if m.readonly {
				m.SetStatusMessage("File is read-only")
				return m, nil
			}
			pos := m.buffer.CursorPos()
			text := string(msg.Runes)
			for _, r := range msg.Runes {
				m.buffer.Insert(r)
			}
			m.history.Push(buffer.EditOperation{
				Type:     buffer.OpInsert,
				Position: pos,
				Text:     text,
			})
			m.modified = true
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

	m.modified = true
	m.SetStatusMessage("Line deleted")
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

	m.modified = true

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

		m.modified = true
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

	m.modified = true
	m.SetStatusMessage("Line cut to clipboard")
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

	m.modified = true
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
func (m *Model) renderLineWithSyntax(line string) string {
	lang := syntax.DetectLanguage(m.filename)
	if lang == nil {
		return editorStyle.Render(line)
	}

	highlighter := syntax.New(lang)
	tokens := highlighter.Highlight(line)

	var result strings.Builder
	for _, token := range tokens {
		style := getSyntaxStyle(token.Type)
		result.WriteString(style.Render(token.Text))
	}
	return result.String()
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
	m.modified = true
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

	// Header
	b.WriteString(m.renderHeader())
	b.WriteString("\n")

	// Editor area
	b.WriteString(m.renderEditor())

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

	// Calculate padding
	rightInfo := "UTF-8 LF"
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

// renderEditor renders the main editor area.
func (m *Model) renderEditor() string {
	var b strings.Builder

	// Calculate visible lines (height minus header, status, help)
	visibleLines := m.height - 3
	if visibleLines < 1 {
		visibleLines = 1
	}

	lineCount := m.buffer.LineCount()
	cursorLine := m.buffer.CurrentLine()
	cursorCol := m.buffer.CurrentColumn()

	// Adjust viewport to keep cursor visible with scroll padding
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

	for i := 0; i < visibleLines; i++ {
		lineNum := m.viewportTopLine + i

		if lineNum < lineCount {
			// Line number with current line marker (if enabled)
			if m.showLineNumbers {
				var lineNumStr string
				if lineNum == cursorLine {
					lineNumStr = lineNumberStyle.Render(fmt.Sprintf("→%d", lineNum+1))
				} else {
					lineNumStr = lineNumberStyle.Render(fmt.Sprintf(" %d", lineNum+1))
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
				// Syntax highlighting
				b.WriteString(m.renderLineWithSyntax(lineContent))
			} else {
				b.WriteString(editorStyle.Render(lineContent))
			}
		} else {
			// Empty line indicator
			if m.showLineNumbers {
				lineNumStr := lineNumberStyle.Render("~")
				b.WriteString(lineNumStr)
				b.WriteString(" │")
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

	// Right-aligned mode
	rightInfo := "INS "

	// Calculate content
	leftContent := posInfo + fileInfo + modeInfo
	padding := m.width - len(leftContent) - len(rightInfo)
	if padding < 0 {
		padding = 0
	}

	return statusStyle.Render(leftContent + strings.Repeat(" ", padding) + rightInfo)
}

// renderHelpBar renders the bottom help bar.
func (m *Model) renderHelpBar() string {
	var helps []string

	switch m.mode {
	case ModeQuit:
		helps = []string{
			helpKeyStyle.Render("Y") + helpStyle.Render(" Yes"),
			helpKeyStyle.Render("N") + helpStyle.Render(" No"),
			helpKeyStyle.Render("Esc") + helpStyle.Render(" Cancel"),
		}
	case ModeSaveAs, ModeGoto, ModeSearch, ModeReplace, ModeReplaceConfirm, ModeReplaceAll, ModeReplaceAllConfirm, ModeOpen:
		// Show input prompt
		prompt := m.inputPrompt + m.inputBuffer + "█"
		padding := m.width - len(prompt) - 2
		if padding < 0 {
			padding = 0
		}
		return helpStyle.Render(" " + prompt + strings.Repeat(" ", padding))
	default:
		helps = []string{
			helpKeyStyle.Render("^X") + helpStyle.Render(" Exit"),
			helpKeyStyle.Render("^S") + helpStyle.Render(" Save"),
			helpKeyStyle.Render("^O") + helpStyle.Render(" Open"),
			helpKeyStyle.Render("^W") + helpStyle.Render(" Search"),
			helpKeyStyle.Render("^R") + helpStyle.Render(" Replace"),
			helpKeyStyle.Render("^U") + helpStyle.Render(" Cut"),
			helpKeyStyle.Render("^V") + helpStyle.Render(" Paste"),
		}
	}

	content := " " + strings.Join(helps, "  ")
	padding := m.width - len(content)
	if padding < 0 {
		padding = 0
	}

	return helpStyle.Render(content + strings.Repeat(" ", padding))
}

package app

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/KilimcininKorOglu/gesh/internal/file"
)

// Styles for the UI components
var (
	headerStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#16213e")).
			Foreground(lipgloss.Color("#e94560")).
			Bold(true)

	statusStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#0f3460")).
			Foreground(lipgloss.Color("#eaeaea"))

	helpStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#16213e")).
			Foreground(lipgloss.Color("#a0a0c0"))

	helpKeyStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#16213e")).
			Foreground(lipgloss.Color("#e94560")).
			Bold(true)

	lineNumberStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4a4a6a")).
			Width(4).
			Align(lipgloss.Right)

	editorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#eaeaea"))
)

// Init initializes the model.
func (m *Model) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyMsg(msg)
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
		return m, nil
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

	// Normal mode key handling
	switch msg.String() {
	case "ctrl+x":
		if m.modified {
			m.mode = ModeQuit
			m.SetStatusMessage("Save changes? (Y)es, (N)o, (Esc)Cancel")
			return m, nil
		}
		m.quitting = true
		return m, tea.Quit

	case "ctrl+c":
		m.quitting = true
		return m, tea.Quit

	case "ctrl+s":
		return m.saveFile()

	case "ctrl+g":
		m.mode = ModeGoto
		m.inputBuffer = ""
		m.inputPrompt = fmt.Sprintf("Go to line [1-%d]: ", m.buffer.LineCount())
		return m, nil

	// Navigation
	case "up", "ctrl+p":
		m.moveCursorUp()
	case "down", "ctrl+n":
		m.moveCursorDown()
	case "left", "ctrl+b":
		m.buffer.MoveLeft()
	case "right", "ctrl+f":
		m.buffer.MoveRight()
	case "home", "ctrl+a":
		m.moveToLineStart()
	case "end", "ctrl+e":
		m.moveToLineEnd()

	// Editing
	case "backspace":
		if m.buffer.Delete() != 0 {
			m.modified = true
		}
	case "delete":
		if m.buffer.DeleteForward() != 0 {
			m.modified = true
		}
	case "enter":
		m.buffer.Insert('\n')
		m.modified = true
	case "tab":
		m.buffer.InsertString("    ")
		m.modified = true

	default:
		// Insert printable characters
		if len(msg.Runes) > 0 {
			for _, r := range msg.Runes {
				m.buffer.Insert(r)
			}
			m.modified = true
		}
	}

	return m, nil
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

// saveFile saves the buffer to file.
func (m *Model) saveFile() (tea.Model, tea.Cmd) {
	// If no filepath, enter save-as mode
	if m.filepath == "" {
		m.mode = ModeSaveAs
		m.inputBuffer = ""
		m.inputPrompt = "Save as: "
		return m, nil
	}

	// Save to existing filepath
	err := file.Save(m.filepath, m.Content())
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
	logo := " GESH "

	// Filename with modified indicator
	filename := m.filename
	if m.modified {
		filename += " [Modified]"
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

	// Adjust viewport to keep cursor visible
	if cursorLine < m.viewportTopLine {
		m.viewportTopLine = cursorLine
	}
	if cursorLine >= m.viewportTopLine+visibleLines {
		m.viewportTopLine = cursorLine - visibleLines + 1
	}

	for i := 0; i < visibleLines; i++ {
		lineNum := m.viewportTopLine + i

		if lineNum < lineCount {
			// Line number
			lineNumStr := lineNumberStyle.Render(fmt.Sprintf("%d", lineNum+1))
			b.WriteString(lineNumStr)
			b.WriteString(" │ ")

			// Line content
			lineContent := m.buffer.Line(lineNum)

			// Render with cursor if this is the cursor line
			if lineNum == cursorLine {
				// Insert cursor character
				if cursorCol >= len([]rune(lineContent)) {
					b.WriteString(editorStyle.Render(lineContent))
					b.WriteString("█")
				} else {
					runes := []rune(lineContent)
					before := string(runes[:cursorCol])
					cursor := string(runes[cursorCol])
					after := string(runes[cursorCol+1:])
					b.WriteString(editorStyle.Render(before))
					b.WriteString(lipgloss.NewStyle().Reverse(true).Render(cursor))
					b.WriteString(editorStyle.Render(after))
				}
			} else {
				b.WriteString(editorStyle.Render(lineContent))
			}
		} else {
			// Empty line indicator
			lineNumStr := lineNumberStyle.Render("~")
			b.WriteString(lineNumStr)
			b.WriteString(" │")
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

	// File info
	lineCount := m.buffer.LineCount()
	fileInfo := fmt.Sprintf(" │ %d lines", lineCount)

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
	case ModeSaveAs, ModeGoto:
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
			helpKeyStyle.Render("^W") + helpStyle.Render(" Search"),
			helpKeyStyle.Render("^G") + helpStyle.Render(" Goto"),
		}
	}

	content := " " + strings.Join(helps, "  ")
	padding := m.width - len(content)
	if padding < 0 {
		padding = 0
	}

	return helpStyle.Render(content + strings.Repeat(" ", padding))
}

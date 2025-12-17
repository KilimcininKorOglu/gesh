# GESH Architecture

This document describes the internal architecture of Gesh text editor.

## Overview

Gesh is built using the [Elm Architecture](https://guide.elm-lang.org/architecture/) via [Bubble Tea](https://github.com/charmbracelet/bubbletea), a Go framework for terminal applications.

```
┌────────────────────────────────────────────────────────────────────┐
│                         GESH ARCHITECTURE                          │
├────────────────────────────────────────────────────────────────────┤
│                                                                    │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐         │
│  │   Terminal   │───▶│  Bubble Tea  │───▶│    Model     │         │
│  │    Input     │    │   Runtime    │    │    State     │         │
│  └──────────────┘    └──────────────┘    └──────────────┘         │
│                             │                    │                 │
│                             ▼                    ▼                 │
│                      ┌──────────────┐    ┌──────────────┐         │
│                      │    View      │◀───│   Update     │         │
│                      │   Render     │    │   Handler    │         │
│                      └──────────────┘    └──────────────┘         │
│                             │                    │                 │
│                             ▼                    ▼                 │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐         │
│  │   Terminal   │◀───│   Lipgloss   │    │  Gap Buffer  │         │
│  │   Output     │    │   Styling    │    │    + Undo    │         │
│  └──────────────┘    └──────────────┘    └──────────────┘         │
│                                                                    │
└────────────────────────────────────────────────────────────────────┘
```

---

## Core Components

### 1. Gap Buffer (`internal/buffer`)

The heart of text editing. A gap buffer is an array with a "gap" (empty space) at the cursor position, making local edits O(1).

```
Text: "Hello World"
Cursor after "Hello"

Internal representation:
['H','e','l','l','o',' ', _, _, _, _, 'W','o','r','l','d']
                     ^               ^
                 gapStart         gapEnd

Insert 'X':
['H','e','l','l','o',' ','X', _, _, _, 'W','o','r','l','d']
                         ^           ^
                     gapStart     gapEnd
```

**Key operations:**
- `Insert(rune)` - O(1) amortized
- `Delete()` - O(1)
- `MoveTo(pos)` - O(n) worst case (moves gap)
- `String()` - O(n)

### 2. History/Undo (`internal/buffer`)

Undo/redo uses an operation stack:

```go
type EditOperation struct {
    Type     OpType  // OpInsert or OpDelete
    Position int
    Text     string
}
```

Operations are merged when:
- Same type (consecutive inserts/deletes)
- Within time threshold
- Adjacent positions

### 3. Application Model (`internal/app`)

The Bubble Tea model holds all application state:

```go
type Model struct {
    buffer          *buffer.GapBuffer
    history         *buffer.History
    
    // File state
    filename        string
    filepath        string
    modified        bool
    readonly        bool
    
    // Display options
    showLineNumbers    bool
    wordWrap           bool
    syntaxHighlighting bool
    
    // Editor mode
    mode            Mode  // Normal, Search, Goto, etc.
    
    // Selection
    selecting       bool
    selectionStart  int
    selectionEnd    int
    
    // Search
    searchQuery     string
    
    // UI
    width, height   int
    viewportTopLine int
}
```

### 4. Modes

Editor operates in different modes:

| Mode | Purpose |
|------|---------|
| `ModeNormal` | Standard editing |
| `ModeSearch` | Find text |
| `ModeReplace` | Find & replace |
| `ModeGoto` | Jump to line |
| `ModeSaveAs` | Save with new name |
| `ModeQuit` | Quit confirmation |
| `ModeOpen` | Open file |

### 5. Syntax Highlighting (`internal/syntax`)

Regex-based tokenizer:

```go
type Language struct {
    Name       string
    Extensions []string
    Rules      []Rule
}

type Rule struct {
    Type    TokenType
    Pattern *regexp.Regexp
}
```

**Token types:** Keyword, Type, String, Number, Comment, Operator, Function, Variable, Constant, Builtin

---

## Data Flow

### Input Processing

```
KeyPress → tea.KeyMsg → handleKeyMsg() → Update state → tea.Cmd
                              ↓
                        Mode-specific handler
                              ↓
                        Buffer/Selection/Navigation
```

### Rendering

```
View() called
    ↓
renderHeader()   → Logo, filename, modified flag, encoding
    ↓
renderEditor()   → Line numbers, content, cursor, selection
    ↓
renderStatusBar() → Position, line count, file size, language
    ↓
renderHelpBar()  → Context-sensitive shortcuts
    ↓
lipgloss styling → Terminal output
```

---

## Module Structure

```
gesh/
├── main.go                      # Entry point, CLI parsing
│
├── internal/
│   ├── app/
│   │   ├── model.go            # Model struct, constructors
│   │   └── app.go              # Update, View, handlers
│   │
│   ├── buffer/
│   │   ├── gap_buffer.go       # Gap buffer implementation
│   │   ├── history.go          # Undo/redo stack
│   │   └── gap_buffer_test.go  # Tests (94% coverage)
│   │
│   ├── config/
│   │   └── config.go           # YAML config parsing
│   │
│   ├── file/
│   │   └── file.go             # File I/O operations
│   │
│   ├── syntax/
│   │   ├── highlighter.go      # Tokenization engine
│   │   └── languages/          # 55+ language definitions
│   │       ├── go.go
│   │       ├── python.go
│   │       ├── javascript.go
│   │       └── ...
│   │
│   └── ui/
│       └── styles/
│           └── theme.go        # Theme definitions
│
└── pkg/
    └── version/
        └── version.go          # Version info
```

---

## Key Algorithms

### Cursor Movement

```go
// Moving to position requires shifting the gap
func (b *GapBuffer) MoveTo(pos int) {
    if pos < b.gapStart {
        // Move gap left: copy chars from before gap to after
        shift := b.gapStart - pos
        copy(b.data[b.gapEnd-shift:b.gapEnd], b.data[pos:b.gapStart])
        b.gapStart = pos
        b.gapEnd -= shift
    } else if pos > b.gapStart {
        // Move gap right: copy chars from after gap to before
        shift := pos - b.gapStart
        copy(b.data[b.gapStart:], b.data[b.gapEnd:b.gapEnd+shift])
        b.gapStart = pos
        b.gapEnd += shift
    }
}
```

### Line Calculation

```go
func (b *GapBuffer) CurrentLine() int {
    line := 0
    for i := 0; i < b.gapStart; i++ {
        if b.data[i] == '\n' {
            line++
        }
    }
    return line
}
```

### Selection Bounds

```go
func (m *Model) getSelectionBounds() (int, int) {
    if m.selectionStart < m.selectionEnd {
        return m.selectionStart, m.selectionEnd
    }
    return m.selectionEnd, m.selectionStart
}
```

---

## Performance Considerations

### Current Optimizations

1. **Gap buffer** - O(1) local edits
2. **Viewport rendering** - Only render visible lines
3. **Scroll padding** - Smooth scrolling experience

### Future Optimizations

1. **Lazy loading** - For large files (>10MB)
2. **Syntax cache** - Don't re-tokenize unchanged lines
3. **Incremental render** - Only redraw changed regions
4. **Rope data structure** - For very large files

---

## Dependencies

| Dependency | Purpose |
|------------|---------|
| `bubbletea` | TUI framework (Elm architecture) |
| `lipgloss` | Terminal styling |
| `yaml.v3` | Config file parsing |

---

## Testing Strategy

| Package | Coverage | Focus |
|---------|----------|-------|
| `buffer` | 94% | Core data structures |
| `file` | 93% | File I/O |
| `version` | 100% | Version info |
| `app` | ~2% | Hard to test TUI |

---

## Extension Points

### Adding New Features

1. **New keybinding:** Add case in `handleKeyMsg()`
2. **New mode:** Add to `Mode` enum, create handler
3. **New language:** Create file in `syntax/languages/`
4. **New theme:** Add to `ui/styles/theme.go`

### Plugin System (Future)

Planned architecture:
```go
type Plugin interface {
    Name() string
    Init(editor *Editor)
    HandleKey(key tea.KeyMsg) bool
    Render() string
}
```

---

## References

- [Bubble Tea Documentation](https://github.com/charmbracelet/bubbletea)
- [Gap Buffer Wikipedia](https://en.wikipedia.org/wiki/Gap_buffer)
- [The Craft of Text Editing](https://www.finseth.com/craft/)

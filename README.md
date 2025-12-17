# Gesh (𒄑)

A minimal, fast, nano-like TUI text editor written in Go with [Bubble Tea](https://github.com/charmbracelet/bubbletea).

> **Gesh** (𒄑) is the Sumerian cuneiform sign meaning "wood, tree" and by extension "pen, stylus, writing tool". In ancient Mesopotamia, scribes used reed styluses to write on clay tablets - the earliest form of text editing. This editor carries that legacy into the modern terminal.

## Features

### Core

| Feature | Description |
|---------|-------------|
| Fast startup | < 50ms cold start |
| Low memory | < 10MB empty, < 15MB with 1MB file |
| Single binary | No external dependencies |
| Cross-platform | Linux, macOS, Windows |
| Mouse support | Click positioning, scroll wheel |

### Editing

| Feature | Description |
|---------|-------------|
| Undo/Redo | Operation merging for efficient history |
| Clipboard | System clipboard integration |
| Search & Replace | Incremental search with highlighting |
| Selection | Keyboard and shift+arrow selection |
| Insert/Overwrite | Toggle with Insert key |
| Macros | Record and playback keystrokes |

### User Interface

| Feature | Description |
|---------|-------------|
| Syntax highlighting | 55+ languages supported |
| Color themes | dark, light, monokai, dracula, gruvbox |
| Multi-tab | Multiple files in tabs |
| Split view | Horizontal and vertical splits |
| Line numbers | Toggleable with current line marker |
| Status bar | Position, encoding, language, mode |

### File Operations

| Feature | Description |
|---------|-------------|
| Auto-save | Configurable interval |
| Backup files | Optional .bak creation |
| File watcher | External change detection |
| Encoding | UTF-8, UTF-8 BOM, Latin-1 |
| Line endings | LF, CRLF, CR detection |
| Save options | Trim trailing spaces, final newline |

### Advanced

| Feature | Description |
|---------|-------------|
| Smooth scroll | Animated page up/down |
| Word wrap | Soft line wrapping |
| Configuration | YAML-based settings |
| Large files | Chunked loading for files >10MB |

---

## Installation

### From Source

```bash
go install github.com/KilimcininKorOglu/gesh@latest
```

### Build Locally

```bash
git clone https://github.com/KilimcininKorOglu/gesh.git
cd gesh
go build -o gesh .
```

---

## Usage

```bash
gesh                      # New empty file
gesh filename.txt         # Open file
gesh +42 main.go          # Open at line 42
gesh +42:10 main.go       # Open at line 42, column 10
gesh -r config.yaml       # Read-only mode
gesh -t monokai main.go   # Use specific theme
gesh --version            # Show version
gesh --help               # Show help
```

---

## Keyboard Shortcuts

### File Operations

| Shortcut | Action |
|:---------|:-------|
| `Ctrl+S` | Save |
| `Ctrl+Shift+S` | Save as |
| `Ctrl+O` | Open |
| `Ctrl+X` | Exit |

### Navigation

| Shortcut | Action |
|:---------|:-------|
| `Arrow keys` | Move cursor |
| `Ctrl+Left/Right` | Move by word |
| `Home` / `End` | Line start/end |
| `Ctrl+Home/End` | File start/end |
| `PageUp/Down` | Page navigation |
| `Ctrl+G` | Go to line |

### Editing

| Shortcut | Action |
|:---------|:-------|
| `Backspace` | Delete before |
| `Delete` | Delete after |
| `Ctrl+K` | Delete line |
| `Ctrl+U` | Cut line |
| `Ctrl+V` | Paste |
| `Ctrl+Z` | Undo |
| `Ctrl+Y` | Redo |
| `Insert` | Toggle INS/OVR |

### Selection

| Shortcut | Action |
|:---------|:-------|
| `Ctrl+Space` | Toggle selection |
| `Shift+Arrows` | Extend selection |
| `Ctrl+A` (2x) | Select all |
| `Ctrl+C` | Copy |
| `Ctrl+X` | Cut |

### Search & Replace

| Shortcut | Action |
|:---------|:-------|
| `Ctrl+W` | Search |
| `Ctrl+R` | Replace |
| `Ctrl+Shift+R` | Replace all |
| `F3` | Next match |
| `Shift+F3` | Previous match |

### Tabs

| Shortcut | Action |
|:---------|:-------|
| `Ctrl+T` | New tab |
| `Ctrl+W` | Close tab |
| `Ctrl+Tab` | Next tab |
| `Ctrl+Shift+Tab` | Previous tab |
| `Alt+1-9` | Go to tab N |

### Macros

| Shortcut | Action |
|:---------|:-------|
| `Ctrl+M` | Toggle recording |
| `Ctrl+Shift+M` | Play macro |

### Split View

| Shortcut | Action |
|:---------|:-------|
| `Ctrl+\` | Horizontal split |
| `Ctrl+Shift+-` | Vertical split |
| `Ctrl+Shift+\` | Close split |
| `Alt+Left/H` | Focus left pane |
| `Alt+Right/L` | Focus right pane |

---

## Themes

Available: `dark`, `light`, `monokai`, `dracula`, `gruvbox`

```bash
gesh -t monokai file.go
```

Or in config:
```yaml
theme: dracula
```

---

## Configuration

**Location:**
- Linux/macOS: `~/.config/gesh/gesh.yaml`
- Windows: `%APPDATA%\gesh\gesh.yaml`

**Example:**

```yaml
editor:
  tab_size: 4
  insert_spaces: true
  auto_indent: true
  word_wrap: false
  line_numbers: true
  scroll_padding: 5
  trim_trailing_spaces: false
  final_newline: true
  create_backup: false
  auto_save_interval: 0

theme: dark
```

---

## Documentation

| Document | Description |
|:---------|:------------|
| [INSTALL.md](INSTALL.md) | Installation guide |
| [KEYBINDINGS.md](KEYBINDINGS.md) | Full keyboard reference |
| [CONFIG.md](CONFIG.md) | Configuration options |
| [THEMES.md](THEMES.md) | Theme customization |
| [CONTRIBUTING.md](CONTRIBUTING.md) | Contribution guide |
| [ARCHITECTURE.md](ARCHITECTURE.md) | Technical details |

---

## Architecture

Gesh uses the Elm Architecture via Bubble Tea:

```
┌─────────────────────────────────────────────┐
│ Header: Logo | Filename | Encoding          │
├─────────────────────────────────────────────┤
│ Editor: Line numbers | Content              │
├─────────────────────────────────────────────┤
│ Status: Position | Language | Mode          │
├─────────────────────────────────────────────┤
│ Help: Context-sensitive shortcuts           │
└─────────────────────────────────────────────┘
```

**Core Components:**

| Component | Purpose |
|:----------|:--------|
| Gap Buffer | O(1) local text editing |
| History | Undo/redo with merging |
| TabManager | Multi-tab management |
| SplitManager | Split view panes |
| MacroRecorder | Keystroke recording |
| FileWatcher | Change detection |
| Syntax Highlighter | Tokenization |

---

## Performance

| Metric | Target | Actual |
|:-------|-------:|-------:|
| Startup | < 50ms | ~30ms |
| Keystroke latency | < 16ms | < 10ms |
| Memory (empty) | < 5MB | ~3MB |
| Memory (1MB file) | < 15MB | ~12MB |
| Scroll | 60fps | 60fps |

---

## Testing

```bash
go test ./...           # Run all tests
go test ./... -cover    # With coverage
go test -bench=. ./...  # Benchmarks
```

**Coverage:**

| Package | Coverage |
|:--------|---------:|
| buffer | 94% |
| file | 93% |
| version | 100% |

---

## License

MIT License - see [LICENSE](LICENSE) for details.

---

## Credits

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Terminal styling
- [Cobra](https://github.com/spf13/cobra) - CLI framework

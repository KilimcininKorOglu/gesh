# Gesh (𒄑)

A minimal, fast, nano-like TUI text editor written in Go with [Bubble Tea](https://github.com/charmbracelet/bubbletea).

> **Gesh** (𒄑) - Sumerian word meaning "pen, writing tool"

![Version](https://img.shields.io/badge/version-0.1.0-blue)
![Go](https://img.shields.io/badge/go-%3E%3D1.21-00ADD8)
![License](https://img.shields.io/badge/license-MIT-green)

## ✨ Features

### Core
- 🚀 **Fast startup** (< 50ms)
- 💾 **Low memory** (< 10MB empty, < 15MB with 1MB file)
- 📦 **Single binary**, no dependencies
- 🖥️ **Cross-platform** (Linux, macOS, Windows)
- 🖱️ **Mouse support** (click, scroll)

### Editing
- ↩️ **Undo/Redo** with operation merging
- 📋 **Clipboard** integration (copy, cut, paste)
- 🔍 **Search & Replace** with highlighting
- ✏️ **Selection mode** (keyboard & shift+arrows)
- 🔄 **Insert/Overwrite** mode toggle
- 📼 **Macro recording** and playback

### UI/UX
- 🎨 **Syntax highlighting** for 55+ languages
- 🌈 **5 color themes** (dark, light, monokai, dracula, gruvbox)
- 📑 **Multi-tab support**
- 🪟 **Split view** (horizontal/vertical)
- 📏 **Line numbers** (toggleable)
- 🔢 **Current line marker**
- 📊 **Status bar** with file info, encoding, position

### File Operations
- 💾 **Auto-save** with configurable interval
- 📝 **Backup files** (.bak) before saving
- 🔍 **File watcher** for external changes
- 📄 **Encoding detection** (UTF-8, UTF-8 BOM, Latin-1)
- ↵ **Line ending detection** (LF, CRLF, CR)
- ✂️ **Trim trailing spaces** on save
- ⏎ **Final newline** enforcement

### Advanced
- 🎯 **Smooth scroll** animation
- 📜 **Word wrap** support
- 🔧 **Configurable** via YAML
- 📂 **Large file support** (chunked loading >10MB)

## 📥 Installation

### From Source

```bash
go install github.com/KilimcininKorOglu/gesh@latest
```

### Build from source

```bash
git clone https://github.com/KilimcininKorOglu/gesh.git
cd gesh
go build -o gesh .
```

## 🚀 Usage

```bash
# New empty file
gesh

# Open file
gesh filename.txt

# Open at specific line
gesh +42 main.go

# Open at specific line and column
gesh +42:10 main.go

# Open in read-only mode
gesh -r config.yaml

# Use specific theme
gesh -t monokai main.go

# Show version
gesh --version

# Show help
gesh --help
```

## ⌨️ Keyboard Shortcuts

### File Operations

| Shortcut | Action |
|----------|--------|
| `Ctrl+S` | Save file |
| `Ctrl+Shift+S` | Save as |
| `Ctrl+O` | Open file |
| `Ctrl+X` | Exit (prompts if unsaved) |

### Navigation

| Shortcut | Action |
|----------|--------|
| `↑↓←→` | Move cursor |
| `Ctrl+←/→` | Move by word |
| `Home` / `End` | Start/end of line |
| `Ctrl+Home/End` | Start/end of file |
| `PageUp/Down` | Page up/down (smooth scroll) |
| `Ctrl+G` | Go to line |

### Editing

| Shortcut | Action |
|----------|--------|
| `Backspace` | Delete before cursor |
| `Delete` | Delete after cursor |
| `Ctrl+K` | Delete line |
| `Ctrl+U` | Cut line |
| `Ctrl+V` | Paste |
| `Ctrl+Z` | Undo |
| `Ctrl+Y` | Redo |
| `Insert` | Toggle INS/OVR mode |

### Selection

| Shortcut | Action |
|----------|--------|
| `Ctrl+Space` | Toggle selection mode |
| `Shift+Arrows` | Select text |
| `Ctrl+A` (2x) | Select all |
| `Ctrl+C` | Copy |
| `Ctrl+X` | Cut selection |

### Search & Replace

| Shortcut | Action |
|----------|--------|
| `Ctrl+W` | Search |
| `Ctrl+R` | Replace one |
| `Ctrl+Shift+R` | Replace all |
| `F3` / `Shift+F3` | Next/previous match |

### Tabs

| Shortcut | Action |
|----------|--------|
| `Ctrl+T` | New tab |
| `Ctrl+W` | Close tab |
| `Ctrl+Tab` | Next tab |
| `Ctrl+Shift+Tab` | Previous tab |
| `Alt+1-9` | Go to tab N |

### Macros

| Shortcut | Action |
|----------|--------|
| `Ctrl+M` | Start/stop recording |
| `Ctrl+Shift+M` | Play macro |

### Split View

| Shortcut | Action |
|----------|--------|
| `Ctrl+\` | Horizontal split |
| `Ctrl+Shift+-` | Vertical split |
| `Ctrl+Shift+\` | Close split |
| `Alt+Left/H` | Focus left/top pane |
| `Alt+Right/L` | Focus right/bottom pane |

## 🎨 Themes

Available themes: `dark`, `light`, `monokai`, `dracula`, `gruvbox`

```bash
# Use theme from CLI
gesh -t monokai file.go

# Or set in config file
echo "theme: dracula" >> ~/.config/gesh/gesh.yaml
```

## ⚙️ Configuration

Config file location:
- Linux/macOS: `~/.config/gesh/gesh.yaml`
- Windows: `%APPDATA%\gesh\gesh.yaml`

Example configuration:

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
  auto_save_interval: 0  # seconds, 0 = disabled

theme: dark
```

## 📖 Documentation

- [INSTALL.md](INSTALL.md) - Installation guide
- [KEYBINDINGS.md](KEYBINDINGS.md) - Full keyboard shortcuts
- [CONFIG.md](CONFIG.md) - Configuration options
- [THEMES.md](THEMES.md) - Theme customization
- [CONTRIBUTING.md](CONTRIBUTING.md) - Contribution guide
- [ARCHITECTURE.md](ARCHITECTURE.md) - Technical details

## 🏗️ Architecture

Gesh uses the [Elm Architecture](https://guide.elm-lang.org/architecture/) via Bubble Tea:

```
┌─────────────────────────────────────────────┐
│                   View                       │
│  ┌─────────────────────────────────────┐    │
│  │ Header: Logo | Filename | Encoding  │    │
│  ├─────────────────────────────────────┤    │
│  │ Editor: Line numbers | Content      │    │
│  ├─────────────────────────────────────┤    │
│  │ Status: Position | Lang | Mode      │    │
│  ├─────────────────────────────────────┤    │
│  │ Help: Context-sensitive shortcuts   │    │
│  └─────────────────────────────────────┘    │
└─────────────────────────────────────────────┘
```

### Core Components

- **Gap Buffer** - O(1) local text editing
- **History** - Undo/redo with operation merging
- **TabManager** - Multi-tab buffer management
- **SplitManager** - Split view pane management
- **MacroRecorder** - Keystroke recording/playback
- **FileWatcher** - External change detection
- **Syntax Highlighter** - Regex-based tokenization

## 📊 Performance

| Metric | Target | Actual |
|--------|--------|--------|
| Startup time | < 50ms | ~30ms |
| Keystroke latency | < 16ms | < 10ms |
| Memory (empty) | < 5MB | ~3MB |
| Memory (1MB file) | < 15MB | ~12MB |
| Scroll | 60fps | 60fps |

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test ./... -cover

# Run benchmarks
go test -bench=. ./...
```

Test coverage:
- `buffer`: 94%
- `file`: 93%
- `version`: 100%

## 📝 License

MIT License - see [LICENSE](LICENSE) for details.

## 🙏 Credits

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Terminal styling
- [Cobra](https://github.com/spf13/cobra) - CLI framework

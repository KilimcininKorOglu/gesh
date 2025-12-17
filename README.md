# Gesh (𒄑)

A minimal TUI text editor written in Go with [Bubble Tea](https://github.com/charmbracelet/bubbletea).

> **Gesh** (𒄑) - Sumerian word meaning "pen, writing tool"

## Features

- 🚀 Fast startup (< 50ms)
- 💾 Low memory footprint (< 10MB)
- 📦 Single binary, no dependencies
- 🖥️ Cross-platform (Linux, macOS, Windows)
- ↩️ Undo/Redo with operation merging
- 🔍 Search with F3/Shift+F3 navigation
- 📝 nano-style keyboard shortcuts

## Installation

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

## Usage

```bash
# Open new file
gesh

# Open existing file
gesh filename.txt

# Show version
gesh --version

# Show help
gesh --help
```

## Keyboard Shortcuts

### File Operations

| Shortcut | Action |
|----------|--------|
| `Ctrl+Alt+N` | New file |
| `Ctrl+O` | Open file |
| `Ctrl+S` | Save file |
| `Ctrl+Shift+S` | Save as |
| `Ctrl+X` | Exit (or cut if selecting) |
| `Ctrl+C` | Copy selection / Force quit |

### Navigation

| Shortcut | Action |
|----------|--------|
| `↑` `↓` `←` `→` | Move cursor |
| `Ctrl+P` / `Ctrl+N` | Up / Down (nano style) |
| `Ctrl+B` / `Ctrl+F` | Left / Right (nano style) |
| `Ctrl+←` / `Ctrl+→` | Move by word |
| `Home` / `Ctrl+A` | Start of line (2x = select all) |
| `End` / `Ctrl+E` | End of line |
| `Ctrl+Home` | Start of file |
| `Ctrl+End` | End of file |
| `PageUp` | Page up |
| `PageDown` | Page down |
| `Ctrl+G` | Go to line |

### Editing

| Shortcut | Action |
|----------|--------|
| `Backspace` | Delete character before cursor |
| `Delete` | Delete character after cursor |
| `Ctrl+K` | Delete current line |
| `Ctrl+U` | Cut line to clipboard |
| `Ctrl+V` | Paste from clipboard |
| `Ctrl+Z` | Undo |
| `Ctrl+Y` | Redo |
| `Tab` | Insert 4 spaces |

### Selection

| Shortcut | Action |
|----------|--------|
| `Ctrl+Space` | Toggle selection mode |
| `Shift+↑↓←→` | Select text |
| `Ctrl+A` (2x) | Select all |
| `Ctrl+C` | Copy selection |
| `Ctrl+X` | Cut selection |

### Search & Replace

| Shortcut | Action |
|----------|--------|
| `Ctrl+W` | Search |
| `Ctrl+R` | Replace (one at a time) |
| `Ctrl+Shift+R` | Replace all |
| `F3` | Next match |
| `Shift+F3` | Previous match |

## Architecture

Gesh uses the [Elm Architecture](https://guide.elm-lang.org/architecture/) via Bubble Tea:

- **Model**: Application state (buffer, cursor, mode)
- **Update**: Handle keyboard/mouse events
- **View**: Render the UI

### Core Components

- **Gap Buffer**: Efficient text editing data structure with O(1) local edits
- **History**: Undo/redo stack with operation merging
- **Lipgloss**: Terminal styling

## Project Structure

```
gesh/
├── main.go                 # Entry point
├── internal/
│   ├── app/                # Bubble Tea model
│   ├── buffer/             # Gap buffer & history
│   ├── file/               # File I/O
│   └── ...
├── pkg/
│   └── version/            # Version info
└── configs/                # Example configs
```

## Development

```bash
# Run tests
go test ./...

# Run with coverage
go test ./... -cover

# Build
go build -o gesh .

# Run
./gesh
```

## License

MIT License - see [LICENSE](LICENSE) for details.

## Credits

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Terminal styling

<p align="center">
  <h1 align="center">𒄑 GESH</h1>
  <p align="center">
    <strong>A nano-compatible terminal text editor written in Go</strong>
  </p>
  <p align="center">
    <a href="#installation">Installation</a> •
    <a href="#features">Features</a> •
    <a href="#keyboard-shortcuts">Shortcuts</a> •
    <a href="#configuration">Config</a> •
    <a href="#documentation">Docs</a>
  </p>
</p>

---

**Gesh** (𒄑) is the Sumerian cuneiform sign meaning "wood, tree" and by extension "pen, stylus, writing tool". In ancient Mesopotamia, scribes used reed styluses to write on clay tablets — the earliest form of text editing.

Gesh brings that legacy to your modern terminal with **nano-compatible keybindings**, making it instantly familiar while adding powerful features like syntax highlighting, multiple tabs, split views, and macros.

## Why Gesh?

- **Fast** -- Starts in <50ms, keystroke latency <10ms
- **Lightweight** -- Single binary, ~3MB memory usage
- **Nano-compatible** -- Same shortcuts you already know
- **Syntax highlighting** -- 55+ languages out of the box
- **Tabs & splits** -- Edit multiple files efficiently
- **Mouse support** -- Click, scroll, drag to select
- **Configurable** -- YAML-based settings, 5 themes

---

## Installation

### Pre-built Binaries

Download from [Releases](https://github.com/KilimcininKorOglu/gesh/releases) for:
- Linux (amd64, arm64, arm)
- macOS (amd64, arm64)
- Windows (amd64, arm64)
- FreeBSD (amd64)

### From Source

```bash
# With Go installed
go install github.com/KilimcininKorOglu/gesh@latest

# Or build locally
git clone https://github.com/KilimcininKorOglu/gesh.git
cd gesh
make build                # Linux/macOS: build for current platform
.\build.bat build         # Windows: build for current platform
```

### Package Managers

```bash
# Coming soon
# brew install gesh
# apt install gesh
# scoop install gesh
```

---

## Quick Start

```bash
gesh                      # New file
gesh README.md            # Open file
gesh +100 main.go         # Open at line 100
gesh -r config.yaml       # Read-only mode
gesh --theme dracula      # With theme
```

---

## Features

### Core Editor
| Feature          | Description                           |
|------------------|---------------------------------------|
| Gap Buffer       | O(1) insertions at cursor position    |
| Undo/Redo        | Intelligent operation merging         |
| Search & Replace | Incremental search with highlighting  |
| Selection        | Keyboard, shift+arrows, or mouse drag |
| Auto-indent      | Preserves indentation on Enter        |

### Multi-File Editing
| Feature      | Description                                 |
|--------------|---------------------------------------------|
| Tabs         | `Ctrl+T` new, `Ctrl+Tab` switch             |
| Split View   | Horizontal (`Alt+\\`) or vertical (`Alt+-`) |
| File Watcher | Detects external changes                    |

### Interface
| Feature             | Description                            |
|---------------------|----------------------------------------|
| Syntax Highlighting | 55+ languages                          |
| Themes              | dark, light, monokai, dracula, gruvbox |
| Line Numbers        | With current line marker               |
| Status Bar          | Position, encoding, language           |
| Help Bar            | Context-sensitive nano-style shortcuts |

### Mouse Support
| Action          | Mouse                      |
|-----------------|----------------------------|
| Position cursor | Left click                 |
| Select text     | Left click + drag          |
| Copy selection  | Right click                |
| Paste           | Right click (no selection) |
| Scroll          | Mouse wheel                |

### File Handling
| Feature      | Description                |
|--------------|----------------------------|
| Encodings    | UTF-8, UTF-8 BOM, Latin-1  |
| Line Endings | LF, CRLF, CR (auto-detect) |
| Auto-save    | Configurable interval      |
| Backup Files | Optional .bak creation     |
| Large Files  | Chunked loading for >10MB  |

---

## Keyboard Shortcuts

Gesh uses **nano-compatible** keybindings. If you know nano, you know Gesh.

### Essential

| Shortcut | Action           |
|----------|------------------|
| `Ctrl+X` | Exit             |
| `Ctrl+O` | Save (Write Out) |
| `Ctrl+R` | Read/Insert file |
| `Ctrl+G` | Help             |

### Editing

| Shortcut | Action              |
|----------|---------------------|
| `Ctrl+K` | Cut line/selection  |
| `Ctrl+U` | Paste (Uncut)       |
| `Alt+6`  | Copy line/selection |
| `Alt+U`  | Undo                |
| `Alt+E`  | Redo                |

### Navigation

| Shortcut           | Action            |
|--------------------|-------------------|
| `Ctrl+Y`           | Page Up           |
| `Ctrl+V`           | Page Down         |
| `Ctrl+A`           | Beginning of line |
| `Ctrl+E`           | End of line       |
| `Alt+\`            | Beginning of file |
| `Alt+/`            | End of file       |
| `Ctrl+_` / `Alt+G` | Go to line        |

### Search

| Shortcut           | Action            |
|--------------------|-------------------|
| `Ctrl+W`           | Search (Where Is) |
| `Alt+W`            | Find next         |
| `Ctrl+Q`           | Find previous     |
| `Ctrl+\` / `Alt+R` | Replace           |

### Selection

| Shortcut       | Action                     |
|----------------|----------------------------|
| `Alt+A`        | Set mark (start selection) |
| `Shift+Arrows` | Extend selection           |

### Display

| Shortcut | Action               |
|----------|----------------------|
| `Ctrl+C` | Show cursor position |
| `Alt+N`  | Toggle line numbers  |
| `Ctrl+L` | Refresh screen       |

### Extensions (Beyond Nano)

| Shortcut   | Action           |
|------------|------------------|
| `Ctrl+T`   | New tab          |
| `Ctrl+Tab` | Next tab         |
| `Alt+\\`   | Horizontal split |
| `Alt+-`    | Vertical split   |
| `F4`       | Record macro     |
| `F5`       | Play macro       |

> Full reference: [KEYBINDINGS.md](docs/KEYBINDINGS.md)

---

## Themes

```bash
gesh --theme monokai file.go
```

| Theme     | Description        |
|-----------|--------------------|
| `dark`    | Default dark theme |
| `light`   | Light background   |
| `monokai` | Classic Monokai    |
| `dracula` | Dracula colors     |
| `gruvbox` | Gruvbox palette    |

---

## Configuration

**Location:**
- Linux/macOS: `~/.config/gesh/gesh.yaml`
- Windows: `%APPDATA%\gesh\gesh.yaml`

```yaml
editor:
  tab_size: 4
  insert_spaces: true
  auto_indent: true
  word_wrap: false
  line_numbers: true
  trim_trailing_spaces: false
  final_newline: true
  auto_save_interval: 0  # seconds, 0 = disabled

theme: dark
```

> Full options: [CONFIG.md](docs/CONFIG.md)

---

## Performance

| Metric            | Target | Actual |
|-------------------|--------|--------|
| Startup           | <50ms  | ~30ms  |
| Keystroke latency | <16ms  | <10ms  |
| Memory (empty)    | <5MB   | ~3MB   |
| Memory (1MB file) | <15MB  | ~12MB  |

---

## Documentation

| Document                           | Description                 |
|------------------------------------|-----------------------------|
| [INSTALL.md](docs/INSTALL.md)           | Installation guide          |
| [KEYBINDINGS.md](docs/KEYBINDINGS.md)   | Complete keyboard reference |
| [CONFIG.md](docs/CONFIG.md)             | Configuration options       |
| [THEMES.md](docs/THEMES.md)             | Theme customization         |
| [CONTRIBUTING.md](CONTRIBUTING.md)       | How to contribute           |
| [ARCHITECTURE.md](docs/ARCHITECTURE.md) | Technical internals         |

---

## Building

All builds and tests go through `make` (Linux/macOS) or `build.bat` (Windows).

```bash
# Linux/macOS
make build                # Build for current platform
make build-all-platforms  # Build for all platforms
make test                 # Run all tests
make check                # Run fmt + vet + lint + test

# Windows
.\build.bat build         # Build for current platform
.\build.bat build-all     # Build for all platforms
.\build.bat test          # Run all tests
.\build.bat check         # Run fmt + vet + lint + test
```

---

## Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) first.

```bash
# Fork and clone
git clone https://github.com/YOUR_USERNAME/gesh.git

# Create branch
git checkout -b feature/amazing-feature

# Make changes, test, build
make check

# Commit and push
git commit -m "feat: add amazing feature"
git push origin feature/amazing-feature
```

---

## License

MIT License — see [LICENSE](LICENSE) for details.

---

## Acknowledgments

Built with these excellent libraries:

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) — TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) — Terminal styling

---

<p align="center">
  <sub>Made for the terminal</sub>
</p>

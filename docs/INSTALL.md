# GESH Installation Guide

## Quick Install

### Pre-built Binaries

Download the latest release from [GitHub Releases](https://github.com/KilimcininKorOglu/gesh/releases).

#### Linux (amd64)
```bash
curl -LO https://github.com/KilimcininKorOglu/gesh/releases/latest/download/gesh-linux-amd64.tar.gz
tar -xzf gesh-linux-amd64.tar.gz
sudo mv gesh /usr/local/bin/
```

#### macOS (amd64/arm64)
```bash
# Intel Mac
curl -LO https://github.com/KilimcininKorOglu/gesh/releases/latest/download/gesh-darwin-amd64.tar.gz

# Apple Silicon
curl -LO https://github.com/KilimcininKorOglu/gesh/releases/latest/download/gesh-darwin-arm64.tar.gz

tar -xzf gesh-darwin-*.tar.gz
sudo mv gesh /usr/local/bin/
```

#### Windows
```powershell
# Download gesh-windows-amd64.zip from releases
# Extract and add to PATH
```

---

## Build from Source

### Prerequisites

- Go 1.24 or later
- Git
- Make (Linux/macOS) or Windows shell for build.bat

### Steps

```bash
# Clone the repository
git clone https://github.com/KilimcininKorOglu/gesh.git
cd gesh

# Build for current platform (Linux/macOS)
make build

# Build for current platform (Windows)
.\build.bat build

# Install to GOPATH/bin
make install
```

The Makefile and build.bat automatically inject version info (version, commit hash, build date) via ldflags.

---

## Verify Installation

```bash
gesh --version
# Output: Gesh v0.1.0 (commit: abc1234)

gesh --help
```

---

## Configuration

On first run, Gesh uses default settings. To customize:

```bash
# Linux/macOS
mkdir -p ~/.config/gesh
touch ~/.config/gesh/gesh.yaml

# Windows
mkdir %APPDATA%\gesh
```

See [CONFIG.md](CONFIG.md) for configuration options.

---

## Uninstall

```bash
# Linux/macOS
sudo rm /usr/local/bin/gesh
rm -rf ~/.config/gesh

# Windows
# Remove gesh.exe from your PATH
# Delete %APPDATA%\gesh folder
```

---

## Troubleshooting

### "gesh: command not found"

Make sure the binary is in your PATH:
```bash
which gesh
echo $PATH
```

### Permission denied

```bash
chmod +x /usr/local/bin/gesh
```

### Terminal compatibility issues

Gesh requires a terminal with:
- UTF-8 support
- 256 color support (for themes)
- Mouse support (optional)

Recommended terminals:
- **Linux**: Alacritty, Kitty, GNOME Terminal
- **macOS**: iTerm2, Terminal.app
- **Windows**: Windows Terminal, ConEmu

---

## Next Steps

- Read [KEYBINDINGS.md](KEYBINDINGS.md) for keyboard shortcuts
- Customize your setup in [CONFIG.md](CONFIG.md)
- Create custom themes with [THEMES.md](THEMES.md)

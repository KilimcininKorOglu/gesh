# GESH Configuration Guide

## Configuration File Location

| Platform | Path |
|----------|------|
| Linux | `~/.config/gesh/gesh.yaml` |
| macOS | `~/.config/gesh/gesh.yaml` |
| Windows | `%APPDATA%\gesh\gesh.yaml` |

---

## Default Configuration

```yaml
# gesh.yaml - Gesh Text Editor Configuration

editor:
  # Tab size in spaces
  tab_size: 4
  
  # Use spaces instead of tabs
  insert_spaces: true
  
  # Auto-indent on Enter
  auto_indent: true
  
  # Word wrap long lines
  word_wrap: false
  
  # Show line numbers
  line_numbers: true
  
  # Scroll padding (lines from edge)
  scroll_padding: 5
  
  # Trim trailing whitespace on save
  trim_trailing_spaces: false
  
  # Ensure file ends with newline
  final_newline: false
  
  # Create backup file (.bak) before saving
  create_backup: false

# Theme name: dark, light, monokai, dracula, gruvbox
theme: dark
```

---

## Configuration Options

### Editor Settings

#### `tab_size`
- **Type:** Integer
- **Default:** `4`
- **Description:** Number of spaces per tab character

#### `insert_spaces`
- **Type:** Boolean
- **Default:** `true`
- **Description:** Insert spaces when Tab is pressed instead of tab character

#### `auto_indent`
- **Type:** Boolean
- **Default:** `true`
- **Description:** Automatically indent new lines based on previous line

#### `word_wrap`
- **Type:** Boolean
- **Default:** `false`
- **Description:** Wrap long lines at window edge

#### `line_numbers`
- **Type:** Boolean
- **Default:** `true`
- **Description:** Show line numbers in the gutter

#### `scroll_padding`
- **Type:** Integer
- **Default:** `5`
- **Description:** Minimum lines between cursor and window edge when scrolling

#### `trim_trailing_spaces`
- **Type:** Boolean
- **Default:** `false`
- **Description:** Remove trailing whitespace when saving

#### `final_newline`
- **Type:** Boolean
- **Default:** `false`
- **Description:** Ensure file ends with a newline character

#### `create_backup`
- **Type:** Boolean
- **Default:** `false`
- **Description:** Create a backup file (.bak) before saving

---

### Theme Settings

#### `theme`
- **Type:** String
- **Default:** `dark`
- **Options:** `dark`, `light`, `monokai`, `dracula`, `gruvbox`
- **Description:** Color theme for the editor

---

## Built-in Themes

### Dark (default)
Classic dark theme with blue/pink accents.

### Light
Light background theme for bright environments.

### Monokai
Popular Monokai color scheme.

### Dracula
Purple-tinted Dracula theme.

### Gruvbox
Warm, retro Gruvbox colors.

---

## CLI Overrides

Configuration can be overridden via command line:

```bash
# Use light theme
gesh --theme light file.txt

# Disable line numbers
gesh --no-line-numbers file.txt

# Disable syntax highlighting
gesh --no-syntax file.txt

# Skip loading config file
gesh --norc file.txt
```

---

## Example Configurations

### Minimal Config
```yaml
theme: dark
editor:
  line_numbers: true
```

### Python Developer
```yaml
theme: monokai
editor:
  tab_size: 4
  insert_spaces: true
  auto_indent: true
  trim_trailing_spaces: true
  final_newline: true
```

### Go Developer
```yaml
theme: dracula
editor:
  tab_size: 4
  insert_spaces: false  # Go uses tabs
  auto_indent: true
  trim_trailing_spaces: true
  final_newline: true
```

### Writer/Markdown
```yaml
theme: light
editor:
  word_wrap: true
  line_numbers: false
  scroll_padding: 10
```

---

## Troubleshooting

### Config not loading

1. Check file exists:
   ```bash
   cat ~/.config/gesh/gesh.yaml
   ```

2. Validate YAML syntax:
   ```bash
   python -c "import yaml; yaml.safe_load(open('~/.config/gesh/gesh.yaml'))"
   ```

3. Check for typos in option names

### Reset to defaults

Delete or rename the config file:
```bash
mv ~/.config/gesh/gesh.yaml ~/.config/gesh/gesh.yaml.bak
```

---

## See Also

- [THEMES.md](THEMES.md) - Create custom themes
- [KEYBINDINGS.md](KEYBINDINGS.md) - Keyboard shortcuts

# GESH Themes Guide

## Built-in Themes

Gesh comes with 5 built-in themes:

| Theme | Description |
|-------|-------------|
| `dark` | Default dark theme with blue/pink accents |
| `light` | Clean light theme for bright environments |
| `monokai` | Classic Monokai color scheme |
| `dracula` | Popular purple-tinted Dracula theme |
| `gruvbox` | Warm, retro Gruvbox colors |

---

## Using Themes

### Via Config File

```yaml
# ~/.config/gesh/gesh.yaml
theme: monokai
```

### Via Command Line

```bash
gesh --theme dracula file.txt
gesh -t gruvbox file.txt
```

---

## Theme Preview

### Dark Theme
```
Background: #1a1a2e (dark blue-gray)
Header:     #16213e with #e94560 text (pink)
Editor:     #eaeaea text
Line nums:  #4a4a6a (muted purple)
Selection:  #3a3a5a
```

### Light Theme
```
Background: White
Header:     #e8e8e8 with #2c3e50 text
Editor:     #1a1a1a text
Line nums:  #808080 (gray)
Selection:  #b0c4de (light blue)
```

### Monokai Theme
```
Background: #272822 (dark brown-green)
Keywords:   #f92672 (pink)
Strings:    #e6db74 (yellow)
Comments:   #75715e (brown-gray)
Functions:  #a6e22e (green)
```

### Dracula Theme
```
Background: #282a36 (dark purple-gray)
Keywords:   #ff79c6 (pink)
Strings:    #f1fa8c (yellow)
Comments:   #6272a4 (muted blue)
Functions:  #50fa7b (green)
Types:      #bd93f9 (purple)
```

### Gruvbox Theme
```
Background: #282828 (dark brown)
Keywords:   #fb4934 (red)
Strings:    #b8bb26 (green)
Comments:   #928374 (gray)
Functions:  #fabd2f (yellow)
Types:      #83a598 (blue)
```

---

## Syntax Highlighting Colors

Each theme defines colors for syntax elements:

| Element | Description |
|---------|-------------|
| `keyword` | Language keywords (if, for, func, etc.) |
| `type` | Type names (int, string, etc.) |
| `string` | String literals |
| `number` | Numeric literals |
| `comment` | Comments |
| `function` | Function names |
| `variable` | Variables |
| `constant` | Constants (true, false, nil) |
| `builtin` | Built-in functions |
| `operator` | Operators (+, -, =, etc.) |

---

## Current Syntax Colors (Dracula-inspired)

The syntax highlighting uses these colors regardless of theme:

```
Keyword:   #ff79c6 (pink)
Type:      #8be9fd (cyan)
String:    #f1fa8c (yellow)
Number:    #bd93f9 (purple)
Comment:   #6272a4 (muted blue)
Operator:  #ff79c6 (pink)
Function:  #50fa7b (green)
Variable:  #ffb86c (orange)
Constant:  #bd93f9 (purple)
Builtin:   #8be9fd (cyan)
```

---

## Tips

### Choosing a Theme

- **Dark environments:** Use `dark`, `monokai`, `dracula`, or `gruvbox`
- **Bright environments:** Use `light`
- **Eye strain:** `gruvbox` has warmer, less harsh colors
- **High contrast:** `monokai` offers good contrast

### Terminal Compatibility

For best results, use a terminal with:
- True color (24-bit) support
- A font with good Unicode coverage

Recommended terminals:
- **Linux:** Alacritty, Kitty, GNOME Terminal
- **macOS:** iTerm2
- **Windows:** Windows Terminal

### Font Recommendations

Monospace fonts that work well with Gesh:
- JetBrains Mono
- Fira Code
- Source Code Pro
- Cascadia Code
- Hack

---

## Future: Custom Themes

Custom theme support via YAML files is planned for a future release.

Example of planned format:
```yaml
# ~/.config/gesh/themes/mytheme.yaml
name: mytheme

colors:
  background: "#1e1e2e"
  foreground: "#cdd6f4"
  
  ui:
    header_bg: "#181825"
    header_fg: "#f38ba8"
    status_bg: "#313244"
    line_number: "#6c7086"
    selection: "#45475a"
  
  syntax:
    keyword: "#cba6f7"
    string: "#a6e3a1"
    number: "#fab387"
    comment: "#6c7086"
    function: "#89b4fa"
```

---

## See Also

- [CONFIG.md](CONFIG.md) - Configuration options
- [KEYBINDINGS.md](KEYBINDINGS.md) - Keyboard shortcuts

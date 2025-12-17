# GESH Keybindings

## File Operations

| Action | Shortcut | Description |
|--------|----------|-------------|
| New File | `Ctrl+Alt+N` | Create new empty buffer |
| New Tab | `Ctrl+T` | Open new tab |
| Close Tab | `Ctrl+W` | Close current tab (search if only one tab) |
| Open File | `Ctrl+O` | Open file dialog |
| Save | `Ctrl+S` | Save current file |
| Save As | `Ctrl+Shift+S` | Save with new name |
| Exit | `Ctrl+X` | Exit editor (or cut if selecting) |

---

## Tab Management

| Action | Shortcut | Description |
|--------|----------|-------------|
| New Tab | `Ctrl+T` | Create new empty tab |
| Close Tab | `Ctrl+W` | Close current tab |
| Next Tab | `Ctrl+Tab` | Switch to next tab |
| Next Tab | `Ctrl+PageDown` | Switch to next tab (alternate) |
| Previous Tab | `Ctrl+Shift+Tab` | Switch to previous tab |
| Previous Tab | `Ctrl+PageUp` | Switch to previous tab (alternate) |
| Go to Tab 1-9 | `Alt+1` to `Alt+9` | Switch to specific tab |

---

## Macro Recording

| Action | Shortcut | Description |
|--------|----------|-------------|
| Toggle Recording | `Ctrl+M` | Start/stop recording macro |
| Play Macro | `Ctrl+Shift+M` | Play recorded macro |

When recording, "REC" is shown in the status bar instead of "INS".

---

## Split View

| Action | Shortcut | Description |
|--------|----------|-------------|
| Horizontal Split | `Ctrl+\` | Split side by side (left/right) |
| Vertical Split | `Ctrl+Shift+-` | Split stacked (top/bottom) |
| Close Split | `Ctrl+Shift+\` | Close split view |
| Focus Left/Top | `Alt+Left` or `Alt+H` | Switch to left/top pane |
| Focus Right/Bottom | `Alt+Right` or `Alt+L` | Switch to right/bottom pane |

---

## Navigation

### Basic Movement

| Action | Shortcut | Alternative |
|--------|----------|-------------|
| Move Up | `↑` | `Ctrl+P` |
| Move Down | `↓` | `Ctrl+N` |
| Move Left | `←` | `Ctrl+B` |
| Move Right | `→` | `Ctrl+F` |

### Line Navigation

| Action | Shortcut | Alternative |
|--------|----------|-------------|
| Line Start | `Home` | `Ctrl+A` |
| Line End | `End` | `Ctrl+E` |

### Word Navigation

| Action | Shortcut |
|--------|----------|
| Previous Word | `Ctrl+←` |
| Next Word | `Ctrl+→` |

### Document Navigation

| Action | Shortcut |
|--------|----------|
| Document Start | `Ctrl+Home` |
| Document End | `Ctrl+End` |
| Page Up | `PageUp` |
| Page Down | `PageDown` |
| Go to Line | `Ctrl+G` |

---

## Editing

### Basic Editing

| Action | Shortcut |
|--------|----------|
| Delete Char (back) | `Backspace` |
| Delete Char (forward) | `Delete` |
| Delete Line | `Ctrl+K` |
| Cut Line | `Ctrl+U` |
| New Line | `Enter` |
| Insert Tab | `Tab` |

### Undo/Redo

| Action | Shortcut |
|--------|----------|
| Undo | `Ctrl+Z` |
| Redo | `Ctrl+Y` |

---

## Selection

| Action | Shortcut | Description |
|--------|----------|-------------|
| Toggle Selection | `Ctrl+Space` | Start/stop selection mode |
| Select All | `Ctrl+A` (2x) | Press twice to select all |
| Extend Up | `Shift+↑` | Select while moving up |
| Extend Down | `Shift+↓` | Select while moving down |
| Extend Left | `Shift+←` | Select while moving left |
| Extend Right | `Shift+→` | Select while moving right |

---

## Clipboard

| Action | Shortcut | Description |
|--------|----------|-------------|
| Copy | `Ctrl+C` | Copy selection to clipboard |
| Cut | `Ctrl+X` | Cut selection (when selecting) |
| Paste | `Ctrl+V` | Paste from clipboard |

> **Note:** `Ctrl+X` exits the editor when nothing is selected, and cuts when there's a selection.

---

## Search & Replace

| Action | Shortcut |
|--------|----------|
| Search | `Ctrl+W` |
| Find Next | `F3` |
| Find Previous | `Shift+F3` |
| Replace One | `Ctrl+R` |
| Replace All | `Ctrl+Shift+R` |

### In Search Mode

| Action | Shortcut |
|--------|----------|
| Confirm | `Enter` |
| Cancel | `Esc` |
| Next Match | `F3` |
| Previous Match | `Shift+F3` |

---

## Mouse Support

| Action | Mouse |
|--------|-------|
| Position Cursor | Left Click |
| Scroll Up | Scroll Wheel Up |
| Scroll Down | Scroll Wheel Down |

---

## Mode-Specific Keys

### Normal Mode

Standard editing keybindings apply.

### Search Mode (`Ctrl+W`)

| Key | Action |
|-----|--------|
| `Enter` | Find and highlight matches |
| `Esc` | Cancel search |
| `F3` | Next match |
| `Shift+F3` | Previous match |

### Replace Mode (`Ctrl+R`)

| Key | Action |
|-----|--------|
| `Enter` | Enter replacement text |
| `Y` | Replace current match |
| `N` | Skip to next match |
| `Esc` | Cancel |

### Go to Line Mode (`Ctrl+G`)

| Key | Action |
|-----|--------|
| `Enter` | Go to entered line number |
| `Esc` | Cancel |

### Quit Confirmation Mode

| Key | Action |
|-----|--------|
| `Y` | Save and quit |
| `N` | Quit without saving |
| `Esc` | Cancel quit |

---

## Quick Reference Card

```
┌─────────────────────────────────────────────────────────────┐
│                    GESH QUICK REFERENCE                     │
├─────────────────────────────────────────────────────────────┤
│  FILE        │  EDIT         │  NAVIGATE      │  SEARCH    │
│  ^T New Tab  │  ^Z Undo      │  ^G Goto       │  ^W Find   │
│  ^O Open     │  ^Y Redo      │  ^Home Start   │  F3 Next   │
│  ^S Save     │  ^K Del Line  │  ^End End      │  ^R Replace│
│  ^X Exit     │  ^U Cut Line  │  PgUp/Dn Page  │            │
├─────────────────────────────────────────────────────────────┤
│  TABS               │  CLIPBOARD                            │
│  ^Tab Next Tab      │  ^C Copy                              │
│  ^Shift+Tab Prev    │  ^X Cut (with selection)              │
│  Alt+1-9 Go to Tab  │  ^V Paste                             │
└─────────────────────────────────────────────────────────────┘
```

---

## Nano Compatibility

Gesh is designed to be familiar to nano users:

| Nano | Gesh | Notes |
|------|------|-------|
| `^G` Help | `^G` Goto | Different function |
| `^X` Exit | `^X` Exit/Cut | Same |
| `^O` Write | `^S` Save | Different key |
| `^R` Read | `^O` Open | Different key |
| `^W` Search | `^W` Search | Same |
| `^K` Cut | `^K` Delete Line | Similar |
| `^U` Paste | `^U` Cut Line | Different |
| `^C` Position | `^C` Copy | Different |

---

## Tips

1. **Double `Ctrl+A`**: First press goes to line start, second press selects all text.

2. **`Ctrl+X` dual function**: Without selection it exits, with selection it cuts.

3. **Auto-indent**: When you press Enter, indentation from the current line is preserved.

4. **Scroll padding**: Cursor stays 5 lines away from edges when scrolling.

5. **Mouse**: Click anywhere to position cursor, scroll wheel to navigate.

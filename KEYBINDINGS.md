# GESH Keybindings (Nano Compatible)

GESH uses keyboard shortcuts compatible with GNU nano. Users familiar with nano will feel right at home.

## File Operations

| Action | Shortcut | Description |
|--------|----------|-------------|
| Exit | `Ctrl+X` | Exit (prompts to save if modified) |
| Write Out (Save) | `Ctrl+O` | Save current file |
| Read File | `Ctrl+R` | Insert file at cursor |
| Help | `Ctrl+G` | Toggle help bar visibility |

---

## Editing

| Action | Shortcut | Description |
|--------|----------|-------------|
| Cut Line/Selection | `Ctrl+K` | Cut current line or selection |
| Paste (Uncut) | `Ctrl+U` | Paste from clipboard |
| Copy Line/Selection | `Alt+6` | Copy current line or selection |
| Undo | `Alt+U` | Undo last action |
| Redo | `Alt+E` | Redo last undone action |
| Delete Char Left | `Backspace` / `Ctrl+H` | Delete character before cursor |
| Delete Char Right | `Delete` / `Ctrl+D` | Delete character under cursor |
| Delete Word Left | `Alt+Backspace` | Delete word to the left |
| Delete Word Right | `Ctrl+Delete` | Delete word to the right |
| New Line | `Enter` / `Ctrl+M` | Insert newline with auto-indent |
| Insert Tab | `Tab` / `Ctrl+I` | Insert 4 spaces |

---

## Navigation

| Action | Shortcut | Description |
|--------|----------|-------------|
| Up | `↑` / `Ctrl+P` | Move cursor up |
| Down | `↓` / `Ctrl+N` | Move cursor down |
| Left | `←` / `Ctrl+B` | Move cursor left |
| Right | `→` / `Ctrl+F` | Move cursor right |
| Line Start | `Home` / `Ctrl+A` | Go to beginning of line |
| Line End | `End` / `Ctrl+E` | Go to end of line |
| Word Left | `Ctrl+←` / `Alt+Space` | Move to previous word |
| Word Right | `Ctrl+→` / `Ctrl+Space` | Move to next word |
| Page Up | `PageUp` / `Ctrl+Y` | Scroll up one page |
| Page Down | `PageDown` / `Ctrl+V` | Scroll down one page |
| File Start | `Ctrl+Home` / `Alt+\` | Go to beginning of file |
| File End | `Ctrl+End` / `Alt+/` | Go to end of file |
| Go to Line | `Ctrl+_` / `Alt+G` | Jump to specific line |

---

## Search & Replace

| Action | Shortcut | Description |
|--------|----------|-------------|
| Search (Where Is) | `Ctrl+W` | Start search |
| Search Next | `Alt+W` / `F3` | Find next match |
| Search Prev | `Ctrl+Q` / `Shift+F3` | Find previous match |
| Replace | `Ctrl+\` / `Alt+R` | Search and replace |

---

## Selection (Mark)

| Action | Shortcut | Description |
|--------|----------|-------------|
| Set Mark | `Alt+A` / `Ctrl+6` | Start/toggle selection |
| Extend Up | `Shift+↑` | Select while moving up |
| Extend Down | `Shift+↓` | Select while moving down |
| Extend Left | `Shift+←` | Select while moving left |
| Extend Right | `Shift+→` | Select while moving right |

---

## Display

| Action | Shortcut | Description |
|--------|----------|-------------|
| Cursor Position | `Ctrl+C` | Show current position info |
| Toggle Line Numbers | `Alt+N` | Show/hide line numbers |
| Toggle Help | `Ctrl+G` / `Alt+X` | Show/hide help bar |
| Refresh Screen | `Ctrl+L` | Redraw screen |
| Toggle Insert/Overwrite | `Insert` | Switch INS/OVR mode |

---

## Tab Management (Extension)

| Action | Shortcut | Description |
|--------|----------|-------------|
| New Tab | `Ctrl+T` | Create new empty tab |
| Next Tab | `Ctrl+Tab` / `Ctrl+PageDown` | Switch to next tab |
| Previous Tab | `Ctrl+Shift+Tab` / `Ctrl+PageUp` | Switch to previous tab |

---

## Split View (Extension)

| Action | Shortcut | Description |
|--------|----------|-------------|
| Horizontal Split | `Alt+\\` | Split side by side |
| Vertical Split | `Alt+-` | Split top/bottom |
| Close Split | `Alt+C` | Close current split |
| Focus Left/Up | `Alt+Left` / `Alt+H` | Switch to left/top pane |
| Focus Right/Down | `Alt+Right` / `Alt+L` | Switch to right/bottom pane |

---

## Macro Recording (Extension)

| Action | Shortcut | Description |
|--------|----------|-------------|
| Record Macro | `F4` | Start/stop recording |
| Play Macro | `F5` | Play recorded macro |

---

## Mode-Specific Keys

### Save Confirmation (Ctrl+X with unsaved changes)

| Key | Action |
|-----|--------|
| `Y` | Save and exit |
| `N` | Exit without saving |
| `C` / `Esc` | Cancel, return to editing |

### Search Mode (Ctrl+W)

| Key | Action |
|-----|--------|
| `Enter` | Find and highlight matches |
| `Esc` | Cancel search |
| `Alt+W` / `F3` | Next match |
| `Ctrl+Q` / `Shift+F3` | Previous match |

### Go to Line Mode (Ctrl+_ / Alt+G)

| Key | Action |
|-----|--------|
| `Enter` | Go to entered line number |
| `Esc` | Cancel |

---

## Quick Reference Card

```
┌─────────────────────────────────────────────────────────────┐
│                    GESH NANO-STYLE SHORTCUTS                │
├─────────────────────────────────────────────────────────────┤
│  FILE         │  EDIT          │  SEARCH       │  NAV      │
│  ^O Save      │  ^K Cut        │  ^W Search    │  ^Y PgUp  │
│  ^R Read      │  ^U Paste      │  M-W Next     │  ^V PgDn  │
│  ^X Exit      │  M-6 Copy      │  ^\ Replace   │  ^_ Goto  │
│  ^G Help      │  M-U Undo      │  ^Q Prev      │  M-\ Top  │
│               │  M-E Redo      │               │  M-/ End  │
├─────────────────────────────────────────────────────────────┤
│  MOVE         │  DELETE        │  DISPLAY      │  MARK     │
│  ^P/^N Up/Dn  │  ^H Backspace  │  ^C Position  │  M-A Mark │
│  ^B/^F Lt/Rt  │  ^D Delete     │  M-N LineNums │  ^6 Mark  │
│  ^A/^E Home/E │  M-Bs Word←    │  ^L Refresh   │           │
└─────────────────────────────────────────────────────────────┘
```

---

## Differences from GNU nano

While GESH aims for nano compatibility, there are some differences:

| Feature | nano | GESH |
|---------|------|------|
| Spell Check | `Ctrl+T` | Not available (new tab instead) |
| Justify | `Ctrl+J` | Not implemented |
| Where Was (back search) | `Ctrl+Q` | Previous match |
| Execute Command | `Ctrl+T` | New tab |
| Browser | `Ctrl+B` | Move left |

---

## Tips

1. **Auto-indent**: When you press Enter, indentation from the current line is preserved.

2. **Scroll padding**: Cursor stays 5 lines away from edges when scrolling.

3. **Mouse**: Click anywhere to position cursor, scroll wheel to navigate.

4. **Syntax highlighting**: Automatic for 55+ languages based on file extension.

5. **Themes**: Use `--theme` flag or config file to change colors (dark, light, monokai, dracula, gruvbox).

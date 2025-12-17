# Gesh Plugin System

This document describes the plugin architecture for Gesh text editor. The plugin system enables users to extend editor functionality through Lua scripts.

## Overview

Gesh uses a Lua-based plugin system inspired by Neovim. Plugins can:

- Register custom commands
- Define keyboard shortcuts
- Hook into editor events
- Modify buffer content
- Execute external programs
- Customize the status bar

```
┌─────────────────────────────────────────────────────────────────┐
│                      GESH PLUGIN ARCHITECTURE                    │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  ┌──────────────┐     ┌──────────────┐     ┌──────────────┐    │
│  │    Gesh      │     │   Plugin     │     │    Lua VM    │    │
│  │   Core       │◄───►│   Manager    │◄───►│  (gopher-    │    │
│  │              │     │              │     │   lua)       │    │
│  └──────────────┘     └──────────────┘     └──────────────┘    │
│         │                    │                    │             │
│         │                    ▼                    ▼             │
│         │             ┌──────────────┐     ┌──────────────┐    │
│         │             │    Hook      │     │   Plugin     │    │
│         └────────────►│   Registry   │     │   Scripts    │    │
│                       │              │     │   (.lua)     │    │
│                       └──────────────┘     └──────────────┘    │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

---

## Directory Structure

```
~/.config/gesh/
├── gesh.yaml                 # Main configuration
└── plugins/
    ├── enabled.yaml          # List of enabled plugins
    │
    ├── comment.lua           # Single-file plugin
    ├── autopairs.lua         # Single-file plugin
    │
    └── git/                   # Multi-file plugin
        ├── plugin.yaml       # Plugin manifest
        ├── init.lua          # Entry point
        └── commands.lua      # Additional modules
```

### enabled.yaml

```yaml
# List of enabled plugins
plugins:
  - comment
  - autopairs
  - surround
  - git
  - whitespace

# Plugin-specific settings
settings:
  autopairs:
    enabled_pairs: ["()", "[]", "{}", '""']
  git:
    show_branch: true
    show_status: true
```

### Plugin Manifest (plugin.yaml)

For multi-file plugins:

```yaml
name: git
version: 1.0.0
description: Git integration for Gesh
author: Your Name
license: MIT

# Entry point
main: init.lua

# Dependencies (other plugins)
dependencies: []

# Activation events
activation:
  - onStartup           # Load on editor start
  - onLanguage:go       # Load for specific language
  - onCommand:git.*     # Load when git command invoked

# Contributions
contributes:
  commands:
    - id: git.status
      title: Show Git Status
    - id: git.blame
      title: Show Git Blame
    - id: git.diff
      title: Show Git Diff
  
  keybindings:
    - key: ctrl+g s
      command: git.status
    - key: ctrl+g b
      command: git.blame
  
  statusBar:
    - id: git.branch
      position: right
      priority: 100
```

---

## Lua API Reference

### Global Object: `gesh`

The `gesh` global object provides access to all editor functionality.

### Event Hooks

#### `gesh.on(event, callback)`

Register a callback for an editor event.

```lua
-- Available events:
-- "buffer_open"      - Buffer opened
-- "buffer_close"     - Buffer closed
-- "buffer_save"      - Before buffer save
-- "buffer_saved"     - After buffer save
-- "cursor_move"      - Cursor position changed
-- "text_change"      - Buffer content changed
-- "mode_change"      - Editor mode changed
-- "key_press"        - Key pressed (can intercept)
-- "startup"          - Editor started
-- "shutdown"         - Editor closing

gesh.on("buffer_open", function(ctx)
    print("Opened: " .. ctx.buffer.path)
end)

gesh.on("buffer_save", function(ctx)
    -- Return false to cancel save
    if ctx.buffer.path:match("%.min%.js$") then
        gesh.message("Cannot save minified files!")
        return false
    end
    return true
end)

gesh.on("key_press", function(ctx)
    -- ctx.key = pressed key string
    -- Return true to consume the key (prevent default handling)
    if ctx.key == "ctrl+d" then
        duplicate_line()
        return true
    end
    return false
end)
```

#### Event Context Objects

```lua
-- buffer_open, buffer_close, buffer_save, buffer_saved
ctx = {
    buffer = {
        path = "/path/to/file.go",
        filename = "file.go",
        language = "go",
        modified = false,
        readonly = false,
        line_count = 150,
        encoding = "utf-8",
        line_ending = "lf"
    }
}

-- cursor_move
ctx = {
    buffer = { ... },
    cursor = {
        line = 10,      -- 0-indexed
        column = 5,     -- 0-indexed
        offset = 234    -- Absolute position
    },
    previous = {
        line = 9,
        column = 5,
        offset = 200
    }
}

-- text_change
ctx = {
    buffer = { ... },
    change = {
        type = "insert",  -- "insert", "delete", "replace"
        start_line = 10,
        start_column = 5,
        end_line = 10,
        end_column = 15,
        text = "inserted text"
    }
}

-- key_press
ctx = {
    key = "ctrl+shift+k",  -- Normalized key string
    raw = { ... }          -- Raw key event data
}

-- mode_change
ctx = {
    mode = "search",       -- Current mode
    previous = "normal"    -- Previous mode
}
```

### Commands

#### `gesh.command(name, callback)`

Register a custom command.

```lua
gesh.command("hello", function(args)
    local name = args[1] or "World"
    gesh.message("Hello, " .. name .. "!")
end)

-- Usage: :hello or :hello John
```

#### `gesh.run(command)`

Execute a registered command.

```lua
gesh.run("hello John")
gesh.run("save")
gesh.run("goto 42")
```

### Keybindings

#### `gesh.keymap(key, command_or_function)`

Map a key combination to a command or function.

```lua
-- Map to command
gesh.keymap("ctrl+shift+d", "duplicate_line")

-- Map to function
gesh.keymap("ctrl+shift+k", function()
    gesh.delete_line()
end)

-- Map to command with arguments
gesh.keymap("ctrl+/", "toggle_comment")
```

#### Key Format

```
Modifiers: ctrl, shift, alt, meta (cmd on macOS)
Keys: a-z, 0-9, f1-f12, enter, tab, space, backspace, delete,
      home, end, pageup, pagedown, up, down, left, right,
      escape, insert

Examples:
  "ctrl+s"
  "ctrl+shift+p"
  "alt+up"
  "f5"
  "ctrl+k ctrl+c"  -- Chord (press ctrl+k, then ctrl+c)
```

### Buffer Operations

#### `gesh.current_buffer()`

Get the current buffer information.

```lua
local buf = gesh.current_buffer()
print(buf.path)        -- "/path/to/file.go"
print(buf.filename)    -- "file.go"
print(buf.language)    -- "go"
print(buf.modified)    -- true/false
print(buf.line_count)  -- 150
```

#### `gesh.get_line([line_number])`

Get line content. Default: current line.

```lua
local current = gesh.get_line()
local line_10 = gesh.get_line(10)  -- 0-indexed
```

#### `gesh.set_line(content, [line_number])`

Set line content. Default: current line.

```lua
gesh.set_line("new content")
gesh.set_line("content", 10)
```

#### `gesh.get_lines(start_line, end_line)`

Get multiple lines.

```lua
local lines = gesh.get_lines(0, 10)
for i, line in ipairs(lines) do
    print(i, line)
end
```

#### `gesh.insert(text)`

Insert text at cursor position.

```lua
gesh.insert("Hello, World!")
gesh.insert("\n")  -- New line
```

#### `gesh.delete([count])`

Delete characters. Positive = forward, negative = backward.

```lua
gesh.delete(1)    -- Delete char after cursor
gesh.delete(-1)   -- Delete char before cursor (backspace)
gesh.delete(10)   -- Delete 10 chars forward
```

#### `gesh.delete_line([line_number])`

Delete a line. Default: current line.

```lua
gesh.delete_line()
gesh.delete_line(5)
```

#### `gesh.get_selection()`

Get selected text.

```lua
local sel = gesh.get_selection()
if sel ~= "" then
    print("Selected: " .. sel)
end
```

#### `gesh.set_selection(start_line, start_col, end_line, end_col)`

Set selection range.

```lua
gesh.set_selection(0, 0, 0, 10)  -- Select first 10 chars of first line
```

#### `gesh.replace_selection(text)`

Replace selected text.

```lua
local sel = gesh.get_selection()
gesh.replace_selection(sel:upper())
```

#### `gesh.get_text()`

Get entire buffer content.

```lua
local content = gesh.get_text()
```

#### `gesh.set_text(content)`

Replace entire buffer content.

```lua
gesh.set_text("new content")
```

### Cursor Operations

#### `gesh.cursor()`

Get cursor position.

```lua
local pos = gesh.cursor()
print(pos.line, pos.column, pos.offset)
```

#### `gesh.goto(line, [column])`

Move cursor to position.

```lua
gesh.goto(10)       -- Go to line 10
gesh.goto(10, 5)    -- Go to line 10, column 5
```

#### `gesh.move(direction, [count])`

Move cursor relatively.

```lua
gesh.move("up", 5)
gesh.move("down", 1)
gesh.move("left", 10)
gesh.move("right", 1)
gesh.move("word_left")
gesh.move("word_right")
gesh.move("line_start")
gesh.move("line_end")
gesh.move("file_start")
gesh.move("file_end")
```

### UI Operations

#### `gesh.message(text, [type])`

Show message in status bar.

```lua
gesh.message("File saved!")
gesh.message("Error occurred!", "error")
gesh.message("Warning!", "warning")
gesh.message("Info", "info")
```

#### `gesh.input(prompt, [default])`

Show input prompt and get user input.

```lua
local name = gesh.input("Enter name: ")
local value = gesh.input("Enter value: ", "default")

if name then
    print("User entered: " .. name)
else
    print("User cancelled")
end
```

#### `gesh.confirm(message)`

Show yes/no confirmation dialog.

```lua
if gesh.confirm("Delete this file?") then
    -- User confirmed
end
```

#### `gesh.select(items, [prompt])`

Show selection popup.

```lua
local items = {"Option 1", "Option 2", "Option 3"}
local choice = gesh.select(items, "Choose an option:")

if choice then
    print("Selected: " .. choice)  -- Returns selected string
end
```

#### `gesh.popup(content, [options])`

Show popup window.

```lua
gesh.popup("This is a popup message")

gesh.popup([[
Multi-line
popup
content
]], {
    title = "Popup Title",
    width = 40,
    height = 10,
    position = "center"  -- "center", "cursor", "top", "bottom"
})
```

#### `gesh.statusbar_set(id, content)`

Set status bar section content.

```lua
gesh.statusbar_set("git", " main")
gesh.statusbar_set("custom", "[Custom]")
```

### External Commands

#### `gesh.exec(command, [args...])`

Execute external command and return output.

```lua
local output = gesh.exec("git", "status", "--short")
print(output)

local formatted = gesh.exec("gofmt", "-s", gesh.current_buffer().path)
```

#### `gesh.exec_async(command, callback, [args...])`

Execute command asynchronously.

```lua
gesh.exec_async("go", function(output, error, exit_code)
    if exit_code == 0 then
        gesh.message("Build successful!")
    else
        gesh.message("Build failed: " .. error, "error")
    end
end, "build", "./...")
```

#### `gesh.shell(command)`

Execute shell command (with shell interpretation).

```lua
local result = gesh.shell("ls -la | grep go")
```

### File Operations

#### `gesh.read_file(path)`

Read file content.

```lua
local content = gesh.read_file("/path/to/file")
```

#### `gesh.write_file(path, content)`

Write file content.

```lua
gesh.write_file("/path/to/file", "content")
```

#### `gesh.file_exists(path)`

Check if file exists.

```lua
if gesh.file_exists("go.mod") then
    print("Go module found")
end
```

#### `gesh.open(path)`

Open file in new buffer/tab.

```lua
gesh.open("/path/to/file.go")
```

### Utility Functions

#### `gesh.log(message)`

Write to plugin log file.

```lua
gesh.log("Debug: something happened")
-- Writes to ~/.config/gesh/plugins.log
```

#### `gesh.config(key, [default])`

Get plugin configuration value.

```lua
local tab_size = gesh.config("editor.tab_size", 4)
local my_setting = gesh.config("plugins.myplugin.option", "default")
```

#### `gesh.set_config(key, value)`

Set configuration value (runtime only).

```lua
gesh.set_config("plugins.myplugin.option", "new_value")
```

---

## Example Plugins

### 1. Toggle Comment

```lua
-- comment.lua
-- Toggle line comments for various languages

local comment_chars = {
    go = "//",
    python = "#",
    lua = "--",
    javascript = "//",
    typescript = "//",
    c = "//",
    cpp = "//",
    rust = "//",
    java = "//",
    shell = "#",
    bash = "#",
    yaml = "#",
    toml = "#",
    ruby = "#",
    perl = "#",
    php = "//",
    html = "<!--",
    css = "/*",
}

local comment_end = {
    html = "-->",
    css = "*/",
}

local function toggle_comment()
    local buf = gesh.current_buffer()
    local char = comment_chars[buf.language]
    
    if not char then
        gesh.message("No comment style for " .. buf.language, "warning")
        return
    end
    
    local line = gesh.get_line()
    local end_char = comment_end[buf.language]
    
    -- Check if already commented
    local pattern = "^(%s*)" .. char:gsub("([%-%.%+%[%]%(%)%$%^%%%?%*])", "%%%1")
    
    if line:match(pattern) then
        -- Uncomment
        if end_char then
            line = line:gsub(pattern .. "%s?", "%1")
            line = line:gsub("%s?" .. end_char:gsub("([%-%.%+%[%]%(%)%$%^%%%?%*])", "%%%1") .. "$", "")
        else
            line = line:gsub(pattern .. "%s?", "%1")
        end
    else
        -- Comment
        local indent = line:match("^(%s*)")
        local content = line:sub(#indent + 1)
        if end_char then
            line = indent .. char .. " " .. content .. " " .. end_char
        else
            line = indent .. char .. " " .. content
        end
    end
    
    gesh.set_line(line)
end

gesh.command("toggle_comment", toggle_comment)
gesh.keymap("ctrl+/", "toggle_comment")
```

### 2. Auto Pairs

```lua
-- autopairs.lua
-- Automatically close brackets and quotes

local pairs = {
    ["("] = ")",
    ["["] = "]",
    ["{"] = "}",
    ['"'] = '"',
    ["'"] = "'",
    ["`"] = "`",
}

local skip_pairs = {
    [")"] = true,
    ["]"] = true,
    ["}"] = true,
    ['"'] = true,
    ["'"] = true,
    ["`"] = true,
}

gesh.on("key_press", function(ctx)
    local key = ctx.key
    
    -- Auto-close opening brackets
    local close = pairs[key]
    if close then
        gesh.insert(key .. close)
        gesh.move("left", 1)
        return true
    end
    
    -- Skip over closing brackets if already there
    if skip_pairs[key] then
        local line = gesh.get_line()
        local col = gesh.cursor().column
        if line:sub(col + 1, col + 1) == key then
            gesh.move("right", 1)
            return true
        end
    end
    
    -- Handle backspace - delete pair
    if key == "backspace" then
        local line = gesh.get_line()
        local col = gesh.cursor().column
        if col > 0 then
            local before = line:sub(col, col)
            local after = line:sub(col + 1, col + 1)
            if pairs[before] == after then
                gesh.delete(1)   -- Delete closing
                gesh.delete(-1)  -- Delete opening
                return true
            end
        end
    end
    
    return false
end)
```

### 3. Surround

```lua
-- surround.lua
-- Surround selection with characters

local pair_map = {
    ["("] = "()",
    [")"] = "()",
    ["["] = "[]",
    ["]"] = "[]",
    ["{"] = "{}",
    ["}"] = "{}",
    ["<"] = "<>",
    [">"] = "<>",
    ['"'] = '""',
    ["'"] = "''",
    ["`"] = "``",
}

local function surround(char)
    local sel = gesh.get_selection()
    if sel == "" then
        gesh.message("No selection", "warning")
        return
    end
    
    local wrap = pair_map[char]
    if not wrap then
        wrap = char .. char
    end
    
    local open = wrap:sub(1, 1)
    local close = wrap:sub(2, 2)
    
    gesh.replace_selection(open .. sel .. close)
end

local function delete_surround()
    -- Find surrounding pair and delete
    local line = gesh.get_line()
    local col = gesh.cursor().column
    
    -- Simple implementation: find matching pair on current line
    for open, close in pairs({["("]=")", ["["]="]", ["{"]="}",  ['"']='"', ["'"]="'"}) do
        local start_pos = line:find(open, 1, true)
        local end_pos = line:find(close, col + 1, true)
        
        if start_pos and end_pos and start_pos < col and end_pos > col then
            local before = line:sub(1, start_pos - 1)
            local content = line:sub(start_pos + 1, end_pos - 1)
            local after = line:sub(end_pos + 1)
            gesh.set_line(before .. content .. after)
            return
        end
    end
    
    gesh.message("No surrounding pair found", "warning")
end

gesh.command("surround", function(args)
    local char = args[1] or '"'
    surround(char)
end)

gesh.command("delete_surround", delete_surround)

-- Keymaps
gesh.keymap("ctrl+s (", function() surround("(") end)
gesh.keymap("ctrl+s [", function() surround("[") end)
gesh.keymap("ctrl+s {", function() surround("{") end)
gesh.keymap('ctrl+s "', function() surround('"') end)
gesh.keymap("ctrl+s '", function() surround("'") end)
gesh.keymap("ctrl+s d", "delete_surround")
```

### 4. Duplicate Line

```lua
-- duplicate.lua
-- Duplicate current line or selection

local function duplicate_line()
    local sel = gesh.get_selection()
    
    if sel ~= "" then
        -- Duplicate selection
        local pos = gesh.cursor()
        gesh.move("right", #sel)  -- Move to end of selection
        gesh.insert(sel)
    else
        -- Duplicate line
        local line = gesh.get_line()
        gesh.move("line_end")
        gesh.insert("\n" .. line)
    end
end

gesh.command("duplicate_line", duplicate_line)
gesh.keymap("ctrl+shift+d", "duplicate_line")
```

### 5. Whitespace Cleanup

```lua
-- whitespace.lua
-- Clean up whitespace issues

local function trim_trailing()
    local line_count = gesh.current_buffer().line_count
    local trimmed = 0
    
    for i = 0, line_count - 1 do
        local line = gesh.get_line(i)
        local cleaned = line:gsub("%s+$", "")
        if cleaned ~= line then
            gesh.set_line(cleaned, i)
            trimmed = trimmed + 1
        end
    end
    
    gesh.message("Trimmed " .. trimmed .. " lines")
end

local function remove_blank_lines()
    local content = gesh.get_text()
    local cleaned = content:gsub("\n\n\n+", "\n\n")
    gesh.set_text(cleaned)
    gesh.message("Removed extra blank lines")
end

local function tabs_to_spaces()
    local tab_size = gesh.config("editor.tab_size", 4)
    local spaces = string.rep(" ", tab_size)
    
    local line_count = gesh.current_buffer().line_count
    local converted = 0
    
    for i = 0, line_count - 1 do
        local line = gesh.get_line(i)
        local cleaned = line:gsub("\t", spaces)
        if cleaned ~= line then
            gesh.set_line(cleaned, i)
            converted = converted + 1
        end
    end
    
    gesh.message("Converted tabs in " .. converted .. " lines")
end

gesh.command("trim_trailing", trim_trailing)
gesh.command("remove_blank_lines", remove_blank_lines)
gesh.command("tabs_to_spaces", tabs_to_spaces)

-- Auto trim on save (optional)
gesh.on("buffer_save", function(ctx)
    if gesh.config("plugins.whitespace.trim_on_save", false) then
        trim_trailing()
    end
    return true
end)
```

### 6. Git Integration

```lua
-- git/init.lua
-- Git integration for Gesh

local M = {}

-- Get current branch
function M.get_branch()
    local branch = gesh.exec("git", "rev-parse", "--abbrev-ref", "HEAD")
    return branch:gsub("%s+$", "")
end

-- Check if in git repo
function M.is_git_repo()
    local result = gesh.exec("git", "rev-parse", "--git-dir")
    return result ~= ""
end

-- Update status bar with branch
local function update_statusbar()
    if M.is_git_repo() then
        local branch = M.get_branch()
        local status = gesh.exec("git", "status", "--porcelain")
        local modified = status ~= "" and "*" or ""
        gesh.statusbar_set("git", " " .. branch .. modified)
    else
        gesh.statusbar_set("git", "")
    end
end

-- Commands
gesh.command("git_status", function()
    local status = gesh.exec("git", "status", "--short")
    if status == "" then
        status = "Working tree clean"
    end
    gesh.popup(status, { title = "Git Status" })
end)

gesh.command("git_diff", function()
    local buf = gesh.current_buffer()
    local diff = gesh.exec("git", "diff", buf.path)
    if diff == "" then
        gesh.message("No changes")
    else
        gesh.popup(diff, { title = "Git Diff: " .. buf.filename })
    end
end)

gesh.command("git_blame", function()
    local buf = gesh.current_buffer()
    local line = gesh.cursor().line + 1
    local blame = gesh.exec("git", "blame", "-L", line .. "," .. line, "--", buf.path)
    gesh.message(blame:gsub("%s+$", ""))
end)

gesh.command("git_log", function()
    local log = gesh.exec("git", "log", "--oneline", "-20")
    gesh.popup(log, { title = "Git Log (last 20)" })
end)

gesh.command("git_add", function()
    local buf = gesh.current_buffer()
    gesh.exec("git", "add", buf.path)
    gesh.message("Added: " .. buf.filename)
    update_statusbar()
end)

gesh.command("git_checkout", function()
    local buf = gesh.current_buffer()
    if gesh.confirm("Discard changes to " .. buf.filename .. "?") then
        gesh.exec("git", "checkout", "--", buf.path)
        gesh.run("reload")
        gesh.message("Reverted: " .. buf.filename)
    end
end)

-- Keymaps
gesh.keymap("ctrl+g s", "git_status")
gesh.keymap("ctrl+g d", "git_diff")
gesh.keymap("ctrl+g b", "git_blame")
gesh.keymap("ctrl+g l", "git_log")
gesh.keymap("ctrl+g a", "git_add")

-- Events
gesh.on("buffer_open", function(ctx)
    update_statusbar()
end)

gesh.on("buffer_saved", function(ctx)
    update_statusbar()
end)

gesh.on("startup", function()
    update_statusbar()
end)

return M
```

### 7. Format on Save

```lua
-- format.lua
-- Auto-format files on save

local formatters = {
    go = { cmd = "gofmt", args = {"-w"} },
    python = { cmd = "black", args = {} },
    javascript = { cmd = "prettier", args = {"--write"} },
    typescript = { cmd = "prettier", args = {"--write"} },
    json = { cmd = "prettier", args = {"--write"} },
    rust = { cmd = "rustfmt", args = {} },
    c = { cmd = "clang-format", args = {"-i"} },
    cpp = { cmd = "clang-format", args = {"-i"} },
}

local function format_buffer()
    local buf = gesh.current_buffer()
    local formatter = formatters[buf.language]
    
    if not formatter then
        gesh.message("No formatter for " .. buf.language, "warning")
        return false
    end
    
    -- Save first
    gesh.run("save")
    
    -- Run formatter
    local args = {}
    for _, arg in ipairs(formatter.args) do
        table.insert(args, arg)
    end
    table.insert(args, buf.path)
    
    local output = gesh.exec(formatter.cmd, table.unpack(args))
    
    -- Reload buffer
    gesh.run("reload")
    gesh.message("Formatted with " .. formatter.cmd)
    
    return true
end

gesh.command("format", format_buffer)
gesh.keymap("ctrl+shift+f", "format")

-- Auto-format on save (configurable)
gesh.on("buffer_save", function(ctx)
    if gesh.config("plugins.format.on_save", false) then
        local buf = gesh.current_buffer()
        if formatters[buf.language] then
            -- Let save complete, then format
            gesh.defer(function()
                format_buffer()
            end)
        end
    end
    return true
end)
```

### 8. Snippets

```lua
-- snippets.lua
-- Simple snippet expansion

local snippets = {
    go = {
        ["fn"] = "func ${1:name}(${2:params}) ${3:returnType} {\n\t${0}\n}",
        ["ife"] = "if err != nil {\n\t${0}\n}",
        ["fori"] = "for ${1:i} := 0; ${1:i} < ${2:n}; ${1:i}++ {\n\t${0}\n}",
        ["forr"] = "for ${1:i}, ${2:v} := range ${3:collection} {\n\t${0}\n}",
        ["main"] = "func main() {\n\t${0}\n}",
        ["pkg"] = "package ${1:main}",
    },
    python = {
        ["def"] = "def ${1:name}(${2:params}):\n\t${0:pass}",
        ["class"] = "class ${1:Name}:\n\tdef __init__(self${2:, params}):\n\t\t${0:pass}",
        ["ifmain"] = 'if __name__ == "__main__":\n\t${0:main()}',
        ["fori"] = "for ${1:i} in range(${2:n}):\n\t${0}",
        ["forr"] = "for ${1:item} in ${2:items}:\n\t${0}",
    },
    javascript = {
        ["fn"] = "function ${1:name}(${2:params}) {\n\t${0}\n}",
        ["afn"] = "const ${1:name} = (${2:params}) => {\n\t${0}\n}",
        ["cl"] = "console.log(${0})",
        ["fori"] = "for (let ${1:i} = 0; ${1:i} < ${2:n}; ${1:i}++) {\n\t${0}\n}",
        ["forof"] = "for (const ${1:item} of ${2:items}) {\n\t${0}\n}",
    },
}

local function expand_snippet()
    local buf = gesh.current_buffer()
    local lang_snippets = snippets[buf.language]
    
    if not lang_snippets then
        return false
    end
    
    -- Get word before cursor
    local line = gesh.get_line()
    local col = gesh.cursor().column
    local word_start = col
    
    while word_start > 0 and line:sub(word_start, word_start):match("%w") do
        word_start = word_start - 1
    end
    
    local word = line:sub(word_start + 1, col)
    local snippet = lang_snippets[word]
    
    if snippet then
        -- Delete trigger word
        gesh.set_selection(gesh.cursor().line, word_start, gesh.cursor().line, col)
        
        -- Simple expansion (without tabstops for now)
        local expanded = snippet:gsub("%${%d:([^}]*)}", "%1"):gsub("%${%d}", "")
        gesh.replace_selection(expanded)
        
        return true
    end
    
    return false
end

gesh.on("key_press", function(ctx)
    if ctx.key == "tab" then
        if expand_snippet() then
            return true  -- Consumed
        end
    end
    return false
end)

-- List available snippets
gesh.command("snippets", function()
    local buf = gesh.current_buffer()
    local lang_snippets = snippets[buf.language]
    
    if not lang_snippets then
        gesh.message("No snippets for " .. buf.language)
        return
    end
    
    local list = {}
    for trigger, _ in pairs(lang_snippets) do
        table.insert(list, trigger)
    end
    table.sort(list)
    
    gesh.popup(table.concat(list, "\n"), { title = "Snippets for " .. buf.language })
end)
```

---

## Implementation Architecture

### Go Package Structure

```
gesh/
└── internal/
    └── plugin/
        ├── api.go          # GeshAPI struct and methods
        ├── manager.go      # PluginManager
        ├── loader.go       # Plugin discovery and loading
        ├── lua.go          # Lua VM setup and bindings
        ├── hooks.go        # Event hook registry
        ├── commands.go     # Command registry
        ├── keymaps.go      # Keymap registry
        └── sandbox.go      # Security sandbox for plugins
```

### Core Types

```go
// internal/plugin/api.go

package plugin

// HookType defines available hook events
type HookType string

const (
    HookBufferOpen   HookType = "buffer_open"
    HookBufferClose  HookType = "buffer_close"
    HookBufferSave   HookType = "buffer_save"
    HookBufferSaved  HookType = "buffer_saved"
    HookCursorMove   HookType = "cursor_move"
    HookTextChange   HookType = "text_change"
    HookModeChange   HookType = "mode_change"
    HookKeyPress     HookType = "key_press"
    HookStartup      HookType = "startup"
    HookShutdown     HookType = "shutdown"
)

// HookContext provides context to hook callbacks
type HookContext struct {
    Buffer   *BufferInfo
    Cursor   *CursorInfo
    Previous interface{}
    Key      string
    Change   *ChangeInfo
    Mode     string
}

// BufferInfo represents buffer state for plugins
type BufferInfo struct {
    Path       string
    Filename   string
    Language   string
    Modified   bool
    Readonly   bool
    LineCount  int
    Encoding   string
    LineEnding string
}

// Plugin represents a loaded plugin
type Plugin struct {
    Name        string
    Path        string
    Manifest    *PluginManifest
    LuaState    *lua.LState
    Enabled     bool
}

// PluginManager manages all plugins
type PluginManager struct {
    plugins     map[string]*Plugin
    hooks       map[HookType][]HookCallback
    commands    map[string]CommandFunc
    keymaps     map[string]string
    statusItems map[string]string
    luaPool     *LuaStatePool
    editor      EditorInterface
    config      *config.Config
    mu          sync.RWMutex
}
```

### Lua Bindings

```go
// internal/plugin/lua.go

func (pm *PluginManager) setupLuaAPI(L *lua.LState) {
    // Create gesh table
    gesh := L.NewTable()
    L.SetGlobal("gesh", gesh)
    
    // Events
    L.SetField(gesh, "on", L.NewFunction(pm.luaOn))
    
    // Commands
    L.SetField(gesh, "command", L.NewFunction(pm.luaCommand))
    L.SetField(gesh, "run", L.NewFunction(pm.luaRun))
    
    // Keymaps
    L.SetField(gesh, "keymap", L.NewFunction(pm.luaKeymap))
    
    // Buffer
    L.SetField(gesh, "current_buffer", L.NewFunction(pm.luaCurrentBuffer))
    L.SetField(gesh, "get_line", L.NewFunction(pm.luaGetLine))
    L.SetField(gesh, "set_line", L.NewFunction(pm.luaSetLine))
    L.SetField(gesh, "get_lines", L.NewFunction(pm.luaGetLines))
    L.SetField(gesh, "insert", L.NewFunction(pm.luaInsert))
    L.SetField(gesh, "delete", L.NewFunction(pm.luaDelete))
    L.SetField(gesh, "delete_line", L.NewFunction(pm.luaDeleteLine))
    L.SetField(gesh, "get_selection", L.NewFunction(pm.luaGetSelection))
    L.SetField(gesh, "set_selection", L.NewFunction(pm.luaSetSelection))
    L.SetField(gesh, "replace_selection", L.NewFunction(pm.luaReplaceSelection))
    L.SetField(gesh, "get_text", L.NewFunction(pm.luaGetText))
    L.SetField(gesh, "set_text", L.NewFunction(pm.luaSetText))
    
    // Cursor
    L.SetField(gesh, "cursor", L.NewFunction(pm.luaCursor))
    L.SetField(gesh, "goto", L.NewFunction(pm.luaGoto))
    L.SetField(gesh, "move", L.NewFunction(pm.luaMove))
    
    // UI
    L.SetField(gesh, "message", L.NewFunction(pm.luaMessage))
    L.SetField(gesh, "input", L.NewFunction(pm.luaInput))
    L.SetField(gesh, "confirm", L.NewFunction(pm.luaConfirm))
    L.SetField(gesh, "select", L.NewFunction(pm.luaSelect))
    L.SetField(gesh, "popup", L.NewFunction(pm.luaPopup))
    L.SetField(gesh, "statusbar_set", L.NewFunction(pm.luaStatusbarSet))
    
    // External
    L.SetField(gesh, "exec", L.NewFunction(pm.luaExec))
    L.SetField(gesh, "exec_async", L.NewFunction(pm.luaExecAsync))
    L.SetField(gesh, "shell", L.NewFunction(pm.luaShell))
    
    // Files
    L.SetField(gesh, "read_file", L.NewFunction(pm.luaReadFile))
    L.SetField(gesh, "write_file", L.NewFunction(pm.luaWriteFile))
    L.SetField(gesh, "file_exists", L.NewFunction(pm.luaFileExists))
    L.SetField(gesh, "open", L.NewFunction(pm.luaOpen))
    
    // Utility
    L.SetField(gesh, "log", L.NewFunction(pm.luaLog))
    L.SetField(gesh, "config", L.NewFunction(pm.luaConfig))
    L.SetField(gesh, "set_config", L.NewFunction(pm.luaSetConfig))
    L.SetField(gesh, "defer", L.NewFunction(pm.luaDefer))
}

// Example binding implementation
func (pm *PluginManager) luaGetLine(L *lua.LState) int {
    lineNum := L.OptInt(1, -1)
    
    var line string
    if lineNum < 0 {
        line = pm.editor.GetCurrentLine()
    } else {
        line = pm.editor.GetLine(lineNum)
    }
    
    L.Push(lua.LString(line))
    return 1
}

func (pm *PluginManager) luaOn(L *lua.LState) int {
    event := L.CheckString(1)
    callback := L.CheckFunction(2)
    
    hookType := HookType(event)
    pm.registerHook(hookType, func(ctx *HookContext) interface{} {
        // Call Lua function
        L.Push(callback)
        L.Push(pm.contextToLua(L, ctx))
        L.Call(1, 1)
        result := L.Get(-1)
        L.Pop(1)
        return pm.luaToGo(result)
    })
    
    return 0
}
```

---

## Security Considerations

### Sandboxing

```go
// internal/plugin/sandbox.go

// RestrictedEnv creates a sandboxed Lua environment
func RestrictedEnv(L *lua.LState) {
    // Remove dangerous functions
    L.SetGlobal("os", lua.LNil)
    L.SetGlobal("io", lua.LNil)
    L.SetGlobal("loadfile", lua.LNil)
    L.SetGlobal("dofile", lua.LNil)
    
    // Whitelist safe os functions
    os := L.NewTable()
    L.SetField(os, "time", L.GetGlobal("os").(*lua.LTable).RawGetString("time"))
    L.SetField(os, "date", L.GetGlobal("os").(*lua.LTable).RawGetString("date"))
    L.SetGlobal("os", os)
}

// AllowedCommands restricts which external commands can be run
var AllowedCommands = map[string]bool{
    "git":          true,
    "gofmt":        true,
    "gopls":        true,
    "black":        true,
    "prettier":     true,
    "rustfmt":      true,
    "clang-format": true,
}
```

### Configuration

```yaml
# gesh.yaml
plugins:
  # Enable/disable plugin system
  enabled: true
  
  # Security settings
  security:
    # Allow plugins to run external commands
    allow_exec: true
    
    # Allowed commands (empty = all allowed)
    allowed_commands:
      - git
      - gofmt
      - prettier
    
    # Allow plugins to read/write files outside buffer
    allow_file_access: false
    
    # Allow network access
    allow_network: false
```

---

## Dependencies

```go
// go.mod additions
require (
    github.com/yuin/gopher-lua v1.1.0
    layeh.com/gopher-luar v1.0.10  // Go struct <-> Lua conversion
)
```

---

## Development Timeline

| Phase | Task | Duration |
|-------|------|----------|
| 1 | Plugin API design & types | 1 day |
| 2 | Hook system implementation | 2 days |
| 3 | Lua VM integration | 2-3 days |
| 4 | Core API bindings (buffer, cursor) | 2-3 days |
| 5 | UI API bindings (message, popup) | 2 days |
| 6 | External command execution | 1 day |
| 7 | Plugin loader & manager | 1-2 days |
| 8 | Built-in plugins | 2 days |
| 9 | Documentation | 1 day |
| 10 | Testing & debugging | 2 days |
| **Total** | | **16-18 days** |

---

## Future Enhancements

1. **Plugin Repository** - Central plugin registry like vim-plug
2. **Plugin Dependencies** - Automatic dependency resolution
3. **Hot Reload** - Reload plugins without restart
4. **TypeScript Plugins** - Alternative to Lua using goja
5. **Plugin UI** - Custom UI components (sidebars, panels)
6. **Remote Plugins** - Plugins running in separate processes
7. **Plugin Marketplace** - In-editor plugin browser

---

## References

- [Neovim Lua API](https://neovim.io/doc/user/lua.html)
- [VS Code Extension API](https://code.visualstudio.com/api)
- [gopher-lua](https://github.com/yuin/gopher-lua)
- [Micro Editor Plugins](https://github.com/zyedidia/micro/blob/master/runtime/help/plugins.md)

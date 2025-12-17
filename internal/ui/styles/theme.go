// Package styles provides theming support for the editor.
package styles

import "github.com/charmbracelet/lipgloss"

// Theme represents a color theme for the editor.
type Theme struct {
	Name string

	// Background colors
	HeaderBg    lipgloss.Color
	StatusBg    lipgloss.Color
	HelpBg      lipgloss.Color
	EditorBg    lipgloss.Color
	SelectionBg lipgloss.Color

	// Foreground colors
	HeaderFg     lipgloss.Color
	StatusFg     lipgloss.Color
	HelpFg       lipgloss.Color
	HelpKeyFg    lipgloss.Color
	EditorFg     lipgloss.Color
	LineNumberFg lipgloss.Color
	SelectionFg  lipgloss.Color

	// Special colors
	CursorBg     lipgloss.Color
	ModifiedFlag lipgloss.Color
	LogoColor    lipgloss.Color
}

// DarkTheme is the default dark theme.
var DarkTheme = Theme{
	Name: "dark",

	HeaderBg:    lipgloss.Color("#0f3460"),
	StatusBg:    lipgloss.Color("#16213e"),
	HelpBg:      lipgloss.Color("#16213e"),
	EditorBg:    lipgloss.Color(""),
	SelectionBg: lipgloss.Color("#3a3a5a"),

	HeaderFg:     lipgloss.Color("#e94560"),
	StatusFg:     lipgloss.Color("#a0a0c0"),
	HelpFg:       lipgloss.Color("#a0a0c0"),
	HelpKeyFg:    lipgloss.Color("#e94560"),
	EditorFg:     lipgloss.Color("#eaeaea"),
	LineNumberFg: lipgloss.Color("#4a4a6a"),
	SelectionFg:  lipgloss.Color("#ffffff"),

	CursorBg:     lipgloss.Color("#ffffff"),
	ModifiedFlag: lipgloss.Color("#ff6b6b"),
	LogoColor:    lipgloss.Color("#e94560"),
}

// LightTheme is a light theme.
var LightTheme = Theme{
	Name: "light",

	HeaderBg:    lipgloss.Color("#e8e8e8"),
	StatusBg:    lipgloss.Color("#d0d0d0"),
	HelpBg:      lipgloss.Color("#d0d0d0"),
	EditorBg:    lipgloss.Color(""),
	SelectionBg: lipgloss.Color("#b0c4de"),

	HeaderFg:     lipgloss.Color("#2c3e50"),
	StatusFg:     lipgloss.Color("#333333"),
	HelpFg:       lipgloss.Color("#333333"),
	HelpKeyFg:    lipgloss.Color("#c0392b"),
	EditorFg:     lipgloss.Color("#1a1a1a"),
	LineNumberFg: lipgloss.Color("#808080"),
	SelectionFg:  lipgloss.Color("#000000"),

	CursorBg:     lipgloss.Color("#000000"),
	ModifiedFlag: lipgloss.Color("#e74c3c"),
	LogoColor:    lipgloss.Color("#c0392b"),
}

// MonokaiTheme is a Monokai-inspired theme.
var MonokaiTheme = Theme{
	Name: "monokai",

	HeaderBg:    lipgloss.Color("#272822"),
	StatusBg:    lipgloss.Color("#3e3d32"),
	HelpBg:      lipgloss.Color("#3e3d32"),
	EditorBg:    lipgloss.Color(""),
	SelectionBg: lipgloss.Color("#49483e"),

	HeaderFg:     lipgloss.Color("#f8f8f2"),
	StatusFg:     lipgloss.Color("#a6e22e"),
	HelpFg:       lipgloss.Color("#75715e"),
	HelpKeyFg:    lipgloss.Color("#f92672"),
	EditorFg:     lipgloss.Color("#f8f8f2"),
	LineNumberFg: lipgloss.Color("#75715e"),
	SelectionFg:  lipgloss.Color("#f8f8f2"),

	CursorBg:     lipgloss.Color("#f8f8f2"),
	ModifiedFlag: lipgloss.Color("#f92672"),
	LogoColor:    lipgloss.Color("#66d9ef"),
}

// DraculaTheme is a Dracula-inspired theme.
var DraculaTheme = Theme{
	Name: "dracula",

	HeaderBg:    lipgloss.Color("#282a36"),
	StatusBg:    lipgloss.Color("#44475a"),
	HelpBg:      lipgloss.Color("#44475a"),
	EditorBg:    lipgloss.Color(""),
	SelectionBg: lipgloss.Color("#44475a"),

	HeaderFg:     lipgloss.Color("#f8f8f2"),
	StatusFg:     lipgloss.Color("#6272a4"),
	HelpFg:       lipgloss.Color("#f8f8f2"),
	HelpKeyFg:    lipgloss.Color("#ff79c6"),
	EditorFg:     lipgloss.Color("#f8f8f2"),
	LineNumberFg: lipgloss.Color("#6272a4"),
	SelectionFg:  lipgloss.Color("#f8f8f2"),

	CursorBg:     lipgloss.Color("#f8f8f2"),
	ModifiedFlag: lipgloss.Color("#ff5555"),
	LogoColor:    lipgloss.Color("#bd93f9"),
}

// GruvboxTheme is a Gruvbox-inspired theme.
var GruvboxTheme = Theme{
	Name: "gruvbox",

	HeaderBg:    lipgloss.Color("#3c3836"),
	StatusBg:    lipgloss.Color("#504945"),
	HelpBg:      lipgloss.Color("#504945"),
	EditorBg:    lipgloss.Color(""),
	SelectionBg: lipgloss.Color("#504945"),

	HeaderFg:     lipgloss.Color("#ebdbb2"),
	StatusFg:     lipgloss.Color("#a89984"),
	HelpFg:       lipgloss.Color("#ebdbb2"),
	HelpKeyFg:    lipgloss.Color("#fe8019"),
	EditorFg:     lipgloss.Color("#ebdbb2"),
	LineNumberFg: lipgloss.Color("#928374"),
	SelectionFg:  lipgloss.Color("#ebdbb2"),

	CursorBg:     lipgloss.Color("#ebdbb2"),
	ModifiedFlag: lipgloss.Color("#fb4934"),
	LogoColor:    lipgloss.Color("#fabd2f"),
}

// BuiltinThemes contains all built-in themes.
var BuiltinThemes = map[string]Theme{
	"dark":    DarkTheme,
	"light":   LightTheme,
	"monokai": MonokaiTheme,
	"dracula": DraculaTheme,
	"gruvbox": GruvboxTheme,
}

// GetTheme returns a theme by name, or the default dark theme if not found.
func GetTheme(name string) Theme {
	if theme, ok := BuiltinThemes[name]; ok {
		return theme
	}
	return DarkTheme
}

// ListThemes returns a list of available theme names.
func ListThemes() []string {
	names := make([]string, 0, len(BuiltinThemes))
	for name := range BuiltinThemes {
		names = append(names, name)
	}
	return names
}

// Gesh (𒄑) - A minimal TUI text editor written in Go with Bubble Tea.
// The name comes from Sumerian word meaning "pen, writing tool".
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/KilimcininKorOglu/gesh/internal/app"
	"github.com/KilimcininKorOglu/gesh/internal/config"
	"github.com/KilimcininKorOglu/gesh/pkg/version"
)

func main() {
	var filepath string
	var startLine, startCol int
	var readonly bool
	var themeName string
	var noConfig bool

	// Parse arguments
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		arg := args[i]

		switch {
		case arg == "-v" || arg == "--version":
			fmt.Println(version.Full())
			os.Exit(0)

		case arg == "-h" || arg == "--help":
			printHelp()
			os.Exit(0)

		case arg == "-r" || arg == "--readonly":
			readonly = true

		case arg == "-t" || arg == "--theme":
			// Get theme name from next argument
			if i+1 < len(args) {
				i++
				themeName = args[i]
			} else {
				fmt.Fprintln(os.Stderr, "Error: --theme requires a theme name")
				fmt.Fprintln(os.Stderr, "Available themes: dark, light, monokai, dracula, gruvbox")
				os.Exit(4)
			}

		case strings.HasPrefix(arg, "--theme="):
			themeName = strings.TrimPrefix(arg, "--theme=")

		case arg == "-n" || arg == "--norc":
			noConfig = true

		case strings.HasPrefix(arg, "+"):
			// Parse +N or +N:M
			pos := arg[1:]
			if strings.Contains(pos, ":") {
				parts := strings.SplitN(pos, ":", 2)
				if line, err := strconv.Atoi(parts[0]); err == nil {
					startLine = line
				}
				if col, err := strconv.Atoi(parts[1]); err == nil {
					startCol = col
				}
			} else {
				if line, err := strconv.Atoi(pos); err == nil {
					startLine = line
				}
			}

		case !strings.HasPrefix(arg, "-"):
			filepath = arg

		default:
			fmt.Fprintf(os.Stderr, "Unknown option: %s\n", arg)
			fmt.Fprintln(os.Stderr, "Use --help for usage information")
			os.Exit(4) // Exit code 4: Invalid argument
		}
	}

	// Load config unless --norc
	var cfg *config.Config
	if !noConfig {
		var err error
		cfg, err = config.Load()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to load config: %v\n", err)
			cfg = config.DefaultConfig()
		}
	} else {
		cfg = config.DefaultConfig()
	}

	// Apply theme: CLI flag takes precedence over config
	if themeName != "" {
		app.SetTheme(themeName)
	} else {
		app.SetTheme(cfg.Theme)
	}

	// Create the model
	var model *app.Model

	if filepath != "" {
		content, err := os.ReadFile(filepath)
		if err != nil {
			if os.IsNotExist(err) {
				// New file
				model = app.NewFromFile(filepath, filepath, "")
			} else if os.IsPermission(err) {
				fmt.Fprintf(os.Stderr, "Permission denied: %s\n", filepath)
				os.Exit(3) // Exit code 3: Permission error
			} else {
				fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
				os.Exit(2) // Exit code 2: File not found / read error
			}
		} else {
			model = app.NewFromFile(filepath, filepath, string(content))
		}
	} else {
		// New empty file
		model = app.New()
	}

	// Set readonly mode
	if readonly {
		model.SetReadonly(true)
	}

	// Go to specific line/column if specified
	if startLine > 0 {
		model.GotoLine(startLine, startCol)
	}

	// Create and run the program with mouse support
	p := tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseCellMotion())

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running program: %v\n", err)
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("Gesh (𒄑) - A minimal TUI text editor")
	fmt.Println()
	fmt.Println("Usage: gesh [options] [+line[:col]] [file]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -h, --help         Show this help message")
	fmt.Println("  -v, --version      Show version information")
	fmt.Println("  -r, --readonly     Open file in read-only mode")
	fmt.Println("  -t, --theme NAME   Set color theme (dark, light, monokai, dracula, gruvbox)")
	fmt.Println("  -n, --norc         Do not load config file")
	fmt.Println("  +N                 Open at line N")
	fmt.Println("  +N:M               Open at line N, column M")
	fmt.Println()
	fmt.Println("Keyboard shortcuts:")
	fmt.Println("  Ctrl+Alt+N  New file        Ctrl+O    Open file")
	fmt.Println("  Ctrl+S      Save            Ctrl+X    Exit/Cut")
	fmt.Println("  Ctrl+Z      Undo            Ctrl+Y    Redo")
	fmt.Println("  Ctrl+W      Search          Ctrl+R    Replace one")
	fmt.Println("  Ctrl+Shift+R Replace all    Ctrl+G    Go to line")
	fmt.Println("  Ctrl+K      Delete line     Ctrl+U    Cut line")
	fmt.Println("  Ctrl+V      Paste           Ctrl+C    Copy/Quit")
	fmt.Println()
	fmt.Println("Navigation (nano-style):")
	fmt.Println("  Arrow keys       Move cursor")
	fmt.Println("  Ctrl+P/N/B/F     Up/Down/Left/Right")
	fmt.Println("  Ctrl+Left/Right  Move by word")
	fmt.Println("  Home/Ctrl+A      Start of line (2x=select all)")
	fmt.Println("  End/Ctrl+E       End of line")
	fmt.Println("  Ctrl+Home/End    Start/end of file")
	fmt.Println("  PageUp/PageDown  Page navigation")
	fmt.Println()
	fmt.Println("Selection:")
	fmt.Println("  Ctrl+Space       Toggle selection")
	fmt.Println("  Shift+Arrows     Select text")
	fmt.Println()
	fmt.Println("For more information: https://github.com/KilimcininKorOglu/gesh")
}

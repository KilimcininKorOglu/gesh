// Gesh (𒄑) - A minimal TUI text editor written in Go with Bubble Tea.
// The name comes from Sumerian word meaning "pen, writing tool".
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/KilimcininKorOglu/gesh/internal/app"
	"github.com/KilimcininKorOglu/gesh/pkg/version"
)

func main() {
	// Handle version flag
	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		fmt.Println(version.Full())
		os.Exit(0)
	}

	// Handle help flag
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		printHelp()
		os.Exit(0)
	}

	// Create the model
	var model *app.Model

	if len(os.Args) > 1 {
		// Open file from argument
		filepath := os.Args[1]
		content, err := os.ReadFile(filepath)
		if err != nil {
			if os.IsNotExist(err) {
				// New file
				model = app.NewFromFile(filepath, filepath, "")
			} else {
				fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
				os.Exit(1)
			}
		} else {
			model = app.NewFromFile(filepath, filepath, string(content))
		}
	} else {
		// New empty file
		model = app.New()
	}

	// Create and run the program
	p := tea.NewProgram(model, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running program: %v\n", err)
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("Gesh - A minimal TUI text editor")
	fmt.Println()
	fmt.Println("Usage: gesh [options] [file]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -h, --help     Show this help message")
	fmt.Println("  -v, --version  Show version information")
	fmt.Println()
	fmt.Println("Keyboard shortcuts:")
	fmt.Println("  Ctrl+X    Exit")
	fmt.Println("  Ctrl+S    Save (not yet implemented)")
	fmt.Println("  Ctrl+W    Search (not yet implemented)")
	fmt.Println("  Ctrl+G    Go to line (not yet implemented)")
	fmt.Println()
	fmt.Println("Navigation:")
	fmt.Println("  Arrow keys     Move cursor")
	fmt.Println("  Home/End       Start/end of line")
	fmt.Println("  Ctrl+A/E       Start/end of line")
	fmt.Println()
	fmt.Println("For more information, visit: https://github.com/KilimcininKorOglu/gesh")
}

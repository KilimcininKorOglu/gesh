# Contributing to GESH

Thank you for your interest in contributing to Gesh! This document provides guidelines and information for contributors.

## Code of Conduct

Be respectful and constructive. We're all here to build a great text editor.

---

## Getting Started

### Prerequisites

- Go 1.21 or later
- Git
- A terminal with UTF-8 support

### Setup

```bash
# Fork and clone
git clone https://github.com/YOUR_USERNAME/gesh.git
cd gesh

# Add upstream remote
git remote add upstream https://github.com/KilimcininKorOglu/gesh.git

# Install dependencies
go mod download

# Build
go build -o gesh .

# Run tests
go test ./...
```

---

## Development Workflow

### 1. Create a Branch

```bash
git checkout -b feature/your-feature-name
# or
git checkout -b fix/bug-description
```

### 2. Make Changes

- Write code
- Add tests
- Update documentation if needed

### 3. Test

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/buffer/...

# Run benchmarks
go test -bench=. ./internal/buffer/...
```

### 4. Commit

Follow conventional commit messages:

```
feat: add word wrap support
fix: cursor position after undo
docs: update keybindings documentation
refactor: simplify gap buffer implementation
test: add selection tests
```

### 5. Push and Create PR

```bash
git push origin feature/your-feature-name
```

Then create a Pull Request on GitHub.

---

## Project Structure

```
gesh/
├── main.go                 # Entry point
├── internal/
│   ├── app/                # Bubble Tea model, update, view
│   ├── buffer/             # Gap buffer implementation
│   ├── config/             # Configuration parsing
│   ├── file/               # File I/O operations
│   ├── syntax/             # Syntax highlighting
│   │   └── languages/      # Language definitions
│   └── ui/
│       └── styles/         # Theme definitions
├── pkg/
│   └── version/            # Version info
└── configs/                # Example configs
```

---

## Coding Guidelines

### Go Style

- Follow standard Go formatting (`gofmt`)
- Use meaningful variable names
- Keep functions small and focused
- Add comments for exported functions

### Package Guidelines

| Package | Purpose |
|---------|---------|
| `internal/app` | Main application logic, Bubble Tea model |
| `internal/buffer` | Text buffer operations, undo/redo |
| `internal/file` | File reading/writing |
| `internal/syntax` | Syntax highlighting engine |
| `internal/config` | Configuration parsing |
| `internal/ui/styles` | UI themes |

### Testing

- Write unit tests for new functionality
- Aim for 80%+ coverage on new code
- Use table-driven tests where appropriate

```go
func TestSomething(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"empty", "", ""},
        {"simple", "hello", "HELLO"},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Something(tt.input)
            if result != tt.expected {
                t.Errorf("got %q, want %q", result, tt.expected)
            }
        })
    }
}
```

---

## Adding a New Language

To add syntax highlighting for a new language:

1. Create `internal/syntax/languages/yourlang.go`

```go
package languages

import (
    "regexp"
    "github.com/KilimcininKorOglu/gesh/internal/syntax"
)

func init() {
    syntax.RegisterLanguage(YourLang)
}

var YourLang = &syntax.Language{
    Name:       "YourLang",
    Extensions: []string{".yl", ".yourlang"},
    Rules: []syntax.Rule{
        // Comments first
        {Type: syntax.TokenComment, Pattern: regexp.MustCompile(`//.*$`)},
        
        // Strings
        {Type: syntax.TokenString, Pattern: regexp.MustCompile(`"[^"]*"`)},
        
        // Keywords
        {Type: syntax.TokenKeyword, Pattern: regexp.MustCompile(`\b(if|else|for)\b`)},
        
        // Numbers
        {Type: syntax.TokenNumber, Pattern: regexp.MustCompile(`\b[0-9]+\b`)},
    },
}
```

2. Build and test
3. Submit PR

---

## Pull Request Guidelines

### Before Submitting

- [ ] Code builds without errors
- [ ] All tests pass
- [ ] New code has tests
- [ ] Documentation updated if needed
- [ ] Commit messages follow convention

### PR Description

Include:
- What the change does
- Why it's needed
- How to test it
- Screenshots (for UI changes)

### Review Process

1. Automated checks run (build, tests)
2. Maintainer reviews code
3. Address feedback if any
4. PR merged when approved

---

## Reporting Issues

### Bug Reports

Include:
- Gesh version (`gesh --version`)
- OS and terminal
- Steps to reproduce
- Expected vs actual behavior
- Error messages or screenshots

### Feature Requests

Include:
- Use case description
- Proposed solution
- Alternatives considered

---

## Areas for Contribution

### Good First Issues

- Documentation improvements
- Adding new language syntax support
- Test coverage improvements
- Bug fixes

### Medium Difficulty

- Config option implementations
- UI improvements
- Performance optimizations

### Advanced

- Multi-buffer support
- Split view
- LSP integration
- Plugin system

---

## Communication

- **Issues:** Bug reports and feature requests
- **Pull Requests:** Code contributions
- **Discussions:** General questions and ideas

---

## License

By contributing, you agree that your contributions will be licensed under the same license as the project.

---

Thank you for contributing to Gesh! 🎉

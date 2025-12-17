// Package buffer provides text buffer implementations for the editor.
package buffer

const (
	// defaultGapSize is the initial size of the gap when creating a new buffer
	// or when the gap needs to be expanded.
	defaultGapSize = 64
)

// GapBuffer is a data structure optimized for text editing operations.
// It maintains a "gap" (empty space) at the cursor position, allowing
// O(1) insertions and deletions at that position.
//
// Structure example for "Hello World" with cursor after "Hello ":
//
//	data: ['H','e','l','l','o',' ', _, _, _, _, 'W','o','r','l','d']
//	                             ^           ^
//	                         gapStart     gapEnd
type GapBuffer struct {
	data     []rune // Character storage including the gap
	gapStart int    // Index where the gap begins (cursor position)
	gapEnd   int    // Index where the gap ends (first char after gap)
}

// New creates a new empty GapBuffer with default gap size.
func New() *GapBuffer {
	return &GapBuffer{
		data:     make([]rune, defaultGapSize),
		gapStart: 0,
		gapEnd:   defaultGapSize,
	}
}

// NewFromString creates a GapBuffer initialized with the given string.
// The cursor is positioned at the end of the text.
func NewFromString(s string) *GapBuffer {
	runes := []rune(s)
	textLen := len(runes)
	totalSize := textLen + defaultGapSize

	data := make([]rune, totalSize)
	copy(data, runes)

	return &GapBuffer{
		data:     data,
		gapStart: textLen,
		gapEnd:   totalSize,
	}
}

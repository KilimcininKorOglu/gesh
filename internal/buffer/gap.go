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

// gapSize returns the current size of the gap.
func (gb *GapBuffer) gapSize() int {
	return gb.gapEnd - gb.gapStart
}

// expandGap grows the gap when it becomes too small.
// It allocates a new larger array and copies the data.
func (gb *GapBuffer) expandGap(minSize int) {
	newGapSize := defaultGapSize
	if minSize > newGapSize {
		newGapSize = minSize
	}

	newData := make([]rune, len(gb.data)+newGapSize)

	// Copy text before the gap
	copy(newData, gb.data[:gb.gapStart])

	// Copy text after the gap to the new position
	newGapEnd := gb.gapStart + gb.gapSize() + newGapSize
	copy(newData[newGapEnd:], gb.data[gb.gapEnd:])

	gb.data = newData
	gb.gapEnd = gb.gapStart + gb.gapSize() + newGapSize
}

// Insert adds a single rune at the cursor position (gapStart).
// The cursor moves one position to the right after insertion.
func (gb *GapBuffer) Insert(r rune) {
	if gb.gapSize() == 0 {
		gb.expandGap(1)
	}

	gb.data[gb.gapStart] = r
	gb.gapStart++
}

// InsertString adds a string at the cursor position.
// The cursor moves to the end of the inserted string.
func (gb *GapBuffer) InsertString(s string) {
	runes := []rune(s)
	if len(runes) == 0 {
		return
	}

	if gb.gapSize() < len(runes) {
		gb.expandGap(len(runes))
	}

	copy(gb.data[gb.gapStart:], runes)
	gb.gapStart += len(runes)
}

// Delete removes the rune before the cursor (backspace behavior).
// Returns the deleted rune, or 0 if there's nothing to delete.
func (gb *GapBuffer) Delete() rune {
	if gb.gapStart == 0 {
		return 0
	}

	gb.gapStart--
	return gb.data[gb.gapStart]
}

// DeleteForward removes the rune after the cursor (delete key behavior).
// Returns the deleted rune, or 0 if there's nothing to delete.
func (gb *GapBuffer) DeleteForward() rune {
	if gb.gapEnd >= len(gb.data) {
		return 0
	}

	r := gb.data[gb.gapEnd]
	gb.gapEnd++
	return r
}

// Len returns the number of runes in the buffer (excluding the gap).
func (gb *GapBuffer) Len() int {
	return len(gb.data) - gb.gapSize()
}

// CursorPos returns the current cursor position (0-indexed).
func (gb *GapBuffer) CursorPos() int {
	return gb.gapStart
}

// MoveLeft moves the cursor one position to the left.
// Returns false if already at the beginning.
func (gb *GapBuffer) MoveLeft() bool {
	if gb.gapStart == 0 {
		return false
	}

	// Move one character from before the gap to after the gap
	gb.gapEnd--
	gb.gapStart--
	gb.data[gb.gapEnd] = gb.data[gb.gapStart]
	return true
}

// MoveRight moves the cursor one position to the right.
// Returns false if already at the end.
func (gb *GapBuffer) MoveRight() bool {
	if gb.gapEnd >= len(gb.data) {
		return false
	}

	// Move one character from after the gap to before the gap
	gb.data[gb.gapStart] = gb.data[gb.gapEnd]
	gb.gapStart++
	gb.gapEnd++
	return true
}

// MoveTo moves the cursor to the specified position.
// Position is clamped to valid range [0, Len()].
func (gb *GapBuffer) MoveTo(pos int) {
	// Clamp position to valid range
	if pos < 0 {
		pos = 0
	}
	textLen := gb.Len()
	if pos > textLen {
		pos = textLen
	}

	// Move cursor to target position
	for gb.gapStart > pos {
		gb.MoveLeft()
	}
	for gb.gapStart < pos {
		gb.MoveRight()
	}
}

// MoveToStart moves the cursor to the beginning of the buffer.
func (gb *GapBuffer) MoveToStart() {
	gb.MoveTo(0)
}

// MoveToEnd moves the cursor to the end of the buffer.
func (gb *GapBuffer) MoveToEnd() {
	gb.MoveTo(gb.Len())
}

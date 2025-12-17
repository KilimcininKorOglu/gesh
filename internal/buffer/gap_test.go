package buffer

import (
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	gb := New()

	if gb == nil {
		t.Fatal("New() returned nil")
	}

	if gb.gapStart != 0 {
		t.Errorf("gapStart = %d, want 0", gb.gapStart)
	}

	if gb.gapEnd != defaultGapSize {
		t.Errorf("gapEnd = %d, want %d", gb.gapEnd, defaultGapSize)
	}

	if len(gb.data) != defaultGapSize {
		t.Errorf("len(data) = %d, want %d", len(gb.data), defaultGapSize)
	}
}

func TestNewFromString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantLen  int
		wantGap  int // expected gapStart (cursor at end)
	}{
		{
			name:    "empty string",
			input:   "",
			wantLen: defaultGapSize,
			wantGap: 0,
		},
		{
			name:    "simple string",
			input:   "Hello",
			wantLen: 5 + defaultGapSize,
			wantGap: 5,
		},
		{
			name:    "unicode string",
			input:   "Merhaba dünya",
			wantLen: 13 + defaultGapSize, // 13 runes
			wantGap: 13,
		},
		{
			name:    "emoji string",
			input:   "Hello 🌍",
			wantLen: 7 + defaultGapSize, // 7 runes (emoji is 1 rune)
			wantGap: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gb := NewFromString(tt.input)

			if gb == nil {
				t.Fatal("NewFromString() returned nil")
			}

			if len(gb.data) != tt.wantLen {
				t.Errorf("len(data) = %d, want %d", len(gb.data), tt.wantLen)
			}

			if gb.gapStart != tt.wantGap {
				t.Errorf("gapStart = %d, want %d", gb.gapStart, tt.wantGap)
			}

			expectedGapEnd := tt.wantGap + defaultGapSize
			if gb.gapEnd != expectedGapEnd {
				t.Errorf("gapEnd = %d, want %d", gb.gapEnd, expectedGapEnd)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	gb := New()

	gb.Insert('H')
	gb.Insert('i')

	if gb.gapStart != 2 {
		t.Errorf("gapStart = %d, want 2", gb.gapStart)
	}

	// Verify the data before gap
	if gb.data[0] != 'H' || gb.data[1] != 'i' {
		t.Errorf("data = %c%c, want Hi", gb.data[0], gb.data[1])
	}
}

func TestInsertString(t *testing.T) {
	tests := []struct {
		name        string
		initial     string
		insert      string
		wantGapStart int
	}{
		{
			name:        "insert into empty",
			initial:     "",
			insert:      "Hello",
			wantGapStart: 5,
		},
		{
			name:        "insert empty string",
			initial:     "Hello",
			insert:      "",
			wantGapStart: 5,
		},
		{
			name:        "insert unicode",
			initial:     "",
			insert:      "Merhaba",
			wantGapStart: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gb *GapBuffer
			if tt.initial == "" {
				gb = New()
			} else {
				gb = NewFromString(tt.initial)
			}

			gb.InsertString(tt.insert)

			if gb.gapStart != tt.wantGapStart {
				t.Errorf("gapStart = %d, want %d", gb.gapStart, tt.wantGapStart)
			}
		})
	}
}

func TestInsertStringLarge(t *testing.T) {
	gb := New()

	// Insert a string larger than defaultGapSize to test gap expansion
	largeString := strings.Repeat("x", defaultGapSize+10)
	gb.InsertString(largeString)

	if gb.gapStart != len(largeString) {
		t.Errorf("gapStart = %d, want %d", gb.gapStart, len(largeString))
	}

	// Gap should have been expanded
	if gb.gapSize() < 0 {
		t.Errorf("gapSize = %d, should be >= 0", gb.gapSize())
	}
}

func TestDelete(t *testing.T) {
	gb := NewFromString("Hello")

	// Delete from end (cursor is at position 5)
	r := gb.Delete()
	if r != 'o' {
		t.Errorf("Delete() = %c, want 'o'", r)
	}

	if gb.gapStart != 4 {
		t.Errorf("gapStart = %d, want 4", gb.gapStart)
	}

	// Delete again
	r = gb.Delete()
	if r != 'l' {
		t.Errorf("Delete() = %c, want 'l'", r)
	}
}

func TestDeleteFromEmpty(t *testing.T) {
	gb := New()

	r := gb.Delete()
	if r != 0 {
		t.Errorf("Delete() from empty = %c, want 0", r)
	}
}

func TestDeleteForward(t *testing.T) {
	gb := NewFromString("Hello")
	// Move cursor to beginning by creating fresh buffer with cursor at start
	gb = New()
	gb.InsertString("Hello")
	// Now cursor is at end, we need to test DeleteForward
	// For this test, let's create a buffer where cursor is NOT at the end

	// Actually, let's create a proper test:
	// Create buffer with "Hello", cursor at end (position 5)
	// DeleteForward should return 0 (nothing after cursor)
	gb = NewFromString("Hello")
	r := gb.DeleteForward()
	if r != 0 {
		t.Errorf("DeleteForward() at end = %c, want 0", r)
	}
}

func TestDeleteForwardFromBeginning(t *testing.T) {
	// Create an empty buffer and insert with cursor movement
	// We'll need MoveTo to properly test this - for now test edge case
	gb := New()

	// Insert some text at position 0
	// Gap is at start, so text after gap doesn't exist yet
	r := gb.DeleteForward()
	if r != 0 {
		t.Errorf("DeleteForward() from empty = %c, want 0", r)
	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantLen int
	}{
		{"empty", "", 0},
		{"simple", "Hello", 5},
		{"unicode", "Merhaba", 7},
		{"emoji", "Hi 🌍", 4}, // 4 runes: H, i, space, emoji
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gb *GapBuffer
			if tt.input == "" {
				gb = New()
			} else {
				gb = NewFromString(tt.input)
			}

			if got := gb.Len(); got != tt.wantLen {
				t.Errorf("Len() = %d, want %d", got, tt.wantLen)
			}
		})
	}
}

func TestCursorPos(t *testing.T) {
	gb := New()
	if gb.CursorPos() != 0 {
		t.Errorf("CursorPos() on empty = %d, want 0", gb.CursorPos())
	}

	gb.InsertString("Hello")
	if gb.CursorPos() != 5 {
		t.Errorf("CursorPos() after insert = %d, want 5", gb.CursorPos())
	}
}

func TestMoveLeft(t *testing.T) {
	gb := NewFromString("Hello")

	// Cursor starts at end (position 5)
	if gb.CursorPos() != 5 {
		t.Errorf("Initial CursorPos() = %d, want 5", gb.CursorPos())
	}

	// Move left
	ok := gb.MoveLeft()
	if !ok {
		t.Error("MoveLeft() returned false, want true")
	}
	if gb.CursorPos() != 4 {
		t.Errorf("CursorPos() after MoveLeft = %d, want 4", gb.CursorPos())
	}

	// Move to beginning
	for gb.MoveLeft() {
	}
	if gb.CursorPos() != 0 {
		t.Errorf("CursorPos() at beginning = %d, want 0", gb.CursorPos())
	}

	// Try to move left at beginning
	ok = gb.MoveLeft()
	if ok {
		t.Error("MoveLeft() at beginning returned true, want false")
	}
}

func TestMoveRight(t *testing.T) {
	gb := NewFromString("Hello")

	// Move cursor to beginning first
	gb.MoveToStart()
	if gb.CursorPos() != 0 {
		t.Errorf("CursorPos() after MoveToStart = %d, want 0", gb.CursorPos())
	}

	// Move right
	ok := gb.MoveRight()
	if !ok {
		t.Error("MoveRight() returned false, want true")
	}
	if gb.CursorPos() != 1 {
		t.Errorf("CursorPos() after MoveRight = %d, want 1", gb.CursorPos())
	}

	// Move to end
	for gb.MoveRight() {
	}
	if gb.CursorPos() != 5 {
		t.Errorf("CursorPos() at end = %d, want 5", gb.CursorPos())
	}

	// Try to move right at end
	ok = gb.MoveRight()
	if ok {
		t.Error("MoveRight() at end returned true, want false")
	}
}

func TestMoveTo(t *testing.T) {
	gb := NewFromString("Hello World")

	tests := []struct {
		name     string
		pos      int
		wantPos  int
	}{
		{"move to middle", 5, 5},
		{"move to start", 0, 0},
		{"move to end", 11, 11},
		{"negative clamped to 0", -5, 0},
		{"beyond end clamped", 100, 11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gb.MoveTo(tt.pos)
			if gb.CursorPos() != tt.wantPos {
				t.Errorf("CursorPos() after MoveTo(%d) = %d, want %d", tt.pos, gb.CursorPos(), tt.wantPos)
			}
		})
	}
}

func TestMoveToStartEnd(t *testing.T) {
	gb := NewFromString("Hello")

	// Move to start
	gb.MoveToStart()
	if gb.CursorPos() != 0 {
		t.Errorf("CursorPos() after MoveToStart = %d, want 0", gb.CursorPos())
	}

	// Move to end
	gb.MoveToEnd()
	if gb.CursorPos() != 5 {
		t.Errorf("CursorPos() after MoveToEnd = %d, want 5", gb.CursorPos())
	}
}

func TestDeleteForwardWithMoveTo(t *testing.T) {
	gb := NewFromString("Hello")

	// Move cursor to beginning
	gb.MoveToStart()

	// Now DeleteForward should delete 'H'
	r := gb.DeleteForward()
	if r != 'H' {
		t.Errorf("DeleteForward() = %c, want 'H'", r)
	}

	if gb.Len() != 4 {
		t.Errorf("Len() after delete = %d, want 4", gb.Len())
	}
}

func TestInsertInMiddle(t *testing.T) {
	gb := NewFromString("Helo")

	// Move cursor to position 3 (after "Hel")
	gb.MoveTo(3)

	// Insert 'l'
	gb.Insert('l')

	// Move to end to read full string
	gb.MoveToEnd()

	// Verify length
	if gb.Len() != 5 {
		t.Errorf("Len() = %d, want 5", gb.Len())
	}
}

func TestRuneAt(t *testing.T) {
	gb := NewFromString("Hello")

	tests := []struct {
		pos  int
		want rune
	}{
		{0, 'H'},
		{1, 'e'},
		{4, 'o'},
		{-1, 0},  // out of bounds
		{5, 0},   // out of bounds
		{100, 0}, // out of bounds
	}

	for _, tt := range tests {
		got := gb.RuneAt(tt.pos)
		if got != tt.want {
			t.Errorf("RuneAt(%d) = %c, want %c", tt.pos, got, tt.want)
		}
	}
}

func TestRuneAtWithCursorInMiddle(t *testing.T) {
	gb := NewFromString("Hello")

	// Move cursor to middle (after "He")
	gb.MoveTo(2)

	// RuneAt should still work correctly regardless of cursor position
	tests := []struct {
		pos  int
		want rune
	}{
		{0, 'H'},
		{1, 'e'},
		{2, 'l'},
		{3, 'l'},
		{4, 'o'},
	}

	for _, tt := range tests {
		got := gb.RuneAt(tt.pos)
		if got != tt.want {
			t.Errorf("RuneAt(%d) with cursor at 2 = %c, want %c", tt.pos, got, tt.want)
		}
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"empty", ""},
		{"simple", "Hello"},
		{"unicode", "Merhaba dünya"},
		{"emoji", "Hello 🌍 World"},
		{"multiline", "Line 1\nLine 2\nLine 3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gb *GapBuffer
			if tt.input == "" {
				gb = New()
			} else {
				gb = NewFromString(tt.input)
			}

			got := gb.String()
			if got != tt.input {
				t.Errorf("String() = %q, want %q", got, tt.input)
			}
		})
	}
}

func TestStringAfterEdits(t *testing.T) {
	gb := NewFromString("Hello")

	// Delete last char
	gb.Delete()

	// Insert new chars
	gb.InsertString(" World")

	want := "Hell World"
	got := gb.String()
	if got != want {
		t.Errorf("String() after edits = %q, want %q", got, want)
	}
}

func TestStringWithCursorInMiddle(t *testing.T) {
	gb := NewFromString("Hello World")

	// Move cursor to middle
	gb.MoveTo(5)

	// String should return complete content
	want := "Hello World"
	got := gb.String()
	if got != want {
		t.Errorf("String() with cursor in middle = %q, want %q", got, want)
	}
}

func TestSlice(t *testing.T) {
	gb := NewFromString("Hello World")

	tests := []struct {
		name  string
		start int
		end   int
		want  string
	}{
		{"full", 0, 11, "Hello World"},
		{"first word", 0, 5, "Hello"},
		{"second word", 6, 11, "World"},
		{"middle", 2, 9, "llo Wor"},
		{"single char", 0, 1, "H"},
		{"empty range", 5, 5, ""},
		{"negative start clamped", -5, 5, "Hello"},
		{"end beyond length clamped", 6, 100, "World"},
		{"invalid range", 10, 5, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gb.Slice(tt.start, tt.end)
			if got != tt.want {
				t.Errorf("Slice(%d, %d) = %q, want %q", tt.start, tt.end, got, tt.want)
			}
		})
	}
}

func TestSliceWithCursorInMiddle(t *testing.T) {
	gb := NewFromString("Hello World")

	// Move cursor to middle
	gb.MoveTo(5)

	// Slice should work regardless of cursor position
	got := gb.Slice(0, 5)
	want := "Hello"
	if got != want {
		t.Errorf("Slice(0, 5) with cursor at 5 = %q, want %q", got, want)
	}

	got = gb.Slice(6, 11)
	want = "World"
	if got != want {
		t.Errorf("Slice(6, 11) with cursor at 5 = %q, want %q", got, want)
	}
}

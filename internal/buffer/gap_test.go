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

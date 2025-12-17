package buffer

import "testing"

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

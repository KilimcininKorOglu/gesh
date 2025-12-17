package buffer

import (
	"testing"
	"time"
)

func TestNewHistory(t *testing.T) {
	h := NewHistory()
	if h == nil {
		t.Fatal("NewHistory() returned nil")
	}
	if h.CanUndo() {
		t.Error("New history should not have undo")
	}
	if h.CanRedo() {
		t.Error("New history should not have redo")
	}
}

func TestPushAndUndo(t *testing.T) {
	h := NewHistory()

	// Push an insert operation
	h.Push(EditOperation{
		Type:     OpInsert,
		Position: 0,
		Text:     "Hello",
	})

	if !h.CanUndo() {
		t.Error("Should be able to undo after push")
	}

	op := h.Undo()
	if op == nil {
		t.Fatal("Undo() returned nil")
	}
	if op.Type != OpInsert {
		t.Errorf("op.Type = %d, want OpInsert", op.Type)
	}
	if op.Text != "Hello" {
		t.Errorf("op.Text = %q, want %q", op.Text, "Hello")
	}
}

func TestRedo(t *testing.T) {
	h := NewHistory()

	h.Push(EditOperation{
		Type:     OpInsert,
		Position: 0,
		Text:     "Hello",
	})

	// Undo
	h.Undo()

	if !h.CanRedo() {
		t.Error("Should be able to redo after undo")
	}

	// Redo
	op := h.Redo()
	if op == nil {
		t.Fatal("Redo() returned nil")
	}
	if op.Text != "Hello" {
		t.Errorf("op.Text = %q, want %q", op.Text, "Hello")
	}

	if !h.CanUndo() {
		t.Error("Should be able to undo after redo")
	}
}

func TestNewOperationClearsRedo(t *testing.T) {
	h := NewHistory()

	h.Push(EditOperation{Type: OpInsert, Position: 0, Text: "A"})
	h.Undo()

	if !h.CanRedo() {
		t.Error("Should be able to redo")
	}

	// New operation should clear redo
	h.Push(EditOperation{Type: OpInsert, Position: 0, Text: "B"})

	if h.CanRedo() {
		t.Error("Redo should be cleared after new operation")
	}
}

func TestMergeConsecutiveInserts(t *testing.T) {
	h := NewHistory()

	// Simulate fast typing
	h.Push(EditOperation{Type: OpInsert, Position: 0, Text: "H"})
	h.Push(EditOperation{Type: OpInsert, Position: 1, Text: "e"})
	h.Push(EditOperation{Type: OpInsert, Position: 2, Text: "l"})
	h.Push(EditOperation{Type: OpInsert, Position: 3, Text: "l"})
	h.Push(EditOperation{Type: OpInsert, Position: 4, Text: "o"})

	// Should be merged into one operation
	op := h.Undo()
	if op == nil {
		t.Fatal("Undo() returned nil")
	}
	if op.Text != "Hello" {
		t.Errorf("Merged text = %q, want %q", op.Text, "Hello")
	}

	// No more undos after single merged undo
	if h.CanUndo() {
		t.Error("Should not have more undo after undoing merged operation")
	}
}

func TestNoMergeAfterTimeout(t *testing.T) {
	h := NewHistory()
	h.mergeTimeout = 10 * time.Millisecond

	h.Push(EditOperation{Type: OpInsert, Position: 0, Text: "A"})
	time.Sleep(20 * time.Millisecond)
	h.Push(EditOperation{Type: OpInsert, Position: 1, Text: "B"})

	// Should be two separate operations
	op1 := h.Undo()
	if op1.Text != "B" {
		t.Errorf("First undo text = %q, want %q", op1.Text, "B")
	}

	op2 := h.Undo()
	if op2.Text != "A" {
		t.Errorf("Second undo text = %q, want %q", op2.Text, "A")
	}
}

func TestMergeBackspaceDeletes(t *testing.T) {
	h := NewHistory()

	// Simulate backspace: positions decrease
	h.Push(EditOperation{Type: OpDelete, Position: 4, Text: "o"})
	h.Push(EditOperation{Type: OpDelete, Position: 3, Text: "l"})
	h.Push(EditOperation{Type: OpDelete, Position: 2, Text: "l"})

	op := h.Undo()
	if op == nil {
		t.Fatal("Undo() returned nil")
	}
	if op.Text != "llo" {
		t.Errorf("Merged delete text = %q, want %q", op.Text, "llo")
	}
	if op.Position != 2 {
		t.Errorf("Merged delete position = %d, want 2", op.Position)
	}
}

func TestClear(t *testing.T) {
	h := NewHistory()

	h.Push(EditOperation{Type: OpInsert, Position: 0, Text: "A"})
	h.Undo()

	h.Clear()

	if h.CanUndo() {
		t.Error("CanUndo should be false after Clear")
	}
	if h.CanRedo() {
		t.Error("CanRedo should be false after Clear")
	}
}

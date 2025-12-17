package buffer

import "time"

// OpType represents the type of edit operation.
type OpType int

const (
	// OpInsert represents an insert operation.
	OpInsert OpType = iota
	// OpDelete represents a delete operation.
	OpDelete
)

// EditOperation represents a single edit operation for undo/redo.
type EditOperation struct {
	Type      OpType
	Position  int
	Text      string
	Timestamp time.Time
}

// History manages undo/redo stacks.
type History struct {
	undoStack    []EditOperation
	redoStack    []EditOperation
	maxSize      int
	mergeTimeout time.Duration
}

// NewHistory creates a new History with default settings.
func NewHistory() *History {
	return &History{
		undoStack:    make([]EditOperation, 0),
		redoStack:    make([]EditOperation, 0),
		maxSize:      1000,
		mergeTimeout: 500 * time.Millisecond,
	}
}

// Push adds an operation to the undo stack.
// Similar consecutive operations within mergeTimeout are merged.
func (h *History) Push(op EditOperation) {
	op.Timestamp = time.Now()

	// Clear redo stack on new operation
	h.redoStack = nil

	// Try to merge with previous operation
	if len(h.undoStack) > 0 {
		last := &h.undoStack[len(h.undoStack)-1]
		if h.canMerge(last, &op) {
			h.merge(last, &op)
			return
		}
	}

	// Add new operation
	h.undoStack = append(h.undoStack, op)

	// Limit stack size
	if len(h.undoStack) > h.maxSize {
		h.undoStack = h.undoStack[1:]
	}
}

// canMerge checks if two operations can be merged.
func (h *History) canMerge(last, new *EditOperation) bool {
	// Must be same type
	if last.Type != new.Type {
		return false
	}

	// Must be within timeout
	if new.Timestamp.Sub(last.Timestamp) > h.mergeTimeout {
		return false
	}

	// For inserts: must be consecutive
	if last.Type == OpInsert {
		return new.Position == last.Position+len([]rune(last.Text))
	}

	// For deletes: must be at same position (backspace) or consecutive (delete key)
	if last.Type == OpDelete {
		// Backspace: new position is one less
		if new.Position == last.Position-1 {
			return true
		}
		// Delete key: same position
		if new.Position == last.Position {
			return true
		}
	}

	return false
}

// merge combines two operations.
func (h *History) merge(last, new *EditOperation) {
	if last.Type == OpInsert {
		last.Text += new.Text
	} else if last.Type == OpDelete {
		if new.Position == last.Position-1 {
			// Backspace: prepend text
			last.Text = new.Text + last.Text
			last.Position = new.Position
		} else {
			// Delete key: append text
			last.Text += new.Text
		}
	}
	last.Timestamp = new.Timestamp
}

// Undo returns the operation to undo, or nil if nothing to undo.
func (h *History) Undo() *EditOperation {
	if len(h.undoStack) == 0 {
		return nil
	}

	// Pop from undo stack
	op := h.undoStack[len(h.undoStack)-1]
	h.undoStack = h.undoStack[:len(h.undoStack)-1]

	// Push to redo stack
	h.redoStack = append(h.redoStack, op)

	return &op
}

// Redo returns the operation to redo, or nil if nothing to redo.
func (h *History) Redo() *EditOperation {
	if len(h.redoStack) == 0 {
		return nil
	}

	// Pop from redo stack
	op := h.redoStack[len(h.redoStack)-1]
	h.redoStack = h.redoStack[:len(h.redoStack)-1]

	// Push to undo stack (without merging)
	h.undoStack = append(h.undoStack, op)

	return &op
}

// CanUndo returns true if there are operations to undo.
func (h *History) CanUndo() bool {
	return len(h.undoStack) > 0
}

// CanRedo returns true if there are operations to redo.
func (h *History) CanRedo() bool {
	return len(h.redoStack) > 0
}

// Clear clears all history.
func (h *History) Clear() {
	h.undoStack = nil
	h.redoStack = nil
}

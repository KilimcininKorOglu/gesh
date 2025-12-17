// Package app provides macro recording and playback functionality.
package app

import (
	tea "github.com/charmbracelet/bubbletea"
)

// MacroRecorder handles macro recording and playback.
type MacroRecorder struct {
	recording  bool
	keys       []tea.KeyMsg
	playing    bool
	playIndex  int
}

// NewMacroRecorder creates a new macro recorder.
func NewMacroRecorder() *MacroRecorder {
	return &MacroRecorder{
		keys: make([]tea.KeyMsg, 0),
	}
}

// IsRecording returns true if currently recording.
func (mr *MacroRecorder) IsRecording() bool {
	return mr.recording
}

// IsPlaying returns true if currently playing.
func (mr *MacroRecorder) IsPlaying() bool {
	return mr.playing
}

// StartRecording begins macro recording.
func (mr *MacroRecorder) StartRecording() {
	mr.recording = true
	mr.keys = make([]tea.KeyMsg, 0)
}

// StopRecording ends macro recording.
func (mr *MacroRecorder) StopRecording() {
	mr.recording = false
}

// ToggleRecording toggles macro recording.
func (mr *MacroRecorder) ToggleRecording() bool {
	if mr.recording {
		mr.StopRecording()
		return false
	}
	mr.StartRecording()
	return true
}

// RecordKey records a key press.
func (mr *MacroRecorder) RecordKey(key tea.KeyMsg) {
	if mr.recording && !mr.playing {
		// Don't record macro control keys
		keyStr := key.String()
		if keyStr == "ctrl+m" || keyStr == "ctrl+shift+m" {
			return
		}
		mr.keys = append(mr.keys, key)
	}
}

// Play starts macro playback.
func (mr *MacroRecorder) Play() bool {
	if len(mr.keys) == 0 || mr.recording {
		return false
	}
	mr.playing = true
	mr.playIndex = 0
	return true
}

// NextKey returns the next key in the macro, or nil if done.
func (mr *MacroRecorder) NextKey() *tea.KeyMsg {
	if !mr.playing || mr.playIndex >= len(mr.keys) {
		mr.playing = false
		return nil
	}
	key := mr.keys[mr.playIndex]
	mr.playIndex++
	return &key
}

// Clear clears the recorded macro.
func (mr *MacroRecorder) Clear() {
	mr.keys = make([]tea.KeyMsg, 0)
	mr.recording = false
	mr.playing = false
}

// KeyCount returns the number of recorded keys.
func (mr *MacroRecorder) KeyCount() int {
	return len(mr.keys)
}

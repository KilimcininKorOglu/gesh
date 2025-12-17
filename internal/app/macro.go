// Package app provides macro recording and playback functionality.
package app

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"

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

// SerializedKey represents a key for JSON serialization.
type SerializedKey struct {
	Type  string   `json:"type"`
	Runes []rune   `json:"runes,omitempty"`
	Alt   bool     `json:"alt,omitempty"`
	Paste bool     `json:"paste,omitempty"`
}

// MacroFile represents the macro file format.
type MacroFile struct {
	Version string                   `json:"version"`
	Macros  map[string][]SerializedKey `json:"macros"`
}

// getMacrosDir returns the macros directory path.
func getMacrosDir() string {
	var configDir string
	if runtime.GOOS == "windows" {
		configDir = os.Getenv("APPDATA")
		if configDir == "" {
			configDir = filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming")
		}
	} else {
		configDir = os.Getenv("XDG_CONFIG_HOME")
		if configDir == "" {
			configDir = filepath.Join(os.Getenv("HOME"), ".config")
		}
	}
	return filepath.Join(configDir, "gesh")
}

// getMacrosFilePath returns the macros file path.
func getMacrosFilePath() string {
	return filepath.Join(getMacrosDir(), "macros.json")
}

// serializeKey converts a tea.KeyMsg to SerializedKey.
func serializeKey(key tea.KeyMsg) SerializedKey {
	return SerializedKey{
		Type:  key.String(),
		Runes: key.Runes,
		Alt:   key.Alt,
		Paste: key.Paste,
	}
}

// deserializeKey converts a SerializedKey to tea.KeyMsg.
func deserializeKey(sk SerializedKey) tea.KeyMsg {
	return tea.KeyMsg{
		Type:  parseKeyType(sk.Type),
		Runes: sk.Runes,
		Alt:   sk.Alt,
		Paste: sk.Paste,
	}
}

// parseKeyType parses a key string to tea.KeyType.
func parseKeyType(s string) tea.KeyType {
	// Map common key strings to KeyType
	keyMap := map[string]tea.KeyType{
		"enter":      tea.KeyEnter,
		"tab":        tea.KeyTab,
		"backspace":  tea.KeyBackspace,
		"delete":     tea.KeyDelete,
		"up":         tea.KeyUp,
		"down":       tea.KeyDown,
		"left":       tea.KeyLeft,
		"right":      tea.KeyRight,
		"home":       tea.KeyHome,
		"end":        tea.KeyEnd,
		"pgup":       tea.KeyPgUp,
		"pgdown":     tea.KeyPgDown,
		"esc":        tea.KeyEsc,
		"space":      tea.KeySpace,
		"ctrl+a":     tea.KeyCtrlA,
		"ctrl+b":     tea.KeyCtrlB,
		"ctrl+c":     tea.KeyCtrlC,
		"ctrl+d":     tea.KeyCtrlD,
		"ctrl+e":     tea.KeyCtrlE,
		"ctrl+f":     tea.KeyCtrlF,
		"ctrl+g":     tea.KeyCtrlG,
		"ctrl+h":     tea.KeyCtrlH,
		"ctrl+i":     tea.KeyCtrlI,
		"ctrl+j":     tea.KeyCtrlJ,
		"ctrl+k":     tea.KeyCtrlK,
		"ctrl+l":     tea.KeyCtrlL,
		"ctrl+n":     tea.KeyCtrlN,
		"ctrl+o":     tea.KeyCtrlO,
		"ctrl+p":     tea.KeyCtrlP,
		"ctrl+q":     tea.KeyCtrlQ,
		"ctrl+r":     tea.KeyCtrlR,
		"ctrl+s":     tea.KeyCtrlS,
		"ctrl+t":     tea.KeyCtrlT,
		"ctrl+u":     tea.KeyCtrlU,
		"ctrl+v":     tea.KeyCtrlV,
		"ctrl+w":     tea.KeyCtrlW,
		"ctrl+x":     tea.KeyCtrlX,
		"ctrl+y":     tea.KeyCtrlY,
		"ctrl+z":     tea.KeyCtrlZ,
		"f1":         tea.KeyF1,
		"f2":         tea.KeyF2,
		"f3":         tea.KeyF3,
		"f4":         tea.KeyF4,
		"f5":         tea.KeyF5,
		"f6":         tea.KeyF6,
		"f7":         tea.KeyF7,
		"f8":         tea.KeyF8,
		"f9":         tea.KeyF9,
		"f10":        tea.KeyF10,
		"f11":        tea.KeyF11,
		"f12":        tea.KeyF12,
		"insert":     tea.KeyInsert,
	}

	if keyType, ok := keyMap[s]; ok {
		return keyType
	}

	// For regular character keys, return Runes type
	return tea.KeyRunes
}

// SaveMacro saves a named macro to the macros file.
func (mr *MacroRecorder) SaveMacro(name string) error {
	if len(mr.keys) == 0 {
		return nil
	}

	// Ensure directory exists
	macrosDir := getMacrosDir()
	if err := os.MkdirAll(macrosDir, 0755); err != nil {
		return err
	}

	// Load existing macros
	macroFile := &MacroFile{
		Version: "1.0",
		Macros:  make(map[string][]SerializedKey),
	}

	filePath := getMacrosFilePath()
	if data, err := os.ReadFile(filePath); err == nil {
		json.Unmarshal(data, macroFile)
	}

	// Convert current keys to serialized format
	serialized := make([]SerializedKey, len(mr.keys))
	for i, key := range mr.keys {
		serialized[i] = serializeKey(key)
	}

	// Save macro
	macroFile.Macros[name] = serialized

	// Write to file
	data, err := json.MarshalIndent(macroFile, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}

// LoadMacro loads a named macro from the macros file.
func (mr *MacroRecorder) LoadMacro(name string) error {
	filePath := getMacrosFilePath()
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var macroFile MacroFile
	if err := json.Unmarshal(data, &macroFile); err != nil {
		return err
	}

	serialized, ok := macroFile.Macros[name]
	if !ok {
		return os.ErrNotExist
	}

	// Convert to tea.KeyMsg
	mr.keys = make([]tea.KeyMsg, len(serialized))
	for i, sk := range serialized {
		mr.keys[i] = deserializeKey(sk)
	}

	return nil
}

// ListSavedMacros returns a list of saved macro names.
func ListSavedMacros() ([]string, error) {
	filePath := getMacrosFilePath()
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}

	var macroFile MacroFile
	if err := json.Unmarshal(data, &macroFile); err != nil {
		return nil, err
	}

	names := make([]string, 0, len(macroFile.Macros))
	for name := range macroFile.Macros {
		names = append(names, name)
	}
	return names, nil
}

// DeleteSavedMacro deletes a saved macro by name.
func DeleteSavedMacro(name string) error {
	filePath := getMacrosFilePath()
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var macroFile MacroFile
	if err := json.Unmarshal(data, &macroFile); err != nil {
		return err
	}

	delete(macroFile.Macros, name)

	newData, err := json.MarshalIndent(macroFile, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, newData, 0644)
}

// HasKeys returns true if the macro has recorded keys.
func (mr *MacroRecorder) HasKeys() bool {
	return len(mr.keys) > 0
}

// Package file provides file watching functionality.
package file

import (
	"os"
	"sync"
	"time"
)

// FileWatcher monitors a file for external changes.
type FileWatcher struct {
	path        string
	lastModTime time.Time
	lastSize    int64
	running     bool
	interval    time.Duration
	mu          sync.RWMutex
	onChange    func()
	stopCh      chan struct{}
}

// NewFileWatcher creates a new file watcher.
func NewFileWatcher(path string, interval time.Duration) *FileWatcher {
	return &FileWatcher{
		path:     path,
		interval: interval,
		stopCh:   make(chan struct{}),
	}
}

// SetPath updates the watched file path.
func (fw *FileWatcher) SetPath(path string) {
	fw.mu.Lock()
	defer fw.mu.Unlock()
	fw.path = path
	fw.updateStats()
}

// SetOnChange sets the callback for when file changes.
func (fw *FileWatcher) SetOnChange(callback func()) {
	fw.mu.Lock()
	defer fw.mu.Unlock()
	fw.onChange = callback
}

// Start begins watching the file.
func (fw *FileWatcher) Start() {
	fw.mu.Lock()
	if fw.running {
		fw.mu.Unlock()
		return
	}
	fw.running = true
	fw.updateStats()
	fw.mu.Unlock()

	go fw.watch()
}

// Stop stops watching the file.
func (fw *FileWatcher) Stop() {
	fw.mu.Lock()
	defer fw.mu.Unlock()
	if !fw.running {
		return
	}
	fw.running = false
	close(fw.stopCh)
	fw.stopCh = make(chan struct{})
}

// IsRunning returns true if the watcher is running.
func (fw *FileWatcher) IsRunning() bool {
	fw.mu.RLock()
	defer fw.mu.RUnlock()
	return fw.running
}

// Check manually checks for file changes.
// Returns true if the file has changed externally.
func (fw *FileWatcher) Check() bool {
	fw.mu.Lock()
	defer fw.mu.Unlock()

	if fw.path == "" {
		return false
	}

	info, err := os.Stat(fw.path)
	if err != nil {
		return false
	}

	modTime := info.ModTime()
	size := info.Size()

	// Check if file has changed
	if !modTime.Equal(fw.lastModTime) || size != fw.lastSize {
		// File changed externally
		return true
	}

	return false
}

// UpdateStats updates the stored file stats (call after saving).
func (fw *FileWatcher) UpdateStats() {
	fw.mu.Lock()
	defer fw.mu.Unlock()
	fw.updateStats()
}

// updateStats updates stats (must be called with lock held).
func (fw *FileWatcher) updateStats() {
	if fw.path == "" {
		return
	}

	info, err := os.Stat(fw.path)
	if err != nil {
		return
	}

	fw.lastModTime = info.ModTime()
	fw.lastSize = info.Size()
}

// watch is the background goroutine that checks for changes.
func (fw *FileWatcher) watch() {
	ticker := time.NewTicker(fw.interval)
	defer ticker.Stop()

	for {
		select {
		case <-fw.stopCh:
			return
		case <-ticker.C:
			if fw.Check() {
				fw.mu.RLock()
				callback := fw.onChange
				fw.mu.RUnlock()

				if callback != nil {
					callback()
				}

				// Update stats after notifying
				fw.UpdateStats()
			}
		}
	}
}

// GetLastModTime returns the last known modification time.
func (fw *FileWatcher) GetLastModTime() time.Time {
	fw.mu.RLock()
	defer fw.mu.RUnlock()
	return fw.lastModTime
}

// HasChanged checks if the file has changed since last check.
// This is a non-blocking check that doesn't update stats.
func (fw *FileWatcher) HasChanged() bool {
	fw.mu.RLock()
	path := fw.path
	lastMod := fw.lastModTime
	lastSize := fw.lastSize
	fw.mu.RUnlock()

	if path == "" {
		return false
	}

	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !info.ModTime().Equal(lastMod) || info.Size() != lastSize
}

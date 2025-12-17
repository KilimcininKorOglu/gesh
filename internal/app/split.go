// Package app provides split view functionality for the editor.
package app

// SplitDirection represents the direction of a split.
type SplitDirection int

const (
	SplitNone SplitDirection = iota
	SplitHorizontal // Side by side (left/right)
	SplitVertical   // Stacked (top/bottom)
)

// Pane represents a single editor pane in a split view.
type Pane struct {
	// Tab index this pane is showing
	tabIndex int

	// Viewport state for this pane
	viewportTopLine    int
	viewportLeftColumn int

	// Selection state (separate from tab's selection)
	cursorPos int

	// Dimensions (calculated during render)
	x, y          int
	width, height int
}

// SplitManager manages split views.
type SplitManager struct {
	// Split configuration
	direction SplitDirection

	// Panes (max 2 for now)
	panes []*Pane

	// Active pane index
	activePane int

	// Split ratio (0.0 to 1.0, default 0.5)
	splitRatio float64
}

// NewSplitManager creates a new split manager with no splits.
func NewSplitManager() *SplitManager {
	return &SplitManager{
		direction:  SplitNone,
		panes:      []*Pane{{tabIndex: 0}},
		activePane: 0,
		splitRatio: 0.5,
	}
}

// IsSplit returns true if the view is split.
func (sm *SplitManager) IsSplit() bool {
	return sm.direction != SplitNone
}

// Direction returns the current split direction.
func (sm *SplitManager) Direction() SplitDirection {
	return sm.direction
}

// ActivePaneIndex returns the index of the active pane.
func (sm *SplitManager) ActivePaneIndex() int {
	return sm.activePane
}

// ActivePane returns the currently active pane.
func (sm *SplitManager) ActivePane() *Pane {
	if sm.activePane >= 0 && sm.activePane < len(sm.panes) {
		return sm.panes[sm.activePane]
	}
	return nil
}

// PaneCount returns the number of panes.
func (sm *SplitManager) PaneCount() int {
	return len(sm.panes)
}

// Panes returns all panes.
func (sm *SplitManager) Panes() []*Pane {
	return sm.panes
}

// SplitHorizontal creates a horizontal split (left/right).
func (sm *SplitManager) SplitHorizontal(tabIndex int) bool {
	if sm.IsSplit() {
		return false // Already split
	}

	sm.direction = SplitHorizontal
	// Create second pane showing the same or different tab
	sm.panes = append(sm.panes, &Pane{tabIndex: tabIndex})
	return true
}

// SplitVertical creates a vertical split (top/bottom).
func (sm *SplitManager) SplitVertical(tabIndex int) bool {
	if sm.IsSplit() {
		return false // Already split
	}

	sm.direction = SplitVertical
	// Create second pane showing the same or different tab
	sm.panes = append(sm.panes, &Pane{tabIndex: tabIndex})
	return true
}

// CloseSplit closes the split and keeps the active pane.
func (sm *SplitManager) CloseSplit() bool {
	if !sm.IsSplit() {
		return false
	}

	// Keep the active pane's tab
	activeTabIndex := sm.panes[sm.activePane].tabIndex

	sm.direction = SplitNone
	sm.panes = []*Pane{{tabIndex: activeTabIndex}}
	sm.activePane = 0
	return true
}

// NextPane switches to the next pane.
func (sm *SplitManager) NextPane() {
	if len(sm.panes) > 1 {
		sm.activePane = (sm.activePane + 1) % len(sm.panes)
	}
}

// PrevPane switches to the previous pane.
func (sm *SplitManager) PrevPane() {
	if len(sm.panes) > 1 {
		sm.activePane = (sm.activePane - 1 + len(sm.panes)) % len(sm.panes)
	}
}

// SetActivePane sets the active pane by index.
func (sm *SplitManager) SetActivePane(index int) bool {
	if index >= 0 && index < len(sm.panes) {
		sm.activePane = index
		return true
	}
	return false
}

// SetPaneTab sets which tab a pane is showing.
func (sm *SplitManager) SetPaneTab(paneIndex, tabIndex int) bool {
	if paneIndex >= 0 && paneIndex < len(sm.panes) {
		sm.panes[paneIndex].tabIndex = tabIndex
		return true
	}
	return false
}

// GetActiveTabIndex returns the tab index of the active pane.
func (sm *SplitManager) GetActiveTabIndex() int {
	if pane := sm.ActivePane(); pane != nil {
		return pane.tabIndex
	}
	return 0
}

// SetSplitRatio sets the split ratio (0.0 to 1.0).
func (sm *SplitManager) SetSplitRatio(ratio float64) {
	if ratio < 0.1 {
		ratio = 0.1
	}
	if ratio > 0.9 {
		ratio = 0.9
	}
	sm.splitRatio = ratio
}

// GetSplitRatio returns the current split ratio.
func (sm *SplitManager) GetSplitRatio() float64 {
	return sm.splitRatio
}

// IncreaseSplitRatio increases the split ratio by 10%.
func (sm *SplitManager) IncreaseSplitRatio() {
	sm.SetSplitRatio(sm.splitRatio + 0.1)
}

// DecreaseSplitRatio decreases the split ratio by 10%.
func (sm *SplitManager) DecreaseSplitRatio() {
	sm.SetSplitRatio(sm.splitRatio - 0.1)
}

// CalculatePaneDimensions calculates dimensions for each pane.
func (sm *SplitManager) CalculatePaneDimensions(totalWidth, totalHeight int) {
	if !sm.IsSplit() || len(sm.panes) < 2 {
		// Single pane takes full area
		if len(sm.panes) > 0 {
			sm.panes[0].x = 0
			sm.panes[0].y = 0
			sm.panes[0].width = totalWidth
			sm.panes[0].height = totalHeight
		}
		return
	}

	switch sm.direction {
	case SplitHorizontal:
		// Side by side
		leftWidth := int(float64(totalWidth) * sm.splitRatio)
		rightWidth := totalWidth - leftWidth - 1 // -1 for separator

		sm.panes[0].x = 0
		sm.panes[0].y = 0
		sm.panes[0].width = leftWidth
		sm.panes[0].height = totalHeight

		sm.panes[1].x = leftWidth + 1 // +1 for separator
		sm.panes[1].y = 0
		sm.panes[1].width = rightWidth
		sm.panes[1].height = totalHeight

	case SplitVertical:
		// Stacked
		topHeight := int(float64(totalHeight) * sm.splitRatio)
		bottomHeight := totalHeight - topHeight - 1 // -1 for separator

		sm.panes[0].x = 0
		sm.panes[0].y = 0
		sm.panes[0].width = totalWidth
		sm.panes[0].height = topHeight

		sm.panes[1].x = 0
		sm.panes[1].y = topHeight + 1 // +1 for separator
		sm.panes[1].width = totalWidth
		sm.panes[1].height = bottomHeight
	}
}

// SavePaneState saves viewport state to a pane.
func (sm *SplitManager) SavePaneState(paneIndex, topLine, leftCol, cursorPos int) {
	if paneIndex >= 0 && paneIndex < len(sm.panes) {
		sm.panes[paneIndex].viewportTopLine = topLine
		sm.panes[paneIndex].viewportLeftColumn = leftCol
		sm.panes[paneIndex].cursorPos = cursorPos
	}
}

// RestorePaneState returns viewport state from a pane.
func (sm *SplitManager) RestorePaneState(paneIndex int) (topLine, leftCol, cursorPos int) {
	if paneIndex >= 0 && paneIndex < len(sm.panes) {
		pane := sm.panes[paneIndex]
		return pane.viewportTopLine, pane.viewportLeftColumn, pane.cursorPos
	}
	return 0, 0, 0
}

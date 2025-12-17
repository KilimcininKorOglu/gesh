package version

import (
	"strings"
	"testing"
)

func TestInfo(t *testing.T) {
	info := Info()
	if info == "" {
		t.Error("Info() returned empty string")
	}
}

func TestFull(t *testing.T) {
	full := Full()

	if !strings.HasPrefix(full, "gesh ") {
		t.Errorf("Full() should start with 'gesh ', got: %s", full)
	}

	if !strings.Contains(full, "commit:") {
		t.Errorf("Full() should contain 'commit:', got: %s", full)
	}

	if !strings.Contains(full, "built:") {
		t.Errorf("Full() should contain 'built:', got: %s", full)
	}
}

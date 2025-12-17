// Package version provides version information for the application.
package version

// Version information set at build time via ldflags.
// Example: go build -ldflags "-X github.com/KilimcininKorOglu/gesh/pkg/version.Version=1.0.0"
var (
	// Version is the semantic version of the application.
	Version = "0.1.0-dev"

	// Commit is the git commit hash.
	Commit = "unknown"

	// BuildDate is the date when the binary was built.
	BuildDate = "unknown"
)

// Info returns a formatted version string.
func Info() string {
	return Version
}

// Full returns complete version information.
func Full() string {
	return "gesh " + Version + " (commit: " + Commit + ", built: " + BuildDate + ")"
}

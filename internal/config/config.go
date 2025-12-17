// Package config provides configuration file support.
package config

import (
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

// Config represents the editor configuration.
type Config struct {
	// Editor settings
	Editor EditorConfig `yaml:"editor"`

	// Theme settings
	Theme string `yaml:"theme"`
}

// EditorConfig contains editor-specific settings.
type EditorConfig struct {
	TabSize            int  `yaml:"tab_size"`
	InsertSpaces       bool `yaml:"insert_spaces"`
	AutoIndent         bool `yaml:"auto_indent"`
	WordWrap           bool `yaml:"word_wrap"`
	LineNumbers        bool `yaml:"line_numbers"`
	ScrollPadding      int  `yaml:"scroll_padding"`
	TrimTrailingSpaces bool `yaml:"trim_trailing_spaces"`
	FinalNewline       bool `yaml:"final_newline"`
	CreateBackup       bool `yaml:"create_backup"`
}

// DefaultConfig returns the default configuration.
func DefaultConfig() *Config {
	return &Config{
		Editor: EditorConfig{
			TabSize:            4,
			InsertSpaces:       true,
			AutoIndent:         true,
			WordWrap:           false,
			LineNumbers:        true,
			ScrollPadding:      5,
			TrimTrailingSpaces: false,
			FinalNewline:       false,
			CreateBackup:       false,
		},
		Theme: "dark",
	}
}

// GetConfigDir returns the configuration directory path.
func GetConfigDir() string {
	switch runtime.GOOS {
	case "windows":
		appData := os.Getenv("APPDATA")
		if appData != "" {
			return filepath.Join(appData, "gesh")
		}
		return filepath.Join(os.Getenv("USERPROFILE"), ".config", "gesh")
	default:
		// Linux, macOS
		xdgConfig := os.Getenv("XDG_CONFIG_HOME")
		if xdgConfig != "" {
			return filepath.Join(xdgConfig, "gesh")
		}
		home, _ := os.UserHomeDir()
		return filepath.Join(home, ".config", "gesh")
	}
}

// GetConfigPath returns the full path to the config file.
func GetConfigPath() string {
	return filepath.Join(GetConfigDir(), "gesh.yaml")
}

// Load loads configuration from file.
func Load() (*Config, error) {
	configPath := GetConfigPath()

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return DefaultConfig(), nil
		}
		return nil, err
	}

	cfg := DefaultConfig()
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// Save saves configuration to file.
func Save(cfg *Config) error {
	configDir := GetConfigDir()
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(GetConfigPath(), data, 0644)
}

// CreateDefaultConfig creates a default config file if it doesn't exist.
func CreateDefaultConfig() error {
	configPath := GetConfigPath()

	if _, err := os.Stat(configPath); err == nil {
		// Config already exists
		return nil
	}

	return Save(DefaultConfig())
}

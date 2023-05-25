package config

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetConfigPath() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}

	configPath := filepath.Join(configDir, "seekr", "config.toml")
	return configPath, nil
}

func GetConfigDir() (string, error) {
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome != "" {
		return filepath.Join(xdgConfigHome, "seekr"), nil
	}

	homeDir, err := GetHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".config", "seekr"), nil
}

func GetHomeDir() (string, error) {
	if runtime.GOOS == "windows" {
		return os.UserHomeDir()
	} else {
		homeDir := os.Getenv("HOME")
		if homeDir == "" {
			return homeDir, ErrHome
		}
		return homeDir, nil
	}
}

package config

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestGetConfigPath(t *testing.T) {
	tempDir := t.TempDir()
	os.Setenv("XDG_CONFIG_HOME", tempDir)

	configPath, err := GetConfigPath()
	if err != nil {
		t.Errorf("GetConfigPath returned an error: %v", err)
	}
	expectedPath := filepath.Join(tempDir, "seekr", "config.toml")
	if configPath != expectedPath {
		t.Errorf("GetConfigPath returned %s, expected %s", configPath, expectedPath)
	}
}

func TestGetConfigDir(t *testing.T) {
	tempDir := t.TempDir()
	os.Setenv("XDG_CONFIG_HOME", tempDir)

	configDir, err := GetConfigDir()
	if err != nil {
		t.Errorf("GetConfigDir returned an error: %v", err)
	}
	expectedDir := filepath.Join(tempDir, "seekr")
	if configDir != expectedDir {
		t.Errorf("GetConfigDir returned %s, expected %s", configDir, expectedDir)
	}
}

func TestGetHomeDir(t *testing.T) {
	// Test GetHomeDir on Linux
	homeDir, err := GetHomeDir()
	if err != nil {
		t.Errorf("GetHomeDir returned an error: %v", err)
	}
	expectedDir := os.Getenv("HOME")
	if homeDir != expectedDir {
		t.Errorf("GetHomeDir returned %s, expected %s", homeDir, expectedDir)
	}
}

func TestGenDocs(t *testing.T) {
	writeDocs()
}

func writeDocs() {
	file, err := os.Create("doc.md")
	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(DefaultConfig().Markdown())
	if err != nil {
		fmt.Printf("Error when writing to file: %e\n", err)
		return
	}
}

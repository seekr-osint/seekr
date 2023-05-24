package config

import (
	"os"
	"path/filepath"

	"errors"
	"github.com/pelletier/go-toml"
	"runtime"
)

var (
	EmptyConfig = Config{}
	ErrHome     = errors.New("unable to determine home directory")
)

func LoadConfig() (*Config, error) {
	defaultConfig := &Config{
		Server: Server{
			Ip:   "localhost",
			Port: 8569,
		},
		General: General{
			ForcePort: false,
			Browser:   true,
		},
	}

	var homeDir string
	var err error

	if runtime.GOOS == "windows" {
		homeDir, err = os.UserHomeDir()
		if err != nil {
			return nil, err
		}
	} else {
		homeDir = os.Getenv("HOME")
		if homeDir == "" {
			return nil, ErrHome
		}
	}

	configPath := ""
	if runtime.GOOS == "windows" {
		configPath = filepath.Join(homeDir, "AppData", "Local", "seekr", "config.toml")
	} else {
		configPath = filepath.Join(homeDir, ".config", "seekr", "config.toml")
	}

	config, err := toml.LoadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return defaultConfig, nil
		}
		return nil, err
	}

	var cfg = *defaultConfig

	if err := config.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

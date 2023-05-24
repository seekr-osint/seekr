package config

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"
)

var (
	EmptyConfig = Config{}
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
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(homeDir, ".config", "seekr", "config.toml")
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

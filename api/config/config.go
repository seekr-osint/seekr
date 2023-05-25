package config

import (
	"os"

	"errors"
	"github.com/pelletier/go-toml"
)

var (
	EmptyConfig = Config{}
	ErrHome     = errors.New("unable to determine home directory")
)

func LoadConfig() (*Config, error) {
	cfg := DefaultConfig()

	configPath, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	config, err := toml.LoadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, err
	}

	if err := config.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

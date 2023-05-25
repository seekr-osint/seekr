package config

import (
	"github.com/mcuadros/go-defaults"
)

func DefaultConfig() *Config {
	config := new(Config)
	defaults.SetDefaults(config)

	return config
}

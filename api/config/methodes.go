package config

import (
	"fmt"
	"strings"

	"github.com/mcuadros/go-defaults"
	"github.com/pelletier/go-toml"
)

func DefaultConfig() Config {
	config := new(Config)
	defaults.SetDefaults(config)

	return *config
}

func (c Config) String() string {
	tomlString, err := toml.Marshal(c)
	if err != nil {
		return ""
	}
	return string(tomlString)
}
func (c Config) Markdown() string {
	var sb strings.Builder
	if c.String() == "" {
		return ""
	}

	sb.WriteString(fmt.Sprintf("```toml\n%s\n```", c.String()))

	return sb.String()
}

func (config Config) WriteConfig() error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}
	err = createFolderAndFile(configPath, config.String())
	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)
	} else {
		fmt.Printf("Config file sucessfully created at %s.\n", configPath)
	}
	return nil
}

func (c Config) Validate() error {
	return nil
}

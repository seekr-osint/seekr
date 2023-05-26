package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"errors"

	"github.com/pelletier/go-toml"
)

var (
	EmptyConfig     = Config{}
	ErrHome         = errors.New("unable to determine home directory")
	ErrNoConfigFile = errors.New("no config file")
)

func LoadConfig() (Config, error) {
	cfg := DefaultConfig()

	configPath, err := GetConfigPath()
	if err != nil {
		return cfg, err
	}

	config, err := toml.LoadFile(configPath)
	if err != nil {
		// This is bad, but os.IsNotExist() doesnt work here
		fmt.Printf("Error loading config file: %s\n", err)
		return cfg, ErrNoConfigFile
	}

	if err := config.Unmarshal(&cfg); err != nil {
		return DefaultConfig(), err
	}

	return cfg, nil
}

func CreateConfig() error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}
	err = createFolderAndFile(configPath, DefaultConfig().String())
	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)
	} else {
		fmt.Printf("Config file sucessfully created at %s.\n", configPath)
	}
	return nil
}

func createFolderAndFile(filePath string, text string) error {
	fmt.Printf("Creating file at %s.\n", filePath)

	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, []byte(text), 0644)
	if err != nil {
		return err
	}

	return nil
}

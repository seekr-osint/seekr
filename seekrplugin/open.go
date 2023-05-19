package seekrplugin

import (
	"errors"
	"fmt"
	"log"
	"os"
	"plugin"

	"github.com/seekr-osint/seekr/api"
)

var (
	ErrCantFindEntrySymbol = errors.New("can't find entry Symbol")
	ErrOpeningPlugin       = errors.New("error opening plugin")
)

func Open(paths []string, apiConfig api.ApiConfig) (api.ApiConfig, error) {
	var err error
	if apiConfig.Parsers == nil {
		apiConfig.Parsers = []func(api.ApiConfig) (api.ApiConfig, error){}
	}
	for _, path := range paths {
		apiConfig, err = load(apiConfig, path)
		if err != nil {
			return apiConfig, err
		}
	}
	return apiConfig, nil
}

func load(apiConfig api.ApiConfig, bundle string) (api.ApiConfig, error) {

	tempDir, err := unpackFile(bundle, "api/plugin.so")
	if err != nil {
		panic(err)
	}
	path := fmt.Sprintf("%s/api/plugin.so", tempDir)
	defer os.RemoveAll(tempDir)
	fmt.Printf("loading plugin: %s\n", path)
	loadedPlugin, err := plugin.Open(path)
	if err != nil {
		log.Printf("load plugin error: %s", err)
		return apiConfig, ErrOpeningPlugin
	}
	entry, err := loadedPlugin.Lookup("Main")
	if err != nil {
		return apiConfig, ErrCantFindEntrySymbol
	}
	err = entry.(func() error)()
	if err != nil {
		return apiConfig, err
	}
	preParser, err := loadedPlugin.Lookup("PreParser")
	if err == nil {
		parsedApiConfig, err := preParser.(func(api.ApiConfig) (api.ApiConfig, error))(apiConfig)
		if err != nil {
			return apiConfig, err
		}
		apiConfig, err = parsedApiConfig.Parse()
		if err != nil {
			return apiConfig, err
		}
	}
	postParser, err := loadedPlugin.Lookup("PostParser")
	if err == nil {
		parser := postParser.(func(api.ApiConfig) (api.ApiConfig, error))
		apiConfig.Parsers = append(apiConfig.Parsers, parser)
	}
	return apiConfig, nil
}

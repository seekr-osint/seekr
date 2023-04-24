package seekrplugin

import (
	"errors"
	"plugin"

	"github.com/seekr-osint/seekr/api"
)

var (
	ErrCantFindEntrySymbol = errors.New("can't find entry Symbol")
	ErrOpeningPlugin = errors.New("error opening plugin")
)

func Open(paths []string, apiConfig api.ApiConfig) (api.ApiConfig, error) {
	for _, path := range paths {

		loadedPlugin, err := plugin.Open(path)
		if err != nil {
			return apiConfig, ErrOpeningPlugin
		}
		entry, err := loadedPlugin.Lookup("Entry")
		if err != nil {
			return apiConfig,ErrCantFindEntrySymbol
		}
		err = entry.(func() error )()
		if err != nil {
			return apiConfig,err
		}
		configParser, err := loadedPlugin.Lookup("ConfigParser")
		if err == nil {
			parsedApiConfig,err := configParser.(func(api.ApiConfig) (api.ApiConfig,error) )(apiConfig)
			if err != nil {
				return apiConfig,err
			}
			apiConfig,err = parsedApiConfig.Parse()
			if err != nil {
				return apiConfig,err
			}
		}
	}
	return apiConfig, nil
}

package seekrdhandler

import (
	"errors"

	"github.com/seekr-osint/seekr/api"
	"github.com/seekr-osint/seekr/api/seekrd"
)

var (
	ErrConv = errors.New("apiConfig cannot be converted to api.ApiConfig")
)

func Handler(seekrdFunc func(*api.ApiConfig) error) seekrd.SeekrdFunc {
	//fmt.Printf("hello")
	return func(apiConfig seekrd.ApiConfig) (seekrd.ApiConfig, error) {
		var err error

		convertedApiConfig, ok := apiConfig.(*api.ApiConfig)
		if !ok {
			// Handle the case where the conversion is not possible
			return nil, ErrConv
		}
		err = seekrdFunc(convertedApiConfig)
		// parse the config
		apiConfig = convertedApiConfig

		return apiConfig, err
	}
}

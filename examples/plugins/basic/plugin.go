package main

import (
	"fmt"

	"github.com/seekr-osint/seekr/api"
)
func Entry() error {
	fmt.Printf("Entry")
	return nil
}

func ConfigParser(apiConfig api.ApiConfig) (api.ApiConfig,error) {
	fmt.Printf("running config parser")
	apiConfig.Server.Port = uint16(8080)
	return apiConfig,nil
}

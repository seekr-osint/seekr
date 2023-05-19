package main

import (
	"fmt"

	"net/http"

	"github.com/seekr-osint/seekr/api"

	"github.com/gin-gonic/gin"
)

func Main() error {
	return nil
}

func PreParser(apiConfig api.ApiConfig) (api.ApiConfig, error) {
	fmt.Printf("running config parser\n")
	apiConfig.Server.Port = uint16(8080)
	return apiConfig, nil
}

func PostParser(apiConfig api.ApiConfig) (api.ApiConfig, error) {
	fmt.Printf("running post parse config parser\nadded /plug api call\n")
	apiConfig.GinRouter.GET("/plug", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})
	return apiConfig, nil
}

package main

import (
	"fmt"

	"github.com/seekr-osint/seekr/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Entry() error {
	fmt.Printf("Entry")
	return nil
}

func ConfigParser(apiConfig api.ApiConfig) (api.ApiConfig, error) {
	fmt.Printf("running config parser\n")
	apiConfig.Server.Port = uint16(8080)
	return apiConfig, nil
}
func PostParseConfigParser(apiConfig api.ApiConfig) (api.ApiConfig, error) {
	fmt.Printf("running post parse config parser\n")
	apiConfig.GinRouter.GET("/plug", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})
	return apiConfig, nil
}
func main() {

}

package main

import (
	"bytes"
	"fmt"
	"io"

	"net/http"

	"github.com/seekr-osint/seekr/api"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/projectdiscovery/subfinder/v2/pkg/resolve"
	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
)

func Main() error {
	return nil
}

func PostParser(apiConfig api.ApiConfig) (api.ApiConfig, error) {
	fmt.Printf("running post parse config parser\n")
	apiConfig.GinRouter.GET("/deep/subfinder/:url", getSubdomains)
	return apiConfig, nil
}

func getSubdomains(c *gin.Context) {
	url := c.Param("url")

	runnerInstance, err := runner.NewRunner(&runner.Options{
		Threads:            10,
		Timeout:            30,
		MaxEnumerationTime: 10,
		Resolvers:          resolve.DefaultResolvers,
		ResultCallback: func(s *resolve.HostEntry) {
			log.Println(s.Host, s.Source)
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	buf := bytes.Buffer{}
	err = runnerInstance.EnumerateSingleDomain(url, []io.Writer{&buf})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := io.ReadAll(&buf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"subdomains": string(data),
	})
}

package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"net/url"
)

func (config ApiConfig) ServeTempMail() {
	router := gin.Default()

	// Add CORS middleware to allow all requests
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	// Reverse proxy to https://www.developermail.com/api/v1/
	apiURL, _ := url.Parse("https://www.developermail.com/api/v1/")
	proxy := httputil.NewSingleHostReverseProxy(apiURL)

	// Handler function for all API requests
	router.Any("/api/*path", func(c *gin.Context) {
		// Modify the request to preserve the original URL path
		c.Request.URL.Path = c.Param("path")

		// Forward the request to the remote API endpoint
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	// Run the server on port 8080
	router.Run(config.TempMailIp)
}

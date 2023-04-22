package api

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (config ApiConfig) ServeTempMail() {
	router := config.GinRouter

	// Add CORS middleware to allow all requests
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{
		"X-MailboxToken",
		"Content-Type",
		//"accept",
	}
	router.Use(cors.New(corsConfig))

	// Reverse proxy to https://www.developermail.com/api/v1/
	apiURL, _ := url.Parse("https://www.developermail.com/api/v1/")
	proxy := httputil.NewSingleHostReverseProxy(apiURL)

	// Handler function for all API requests
	router.Any("/developermail/api/*path", func(c *gin.Context) {
		// Modify the request to preserve the original URL path
		c.Request.URL.Path = c.Param("path")

		// Forward the request to the remote API endpoint
		proxy.ServeHTTP(c.Writer, c.Request)
	})
}

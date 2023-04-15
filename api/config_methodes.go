package api

import (
	"github.com/gin-gonic/gin"
	//"github.com/gin-contrib/static"
	"net/http"
)

// Parsing

func (config ApiConfig) Parse() (ApiConfig, error) {
	var err error
	config.DataBase, err = config.DataBase.Parse(config) // Parse the database
	return config, err
}

// Web Server
func (config ApiConfig) SetupWebServer() {
	config.GinRouter.GET("/web/*filepath", func(c *gin.Context) {
		http.FileServer(http.FS(config.WebServerFS)).ServeHTTP(c.Writer, c.Request)
	})
}

// DB methodes

func (config ApiConfig) SaveDB() error {
	return config.SaveDBFunc(config)
}

func (config ApiConfig) LoadDB() (ApiConfig, error) {
	return config.LoadDBFunc(config)
}

package api

import (
	"github.com/gin-gonic/gin"
	//"github.com/gin-contrib/static"
	"net/http"
)

// Parsing

func (config *ApiConfig) ParsePointer() error {
	parsedConfig, err := config.ConfigParse()
	if err != nil {
		return err
	}
	config = &parsedConfig

	config.DataBase, err = parsedConfig.DataBase.Parse(*config) // Parse the database
	if err != nil {
		return err
	}

	return nil
}

func (config ApiConfig) Parse() (ApiConfig, error) {
	config, err := config.ConfigParse()
	if err != nil {
		return config, err
	}
	config.DataBase, err = config.DataBase.Parse(config) // Parse the database
	if err != nil {
		return config, err
	}

	return config, nil
}

func (config ApiConfig) ConfigParse() (ApiConfig, error) {
	err := config.Validate()
	if err != nil {
		return config, err
	}
	config.Server, err = config.Server.Parse()
	if err != nil {
		return config, err
	}
	return config, nil
}

// Validation (Excluding the database)

func (config ApiConfig) Validate() error {
	err := config.Server.Validate() // Parse the database
	if err != nil {
		return err
	}
	return nil
}

// Web Server

func (config ApiConfig) SetupWebServer() {
	config.GinRouter.GET("/web/*filepath", func(c *gin.Context) {
		http.FileServer(http.FS(config.Server.WebServer.FileSystem)).ServeHTTP(c.Writer, c.Request)
	})
	config.GinRouter.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/web")
	})

	config.GinRouter.GET("/index.html", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/web/index.html")
	})
}

// DB methodes

func (config ApiConfig) SaveDB() error {
	return config.SaveDBFunc(config)
}

func (config ApiConfig) LoadDB() (ApiConfig, error) {
	return config.LoadDBFunc(config)
}

func (config *ApiConfig) LoadDBPointer() error {
	newConfig, err := config.LoadDBFunc(*config)
	if err != nil {
		return err
	}
	config = &newConfig
	return nil
}

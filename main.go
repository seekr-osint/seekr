package main

import (
	"embed"
	"fmt"

	api "github.com/niteletsplay/seekr/api"
	webServer "github.com/niteletsplay/seekr/webServer"
)

// Web server content
//
//go:embed web
var content embed.FS

var persons = make(api.DataBase)

var config = webServer.WebServerConfig{
	Type:    webServer.SingleBinary,
	Content: content,
	Ip:      ":5050",
}

func main() {
	go api.ServeApi(persons, ":8080", "data.json") // TODO config parsing stuff
	webServer.ParseConfig(config)
	fmt.Println(api.ServicesHandler(api.DefaultServices, "9glenda"))
}

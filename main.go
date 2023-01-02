package main

////import "C"
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

var people = make(api.DataBase)

func main() {

	go api.ServeApi(people, ":8080", "data.json") // TODO config parsing stuff
	RunWebServer()
	fmt.Println(api.ServicesHandler(api.DefaultServices, "9glenda"))
}

// //export RunWebServer
func RunWebServer() {

	var config = webServer.WebServerConfig{
		Type:    webServer.SingleBinary,
		Content: content,
		Ip:      ":5050",
	}
	webServer.ParseConfig(config)
}

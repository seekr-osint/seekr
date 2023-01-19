package main

////import "C"
import (
	"embed"
	"fmt"
	"strconv"

	api "github.com/seekr-osint/seekr/api"
	webServer "github.com/seekr-osint/seekr/webServer"
)

// Web server content
//
//go:embed web
var content embed.FS

var people = make(api.DataBase)

func main() {
	api.GithubInfoDeep("niteletsplay")
	fmt.Println("Welcome to seekr a powerful OSINT tool able to scan the web for " + strconv.Itoa(len(api.DefaultServices)) + "")
	go api.ServeApi(people, ":8080", "data.json") // TODO config parsing stuff
	RunWebServer()
	fmt.Println(api.ServicesHandler(api.DefaultServices, "9glenda"))
}

// //export RunWebServer
func RunWebServer() {

	var config = webServer.WebServerConfig{
		Type:    webServer.LiveServer,
		Content: content,
		Dir:     "./web",
		Ip:      ":5050",
	}
	webServer.ParseConfig(config)
}

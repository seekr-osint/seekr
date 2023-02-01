package main

import (
	"embed"
	"flag"
	"os/exec"
	"runtime"

	//"fmt"
	//"log"
	//"strconv"

	api "github.com/seekr-osint/seekr/api"
	webServer "github.com/seekr-osint/seekr/webServer"
)

// Web server content
//
//go:embed web/*
var content embed.FS

var people = make(api.DataBase)

func main() {

	liveServer := flag.Bool("live", false, "serve html files from seekr source code")
	dir := flag.String("dir", "./web", "dir where the html source code is located")
	ip := flag.String("ip", "localhost:5050", "Ip to serve the web server on")
	data := flag.String("dataJson", "data.json", "Database file")
	apiIp := flag.String("apiIp", "localhost:8080", "Ip to serve the api on")
	browser := flag.Bool("browser", true, "open up the html interface in the default web browser")

	flag.Parse()

	if *browser {
		openbrowser(*ip)
	}

	var apiConfig = api.ApiConfig{
		Ip:            *apiIp,
		LogFile:       "seekr.log",
		DataBaseFile:  *data,
		DataBase:      people,
		SetCORSHeader: true,
	}
	var config = webServer.WebServerConfig{
		Content: content,
		Dir:     *dir,
		Ip:      *ip,
	}
	if *liveServer {
		config.Type = webServer.LiveServer
	} else {
		config.Type = webServer.SingleBinary
	}

	//fmt.Println("Welcome to seekr a powerful OSINT tool able to scan the web for " + strconv.Itoa(len(api.DefaultServices)) + "services")
	go api.Seekrd(api.DefaultSeekrdServices, 30) // run every 30 minutes
	go api.ServeApi(apiConfig)
	RunWebServer(config)
}

// //export RunWebServer
func RunWebServer(config webServer.WebServerConfig) {
	webServer.ParseConfig(config)
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}
	api.Check(err)
}

package main

import (
	"embed"
	"flag"
	"fmt"
	"os/exec"
	"runtime"

	api "github.com/seekr-osint/seekr/api"
	webServer "github.com/seekr-osint/seekr/webServer"
)

// Web server content
//
//go:embed web/*
var content embed.FS

var dataBase = make(api.DataBase)

func main() {

	liveServer := flag.Bool("live", false, "serve html files from seekr source code")
	dir := flag.String("dir", "./web", "dir where the html source code is located")
	ip := flag.String("ip", "localhost", "Ip to serve api + webServer on (0.0.0.0 or localhost usually)")
	data := flag.String("db", "data", "Database location")
	apiPort := flag.String("apiPort", "8080", "Port to serve API on")
	webserverPort := flag.String("webserverPort", "5050", "Port to serve webserver on")
	browser := flag.Bool("browser", true, "open up the html interface in the default web browser")

	flag.Parse()

	if *browser {
		openbrowser(fmt.Sprintf("http://%s:%s/web/index.html", *ip, *apiPort))
	}

	var apiConfig = api.ApiConfig{
		WebServerFS:   content,
		Ip:            fmt.Sprintf("%s:%s", *ip, *apiPort),
		LogFile:       "seekr.log",
		DataBaseFile:  *data,
		DataBase:      dataBase,
		SetCORSHeader: true,
		SaveDBFunc:    api.DefaultSaveDB,
		LoadDBFunc:    api.DefaultLoadDB,
		WebServer:     true,
	}
	var config = webServer.WebServerConfig{
		Content: content,
		Dir:     *dir,
		Ip:      fmt.Sprintf("%s:%s", *ip, *webserverPort),
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
		err = exec.Command("cmd", "/c", "start", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}
	api.Check(err)
}

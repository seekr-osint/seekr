package main

import (
	"embed"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	api "github.com/seekr-osint/seekr/api"
	"github.com/seekr-osint/seekr/api/server"
	"github.com/seekr-osint/seekr/api/webserver"
)

// Web server content
//
//go:embed web/*
var content embed.FS

var dataBase = make(api.DataBase)

var version string

func checkVer() {
	destPath := os.Getenv("_SEEKR_UPDATE_BINARY")
	if destPath != "" {
		exePath, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exeFile, err := os.Open(exePath)
		if err != nil {
			panic(err)
		}
		defer exeFile.Close()

		destFile, err := os.Create(filepath.Join(destPath, filepath.Base(exePath)))
		if err != nil {
			panic(err)
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, exeFile)
		if err != nil {
			panic(err)
		}
	}
}
func main() {
	if version != "" {
		fmt.Printf("Welcome to seekr v%s\n", version)
		schematicVersion, err := ParseSchematicVersion(version)
		if err != nil {
			log.Panicf("error checking version: %s\n", version)
		}
		latestVersion, err := GetLatestSeekrVersion()
		if err != nil {
			log.Printf("error getting latest seekr version: %s\n", version)
		}
		if !schematicVersion.Latest(latestVersion) {
			downloadUrl := fmt.Sprintf("https://github.com/seekr-osint/seekr/releases/download/%s/%s", latestVersion, GetBinaryName(latestVersion))
			fmt.Printf("You are running an old seekr version.\nDownload the latest seekr version at: %s\n", downloadUrl)
			if promptYesNo("Update seekr") {
				err := updateBinary(downloadUrl)
				if err != nil {
					log.Panicf("error downloading seekr update: %s\n", err)
				}
				os.Exit(0)
			}
		}

	} else {
		fmt.Printf("Welcome to seekr unstable\nplease note that this version of seekr is NOT officially supported\n")
	}

	// dir := flag.String("dir", "./web", "dir where the html source code is located")
	ip := flag.String("ip", "localhost", "Ip to serve api + webServer on (0.0.0.0 or localhost usually)")
	data := flag.String("db", "data", "Database location")
	port := flag.Uint64("port", 8569, "Port to serve API on")
	enableWebserver := flag.Bool("webserver", true, "Enable the webserver")

	forcePort := flag.Bool("forcePort", false, "forcePort")
	//enableWebserver := flag.Bool("webserver", true, "Enable the webserver")
	enableApiServer := true
	// webserverPort := flag.String("webserverPort", "5050", "Port to serve webserver on")
	browser := flag.Bool("browser", true, "open up the html interface in the default web browser")

	flag.Parse()

	apiConfig, err := api.ApiConfig{
		Server: server.Server{
			Ip:        *ip,
			Port:      uint16(*port),
			ForcePort: *forcePort,
			WebServer: webserver.Webserver{
				Disable:    !*enableWebserver,
				FileSystem: content,
			},
			ApiServer: server.ApiServer{
				Disable: !enableApiServer,
			},
		},
		LogFile:       "seekr.log",
		DataBaseFile:  *data,
		DataBase:      dataBase,
		SetCORSHeader: true,
		SaveDBFunc:    api.DefaultSaveDB,
		LoadDBFunc:    api.DefaultLoadDB,
	}.ConfigParse()
	if err != nil {
		log.Panicf("error: %s", err)
	}
	if *browser && !apiConfig.Server.WebServer.Disable {
		openbrowser(fmt.Sprintf("http://%s:%d/web/index.html", apiConfig.Server.Ip, apiConfig.Server.Port))
	}
	//fmt.Println("Welcome to seekr a powerful OSINT tool able to scan the web for " + strconv.Itoa(len(api.DefaultServices)) + "services")
	go api.Seekrd(api.DefaultSeekrdServices, 30) // run every 30 minutes
	api.ServeApi(apiConfig)
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

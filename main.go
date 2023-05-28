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
	"strings"

	api "github.com/seekr-osint/seekr/api"
	"github.com/seekr-osint/seekr/api/config"

	"github.com/seekr-osint/seekr/api/discord"
	"github.com/seekr-osint/seekr/api/server"
	"github.com/seekr-osint/seekr/api/webserver"
	"github.com/seekr-osint/seekr/seekrplugin"
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

	cfg, err := config.LoadConfig()
	if err != nil && err != config.ErrNoConfigFile {
		fmt.Printf("Failed to load config: %s\n", err)
		return
	}
	configError := err
	// dir := flag.String("dir", "./web", "dir where the html source code is located")
	ip := flag.String("ip", cfg.Server.Ip, "Ip to serve api + webServer on (0.0.0.0 or localhost usually)")
	data := flag.String("db", "data", "Database location")
	port := flag.Uint64("port", cfg.Server.Port, "Port to serve the API on")
	enableWebserver := flag.Bool("webserver", true, "Enable the webserver")

	browser := flag.Bool("browser", cfg.General.Browser, "open up the html interface in the default web browser")
	forcePort := flag.Bool("forcePort", cfg.General.ForcePort, "forcePort")

	createConfig := flag.Bool("writeDefaultConfig", false, "create toml config file containing the default config if the config is invalid or doesn't exsist")

	enableRichCord := flag.Bool("discord", cfg.General.Discord, "Enable the discord rich appearance")
	//enableWebserver := flag.Bool("webserver", true, "Enable the webserver")
	enableApiServer := true
	// webserverPort := flag.String("webserverPort", "5050", "Port to serve webserver on")
	pluginList := os.Getenv("SEEKR_PLUGINS")
	plugins := []string{}
	if pluginList != "" {
		plugins = strings.Split(pluginList, ",")
	}
	flag.Parse()
	if configError == config.ErrNoConfigFile && *createConfig {
		err = config.CreateDefaultConfig()
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}
	}
	if *enableRichCord {
		err := discord.Rich()
		if err == nil {
			// No error printing due it printing an error if discord is not running / installed
			//fmt.Printf("%s\n", err)
			fmt.Printf("Setting discord rich presence\n")
		}
	}

	apiConfig, err := api.ApiConfig{
		Config:  cfg,
		Version: version,
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
	apiConfig, err = seekrplugin.Open(plugins, apiConfig)
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

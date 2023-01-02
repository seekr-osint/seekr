package web

import (
	"embed"
	"log"
	"net/http"
)

type ServerType int

const (
	SingleBinary ServerType = iota
	LiveServer   ServerType = iota
	NotEnabled   ServerType = iota
)

type WebServerConfig struct {
	Type    ServerType
	Content embed.FS
	Dir string
	Ip      string
}

func serveSingleBinary(config WebServerConfig) {
	// Serve files from static folder
	http.Handle("/", http.FileServer(http.FS(config.Content)))

	println("web server running" + config.Ip)
	log.Fatal(http.ListenAndServe(config.Ip, nil))
}

func serveLive(config WebServerConfig) {
	// Serve files from static folder
	http.Handle("/", http.FileServer(http.Dir(config.Dir)))

	println("web server running" + config.Ip)
	log.Fatal(http.ListenAndServe(config.Ip, nil))
}


func ParseConfig(config WebServerConfig) {
	switch config.Type {
	case SingleBinary:
		serveSingleBinary(config)
	case LiveServer:
   serveLive(config) 
	case NotEnabled:
	default:
		panic("WebServerConfig.Type is neither SingleBinary, LiveServer or NotEnabled")
	}
}

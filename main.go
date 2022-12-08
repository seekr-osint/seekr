package main

import (
	"embed"
	api "github.com/niteletsplay/seekr/api"
	"log"
	"net/http"
)

// web holds our web server content.
//
//go:embed web
var content embed.FS
var persons = make(api.DataBase)

func main() {
	go api.ServeApi(persons, ":8080", "data.json")
	// Serve files from static folder
	http.Handle("/", http.FileServer(http.FS(content)))

	println("web server running http://localhost:5050/web")
	log.Fatal(http.ListenAndServe(":5050", nil))

}

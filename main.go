package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	api "github.com/niteletsplay/seekr/api"
)

// web holds our web server content.
//
//go:embed web
var content embed.FS
var persons = make(api.DataBase)
type Config struct {
  apiServer bool
  webServer bool
}

func main() {
	go api.ServeApi(persons, ":8080", "data.json")
  fmt.Println(api.CheckUsername("9glenda"))

  fmt.Println(api.CheckUsername("9glenda22"))
	// Serve files from static folder
	http.Handle("/", http.FileServer(http.FS(content)))

	println("web server running http://localhost:5050/web")
	log.Fatal(http.ListenAndServe(":5050", nil))

}

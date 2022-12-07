package main

import (
	"net/http"
  "embed"
  "log"
)

// web holds our web server content.
//go:embed web
var content embed.FS

func main() {


	// Serve files from static folder
  http.Handle("/", http.FileServer(http.FS(content)))
  //http.Handle("/", http.StripPrefix("/", http.FileServer(http.FS(content))))

  println("web server running http://localhost:5050/web")
	// Start server on port specified above
  log.Fatal(http.ListenAndServe(":5050", nil))

}


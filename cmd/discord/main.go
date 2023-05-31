package main

import (
	"log"
	"time"

	"github.com/seekr-osint/seekr/api/discord"
)

func main() {
	for {
		err := discord.Rich()
		if err == nil {
			// No error printing due it printing an error if discord is not running / installed
			//fmt.Printf("%s\n", err)
			log.Printf("Setting discord rich presence\n")
		}
		time.Sleep(10 * time.Second)
	}
}

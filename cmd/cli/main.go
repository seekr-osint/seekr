package main

import (
	"fmt"

	"github.com/seekr-osint/seekr/api/client"
)

func main() {
	c := client.NewClient("localhost", 8569)
	ping, err := c.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("ping: %s", ping)
}

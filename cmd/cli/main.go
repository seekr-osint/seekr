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
	fmt.Printf("ping: %s\n", ping)

	db, err := c.GetDB() 
	if err != nil {
		panic(err)
	}
	fmt.Printf("db: %s\n",db)

	person, err := c.GetPerson("1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("person: %s\n",person.Markdown())
}

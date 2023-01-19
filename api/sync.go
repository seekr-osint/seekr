package api

import (
	"log"
	"sync"
)

func Emails(email string) {
	wg := &sync.WaitGroup{}

	var email_services = map[string]bool{}
	wg.Add(1)
	go func() {
		// Do something
		log.Println("hello")
		email_services["discord"] = Discord(email)
		wg.Done()
	}()
	log.Println("hello2")

	wg.Wait()
	log.Println(email_services["discord"])
}

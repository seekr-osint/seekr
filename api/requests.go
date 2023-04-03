package api

import (
	"log"
	"net/http"
)

func GetStatusCode(url string) int { // FIXME config
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
	}
	return resp.StatusCode
}
func GetStatusCodeNew(url string,config ApiConfig) (error,int) { // FIXME config
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
    return err,0
	}
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
    return err,0
	}
	return nil,resp.StatusCode
}

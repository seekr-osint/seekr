package api

import (
	"log"
	"math/rand"
	"time"
)

func RandomChar() string {
	rand.Seed(time.Now().UnixNano())
	charset := "abcdefghijklmnopqrstuvwxyz1234567890"
	return string(charset[rand.Intn(len(charset))])
}
func RandomString(cnt int) string {
	str := ""
	for i := 0; i < cnt; i++ {
		str = str + RandomChar()
	}
	return str
}
func Check(err error) {
	if err != nil {
		log.Println(err)
	}
}

package api

import (
	"log"
	"math/rand"
	"os/exec"
	"runtime"
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
func Openbrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
	}
	Check(err)

}

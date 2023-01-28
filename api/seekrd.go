package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

type SeekrdServices []SeekrdService
type SeekrdService struct {
	Name     string     // example: "GitHub"
	Function SeekrdFunc // example: foo(DataBase) DataBase  func(db DataBase) DataBase { return db }
}
type SeekrdFunc func(DataBase) DataBase

var DefaultSeekrdServices = SeekrdServices{
	SeekrdService{
		Name:     "test",
		Function: func(db DataBase) DataBase { log.Println("test seekrd"); return db },
	},
}

func Seekrd(seekrdServices SeekrdServices, interval int) { // Seekrd(DefaultSeekrdServices,30) 30 is in minutes
	ticker := time.NewTicker(time.Duration(interval) * time.Minute)
	for range ticker.C {
		log.Println("Seekrd...")
		var db DataBase
		file, _ := ioutil.ReadFile("data.json")
		err := json.Unmarshal(file, &db)
		Check(err)
		for _, seekrdService := range seekrdServices {
			db = seekrdService.Function(db)
		}
		SaveJson(db)
		// FIXME lock file
	}
}

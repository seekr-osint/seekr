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
		Name: "birthday",
		//Function: func(db DataBase) DataBase { log.Println("test seekrd"); return db },
		Function: Birthday,
	},
}

func IsValidDate(dateString string) (bool, time.Time) {
	dateFormats := []string{
		//"2006-01-02",
		//"02-01-2006",
		//"01/02/2006",
		"01.02.2006",
	}
	var err error
	var date time.Time
	for _, format := range dateFormats {
		date, err = time.Parse(format, dateString)
		if err != nil {
			return false, time.Now()
		}
	}
	return true, date
}

func Birthday(db DataBase) DataBase {
	for i, person := range db {
		if person.Birthday != "" {
			value, date := IsValidDate(person.Birthday)
			if value {
				now := time.Now()
				age := now.Year() - date.Year()
				if now.YearDay() < date.YearDay() {
					age--
				}
				log.Println(person.Birthday)
				log.Println(age)
				person.Age = Age(age)
			}
		}
		db[i] = person
	}
	return db
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

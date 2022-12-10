package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func SaveJson(people DataBase) {
	jsonBytes, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error:", err)
	}
	ioutil.WriteFile(DatabaseFile, jsonBytes, 0644)
}

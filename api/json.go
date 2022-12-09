package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func SaveJson(persons DataBase) {
	jsonBytes, err := json.Marshal(persons)
	if err != nil {
		fmt.Println("error:", err)
	}
	ioutil.WriteFile(DatabaseFile, jsonBytes, 0644)
}

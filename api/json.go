package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func SaveJson(people DataBase) {
	jsonBytes, err := json.MarshalIndent(people, "", "\t")

	if err != nil {
		fmt.Println("error:", err)
	}
	ioutil.WriteFile(DatabaseFile, jsonBytes, 0644)
}

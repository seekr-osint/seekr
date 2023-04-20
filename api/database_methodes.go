package api

import (
	"errors"

	"github.com/seekr-osint/seekr/api/functions"
)

// Parsing

func (dataBase DataBase) Parse(config ApiConfig) (DataBase, error) {
	var err error
	for _, name := range functions.SortMapKeys(map[string]Person(dataBase)) {
		dataBase[name], err = dataBase[name].Parse(config)
	}
	return dataBase, err
}

// GetPerson

func (config ApiConfig) GetPerson(id string) (Person, error) {
	if _, ok := config.DataBase[id]; ok {
		return config.DataBase[id], nil
	}
	return Person{}, errors.New("person does not exsist")
}

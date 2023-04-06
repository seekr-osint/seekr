package api

import (
	"errors"
)

// Errors:

func (config ApiConfig) GetPerson(id string) (Person, error) {
	if _, ok := config.DataBase[id]; ok {
		return config.DataBase[id], nil
	}
	return Person{}, errors.New("person does not exsist")
}

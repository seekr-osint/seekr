package api

import (
	"fmt"
	"strings"
)

func (config ApiConfig) SaveDB() error {
	return config.SaveDBFunc(config)
}

func (config ApiConfig) LoadDB() (ApiConfig, error) {
	return config.LoadDBFunc(config)
}

func (config ApiConfig) Parse() (ApiConfig, error) {
	var err error
	config.DataBase, err = config.DataBase.Parse(config)
	return config, err
}

func (dataBase DataBase) Parse(config ApiConfig) (DataBase, error) {
	var err error
	for _, name := range SortMapKeys(map[string]Person(dataBase)) {
		dataBase[name], err = dataBase[name].Parse(config)
	}
	return dataBase, err
}

func (person Person) Markdown() string {
	var sb strings.Builder
	if person.Name != "" {
		sb.WriteString(fmt.Sprintf("# %s\n", person.Name))
	} else {
		sb.WriteString(fmt.Sprintf("# ID:%s\n", person.ID))
	}

	sb.WriteString(person.Gender.Markdown())
	sb.WriteString(person.Age.Markdown())
	sb.WriteString(person.Civilstatus.Markdown())
	sb.WriteString(person.Religion.Markdown())
	sb.WriteString(person.Phone.Markdown())
	if len(person.Email) >= 1 {
		sb.WriteString(fmt.Sprintf("## Email\n%s\n", person.Email.Markdown()))
	}
	return sb.String()
}

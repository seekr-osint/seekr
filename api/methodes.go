package api

import (
	"fmt"
	"strings"

	//"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type prefixedFileSystem struct {
	fs     http.FileSystem
	prefix string
}

func (p *prefixedFileSystem) Open(name string) (http.File, error) {
	if !strings.HasPrefix(name, p.prefix) {
		return nil, os.ErrNotExist
	}
	return p.fs.Open(name[len(p.prefix):])
}

func (config ApiConfig) RemovePrefix(prefix string) http.FileSystem {
	fs := config.WebServerFS
	return &prefixedFileSystem{fs, prefix}
}

func (config ApiConfig) SetupWebServer() {
	config.GinRouter.StaticFS("/web", config.RemovePrefix("/web"))
}

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

//go:generate go run generate.go

package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/seekr-osint/seekr/api"
	"github.com/seekr-osint/seekr/api/config"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

func main() {
	GenType(api.Person{})
	GenType(config.Config{})

}

func GenType(toConvert interface{}) error {

	converter := typescriptify.New()
	converter.CreateConstructor = true
	converter.Indent = "    "
	converter.BackupDir = ""

	converter.Add(toConvert)

	fileName := fmt.Sprintf("../web/ts-gen/%s.ts", strings.ToLower(reflect.TypeOf(toConvert).Name()))
	err := os.MkdirAll("../web/ts-gen", os.ModePerm)
	if err != nil {
		return err
	}
	err = converter.ConvertToFile(fileName)
	if err != nil {
		return err
	}
	return nil
}

//go:generate go run generate.go

package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"sync"

	//"github.com/seekr-osint/seekr/api"
	"github.com/seekr-osint/seekr/api"
	"github.com/seekr-osint/seekr/api/config"
	"github.com/seekr-osint/seekr/api/services"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

func main() {
	wg := &sync.WaitGroup{}
	GenType(api.Person{}, wg)
	GenType(config.Config{}, wg)
	GenType(services.ServiceCheckResult{}, wg)
	wg.Wait()
}

func GenType(toConvert interface{}, wg *sync.WaitGroup) error {
	wg.Add(1)
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
	wg.Done()
	return nil
}

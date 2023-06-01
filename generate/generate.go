//go:generate go run generate.go

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/seekr-osint/seekr/api/config"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

func main() {
	converter := typescriptify.New()
	//converter.CreateConstructor = true
	converter.Indent = "    "
	converter.BackupDir = ""

	converter.Add(config.Config{})
	//converter.CreateInterface = true

	fileName := "../web/ts-gen/config.ts"
	err := os.MkdirAll(filepath.Dir(fileName), os.ModePerm)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}
	err = converter.ConvertToFile(fileName)
	if err != nil {
		panic(err.Error())
	}

}

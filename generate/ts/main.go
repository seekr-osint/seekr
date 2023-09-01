//go:generate go run main.go
package main

import (
	"fmt"

	"github.com/gzuidhof/tygo/tygo"
)

func main() {
	fmt.Println("generating typescript")
	config := &tygo.Config{
		Packages: []*tygo.PackageConfig{
			{
				Path:       "github.com/seekr-osint/seekr/api/person",
				OutputPath: "./../../web/ts-gen/person.ts",
			},
		},
	}
	gen := tygo.New(config)
	err := gen.Generate()
	if err != nil {
		panic(err)
	}
}

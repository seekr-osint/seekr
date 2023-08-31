//go:generate go run main.go
package main

import (
	"github.com/gzuidhof/tygo/tygo"
)

func main() {
	config := &tygo.Config{
		Packages: []*tygo.PackageConfig{
			{
				Path:       "github.com/seekr-osint/seekrstack/api/person",
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

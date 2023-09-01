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
				Frontmatter: `import * as enums from "enums.js"
import * as services from "services.js"

import * as hobbies from "hobbies.js"
import * as ips from "ips.js"
import * as clubs from "clubs.js"
import * as sources from "sources.js"

`,
			},
			{
				Path:       "github.com/seekr-osint/seekr/api/enums",
				OutputPath: "./../../web/ts-gen/enums.ts",
			},
			{
				Path:         "github.com/seekr-osint/seekr/api/services",
				OutputPath:   "./../../web/ts-gen/services.ts",
				ExcludeFiles: []string{"functypes.go"},
			},

			{
				Path:       "github.com/seekr-osint/seekr/api/types/clubs",
				OutputPath: "./../../web/ts-gen/clubs.ts",
			},
			{
				Path:       "github.com/seekr-osint/seekr/api/types/sources",
				OutputPath: "./../../web/ts-gen/sources.ts",
			},
			{
				Path:       "github.com/seekr-osint/seekr/api/types/hobbies",
				OutputPath: "./../../web/ts-gen/hobbies.ts",
			},
			{
				Path:       "github.com/seekr-osint/seekr/api/types/ips",
				OutputPath: "./../../web/ts-gen/ips.ts",
			},
			// {
			// 	Path:       "github.com/seekr-osint/seekr/api/history",
			// 	OutputPath: "./../../web/ts-gen/history.ts",
			// },
		},
	}
	gen := tygo.New(config)
	err := gen.Generate()
	if err != nil {
		panic(err)
	}
}

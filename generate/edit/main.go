// UNUSED CODE 
package main

import (
	// "fmt"
	"os"
	"text/template"
	// "strings"

	"github.com/seekr-osint/seekr/api/person"
	"github.com/seekr-osint/seekr/generate/tsparse"
)

// import (
// 	"os"
// 	"text/template"

// 	"github.com/seekr-osint/seekrstack/api/db"
// 	"github.com/seekr-osint/seekrstack/api/person"
// )

func main() {

	data, err := os.ReadFile("../../web/ts-tmpl/edit.tmpl")
	if err != nil {
		panic(err)
	}
	// template.FuncMap{"createTable": func(p person.Person) string { return db.GenerateCreateTableSQL(p, "people") }},
	tmpl, err := template.New("tstmpl").Funcs(template.FuncMap{}).Parse(string(data))
	if err != nil {
		panic(err)
	}
	outputFile, err := os.Create("../../web/ts-gen/edit.ts")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	fieldNames := tsparse.GetFieldNames(person.Person{})

	in := struct {
		Fields []tsparse.FieldInfo
	}{
		Fields: fieldNames,
	}

	err = tmpl.Execute(outputFile, in)
	if err != nil {
		panic(err)
	}
}

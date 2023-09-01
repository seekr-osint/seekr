package enum

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

type TemplateData struct {
	TypeName      string
	LowerTypeName string
	TSEnum        string
	Values        []string
	NullValue     string
}

const codeTemplate = `
import (
	"github.com/seekr-osint/seekr/api/enum"
)
type {{.TypeName}} string
type {{.TypeName}}Enum struct{
	{{.TypeName}} enum.Enum[{{.TypeName}}] ` + "`" + `json:"{{.LowerTypeName}}" tstype:"{{.TSEnum}}" example:"{{index .Values 0}}"` + "`" + `
}

// func ({{.LowerTypeName}} {{.TypeName}}Enum) Value() (driver.Value, error) {
// 	return {{.LowerTypeName}}.{{.TypeName}}.Vlaue()
// }

// func ({{.LowerTypeName}} {{.TypeName}}Enum) Scan(value interface{}) error {
// 	return {{.LowerTypeName}}.{{.TypeName}}.Scan(value)
// }


func ({{.LowerTypeName}} {{.TypeName}}) Values() []{{.TypeName}} {
	return []{{.TypeName}}{ {{range $index, $value := .Values}}{{if $index}}, {{end}}"{{$value}}"{{end}} }
}

func ({{.LowerTypeName}} {{.TypeName}}) NullValue() {{.TypeName}} {
	return "{{.NullValue}}"
}

`

func CreateTemplateData(typeName string, values []string, nullValue string) TemplateData {
	tsenum := ""
	for _, i := range values {
		tsenum = fmt.Sprintf("%s'%s' | ", tsenum, i)
	}
	tsenum = fmt.Sprintf("%s'%s'", tsenum, nullValue)
	return TemplateData{
		TypeName:      typeName,
		LowerTypeName: strings.ToLower(typeName),
		Values:        values,
		NullValue:     nullValue,
		TSEnum:        tsenum,
	}
}

func (d TemplateData) GenerateEnumCode() (string, error) {

	tmpl, err := template.New("codeTemplate").Parse(codeTemplate)
	if err != nil {
		return "", err
	}

	var code bytes.Buffer
	err = tmpl.Execute(&code, d)
	if err != nil {
		return "", err
	}

	return code.String(), nil
}

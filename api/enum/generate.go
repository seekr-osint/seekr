package enum

import (
	"bytes"
	"strings"
	"text/template"
)

type TemplateData struct {
	TypeName      string
	LowerTypeName string
	Values        []string
	NullValue     string
}

const codeTemplate = `
type {{.TypeName}} string

func ({{.LowerTypeName}} {{.TypeName}}) Values() []{{.TypeName}} {
	return []{{.TypeName}}{ {{range $index, $value := .Values}}{{if $index}}, {{end}}"{{$value}}"{{end}} }
}

func ({{.LowerTypeName}} {{.TypeName}}) NullValue() {{.TypeName}} {
	return "{{.NullValue}}"
}
`

func CreateTemplateData(typeName string, values []string, nullValue string) TemplateData {
	return TemplateData{
		TypeName:      typeName,
		LowerTypeName: strings.ToLower(typeName),
		Values:        values,
		NullValue:     nullValue,
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

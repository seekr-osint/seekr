//go:generate go run main.go
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/seekr-osint/seekrstack/api/enum"
)
func writeToFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func WriteEnum(typeName string, values []string, nullValue string) error {
	templateData := enum.CreateTemplateData(typeName, values, nullValue)
	code, err := templateData.GenerateEnumCode()
	if err != nil {
		fmt.Println("Error generating code:", err)
		return err
	}

	filename := fmt.Sprintf("../../api/enums/%s_enum.go", strings.ToLower(templateData.LowerTypeName))
	err = writeToFile(filename, fmt.Sprintf("package enums\n%s\n", code))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	fmt.Printf("Code written to %s\n", filename)
	return nil
}

func main() {
	err := WriteEnum("Gender", []string{"male", "female", "other"}, "")
	if err != nil {
		panic(err)
	}

	err = WriteEnum("Ethnicity", []string{"African", "Asian", "Caucasian/White", "Hispanic/Latino", "Indigenous/Native American", "Multiracial/Mixed"}, "")
	if err != nil {
		panic(err)
	}


}

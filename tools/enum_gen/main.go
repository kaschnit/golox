package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
	"time"
)

const enumFileTemplate = `// GENERATED - DO NOT MODIFY THIS FILE DIRECTLY.
// To modify this file, edit the corresponding JSON file "{{.SourceFileName}}".
// This file was generated on {{.Date}} from source {{.SourceFileName}}
package {{.Data.Package}}

type {{.Data.Type}} int

const (
	{{range $index, $enum := .Data.Enumerations -}}
	{{if eq $index 0}}{{$enum}} {{$.Data.Type}} = iota{{else}}{{$enum}}{{end}}
	{{end}}
)

`

const stringFileTemplate = `// GENERATED - DO NOT MODIFY THIS FILE DIRECTLY.
// To modify this file, edit the corresponding JSON file "{{.SourceFileName}}".
// This file was generated on {{.Date}} from source {{.SourceFileName}}
package {{.Data.Package}}

func ({{.Data.Type | LowerCase}} {{.Data.Type}}) String() string {
	switch {{.Data.Type | LowerCase}} {
	{{range $enum := .Data.Enumerations -}}
	case {{$enum}}:
		return "{{$enum}}"
	{{end -}}
	default:
		return ""
	}
}
`

type TemplateData struct {
	Type         string   `json:"type"`
	Package      string   `json:"package"`
	Enumerations []string `json:"enumerations"`
}

type TemplateParams struct {
	Data           TemplateData
	SourceFileName string
	Date           string
}

func loadJson(jsonPath string) (*TemplateParams, error) {
	jsonFile, err := os.Open(jsonPath)
	defer jsonFile.Close()
	if err != nil {
		return nil, err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	// Load the JSON
	var data TemplateData
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, err
	}

	return &TemplateParams{
		Data:           data,
		SourceFileName: path.Base(jsonPath),
		Date:           time.Now().String(),
	}, nil
}

func lowercase(s string) string {
	result := []rune(s)
	replacement := strings.ToLower(string(result[0]))
	result[0] = []rune(replacement)[0]
	return string(result)
}

func renderTemplate(templateStr string, params *TemplateParams) ([]byte, error) {
	funcMap := template.FuncMap{
		"LowerCase": lowercase,
	}

	// Render the templates
	tmpl, err := template.New(params.Data.Package).Funcs(funcMap).Parse(templateStr)
	if err != nil {
		return nil, err
	}

	// Format the code with gofmt
	var buf bytes.Buffer
	tmpl.Execute(&buf, params)
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, err
	}

	return formatted, nil
}

func getEnumFilePath(directory string, packageName string) string {
	return path.Join(directory, fmt.Sprintf("%s.go", packageName))
}

func getStringFilePath(directory string, packageName string) string {
	return path.Join(directory, fmt.Sprintf("%s_string.go", packageName))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Command-line args
	jsonPath := flag.String("json", "", "The path to the JSON used to generate the enum code.")
	outputPath := flag.String("output", "", "The output directory where the the generated code will be written.")
	flag.Parse()

	// Check that the args were passed
	if *jsonPath == "" || *outputPath == "" {
		println("Required args not provided")
		flag.Usage()
		os.Exit(1)
	}

	// Load the json definition
	params, err := loadJson(*jsonPath)
	check(err)

	// Render the templates to generate the code
	enumRendered, err := renderTemplate(enumFileTemplate, params)
	check(err)
	stringRendered, err := renderTemplate(stringFileTemplate, params)
	check(err)

	// Write the generated code
	err = os.MkdirAll(*outputPath, 0755)
	check(err)
	err = ioutil.WriteFile(getEnumFilePath(*outputPath, params.Data.Package), enumRendered, 0644)
	check(err)
	err = ioutil.WriteFile(getStringFilePath(*outputPath, params.Data.Package), stringRendered, 0644)
	check(err)
}

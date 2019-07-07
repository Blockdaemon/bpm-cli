package templates

import (
	"bytes"
	"text/template"
	"io/ioutil"
	"fmt"

	"gitlab.com/Blockdaemon/bpm/pkg/models"
)

// TemplateRendered renders a template with node confguration and writes it to disk if it doesn't exist yet
//
// In order to allow comma separated lists in the template it defines the template
// function `notLast` which can be used like this:
//
//		{{range $index, $id:= .Config.core.quorum_set_ids -}}
//		"${{ $id }}"{{if notLast $index $.Config.core.quorum_set_ids}},{{end}}
//		{{end -}}
//
func TemplateRendered(outputFilename, templateContent string, configuration models.NodeConfiguration) error {
	exists, err := models.FileExists(outputFilename)
	if err != nil {
		return err
	}

	if exists {
		fmt.Printf("File '%s' already exists, skipping rendering of template\n", outputFilename)
		return nil
	}

	fmt.Printf("Rendering template to file '%s'\n", outputFilename)

	var templateFunctions = template.FuncMap{
		"notLast": func(x int, a []interface{}) bool {
			return x != len(a)-1
		},
	}

	tmpl, err := template.New(outputFilename).Funcs(templateFunctions).Parse(templateContent)
	if err != nil {
		return err
	}

	output := bytes.NewBufferString("")

	err = tmpl.Execute(output, configuration)
	if err != nil {
		return err
	}


	if err := ioutil.WriteFile(outputFilename, output.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}



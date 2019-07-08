package configuration

import (
	"bytes"
	"text/template"
	"io/ioutil"
	"fmt"
	"os"
	"path"

	"gitlab.com/Blockdaemon/bpm/pkg/node"
	"gitlab.com/Blockdaemon/bpm/internal/bpm/util"
)

// ConfigurationFileRendered renders a template with node confguration and writes it to disk if it doesn't exist yet
//
// In order to allow comma separated lists in the template it defines the template
// function `notLast` which can be used like this:
//
//		{{range $index, $id:= .Config.core.quorum_set_ids -}}
//		"${{ $id }}"{{if notLast $index $.Config.core.quorum_set_ids}},{{end}}
//		{{end -}}
//
func ConfigurationFileRendered(filename, templateContent string, node node.Node) error {
	configsDir, err := node.MakeConfigsDirectory()
	if err != nil {
		return err
	}

	outputFilename := path.Join(configsDir, filename)

	exists, err := util.FileExists(outputFilename)
	if err != nil {
		return err
	}

	if exists {
		fmt.Printf("Configuration file '%s' already exists, skipping creation\n", outputFilename)
		return nil
	}

	fmt.Printf("Writing configuration file '%s'\n", outputFilename)

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

	err = tmpl.Execute(output, node)
	if err != nil {
		return err
	}


	if err := ioutil.WriteFile(outputFilename, output.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}

func ConfigurationFileAbsent(filename string, node node.Node) error {
	configsDir := node.ConfigsDirectory()

	filePath := path.Join(configsDir, filename)

	exists, err := util.FileExists(filePath)
	if err != nil {
		return err
	}

	if !exists {
		fmt.Printf("Cannot find configuration file '%s', skipping removal\n", filePath)
		return nil
	}

	fmt.Printf("Removing configuration file '%s'\n", filePath)
	return os.Remove(filePath)
}


package tasks

import (
	"bytes"

	"github.com/landoop/tableprinter"
	"gitlab.com/Blockdaemon/runner/pkg/models"
)

// List contains functionality for the `list` cmd
//
// This has been seperated out into a function to make it easily testable
func List(apiKey, baseDir, pluginURL string) (string, error) {
	versionInfoExists, err := models.CheckVersionInfoExists(baseDir)
	if err != nil {
		return "", err
	}

	if !versionInfoExists {
		return VERSION_INFO_MISSING, nil
	}

	pluginListItems, err := models.ListPlugins(baseDir, pluginURL)
	if err != nil {
		return "", err
	}

	output := bytes.NewBufferString("")
	tableprinter.Print(output, pluginListItems)

	return output.String(), nil

}

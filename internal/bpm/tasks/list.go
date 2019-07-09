package tasks

import (
	"bytes"

	"github.com/landoop/tableprinter"
	"gitlab.com/Blockdaemon/bpm/internal/bpm/plugin"
)

// List contains functionality for the `list` cmd
//
// This has been seperated out into a function to make it easily testable
func List(baseDir, pluginURL string) (string, error) {
	if err := plugin.DownloadVersionInfo(pluginURL, baseDir); err != nil {
		return "", err
	}

	pluginListItems, err := plugin.ListPlugins(baseDir, pluginURL)
	if err != nil {
		return "", err
	}

	output := bytes.NewBufferString("")
	tableprinter.Print(output, pluginListItems)

	return output.String(), nil

}

package tasks

import (
	"gitlab.com/Blockdaemon/runner/models"
)

// Run contains functionality for the `run` cmd
//
// This has been seperated out into a function to make it easily testable
func Run(apiKey, baseDir, pluginURL, pluginName string) (string, error) {
	versionInfoExists, err := models.CheckVersionInfoExists(baseDir)
	if err != nil {
		return "", err
	}

	if !versionInfoExists {
		return VERSION_INFO_MISSING, nil
	}

	plugin, err := models.LoadPlugin(baseDir, pluginURL, pluginName)
	if err != nil {
		return "", err
	}

	return "", plugin.RunPlugin()
}

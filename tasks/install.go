package tasks

import (
	"gitlab.com/Blockdaemon/runner/models"
)

const VERSION_INFO_MISSING = "The version info list does not exist. Please run `refresh` first."

// Install contains functionality for the `install` cmd
//
// This has been seperated out into a function to make it easily testable
func Install(apiKey, baseDir, pluginURL, pluginName, version string) (string, error) {
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

	if len(version) > 0 {
		return "", plugin.InstallVersion(apiKey, version)
	}

	if err = plugin.InstallLatest(apiKey); err != nil {
		return "", err
	}

	return "Plugin succesfully installed", nil
}

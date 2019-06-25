package tasks

import (
	"gitlab.com/Blockdaemon/bpm/pkg/models"
)


// Install contains functionality for the `install` cmd
//
// This has been seperated out into a function to make it easily testable
func Install(apiKey, baseDir, pluginURL, pluginName, pluginVersion, runnerVersion string) (string, error) {
	if err := models.DownloadVersionInfo(apiKey, pluginURL, baseDir); err != nil {
		return "", err
	}

	upgradable, err := models.CheckRunnerUpgradable(baseDir, runnerVersion)
	if err != nil {
		return "", err
	}
	if upgradable {
		return TEXT_NEW_BPM_VERSION, nil
	}

	plugin, err := models.LoadPlugin(baseDir, pluginURL, pluginName)
	if err != nil {
		return "", err
	}

	if len(pluginVersion) > 0 {
		return "", plugin.InstallVersion(apiKey, pluginVersion)
	}

	if err = plugin.InstallLatest(apiKey); err != nil {
		return "", err
	}

	return TEXT_PLUGIN_INSTALLED, nil
}

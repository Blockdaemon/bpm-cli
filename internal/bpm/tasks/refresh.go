package tasks

import (
	"gitlab.com/Blockdaemon/runner/pkg/models"
)

// Refresh contains functionality for the `refresh` cmd
//
// This has been seperated out into a function to make it easily testable
func Refresh(apiKey, baseDir, pluginURL, runnerVersion string) (string, error) {
	if err := models.DownloadVersionInfo(apiKey, pluginURL, baseDir); err != nil {
		return "", err
	}

	upgradable, err := models.CheckRunnerUpgradable(baseDir, runnerVersion)
	if err != nil {
		return "", err
	}
	if upgradable {
		return "A new version of the runner is available, please upgrade!", nil
	}

	return "Succesfully downloaded version information. The runner is up to date!", nil
}

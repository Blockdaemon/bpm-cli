package tasks

import (
	"gitlab.com/Blockdaemon/bpm/internal/bpm/plugin"
)


// Install contains functionality for the `install` cmd
//
// This has been seperated out into a function to make it easily testable
func Install(baseDir, pluginURL, pluginName, pluginVersion, runnerVersion string) (string, error) {
	if err := plugin.DownloadVersionInfo(pluginURL, baseDir); err != nil {
		return "", err
	}

	upgradable, err := plugin.CheckRunnerUpgradable(baseDir, runnerVersion)
	if err != nil {
		return "", err
	}
	if upgradable {
		return TEXT_NEW_BPM_VERSION, nil
	}

	pluginToInstall, err := plugin.LoadPlugin(baseDir, pluginURL, pluginName)
	if err != nil {
		return "", err
	}

	if len(pluginVersion) > 0 {
		return "", pluginToInstall.InstallVersion(pluginVersion)
	}

	if err = pluginToInstall.InstallLatest(); err != nil {
		return "", err
	}

	return TEXT_PLUGIN_INSTALLED, nil
}

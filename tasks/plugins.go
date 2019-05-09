package tasks

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func getPluginFilename(baseDir, pluginName string) (string, error) {
	pluginDir, err := makeDirectory(baseDir, "plugins")
	if err != nil {
		return "", err
	}

	return filepath.Join(pluginDir, pluginName), nil
}

func getPluginURL(baseURL, apiKey, pluginName, version, GOOS, GOARCH string) string {
	path := fmt.Sprintf("%s-%s-%s-%s", pluginName, version, GOOS, GOARCH)

	return buildURL(baseURL, path, apiKey)
}

func InstallPluginVersion(baseDir, baseURL, apiKey, pluginName, version string) error {
	pluginFilename, err := getPluginFilename(baseDir, pluginName)
	if err != nil {
		return err
	}

	pluginURL := getPluginURL(baseURL, apiKey, pluginName, version, runtime.GOOS, runtime.GOARCH)

	if err := downloadFile(pluginFilename, pluginURL); err != nil {
		return err
	}

	return os.Chmod(pluginFilename, 0700)
}

func InstallPluginLatest(baseDir, baseURL, apiKey, pluginName string) error {
	versionInfo, err := LoadVersionInfo(baseDir)
	if err != nil {
		return err
	}

	pluginInfo, ok := versionInfo.GetPluginInfo(pluginName)

	if !ok {
		return fmt.Errorf("unknown plugin: %s", pluginName)
	}

	return InstallPluginVersion(baseDir, baseURL, apiKey, pluginName, pluginInfo.Version)
}

func ListPlugins(baseDir string) ([]PluginListItem, error) {
	versionInfo, err := LoadVersionInfo(baseDir)
	if err != nil {
		return nil, err
	}

	pluginListItems := []PluginListItem{}

	for _, info := range versionInfo.Plugins {
		plugin := NewPlugin(info, baseDir)
		installedVersion, err := plugin.RunVersionCommand()
		if err != nil {
			return pluginListItems, fmt.Errorf("cannot get installed version of plugin '%s': %s", plugin.Info.Name, err)
		}

		pluginListItems = append(pluginListItems, PluginListItem{
			Name:             plugin.Info.Name,
			AvailableVersion: plugin.Info.Version,
			InstalledVersion: installedVersion,
		})

	}

	return pluginListItems, nil
}

func CheckPluginUpgradable(baseDir, pluginName string) (bool, error) {
	plugin, err := LoadPlugin(baseDir, pluginName)
	if err != nil {
		return false, err
	}

	return plugin.NeedsUpgrade()
}

func RunPlugin(baseDir, pluginName string) error {
	plugin, err := LoadPlugin(baseDir, pluginName)
	if err != nil {
		return err
	}

	// TODO: Might need changes depending on the plugin

	_, _ = plugin.RunCommand("create-secrets")
	_, _ = plugin.RunCommand("pull-config")
	_, _ = plugin.RunCommand("configure")
	_, _ = plugin.RunCommand("validate")
	_, _ = plugin.RunCommand("start")

	return nil

}

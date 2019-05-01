package tasks

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func downloadFile(filepath string, url string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

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

	for _, pluginInfo := range versionInfo.Plugins {
		installedVersion, err := pluginInfo.RunVersionCommand(baseDir)
		if err != nil {
			return pluginListItems, fmt.Errorf("cannot get installed version of plugin '%s': %s", pluginInfo.Name, err)
		}

		pluginListItems = append(pluginListItems, PluginListItem{
			Name:             pluginInfo.Name,
			AvailableVersion: pluginInfo.Version,
			InstalledVersion: installedVersion,
		})

	}

	return pluginListItems, nil
}

func CheckPluginUpgradable(baseDir, pluginName string) (bool, error) {
	pluginInfo, err := LoadPluginInfo(baseDir, pluginName)
	if err != nil {
		return false, err
	}

	return pluginInfo.NeedsUpgrade(baseDir)
}

func RunPlugin(baseDir, pluginName string) error {
	pluginInfo, err := LoadPluginInfo(baseDir, pluginName)
	if err != nil {
		return err
	}

	// TODO: Might need changes depending on the plugin

	_, _ = pluginInfo.RunCommand(baseDir, "create-secrets")
	_, _ = pluginInfo.RunCommand(baseDir, "pull-config")
	_, _ = pluginInfo.RunCommand(baseDir, "configure")
	_, _ = pluginInfo.RunCommand(baseDir, "validate")
	_, _ = pluginInfo.RunCommand(baseDir, "start")

	return nil

}

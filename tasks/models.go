package tasks

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type PluginInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (i PluginInfo) RunCommand(baseDir, command string) (string, error) {
	filename, err := getPluginFilename(baseDir, i.Name)
	if err != nil {
		return "", err
	}

	cmd := exec.Command(filename, command)
	output, err := cmd.CombinedOutput()

	if err != nil {
		_, isPathError := err.(*os.PathError)

		if isPathError {
			// Looks like that plugin isn't installed
			return "", nil
		}

		// Plugin is installed but something else is wrong
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func (i PluginInfo) RunVersionCommand(baseDir string) (string, error) {
	return i.RunCommand(baseDir, "version")
}

func (i PluginInfo) NeedsUpgrade(baseDir string) (bool, error) {
	installedVersionStr, err := i.RunVersionCommand(baseDir)
	if err != nil {
		return false, fmt.Errorf("cannot get installed version of plugin '%s': %s", i.Name, err)
	}

	return needsUpgrade(installedVersionStr, i.Version)
}

type VersionInfo struct {
	RunnerVersion string       `json:"runner-version"`
	Plugins       []PluginInfo `json:"plugins"`
}

func (v VersionInfo) GetPluginInfo(pluginName string) (PluginInfo, bool) {
	for _, pluginInfo := range v.Plugins {
		if pluginName == pluginInfo.Name {
			return pluginInfo, true
		}
	}

	return PluginInfo{}, false
}

type PluginListItem struct {
	Name             string `header:"Name"`
	InstalledVersion string `header:"Installed Version"`
	AvailableVersion string `header:"Available Version"`
}

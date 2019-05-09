package tasks

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// VersionInfo contains information about versions for the plugins and the runner
type VersionInfo struct {
	RunnerVersion string       `json:"runner-version"`
	Plugins       []PluginInfo `json:"plugins"`
}

// GetPluginInfo returns a particular PluginInfo
func (v VersionInfo) GetPluginInfo(pluginName string) (PluginInfo, bool) {
	for _, pluginInfo := range v.Plugins {
		if pluginName == pluginInfo.Name {
			return pluginInfo, true
		}
	}

	return PluginInfo{}, false
}

// PluginInfo contains information about an available (but not necessarily installed) plugin
type PluginInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Plugin contains information and functions for an installed plugin
type Plugin struct {
	Info    PluginInfo
	baseDir string
}

// RunCommand runs a particular command with this plugin
func (i Plugin) RunCommand(command string) (string, error) {
	filename, err := getPluginFilename(i.baseDir, i.Info.Name)
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

// RunVersionCommand runs the `version` command on the plugin
func (i Plugin) RunVersionCommand() (string, error) {
	return i.RunCommand("version")
}

// NeedsUpgrade checks if this plugin needs to be upgraded
func (i Plugin) NeedsUpgrade() (bool, error) {
	installedVersionStr, err := i.RunVersionCommand()
	if err != nil {
		return false, fmt.Errorf("cannot get installed version of plugin '%s': %s", i.Info.Name, err)
	}

	return needsUpgrade(installedVersionStr, i.Info.Version)
}

// NewPlugin creates a new plugin from a PluginInfo
func NewPlugin(info PluginInfo, baseDir string) Plugin {
	return Plugin{
		Info:    info,
		baseDir: baseDir,
	}
}

// PluginListItem contains the values used to print a row of the `list` command
type PluginListItem struct {
	Name             string `header:"Name"`
	InstalledVersion string `header:"Installed Version"`
	AvailableVersion string `header:"Available Version"`
}

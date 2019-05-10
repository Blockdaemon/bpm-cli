package tasks

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
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

// Plugin contains information and functions for an installed (or to be installed) plugin
type Plugin struct {
	Info    PluginInfo
	baseDir string
	baseURL string
}

func (i Plugin) getPluginFilename() (string, error) {
	pluginDir, err := makeDirectory(i.baseDir, "plugins")
	if err != nil {
		return "", err
	}

	return filepath.Join(pluginDir, i.Info.Name), nil
}

func (i Plugin) getPluginURL(apiKey, version, GOOS, GOARCH string) string {
	path := fmt.Sprintf("%s-%s-%s-%s", i.Info.Name, version, GOOS, GOARCH)

	return buildURL(i.baseURL, path, apiKey)
}

// InstallVersion installs a particular version of the plugin
func (i Plugin) InstallVersion(apiKey, version string) error {
	pluginFilename, err := i.getPluginFilename()
	if err != nil {
		return err
	}

	pluginURL := i.getPluginURL(apiKey, version, runtime.GOOS, runtime.GOARCH)

	if err := downloadFile(pluginFilename, pluginURL); err != nil {
		return err
	}

	return os.Chmod(pluginFilename, 0700)
}

// InstallLatest installs the latest version of the plugin
func (i Plugin) InstallLatest(apiKey string) error {
	return i.InstallVersion(apiKey, i.Info.Version)
}

// RunCommand runs a particular command with this plugin
func (i Plugin) RunCommand(command string) (string, error) {
	fmt.Printf("Running plugin %s with command %s\n", i.Info.Name, command)
	filename, err := i.getPluginFilename()
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

// RunPlugin runs through the plugin lifecycle
func (i Plugin) RunPlugin() error {

	// TODO: Might need changes depending on the plugin

	_, _ = i.RunCommand("create-secrets")
	_, _ = i.RunCommand("pull-config")
	_, _ = i.RunCommand("configure")
	_, _ = i.RunCommand("validate")
	_, _ = i.RunCommand("start")

	return nil

}

// NewPlugin creates a new plugin from a PluginInfo
func NewPlugin(info PluginInfo, baseDir, baseURL string) Plugin {
	return Plugin{
		Info:    info,
		baseDir: baseDir,
		baseURL: baseURL,
	}
}

// PluginListItem contains the values used to print a row of the `list` command
type PluginListItem struct {
	Name             string `header:"Name"`
	InstalledVersion string `header:"Installed Version"`
	AvailableVersion string `header:"Available Version"`
}

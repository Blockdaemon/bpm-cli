package models

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

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

func (i Plugin) IsInstalled() (bool, error) {
	filename, err := i.getPluginFilename()
	if err != nil {
		return false, err
	}

	return FileExists(filename)
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
func (i Plugin) RunCommand(command, nodeGID string) (string, error) {
	fmt.Printf("Running plugin %s with command %s\n", i.Info.Name, command)

	filename, err := i.getPluginFilename()
	if err != nil {
		return "", err
	}

	var cmd *exec.Cmd

	if nodeGID != "" {
		cmd = exec.Command(filename, command, nodeGID)
	} else {
		cmd = exec.Command(filename, command)
	}
	output, err := cmd.CombinedOutput()

	// if err != nil {
	// 	_, isPathError := err.(*os.PathError)

	// 	if isPathError {
	// 		// Looks like that plugin isn't installed
	// 		return "", nil
	// 	}

	// 	// Plugin is installed but something else is wrong
	// 	return "", err
	// }

	return strings.TrimSpace(string(output)), err
}

// RunVersionCommand runs the `version` command on the plugin
func (i Plugin) RunVersionCommand() (string, error) {
	return i.RunCommand("version", "")
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
func (i Plugin) RunPlugin(nodeGID string) error {

	output, err := i.RunCommand("create-secrets", nodeGID)
	fmt.Println(indent(output, "    "))
	if err != nil {
		return err
	}

	output, err = i.RunCommand("create-configurations", nodeGID)
	fmt.Println(indent(output, "    "))
	if err != nil {
		return err
	}

	output, err = i.RunCommand("initialize", nodeGID)
	fmt.Println(indent(output, "    "))
	if err != nil {
		return err
	}

	output, err = i.RunCommand("start", nodeGID)
	fmt.Println(indent(output, "    "))
	if err != nil {
		return err
	}

	configuration, err := LoadConfiguration(i.baseDir, nodeGID)
	if err != nil {
		return err
	}

	version, err := i.RunVersionCommand()
	if err != nil {
		return err
	}

	// After everything is done, write the current version so we know where we are in case of upgrades
	configuration.WritePluginVersion(version)

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

// LoadPlugin loads plugin information from disk
func LoadPlugin(baseDir, baseURL, pluginName string) (Plugin, error) {
	versionInfo, err := LoadVersionInfo(baseDir)
	if err != nil {
		return Plugin{}, err
	}

	info, ok := versionInfo.GetPluginInfo(pluginName)
	if !ok {
		return Plugin{}, fmt.Errorf("unknown plugin: %s", pluginName)
	}

	plugin := NewPlugin(info, baseDir, baseURL)

	return plugin, nil
}

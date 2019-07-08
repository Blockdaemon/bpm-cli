package models

import (
	"fmt"
)

// PluginListItem contains the values used to print a row of the `list` command
type PluginListItem struct {
	Name             string `header:"Name"`
	InstalledVersion string `header:"Installed Version"`
	AvailableVersion string `header:"Available Version"`
}

// ListPlugins lists all plugins with currently installed version and available version
func ListPlugins(baseDir, baseURL string) ([]PluginListItem, error) {
	versionInfo, err := LoadVersionInfo(baseDir)
	if err != nil {
		return nil, err
	}

	pluginListItems := []PluginListItem{}

	for _, info := range versionInfo.Plugins {
		plugin := NewPlugin(info, baseDir, baseURL)


		installed, err := plugin.IsInstalled()
		if err != nil {
			return pluginListItems, err
		}

		installedVersion := "not installed"

		if installed {
			installedVersion, err = plugin.RunVersionCommand()
			if err != nil {
				return pluginListItems, fmt.Errorf("cannot get installed version of plugin '%s': %s", plugin.Info.Name, err)
			}
		}

		pluginListItems = append(pluginListItems, PluginListItem{
			Name:             plugin.Info.Name,
			AvailableVersion: plugin.Info.Version,
			InstalledVersion: installedVersion,
		})

	}

	return pluginListItems, nil
}

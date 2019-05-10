package tasks

import (
	"fmt"
)

func ListPlugins(baseDir, baseURL string) ([]PluginListItem, error) {
	versionInfo, err := LoadVersionInfo(baseDir)
	if err != nil {
		return nil, err
	}

	pluginListItems := []PluginListItem{}

	for _, info := range versionInfo.Plugins {
		plugin := NewPlugin(info, baseDir, baseURL)
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

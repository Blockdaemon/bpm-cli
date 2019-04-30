package tasks

type PluginInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
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

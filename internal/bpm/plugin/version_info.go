package plugin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"gitlab.com/Blockdaemon/bpm/internal/bpm/util"
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

// LoadVersionInfo reads the version info from disk
func LoadVersionInfo(baseDir string) (VersionInfo, error) {
	var versionInfo VersionInfo

	versionFilePath, err := util.GetVersionInfoFilename(baseDir)
	if err != nil {
		return versionInfo, err
	}
	data, err := ioutil.ReadFile(versionFilePath)
	if err != nil {
		return versionInfo, err
	}

	if err = json.Unmarshal(data, &versionInfo); err != nil {
		return versionInfo, err
	}

	return versionInfo, nil
}

// CheckRunnerUpgradable checks if there is a new version of the runner according to the version info on disk
func CheckRunnerUpgradable(baseDir string, runnerVersion string) (string, error) {
	if runnerVersion == "development" || runnerVersion == "master" {
		fmt.Printf("Skipping check if runner is upgradable during development!\n")
		return "", nil
	}

	versionInfo, err := LoadVersionInfo(baseDir)
	if err != nil {
		return "", err
	}

	upgradable, err := util.NeedsUpgrade(runnerVersion, versionInfo.RunnerVersion) 
	if err != nil {
		return "", err
	}

	if upgradable {
		return versionInfo.RunnerVersion, nil
	}

	return "", nil
}

// DownloadVersionInfo downloads the version info onto disk
func DownloadVersionInfo(baseURL string, baseDir string) error {
	fullURL := util.BuildURL(baseURL, "version-info.json", "")

	fmt.Printf("Downloading version info from %s\n", fullURL)

	resp, err := http.Get(fullURL)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s (%d)", resp.Status, resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Check if it is valid JSON and can be unmarshalled
	var versionInfo VersionInfo
	if err = json.Unmarshal(body, &versionInfo); err != nil {
		return err
	}

	filePath, err := util.GetVersionInfoFilename(baseDir)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, body, 0644)
}

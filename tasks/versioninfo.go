package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

	versionFilePath, err := getVersionInfoFilename(baseDir)
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

// CheckVersionInfoExists checks if the version info file exists
func CheckVersionInfoExists(baseDir string) (bool, error) {
	versionFilePath, err := getVersionInfoFilename(baseDir)
	if err != nil {
		return false, err
	}

	_, err = os.Stat(versionFilePath)

	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// CheckRunnerUpgradable checks if there is a new version of the runner according to the version info on disk
func CheckRunnerUpgradable(baseDir string, runnerVersion string) (bool, error) {
	if runnerVersion == "development" {
		fmt.Printf("Skpping check if runner is upgradable during development!\n")
		return false, nil
	}

	versionInfo, err := LoadVersionInfo(baseDir)
	if err != nil {
		return false, err
	}

	return needsUpgrade(runnerVersion, versionInfo.RunnerVersion)
}

// DownloadVersionInfo downloads the version info onto disk
func DownloadVersionInfo(apiKey string, baseURL string, baseDir string) error {
	fullURL := buildURL(baseURL, "version-info.json", apiKey)
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

	filePath, err := getVersionInfoFilename(baseDir)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, body, 0644)
}

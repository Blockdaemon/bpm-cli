package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func getVersionInfoFilename(baseDir string) (string, error) {
	configDir, err := makeDirectory(baseDir, "config")
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, "version-info.json"), nil
}

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

func LoadPluginInfo(baseDir, pluginName string) (PluginInfo, error) {
	versionInfo, err := LoadVersionInfo(baseDir)
	if err != nil {
		return PluginInfo{}, err
	}

	pluginInfo, ok := versionInfo.GetPluginInfo(pluginName)
	if !ok {
		return PluginInfo{}, fmt.Errorf("unknown plugin: %s", pluginName)
	}

	return pluginInfo, nil
}

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

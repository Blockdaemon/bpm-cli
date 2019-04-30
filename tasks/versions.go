package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/coreos/go-semver/semver"
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

func CheckRunnerUpgradable(baseDir string, runnerVersion string) (bool, error) {
	if runnerVersion == "development" {
		fmt.Printf("Skpping check if runner is upgradable during development!\n")
		return false, nil
	}

	versionInfo, err := LoadVersionInfo(baseDir)
	if err != nil {
		return false, err
	}

	// Since current version is set at build-time we do not need to check
	// for an error explicitely. It will panic if it's not a valid version string
	currentVersion := semver.New(runnerVersion)

	availableVersion, err := semver.NewVersion(versionInfo.RunnerVersion)
	if err != nil {
		return false, err
	}

	if currentVersion.LessThan(*availableVersion) {
		return true, nil
	}

	return false, nil
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

package tasks

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/coreos/go-semver/semver"
	homedir "github.com/mitchellh/go-homedir"
)

func downloadFile(filepath string, url string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func makeDirectory(baseDir, subDir string) (string, error) {
	expandedBaseDir, err := homedir.Expand(baseDir)
	if err != nil {
		return "", err
	}

	path := filepath.Join(expandedBaseDir, subDir)

	// Create directory structure if it doesn't exist
	err = os.MkdirAll(path, os.ModePerm)
	return path, err
}

func buildURL(baseURL, path, apiKey string) string {
	var result = baseURL

	if !strings.HasSuffix(result, "/") {
		result = result + "/"
	}

	return result + path + "?apiKey=" + apiKey
}

func needsUpgrade(currentVersionStr, availableVersionStr string) (bool, error) {
	currentVersion, err := semver.NewVersion(currentVersionStr)
	if err != nil {
		return false, err
	}

	availableVersion, err := semver.NewVersion(availableVersionStr)
	if err != nil {
		return false, err
	}

	if currentVersion.LessThan(*availableVersion) {
		return true, nil
	}

	return false, nil
}

func getVersionInfoFilename(baseDir string) (string, error) {
	configDir, err := makeDirectory(baseDir, "config")
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, "version-info.json"), nil
}

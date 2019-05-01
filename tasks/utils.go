package tasks

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/coreos/go-semver/semver"
	homedir "github.com/mitchellh/go-homedir"
)

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

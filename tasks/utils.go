package tasks

import (
	"os"
	"path/filepath"
	"strings"

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

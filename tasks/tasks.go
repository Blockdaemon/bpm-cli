package tasks

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/logrusorgru/aurora"
	homedir "github.com/mitchellh/go-homedir"
)

var version string
var baseDir string

// overwriteBaseDir allows overwriting the baseDir for testing purposes
var overwriteBaseDir string

func getBaseDir() string {
	if overwriteBaseDir != "" {
		return overwriteBaseDir
	}

	return baseDir
}

func getDirectory(subDir string) (string, error) {
	expandedBaseDir, err := homedir.Expand(getBaseDir())
	if err != nil {
		return "", err
	}

	path := filepath.Join(expandedBaseDir, subDir)

	// Create directory structure if it doesn't exist
	err = os.MkdirAll(path, os.ModePerm)
	return path, err
}

//func getPluginsDir() (string, error) {
//	return getDirectory("plugins")
//}
//
//func getSecretsDir() (string, error) {
//	return getDirectory("secrets")
//}

func getConfigDir() (string, error) {
	return getDirectory("config")
}

func getPluginListFilename() (string, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, "available-plugins.json"), nil
}

func GetVersion() string {
	if version == "" {
		return "development"
	}

	return version
}

func CheckVersion() {
	fmt.Printf("%s Check if upgrade is necessary. Current version: %s\n", aurora.Red("TODO:"), GetVersion())
}

func DownloadPluginList(apiKey string) error {
	fmt.Printf("%s Fetch plugin list using '%s' as api key\n", aurora.Red("TODO:"), apiKey)

	mockData := `
	[
        {
            "name": "stellar-horizon",
            "description": "blabla",
            "version": "1.0.2",
            "download-url": "..."
        },
        {
            "name": "...",
            "description": "...",
            "version": "..."
            "download-url": "..."
        },
    ]
	`

	pluginListPath, err := getPluginListFilename()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(pluginListPath, []byte(mockData), 0644)
}

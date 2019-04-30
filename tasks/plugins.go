package tasks

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
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

func getPluginFilename(baseDir, pluginName string) (string, error) {
	pluginDir, err := makeDirectory(baseDir, "plugins")
	if err != nil {
		return "", err
	}

	return filepath.Join(pluginDir, pluginName), nil
}

func getPluginURL(baseURL, apiKey, pluginName, version, GOOS, GOARCH string) string {
	path := fmt.Sprintf("%s-%s-%s-%s", pluginName, version, GOOS, GOARCH)

	return buildURL(baseURL, path, apiKey)
}

func InstallPlugin(baseDir, baseURL, apiKey, pluginName string) error {
	versionInfo, err := LoadVersionInfo(baseDir)
	if err != nil {
		return err
	}

	pluginInfo, ok := versionInfo.GetPluginInfo(pluginName)

	if !ok {
		return fmt.Errorf("unknown plugin: %s", pluginName)
	}

	pluginFilename, err := getPluginFilename(baseDir, pluginName)
	if err != nil {
		return err
	}

	pluginURL := getPluginURL(baseURL, apiKey, pluginName, pluginInfo.Version, runtime.GOOS, runtime.GOARCH)

	if err := downloadFile(pluginFilename, pluginURL); err != nil {
		return err
	}

	return os.Chmod(pluginFilename, 0700)
}

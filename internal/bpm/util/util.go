package util

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

func DownloadFile(filepath string, url string) error {
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

func MakeDirectory(baseDir string, subDirs ...string) (string, error) {
	expandedBaseDir, err := homedir.Expand(baseDir)
	if err != nil {
		return "", err
	}

	subDirs = append([]string{expandedBaseDir}, subDirs...)	

	path := filepath.Join(subDirs...)

	// Create directory structure if it doesn't exist
	err = os.MkdirAll(path, os.ModePerm)
	return path, err
}

func BuildURL(baseURL, path, apiKey string) string {
	var result = baseURL

	if !strings.HasSuffix(result, "/") {
		result = result + "/"
	}

	return result + path + "?apiKey=" + apiKey
}

func NeedsUpgrade(currentVersionStr, availableVersionStr string) (bool, error) {
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

func GetVersionInfoFilename(baseDir string) (string, error) {
	configDir, err := MakeDirectory(baseDir, "config")
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, "version-info.json"), nil
}

func Indent(text, indent string) string {
	if text == "" {
		return ""
	}

	if text[len(text)-1:] == "\n" {
		result := ""
		for _, j := range strings.Split(text[:len(text)-1], "\n") {
			result += indent + j + "\n"
		}
		return result
	}
	result := ""
	for _, j := range strings.Split(strings.TrimRight(text, "\n"), "\n") {
		result += indent + j + "\n"
	}
	return result[:len(result)-1]
}


func FileExists(name string) (bool, error) {
    if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            return false, nil
        }

        return false, err
    }
    return true, nil
}

package manager

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadToFile(path, filename, url string) error {
	fullpath := filepath.Join(path, filename)

	// Create the file
	file, err := os.Create(fullpath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := os.Chmod(fullpath, 0700); err != nil {
		return err
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Bad Status: %s", resp.Status)
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

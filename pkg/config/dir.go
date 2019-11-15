package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func MakeDir(homeDir string, subDirs ...string) error {
	dirs := append([]string{homeDir}, subDirs...)

	// Check for directories
	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				return err
			}
		}
	}

	return nil
}

func ReadDirs(path string) ([]os.FileInfo, error) {
	var dirs []os.FileInfo

	infos, err := ioutil.ReadDir(path)
	if err != nil {
		return dirs, err
	}

	for _, info := range infos {
		if info.IsDir() {
			dirs = append(dirs, info)
		}
	}

	return dirs, nil
}

func NodesDir(homeDir string) string {
	return filepath.Join(homeDir, "nodes")
}

func PluginsDir(homeDir string) string {
	return filepath.Join(homeDir, "plugins")
}

func Init(homeDir string) error {
	return MakeDir(homeDir, NodesDir(homeDir), PluginsDir(homeDir))
}

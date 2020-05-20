package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// DeleteFileOrDirectory delets a file or directory
func DeleteFileOrDirectory(path, filename string) error {
	return os.RemoveAll(filepath.Join(path, filename))
}

// FileExists checks if a file exists in a particular path
func FileExists(path, filename string) bool {
	return PathExists(filepath.Join(path, filename))
}

// PathExists checks if a path exists
func PathExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		switch {
		case os.IsNotExist(err):
			fallthrough
		default:
			return false
		}
	}

	return true
}

// Read returns the contents of a file in a particular path
func Read(path, filename string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(path, filename))
}

// ReadFile returns the unmarshalled contents of a json file in a particular path
func ReadFile(path, filename string, v interface{}) error {
	data, err := Read(path, filename)
	if err != nil {
		return err
	}

	// Decode JSON
	return json.Unmarshal(data, v)
}

// WriteFile writes the json-marshalled content of an interface into a file in a particular path
func WriteFile(path, filename string, v interface{}) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(
		filepath.Join(path, filename),
		data,
		os.ModePerm,
	)
}

// Walk is a wrapper around filepath.Walk
//
// This can probably get refactored because it literally does nothing other than
// call filepath.Walk
func Walk(path string, f filepath.WalkFunc) error {
	return filepath.Walk(path, f)
}

// CopyFile copies a file
func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

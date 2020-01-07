package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func DeleteFile(path, filename string) error {
	return os.Remove(filepath.Join(path, filename))
}

func FileExists(path, filename string) bool {
	if _, err := os.Stat(filepath.Join(path, filename)); err != nil {
		switch {
		case os.IsNotExist(err):
			fallthrough
		default:
			return false
		}
	}

	return true
}

func Read(path, filename string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(path, filename))
}

func ReadFile(path, filename string, v interface{}) error {
	data, err := Read(path, filename)
	if err != nil {
		return err
	}

	// Decode JSON
	return json.Unmarshal(data, v)
}

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

func Walk(path string, f filepath.WalkFunc) error {
	return filepath.Walk(path, f)
}

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

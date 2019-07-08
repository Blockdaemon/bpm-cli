package plugin

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

func assertNoError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func assertError(err error, t *testing.T) {
	if err == nil {
		t.Errorf("expected error but there is none")
	}
}

func assertFileExists(path string, t *testing.T) {
	_, err := os.Stat(path)
	if err != nil {
		t.Errorf("cannot stat file %s: %s", path, err)
	}
}

func assertFileExistsWithPermissions(path string, mode os.FileMode, t *testing.T) {
	info, err := os.Stat(path)
	if err != nil {
		t.Errorf("cannot stat file %s: %s", path, err)
	}
	if info.Mode() != mode {
		t.Errorf("wrong permissions on file %s: %#o", path, info.Mode())
	}
}

func assertEqual(object interface{}, expected interface{}, t *testing.T) {
	if object != expected {
		t.Errorf("expected '%s' but got '%s'", expected, object)
	}
}

func setupBaseDir(t *testing.T) string {
	baseDir := "/tmp/blockdaemon-test"

	if err := os.MkdirAll(baseDir, os.ModePerm); err != nil {
		t.Error(fmt.Sprintf("cannot create base directory: %s", err))
	}

	return baseDir
}

func setupTestPlugin(baseDir string, mockPlugin string, t *testing.T) {
	if mockPlugin == "" {
		// If empty, use a default
		mockPlugin = `#!/bin/bash
		case "$1" in
		version)
			echo "1.0.0"
			;;
		esac`
	}

	pluginsPath := path.Join(baseDir, "plugins")

	if err := os.MkdirAll(pluginsPath, os.ModePerm); err != nil {
		t.Error(fmt.Sprintf("cannot create plugin directory: %s", err))
	}

	pluginFile := path.Join(pluginsPath, "test")

	if err := ioutil.WriteFile(pluginFile, []byte(mockPlugin), 0700); err != nil {
		t.Errorf("cannot write mock plugin: %s", err)
	}
}

func setupVersionInfo(baseDir string, mockData string, t *testing.T) {
	if mockData == "" {
		// If empty, use a default
		mockData = `
		{
			"runner-version": "1.2.3",
			"plugins": [
				{
					"name": "stellar-horizon",
					"version": "1.2.3"
				},
				{
					"name": "test",
					"version": "1.1.0"
				}
			]
		}`
	}

	configPath := path.Join(baseDir, "config")

	if err := os.MkdirAll(configPath, os.ModePerm); err != nil {
		t.Error(fmt.Sprintf("cannot create config directory: %s", err))
	}

	versionInfoPath := path.Join(configPath, "version-info.json")

	if err := ioutil.WriteFile(versionInfoPath, []byte(mockData), 0644); err != nil {
		t.Errorf("cannot write mock data: %s", err)
	}
}

func teardown(baseDir string, server *httptest.Server, t *testing.T) {
	if baseDir != "" {
		if err := os.RemoveAll(baseDir); err != nil {
			t.Error(fmt.Sprintf("cannot delete base directory: %s", err))
		}
	}

	if server != nil {
		server.Close()
	}
}

func setupMockHTTPServer(expectedPath string, expectedQuery string, mockBody []byte, t *testing.T) *httptest.Server {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path != expectedPath {
			t.Errorf("expected path '%s' but got '%s'", expectedPath, req.URL.Path)
		}
		if req.URL.RawQuery != expectedQuery {
			t.Errorf("expected query '%s' but got '%s'", expectedQuery, req.URL.RawQuery)
		}
		_, err := res.Write(mockBody)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	}))

	return testServer
}

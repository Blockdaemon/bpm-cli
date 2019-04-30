package tasks

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

func setupBaseDir(t *testing.T) string {
	baseDir := "/tmp/blockdaemon-test"

	if err := os.MkdirAll(baseDir, os.ModePerm); err != nil {
		t.Error(fmt.Sprintf("cannot create base directory: %s", err))
	}

	return baseDir
}

func setupVersionInfo(baseDir string, mockData string, t *testing.T) {
	if mockData == "" {
		mockData = `
		{
			"runner-version": "1.2.3",
			"plugins": [
				{
					"name": "stellar-horizon",
					"version": "1.2.3"
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
		res.Write(mockBody)
	}))

	return testServer
}

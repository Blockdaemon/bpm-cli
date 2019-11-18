package plugin

import (
	"fmt"
	"github.com/Blockdaemon/bpm/pkg/config"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const (
	testBaseDir = "/tmp/bpm-test"

	testPluginVersionResponse = `{
	    "data": {
	        "os": {
	            "name": "Linux",
	            "type": "linux"
	        },
	        "package": {
	            "description": "A testplugin",
	            "name": "testplugin"
	        },
	        "registryUrl": "%s/v1/install/testplugin-1.0.0-linux-amd64",
	        "version": "1.0.0"
	    }
	}
	`

	testPluginInstallResponse = `#!/bin/bash
		case "$1" in
		version)
		    1.0.0
		    ;;
		esac
		`
)

type TestContext struct {
	BaseDir string
	Server *httptest.Server
}

func setupHTTPServer(t *testing.T) *httptest.Server {
	var serverURL string
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		responsePayload := ""

		// return different responses based on the path & parameters
		if req.URL.Path == "/v1/packages/testplugin/versions/1.0.0" && req.URL.Query().Get("expand") == "package" {
			responsePayload = fmt.Sprintf(testPluginVersionResponse, serverURL)
		} else if req.URL.Path == "/v1/packages/testplugin/versions/latest" && req.URL.Query().Get("expand") == "package" {
			responsePayload = fmt.Sprintf(testPluginVersionResponse, serverURL)
		} else if req.URL.Path == "/v1/install/testplugin-1.0.0-linux-amd64" {
			responsePayload = testPluginInstallResponse
		}

		_, err := res.Write([]byte(responsePayload))
		if err != nil {
			t.Error(err)
		}
	}))

	serverURL = testServer.URL

	return testServer
}

func setupUnittest(t *testing.T) (PluginCmdContext, TestContext) {
	server := setupHTTPServer(t)
	if err := config.Init(testBaseDir); err != nil {
		t.Error(err)
	}

	manifest, err := config.LoadManifest(testBaseDir)
	if err != nil {
		t.Error(err)
	}

	return PluginCmdContext{
		HomeDir: testBaseDir,
		Manifest: manifest,
		RuntimeOS: "linux", // pretend to be linux during testing, this works on osx too!
		RegistryURL: server.URL,
		Debug: false,
	}, TestContext{
		BaseDir: testBaseDir,
		Server: server,
	} 
}

func teardownUnittest(testContext TestContext, t *testing.T) {
	if testContext.BaseDir != "" {
		if err := os.RemoveAll(testContext.BaseDir); err != nil {
			t.Error(fmt.Sprintf("cannot delete base directory: %s", err))
		}
	}

	if testContext.Server != nil {
		testContext.Server.Close()
	}
}

func assertEqual(actual string, expected string, t *testing.T) {
	if actual != expected {
		t.Errorf("expected '%s' but got '%s'", expected, actual)
	}
}

func assertError(err error, text string, t *testing.T) {
	assertEqual(err.Error(), text, t)
}

package plugin

import (
	"net/http"
	"net/http/httptest"
	"path"
	"testing"
)

func TestDownloadVersionInfo(t *testing.T) {
	baseDir := setupBaseDir(t)
	testServer := setupMockHTTPServer("/version-info.json", "", []byte("{}"), t)
	defer teardown(baseDir, testServer, t)

	err := DownloadVersionInfo(testServer.URL, baseDir)

	assertNoError(err, t)
	assertFileExists(path.Join(baseDir, "config", "version-info.json"), t)
}

func TestDownloadVersionInfoUnauthorized(t *testing.T) {
	// Mock http server
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusUnauthorized)
	}))
	defer func() { testServer.Close() }()

	err := DownloadVersionInfo(testServer.URL, "/tmp")
	assertError(err, t)
}

func TestDownloadVersionInfoInvalidJSON(t *testing.T) {
	baseDir := setupBaseDir(t)
	setupVersionInfo(baseDir, "", t)
	testServer := setupMockHTTPServer("/version-info.json", "", []byte("invalid}"), t)
	defer teardown(baseDir, testServer, t)

	err := DownloadVersionInfo(testServer.URL, baseDir)
	assertError(err, t)
}

func TestCheckRunnerUpgradable(t *testing.T) {
	baseDir := setupBaseDir(t)
	setupVersionInfo(baseDir, "", t)
	defer teardown(baseDir, nil, t)

	upgradable, err := CheckRunnerUpgradable(baseDir, "1.2.2")
	assertNoError(err, t)
	if !upgradable {
		t.Errorf("no upgrade available")
	}

	upgradable, err = CheckRunnerUpgradable(baseDir, "1.2.3")
	assertNoError(err, t)
	if upgradable {
		t.Errorf("unexpecedly there is an upgrade available")
	}
}

func TestCheckRunnerUpgradableSkipInDevelopment(t *testing.T) {
	baseDir := setupBaseDir(t)
	defer teardown(baseDir, nil, t)

	upgradable, err := CheckRunnerUpgradable(baseDir, "development")
	assertNoError(err, t)
	if upgradable {
		t.Errorf("unexpectedly there is an upgrade available")
	}
}

func TestCheckRunnerUpgradableNoFile(t *testing.T) {
	// Version file does not exist in /tmp!
	_, err := CheckRunnerUpgradable("/tmp", "1.2.2")
	assertError(err, t)
}

func TestCheckRunnerUpgradableInvalidJSON(t *testing.T) {
	baseDir := setupBaseDir(t)
	setupVersionInfo(baseDir, `"invalid json" }`, t)
	defer teardown(baseDir, nil, t)

	_, err := CheckRunnerUpgradable(baseDir, "1.2.2")
	assertError(err, t)
}

func TestCheckRunnerUpgradableInvalidVersion(t *testing.T) {
	baseDir := setupBaseDir(t)
	setupVersionInfo(baseDir, `{ "runner-version": "1-2-3" }`, t)
	defer teardown(baseDir, nil, t)

	_, err := CheckRunnerUpgradable(baseDir, "1.2.2")
	assertError(err, t)
}

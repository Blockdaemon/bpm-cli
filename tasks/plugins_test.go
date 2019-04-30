package tasks

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"testing"
)

func TestInstallPlugin(t *testing.T) {
	baseDir := setupBaseDir(t)
	setupVersionInfo(baseDir, "", t)
	expectedPath := fmt.Sprintf("/stellar-horizon-1.2.3-%s-%s", runtime.GOOS, runtime.GOARCH)
	testServer := setupMockHTTPServer(expectedPath, "apiKey=test", []byte("asdf"), t)
	defer teardown(baseDir, testServer, t)

	if err := InstallPlugin(baseDir, testServer.URL, "test", "stellar-horizon"); err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	pluginFile := path.Join(baseDir, "plugins", "stellar-horizon")
	info, err := os.Stat(pluginFile)
	if err != nil {
		t.Errorf("cannot stat file %s: %s", pluginFile, err)
	}
	if info.Mode() != 0700 {
		t.Errorf("wrong permissions: %#o", info.Mode())
	}
}

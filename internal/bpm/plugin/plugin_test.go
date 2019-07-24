package plugin

import (
	"fmt"
	"path"
	"runtime"
	"testing"
)

func TestPluginInstallLatest(t *testing.T) {
	baseDir := setupBaseDir(t)
	setupVersionInfo(baseDir, "", t)
	expectedPath := fmt.Sprintf("/stellar-horizon-1.2.3-%s-%s", runtime.GOOS, runtime.GOARCH)
	testServer := setupMockHTTPServer(expectedPath, "", []byte("asdf"), t)
	defer teardown(baseDir, testServer, t)

	plugin, err := LoadPlugin(baseDir, testServer.URL, "stellar-horizon")
	assertNoError(err, t)

	err = plugin.InstallLatest()
	assertNoError(err, t)

	pluginFile := path.Join(baseDir, "plugins", "stellar-horizon")
	assertFileExistsWithPermissions(pluginFile, 0700, t)
}

func TestPluginInstallVersion(t *testing.T) {
	baseDir := setupBaseDir(t)
	setupVersionInfo(baseDir, "", t)
	expectedPath := fmt.Sprintf("/stellar-horizon-1.2.3-%s-%s", runtime.GOOS, runtime.GOARCH)
	testServer := setupMockHTTPServer(expectedPath, "", []byte("asdf"), t)
	defer teardown(baseDir, testServer, t)

	plugin, err := LoadPlugin(baseDir, testServer.URL, "stellar-horizon")
	assertNoError(err, t)

	err = plugin.InstallVersion("1.2.3")
	assertNoError(err, t)

	pluginFile := path.Join(baseDir, "plugins", "stellar-horizon")
	assertFileExistsWithPermissions(pluginFile, 0700, t)
}

func TestPluginNeedsUpgrade(t *testing.T) {
	baseDir := setupBaseDir(t)
	setupVersionInfo(baseDir, `
	{
		"runner-version": "1.2.3",
		"plugins": [
			{
				"name": "test",
				"version": "1.1.0"
			}
		]
	}`, t)
	setupTestPlugin(baseDir, "", t)
	defer teardown(baseDir, nil, t)

	// Test plugin is version 1.0.0
	// First, test if it is upgradable
	plugin, err := LoadPlugin(baseDir, "", "test")
	assertNoError(err, t)
	upgradeVersion, err := plugin.NeedsUpgrade()
	assertNoError(err, t)
	if upgradeVersion == "" {
		t.Errorf("expected the plugin to be upgradable but it is not")
	}

	// Now, test the opposite: Not upgradable
	setupVersionInfo(baseDir, `
	{
		"runner-version": "1.2.3",
		"plugins": [
			{
				"name": "test",
				"version": "1.0.0"
			}
		]
	}`, t)

	plugin, err = LoadPlugin(baseDir, "", "test")
	assertNoError(err, t)
	upgradeVersion, err = plugin.NeedsUpgrade()
	assertNoError(err, t)
	if upgradeVersion != "" {
		t.Errorf("expected the plugin to NOT be upgradable but it is")
	}
}

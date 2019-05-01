package tasks

import (
	"fmt"
	"path"
	"runtime"
	"testing"
)

func TestInstallPluginLatest(t *testing.T) {
	baseDir := setupBaseDir(t)
	setupVersionInfo(baseDir, "", t)
	expectedPath := fmt.Sprintf("/stellar-horizon-1.2.3-%s-%s", runtime.GOOS, runtime.GOARCH)
	testServer := setupMockHTTPServer(expectedPath, "apiKey=test", []byte("asdf"), t)
	defer teardown(baseDir, testServer, t)

	if err := InstallPluginLatest(baseDir, testServer.URL, "test", "stellar-horizon"); err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	pluginFile := path.Join(baseDir, "plugins", "stellar-horizon")
	assertFileExistsWithPermissions(pluginFile, 0700, t)
}

func TestInstallPluginVersion(t *testing.T) {
	baseDir := setupBaseDir(t)
	expectedPath := fmt.Sprintf("/stellar-horizon-1.2.3-%s-%s", runtime.GOOS, runtime.GOARCH)
	testServer := setupMockHTTPServer(expectedPath, "apiKey=test", []byte("asdf"), t)
	defer teardown(baseDir, testServer, t)

	if err := InstallPluginVersion(baseDir, testServer.URL, "test", "stellar-horizon", "1.2.3"); err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	pluginFile := path.Join(baseDir, "plugins", "stellar-horizon")
	assertFileExistsWithPermissions(pluginFile, 0700, t)
}

func TestListPlugins(t *testing.T) {
	baseDir := setupBaseDir(t)
	setupVersionInfo(baseDir, "", t)
	setupTestPlugin(baseDir, "", t)
	defer teardown(baseDir, nil, t)

	pluginListItems, err := ListPlugins(baseDir)
	assertNoError(err, t)

	assertEqual(pluginListItems[0].Name, "stellar-horizon", t)
	assertEqual(pluginListItems[0].AvailableVersion, "1.2.3", t)
	assertEqual(pluginListItems[0].InstalledVersion, "", t) // not installed
	assertEqual(pluginListItems[1].Name, "test", t)
	assertEqual(pluginListItems[1].AvailableVersion, "1.1.0", t)
	assertEqual(pluginListItems[1].InstalledVersion, "1.0.0", t)
}

func TestListPluginsInvalidPlugin(t *testing.T) {
	baseDir := setupBaseDir(t)
	setupVersionInfo(baseDir, "", t)

	setupTestPlugin(baseDir, `#!/bin/bash
		exit 1`, t) // This plugin always exits with 1
	defer teardown(baseDir, nil, t)

	_, err := ListPlugins(baseDir)
	assertError(err, t)
}

func TestCheckPluginUpgradable(t *testing.T) {
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
	upgradable, err := CheckPluginUpgradable(baseDir, "test")
	assertNoError(err, t)
	if !upgradable {
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

	upgradable, err = CheckPluginUpgradable(baseDir, "test")
	assertNoError(err, t)
	if upgradable {
		t.Errorf("expected the plugin to NOT be upgradable but it is")
	}
}

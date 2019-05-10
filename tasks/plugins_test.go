package tasks

import (
	"testing"
)

func TestListPlugins(t *testing.T) {
	baseDir := setupBaseDir(t)
	setupVersionInfo(baseDir, "", t)
	setupTestPlugin(baseDir, "", t)
	defer teardown(baseDir, nil, t)

	pluginListItems, err := ListPlugins(baseDir, "dummy-url")
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

	_, err := ListPlugins(baseDir, "dummy-url")
	assertError(err, t)
}

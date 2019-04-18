package tasks

import (
	"os"
	"testing"
)

func TestDownloadPluginList(t *testing.T) {
	overwriteBaseDir = "/tmp/blockdaemon-test"
	defer os.RemoveAll(overwriteBaseDir)

	if err := DownloadPluginList("testKey"); err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	_, err := os.Stat("/tmp/blockdaemon-test/config/available-plugins.json")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

package plugin

import (
	"testing"

	"github.com/kami-zh/go-capturer"
)

func TestUninstall(t *testing.T) {
	cmdContext, testContext := setupUnittest(t)
	defer teardownUnittest(testContext, t)

	if err := cmdContext.Install("testplugin", "1.0.0"); err != nil {
		t.Error(err)
	}

	out := capturer.CaptureOutput(func() {
		if err := cmdContext.Uninstall("testplugin"); err != nil {
			t.Error(err)
		}
	})

	assertEqual(out, `The package "testplugin" has been uninstalled.`, t)
}

func TestUninstallNotInstalled(t *testing.T) {
	cmdContext, testContext := setupUnittest(t)
	defer teardownUnittest(testContext, t)

	err := cmdContext.Uninstall("testplugin")
	assertError(err, `The package "testplugin" is currently not installed.`, t)
}

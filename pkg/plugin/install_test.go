package plugin

import (
	"testing"

	"github.com/kami-zh/go-capturer"
)

func TestInstall(t *testing.T) {
	cmdContext, testContext := setupUnittest(t)
	defer teardownUnittest(testContext, t)

	out := capturer.CaptureOutput(func() {
		if err := cmdContext.Install("testplugin", "1.0.0"); err != nil {
			t.Error(err)
		}
	})

	assertEqual(out, `The package "testplugin" has been installed.`, t)
}

func TestInstallAlreadyInstalled(t *testing.T) {
	cmdContext, testContext := setupUnittest(t)
	defer teardownUnittest(testContext, t)

	if err := cmdContext.Install("testplugin", "1.0.0"); err != nil {
		t.Error(err)
	}

	err := cmdContext.Install("testplugin", "1.0.0")
	assertError(err, `"testplugin" version "1.0.0" has already been installed.`, t)
}

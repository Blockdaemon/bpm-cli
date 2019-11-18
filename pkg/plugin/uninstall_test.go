package plugin

import (
	"testing"
	"fmt"

	"github.com/kami-zh/go-capturer"
)

func TestUninstall(t *testing.T) {
	cmdContext, testContext := setupUnittest(t)
	defer teardownUnittest(testContext, t)

	_, err := cmdContext.Install("testplugin", "1.0.0")
	if err != nil {
		t.Error(err)
	}

	out := capturer.CaptureOutput(func() {
		output, err := cmdContext.Uninstall("testplugin")
		if err != nil {
			t.Error(err)
		}

		fmt.Print(output)
	})

	assertEqual(out, `The package "testplugin" has been uninstalled.`, t)
}

func TestUninstallNotInstalled(t *testing.T) {
	cmdContext, testContext := setupUnittest(t)
	defer teardownUnittest(testContext, t)

	_, err := cmdContext.Uninstall("testplugin")
	assertError(err, `The package "testplugin" is currently not installed.`, t)
}

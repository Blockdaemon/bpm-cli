package plugin

import (
	"testing"
	"fmt"

	"github.com/kami-zh/go-capturer"
)

func TestInstall(t *testing.T) {
	cmdContext, testContext := setupUnittest(t)
	defer teardownUnittest(testContext, t)

	out := capturer.CaptureOutput(func() {
		output, err := cmdContext.Install("testplugin", "1.0.0")
		if err != nil {
			t.Error(err)
		}

		fmt.Print(output)
	})

	assertEqual(out, `The package "testplugin" has been installed.`, t)
}

func TestInstallAlreadyInstalled(t *testing.T) {
	cmdContext, testContext := setupUnittest(t)
	defer teardownUnittest(testContext, t)

	_, err := cmdContext.Install("testplugin", "1.0.0")
	if err != nil {
		t.Error(err)
	}

	_, err = cmdContext.Install("testplugin", "1.0.0")
	assertError(err, `"testplugin" version "1.0.0" has already been installed.`, t)
}

package plugin

import (
	"testing"

	"github.com/kami-zh/go-capturer"
)

func TestConfigure(t *testing.T) {
	cmdContext, testContext := setupUnittest(t)
	defer teardownUnittest(testContext, t)

	if err := cmdContext.Install("testplugin", "1.0.0"); err != nil {
		t.Error(err)
	}

	out := capturer.CaptureOutput(func() {
		if err := cmdContext.Configure("testplugin", map[string]string{"subtype": "validator"}, map[string]bool{}, false); err != nil {
			t.Error(err)
		}
	})

	assertRegEx(out, `Node with id "(.*)" has been initialized\.`, t)
}

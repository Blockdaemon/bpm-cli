package plugin

import (
	"testing"
	"fmt"

	"github.com/kami-zh/go-capturer"
)

// Note: There is important whitespace in these strings!

const (
	expected = `  NAME | INSTALLED VERSION | AVAILABLE VERSION  
+------+-------------------+-------------------+
`

    expected2 = `     NAME    | INSTALLED VERSION | AVAILABLE VERSION  
+------------+-------------------+-------------------+
  testplugin | 1.0.0             | 1.0.0              
`
)


func TestList(t *testing.T) {
	cmdContext, testContext := setupUnittest(t)
	defer teardownUnittest(testContext, t)

	out := capturer.CaptureOutput(func() {
		output, err := cmdContext.List()
		if err != nil {
			t.Error(err)
		}
		fmt.Print(output)
	})

	assertEqual(out, expected, t)

	_, err := cmdContext.Install("testplugin", "1.0.0")
	if err != nil {
		t.Error(err)
	}

	out2 := capturer.CaptureOutput(func() {
		output, err := cmdContext.List()
		if err != nil {
			t.Error(err)
		}
		fmt.Print(output)
	})
	assertEqual(out2, expected2, t)
}


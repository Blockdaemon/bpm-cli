package manager

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// ExecCmd runs a particular command with this plugin
func ExecCmd(debug bool, pluginFilename string, args ...string) (string, error) {
	if debug {
		fmt.Printf("Running: %s %s\n", pluginFilename, strings.Join(args, " "))
	}

	cmd := exec.Command(pluginFilename, args...)
	output, err := cmd.CombinedOutput()

	cleanOutput := string(bytes.TrimSpace(output))

	if err != nil {
		if debug {
			fmt.Printf("Error: %s\n", err)
			fmt.Println(cleanOutput)
		}
		return cleanOutput, err
	}

	return cleanOutput, nil
}

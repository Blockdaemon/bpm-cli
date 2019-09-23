package manager

import (
	"bytes"
	"fmt"
	"os/exec"
)

// ExecCmd runs a particular command with this plugin
func ExecCmd(pluginName, pluginFilename string, args ...string) (string, error) {
	fmt.Printf("Running plugin %s with command %s\n", pluginName, args[0])

	cmd := exec.Command(pluginFilename, args...)
	output, err := cmd.CombinedOutput()

	cleanOutput := string(bytes.TrimSpace(output))

	if err != nil {
		return cleanOutput, err
	}

	return cleanOutput, nil
}

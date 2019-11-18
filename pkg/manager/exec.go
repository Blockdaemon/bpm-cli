package manager

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ExecCmd runs a particular command with this plugin
func ExecCmd(debug bool, pluginFilename string, args ...string) error {
	if debug {
		fmt.Printf("Running: %s %s\n", pluginFilename, strings.Join(args, " "))
	}

	cmd := exec.Command(pluginFilename, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// ExecCmdCapture runs a particular command with this plugin and returns it's output if succefull
func ExecCmdCapture(debug bool, pluginFilename string, args ...string) (string, error) {
	if debug {
		fmt.Printf("Running: %s %s\n", pluginFilename, strings.Join(args, " "))
	}

	cmd := exec.Command(pluginFilename, args...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	cleanOutput := string(bytes.TrimSpace(output))
	return cleanOutput, nil
}

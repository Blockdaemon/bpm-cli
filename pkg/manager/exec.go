package manager

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ExecCmd runs a particular command with this plugin
func ExecCmd(debug bool, executable string, args ...string) error {
	if debug {
		fmt.Printf("Running: %s %s\n", executable, strings.Join(args, " "))
	}

	cmd := exec.Command(executable, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// ExecCmdCapture runs a particular command with this plugin and returns it's output if succefull
func ExecCmdCapture(debug bool, executable string, args ...string) (string, error) {
	if debug {
		fmt.Printf("Running: %s %s\n", executable, strings.Join(args, " "))
	}

	cmd := exec.Command(executable, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("%s: %s", err, stderr.String())
	}

	cleanOutput := string(bytes.TrimSpace(stdout.Bytes()))
	return cleanOutput, nil
}

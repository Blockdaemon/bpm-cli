package tasks

import (
	"os"
	"fmt"
	"gitlab.com/Blockdaemon/bpm/internal/bpm/plugin"
	"gitlab.com/Blockdaemon/bpm/internal/bpm/util"
)

// Purge contains functionality for the `purge` cmd
//
// This has been seperated out into a function to make it easily testable
func Purge(baseDir, pluginURL, pluginName, runnerVersion string) (string, error) {
	pluginToRun, err := plugin.LoadPlugin(baseDir, pluginURL, pluginName)
	if err != nil {
		return "", err
	}

	gid := os.Getenv("MOCK_GID")
	if gid == "" {
		return "", fmt.Errorf("env variable `MOCK_GID` isn't set. This is just used temporarily until we get the token from the BPG")
	}

	output, err := pluginToRun.RunCommand("purge", gid)
	fmt.Println(util.Indent(output, "    "))
	return "", err
}

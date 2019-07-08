package tasks

import (
	"os"
	"fmt"
	"gitlab.com/Blockdaemon/bpm/internal/bpm/plugin"
	"gitlab.com/Blockdaemon/bpm/internal/bpm/util"
)

// Remove contains functionality for the `remove` cmd
//
// This has been seperated out into a function to make it easily testable
func Remove(apiKey, baseDir, pluginURL, pluginName, runnerVersion string) (string, error) {
	pluginToRun, err := plugin.LoadPlugin(baseDir, pluginURL, pluginName)
	if err != nil {
		return "", err
	}

	gid := os.Getenv("MOCK_GID")
	if gid == "" {
		return "", fmt.Errorf("env variable `MOCK_GID` isn't set. This is just used temporarily until we get the token from the BPG")
	}

	output, err := pluginToRun.RunCommand("remove", gid)
	fmt.Println(util.Indent(output, "    "))
	return "", err
}

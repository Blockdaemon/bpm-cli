package tasks

import (
	"fmt"
	"gitlab.com/Blockdaemon/bpm-sdk/pkg/node"
	"gitlab.com/Blockdaemon/bpm/internal/bpm/plugin"
	"gitlab.com/Blockdaemon/bpm/internal/bpm/util"
	"os"
)

// Remove contains functionality for the `remove` cmd
//
// This has been seperated out into a function to make it easily testable
func Remove(baseDir, pluginURL, pluginName, runnerVersion string, purge bool) (string, error) {
	pluginToRun, err := plugin.LoadPlugin(baseDir, pluginURL, pluginName)
	if err != nil {
		return "", err
	}

	gid := os.Getenv("MOCK_GID")
	if gid == "" {
		return "", fmt.Errorf("env variable `MOCK_GID` isn't set. This is just used temporarily until we get the token from the BPG")
	}

	var output string
	if purge {
		output, err = pluginToRun.RunCommand("remove", gid, "--purge")
	} else {
		output, err = pluginToRun.RunCommand("remove", gid)
	}

	output = util.Indent(output, "    ")

	if purge {
		currentNode, err := node.LoadNode(baseDir, gid)
		if err != nil {
			return output, err
		}

		output = output + fmt.Sprintf("\nThe directory %s has not been removed. It may contain private keys that protect sensitive information/funds. Please remove it manually if it is not needed anymore.", currentNode.SecretsDirectory())
	}

	return output, err
}

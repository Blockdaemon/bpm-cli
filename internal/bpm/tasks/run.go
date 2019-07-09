package tasks

import (
	"io/ioutil"
	"os"
	"path"
	"fmt"

	homedir "github.com/mitchellh/go-homedir"
	"gitlab.com/Blockdaemon/bpm/internal/bpm/plugin"
)

// Run contains functionality for the `run` cmd
//
// This has been seperated out into a function to make it easily testable
func Run(apiKey, baseDir, pluginURL, pluginName, runnerVersion string) (string, error) {
	if err := plugin.DownloadVersionInfo(pluginURL, baseDir); err != nil {
		return "", err
	}

	bpmUpgradable, err := plugin.CheckRunnerUpgradable(baseDir, runnerVersion)
	if err != nil {
		return "", err
	}
	if bpmUpgradable {
		return TEXT_NEW_BPM_VERSION, nil
	}

	pluginToRun, err := plugin.LoadPlugin(baseDir, pluginURL, pluginName)
	if err != nil {
		return "", err
	}

	pluginUpgradable, err := pluginToRun.NeedsUpgrade()
	if err != nil {
		return "", err
	}
	if pluginUpgradable {
		return TEXT_NEW_PLUGIN_VERSION, nil
	}

	// TODO: Fetch the config based on the api key from the PBG
	gid := os.Getenv("MOCK_GID")
	if gid == "" {
		return "", fmt.Errorf("env variable `MOCK_GID` isn't set. This is just used temporarily until we get the token from the BPG")
	}

	mockNodeFile := os.Getenv("MOCK_NODE_FILE")
	if mockNodeFile == "" {
		return "", fmt.Errorf("env variable `MOCK_NODE_FILE` isn't set. This is just used temporarily until we get the token from the BPG")
	}

	content, err := ioutil.ReadFile(mockNodeFile)
	if err != nil {
		return "", err
	}

	expandedBaseDir, err := homedir.Expand(baseDir)
	if err != nil {
		return "", err
	}
	nodePath := path.Join(expandedBaseDir, "nodes", gid)

	if err := os.MkdirAll(nodePath, os.ModePerm); err != nil {
		return "", err
	}

	nodeConfigPath := path.Join(nodePath, "node.json")

	if err := ioutil.WriteFile(nodeConfigPath, content, 0644); err != nil {
		return "", err
	}

	return "", pluginToRun.RunPlugin(gid)
}

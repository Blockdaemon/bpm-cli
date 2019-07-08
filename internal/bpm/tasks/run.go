// This is WIP, please don't code review this file yet
package tasks

import (
	"io/ioutil"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"gitlab.com/Blockdaemon/bpm/internal/bpm/plugin"
)

// Run contains functionality for the `run` cmd
//
// This has been seperated out into a function to make it easily testable
func Run(apiKey, baseDir, pluginURL, pluginName, runnerVersion string) (string, error) {
	if err := plugin.DownloadVersionInfo(apiKey, pluginURL, baseDir); err != nil {
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
	mockGID := "12345"
	mockData := `
	{
		"node_gid": "` + mockGID + `",
		"blockchain_gid": "6789gid",
		"environment": "mainnet",
		"network_type": "public",
		"node_subtype": "watcher",
		"protocol_type": "stellar-horizon",
		"config": {
			"core": {
				"full_history": true,
				"nodes": [
					{
						"id": "sdf1",
						"publicKey": "GCGB2S2KGYARPVIA37HYZXVRM2YZUEXA6S33ZU5BUDC6THSB62LZSTYH",
						"host": "core-live-a.stellar.org:11625"
					},
					{
						"id": "sdf2",
						"publicKey": "GCM6QMP3DLRPTAZW2UZPCPX2LF3SXWXKPMP3GKFZBDSF3QZGV2G5QSTK",
						"host": "core-live-b.stellar.org:11625"
					},
					{
						"id": "satoshipay-de",
						"publicKey": "GC5SXLNAM3C4NMGK2PXK4R34B5GNZ47FYQ24ZIBFDFOCU6D4KBN4POAE",
						"host": "stellar-de-fra.satoshipay.io:11625"
					},
					{
						"id": "satoshipay-sg",
						"publicKey": "GBJQUIXUO4XSNPAUT6ODLZUJRV2NPXYASKUBY4G5MYP3M47PCVI55MNT",
						"host": "stellar-sg-sin.satoshipay.io:11625"
					},
					{
						"id": "ibm-uk",
						"publicKey": "GAENPO2XRTTMAJXDWM3E3GAALNLG4HVMKJ4QF525TR25RI42YPEDULOW",
						"host": "uk.stellar.ibm.com:11625"
					}
				],
				"quorum_set_ids": [
					"sdf1",
					"sdf2",
					"satoshipay-de",
					"satoshipay-sg",
					"ibm-uk"
				]
			}
		}
	}`

	expandedBaseDir, err := homedir.Expand(baseDir)
	if err != nil {
		return "", err
	}
	nodePath := path.Join(expandedBaseDir, "nodes", mockGID)

	if err := os.MkdirAll(nodePath, os.ModePerm); err != nil {
		return "", err
	}

	nodeConfigPath := path.Join(nodePath, "node.json")

	if err := ioutil.WriteFile(nodeConfigPath, []byte(mockData), 0644); err != nil {
		return "", err
	}

	return "", pluginToRun.RunPlugin(mockGID)
}

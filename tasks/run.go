// This is WIP, please don't code review this file yet
package tasks

import (
	"io/ioutil"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"gitlab.com/Blockdaemon/runner/models"
)

// Run contains functionality for the `run` cmd
//
// This has been seperated out into a function to make it easily testable
func Run(apiKey, baseDir, pluginURL, pluginName string) (string, error) {
	versionInfoExists, err := models.CheckVersionInfoExists(baseDir)
	if err != nil {
		return "", err
	}

	if !versionInfoExists {
		return VERSION_INFO_MISSING, nil
	}

	// TODO: Auto upgrade at this point?

	// TODO: Fetch the config based on the api key

	mockData := `
	{
		"node_gid": "12345gid",
		"blockchain_gid": "6789gid",
		"containers": [
			{
				"name": "stellar-horizon",
				"image": "blockdaemon/docker-stellar-horizon",
				"version": "0.17.5-1"
			},
			{
				"name": "stellar-core",
				"image": "blockdaemon/docker-stellar-core",
				"version": "11.0.0-1"
			},
			{
				"name": "postgres",
				"image": "blockdaemon/docker-stellar-core",
				"version": "11.1"
			},
			{
				"name": "nodestate",
				"image": "blockdaemon/nodestate",
				"version": "1.2.0"
			}
		],
		"environment": "mainnet",
		"network_type": "public",
		"node_subtype": "watcher",
		"protocol_type": "stellar-horizon",
		"config": {
			"core": {
				"catchup_complete": false,
				"catchup_recent": 0,
				"failure_safety": -1,
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
				],
				"testnet": false,
				"unsafe_quorum": false
			},
			"horizon": {
				"history_retention_count": 0
			}
		}
	}`

	expandedBaseDir, err := homedir.Expand(baseDir)
	if err != nil {
		return "", err
	}
	nodePath := path.Join(expandedBaseDir, "nodes", "12345")

	if err := os.MkdirAll(nodePath, os.ModePerm); err != nil {
		return "", err
	}

	nodeConfigPath := path.Join(nodePath, "config.json")

	if err := ioutil.WriteFile(nodeConfigPath, []byte(mockData), 0644); err != nil {
		return "", err
	}

	plugin, err := models.LoadPlugin(baseDir, pluginURL, pluginName)
	if err != nil {
		return "", err
	}

	return "", plugin.RunPlugin()
}

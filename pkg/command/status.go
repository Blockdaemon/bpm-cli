package command

import (
	"strconv"
	"os"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/kataras/tablewriter"
)

// Status returns the status of a particular node
func (p *CmdContext) Status() error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetHeader([]string{
		"NODE ID",
		"PACKAGE",
		"STATUS",
		"SECRETS",
	})

	// List files in config directory
	nodeDirs, err := config.ReadDirs(config.NodesDir(p.HomeDir))
	if err != nil {
		return err
	}

	for _, nodeDir := range nodeDirs {
		nodeID := nodeDir.Name()

		n, err := node.Load(config.NodeFile(p.HomeDir, nodeID))
		if err != nil {
			return err
		}

		status := "unknown (package not installed)"
		if p.isInstalled(n.PluginName) {
			status, err = p.execCmdCapture(n, "status")
			if err != nil {
				return err
			}
		}

		secrets := strconv.Itoa(len(n.Secrets))

		table.Append([]string{
			nodeID,
			status,
			n.PluginName,
			secrets,
		})
	}

	table.Render()
	return nil
}

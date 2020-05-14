package command

import (
	"fmt"
	"os"

	"github.com/kataras/tablewriter"
	"go.blockdaemon.com/bpm/cli/pkg/config"
	"go.blockdaemon.com/bpm/sdk/pkg/node"
)

// Status returns the status of a particular node
func (p *CmdContext) Status() error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetHeader([]string{
		"NODE NAME",
		"PACKAGE",
		"STATUS",
	})

	// List files in config directory
	nodeDirs, err := config.ReadDirs(config.NodesDir(p.HomeDir))
	if err != nil {
		return err
	}

	for _, nodeDir := range nodeDirs {
		nodeName := nodeDir.Name()

		n, err := node.Load(config.NodeFile(p.HomeDir, nodeName))
		if err != nil {
			return err
		}

		status := "unknown (package not installed)"
		if p.isInstalled(n.PluginName) {
			status, err = p.execCmdCapture(n, "status")
			if err != nil {
				status = fmt.Sprintf("error: %s", err)
			}
		}

		table.Append([]string{
			nodeName,
			n.PluginName,
			status,
		})
	}

	table.Render()
	return nil
}

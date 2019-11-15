package plugin

import (
	"bytes"
	"strconv"

	"github.com/kataras/tablewriter"
	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm-sdk/pkg/node"
)

// Status returns the status of a particular node
func (p *PluginCmdContext) Status() (string, error) {
	var buf bytes.Buffer

	table := tablewriter.NewWriter(&buf)
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
		return "", err
	}

	for _, nodeDir := range nodeDirs {
		nodeID := nodeDir.Name()

		n, err := node.Load(config.NodesDir(p.HomeDir), nodeID)
		if err != nil {
			return "", err
		}

		status, err := p.execNodeCommand(n, "status")
		if err != nil {
			return "", err
		}

		secrets := strconv.Itoa(len(n.Secrets))

		table.Append([]string{
			nodeID,
			status,
			n.Protocol, // TODO: Wrong name
			secrets,
		})
	}

	table.Render()
	return buf.String(), nil
}

package plugin

import (
	"fmt"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/Blockdaemon/bpm/pkg/config"
)

func (p *PluginCmdContext) Stop(nodeID string) (string, error) {
	n, err := node.Load(config.NodesDir(p.HomeDir), nodeID)
	if err != nil {
		return "", err
	}

	output, err := p.execNodeCommand(n, "stop")
	if err != nil {
		return output, err
	}

	return fmt.Sprintf("%s\nThe node %q has been stopped.\n", output, nodeID), nil
}

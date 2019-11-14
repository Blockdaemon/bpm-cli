package plugin

import (
	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/Blockdaemon/bpm/pkg/config"
)

func (p *PluginCmdContext) Test(nodeID string) (string, error) {
	n, err := node.Load(config.NodesDir(p.HomeDir), nodeID)
	if err != nil {
		return "", err
	}

	return p.execNodeCommand(n, "test")
}

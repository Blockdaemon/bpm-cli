package command

import (
	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/Blockdaemon/bpm/pkg/config"
)

func (p *CmdContext) Test(nodeID string) error {
	n, err := node.Load(config.NodeFile(p.HomeDir, nodeID))
	if err != nil {
		return err
	}

	return p.execCmd(n, "test")
}

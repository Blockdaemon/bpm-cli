package command

import (
	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/Blockdaemon/bpm/pkg/config"
)

func (p *CmdContext) Test(nodeName string) error {
	n, err := node.Load(config.NodeFile(p.HomeDir, nodeName))
	if err != nil {
		return err
	}

	return p.execCmd(n, "test")
}

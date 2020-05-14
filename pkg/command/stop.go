package command

import (
	"fmt"

	"go.blockdaemon.com/bpm/cli/pkg/config"
	"go.blockdaemon.com/bpm/sdk/pkg/node"
)

func (p *CmdContext) Stop(nodeName string) error {
	n, err := node.Load(config.NodeFile(p.HomeDir, nodeName))
	if err != nil {
		return err
	}

	if err := p.execCmd(n, "stop"); err != nil {
		return err
	}

	fmt.Printf("The node %q has been stopped.\n", nodeName)
	return nil
}

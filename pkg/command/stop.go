package command

import (
	"fmt"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/Blockdaemon/bpm/pkg/config"
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

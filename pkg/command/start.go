package command

import (
	"fmt"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/Blockdaemon/bpm/pkg/config"
)

func (p *CmdContext) Start(nodeName string) error {
	n, err := node.Load(config.NodeFile(p.HomeDir, nodeName))
	if err != nil {
		return err
	}

	// Check if the plugin is the same version as used to configure the node
	packageVersion := p.getInstalledVersion(n.PluginName)
	if n.Version != packageVersion {
		return fmt.Errorf("cannot start node with currently installed package because it was installed with version %s. Have you considered upgrading it?", n.Version)
	}

	if err := p.execCmd(n, "start"); err != nil {
		return err
	}

	fmt.Printf("The node %q has been started.\n", nodeName)
	return nil
}

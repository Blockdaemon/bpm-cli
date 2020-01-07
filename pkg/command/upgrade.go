package command

import (
	"fmt"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/Blockdaemon/bpm/pkg/config"
)

func (p *CmdContext) Upgrade(nodeID string) error {
	n, err := node.Load(config.NodeFile(p.HomeDir, nodeID))
	if err != nil {
		return err
	}

	// Check if the plugin is the same version as used to configure the node
	packageVersion := p.getInstalledVersion(n.PluginName)
	if n.Version == packageVersion {
		return fmt.Errorf("package and node version are identical (%s). Have you considered upgrading the package?", n.Version)
	}

	if err := p.execCmd(n, "upgrade"); err != nil {
		return err
	}

	// Save new version in node.json
	n.Version = packageVersion
	if err := n.Save(); err != nil {
		return err
	}

	fmt.Printf("The node %q has been upgraded.\n", nodeID)
	return nil
}

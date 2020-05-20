package command

import (
	"fmt"

	"go.blockdaemon.com/bpm/cli/pkg/config"
	"go.blockdaemon.com/bpm/sdk/pkg/node"
	"go.blockdaemon.com/bpm/sdk/pkg/plugin"
)

func (p *CmdContext) Upgrade(nodeName string) error {
	n, err := node.Load(config.NodeFile(p.HomeDir, nodeName))
	if err != nil {
		return err
	}

	// Check if upgrades are supported
	meta, err := p.getMetaFromManifest(n.PluginName)
	if err != nil {
		return err
	}
	if !meta.Supports(plugin.SupportsUpgrade) {
		fmt.Printf("Package %q does not support upgrades. Skipping!\n", n.PluginName)
		return nil
	}

	// Check if the plugin is the same version as used to configure the node
	packageVersion := p.getInstalledVersion(n.PluginName)
	if n.Version == packageVersion {
		fmt.Printf("package and node version are identical (%s). Have you considered upgrading the package?\n", n.Version)
		return nil
	}

	if err := p.execCmd(n, "upgrade"); err != nil {
		return err
	}

	// Save new version in node.json
	n.Version = packageVersion
	if err := n.Save(); err != nil {
		return err
	}

	fmt.Printf("The node %q has been upgraded.\n", nodeName)
	return nil
}

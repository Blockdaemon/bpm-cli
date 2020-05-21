package command

import (
	"fmt"

	"go.blockdaemon.com/bpm/cli/pkg/config"
	"go.blockdaemon.com/bpm/sdk/pkg/node"
	"go.blockdaemon.com/bpm/sdk/pkg/plugin"
)

func (p *CmdContext) Test(nodeName string) error {
	n, err := node.Load(config.NodeFile(p.HomeDir, nodeName))
	if err != nil {
		return err
	}

	// Check if tests are supported
	meta, err := p.getMetaFromManifest(n.PluginName)
	if err != nil {
		return err
	}
	if !meta.Supports(plugin.SupportsTest) {
		fmt.Printf("Package %q does not support tests. Skipping!\n", n.PluginName)
		return nil
	}

	return p.execCmd(n, "test")
}

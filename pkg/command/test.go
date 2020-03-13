package command

import (
	"fmt"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/Blockdaemon/bpm-sdk/pkg/plugin"
	"github.com/Blockdaemon/bpm/pkg/config"
)

func (p *CmdContext) Test(nodeName string) error {
	n, err := node.Load(config.NodeFile(p.HomeDir, nodeName))
	if err != nil {
		return err
	}

	// Check if tests are supported
	meta, err := p.getMeta(n.PluginName)
	if err != nil {
		return err
	}
	if !meta.Supports(plugin.SupportsTest) {
		fmt.Printf("Package %q does not support tests. Skipping!\n", n.PluginName)
		return nil
	}

	return p.execCmd(n, "test")
}

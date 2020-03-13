package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/Blockdaemon/bpm-sdk/pkg/plugin"
	bpmconfig "github.com/Blockdaemon/bpm/pkg/config"
)

// Remove removes the node or parts of it
func (p *CmdContext) Remove(nodeName string, all bool, data bool, config bool, runtime bool, identity bool) error {
	n, err := node.Load(bpmconfig.NodeFile(p.HomeDir, nodeName))
	if err != nil {
		return err
	}

	if config || all {
		if err := p.execCmd(n, "remove-config"); err != nil {
			return err
		}
	}

	if runtime || data || all {
		if err := p.execCmd(n, "remove-runtime"); err != nil {
			return err
		}
	}

	if data || all {
		if err := p.execCmd(n, "remove-data"); err != nil {
			return err
		}
	}

	if identity || all {
		meta, err := p.getMeta(n.PluginName)
		if err != nil {
			return err
		}
		if meta.Supports(plugin.SupportsIdentity) {
			if err := p.execCmd(n, "remove-identity"); err != nil {
				return err
			}
		} else {
			fmt.Printf("Package %q does not support managing identities. Skipping removal!\n", n.PluginName)
		}
	}

	if all {
		baseDir := bpmconfig.NodesDir(p.HomeDir)
		nodeDir := filepath.Join(baseDir, nodeName)
		fmt.Printf("\nRemoving directory %q\n", nodeDir)
		if err := os.RemoveAll(nodeDir); err != nil {
			return err
		}
	}

	return nil
}

package command

import (
	"fmt"

	bpmconfig "go.blockdaemon.com/bpm/cli/pkg/config"
	"go.blockdaemon.com/bpm/sdk/pkg/node"
	"go.blockdaemon.com/bpm/sdk/pkg/plugin"
)

// Remove removes the node or parts of it
func (p *CmdContext) Remove(nodeName string, all bool, data bool, config bool, runtime bool, identity bool) error {
	n, err := node.Load(bpmconfig.NodeFile(p.HomeDir, nodeName))
	if err != nil {
		return err
	}
	meta, err := p.getMetaFromManifest(n.PluginName)
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
		if meta.Supports(plugin.SupportsIdentity) {
			if err := p.execCmd(n, "remove-identity"); err != nil {
				return err
			}
		} else {
			fmt.Printf("Package %q does not support managing identities. Skipping removal!\n", n.PluginName)
		}
	}

	if all {
		// Tear down runtime environment
		if meta.ProtocolVersionGreaterEqualThan("1.2.0") {
			if err := p.execCmd(n, "tear-down-environment"); err != nil {
				return err
			}
		}

		fmt.Printf("\nRemoving node %q\n", nodeName)
		if err := n.Remove(); err != nil {
			return err
		}
	}

	return nil
}

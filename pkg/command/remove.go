package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
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
		// remove-identity has been introduced in protocol version 1.1.0
		meta, err := p.getMeta(n.PluginName)
		if err != nil {
			return err
		}
		if meta.ProtocolVersion == "1.0.0" {
			fmt.Println("You are using an outdated package which doesn't support `remove-identity`. Skipping!")
		} else {
			if err := p.execCmd(n, "remove-identity"); err != nil {
				return err
			}
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

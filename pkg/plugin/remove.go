package plugin

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	bpmconfig "github.com/Blockdaemon/bpm/pkg/config"
)

func (p *PluginCmdContext) Remove(nodeID string, all bool, data bool, config bool) error {
	n, err := node.Load(bpmconfig.NodesDir(p.HomeDir), nodeID)
	if err != nil {
		return err
	}

	if config || all {
		if err := p.execCmd(n, "remove-config"); err != nil {
			return err
		}
	}

	if data || all {
		if err := p.execCmd(n, "stop"); err != nil {
			return err
		}

		if err := p.execCmd(n, "remove-data"); err != nil {
			return err
		}
	}

	if all {
		if err := p.execCmd(n, "remove-node"); err != nil {
			return err
		}

		baseDir := bpmconfig.NodesDir(p.HomeDir)
		nodeDir := filepath.Join(baseDir, nodeID)
		fmt.Printf("\nRemoving directory %q\n", nodeDir)
		if err := os.RemoveAll(nodeDir); err != nil {
			return err
		}
	}

	return nil
}

package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	bpmconfig "github.com/Blockdaemon/bpm/pkg/config"
)

func (p *CmdContext) Remove(nodeID string, all bool, data bool, config bool, runtime bool) error {
	n, err := node.Load(bpmconfig.NodeFile(p.HomeDir, nodeID))
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

	if all {
		baseDir := bpmconfig.NodesDir(p.HomeDir)
		nodeDir := filepath.Join(baseDir, nodeID)
		fmt.Printf("\nRemoving directory %q\n", nodeDir)
		if err := os.RemoveAll(nodeDir); err != nil {
			return err
		}
	}

	return nil
}

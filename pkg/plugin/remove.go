package plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"bytes"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	bpmconfig "github.com/Blockdaemon/bpm/pkg/config"
)

func (p *PluginCmdContext) Remove(nodeID string, all bool, data bool, config bool) (string, error) {
	n, err := node.Load(bpmconfig.NodesDir(p.HomeDir), nodeID)
	if err != nil {
		return "", err
	}

	fullOutput := bytes.NewBufferString("")

	if config || all {
		output, err := p.execNodeCommand(n, "remove-config")
		fullOutput.WriteString("\n" + output)

		if err != nil {
			return fullOutput.String(), err
		}
	}

	if data || all {
		output, err := p.execNodeCommand(n, "stop")
		fullOutput.WriteString("\n" + output)

		if err != nil {
			return fullOutput.String(), err
		}

		output, err = p.execNodeCommand(n, "remove-data")
		fullOutput.WriteString("\n" + output)

		if err != nil {
			return fullOutput.String(), err
		}
	}

	if all {
		output, err := p.execNodeCommand(n, "remove-node")
		fullOutput.WriteString("\n" + output)

		if err != nil {
			return fullOutput.String(), err
		}

		baseDir := bpmconfig.NodesDir(p.HomeDir)
		nodeDir := filepath.Join(baseDir, nodeID)
		fullOutput.WriteString(fmt.Sprintf("\nRemoving directory %q\n", nodeDir))
		os.RemoveAll(nodeDir)
	}

	return fullOutput.String(), nil
}

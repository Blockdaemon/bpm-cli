package plugin

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/Blockdaemon/bpm/pkg/config"
)

func (p *PluginCmdContext) ShowConfig(nodeID string) error {
	// Check if node exists
	if !config.FileExists(
		filepath.Join(config.NodesDir(p.HomeDir), nodeID),
		"node.json",
	) {
		return fmt.Errorf("Node %q does not exist\n", nodeID)
	}

	// Get the node
	n, err := node.Load(config.NodesDir(p.HomeDir), nodeID)
	if err != nil {
		return err
	}

	var buf bytes.Buffer

	// List files in config directory
	if err := config.Walk(n.ConfigsDirectory(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			writeHeader(path, &buf)

			// Filename is empty because the path contains the file
			data, err := config.Read(path, "")
			if err != nil {
				return nil
			}

			buf.Write(data)
		}

		return nil
	}); err != nil {
		return err
	}

	fmt.Println(buf.String())

	return nil
}

func (p *PluginCmdContext) ShowNode(nodeID string) error {
	// Check if node exists
	if !config.FileExists(
		filepath.Join(config.NodesDir(p.HomeDir), nodeID),
		"node.json",
	) {
		return fmt.Errorf("Node %q does not exist\n", nodeID)
	}

	// Get the node
	n, err := node.Load(config.NodesDir(p.HomeDir), nodeID)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	writeHeader(n.NodeFile(), &buf)

	data, err := config.Read(n.NodeFile(), "")
	if err != nil {
		return err
	}

	buf.Write(data)
	fmt.Println(buf.String())

	return nil
}

func writeHeader(path string, buf *bytes.Buffer) {
	buf.WriteString("\n")
	buf.WriteString("--- ")
	buf.WriteString(path)
	buf.WriteString(":\n\n")
}

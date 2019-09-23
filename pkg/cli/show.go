package cli

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/bpm/pkg/config"
)

func newShowCmd(c *command) *cobra.Command {
	showCmd := &cobra.Command{
		Use:   "show <resource>",
		Short: "Print a resource to stdout",
	}

	showConfigCmd := &cobra.Command{
		Use:   "config <id>",
		Short: "Display config files for a node",
		Args:  cobra.MinimumNArgs(1),
		RunE: c.Wrap(func(homeDir string, _ config.Manifest, args []string) error {
			id := args[0]

			// Check if node exists
			if !config.FileExists(
				filepath.Join(config.NodesDir(homeDir), id),
				"node.json",
			) {
				fmt.Printf("Node %q does not exist\n", id)
				return nil
			}

			// Get the node
			n, err := node.Load(config.NodesDir(homeDir), id)
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

		}),
	}

	showNodeCmd := &cobra.Command{
		Use:   "node <id>",
		Short: "Display the node.json config",
		Args:  cobra.MinimumNArgs(1),
		RunE: c.Wrap(func(homeDir string, _ config.Manifest, args []string) error {
			id := args[0]

			// Check if node exists
			if !config.FileExists(
				filepath.Join(config.NodesDir(homeDir), id),
				"node.json",
			) {
				fmt.Printf("Node %q does not exist\n", id)
				return nil
			}

			// Get the node
			n, err := node.Load(config.NodesDir(homeDir), id)
			if err != nil {
				fmt.Println(err)
				return err
			}

			var buf bytes.Buffer
			writeHeader(n.NodeFile(), &buf)

			data, err := config.Read(n.NodeFile(), "")
			if err != nil {
				return nil
			}

			buf.Write(data)
			fmt.Println(buf.String())

			return nil
		}),
	}

	showCmd.AddCommand(showConfigCmd)
	showCmd.AddCommand(showNodeCmd)

	return showCmd
}

func writeHeader(path string, buf *bytes.Buffer) {
	buf.WriteString("\n")
	buf.WriteString("--- ")
	buf.WriteString(path)
	buf.WriteString(":\n\n")
}

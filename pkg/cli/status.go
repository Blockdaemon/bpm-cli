package cli

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/kataras/tablewriter"
	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"gitlab.com/Blockdaemon/bpm/pkg/plugin"
)

func newStatusCmd(c *command) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Display statuses of configured nodes",
		RunE: c.Wrap(func(homeDir string, _ config.Manifest, args []string) error {
			var buf bytes.Buffer

			table := tablewriter.NewWriter(&buf)
			table.SetBorder(false)
			table.SetHeader([]string{
				"NODE ID",
				"PLUGIN",
				"STATUS",
				"SECRETS",
			})

			// List files in config directory
			nodeDirs, err := config.ReadDirs(config.NodesDir(homeDir))
			if err != nil {
				return err
			}

			for _, nodeDir := range nodeDirs {
				nodeID := nodeDir.Name()
				n, err := node.Load(config.NodesDir(homeDir), nodeID)
				if err != nil {
					return err
				}

				pluginName := n.Protocol

				status, err := plugin.Status(homeDir, pluginName, nodeID, c.debug)
				if err != nil {
					return err
				}

				secrets := strconv.Itoa(len(n.Secrets))

				table.Append([]string{
					nodeID,
					status,
					pluginName,
					secrets,
				})
			}

			table.Render()
			fmt.Println(buf.String())

			return nil

		}),
	}
}

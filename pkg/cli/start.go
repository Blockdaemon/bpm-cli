package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"gitlab.com/Blockdaemon/bpm/pkg/plugin"
)

func newStartCmd(c *command) *cobra.Command {
	return &cobra.Command{
		Use:   "start <id>",
		Short: "Start a blockchain node",
		Args:  cobra.MinimumNArgs(1),
		RunE: c.Wrap(func(homeDir string, m config.Manifest, args []string) error {
			id := args[0]

			n, err := node.Load(config.NodesDir(homeDir), id)
			if err != nil {
				return err
			}
			pluginName := n.Protocol

			// Check if plugin is installed
			if _, ok := m.Plugins[pluginName]; !ok {
				fmt.Printf("The package %q is currently not installed.\n", pluginName)
				return nil
			}

			// Run the plugin
			if err := plugin.Start(homeDir, pluginName, id, c.debug); err != nil {
				return err
			}

			fmt.Printf("The node %q has been started.\n", id)

			return nil
		}),
	}
}

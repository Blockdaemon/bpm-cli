package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"gitlab.com/Blockdaemon/bpm/pkg/plugin"
)

func newStopCmd(c *command) *cobra.Command {
	var purge bool

	cmd := &cobra.Command{
		Use:   "stop <id>",
		Short: "Stops a running blockchain node",
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

			// Remove plugin
			if err := plugin.Stop(homeDir, pluginName, id, purge, c.debug); err != nil {
				return err
			}

			fmt.Printf("The node %q has been stopped.\n", id)

			return nil
		}),
	}

	cmd.Flags().BoolVar(&purge, "purge", false, "Purge all data and configuration files. Secrets (e.g. private keys) will not be removed because they may protect sensitive information/funds.")

	return cmd
}

package cli

import (
	"fmt"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/spf13/cobra"
	bpmconfig "github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm/pkg/plugin"
)

func newRemoveCmd(c *command) *cobra.Command {
	var (
		all    bool
		data   bool
		config bool
	)

	cmd := &cobra.Command{
		Use:   "remove <id>",
		Short: "Remove blockchain node data and configuration",
		Args:  cobra.MinimumNArgs(1),
		RunE: c.Wrap(func(homeDir string, m bpmconfig.Manifest, args []string) error {
			id := args[0]

			if !(all || data || config) {
				return fmt.Errorf("flag missing to specify what to remove. Use `--help` for details!")
			}

			n, err := node.Load(bpmconfig.NodesDir(homeDir), id)
			if err != nil {
				return err
			}
			pluginName := n.Protocol

			// Check if plugin is installed
			if _, ok := m.Plugins[pluginName]; !ok {
				fmt.Printf("The package %q is currently not installed.\n", pluginName)
				return nil
			}

			if err := plugin.Remove(homeDir, pluginName, id, c.debug, all, data, config); err != nil {
				return err
			}

			return nil
		}),
	}

	cmd.Flags().BoolVar(&all, "all", false, "Remove all data, configuration files and node information")
	cmd.Flags().BoolVar(&config, "config", false, "Remove all configuration files but keep data and node information")
	cmd.Flags().BoolVar(&data, "data", false, "Remove all data but keep configuration files and node information")

	return cmd

}

package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	bpmconfig "github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm/pkg/plugin"
)

func newRemoveCmd(c *command, runtimeOS string) *cobra.Command {
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

			// TODO: Why do we have three ways of passing down variables?
			cmdContext := plugin.PluginCmdContext{
				HomeDir: homeDir,
				Manifest: m,
				RuntimeOS: runtimeOS,
				RegistryURL: c.registry,
				Debug: c.debug,
			}

			output, err := cmdContext.Remove(id, all, data, config)
			fmt.Println(output)
			return err
		}),
	}

	cmd.Flags().BoolVar(&all, "all", false, "Remove all data, configuration files and node information")
	cmd.Flags().BoolVar(&config, "config", false, "Remove all configuration files but keep data and node information")
	cmd.Flags().BoolVar(&data, "data", false, "Remove all data but keep configuration files and node information")

	return cmd

}

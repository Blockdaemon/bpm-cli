package cli

import (
	"fmt"

	"github.com/Blockdaemon/bpm/pkg/plugin"
	"github.com/spf13/cobra"
	"github.com/Blockdaemon/bpm/pkg/config"
)

func newStartCmd(c *command, runtimeOS string) *cobra.Command {
	return &cobra.Command{
		Use:   "start <id>",
		Short: "Start a blockchain node",
		Args:  cobra.MinimumNArgs(1),
		RunE: c.Wrap(func(homeDir string, m config.Manifest, args []string) error {
			id := args[0]

			// TODO: Why do we have three ways of passing down variables?
			cmdContext := plugin.PluginCmdContext{
				HomeDir: homeDir,
				Manifest: m,
				RuntimeOS: runtimeOS,
				RegistryURL: c.registry,
				Debug: c.debug,
			}

			output, err := cmdContext.Start(id)
			fmt.Println(output)
			if err != nil {
				return err
			}
			fmt.Printf("The node %q has been started.\n", id)
			return nil
		}),
	}
}

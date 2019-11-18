package cli

import (
	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm/pkg/plugin"
	"github.com/spf13/cobra"
)

func newTestCmd(c *command, runtimeOS string) *cobra.Command {
	return &cobra.Command{
		Use:   "test <id>",
		Short: "Tests a running blockchain node",
		Args:  cobra.MinimumNArgs(1),
		RunE: c.Wrap(func(homeDir string, m config.Manifest, args []string) error {
			id := args[0]

			// TODO: Why do we have three ways of passing down variables?
			cmdContext := plugin.PluginCmdContext{
				HomeDir:     homeDir,
				Manifest:    m,
				RuntimeOS:   runtimeOS,
				RegistryURL: c.registry,
				Debug:       c.debug,
			}

			return cmdContext.Test(id)
		}),
	}
}

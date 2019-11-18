package cli

import (
	"github.com/Blockdaemon/bpm/pkg/plugin"
	"github.com/spf13/cobra"
	"github.com/Blockdaemon/bpm/pkg/config"
)

func newStopCmd(c *command, runtimeOS string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop <id>",
		Short: "Stops a running blockchain node",
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

			return cmdContext.Stop(id)
		}),
	}

	return cmd
}

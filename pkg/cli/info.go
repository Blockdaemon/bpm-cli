package cli

import (
	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm/pkg/plugin"
	"github.com/spf13/cobra"
)

func newInfoCmd(c *command, os string) *cobra.Command {
	return &cobra.Command{
		Use:   "info <package>",
		Short: "Show information about a package",
		Args:  cobra.MinimumNArgs(1),
		RunE: c.Wrap(func(homeDir string, m config.Manifest, args []string) error {
			pluginName := args[0]

			// TODO: Why do we have three ways of passing down variables?
			cmdContext := plugin.PluginCmdContext{
				HomeDir:     homeDir,
				Manifest:    m,
				RuntimeOS:   os,
				RegistryURL: c.registry,
				Debug:       c.debug,
			}

			return cmdContext.Info(pluginName)
		}),
	}
}

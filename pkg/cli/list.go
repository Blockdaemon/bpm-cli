package cli

import (
	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm/pkg/plugin"
	"github.com/spf13/cobra"
)

func newListCmd(c *command, os string) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List installed packages",
		RunE: c.Wrap(func(homeDir string, m config.Manifest, args []string) error {
			// TODO: Why do we have three ways of passing down variables?
			cmdContext := plugin.PluginCmdContext{
				HomeDir:     homeDir,
				Manifest:    m,
				RuntimeOS:   os,
				RegistryURL: c.registry,
				Debug:       c.debug,
			}

			return cmdContext.List()
		}),
	}
}

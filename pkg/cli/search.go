package cli

import (
	"github.com/spf13/cobra"
	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm/pkg/plugin"
)

func newSearchCmd(c *command, os string) *cobra.Command {
	return &cobra.Command{
		Use:   "search <package>",
		Short: "Search available packages",
		RunE: c.Wrap(func(homeDir string, m config.Manifest, args []string) error {
			cmdContext := plugin.PluginCmdContext{
				HomeDir: homeDir,
				Manifest: m,
				RuntimeOS: os,
				RegistryURL: c.registry,
				Debug: c.debug,
			}

			query := ""
			if len(args) > 0 {
				query = args[0]
			}

			return cmdContext.Search(query)
		}),
	}
}

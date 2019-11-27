package cli

import (
	"github.com/Blockdaemon/bpm/pkg/plugin"
	"github.com/spf13/cobra"
)

func newSearchCmd(cmdContext plugin.PluginCmdContext) *cobra.Command {
	return &cobra.Command{
		Use:   "search [package]",
		Short: "Search available packages",
		RunE: func(cmd *cobra.Command, args []string) error {
			query := ""
			if len(args) > 0 {
				query = args[0]
			}

			return cmdContext.Search(query)
		},
	}
}

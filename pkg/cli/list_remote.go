package cli

import (
	"github.com/Blockdaemon/bpm/pkg/command"
	"github.com/spf13/cobra"
)

func newSearchCmd(cmdContext command.CmdContext) *cobra.Command {
	return &cobra.Command{
		Use:     "list-remote [package]",
		Short:   "Search available packages",
		Aliases: []string{"search"},
		RunE: func(cmd *cobra.Command, args []string) error {
			query := ""
			if len(args) > 0 {
				query = args[0]
			}

			return cmdContext.Search(query)
		},
	}
}

package cli

import (
	"github.com/spf13/cobra"
	"go.blockdaemon.com/bpm/cli/pkg/command"
)

func newListCmd(cmdContext command.CmdContext) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List installed packages",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmdContext.List()
		},
	}
}

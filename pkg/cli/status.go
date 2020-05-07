package cli

import (
	"github.com/spf13/cobra"
	"go.blockdaemon.com/bpm/cli/pkg/command"
)

func newStatusCmd(cmdContext command.CmdContext) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Display statuses of configured nodes",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmdContext.Status()
		},
	}
}

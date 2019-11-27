package cli

import (
	"github.com/Blockdaemon/bpm/pkg/command"
	"github.com/spf13/cobra"
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

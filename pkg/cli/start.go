package cli

import (
	"github.com/Blockdaemon/bpm/pkg/command"
	"github.com/spf13/cobra"
)

func newStartCmd(cmdContext command.CmdContext) *cobra.Command {
	return &cobra.Command{
		Use:   "start <id>",
		Short: "Start a blockchain node",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			return cmdContext.Start(id)
		},
	}
}

package cli

import (
	"github.com/spf13/cobra"
	"go.blockdaemon.com/bpm/cli/pkg/command"
)

func newTestCmd(cmdContext command.CmdContext) *cobra.Command {
	return &cobra.Command{
		Use:   "test <name>",
		Short: "Tests a running blockchain node",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			return cmdContext.Test(name)
		},
	}
}

package cli

import (
	"github.com/spf13/cobra"
	"go.blockdaemon.com/bpm/cli/pkg/command"
)

func newStopCmd(cmdContext command.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop <name>",
		Short: "Stops a running blockchain node",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			return cmdContext.Stop(name)
		},
	}

	return cmd
}

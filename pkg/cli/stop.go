package cli

import (
	"github.com/Blockdaemon/bpm/pkg/command"
	"github.com/spf13/cobra"
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

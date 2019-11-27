package cli

import (
	"github.com/Blockdaemon/bpm/pkg/plugin"
	"github.com/spf13/cobra"
)

func newStopCmd(cmdContext plugin.PluginCmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop <id>",
		Short: "Stops a running blockchain node",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			return cmdContext.Stop(id)
		},
	}

	return cmd
}

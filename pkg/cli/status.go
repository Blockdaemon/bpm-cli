package cli

import (
	"github.com/Blockdaemon/bpm/pkg/plugin"
	"github.com/spf13/cobra"
)

func newStatusCmd(cmdContext plugin.PluginCmdContext) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Display statuses of configured nodes",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmdContext.Status()
		},
	}
}

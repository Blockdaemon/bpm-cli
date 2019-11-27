package cli

import (
	"github.com/Blockdaemon/bpm/pkg/plugin"
	"github.com/spf13/cobra"
)

func newListCmd(cmdContext plugin.PluginCmdContext) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List installed packages",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmdContext.List()
		},
	}
}

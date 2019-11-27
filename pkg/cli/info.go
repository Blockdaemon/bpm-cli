package cli

import (
	"github.com/Blockdaemon/bpm/pkg/plugin"
	"github.com/spf13/cobra"
)

func newInfoCmd(cmdContext plugin.PluginCmdContext) *cobra.Command {
	return &cobra.Command{
		Use:   "info <package>",
		Short: "Show information about a package",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			pluginName := args[0]
			return cmdContext.Info(pluginName)
		},
	}
}

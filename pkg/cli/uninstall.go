package cli

import (
	"github.com/spf13/cobra"
	"go.blockdaemon.com/bpm/cli/pkg/command"
)

func newUninstallCmd(cmdContext command.CmdContext) *cobra.Command {
	return &cobra.Command{
		Use:   "uninstall",
		Short: "Uninstall a package. Existing nodes will not be removed.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			pluginName := args[0]

			return cmdContext.Uninstall(pluginName)
		},
	}
}

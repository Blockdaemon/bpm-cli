package cli

import (
	"github.com/spf13/cobra"
	"go.blockdaemon.com/bpm/cli/pkg/command"
)

func newUpgradeCmd(cmdContext command.CmdContext) *cobra.Command {
	return &cobra.Command{
		Use:   "upgrade <id>",
		Short: "Upgrade a blockchain node to the current version of the package",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			return cmdContext.Upgrade(name)
		},
	}
}

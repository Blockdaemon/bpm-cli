package cli

import (
	"github.com/Blockdaemon/bpm/pkg/command"
	"github.com/spf13/cobra"
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

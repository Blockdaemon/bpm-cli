package cli

import (
	"github.com/Blockdaemon/bpm/pkg/command"
	"github.com/spf13/cobra"
)

func newShowCmd(cmdContext command.CmdContext) *cobra.Command {
	showCmd := &cobra.Command{
		Use:   "show <resource>",
		Short: "Print a resource to stdout",
	}

	showConfigCmd := &cobra.Command{
		Use:   "config <name>",
		Short: "Display config files for a node",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			return cmdContext.ShowConfig(name)
		},
	}

	showNodeCmd := &cobra.Command{
		Use:   "node <name>",
		Short: "Display the node.json config",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			return cmdContext.ShowNode(name)
		},
	}

	showCmd.AddCommand(showConfigCmd)
	showCmd.AddCommand(showNodeCmd)

	return showCmd
}

package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.blockdaemon.com/bpm/cli/pkg/command"
)

func newRemoveCmd(cmdContext command.CmdContext) *cobra.Command {
	var (
		all      bool
		data     bool
		config   bool
		runtime  bool
		identity bool
	)

	cmd := &cobra.Command{
		Use:   "remove <name>",
		Short: "Remove blockchain node data and configuration. Select one of the required flags for the remove command.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			if !(all || data || config || runtime || identity) {
				return fmt.Errorf("flag missing to specify what to remove. Use `--help` for details")
			}

			return cmdContext.Remove(name, all, data, config, runtime, identity)
		},
	}

	cmd.Flags().BoolVar(&all, "all", false, "[Required] Remove all data, configuration files and node information. Linux only: To avoid file permission denied errors on Linux use 'sudo' with this command")
	cmd.Flags().BoolVar(&config, "config", false, "[Required] Remove all configuration files but keep data and node information")
	cmd.Flags().BoolVar(&data, "data", false, "[Required] Remove all data but keep configuration files and node information. Linux only: To avoid file permission denied errors on Linux use 'sudo' with this command")
	cmd.Flags().BoolVar(&runtime, "runtime", false, "[Required] Remove all runtimes but keep configuration files and node information")
	cmd.Flags().BoolVar(&identity, "identity", false, "[Required] Remove the identity of the node")

	return cmd

}

package cli

import (
	"fmt"

	"github.com/Blockdaemon/bpm/pkg/command"
	"github.com/spf13/cobra"
)

func newRemoveCmd(cmdContext command.CmdContext) *cobra.Command {
	var (
		all     bool
		data    bool
		config  bool
		runtime bool
	)

	cmd := &cobra.Command{
		Use:   "remove <name>",
		Short: "Remove blockchain node data and configuration. Select one of the required flags for the remove command.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			if !(all || data || config || runtime) {
				return fmt.Errorf("flag missing to specify what to remove. Use `--help` for details")
			}

			return cmdContext.Remove(name, all, data, config, runtime)
		},
	}

	cmd.Flags().BoolVar(&all, "all", false, "[Required] Remove all data, configuration files and node information")
	cmd.Flags().BoolVar(&config, "config", false, "[Required] Remove all configuration files but keep data and node information")
	cmd.Flags().BoolVar(&data, "data", false, "[Required] Remove all data but keep configuration files and node information")
	cmd.Flags().BoolVar(&runtime, "runtime", false, "[Required] Remove all runtimes but keep configuration files and node information")

	return cmd

}

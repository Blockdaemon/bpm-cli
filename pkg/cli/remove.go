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
		Use:   "remove <id>",
		Short: "Remove blockchain node data and configuration",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]

			if !(all || data || config || runtime) {
				return fmt.Errorf("flag missing to specify what to remove. Use `--help` for details!")
			}

			return cmdContext.Remove(id, all, data, config, runtime)
		},
	}

	cmd.Flags().BoolVar(&all, "all", false, "Remove all data, configuration files and node information")
	cmd.Flags().BoolVar(&config, "config", false, "Remove all configuration files but keep data and node information")
	cmd.Flags().BoolVar(&data, "data", false, "Remove all data but keep configuration files and node information")
	cmd.Flags().BoolVar(&runtime, "runtime", false, "Remove all runtimes but keep configuration files and node information")

	return cmd

}

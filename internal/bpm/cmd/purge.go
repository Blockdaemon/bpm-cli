package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/bpm/internal/bpm/tasks"
)

var purgeCmd = &cobra.Command{
	Use:   "purge <plugin>",
	Short: "Purge the configuration and data of a blockchain client",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pluginName := args[0]

		output, err := tasks.Purge(baseDir, pluginURL, pluginName, runnerVersion)
		if err != nil {
			return err
		}

		fmt.Println(output)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)
}

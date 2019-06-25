package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/bpm/internal/bpm/tasks"
)

var runCmd = &cobra.Command{
	Use:   "run <plugin>",
	Short: "Run an installed plugin",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pluginName := args[0]

		output, err := tasks.Run(apiKey, baseDir, pluginURL, pluginName, runnerVersion)
		if err != nil {
			return err
		}

		fmt.Println(output)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	addAPIKeyFlag(runCmd)
}

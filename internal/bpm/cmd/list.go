package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/bpm/internal/bpm/tasks"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available and installed blockchain protocols",
	RunE: func(cmd *cobra.Command, args []string) error {
		output, err := tasks.List(apiKey, baseDir, pluginURL)
		if err != nil {
			return err
		}

		fmt.Println(output)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

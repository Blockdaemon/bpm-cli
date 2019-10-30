package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"gitlab.com/Blockdaemon/bpm/pkg/plugin"
)

func newInfoCmd(c *command, os string) *cobra.Command {
	return &cobra.Command{
		Use:   "info <package>",
		Short: "Show information about a package",
		Args:  cobra.MinimumNArgs(1),
		RunE: c.Wrap(func(homeDir string, m config.Manifest, args []string) error {

			query := ""
			if len(args) > 0 {
				query = strings.ToLower(args[0])
			}

			output, err := plugin.Info(c.registry, query, os, m, homeDir, c.debug)
			if err != nil {
				return err
			}

			fmt.Println(output)

			return nil
		}),
	}
}

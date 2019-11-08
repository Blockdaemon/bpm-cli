package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm/pkg/plugin"
)

func newSearchCmd(c *command, os string) *cobra.Command {
	return &cobra.Command{
		Use:   "search <package>",
		Short: "Search available packages",
		RunE: c.Wrap(func(homeDir string, m config.Manifest, args []string) error {
			query := ""
			if len(args) > 0 {
				query = strings.ToLower(args[0])
			}

			output, err := plugin.Search(c.registry, query, os, m)
			if err != nil {
				return err
			}

			fmt.Println(output)

			return nil
		}),
	}
}

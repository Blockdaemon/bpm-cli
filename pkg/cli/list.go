package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm/pkg/plugin"
)

func newListCmd(c *command, os string) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List installed packages",
		RunE: c.Wrap(func(homeDir string, m config.Manifest, args []string) error {
			output, err := plugin.List(c.registry, m, os)
			if err != nil {
				return err
			}

			fmt.Println(output)

			return nil
		}),
	}
}

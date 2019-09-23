package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"gitlab.com/Blockdaemon/bpm/pkg/plugin"
)

func newListCmd(c *command) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List available and installed blockchain protocols",
		RunE: c.Wrap(func(homeDir string, m config.Manifest, args []string) error {
			output, err := plugin.List(m)
			if err != nil {
				return err
			}

			fmt.Println(output)

			return nil
		}),
	}
}

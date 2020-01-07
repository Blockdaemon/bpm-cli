package cli

import (
	"fmt"
	"github.com/Blockdaemon/bpm/pkg/command"
	"github.com/spf13/cobra"
)

func newUninstallCmd(cmdContext command.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Uninstall a package. Data and configuration will not be removed.",
	}
	for name := range cmdContext.Manifest.Plugins {
		pluginCmd := &cobra.Command{
			Use:   name,
			Short: fmt.Sprintf("Uninstall the %q package.", name),
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmdContext.Uninstall(name)
			},
		}

		cmd.AddCommand(pluginCmd)
	}

	return cmd
}

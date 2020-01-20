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
		// copy to break closure
		nameCopy := name

		pluginCmd := &cobra.Command{
			Use:   nameCopy,
			Short: fmt.Sprintf("Uninstall the %q package.", nameCopy),
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmdContext.Uninstall(nameCopy)
			},
		}

		cmd.AddCommand(pluginCmd)
	}

	return cmd
}

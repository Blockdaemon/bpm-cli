package cli

import (
	"github.com/Blockdaemon/bpm/pkg/command"
	"github.com/spf13/cobra"
)

// newInstallCmd downloads and install a plugin from the PBR to the plugins directory
func newInstallCmd(cmdContext command.CmdContext) *cobra.Command {
	var fromFile string

	installCmd := &cobra.Command{
		Use:   "install <package> [version]",
		Short: "Installs or upgrades a package to a specific version or latest if no version is specified",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			pluginName := args[0]

			if len(fromFile) > 0 {
				return cmdContext.InstallFile(pluginName, fromFile)
			}

			if len(args) > 1 {
				versionToInstall := args[1]
				return cmdContext.Install(pluginName, versionToInstall)
			}

			return cmdContext.InstallLatest(pluginName)
		},
	}
	installCmd.Flags().StringVar(&fromFile, "from-file", "", "Install a package from a file instead of downloading it")
	return installCmd
}

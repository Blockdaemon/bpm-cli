package cli

import (
	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm/pkg/plugin"
	"github.com/spf13/cobra"
)

// newInstallCmd downloads and install a plugin from the PBR to the plugins directory
func newInstallCmd(c *command, os string) *cobra.Command {
	return &cobra.Command{
		Use:   "install <package> [version]",
		Short: "Installs or upgrades a package to a specific version or latest if no version is specified",
		Args:  cobra.MinimumNArgs(1),
		RunE: c.Wrap(func(homeDir string, m config.Manifest, args []string) error {
			pluginName := args[0]

			// TODO: Why do we have three ways of passing down variables?
			cmdContext := plugin.PluginCmdContext{
				HomeDir:     homeDir,
				Manifest:    m,
				RuntimeOS:   os,
				RegistryURL: c.registry,
				Debug:       c.debug,
			}

			if len(args) > 1 {
				versionToInstall := args[1]
				return cmdContext.Install(pluginName, versionToInstall)
			}

			return cmdContext.InstallLatest(pluginName)
		}),
	}
}

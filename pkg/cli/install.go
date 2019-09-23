package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"gitlab.com/Blockdaemon/bpm/pkg/plugin"
)

// newInstallCmd downloads and install a plugin from the PBR to the plugins directory
func newInstallCmd(c *command, os string) *cobra.Command {
	return &cobra.Command{
		Use:   "install <package> <version>",
		Short: "Installs or upgrades a package",
		Args:  cobra.MinimumNArgs(2),
		RunE: c.Wrap(func(homeDir string, m config.Manifest, args []string) error {
			pluginName := strings.ToLower(args[0])
			version := args[1]

			// Check if plugin is already installed
			if p, ok := m.Plugins[pluginName]; ok {
				if version == p.Version {
					fmt.Printf("%q has already been installed.\n", pluginName)
					return nil
				}
			}

			// Download plugin from registry
			installedVersion, err := plugin.Install(homeDir, pluginName, version, os)
			if err != nil {
				return err
			}

			// Add plugin to manifest
			m.Plugins[installedVersion.Package.Name] = config.Plugin{
				Environment: installedVersion.Package.Environment,
				NetworkType: installedVersion.Package.NetworkType,
				Protocol:    installedVersion.Package.Name,
				Subtype:     installedVersion.Package.Subtype,
				Version:     installedVersion.Version,
			}

			if err := config.WriteFile(
				homeDir,
				config.ManifestFilename,
				m,
			); err != nil {
				return err
			}

			fmt.Printf("The package %q has been installed.\n", installedVersion.Package.Name)

			return nil
		}),
	}
}

package cli

import (
	"github.com/spf13/cobra"
)

func New(os, version string) *cobra.Command {
	c := &command{}

	rootCmd := &cobra.Command{
		Use:          "bpm",
		Short:        "Blockchain Package Manager (BPM) manages blockchain nodes on your own infrastructure.",
		SilenceUsage: true,
	}

	pf := rootCmd.PersistentFlags()
	pf.StringVar(&c.baseDir, "base-dir", "~/.bpm/", "The directory plugins and configuration are stored")
	pf.StringVar(&c.registry, "package-registry", "https://dev.registry.blockdaemon.com", "The package registry provides packages to install")
	pf.BoolVar(&c.debug, "debug", false, "Enable debug output")

	// Commands
	rootCmd.AddCommand(
		newConfigureCmd(c, os),
		newInstallCmd(c, os),
		newListCmd(c, os),
		newShowCmd(c, os),
		newStartCmd(c, os),
		newStatusCmd(c, os),
		newStopCmd(c, os),
		newUninstallCmd(c, os),
		newVersionCmd(version),
		newTestCmd(c, os),
		newSearchCmd(c, os),
		newInfoCmd(c, os),
		newRemoveCmd(c, os),
	)

	return rootCmd
}

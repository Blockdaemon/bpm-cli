package cli

import (
	stdos "os"
	"fmt"
	"time"

	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm/pkg/pbr"
	pkgversion "github.com/Blockdaemon/bpm/pkg/version"
	"github.com/mitchellh/go-homedir"
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

	// Cobra parses all parameters after the commands are added but we need the base-dir before that to load the manifest.
	// Intentionally ignoring the error here. At this stage there is no way of knowing if the user tried to pass --base-dir
	// but made a typo or if it's a legitimate flag for one of the subcommands.
	_ = rootCmd.PersistentFlags().Parse(stdos.Args)

	// Initialize
	homeDir, err := homedir.Expand(c.baseDir)
	if err != nil {
		exitWithError(err, rootCmd)
	}
	if !config.ManifestExists(homeDir) {
		if ask4confirm(fmt.Sprintf("Looks like bpm isn't initialized correctly in %q, do you want to do that now?", homeDir)) {
			if err := config.Init(homeDir); err != nil {
				exitWithError(err, rootCmd)
			}
		} else {
			exitWithError(fmt.Errorf("manifest not found in %q", homeDir), rootCmd)
		}
	}

	// Get manifest
	m, err := config.LoadManifest(homeDir)
	if err != nil {
		exitWithError(err, rootCmd)
	}

	// Check if version is up to date
	if time.Now().Sub(m.LatestCLIVersionUpdatedAt) > 12 * time.Hour {
		client := pbr.New(c.registry)
		ver, err := client.GetCLIVersion(os)
		if err != nil {
			// Could be an intermittant connectivity issue - Do not exit
			fmt.Printf("Cannot get latest bpm cli version from %q: %s\n", c.registry, err)
		} else {
			m.LatestCLIVersion = ver.Version
			m.LatestCLIVersionUpdatedAt = time.Now()
			m.Write()
		}
	}

	needsUpgrade, err := pkgversion.NeedsUpgrade(version, m.LatestCLIVersion)
	if err != nil {
		exitWithError(err, rootCmd)
	}
	if needsUpgrade {
		fmt.Printf("bpm version %q is available. Please upgrade as soon as possible!\n", m.LatestCLIVersion)
	}

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
		newTestCmd(c, os),
		newSearchCmd(c, os),
		newInfoCmd(c, os),
		newRemoveCmd(c, os),
		newVersionCmd(version),
	)

	return rootCmd
}

func exitWithError(err error, cmd *cobra.Command) {
	// Immitate cobra error handling with SilenceUsage = true
	cmd.Printf("Error: %s\n", err.Error())
	stdos.Exit(1)
}

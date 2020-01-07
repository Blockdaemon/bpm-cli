package cli

import (
	"fmt"
	stdos "os"
	"strings"
	"time"

	"github.com/Blockdaemon/bpm/pkg/command"
	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm/pkg/pbr"
	pkgversion "github.com/Blockdaemon/bpm/pkg/version"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path/filepath"
)

const (
	baseDirFlag  = "base-dir"
	registryFlag = "package-registry"
	debugFlag    = "debug"
	yesFlag      = "yes"
)

func New(os, version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:          "bpm",
		Short:        "Blockchain Package Manager (BPM) manages blockchain nodes on your own infrastructure.",
		SilenceUsage: true,
	}

	pf := rootCmd.PersistentFlags()
	pf.String(baseDirFlag, "~/.bpm/", "The directory plugins and configuration are stored")
	pf.String(registryFlag, "https://dev.registry.blockdaemon.com", "The package registry provides packages to install")
	pf.Bool(debugFlag, false, "Enable debug output")
	pf.BoolP(yesFlag, "y", false, `Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively`)
	if err := viper.BindPFlags(pf); err != nil {
		exitWithError(err, rootCmd)
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetEnvPrefix("bpm")
	viper.AutomaticEnv()

	// Cobra parses all parameters after the commands are added but we need the base-dir before that to load the manifest.
	// Intentionally ignoring the error here. At this stage there is no way of knowing if the user tried to pass --base-dir
	// but made a typo or if it's a legitimate flag for one of the subcommands.
	_ = rootCmd.PersistentFlags().Parse(stdos.Args)

	baseDir := viper.GetString(baseDirFlag)
	registry := viper.GetString(registryFlag)
	debug := viper.GetBool(debugFlag)
	yes := viper.GetBool(yesFlag)

	if debug {
		fmt.Println("Debug mode enabled. Output will be more verbose")
	}

	// Initialize
	homeDir, err := homedir.Expand(baseDir)
	if err != nil {
		exitWithError(err, rootCmd)
	}
	if !config.ManifestExists(homeDir) {
		if yes || ask4confirm(fmt.Sprintf("Looks like bpm isn't initialized correctly in %q, do you want to do that now?", homeDir)) {
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
	if debug {
		fmt.Println("Getting latest bpm cli version")
	}
	if time.Since(m.LatestCLIVersionUpdatedAt) > 12*time.Hour {
		client := pbr.New(registry)
		ver, err := client.GetCLIVersion(os)
		if err != nil {
			// Could be an intermittant connectivity issue - Do not exit
			fmt.Printf("Cannot get latest bpm cli version from %q: %s\n", registry, err)
		} else {
			m.LatestCLIVersion = ver.Version
			m.LatestCLIVersionUpdatedAt = time.Now()
			if err := m.Write(); err != nil {
				exitWithError(err, rootCmd)
			}
		}
	}

	if debug {
		fmt.Println("Checking if bpm cli needs to be upgraded")
	}
	if pkgversion.IsValidVersion(version) {
		needsUpgrade, err := pkgversion.NeedsUpgrade(version, m.LatestCLIVersion)
		if err != nil {
			exitWithError(err, rootCmd)
		}
		if needsUpgrade {
			fmt.Printf("bpm version %q is available. Please upgrade as soon as possible!\n", m.LatestCLIVersion)
		}
	} else {
		fmt.Printf("bpm version %q is not a valid semantic version. Assuming this is a development release and skipping upgrade check\n", version)
	}

	// Create shared context that holds data common used by (nearly) all commands
	absHomeDir, err := filepath.Abs(homeDir)
	if err != nil {
		exitWithError(err, rootCmd)
	}
	cmdContext := command.CmdContext{
		HomeDir:     absHomeDir,
		Manifest:    m,
		RuntimeOS:   os,
		RegistryURL: registry,
		Debug:       debug,
	}

	// Commands
	nodesCmd := &cobra.Command{
		Use:   "nodes",
		Short: "Manage blockchain nodes",
	}
	nodesCmd.AddCommand(
		newConfigureCmd(cmdContext),
		newShowCmd(cmdContext),
		newStartCmd(cmdContext),
		newUpgradeCmd(cmdContext),
		newStatusCmd(cmdContext),
		newStopCmd(cmdContext),
		newTestCmd(cmdContext),
		newRemoveCmd(cmdContext),
	)

	packagesCmd := &cobra.Command{
		Use:   "packages",
		Short: "Manage packages",
	}
	packagesCmd.AddCommand(
		newInstallCmd(cmdContext),
		newListCmd(cmdContext),
		newUninstallCmd(cmdContext),
		newSearchCmd(cmdContext),
		newInfoCmd(cmdContext),
	)

	rootCmd.AddCommand(
		nodesCmd,
		packagesCmd,
		newVersionCmd(version),
	)

	return rootCmd
}

func exitWithError(err error, cmd *cobra.Command) {
	// Immitate cobra error handling with SilenceUsage = true
	cmd.Printf("Error: %s\n", err.Error())
	stdos.Exit(1)
}

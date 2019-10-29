package cli

import (
	"strconv"
	"strings"

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
		newShowCmd(c),
		newStartCmd(c),
		newStatusCmd(c),
		newStopCmd(c),
		newUninstallCmd(c),
		newVersionCmd(version),
		newTestCmd(c),
		newSearchCmd(c, os),
		newInfoCmd(c, os),
	)

	return rootCmd
}

func parseKeyPairs(fields []string) map[string]interface{} {
	pairs := make(map[string]interface{}, len(fields))

	for _, field := range fields {
		pair := strings.Split(field, "=")

		// TODO: Add validation error
		if len(pair) > 1 {
			key := pair[0]
			value := pair[1]

			// Check if string is a float
			if f, err := strconv.ParseFloat(value, 64); err == nil {
				pairs[key] = f
				continue
			}

			// Check if string is an int
			if i, err := strconv.ParseInt(value, 10, 64); err == nil {
				pairs[key] = i
				continue
			}

			// Check if string is a bool
			if b, err := strconv.ParseBool(value); err == nil {
				pairs[key] = b
				continue
			}

			pairs[key] = value
		}
	}

	return pairs
}

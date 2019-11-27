package cli

import (
	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"path/filepath"
)

type cmdFunc func(string, config.Manifest, []string) error
type runEFunc func(*cobra.Command, []string) error

type command struct {
	baseDir  string
	registry string
	debug    bool
	yes		 bool
}

func (c *command) Wrap(f cmdFunc) runEFunc {
	return func(cmd *cobra.Command, args []string) error {
		// Initialize
		homeDir, err := homedir.Expand(c.baseDir)
		if err != nil {
			return err
		}
		absHomeDir, err := filepath.Abs(homeDir)
		if err != nil {
			return err
		}

		// Get manifest
		m, err := config.LoadManifest(homeDir)
		if err != nil {
			return err
		}

		return f(absHomeDir, m, args)
	}
}

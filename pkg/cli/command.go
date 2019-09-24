package cli

import (
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"golang.org/x/xerrors"
)

type cmdFunc func(string, config.Manifest, []string) error
type runEFunc func(*cobra.Command, []string) error

type command struct {
	baseDir string
	debug   bool
}

func (c *command) Wrap(f cmdFunc) runEFunc {
	return func(cmd *cobra.Command, args []string) error {
		homeDir, err := homedir.Expand(c.baseDir)
		if err != nil {
			return err
		}

		// Get manifest
		m := config.Manifest{}
		if err := config.ReadFile(
			homeDir,
			config.ManifestFilename,
			&m,
		); err != nil {
			// Create directories if manifest does not exist
			// Will recreate directories if the base dir is changed
			var pathError *os.PathError
			switch {
			case xerrors.As(err, &pathError):
				// Create directories
				if err := config.MakeDir(
					homeDir,
					config.NodesDir(homeDir),
					config.PluginsDir(homeDir),
				); err != nil {
					return err
				}

				// Create an empty manifest file
				m = config.Manifest{
					Plugins: map[string]config.Plugin{},
				}

				if err := config.WriteFile(
					homeDir,
					config.ManifestFilename,
					&m,
				); err != nil {
					return err
				}
			default:
				return err
			}
		}

		return f(homeDir, m, args)
	}
}

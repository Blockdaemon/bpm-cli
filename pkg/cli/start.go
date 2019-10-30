package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"gitlab.com/Blockdaemon/bpm/pkg/plugin"
)

func newStartCmd(c *command) *cobra.Command {
	return &cobra.Command{
		Use:   "start <id>",
		Short: "Start a blockchain node",
		Args:  cobra.MinimumNArgs(1),
		RunE: c.Wrap(func(homeDir string, m config.Manifest, args []string) error {
			id := args[0]

			n, err := node.Load(config.NodesDir(homeDir), id)
			if err != nil {
				return err
			}
			pluginName := n.Protocol

			// Check if plugin is installed
			if _, ok := m.Plugins[pluginName]; !ok {
				fmt.Printf("The package %q is currently not installed.\n", pluginName)
				return nil
			}

			// Check if manual intervention is necessary in the configs
			// This is the case if a string like, e.g. {% ADD NODE KEY HERE %} is found in the files.
			// Until we have a better way of getting this information via the CLI, the users can edit the files manually.
			var ff = func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if info.IsDir() {
					return nil // skip dirs
				}

				content, err := ioutil.ReadFile(path)

				var substitute = regexp.MustCompile(`{%[^%]*%}`)

				matches := substitute.FindSubmatch(content)

				if len(matches) > 0 {
					return fmt.Errorf("The string %q needs to be replaced with a suitable value in %q", string(matches[0]), path)
				}

				return nil
			}

			if err := filepath.Walk(n.ConfigsDirectory(), ff); err != nil {
				return err
			}

			// Run the plugin
			if err := plugin.Start(homeDir, pluginName, id, c.debug); err != nil {
				return err
			}

			fmt.Printf("The node %q has been started.\n", id)

			return nil
		}),
	}
}

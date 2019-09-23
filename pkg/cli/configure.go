package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/rs/xid"
	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"golang.org/x/xerrors"
)

func newConfigureCmd(c *command) *cobra.Command {
	var fields []string

	cmd := &cobra.Command{
		Use:   "configure <package>",
		Short: "Configure a new blockchain node",
		Args:  cobra.MinimumNArgs(1),
		RunE: c.Wrap(func(homeDir string, m config.Manifest, args []string) error {
			pluginName := strings.ToLower(args[0])

			// Generate instance id
			id := xid.New().String()

			// Check if plugin is installed
			p, ok := m.Plugins[pluginName]
			if !ok {
				fmt.Printf("The package %q is currently not installed.\n", pluginName)
				return nil
			}

			// Create node config
			n, err := node.Load(config.NodesDir(homeDir), id)
			if err != nil {
				var pathError *os.PathError
				switch {
				case xerrors.As(err, &pathError):
					// Write node json if it was the first run
					n.Environment = p.Environment
					n.Protocol = p.Protocol
					n.NetworkType = p.NetworkType
					n.Subtype = p.Subtype
					n.Version = p.Version
					n.Config = parseKeyPairs(fields)

					// Only temporary until we find a better solution to distribute the certs
					n.Collection.Host = "dev-1.logstash.blockdaemon.com:5044"
					n.Collection.Cert = "~/.bpm/beats/beat.crt"
					n.Collection.CA = "~/.bpm/beats/ca.crt"
					n.Collection.Key = "~/.bpm/beats/beat.key"

					if err := config.WriteFile(
						n.NodeDirectory(),
						"node.json",
						n,
					); err != nil {
						return err
					}
				default:
					return err
				}
			}

			fmt.Printf("Node with id %q has been initialized, add your configuration (node.json) and secrets here:\n%s\n", id, n.NodeDirectory())

			return nil
		}),
	}

	cmd.Flags().StringSliceVar(&fields, "field", []string{}, "Custom fields to add to node.json")

	return cmd
}

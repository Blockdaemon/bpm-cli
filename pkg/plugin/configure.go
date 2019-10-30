package plugin

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/rs/xid"
	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"gitlab.com/Blockdaemon/bpm/pkg/pbr"
	"gitlab.com/Blockdaemon/bpm/pkg/version"
	"golang.org/x/xerrors"
)

func Configure(pluginName string, homeDir string, m config.Manifest, runtimeOS string, registry string, fields []string, skipUpgradeCheck bool) (string, error) {
	// Generate instance id
	id := xid.New().String()

	// Check if plugin is installed
	p, ok := m.Plugins[pluginName]
	if !ok {
		return fmt.Sprintf("The package %q is currently not installed.\n", pluginName), nil
	}

	if !skipUpgradeCheck {
		// Check if plugin is using the latest version
		client := pbr.New(registry)
		packageVersion, err := client.GetLatestPackageVersion(pluginName, runtimeOS)
		if err != nil {
			return "", err
		}
		latestVersion := packageVersion.Version

		needsUpgrade, err := version.NeedsUpgrade(p.Version, latestVersion)
		if err != nil {
			return "", err
		}

		if needsUpgrade {
			return fmt.Sprintf("A new version of package %q is available. Please install using \"bpm install %s %s\".\n", pluginName, pluginName, latestVersion), nil
		}
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
				return "", err
			}
		default:
			return "", err
		}
	}

	return fmt.Sprintf("Node with id %q has been initialized, add your configuration (node.json) and secrets here:\n%s\n", id, n.NodeDirectory()), nil
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

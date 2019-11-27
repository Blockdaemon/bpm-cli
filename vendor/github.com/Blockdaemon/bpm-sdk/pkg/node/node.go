// Package node provides an easy way to access node related information.
//
// Utility functions to generate names and directory paths encapsulate the package conventions.
// It is highly recommended to use this package when implementing a new package to achieve consistency
// across packages.
package node

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"os"

	"github.com/Blockdaemon/bpm-sdk/internal/util"
	homedir "github.com/mitchellh/go-homedir"
)

// Node represents a blockchain node, it's configuration and related information
type Node struct {
	baseDir string

	// The global ID of this node
	ID string `json:"id"`

	// The plugin name
	PluginName string `json:"plugin"`

	// Dynamic (i.e. defined by the plugin) string parameters
	StrParameters map[string]string `json:"str_parameters"`

	// Dynamic bool parameters
	BoolParameters map[string]bool `json:"bool_parameters"`

	// Describes the collection configuration
	Collection Collection `json:"collection"`

	// Secrets (Example: Private keys)
	Secrets map[string]interface{} `json:"-"` // No json here, never serialize secrets!

	// The package version used to install this node (if installed yet)
	// This is useful to know in order to run migrations on upgrades.
	Version string `json:"version"`
}

// Collection is config for log and node data collection
type Collection struct {
	CA   string `json:"ca"`
	Cert string `json:"cert"`
	Host string `json:"host"`
	Key  string `json:"key"`
}

// NamePrefix returns the prefix used as a convention when naming containers, volumes, networks, etc.
func (c Node) NamePrefix() string {
	return fmt.Sprintf("bpm-%s-", c.ID)
}

// NodeDirectory returns the base directory under which all configuration, secrets and meta-data for this node is stored
func (c Node) NodeDirectory() string {
	expandedBaseDir, err := homedir.Expand(c.baseDir)
	if err != nil {
		panic(err) // Should never happen because, at this stage, the directory should already be created
	}

	return path.Join(expandedBaseDir, c.ID)
}

// NodeFile returns the filepath in which the base configuration as well as meta-data from the PBG is stored
func (c Node) NodeFile() string {
	return path.Join(c.NodeDirectory(), "node.json")
}

// ConfigsDirectorys returns the directory under which all configuration for the blockchain client is stored
func (c Node) ConfigsDirectory() string {
	return path.Join(c.NodeDirectory(), "configs")
}

// ConfigsDirectorys returns the directory under which all secrets for the blockchain client is stored
func (c Node) SecretsDirectory() string {
	return path.Join(c.NodeDirectory(), "secrets")
}

// Save the node data
func (c Node) Save() error {
	// Create node directories if they don't exist yet
	_, err := util.MakeDirectory(c.SecretsDirectory())
	if err != nil {
		return err
	}

	_, err = util.MakeDirectory(c.ConfigsDirectory())
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(
		c.NodeFile(),
		data,
		os.ModePerm,
	)
}

func New(baseDir, id string) Node {
	return Node{
		baseDir: baseDir,
		ID:      id,
	}
}

// Load all the data for a particular node and creates all required directories
func Load(baseDir, id string) (Node, error) {
	node := New(baseDir, id)

	// Load node data
	nodeData, err := ioutil.ReadFile(node.NodeFile())
	if err != nil {
		return node, err
	}

	if err = json.Unmarshal(nodeData, &node); err != nil {
		return node, err
	}

	// Load secrets
	node.Secrets = make(map[string]interface{})

	files, err := ioutil.ReadDir(node.SecretsDirectory())
	if err != nil {
		return node, err
	}

	for _, f := range files {
		if !f.IsDir() {
			secret, err := ioutil.ReadFile(path.Join(node.SecretsDirectory(), f.Name()))
			if err != nil {
				return node, err
			}

			node.Secrets[f.Name()] = string(secret)
		}
	}

	return node, nil
}


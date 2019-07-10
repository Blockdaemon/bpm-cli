package node

import (
	"io/ioutil"
	"encoding/json"
	"path"
	homedir "github.com/mitchellh/go-homedir"
	"gitlab.com/Blockdaemon/blockchain/bpm-lib/internal/util"
)

type Node struct {
	NodeGID string `json:"node_gid"`
	BlockchainGID string `json:"blockchain_gid"`

	Environment  string `json:"environment"`
	NetworkType  string `json:"network_type"`
	NodeSubtype  string `json:"node_subtype"`
	ProtocolType string `json:"protocol_type"`

	Config map[string]interface{} `json:"config"`
	Secrets map[string]interface{} // No json here, never serialize secrets!

	baseDir string
}

func (c Node) DockerNetworkName() string {
	return "bd-" + c.NodeGID
}

func (c Node) ContainerName(containerName string) string {
	return "bd-" + c.NodeGID + "-" + containerName
}

func (c Node) VolumeName(volumeName string) string {
	return "bd-" + c.NodeGID + "-" + volumeName
}

func (c Node) NodeDirectory(baseDir string) string {
	expandedBaseDir, err := homedir.Expand(baseDir)
	if err != nil {
		panic(err) // Should never happen because, at this stage, the directory should already be created
	}

	return path.Join(expandedBaseDir, "nodes", c.NodeGID)
}

func (c Node) ConfigsDirectory() string {
	return path.Join(c.NodeDirectory(c.baseDir), "configs")
}

func (c Node) SecretsDirectory() string {
	return path.Join(c.NodeDirectory(c.baseDir), "secrets")
}


func (c Node) MakeConfigsDirectory() (string, error) {
	return util.MakeDirectory(c.ConfigsDirectory())
}

func (c Node) MakeSecretsDirectory() (string, error) {
	return util.MakeDirectory(c.SecretsDirectory())
}

func (c Node) WritePluginVersion(version string) error {
	configsDir, err := c.MakeConfigsDirectory()

	if err != nil {
		return err
	}


	return ioutil.WriteFile(configsDir, []byte(version), 0644)
}

func LoadNode(baseDir, nodeGID string) (Node, error) {
	var node Node

	// Prepare directories
	configDir, err := util.MakeDirectory(baseDir, "nodes")
	if err != nil {
		return node, err
	}
	nodeDir, err := util.MakeDirectory(configDir, nodeGID)
	if err != nil {
		return node, err
	}

	// Load config
	configPath := path.Join(nodeDir, "node.json")
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return node, err
	}

	err = json.Unmarshal(configData, &node)
	if err != nil {
		return node, err
	}

	// Load secrets
	node.Secrets = make(map[string]interface{})

	secretsDir, err := util.MakeDirectory(nodeDir, "secrets")
	if err != nil {
		return Node{}, err
	}

	files, err := ioutil.ReadDir(secretsDir)
    if err != nil {
    	return node, err
    }

    for _, f := range files {
    	if !f.IsDir() {
	    	secret, err := ioutil.ReadFile(path.Join(secretsDir, f.Name()))
	    	if err != nil {
	    		return node, err
	    	}

	    	node.Secrets[f.Name()] = string(secret)
    	}
    }

    node.baseDir = baseDir

	return node, nil
}


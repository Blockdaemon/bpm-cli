package models

import (
	"io/ioutil"
	"encoding/json"
	"path"
	homedir "github.com/mitchellh/go-homedir"
)

type NodeConfiguration struct {
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

func (c NodeConfiguration) DockerNetworkName() string {
	return "bd-" + c.NodeGID
}

func (c NodeConfiguration) ContainerName(containerName string) string {
	return "bd-" + c.NodeGID + "-" + containerName
}


func (c NodeConfiguration) NodeDirectory(baseDir string) string {
	expandedBaseDir, err := homedir.Expand(baseDir)
	if err != nil {
		panic(err) // Should never happen because, at this stage, the directory should already be created
	}

	return path.Join(expandedBaseDir, "nodes", c.NodeGID)
}

func (c NodeConfiguration) ConfigsDirectory() string {
	return path.Join(c.NodeDirectory(c.baseDir), "configs")
}

func (c NodeConfiguration) SecretsDirectory() string {
	return path.Join(c.NodeDirectory(c.baseDir), "secrets")
}


func (c NodeConfiguration) MakeConfigsDirectory() (string, error) {
	return makeDirectory(c.ConfigsDirectory())
}

func (c NodeConfiguration) MakeSecretsDirectory() (string, error) {
	return makeDirectory(c.SecretsDirectory())
}

func (c NodeConfiguration) WritePluginVersion(version string) error {
	configsDir, err := c.MakeConfigsDirectory()

	if err != nil {
		return err
	}


	return ioutil.WriteFile(configsDir, []byte(version), 0644)
}

func LoadConfiguration(baseDir, nodeGID string) (NodeConfiguration, error) {
	var configuration NodeConfiguration

	// Prepare directories
	configDir, err := makeDirectory(baseDir, "nodes")
	if err != nil {
		return configuration, err
	}
	nodeDir, err := makeDirectory(configDir, nodeGID)
	if err != nil {
		return configuration, err
	}

	// Load config
	configPath := path.Join(nodeDir, "config.json")
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return configuration, err
	}

	err = json.Unmarshal(configData, &configuration)
	if err != nil {
		return configuration, err
	}

	// Load secrets
	configuration.Secrets = make(map[string]interface{})

	secretsDir, err := makeDirectory(nodeDir, "secrets")
	if err != nil {
		return NodeConfiguration{}, err
	}

	files, err := ioutil.ReadDir(secretsDir)
    if err != nil {
    	return configuration, err
    }

    for _, f := range files {
    	if !f.IsDir() {
	    	secret, err := ioutil.ReadFile(path.Join(secretsDir, f.Name()))
	    	if err != nil {
	    		return configuration, err
	    	}

	    	configuration.Secrets[f.Name()] = string(secret)
    	}
    }

    configuration.baseDir = baseDir

	return configuration, nil
}


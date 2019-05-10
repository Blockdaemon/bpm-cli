package models

type Container struct {
	Image   string `json:"image"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type NodeConfiguration struct {
	NodeGID string `json:"node_gid"`

	Containers []Container `json:"containers"`

	Environment  string `json:"environment"`
	NetworkType  string `json:"network_type"`
	NodeSubtype  string `json:"node_subtype"`
	ProtocolType string `json:"protocol_type"`

	BlockchainGID string `json:"blockchain_gid"`

	Config map[string]interface{} `json:"config"`
}

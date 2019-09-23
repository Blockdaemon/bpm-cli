package config

const ManifestFilename = "manifest.json"

type Plugin struct {
	Environment string `json:"environment"`
	NetworkType string `json:"networkType"`
	Protocol    string `json:"protocol"`
	Subtype     string `json:"subtype"`
	Version     string `json:"version"`
}

type Manifest struct {
	// Plugins are a map of package name -> version
	Plugins map[string]Plugin `json:"plugins"`
}

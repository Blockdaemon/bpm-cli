package config

import (
	"time"

	"go.blockdaemon.com/bpm/sdk/pkg/plugin"
)

// ManifestFilename contains the filename of the manifest file
const ManifestFilename = "manifest.json"

// Manifest contains information about the installed plugins and cli
type Manifest struct {
	// Plugins are a map of package name -> version
	Plugins                   map[string]plugin.MetaInfo `json:"plugins"`
	LatestCLIVersion          string                     `json:"latest_cli_version"`
	LatestCLIVersionUpdatedAt time.Time                  `json:"latest_cli_version_updated_at"`

	// this could be internal (lower case) but golanglint-ci will complain
	Path string `json:"-"`
}

// ManifestExists returns true if the manifest already exists
func ManifestExists(path string) bool {
	return FileExists(path, ManifestFilename)
}

// LoadManifest reads the manifest from a file
func LoadManifest(path string) (Manifest, error) {
	m := Manifest{Path: path}

	err := ReadFile(path, ManifestFilename, &m)
	return m, err
}

// Init initializes a new manifest
func Init(path string) error {
	err := MakeDir(path, NodesDir(path), PluginsDir(path))
	if err != nil {
		return err
	}
	m := Manifest{
		Plugins:          map[string]plugin.MetaInfo{}, // initialize with empty map to avoid `assignment to entry in nil map`
		LatestCLIVersion: "0.0.0",                      // avoid "is not in dotted-tri format" errors
		Path:             path,
	}
	return m.Write()
}

// Write writes the manifest to disk
func (m *Manifest) Write() error {
	return WriteFile(m.Path, ManifestFilename, m)
}

// UpdatePlugin updates information about one particular plugin
//
// To write the changes to disk use `Write`
func (m *Manifest) UpdatePlugin(pluginName string, pluginInfo plugin.MetaInfo) error {
	m.Plugins[pluginName] = pluginInfo
	return m.Write()
}

// RemovePlugin removes a plugin from the manifest
//
// To write the changes to disk use `Write`
func (m *Manifest) RemovePlugin(pluginName string) error {
	delete(m.Plugins, pluginName)
	return m.Write()
}

// HasPluginsInstalled returns true if there is at least one installed plugin
func (m *Manifest) HasPluginsInstalled() bool {
	return len(m.Plugins) > 0
}

package config

import (
	"go.blockdaemon.com/bpm/sdk/pkg/plugin"
	"time"
)

const ManifestFilename = "manifest.json"

type Manifest struct {
	// Plugins are a map of package name -> version
	Plugins                   map[string]plugin.MetaInfo `json:"plugins"`
	LatestCLIVersion          string                     `json:"latest_cli_version"`
	LatestCLIVersionUpdatedAt time.Time                  `json:"latest_cli_version_updated_at"`

	// this could be internal (lower case) but golanglint-ci will complain
	Path string `json:"-"`
}

func ManifestExists(path string) bool {
	return FileExists(path, ManifestFilename)
}

func LoadManifest(path string) (Manifest, error) {
	m := Manifest{Path: path}

	err := ReadFile(path, ManifestFilename, &m)
	return m, err
}

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

func (m *Manifest) Write() error {
	return WriteFile(m.Path, ManifestFilename, m)
}

func (m *Manifest) UpdatePlugin(pluginName string, meta plugin.MetaInfo) error {
	m.Plugins[pluginName] = meta
	return m.Write()
}

func (m *Manifest) RemovePlugin(pluginName string) error {
	delete(m.Plugins, pluginName)
	return m.Write()
}

func (m *Manifest) HasPluginsInstalled() bool {
	return len(m.Plugins) > 0
}

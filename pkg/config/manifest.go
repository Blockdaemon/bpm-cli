package config

const ManifestFilename = "manifest.json"

type Plugin struct {
	Version string `json:"version"`
}

type Manifest struct {
	// Plugins are a map of package name -> version
	Plugins map[string]Plugin `json:"plugins"`

	// this could be internal (lower case) but golanglint-ci will complain
	Path string `json:"-"`
}

func LoadManifest(path string) (Manifest, error) {
	m := Manifest{
		Plugins: map[string]Plugin{}, // initialize with empty map to avoid `assignment to entry in nil map`
		Path:    path,
	}

	var err error
	if FileExists(path, ManifestFilename) {
		err = ReadFile(path, ManifestFilename, &m)
	}
	return m, err
}

func (m *Manifest) Write() error {
	return WriteFile(m.Path, ManifestFilename, m)
}

func (m *Manifest) UpdatePlugin(pluginName, version string) error {
	m.Plugins[pluginName] = Plugin{Version: version}
	return m.Write()
}

func (m *Manifest) RemovePlugin(pluginName string) error {
	delete(m.Plugins, pluginName)
	return m.Write()
}

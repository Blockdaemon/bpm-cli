package plugin

import (
	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm/pkg/manager"
	"github.com/Blockdaemon/bpm/pkg/pbr"
)

func Install(homeDir, registry, name, version, opSys string) (pbr.Version, error) {
	client := pbr.New(registry)

	// Get the specified version from the registry
	ver, err := client.GetPackageVersion(name, version, opSys)
	if err != nil {
		return ver, err
	}

	// Download the plugin file
	if err := manager.DownloadToFile(
		config.PluginsDir(homeDir),
		ver.Package.Name,
		ver.RegistryURL,
	); err != nil {
		return ver, err
	}

	return ver, nil
}

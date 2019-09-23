package plugin

import (
	"os"

	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"gitlab.com/Blockdaemon/bpm/pkg/manager"
	"gitlab.com/Blockdaemon/bpm/pkg/pbr"
)

func Install(homeDir, name, version, opSys string) (pbr.Version, error) {
	client := pbr.New(os.Getenv("BPM_REGISTRY_ADDR"))

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

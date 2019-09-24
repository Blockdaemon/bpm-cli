package plugin

import (
	"path/filepath"

	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"gitlab.com/Blockdaemon/bpm/pkg/manager"
)

func Stop(homeDir, name, id string, purge bool, debug bool) error {
	// Run plugin commands
	pluginFilename := filepath.Join(config.PluginsDir(homeDir), name)
	baseDirArgs := []string{"--base-dir", config.NodesDir(homeDir)}

	// Secrets
	stopArgs := append([]string{"stop", id}, baseDirArgs...)

	if purge {
		stopArgs = append(stopArgs, "--purge")
	}

	_, err := manager.ExecCmd(debug, pluginFilename, stopArgs...)
	if err != nil {
		return err
	}

	return nil
}

package plugin

import (
	"fmt"
	"path/filepath"

	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"gitlab.com/Blockdaemon/bpm/pkg/manager"
)

func Stop(homeDir, name, id string, purge bool) error {
	// Run plugin commands
	pluginFilename := filepath.Join(config.PluginsDir(homeDir), name)
	baseDirArgs := []string{"--base-dir", config.NodesDir(homeDir)}

	// Secrets
	stopArgs := append([]string{"stop", id}, baseDirArgs...)

	if purge {
		stopArgs = append(stopArgs, "--purge")
	}

	output, err := manager.ExecCmd(name, pluginFilename, stopArgs...)
	if err != nil {
		fmt.Println(output)
		return err
	}

	return nil
}

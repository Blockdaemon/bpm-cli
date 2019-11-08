package plugin

import (
	"fmt"
	"path/filepath"

	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm/pkg/manager"
)

func Stop(homeDir, name, id string, debug bool) error {
	// Run plugin commands
	pluginFilename := filepath.Join(config.PluginsDir(homeDir), name)
	baseDirArgs := []string{"--base-dir", config.NodesDir(homeDir)}

	// Secrets
	stopArgs := append([]string{"stop", id}, baseDirArgs...)

	output, err := manager.ExecCmd(debug, pluginFilename, stopArgs...)
	if err != nil {
		return err
	}

	fmt.Println(output)

	return nil
}

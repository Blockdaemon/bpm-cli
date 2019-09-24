package plugin

import (
	"path/filepath"

	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"gitlab.com/Blockdaemon/bpm/pkg/manager"
)

// Status returns the status of a particular node
func Status(homeDir, pluginName, nodeID string, debug bool) (string, error) {
	// Run plugin commands
	pluginFilename := filepath.Join(config.PluginsDir(homeDir), pluginName)
	baseDirArgs := []string{"--base-dir", config.NodesDir(homeDir)}
	statusArgs := append([]string{"status", nodeID}, baseDirArgs...)

	return manager.ExecCmd(debug, pluginFilename, statusArgs...)
}

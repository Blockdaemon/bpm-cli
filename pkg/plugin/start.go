package plugin

import (
	"fmt"
	"path/filepath"

	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"gitlab.com/Blockdaemon/bpm/pkg/manager"
)

func Start(homeDir, name, id string, debug bool) error {
	// Run plugin commands
	pluginFilename := filepath.Join(config.PluginsDir(homeDir), name)
	baseDirArgs := []string{"--base-dir", config.NodesDir(homeDir)}

	// Start
	startArgs := append([]string{"start", id}, baseDirArgs...)
	output, err := manager.ExecCmd(debug, pluginFilename, startArgs...)

	// This needs to be printed even if err != nil because the output typically contains more information about what went wrong
	fmt.Println(output)

	if err != nil {
		return err
	}

	return nil
}

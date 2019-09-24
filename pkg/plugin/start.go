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

	// Secrets
	secretArgs := append([]string{"create-secrets", id}, baseDirArgs...)
	output, err := manager.ExecCmd(debug, pluginFilename, secretArgs...)
	if err != nil {
		return err
	}

	fmt.Println(output)

	// Config
	configArgs := append([]string{"create-configurations", id}, baseDirArgs...)
	output, err = manager.ExecCmd(debug, pluginFilename, configArgs...)
	if err != nil {
		return err
	}

	fmt.Println(output)

	// Start
	startArgs := append([]string{"start", id}, baseDirArgs...)
	output, err = manager.ExecCmd(debug, pluginFilename, startArgs...)
	if err != nil {
		return err
	}

	fmt.Println(output)

	return nil
}

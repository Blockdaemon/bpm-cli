package plugin

import (
	"fmt"
	"path/filepath"

	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm/pkg/manager"
)

func Test(homeDir, name, id string, debug bool) error {
	// Run plugin commands
	pluginFilename := filepath.Join(config.PluginsDir(homeDir), name)
	baseDirArgs := []string{"--base-dir", config.NodesDir(homeDir)}

	// Secrets
	testArgs := append([]string{"test", id}, baseDirArgs...)

	output, err := manager.ExecCmd(debug, pluginFilename, testArgs...)

	fmt.Println(output)

	if err != nil {
		return err
	}

	return nil
}

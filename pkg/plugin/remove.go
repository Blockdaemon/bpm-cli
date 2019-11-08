package plugin

import (
	"fmt"
	"os"
	"path/filepath"

	bpmconfig "github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm/pkg/manager"
)

func Remove(homeDir, name, id string, debug bool, all bool, data bool, config bool) error {
	baseDir := bpmconfig.NodesDir(homeDir)
	nodeDir := filepath.Join(baseDir, id)
	// Run plugin commands
	pluginFilename := filepath.Join(bpmconfig.PluginsDir(homeDir), name)
	baseDirArgs := []string{"--base-dir", baseDir}

	if config || all {
		cmdArgs := append([]string{"remove-config", id}, baseDirArgs...)

		output, err := manager.ExecCmd(debug, pluginFilename, cmdArgs...)

		fmt.Println(output)

		if err != nil {
			return err
		}
	}

	if data || all {
		cmdArgs := append([]string{"stop", id}, baseDirArgs...)

		output, err := manager.ExecCmd(debug, pluginFilename, cmdArgs...)

		fmt.Println(output)

		if err != nil {
			return err
		}

		cmdArgs = append([]string{"remove-data", id}, baseDirArgs...)

		output, err = manager.ExecCmd(debug, pluginFilename, cmdArgs...)

		fmt.Println(output)

		if err != nil {
			return err
		}
	}

	if all {
		cmdArgs := append([]string{"remove-node", id}, baseDirArgs...)

		output, err := manager.ExecCmd(debug, pluginFilename, cmdArgs...)

		fmt.Println(output)

		if err != nil {
			return err
		}

		fmt.Printf("Removing directory %q\n", nodeDir)
		os.RemoveAll(nodeDir)
	}

	return nil
}

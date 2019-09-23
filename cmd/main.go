package main

import (
	"log"
	"os"
	"runtime"

	"gitlab.com/Blockdaemon/bpm/pkg/cli"
)

// Set on compile: -ldflags "-X main.version=dev"
var version string

const versionDev = "0.0.0-dev"

func main() {
	logger := log.New(os.Stdout, "", 0)

	// Check the runner version
	/*runnerUpgradeVersion, err := plugin.CheckRunnerUpgradable(baseDir, runnerVersion)
	if err != nil {
		return "", err
	}
	if len(runnerUpgradeVersion) > 0 {
		return fmt.Sprintf(TEXT_NEW_BPM_VERSION, runnerUpgradeVersion), nil
	}*/

	if version == "" {
		version = versionDev
	}

	// Init cli and exec command
	rootCmd := cli.New(runtime.GOOS, version)
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal(err)
	}
}

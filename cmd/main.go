package main

import (
	"os"
	"runtime"

	"go.blockdaemon.com/bpm/cli/pkg/cli"
)

// Set on compile: -ldflags "-X main.version=dev"
var version string

const versionDev = "0.0.0"

func main() {
	if version == "" {
		version = versionDev
	}

	// Init cli and exec command
	rootCmd := cli.New(runtime.GOOS, version)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

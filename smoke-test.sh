#!/bin/bash
# This is a simple test that runs through all bpm commands.

basedir=./build/bpm

function getStatusColumn() {
	echo $(go run cmd/main.go --base-dir $basedir nodes status | cut -d'|' -f$1 | tail -n1 | tr -d " ")
}

function checkStatus() {
	status=$(getStatusColumn 3)

	if [[ $status != $1 ]];then
		exit 1
	fi
}

function cleanup {
	echo ">>> CLEANING UP - IGNORE ANY ERRORS AFTER THIS"
	set +e # Run cleanup regardless of errors
	set +x # Make output a bit nicer

	# Stop node if it is still running
	go run cmd/main.go --base-dir $basedir nodes stop $(getStatusColumn 1) &> /dev/null
}

function setup {
	# Clean up from previous test runs
	if [[ -d $basedir ]]; then
		/bin/rm -r $basedir
	fi

	trap cleanup EXIT

	set -e # Stop on first error
	set -x # Print commands

}

setup

go run cmd/main.go --yes --base-dir $basedir version
go run cmd/main.go --yes --base-dir $basedir packages list
go run cmd/main.go --yes --base-dir $basedir packages search polkadot
go run cmd/main.go --yes --base-dir $basedir packages install polkadot 1.1.0
go run cmd/main.go --yes --base-dir $basedir packages info polkadot
go run cmd/main.go --yes --base-dir $basedir nodes configure polkadot --name test-node
checkStatus "stopped"
go run cmd/main.go --yes --base-dir $basedir nodes start test-node
checkStatus "running"
go run cmd/main.go --yes --base-dir $basedir nodes show node test-node
go run cmd/main.go --yes --base-dir $basedir nodes show config test-node
# Commented out because it doesn't work in the CI due to networking issues
# go run cmd/main.go --yes --base-dir $basedir test $nodeID
go run cmd/main.go --yes --base-dir $basedir nodes stop test-node
go run cmd/main.go --yes --base-dir $basedir nodes remove --config test-node
go run cmd/main.go --yes --base-dir $basedir nodes remove --data test-node
go run cmd/main.go --yes --base-dir $basedir nodes remove --all test-node
go run cmd/main.go --yes --base-dir $basedir packages uninstall polkadot

echo ">>> DONE, ALL TESTS RAN SUCCESSFUL"

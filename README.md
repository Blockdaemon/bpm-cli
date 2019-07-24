# Description

BPM is the Blockchain Package Manager by Blockdaemon. It allows easy and uniform deployment, maintenance and upgrades of blockchain nodes.

BPM itself provides the framework, the actual deployment is performed by plugins.

# Requirements

- Linux or OSX
- Running Docker

`wget` must be installed to follow these instructions. The simplest way to install it on Ubuntu is:

	sudo apt install wget

and on OSX (using the [homebrew package manager](https://brew.sh/)):

	brew install wget

Please see the official [wget homepage](https://www.gnu.org/software/wget/) for further details.

# Installation

**These instructions are just temporary until we have proper packages!**

Download and copy the `bpm` binary into `/usr/local/bin` and make sure it is executable:

	wget https://runner-test.sfo2.digitaloceanspaces.com/bpm-<VERSION>-<OS>-amd64
	sudo cp bpm-<VERSION>-<OS>-amd64 /usr/local/bin/bpm
	sudo chmod 755 /usr/local/bin/bpm

Replace `<VERSION>` with the actual version of BPM (e.g. `0.2.0`) and `<OS>` with either `linux` or `darwin`.

# Usage

At any time, type `bpm --help` for a full list of all available commands.

First, let's list all available plugins:

	bpm list

Which should return a list of available plugins:

    NAME       INSTALLED VERSION   AVAILABLE VERSION
    ---------- ------------------- -------------------
    stellar    not installed       0.4.1
    polkadot   not installed       0.5.1

Then, install the latest version of one of the plugins:

	bpm install polkadot

During the development of a plugin it might be useful to specify which version to install. See `bpm install --help` for details.

Downloaded plugins are stored in `~/.blockdaemon/plugins`.

----

**The public blockchain gateway which would provide configuration for the nodes doesn't exist yet. Until it does, we have to mock it. This process will change!**

1. Download the mock config:
	- https://gitlab.com/Blockdaemon/bpm-stellar/raw/master/mock_node.json
	- https://gitlab.com/Blockdaemon/bpm-polkadot/raw/master/mock_node.json
2. Set the mock environment variables pointing to the config, node id and data ingress token:
```
export MOCK_NODE_FILE=${PWD}/mock_node.json
export MOCK_GID=test123
export DATA_INGRESS_TOKEN=XXXX
```

**End of workaround**

----

Let's start the blockchain client:

	bpm run polkadot --api-key test

**Because of the missing gateway, it doesn't matter what we pass in as api-key. This will change in the future!**

You should now see docker containers being started:

    docker ps

To remove the blockchain client, run:

	bpm remove polkadot

Please note that this doesn't remove data volumes or configuration. To force the removal of all this data, run:

	bpm remove polkadot --purge

Be careful with the `--purge` parameters. If you purge an already fully synced blockchain you loose all data and have to re-sync from scratch. Any manual customisations to the configuration files will be lost as well.

# Code structure

For easier testability we separate business logic from Cobra commands.

- `cmd/bpm/main.go` is the main entrypint
- `internal/bpm/cmd/` contains the Cobra commands. These are only responsible for:
	- Parsing arguments
	- Calling a function in `internal/bpm/tasks`
	- Returning an error or printing the output
- `internal/bpm/tasks/` contains the business logic. Each function performs a single task and return either an error or a string containing the output of the task. This makes it very simple to test the tasks.
- `internal/bpm/plugin/` contains the low-level functionality and types. These are mostly used in the higher-level business logic in `internal/bpm/tasks`.

# bpm

Blockchain Package Manager by [Blockdaemon](https://blockdaemon.com/). Deploy, maintain, and upgrade blockchain nodes on your own infrastructure.

## Installation

### Requirements

- [Docker](https://www.docker.com/)

### Download bpm binary

`wget` must be installed to follow these instructions. The simplest way to install it on Ubuntu is:

```bash
sudo apt install wget
```

and on MacOS (using the [homebrew package manager](https://brew.sh/)):

```bash
brew install wget
```

Please see the official [wget homepage](https://www.gnu.org/software/wget/) for further details.

```bash
wget https://runner-test.sfo2.digitaloceanspaces.com/bpm-<VERSION>-<OS>-amd64
sudo cp bpm-<VERSION>-<OS>-amd64 /usr/local/bin/bpm
sudo chmod 755 /usr/local/bin/bpm
```

Replace `<VERSION>` with the actual version of BPM (e.g. `0.2.0`) and `<OS>` with either `linux` or `darwin`.

## Usage

Set the current registry using the `BPM_REGISTRY_URL` env var.

```bash
Usage:
  bpm [command]

Available Commands:
  configure   Configure a new blockchain node
  help        Help about any command
  install     Installs or upgrades a package
  list        List available and installed blockchain protocols
  show        Print a resource to stdout
  start       Start a blockchain node
  status      Display statuses of configured nodes
  stop        Removes a running blockchain client. Data and configuration will not be removed.
  uninstall   Uninstall a package. Data and configuration will not be removed.
  version     Print the version

Flags:
      --base-dir string   The directory plugins and configuration are stored (default "~/.bpm/")
  -h, --help              help for bpm

Use "bpm [command] --help" for more information about a command.
```

### Example (Polkadot)

Install the polkadot blockchain:

```bash
export BPM_REGISTRY_URL= 
bpm install polkadot 1.0.0
```

Configure the blockchain node and optionally pass additional fields:

```bash
bpm configure polkadot --field name=polkadot
Node with id "bm0lmirmvbaj4is78gtg" has been initialized, add your configuration (node.json) and secrets here:
...
```

Add your configs and secrets to the directory above then:

```bash
bpm start polkadot bm0lmirmvbaj4is78gtg
```

Check for running nodes with the status command:

```bash
bpm status
        NODE ID        | PLUGIN  |  STATUS  | SECRETS
+----------------------+---------+----------+---------+
  bm0lmirmvbaj4is78gtg | running | polkadot |       0
```

To remove the blockchain client, run:

```bash
bpm stop polkadot
```

Please note that the above does not remove data volumes or configurations. To force the removal of all data use:

```bash
bpm stop polkadot --purge
```

> Be careful with the `--purge` parameters. If you purge an already fully synced blockchain you loose all data and have to re-sync from scratch. Any manual customisations to the configuration files will be lost as well.

```bash
bpm uninstall polkadot
```

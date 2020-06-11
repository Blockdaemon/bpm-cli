# Overview

BPM provides tools and packages to run blockchain nodes on you own infrastructure using a unified simple command line interface.

* Deploy
* Maintain
* Monitor
* Upgrade

Get started launching nodes now by following the [Quickstart Tutorials](#section/Quickstart).

## System Requirements

- OSX or Linux operating system
- [Docker](https://docs.docker.com/install/) needs to be installed and running for most packages

# Installation

Depending on your operatin system, there are different ways to install BPM.

In the following instructions, please replace `<version>` with the latest BPM version.

## Check the integrity and authenticity of the installation files

With the exception of Homebrew, [which doesn't support packages signing](https://github.com/Homebrew/brew/pull/4120#issuecomment-406879969) all installation files are signed using
[GnuPG](https://gnupg.org). Before installing them it is strongly recommended to verify the
signature.

GPG needs to be installed to check the integrity and authenticity of the package.

Blockdaemon uses [keys.openpgpg.org](https://keys.openpgpg.org) to publish the release GPG key.

To verify the integrity and authenticity of a file, download the file itself as well as the corresponding signature `*.sig` file from the [BPM releases page](https://gitlab.com/Blockdaemon/bpm-cli/-/releases).

Then run the following command with `<filename>` replaced with the actual filename.

```bash
gpg --auto-key-locate hkps://keys.openpgp.org --auto-key-retrieve --verify <filename>.sig <filename>
```

You should see an output like this:

```bash
gpg: Signature made Fri 08 May 2020 05:45:41 PM UTC
gpg:                using RSA key F08C9DBB7215EC0FBC887A9E30F518E8C905E87B
gpg: requesting key 30F518E8C905E87B from hkps server keys.openpgp.org
gpg: /home/jonas/.gnupg/trustdb.gpg: trustdb created
gpg: key 30F518E8C905E87B: public key "Blockdaemon <support@blockdaemon.com>" imported
gpg: Total number processed: 1
gpg:               imported: 1
gpg: Good signature from "Blockdaemon <support@blockdaemon.com>" [unknown]
gpg: WARNING: This key is not certified with a trusted signature!
gpg:          There is no indication that the signature belongs to the owner.
Primary key fingerprint: F08C 9DBB 7215 EC0F BC88  7A9E 30F5 18E8 C905 E87B
```

The warning appears because there is no [web-of-trust](https://en.wikipedia.org/wiki/Web_of_trust) configured for this key. Instead you can verify the authenticity of the key by checking that the key fingerprint is `F08C 9DBB 7215 EC0F BC88  7A9E 30F5 18E8 C905 E87B` which ensures that this key is actually the official Blockdaemon release key. If there are any doubts, feel free to always reach out to support@blockdaemon.com.

## Install on OSX using Homebrew

The [Homebrew package manager](https://brew.sh/) needs to be installed.

In a terminal, run:

```bash
brew tap blockdaemon/homebrew-tap https://gitlab.com/Blockdaemon/homebrew-tap
brew install bpm-cli
```

## Install on Ubuntu using dpkg

1. Download the latest `bpm-cli_<version>_linux_amd64.deb` package from the [release page](https://gitlab.com/Blockdaemon/bpm-cli)
2. Verify the file as described [above](#section/Installation/Check-the-integrity-and-authenticity-of-the-installation-files)
3. Install the package: `sudo dpkg -i bpm-cli_<version>_linux_amd64.deb`

## Install on RedHat using rpm

1. Download the latest `bpm-cli_<version>_linux_amd64.rpm` package from the [release page](https://gitlab.com/Blockdaemon/bpm-cli)
2. Verify the file as described [above](#section/Installation/Check-the-integrity-and-authenticity-of-the-installation-files)
3. Install the package: `sudo rpm -ivh bpm-cli_<version>_linux_amd64.rpm`

## Install on any Linux system using the tar.gz package

1. Download the latest `bpm-cli_<version>_linux_amd64.tar.gz` package from the [release page](https://gitlab.com/Blockdaemon/bpm-cli)
2. Verify the file as described [above](#section/Installation/Check-the-integrity-and-authenticity-of-the-installation-files)
3. Change the working directory to where you want to install it: `cd /usr/local/bin`
4. Extract the BPM CLI from the package: `sudo tar -xvzf ~/Downloads/bpm-cli_<version>_linux_amd64.tar.gz bpm`
5. Ensure file permissions are set correctly: `sudo chmod 0755 ./bpm`
    
# Quickstart

The following tutorials give you a quick overview of BPM. We will install the BPM CLI and use it to configure and run Blockchain nodes.

The BPM CLI serves two main purposes:

1. Manage `packages`. A BPM `package` contains the functionality to launch a specific Blockchain node. There is typically one package per supported Blockchain client.
2. Manage `nodes`. A BPM `node` is a combination of the configuration, data, runtime as well as monitoring agents of a particular Blockchain node. Nodes can be started and stopped.

**Note:** While it is possible to run multiple Blockchain nodes at the same time we typically recomend to only ever have one node started at any time. There are multiple reasons for this but the main ones are:

* Blockchain nodes are resource intensive and running more than one can easily overwhelm the server.
* A lot of Blockchain protocols use the same common default ports which causes conflicts if more than one is run at the same time. While it is possible to change those ports, care has to be taken to avoid using the same ones.

## Tutorial 1: Managing packages

In this first tutorial you will learn how to install and manage BPM packages. To get an overview over what is possible, let's start with listing all commands available for managing packages:

```bash
bpm packages --help
```

The very first time you run `bpm` it needs to initialize it's directories, you should see the following output:

```
Looks like bpm isn't initialized correctly in "/Users/blockdaemon/.bpm", do you want to do that now? (y/N): y
```

Just type `y` to start the BPM initialization procedure.

You should see the following output:

```
Manage packages

Usage:
    bpm packages [command]

Available Commands:
    info        Show information about a package
    install     Installs or upgrades a package to a specific version or latest if no version is specified
    list        List installed packages
    search      Search available packages
    uninstall   Uninstall a package. Data and configuration will not be removed.

Flags:
    -h, --help   help for packages

Global Flags:
        --base-dir string           The directory plugins and configuration are stored (default "~/.bpm/")
        --debug                     Enable debug output
        --package-registry string   The package registry provides packages to install (default "https://dev.registry.blockdaemon.com")
    -y, --yes                       Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively

Use "bpm packages [command] --help" for more information about a command.
```

The first step is usually to search for a package to install. To list all available packages, just run the `search` command without a search parameter:

```bash
bpm packages search
```

You can also search for specific protocol packages by providing a search parameter (e.g. parity):
```bash
bpm packages search parity
```

After you have found a suitable packages, install it:

```bash
bpm packages install parity
```

This can take a few minutes while it is downloading the package.

To check which packages have already beeen installed, run the `list` command:

```bash
bpm packages list
```

You should see similar output:

```
    NAME  | INSTALLED VERSION | RECOMMENDED VERSION
+--------+-------------------+---------------------+
    parity | 1.0.0             | 1.0.0
```

This tells us that the parity package is installed using the latest recommended version.

To show more information about a package, use the `info` command:

```bash
bpm packages info parity
```

Finally, if a package isn't needed anymore, it can be uninstalled:

```bash
bpm packages uninstall parity
```

## Tutorial 2: Starting a Polkadot testnet node

First, let's list all commands available for managing nodes:

```bash
bpm nodes --help
```

You should see the following output:

```
Manage blockchain nodes

Usage:
    bpm nodes [command]

Available Commands:
    configure   Configure a new blockchain node
    remove      Remove blockchain node data and configuration
    show        Print a resource to stdout
    start       Start a blockchain node
    status      Display statuses of configured nodes
    stop        Stops a running blockchain node
    test        Tests a running blockchain node

Flags:
    -h, --help   help for nodes

Global Flags:
        --base-dir string           The directory plugins and configuration are stored (default "~/.bpm/")
        --debug                     Enable debug output
        --package-registry string   The package registry provides packages to install (default "https://dev.registry.blockdaemon.com")
    -y, --yes                       Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively

Use "bpm nodes [command] --help" for more information about a command.
```

### Installing the polkadot package

Before diving into the actual node management we need to install the polkadot package:

```bash
bpm packages install polkadot
```

### Configuring the node

Run the configure command with `--help` to see the available parameters:

```
bpm nodes configure polkadot --help
```

There are two parameters that are specific to the polkadot package:

```
--subtype watcher        The type of node. Must be either watcher or `validator` (default "watcher")
--validator-key string   The key used for a validator (required if subtype = validator)
```

For now we will leave the defaults which creates a configuration for a non-validating polkadot node. Run the configure command to create it now:

```bash
bpm nodes configure polkadot
```

You should see an output similar to this (but with different IDs):

```
Writing file '/Users/jonas/.bpm/nodes/little-glade-7812/configs/collector.env'
Writing file '/Users/jonas/.bpm/nodes/little-glade-7812/polkadot.dockercmd'

Node with id "little-glade-7812" has been initialized.

To change the configuration, modify the files here:
    /Users/jonas/.bpm/nodes/little-glade-7812
To start the node, run:
    bpm nodes start little-glade-7812
To see the status of configured nodes, run:
    bpm nodes status
```

### Node status

You can verify the status of all nodes by running the `status` command:

```bash
bpm nodes status
```

### Starting and stopping the node

To start the node, run the `node  start` command. Replace `<node-id>` with the ID outputed by the `configure` or `status` command:

```bash
bpm nodes start <node-id>
```

You should get an output similar to:

```
Network 'bpm' already exists, skipping creation
Creating container 'bpm-little-glade-7812-polkadot'
Starting container 'bpm-little-glade-7812-polkadot'
Creating container 'bpm-little-glade-7812-collector'
Starting container 'bpm-little-glade-7812-collector'
The node "little-glade-7812" has been started.
```

This shows BPM creating and starting the Docker containers for Polkadot and the monitoring agents.

Verify the node status by running:

```bash
bpm nodes status
```

The node should now show up as `running`, similar to below:

```
         NODE NAME         | PACKAGE  | STATUS
+--------------------------+----------+---------+
  little-glade-7812        | polkadot | running
```

To stop it temporarily, run:

```bash
bpm nodes stop <node-id>
```

### Removing a node

When removing a node you need to consider the following. A node consists of:

1. Secrets (e.g. accounts, passwords, private keys)
2. Node configuration
3. Runtime (e.g. Docker networks and containers)
4. Data (typically the parts of the Blockchain that have already been synced)

Depending on the use-case it may be desirable to remove all or only parts of the node. For example:

* In order to re-configure a node one might only want to remove the configuration but leave the data intact to avoid having to re-sync the Blockchain
* If the node crashed due to an unexpected error it can make sense to remove the runtime and start it again but keep the configuration and data
* If something went wrong during the initial sync it can help to remove the data and then start the node again to start syncing from scratch

To support the above use-cases plus others we have allowed parameters/flags to be used with our `remove` command.

You can view all the available parameters/flags by running the following BPM command:

```bash
bpm nodes remove --help
```

Which should return:

```
Remove blockchain node data and configuration

Usage:
    bpm nodes remove <id> [flags]

Flags:
        --all       Remove all data, configuration files and node information
        --config    Remove all configuration files but keep data and node information
        --data      Remove all data but keep configuration files and node information
    -h, --help      help for remove
        --runtime   Remove all runtimes but keep configuration files and node information

Global Flags:
        --base-dir string           The directory plugins and configuration are stored (default "~/.bpm/")
        --debug                     Enable debug output
        --package-registry string   The package registry provides packages to install (default "https://dev.registry.blockdaemon.com")
    -y, --yes                       Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively
```

For now, let's remove the whole node:

```bash
bpm nodes remove <node-id> --all
```

**Linux only:** Depending on the package, it is possible that the data created is inaccesible by your user. If that's the case you will get `permission denied` errors during the removal. To avoid the errors, run the bpm command using `sudo`.

```bash
sudo bpm nodes remove <node-id> --all
```

# Usage

BPM CLI comes with a built-in help functionality for all commands. You can view this information by adding `--help` to any command.

Examples:

```bash
bpm --help
bpm packages --help
bpm nodes --help
bpm nodes configure --help
```

The intention of this *Usage* guide is to provide additional information that is not already covered by `--help` on any given command. For examples on how to use BPM end-to-end please consult the [Quickstart tutorials](#section/Quickstart).

## Terms and concepts

bpm
: The Blockchain Package Manager is a combination of tools to simplify configuring and running Blockchain nodes in an enterprise environment.


bpm-cli
: The bpm-cli is the main user interface with which users will manage packages and nodes. It can be invoked by running `bpm <command>`.


bpm-sdk
: The bpm-sdk is a software library written in Go that makes it easy to develop new packages.


bpr
: The Blockchain Package Registry is a server component that allows to search for packages and versions of packages. The main interface to interact with bpr are the `bpm packages` commands within the bpm-cli.


package
: A BPM package contains the functionality to launch a specific Blockchain node. This typically involves creating secrets, configuration files and starting the Blockchain client as well as monitoring agents. Packages can be managed via the `bpm packages` commands.


node
: A BPM node is a combination of the configuration, data, runtime as well as monitoring agents of a particular Blockchain node. Nodes can be started and stopped. Nodes can be managed via the `bpm nodes` commands.


monitoring
: A key component of BPM is that all nodes are deployed with monitoring agents. Monitoring data can be sent to any [Logstash](https://www.elastic.co/products/logstash) instance or to the proprietary Blockdaemon monitoring system.

monitoring-pack
: A monitoring pack contains configuration and supporting files to send logs & monitoring data to an external monitoring system.

node-id
: Each BPM node gets assigned a unique ID that is used with the `bpm nodes` commands. (Example: `bnn4gnjo6him0f09h07g`)

## Removing nodes

When removing a node you need to consider the following. A node consists of:

1. Identity and similar secrets (e.g. accounts, private keys, certificates, ...)
2. Node configuration
3. Runtime (e.g. Docker networks and containers)
4. Data (typically the parts of the Blockchain that have already been synced)

Depending on the use-case it may be desirable to remove all or only parts of the node. For example:

* In order to re-configure a node one might only want to remove the configuration but leave the data intact to avoid having to re-sync the Blockchain
* If the node crashed due to an unexpected error it can make sense to remove the runtime and start it again but keep the configuration and data
* If something went wrong during the initial sync it can help to remove the data and then start the node again to start syncing from scratch

To support the above use-cases plus others we have allowed parameters/flags to be used with our `remove` command.

You can view all the available parameters/flags by running the following BPM command:

```bash
bpm nodes remove --help
```

Which should return:

```
Remove blockchain node data and configuration. Select one of the required flags for the remove command.

Usage:
    bpm nodes remove <name> [flags]

Flags:
        --all        [Required] Remove all data, configuration files and node information. Linux only: To avoid file permission denied errors on Linux use 'sudo' with this command
        --config     [Required] Remove all configuration files but keep data and node information
        --data       [Required] Remove all data but keep configuration files and node information. Linux only: To avoid file permission denied errors on Linux use 'sudo' with this command
    -h, --help       help for remove
        --identity   [Required] Remove the identity of the node
        --runtime    [Required] Remove all runtimes but keep configuration files and node information
```

**Linux only:** Depending on the package, it is possible that the data created is inaccesible by your user. If that's the case you will get `permission denied` errors during the removal. To avoid the errors, run the bpm command using `sudo`. Depending on how `sudo` is configured it is possible that this changes the home directory. Because the default location for the bpm data is in `$HOME/.bpm` there is a possibility that running bpm with `sudo` will target the wrong directory. To ensure that this doesn't happen it is advisable to use the `--base-dir` parameter to point to the absolute path of the bpm directory.


## Upgrades

Even considering the general speed of changes in information technology, Blockchain is an outlier because it changes even faster. Some Blockchains release new versions weekly and it is important to stay up-to-date for security reasons, to use the latest features as well as to be able to connect to the Blockchain. Not upgrading carries the risk of eventually being left behind or connecting to a fork of the Blockchain.

The BPM CLI regularly checks if it is running with the latest version. If you see a message like:

```
bpm version "0.8.0" is available. Please upgrade as soon as possible!
```

It is advised to install this newer version of the BPM CLI as soon as possible. Please see [Installation](#section/Installation) for installation instructions.

Similarly the BPM CLI will check if a particular package is up-to-date before configuring a new node. Since it is so crucial to always use the latest version of a Blockchain client, BPM will not allow configuring of new nodes with outdated packages unless specifically instructed by the user.

```bash
Error: A new version of package "parity" is available. Please install using "bpm install parity" or skip this check using "--skip-upgrade-check".
```

As described in the output one can always add `--skip-upgrade-check` to run an older version of a package. Do so at your own risk!

To upgrade a package, just install it again:

```bash
bpm packages install <package>
```

For development purposes it is possible to install a specific version using:

```bash
bpm packages install <package> <version>
```

To view a list of installed packages with their version and the recommended version to use, run:

```bash
bpm packages list
```

## Directories

**Note:** BPM doesn't dictate where a Blockchain node should be run or how it is being configured. At the moment most packages create configuration files and will start Docker containers. In the future there may be packages that, for example, create [Kubernetes secrets and configuration](https://kubernetes.io/docs/concepts/configuration/secret) instead of configuration directories.

By default BPM writes it's configuration to `~/.bpm`. This default can be changed by supplying a `--base-dir` parameter to all commands.

`~/.bpm/manifest.json`
: Contains version information and details about the installed plugins. You should never need to edit this file manually.


`~/.bpm/plugins`
: Contains binaries for the installed packages. While it is possible to run these binaries directly for development purposes it is recommended to use the `bpm` command as the main interface for end-users.


`~/.bpm/nodes/<node-id>/node.json`
: Contains configuration parameters and meta-information about an existing node.


`~/.bpm/nodes/<node-id>/configs/*`
: Contains configuration files for a particular node. Feel free to manually edit these files before launching a node if you need special configuration but also be aware that this may break things.

## Setting defaults for the BPM CLI

To avoid having to set parameters with every single `bpm` call, it is possible to definee them using environment variables.

Instead of specifying `--base-dir ~/.other-bpm` one can set an environment variable like this:

```bash
export BPM_BASE_DIR=~/.other-bpm
```

To enable debug mode all the time set this environment variable:

```bash
export BPM_DEBUG=true
```

## Automation

The BPM CLI sometimes asks for user confirmation. To use it in an automated non-interactive environment use the `--yes` parameter to always answer *yes* when prompted. Be careful and use only if you know what you are doing!

## Troubleshooting

* If somethign goes wrong, the first step should usually be to run the command again using the `--debug` parameter. This will print additional debug information which can aid in troubleshooting the issue.

* If a `bpm nodes start` command fails there is a good chance that something went wrong with the docker containers. Use the typical docker commands like `docker ps` or `docker logs` to see if all containers are running and whether there are any errors in the logs.

* Temporary docker issues can sometimes be resolved by removing all containers (`bpm nodes remove <node-id> --runtime`) followed by starting them again `bpm nodes start <node-id>`.

Please report any issues to [support@blockdaemon.com](mailto:support@blockdaemon.com).

# Logs and Monitoring

Most BPM packages come with monitoring agents included. By default BPM will collect logs and monitoring information and output them locally on the console. BPM will never send monitoring data or logs to a 3rd party unsolicited!

To connect a BPM node to an external monitoring system we need to specify how to connect, where to connect to, how to authenticate, etc. All of this information is stored in `monitoring-packs`. A monitoring-pack is a package containing configuration and supporting files to connect to a monitoring system.

To enable sending logs & monitoring data, specify the monitoring pack on the command line like this:

```bash
bpm nodes configure <package> --monitoring-pack <path-to-monitoring pack>
```

For example to connect a Tezos node to the Blockdaemon monitoring system the command would look like this:

```bash
bpm nodes configure tezos --monitoring-pack ~/Download/blockdaemon-monitoring-pack.tgz
```

If you are interested in using the Blockdaemon monitoring services and need a monitoring pack for this, please contact support@blockdaemon.com

# Packages

## Celo

[Celo](https://docs.celo.org/v/master/) is an ambitious staking protocol that hopes to solve issues with DPOS centralisation whilst also bringing financial services to those that are left out of the loop currently. Using the Celo Protocol there are 5 different types of nodes that make up the Celo ecosystem:

- Validator (including proxy nodes)
- Attestations (including attestation service node)
- Fullnodes

### Installing the Celo package
In order to manage Celo nodes we will need to install the Celo package:

```bash
bpm packages install celo
```

Configuring Celo nodes require specific parameters when running the configure command. Run the configure command with `--help` to see the available parameters for Celo nodes. Note not all parameters are needed for each node type:

```bash
bpm nodes configure celo --help
```

### Configure a Celo Fullnode

In order to configure a fullnode you will need to provide parameters to the configure command. These are the minimum parameters needed to configure a fullnode

```
--subtype string                  The type of node. Must be either validator, `proxy` or `fullnode` (default "fullnode")
--network string                  Mainnet or baklava testnet (default "baklava")
--networkid string                The current Celo network id
--account string                  The account to send rewards to
--bootnodes string                List of bootnodes to connect to
```
To view all available parameters run the configure command with `--help`:
```bash
bpm nodes configure celo --help
```

Configure the fullnode:
```bash
bpm nodes configure celo --subtype fullnode --network <network> --networkid <networkid> --account <account_address> --bootnodes "<enode_URLs>"
```

You should get an output similar to:
```
Writing file '/Users/blockdaemon/.bpm/nodes/small-water-3910/celo.dockercmd'
Writing file '/Users/blockdaemon/.bpm/nodes/small-water-3910/configs/collector.env'

Node with id "small-water-3910" has been initialized.

To change the configuration, modify the files here:
    /Users/blockdaemon/.bpm/nodes/small-water-3910
To start the node, run:
    bpm nodes start small-water-3910
To see the status of configured nodes, run:
    bpm nodes status
```

Start the fullnode with the provided command in the output above. 
```bash
bpm nodes start <node_id>
```

### Configure a Celo Validator Node

When running a validator node you will first need to configure a proxy node.

These are the minimum parameters needed to configure the proxy node:
```
--subtype string                  The type of node. Must be either validator, `proxy` or `fullnode` (default "fullnode")
--network string                  Mainnet or baklava testnet (default "baklava")
--networkid string                The current Celo network id
--signer string                   The signer address
--bootnodes string                List of bootnodes to connect to
```

To view all available parameters run the configure command with `--help`:
```bash
bpm nodes configure celo --help
```

Configure the proxy node:
```bash
bpm nodes configure celo --subtype proxy --network <network> --networkid <networkid> --signer <signer_address> --bootnodes "<enode_URLs>"
```

You should get an output similar to:
```
Writing file '/Users/blockdaemon/.bpm/nodes/damp-haze-2534/celo.dockercmd'
Writing file '/Users/blockdaemon/.bpm/nodes/damp-haze-2534/configs/collector.env'

Node with id "damp-haze-2534" has been initialized.

To change the configuration, modify the files here:
    /Users/blockdaemon/.bpm/nodes/damp-haze-2534
To start the node, run:
    bpm nodes start damp-haze-2534
To see the status of configured nodes, run:
    bpm nodes status
```

Start the proxy node with the provided command in the output above. 
```bash
bpm nodes start <node_id>
```

Next you will need to obtain the proxy enode url, the proxy ip, and your public ip in order to configure your validator.

The proxy enode url can be found with the following command:
```bash
docker exec bpm-<node_id>-proxy geth --exec "admin.nodeInfo['enode'].split('//')[1].split('@')[0]" attach | tr -d '"'
```

The proxy ip can be found with the following command:
```bash
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' bpm-<node_id>-proxy
```

Finally your public ip can be obtain with:
```bash
dig +short myip.opendns.com @resolver1.opendns.com
```

These are the minimum parameters needed to configure the validator node:
```
--subtype string                  The type of node. Must be either validator, `proxy` or `fullnode` (default "fullnode")
--network string                  Mainnet or baklava testnet (default "baklava")
--networkid string                The current Celo network id
--signer string                   The signer address
--keystore-file string            Location of the signer keystore json
--keystore-pass string            The password for the keystore json
--enode string                    The proxy enode id
--proxy_external string           The external proxy ip
--proxy_internal string           The internal proxy ip, use external if none
--port string                     Port to listen to (default "30303")
--bootnodes string                List of bootnodes to connect to
```

Now we are ready to configure our validator node:
```bash
bpm nodes configure celo --subtype validator --network <network> --networkid <networkid> --signer <signer_address> --keystore-file <path_to_keystore_file> --keystore-pass <path_to_keystore_password> --enode <proxy_enode> --proxy_internal <proxy_ip> --proxy_external <host_public_ip> --port <port> --bootnodes "<enode_URLs>"
```
Note: If both the proxy and validator are running on the same instance, you will need to change the port value for the validator to something other than 30303 since the proxy needs to communicate with the interweb.

You should get an output similar to :
```
Writing file '/Users/blockdaemon/.bpm/nodes/snowy-morning-3844/celo.dockercmd'
Writing file '/Users/blockdaemon/.bpm/nodes/snowy-morning-3844/configs/collector.env'
Writing file '/Users/blockdaemon/.bpm/nodes/snowy-morning-3844/configs/keystore/validator.signer-<address>.json'
Writing file '/Users/blockdaemon/.bpm/nodes/snowy-morning-3844/configs/.password.secret'

Node with id "snowy-morning-3844" has been initialized.

To change the configuration, modify the files here:
    /Users/blockdaemon/.bpm/nodes/snowy-morning-3844
To start the node, run:
    bpm nodes start snowy-morning-3844
To see the status of configured nodes, run:
    bpm nodes status
``` 

Start the validator with the provided command in the output above.
```bash
bpm nodes start <node_id>
```

### Configure a Celo Attestation Node

When running an attestation node you will also need to run an attestation-service node. First we will configure our attestation node.


These are the minimum parameters needed to configure the attestation node:
```
--subtype string                  The type of node. Must be either validator, `proxy` or `fullnode` (default "fullnode")
--network string                  Mainnet or baklava testnet (default "baklava")
--networkid string                The current Celo network id
--signer string                   The signer address
--keystore-file string            Location of the signer keystore json
--keystore-pass string            The password for the keystore json
--bootnodes string                List of bootnodes to connect to
```

To view all available parameters run the configure command with `--help`:
```bash
bpm nodes configure celo --help
```

Configure the attestation node:
```bash
bpm nodes configure celo --subtype attestation-node --network <network> --networkid <networkid> --signer <signer_address> --keystore-file <path_to_keystore_file> --keystore-pass <path_to_keystore_password> --bootnodes "<enode_URLs>"
```

You should get an output similar to:
```
Writing file '/Users/blockdaemon/.bpm/nodes/wild-bird-8942/configs/.password.secret'
Writing file '/Users/blockdaemon/.bpm/nodes/wild-bird-8942/celo.dockercmd'
Writing file '/Users/blockdaemon/.bpm/nodes/wild-bird-8942/configs/collector.env'
Writing file '/Users/blockdaemon/.bpm/nodes/wild-bird-8942/configs/keystore/attestation-node.signer.<address>.json'

Node with id "wild-bird-8942" has been initialized.

To change the configuration, modify the files here:
    /Users/blockdaemon/.bpm/nodes/wild-bird-8942
To start the node, run:
    bpm nodes start wild-bird-8942
To see the status of configured nodes, run:
    bpm nodes status
```

Start the attestation node with the provided command in the output above.
```bash
bpm nodes start <node_id>
```

Before configuring the attestation-service node the attestation node needs to be fully synced.

These are the minimum parameters needed to configure the attestation-service node:
```
--subtype string                  The type of node. Must be either validator, `proxy` or `fullnode` (default "fullnode")
--network string                  Mainnet or baklava testnet (default "baklava")
--networkid string                The current Celo network id
--signer string                   The signer address
--validator string                The validator address
--node_url string                 Attestation node url, eg http://bpm-flower-pot-1234-attestattion-node:8545
--db_user string                  Database user for attestation service postgres
--db_password string              Database password for attestation service postgres
--twilio_account_sid string       Twilio account SID for attesation service
--twilio_auth_token string        Auth token for Twilio
--twilio_service_sid string       Twilio messaging service SID for attestation services
--port string                     Port to listen to (default "30303")
```

To view all available parameters run the configure command with `--help`:
```bash
bpm nodes configure celo --help
```

Configure the attestation-service node:
```bash
bpm nodes configure celo --subtype attestation-service --network <network> --networkid <networkid> --signer <signer_address> --validator <validator_address> --node_url http://bpm-<attestation_node_id>-attestation-node:8545 --db_user <user> --db_password <password> --twilio_service_sid <twilio_service_sid> --twilio_account_sid <twilio_account_sid> --twilio_auth_token <twilio_token> --port <port>
```

You hsould get an output similar to:
```
Writing file '/Users/blockdaemon/.bpm/nodes/sparkling-field-360/celo.dockercmd'
Writing file '/Users/blockdaemon/.bpm/nodes/sparkling-field-360/configs/attestation-service.env'
Writing file '/Users/blockdaemon/.bpm/nodes/sparkling-field-360/configs/postgres.env'

Node with id "sparkling-field-360" has been initialized.

To change the configuration, modify the files here:
    /Users/blockdaemon/.bpm/nodes/sparkling-field-360
To start the node, run:
    bpm nodes start sparkling-field-360
To see the status of configured nodes, run:
    bpm nodes status
```

Start the attestation-service node with the provided command in the output above.
```bash
bpm nodes start <node_id>
```

Note your attestation-service node will show errors until your attenstation node is fully synced.

#### Node status

You can verify the status of all nodes by running the `status` command:

```bash
bpm nodes status
```

#### Starting and stopping a node

To start the node, run the `node start` command. Replace `<node-id>` with the ID outputed by the `configure` or `status` command:

```bash
bpm nodes start <node-id>
```

You should get an output similar to:

```
Creating container 'bpm-solitary-bird-663-celoinit'
Starting container 'bpm-solitary-bird-663-celoinit'
Stopping container 'bpm-solitary-bird-663-celoinit'
Removing container 'bpm-solitary-bird-663-celoinit'
Creating container 'bpm-solitary-bird-663-filebeat'
Starting container 'bpm-solitary-bird-663-filebeat'
Creating container 'bpm-solitary-bird-663-fullnode'
Starting container 'bpm-solitary-bird-663-fullnode'
Creating container 'bpm-solitary-bird-663-collector'
Starting container 'bpm-solitary-bird-663-collector'
The node "solitary-bird-663" has been started.
```

This shows BPM creating and starting the Docker containers for Celo itself as well as monitoring agents.

Verify the node status by running:

```bash
bpm nodes status
```

The node should now show up as `running`, similar to below:

```
        NODE NAME        | PACKAGE | STATUS
+------------------------+---------+---------+
  solitary-bird-663      | celo    | running
```

To stop it temporarily, run:

```bash
bpm nodes stop <node-id>
```

#### Testing a node

The nodes command allows you to test your running node. The tests being run are simple and should give you an indication if your node is running properly. Please allow the node to sync for a bit before running this command, this will ensure that you get the correct results.

You can test your node by running the following command:
```bash
bpm nodes test <node_id>
```

You should get an output similar to:
```
testing container: bpm-<node_id>
    Test [Container is running]   => true
    Test [Peer Count]   => 19
Total failed tests: 0
Total passed tests: 2
```

#### Removing a node

When removing a node you need to consider the following. A node consists of:

1. Node configuration and secrets (e.g. accounts, passwords, private keys)
2. Runtime (e.g. Docker networks and containers)
3. Data (typically the parts of the Blockchain that have already been synced)

Depending on the use-case it may be desirable to remove all or only parts of the node. For example:

* In order to re-configure a node one might only want to remove the configuration but leave the data intact to avoid having to re-sync the Blockchain
* If the node crashed due to an unexpected error it can make sense to remove the runtime and start it again but keep the configuration and data
* If something went wrong during the initial sync it can help to remove the data and then start the node again to start syncing from scratch

To support the above use-cases plus others we have allowed parameters/flags to be used with our `remove` command.

You can view all the available parameters/flags by running the following BPM command:

```bash
bpm nodes remove --help
```

For now, let's remove the complete node:

```bash
bpm nodes remove <node-id> --all
```

!!! warning
    By removing the whole node you will also remove the node identity. It's always advisable to backup the `~/.bpm/nodes/<node-id>` directory to a safe place before doing anything with the node.

## Parity

[Parity](https://www.parity.io/ethereum/) is a fast and feature-rich multi-network Ethereum client.

The Parity package is built specifically to run a node for a permissioned Ethereum Blockchain network.

1. A comma separated list of [enode URLs](https://github.com/ethereum/wiki/wiki/enode-url-format) that can act as *bootnodes*. These will serve as a list of peers that a new BPM Parity node will connect to.
2. A [chain specification](https://wiki.parity.io/Chain-specification) that contains the configuration for the permissioned Blockchain network.

**Note:** If you've used Blockdaemon to launch your Blockchain network you can find both parameters on your Blockdaemon dashboard. Otherwise please ask the administrator of the permissioned Blockchain network for parameters.

### Running a Parity node

First make sure that the latest version of the parity package is installed:

```bash
bpm packages install parity
```

Next, configure the node:

```bash
bpm nodes configure parity --bootnodes "<enode_URLs>" --chain-spec <chain-spec_filepath>
```

Depending on your specific parameters this could look something like:

```bash
bpm nodes configure parity --bootnodes "enode://8b1dfdfb03fbf40b7094aefa3e46a67b6e97cd06f28196413e76963935900b3c60bfc2c2ad7b14eec31245f7788cbf87cc8ca2352d138a736c081d7243e43163@33.236.162.172:30303,enode://aa32c91c962fad2cea3d6fac8a3ee9a86732e49423c78237a6827903ebba9d7b6a151d6dc14c2ce44663f140942afadd365290624f86c11c5a3adf4991c8ded7@19.199.180.203:30303" --chain-spec ~/Downloads/chain.json
```

You should see an output similar to this (but with different IDs):

```
generating accounts...
generating accounts...
Copying chain spec file to '/Users/blockdaemon/.bpm/nodes/bnn4trbo6him2ii9uca0/configs/chain.json'
Copying signer account file to '/Users/blockdaemon/.bpm/nodes/bnn4trbo6him2ii9uca0/configs/keys/DemoPoA/signer.json'
Copying signer password file file to '/Users/blockdaemon/.bpm/nodes/bnn4trbo6him2ii9uca0/configs/password.secret'
Writing file '/Users/blockdaemon/.bpm/nodes/bnn4trbo6him2ii9uca0/configs/parity.dockercmd'
Writing file '/Users/blockdaemon/.bpm/nodes/bnn4trbo6him2ii9uca0/configs/filebeat.yml'

Node with id "bnn4trbo6him2ii9uca0" has been initialized.

To change the configuration, modify the files here:
    /Users/blockdaemon/.bpm/nodes/bnn4trbo6him2ii9uca0/configs
To start the node, run:
    bpm nodes start bnn4trbo6him2ii9uca0
To see the status of configured nodes, run:
    bpm nodes status
```

Starting, stopping and removing the node is covered in detail in [Tutorial 2](#section/Quickstart/Tutorial-2:-Starting-a-Polkadot-testnet-node).

To summarize, use the following commands to manage the node:

```bash
bpm nodes start <node-id>
bpm nodes stop <node-id>
bpm nodes remove <node-id> --all
```

## Polkadot

[Polkadot](https://polkadot.network/) is a sharded protocol that enables blockchain networks to operate together seamlessly

See [Tutorial 2: Starting a Polkadot testnet node](#section/Quickstart/Tutorial-2:-Starting-a-Polkadot-testnet-node) for details on how to use the polkadot package.

## Tezos

[Tezos](https://tezos.gitlab.io/) is a distributed consensus platform with meta-consensus capability. Tezos not only comes to consensus about the state of its ledger, like Bitcoin or Ethereum. It also attempts to come to consensus about how the protocol and the nodes should adapt and upgrade.

The Tezos BPM package supports running and maintaining a Tezos node on mainnet or the carthagenet testnet.

### Configuring a Tezos node

Before diving into the actual node management we need to install the tezos package:

```bash
bpm packages install tezos
```

Run the configure command with `--help` to see the available parameters:

```
bpm nodes configure tezos --help
```

There are two parameters that are specific to the tezos package:

```
--network mainnet          The network. Can be either mainnet or `carthagenet` (default "mainnet")
--subtype watcher          The type of node. Only watcher supported currently (default "watcher")
```

For now we will create a configuration for a non-validating tezos node on mainnet. Run the configure command to create it now:

```bash
bpm nodes configure tezos --network mainnet
```

This will take a few minutes while we generate the node identity. The identity file is stored in `~/.bpm/nodes/<node-id>/identity`. Please make sure to backup this directory and keep it somewhere safe. This allows you to re-create the node.

If you already have an existing identity file that you want to use you can overwrite the identity file in `~/.bpm/nodes/<node-id>/identity` before starting the node.

Once the command is finished, you should see an output similar to this (but with different IDs):

```
Network 'bpm' already exists, skipping creation
Forwarding of monitoring is disabled. Specify `--monitoring-pack` to enable it.
/Users/jonas/.bpm/nodes/solitary-shadow-8810/identity-config/minimal-config.json
Creating container 'bpm-solitary-shadow-8810-tezos-init'
Starting container 'bpm-solitary-shadow-8810-tezos-init'
Stopping container 'bpm-solitary-shadow-8810-tezos-init'
Removing container 'bpm-solitary-shadow-8810-tezos-init'
�Generating a new identity... (level: 26.00)
YStored the new identity (idsWYEsCT91Vb6e6tmeGuxscWVRz7V) into '/identity/identity.json'.

Writing file '/Users/jonas/.bpm/nodes/solitary-shadow-8810/configs/config.json'
Writing file '/Users/jonas/.bpm/nodes/solitary-shadow-8810/configs/collector.env'

Node with id "solitary-shadow-8810" has been initialized.

To change the configuration, modify the files here:
    /Users/jonas/.bpm/nodes/solitary-shadow-8810
To start the node, run:
    bpm nodes start solitary-shadow-8810
To see the status of configured nodes, run:
    bpm nodes status
```

#### Node status

You can verify the status of all nodes by running the `status` command:

```bash
bpm nodes status
```

#### Starting and stopping the node

To start the node, run the `node start` command. Replace `<node-id>` with the ID outputed by the `configure` or `status` command:

```bash
bpm nodes start <node-id>
```

You should get an output similar to:

```
Creating container 'bpm-solitary-shadow-8810-filebeat'
Starting container 'bpm-solitary-shadow-8810-filebeat'
Creating container 'bpm-solitary-shadow-8810-tezos'
Starting container 'bpm-solitary-shadow-8810-tezos'
Creating container 'bpm-solitary-shadow-8810-collector'
Starting container 'bpm-solitary-shadow-8810-collector'
The node "solitary-shadow-8810" has been started.
```

This shows BPM creating and starting the Docker containers for Tezos itself as well as monitoring agents.

Verify the node status by running:

```bash
bpm nodes status
```

The node should now show up as `running`, similar to below:

```
         NODE NAME         | PACKAGE | STATUS
+--------------------------+---------+---------+
  solitary-shadow-8810     | tezos   | running
```

To stop it temporarily, run:

```bash
bpm nodes stop <node-id>
```

#### Removing the node

BPM allows fine-grained control over what gets deleted. Please see [Removing nodes](#section/Usage/Removing-nodes) for a full explanation of the `remove` command.

For now, let's remove the complete node:

```bash
bpm nodes remove <node-id> --all
```

# Architecture Considerations

This is a working document and subject to change. It captures the current architecture of BPM.

## Design Goals

**We embrace decentralization**

A BPM managed node should be run anywhere by anyone. Never should it be possible for a "service provider" (hint: we are one!) to change or modify a BPM node without the administrators explicit consent.

This has certain implications:

* BPM will not provide a way to manage the nodes "from outside"
* Monitoring data can flow from the node to a central point but not vice-versa.
* BPM will not auto-upgrade nodes without a users explicit consent.

**One interface to manage all blockchains while keeping most flexilibity**

Packages abstract away the details of a Blockchain client and only expose a clearly defined interface.

Blockchain is a bleeding-edge technology that constantly changes. While a clearly defined interface is desirable most of the time we want to enable users to edit configuration files easily and run one-off tasks where necessary for highly specialized use-cases. This is achieved by a two-step approach that creates a default configuration first, optionally let's the user manually edit this configuration and only then does the node get started.

**Simple by default; Flexible when necessary**

Blockchain protocols/clients vary greatly; BPM must be flexible enough to accompany them all.

This is achieved by:

* Keeping the `bpm-cli` itself as simple as possible and putting most of the actual logic in the packages allows to change nearly everything a package can do.
* Depending on the level of customization necessary, the BPM SDK provides different entry points. For exapmle there are pre-defined functions to create a Docker based package but one can always supply their own customized functions.
* The package interface is clearly defined. If the BPM SDK proves to be too limiting for a specific task it is always possible to write a package from scratch using an programming language.

**Monitoring is a first class citizen**

Each package must come with monitoring agents. Monitoring is the first step towards more reliable Blockchain infrastructure.

**Run nodes everywhere**

BPM is not designed with a paritcular deployment target in mind. While currently most Blockchain deploy to Docker it is possible to write packages that run directly on bare-metal or deploy to Kuberneters. This is achieved by putting most logic into the package or BPM SDK.

**Fast development cycle for new packages**

Packages can be developed by launching them directly (not advised for end-users!) and individual steps can be re-run easily (see *idempotency* below). This allows to quickly iterate during the development of packages.

Allowing developers (as well as users!) to edit configuration files manually helps with quickly testing out new ideas and configurations.

**Fully scriptable/automatable**

In enterprise environments nodes are typically launched automatically. It must always be possible to run BPM without human intervention.

This involves things like:

- Make sure to return correct error codes
- If a interactive prompt is used, provide a way to also use it non-interactively (`--yes`)
- Make use of idempotency to allow the same step to be re-run without accidentally overwriting

> An operation is idempotent if the result of performing it once is exactly the same as the result of performing it repeatedly without any intervening actions. [[1]](https://docs.ansible.com/ansible/latest/reference_appendices/glossary.html)

**Make it easy to backup and restore a node**

All information (e.g. secret keys) to restore a node must be easy to backup and restore. This is currently done by manually backing up the `~/.bpm/nodes/<node-id>/` directory.

In the future we plan much more backup&recovery functionality!

## Data flow

![Data flow](https://storage.googleapis.com/stg-blockdaemon-cli-bpm-docs/dataflow.png)

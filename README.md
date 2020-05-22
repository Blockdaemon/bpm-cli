# bpm

Blockchain Package Manager by [Blockdaemon](https://blockdaemon.com/). Deploy, maintain, and upgrade blockchain nodes on your own infrastructure.

Further reading:

* End-user documentation: https://cli.bpm.docs.blockdaemon.com/
* Developer documentation: https://sdk.bpm.docs.blockdaemon.com/

# Contributing

Pleaes use [conventional commits](https://www.conventionalcommits.org) for you commit messages. This will help us in the future to auto-generate changelogs.

New features should be developed in a branch and merged after a code review.

# Building from source

## Requirements

Make sure you have the following tools:

- [Go](https://golang.org/) is the main pogramming language. It needs to be installed
- [Docker](https://www.docker.com/) is used to run packages. It needs to installed and running
- [goreleaser](https://goreleaser.com/) is used to build binary packages. It needs to be installed
- [golangci-lint](https://github.com/golangci/golangci-lint) is used to do static code checks. It needs to be installed
- [GPG](https://gnupg.org/) is used to sign build artifacts. It needs to be installed

## Building during development

Make sure Docker is running, otherwise the tests will fail.

To build during development without creating a version or publishing binaries, run:

    make dev-release

## Releasing a new version

Make sure Docker is running, otherwise the tests will fail.

You need the Blockdaemon release GPG key imported into your GPG keyring.

    make version=<VERSION> release

`<VERSION>` needs to be a valid [semantic version](https://semver.org/). Do **not prefix with `v`**, the script does that automatically.

# Writing documentation

The documentation is built using [redoc](https://redocly.github.io/redoc/) which renders and openapi/swagger file. Most of the text is stored in the `docs/documentation.md` file.

In order to render the documentation locally while working on it, run:

    make serve-docs

and navigate to http://localhost:8080 to see a live rendering of the documentation.

# bpm

Blockchain Package Manager by [Blockdaemon](https://blockdaemon.com/). Deploy, maintain, and upgrade blockchain nodes on your own infrastructure.

Further reading:

* End-user documentation: https://cli.bpm.docs.blockdaemon.com/
* Developer documentation: https://sdk.bpm.docs.blockdaemon.com/

# Building from source

## Requirements

Make sure you have the following tools:

- [Go](https://golang.org/) is the main pogramming language. It needs to be installed
- [Docker](https://www.docker.com/) is used to run packages. It needs to installed and running
- [goreleaser](https://goreleaser.com/) is used to build binary packages. It needs to be installed
- [golangci-lint](https://github.com/golangci/golangci-lint) is used to do static code checks. It needs to be installed
- [GPG](https://gnupg.org/) is used to sign build artifacts. It needs to be installed

## Building during development

To build during development without creating a version or publishing binaries, run:

  goreleaser --snapshot --skip-publish --rm-dist

## Releasing a new version

1. Check that `CHANGELOG.md` is up-to-date

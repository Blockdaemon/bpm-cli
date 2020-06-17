# 0.14.1


### Bug Fixes

* linting complains about missing error handling ([d5b7bc9](https://gitlab.com/Blockdaemon/bpm-cli/commit/d5b7bc9ddaa90fb803d955bb0128b23de3f5298f))
* rpm and deb packages put binary in an incorrect directory ([5922234](https://gitlab.com/Blockdaemon/bpm-cli/commit/59222346f0be968f99f76304a83bfe2f49cb02ab))
* unable to install from file if there is a previous installation ([e4ecc69](https://gitlab.com/Blockdaemon/bpm-cli/commit/e4ecc6931119e937f0818b11c6792c4f2bc6c9b5))



# 0.14.0

### Bug Fixes

* remove nodes that failed during configuration ([dd9458b](https://gitlab.com/Blockdaemon/bpm-cli/commit/dd9458be69704ec9f2bd6e5f7faa24f1b8c7b926))

### Features

* support bpm-sdk protocol version 1.2.0 ([aa26d73](https://gitlab.com/Blockdaemon/bpm-cli/commit/aa26d73a89c6bddb17449cd167d369faa60cff0f))
* support plugins packaged in tar.gz ([f8c1925](https://gitlab.com/Blockdaemon/bpm-cli/commit/f8c192590a3d974ea0f102695d710b0026c88769))

# 0.10.4

Development improvements:

* Build process based on goreleaser. This creates tgz,dev,rpm and brew packages as well as checksums and gpg signatures
* Change project location to `go.blockdaemon.com/bpm/cli` which redirects to the actual repository
* Upgrade to bpm-sdk v0.13.0

Bug fixes:

* Add documentation and help to deal with file permission problems on Linux

# 0.10.2

Bug fixes:

* Clean up (i.e. remove) node file & directory if the parameter validation fails during `configure`

# 0.10.1

Bug fixes:

* Do not crash on `nodes status` if the package `status` command crashes. Show error instead

# 0.10.0

New functionality:

* Instead of cryptic ids, nodes are now named using a human readable name generator.

* The `nodes configure` command now has a new parameter `--name` that allows to name a node.

* The `nodes remove` command now has a new parameter `--identity` to only remove the identity of a node.

* The `packages uninstall` command is now always available and shows up in the help. Previously it was only available and shown in the help when at least one package was installed. This behaviour was a-typical for a cli tool and surprising for users.

* The `nodes configure` command now always shows up in the help. The reasoning is the same as with `packages uninstall`.

* Rename `packages search` to `packages list-remote` for better consistency with the `packages list` command. `packages search` is still supported for compatibility reasons.

* Both `packages list-remote` and `packages list` now show the same information (`NAME`, `DESCRIPTION`, `INSTALLED VERSION`, `RECOMMENDED VERSION`) for better consistency.

Bug fixes:

* The `nodes status` command now works even if a particular package is not installed anymore. Instead of crashing it will show `unknown (package not installed)` for every node for which the package isn't installed.

* The `nodes status` command now shows the values for PACKAGE and STATUS in the correct columns.

* Due to a race condition the `packages uninstall` command sometimes uninstalled the wrong package. This is now fixed.

* The help text for the `nodes remove` command now properly denotes which flags are required.

Development improvements:

* Support for new plugin protocol version 1.1.0
	* Rename the `create-secrets` call to `create-identity`. This makes the intend more clear.
	* Add a `remove-identity` call similar to `remove-config`
	* Add a `validate-parameters` call (parameters used to be validated implicitely when creating the configurations)
	* Adds plugin name to the plugin meta information

* Remove vendor directory: This makes it easier to diff github commits. The downside of having to rely on third-parties for dependencies is somewhat mitigated by the official [Google module mirror](https://proxy.golang.org/). We'll add our own mirror soon to have better control over dependencies and perform security checks on the dependencies.

* Upgraded internal dependencies


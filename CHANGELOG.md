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


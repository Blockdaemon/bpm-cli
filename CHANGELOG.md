# 0.10.0

New functionality:

* The `packages uninstall` command is now always available and shows up in the help. Previously it was only available and shown in the help when at least one package was installed. This behaviour was a-typical for a cli tool and surprising for users.

* The `nodes configure` command now always shows up in the help. The reasoning is the same as with `packages uninstall`.

* Rename `packages search` to `packages list-remote` for better consistency with the `packages list` command. `packages search` is still supported for compatibility reasons.

* Both `packages list-remote` and `packages list` now show the same information (`NAME`, `DESCRIPTION`, `INSTALLED VERSION`, `RECOMMENDED VERSION`) for better consistency.

Bug fixes:

* The `nodes status` command now works even if a particular package is not installed anymore. Instead of crashing it will show `unknown (package not installed)` for every node for which the package isn't installed.

* Due to a race condition the `packages uninstall` command sometimes uninstalled the wrong package. This is now fixed.

* The help text for the `nodes remove` command now properly denotes which flags are required.

Development improvements:

* Support for new plugin protocol version 1.1.0: This version removes the `create-secrets` call. There was a lot of confusion between `create-configurations` and `create-secrets` because most secrets needed to be copied into the configuration anyway to be used. The only benefit was the ability to backup all secrets separately from the configurations. This doesn't warrant the added complexity.
Compatibility with the old protocol version 1.0.0 is kept.

* Remove vendor directory: This makes it easier to diff github commits. The downside of having to rely on third-parties for dependencies is somewhat mitigated by the official [Google module mirror](https://proxy.golang.org/). We'll add our own mirror soon to have better control over dependencies and perform security checks on the dependencies.

* Upgraded internal dependencies


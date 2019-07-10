module gitlab.com/Blockdaemon/bpm

go 1.12

require (
	github.com/coreos/go-semver v0.3.0
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/kataras/tablewriter v0.0.0-20180708051242-e063d29b7c23 // indirect
	github.com/landoop/tableprinter v0.0.0-20180806200924-8bd8c2576d27
	github.com/mattn/go-runewidth v0.0.4 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3 // indirect
	gitlab.com/Blockdaemon/blockchain/bpm-lib v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.2.2 // indirect
)

replace gitlab.com/Blockdaemon/blockchain/bpm-lib => gitlab.com/Blockdaemon/blockchain/bpm-lib.git v0.0.0-20190710143909-67f7c2b4c0f9

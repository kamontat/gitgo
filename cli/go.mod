module github.com/kamontat/gitgo/cli

go 1.16

replace (
	github.com/kamontat/gitgo/config v0.0.0-local => ../config
	github.com/kamontat/gitgo/core v0.0.0-local => ../core
	github.com/kamontat/gitgo/git v0.0.0-local => ../git
	github.com/kamontat/gitgo/prompt v0.0.0-local => ../prompt
	github.com/kamontat/gitgo/utils v0.0.0-local => ../utils
)

require (
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/kamontat/gitgo/config v0.0.0-local
	github.com/kamontat/gitgo/core v0.0.0-local
	github.com/kamontat/gitgo/git v0.0.0-local
	github.com/kamontat/gitgo/prompt v0.0.0-local
	github.com/kamontat/gitgo/utils v0.0.0-local
	github.com/magiconair/properties v1.8.4 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/spf13/afero v1.5.1 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.1.3 // direct
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.7.1
	golang.org/x/crypto v0.0.0-20210317152858-513c2a44f670 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/text v0.3.5 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

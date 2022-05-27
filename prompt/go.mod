module github.com/kamontat/gitgo/prompt

go 1.16

replace (
	github.com/kamontat/gitgo/config v0.0.0-local => ../config
	github.com/kamontat/gitgo/git v0.0.0-local => ../git
	github.com/kamontat/gitgo/utils v0.0.0-local => ../utils
)

require (
	github.com/AlecAivazis/survey/v2 v2.2.9
	github.com/kamontat/gitgo/config v0.0.0-local
	github.com/kamontat/gitgo/git v0.0.0-local
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
)

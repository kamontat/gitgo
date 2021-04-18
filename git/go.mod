module github.com/kamontat/gitgo/git

go 1.16

replace github.com/kamontat/gitgo/utils v0.0.0-local => ../utils

require (
	github.com/go-git/go-billy/v5 v5.0.0
	github.com/go-git/go-git/v5 v5.2.0
	github.com/kamontat/gitgo/utils v0.0.0-local
)

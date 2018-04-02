package flag

import "github.com/urfave/cli"

var repo string

func GetRepository() string {
	if repo == "" {
		return "origin"
	}
	return repo
}

func CustomRepoFlag() cli.Flag {
	return cli.StringFlag{
		Name:        "repository, repo, r",
		Usage:       "change default `REPO`, default is 'origin'",
		Destination: &repo,
	}
}

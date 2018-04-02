package flag

import "github.com/urfave/cli"

var branch string

func GetBranch() string {
	if branch == "" {
		return "master"
	}
	return branch
}

func CustomBranchFlag() cli.Flag {
	return cli.StringFlag{
		Name:        "branch, b",
		Usage:       "change default `BRANCH`, default is 'master'",
		Destination: &branch,
	}
}

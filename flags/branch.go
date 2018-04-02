package flag

import (
	"strings"

	"github.com/urfave/cli"
)

var branch string

func GetBranchs() []string {
	if branch == "" {
		return []string{"master"}
	}
	return strings.Split(branch, " ")
}

func CustomBranchFlag() cli.Flag {
	return cli.StringFlag{
		Name:        "branch, b",
		Usage:       "change default `BRANCH`, default is 'master'",
		Destination: &branch,
	}
}

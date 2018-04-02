package flag

import "github.com/urfave/cli"

var isAll = false

func IsAll() bool {
	return isAll
}

func AllFlag() cli.Flag {
	return cli.BoolFlag{
		Name:        "all, a, A",
		Usage:       "get all",
		Destination: &isAll,
	}
}

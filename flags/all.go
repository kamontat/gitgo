package flag

import "github.com/urfave/cli"

var isAll = false

func IsAll() bool {
	return isAll
}

func AllFlag() cli.Flag {
	return cli.BoolFlag{
		Name:        "all, A",
		Usage:       "get all",
		Destination: &isAll,
	}
}

func AllFlagCustom(msg string) cli.Flag {
	return cli.BoolFlag{
		Name:        "all, A",
		Usage:       msg,
		Destination: &isAll,
	}
}

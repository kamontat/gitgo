package flag

import "github.com/urfave/cli"

var add = false

func IsNeedAdd() bool {
	return add
}

func AddAddFlag() cli.Flag {
	return cli.BoolFlag{
		Name:        "add, a",
		Usage:       "include add flag into commit",
		Destination: &add,
	}
}

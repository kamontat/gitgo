package flag

import "github.com/urfave/cli"

var isList = false

func NeedToListVersion() bool {
	return isList
}

func ListVersionFlag() cli.Flag {
	return cli.BoolFlag{
		Name:        "list-version, L",
		Usage:       "list every exist versions",
		Destination: &isList,
	}
}

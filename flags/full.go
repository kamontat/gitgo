package flag

import "github.com/urfave/cli"

var isFull = false

func IsFull() bool {
	return isFull
}

func FullFlag() cli.Flag {
	return cli.BoolFlag{
		Name:        "full, F",
		Usage:       "show full output",
		Destination: &isFull,
	}
}

package flag

import "github.com/urfave/cli"

var auto = false

func IsAuto() bool {
	return auto
}

func AddAutoFlag() cli.Flag {
	return cli.BoolFlag{
		Name:        "auto, t",
		Usage:       "auto choose, if prompt is exist",
		Destination: &auto,
	}
}

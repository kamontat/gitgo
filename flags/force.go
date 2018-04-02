package flag

import "github.com/urfave/cli"

var isForce = false

func IsForce() bool {
	return isForce
}

func ForceFlag() cli.Flag {
	return cli.BoolFlag{
		Name:        "force, f",
		Usage:       "force to do something, default: false",
		Destination: &isForce,
	}
}

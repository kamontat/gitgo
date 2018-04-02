package flag

import (
	"fmt"

	"github.com/urfave/cli"
)

var isForce = false

func IsForce() bool {
	return isForce
}

func ForceFlag(use string) cli.Flag {
	return cli.BoolFlag{
		Name:        "force, f",
		Usage:       fmt.Sprintf("force to %s, default: false", use),
		Destination: &isForce,
	}
}

package flag

import "github.com/urfave/cli"

func FileAndFolderFlag() cli.Flag {
	return cli.BoolFlag{
		Name:        "force, f",
		Usage:       "force to do something, default: false",
		Destination: &isForce,
	}
}

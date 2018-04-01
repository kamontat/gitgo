package global

import "github.com/urfave/cli"

var isForce = false

func IsForce() bool {
	return isForce
}

// type isForce bool

func AddForceAsGlobalFlag(app *cli.App) {
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "force, f",
			Usage:       "force to do something, default: false",
			Destination: &isForce,
		},
	}
}

package command

import (
	"gitgo/models"

	"github.com/urfave/cli"
)

func AddVersion(appConfig models.AppConfig) cli.Command {
	var full bool
	return cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "show version, same as --version",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:        "full, F",
				Usage:       "show full output",
				Destination: &full,
			},
		},
		Action: func(c *cli.Context) error {
			if full {
				appConfig.LatestVersion().PrintFullVersion(appConfig.Name)
			} else {
				appConfig.LatestVersion().PrintVersion(appConfig.Name)
			}
			return nil
		},
	}
}

func AddListVersion(appConfig models.AppConfig) cli.Command {
	var full bool
	return cli.Command{
		Name:    "list-version",
		Aliases: []string{"L"},
		Usage:   "list every version, same as --list-version",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:        "full, F",
				Usage:       "show full output",
				Destination: &full,
			},
		},
		Action: func(c *cli.Context) error {
			if full {
				appConfig.PrintFullEveryVersions()
			} else {
				appConfig.PrintEveryVersions()
			}
			return nil
		},
	}
}

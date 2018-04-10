package command

import (
	flag "github.com/kamontat/gitgo/flags"
	"github.com/kamontat/gitgo/models"

	"github.com/urfave/cli"
)

// AddVersion add cli command of show version
func AddVersion(appConfig models.AppConfig) cli.Command {
	return cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "show version, same as --version",
		Flags: []cli.Flag{
			flag.FullFlag(),
		},
		Action: func(c *cli.Context) error {
			appConfig.LatestVersion().ChooseToPrintVersion(flag.IsFull())
			return nil
		},
	}
}

// AddListVersion add cli command of show every versions
func AddListVersion(appConfig models.AppConfig) cli.Command {
	return cli.Command{
		Name:    "list-version",
		Aliases: []string{"L"},
		Usage:   "list every version, same as --list-version",
		Flags: []cli.Flag{
			flag.FullFlag(),
		},
		Action: func(c *cli.Context) error {
			appConfig.ChooseToPrintEveryVersions(flag.IsFull())
			return nil
		},
	}
}

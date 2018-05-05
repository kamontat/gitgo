package command

import (
	"fmt"

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
			flag.AllFlag(),
		},
		Action: func(c *cli.Context) error {
			if flag.IsAll() {
				fmt.Println(appConfig.GetVersionLong(0))
			} else {
				fmt.Println(appConfig.GetVersionShort(0))
			}
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
			flag.AllFlag(),
		},
		Action: func(c *cli.Context) error {
			if flag.IsAll() {
				appConfig.PrintAllVersionLong()
			} else {
				appConfig.PrintAllVersionShort()
			}
			return nil
		},
	}
}

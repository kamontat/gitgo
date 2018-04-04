package command

import (
	"github.com/urfave/cli"
)

// AddConfig add command of setting(s)
func AddConfig() cli.Command {
	return cli.Command{
		Name:      "configuration",
		Aliases:   []string{"config", "g"},
		Category:  "Setting",
		Usage:     "Get config commands",
		UsageText: "gitgo config|g ",
		Subcommands: []cli.Command{
			AddConfigLocation(),
		},
		Action: func(c *cli.Context) error {
			return nil
		},
	}
}

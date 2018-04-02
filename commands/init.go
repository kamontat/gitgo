package command

import (
	"gitgo/flags"

	"gitgo/client"

	"github.com/urfave/cli"
)

func InitGit() cli.Command {
	return cli.Command{
		Name:      "init",
		Aliases:   []string{"i"},
		Category:  "Setting",
		Usage:     "Inital git",
		UsageText: "gitgo init|i [--force|-f]",
		Flags: []cli.Flag{
			flag.ForceFlag("initial git"),
		},
		Action: func(c *cli.Context) error {
			if client.GitIsNotInit() || flag.IsForce() {
				client.GitInit()
			} else {
				return cli.NewExitError("Initial already!, GitAdd --force", 4)
			}
			return nil
		},
	}
}

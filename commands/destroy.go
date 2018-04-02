package command

import (
	"gitgo/client"

	"github.com/urfave/cli"
)

func DestroyGit() cli.Command {
	return cli.Command{
		Name:     "destroy",
		Aliases:  []string{"d"},
		Category: "Setting",
		Usage:    "Destroy git",
		Action: func(c *cli.Context) error {
			if client.GitIsInit() {
				client.GitDelete()
			} else {
				return cli.NewExitError("Never initial!", 4)
			}
			return nil
		},
	}
}

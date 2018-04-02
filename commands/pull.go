package command

import (
	"gitgo/flags"

	"gitgo/client"

	"github.com/urfave/cli"
)

func PullGit() cli.Command {
	return cli.Command{
		Name:     "pull",
		Aliases:  []string{"P"},
		Category: "Server",
		Usage:    "pull server git to local",
		Flags: []cli.Flag{
			flag.ForceFlag("pull server code"),
		},
		Action: func(c *cli.Context) error {
			if client.GitIsInit() {
				client.GitInit()
			} else {
				return cli.NewExitError("Never initial!", 4)
			}
			return nil
		},
	}
}

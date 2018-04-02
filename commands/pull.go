package command

import (
	"gitgo/flags"

	"gitgo/client"

	"github.com/urfave/cli"
)

func PullGit() cli.Command {
	return cli.Command{
		Name:      "pull",
		Aliases:   []string{"P"},
		Category:  "Server",
		Usage:     "pull server git to local",
		UsageText: "gitgo pull|P [--force|-f] [--repo|-r <repo>] [<branch>]",
		Flags: []cli.Flag{
			flag.ForceFlag("pull server code"),
		},
		Action: func(c *cli.Context) error {
			if client.GitIsNotInit() {
				return cli.NewExitError("Never initial!", 4)
			}
			if client.GitDontHaveRemote() {
				return cli.NewExitError("Never set git remote!", 4)
			}
			err := client.GitPull(flag.IsForce(), flag.GetRepository(), c.Args().Tail())
			if err != nil {
				return cli.NewExitError(err, 4)
			}
			return nil
		},
	}
}

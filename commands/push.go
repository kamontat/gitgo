package command

import (
	"github.com/kamontat/gitgo/flags"

	"github.com/kamontat/gitgo/client"

	"github.com/urfave/cli"
)

// PushGit will generate cli command of 'git push'
func PushGit() cli.Command {
	return cli.Command{
		Name:      "push",
		Aliases:   []string{"p"},
		Category:  "Server",
		Usage:     "push local git to server",
		UsageText: "gitgo push|p [--force|-f] [--repo|-r <repo>] [<branch>]",
		Flags: []cli.Flag{
			flag.ForceFlag("push code to server"),
			flag.CustomRepoFlag(),
		},
		Subcommands: []cli.Command{
			AddSetPush(),
			AddDeployment(),
		},
		Action: func(c *cli.Context) error {
			if client.GitIsNotInit() {
				return cli.NewExitError("Never initial!", 4)
			}
			if client.GitDontHaveRemote() {
				return cli.NewExitError("Never set git remote!", 4)
			}

			err := client.GitPush(flag.IsForce(), flag.GetRepository(), c.Args())
			if err != nil {
				return cli.NewExitError(err, 4)
			}
			return nil
		},
	}
}

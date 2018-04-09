package command

import (
	"github.com/kamontat/gitgo/client"
	flag "github.com/kamontat/gitgo/flags"

	"github.com/urfave/cli"
)

// AddSetPush will generate cli command of 'git remote add' and 'git push -u'
func AddSetPush() cli.Command {
	return cli.Command{
		Name:      "set",
		Aliases:   []string{"s"},
		Usage:     "set push server and remote, and push code",
		UsageText: "gitgo push|p set|s [--force|-f] [--repo|-r <repo>] [--branch|-b <branch>] <link>",
		Flags: []cli.Flag{
			flag.ForceFlag("setup and push code"),
			flag.CustomRepoFlag(),
			flag.CustomBranchFlag(),
		},
		Action: func(c *cli.Context) error {
			// error not inital
			if client.GitIsNotInit() {
				return cli.NewExitError("Never initial!", 4)
			}
			// remote exist, must have --force to continue
			if client.GitHasRemote(flag.GetRepository()) {
				if !flag.IsForce() {
					return cli.NewExitError("Remote exist!", 4)
				}
				client.GitReAddRemote(flag.GetRepository(), c.Args().First())
			} else {
				client.GitAddRemote(flag.GetRepository(), c.Args().First())
			}

			err := client.GitSetupPush(flag.IsForce(), true, flag.GetRepository(), flag.GetBranchs())
			if err != nil {
				return cli.NewExitError(err, 4)
			}
			return nil
		},
	}
}

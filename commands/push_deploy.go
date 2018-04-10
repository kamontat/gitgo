package command

import (
	"github.com/kamontat/gitgo/client"
	flag "github.com/kamontat/gitgo/flags"

	"github.com/urfave/cli"
)

// AddDeployment will generate cli command of deploy git to server
func AddDeployment() cli.Command {
	return cli.Command{
		Name:      "deployment",
		Aliases:   []string{"deploy", "d"},
		Usage:     "deploy current code, should use after 'gitgo commit release'",
		UsageText: "gitgo push|p deploy|d",
		// Flags:     []cli.Flag{
		// flag.ForceFlag("setup and push code"),
		// flag.CustomRepoFlag(),
		// flag.CustomBranchFlag(),
		// },
		Action: func(c *cli.Context) error {
			// error not inital
			if client.GitIsNotInit() {
				return cli.NewExitError("Never initial!", 4)
			}
			// error remote not exist
			if !client.GitHasRemote(flag.GetRepository()) {
				return cli.NewExitError("Remote not exist!", 4)
			}

			return cli.NewExitError("NOT implement yet!", 199)
			// err := client.GitSetupPush(flag.IsForce(), true, flag.GetRepository(), flag.GetBranchs())
			// if err != nil {

			// }
			// return nil
		},
	}
}

package command

import (
	"github.com/kamontat/gitgo/client"
	flag "github.com/kamontat/gitgo/flags"

	"github.com/urfave/cli"
)

// AddCommitInital add command of 'commit initial'
func AddCommitInital(emoji bool) cli.Command {
	return cli.Command{
		Name:      "initial",
		Aliases:   []string{"i", "init"},
		Usage:     "Create default initial commit",
		UsageText: "gitgo commit|cm|c init|i",
		Flags: []cli.Flag{
			flag.AllFlagCustom("add all before commit"),
		},
		Action: func(c *cli.Context) error {
			if client.GitIsNotInit() {
				return cli.NewExitError("Never initial!", 4)
			}

			if flag.IsAll() {
				client.GitAddAll()
			}

			return client.BypassCommit(emoji, "init")
		},
	}
}

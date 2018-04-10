package command

import (
	"github.com/kamontat/gitgo/client"

	"github.com/urfave/cli"
)

// AddCommitRelease add command of release version commit
func AddCommitRelease(emoji bool) cli.Command {
	return cli.Command{
		Name:      "release",
		Aliases:   []string{"r"},
		Usage:     "Create release commit",
		UsageText: "gitgo commit|cm|c release|r",
		Action: func(c *cli.Context) error {
			if client.GitIsNotInit() {
				return cli.NewExitError("Never initial!", 4)
			}

			return cli.NewExitError("NOT implement yet!", 199)
			// return client.BypassInitialCommit(emoji, "init")
		},
	}
}

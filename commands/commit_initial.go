package command

import (
	"github.com/kamontat/gitgo/client"

	"github.com/urfave/cli"
)

// AddCommitInital add command of 'commit initial'
func AddCommitInital(emoji bool) cli.Command {
	return cli.Command{
		Name:      "initial",
		Aliases:   []string{"i", "init"},
		Usage:     "Create default initial commit",
		UsageText: "gitgo commit|cm|c init|i",
		Action: func(c *cli.Context) error {
			if emoji {
				// fmt.Print("create as emoji")
				client.MakeGitCommitWithEmoji(true, "🎉", "Initial commit")
			} else {
				// fmt.Print("create as test")
				client.MakeGitCommitWithText(true, "init", "Initial commit")
			}
			return nil
		},
	}
}

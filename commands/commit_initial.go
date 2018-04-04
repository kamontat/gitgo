package command

import (
	"gitgo/client"

	"github.com/urfave/cli"
)

func AddCommitInital(emoji bool) cli.Command {
	return cli.Command{
		Name:    "initial",
		Aliases: []string{"i", "init"},
		Action: func(c *cli.Context) error {
			if emoji {
				client.MakeGitCommitWithEmoji(true, "🎉", "Initial commit")
			} else {
				client.MakeGitCommitWithText(true, "init", "Initial commit")
			}
			return nil
		},
	}
}

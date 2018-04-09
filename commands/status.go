package command

import (
	"github.com/kamontat/gitgo/client"

	"github.com/urfave/cli"
)

// AddGitStatus add command of 'git status'
func AddGitStatus() cli.Command {
	return cli.Command{
		Name:      "status",
		Aliases:   []string{"s"},
		Category:  "Setting",
		Usage:     "Show status of the git",
		UsageText: "gitgo status|s",
		Action: func(c *cli.Context) error {
			if client.GitIsNotInit() {
				return cli.NewExitError("Never initial git!", 4)
			}
			client.Status()
			return nil
		},
	}
}

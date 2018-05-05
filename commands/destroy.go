package command

import (
	"fmt"

	"github.com/kamontat/gitgo/client"
	flag "github.com/kamontat/gitgo/flags"
	survey "gopkg.in/AlecAivazis/survey.v1"

	"github.com/urfave/cli"
)

// DestroyGit will generate cli command of 'rm -rf .git'
func DestroyGit() cli.Command {
	return cli.Command{
		Name:      "destroy",
		Aliases:   []string{"d"},
		Category:  "Setting",
		Usage:     "Destroy git",
		UsageText: "gitgo destroy|d [--force|-f]",
		Flags: []cli.Flag{
			flag.ForceFlag("delete git without prompt"),
		},
		Action: func(c *cli.Context) (err error) {
			if client.GitIsInit() {
				if flag.IsForce() {
					return client.GitDelete()
				}
				result := false
				prompt := &survey.Confirm{
					Message: "Do you sure?",
					Help:    "git directory will be deleted by 'rm -r' command",
				}
				survey.AskOne(prompt, &result, nil)
				if result {
					return client.GitDelete()
				}
				fmt.Println("Cancel delete git!")
			} else {
				return cli.NewExitError("Never initial!", 4)
			}
			return nil
		},
	}
}

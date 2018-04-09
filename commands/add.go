package command

import (
	client "github.com/kamontat/gitgo/client"
	flag "github.com/kamontat/gitgo/flags"

	"github.com/urfave/cli"
)

func _addAll() cli.Command {
	return cli.Command{
		Name:      "all",
		Aliases:   []string{"a"},
		Usage:     "GitAdd every files to git",
		UsageText: "gitgo add|a all|a",
		Action: func(c *cli.Context) error {
			if client.GitIsNotInit() {
				return cli.NewExitError("Never initial git", 5)
			}
			client.GitAddAll()
			return nil
		},
	}
}

// AddGit will generate cli command of 'git add'
func AddGit() cli.Command {
	return cli.Command{
		Name:      "add",
		Aliases:   []string{"a"},
		Category:  "Action",
		Usage:     "Add file/folder to git",
		HelpName:  "gitgo add|a [--force|-f] [--all|-a]",
		ArgsUsage: "[<files|folder>]",
		Flags: []cli.Flag{
			flag.AllFlag(),
			// flag.FileAndFolderFlag(),
		},
		Action: func(c *cli.Context) error {
			if client.GitIsNotInit() {
				return cli.NewExitError("Never initial git", 5)
			}

			if flag.IsAll() {
				client.GitAddAll()
			} else {
				if c.NArg() == 0 {
					return cli.NewExitError("no args exist, add must have argument", 5)
				}
				client.GitAdd(c.Args()...)
			}
			return nil
		},
		Subcommands: []cli.Command{
			_addAll(),
		},
	}
}

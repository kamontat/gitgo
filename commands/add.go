package command

import (
	client "gitgo/client"
	flag "gitgo/flags"

	"github.com/urfave/cli"
)

func _addAll() cli.Command {
	return cli.Command{
		Name:      "all",
		Aliases:   []string{"a"},
		Usage:     "GitAdd every files to git",
		UsageText: "gitgo add|a all|a",
		Action: func(c *cli.Context) error {
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
			if flag.IsAll() {
				client.GitAddAll()
			} else {
				client.GitAdd(c.Args()...)
			}
			return nil
		},
		Subcommands: []cli.Command{
			_addAll(),
		},
	}
}

package command

import (
	"fmt"
	client "gitgo/client"
	flag "gitgo/flags"

	"github.com/urfave/cli"
)

func _addAll() cli.Command {
	return cli.Command{
		Name:    "all",
		Aliases: []string{"a"},
		Usage:   "GitAdd every files to git",
		Action: func(c *cli.Context) error {
			client.GitAddAll()
			fmt.Println("GitAdd all")
			return nil
		},
	}
}

func AddGit() cli.Command {
	return cli.Command{
		Name:      "add",
		Aliases:   []string{"a"},
		Category:  "Action",
		Usage:     "GitAdd file/folder to git",
		UsageText: "gitgo add [all|-all]",
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

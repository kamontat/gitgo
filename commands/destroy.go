package command

import (
	"fmt"
	"gitgo/client"
	flag "gitgo/flags"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
)

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
		Action: func(c *cli.Context) error {
			if client.GitIsInit() {
				if flag.IsForce() {
					client.GitDelete()
				} else {
					prompt := promptui.Prompt{
						Label:     "Delete git",
						IsVimMode: true,
						IsConfirm: true,
					}
					result, err := prompt.Run()
					if err == nil && result == "y" {
						client.GitDelete()
					} else {
						fmt.Println("Cancel delete git!")
					}
				}
			} else {
				return cli.NewExitError("Never initial!", 4)
			}
			return nil
		},
	}
}

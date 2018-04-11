package command

import (
	"fmt"

	"github.com/kamontat/gitgo/client"
	flag "github.com/kamontat/gitgo/flags"
	"github.com/manifoldco/promptui"

	"github.com/urfave/cli"
)

func promptTag(tag string) error {
	prompt := promptui.Prompt{
		Label:     "Auto create tag name: " + tag,
		IsVimMode: true,
		IsConfirm: true,
	}
	result, err := prompt.Run()
	if err == nil && result == "y" {
		return client.SetTag(tag)
	}
	fmt.Println("Not create new tag")
	return nil
}

// AddCommitRelease add command of release version commit
func AddCommitRelease(emoji bool) cli.Command {
	return cli.Command{
		Name:      "release",
		Aliases:   []string{"r"},
		Usage:     "Create release commit",
		UsageText: "gitgo commit|cm|c release|r <version_tag> [--auto]",
		Flags: []cli.Flag{
			flag.AddAutoFlag(),
		},
		Action: func(c *cli.Context) (err error) {
			if client.GitIsNotInit() {
				return cli.NewExitError("Never initial!", 4)
			}
			tag := c.Args().First()
			if tag == "" {
				return cli.NewExitError("input tag MUST be exist!", 5)
			}
			err = client.BypassCommit(emoji, "release", tag)
			if flag.IsAuto() {
				return client.SetTag(tag)
			}
			return promptTag(tag)
		},
	}
}

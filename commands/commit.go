package command

import (
	"os"

	"github.com/kamontat/gitgo/client"
	"github.com/kamontat/gitgo/flags"
	"github.com/kamontat/gitgo/models"

	"github.com/urfave/cli"
)

func commitAsText() cli.Command {
	return cli.Command{
		Name:      "text",
		Aliases:   []string{"t"},
		Usage:     "Commit as text key",
		UsageText: "gitgo commit|cm|c text|t [--add|-a] [--all|-A] [--key|k <key>] [--title|t <title>] [<message>]",
		Flags: []cli.Flag{
			flag.AddAddFlag(),
			flag.AllFlagCustom("add all before commit"),
			flag.AddTitleFlag("text commit"),
			flag.AddKeyFlag("text commit"),
		},
		Subcommands: []cli.Command{
			AddCommitInital(false),
			AddCommitRelease(false),
		},
		Action: func(c *cli.Context) error {
			if client.GitIsNotInit() {
				return cli.NewExitError("Never initial git", 3)
			}
			if flag.IsAll() {
				client.GitAddAll()
			}

			return client.MakeGitCommitWithText(flag.IsNeedAdd(), flag.GetKey(), flag.GetTitle(), c.Args()...)
		},
	}
}

func commitAsEmoji() cli.Command {
	return cli.Command{
		Name:      "emoji",
		Aliases:   []string{"moji", "e"},
		Usage:     "Commit as emoji key",
		UsageText: "gitgo commit|cm|c emoji|moji|e [--add|-a] [--all|-A] [--key|k <key>] [--title|t <title>] [<message>]",
		Flags: []cli.Flag{
			flag.AddAddFlag(),
			flag.AllFlagCustom("add all before commit"),
			flag.AddTitleFlag("emoji commit"),
			flag.AddKeyFlag("emoji commit"),
		},
		Subcommands: []cli.Command{
			AddCommitInital(true),
			AddCommitRelease(true),
		},
		Action: func(c *cli.Context) error {
			if client.GitIsNotInit() {
				return cli.NewExitError("Never initial git", 3)
			}

			if flag.IsAll() {
				client.GitAddAll()
			}

			return client.MakeGitCommitWithEmoji(flag.IsNeedAdd(), flag.GetKey(), flag.GetTitle(), c.Args()...)
		},
	}
}

// CommitGit will generate cli command of 'git commit'
func CommitGit() cli.Command {
	return cli.Command{
		Name:      "commit",
		Aliases:   []string{"cm", "c"},
		Category:  "Action",
		Usage:     "Commit changes",
		UsageText: "gitgo commit [--add|-a] [--all|-A] [subcommand] [--key|k <key>] [--title|t <title>] [<message>]",
		Flags: []cli.Flag{
			flag.AddAddFlag(),
			flag.AllFlagCustom("add all before commit"),
			flag.AddKeyFlag("commit"),
			flag.AddTitleFlag("commit"),
		},
		Subcommands: []cli.Command{
			commitAsText(),
			commitAsEmoji(),
			AddCommitInital(models.GetUserConfig().IsEmojiType()),
			AddCommitRelease(models.GetUserConfig().IsEmojiType()),
		},
		Action: func(c *cli.Context) error {
			if client.GitIsNotInit() {
				return cli.NewExitError("Never initial git", 3)
			}
			commit := models.GetUserConfig().Config.Commit.Type
			if commit == "" {
				commit = os.Getenv("COMMIT_TYPE")
			}
			if commit == "" {
				commit = os.Getenv("COMMITTYPE")
			}
			if commit == "" {
				return cli.NewExitError("COMMIT_TYPE not exist, please call commit with subcommand instead.", 5)
			}

			if flag.IsAll() {
				client.GitAddAll()
			}

			if commit == "e" || commit == "emoji" || commit == "moji" {
				return client.MakeGitCommitWithEmoji(flag.IsNeedAdd(), flag.GetKey(), flag.GetTitle(), c.Args()...)
			} else if commit == "text" || commit == "t" {
				return client.MakeGitCommitWithText(flag.IsNeedAdd(), flag.GetKey(), flag.GetTitle(), c.Args()...)
			} else {
				return cli.NewExitError("COMMIT_TYPE must be 'text' or 'emoji'", 5)
			}
		},
	}
}

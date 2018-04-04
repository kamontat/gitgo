package command

import (
	"gitgo/client"
	"gitgo/flags"
	"gitgo/models"
	"os"

	"github.com/urfave/cli"
)

func commitAsText() cli.Command {
	return cli.Command{
		Name:      "text",
		Aliases:   []string{"t"},
		Usage:     "Commit as text key",
		UsageText: "gitgo commit|cm|c text|t [--key|k <key>] [--title|t <title>] [<message>]",
		Flags: []cli.Flag{
			flag.AddTitleFlag("text commit"),
			flag.AddKeyFlag("text commit"),
		},
		Action: func(c *cli.Context) error {
			client.MakeGitCommitWithText(flag.GetKey(), flag.GetTitle(), c.Args().First())
			return nil
		},
	}
}

func commitAsEmoji() cli.Command {
	return cli.Command{
		Name:      "emoji",
		Aliases:   []string{"moji", "e"},
		Usage:     "Commit as emoji key",
		UsageText: "gitgo commit|cm|c emoji|moji|e [--key|k <key>] [--title|t <title>] [<message>]",
		Flags: []cli.Flag{
			flag.AddTitleFlag("emoji commit"),
			flag.AddKeyFlag("emoji commit"),
		},
		Action: func(c *cli.Context) error {
			client.MakeGitCommitWithEmoji(flag.GetKey(), flag.GetTitle(), c.Args().First())
			return nil
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
		UsageText: "gitgo commit|cm|c [text|emoji] [--key|k <key>] [--title|t <title>] [<message>]",
		Flags: []cli.Flag{
			flag.AddKeyFlag("commit"),
			flag.AddTitleFlag("commit"),
		},
		Subcommands: []cli.Command{
			commitAsText(),
			commitAsEmoji(),
		},
		Action: func(c *cli.Context) error {
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

			if commit == "e" || commit == "emoji" || commit == "moji" {
				client.MakeGitCommitWithEmoji(flag.GetKey(), flag.GetTitle(), c.Args().First())
			} else if commit == "text" || commit == "t" {
				client.MakeGitCommitWithText(flag.GetKey(), flag.GetTitle(), c.Args().First())
			} else {
				return cli.NewExitError("COMMIT_TYPE must be 'text' or 'emoji'", 5)
			}
			// if client.GitIsNotInit() || flag.IsForce() {
			// 	client.GitInit()
			// } else {
			// 	return cli.NewExitError("Initial already!, GitAdd --force", 4)
			// }
			return nil
		},
	}
}

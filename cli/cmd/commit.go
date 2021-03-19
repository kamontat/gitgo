package cmd

import (
	"github.com/kamontat/gitgo/git"
	"github.com/kamontat/gitgo/prompt"
	"github.com/kamontat/gitgo/utils/phase"

	"github.com/spf13/cobra"
)

type CommitOption struct {
	enabledMessage bool
	dryrun         bool
}

var commitOption = &CommitOption{}

var commitCmd = &cobra.Command{
	Use:     "commit",
	Aliases: []string{"c"},
	Short:   "commit management",
	Run: func(cmd *cobra.Command, args []string) {
		phase.OnCommandPhase()

		if commitOption.enabledMessage {
			phase.Info("Overrided: enabled message")
			configuration.Settings.Commit.EnabledMessage()
		}

		if commitOption.dryrun {
			phase.Info("Dryrun")
		}

		repo, err := git.New(pwdPath)
		phase.Error(err)

		phase.Debug("initial repository: %s", repo.Path())

		msg, err := prompt.CommitMessage(configuration.Settings.Commit)
		phase.Error(err)

		if !commitOption.dryrun {
			hash, err := repo.Commit(&msg)
			phase.Error(err)

			phase.Info(hash)
		} else {
			message, err := msg.Formatted()
			phase.Error(err)

			phase.Info("git commit -m '%s'", message)
		}
	},
}

func init() {
	root.AddCommand(commitCmd)

	commitCmd.Flags().BoolVarP(&commitOption.enabledMessage, "message", "M", false, "override enabled message key on config file")
	commitCmd.Flags().BoolVarP(&commitOption.dryrun, "dry-run", "D", false, "not create any commit to current worktree")
}

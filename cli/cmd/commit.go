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
	allowEmpty     bool
}

var commitOption = &CommitOption{}

var commitCmd = &cobra.Command{
	Use:     "commit",
	Aliases: []string{"c"},
	Short:   "commit management",
	//lint:ignore SA4009 we don't care what is user arguments
	Run: func(cmd *cobra.Command, args []string) {
		phase.OnCommandPhase()

		if commitOption.enabledMessage {
			phase.Info("Overrided: enabled message")
			configuration.Settings.Commit.EnabledMessage()
		}

		repo, err := git.New(configOption.WdPath)
		phase.Error(err)

		phase.Debug("initial repository: %s", repo.Path())

		msg, err := prompt.CommitMessage(configuration.Settings.Commit)
		phase.Error(err)

		// TODO: change whether has files staged or not
		// this code only check whether has files modified or not
		// if !commitOption.allowEmpty && repo.IsClean() {
		// 	phase.Error(errors.New("cannot create commit because it not allow empty"))
		// }

		if !commitOption.dryrun {
			// This should go away if go-git support sign auto
			if configuration.Settings.Hack {
				phase.Debug("run git commit with hack mode")

				args = make([]string, 0)
				if commitOption.allowEmpty {
					args = append(args, "--allow-empty")
				}

				_, err := repo.HackCommit(&msg, args...)
				phase.Error(err)
			} else {
				hash, err := repo.Commit(&msg)
				phase.Error(err)

				phase.Info(hash)
			}
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
	commitCmd.Flags().BoolVarP(&commitOption.allowEmpty, "empty", "E", false, "create whether empty tree or not")
}

package cmd

import (
	"github.com/kamontat/gitgo/git"
	"github.com/kamontat/gitgo/git/constants"
	"github.com/kamontat/gitgo/utils/phase"

	"github.com/spf13/cobra"
)

var commitInitCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i"},
	Short:   "create initial project commit",
	Run: func(cmd *cobra.Command, args []string) {
		phase.OnCommandPhase()

		repo, err := git.New(pwdPath)
		phase.Error(err)

		phase.Debug("initial repository: %s", repo.Path())

		hash, err := repo.Commit(constants.InitialCommitMessage)
		phase.Error(err)

		phase.Info(hash)
	},
}

func init() {
	commitCmd.AddCommand(commitInitCmd)
}

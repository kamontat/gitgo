package cmd

import (
	"github.com/kamontat/gitgo/git"
	"github.com/kamontat/gitgo/git/constants"
	"github.com/kamontat/gitgo/utils/phase"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:     "new",
	Aliases: []string{"N"},
	Short:   "create new git repository",
	Run: func(cmd *cobra.Command, args []string) {
		phase.OnCommandPhase()

		// git init
		repo, err := git.Create(configOption.WdPath)
		phase.Error(err)
		phase.Debug("initial repository: %s", repo.Path())

		hash, err := repo.Commit(constants.InitialCommitMessage)
		phase.Error(err)
		phase.Debug("initial with commit: %s", hash)
	},
}

func init() {
	root.AddCommand(newCmd)
}

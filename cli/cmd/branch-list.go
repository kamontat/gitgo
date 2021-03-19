package cmd

import (
	"github.com/kamontat/gitgo/git"
	"github.com/kamontat/gitgo/utils/phase"
	"github.com/spf13/cobra"
)

var branchListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List all local branch",
	Run: func(cmd *cobra.Command, args []string) {
		phase.OnCommandPhase()

		repo, err := git.New(configOption.WdPath)
		phase.Error(err)

		phase.Debug("initial repository: ", repo.Path())

		for i, branch := range repo.Branches() {
			phase.Format("%02d: '%s'", i+1, branch)
		}
	},
}

func init() {
	branchCmd.AddCommand(branchListCmd)
}

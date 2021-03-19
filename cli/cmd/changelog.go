package cmd

import (
	"github.com/kamontat/gitgo/git"
	"github.com/kamontat/gitgo/utils/phase"
	"github.com/spf13/cobra"
)

var changelogCmd = &cobra.Command{
	Use:     "changelog",
	Aliases: []string{"cl"},
	Short:   "generate changelog file",
	Run: func(cmd *cobra.Command, args []string) {
		phase.OnCommandPhase()

		repo, err := git.New(configOption.WdPath)
		phase.Error(err)

		phase.Debug("initial repository: %s", repo.Path())

		changelog, err := repo.Changelog(&git.ChangelogOption{})
		phase.Error(err)

		phase.Info("%s", changelog)
	},
}

func init() {
	root.AddCommand(changelogCmd)
}

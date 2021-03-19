package cmd

import (
	"github.com/kamontat/gitgo/utils/phase"
	"github.com/spf13/cobra"
)

var configPathCmd = &cobra.Command{
	Use:     "path",
	Aliases: []string{"p"},
	Short:   "show current used config file path",
	Run: func(cmd *cobra.Command, args []string) {
		phase.OnCommandPhase()
		phase.Log(configPath)
	},
}

func init() {
	configCmd.AddCommand(configPathCmd)
}

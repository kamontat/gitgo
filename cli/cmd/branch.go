package cmd

import (
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:     "branch",
	Aliases: []string{"b"},
	Short:   "branch management",
}

func init() {
	root.AddCommand(branchCmd)
}

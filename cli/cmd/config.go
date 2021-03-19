package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"C"},
	Short:   "Configuration management",
}

func init() {
	root.AddCommand(configCmd)
}

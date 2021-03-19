package cmd

import (
	"github.com/kamontat/gitgo/utils/phase"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var force bool

var configInitCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i"},
	Short:   "Initital new config file on current directory",
	Run: func(cmd *cobra.Command, args []string) {
		phase.OnCommandPhase()

		phase.Debug("starting create new config")
		phase.Debug("create config at %s", configPath)
		if force {
			phase.Debug("start with force mode")
		}

		viper.Set("settings", configuration.Settings)
		viper.Set("version", configuration.Version)

		var err error
		if force {
			err = viper.WriteConfig()
			phase.Error(err)
		} else {
			err = viper.SafeWriteConfig()
			phase.Warn(err)
		}
	},
}

func init() {
	configCmd.AddCommand(configInitCmd)
	configInitCmd.Flags().BoolVarP(&force, "force", "F", false, "Force create config even it exist")
}

package cmd

import (
	"os"

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

		// Update viper instance
		viper.Set("settings", configuration.Settings)
		viper.Set("version", configuration.Version)

		// Log debug information
		phase.Debug("starting create new config")
		phase.Debug("create config at %s", configOption.GetConfigPath())
		if force {
			phase.Debug("start with force mode")
		}

		var err error

		err = os.MkdirAll(configOption.Setting.DefaultDirectoryPath(), os.ModePerm)
		phase.Warn(err)

		if force {
			err = viper.WriteConfigAs(configOption.GetConfigPath())
			phase.Error(err)
		} else {
			err = viper.SafeWriteConfigAs(configOption.GetConfigPath())
			phase.Warn(err)
		}
	},
}

func init() {
	configCmd.AddCommand(configInitCmd)
	configInitCmd.Flags().BoolVarP(&force, "force", "F", false, "Force create config even it exist")
}

package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/kamontat/gitgo/cli/helpers"
	"github.com/kamontat/gitgo/config"
	"github.com/kamontat/gitgo/config/constants"
	"github.com/kamontat/gitgo/config/models"
	"github.com/kamontat/gitgo/core"
	"github.com/kamontat/gitgo/utils"
	"github.com/kamontat/gitgo/utils/logger"
	"github.com/kamontat/gitgo/utils/phase"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configuration models.Configuration = *config.Default()
var configPath string
var pwdPath string

var root = &cobra.Command{
	Use:   "gitgo",
	Short: "gitgo command for organized git-flow",
	Long: `# Gitgo

This command create by golang with cobra cli.
Force everyone to create the exect same templates of commit and branch

v5.0.0-alpha.1 - create new project with new golang version (1.10 -> 1.16)

Motivated by gitmoji and GitFlow.`,
	Version: core.Version,
}

func initConfig() {
	phase.OnInitialPhase()

	var err error
	pwdPath, err = os.Getwd()
	if err != nil {
		phase.Error(err)
	}

	directories, errs := helpers.ListConfigDirectories()

	for _, dir := range directories {
		viper.AddConfigPath(dir)
	}

	for _, err := range errs {
		phase.Warn(err)
	}

	viper.AddConfigPath(".gitgo")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	if !viper.GetBool(constants.SettingDisabledConfig) {
		// read configuration from files
		err = viper.ReadInConfig()
		switch t := err.(type) {
		case viper.ConfigFileNotFoundError:
			phase.Warn(errors.New("config file not found"))
		default:
			phase.Warn(t)
		}
	}

	// read configuration from env (prefix with GG)
	viper.SetEnvPrefix("GG")
	viper.AutomaticEnv()
}

func initConfigPath() {
	configPath = viper.ConfigFileUsed()
}

func validateVersion() {
	utils.VersionChecker(viper.GetString("version"), core.Version)

}

func initLogger() {
	err := viper.Unmarshal(&configuration)
	phase.Error(err)

	logger.SetLevelStr(configuration.Settings.Log.Level)
	phase.Debug(configuration.String())
}

func init() {
	cobra.OnInitialize(initConfig, validateVersion, initLogger, initConfigPath)

	// bind log flags with viper configuration
	root.PersistentFlags().StringP("log-level", "L", "", "set log level")
	viper.BindPFlag(constants.SettingLogLevel, root.PersistentFlags().Lookup("log-level"))
	viper.SetDefault(constants.SettingLogLevel, "info")

	root.PersistentFlags().BoolP("no-config", "N", false, "will not load config from file")
	viper.BindPFlag(constants.SettingDisabledConfig, root.PersistentFlags().Lookup("no-config"))
	viper.SetDefault(constants.SettingDisabledConfig, false)
}

// Execute will run commandline interface
func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

package cmd

import (
	"errors"
	"fmt"
	"os"

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
var configOption models.ConfigurationOption = *config.DefaultOption()

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

func initLocation() {
	phase.OnInitialPhase()

	var path string
	var err error

	if viper.GetString(constants.SettingWdPath) != "" {
		configOption.SetWdPath(viper.GetString(constants.SettingWdPath))
	} else {
		path, err = os.Getwd()
		if err != nil {
			phase.Error(err)
		}
		configOption.SetWdPath(path)
	}

	path, err = os.Executable()
	if err != nil {
		phase.Warn(err)
	} else {
		configOption.AddPath(path)
	}

	path, err = os.UserHomeDir()
	if err != nil {
		phase.Warn(err)
	} else {
		configOption.AddPath(path)
	}
}

func initConfig() {
	for _, path := range configOption.Setting.ConfigDirectoryPaths() {
		viper.AddConfigPath(path)
	}

	viper.SetConfigName(configOption.Setting.FileName)
	viper.SetConfigType(configOption.Setting.FileType)

	if !viper.GetBool(constants.SettingDisabledConfig) {
		// read configuration from files
		err := viper.ReadInConfig()
		switch t := err.(type) {
		case viper.ConfigFileNotFoundError:
			phase.Warn(errors.New("config file not found"))
		default:
			phase.Warn(t)
		}
	}

	viper.SetEnvPrefix(configOption.Setting.EnvPrefix)
	viper.AutomaticEnv()
}

func initConfigPath() {
	configOption.SetConfigPath(viper.ConfigFileUsed())
}

func validateVersion() {
	err := utils.VersionChecker(viper.GetString("version"), core.Version)
	phase.Error(err)
}

func postConfig() {
	if configuration.Settings.Config.Disabled {
		phase.Debug("rollback config when user disable config")

		configuration = *config.Default()
	}
}

func initLogger() {
	err := viper.Unmarshal(&configuration)
	phase.Error(err)

	logger.SetLevelStr(configuration.Settings.Log.Level)
	phase.Debug(configuration.String())
}

func init() {
	cobra.OnInitialize(initLocation, initConfig, validateVersion, initLogger, postConfig, initConfigPath)

	// bind log flags with viper configuration
	root.PersistentFlags().StringP("log-level", "L", "", "set log level")
	viper.BindPFlag(constants.SettingLogLevel, root.PersistentFlags().Lookup("log-level"))
	viper.SetDefault(constants.SettingLogLevel, "info")

	root.PersistentFlags().BoolP("no-config", "N", false, "will not load config from file")
	viper.BindPFlag(constants.SettingDisabledConfig, root.PersistentFlags().Lookup("no-config"))
	viper.SetDefault(constants.SettingDisabledConfig, false)

	root.PersistentFlags().StringP("wd", "W", "", "custom current directory")
	viper.BindPFlag(constants.SettingWdPath, root.PersistentFlags().Lookup("wd"))
	viper.SetDefault(constants.SettingWdPath, "")
}

// Execute will run commandline interface
func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

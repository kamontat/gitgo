// Copyright © 2018 Kamontat Chantrachirathumrong <kamontat.c@hotmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/kamontat/gitgo/model"

	"github.com/kamontat/go-error-manager"
	"github.com/kamontat/go-log-manager"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var globalList *viper.Viper
var localList *viper.Viper

var repo *model.Repo
var debug bool
var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gitgo",
	Short:   "gitgo command by Kamontat Chantrachurathumrong",
	Version: "2.0.1",
	Long:    ``,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initLogger, initConfig, setOutput, initGlobalList, initLocalList, initRepository)

	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "D", false, "add debug output")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "V", false, "add verbose output")
}

func initLogger() {
	om.SetupLogger(&om.Setting{
		Color: true,
		Level: om.LLevelInfo,
		To:    &om.OutputTo{Stdout: true, File: true, FileName: "/tmp/gitgo/log.log"},
	}, nil)

	om.SetupPrinter(&om.Setting{
		Color: true,
		Level: om.LLevelInfo,
		To:    &om.OutputTo{Stdout: true, File: true, FileName: "/tmp/gitgo/print.log"},
	}, nil)
}

func setOutput() {
	if debug {
		om.Log().Setting().SetMaximumLevel(om.LLevelDebug)
		om.Log().ToDebug("set", "debug mode")
	}

	if verbose {
		om.Log().Setting().SetMaximumLevel(om.LLevelVerbose)
		om.Log().ToVerbose("set", "verbose mode")
	}
}

func initGlobalList() {
	om.Log().ToVerbose("init", "global list")

	home, err := manager.ResetError().E2P(homedir.Dir()).GetResult()
	err.ShowMessage(nil).Exit()

	globalList = viper.New()
	globalList.SetConfigFile(home.(string) + "/.gitgo/list.yaml")

	if !manager.ResetError().E1P(globalList.ReadInConfig()).HaveError() {
		om.Log().ToDebug("Global list", globalList.ConfigFileUsed())
		configVersionChecker(globalList)
	}
}

func initLocalList() {
	om.Log().ToVerbose("init", "local list")

	home, err := manager.ResetError().E2P(filepath.Abs(".")).GetResult()
	err.ShowMessage(nil).Exit()

	localList = viper.New()
	localList.SetConfigFile(home.(string) + "/.gitgo/list.yaml")

	if !manager.ResetError().E1P(localList.ReadInConfig()).HaveError() {
		om.Log().ToDebug("Local list", localList.ConfigFileUsed())
		configVersionChecker(localList)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	om.Log().ToVerbose("init", "config")
	// Find home directory.
	home, err := manager.ResetError().E2P(homedir.Dir()).GetResult()
	err.ShowMessage(nil).Exit()

	// Search config in home directory with name ".xyz" (without extension).
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./.gitgo")
	viper.AddConfigPath(home.(string) + "/.gitgo")

	viper.SetEnvPrefix("GG")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		om.Log().ToDebug("Config file", viper.ConfigFileUsed())
		if viper.Get("log") != nil && !viper.GetBool("log") {
			om.Log().Setting().SetMaximumLevel(om.LLevelNone)
		}

		configVersionChecker(nil)
	}
}

func initRepository() {
	om.Log().ToVerbose("init", "repository")
	repo = model.NewRepo()
	repo.Setup()
}

func configVersionChecker(vp *viper.Viper) bool {
	var v string
	var cv string

	v = rootCmd.Version
	if vp == nil {
		cv = viper.GetString("version")
	} else {
		cv = vp.GetString("version")
	}
	m, _ := regexp.MatchString(cv, v)
	if !m {
		manager.
			ResetError().
			AddNewErrorMessage(`config version not matches ( ` + v + ` !== ` + cv + ` )`).
			Throw().
			ShowMessage(nil).
			Exit()

		return false
	}
	return true
}

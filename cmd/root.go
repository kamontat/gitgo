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
	"errors"
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

var version = "3.1.1"

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "gitgo",
	Short: "gitgo command by Kamontat Chantrachurathumrong",
	Long: `Gitgo: git commit for organize user.
  
This command create by golang with cobra cli. 

Motivation by gitmoji, 
I used to like gitmoji but emoji isn't made for none developer.
And the problem I got is I forget which emoji is represent what.
And hard to generate changelog file. 
So I think 'short key text' is the solution of situation.

3.1.1 -> Change default and local configuration list
3.1.0 -> Add --tag to changelog generator
3.0.1 -> Add README file to local config
3.0.0 -> Change commit format and refactor code
2.4.0 -> Add --empty to allow empty changes to commit code
2.3.2 -> Issue hash tag will be always add if setting is true
2.3.1 -> Fix branch creator error, and improve logger
2.3.0 -> Add changelog command with initial changelog
2.2.1 -> Improve branch creator and commit creator
  `,
	Version: version,
}

// Execute is execute method that call by cobra cli.
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
	om.ConstrantExitOnError = false
	om.ConstrantAppName = "gitgo"

	om.SetupNewLogger(&om.Setting{
		Color: true,
		Level: om.LLevelInfo,
	})
}

func setOutput() {
	if debug {
		om.Log.Setting().SetMaximumLevel(om.LLevelDebug)
		om.Log.ToDebug("set", "debug mode")
	}

	if verbose {
		om.Log.Setting().SetMaximumLevel(om.LLevelVerbose)
		om.Log.ToVerbose("set", "verbose mode")
	}
}

func initGlobalList() {
	om.Log.ToVerbose("init", "global list")

	manager.StartResultManager().Exec02(homedir.Dir).IfError(func(t *manager.Throwable) {
		t.ShowMessage().ExitWithCode(3)
	}).IfResult(func(home string) {
		globalList = viper.New()
		globalList.SetConfigFile(home + "/.gitgo/list.yaml")

		if !manager.NewE().Add(globalList.ReadInConfig()).HaveError() {
			om.Log.ToDebug("Global list", globalList.ConfigFileUsed())

			configVersionChecker(globalList)
		}
	})
}

func initLocalList() {
	om.Log.ToVerbose("init", "local list")

	manager.StartResultManager().Exec12(filepath.Abs, ".").IfError(func(t *manager.Throwable) {
		t.ShowMessage().ExitWithCode(4)
	}).IfResult(func(home string) {
		localList = viper.New()
		localList.SetConfigFile(home + "/.gitgo/list.yaml")

		if !manager.NewE().Add(localList.ReadInConfig()).HaveError() {
			om.Log.ToDebug("Local list", localList.ConfigFileUsed())

			configVersionChecker(localList)
		}
	})
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	om.Log.ToVerbose("init", "config")

	manager.StartResultManager().Exec02(homedir.Dir).IfError(func(t *manager.Throwable) {
		t.ShowMessage().ExitWithCode(2)
	}).IfResult(func(home string) {
		// Search config in home directory with name ".xyz" (without extension).
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
		viper.AddConfigPath("./.gitgo")
		viper.AddConfigPath(home + "/.gitgo")

		viper.SetEnvPrefix("GG")
		viper.AutomaticEnv() // read in environment variables that match

		if !manager.NewE().Add(viper.ReadInConfig()).HaveError() {
			om.Log.ToDebug("Config file", viper.ConfigFileUsed())

			manager.Wrap(viper.Get("log")).UnwrapNext(func(i interface{}) interface{} {
				return viper.GetBool("log")
			}).Unwrap(func(log interface{}) {
				if !log.(bool) {
					om.Log.ToVerbose("log setting", "none of output will be print")
					om.Log.Setting().SetMaximumLevel(om.LLevelNone)
				}
			})

			configVersionChecker(nil)
		}
	})
}

func initRepository() {
	om.Log.ToVerbose("init", "repository")
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
	m, e := regexp.MatchString(cv, v)
	var ee error
	if !m {
		ee = errors.New(`config version not matches ( ` + v + ` !== ` + cv + ` )`)
	}

	manager.NewE().AddNewError(e).AddNewError(ee).Throw().
		ShowMessage().ExitWithCode(10)

	return true
}

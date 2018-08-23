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
	"os"
	"path/filepath"

	"github.com/kamontat/gitgo/model"

	manager "github.com/kamontat/go-error-manager"
	"github.com/kamontat/go-log-manager"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var initialForce bool

// configInitialCmd represents the configInitial command
var configInitialCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i", "initial"},
	Short:   "Create and initial gitgo configuration files",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		om.Log.ToLog("config", "initial start...")
		yaml := model.GeneratorYAML()

		if initialForce {
			om.Log.ToVerbose("config", "initial with force")
		}

		init := false
		var file *manager.ResultWrapper

		if inLocal {
			file = getFile(getLPath("config.yaml"))
			file.Unwrap(func(i interface{}) {
				writeTo(i.(*os.File), yaml.GDefaultConfig())
			}).Catch(func() error {
				return errors.New("Cannot save config.yaml in local")
			}, throw)

			file = getFile(getLPath("list.yaml"))
			file.Unwrap(func(i interface{}) {
				writeTo(i.(*os.File), yaml.LEmptyList())
			}).Catch(func() error {
				return errors.New("Cannot save list.yaml in local")
			}, throw)

			init = true
		}

		if inGlobal {
			file = getFile(getGPath("config.yaml"))
			file.Unwrap(func(i interface{}) {
				writeTo(i.(*os.File), yaml.GDefaultConfig())
			}).Catch(func() error {
				return errors.New("Cannot save config.yaml in global")
			}, throw)

			file = getFile(getGPath("list.yaml"))
			file.Unwrap(func(i interface{}) {
				writeTo(i.(*os.File), yaml.GDefaultList())
			}).Catch(func() error {
				return errors.New("Cannot save list.yaml in global")
			}, throw)

			init = true
		}

		if !init {
			file = getFile(getGPath("config.yaml"))
			file.Unwrap(func(i interface{}) {
				writeTo(i.(*os.File), yaml.GDefaultConfig())
			}).Catch(func() error {
				return errors.New("Cannot save config.yaml in global")
			}, throw)

			file = getFile(getGPath("list.yaml"))
			file.Unwrap(func(i interface{}) {
				writeTo(i.(*os.File), yaml.GDefaultList())
			}).Catch(func() error {
				return errors.New("Cannot save list.yaml in global")
			}, throw)
		}
	},
}

func getGPath(filename string) *manager.ResultWrapper {
	return manager.StartResultManager().Exec02(homedir.Dir).IfResultThen(func(home string) interface{} {
		path := filepath.Join(home, ".gitgo", filename)
		os.MkdirAll(filepath.Dir(path), os.ModePerm)
		return path
	})
}

func getLPath(filename string) *manager.ResultWrapper {
	return manager.StartResultManager().Exec12(filepath.Abs, ".").IfResultThen(func(home string) interface{} {
		path := filepath.Join(home, ".gitgo", filename)
		os.MkdirAll(filepath.Dir(path), os.ModePerm)
		return path
	})
}

func getFile(path *manager.ResultWrapper) *manager.ResultWrapper {
	return path.UnwrapNext(func(i interface{}) interface{} {
		om.Log.ToDebug("config", "start initial path ")

		f, _ := os.OpenFile(i.(string), os.O_CREATE|os.O_WRONLY, os.ModePerm)
		return f
	})
}

func isFileExist(file *os.File) *manager.ErrManager {
	i, e := file.Stat()
	m := manager.NewE().Add(e)
	if i.Size() <= 0 {
		m.AddMessage("File size is 0 (file empty)")
	}
	return m
}

func writeTo(file *os.File, str string) {
	e := isFileExist(file)
	if initialForce || !e.HaveError() {
		_, err := file.WriteString(str)

		e.Add(err)
		e.Throw().ShowMessage().ExitWithCode(11)

		om.Log.ToInfo("config", "Done @"+file.Name())
	} else {
		om.Log.ToWarn("config", "Exist @"+file.Name())
	}
}

func throw(throw *manager.Throwable) {
	throw.ShowMessage().ExitWithCode(len(throw.ListErrors()))
}

func init() {
	configCmd.AddCommand(configInitialCmd)

	configInitialCmd.Flags().BoolVarP(&initialForce, "force", "f", false, "force initial even file exist")
}

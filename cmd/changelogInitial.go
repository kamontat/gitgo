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
	"os"
	"path"
	"path/filepath"

	"gopkg.in/AlecAivazis/survey.v2"

	e "github.com/kamontat/gitgo/exception"
	"github.com/kamontat/gitgo/util"
	homedir "github.com/mitchellh/go-homedir"

	om "github.com/kamontat/go-log-manager"
	"github.com/spf13/cobra"
)

// changelogInitialCmd represents the changelogInitial command
var changelogInitialCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i", "initial"},
	Short:   "Create and initial changelog setting and libraries",
	Run: func(cmd *cobra.Command, args []string) {
		om.Log.ToLog("changelog", "initial start...")

		gitgoStr := path.Dir(listYAML.ConfigFileUsed())
		_, err := os.Open(gitgoStr)
		if err != nil {
			home, err := homedir.Dir()
			yaml := util.GeneratorYAML()

			if err == nil {
				path := filepath.Join(home, ".gitgo", "config.yml")
				os.MkdirAll(filepath.Dir(path), os.ModePerm)
				f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.ModePerm)
				if err == nil {
					writeTo(f, yaml.Config())
				}
			}
		}

		if initialForce {
			om.Log.ToVerbose("config", "initial with force")
		}

		yaml := util.GeneratorYAML()

		conf := openFile(gitgoStr, "config.yml")
		tmp := openFile(gitgoStr, "CHANGELOG.tpl.md")

		var style string
		survey.AskOne(&survey.Select{
			Default: "github",
			Options: []string{"github", "gitlab", "bitbucket", "none"},
		}, &style, nil)

		var url string
		survey.AskOne(&survey.Input{
			Message: "Enter repository url",
			Help:    "Add full path, include http:// or https://",
		}, &url, nil)

		writeTo(conf, yaml.ChgLogConfig(style, url))
		writeTo(tmp, yaml.ChgLogTpl())
	},
}

func getChgLogConfig(parent, file string) string {
	return path.Join(parent, "chglog", file)
}

func openFile(parent, file string) *os.File {
	filepath := getChgLogConfig(parent, file)
	d := path.Dir(filepath)
	os.MkdirAll(d, os.ModePerm)

	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, os.ModePerm)

	e.ShowAndExit(e.Error(e.IsInitial, err))
	return f
}

func init() {
	changelogCmd.AddCommand(changelogInitialCmd)

	changelogInitialCmd.Flags().BoolVarP(&initialForce, "force", "f", false, "force initial even file exist")
}

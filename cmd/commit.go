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

// Package cmd is the default package of commands provide by cobra cli.
package cmd

import (
	"github.com/kamontat/gitgo/model"
	"github.com/kamontat/gitgo/util"
	"github.com/spf13/viper"
	"gopkg.in/AlecAivazis/survey.v2"

	om "github.com/kamontat/go-log-manager"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:     "commit",
	Aliases: []string{"c"},
	Short:   "Git commit with format string",
	Run: func(cmd *cobra.Command, args []string) {
		om.Log.ToLog("commit", "start...")

		config := model.LoadCommitConfiguration(viper.GetViper())
		om.Log.ToLog("commit config", config)

		commitMessage := model.CommitMessage{}

		keyQuestion := util.GenerateQuestionViaTypeConfig("Please enter commit type", config.Key, listYAML)
		survey.AskOne(keyQuestion, &commitMessage.Type)

		scopeQuestion := util.GenerateQuestionViaTypeConfig("Please enter commit scope", config.Scope, listYAML)
		survey.AskOne(scopeQuestion, &commitMessage.Type)

		titleQuestion := util.GenerateQuestionViaTypeConfig("Please enter commit title", config.Title, listYAML)
		survey.AskOne(titleQuestion, &commitMessage.Title)

		messageQuestion := util.GenerateQuestionViaTypeConfig("Please enter commit message", config.Message, listYAML)
		survey.AskOne(messageQuestion, &commitMessage.Title)

		// hasMessage := viper.GetBool("commit.message")
		// if hasMessage {
		// 	om.Log.ToVerbose("commit", "with message")
		// } else {
		// 	om.Log.ToVerbose("commit", "without message")
		// }

		// commit := repo.GetCommit()

		// commit.SetSettings(
		// 	viper.GetInt("commit.scope.size"),
		// 	viper.GetInt("commit.message.size"),
		// )

		// commit.KeyList.Merge(listYAML)
		// commit.Commit(customKey, model.CommitOption{
		// 	Add:     false,
		// 	Empty:   empty,
		// 	Message: hasMessage,
		// 	Dry:     dry,
		// })
	},
}

var empty bool
var dry bool

var customKey string

func init() {
	rootCmd.AddCommand(commitCmd)

	commitCmd.Flags().StringVarP(&customKey, "type", "t", "", "Custom commit type [shouldn't use]")

	commitCmd.Flags().BoolVarP(&dry, "dry", "d", false, "dry run with never commit anything in git")

	commitCmd.Flags().BoolVarP(&empty, "empty", "m", false, "Commit with --allow-empty flag")
}

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
	"github.com/kamontat/gitgo/model"
	om "github.com/kamontat/go-log-manager"

	"github.com/spf13/cobra"
)

// commitInitialCmd represents the commitInitial command
var commitInitialCmd = &cobra.Command{
	Use:     "initial",
	Aliases: []string{"i", "init", "start", "s"},
	Short:   "Start commit and add every file if not add",
	Run: func(cmd *cobra.Command, args []string) {
		om.Log.ToVerbose("Commit initial", "start...")

		repo.AddAll().ShowMessage().Exit()
		om.Log.ToVerbose("Status", "Added all files or folders")

		om.Log.ToVerbose("Status", "Getting commit object")
		commit := repo.GetCommit()

		om.Log.ToVerbose("Commit", "Start create commit")
		commit.Make(model.CommitMessage{
			Type:    key,
			Scope:   scope,
			Title:   title,
			Message: message,
		}, model.CommitOptions{
			Dry:   false,
			Empty: true,
		})
	},
}

var key string
var scope string
var title string
var message string

func init() {
	commitCmd.AddCommand(commitInitialCmd)

	commitInitialCmd.Flags().StringVarP(&key, "key", "k", "init", "Custom commit key")
	commitInitialCmd.Flags().StringVarP(&key, "scope", "s", "project", "Custom commit scope")
	commitInitialCmd.Flags().StringVarP(&title, "title", "t", "Initial commit", "Custom commit title.")
	commitInitialCmd.Flags().StringVarP(&message, "message", "m", `We create this commit for initial repostiory.
Commit 'gitgo' project [https://github.com/kamontat/gitgo].`, "Custom commit message.")
}

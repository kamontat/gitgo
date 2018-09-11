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

	"github.com/kamontat/gitgo/model"

	"github.com/kamontat/go-log-manager"
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

		if commit.CanCommit() {
			om.Log.ToVerbose("Commit", "Start create commit")

			if key == "" {
				key = "init"
			}

			if title == "" {
				title = "Initial commit with files"
			}

			message := `We create this commit for initial repostiory.
      Commit 'gitgo' project [https://github.com/kamontat/gitgo].`

			if !hasMessage {
				message = ""
			}

			commit.CustomCommit(true, model.CommitMessage{
				Key:     key,
				Title:   title,
				Message: message,
			})
		} else {
			om.Log.ToError("Repository", "Cannot create commit")
			os.Exit(1)
		}
	},
}

var key string
var title string
var hasMessage bool

func init() {
	commitCmd.AddCommand(commitInitialCmd)

	commitInitialCmd.Flags().StringVarP(&key, "key", "k", "init", "Custom commit key [default=init]")
	commitInitialCmd.Flags().StringVarP(&title, "title", "t", "Initial commit with files", "Custom commit title [default=Initial commit with files]")
	commitInitialCmd.Flags().BoolVarP(&hasMessage, "no-message", "N", false, "Force commit to not use message. [default=false]")
}

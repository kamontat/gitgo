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
	"strings"

	"github.com/spf13/viper"

	e "github.com/kamontat/gitgo/exception"
	"github.com/kamontat/go-log-manager"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:     "commit",
	Aliases: []string{"c"},
	Short:   "Git commit with format string",
	Run: func(cmd *cobra.Command, args []string) {
		om.Log.ToLog("commit", "start...")

		if all {
			om.Log.ToVerbose("git add", "git add all files by git add --all")
			t := repo.AddAll()
			e.Show(t)
		} else {
			if len(each) > 0 {
				om.Log.ToDebug("git add", "add files ["+strings.Join(each, ", ")+"]")
				t := repo.Add(each)
				e.Show(t)
			}
		}

		hasMessage := viper.GetBool("commit.message")
		if hasMessage {
			om.Log.ToVerbose("commit", "with message")
		} else {
			om.Log.ToVerbose("commit", "without message")
		}

		commit := repo.GetCommit()

		commit.KeyList.Load(globalList).Merge(localList)
		commit.Commit(all, hasMessage, customKey)
	},
}

var each []string
var add bool
var all bool

var customKey string

func init() {
	rootCmd.AddCommand(commitCmd)

	commitCmd.Flags().StringVarP(&customKey, "key", "k", "", "Custom commit key [shouldn't use]")

	commitCmd.Flags().StringArrayVarP(&each, "each", "e", []string{}, "Commit with add [multiple use]")
	commitCmd.Flags().BoolVarP(&all, "all", "A", false, "Commit with add all")
	commitCmd.Flags().BoolVarP(&add, "add", "a", false, "Commit with -a flag")
}

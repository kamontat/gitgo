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
	"strings"

	"github.com/spf13/viper"

	"github.com/kamontat/go-log-manager"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:     "commit",
	Aliases: []string{"c"},
	Short:   "commit [-a <files>|--add <files>] [-A|--all]",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		om.Log().ToInfo("commit", "start...")

		if all {
			exception := repo.AddAll()
			exception.Throw().ShowMessage(nil).ExitWithCode(10)
		} else {
			if len(add) > 0 {
				om.Log().ToDebug("commit", "add files ["+strings.Join(add, ", ")+"]")
				repo.Add(add)
			}
		}

		hasMessage := viper.GetBool("commit.message")
		if hasMessage {
			om.Log().ToVerbose("commit", "with message")
		} else {
			om.Log().ToVerbose("commit", "without message")
		}
		repo.GetCommit().LoadList(globalList).MergeList(localList).Commit(hasMessage)
	},
}

var add []string
var all bool

func init() {
	rootCmd.AddCommand(commitCmd)

	commitCmd.PersistentFlags().StringArrayVarP(&add, "add", "a", []string{}, "Commit with add")
	commitCmd.PersistentFlags().BoolVarP(&all, "all", "A", false, "Commit with add all")
}

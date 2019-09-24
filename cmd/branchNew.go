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
	"github.com/kamontat/go-log-manager"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// branchNewCmd represents the branchNew command
var branchNewCmd = &cobra.Command{
	Use:     "new",
	Aliases: []string{"n", "c", "create", "s", "start", "i", "init", "initial"},
	Short:   "create new branch and checkout",
	Run: func(cmd *cobra.Command, args []string) {
		om.Log.ToVerbose("create Branch", "start...")

		isRequireIteration := viper.GetBool("branch.iteration.require")
		isRequireDesc := viper.GetBool("branch.description.require")
		isRequireIssue := viper.GetBool("branch.issue.require")
		allowIssueHashtag := viper.GetBool("branch.issue.hashtag")

		branch := repo.GetBranch()
		branch.SetDryrun(dry)

		if name == "" {
			branch.KeyList.Load(globalList).Merge(localList)
			om.Log.ToVerbose("branch", "ask for branch name")
			branch.AskCreate(
				isRequireDesc,
				isRequireIteration,
				isRequireIssue,
				allowIssueHashtag,
			)
		} else {
			om.Log.ToVerbose("branch", "create branch name "+name)
			branch.Create(name)
		}
		if !disableCheckout {
			om.Log.ToDebug("branch", "create and checkout")
			branch.CheckoutD()
		}
	},
}

var disableCheckout = false
var name string

func init() {
	branchCmd.AddCommand(branchNewCmd)

	branchNewCmd.Flags().BoolVarP(&disableCheckout, "no-checkout", "C", false, "not checkout to new branch")
	branchNewCmd.Flags().BoolVarP(&dry, "dry", "d", false, "dry run")
	branchNewCmd.Flags().StringVarP(&name, "name", "n", "", "custom branch name [shouldn't use]")
}

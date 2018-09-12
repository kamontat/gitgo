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

	"github.com/fatih/color"
	"github.com/kamontat/go-log-manager"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// branchListCmd represents the branchList command
var branchListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List all branch in local",
	Run: func(cmd *cobra.Command, args []string) {
		om.Log.ToVerbose("Branch List", "start...")

		branch := repo.GetBranch()
		branch.List(remotes, listFn)
	},
}

var listFn = func(title string, i int, r *plumbing.Reference) {
	var s string
	if title != "branch" {
		s = color.RedString(r.Name().Short())
	} else if r.Name().Short() == repo.GetBranch().CurrentBranch().Short() {
		s = color.GreenString(r.Name().Short())
	} else {
		s = r.Name().Short()
	}

	om.Log.ToInfo(fmt.Sprintf("%s: %d)", title, i+1), s)
}

var remotes bool

func init() {
	branchCmd.AddCommand(branchListCmd)

	branchListCmd.Flags().BoolVarP(&remotes, "all", "a", false, "List both local and remote branches [WIP]")
}

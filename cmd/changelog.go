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
	"os/exec"
	"path"

	"github.com/kamontat/gitgo/exception"

	"github.com/kamontat/go-log-manager"
	"github.com/spf13/cobra"
)

// changelogCmd represents the changelog command
var changelogCmd = &cobra.Command{
	Use:     "changelog",
	Aliases: []string{"change", "clog", "cl"},
	Short:   "Create changelog",
	Run: func(cmd *cobra.Command, args []string) {
		om.Log.ToLog("changelog", "create start...")

		c := exec.Command("git-chglog", "--version")
		c.Stdout = os.Stdout
		err := c.Run()
		e.ShowAndExit(e.Error(e.IsChangelog, err))

		gitgoFolder := path.Dir(localList.ConfigFileUsed())
		config := path.Join(gitgoFolder, "chglog", "config.yml")
		_, err = os.Open(config)
		if err != nil {
			e.ShowAndExit(e.ErrorMessage(e.IsChangelog, "Config file not exist in .gitgo folder"))
		}
		
		args := []string{"--config", config, "--output", changelogName}
		if nextTag != "" {
			args = append(args, "--next-tag")
			args = append(args, nextTag)
		}

		c = exec.Command("git-chglog", args...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		err = c.Run()
		e.ShowAndExit(e.Error(e.IsChangelog, err))
	},
}

var changelogName = "./CHANGELOG.md"
var nextTag = ""

func init() {
	rootCmd.AddCommand(changelogCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// changelogCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	changelogCmd.Flags().StringVarP(&changelogName, "location", "l", "./CHANGELOG.md", "Output file location")
	
	changelogCmd.Flags().StringVarP(&nextTag, "tag", "t", "", "custom tag instead of git-tag")
}

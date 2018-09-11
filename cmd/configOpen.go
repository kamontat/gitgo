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
	"path"

	"github.com/kamontat/go-log-manager"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var parent bool

// configOpenCmd represents the configOpen command
var configOpenCmd = &cobra.Command{
	Use:     "open",
	Aliases: []string{"o"},
	Short:   "open config in your default editor",
	Run: func(cmd *cobra.Command, args []string) {
		om.Log.ToLog("config", "open start...")

		if inLocal {
			if parent {
				om.Log.ToVerbose("config", "On local parent")
				open.Run(path.Dir(localList.ConfigFileUsed()))
			} else {
				om.Log.ToVerbose("config", "On local list")
				open.Run(localList.ConfigFileUsed())
			}
		} else if inGlobal {
			if parent {
				om.Log.ToVerbose("config", "on global parent")
				open.Run(path.Dir(globalList.ConfigFileUsed()))
			} else {
				om.Log.ToVerbose("config", "on global list")
				open.Run(globalList.ConfigFileUsed())
			}
		} else {
			om.Log.ToVerbose("config", "on using config")
			open.Run(viper.ConfigFileUsed())
		}
	},
}

func init() {
	configCmd.AddCommand(configOpenCmd)

	configOpenCmd.Flags().BoolVarP(&parent, "parent", "p", false, "open folder instead of config file")
}

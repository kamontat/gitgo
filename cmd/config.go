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
	"github.com/skratchdot/open-golang/open"

	"github.com/kamontat/go-log-manager"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	inGlobal bool
	inLocal  bool
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"C", "configuration"},
	Short:   "config [initial|set|get|open]",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		om.Log().ToInfo("config", "start...")

		open.Run(viper.ConfigFileUsed())
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.PersistentFlags().BoolVarP(&inLocal, "local", "l", false, "initial configuration file in local")
	configCmd.PersistentFlags().BoolVarP(&inGlobal, "global", "g", false, "initial configuration file in global")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

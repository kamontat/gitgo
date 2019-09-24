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
	"github.com/spf13/viper"

	om "github.com/kamontat/go-log-manager"
	"github.com/spf13/cobra"
)

// configOpenCmd represents the configOpen command
var configOpenCmd = &cobra.Command{
	Use:     "path",
	Aliases: []string{"p"},
	Short:   "return current config path",
	Long:    `configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		om.Log.ToLog("config", "path start...")

		om.Log.ToInfo("Configuration", viper.ConfigFileUsed())
		om.Log.ToInfo("List", listYAML.ConfigFileUsed())
	},
}

func init() {
	configCmd.AddCommand(configOpenCmd)
}

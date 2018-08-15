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
	"path/filepath"

	manager "github.com/kamontat/go-error-manager"
	"github.com/kamontat/go-log-manager"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var initialForce bool

// configInitialCmd represents the configInitial command
var configInitialCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i", "initial"},
	Short:   "Create and initial gitgo configuration files",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		om.Log().ToLog("config", "initial start...")

		yaml := `version: 2
log: true
commit:
  message: false
`

		listYaml := `version: 2
list:
  - key: feature
    value: Introducing new features.
  - key: improve
    value: Improving user experience / usability / performance.
  - key: fix
    value: Fixing a bug.
  - key: refactor
    value: Refactoring code.
  - key: file
    value: Updating file(s) or folder(s).
  - key: doc
    value: Documenting source code / user manual.
`

		emptyListYaml := `version: 2
list:
  - key: empty
    value: Update this commit header
`

		if initialForce {
			om.Log().ToVerbose("config", "initial with force")
		}

		init := false

		if inLocal {
			path := getLocalConfigPath("config.yaml")
			file := getFileFromPath(path)
			writeTo(file, yaml)

			path = getLocalConfigPath("list.yaml")
			file = getFileFromPath(path)
			writeTo(file, emptyListYaml)

			init = true
		}

		if inGlobal {
			path := getGlobalConfigPath("config.yaml")
			file := getFileFromPath(path)
			writeTo(file, yaml)

			path = getGlobalConfigPath("list.yaml")
			file = getFileFromPath(path)
			writeTo(file, listYaml)

			init = true
		}

		if !init {
			path := getGlobalConfigPath("config.yaml")
			file := getFileFromPath(path)
			writeTo(file, yaml)

			path = getGlobalConfigPath("list.yaml")
			file = getFileFromPath(path)
			writeTo(file, listYaml)
		}
	},
}

func getGlobalConfigPath(filename string) string {
	home, err := manager.GetManageError().E2P(homedir.Dir()).GetResult()
	err.ShowMessage(nil).Exit()

	path := filepath.Join(home.(string), ".gitgo", filename)
	os.MkdirAll(filepath.Dir(path), os.ModePerm)
	return path
}

func getLocalConfigPath(filename string) string {
	path := filepath.
		Join(manager.ResetError().E2P(filepath.Abs(".")).GetResultOnly().(string),
			".gitgo",
			filename,
		)
	os.MkdirAll(filepath.Dir(path), os.ModePerm)
	return path
}

func getFileFromPath(path string) *os.File {
	om.Log().ToDebug("config", "start initial path ")

	result, err := manager.ResetError().
		E2P(os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.ModePerm)).GetResult()
	err.ShowMessage(nil).Exit()

	file, ok := result.(*os.File)
	if !ok {
		manager.ResetError().AddNewErrorMessage("result file not exist!").Throw().ShowMessage(nil).Exit()
	}
	return file
}

func isFileExist(file *os.File) bool {
	i, e := file.Stat()
	manager.ResetError().AddNewError(e).Throw().ShowMessage(nil).Exit()
	return i.Size() > 0
}

func writeTo(file *os.File, str string) {
	if initialForce || !isFileExist(file) {

		_, e := file.WriteString(str)
		manager.ResetError().AddNewError(e).Throw().ShowMessage(nil).ExitWithCode(155)
		om.Log().ToInfo("config", "Done @"+file.Name())
	} else {
		om.Log().ToWarn("config", "Exist @"+file.Name())
	}
}

func init() {
	configCmd.AddCommand(configInitialCmd)

	configInitialCmd.Flags().BoolVarP(&initialForce, "force", "f", false, "force initial even file exist")
}

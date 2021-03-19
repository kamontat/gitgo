package helpers

import (
	"os"
	"path/filepath"
)

// ConfigDirectoryName is gitgo config directory name
const ConfigDirectoryName = ".gitgo"

// ListConfigDirectories will list all possible directory
func ListConfigDirectories() (result []string, errors []error) {

	wd, err := os.Getwd()
	if err != nil {
		errors = append(errors, err)
	} else {
		result = append(result, filepath.Join(wd, ConfigDirectoryName))
	}

	executable, err := os.Executable()
	if err != nil {
		errors = append(errors, err)
	} else {
		result = append(result, filepath.Join(filepath.Dir(executable), ConfigDirectoryName))
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		errors = append(errors, err)
	} else {
		result = append(result, filepath.Join(homedir, ConfigDirectoryName))
	}

	return
}

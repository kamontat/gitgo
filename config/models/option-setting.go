package models

import (
	"fmt"
	"path/filepath"
)

// OptionSetting is setting of configuration option
type OptionSetting struct {
	FileName      string
	FileType      string
	DirectoryName string
	EnvPrefix     string
	Paths         []string
}

func (o *OptionSetting) ConfigDirectoryPaths() (result []string) {
	for i := range o.Paths {
		result = append(result, o.ConfigDirectoryPath(i))
	}

	return
}

func (o *OptionSetting) ConfigDirectoryPath(index int) string {
	length := len(o.Paths)
	if index < 0 && index >= length {
		return o.DirectoryName
	} else {
		path := o.Paths[0]
		dirName := o.DirectoryName

		return filepath.Join(path, dirName)
	}
}

func (o *OptionSetting) DefaultDirectoryPath() string {
	return o.ConfigDirectoryPath(0)
}

func (o *OptionSetting) DefaultConfigFilePath() string {
	filename := fmt.Sprintf("%s.%s", o.FileName, o.FileType)
	return filepath.Join(o.DefaultDirectoryPath(), filename)
}

package models

import (
	"fmt"
	"path/filepath"
)

type ConfigFile struct {
	Disabled      bool
	FileName      string
	FileType      string
	DirectoryName string
	EnvPrefix     string
	Path          string
}

func (c *ConfigFile) SetUsedPath(path string) {
	c.Path = path
}

func (c *ConfigFile) Name() string {
	return filepath.Join(c.DirectoryName, fmt.Sprintf("%s.%s", c.FileName, c.FileType))
}

func DefaultConfigFile() *ConfigFile {
	return &ConfigFile{
		Disabled:      false,
		FileName:      "config",
		FileType:      "yml",
		DirectoryName: ".gitgo",
		EnvPrefix:     "GG",
	}
}

package models

import "path/filepath"

type Location struct {
	// Root path (all relative path will resolve base on this config)
	Paths      []string
	ConfigFile *ConfigFile
}

// Add new path to top level
func (l *Location) AddPath(path string) {
	// https://github.com/golang/go/wiki/SliceTricks#insert
	l.Paths = append(l.Paths, "")
	copy(l.Paths[1:], l.Paths)
	l.Paths[0] = path
}

func (l *Location) ConfigPaths() (fpath []string) {
	fpath = []string{}
	for _, path := range l.Paths {
		fpath = append(fpath, filepath.Join(path, l.ConfigFile.Name()))
	}

	return
}

func DefaultLocation() *Location {
	return &Location{
		Paths:      []string{},
		ConfigFile: DefaultConfigFile(),
	}
}

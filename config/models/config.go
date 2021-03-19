package models

import "fmt"

// Configuration TODO
type Configuration struct {
	Version  uint8
	Settings *Setting
}

func (c *Configuration) String() string {
	return fmt.Sprintf("version: %d\nsettings: \n%s", c.Version, c.Settings.String())
}

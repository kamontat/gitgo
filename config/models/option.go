package models

type ConfigurationOption struct {
	// WdPath is current directory path
	WdPath string
	// ConfigPath is current used configuration filepath
	ConfigPath string

	Setting *OptionSetting
}

func (c *ConfigurationOption) GetConfigPath() string {
	if c.IsConfigExist() {
		return c.ConfigPath
	} else {
		return c.Setting.DefaultConfigFilePath()
	}
}

func (c *ConfigurationOption) IsConfigExist() bool {
	return c.ConfigPath != ""
}

func (c *ConfigurationOption) SetConfigPath(path string) {
	c.ConfigPath = path
}

func (c *ConfigurationOption) SetWdPath(path string) {
	c.WdPath = path
	c.AddPath(path)
}

func (c *ConfigurationOption) AddPath(path string) {
	c.Setting.Paths = append(c.Setting.Paths, path)
}

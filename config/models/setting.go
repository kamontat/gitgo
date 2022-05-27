package models

import "fmt"

// Setting TODO
type Setting struct {
	Engine SettingEngine
	Config *ConfigFileSetting
	Log    *LogSetting
	Commit *CommitSetting
}

func (s *Setting) String() string {
	return fmt.Sprintf("  engine: '%s'\n  config: '%t'\n  log: '%s'\n  commit: %s\n", s.Engine, !s.Config.Disabled, s.Log.Level, s.Commit.String())
}

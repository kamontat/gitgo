package models

import "fmt"

// Setting TODO
type Setting struct {
	Hack   bool
	Config *ConfigFileSetting
	Log    *LogSetting
	Commit *CommitSetting
}

func (s *Setting) String() string {
	return fmt.Sprintf("  config: '%t'\n  log: '%s'\n  commit: %s\n", !s.Config.Disabled, s.Log.Level, s.Commit.String())
}

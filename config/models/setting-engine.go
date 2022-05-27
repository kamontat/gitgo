package models

type SettingEngine string

const (
	Native SettingEngine = "native" // using golang native git parser
	Cli    SettingEngine = "cli"    // using git cli
)

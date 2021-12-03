package models

type SettingV2 struct {
	LogLevel string
	Commit   *CommitSettingV2
}

func DefaultSettingV2() *SettingV2 {
	return &SettingV2{
		LogLevel: "INFO",
		Commit:   DefaultCommitSettingV2(),
	}
}

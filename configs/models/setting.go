package models

type Setting struct {
	LogLevel string
	Commit   *CommitSetting
}

func DefaultSetting() *Setting {
	return &Setting{
		LogLevel: "INFO",
		Commit:   DefaultCommitSetting(),
	}
}

package models

type Base struct {
	Version       uint8
	Settings      *Setting
	CommitMessage *CommitMessage
	Location      *Location
	Configs       *ConfigFile
}

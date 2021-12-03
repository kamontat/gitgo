package models

type ConfigurationV2 struct {
	Version       uint8
	Settings      *SettingV2
	CommitMessage *CommitMessageV2
	Location      *LocationV2
	Configs       *ConfigFileV2
}

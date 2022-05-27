package models

// CommitPrompt TODO
type CommitPrompt struct {
	Select *CommitSelectPromptSetting `yaml:"select,omitempty"`
	Input  *CommitInputPromptSetting  `yaml:"input,omitempty"`
}

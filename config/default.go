package config

import "github.com/kamontat/gitgo/config/models"

func DefaultOption() *models.ConfigurationOption {
	return &models.ConfigurationOption{
		WdPath:     "",
		ConfigPath: "",
		Setting: &models.OptionSetting{
			FileName:      "config",
			FileType:      "yml",
			DirectoryName: ".gitgo",
			EnvPrefix:     "GG",
			Paths:         []string{},
		},
	}
}

func Default() *models.Configuration {
	return &models.Configuration{
		Version: 5,
		Settings: &models.Setting{
			Hack: true, // enabled this be default untils go-git commit support allow empty and auto sign key
			Config: &models.ConfigFileSetting{
				Disabled: false,
			},
			Log: &models.LogSetting{
				Level: "info",
			},
			Commit: &models.CommitSetting{
				Key: &models.CommitTypeSetting{
					Enabled:  true,
					Required: true,
					Prompt: &models.CommitPrompt{
						Select: &models.CommitSelectPromptSetting{
							Page: 5,
							Values: []*models.CommitSelectValue{
								{
									Key:  "feat",
									Text: "Introducing new features",
								},
								{
									Key:  "perf",
									Text: "Improving user experience / usability / reliablity",
								},
								{
									Key:  "fix",
									Text: "Fixing a bug",
								},
								{
									Key:  "chore",
									Text: "Other changes unrelated to user/client",
								},
							},
						},
					},
				},
				Scope: &models.CommitTypeSetting{
					Enabled:  true,
					Required: false,
					Prompt: &models.CommitPrompt{
						Select: &models.CommitSelectPromptSetting{
							Page: 5,
							Values: []*models.CommitSelectValue{
								{
									Key:  "core",
									Text: "Core modules of application",
								},
								{
									Key:  "model",
									Text: "Application models",
								},
								{
									Key:  "api",
									Text: "Application program interface",
								},
								{
									Key:  "deps",
									Text: "Application dependencies",
								},
							},
						},
					},
				},
				Title: &models.CommitTypeSetting{
					Enabled:  true,
					Required: true,
					Prompt: &models.CommitPrompt{
						Input: &models.CommitInputPromptSetting{
							Max: 75,
						},
					},
				},
				Message: &models.CommitTypeSetting{
					Enabled:  false,
					Required: false,
					Prompt: &models.CommitPrompt{
						Input: &models.CommitInputPromptSetting{
							Max:       200,
							Multiline: true,
						},
					},
				},
			},
		},
	}
}

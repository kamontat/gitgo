package model

import (
	"github.com/spf13/viper"
)

// TypeConfig is a type of configuration inside commit and branch config
type TypeConfig struct {
	Key     string
	Type    string
	enable  bool
	require bool
	size    int
	page    int
}

// IsList for check type is list type ?
func (t *TypeConfig) IsList() bool {
	return t.Type == "list"
}

// IsInput for check type is custom type ?
func (t *TypeConfig) IsInput() bool {
	return t.Type == "input"
}

// IsMultiline for check type is custom type ?
func (t *TypeConfig) IsMultiline() bool {
	return t.Type == "multiline"
}

// IsMix for check type is mix type ?
func (t *TypeConfig) IsMix() bool {
	return t.Type == "mix"
}

// Enable will return is this settings enable or not
func (t *TypeConfig) Enable() bool {
	return t.enable
}

// Require will return is setting can be empty string
func (t *TypeConfig) Require() bool {
	return t.require
}

// Size will return size or default size if not exist
func (t *TypeConfig) Size() int {
	if t.size == 0 {
		return 100
	}

	return t.size
}

// Page will return page size or default page size if not exist
func (t *TypeConfig) Page() int {
	if t.page == 0 {
		return 5
	}

	return t.page
}

// CommitConfig is a structure of commit configuration
type CommitConfig struct {
	Version int
	Key     *TypeConfig
	Scope   *TypeConfig
	Title   *TypeConfig
	Message *TypeConfig
}

// BranchConfig is a structure of branch configuration
type BranchConfig struct {
	Version     int
	Iteration   *TypeConfig
	Key         *TypeConfig
	Title       *TypeConfig
	Description *TypeConfig
}

// LoadCommitConfiguration will return Commit config object from yaml config
func LoadCommitConfiguration(vp *viper.Viper) *CommitConfig {
	key := []string{"key", "scope", "title", "message"}

	// default enable  is true
	// default require is false
	// default type    is custom
	// default size    is 100
	// default page    is 5
	for _, v := range key {
		vp.SetDefault("commit."+v+".enable", true)
		vp.SetDefault("commit."+v+".require", false)
		vp.SetDefault("commit."+v+".type", "input")
		vp.SetDefault("commit."+v+".size", 100)
		vp.SetDefault("commit."+v+".page", 5)
	}

	return &CommitConfig{
		Version: vp.GetInt("version"),
		Key: &TypeConfig{
			Key:     "commit.keys",
			Type:    vp.GetString("commit.key.type"),
			enable:  vp.GetBool("commit.key.enable"),
			require: vp.GetBool("commit.key.require"),
			size:    vp.GetInt("commit.key.size"),
			page:    vp.GetInt("commit.key.page"),
		},
		Scope: &TypeConfig{
			Key:     "commit.scopes",
			Type:    vp.GetString("commit.scope.type"),
			enable:  vp.GetBool("commit.scope.enable"),
			require: vp.GetBool("commit.scope.require"),
			size:    vp.GetInt("commit.scope.size"),
			page:    vp.GetInt("commit.scope.page"),
		},
		Title: &TypeConfig{
			Key:     "commit.titles",
			Type:    vp.GetString("commit.title.type"),
			enable:  vp.GetBool("commit.title.enable"),
			require: vp.GetBool("commit.title.require"),
			size:    vp.GetInt("commit.title.size"),
			page:    vp.GetInt("commit.title.page"),
		},
		Message: &TypeConfig{
			Key:     "commit.messages",
			Type:    vp.GetString("commit.message.type"),
			enable:  vp.GetBool("commit.message.enable"),
			require: vp.GetBool("commit.message.require"),
			size:    vp.GetInt("commit.message.size"),
			page:    vp.GetInt("commit.message.page"),
		},
	}
}

// LoadBranchConfiguration will return Branch config object from yaml config
func LoadBranchConfiguration(vp *viper.Viper) *BranchConfig {
	return &BranchConfig{
		Version: vp.GetInt("version"),
		Iteration: &TypeConfig{
			Key:     "branch.iterations",
			Type:    vp.GetString("branch.iteration.type"),
			enable:  vp.GetBool("branch.iteration.enable"),
			require: vp.GetBool("branch.iteration.require"),
			size:    vp.GetInt("branch.iteration.size"),
			page:    vp.GetInt("branch.iteration.page"),
		},
		Key: &TypeConfig{
			Key:     "branch.keys",
			Type:    vp.GetString("branch.key.type"),
			enable:  vp.GetBool("branch.key.enable"),
			require: vp.GetBool("branch.key.require"),
			size:    vp.GetInt("branch.key.size"),
			page:    vp.GetInt("branch.key.page"),
		},
		Title: &TypeConfig{
			Key:     "branch.titles",
			Type:    vp.GetString("branch.title.type"),
			enable:  vp.GetBool("branch.title.enable"),
			require: vp.GetBool("branch.title.require"),
			size:    vp.GetInt("branch.title.size"),
			page:    vp.GetInt("branch.title.page"),
		},
		Description: &TypeConfig{
			Key:     "branch.descriptions",
			Type:    vp.GetString("branch.description.type"),
			enable:  vp.GetBool("branch.description.enable"),
			require: vp.GetBool("branch.description.require"),
			size:    vp.GetInt("branch.description.size"),
			page:    vp.GetInt("branch.description.page"),
		},
	}
}

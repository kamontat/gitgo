package git

import (
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/kamontat/gitgo/git/models"
	"github.com/kamontat/gitgo/utils/exec"
)

type Repo struct {
	git  *git.Repository
	tree *git.Worktree
}

type ChangelogOption struct {
	// Start reference to filter commits (default is latest tag)
	Start string

	// Stop reference to filter commits (default is HEAD)
	Stop string

	// Version is custom version tag of Stop reference
	Version string
}

func (r *Repo) Changelog(option *ChangelogOption) (*models.ChangeLog, error) {
	iter, err := r.git.Log(&git.LogOptions{
		Order: git.LogOrderCommitterTime,
	})
	if err != nil {
		return nil, err
	}

	commits := make([]*object.Commit, 0)
	err = iter.ForEach(func(c *object.Commit) error {
		commits = append(commits, c)
		return nil
	})

	return models.NewChangelog("v1.0.0", commits), err
}

// Commit create commit message to repository and return commit hash
func (r *Repo) Commit(msg *models.CommitMessage) (string, error) {
	message, err := msg.Formatted()
	if err != nil {
		return "", err
	}

	// TODO: Update this commit to create with sign key if exist (https://github.com/go-git/go-git/pull/214)
	hash, err := r.tree.Commit(message, &git.CommitOptions{
		All: false,
	})

	if err != nil {
		return "", err
	} else if hash.IsZero() {
		return "", errors.New("hash data is not exist")
	} else {
		return hash.String(), nil
	}
}

// CliCommit will execute raw git command instead go-git module
func (r *Repo) CliCommit(msg *models.CommitMessage, args ...string) (string, error) {
	message, err := msg.Formatted()
	if err != nil {
		return "", err
	}

	arguments := make([]string, 0)
	arguments = append(arguments, "commit", "-m", message)
	if len(args) > 0 {
		arguments = append(arguments, args...)
	}

	option := exec.NewOption()
	err = exec.Run(option, "git", arguments...)
	return "", err
}

func (r *Repo) Branches() (branches []string) {
	iter, err := r.git.Branches()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	iter.ForEach(func(ref *plumbing.Reference) error {
		branches = append(branches, ref.Name().Short())
		return nil
	})

	return
}

func (r *Repo) IsClean() bool {
	s, err := r.tree.Status()
	if err != nil {
		return false
	}

	return s.IsClean()
}

func (r *Repo) Path() string {
	return r.tree.Filesystem.Root()
}

func (r *Repo) String() string {
	branches, err := r.git.Branches()
	if err != nil {
		fmt.Println(err.Error())
	}

	branch, err := branches.Next()
	if err != nil {
		fmt.Println(err.Error())
	}

	return branch.Name().String()
}

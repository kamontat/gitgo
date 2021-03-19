package git

import (
	"errors"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

const defaultSearchLimit = 10

func new(path string, limit int) (*Repo, error) {
	if limit < 1 {
		return nil, errors.New("too many searches, cannot find git repository")
	} else if path == "" {
		return nil, git.ErrRepositoryNotExists
	}

	repo, err := git.PlainOpen(path)
	if err == git.ErrRepositoryNotExists {
		dir := filepath.Dir(path)
		return new(dir, limit-1)
	} else if err != nil {
		return nil, err
	}

	tree, err := repo.Worktree()
	if err != nil {
		return nil, err
	}

	return &Repo{
		git:  repo,
		tree: tree,
	}, nil
}

func New(path string) (*Repo, error) {
	return new(path, defaultSearchLimit)
}

package model

import (
	"path/filepath"

	"github.com/kamontat/go-error-manager"
	git "gopkg.in/src-d/go-git.v4"
)

type Repo struct {
	path       string
	repo       *git.Repository
	gitCommand *GitCommand
}

func NewRepo() *Repo {
	err := manager.StartNewManageError()

	return &Repo{
		path:       err.ExecuteWith2Parameters(filepath.Abs(".")).GetResultOnly().(string),
		gitCommand: Git(),
	}
}

func (r *Repo) Setup() {
	result, err := git.PlainOpen(r.path)
	manager.ResetError().AddNewError(err).Throw().ShowMessage(nil).ExitWithCode(5)

	r.repo = result
}

func (r *Repo) Status() git.Status {
	result, err := r.GetWorktree().Status()
	manager.ResetError().AddNewError(err).Throw().ShowMessage(nil).ExitWithCode(5)
	return result
}

func (r *Repo) Add(filepath []string) *manager.ErrManager {
	exception := manager.StartNewManageError()
	for _, f := range filepath {
		_, err := r.GetWorktree().Add(f)
		exception.AddNewError(err)
	}
	exception.Throw().ShowMessage(nil).ExitWithCode(10)
	return exception
}

func (r *Repo) AddAll() *manager.ErrManager {
	return r.gitCommand.Exec("add", "-A")
}

func (r *Repo) GetGitRepository() *git.Repository {
	return r.repo
}

func (r *Repo) GetWorktree() *git.Worktree {
	repo := r.GetGitRepository()
	result, err := repo.Worktree()
	manager.ResetError().AddNewError(err).Throw().ShowMessage(nil).ExitWithCode(5)

	return result
}

func (r *Repo) GetCommit() *Commit {
	return &Commit{
		repo: r,
	}
}

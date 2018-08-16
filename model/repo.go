package model

import (
	"path/filepath"

	"github.com/kamontat/go-error-manager"
	git "gopkg.in/src-d/go-git.v4"
)

// Repo is git repository object for gitgo.
type Repo struct {
	path       string
	repo       *git.Repository
	gitCommand *GitCommand
}

// NewRepo will return new repository with current path.
// you must call setup for load git repository to memory.
func NewRepo() *Repo {
	err := manager.StartNewManageError()

	return &Repo{
		path:       err.ExecuteWith2Parameters(filepath.Abs(".")).GetResultOnly().(string),
		gitCommand: Git(),
	}
}

// Setup will load git repository to memory.
// If any error occurred, exit with code 5.
func (r *Repo) Setup() {
	result, err := git.PlainOpen(r.path)
	manager.ResetError().AddNewError(err).Throw().ShowMessage(nil).ExitWithCode(5)

	r.repo = result
}

// Status will return *git.Status.
func (r *Repo) Status() git.Status {
	result, err := r.GetWorktree().Status()
	manager.ResetError().AddNewError(err).Throw().ShowMessage(nil).ExitWithCode(5)
	return result
}

// Add get array of filepath, and return ErrManager.
// anyway, It's will run os.Exit with code 10 if any error occurred.
func (r *Repo) Add(filepath []string) *manager.ErrManager {
	exception := manager.StartNewManageError()
	for _, f := range filepath {
		_, err := r.GetWorktree().Add(f)
		exception.AddNewError(err)
	}
	exception.Throw().ShowMessage(nil).ExitWithCode(10)
	return exception
}

// AddAll will run git add -A command in cli.
func (r *Repo) AddAll() *manager.ErrManager {
	return r.gitCommand.Exec("add", "-A")
}

// GetGitRepository will return git.Repository of this Repo
func (r *Repo) GetGitRepository() *git.Repository {
	return r.repo
}

// GetWorktree is getter method, which get git.Worktree from git.Repository.
// It's will Exit with code 5 if any error occurred.
func (r *Repo) GetWorktree() *git.Worktree {
	repo := r.GetGitRepository()
	result, err := repo.Worktree()

	manager.ResetError().AddNewError(err).Throw().ShowMessage(nil).ExitWithCode(5)

	return result
}

// GetCommit will return Commit object.
func (r *Repo) GetCommit() *Commit {
	return &Commit{
		repo: r,
	}
}

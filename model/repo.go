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
	Manager    *manager.ResultManager
}

// NewRepo will return new repository with current path.
// you must call setup for load git repository to memory.
func NewRepo() *Repo {
	return CustomRepo(".")
}

// CustomRepo will return repo of custom path
func CustomRepo(path string) *Repo {
	management := manager.StartResultManager()

	return &Repo{
		path:       management.Execute1ParametersB(filepath.Abs, path).GetResult(),
		gitCommand: Git(),
		Manager:    management,
	}
}

// Setup will load git repository to memory.
// If any error occurred, exit with code 5.
func (r *Repo) Setup() {
	result, err := git.PlainOpen(r.path)
	r.Manager.Save("", err)

	r.Manager.IfNoError(func() {
		r.repo = result
	})
}

// GetGitRepository will return git.Repository of this Repo
func (r *Repo) GetGitRepository() *manager.ResultWrapper {
	return r.Manager.IfNoErrorThen(func() interface{} {
		return r.repo
	})
}

// GetWorktree is getter method, which get git.Worktree from git.Repository.
// It's will Exit with code 5 if any error occurred.
func (r *Repo) GetWorktree() *manager.ResultWrapper {
	resultWrapper := r.GetGitRepository()
	return resultWrapper.UnwrapNext(func(i interface{}) interface{} {
		worktree, err := i.(*git.Repository).Worktree()
		r.Manager.Save("", err)
		if r.Manager.NoError() {
			return worktree
		}
		return nil
	})
}

// Status will return *git.Status.
func (r *Repo) Status() *manager.ResultWrapper {
	resultWrapper := r.GetWorktree()
	return resultWrapper.UnwrapNext(func(i interface{}) interface{} {
		status, err := i.(*git.Worktree).Status()
		r.Manager.Save("", err)
		if r.Manager.NoError() {
			return status
		}
		return nil
	})
}

// Add get array of filepath, and return ErrManager.
// anyway, It's will run os.Exit with code 10 if any error occurred.
func (r *Repo) Add(filepath []string) *manager.Throwable {
	worktree := r.GetWorktree()

	worktree.Unwrap(func(i interface{}) {
		work := i.(*git.Worktree)
		for _, f := range filepath {
			work.Add(f)
		}
	})

	return r.Manager.Throw()
}

// AddAll will run git add -A command in cli.
func (r *Repo) AddAll() *manager.ErrManager {
	return r.gitCommand.Exec("add", "-A")
}

// GetCommit will return Commit object.
func (r *Repo) GetCommit() *Commit {
	return &Commit{
		repo: r,
	}
}

package model

import (
	"path/filepath"

	e "github.com/kamontat/gitgo/exception"

	manager "github.com/kamontat/go-error-manager"
	om "github.com/kamontat/go-log-manager"

	git "gopkg.in/src-d/go-git.v4"
)

// Repo is git repository object for gitgo.
type Repo struct {
	isSetup    bool
	path       string
	repo       *git.Repository
	worktree   *git.Worktree
	gitCommand *GitCommand
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
		isSetup:    false,
		path:       management.Execute1ParametersB(filepath.Abs, path).GetResult(),
		gitCommand: Git(),
		repo:       nil,
	}
}

// Setup will load git repository to memory.
// If any error occurred, exit with code 5.
func (r *Repo) Setup() {
	if r.isSetup {
		return
	}

	om.Log.ToVerbose("Repository", "initial path "+r.path)
	result, err := git.PlainOpen(r.path)

	e.ShowAndExit(e.Error(e.IsInitial, err))

	if err == nil {
		r.repo = result

		r.worktree, err = r.repo.Worktree()
		e.ShowAndExit(e.Error(e.IsInitial, err))

		r.isSetup = result != nil && r.worktree != nil
	}
}

// GetRawGitRepository will return git.Repository of this Repo, can be nil
func (r *Repo) GetRawGitRepository() *git.Repository {
	r.Setup()

	if r.isSetup {
		return r.repo
	}
	return nil
}

// GetGitRepository will return git.Repository of this Repo
func (r *Repo) GetGitRepository() *manager.ResultWrapper {
	return manager.Wrap(r.GetRawGitRepository())
}

// GetRawWorktree is getter to get worktree, this method can return nil value
func (r *Repo) GetRawWorktree() *git.Worktree {
	r.Setup()

	if r.isSetup {
		return r.worktree
	}
	return nil
}

// GetWorktree is getter method, which get git.Worktree from git.Repository.
func (r *Repo) GetWorktree() *manager.ResultWrapper {
	return manager.Wrap(r.GetRawWorktree())
}

// Status will return *git.Status.
func (r *Repo) Status() *manager.ResultWrapper {
	resultWrapper := r.GetWorktree()
	return resultWrapper.UnwrapNext(func(i interface{}) interface{} {
		if r.Throw().CanBeThrow() {
			return nil
		}

		status, err := i.(*git.Worktree).Status()
		if err == nil {
			return status
		}
		return nil
	})
}

// Add get array of filepath, and return ErrManager.
// anyway, It's will run os.Exit with code 10 if any error occurred.
func (r *Repo) Add(filepath []string) (t *manager.Throwable) {
	t = r.Throw()
	if t.CanBeThrow() {
		return
	}

	worktree := r.GetRawWorktree()
	for _, f := range filepath {
		worktree.Add(f)
	}

	return
}

// AddAll will run git add -A command in cli.
func (r *Repo) AddAll() *manager.Throwable {
	if t := r.Throw(); t.CanBeThrow() {
		return t
	}

	t := r.gitCommand.Exec("add", "-A").Throw()
	manager := manager.NewR()
	for _, e := range t.ListErrors() {
		manager.Save("", e)
	}

	return manager.Throw()
}

// GetBranch will return config Branch struct
func (r *Repo) GetBranch() *Branch {
	ref, err := r.repo.Head()
	e.ShowAndExit(e.Error(e.IsInitial, err))

	return &Branch{
		KeyList:    (&List{Key: "branches"}),
		Repository: r.repo,
		Worktree:   r.GetRawWorktree(),
		HEAD:       ref,
	}
}

// GetCommit will return Commit object.
func (r *Repo) GetCommit() *Commit {
	return &Commit{
		throwable: r.Throw(),
	}
}

// Throw method will check has anything to throw and return Throwable object
func (r *Repo) Throw() *manager.Throwable {
	if !r.isSetup {
		return e.ErrorMessage(e.IsInitial, "This repository is not setup yet, or have error while setting.")
	}
	return manager.NewE().Throw()
}

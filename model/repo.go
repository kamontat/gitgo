package model

import (
	"path/filepath"

	"github.com/kamontat/gitgo/exception"

	"github.com/kamontat/go-error-manager"
	"github.com/kamontat/go-log-manager"

	git "gopkg.in/src-d/go-git.v4"
)

// Repo is git repository object for gitgo.
type Repo struct {
	isSetup    bool
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
		isSetup:    false,
		path:       management.Execute1ParametersB(filepath.Abs, path).GetResult(),
		gitCommand: Git(),
		Manager:    management,
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

	e.ShowAndExit(e.ThrowE(e.InitialError, err))
	if err == nil {
		r.isSetup = result != nil
		r.repo = result
	}
}

// GetGitRepository will return git.Repository of this Repo
func (r *Repo) GetGitRepository() *manager.ResultWrapper {
	r.Setup()

	if r.isSetup {
		return manager.Wrap(r.repo)
	}
	return manager.WrapNil()
}

// GetRawWorktree is getter to get worktree, this method can return nil value
func (r *Repo) GetRawWorktree() *git.Worktree {
	if r.isSetup {
		work, err := r.repo.Worktree()
		e.ShowAndExit(e.ThrowE(e.InitialError, err))
		if err == nil {
			return work
		}
		return nil
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
func (r *Repo) AddAll() *manager.Throwable {
	r.Setup()
	t := r.gitCommand.Exec("add", "-A").Throw()
	for _, e := range t.ListErrors() {
		r.Manager.Save("", e)
	}

	return r.Manager.Throw()
}

func (r *Repo) GetBranch() *Branch {
	ref, err := r.repo.Head()
	e.ShowAndExit(e.ThrowE(e.InitialError, err))

	return &Branch{
		KeyList:    (&List{}).Setup("branches"),
		Repository: r.repo,
		Worktree:   r.GetRawWorktree(),
		HEAD:       ref,
	}
}

// GetCommit will return Commit object.
func (r *Repo) GetCommit() *Commit {
	r.Setup()

	return &Commit{
		KeyList:   (&List{}).Setup("commits"),
		throwable: r.Manager.Throw(),
	}
}

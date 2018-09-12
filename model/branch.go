package model

import (
	"errors"

	"github.com/kamontat/gitgo/exception"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type Branch struct {
	Repository *git.Repository
	Worktree   *git.Worktree
	HEAD       *plumbing.Reference
	Reference  *plumbing.Reference
}

func (b *Branch) check() {
	if b.Repository == nil {
		e.ShowAndExit(e.Throw(e.BranchError, "cannot get repository"))
	}

	if b.Worktree == nil {
		e.ShowAndExit(e.Throw(e.BranchError, "cannot get git worktree"))
	}

	if b.HEAD == nil {
		e.ShowAndExit(e.Throw(e.BranchError, "cannot get commit HEAD"))
	}
}

func (b *Branch) CurrentBranch() plumbing.ReferenceName {
	return b.HEAD.Name()
}

func (b *Branch) Create(name string) *Branch {
	b.check()

	if b.Exist(name) {
		e.ShowAndExit(e.Throw(e.BranchError, "This branch name is exist."))
	}

	branchName := (plumbing.ReferenceName)("refs/heads/" + name)
	b.Reference = plumbing.NewHashReference(branchName, b.HEAD.Hash())

	err := b.Repository.Storer.SetReference(b.Reference)
	e.ShowAndExit(e.ThrowE(e.BranchError, err))

	return b
}

func (b *Branch) CheckoutD() {
	if b.Reference == nil {
		b.Checkout(b.HEAD)
	} else {
		b.Checkout(b.Reference)
	}
}

func (b *Branch) Checkout(ref *plumbing.Reference) {
	b.check()

	b.Worktree.Checkout(&git.CheckoutOptions{
		Branch: ref.Name(),
		Create: false,
	})
}

func (b *Branch) Exist(name string) bool {
	i, err := b.Repository.Branches()
	e.ShowAndExit(e.ThrowE(e.BranchError, err))

	err = i.ForEach(func(r *plumbing.Reference) error {
		if r.Name().Short() == name {
			return errors.New("error")
		}
		return nil
	})

	return err != nil
}

func (b *Branch) List(all bool, fn func(title string, i int, r *plumbing.Reference)) {
	b.check()

	i, err := b.Repository.Branches()
	e.ShowAndExit(e.ThrowE(e.BranchError, err))

	var index = 0
	i.ForEach(func(r *plumbing.Reference) error {
		fn("branch", index, r)
		index++
		return nil
	})

	if all {
		var index = 0
		rs, err := b.Repository.Remotes()
		e.ShowAndExit(e.ThrowE(e.BranchError, err))

		// list all remote
		for _, remote := range rs {
			refs, err := remote.List(&git.ListOptions{})
			e.ShowAndExit(e.ThrowE(e.BranchError, err))
			// list all ref
			for _, ref := range refs {
				if ref.Name().IsBranch() {
					fn(remote.Config().Name, index, ref)
					index++
				}
			}
		}
	}
}

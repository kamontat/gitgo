package model

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/kamontat/gitgo/exception"
	manager "github.com/kamontat/go-error-manager"
	"github.com/kamontat/go-log-manager"
	survey "gopkg.in/AlecAivazis/survey.v1"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type Branch struct {
	Repository *git.Repository
	Worktree   *git.Worktree
	HEAD       *plumbing.Reference
	Reference  *plumbing.Reference
	KeyList    *List
}

func (b *Branch) check() {
	if b.Repository == nil {
		e.ShowAndExit(e.ErrorMessage(e.IsBranch, "cannot get repository"))
	}

	if b.Worktree == nil {
		e.ShowAndExit(e.ErrorMessage(e.IsBranch, "cannot get git worktree"))
	}

	if b.HEAD == nil {
		e.ShowAndExit(e.ErrorMessage(e.IsBranch, "cannot get commit HEAD"))
	}
}

func (b *Branch) getQuestion(requireDesc, requireIter, requireIssue, issueHashtag bool) []*survey.Question {
	if !b.KeyList.IsContain() {
		e.ShowAndExit(e.ErrorMessage(e.IsInitial, "No key list for branch"))
	}

	var qs = []*survey.Question{}

	if requireIter {
		qs = append(qs, &survey.Question{
			Name: "iter",
			Prompt: &survey.Input{
				Message: "Enter iteration 'number'",
				Help:    "Iteration number is number represent each split, this will add to parent of branch name",
			},
			Validate: func(ans interface{}) error {
				str, _ := ans.(string)
				if ans == nil || str == "" {
					return nil
				}

				_, err := strconv.Atoi(str)
				return err
			},
		})
	}

	qs = append(qs, &survey.Question{
		Name: "key",
		Prompt: &survey.Select{
			Message:  "Select branch header",
			Options:  b.KeyList.MakeList(),
			Help:     "Header will represent 'one word' of the action",
			PageSize: 5,
			VimMode:  true,
		},
		Validate: survey.Required,
	})

	qs = append(qs, &survey.Question{
		Name: "title",
		Prompt: &survey.Input{
			Message: "Enter branch title",
			Help:    "Title will represent 'one word' of the result of action",
		},
		Validate: survey.Required,
	})

	if requireDesc {
		qs = append(qs, &survey.Question{
			Name: "desc",
			Prompt: &survey.Input{
				Message: "Enter branch description",
				Help:    "Title will represent 'one-two word' to descript branch",
			},
		})
	}

	if requireIssue {
		qs = append(qs, &survey.Question{
			Name: "issue",
			Prompt: &survey.Input{
				Message: "Enter issue 'number'",
				Help:    "Issue is issue number that this branch resolve or work on",
			},
			Validate: func(ans interface{}) error {
				str, _ := ans.(string)
				if ans == nil || str == "" {
					return nil
				}

				_, err := strconv.Atoi(str)
				if err == nil {
					return nil
				}

				ok, err := regexp.MatchString("#[0-9]+", str)
				if err != nil {
					return err
				}

				if !ok {
					return errors.New("your issue number is not matches our regex")
				}
				return nil
			},
			Transform: func(ans interface{}) interface{} {
				if !issueHashtag {
					return strings.Trim(ans.(string), "#")
				}
				return ans
			},
		})
	}
	return qs
}

func (b *Branch) CurrentBranch() plumbing.ReferenceName {
	return b.HEAD.Name()
}

func (b *Branch) AskCreate(requireDesc, requireIter, requireIssue, issueHashtag bool) *Branch {
	var qs = b.getQuestion(requireDesc, requireIter, requireIssue, issueHashtag)
	om.Log.ToDebug("question list", len(qs))

	name := BranchName{}
	manager.StartResultManager().Save("", survey.Ask(qs, &name)).IfNoError(func() {
		om.Log.ToVerbose("branch key", name.GetKey())
		om.Log.ToVerbose("branch title", name.Title)
		om.Log.ToVerbose("branch descript", name.Desc)
		om.Log.ToVerbose("branch issue", name.Issue)

		om.Log.ToDebug("Branch name", name.Name())

		b.Create(name.Name())
	}).IfError(func(t *manager.Throwable) {
		e.ShowAndExit(e.Update(t, e.IsUser))
	})

	return b
}

func (b *Branch) Create(name string) *Branch {
	b.check()

	if b.Exist(name) {
		e.ShowAndExit(e.ErrorMessage(e.IsBranch, "This branch name is exist."))
	}

	normalize := b.NormalizeBranchName(name)
	branchName := (plumbing.ReferenceName)("refs/heads/" + normalize)
	om.Log.ToVerbose("branch", "name "+branchName)
	b.Reference = plumbing.NewHashReference(branchName, b.HEAD.Hash())

	err := b.Repository.Storer.SetReference(b.Reference)
	e.ShowAndExit(e.Error(e.IsBranch, err))

	return b
}

func (b *Branch) CheckoutD() {
	if b.Reference == nil {
		om.Log.ToDebug("checkout", "input branch not exist, checkout to head")
		b.Checkout(b.HEAD)
	} else {
		om.Log.ToDebug("checkout", "checkout to "+b.Reference.Name().String())
		b.Checkout(b.Reference)
	}
}

func (b *Branch) Checkout(ref *plumbing.Reference) {
	b.check()

	err := b.Worktree.Checkout(&git.CheckoutOptions{
		Branch: ref.Name(),
		Create: false,
	})

	e.ShowAndExit(e.Error(e.IsCheckout, err))
}

func (b *Branch) Exist(name string) bool {
	i, err := b.Repository.Branches()
	e.ShowAndExit(e.Error(e.IsBranch, err))

	err = i.ForEach(func(r *plumbing.Reference) error {
		if r.Name().Short() == name {
			om.Log.ToWarn("branch", "branch name "+r.Name().Short()+" is exist.")
			return errors.New("error")
		}
		return nil
	})

	return err != nil
}

func (b *Branch) List(all bool, fn func(title string, i int, r *plumbing.Reference)) {
	b.check()

	i, err := b.Repository.Branches()
	e.ShowAndExit(e.Error(e.IsBranch, err))

	var index = 0
	i.ForEach(func(r *plumbing.Reference) error {
		fn("branch", index, r)
		index++
		return nil
	})

	if all {
		var index = 0
		rs, err := b.Repository.Remotes()
		e.ShowAndExit(e.Error(e.IsBranch, err))

		// list all remote
		for _, remote := range rs {
			refs, err := remote.List(&git.ListOptions{})
			e.ShowAndExit(e.Error(e.IsBranch, err))
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

// NormalizeBranchName will make name as lower case and remove spacebar
func (b *Branch) NormalizeBranchName(name string) (newname string) {
	newname = strings.ToLower(name)
	newname = strings.Replace(newname, " ", "-", -1)
	return
}

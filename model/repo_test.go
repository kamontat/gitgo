package model_test

import (
	"errors"
	"testing"

	"github.com/kamontat/go-error-manager"

	"github.com/kamontat/gitgo/model"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRepoObject(t *testing.T) {
	Convey("Given Repository", t, func() {
		Convey("When try to create new repo", func() {
			// start path is /model
			r := model.CustomRepo("..")

			Convey("And Setup the git repository", func() {
				r.Setup()

				Convey("Then can get git repository", func() {
					repo := r.GetGitRepository()

					t := repo.Unwrap(func(i interface{}) {
						So(i, ShouldNotBeNil)
					}).Catch(func() error {
						return nil
					}, nil)

					So(repo.Exist(), ShouldBeTrue)
					So(t.CanBeThrow(), ShouldBeFalse)
				})

				Convey("Then can get git worktree", func() {
					worktree := r.GetWorktree()

					So(r.Manager.HaveError(), ShouldBeFalse)

					t := worktree.Unwrap(func(i interface{}) {
						So(i, ShouldNotBeNil)
					}).Catch(func() error {
						return errors.New("throw error")
					}, func(t *manager.Throwable) {
						// This should be run
						So(t.CanBeThrow(), ShouldBeTrue)
					})

					So(worktree.Exist(), ShouldBeTrue)
					So(t.CanBeThrow(), ShouldBeFalse)
				})

				Convey("Then can add file", nil)

				Convey("Then can get every files/folders", nil)

				Convey("Then can show git status", func() {
					status := r.Status()

					So(status.Exist(), ShouldBeTrue)

					status.Catch(func() error {
						return errors.New("status not exist")
					}, func(t *manager.Throwable) {
						So(t.GetMessage(), ShouldBeEmpty)
					})
				})

				Convey("Then commit will return Commit object", nil)

			})
		})

		Convey("When create new not exist repo", func() {
			// new repo will setup repo on this folder (/model)
			r := model.NewRepo()

			Convey("Then cannot get any git repository", func() {
				repo := r.GetGitRepository()
				So(repo.NotExist(), ShouldBeTrue)
			})

			Convey("Then cannot get any git worktree", func() {
				worktree := r.GetWorktree()
				So(worktree.NotExist(), ShouldBeTrue)
			})

			Convey("Then cannot add", func() {
				t := r.Add([]string{"/abc/def"})

				So(t.CanBeThrow(), ShouldBeTrue)
			})

			Convey("Then cannot add all", func() {
				t := r.AddAll()

				So(t.CanBeThrow(), ShouldBeTrue)
			})

			Convey("Then cannot show git status", func() {
				status := r.Status()

				So(status.NotExist(), ShouldBeTrue)

				status.Catch(func() error {
					return errors.New("status not exist")
				}, func(t *manager.Throwable) {
					So(t.GetMessage(), ShouldContainSubstring, "status not exist")
				})
			})

			Convey("Then commit shouldn't exist", func() {
				commit := r.GetCommit()
				So(commit.CanCommit(), ShouldBeFalse)
			})
		})
	})
}

package model_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/bouk/monkey"
	"github.com/kamontat/gitgo/model"

	. "github.com/smartystreets/goconvey/convey"
)

type Helper struct {
	ps []*monkey.PatchGuard
}

var helper = &Helper{
	ps: []*monkey.PatchGuard{},
}

func StartHelp() *Helper {
	return helper
}

func (h *Helper) MutePrint() *Helper {
	patchPrint := monkey.Patch(fmt.Print, func(...interface{}) (i int, e error) { return })
	patchPrintln := monkey.Patch(fmt.Println, func(...interface{}) (i int, e error) { return })

	patchFPrint := monkey.Patch(fmt.Fprint, func(io.Writer, ...interface{}) (i int, e error) { return })
	patchFPrintln := monkey.Patch(fmt.Fprintln, func(io.Writer, ...interface{}) (i int, e error) { return })

	h.ps = append(h.ps, patchPrint, patchPrintln, patchFPrint, patchFPrintln)
	return h
}

func (h *Helper) FakeExit() (buf bytes.Buffer) {
	fakeExit := monkey.Patch(os.Exit, func(i int) {
		fmt.Fprint(&buf, i)
	})
	h.ps = append(h.ps, fakeExit)
	return
}

func (h *Helper) Unpatch() {
	for _, p := range h.ps {
		p.Unpatch()
	}

	h.ps = []*monkey.PatchGuard{}
}

func TestRepoObject(t *testing.T) {
	Convey("Given Current repo object", t, func() {
		Convey("Before call Setup", func() {
			Convey("Should fail to get commit status", func() {
				buf := StartHelp().MutePrint().FakeExit()

				repo := model.NewRepo()

				So(repo.Status(), ShouldBeNil)
				So(buf.String(), ShouldEqual, "155")

				StartHelp().Unpatch()
			})
		})

		Convey("After call Setup", func() {
			repo := model.NewRepo()

			repo.Setup()
			Convey("Should get commit status", func() {
				So(repo.Status().String(), ShouldNotBeNil)
			})
		})
	})

	Convey("Given Repo object", t, func() {

		Convey("Before call setup", func() {

			Convey("Should throw error, if call Repo methods", nil)

			Convey("Shouldn't get any memory repository", nil)

			Convey("Shouldn't get git repository worktree", nil)

		})

		Convey("After call setup", func() {

			Convey("Git repo should create in memory", nil)

			Convey("Should return git.Repository", nil)

			Convey("Should able to get git Worktree", nil)

			Convey("Should show git status", nil)

			Convey("Should add new file to git", nil)

			Convey("Should add all files and folder to git", nil)

			Convey("Should get commit object", nil)

		})

	})
}

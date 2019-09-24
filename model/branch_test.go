package model_test

import (
	"testing"

	"github.com/kamontat/gitgo/model"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBranch(t *testing.T) {
	Convey("Given Branch", t, func() {
		b := &model.Branch{}

		Convey("When normalize branch name", func() {
			newName := b.NormalizeBranchName("soMeBRanChNAmE")
			Convey("Then all charector should be lowercase", func() {

				So(newName, ShouldEqual, "somebranchname")
			})
		})

		Convey("When normalize name with space", func() {
			newName := b.NormalizeBranchName("i he she it")

			Convey("Then replace all spacebar to dash", func() {
				So(newName, ShouldEqual, "i-he-she-it")
			})
		})
	})
}

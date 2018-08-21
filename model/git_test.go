package model_test

import (
	"bytes"
	"testing"

	"github.com/kamontat/go-log-manager"

	"github.com/kamontat/gitgo/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGitCommand(t *testing.T) {
	om.SetupLogger(&om.Setting{Color: false, Level: om.LLevelVerbose, To: &om.OutputTo{Stdout: true, File: false}}, nil)
	om.Log().Setting().SetMaximumLevel(om.LLevelNone)

	Convey("Given GitCommand", t, func() {
		g := model.Git()

		Convey("When get the command", func() {
			Convey("Then shouldn't be nil", func() {
				So(g, ShouldNotBeNil)
			})

			Convey("the result should be singleton", func() {
				So(model.Git(), ShouldEqual, g)
			})
		})

		Convey("When setup output", func() {
			var buf bytes.Buffer
			g.SetOutWriter(&buf).Exec("status")

			Convey("Then the output should write to custom Writer", func() {
				So(buf.String(), ShouldNotBeEmpty)
				So(buf.String(), ShouldNotBeNil)
				So(buf.String(), ShouldContainSubstring, "commit")
			})
		})

		Convey("When reset both out and err", func() {
			model.Git().SetErrWriter(nil)
			model.Git().SetOutWriter(nil)
		})

		Convey("When setup error output", func() {
			var buf bytes.Buffer
			g.SetErrWriter(&buf).Exec("abc")

			Convey("Then the output should write to custom Writer", func() {
				So(buf.String(), ShouldNotBeEmpty)
				So(buf.String(), ShouldNotBeNil)
				So(buf.String(), ShouldContainSubstring, "abc")
			})
		})

		Convey("When setup input", func() {
			var buf bytes.Buffer
			var err bytes.Buffer
			buf.WriteString("Hello world")
			g.SetReader(&buf).SetErrWriter(&err).Exec("xyz")

			Convey("Then the output should write to custom Writer", func() {
				So(err.String(), ShouldNotBeEmpty)
				So(err.String(), ShouldNotBeNil)
				So(err.String(), ShouldContainSubstring, "xyz")
			})
		})
	})
}

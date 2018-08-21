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
		Convey("When get the command", func() {
      g := model.Git()

      Convey("Then shouldn't be nil" func() {
        So(g, ShouldNotBeNil)
      })

      Convey("When setup output", func() {
			  var buf bytes.Buffer
        g.SetOutWriter(&buf).Exec("status")
      
        Convey("Then the output should write to custom Writer",func() {
          So(buf.String(), ShouldNotBeEmpty)
          So(buf.String(), ShouldNotBeNil)
          So(buf.String(), ShouldContainSubstring, "commit")
        })
      })

      Convey("When setup error output", func() {
			  var buf bytes.Buffer
        g.SetErrWriter(&buf).Exec("abc")
      
        Convey("Then the output should write to custom Writer",func() {
          So(buf.String(), ShouldNotBeEmpty)
          So(buf.String(), ShouldNotBeNil)
          So(buf.String(), ShouldContainSubstring, "abc")
        })
      })
		})
	})
}

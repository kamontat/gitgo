package util_test

import (
	"testing"

	"github.com/kamontat/gitgo/util"
	. "github.com/smartystreets/goconvey/convey"
)

func TestConfig(t *testing.T) {
	Convey("Given config yaml object", t, func() {
		yaml := util.GeneratorYAML()

		Convey("When check", func() {
			Convey("Then shouldn't be empty or nil", func() {
				So(yaml, ShouldNotBeNil)
			})
		})

		Convey("When check global default config", func() {
			g := yaml.Config()

			Convey("Then it should contain version", func() {
				So(g, ShouldContainSubstring, "version: ")
			})

			Convey("Then it should contain log setting", func() {
				So(g, ShouldContainSubstring, "log: info")
			})

			Convey("Then it should contain settings setting", func() {
				So(g, ShouldContainSubstring, "settings:")
			})

			Convey("Then it should contain commit.message setting", func() {
				So(g, ShouldContainSubstring, "message:")
			})
		})

		Convey("When check local empty list", func() {
			g := yaml.ListConfig()

			Convey("Then it should contain version", func() {
				So(g, ShouldContainSubstring, "version: ")
			})

			Convey("Then it should contain commits array", func() {
				So(g, ShouldContainSubstring, "commit:")
				So(g, ShouldContainSubstring, "- type: ")
				So(g, ShouldContainSubstring, "value: ")
			})

			Convey("Then it should contain branch array", func() {
				So(g, ShouldContainSubstring, "branch:")
				So(g, ShouldContainSubstring, "- type: ")
				So(g, ShouldContainSubstring, "value: ")
			})
		})
	})
}

package model_test

import (
	"testing"

	"github.com/kamontat/gitgo/model"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConfig(t *testing.T) {
	Convey("Given config yaml object", t, func() {
		yaml := model.GeneratorYAML()

		Convey("When check", func() {
			Convey("Then shouldn't be empty or nil", func() {
				So(yaml, ShouldNotBeNil)
			})
		})

		Convey("When check global default config", func() {
			g := yaml.GDefaultConfig()

			Convey("Then it should contain version", func() {
				So(g, ShouldContainSubstring, "version: ")
			})

			Convey("Then it should contain log setting", func() {
				So(g, ShouldContainSubstring, "log: true")
			})

			Convey("Then it should contain commit.message setting", func() {
				So(g, ShouldContainSubstring, "message: false")
			})
		})

		Convey("When check global default list", func() {
			g := yaml.GDefaultList()

			Convey("Then it should contain version", func() {
				So(g, ShouldContainSubstring, "version: ")
			})

			Convey("Then it should contain commit array", func() {
				So(g, ShouldContainSubstring, "commits:")
				So(g, ShouldContainSubstring, "- key: ")
				So(g, ShouldContainSubstring, "value: ")
			})

			Convey("Then it should contain some default commit list", func() {
				So(g, ShouldContainSubstring, "feature")
				So(g, ShouldContainSubstring, "improve")
				So(g, ShouldContainSubstring, "fix")
			})
		})

		Convey("When check local empty list", func() {
			g := yaml.LEmptyList()

			Convey("Then it should contain version", func() {
				So(g, ShouldContainSubstring, "version: ")
			})

			Convey("Then it should contain commits array", func() {
				So(g, ShouldContainSubstring, "commits:")
				So(g, ShouldContainSubstring, "- key: ")
				So(g, ShouldContainSubstring, "value: ")
			})

			Convey("Then it should contain branch array", func() {
				So(g, ShouldContainSubstring, "branches:")
				So(g, ShouldContainSubstring, "- key: ")
				So(g, ShouldContainSubstring, "value: ")
			})
		})
	})
}

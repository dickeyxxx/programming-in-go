package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCliApp(t *testing.T) {
	Convey("with an app", t, func() {
		app := &CliApp{}
		Convey("Args()", func() {
			Convey("there are some", func() {
				So(app.Args(), ShouldNotBeEmpty)
			})
		})
	})
}

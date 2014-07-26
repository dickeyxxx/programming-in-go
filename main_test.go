package main

import (
	"bytes"
	"io"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type fakeApp struct {
	args     []string
	exitCode int
}

func (a *fakeApp) Exit(code int) {
	a.exitCode = code
}

func (a *fakeApp) Args() []string {
	return a.args
}

func TestMain(t *testing.T) {
	Convey("with stdout and stderr", t, func() {
		stderr = bytes.NewBuffer([]byte{})
		stdout = bytes.NewBuffer([]byte{})

		Convey("running with no command", func() {
			app = &fakeApp{args: []string{"foo"}}
			main()

			Convey("it prints the usage", func() {
				So(output(stderr), ShouldStartWith, "USAGE: foo [command]")
			})

			Convey("has an exit code 1", func() {
				So(exitCode(), ShouldEqual, 1)
			})
		})

		Convey("running with a nonexistant command", func() {
			app = &fakeApp{args: []string{"foo", "???"}}
			main()

			Convey("it prints the usage", func() {
				So(output(stderr), ShouldStartWith, "USAGE: foo [command]")
			})

			Convey("has an exit code 1", func() {
				So(exitCode(), ShouldEqual, 1)
			})
		})

		Convey(`run command`, func() {
			app = &fakeApp{args: []string{"foo", "run"}}
			main()

			Convey("it prints the usage", func() {
				So(output(stdout), ShouldEqual, "running!\n")
			})

			Convey("has an exit code 0", func() {
				So(exitCode(), ShouldEqual, 0)
			})
		})
	})
}

func output(buffer io.Writer) string {
	return buffer.(*bytes.Buffer).String()
}

func exitCode() int {
	return app.(*fakeApp).exitCode
}

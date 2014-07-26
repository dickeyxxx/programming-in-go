package main

import (
	"bytes"
	"io"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPrintUsage(t *testing.T) {
	Convey("with stderr", t, func() {
		stderr = bytes.NewBuffer([]byte{})

		Convey("it prints the usage", func() {
			printUsage()
			So(output(stderr), ShouldStartWith, "USAGE:")
		})
	})
}

func output(buffer io.Writer) string {
	return buffer.(*bytes.Buffer).String()
}

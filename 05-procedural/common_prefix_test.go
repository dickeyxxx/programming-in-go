package procedural

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCommonPrefix(t *testing.T) {
	Convey("with test data", t, func() {
		tests := []struct {
			prefix string
			paths  []string
		}{
			{"", []string{}},
			{"/home/user/goeg", []string{"/home/user/goeg", "/home/user/goeg/prefix", "/home/user/goeg/prefix/extra"}},
			{"/home/user/", []string{"/home/user/goeg", "/home/user/goeg/prefix", "/home/user/prefix/extra"}},
			{"/pecan/π/", []string{"/pecan/π/goeg", "/pecan/π/goeg/prefix", "/pecan/π/prefix/extra"}},
			{"/pecan/π/circle", []string{"/pecan/π/circle", "/pecan/π/circle/prefix", "/pecan/π/circle/prefix/extra"}},
			{"/home/user", []string{"/home/user/goeg", "/home/users/goeg", "/home/userspace/goeg"}},
			{"/", []string{"/home/user/goeg", "/tmp/user", "/var/log"}},
			{"/home/", []string{"/home/mark/goeg", "/home/user/goeg"}},
			{"", []string{"home/user/goeg", "/tmp/user", "/var/log"}},
		}

		for _, test := range tests {
			So(CommonPrefix(test.paths), ShouldEqual, test.prefix)
		}
	})
}

func TestPathPrefix(t *testing.T) {
	Convey("with test data", t, func() {
		tests := []struct {
			prefix string
			paths  []string
		}{
			{"", []string{}},
			{"/home/user/goeg", []string{"/home/user/goeg", "/home/user/goeg/prefix", "/home/user/goeg/prefix/extra"}},
			{"/home/user", []string{"/home/user/goeg", "/home/user/goeg/prefix", "/home/user/prefix/extra"}},
			{"/pecan/π", []string{"/pecan/π/goeg", "/pecan/π/goeg/prefix", "/pecan/π/prefix/extra"}},
			{"/pecan/π/circle", []string{"/pecan/π/circle", "/pecan/π/circle/prefix", "/pecan/π/circle/prefix/extra"}},
			{"/home", []string{"/home/user/goeg", "/home/users/goeg", "/home/userspace/goeg"}},
			{"", []string{"/home/user/goeg", "/tmp/user", "/var/log"}},
			{"/home", []string{"/home/mark/goeg", "/home/user/goeg"}},
			{"", []string{"home/user/goeg", "/tmp/user", "/var/log"}},
		}

		for _, test := range tests {
			So(CommonPathPrefix(test.paths), ShouldEqual, test.prefix)
		}
	})
}

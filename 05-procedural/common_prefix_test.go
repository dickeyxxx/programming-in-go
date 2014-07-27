package procedural

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCommonPrefix(t *testing.T) {
	Convey("with no strings", t, func() {
		strings := []string{}
		result := CommonPrefix(strings)
		So(result, ShouldEqual, "")
	})

	Convey("with unlike strings", t, func() {
		strings := []string{"foo", "bar"}
		result := CommonPrefix(strings)
		So(result, ShouldEqual, "")
	})

	Convey("with one string", t, func() {
		strings := []string{"foobar"}
		result := CommonPrefix(strings)
		So(result, ShouldEqual, "foobar")
	})

	Convey("with same strings", t, func() {
		strings := []string{"foobar", "foobar"}
		result := CommonPrefix(strings)
		So(result, ShouldEqual, "foobar")
	})

	Convey("with common prefixes", t, func() {
		strings := []string{"foobar", "foobaz", "foob", "fooer"}
		result := CommonPrefix(strings)
		So(result, ShouldEqual, "foo")
	})

	Convey("with unicode", t, func() {
		strings := []string{"foΩobar", "foΩobaz", "foΩob", "foΩoer"}
		result := CommonPrefix(strings)
		So(result, ShouldEqual, "foΩo")
	})
}

func TestPathPrefix(t *testing.T) {
	Convey("with no strings", t, func() {
		strings := []string{}
		result := CommonPathPrefix(strings)
		So(result, ShouldEqual, "")
	})

	Convey("with 3 paths", t, func() {
		strings := []string{"/foo/bar/boy", "/foo/bar/baz", "/foo/bar/zaz"}
		result := CommonPathPrefix(strings)
		So(result, ShouldEqual, "/foo/bar")
	})

	Convey("with test data", t, func() {
		paths := [][]string{
			{"/home/user/goeg", "/home/user/goeg/prefix", "/home/user/goeg/prefix/extra"},
			{"/home/user/goeg", "/home/user/goeg/prefix", "/home/user/prefix/extra"},
			{"/pecan/π/goeg", "/pecan/π/goeg/prefix", "/pecan/π/prefix/extra"},
			{"/pecan/π/circle", "/pecan/π/circle/prefix", "/pecan/π/circle/prefix/extra"},
			{"/home/user/goeg", "/home/users/goeg", "/home/userspace/goeg"},
			{"/home/user/goeg", "/tmp/user", "/var/log"},
			{"/home/mark/goeg", "/home/user/goeg"},
			{"home/user/goeg", "/tmp/user", "/var/log"},
		}

		expectations := []string{
			"/home/user/goeg",
			"/home/user",
			"/pecan/π",
			"/pecan/π/circle",
			"/home",
			"",
			"/home",
			"",
		}

		for i := range paths {
			So(CommonPathPrefix(paths[i]), ShouldEqual, expectations[i])
		}
	})
}

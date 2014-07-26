package strings

import (
	"io/ioutil"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseIni(t *testing.T) {
	Convey("with a newline-split mozilla.ini file", t, func() {
		ini, _ := ioutil.ReadFile("mozilla.ini")

		Convey("it parses it into a map", func() {
			results := ParseIni(strings.Split(string(ini), "\n"))
			expectedResults := map[string]map[string]string{
				"General": {"User": "dickeyxxx"},
				"Gecko":   {"MinVersion": "1.9.1", "MaxVersion": "1.9.1.*"},
				"XRE":     {"EnableProfileMigrator": "0", "EnableExtensionManager": "1"},
				"App":     {"Vendor": "Mozilla", "Profile": "mozilla/firefox", "Name": "Iceweasel", "Version": "3.5.16"},
			}

			So(results, ShouldResemble, expectedResults)
		})
	})

	Convey("with a newline-split gitconfig.ini file", t, func() {
		ini, _ := ioutil.ReadFile("gitconfig.ini")

		Convey("it parses it into a map", func() {
			results := ParseIni(strings.Split(string(ini), "\n"))
			expectedResults := map[string]map[string]string{
				"core":       {"excludesfile": "/Users/dickeyxxx/.gitignore", "editor": "/usr/local/bin/vim", "autocrlf": "input"},
				"color":      {"ui": "auto"},
				"push":       {"default": "current"},
				"help":       {"autocorrect": "1"},
				"alias":      {"a": "add", "submodules": "submodule"},
				"user":       {"email": "jeff@dickeyxxx.com", "name": "dickeyxxx"},
				"credential": {"helper": "osxkeychain"},
				"rerere":     {"enabled": "1"},
				"merge":      {"ff": "false"},
			}

			So(results, ShouldResemble, expectedResults)
		})
	})
}

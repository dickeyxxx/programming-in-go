package strings

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const mozillaIni = `[App]
Name=Iceweasel
Profile=mozilla/firefox
Vendor=Mozilla
Version=3.5.16
[Gecko]
MaxVersion=1.9.1.*
MinVersion=1.9.1
[General]
User=dickeyxxx
[XRE]
EnableExtensionManager=1
EnableProfileMigrator=0
`

func TestPrintIni(t *testing.T) {
	Convey("it works", t, func() {
		results := PrintIni(map[string]map[string]string{
			"General": {"User": "dickeyxxx"},
			"Gecko":   {"MinVersion": "1.9.1", "MaxVersion": "1.9.1.*"},
			"XRE":     {"EnableProfileMigrator": "0", "EnableExtensionManager": "1"},
			"App":     {"Vendor": "Mozilla", "Profile": "mozilla/firefox", "Name": "Iceweasel", "Version": "3.5.16"},
		})

		So(results, ShouldEqual, mozillaIni)
	})
}
